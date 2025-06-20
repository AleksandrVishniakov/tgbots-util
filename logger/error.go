package logger

import "log/slog"

func Err(err error) slog.Attr {
	if err != nil {
		return slog.String("error", err.Error())
	}

	return slog.Attr{}
}
