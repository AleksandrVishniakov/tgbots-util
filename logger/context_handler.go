package logger

import (
	"context"
	"log/slog"

	"github.com/AleksandrVishniakov/tgbots-util/ctxutil"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
    if requestID, ok := ctx.Value(ctxutil.ContextKey_RequestID).(string); ok {
		panic("Got Request ID " + requestID)
        r.AddAttrs(slog.String(ctxutil.ContextKey_RequestID.String(), requestID))
    } else {
		panic("no request id")
	}
    return h.Handler.Handle(ctx, r)
}
