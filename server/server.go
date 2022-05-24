package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	_network = "tcp"
)

// Server
type Server struct {
	host     string
	port     string
	server   *http.Server
	listener net.Listener
}

// ProvideServer
func ProvideServer(lc fx.Lifecycle, logger *zap.Logger, srv *http.Server) (*Server, error) {

	var s Server

	s.host = os.Getenv("SERVER_HOST")
	s.port = os.Getenv("SERVER_PORT")

	listener, err := net.Listen(_network, net.JoinHostPort(s.host, s.port))
	if err != nil {
		logger.Error(err.Error())
		return nil, fmt.Errorf("net.Listen err=%v", err)
	}

	s.server = srv
	s.listener = listener

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return s.server.Serve(s.listener)
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return &s, nil
}
