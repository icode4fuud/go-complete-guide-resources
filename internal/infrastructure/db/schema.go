// Why this schema.go better:
// You can add more migration files later
// They run in order
// You get clear logs
// Errors are descriptive
// Versioned migrations
//Name files like:
//DDL/001_init.sql
//DDL/002_add_index_on_events_datetime.sql

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type migration struct {
	Version string
	Path    string
}

func RunMigrations() error {
	if err := ensureMigrationsTable(); err != nil {
		return err
	}

	applied, err := loadAppliedVersions()
	if err != nil {
		return err
	}

	pending, err := discoverMigrations("../../internal/infrastructure/db/DDL")
	if err != nil {
		return err
	}

	for _, m := range pending {
		if applied[m.Version] {
			continue
		}

		log.Println("[DB] Applying migration:", m.Version, m.Path)

		sqlBytes, err := os.ReadFile(m.Path)
		if err != nil {
			return fmt.Errorf("read %s: %w", m.Path, err)
		}

		if _, err := DB.Exec(string(sqlBytes)); err != nil {
			return fmt.Errorf("exec %s: %w", m.Path, err)
		}

		if err := recordMigration(m.Version); err != nil {
			return err
		}
	}

	return nil
}

func ensureMigrationsTable() error {
	_, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS schema_migrations (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            version TEXT NOT NULL UNIQUE,
            applied_at TEXT NOT NULL
        );
    `)
	return err
}

func loadAppliedVersions() (map[string]bool, error) {
	rows, err := DB.Query(`SELECT version FROM schema_migrations`)
	if err != nil {
		if err == sql.ErrNoRows {
			return map[string]bool{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	m := make(map[string]bool)
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			return nil, err
		}
		m[v] = true
	}
	return m, rows.Err()
}

func discoverMigrations(dir string) ([]migration, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var ms []migration
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		// e.g. "001_init.sql" → version "001_init"
		version := name[:len(name)-len(filepath.Ext(name))]
		ms = append(ms, migration{
			Version: version,
			Path:    filepath.Join(dir, name),
		})
	}

	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Version < ms[j].Version
	})

	return ms, nil
}

func recordMigration(version string) error {
	_, err := DB.Exec(
		`INSERT INTO schema_migrations (version, applied_at) VALUES (?, ?)`,
		version,
		time.Now().Format(time.RFC3339),
	)
	return err
}
