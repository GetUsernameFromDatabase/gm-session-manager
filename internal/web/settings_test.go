package web

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/web/views"
)

// writeBook puts a downloaded book's book.json into the app's compendium.
func writeBook(t *testing.T, a *App, folder string, b compendium.Book) {
	t.Helper()
	dir := filepath.Join(a.compendiumDir(), compendium.FiveToolsPrefix+folder)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		t.Fatal(err)
	}
	body, _ := json.Marshal(b)
	if err := os.WriteFile(filepath.Join(dir, "book.json"), body, 0o600); err != nil {
		t.Fatal(err)
	}
}

// TestSettingsFile: no file is the defaults; a saved file reads back and
// leaves no temp file behind.
func TestSettingsFile(t *testing.T) {
	dir := t.TempDir()
	if st, err := loadSettings(dir); err != nil || st.FiveToolsAutoRedownload {
		t.Fatalf("defaults: %+v %v", st, err)
	}
	if err := saveSettings(dir, settings{FiveToolsAutoRedownload: true}); err != nil {
		t.Fatal(err)
	}
	if st, err := loadSettings(dir); err != nil || !st.FiveToolsAutoRedownload {
		t.Errorf("read back: %+v %v", st, err)
	}
	if files, _ := os.ReadDir(dir); len(files) != 1 {
		t.Errorf("files left: %v", files)
	}
}

// TestStartUpBrokenSettings: a broken settings.json still lists the stale books.
func TestStartUpBrokenSettings(t *testing.T) {
	e := newTestEnv(t)
	if err := os.WriteFile(filepath.Join(e.app.DataDir, settingsFile), []byte("{"), 0o600); err != nil {
		t.Fatal(err)
	}
	writeBook(t, e.app, "xge", compendium.Book{ID: "xge"})
	if err := e.app.StartUp(t.Context()); err != nil {
		t.Fatal(err)
	}
	if _, stale := e.app.state.get(); len(stale) != 1 {
		t.Errorf("stale: %v", stale)
	}
}

// TestStaleBooksBanner: the GM sees the stale books on every page, except
// while an import runs; players never do; the switch is stored.
func TestStaleBooksBanner(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)
	writeBook(t, e.app, "xge", compendium.Book{ID: "xge", Name: "Xanathar"})
	writeBook(t, e.app, "tce", compendium.Book{ID: "tce", Name: "Tasha", Version: compendium.DownloadVersion})
	if err := e.app.StartUp(t.Context()); err != nil {
		t.Fatal(err)
	}
	if e.app.fiveTools.get().Running {
		t.Error("switch off, yet the download started")
	}
	banner := func(cookie string) string {
		body := e.do(t, http.MethodGet, "/", cookie, nil, false).Body.String()
		if !strings.Contains(body, "downloading again") {
			return ""
		}
		return body
	}
	for _, tc := range []struct {
		name, cookie string
		want         bool
	}{{"gm", gm, true}, {"player", player, false}} {
		if body := banner(tc.cookie); (body != "") != tc.want {
			t.Errorf("%s: banner %v", tc.name, body != "")
		} else if tc.want && (!strings.Contains(body, "XGE") || strings.Contains(body, "TCE")) {
			t.Errorf("%s: wrong books in the banner", tc.name)
		}
	}
	e.app.fiveTools.start(1)
	if banner(gm) != "" {
		t.Error("banner during an import")
	}
	e.app.fiveTools.update(func(j *views.FiveToolsJob) { j.Running = false })
	if w := e.do(t, http.MethodPost, "/compendium/5etools/auto", gm, url.Values{"on": {"true"}}, true); w.Code != http.StatusNoContent {
		t.Fatalf("switch: %d", w.Code)
	}
	if st, err := loadSettings(e.app.DataDir); err != nil || !st.FiveToolsAutoRedownload {
		t.Errorf("switch not saved: %+v %v", st, err)
	}
}

// TestStartUpRedownloads: with the switch on, the start downloads the stale
// books again (a bad id fails before any network call).
func TestStartUpRedownloads(t *testing.T) {
	e := newTestEnv(t)
	if err := saveSettings(e.app.DataDir, settings{FiveToolsAutoRedownload: true}); err != nil {
		t.Fatal(err)
	}
	writeBook(t, e.app, "bad", compendium.Book{ID: "bad id"})
	if err := e.app.StartUp(t.Context()); err != nil {
		t.Fatal(err)
	}
	for deadline := time.Now().Add(5 * time.Second); e.app.fiveTools.get().Running && time.Now().Before(deadline); {
		time.Sleep(10 * time.Millisecond)
	}
	if j := e.app.fiveTools.get(); !strings.Contains(j.Msg, "bad 5etools book id") {
		t.Errorf("job: %+v", j)
	}
}
