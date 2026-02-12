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
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const migrationsDir = "../../internal/infrastructure/db/DDL"

type Migration struct {
	Version  string
	UpPath   string
	DownPath string
	Checksum string
}

func RunMigrations() error {
	if err := ensureMigrationTable(); err != nil {
		return err
	}

	migrations, err := discoverMigrations()
	if err != nil {
		return err
	}

	applied, err := loadAppliedMigrations()
	if err != nil {
		return err
	}

	for _, m := range migrations {
		if _, ok := applied[m.Version]; ok {
			// Already applied — verify checksum
			if applied[m.Version] != m.Checksum {
				return fmt.Errorf("checksum mismatch for migration %s", m.Version)
			}
			continue
		}

		fmt.Println("[DB] Applying migration:", m.Version)

		sqlBytes, err := os.ReadFile(m.UpPath)
		if err != nil {
			return fmt.Errorf("read %s: %w", m.UpPath, err)
		}

		if _, err := DB.Exec(string(sqlBytes)); err != nil {
			return fmt.Errorf("exec %s: %w", m.UpPath, err)
		}

		if err := recordMigration(m.Version, m.Checksum); err != nil {
			return err
		}

		fmt.Println("[DB] Migration applied:", m.Version)
	}

	return nil
}

func ensureMigrationTable() error {
	_, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS schema_migrations (
            version TEXT PRIMARY KEY,
            checksum TEXT NOT NULL,
            applied_at TEXT NOT NULL
        );
    `)
	return err
}

func discoverMigrations() ([]Migration, error) {
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return nil, err
	}

	upFiles := map[string]string{}
	downFiles := map[string]string{}

	for _, f := range files {
		name := f.Name()
		if strings.HasSuffix(name, ".up.sql") {
			version := strings.TrimSuffix(name, ".up.sql")
			upFiles[version] = filepath.Join(migrationsDir, name)
		}
		if strings.HasSuffix(name, ".down.sql") {
			version := strings.TrimSuffix(name, ".down.sql")
			downFiles[version] = filepath.Join(migrationsDir, name)
		}
	}

	var migrations []Migration
	for version, up := range upFiles {
		down := downFiles[version]

		checksum, err := computeChecksum(up)
		if err != nil {
			return nil, err
		}

		migrations = append(migrations, Migration{
			Version:  version,
			UpPath:   up,
			DownPath: down,
			Checksum: checksum,
		})
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

func computeChecksum(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

func loadAppliedMigrations() (map[string]string, error) {
	rows, err := DB.Query(`SELECT version, checksum FROM schema_migrations`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	applied := map[string]string{}
	for rows.Next() {
		var version, checksum string
		if err := rows.Scan(&version, &checksum); err != nil {
			return nil, err
		}
		applied[version] = checksum
	}
	return applied, nil
}

func recordMigration(version, checksum string) error {
	_, err := DB.Exec(`
        INSERT INTO schema_migrations (version, checksum, applied_at)
        VALUES (?, ?, ?)
    `, version, checksum, time.Now().Format(time.RFC3339))
	return err
}

func RollbackLastMigration() error {
	rows, err := DB.Query(`
        SELECT version FROM schema_migrations
        ORDER BY applied_at DESC LIMIT 1
    `)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return fmt.Errorf("no migrations to rollback")
	}

	var version string
	if err := rows.Scan(&version); err != nil {
		return err
	}

	downPath := filepath.Join(migrationsDir, version+".down.sql")

	sqlBytes, err := os.ReadFile(downPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", downPath, err)
	}

	if _, err := DB.Exec(string(sqlBytes)); err != nil {
		return fmt.Errorf("exec %s: %w", downPath, err)
	}

	_, err = DB.Exec(`DELETE FROM schema_migrations WHERE version = ?`, version)
	return err
}
