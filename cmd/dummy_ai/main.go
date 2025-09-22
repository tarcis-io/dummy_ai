package main

import (
	"os"

	"dummy_ai/internal/config"
	"dummy_ai/internal/server"
)

func main() {
	config, err := config.New()
	if err != nil {
		os.Exit(1)
	}
	server, err := server.New()
	if err != nil {
		os.Exit(1)
	}
	if err := server.Start(config.ServerAddress); err != nil {
		os.Exit(1)
	}
}
