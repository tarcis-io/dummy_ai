// Package server provides functionality for creating and running an HTTP server.
package server

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"io/fs"
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
		router *http.ServeMux

		// shutdownTimeout is the maximum duration for a graceful shutdown.
		shutdownTimeout time.Duration
	}
)

const (
	staticFilesPathPrefix = "/static/"
	staticFilesDirectory  = "web/static"
	faviconStaticFilePath = "/favicon.ico"
	faviconStaticFile     = "/img/favicon/favicon.ico"
	robotsStaticFilePath  = "/robots.txt"
	robotsStaticFile      = "/config/robots.txt"
	sitemapStaticFilePath = "/sitemap.xml"
	sitemapStaticFile     = "/config/sitemap.xml"

	cacheControlHeaderKey            = "Cache-Control"
	cacheControlHeaderValue          = "public, max-age=86400"
	contentSecurityPolicyHeaderKey   = "Content-Security-Policy"
	contentSecurityPolicyHeaderValue = "default-src 'self';"
	xContentTypeOptionsHeaderKey     = "X-Content-Type-Options"
	xContentTypeOptionsHeaderValue   = "nosniff"
	xFrameOptionsHeaderKey           = "X-Frame-Options"
	xFrameOptionsHeaderValue         = "DENY"
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

// registerStaticFiles registers the static files handler on the server's router.
// It returns an error if the static files directory cannot be opened.
func (server *Server) registerStaticFiles() error {
	staticFilesFS, err := fs.Sub(webFS, staticFilesDirectory)
	if err != nil {
		return fmt.Errorf("failed to open static files directory: %w", err)
	}
	staticFilesHeaders := map[string]string{
		cacheControlHeaderKey:          cacheControlHeaderValue,
		contentSecurityPolicyHeaderKey: contentSecurityPolicyHeaderValue,
		xContentTypeOptionsHeaderKey:   xContentTypeOptionsHeaderValue,
		xFrameOptionsHeaderKey:         xFrameOptionsHeaderValue,
	}
	staticFilesHandler := withHeaders(staticFilesHeaders, http.FileServerFS(staticFilesFS))
	server.router.Handle(staticFilesPathPrefix, http.StripPrefix(staticFilesPathPrefix, staticFilesHandler))
	rootStaticFiles := map[string]string{
		faviconStaticFilePath: faviconStaticFile,
		robotsStaticFilePath:  robotsStaticFile,
		sitemapStaticFilePath: sitemapStaticFile,
	}
	for rootStaticFilePath, rootStaticFile := range rootStaticFiles {
		rootStaticFileHandler := http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
			http.ServeFileFS(responseWriter, request, staticFilesFS, rootStaticFile)
		})
		server.router.Handle(rootStaticFilePath, withHeaders(staticFilesHeaders, rootStaticFileHandler))
	}
	return nil
}

// withHeaders is a middleware function that sets the provided headers on the response.
func withHeaders(headers map[string]string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		for headerKey, headerValue := range headers {
			responseWriter.Header().Set(headerKey, headerValue)
		}
		next.ServeHTTP(responseWriter, request)
	})
}
