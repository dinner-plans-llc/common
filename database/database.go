package database

import (
	"context"
	"database/sql"
)

// InTx
func (db *Database) InTx(ctx context.Context, inTx *sql.Tx) error {

	conn, err := db.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {

		if err1 := tx.Rollback(); err1 != nil {
			return err1
		}

		return err
	}

	return nil
}
