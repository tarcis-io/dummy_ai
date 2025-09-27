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
	Server struct {
		address         string
		router          http.Handler
		shutdownTimeout time.Duration
	}
)

func New(config *config.Config) (*Server, error) {
	router := http.NewServeMux()
	server := &Server{
		address:         config.ServerAddress,
		shutdownTimeout: config.ServerShutdownTimeout,
		router:          router,
	}
	return server, nil
}

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
