package web

import (
	"io"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"gm-session-manager/auth"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/web/views"
)

// Each device's recent requests, kept in memory for its bug reports.
const (
	reqLogSize = 30
	maxErrBody = 300
)

// hostDevice keys the host's browser while it has no session.
const hostDevice = "host"

type ctxKey int

// failMsgKey and failStatusKey hold what fail sent back, which the request
// log prefers over the HTML it was wrapped in. serverErrKey holds the real
// error behind a 500: kept here, not in c.Errors, which gin's logger would
// print a second time.
const (
	failMsgKey ctxKey = iota
	failStatusKey
	serverErrKey
)

type reqEntry struct {
	At     time.Time
	Method string
	Path   string
	Status int
	Took   time.Duration
	// Sent is the start of what the request sent (sentText), if anything.
	Sent string
	// Error is the real server error, else the message or body sent back.
	Error string
}

type reqLog struct {
	mu      sync.Mutex
	devices map[string][]reqEntry
	// lastAuto is when each device last got an automatic report.
	lastAuto map[string]time.Time
}

// autoReportGap keeps a burst of server errors to one report per device.
const autoReportGap = 5 * time.Minute

// autoReportDue reports whether the device may get an automatic report at
// now, and if so counts it as given.
func (l *reqLog) autoReportDue(key string, now time.Time) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if last, ok := l.lastAuto[key]; ok && now.Sub(last) < autoReportGap {
		return false
	}
	if l.lastAuto == nil {
		l.lastAuto = map[string]time.Time{}
	}
	l.lastAuto[key] = now
	return true
}

// deviceKey names the device a request comes from; visitors have none.
func deviceKey(id auth.Identity) (string, bool) {
	if id.Session != nil {
		return sessionKey(*id.Session), true
	}
	return hostDevice, id.Loopback
}

// sessionKey is the session's token hash: random and never reused, unlike
// its row ID, which SQLite hands out again once the newest row is deleted.
func sessionKey(s sqlc.Session) string { return string(s.TokenHash) }

// recordRequest is middleware logging each request of a known device, and
// writing a report by itself when the server fails one.
func (a *App) recordRequest(c *gin.Context) {
	key, ok := deviceKey(auth.FromContext(c))
	path := c.Request.URL.Path
	// Logout ends the device's log; logging it would start a new one.
	if !ok || path == "/ws" || path == views.PathFeedback || path == PathLogout || strings.HasPrefix(path, "/assets/") {
		c.Next()
		return
	}
	start := time.Now()
	w := &errBodyWriter{ResponseWriter: c.Writer}
	c.Writer = w
	var sent *capReader
	if c.Request.Method != http.MethodGet && keepsBody(path) {
		sent = &capReader{r: c.Request.Body}
		c.Request.Body = sent
	}
	c.Next()
	e := reqEntry{At: start.UTC(), Method: c.Request.Method, Path: c.Request.URL.RequestURI(), Status: w.Status(), Took: time.Since(start)}
	if sent != nil {
		e.Sent = sentText(c.ContentType(), sent.kept)
	}
	// htmx gets every failure as 422; log the status it stands for.
	if s := c.GetInt(failStatusKey); s != 0 {
		e.Status = s
	}
	if e.Status >= 400 {
		e.Error = errorText(c, w)
	}
	a.reqs.add(key, e)
	if c.GetString(serverErrKey) != "" && a.reqs.autoReportDue(key, e.At) {
		a.autoReport(c, key, e)
	}
}

// keepsBody is false where a body must not reach a report: credentials,
// and note blocks, which can be a player's private text.
func keepsBody(path string) bool {
	return path != PathAdminLogin && !strings.HasPrefix(path, "/blocks/")
}

// Bounds on what is kept of a request body.
const (
	maxKeptBody = 1 << 10
	maxSent     = 300
)

// capReader passes a body through to the handler, keeping its start.
type capReader struct {
	r    io.ReadCloser
	kept []byte
}

func (c *capReader) Read(p []byte) (int, error) {
	n, err := c.r.Read(p)
	if room := maxKeptBody - len(c.kept); room > 0 {
		c.kept = append(c.kept, p[:min(n, room)]...)
	}
	return n, err
}

func (c *capReader) Close() error { return c.r.Close() }

// sentText shows what a request sent: form fields with sorted keys and
// passwords masked, JSON as is, uploads only as such.
func sentText(contentType string, body []byte) string {
	var s string
	switch {
	case len(body) == 0:
		return ""
	case contentType == binding.MIMEMultipartPOSTForm:
		return "(file upload)"
	case contentType == binding.MIMEPOSTForm:
		form, _ := url.ParseQuery(string(body)) // a cut body still parses up to the cut
		var fields []string
		for _, k := range slices.Sorted(maps.Keys(form)) {
			v := strings.Join(form[k], ",")
			if strings.Contains(strings.ToLower(k), "password") {
				v = "…"
			}
			fields = append(fields, k+"="+v)
		}
		s = strings.Join(fields, " ")
	default:
		s = string(body)
	}
	s = strings.Join(strings.Fields(s), " ")
	if r := []rune(s); len(r) > maxSent {
		s = string(r[:maxSent]) + "…"
	}
	return s
}

func errorText(c *gin.Context, w *errBodyWriter) string {
	if msg := c.GetString(serverErrKey); msg != "" {
		return msg
	}
	if msg := c.GetString(failMsgKey); msg != "" {
		return msg
	}
	if strings.HasPrefix(w.Header().Get("Content-Type"), "text/html") {
		return ""
	}
	return strings.TrimSpace(string(w.body))
}

func (l *reqLog) add(key string, e reqEntry) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.devices == nil {
		l.devices = map[string][]reqEntry{}
	}
	log := l.devices[key]
	if len(log) == reqLogSize {
		log = append(log[:0], log[1:]...)
	}
	l.devices[key] = append(log, e)
}

// recent returns a device's requests, newest first.
func (l *reqLog) recent(key string) []reqEntry {
	l.mu.Lock()
	defer l.mu.Unlock()
	log := l.devices[key]
	out := make([]reqEntry, len(log))
	for i, e := range log {
		out[len(log)-1-i] = e
	}
	return out
}

// forget drops a device's log when its session ends.
func (l *reqLog) forget(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.devices, key)
	delete(l.lastAuto, key)
}

// errBodyWriter keeps the start of an error response's body.
type errBodyWriter struct {
	gin.ResponseWriter
	body []byte
}

func (w *errBodyWriter) Write(b []byte) (int, error) {
	w.keep(b)
	return w.ResponseWriter.Write(b)
}

func (w *errBodyWriter) WriteString(s string) (int, error) {
	w.keep([]byte(s))
	return w.ResponseWriter.WriteString(s)
}

func (w *errBodyWriter) keep(b []byte) {
	if w.Status() >= 400 && len(w.body) < maxErrBody {
		w.body = append(w.body, b[:min(len(b), maxErrBody-len(w.body))]...)
	}
}
