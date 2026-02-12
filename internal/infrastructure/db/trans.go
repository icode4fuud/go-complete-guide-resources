// db/tx.go
// Add simple helpers so callers don’t repeat boilerplate.
//Usage in a service/repo
//err := db.WithTx(context.Background(), func(tx *sql.Tx) error {
//  _, err := tx.Exec(`INSERT INTO events (...) VALUES (...)`, ...)
//  return err
//})

package db

import (
	"context"
	"database/sql"
	"fmt"
)

func WithTx(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx rollback error: %v (original: %w)", rbErr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
