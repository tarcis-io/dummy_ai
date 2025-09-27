// Package server provides functionality for creating and running an HTTP server.
package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dummy_ai/internal/config"
)

type (
	// Server represents an HTTP server.
	Server struct {
		// address is the host and port for the server to listen on.
		address string

		// router is the HTTP request multiplexer for the server.
		router http.Handler

		// shutdownTimeout is the timeout for the server to shut down gracefully.
		shutdownTimeout time.Duration
	}
)

// New creates and returns a new Server instance.
// It takes a config.Config instance as an argument
// and returns an error if the server cannot be created.
func New(config *config.Config) (*Server, error) {
	router := http.NewServeMux()
	server := &Server{
		address:         config.ServerAddress,
		shutdownTimeout: config.ServerShutdownTimeout,
		router:          router,
	}
	return server, nil
}

// Run starts the HTTP server, listens for incoming requests and handles graceful shutdown.
// This function blocks until the server is shut down.
// It returns an error if an error occurs while starting, running or shutting down the server.
func (server *Server) Run() error {
	httpServer := &http.Server{
		Addr:    server.address,
		Handler: server.router,
	}
	errorChan := make(chan error, 1)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errorChan <- err
		}
	}()
	shutdownSignalChan := make(chan os.Signal, 1)
	signal.Notify(shutdownSignalChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errorChan:
		return fmt.Errorf("runtime server error: %w", err)
	case <-shutdownSignalChan:
		context, cancel := context.WithTimeout(context.Background(), server.shutdownTimeout)
		defer cancel()
		if err := httpServer.Shutdown(context); err != nil {
			return fmt.Errorf("failed to shut down server gracefully: %w", err)
		}
		return nil
	}
}
