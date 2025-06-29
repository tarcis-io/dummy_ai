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

func ListenAndServe() {

	listenAndServe()
}

func serveFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}

func servePage(route string, wasmRoute string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		if request.URL.Path != route {

			servePageError404(responseWriter)
			return
		}

		executeServerTemplate(responseWriter, wasmRoute)
	})
}

func servePageError404(responseWriter http.ResponseWriter) {

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
