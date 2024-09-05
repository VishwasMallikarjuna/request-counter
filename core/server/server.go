package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server represents an HTTP server with a configured http.Server instance.
type Server struct {
	srv *http.Server // Underlying http.Server instance for handling HTTP requests
}

// New creates and returns a new Server instance.
func New(addr string) Server {
	return Server{
		srv: &http.Server{Addr: addr},
	}
}

// Register sets the HTTP handler for the server.
func (s *Server) Register(handler http.Handler) {
	s.srv.Handler = handler
}

// Run starts the HTTP server and begins listening for incoming requests.
// It also handles graceful shutdown when an interrupt signal is received.
func (s *Server) Run() error {
	// Start the server in a goroutine so that it doesn't block
	go func() {
		log.Printf("Server listening on %s", s.srv.Addr)
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-quit
	log.Println("Server is shutting down...")

	// Create a context with a timeout to ensure shutdown completes
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	log.Println("Server stopped gracefully")
	return nil
}
