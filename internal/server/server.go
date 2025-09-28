// Package server provides functionality for creating and running an HTTP server.
package server

import (
	"context"
	"embed"
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
	// Server is an HTTP server.
	Server struct {
		// address is the host and port for the server to listen on.
		address string

		// router is the HTTP request multiplexer.
		router http.Handler

		// shutdownTimeout is the maximum duration for a graceful shutdown.
		shutdownTimeout time.Duration
	}
)

const (
	staticFilesDirectory  = "web/static"
	staticFilesPathPrefix = "/static/"
)

var (
	//go:embed web
	webFS embed.FS
)

// New creates and returns a new Server instance based on the provided configuration.
func New(config *config.Config) (*Server, error) {
	server := &Server{
		address:         config.ServerAddress,
		shutdownTimeout: config.ServerShutdownTimeout,
		router:          http.NewServeMux(),
	}
	if err := server.registerStaticFiles(); err != nil {
		return nil, err
	}
	return server, nil
}

// Run starts the HTTP server and blocks until a graceful shutdown signal is received.
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

func (server *Server) registerStaticFiles() error {
	return nil
}
