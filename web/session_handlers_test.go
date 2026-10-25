package web

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"gm-session-manager/auth"
	"gm-session-manager/db"
	"gm-session-manager/db/sqlc"
)

const (
	testAdminUser = "gm"
	testAdminPass = "correct horse"
	lanAddr       = "192.168.1.20:51000"
	loopbackAddr  = "127.0.0.1:51000"
)

var adminCreds = url.Values{"username": {testAdminUser}, "password": {testAdminPass}}

type testEnv struct {
	router      *gin.Engine
	mgr         *auth.Manager
	app         *App
	campaignID  int64
	characterID int64
}

func newTestEnv(t *testing.T) testEnv {
	t.Helper()
	ctx := t.Context()
	conn, err := db.Open(ctx, filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })

	q := db.NewStore(conn)
	camp, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Test", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	char, err := q.CreateCharacter(ctx, sqlc.CreateCharacterParams{CampaignID: camp.ID, Name: "Alice", System: "dnd5e"})
	if err != nil {
		t.Fatal(err)
	}
	mgr := auth.NewManager(q, testAdminUser, testAdminPass)
	app := New(Deps{Auth: mgr, Store: q, DataDir: t.TempDir(), Port: 8080})
	t.Cleanup(app.Close)
	return testEnv{router: app.Router(), mgr: mgr, app: app, campaignID: camp.ID, characterID: char.ID}
}

// token issues a session of the given role, or returns "" for none.
func (e testEnv) token(t *testing.T, role auth.Role) string {
	t.Helper()
	var tok string
	var err error
	switch role {
	case auth.RolePlayer:
		tok, _, err = e.mgr.IssuePlayer(t.Context(), e.characterID)
	case auth.RoleAdmin:
		tok, _, err = e.mgr.LoginAdmin(t.Context(), testAdminUser, testAdminPass)
	}
	if err != nil {
		t.Fatal(err)
	}
	return tok
}

func (e testEnv) post(t *testing.T, path, remoteAddr string, form url.Values, cookie string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, path, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.RemoteAddr = remoteAddr
	if cookie != "" {
		req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: cookie})
	}
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	return w
}

func (e testEnv) alive(t *testing.T, token string) bool {
	t.Helper()
	_, err := e.mgr.Resolve(t.Context(), token)
	return err == nil
}

func sessionCookie(w *httptest.ResponseRecorder) *http.Cookie {
	for _, c := range w.Result().Cookies() {
		if c.Name == auth.CookieName {
			return c
		}
	}
	return nil
}

func TestAdminLogin(t *testing.T) {
	tests := []struct {
		name       string
		existing   auth.Role
		form       url.Values
		wantStatus int
	}{
		{"correct", "", adminCreds, http.StatusSeeOther},
		{"wrong password", "", url.Values{"username": {testAdminUser}, "password": {"nope"}}, http.StatusUnauthorized},
		{"wrong user", "", url.Values{"username": {"bob"}, "password": {testAdminPass}}, http.StatusUnauthorized},
		{"missing password", "", url.Values{"username": {testAdminUser}}, http.StatusBadRequest},
		{"empty", "", url.Values{}, http.StatusBadRequest},
		{"already a player: log out first", auth.RolePlayer, adminCreds, http.StatusConflict},
		{"already an admin: log out first", auth.RoleAdmin, adminCreds, http.StatusConflict},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := newTestEnv(t)
			existing := e.token(t, tt.existing)

			w := e.post(t, PathAdminLogin, lanAddr, tt.form, existing)
			if w.Code != tt.wantStatus {
				t.Fatalf("status %d, want %d", w.Code, tt.wantStatus)
			}
			if existing != "" && !e.alive(t, existing) {
				t.Error("refused login must not touch the existing session")
			}

			c := sessionCookie(w)
			if wantCookie := tt.wantStatus == http.StatusSeeOther; (c != nil) != wantCookie {
				t.Fatalf("cookie set = %v, want %v", c != nil, wantCookie)
			}
			if c == nil {
				return
			}
			if !c.HttpOnly || c.SameSite != http.SameSiteLaxMode || c.MaxAge <= 0 || c.Path != "/" {
				t.Errorf("cookie attributes: %+v", c)
			}
			if !e.alive(t, c.Value) {
				t.Error("new admin session not usable")
			}
		})
	}
}

func TestLogout(t *testing.T) {
	tests := []struct {
		name       string
		remoteAddr string
		role       auth.Role
	}{
		{"lan player", lanAddr, auth.RolePlayer},
		{"lan admin", lanAddr, auth.RoleAdmin},
		{"loopback player switching back to admin", loopbackAddr, auth.RolePlayer},
		{"no session", lanAddr, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := newTestEnv(t)
			tok := e.token(t, tt.role)

			w := e.post(t, PathLogout, tt.remoteAddr, nil, tok)
			if w.Code != http.StatusSeeOther {
				t.Fatalf("status %d", w.Code)
			}
			if c := sessionCookie(w); c == nil || c.MaxAge >= 0 {
				t.Error("cookie not cleared")
			}
			if tok != "" && e.alive(t, tok) {
				t.Error("session row survived logout")
			}
		})
	}
}

// TestCrossSiteRefused: another site in the host's browser must not act as
// the cookie-less localhost admin.
func TestCrossSiteRefused(t *testing.T) {
	e := newTestEnv(t)
	for site, want := range map[string]int{"cross-site": http.StatusForbidden, "same-origin": http.StatusSeeOther, "": http.StatusSeeOther} {
		req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, fmt.Sprintf("/campaigns/%d/open-join", e.campaignID), nil)
		req.RemoteAddr = loopbackAddr
		if site != "" {
			req.Header.Set("Sec-Fetch-Site", site)
		}
		w := httptest.NewRecorder()
		e.router.ServeHTTP(w, req)
		if w.Code != want {
			t.Errorf("Sec-Fetch-Site %q: status %d, want %d", site, w.Code, want)
		}
	}
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}

func TestCampaignJoinLinks(t *testing.T) {
	e := newTestEnv(t)
	e.app.LANIPs = []net.IP{net.ParseIP("192.168.1.10")}
	gm := e.token(t, auth.RoleAdmin)

	cases := []struct {
		path, url, qr string
	}{
		{"/", "https://192.168.1.10:8080/join<", `src="/join.png"`},
		{fmt.Sprintf("/campaigns/%d", e.campaignID), fmt.Sprintf("https://192.168.1.10:8080/join/%d<", e.campaignID), fmt.Sprintf(`src="/join.png?campaign=%d"`, e.campaignID)},
		{fmt.Sprintf("/campaigns/%d/images", e.campaignID), "Show join QR", fmt.Sprintf(`data-src="/join.png?campaign=%d"`, e.campaignID)},
	}
	for _, tc := range cases {
		body := e.do(t, http.MethodGet, tc.path, gm, nil, false).Body.String()
		if !strings.Contains(body, tc.url) || !strings.Contains(body, tc.qr) {
			t.Errorf("%s: missing %q or %q", tc.path, tc.url, tc.qr)
		}
	}

	for q, want := range map[string]int{"": http.StatusOK, "?campaign=1": http.StatusOK, "?campaign=x": http.StatusBadRequest, "?campaign=0": http.StatusBadRequest} {
		if w := e.do(t, http.MethodGet, "/join.png"+q, gm, nil, false); w.Code != want {
			t.Errorf("/join.png%s: %d, want %d", q, w.Code, want)
		}
	}
}

func TestShortHost(t *testing.T) {
	for in, want := range map[string]string{
		"miras-pc.lan.":       "miras-pc",
		"miras-iphone.local.": "miras-iphone",
		"pixel-8.home.arpa.":  "pixel-8",
		"bare":                "bare",
		"":                    "",
	} {
		if got := shortHost(in); got != want {
			t.Errorf("shortHost(%q) = %q, want %q", in, got, want)
		}
	}
}
