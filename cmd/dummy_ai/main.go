// Package main is the entry point of the dummy_ai application.
package main

import (
	"log/slog"
	"os"

	"dummy_ai/internal/config"
	"dummy_ai/internal/server"
)

// main is the entry point of the dummy_ai application.
// It loads the configurations, creates and starts the server.
func main() {
	slog.Info("loading configurations...")
	config, err := config.New()
	if err != nil {
		slog.Error("failed to load configurations", "error", err)
		os.Exit(1)
	}
	slog.Info("creating server...")
	server, err := server.New()
	if err != nil {
		slog.Error("failed to create server", "error", err)
		os.Exit(1)
	}
	slog.Info("starting server...", "address", config.ServerAddress)
	if err := server.Start(config.ServerAddress); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
