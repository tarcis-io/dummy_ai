// Package main is the entry point for the dummy_ai application.
package main

import (
	"fmt"
	"log/slog"
	"os"

	"dummy_ai/internal/config"
	"dummy_ai/internal/server"
)

// main is the entry point for the dummy_ai application.
func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.Info("Starting application")
	if err := run(); err != nil {
		slog.Error("Application failed", "error", err)
		os.Exit(1)
	}
	slog.Info("Application stopped successfully")
}

// run loads the configurations, creates a new server and runs it.
// It returns an error if any of the steps fail.
func run() error {
	config, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	server, err := server.New(config)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	if err := server.Run(); err != nil {
		return fmt.Errorf("server stopped unexpectedly: %w", err)
	}
	return nil
}
