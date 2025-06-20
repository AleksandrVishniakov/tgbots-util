package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/AleksandrVishniakov/tgbots-util/http/e"
	"github.com/AleksandrVishniakov/tgbots-util/http/json"
	"github.com/AleksandrVishniakov/tgbots-util/logger"
)

func Recovery(log *slog.Logger) func(http.Handler) http.Handler {
	const src = "middlewares.Recovery"
	log = log.With(slog.String("src", src))

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			defer func() {
				if err := recover(); err != nil {
					log.ErrorContext(ctx, "got panic", slog.Any("error", err))

					httpError := e.NewError(http.StatusInternalServerError, "panic")
					if err := json.EncodeResponse(w, httpError, httpError.Code); err != nil {
						log.ErrorContext(ctx, "encode error", logger.Err(err))
					}
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
