package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"gm-session-manager/auth"
	"gm-session-manager/db/sqlc"
)

// session is the device row behind a token.
func (e testEnv) session(t *testing.T, token string) sqlc.Session {
	t.Helper()
	s, err := e.mgr.Resolve(t.Context(), token)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func (e testEnv) device(t *testing.T, token string) string {
	return sessionKey(e.session(t, token))
}

// TestRequestLog: each device keeps its own recent requests, newest first,
// with what went wrong on failures; visitors and assets are not logged.
func TestRequestLog(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, auth.RolePlayer)
	other := e.token(t, auth.RolePlayer)
	charPath := fmt.Sprintf("/characters/%d", e.characterID)

	e.do(t, http.MethodGet, charPath, player, nil, false)
	e.do(t, http.MethodGet, "/notes/9999", player, nil, true)
	e.do(t, http.MethodGet, "/assets/js/app.js", player, nil, false)
	e.do(t, http.MethodGet, charPath, "", nil, false) // a visitor

	got := e.app.reqs.recent(e.device(t, player))
	if len(got) != 2 {
		t.Fatalf("got %d entries, want 2: %+v", len(got), got)
	}
	if got[0].Path != "/notes/9999" || got[0].Status != http.StatusNotFound || got[0].Error != "Not found." {
		t.Errorf("newest = %+v, want the 404 with its message", got[0])
	}
	if got[1].Path != charPath || got[1].Status != http.StatusOK || got[1].Error != "" {
		t.Errorf("oldest = %+v, want the character page", got[1])
	}
	if n := len(e.app.reqs.recent(e.device(t, other))); n != 0 {
		t.Errorf("another device has %d entries, want 0", n)
	}
	if len(e.app.reqs.devices) != 1 {
		t.Errorf("logged %d devices, want only the player's", len(e.app.reqs.devices))
	}

	for range reqLogSize + 5 {
		e.do(t, http.MethodGet, charPath, player, nil, false)
	}
	if n := len(e.app.reqs.recent(e.device(t, player))); n != reqLogSize {
		t.Errorf("kept %d entries, want %d", n, reqLogSize)
	}

	id := e.device(t, player)
	e.do(t, http.MethodPost, PathLogout, player, nil, false)
	if n := len(e.app.reqs.recent(id)); n != 0 {
		t.Errorf("logout kept %d entries", n)
	}
}

// TestRequestLogReusedID: SQLite gives a new session the ID of the newest
// deleted one (here, removed with its character); the new device must not
// inherit the old one's requests.
func TestRequestLogReusedID(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	old := e.token(t, auth.RolePlayer)
	oldSession := e.session(t, old)
	e.do(t, http.MethodGet, "/notes/9999", old, nil, true)
	if err := e.q().DeleteCharacterSessions(ctx, sql.NullInt64{Int64: e.characterID, Valid: true}); err != nil {
		t.Fatal(err)
	}

	fresh := e.token(t, auth.RolePlayer)
	if got := e.session(t, fresh).ID; got != oldSession.ID {
		t.Fatalf("new session got ID %d, want the reused %d", got, oldSession.ID)
	}
	if n := len(e.app.reqs.recent(e.device(t, fresh))); n != 0 {
		t.Errorf("new device inherited %d requests", n)
	}
}

func TestErrorText(t *testing.T) {
	tests := []struct {
		name  string
		setup func(c *gin.Context)
		ctype string
		body  string
		want  string
	}{
		{"server error wins", func(c *gin.Context) {
			c.Set(serverErrKey, "UNIQUE constraint failed")
			c.Set(failMsgKey, "Something went wrong")
		}, "text/html", "<div>", "UNIQUE constraint failed"},
		{"fail message", func(c *gin.Context) { c.Set(failMsgKey, "Not found.") }, "text/html", "<div>Not found.</div>", "Not found."},
		{"JSON body", func(*gin.Context) {}, "application/json", `{"error":"stale"}`, `{"error":"stale"}`},
		{"long body cut", func(*gin.Context) {}, "text/plain", strings.Repeat("x", 1000), strings.Repeat("x", maxErrBody)},
		{"HTML body dropped", func(*gin.Context) {}, "text/html", "<html>…</html>", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			w := &errBodyWriter{ResponseWriter: c.Writer}
			tt.setup(c)
			w.Header().Set("Content-Type", tt.ctype)
			w.WriteHeader(http.StatusConflict)
			_, _ = w.Write([]byte(tt.body))
			if got := errorText(c, w); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

// TestRequestSent: the log keeps the start of what a request sent, except
// credentials and note blocks; a GET sends nothing.
func TestRequestSent(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, auth.RolePlayer)
	cmd := fmt.Sprintf("/characters/%d/cmd", e.characterID)
	e.do(t, http.MethodPost, cmd, player, url.Values{"hp_op": {"damage"}, "hp_amount": {"five"}}, true)
	e.do(t, http.MethodPost, "/blocks/1/steps", player, url.Values{"text": {"my secret"}}, true)
	e.do(t, http.MethodGet, cmd+"?x=1", player, nil, true)

	got := e.app.reqs.recent(e.device(t, player))
	if len(got) != 3 {
		t.Fatalf("got %d entries", len(got))
	}
	if got[2].Sent != "hp_amount=five hp_op=damage" {
		t.Errorf("command sent %q", got[2].Sent)
	}
	if got[1].Sent != "" || got[0].Sent != "" {
		t.Errorf("kept %q for a note block and %q for a GET", got[1].Sent, got[0].Sent)
	}
}

func TestSentText(t *testing.T) {
	tests := []struct{ name, ctype, body, want string }{
		{"form", "application/x-www-form-urlencoded", "b=2&a=1+1", "a=1 1 b=2"},
		{"password", "application/x-www-form-urlencoded", "username=gm&password=hunter2", "password=… username=gm"},
		{"JSON", "application/json", "{\n  \"steps\": [1]\n}", `{ "steps": [1] }`},
		{"upload", "multipart/form-data", "--x\r\n...", "(file upload)"},
		{"long", "text/plain", strings.Repeat("é", maxSent+5), strings.Repeat("é", maxSent) + "…"},
		{"empty", "application/json", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sentText(tt.ctype, []byte(tt.body)); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

// TestAdminLoginNotKept: the login form never reaches the log.
func TestAdminLoginNotKept(t *testing.T) {
	e := newTestEnv(t)
	e.post(t, PathAdminLogin, loopbackAddr, adminCreds, "")
	for _, log := range e.app.reqs.devices {
		for _, en := range log {
			if strings.Contains(en.Sent, testAdminPass) || en.Sent != "" && en.Path == PathAdminLogin {
				t.Errorf("login kept: %+v", en)
			}
		}
	}
}
