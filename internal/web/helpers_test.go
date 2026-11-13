package web

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
)

func TestHandlerErrors(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		status int // a plain request's; htmx gets 422 for every failure
		body   string
	}{
		{"nil", nil, http.StatusOK, ""},
		{"user", userError{errors.New("pick one")}, http.StatusUnprocessableEntity, "pick one"},
		{"field", fieldError{userError{errors.New("too long")}, "name"}, http.StatusUnprocessableEntity, "too long"},
		{"no rows", fmt.Errorf("load: %w", sql.ErrNoRows), http.StatusNotFound, "Not found."},
		{"other", errors.New("disk full"), http.StatusInternalServerError, "Something went wrong"},
	}
	a := &App{}
	for _, tt := range tests {
		for _, htmx := range []bool{false, true} {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/x", nil)
			if htmx {
				c.Request.Header.Set("HX-Request", strTrue)
			}
			a.h(func(*gin.Context) error { return tt.err })(c)
			want := tt.status
			if htmx && tt.err != nil {
				want = http.StatusUnprocessableEntity
			}
			if w.Code != want || !strings.Contains(w.Body.String(), tt.body) {
				t.Errorf("%s (htmx %v): %d %q, want %d containing %q", tt.name, htmx, w.Code, w.Body.String(), want, tt.body)
			}
			if tt.err == nil && w.Body.Len() != 0 {
				t.Errorf("%s (htmx %v): wrote %q", tt.name, htmx, w.Body.String())
			}
		}
	}
}

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
