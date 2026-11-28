package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gm-session-manager/internal/db"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
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
	// player is the player app: no admin, its own cookie.
	player bool
}

// NewManager returns a Manager that checks admin logins against the given
// credentials (from .env).
func NewManager(q *db.Store, adminUser, adminPass string) *Manager {
	return &Manager{q: q, adminUser: adminUser, adminPass: adminPass}
}

// NewPlayerManager returns the player app's Manager: no credentials, and
// loopback is not admin.
func NewPlayerManager(q *db.Store) *Manager {
	return &Manager{q: q, player: true}
}

// IssuePlayer creates a session bound to a character and returns its token.
func (m *Manager) IssuePlayer(
	ctx context.Context,
	characterID int64,
) (string, sqlc.Session, error) {
	return m.issue(ctx, enum.RolePlayer, sql.NullInt64{Int64: characterID, Valid: true})
}

// LoginAdmin checks the credentials and, if they match, creates an admin
// session with no character.
func (m *Manager) LoginAdmin(ctx context.Context, user, pass string) (string, sqlc.Session, error) {
	// The player app has no credentials; two empty pairs would otherwise match.
	if m.adminUser == "" || !credentialsMatch(m.adminUser, m.adminPass, user, pass) {
		return "", sqlc.Session{}, ErrBadCredentials
	}
	return m.issue(ctx, enum.RoleAdmin, sql.NullInt64{})
}

func (m *Manager) issue(
	ctx context.Context,
	role enum.Role,
	characterID sql.NullInt64,
) (string, sqlc.Session, error) {
	token, hash, err := newToken()
	if err != nil {
		return "", sqlc.Session{}, fmt.Errorf("generate token: %w", err)
	}
	sess, err := m.q.CreateSession(ctx, sqlc.CreateSessionParams{
		TokenHash:   hash,
		Role:        role,
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
	// An UPDATE takes the write lock even when it changes nothing, so a fresh
	// row is not touched: requests then never queue behind a long import.
	if time.Since(sess.LastSeenAt) >= touchEvery {
		if err := m.q.TouchSession(ctx, sess.ID); err != nil {
			return sqlc.Session{}, fmt.Errorf("touch session: %w", err)
		}
	}
	return sess, nil
}

// touchEvery matches the one minute in the TouchSession query.
const touchEvery = time.Minute

// Revoke deletes a session. Logout and kick both use this.
func (m *Manager) Revoke(ctx context.Context, sessionID int64) error {
	if err := m.q.DeleteSession(ctx, sessionID); err != nil {
		return fmt.Errorf("revoke session: %w", err)
	}
	return nil
}
