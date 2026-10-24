package auth

import "gm-session-manager/db/sqlc"

// Identity is who is making a request. A device has at most one session,
// whose role was chosen at login; switching role means logging out and in
// again. Localhost needs no session to be admin, so on the host machine
// logging out of a player session is all it takes to be admin again.
type Identity struct {
	// Session is nil when the request carries no valid session cookie.
	Session *sqlc.Session
	// Loopback is true when the TCP peer is this machine.
	Loopback bool
}

// IsAdmin reports whether this request has admin rights: an admin session,
// or localhost without a session. Localhost holding a player session acts as
// that player only, so the admin and player views never mix.
func (id Identity) IsAdmin() bool {
	if id.Session != nil {
		return Role(id.Session.Role) == RoleAdmin
	}
	return id.Loopback
}

// CharacterID returns the character a player session is bound to.
func (id Identity) CharacterID() (int64, bool) {
	if id.Session == nil || !id.Session.CharacterID.Valid {
		return 0, false
	}
	return id.Session.CharacterID.Int64, true
}
