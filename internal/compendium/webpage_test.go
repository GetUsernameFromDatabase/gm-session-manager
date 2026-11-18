package compendium

import (
	"net/http"
	"net/url"
	"testing"
)

// A page fetch follows a redirect only within the same scheme, host and port.
func TestPageClientRedirect(t *testing.T) {
	from, _ := url.Parse("https://wiki.example:8443/page")
	for _, tt := range []struct {
		name, to string
		ok       bool
	}{
		{"https to http", "http://wiki.example:8443/page", false},
		{"other port", "https://wiki.example:9000/page", false},
		{"other host", "https://other.example:8443/page", false},
		{"same origin", "https://WIKI.example:8443/other", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			to, _ := url.Parse(tt.to)
			err := pageClient.CheckRedirect(&http.Request{URL: to}, []*http.Request{{URL: from}})
			if (err == nil) != tt.ok {
				t.Errorf("CheckRedirect = %v, want allowed %v", err, tt.ok)
			}
		})
	}
}
