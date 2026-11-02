// Package auth handles device sessions (the sessions table), the session
// cookie, and admin detection. It knows nothing about game sessions.
package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// tokenBytes is the amount of randomness in a session token.
const tokenBytes = 32

var tokenEncoding = base64.RawURLEncoding

// newToken returns a random token for the cookie and the SHA-256 hash that is
// the only thing stored in the database.
func newToken() (token string, hash []byte, err error) {
	raw := make([]byte, tokenBytes)
	if _, err := rand.Read(raw); err != nil {
		return "", nil, err
	}
	sum := sha256.Sum256(raw)
	return tokenEncoding.EncodeToString(raw), sum[:], nil
}

// hashToken decodes a cookie token and returns its hash. ok is false when the
// value is not a well-formed token, so it never reaches the database.
func hashToken(token string) (hash []byte, ok bool) {
	raw, err := tokenEncoding.DecodeString(token)
	if err != nil || len(raw) != tokenBytes {
		return nil, false
	}
	sum := sha256.Sum256(raw)
	return sum[:], true
}
