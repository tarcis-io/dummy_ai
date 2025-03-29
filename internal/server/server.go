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

	staticFiles := map[string]string{
		"/wasm_exec.js":  "./static/lib/wasm/wasm_exec.js",
		"/wasm_start.js": "./static/lib/wasm/wasm_start.js",
	}

	for route, staticFile := range staticFiles {

		handleFile(route, staticFile)
	}

	listenAndServe()
}

func handleFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}

func listenAndServe() {

	if err := http.ListenAndServe(env.ServerAddress(), nil); err != nil {

		panic(err)
	}
}
