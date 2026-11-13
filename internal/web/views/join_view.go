package views

import "strings"

// JoinInfo is what the GM shows players so they can connect.
type JoinInfo struct {
	URLs        []string
	QR          string // src of the QR code image
	Fingerprint string
}

// fingerprintPairs spaces the hex fingerprint into uppercase byte pairs, as
// Chrome's certificate viewer shows it, so it wraps and can be compared.
func fingerprintPairs(hex string) string {
	var b strings.Builder
	for i := 0; i < len(hex); i += 2 {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(strings.ToUpper(hex[i:min(i+2, len(hex))]))
	}
	return b.String()
}

// Waiting is a visitor's join request as they see it.
type Waiting struct {
	Secret string
	Name   string
	State  string // "waiting", "approved", "denied"
}
