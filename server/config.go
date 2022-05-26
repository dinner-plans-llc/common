package server

import (
	"net"

	"go.uber.org/zap"
)

const (
	_network = "tcp"
)

// Config
type Config struct {
	Host string `env:"SERVER_HOST, default=0.0.0.0" json:"server_host"`
	Port string `env:"SERVER_PORT, default=8080" json:"server_port"`
}

// Load
func (c *Config) Load(log *zap.Logger) (*Server, error) {

	logger := log.Named("http controller").Sugar()

	listener, err := net.Listen(_network, net.JoinHostPort(c.Host, c.Port))
	if err != nil {
		logger.Error("failed to create listener: ", err)
		return nil, err
	}

	return &Server{
		c:        c,
		logger:   logger,
		listener: listener,
	}, nil
}
