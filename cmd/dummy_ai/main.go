// Package main is the entry point for the dummy_ai application.
package main

import (
	"log/slog"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.Info("starting application")
	if err := run(); err != nil {
		slog.Error("application failed", "error", err)
		os.Exit(1)
	}
	slog.Info("application exited")
}

func run() error {
	return nil
}
