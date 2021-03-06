package database

import (
	"context"
	"database/sql"

	"go.uber.org/zap"

	_ "github.com/mattn/go-sqlite3"
)

const (
	_driverName = "sqlite3"
)

// Database
type Database struct {
	db     *sql.DB
	logger *zap.SugaredLogger
	config *Config
}

// Open
func (db *Database) Open() error {

	var err error

	db.db, err = sql.Open(_driverName, db.config.FilePath)
	if err != nil {
		db.logger.Errorf("failed to open database ", err)
		return err
	}

	return nil
}

// Close
func (db *Database) Close() error {
	return db.db.Close()
}

// Conn
func (db *Database) Conn(ctx context.Context) (*sql.Conn, error) {
	return db.db.Conn(ctx)
}
