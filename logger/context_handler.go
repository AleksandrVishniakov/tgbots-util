package logger

import (
	"context"
	"log/slog"

	"github.com/AleksandrVishniakov/tgbots-util/ctxutil"
)

type ContextHandler struct {
	slog.Handler
	attrs []slog.Attr
}

func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{
		Handler: h.Handler.WithAttrs(attrs),
		attrs:   append(h.attrs, attrs...),
	}
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if requestID, ok := ctx.Value(ctxutil.ContextKey_RequestID).(string); ok {
		r.AddAttrs(slog.String(ctxutil.ContextKey_RequestID.String(), requestID))
	}
	return h.Handler.Handle(ctx, r)
}
