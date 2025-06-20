package logger

import (
	"context"
	"log"
	"log/slog"

	"github.com/AleksandrVishniakov/tgbots-util/ctxutil"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
    if requestID, ok := ctx.Value(ctxutil.ContextKey_RequestID).(string); ok {
		log.Println("Got Request ID", requestID)
        r.AddAttrs(slog.String(ctxutil.ContextKey_RequestID.String(), requestID))
    } else {
		log.Println("no request id")
	}
    return h.Handler.Handle(ctx, r)
}
