package web

import (
	"testing"

	"github.com/mileusna/useragent"
)

func TestUserAgent(t *testing.T) {
	tests := []struct {
		ua, summary, browser string
		phone                bool
	}{
		{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0", "Windows · Edge", "Edge 130 on Windows", false},
		{"Mozilla/5.0 (Linux; Android 14; Pixel 8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.6723.58 Mobile Safari/537.36", "Android · Chrome", "Chrome 130 on Android", true},
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Mobile/15E148 Safari/604.1", "iOS · Safari", "Safari 17 on iOS", true},
		{"Mozilla/5.0 (Macintosh; Intel Mac OS X 14.5; rv:131.0) Gecko/20100101 Firefox/131.0", "macOS · Firefox", "Firefox 131 on macOS", false},
		{"Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Mobile/15E148 Safari/604.1", "iOS · Safari", "Safari 17 on iOS", false},
		{"curl/8.0", "curl", "curl 8", false},
		{"", "Unknown device", "", false},
	}
	for _, tt := range tests {
		u := useragent.Parse(tt.ua)
		if got := userAgentSummary(u); got != tt.summary {
			t.Errorf("userAgentSummary(%q) = %q, want %q", tt.ua, got, tt.summary)
		}
		if got := browserName(tt.ua); got != tt.browser {
			t.Errorf("browserName(%q) = %q, want %q", tt.ua, got, tt.browser)
		}
		if u.Mobile != tt.phone {
			t.Errorf("Mobile(%q) = %v, want %v", tt.ua, u.Mobile, tt.phone)
		}
	}
}
