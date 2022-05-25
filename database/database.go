package database

import (
	"context"
	"database/sql"
)

// InTx
func (db *Database) InTx(ctx context.Context, inTx *sql.Tx) error {

	if err := inTx.Commit(); err != nil {
		return inTx.Rollback()
	}

	return nil
}
