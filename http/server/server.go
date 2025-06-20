package server

import (
	"context"
	"fmt"
	"net/http"
)

type HTTPServer struct {
	server *http.Server
}

type Configs struct {
	Port int
	Host string
}

func New(cfg Configs, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler: handler,
		},
	}
}

func (s *HTTPServer) Run() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
