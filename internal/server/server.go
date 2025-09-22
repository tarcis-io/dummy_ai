package server

import (
	"fmt"
	"net/http"
)

type (
	Server struct {
		router *http.ServeMux
	}
)

func New() (*Server, error) {
	server := &Server{
		router: http.NewServeMux(),
	}
	return server, nil
}

func (server *Server) Start(address string) error {
	err := http.ListenAndServe(address, server.router)
	if err != nil {
		return fmt.Errorf("failed to start server error=%w", err)
	}
	return nil
}
