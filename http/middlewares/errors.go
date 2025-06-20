package middlewares

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/AleksandrVishniakov/tgbots-util/http/e"
	"github.com/AleksandrVishniakov/tgbots-util/http/json"
	"github.com/AleksandrVishniakov/tgbots-util/logger"
)

type ErrorHandler func(w http.ResponseWriter, r *http.Request) error

func Error(next ErrorHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := next(w, r)
		if err == nil {
			return
		}

		var httpError *e.HTTPError
		if !errors.As(err, &httpError) {
			httpError = e.NewError(http.StatusInternalServerError, err.Error())
		}
		
		if err := json.EncodeResponse(w, httpError, httpError.Code); err != nil {
			slog.ErrorContext(ctx, "encode error", logger.Err(err))
		}
	})
}
