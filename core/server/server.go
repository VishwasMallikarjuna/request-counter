package server

import (
	"fmt"
	"log"
	"net/http"
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
func (s *Server) Run() error {
	log.Printf("Server listening on %s", s.srv.Addr)
	if err := s.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("server error: %w", err)
	}
	return nil
}
