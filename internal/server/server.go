package server

import (
	"net/http"
	"time"

	"dummy_ai/internal/config"
)

type (
	Server struct {
		address         string
		router          *http.ServeMux
		shutdownTimeout time.Duration
	}
)

func New(config *config.Config) (*Server, error) {
	router := http.NewServeMux()
	server := &Server{
		address:         config.ServerAddress,
		shutdownTimeout: config.ServerShutdownTimeout,
		router:          router,
	}
	return server, nil
}

func (server *Server) Run() error {
	return nil
}
