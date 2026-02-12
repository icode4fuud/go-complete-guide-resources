package db

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Migration struct {
	Version  string
	UpPath   string
	DownPath string
	Checksum string
}

// ------------------------------------------------------------
// 1. Locate project root dynamically (bulletproof path resolver)
// ------------------------------------------------------------
func findProjectRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk upward until we find the internal/ folder
	for i := 0; i < 10; i++ {
		try := filepath.Join(wd, "internal")
		if stat, err := os.Stat(try); err == nil && stat.IsDir() {
			return wd, nil
		}
		wd = filepath.Dir(wd)
	}

	return "", errors.New("could not locate project root (missing internal/ folder)")
}

func migrationsDir() (string, error) {
	root, err := findProjectRoot()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(root, "internal", "infrastructure", "db", "DDL")
	if _, err := os.Stat(dir); err != nil {
		return "", fmt.Errorf("migrations directory not found: %s", dir)
	}

	return dir, nil
}

// ------------------------------------------------------------
// 2. Run migrations
// ------------------------------------------------------------
func RunMigrations() error {
	if err := ensureMigrationTable(); err != nil {
		return err
	}

	dir, err := migrationsDir()
	if err != nil {
		return err
	}

	migrations, err := discoverMigrations(dir)
	if err != nil {
		return err
	}

	applied, err := loadAppliedMigrations()
	if err != nil {
		return err
	}

	for _, m := range migrations {
		if _, ok := applied[m.Version]; ok {
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

// ------------------------------------------------------------
// 3. Discover migrations
// ------------------------------------------------------------
func discoverMigrations(dir string) ([]Migration, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	upFiles := map[string]string{}
	downFiles := map[string]string{}

	for _, f := range entries {
		name := f.Name()
		if strings.HasSuffix(name, ".up.sql") {
			version := strings.TrimSuffix(name, ".up.sql")
			upFiles[version] = filepath.Join(dir, name)
		}
		if strings.HasSuffix(name, ".down.sql") {
			version := strings.TrimSuffix(name, ".down.sql")
			downFiles[version] = filepath.Join(dir, name)
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

// ------------------------------------------------------------
// 4. Checksum
// ------------------------------------------------------------
func computeChecksum(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

// ------------------------------------------------------------
// 5. Migration table
// ------------------------------------------------------------
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

// ------------------------------------------------------------
// 6. Rollback
// ------------------------------------------------------------
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

	dir, err := migrationsDir()
	if err != nil {
		return err
	}

	downPath := filepath.Join(dir, version+".down.sql")

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

// ------------------------------------------------------------
// 7. Status
// ------------------------------------------------------------
func PrintMigrationStatus() error {
	rows, err := DB.Query(`
        SELECT version, applied_at
        FROM schema_migrations
        ORDER BY version
    `)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Applied migrations:")
	for rows.Next() {
		var version, appliedAt string
		if err := rows.Scan(&version, &appliedAt); err != nil {
			return err
		}
		fmt.Printf("  %s  (%s)\n", version, appliedAt)
	}

	return nil
}
