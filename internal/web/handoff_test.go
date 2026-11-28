package web

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

var loopbackIPs = func() []net.IP { return []net.IP{net.IPv4(127, 0, 0, 1)} }

// postPlayer posts a plain form to the player app as the browser at home.
func (e testEnv) postPlayer(t *testing.T, path string, form url.Values) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		path,
		strings.NewReader(form.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("HX-Request", "true")
	req.RemoteAddr = loopbackAddr
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	return w
}

// importedHome imports the sample take-home file and returns the home character's id.
func (e testEnv) importedHome(t *testing.T) int64 {
	t.Helper()
	w := e.importRequest(t, sampleTakeHome(t, nil), "", nil)
	if w.Code != http.StatusNoContent {
		t.Fatalf("import: status %d: %s", w.Code, w.Body)
	}
	sess, err := e.mgr.Resolve(t.Context(), homeCookie(w).Value)
	if err != nil {
		t.Fatal(err)
	}
	return sess.CharacterID.Int64
}

func TestBringBack(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	id := e.importedHome(t)
	q := e.q()
	// A home edit: a new note, unknown to the table.
	ch, err := q.GetCharacter(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	note, err := q.CreateNote(
		ctx,
		sqlc.CreateNoteParams{
			CampaignID:       ch.CampaignID,
			Kind:             enum.NoteGeneric,
			Title:            "Home note",
			OwnerCharacterID: nullID(id),
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := q.CreateBlock(
		ctx,
		sqlc.CreateBlockParams{
			NoteID:            note.ID,
			Visibility:        enum.VisOwner,
			AuthorCharacterID: nullID(id),
		},
	); err != nil {
		t.Fatal(err)
	}

	// The Save button's post checks, then sends the browser to the download.
	w := e.postPlayer(t, route.Path(route.AtHomeBringBack, id), nil)
	if w.Code != http.StatusNoContent ||
		w.Header().Get("HX-Redirect") != route.Path(route.AtHomeBringBack, id) {
		t.Fatalf("save: status %d, redirect %q: %s", w.Code, w.Header().Get("HX-Redirect"), w.Body)
	}
	w = e.get(t, route.Path(route.AtHomeBringBack, id), "")
	if w.Code != http.StatusOK || w.Header().Get("Content-Type") != "application/zip" {
		t.Fatalf("status %d %q: %s", w.Code, w.Header().Get("Content-Type"), w.Body)
	}
	body := w.Body.Bytes()
	f, err := export.ReadTakeHome(bytes.NewReader(body), int64(len(body)), export.KindBringBack)
	if err != nil {
		t.Fatal(err)
	}
	if f.Manifest.Token != "tok" || f.Manifest.ServerID != "srv" || f.Manifest.CharacterID != 7 ||
		f.Manifest.SavedAt.IsZero() {
		t.Errorf("manifest %+v", f.Manifest)
	}
	hc, err := q.GetHomeCopy(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	var m export.HomeManifest
	if err := json.Unmarshal([]byte(hc.ManifestJson), &m); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*f.Maps, m.Maps) {
		t.Errorf("maps %+v, want the stored %+v", *f.Maps, m.Maps)
	}
	if _, known := f.Maps.Notes[note.ID]; known {
		t.Error("the new note has a table id")
	}
	var found bool
	for _, n := range f.Notes {
		found = found || n.ID == note.ID
	}
	if !found {
		t.Error("the new note is not in the file")
	}
	if !strings.Contains(w.Header().Get("Content-Disposition"), "-bring-back-") {
		t.Errorf("disposition %q", w.Header().Get("Content-Disposition"))
	}

	// Steps past the last snapshot: a note is still open elsewhere.
	if _, err := e.conn.ExecContext(
		ctx,
		`UPDATE note_blocks SET version = doc_version + 3 WHERE note_id = ?`,
		note.ID,
	); err != nil {
		t.Fatal(err)
	}
	w = e.postPlayer(t, route.Path(route.AtHomeBringBack, id), nil)
	if w.Code != http.StatusUnprocessableEntity ||
		!strings.Contains(w.Body.String(), "still open in another tab") {
		t.Errorf("open note: status %d, want 422 with the reason: %s", w.Code, w.Body)
	}
}

// The handoff and bring-back routes exist in the player app only.
func TestHandoffRoutesPlayerAppOnly(t *testing.T) {
	g := newTestEnv(t)
	for _, r := range []struct{ method, path string }{
		{http.MethodPost, route.Path(route.AtHomeBringBack, 1)},
		{http.MethodGet, route.Path(route.AtHomeBringBack, 1)},
		{http.MethodPost, route.Path(route.AtHomeSendPhone, 1)},
		{http.MethodPost, route.AtHomeReceive},
		{http.MethodGet, route.AtHomeHandoff},
		{http.MethodPost, route.AtHomeHandoffStop},
		// Without the guard anyone could log in as any table character.
		{http.MethodPost, route.Path(route.AtHomePlay, 1)},
		{http.MethodPost, route.Path(route.AtHomeMakeCopy, 1)},
		{http.MethodPost, route.Path(route.AtHomeFallback, 1)},
		{http.MethodPost, route.AtHomeNew},
	} {
		req := httptest.NewRequestWithContext(t.Context(), r.method, r.path, nil)
		req.RemoteAddr = loopbackAddr
		w := httptest.NewRecorder()
		g.router.ServeHTTP(w, req)
		if w.Code != http.StatusNotFound {
			t.Errorf("%s %s in the GM app: status %d, want 404", r.method, r.path, w.Code)
		}
	}
	// And in the player app they are routed, beside /at-home/:character/...
	p := newPlayerEnv(t)
	p.app.handoffIPsFn = loopbackIPs
	w := p.postPlayer(t, route.AtHomeReceive, nil)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Receive from your phone") {
		t.Errorf("receive: status %d: %s", w.Code, w.Body)
	}
	if w := p.get(t, route.AtHomeHandoff, ""); w.Code != http.StatusOK {
		t.Errorf("panel: status %d", w.Code)
	}
	// The poll names the handoff and state it shows: unchanged is a bare 204,
	// changed (or another handoff, started in another tab) reloads the page.
	pollFor := func(id string) string {
		return route.AtHomeHandoff + "?" + url.Values{
			views.HandoffStateParam: {views.HandoffWaiting},
			views.HandoffIDParam:    {id},
		}.Encode()
	}
	poll := pollFor(p.app.handoffView().ID)
	if w := p.get(
		t,
		poll,
		"",
	); w.Code != http.StatusNoContent ||
		w.Header().Get("HX-Refresh") != "" {
		t.Errorf(
			"poll while waiting: status %d, refresh %q; want a bare 204",
			w.Code,
			w.Header().Get("HX-Refresh"),
		)
	}
	if w := p.get(t, pollFor("another-tab"), ""); w.Header().Get("HX-Refresh") != "true" {
		t.Errorf("poll for a replaced handoff: refresh %q, want true", w.Header().Get("HX-Refresh"))
	}
	p.app.expireHandoff(p.app.handoff)
	if w := p.get(
		t,
		poll,
		"",
	); w.Code != http.StatusNoContent ||
		w.Header().Get("HX-Refresh") != "true" {
		t.Errorf(
			"poll after expiry: status %d, refresh %q; want a reload",
			w.Code,
			w.Header().Get("HX-Refresh"),
		)
	}
	if w := p.postPlayer(
		t,
		route.AtHomeHandoffStop,
		nil,
	); w.Code != http.StatusOK ||
		p.app.handoffView().State != views.HandoffNone {
		t.Errorf("stop: status %d, state %q", w.Code, p.app.handoffView().State)
	}
}

func phoneUpload(t *testing.T, target string, zipBody []byte) *phoneResp {
	t.Helper()
	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	part, err := mw.CreateFormFile("file", "x.zip")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = part.Write(zipBody)
	_ = mw.Close()
	req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, target, &buf)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	return &phoneResp{resp.StatusCode, resp.Header, body}
}

// phoneResp is a response read whole.
type phoneResp struct {
	StatusCode int
	Header     http.Header
	Body       []byte
}

func phoneGet(t *testing.T, target string) (*phoneResp, error) {
	t.Helper()
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, target, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	return &phoneResp{resp.StatusCode, resp.Header, body}, err
}

// eventually polls until cond holds; servers shut down after their response.
func eventually(t *testing.T, what string, cond func() bool) {
	t.Helper()
	for range 100 {
		if cond() {
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Errorf("never: %s", what)
}

// shortLinger closes listeners soon after a transfer, within eventually's wait.
func shortLinger(t *testing.T) {
	old := handoffLinger
	handoffLinger = time.Second
	t.Cleanup(func() { handoffLinger = old })
}

func TestHandoffReceive(t *testing.T) {
	shortLinger(t)
	e := newPlayerEnv(t)
	h := e.app.startHandoff(loopbackIPs(), &handoff{})
	if len(h.urls) != 1 || h.state != views.HandoffWaiting {
		t.Fatalf("handoff %+v", h)
	}
	target := h.urls[0]

	if resp, err := phoneGet(
		t,
		strings.Replace(target, h.secret, "WRONG", 1),
	); err != nil ||
		resp.StatusCode != http.StatusNotFound {
		t.Errorf("wrong secret: %v %v", resp, err)
	}
	if resp, err := phoneGet(
		t,
		target,
	); err != nil || resp.StatusCode != http.StatusOK ||
		!strings.Contains(resp.Header.Get("Content-Security-Policy"), "default-src") {
		t.Errorf("phone page: %v %v", resp, err)
	}
	if resp, err := phoneGet(
		t,
		target+views.PhoneCSS,
	); err != nil ||
		resp.StatusCode != http.StatusOK {
		t.Errorf("stylesheet: %v %v", resp, err)
	}

	// A bring-back file is refused and does not use up the slot.
	bring := sampleTakeHome(t, func(th *export.TakeHome) {
		th.Manifest.Kind = export.KindBringBack
		th.Maps = &export.HomeIDMaps{}
	})
	resp := phoneUpload(t, target, bring)
	msg := resp.Body
	if resp.StatusCode != http.StatusUnprocessableEntity ||
		!strings.Contains(string(msg), "bring it to the table") ||
		e.app.handoffView().State != views.HandoffWaiting {
		t.Fatalf(
			"bring-back upload: status %d, state %q: %s",
			resp.StatusCode,
			e.app.handoffView().State,
			msg,
		)
	}

	resp = phoneUpload(t, target, sampleTakeHome(t, nil))
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("take-home upload: status %d", resp.StatusCode)
	}
	if v := e.app.handoffView(); v.State != views.HandoffReceived || v.Name != "Aria" ||
		v.TakenAt.IsZero() {
		t.Fatalf("view %+v", v)
	}
	// The "Sent" page loads its stylesheet after the transfer, but the transfer is used up.
	if resp, err := phoneGet(
		t,
		target+views.PhoneCSS,
	); err != nil ||
		resp.StatusCode != http.StatusOK {
		t.Errorf("stylesheet after the transfer: %v %v", resp, err)
	}
	if resp := phoneUpload(
		t,
		target,
		sampleTakeHome(t, nil),
	); resp.StatusCode != http.StatusNotFound {
		t.Errorf("second upload: status %d, want 404", resp.StatusCode)
	}
	eventually(
		t,
		"servers closed",
		func() bool { _, err := phoneGet(t, target); return err != nil },
	)

	// The page offers the received file and does not claim there are no characters.
	if body := e.get(
		t,
		route.AtHome,
		"",
	).Body.String(); !strings.Contains(body, "Open Aria") ||
		strings.Contains(body, "No characters here yet") {
		t.Errorf("home page with a received file:\n%s", body)
	}

	tmp := h.file
	if _, err := os.Stat(tmp); err != nil {
		t.Fatalf("temp file gone before the import: %v", err)
	}
	if w := e.postPlayer(
		t,
		route.AtHome,
		url.Values{"handoff": {"1"}},
	); w.Code != http.StatusNoContent {
		t.Fatalf("import from handoff: status %d: %s", w.Code, w.Body)
	}
	if _, err := os.Stat(tmp); !os.IsNotExist(err) {
		t.Errorf("temp file left: %v", err)
	}
	if e.app.handoffView().State != views.HandoffNone {
		t.Error("handoff not cleared after the import")
	}
	if w := e.postPlayer(
		t,
		route.AtHome,
		url.Values{"handoff": {"1"}},
	); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("second import from handoff: status %d, want 422", w.Code)
	}
}

func TestHandoffSend(t *testing.T) {
	shortLinger(t)
	e := newPlayerEnv(t)
	e.app.handoffIPsFn = loopbackIPs
	id := e.importedHome(t)
	w := e.postPlayer(t, route.Path(route.AtHomeSendPhone, id), nil)
	if w.Code != http.StatusOK {
		t.Fatalf("status %d: %s", w.Code, w.Body)
	}
	h := e.app.handoff
	target := h.urls[0]
	if resp, err := phoneGet(t, target); err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("phone page: %v %v", resp, err)
	}
	resp, err := phoneGet(t, target+route.File)
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("download: %v %v", resp, err)
	}
	body := resp.Body
	if int64(len(body)) != h.size {
		t.Fatalf("read %d bytes of %d", len(body), h.size)
	}
	if _, err := export.ReadTakeHome(
		bytes.NewReader(body),
		int64(len(body)),
		export.KindBringBack,
	); err != nil {
		t.Errorf("downloaded file: %v", err)
	}
	eventually(
		t,
		"state sent",
		func() bool { return e.app.handoffView().State == views.HandoffSent },
	)
	eventually(t, "second download refused", func() bool {
		r, err := phoneGet(t, target+route.File)
		return err != nil || r.StatusCode == http.StatusNotFound
	})
}

func TestHandoffExpiresAndReplaces(t *testing.T) {
	e := newPlayerEnv(t)
	old := handoffTTL
	handoffTTL = 50 * time.Millisecond
	t.Cleanup(func() { handoffTTL = old })
	h := e.app.startHandoff(loopbackIPs(), &handoff{})
	eventually(
		t,
		"state closed",
		func() bool { return e.app.handoffView().State == views.HandoffClosed },
	)
	if _, err := phoneGet(t, h.urls[0]); err == nil {
		t.Error("the closed handoff still answers")
	}

	handoffTTL = time.Minute
	first := e.app.startHandoff(loopbackIPs(), &handoff{})
	e.app.startHandoff(loopbackIPs(), &handoff{})
	if first.state != views.HandoffClosed {
		t.Errorf("first handoff state %q after a second started", first.state)
	}
	if _, err := phoneGet(t, first.urls[0]); err == nil {
		t.Error("the replaced handoff still answers")
	}
}

func TestHandoffNoNetwork(t *testing.T) {
	e := newPlayerEnv(t)
	e.app.startHandoff(nil, &handoff{})
	if e.app.handoffView().State != views.HandoffNoNetwork {
		t.Errorf("state %q", e.app.handoffView().State)
	}
}

// homeNote makes a note of character id with one block per visibility and
// returns the note and the block ids.
func (e testEnv) homeNote(
	t *testing.T,
	campaignID, id int64,
	vis ...enum.Visibility,
) (sqlc.Note, []int64) {
	t.Helper()
	ctx := t.Context()
	note, err := e.q().
		CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: campaignID, Kind: enum.NoteGeneric, Title: "Mixed", OwnerCharacterID: nullID(id)})
	if err != nil {
		t.Fatal(err)
	}
	var ids []int64
	for _, v := range vis {
		b, err := e.q().
			CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: v, AuthorCharacterID: nullID(id)})
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, b.ID)
	}
	return note, ids
}

// At home only private blocks can be changed and a note holding shared blocks
// is locked; the table's rules stay as they were.
func TestAtHomeNoteRules(t *testing.T) {
	e := newPlayerEnv(t)
	imp := e.importRequest(t, sampleTakeHome(t, nil), "", nil)
	cookie := homeCookie(imp)
	sess, err := e.mgr.Resolve(t.Context(), cookie.Value)
	if err != nil {
		t.Fatal(err)
	}
	id := sess.CharacterID.Int64
	ch, err := e.q().GetCharacter(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	mixed, blocks := e.homeNote(t, ch.CampaignID, id, enum.VisPlayers, enum.VisOwner)
	shared, private := blocks[0], blocks[1]
	send := func(path, contentType, body string) *httptest.ResponseRecorder {
		req := httptest.NewRequestWithContext(
			t.Context(),
			http.MethodPost,
			path,
			strings.NewReader(body),
		)
		req.Header.Set("Content-Type", contentType) // no HX-Request: a refusal keeps its own status
		req.AddCookie(cookie)
		req.RemoteAddr = loopbackAddr
		w := httptest.NewRecorder()
		e.router.ServeHTTP(w, req)
		return w
	}
	postPlayer := func(path string, form url.Values) *httptest.ResponseRecorder {
		return send(path, "application/x-www-form-urlencoded", form.Encode())
	}
	steps := func(block int64) int {
		return send(
			route.Path(route.BlockSteps, block),
			"application/json",
			`{"version":0,"clientID":"c","steps":[{"stepType":"replace","from":0,"to":0}]}`,
		).Code
	}
	if got := steps(shared); got != http.StatusNotFound {
		t.Errorf("edit shared block: %d, want 404", got)
	}
	if got := steps(private); got == http.StatusNotFound {
		t.Errorf("edit private block: %d, want it allowed", got)
	}
	form := url.Values{"title": {"New"}, "kind": {string(enum.NoteGeneric)}}
	if w := postPlayer(
		route.Path(route.NoteRename, mixed.ID),
		form,
	); w.Code != http.StatusForbidden {
		t.Errorf("rename locked note: %d, want 403", w.Code)
	}
	if w := postPlayer(
		route.Path(route.NoteDelete, mixed.ID),
		nil,
	); w.Code != http.StatusForbidden {
		t.Errorf("delete locked note: %d, want 403", w.Code)
	}
	if w := postPlayer(
		route.Path(route.BlockMove, private),
		url.Values{"dir": {"up"}},
	); w.Code != http.StatusNotFound {
		t.Errorf("reorder locked note: %d, want 404", w.Code)
	}
	if w := postPlayer(
		route.Path(route.NoteBlocks, mixed.ID),
		url.Values{"visibility": {string(enum.VisPlayers)}},
	); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("add shared block: %d, want 422", w.Code)
	}
	if w := postPlayer(
		route.Path(route.NoteBlocks, mixed.ID),
		url.Values{"visibility": {string(enum.VisOwner)}},
	); w.Code >= 400 {
		t.Errorf("add private block: %d: %s", w.Code, w.Body)
	}

	// A note of private blocks only stays fully manageable.
	plain, _ := e.homeNote(t, ch.CampaignID, id, enum.VisOwner)
	if w := postPlayer(route.Path(route.NoteRename, plain.ID), form); w.Code >= 400 {
		t.Errorf("rename private note: %d: %s", w.Code, w.Body)
	}
	if w := postPlayer(route.Path(route.NoteDelete, plain.ID), nil); w.Code >= 400 {
		t.Errorf("delete private note: %d: %s", w.Code, w.Body)
	}
}

// At the table a note with shared blocks is renamed as before.
func TestTableRenameWithSharedBlocks(t *testing.T) {
	e := newTestEnv(t)
	note, err := e.q().
		CreateNote(t.Context(), sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: enum.NoteGeneric, Title: "Plot"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := e.q().
		CreateBlock(t.Context(), sqlc.CreateBlockParams{NoteID: note.ID, Visibility: enum.VisPlayers}); err != nil {
		t.Fatal(err)
	}
	form := url.Values{"title": {"Plot 2"}, "kind": {string(enum.NoteGeneric)}}
	if w := e.postPlayer(t, route.Path(route.NoteRename, note.ID), form); w.Code >= 300 {
		t.Errorf("rename: %d: %s", w.Code, w.Body)
	}
}
