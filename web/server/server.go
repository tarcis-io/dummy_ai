package server

import (
	"net/http"
	"text/template"
)

import (
	"dummy_ai/pkg/env"
)

var (
	serverTemplate = template.Must(template.ParseFiles("./web/server/server.html"))
)

func Start() {

	if err := http.ListenAndServe(env.ServerAddress, nil); err != nil {

		panic(err)
	}
}

func serveFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}

func executeServerTemplate(responseWriter http.ResponseWriter, wasmRoute string) {

	if err := serverTemplate.Execute(responseWriter, wasmRoute); err != nil {

		panic(err)
	}
}
