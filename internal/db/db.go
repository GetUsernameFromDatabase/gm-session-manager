// Package db opens the SQLite database and applies the embedded migrations.
// Queries live in db/queries and are generated into db/sqlc.
package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"sort"
	"strconv"
	"strings"

	_ "modernc.org/sqlite" // registers the pure-Go "sqlite" driver
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

// connParams are applied by the driver to every new connection in the pool,
// so foreign keys can never be silently off on one of them.
var connParams = []string{
	"_pragma=foreign_keys(1)",
	"_pragma=journal_mode(WAL)",
	"_pragma=synchronous(NORMAL)",
	"_pragma=busy_timeout(5000)",
	// Take the write lock at BEGIN so concurrent writers wait on busy_timeout
	// instead of failing when a read transaction tries to upgrade.
	"_txlock=immediate",
	"_time_format=sqlite",
}

// Open opens (creating if needed) the database at path and migrates it.
func Open(ctx context.Context, path string) (*sql.DB, error) {
	// A migration that edits sqlite_master leaves its own connection with the
	// old schema parsed, so migrate on a throwaway pool and serve from a fresh one.
	boot, err := openPool(ctx, path)
	if err != nil {
		return nil, err
	}
	err = Migrate(ctx, boot)
	_ = boot.Close()
	if err != nil {
		return nil, err
	}
	return openPool(ctx, path)
}

func openPool(ctx context.Context, path string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite", path+"?"+strings.Join(connParams, "&"))
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	if err := conn.PingContext(ctx); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("open database %s: %w", path, err)
	}
	return conn, nil
}

type migration struct {
	version int
	name    string
	sql     string
}

// Migrate applies every embedded migration newer than the database's
// PRAGMA user_version, each in its own transaction.
func Migrate(ctx context.Context, conn *sql.DB) error {
	migrations, err := loadMigrations(migrationsFS)
	if err != nil {
		return err
	}

	var current int
	if err := conn.QueryRowContext(ctx, "PRAGMA user_version").Scan(&current); err != nil {
		return fmt.Errorf("read schema version: %w", err)
	}
	if current > len(migrations) {
		return fmt.Errorf("database schema version %d is newer than this binary (%d); refusing to run", current, len(migrations))
	}

	for _, m := range migrations[current:] {
		if err := apply(ctx, conn, m); err != nil {
			return err
		}
	}
	return nil
}

func apply(ctx context.Context, conn *sql.DB, m migration) error {
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("migration %s: %w", m.name, err)
	}
	defer func() { _ = tx.Rollback() }()

	if _, err := tx.ExecContext(ctx, m.sql); err != nil {
		return fmt.Errorf("migration %s: %w", m.name, err)
	}
	// PRAGMA does not accept bound parameters; version is a parsed int.
	if _, err := tx.ExecContext(ctx, "PRAGMA user_version = "+strconv.Itoa(m.version)); err != nil {
		return fmt.Errorf("migration %s: set version: %w", m.name, err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("migration %s: commit: %w", m.name, err)
	}
	return nil
}

// loadMigrations reads NNNN_name.sql files and requires their numbers to run
// 1, 2, 3, ... with no gaps or duplicates.
func loadMigrations(fsys fs.FS) ([]migration, error) {
	paths, err := fs.Glob(fsys, "migrations/*.sql")
	if err != nil {
		return nil, err
	}
	sort.Strings(paths)

	migrations := make([]migration, 0, len(paths))
	for i, p := range paths {
		name := strings.TrimPrefix(p, "migrations/")
		prefix, _, ok := strings.Cut(name, "_")
		version, err := strconv.Atoi(prefix)
		if !ok || err != nil {
			return nil, fmt.Errorf("migration %s: name must look like 0001_description.sql", name)
		}
		if version != i+1 {
			return nil, fmt.Errorf("migration %s: expected version %d", name, i+1)
		}
		body, err := fs.ReadFile(fsys, p)
		if err != nil {
			return nil, err
		}
		migrations = append(migrations, migration{version: version, name: name, sql: string(body)})
	}
	return migrations, nil
}
