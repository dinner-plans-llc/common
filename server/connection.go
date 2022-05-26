package server

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
)

// ServeHTTP
func (s *Server) ServeHTTP(srv *http.Server) error {
	return srv.Serve(s.listener)
}

// Shutdown
func (s *Server) Shutdown(ctx context.Context, srv *http.Server) error {
	return srv.Shutdown(ctx)
}

// ServeGRPC
func (s *Server) ServeGRPC(ctx context.Context, srv *grpc.Server) error {
	return srv.Serve(s.listener)
}
