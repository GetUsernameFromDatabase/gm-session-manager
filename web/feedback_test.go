package web

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"gm-session-manager/auth"
	"gm-session-manager/web/views"
)

// sendFeedback posts a report as a device, with files as name → content.
func (e testEnv) sendFeedback(t *testing.T, cookie string, fields url.Values, files map[string]string) *httptest.ResponseRecorder {
	t.Helper()
	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	for k, vs := range fields {
		for _, v := range vs {
			_ = mw.WriteField(k, v)
		}
	}
	for name, content := range files {
		fw, _ := mw.CreateFormFile("screenshots", name)
		_, _ = fw.Write([]byte(content))
	}
	_ = mw.Close()
	req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, views.PathFeedback, &body)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("HX-Request", "true")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 14) Gecko/131.0 Firefox/131.0")
	req.RemoteAddr = lanAddr
	if cookie != "" {
		req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: cookie})
	}
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	return w
}

// reports returns the report folders and each one's Markdown file.
func (e testEnv) reports(t *testing.T) map[string]string {
	t.Helper()
	entries, err := os.ReadDir(e.app.feedbackDir())
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		t.Fatal(err)
	}
	out := map[string]string{}
	for _, en := range entries {
		mds, _ := filepath.Glob(filepath.Join(e.app.feedbackDir(), en.Name(), "*.md"))
		if len(mds) != 1 {
			t.Fatalf("folder %s holds %d Markdown files, want 1", en.Name(), len(mds))
		}
		b, _ := os.ReadFile(mds[0])
		out[en.Name()+"/"+filepath.Base(mds[0])] = string(b)
	}
	return out
}

func bugForm(title string) url.Values {
	return url.Values{
		"kind": {"bug"}, "title": {title}, "body": {"The sheet went blank."}, "steps": {"Open sheet\nTap HP"},
		"expected": {"HP changes"}, "player": {"Sam"}, "page": {"/characters/1"}, "viewport": {"390×844"}, "dark": {"true"},
	}
}

func TestFeedbackAccess(t *testing.T) {
	e := newTestEnv(t)
	tests := []struct {
		name   string
		cookie string
		want   int
	}{
		{"player", e.token(t, auth.RolePlayer), http.StatusOK},
		{"GM", e.token(t, auth.RoleAdmin), http.StatusOK},
		{"visitor", "", http.StatusUnprocessableEntity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := len(e.reports(t))
			w := e.sendFeedback(t, tt.cookie, bugForm(tt.name), nil)
			if w.Code != tt.want {
				t.Fatalf("status %d, want %d: %s", w.Code, tt.want, w.Body)
			}
			if added := len(e.reports(t)) - before; added != map[bool]int{true: 1, false: 0}[tt.want == http.StatusOK] {
				t.Errorf("added %d reports", added)
			}
		})
	}
}

func TestFeedbackValidation(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, auth.RolePlayer)
	tests := []struct {
		name  string
		edit  func(url.Values)
		files int
	}{
		{"bad kind", func(f url.Values) { f.Set("kind", "rant") }, 0},
		{"no title", func(f url.Values) { f.Del("title") }, 0},
		{"no description", func(f url.Values) { f.Del("body") }, 0},
		{"long title", func(f url.Values) { f.Set("title", strings.Repeat("x", 121)) }, 0},
		{"long steps", func(f url.Values) { f.Set("steps", strings.Repeat("x", 4001)) }, 0},
		{"too many screenshots", func(url.Values) {}, views.MaxScreenshots + 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := bugForm("x")
			tt.edit(f)
			files := map[string]string{}
			for i := range tt.files {
				files[string(rune('a'+i))+".png"] = pngHead
			}
			if w := e.sendFeedback(t, player, f, files); w.Code != http.StatusUnprocessableEntity {
				t.Errorf("status %d, want 422", w.Code)
			}
		})
	}
	if n := len(e.reports(t)); n != 0 {
		t.Errorf("saved %d reports, want none", n)
	}
}

// TestBugReport: a bug report holds what the reporter wrote, their device,
// the requests that failed with their errors, and the pictures sent.
func TestBugReport(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, auth.RolePlayer)
	e.do(t, http.MethodGet, "/notes/9999", player, nil, true)

	f := bugForm("../../Dice | tray")
	f.Set("console", `[
		{"at":"2026-10-26T15:29:50.000Z","level":"log","msg":"sheet loaded","page":"/characters/1"},
		{"at":"2026-10-26T15:29:57.000Z","level":"uncaught","msg":"TypeError: x is null\n    at hp (app.js:12:3)","page":"/characters/1"},
		{"at":"2026-10-26T15:29:58.000Z","level":"warn","msg":"a | b","page":"/characters/1"}
	]`)
	w := e.sendFeedback(t, player, f, map[string]string{"shot.png": pngHead, "notes.txt": "not a picture"})
	if w.Code != http.StatusOK {
		t.Fatalf("status %d: %s", w.Code, w.Body)
	}
	reports := e.reports(t)
	if len(reports) != 1 {
		t.Fatalf("got %d reports", len(reports))
	}
	for path, md := range reports {
		dir, file := filepath.Split(path)
		if !strings.Contains(dir, "_bug_dice-tray_") {
			t.Errorf("folder %q lacks the kind and title slug", dir)
		}
		if file != "bug-dice-tray.md" {
			t.Errorf("report file %q is not named after its title", file)
		}
		for _, want := range []string{
			"# Bug: ../../Dice | tray\n",
			"| From | Alice (Sam) |",
			"| Page | /characters/1 |",
			"| Device | Firefox 131 on Android · 390×844 · dark |",
			"## Steps to reproduce\n\nOpen sheet\nTap HP\n",
			"| GET /notes/9999 |  | 404 | Not found. |",
			"## Recent requests (last 30)",
			// Errors newest first, a stack's first line in the table and all of it below.
			"## Browser errors\n\n| Time (UTC) | Page | Message |\n|---|---|---|\n| 15:29:57 | /characters/1 | TypeError: x is null |\n",
			"```\nTypeError: x is null\n    at hp (app.js:12:3)\n```",
			"| 15:29:58 | warn | /characters/1 | a \\| b |\n| 15:29:57 | uncaught |",
			"| 15:29:50 | log | /characters/1 | sheet loaded |",
			"![screenshot-1.png](screenshot-1.png)",
			"Left out, not a picture: notes.txt (text/plain)",
		} {
			if !strings.Contains(md, want) {
				t.Errorf("report lacks %q:\n%s", want, md)
			}
		}
		if strings.Contains(md, "screenshot-2") {
			t.Error("a file that is not a picture was kept")
		}
		files, _ := os.ReadDir(filepath.Join(e.app.feedbackDir(), dir))
		if len(files) != 2 {
			t.Errorf("folder holds %d files, want the report and one screenshot", len(files))
		}
	}
}

func TestIdeaHasNoRequests(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, auth.RoleAdmin)
	f := bugForm("More dice colors")
	f.Set("kind", "idea")
	w := e.sendFeedback(t, gm, f, nil)
	if w.Code != http.StatusOK {
		t.Fatalf("status %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Idea sent.") {
		t.Errorf("toast does not name an idea: %s", w.Body)
	}
	for _, md := range e.reports(t) {
		if !strings.Contains(md, "| From | GM (Sam) |") || !strings.Contains(md, "## Idea\n\nThe sheet went blank.") {
			t.Errorf("unexpected report:\n%s", md)
		}
		if strings.Contains(md, "requests") || strings.Contains(md, "Steps") {
			t.Errorf("an idea carries bug details:\n%s", md)
		}
	}
}

func TestFolderName(t *testing.T) {
	tests := []struct{ title, slug string }{
		{"Dice tray overlaps sheet", "dice-tray-overlaps-sheet"},
		{"../../etc/passwd", "etc-passwd"},
		{`C:\Windows | x`, "c-windows-x"},
		{"???", "report"},
		{strings.Repeat("ab-", 30), "ab-ab-ab-ab-ab-ab-ab-ab-ab-ab-ab-ab-ab-a"},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			name := folderName(time.Time{}, "bug", tt.title)
			if !strings.HasPrefix(name, "0001-01-01_000000_bug_"+tt.slug+"_") || len(name) != len("0001-01-01_000000_bug_"+tt.slug+"_abcd") {
				t.Errorf("got %q, want slug %q", name, tt.slug)
			}
			if filepath.Base(name) != name {
				t.Errorf("%q is not a single path element", name)
			}
		})
	}
}

func TestMarkdownCell(t *testing.T) {
	if got := mdCell("a | b\n\nc"); got != `a \| b c` {
		t.Errorf("got %q", got)
	}
}

func TestBrowserName(t *testing.T) {
	tests := []struct{ ua, want string }{
		{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0", "Edge 130 on Windows"},
		{"Mozilla/5.0 (Linux; Android 14; Pixel 8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.6723.58 Mobile Safari/537.36", "Chrome 130 on Android"},
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Mobile/15E148 Safari/604.1", "Safari 17 on iPhone"},
		{"Mozilla/5.0 (Macintosh; Intel Mac OS X 14.5; rv:131.0) Gecko/20100101 Firefox/131.0", "Firefox 131 on macOS"},
		{"curl/8.0", "curl/8.0"},
	}
	for _, tt := range tests {
		if got := browserName(tt.ua); got != tt.want {
			t.Errorf("browserName(%q) = %q, want %q", tt.ua, got, tt.want)
		}
	}
}

// TestBadConsole: a console the server cannot use is left out, and the
// report is still saved.
func TestBadConsole(t *testing.T) {
	tooMany := "[" + strings.Repeat(`{"at":"x","level":"log","msg":"m","page":"/"},`, consoleSize) + `{"at":"x","level":"log","msg":"m","page":"/"}]`
	tests := []struct{ name, console string }{
		{"not JSON", "{oops"},
		{"unknown level", `[{"at":"x","level":"debug","msg":"m","page":"/"}]`},
		{"too many", tooMany},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := newTestEnv(t)
			f := bugForm("x")
			f.Set("console", tt.console)
			if w := e.sendFeedback(t, e.token(t, auth.RolePlayer), f, nil); w.Code != http.StatusOK {
				t.Fatalf("status %d", w.Code)
			}
			for _, md := range e.reports(t) {
				if !strings.Contains(md, "unreadable console") || strings.Contains(md, "| log |") {
					t.Errorf("console not left out:\n%s", md)
				}
			}
		})
	}
}

// TestFeedbackIndex: USER_FEEDBACK_NOT_READ lists every report folder,
// newest first, linking its file; it goes away once they are deleted.
func TestFeedbackIndex(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, auth.RolePlayer)
	index := filepath.Join(e.app.DataDir, FeedbackIndexName)
	read := func() string {
		b, err := os.ReadFile(index)
		if err != nil {
			t.Fatal(err)
		}
		return string(b)
	}

	// An older report, placed by hand so the two never share a second.
	old := filepath.Join(e.app.feedbackDir(), "2020-01-01_000000_bug_first-one_abcd")
	if err := os.MkdirAll(old, 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(old, "bug-first-one.md"), []byte("# Bug: First [one]\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	idea := bugForm("Second")
	idea.Set("kind", "idea")
	e.sendFeedback(t, player, idea, nil) // sending rewrites the index

	got := read()
	if !strings.HasPrefix(got, "# Feedback not read: 2\n") {
		t.Errorf("index heading:\n%s", got)
	}
	second, first := strings.Index(got, "Idea: [Second]"), strings.Index(got, `2020-01-01 00:00 UTC · Bug: [First \[one\]]`)
	if second < 0 || first < 0 || second > first {
		t.Errorf("want both reports, newest first:\n%s", got)
	}
	for path := range e.reports(t) {
		if !strings.Contains(got, "(feedback/"+path+")") {
			t.Errorf("index lacks a link to %s:\n%s", path, got)
		}
	}

	if err := os.RemoveAll(e.app.feedbackDir()); err != nil {
		t.Fatal(err)
	}
	if err := e.app.WriteFeedbackIndex(); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(index); !os.IsNotExist(err) {
		t.Errorf("index kept with nothing to read: %v", err)
	}
}

// TestFeedbackTooBig: an upload over the size cap says so, rather than
// blaming the title.
func TestFeedbackTooBig(t *testing.T) {
	e := newTestEnv(t)
	big := strings.Repeat("x", views.MaxScreenshots*maxAttachmentBytes+2<<20)
	w := e.sendFeedback(t, e.token(t, auth.RolePlayer), bugForm("Big"), map[string]string{"big.png": big})
	if w.Code != http.StatusUnprocessableEntity || !strings.Contains(w.Body.String(), "too big") {
		t.Errorf("status %d: %s", w.Code, w.Body)
	}
	if n := len(e.reports(t)); n != 0 {
		t.Errorf("saved %d reports", n)
	}
}

// TestAutoReport: a server error writes a report by itself, once per
// device in a burst, with the device's requests; visitors get none.
func TestAutoReport(t *testing.T) {
	e := newTestEnv(t)
	e.router.GET("/boom", func(c *gin.Context) { e.app.serverError(c, errors.New("disk full")) })
	player := e.token(t, auth.RolePlayer)
	e.do(t, http.MethodGet, "/boom", "", nil, false) // a visitor
	if n := len(e.reports(t)); n != 0 {
		t.Fatalf("a visitor's error wrote %d reports", n)
	}

	e.do(t, http.MethodGet, "/boom", player, nil, true)
	e.do(t, http.MethodGet, "/boom", player, nil, true)
	reports := e.reports(t)
	if len(reports) != 1 {
		t.Fatalf("got %d reports, want 1", len(reports))
	}
	for path, md := range reports {
		if filepath.Base(path) != "error-get-boom.md" {
			t.Errorf("report file %q", path)
		}
		for _, want := range []string{"# Server error: GET /boom\n", "| From | Alice |", "## Error\n\ndisk full\n"} {
			if !strings.Contains(md, want) {
				t.Errorf("report lacks %q:\n%s", want, md)
			}
		}
		if n := strings.Count(md, "| GET /boom |  | 500 | disk full |"); n != 1 {
			t.Errorf("failed table lists the error %d times, want 1 (the second came after):\n%s", n, md)
		}
	}
	index, _ := os.ReadFile(filepath.Join(e.app.DataDir, FeedbackIndexName))
	if !strings.Contains(string(index), "Server error: [GET /boom]") {
		t.Errorf("index lacks the report:\n%s", index)
	}
}
