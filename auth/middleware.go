package auth

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type contextKey int

const identityKey contextKey = iota

// Middleware resolves the session cookie and admin status for every request
// and stores the Identity in the gin context. A cookie that no longer maps to
// a session (logged out, kicked) is cleared.
func (m *Manager) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := Identity{Loopback: IsLoopback(c.Request)}

		if token, err := c.Cookie(CookieName); err == nil {
			sess, err := m.Resolve(c.Request.Context(), token)
			switch {
			case err == nil:
				id.Session = &sess
			case errors.Is(err, ErrNoSession):
				ClearCookie(c)
			default:
				slog.Error("session lookup failed", "err", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}

		c.Set(identityKey, id)
		c.Next()
	}
}

// FromContext returns the Identity set by Middleware.
func FromContext(c *gin.Context) Identity {
	v, _ := c.Get(identityKey)
	id, _ := v.(Identity)
	return id
}

// RequireAdmin rejects requests without admin rights.
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !FromContext(c).IsAdmin() {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
