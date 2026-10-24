package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CookieName is the session cookie.
const CookieName = "gsm_session"

// cookieMaxAge keeps players logged in across host restarts. 400 days is the
// longest lifetime current browsers honour.
const cookieMaxAge = 400 * 24 * time.Hour

// SetCookie stores the session token on the client.
func SetCookie(c *gin.Context, token string) {
	writeCookie(c, token, int(cookieMaxAge.Seconds()))
}

// ClearCookie removes the session cookie from the client.
func ClearCookie(c *gin.Context) {
	writeCookie(c, "", -1)
}

func writeCookie(c *gin.Context, value string, maxAge int) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     CookieName,
		Value:    value,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
}
