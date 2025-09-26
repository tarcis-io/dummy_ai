// Package main is the entry point for the dummy_ai application.
package main

import (
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

func run() error {
	config, err := config.New()
	if err != nil {
		return err
	}
	server, err := server.New(config)
	if err != nil {
		return err
	}
	if err := server.Run(); err != nil {
		return err
	}
	return nil
}
