package web

import (
	"bytes"
	"context"
	"crypto/sha256"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"gm-session-manager/auth"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/hub"
	"gm-session-manager/perm"
	"gm-session-manager/web/views"
)

//go:embed all:assets
var assetsFS embed.FS

func (a *App) mountAssets(r *gin.Engine) {
	sub, err := fs.Sub(assetsFS, "assets")
	if err != nil {
		panic(err) // the embed pattern guarantees the directory
	}
	etags := assetETags(sub)
	files := http.FileServerFS(sub)
	r.GET("/assets/*path", func(c *gin.Context) {
		p := c.Param("path")
		if strings.HasSuffix(p, "/") {
			c.Status(http.StatusNotFound) // no directory listings
			return
		}
		if immutableAsset(p) {
			c.Header("Cache-Control", "public, max-age=31536000, immutable")
		} else {
			c.Header("Cache-Control", "no-cache")
			if etag, ok := etags[p]; ok {
				c.Header("ETag", etag) // embedded files have no ModTime, so this is the only validator
			}
		}
		c.Request.URL.Path = p
		files.ServeHTTP(c.Writer, c.Request)
	})
}

// versionedPath matches a vendored file under a "name@1.2.3" directory.
var versionedPath = regexp.MustCompile(`^/vendor/.*@\d`)

// immutableAsset: content-hashed and versioned files never change at a URL.
// Unversioned vendor files (htmx.min.js, the tiptap-*.js aliases) do on an
// upgrade, so they revalidate like our own.
func immutableAsset(p string) bool {
	return strings.Contains(p, "/shadcn-templ-") || versionedPath.MatchString(p)
}

// assetETags hashes every revalidated asset once, keyed by its "/path". Files
// are streamed into the hash, so only the short tags stay in memory.
func assetETags(sub fs.FS) map[string]string {
	etags := map[string]string{}
	err := fs.WalkDir(sub, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || immutableAsset("/"+p) {
			return err
		}
		f, err := sub.Open(p)
		if err != nil {
			return err
		}
		defer func() { _ = f.Close() }()
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			return err
		}
		etags["/"+p] = fmt.Sprintf(`"%x"`, h.Sum(nil)[:8])
		return nil
	})
	if err != nil {
		panic(err) // embedded files always read
	}
	return etags
}

// who is the requester: their identity, viewer and (for players) character.
type who struct {
	id        auth.Identity
	viewer    perm.Viewer
	character *sqlc.Character
}

func (w who) isGM() bool { return w.viewer.IsGM }

// campaignID is the player's campaign (0 for the GM and visitors).
func (w who) campaignID() int64 {
	if w.character == nil {
		return 0
	}
	return w.character.CampaignID
}

// whoami resolves the requester. A player whose character vanished is
// treated as a visitor (the session row cascaded away with it).
func (a *App) whoami(c *gin.Context) (who, error) {
	id := auth.FromContext(c)
	w := who{id: id}
	if id.IsAdmin() {
		w.viewer = perm.GM
		return w, nil
	}
	if charID, ok := id.CharacterID(); ok {
		ch, err := a.Store.GetCharacter(c.Request.Context(), charID)
		if errors.Is(err, sql.ErrNoRows) {
			return w, nil
		}
		if err != nil {
			return w, err
		}
		w.viewer, w.character = perm.Player(charID), &ch
	}
	return w, nil
}

// campaignAccess loads a campaign the requester may see: the GM sees every
// campaign, a player only their own. It writes the error response itself.
func (a *App) campaignAccess(c *gin.Context, campaignID int64) (who, sqlc.Campaign, bool) {
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return w, sqlc.Campaign{}, false
	}
	if !w.isGM() && (w.character == nil || w.character.CampaignID != campaignID) {
		a.fail(c, http.StatusForbidden, "You are not part of that campaign.")
		return w, sqlc.Campaign{}, false
	}
	camp, err := a.Store.GetCampaign(c.Request.Context(), campaignID)
	if !a.found(c, err) {
		return w, sqlc.Campaign{}, false
	}
	return w, camp, true
}

// found handles a single-row lookup error: false (with a response written)
// for no rows or a real error.
func (a *App) found(c *gin.Context, err error) bool {
	switch {
	case err == nil:
		return true
	case errors.Is(err, sql.ErrNoRows):
		a.fail(c, http.StatusNotFound, "Not found.")
	default:
		a.serverError(c, err)
	}
	return false
}

// idParam parses a numeric route parameter, answering 404 when malformed.
func (a *App) idParam(c *gin.Context, name string) (int64, bool) {
	id, err := strconv.ParseInt(c.Param(name), 10, 64)
	if err != nil || id <= 0 {
		a.fail(c, http.StatusNotFound, "Not found.")
		return 0, false
	}
	return id, true
}

func isHTMX(c *gin.Context) bool { return c.GetHeader("HX-Request") == "true" }

// fail reports an error: as an out-of-band flash for htmx requests (status
// 422, which app.js tells htmx to swap), or as a plain page.
func (a *App) fail(c *gin.Context, status int, msg string) {
	a.failBack(c, status, msg, PathHome, "Back to the start")
}

// failBack is fail with the error page linking to backURL instead of home.
func (a *App) failBack(c *gin.Context, status int, msg, backURL, backLabel string) {
	if isHTMX(c) {
		a.render(c, http.StatusUnprocessableEntity, views.Flash(msg, true))
		c.Abort()
		return
	}
	a.render(c, status, views.ErrorPage(views.Layout{Title: http.StatusText(status)}, msg, backURL, backLabel))
	c.Abort()
}

func (a *App) serverError(c *gin.Context, err error) {
	slog.Error("request failed", "path", c.FullPath(), "err", err)
	a.fail(c, http.StatusInternalServerError, "Something went wrong on the host. Check its log.")
}

func (a *App) render(c *gin.Context, status int, comp templ.Component) {
	var buf bytes.Buffer
	if err := comp.Render(c.Request.Context(), &buf); err != nil {
		slog.Error("render", "path", c.FullPath(), "err", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Data(status, "text/html; charset=utf-8", buf.Bytes())
}

// redirect navigates: HX-Redirect for htmx requests, 303 otherwise.
func redirect(c *gin.Context, url string) {
	if isHTMX(c) {
		c.Header("HX-Redirect", url)
		c.Status(http.StatusNoContent)
		return
	}
	c.Redirect(http.StatusSeeOther, url)
}

// changed publishes topics to the campaign. Every open page, the actor's
// included, gets the changed data over its websocket; the response only
// clears the flash (live.js refetches the page's views itself when its
// websocket is down).
func (a *App) changed(c *gin.Context, campaignID int64, topics ...hub.Topic) {
	a.changedNotice(c, campaignID, "", topics...)
}

// changedNotice is changed with a success message in the flash.
func (a *App) changedNotice(c *gin.Context, campaignID int64, msg string, topics ...hub.Topic) {
	a.Hub.Publish(campaignID, topics...)
	if !isHTMX(c) {
		back := c.GetHeader("Referer")
		if back == "" {
			back = PathHome
		}
		c.Redirect(http.StatusSeeOther, back)
		return
	}
	a.render(c, http.StatusOK, views.Flash(msg, false))
}

// layout fills the frame for a page. Pages that show live data pass a
// websocket URL; the hub position is read here, before the page's data.
func (a *App) layout(w who, title, active string, camp *sqlc.Campaign, live bool) views.Layout {
	l := views.Layout{Title: title, Active: active, IsGM: w.isGM(), HasSession: w.id.Session != nil}
	if w.character != nil {
		l.CharacterID = w.character.ID
	}
	if camp != nil {
		l.CampaignID, l.CampaignName = camp.ID, camp.Name
	}
	// The GM's pages are all live: join requests show wherever they are.
	if live && camp != nil || w.isGM() {
		epoch, seq := a.Hub.Position()
		url := "/ws"
		if w.isGM() {
			url = fmt.Sprintf("/ws?campaign=%d", l.CampaignID)
			l.Requests = a.pendingViews()
		}
		l.WS = &views.WS{URL: url, Epoch: epoch, Seq: seq}
	}
	return l
}

func nullID(id int64) sql.NullInt64 {
	return sql.NullInt64{Int64: id, Valid: id != 0}
}

// formInt reads an optional integer form field.
func formInt(c *gin.Context, name string) (int64, bool) {
	v := strings.TrimSpace(c.PostForm(name))
	if v == "" {
		return 0, false
	}
	n, err := strconv.ParseInt(v, 10, 64)
	return n, err == nil
}

// bindForm binds and validates a form struct, answering 422 on failure.
func (a *App) bindForm(c *gin.Context, dst any) bool {
	if err := c.ShouldBind(dst); err != nil {
		a.fail(c, http.StatusUnprocessableEntity, "Please check the form: "+err.Error())
		return false
	}
	return true
}

// uaPick returns the label of the first pair whose token is in ua.
func uaPick(ua string, pairs [][2]string) string {
	for _, p := range pairs {
		if strings.Contains(ua, p[0]) {
			return p[1]
		}
	}
	return ""
}

// deviceOS names a User-Agent's operating system ("iPhone", "Windows"), or "".
func deviceOS(ua string) string {
	return uaPick(ua, [][2]string{{"iPhone", "iPhone"}, {"iPad", "iPad"}, {"Android", "Android"}, {"Windows", "Windows"}, {"Mac OS", "Mac"}, {"CrOS", "Chromebook"}, {"Linux", "Linux"}})
}

// userAgentSummary turns a User-Agent into "iPhone · Safari".
func userAgentSummary(ua string) string {
	device := deviceOS(ua)
	browser := uaPick(ua, [][2]string{{"Firefox/", "Firefox"}, {"FxiOS", "Firefox"}, {"Edg/", "Edge"}, {"OPR/", "Opera"}, {"CriOS", "Chrome"}, {"Chrome/", "Chrome"}, {"Safari/", "Safari"}})
	switch {
	case device == "" && browser == "":
		return "Unknown device"
	case device == "":
		return browser
	case browser == "":
		return device
	}
	return device + " · " + browser
}

// isPhone reports whether a User-Agent is a phone: browsers mark phones,
// and not tablets, with "Mobi".
func isPhone(ua string) bool { return strings.Contains(ua, "Mobi") }

// hostLookupTimeout bounds the reverse lookup of a joining device's name.
const hostLookupTimeout = time.Second

// lanHostname asks the network (usually the router, from DHCP) for the name
// of the device at remoteAddr, "" when nobody knows it.
// shortcut: synchronous, so the waiting page can take up to hostLookupTimeout when the router does not answer; move to a goroutine and republish if that is noticeable.
func lanHostname(ctx context.Context, remoteAddr string) string {
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return ""
	}
	ctx, cancel := context.WithTimeout(ctx, hostLookupTimeout)
	defer cancel()
	names, err := net.DefaultResolver.LookupAddr(ctx, ip)
	if err != nil || len(names) == 0 {
		return ""
	}
	return shortHost(names[0])
}

// shortHost keeps a hostname's first label: "miras-pc.lan." is "miras-pc".
func shortHost(name string) string {
	name, _, _ = strings.Cut(strings.TrimSuffix(name, "."), ".")
	return name
}

// formIntQuery reads an optional positive integer query parameter, 0 when
// missing or invalid.
func formIntQuery(c *gin.Context, name string) int64 {
	n, err := strconv.ParseInt(c.Query(name), 10, 64)
	if err != nil || n < 0 {
		return 0
	}
	return n
}
