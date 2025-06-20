package middlewares

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/AleksandrVishniakov/tgbots-util/ctxutil"
	"github.com/google/uuid"
)

type responseWriterWrapper struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	status int
	data   []byte
}

func (w *responseWriterWrapper) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *responseWriterWrapper) Write(data []byte) (int, error) {
	w.data = data
	return w.ResponseWriter.Write(data)
}

func Logger(
	log *slog.Logger,
) func(next http.Handler) http.Handler {
	const src = "middlewares.Logger"
	log = log.With(slog.String("src", src))

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rid := requestID(r)
			w.Header().Add("X-Request-ID", rid)
			ctx := r.Context()
			ctx = context.WithValue(ctx, ctxutil.ContextKey_RequestID, rid)

			start := time.Now()

			wrapped := &responseWriterWrapper{
				ResponseWriter: w,
				Flusher:        w.(http.Flusher),
				Hijacker:       w.(http.Hijacker),
				status:         http.StatusOK,
			}

			next.ServeHTTP(wrapped, r.WithContext(ctx))

			attrs := []any{
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", wrapped.status),
				slog.String("duration", time.Since(start).String()),
			}

			if host, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
				attrs = append(attrs, slog.String("ip", host))
			}

			if wrapped.status >= http.StatusBadRequest {
				log.ErrorContext(ctx, "request failed", attrs...)
			} else {
				log.InfoContext(ctx, "request completed", attrs...)
			}
		})
	}
}

func requestID(r *http.Request) string {
	if id := r.Header.Get("X-Request-ID"); id != "" {
		return id
	}

	return uuid.NewString()
}
