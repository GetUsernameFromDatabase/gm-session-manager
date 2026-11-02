package auth

import (
	"crypto/sha256"
	"crypto/subtle"
	"net"
	"net/http"
)

// IsLoopback reports whether the request's real TCP peer is on this machine.
// It reads only RemoteAddr; forwarding headers such as X-Forwarded-For are
// client-controlled and deliberately ignored.
func IsLoopback(r *http.Request) bool {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return false
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}

// credentialsMatch compares both values in constant time. Hashing first makes
// the comparison independent of the lengths involved, and both comparisons
// always run so timing does not reveal which one failed.
func credentialsMatch(wantUser, wantPass, gotUser, gotPass string) bool {
	userOK := constantTimeEqual(wantUser, gotUser)
	passOK := constantTimeEqual(wantPass, gotPass)
	return userOK&passOK == 1
}

func constantTimeEqual(want, got string) int {
	w := sha256.Sum256([]byte(want))
	g := sha256.Sum256([]byte(got))
	return subtle.ConstantTimeCompare(w[:], g[:])
}
