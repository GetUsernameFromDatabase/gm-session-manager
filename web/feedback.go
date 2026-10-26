package web

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"gm-session-manager/auth"
	"gm-session-manager/web/views"
)

// Bug reports and ideas are files on the host for the GM to open: one
// folder per report, with its Markdown file and screenshots.

// kindError marks a report the server wrote itself for a 500.
const kindError = "error"

// maxSlug caps the title's part of a report's folder name.
const maxSlug = 40

// screenshotExts names the screenshot types kept, by sniffed content type.
var screenshotExts = map[string]string{"image/png": ".png", "image/jpeg": ".jpg", "image/gif": ".gif", "image/webp": ".webp"}

type feedbackForm struct {
	Kind     string `form:"kind"     validate:"oneof=bug idea"`
	Title    string `form:"title"    validate:"required,max=120"`
	Body     string `form:"body"     validate:"required,max=4000"`
	Steps    string `form:"steps"    validate:"max=4000"`
	Expected string `form:"expected" validate:"max=4000"`
	Player   string `form:"player"   validate:"max=100"`
	Page     string `form:"page"     validate:"max=512"`
	Viewport string `form:"viewport" validate:"max=20"`
	Dark     bool   `form:"dark"`
	// Console is the tab's recent console as JSON (app.js), read by parseConsole.
	Console string `form:"console" validate:"max=40000"`
}

// consoleSize is how many console entries app.js keeps per tab.
const consoleSize = 50

type consoleEntry struct {
	At    string `json:"at"    validate:"max=40"`
	Level string `json:"level" validate:"oneof=log info warn error uncaught"`
	Msg   string `json:"msg"   validate:"max=500"`
	Page  string `json:"page"  validate:"max=512"`
}

type consoleLog struct {
	Entries []consoleEntry `validate:"max=50,dive"`
}

// parseConsole reads the console field; a bad one is dropped (ok false)
// rather than refusing the report.
func parseConsole(s string) (entries []consoleEntry, ok bool) {
	if s == "" {
		return nil, true
	}
	var l consoleLog
	if json.Unmarshal([]byte(s), &l.Entries) != nil || binding.Validator.ValidateStruct(&l) != nil {
		return nil, false
	}
	slices.Reverse(l.Entries) // newest first, like the requests
	return l.Entries, true
}

func (a *App) feedbackDir() string { return filepath.Join(a.DataDir, "feedback") }

// feedbackPath is feedbackDir in full, for the GM to find on the host.
func (a *App) feedbackPath() string {
	if abs, err := filepath.Abs(a.feedbackDir()); err == nil {
		return abs
	}
	return a.feedbackDir()
}

func (a *App) sendFeedback(c *gin.Context) {
	key, ok := deviceKey(auth.FromContext(c))
	if !ok {
		a.fail(c, http.StatusForbidden, "Join a campaign to send feedback.")
		return
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, views.MaxScreenshots*maxAttachmentBytes+1<<20)
	var f feedbackForm
	if err := c.ShouldBind(&f); err != nil {
		if tooBig := new(http.MaxBytesError); errors.As(err, &tooBig) {
			a.fail(c, http.StatusRequestEntityTooLarge, "The screenshots are too big. Attach smaller ones.")
			return
		}
		a.fail(c, http.StatusBadRequest, "Give the report a title (up to 120 characters) and a description (up to 4000).")
		return
	}
	var shots []*multipart.FileHeader
	if form, err := c.MultipartForm(); err == nil {
		shots = form.File["screenshots"]
	}
	if len(shots) > views.MaxScreenshots {
		a.fail(c, http.StatusBadRequest, fmt.Sprintf("Attach at most %d screenshots.", views.MaxScreenshots))
		return
	}
	r := report{form: f, from: a.reporter(c, f.Player), sent: time.Now().UTC(), browser: browserName(c.Request.UserAgent())}
	if f.Kind == "bug" {
		r.requests = a.reqs.recent(key)
		r.console, r.consoleOK = parseConsole(f.Console)
	}

	if err := a.saveReport(r, shots); err != nil {
		a.serverError(c, err)
		return
	}
	// Named like the dialog's button: "Send report" / "Send idea".
	a.render(c, http.StatusOK, views.Toast(map[string]string{"bug": "Report sent.", "idea": "Idea sent."}[f.Kind]))
}

// saveReport writes r into a new folder (removed again if that fails) and
// updates the unread list.
func (a *App) saveReport(r report, shots []*multipart.FileHeader) error {
	dir := filepath.Join(a.feedbackDir(), folderName(r.sent, r.form.Kind, r.form.Title))
	if err := r.write(dir, shots); err != nil {
		_ = os.RemoveAll(dir)
		return err
	}
	if err := a.WriteFeedbackIndex(); err != nil {
		slog.Error("feedback index", "err", err)
	}
	return nil
}

// reporter names who a report is from: the GM or the character, plus the
// device's player name if given. The role comes from the session, so a
// report still says "a player" when the database cannot name the character
// (often the very error being reported).
func (a *App) reporter(c *gin.Context, player string) string {
	from := "a player"
	if auth.FromContext(c).IsAdmin() {
		from = "GM"
	} else if w, err := a.whoami(c); err == nil && w.character != nil {
		from = w.character.Name
	}
	if p := strings.TrimSpace(player); p != "" {
		from += " (" + p + ")"
	}
	return from
}

// autoReport writes a report for a server error e on its own, so a 500 is
// on file even when nobody says so. The response has already gone out.
func (a *App) autoReport(c *gin.Context, key string, e reqEntry) {
	r := report{
		form:     feedbackForm{Kind: kindError, Title: e.Method + " " + c.Request.URL.Path, Body: e.Error, Page: e.Path},
		from:     a.reporter(c, ""),
		sent:     e.At,
		browser:  browserName(c.Request.UserAgent()),
		requests: a.reqs.recent(key),
	}
	if err := a.saveReport(r, nil); err != nil {
		slog.Error("automatic error report", "err", err)
	}
}

type report struct {
	form     feedbackForm
	from     string
	sent     time.Time
	browser  string
	requests []reqEntry // newest first; bugs only
	// console is newest first; consoleOK is false when the browser sent a bad one.
	console   []consoleEntry
	consoleOK bool
	shots     []string // file names in the report's folder
	dropped   []string // uploads that were not pictures, as "name (type)"
}

// write saves the screenshots, then the report (named after its title, for
// the editor's file search), into a new folder dir.
func (r report) write(dir string, shots []*multipart.FileHeader) error {
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return err
	}
	for _, fh := range shots {
		name, mime, err := saveScreenshot(dir, fmt.Sprintf("screenshot-%d", len(r.shots)+1), fh)
		if err != nil {
			return err
		}
		if name != "" {
			r.shots = append(r.shots, name)
		} else {
			r.dropped = append(r.dropped, fmt.Sprintf("%s (%s)", filepath.Base(fh.Filename), mime))
		}
	}
	return os.WriteFile(filepath.Join(dir, r.form.Kind+"-"+slugify(r.form.Title)+".md"), []byte(r.markdown()), 0o600)
}

// saveScreenshot stores one upload as base plus the extension of its sniffed
// type, returning its file name and type; the name is "" when it is not a
// picture, and the file is not kept.
func saveScreenshot(dir, base string, fh *multipart.FileHeader) (name, mime string, err error) {
	src, err := fh.Open()
	if err != nil {
		return "", "", err
	}
	defer func() { _ = src.Close() }()
	tmp := filepath.Join(dir, base)
	mime, _, err = saveSniffed(tmp, src)
	if err != nil {
		return "", "", err
	}
	ext, ok := screenshotExts[mime]
	if !ok {
		return "", mime, os.Remove(tmp)
	}
	return base + ext, mime, os.Rename(tmp, tmp+ext)
}

func (r report) markdown() string {
	f := r.form
	var b strings.Builder
	kind := map[string]string{"bug": "Bug", "idea": "Idea", kindError: "Server error"}[f.Kind]
	fmt.Fprintf(&b, "# %s: %s\n\n| | |\n|---|---|\n", kind, oneLine(f.Title))
	row := func(k, v string) {
		if v != "" {
			fmt.Fprintf(&b, "| %s | %s |\n", k, mdCell(v))
		}
	}
	row("From", r.from)
	row("Sent", r.sent.Format("2006-01-02 15:04:05 UTC"))
	row("Page", f.Page)
	row("Build", appBuild())
	device := []string{r.browser}
	if f.Viewport != "" {
		device = append(device, f.Viewport)
	}
	if f.Dark {
		device = append(device, "dark")
	}
	row("Device", strings.Join(device, " · "))

	section := func(title, text string) {
		if text = strings.TrimSpace(text); text != "" {
			fmt.Fprintf(&b, "\n## %s\n\n%s\n", title, text)
		}
	}
	switch f.Kind {
	case "bug":
		section("What happened", f.Body)
		section("Steps to reproduce", f.Steps)
		section("Expected", f.Expected)
		r.requestTables(&b)
		r.consoleTables(&b)
	case kindError:
		section("Error", f.Body)
		r.requestTables(&b)
	default:
		section("Idea", f.Body)
	}
	if len(r.shots)+len(r.dropped) > 0 {
		b.WriteString("\n## Screenshots\n\n")
		for _, s := range r.shots {
			fmt.Fprintf(&b, "![%s](%s)\n", s, s)
		}
		if len(r.dropped) > 0 {
			fmt.Fprintf(&b, "\nLeft out, not a picture: %s\n", oneLine(strings.Join(r.dropped, ", ")))
		}
	}
	return b.String()
}

// requestTables lists the device's failed requests first, then all recent ones.
func (r report) requestTables(b *strings.Builder) {
	at := func(e reqEntry) string { return e.At.Format("15:04:05") }
	req := func(e reqEntry) string { return mdCell(e.Method + " " + e.Path) }
	b.WriteString("\n## Failed requests\n\n")
	failed := false
	for _, e := range r.requests {
		if e.Status < 400 {
			continue
		}
		if !failed {
			b.WriteString("| Time (UTC) | Request | Sent | Status | Error |\n|---|---|---|---|---|\n")
			failed = true
		}
		fmt.Fprintf(b, "| %s | %s | %s | %d | %s |\n", at(e), req(e), mdCell(e.Sent), e.Status, mdCell(e.Error))
	}
	if !failed {
		b.WriteString("None.\n")
	}
	fmt.Fprintf(b, "\n## Recent requests (last %d)\n\n", reqLogSize)
	if len(r.requests) == 0 {
		b.WriteString("None since the host started.\n")
		return
	}
	b.WriteString("| Time (UTC) | Request | Sent | Status | ms |\n|---|---|---|---|---|\n")
	for _, e := range r.requests {
		fmt.Fprintf(b, "| %s | %s | %s | %d | %.1f |\n", at(e), req(e), mdCell(e.Sent), e.Status, float64(e.Took.Microseconds())/1000)
	}
}

// consoleTables lists the browser's errors first, then its whole console.
// A table cell shows a message's first line; full stacks follow the errors.
func (r report) consoleTables(b *strings.Builder) {
	if !r.consoleOK {
		b.WriteString("\n## Browser console\n\nThe browser sent an unreadable console, so it was left out.\n")
		return
	}
	at := func(e consoleEntry) string {
		if t, err := time.Parse(time.RFC3339Nano, e.At); err == nil {
			return t.UTC().Format("15:04:05")
		}
		return mdCell(e.At)
	}
	first := func(msg string) string { line, _, _ := strings.Cut(msg, "\n"); return mdCell(line) }
	var stacks []consoleEntry
	for _, e := range r.console {
		if e.Level != "error" && e.Level != "uncaught" {
			continue
		}
		if len(stacks) == 0 {
			b.WriteString("\n## Browser errors\n\n| Time (UTC) | Page | Message |\n|---|---|---|\n")
		}
		fmt.Fprintf(b, "| %s | %s | %s |\n", at(e), mdCell(e.Page), first(e.Msg))
		stacks = append(stacks, e)
	}
	for _, e := range stacks {
		if strings.Contains(e.Msg, "\n") {
			fmt.Fprintf(b, "\n%s\n\n```\n%s\n```\n", at(e), strings.ReplaceAll(e.Msg, "```", "'''"))
		}
	}
	fmt.Fprintf(b, "\n## Browser console (last %d)\n\n", consoleSize)
	if len(r.console) == 0 {
		b.WriteString("Empty.\n")
		return
	}
	b.WriteString("| Time (UTC) | Level | Page | Message |\n|---|---|---|---|\n")
	for _, e := range r.console {
		fmt.Fprintf(b, "| %s | %s | %s | %s |\n", at(e), e.Level, mdCell(e.Page), first(e.Msg))
	}
}

// oneLine keeps user text on one line, for headings and table cells.
func oneLine(s string) string { return strings.Join(strings.Fields(s), " ") }

// mdCell makes s safe inside a Markdown table cell.
func mdCell(s string) string { return strings.ReplaceAll(oneLine(s), "|", `\|`) }

var notSlug = regexp.MustCompile(`[^a-z0-9]+`)

// slugify keeps only [a-z0-9-] of a title, so it is safe in a path.
func slugify(title string) string {
	slug := strings.Trim(notSlug.ReplaceAllString(strings.ToLower(title), "-"), "-")
	if len(slug) > maxSlug {
		slug = strings.TrimRight(slug[:maxSlug], "-")
	}
	if slug == "" {
		return "report"
	}
	return slug
}

// folderName is a report's folder: time, kind and title, plus a random
// suffix. Only [a-z0-9_-] reach the path, so it stays in the feedback folder.
func folderName(t time.Time, kind, title string) string {
	return fmt.Sprintf("%s_%s_%s_%s", t.Format("2006-01-02_150405"), kind, slugify(title), strings.ToLower(rand.Text()[:4]))
}

// FeedbackIndexName is the file in DATA_DIR listing the reports not read
// yet; it exists only while there are some.
const FeedbackIndexName = "USER_FEEDBACK_NOT_READ"

var feedbackIndexMu sync.Mutex

// WriteFeedbackIndex rewrites the list of unread reports: every report
// folder still in feedback/ (the GM deletes a folder once read).
func (a *App) WriteFeedbackIndex() error {
	// Two reports at once would otherwise write the file over each other.
	feedbackIndexMu.Lock()
	defer feedbackIndexMu.Unlock()
	path := filepath.Join(a.DataDir, FeedbackIndexName)
	dirs, err := os.ReadDir(a.feedbackDir())
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	var lines []string
	for _, d := range slices.Backward(dirs) { // folder names start with the time
		if !d.IsDir() {
			continue
		}
		mds, _ := filepath.Glob(filepath.Join(a.feedbackDir(), d.Name(), "*.md"))
		if len(mds) == 0 {
			continue
		}
		title := strings.TrimPrefix(firstLine(mds[0]), "# ")
		kind, name, ok := strings.Cut(title, ": ")
		if !ok {
			kind, name = "Report", title
		}
		rel := filepath.ToSlash(filepath.Join("feedback", d.Name(), filepath.Base(mds[0])))
		sent := d.Name()
		if t, err := time.Parse("2006-01-02_150405", sent[:min(len(sent), 17)]); err == nil {
			sent = t.Format("2006-01-02 15:04")
		}
		lines = append(lines, fmt.Sprintf("- %s UTC · %s: [%s](%s)", sent, kind, mdLinkText(name), rel))
	}
	if len(lines) == 0 {
		if err := os.Remove(path); err != nil && !errors.Is(err, fs.ErrNotExist) {
			return err
		}
		return nil
	}
	body := fmt.Sprintf("# Feedback not read: %d\n\nDelete a report's folder once read. This list is rebuilt when the app starts and when a report arrives.\n\n%s\n",
		len(lines), strings.Join(lines, "\n"))
	return os.WriteFile(path, []byte(body), 0o600)
}

// firstLine is a file's first line, or "" when it cannot be read.
func firstLine(path string) string {
	b, err := os.ReadFile(path) //nolint:gosec // a report file inside our feedback folder
	if err != nil {
		return ""
	}
	line, _, _ := strings.Cut(string(b), "\n")
	return line
}

// mdLinkText keeps a title from closing its Markdown link early.
func mdLinkText(s string) string { return strings.NewReplacer("[", `\[`, "]", `\]`).Replace(s) }

// appBuild is the commit the host was built from, when Go recorded it.
var appBuild = sync.OnceValue(func() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	var rev, at, dirty string
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			rev = s.Value[:min(7, len(s.Value))]
		case "vcs.time":
			at = ", " + strings.SplitN(s.Value, "T", 2)[0]
		case "vcs.modified":
			if s.Value == "true" {
				dirty = " (modified)"
			}
		}
	}
	if rev == "" {
		return info.Main.Version
	}
	return rev + dirty + at
})

// Browser and OS tokens in User-Agent, most specific first.
var (
	uaBrowsers = []struct{ token, name string }{
		{"Edg/", "Edge"}, {"OPR/", "Opera"}, {"SamsungBrowser/", "Samsung Internet"}, {"Firefox/", "Firefox"},
		{"FxiOS/", "Firefox"}, {"CriOS/", "Chrome"}, {"Chrome/", "Chrome"}, {"Version/", "Safari"},
	}
	uaSystems = []struct{ token, name string }{
		{"Android", "Android"}, {"iPhone", "iPhone"}, {"iPad", "iPad"}, {"Windows", "Windows"},
		{"Mac OS X", "macOS"}, {"CrOS", "ChromeOS"}, {"Linux", "Linux"},
	}
)

// browserName shortens a User-Agent to "Firefox 131 on Android", or keeps
// it whole when it is not one it knows.
func browserName(ua string) string {
	for _, br := range uaBrowsers {
		_, after, ok := strings.Cut(ua, br.token)
		if !ok {
			continue
		}
		version, _, _ := strings.Cut(after, ".")
		name := br.name + " " + strings.TrimSpace(version)
		for _, sys := range uaSystems {
			if strings.Contains(ua, sys.token) {
				return name + " on " + sys.name
			}
		}
		return name
	}
	return ua
}
