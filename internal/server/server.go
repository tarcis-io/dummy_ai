package server

import (
	"net/http"
)

import (
	"dummy_ai/internal/env"
)

func listenAndServe() {

	if err := http.ListenAndServe(env.ServerAddress(), nil); err != nil {

		panic(err)
	}
}

func ListenAndServe() {

	listenAndServe()
}
