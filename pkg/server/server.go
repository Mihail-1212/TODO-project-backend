package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type ServerConfig struct {
	address string
	handler http.Handler
}

func NewServerConfig(address string, handler http.Handler) *ServerConfig {
	return &ServerConfig{
		address: address,
		handler: handler,
	}
}

func (s Server) Run(cfg *ServerConfig) error {
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.address,
		Handler:        cfg.handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.httpServer.ListenAndServe()
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
