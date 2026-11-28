package web

// The phone handoff: moving a take-home or bring-back zip between a phone and
// the player app over the home Wi-Fi. A temporary plain-HTTP listener on the
// private addresses (never loopback, never the app's port) serves one
// unguessable path for one transfer or ten minutes. It has no session, cookie
// or database access: the secret is the whole authorization.
//
// shortcut: a send closed by the timeout while a download runs leaves its temp
// file behind on Windows (an open file cannot be deleted); the OS temp
// cleanup takes it.

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"

	"gm-session-manager/internal/export"
	"gm-session-manager/internal/join"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

// handoffTTL is how long a handoff waits for the phone; tests shorten it.
var handoffTTL = 10 * time.Minute

// handoffLinger keeps the listeners up a moment after the transfer, so the
// phone's last page can still load its stylesheet; tests shorten it.
var handoffLinger = 3 * time.Second

// handoff is the one transfer in progress or just ended. Its mutable fields
// are guarded by App.handoffMu.
type handoff struct {
	secret  string
	send    bool
	servers []*http.Server
	urls    []string
	qr      string
	timer   *time.Timer
	state   string // views.Handoff…
	// file is the temp file: the zip to send, or the one received.
	file        string
	name        string
	characterID int64
	takenAt     time.Time
	fileName    string
	size        int64
	expires     time.Time
}

// startHandoff replaces any handoff with a new one listening on ips. A send
// passes the prepared zip in h.file, name, fileName, size and characterID. No
// address that binds leaves a handoff in the no-network state.
func (a *App) startHandoff(ips []net.IP, h *handoff) *handoff {
	a.handoffMu.Lock()
	defer a.handoffMu.Unlock()
	a.stopLocked()

	h.secret = rand.Text()
	h.state = views.HandoffWaiting
	port := "0"
	for _, ip := range ips {
		ln, err := (&net.ListenConfig{}).Listen(context.Background(), "tcp4", net.JoinHostPort(ip.String(), port))
		if err != nil {
			continue
		}
		if port == "0" {
			_, port, _ = net.SplitHostPort(ln.Addr().String())
		}
		srv := &http.Server{Handler: a.handoffMux(h), ReadHeaderTimeout: 10 * time.Second}
		go func() { _ = srv.Serve(ln) }()
		h.servers = append(h.servers, srv)
		h.urls = append(h.urls, "http://"+ln.Addr().String()+"/"+h.secret)
	}
	if len(h.servers) == 0 {
		h.state = views.HandoffNoNetwork
		removeFile(h)
	} else {
		if png, err := qrcode.Encode(h.urls[0], qrcode.Medium, qrSize); err == nil {
			h.qr = "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)
		}
		h.expires = time.Now().Add(handoffTTL)
		h.timer = time.AfterFunc(handoffTTL, func() { a.expireHandoff(h) })
	}
	a.handoff = h
	return h
}

// expireHandoff closes a handoff still waiting when its time is up.
func (a *App) expireHandoff(h *handoff) {
	a.handoffMu.Lock()
	defer a.handoffMu.Unlock()
	if h.state != views.HandoffWaiting {
		return
	}
	h.state = views.HandoffClosed
	closeServers(h)
	removeFile(h)
}

// stopHandoff ends the current handoff, if any, and removes its temp file.
func (a *App) stopHandoff() {
	a.handoffMu.Lock()
	defer a.handoffMu.Unlock()
	a.stopLocked()
}

func (a *App) stopLocked() {
	h := a.handoff
	if h == nil {
		return
	}
	a.handoff = nil
	h.state = views.HandoffClosed
	closeServers(h)
	removeFile(h)
}

func closeServers(h *handoff) {
	if h.timer != nil {
		h.timer.Stop()
	}
	for _, s := range h.servers {
		_ = s.Close()
	}
}

func removeFile(h *handoff) {
	if h.file != "" {
		_ = os.Remove(h.file)
		h.file = ""
	}
}

// shutdownServers closes the listeners after handoffLinger, letting requests
// in flight finish; the handler that calls it runs on one of them. Only the
// stylesheet still answers meanwhile: the state is no longer waiting.
func shutdownServers(h *handoff) {
	time.AfterFunc(handoffLinger, func() {
		for _, s := range h.servers {
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if s.Shutdown(ctx) != nil {
					_ = s.Close()
				}
			}()
		}
	})
}

// handoffReceived opens the file the phone sent; the caller closes it.
func (a *App) handoffReceived() (*os.File, int64, error) {
	a.handoffMu.Lock()
	defer a.handoffMu.Unlock()
	h := a.handoff
	if h == nil || h.state != views.HandoffReceived {
		return nil, 0, userf("Nothing was received from the phone.")
	}
	f, err := os.Open(h.file)
	if err != nil {
		return nil, 0, err
	}
	st, err := f.Stat()
	if err != nil {
		_ = f.Close()
		return nil, 0, err
	}
	return f, st.Size(), nil
}

// handoffView is the panel for the current state.
func (a *App) handoffView() views.HandoffView {
	a.handoffMu.Lock()
	defer a.handoffMu.Unlock()
	h := a.handoff
	if h == nil {
		return views.HandoffView{}
	}
	return views.HandoffView{
		ID: h.secret, State: h.state, Send: h.send, Name: h.name, CharacterID: h.characterID, TakenAt: h.takenAt,
		URLs: h.urls, QR: h.qr, Expires: h.expires,
	}
}

// handoffAddrs are the addresses a handoff listens on; tests set handoffIPsFn.
func (a *App) handoffAddrs() []net.IP {
	if a.handoffIPsFn != nil {
		return a.handoffIPsFn()
	}
	return handoffIPs()
}

// handoffIPs are the private IPv4 addresses the phone could reach, the
// default route's first.
func handoffIPs() []net.IP {
	all, err := join.LANAddresses()
	if err != nil {
		return nil
	}
	var v4 []net.IP
	for _, ip := range all {
		if ip.To4() != nil {
			v4 = append(v4, ip)
		}
	}
	return join.PreferDefaultRoute(v4)
}

// handoffMux serves the phone's pages: the secret path and nothing else.
func (a *App) handoffMux(h *handoff) http.Handler {
	base := "/" + h.secret
	mux := http.NewServeMux()
	page := func(w http.ResponseWriter, r *http.Request, status int, v views.PhoneView) {
		v.Path, v.Send = base, h.send
		if h.send { // a receive's fields change under the lock and the page does not show them
			v.Name, v.FileName, v.Size = h.name, h.fileName, h.size
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(status)
		_ = views.PhonePage(v).Render(r.Context(), w)
	}
	// open is true while the handoff still takes its transfer.
	open := func() bool {
		a.handoffMu.Lock()
		defer a.handoffMu.Unlock()
		return h.state == views.HandoffWaiting
	}
	mux.HandleFunc("GET "+base, func(w http.ResponseWriter, r *http.Request) {
		if !open() {
			http.NotFound(w, r)
			return
		}
		page(w, r, http.StatusOK, views.PhoneView{})
	})
	mux.HandleFunc("GET "+base+views.PhoneCSS, func(w http.ResponseWriter, r *http.Request) {
		// Whatever the state: the page after the transfer needs it too, and it is public.
		css, err := assetsFS.ReadFile("assets/css/output.css")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		_, _ = w.Write(css)
	})
	if h.send {
		mux.HandleFunc("GET "+base+route.File, func(w http.ResponseWriter, r *http.Request) { a.handoffDownload(w, r, h) })
	} else {
		mux.HandleFunc("POST "+base, func(w http.ResponseWriter, r *http.Request) {
			if !open() {
				http.NotFound(w, r)
				return
			}
			if msg := a.handoffUpload(w, r, h); msg != "" {
				page(w, r, http.StatusUnprocessableEntity, views.PhoneView{Error: msg})
				return
			}
			page(w, r, http.StatusOK, views.PhoneView{Done: true})
			shutdownServers(h)
		})
	}
	mux.HandleFunc("/", http.NotFound)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hd := w.Header()
		hd.Set("Content-Security-Policy", contentSecurityPolicy)
		hd.Set("X-Content-Type-Options", "nosniff")
		hd.Set("Cache-Control", "no-store")
		mux.ServeHTTP(w, r)
	})
}

// handoffUpload takes the posted file into a temp file and, when it is a
// take-home file, keeps it as the handoff's. It returns the refusal to show
// ("" for success); the handoff stays open after a refusal.
func (a *App) handoffUpload(w http.ResponseWriter, r *http.Request, h *handoff) string {
	r.Body = http.MaxBytesReader(w, r.Body, maxTakeHomeBytes)
	mr, err := r.MultipartReader()
	if err != nil {
		return "Choose a take-home file (.zip)."
	}
	var part *multipart.Part
	for {
		if part, err = mr.NextPart(); err != nil || part.FormName() == "file" {
			break
		}
	}
	if err != nil {
		return "Choose a take-home file (.zip)."
	}
	tmp, err := os.CreateTemp("", "gsm-handoff-*.zip")
	if err != nil {
		return "The computer could not store the file."
	}
	keep := false
	defer func() {
		_ = tmp.Close()
		if !keep {
			_ = os.Remove(tmp.Name())
		}
	}()
	size, err := io.Copy(tmp, part)
	if err != nil {
		return "The file could not be received."
	}
	kind, err := export.ManifestKind(tmp, size)
	if err != nil {
		return "That is not a take-home file: " + err.Error()
	}
	if kind != export.KindStarter {
		kind = export.KindTakeHome
	}
	file, err := export.ReadTakeHome(tmp, size, kind)
	if err != nil {
		return "That is not a take-home file: " + err.Error()
	}
	name := file.Character.Name
	if kind == export.KindStarter {
		name = file.Manifest.Campaign.Name
	}
	a.handoffMu.Lock()
	defer a.handoffMu.Unlock()
	if h.state != views.HandoffWaiting {
		return "This page is closed. Start again on the computer."
	}
	if h.timer != nil {
		h.timer.Stop()
	}
	h.state, h.file, h.name, h.takenAt, keep = views.HandoffReceived, tmp.Name(), name, file.Manifest.CreatedAt, true
	return ""
}

// handoffDownload writes the prepared zip; once all of it is out the handoff is sent.
func (a *App) handoffDownload(w http.ResponseWriter, r *http.Request, h *handoff) {
	a.handoffMu.Lock()
	path := h.file
	waiting := h.state == views.HandoffWaiting
	a.handoffMu.Unlock()
	if !waiting {
		http.NotFound(w, r)
		return
	}
	f, err := os.Open(path) //nolint:gosec // our own temp file, set by the handler that started this handoff
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer func() { _ = f.Close() }()
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Length", strconv.FormatInt(h.size, 10))
	w.Header().Set("Content-Disposition", `attachment; filename="`+h.fileName+`"`)
	n, err := io.Copy(w, f)
	if err != nil || n != h.size {
		return
	}
	_ = f.Close() // Windows cannot delete an open file
	a.handoffMu.Lock()
	if h.state == views.HandoffWaiting {
		h.state = views.HandoffSent
		if h.timer != nil {
			h.timer.Stop()
		}
		removeFile(h)
	}
	a.handoffMu.Unlock()
	shutdownServers(h)
}

// handoffReceive starts a handoff that takes a take-home file from the phone.
func (a *App) handoffReceive(c *gin.Context) error {
	a.startHandoff(a.handoffAddrs(), &handoff{})
	a.render(c, http.StatusOK, views.HandoffPanel(a.handoffView()))
	return nil
}

// handoffSend writes the copy's bring-back zip to a temp file and offers it to the phone.
func (a *App) handoffSend(c *gin.Context) error {
	id, ok := a.idParam(c, "character")
	if !ok {
		return nil
	}
	th, name, err := a.bringBack(c.Request.Context(), id)
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp("", "gsm-handoff-*.zip")
	if err != nil {
		return err
	}
	err = export.WriteTakeHome(tmp, th, a.attachmentsDir())
	var size int64
	if err == nil {
		var st os.FileInfo
		if st, err = tmp.Stat(); err == nil {
			size = st.Size()
		}
	}
	if cerr := tmp.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		_ = os.Remove(tmp.Name())
		return err
	}
	a.startHandoff(a.handoffAddrs(), &handoff{send: true, file: tmp.Name(), name: th.Character.Name, characterID: id, fileName: name, size: size})
	a.render(c, http.StatusOK, views.HandoffPanel(a.handoffView()))
	return nil
}

// handoffPanel is the panel. A poll names the handoff and state it shows:
// unchanged is a 204, which htmx does not swap (a swap would close the open
// help and flash the time); changed reloads the page, so the character list
// below agrees with the panel, and a tab showing a replaced handoff catches up.
func (a *App) handoffPanel(c *gin.Context) error {
	v := a.handoffView()
	if state, ok := c.GetQuery(views.HandoffStateParam); ok {
		if state != v.State || c.Query(views.HandoffIDParam) != v.ID {
			c.Header("HX-Refresh", strTrue)
		}
		c.Status(http.StatusNoContent)
		return nil
	}
	a.render(c, http.StatusOK, views.HandoffPanel(v))
	return nil
}

// handoffStopRoute ends the handoff and empties the panel.
func (a *App) handoffStopRoute(c *gin.Context) error {
	a.stopHandoff()
	a.render(c, http.StatusOK, views.HandoffPanel(views.HandoffView{}))
	return nil
}
