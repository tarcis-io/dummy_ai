package server

import (
	"net/http"
)

func Start() {

	listenAndServe()
}

func listenAndServe() {

	if err := http.ListenAndServe(":8080", nil); err != nil {

		panic(err)
	}
}
