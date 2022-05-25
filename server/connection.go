package server

import "context"

// ServeHTTP
func (s *Server) ServeHTTP() error {
	return s.server.Serve(s.listener)
}

// Shutdown
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
