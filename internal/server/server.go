package server

import (
	"net/http"
)

type (
	Server struct {
		address string
		router  *http.ServeMux
	}
)

func New(address string) (*Server, error) {
	server := &Server{
		address: address,
		router:  http.NewServeMux(),
	}
	return server, nil
}

func (server *Server) Start() error {
	return http.ListenAndServe(server.address, server.router)
}
