package main

import (
	"dummy_ai/internal/server"
)

func main() {

	server.New().ListenAndServe()
}
