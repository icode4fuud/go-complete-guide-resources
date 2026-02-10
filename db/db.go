package db

//Why this version of db.go is better:
// No shadowing
// Ping verifies the DB is actually usable
// Errors bubble up instead of panicking
// Logging gives you visibility
// Migrations are separated cleanly

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() error {
	var err error

	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Hour)

	// Verify the connection
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	log.Println("[DB] Connected successfully")

	// Run schema setup
	if err := RunMigrations(); err != nil {
		return fmt.Errorf("migration error: %w", err)
	}

	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("[DB] Connection closed")
	}
}
