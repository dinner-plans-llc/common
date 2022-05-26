package server

import (
	"net"

	"go.uber.org/zap"
)

// Server
type Server struct {
	c        *Config
	listener net.Listener
	logger   *zap.SugaredLogger
}
