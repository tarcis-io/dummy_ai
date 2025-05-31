package server

import (
	"net/http"
	"text/template"
)

import (
	"dummy_ai/internal/env"
)

var (
	serverTemplate = template.Must(template.ParseFiles("./static/server/server.html"))
)

func Start() {

	listenAndServe()
}

func listenAndServe() {

	if err := http.ListenAndServe(env.ServerAddress(), nil); err != nil {

		panic(err)
	}
}
