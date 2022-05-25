package server

import (
	"net"
	"net/http"

	"go.uber.org/zap"
)

// Server
type Server struct {
	c        *Config
	server   *http.Server
	listener net.Listener
	logger   *zap.SugaredLogger
}
