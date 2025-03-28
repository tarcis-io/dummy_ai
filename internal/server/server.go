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

	pages := map[string]string{
		"/":      "/home.wasm",
		"/about": "/about.wasm",
	}

	for route, wasmRoute := range pages {

		handlePage(route, wasmRoute)
	}

	listenAndServe()
}

func handleFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}

func handlePage(route string, wasmRoute string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		if request.URL.Path != route {

			handlePageError404(responseWriter)
			return
		}

		executeServerTemplate(responseWriter, wasmRoute)
	})
}

func handlePageError404(responseWriter http.ResponseWriter) {

	responseWriter.WriteHeader(http.StatusNotFound)
	executeServerTemplate(responseWriter, "/error_404.wasm")
}

func executeServerTemplate(responseWriter http.ResponseWriter, wasmRoute string) {

	if err := serverTemplate.Execute(responseWriter, wasmRoute); err != nil {

		panic(err)
	}
}

func listenAndServe() {

	if err := http.ListenAndServe(env.ServerAddress(), nil); err != nil {

		panic(err)
	}
}
