package server

import (
	"time"

	"dummy_ai/internal/config"
)

type (
	Server struct {
		address         string
		shutdownTimeout time.Duration
	}
)

func New(config *config.Config) (*Server, error) {
	server := &Server{
		address:         config.ServerAddress,
		shutdownTimeout: config.ServerShutdownTimeout,
	}
	return server, nil
}

func (server *Server) Run() error {
	return nil
}
