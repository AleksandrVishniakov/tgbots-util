package ctxutil

type ContextKey string

const (
	ContextKey_RequestID ContextKey = "rid"
)

func (k ContextKey) String() string {
	return string(k)
}
