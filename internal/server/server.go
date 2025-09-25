package server

import (
	"dummy_ai/internal/config"
)

type (
	Server struct {
	}
)

func New(config *config.Config) (*Server, error) {
	server := &Server{}
	return server, nil
}

func (server *Server) Run() error {
	return nil
}
