package server

import (
	"net/http"
)

type (
	Server struct {
		router *http.ServeMux
	}
)
