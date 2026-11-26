package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CookieName is the GM app's session cookie.
const CookieName = "gsm_session"

// PlayerCookieName is the player app's: cookies ignore ports, so with one name
// each app would clear the other's cookie on 127.0.0.1.
const PlayerCookieName = "gsm_home"

// cookieMaxAge keeps players logged in across host restarts. 400 days is the
// longest lifetime current browsers honour.
const cookieMaxAge = 400 * 24 * time.Hour

// CookieName returns this app's session cookie name.
func (m *Manager) CookieName() string {
	if m.player {
		return PlayerCookieName
	}
	return CookieName
}

// SetCookie stores the session token on the client.
func (m *Manager) SetCookie(c *gin.Context, token string) {
	m.writeCookie(c, token, int(cookieMaxAge.Seconds()))
}

// ClearCookie removes the session cookie from the client.
func (m *Manager) ClearCookie(c *gin.Context) {
	m.writeCookie(c, "", -1)
}

func (m *Manager) writeCookie(c *gin.Context, value string, maxAge int) {
	http.SetCookie(c.Writer, &http.Cookie{ //nolint:gosec // Secure is off only in the player app: plain HTTP on loopback
		Name:     m.CookieName(),
		Value:    value,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   !m.player, // the player app is plain HTTP on loopback
		SameSite: http.SameSiteLaxMode,
	})
}
