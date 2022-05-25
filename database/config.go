package database

import "go.uber.org/zap"

// Config ...
type Config struct {
	FilePath string `env:"DB_FILE_PATH"`
}

// Load create new database instance with a unopened database
func (c *Config) Load(log *zap.Logger) (*Database, error) {

	logger := log.Named("database").Sugar()

	return &Database{
		config: c,
		logger: logger}, nil
}
