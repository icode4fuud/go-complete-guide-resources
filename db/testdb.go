// db/testdb.go
package db

import (
	"database/sql"
	//"os"
	"testing"

	_ "modernc.org/sqlite"
)

func NewTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("open test db: %v", err)
	}

	DB = db // reuse existing migration logic if you want
	if err := RunMigrations(); err != nil {
		t.Fatalf("migrate test db: %v", err)
	}

	return db
}
