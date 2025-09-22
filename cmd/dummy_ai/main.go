package main

import (
	"log/slog"
	"os"

	"dummy_ai/internal/config"
	"dummy_ai/internal/server"
)

func main() {
	config, err := config.New()
	if err != nil {
		slog.Error("Failed to load configurations", "error", err)
		os.Exit(1)
	}
	server, err := server.New()
	if err != nil {
		slog.Error("Failed to create server", "error", err)
		os.Exit(1)
	}
	if err := server.Start(config.ServerAddress); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
