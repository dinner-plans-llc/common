package server

import (
	"net"
	"net/http"

	"go.uber.org/zap"
)

const (
	_network = "tcp"
)

// Config
type Config struct {
	Host string
	Port string
}

// Load
func (c *Config) Load(log *zap.Logger, handler http.Handler) (*Server, error) {

	logger := log.Named("http controller").Sugar()

	listener, err := net.Listen(_network, net.JoinHostPort(c.Host, c.Port))
	if err != nil {
		logger.Error("failed to create listener: ", err)
		return nil, err
	}

	server := &http.Server{
		ErrorLog: zap.NewStdLog(log),
		Handler:  handler,
	}

	return &Server{
		c:        c,
		logger:   logger,
		listener: listener,
		server:   server,
	}, nil
}
