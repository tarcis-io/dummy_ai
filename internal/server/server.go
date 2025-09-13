package server

import (
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
