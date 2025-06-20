package e

import (
	"fmt"
	"net/http"
	"time"
)

type HTTPError struct {
	Code      int `json:"code"`
	Message   string `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	err error
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("API Error: %d. %s", e.Code, e.Message)
}

func (e *HTTPError) Unwrap() error {
	return e.err
}

func NewError(code int, message string) *HTTPError {
	return &HTTPError{
		Code:      code,
		Message:   message,
		Timestamp: time.Now(),
		err: nil,
	}
}

func Internal(opts ...HTTPErrorOption) error {
	base := NewError(http.StatusInternalServerError, "internal error")
	applyOptions(base, opts...)
	return base
}

func BadRequest(opts ...HTTPErrorOption) error {
	base := NewError(http.StatusBadRequest, "bad request")
	applyOptions(base, opts...)
	return base
}
