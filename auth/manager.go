package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gm-session-manager/db"
	"gm-session-manager/db/sqlc"
)

// Role values stored in sessions.role.
type Role string

// Session roles. The GM is an admin; there is no separate GM role.
const (
	RolePlayer Role = "player"
	RoleAdmin  Role = "admin"
)

var (
	// ErrNoSession means the token does not match any sessions row: it was
	// never issued, or the device was logged out or kicked.
	ErrNoSession = errors.New("auth: no such session")
	// ErrBadCredentials means an admin login did not match .env.
	ErrBadCredentials = errors.New("auth: invalid admin credentials")
)

// Manager issues, resolves and revokes device sessions.
type Manager struct {
	q         *db.Store
	adminUser string
	adminPass string
}

// NewManager returns a Manager that checks admin logins against the given
// credentials (from .env).
func NewManager(q *db.Store, adminUser, adminPass string) *Manager {
	return &Manager{q: q, adminUser: adminUser, adminPass: adminPass}
}

// IssuePlayer creates a session bound to a character and returns its token.
func (m *Manager) IssuePlayer(ctx context.Context, characterID int64) (string, sqlc.Session, error) {
	return m.issue(ctx, RolePlayer, sql.NullInt64{Int64: characterID, Valid: true})
}

// LoginAdmin checks the credentials and, if they match, creates an admin
// session with no character.
func (m *Manager) LoginAdmin(ctx context.Context, user, pass string) (string, sqlc.Session, error) {
	if !credentialsMatch(m.adminUser, m.adminPass, user, pass) {
		return "", sqlc.Session{}, ErrBadCredentials
	}
	return m.issue(ctx, RoleAdmin, sql.NullInt64{})
}

func (m *Manager) issue(ctx context.Context, role Role, characterID sql.NullInt64) (string, sqlc.Session, error) {
	token, hash, err := newToken()
	if err != nil {
		return "", sqlc.Session{}, fmt.Errorf("generate token: %w", err)
	}
	sess, err := m.q.CreateSession(ctx, sqlc.CreateSessionParams{
		TokenHash:   hash,
		Role:        string(role),
		CharacterID: characterID,
	})
	if err != nil {
		return "", sqlc.Session{}, fmt.Errorf("create session: %w", err)
	}
	return token, sess, nil
}

// Resolve looks up the session for a cookie token and records activity.
func (m *Manager) Resolve(ctx context.Context, token string) (sqlc.Session, error) {
	hash, ok := hashToken(token)
	if !ok {
		return sqlc.Session{}, ErrNoSession
	}
	sess, err := m.q.GetSessionByTokenHash(ctx, hash)
	if errors.Is(err, sql.ErrNoRows) {
		return sqlc.Session{}, ErrNoSession
	}
	if err != nil {
		return sqlc.Session{}, fmt.Errorf("resolve session: %w", err)
	}
	if err := m.q.TouchSession(ctx, sess.ID); err != nil {
		return sqlc.Session{}, fmt.Errorf("touch session: %w", err)
	}
	return sess, nil
}

// Revoke deletes a session. Logout and kick both use this.
func (m *Manager) Revoke(ctx context.Context, sessionID int64) error {
	if err := m.q.DeleteSession(ctx, sessionID); err != nil {
		return fmt.Errorf("revoke session: %w", err)
	}
	return nil
}
