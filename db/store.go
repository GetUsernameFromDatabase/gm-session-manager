package db

import (
	"context"
	"database/sql"
	"fmt"

	"gm-session-manager/db/sqlc"
)

// Store bundles the generated queries with the connection pool so callers can
// run single queries directly or group several in a transaction with InTx.
type Store struct {
	*sqlc.Queries
	conn *sql.DB
}

// NewStore wraps an open database.
func NewStore(conn *sql.DB) *Store {
	return &Store{Queries: sqlc.New(conn), conn: conn}
}

// InTx runs fn in a transaction, committing if it returns nil and rolling
// back otherwise. The connection uses _txlock=immediate, so the write lock is
// taken at BEGIN and concurrent writers queue on busy_timeout.
func (s *Store) InTx(ctx context.Context, fn func(q *sqlc.Queries) error) error {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	if err := fn(s.WithTx(tx)); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	return nil
}
