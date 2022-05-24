package database

import (
	"database/sql"
)

const (
	_driverName = "sqlite3"
)

// Database
type Database struct {
	db *sql.DB
}

// ProvideDatabase
func ProvideDatabase(dsn string) (*Database, error) {

	db, err := sql.Open(_driverName, dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return &Database{db: db}, nil
}
