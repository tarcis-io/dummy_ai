package server

import (
	"context"
	"errors"
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
		router          *http.ServeMux
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
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errorChan:
		return err
	case <-shutdownChan:
		context, cancel := context.WithTimeout(context.Background(), server.shutdownTimeout)
		defer cancel()
		if err := httpServer.Shutdown(context); err != nil {
			return err
		}
		return nil
	}
}
