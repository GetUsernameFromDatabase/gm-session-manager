package web

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/coder/websocket"

	"gm-session-manager/auth"
	"gm-session-manager/hub"
	"gm-session-manager/web/views"
)

// dial connects a websocket as a device and says hello at the current
// hub position, watching the given live views.
func dial(t *testing.T, e testEnv, srv *httptest.Server, query, cookie string, watch ...string) *websocket.Conn {
	t.Helper()
	h := http.Header{}
	if cookie != "" {
		h.Set("Cookie", auth.CookieName+"="+cookie)
	}
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()
	conn, res, err := websocket.Dial(ctx, "ws"+strings.TrimPrefix(srv.URL, "http")+"/ws"+query, &websocket.DialOptions{HTTPHeader: h, CompressionMode: websocket.CompressionContextTakeover})
	if res != nil && res.Body != nil {
		_ = res.Body.Close()
	}
	if err != nil {
		t.Fatal(err)
	}
	// Same cap as the server's.
	conn.SetReadLimit(hub.ReadLimit)
	t.Cleanup(func() { _ = conn.CloseNow() })
	epoch, seq := e.app.Hub.Position()
	hello, _ := json.Marshal(map[string]any{"type": "hello", "epoch": epoch, "seq": seq, "watch": watch})
	if err := conn.Write(ctx, websocket.MessageText, hello); err != nil {
		t.Fatal(err)
	}
	return conn
}

// wsMessage is what the hub sends.
type wsMessage struct {
	Type  string `json:"type"`
	Seq   uint64 `json:"seq"`
	Views []struct {
		V       string         `json:"v"`
		Set     map[string]any `json:"set"`
		Del     []string       `json:"del"`
		Refetch bool           `json:"refetch"`
	} `json:"views"`
	raw string
}

// readUntil reads messages until match accepts one.
func readUntil(t *testing.T, conn *websocket.Conn, what string, match func(wsMessage) bool) wsMessage {
	t.Helper()
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()
	for {
		_, data, err := conn.Read(ctx)
		if err != nil {
			t.Fatalf("waiting for %s: %v", what, err)
		}
		var m wsMessage
		if err := json.Unmarshal(data, &m); err != nil {
			t.Fatalf("not JSON: %s", data)
		}
		m.raw = string(data)
		if match(m) {
			return m
		}
	}
}

// patchOf returns the message's entry for view, if any.
func patchOf(m wsMessage, view string) (set map[string]any, ok bool) {
	for _, v := range m.Views {
		if v.V == view {
			return v.Set, true
		}
	}
	return nil, false
}

func TestLiveUpdatesAndKick(t *testing.T) {
	e := newTestEnv(t)
	srv := httptest.NewServer(e.router)
	t.Cleanup(srv.Close)
	player := e.token(t, auth.RolePlayer)
	sess, err := e.mgr.Resolve(t.Context(), player)
	if err != nil {
		t.Fatal(err)
	}
	roster := views.LiveKey(views.LiveRoster, e.campaignID)
	sheet := views.LiveKey(views.LiveCharacter, e.characterID)
	encounters := views.LiveKey(views.LiveEncounters, e.campaignID, 0)
	hello := func(m wsMessage) bool { return m.Type == "patch" && len(m.Views) == 0 }

	// httptest connects from 127.0.0.1: no cookie is the localhost GM.
	gm := dial(t, e, srv, fmt.Sprintf("?campaign=%d", e.campaignID), "", roster, sheet)
	readUntil(t, gm, "the GM's hello", hello)
	// A GM tab on the combat page shows no sheet, so it never gets one.
	combat := dial(t, e, srv, fmt.Sprintf("?campaign=%d", e.campaignID), "", encounters)
	readUntil(t, combat, "the combat tab's hello", hello)
	phone := dial(t, e, srv, "", player, sheet)
	readUntil(t, phone, "the phone's hello", hello)

	// The GM's device list sees the phone come online as one value.
	readUntil(t, gm, "the phone online", func(m wsMessage) bool {
		set, _ := patchOf(m, roster)
		return set[fmt.Sprintf("dev.%d.on", sess.ID)] == true
	})

	cmd := func(c string) {
		t.Helper()
		form := url.Values{"cmd": {c}}
		req, _ := http.NewRequestWithContext(t.Context(), http.MethodPost, srv.URL+fmt.Sprintf("/characters/%d/cmd", e.characterID), strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("HX-Request", "true")
		req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: player})
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		_ = res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Fatalf("%s: %d", c, res.StatusCode)
		}
	}

	// Max HP is not bound: the sheet changed shape, so the hub sends only a
	// refetch order.
	cmd("maxhp =12")
	for _, conn := range []*websocket.Conn{gm, phone} {
		readUntil(t, conn, "the refetch order", func(m wsMessage) bool {
			for _, v := range m.Views {
				if v.V == sheet {
					if len(v.Set) > 0 {
						t.Errorf("refetch carried values: %s", m.raw)
					}
					return v.Refetch
				}
			}
			return false
		})
	}

	// HP is bound: only that one value travels.
	cmd("hp =7")
	for _, conn := range []*websocket.Conn{gm, phone} {
		m := readUntil(t, conn, "the hp patch", func(m wsMessage) bool { _, ok := patchOf(m, sheet); return ok })
		if set, _ := patchOf(m, sheet); len(set) != 1 || set["hp"] != "7" {
			t.Errorf("hp patch %s, want only hp", m.raw)
		}
		if len(m.raw) > 200 {
			t.Errorf("hp patch is %d bytes: %s", len(m.raw), m.raw)
		}
	}

	// The combat tab heard none of it: its first message is its own view.
	if w := e.do(t, http.MethodPost, fmt.Sprintf("/campaigns/%d/encounters", e.campaignID), e.token(t, auth.RoleAdmin), url.Values{"name": {"Ambush"}}, true); w.Code >= 400 {
		t.Fatalf("create encounter: %d", w.Code)
	}
	m := readUntil(t, combat, "any message", func(wsMessage) bool { return true })
	if _, ok := patchOf(m, encounters); !ok || len(m.Views) != 1 {
		t.Errorf("combat tab got %s, want only its encounter list", m.raw)
	}

	e.app.Hub.CloseSession(sess.ID)
	readUntil(t, phone, "the reload", func(m wsMessage) bool { return m.Type == "reload" })
}

var moduleImport = regexp.MustCompile(`(?:from|import)\s*"(/assets/[^"]+|\./[^"]+)"`)

// TestModuleImportsResolve walks the editor's ES module graph through the
// embedded files, so a missing vendored file fails here, not on a phone.
func TestModuleImportsResolve(t *testing.T) {
	sub, _ := fs.Sub(assetsFS, "assets")
	seen := map[string]bool{}
	var visit func(p string)
	visit = func(p string) {
		if seen[p] {
			return
		}
		seen[p] = true
		body, err := fs.ReadFile(sub, p)
		if err != nil {
			t.Errorf("missing module %s", p)
			return
		}
		for _, m := range moduleImport.FindAllStringSubmatch(string(body), -1) {
			spec := m[1]
			if strings.HasPrefix(spec, "./") {
				visit(path.Join(path.Dir(p), spec))
			} else {
				visit(strings.TrimPrefix(spec, "/assets/"))
			}
		}
	}
	visit("js/notes.js")
	if len(seen) < 10 {
		t.Errorf("only %d modules reached; the import pattern may be wrong", len(seen))
	}
	for _, f := range []string{"vendor/htmx.min.js", "vendor/htmx-ws.js", "js/app.js", "css/output.css"} {
		if _, err := fs.Stat(sub, f); err != nil {
			t.Errorf("missing asset %s", f)
		}
	}
}
