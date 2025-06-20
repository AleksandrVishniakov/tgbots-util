package e

type HTTPErrorOption func(e *HTTPError)

func WithCode(code int) HTTPErrorOption {
	return func(e *HTTPError) {
		e.Code = code
	}
}

func WithMessage(message string) HTTPErrorOption {
	return func(e *HTTPError) {
		e.Message = message
	}
}

func WithError(err error) HTTPErrorOption {
	return func(e *HTTPError) {
		e.err = err
	}
}

func applyOptions(e *HTTPError, opts ...HTTPErrorOption) {
	for _, fn := range opts {
		fn(e)
	}
}
