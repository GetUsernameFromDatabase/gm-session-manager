package auth

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"

	"gm-session-manager/db"
	"gm-session-manager/db/sqlc"
)

const (
	testAdminUser = "gm"
	testAdminPass = "correct horse"
	lanAddr       = "192.168.1.20:51000"
	loopbackAddr  = "127.0.0.1:51000"
)

type fixture struct {
	conn        *sql.DB
	q           *db.Store
	mgr         *Manager
	characterID int64
}

func newFixture(t *testing.T) fixture {
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
	return fixture{conn: conn, q: q, mgr: NewManager(q, testAdminUser, testAdminPass), characterID: char.ID}
}

// tokens holds one of each kind of credential a device might present.
type tokens struct {
	player, admin, revoked, unknown string
}

func (f fixture) issueTokens(t *testing.T) tokens {
	t.Helper()
	ctx := t.Context()
	var tk tokens
	var err error
	if tk.player, _, err = f.mgr.IssuePlayer(ctx, f.characterID); err != nil {
		t.Fatal(err)
	}
	if tk.admin, _, err = f.mgr.LoginAdmin(ctx, testAdminUser, testAdminPass); err != nil {
		t.Fatal(err)
	}
	revokedTok, revoked, err := f.mgr.IssuePlayer(ctx, f.characterID)
	if err != nil {
		t.Fatal(err)
	}
	if err := f.mgr.Revoke(ctx, revoked.ID); err != nil {
		t.Fatal(err)
	}
	tk.revoked = revokedTok
	if tk.unknown, _, err = newToken(); err != nil {
		t.Fatal(err)
	}
	return tk
}

func serve(t *testing.T, h http.Handler, remoteAddr, cookie string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/", nil)
	req.RemoteAddr = remoteAddr
	if cookie != "" {
		req.AddCookie(&http.Cookie{Name: CookieName, Value: cookie})
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	return w
}

func isCleared(w *httptest.ResponseRecorder) bool {
	for _, c := range w.Result().Cookies() {
		if c.Name == CookieName && c.MaxAge < 0 {
			return true
		}
	}
	return false
}

func TestTokens(t *testing.T) {
	tok1, hash1, err := newToken()
	if err != nil {
		t.Fatal(err)
	}
	tok2, _, err := newToken()
	if err != nil {
		t.Fatal(err)
	}
	if tok1 == tok2 {
		t.Error("tokens repeat")
	}
	if len(hash1) != 32 {
		t.Errorf("hash length %d, want 32", len(hash1))
	}

	tests := []struct {
		name   string
		token  string
		wantOK bool
	}{
		{"issued token", tok1, true},
		{"empty", "", false},
		{"not base64", "!!!not-a-token!!!", false},
		{"too short", tokenEncoding.EncodeToString(make([]byte, 16)), false},
		{"too long", tokenEncoding.EncodeToString(make([]byte, 64)), false},
		{"padded encoding", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := hashToken(tt.token)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
			}
			if ok && tt.token == tok1 && !bytes.Equal(got, hash1) {
				t.Error("hash mismatch for issued token")
			}
		})
	}
}

func TestIsLoopback(t *testing.T) {
	tests := []struct {
		name       string
		remoteAddr string
		headers    map[string]string
		want       bool
	}{
		{"ipv4 loopback", "127.0.0.1:5000", nil, true},
		{"ipv4 loopback range", "127.0.0.2:5000", nil, true},
		{"ipv6 loopback", "[::1]:5000", nil, true},
		{"ipv4-mapped loopback", "[::ffff:127.0.0.1]:5000", nil, true},
		{"lan peer", lanAddr, nil, false},
		{"unspecified", "0.0.0.0:5000", nil, false},
		{"lan peer spoofing X-Forwarded-For", lanAddr, map[string]string{"X-Forwarded-For": "127.0.0.1"}, false},
		{"lan peer spoofing X-Real-IP", lanAddr, map[string]string{"X-Real-Ip": "127.0.0.1"}, false},
		{"lan peer spoofing Forwarded", lanAddr, map[string]string{"Forwarded": "for=127.0.0.1"}, false},
		{"missing port", "127.0.0.1", nil, false},
		{"hostname", "localhost:5000", nil, false},
		{"empty", "", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/", nil)
			r.RemoteAddr = tt.remoteAddr
			for k, v := range tt.headers {
				r.Header.Set(k, v)
			}
			if got := IsLoopback(r); got != tt.want {
				t.Errorf("IsLoopback = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialsMatch(t *testing.T) {
	tests := []struct {
		name, user, pass string
		want             bool
	}{
		{"correct", testAdminUser, testAdminPass, true},
		{"wrong password", testAdminUser, "wrong", false},
		{"wrong user", "player", testAdminPass, false},
		{"both wrong", "x", "y", false},
		{"empty", "", "", false},
		{"password prefix", testAdminUser, testAdminPass[:3], false},
		{"password with suffix", testAdminUser, testAdminPass + " ", false},
		{"user case differs", "GM", testAdminPass, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := credentialsMatch(testAdminUser, testAdminPass, tt.user, tt.pass); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddlewareIdentity(t *testing.T) {
	f := newFixture(t)
	tk := f.issueTokens(t)

	tests := []struct {
		name          string
		remoteAddr    string
		cookie        string
		wantAdmin     bool
		wantSession   bool
		wantCharacter bool
		wantCleared   bool
	}{
		{name: "lan visitor", remoteAddr: lanAddr},
		{name: "loopback visitor", remoteAddr: loopbackAddr, wantAdmin: true},
		{name: "lan player", remoteAddr: lanAddr, cookie: tk.player, wantSession: true, wantCharacter: true},
		{name: "loopback player acts as player only", remoteAddr: loopbackAddr, cookie: tk.player, wantSession: true, wantCharacter: true},
		{name: "lan admin login", remoteAddr: lanAddr, cookie: tk.admin, wantAdmin: true, wantSession: true},
		{name: "loopback admin login", remoteAddr: loopbackAddr, cookie: tk.admin, wantAdmin: true, wantSession: true},
		{name: "lan revoked (kick/logout)", remoteAddr: lanAddr, cookie: tk.revoked, wantCleared: true},
		{name: "loopback revoked is admin again", remoteAddr: loopbackAddr, cookie: tk.revoked, wantAdmin: true, wantCleared: true},
		{name: "lan never issued", remoteAddr: lanAddr, cookie: tk.unknown, wantCleared: true},
		{name: "lan garbage cookie", remoteAddr: lanAddr, cookie: "garbage", wantCleared: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Identity
			r := gin.New()
			r.Use(f.mgr.Middleware())
			r.GET("/", func(c *gin.Context) { got = FromContext(c) })

			w := serve(t, r, tt.remoteAddr, tt.cookie)
			if w.Code != http.StatusOK {
				t.Fatalf("status %d", w.Code)
			}
			if got.IsAdmin() != tt.wantAdmin {
				t.Errorf("IsAdmin = %v, want %v", got.IsAdmin(), tt.wantAdmin)
			}
			if (got.Session != nil) != tt.wantSession {
				t.Errorf("has session = %v, want %v", got.Session != nil, tt.wantSession)
			}
			charID, hasChar := got.CharacterID()
			if hasChar != tt.wantCharacter || (hasChar && charID != f.characterID) {
				t.Errorf("CharacterID = (%d, %v), want character %v", charID, hasChar, tt.wantCharacter)
			}
			if cleared := isCleared(w); cleared != tt.wantCleared {
				t.Errorf("cookie cleared = %v, want %v", cleared, tt.wantCleared)
			}
		})
	}
}

func TestRequireAdmin(t *testing.T) {
	f := newFixture(t)
	tk := f.issueTokens(t)

	tests := []struct {
		name       string
		remoteAddr string
		cookie     string
		want       int
	}{
		{"loopback", loopbackAddr, "", http.StatusOK},
		{"lan admin session", lanAddr, tk.admin, http.StatusOK},
		{"lan player", lanAddr, tk.player, http.StatusForbidden},
		{"lan anonymous", lanAddr, "", http.StatusForbidden},
		{"loopback player", loopbackAddr, tk.player, http.StatusForbidden},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.Use(f.mgr.Middleware())
			r.GET("/", RequireAdmin(), func(c *gin.Context) { c.Status(http.StatusOK) })
			if w := serve(t, r, tt.remoteAddr, tt.cookie); w.Code != tt.want {
				t.Errorf("status %d, want %d", w.Code, tt.want)
			}
		})
	}
}

func TestOnlyHashIsStored(t *testing.T) {
	f := newFixture(t)
	tok, sess, err := f.mgr.IssuePlayer(t.Context(), f.characterID)
	if err != nil {
		t.Fatal(err)
	}
	raw, _ := tokenEncoding.DecodeString(tok)
	if bytes.Equal(sess.TokenHash, raw) || bytes.Equal(sess.TokenHash, []byte(tok)) {
		t.Error("raw token stored in database")
	}
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}
