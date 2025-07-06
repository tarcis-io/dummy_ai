package server

import (
	"net/http"

	"dummy_ai/internal/env"
)

var (
	staticFiles = map[string]string{
		"/favicon.ico": "./static/image/favicon/favicon.ico",
		"/favicon.svg": "./static/image/favicon/favicon.svg",
	}
)

func ListenAndServe() {
	for route, file := range staticFiles {
		serveFile(route, file)
	}
	listenAndServe()
}

func serveFile(route string, file string) {
	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {
		http.ServeFile(responseWriter, request, file)
	})
}

func listenAndServe() {
	if err := http.ListenAndServe(env.ServerAddress(), nil); err != nil {
		panic(err)
	}
}
