package db

//Why this schema.go better:
// You can add more migration files later
//They run in order
//You get clear logs
//Errors are descriptive

import (
	"fmt"
	"log"
	"os"
)

func RunMigrations() error {
	files := []string{
		"db/DDL/001_init.sql",
	}

	for _, file := range files {
		log.Println("[DB] Applying:", file)

		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", file, err)
		}

		_, err = DB.Exec(string(sqlBytes))
		if err != nil {
			return fmt.Errorf("failed executing %s: %w", file, err)
		}
	}

	log.Println("[DB] All migrations applied successfully")
	return nil
}
