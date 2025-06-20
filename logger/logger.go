package logger

import (
	"io"
	"log/slog"
)

func New(w io.Writer, dev bool) *slog.Logger {
	var handler slog.Handler
	if dev {
		handler = slog.NewTextHandler(w, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	} else {
		handler = slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	handler = &ContextHandler{Handler: handler}

	return slog.New(handler)
}
