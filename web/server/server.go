package server

import (
	"net/http"
)

import (
	"dummy_ai/pkg/env"
)

func Run() {

	if err := http.ListenAndServe(env.ServerAddress, nil); err != nil {

		panic(err)
	}
}
