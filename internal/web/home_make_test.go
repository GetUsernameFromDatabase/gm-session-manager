package web

import (
	"bytes"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"testing"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/route"
)

// homeCharacterByName finds a live character on this computer.
func (e testEnv) homeCharacterByName(t *testing.T, name string) sqlc.ListHomeCharactersRow {
	t.Helper()
	rows, err := e.q().ListHomeCharacters(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range rows {
		if r.Name == name {
			return r
		}
	}
	t.Fatalf("no character %q on this computer", name)
	return sqlc.ListHomeCharactersRow{}
}

// homeFileOf reads the file the Save button's download gives for a character.
func (e testEnv) homeFileOf(
	t *testing.T,
	id int64,
	want export.Kind,
) (*export.TakeHomeFile, *httptest.ResponseRecorder) {
	t.Helper()
	w := e.get(t, route.Path(route.AtHomeBringBack, id), "")
	if w.Code != http.StatusOK {
		return nil, w
	}
	body := w.Body.Bytes()
	f, err := export.ReadTakeHome(bytes.NewReader(body), int64(len(body)), want)
	if err != nil {
		t.Fatal(err)
	}
	return f, w
}

func TestHomeMakeCopy(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	q := e.q()
	id := e.importedHome(t)
	orig, err := q.GetCharacter(ctx, id)
	if err != nil {
		t.Fatal(err)
	}

	w := e.postPlayer(t, route.Path(route.AtHomeMakeCopy, id), nil)
	cp := e.homeCharacterByName(t, "Aria (copy)")
	if w.Code != http.StatusNoContent ||
		w.Header().Get("HX-Redirect") != route.Path(route.Character, cp.ID) {
		t.Fatalf("copy: status %d, redirect %q: %s", w.Code, w.Header().Get("HX-Redirect"), w.Body)
	}
	if c := homeCookie(w); c == nil {
		t.Error("the copy did not log the browser in")
	} else if s, err := e.mgr.Resolve(ctx, c.Value); err != nil || s.CharacterID.Int64 != cp.ID {
		t.Errorf("session %+v, %v", s, err)
	}

	// Its own campaign, with the original's sources and links.
	if cp.CampaignID == orig.CampaignID {
		t.Error("the copy shares the original's campaign")
	}
	origSources, _ := q.ListCampaignSources(ctx, orig.CampaignID)
	copySources, _ := q.ListCampaignSources(ctx, cp.CampaignID)
	if len(origSources) == 0 || !slices.Equal(origSources, copySources) {
		t.Errorf("sources %v, want %v", copySources, origSources)
	}
	if hb, err := q.ListCampaignHomeHomebrew(
		ctx,
		cp.CampaignID,
	); err != nil || len(hb) != 1 ||
		hb[0].Slug != "zap" {
		t.Errorf("homebrew seen by the copy: %v, %v", hb, err)
	}
	// The original keeps owning its homebrew; the copy owns none.
	if own, _ := q.ListOwnHomebrew(ctx, nullID(id)); len(own) != 1 {
		t.Errorf("the original owns %d homebrew entries, want 1", len(own))
	}
	if own, _ := q.ListOwnHomebrew(ctx, nullID(cp.ID)); len(own) != 0 {
		t.Errorf("the copy owns %d homebrew entries, want 0", len(own))
	}
	if _, err := q.GetHomeCopy(ctx, cp.ID); !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("the copy has a home copy row: %v", err)
	}

	// Level-ups, looks and notes, with files and attachment rows of their own.
	if n, _ := q.CountLevelUps(ctx, cp.ID); n != 1 {
		t.Errorf("%d level-ups, want 1", n)
	}
	looks, _ := q.ListCharacterImages(ctx, cp.ID)
	if len(looks) != 1 || !looks[0].AttachmentID.Valid {
		t.Fatalf("looks %+v", looks)
	}
	origLooks, _ := q.ListCharacterImages(ctx, id)
	lookAtt, _ := q.GetAttachment(ctx, looks[0].AttachmentID.Int64)
	origAtt, _ := q.GetAttachment(ctx, origLooks[0].AttachmentID.Int64)
	if lookAtt.ID == origAtt.ID || lookAtt.Path == origAtt.Path {
		t.Errorf("the copy's look shares the original's attachment: %+v %+v", lookAtt, origAtt)
	}
	for _, a := range []sqlc.Attachment{lookAtt, origAtt} {
		if _, err := os.Stat(filepath.Join(e.app.attachmentsDir(), a.Path)); err != nil {
			t.Errorf("file of attachment %d: %v", a.ID, err)
		}
	}
	atts, _ := q.ListOwnNoteAttachments(ctx, nullID(cp.ID))
	origAtts, _ := q.ListOwnNoteAttachments(ctx, nullID(id))
	blocks, _ := q.ListOwnNoteBlocks(ctx, nullID(cp.ID))
	if len(atts) != 1 || len(origAtts) != 1 || len(blocks) != 1 ||
		atts[0].Path == origAtts[0].Path {
		t.Fatalf("note attachments %+v (original %+v), blocks %d", atts, origAtts, len(blocks))
	}
	if !strings.Contains(blocks[0].DocJson, route.Path(route.Attachment, atts[0].ID)) ||
		strings.Contains(blocks[0].DocJson, route.Path(route.Attachment, origAtts[0].ID)) {
		t.Errorf("the copy's block does not point at its own attachment: %s", blocks[0].DocJson)
	}

	// Steps past the last snapshot (a note open in another tab) would be lost.
	if _, err := e.conn.ExecContext(
		ctx,
		`UPDATE note_blocks SET version = doc_version + 1 WHERE id = ?`,
		blocks[0].ID,
	); err != nil {
		t.Fatal(err)
	}
	if w := e.postPlayer(
		t,
		route.Path(route.AtHomeMakeCopy, cp.ID),
		nil,
	); w.Code != http.StatusUnprocessableEntity ||
		!strings.Contains(w.Body.String(), "still open in another tab") {
		t.Errorf("copy with an open note: status %d: %s", w.Code, w.Body)
	}

	// A fallback is not copied.
	if w := e.postPlayer(
		t,
		route.Path(route.AtHomeFallback, id),
		nil,
	); w.Code != http.StatusNoContent {
		t.Fatalf("fallback: status %d: %s", w.Code, w.Body)
	}
	fb := e.homeCharacterByName(t, "Aria (fallback)")
	if w := e.postPlayer(
		t,
		route.Path(route.AtHomeMakeCopy, fb.ID),
		nil,
	); w.Code != http.StatusUnprocessableEntity ||
		!strings.Contains(w.Body.String(), "A fallback is not copied") {
		t.Errorf("copy of a fallback: status %d: %s", w.Code, w.Body)
	}
}

func TestHomeFallback(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	q := e.q()
	id := e.importedHome(t)
	orig, _ := q.GetCharacter(ctx, id)

	w := e.postPlayer(t, route.Path(route.AtHomeFallback, id), nil)
	fb := e.homeCharacterByName(t, "Aria (fallback)")
	if w.Code != http.StatusNoContent ||
		w.Header().Get("HX-Redirect") != route.Path(route.Character, fb.ID) {
		t.Fatalf(
			"fallback: status %d, redirect %q: %s",
			w.Code,
			w.Header().Get("HX-Redirect"),
			w.Body,
		)
	}
	if fb.CampaignID != orig.CampaignID {
		t.Error("the fallback is not in its character's campaign")
	}
	hc, _ := q.GetHomeCopy(ctx, id)
	if hc.FallbackCharacterID.Int64 != fb.ID {
		t.Errorf("link %+v, want %d", hc.FallbackCharacterID, fb.ID)
	}
	// The build only: sheet and level-ups.
	if n, _ := q.CountLevelUps(ctx, fb.ID); n != 1 {
		t.Errorf("%d level-ups, want 1", n)
	}
	if looks, _ := q.ListCharacterImages(ctx, fb.ID); len(looks) != 0 {
		t.Errorf("the fallback has %d looks", len(looks))
	}
	if notes, _ := q.ListOwnNotes(ctx, nullID(fb.ID)); len(notes) != 0 {
		t.Errorf("the fallback has %d notes", len(notes))
	}
	if _, err := q.GetHomeCopy(ctx, fb.ID); !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("the fallback has a home copy row: %v", err)
	}
	if w := e.postPlayer(
		t,
		route.Path(route.AtHomeFallback, id),
		nil,
	); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("a second fallback: status %d", w.Code)
	}
	// htmx answers a 404 as a 422 flash.
	if w := e.postPlayer(
		t,
		route.Path(route.AtHomeFallback, fb.ID),
		nil,
	); w.Code != http.StatusUnprocessableEntity ||
		!strings.Contains(w.Body.String(), "Not found") {
		t.Errorf("a fallback of a fallback: status %d: %s", w.Code, w.Body)
	}

	// The bring-back carries it.
	f, _ := e.homeFileOf(t, id, export.KindBringBack)
	if f.Fallback == nil || len(f.Fallback.LevelUps) != 1 || f.Fallback.MaxHP != orig.MaxHp ||
		len(fileSheets(f)) != 4 {
		t.Fatalf("fallback in the file: %+v", f.Fallback)
	}
	// A deleted fallback is left out, and a new one can be made.
	if _, err := e.conn.ExecContext(
		ctx,
		`UPDATE characters SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?`,
		fb.ID,
	); err != nil {
		t.Fatal(err)
	}
	if f, _ := e.homeFileOf(t, id, export.KindBringBack); f.Fallback != nil {
		t.Error("a deleted fallback went into the file")
	}
	if w := e.postPlayer(
		t,
		route.Path(route.AtHomeFallback, id),
		nil,
	); w.Code != http.StatusNoContent {
		t.Fatalf("a new fallback after a deleted one: status %d: %s", w.Code, w.Body)
	}

	// Deleting the character at home takes its fallback with it.
	fb = e.homeCharacterByName(t, "Aria (fallback)")
	fbCookie := homeCookie(
		e.postPlayer(t, route.Path(route.AtHomePlay, fb.ID), nil),
	) // a second browser on the fallback
	req := httptest.NewRequestWithContext(
		ctx,
		http.MethodPost,
		route.Path(route.CharacterRemove, id),
		nil,
	)
	req.Header.Set("HX-Request", "true")
	req.RemoteAddr = loopbackAddr
	req.AddCookie(homeCookie(e.postPlayer(t, route.Path(route.AtHomePlay, id), nil)))
	w = httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Fatalf("delete: status %d: %s", w.Code, w.Body)
	}
	if _, err := q.GetCharacter(ctx, fb.ID); !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("the fallback outlived its character: %v", err)
	}
	if s, err := e.mgr.Resolve(ctx, fbCookie.Value); err == nil {
		t.Errorf("the fallback's device is still logged in: %+v", s)
	}
}

// A replace import takes the fallback with the old campaign, and keeps a test copy.
func TestHomeReplaceDropsFallback(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	q := e.q()
	id := e.importedHome(t)
	e.postPlayer(t, route.Path(route.AtHomeMakeCopy, id), nil)
	e.postPlayer(t, route.Path(route.AtHomeFallback, id), nil)
	cp := e.homeCharacterByName(t, "Aria (copy)")
	fb := e.homeCharacterByName(t, "Aria (fallback)")

	// Without the confirmation the question names the fallback.
	w := e.importRequest(t, sampleTakeHome(t, nil), "", nil)
	if !strings.Contains(w.Body.String(), "fallback build (Aria (fallback)) goes too") {
		t.Errorf("replace question: %s", w.Body)
	}
	if w := e.importRequest(
		t,
		sampleTakeHome(t, nil),
		"",
		map[string]string{formReplace: "1"},
	); w.Code != http.StatusNoContent {
		t.Fatalf("replace: status %d: %s", w.Code, w.Body)
	}
	// By name: SQLite reuses the freed ids.
	rows, err := q.ListHomeCharacters(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, r := range rows {
		names = append(names, r.Name)
	}
	slices.Sort(names)
	if !slices.Equal(names, []string{"Aria", "Aria (copy)"}) {
		t.Errorf(
			"characters after the replace: %v, want the new Aria and the test copy (%d, %d went/stayed)",
			names,
			fb.ID,
			cp.ID,
		)
	}
	if hc, err := q.GetHomeCopy(
		ctx,
		e.homeCharacterByName(t, "Aria").ID,
	); err != nil ||
		hc.FallbackCharacterID.Valid {
		t.Errorf("the new copy: %+v, %v", hc, err)
	}
}

func TestHomeNewCharacter(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	q := e.q()
	id := e.importedHome(t)
	orig, _ := q.GetCharacter(ctx, id)
	e.homeBook(t, "homebook", "2024")
	if _, err := q.AllowCampaignSource(
		ctx,
		sqlc.AllowCampaignSourceParams{CampaignID: orig.CampaignID, SourceKey: "homebook"},
	); err != nil {
		t.Fatal(err)
	}

	form := func(name string, from int64) url.Values {
		return url.Values{"name": {name}, "from": {strconv.FormatInt(from, 10)}}
	}
	for _, bad := range []url.Values{form("   ", orig.CampaignID), form("Nia", 0), form("Nia", 999), form(strings.Repeat("x", 101), orig.CampaignID)} {
		if w := e.postPlayer(t, route.AtHomeNew, bad); w.Code != http.StatusUnprocessableEntity {
			t.Errorf("%v: status %d, want 422", bad, w.Code)
		}
	}
	w := e.postPlayer(t, route.AtHomeNew, form("  Nia ", orig.CampaignID))
	nia := e.homeCharacterByName(t, "Nia")
	if w.Code != http.StatusNoContent ||
		w.Header().Get("HX-Redirect") != wizardPath(nia.ID, wizardSteps[0].Key) {
		t.Fatalf("new: status %d, redirect %q: %s", w.Code, w.Header().Get("HX-Redirect"), w.Body)
	}
	if nia.CampaignID == orig.CampaignID {
		t.Error("the new character shares the table character's campaign")
	}
	if got, _ := q.ListCampaignSources(
		ctx,
		nia.CampaignID,
	); !slices.Contains(got, "homebook") ||
		!slices.Contains(got, "src1") {
		t.Errorf("sources %v", got)
	}
	if _, err := q.GetHomeCopy(ctx, nia.ID); !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("the new character has a home copy row: %v", err)
	}
	// Only a table character's campaign (or a starter's) can be the model.
	if w := e.postPlayer(
		t,
		route.AtHomeNew,
		form("Zed", nia.CampaignID),
	); w.Code != http.StatusUnprocessableEntity ||
		!strings.Contains(w.Body.String(), "Choose a campaign") {
		t.Errorf("model that is not a table character's campaign: status %d: %s", w.Code, w.Body)
	}
	if rows, _ := q.ListHomeCharacters(ctx); len(rows) != 2 {
		t.Errorf("%d characters, want the original and Nia only", len(rows))
	}

	// Its file is a new-character file.
	f, w := e.homeFileOf(t, nia.ID, export.KindNewCharacter)
	if f == nil {
		t.Fatalf("file: status %d: %s", w.Code, w.Body)
	}
	if m := f.Manifest; m.Token != "" || m.CharacterID != 0 || m.MadeAtHome == "" ||
		m.ServerID != "srv" ||
		m.Campaign.Name != "Table" ||
		m.SavedAt.IsZero() {
		t.Errorf("manifest %+v", m)
	}
	if f.Maps != nil || f.Fallback != nil ||
		!strings.Contains(w.Header().Get("Content-Disposition"), "-new-character-") {
		t.Errorf(
			"maps %v, fallback %v, disposition %q",
			f.Maps,
			f.Fallback,
			w.Header().Get("Content-Disposition"),
		)
	}
	if !slices.ContainsFunc(
		f.Sources,
		func(s export.TakeHomeSource) bool { return s.SourceKey == "homebook" },
	) {
		t.Errorf("sources %+v lack the home source", f.Sources)
	}
	if !slices.ContainsFunc(
		f.Homebrew,
		func(h export.TakeHomeEntry) bool { return h.Slug == "zap" },
	) {
		t.Errorf("homebrew %+v", f.Homebrew)
	}

	// The table character still gets a bring-back; a fallback has no file of its own.
	if f, _ := e.homeFileOf(t, id, export.KindBringBack); f.Manifest.Kind != export.KindBringBack {
		t.Errorf("kind %q", f.Manifest.Kind)
	}
	e.postPlayer(t, route.Path(route.AtHomeFallback, id), nil)
	fb := e.homeCharacterByName(t, "Aria (fallback)")
	if _, w := e.homeFileOf(
		t,
		fb.ID,
		export.KindBringBack,
	); w.Code != http.StatusUnprocessableEntity ||
		!strings.Contains(w.Body.String(), "goes back inside Aria") {
		t.Errorf("file of a fallback: status %d: %s", w.Code, w.Body)
	}
}

// The page offers a fallback only when the build pins a source added at home,
// and lists copies apart from table characters.
func TestHomePageFallbackOffer(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	id := e.importedHome(t)
	const offer = "Prepare a fallback"
	if body := e.get(
		t,
		route.AtHome,
		"",
	).Body.String(); strings.Contains(body, offer) || strings.Contains(body, "Only on this computer") || !strings.Contains(body, "From the table") ||
		!strings.Contains(body, "New character") {
		t.Errorf("home page without a home source:\n%s", body)
	}

	e.homeBook(t, "homebook", "2024")
	// A spell from the home source is not part of the build.
	if _, err := e.conn.ExecContext(
		ctx,
		`UPDATE characters SET sheet_json = json_set(sheet_json, '$.spells', json('[{"name":"Zap","slug":"zap","level":1,"source":"homebook"}]')) WHERE id = ?`,
		id,
	); err != nil {
		t.Fatal(err)
	}
	if body := e.get(t, route.AtHome, "").Body.String(); strings.Contains(body, offer) {
		t.Error("a fallback is offered for a spell")
	}
	if _, err := e.conn.ExecContext(
		ctx,
		`UPDATE characters SET sheet_json = json_set(sheet_json, '$.species_source', 'homebook') WHERE id = ?`,
		id,
	); err != nil {
		t.Fatal(err)
	}
	if body := e.get(
		t,
		route.AtHome,
		"",
	).Body.String(); !strings.Contains(body, offer) ||
		!strings.Contains(body, "uses homebook, added here") {
		t.Errorf("no offer for a build that uses a home source:\n%s", body)
	}

	w := e.postPlayer(t, route.Path(route.AtHomeFallback, id), nil)
	body := e.get(t, route.AtHome, "").Body.String()
	if strings.Contains(body, offer) || !strings.Contains(body, "Fallback build") ||
		strings.Count(body, "Aria (fallback)") != 1 {
		t.Errorf("home page with a fallback:\n%s", body)
	}
	e.postPlayer(t, route.Path(route.AtHomeMakeCopy, id), nil)
	if body := e.get(
		t,
		route.AtHome,
		"",
	).Body.String(); !strings.Contains(body, "Only on this computer") ||
		!strings.Contains(body, "Save as new character file") {
		t.Errorf("home page with a copy:\n%s", body)
	}

	// The fallback's sheet says what it is.
	fb := e.homeCharacterByName(t, "Aria (fallback)")
	sheet := e.get(t, route.Path(route.Character, fb.ID), homeCookie(w).Value).Body.String()
	if !strings.Contains(sheet, "Fallback build for Aria.") ||
		!strings.Contains(sheet, "still uses homebook") {
		t.Errorf("fallback sheet lacks its banner:\n%s", sheet)
	}
}

func TestHomeName(t *testing.T) {
	long := strings.Repeat("é", 120)
	for _, tt := range []struct{ base, suffix, want string }{
		{"Aria", copySuffix, "Aria (copy)"},
		{long, copySuffix, strings.Repeat("é", maxCharacterName-len([]rune(copySuffix))) + copySuffix},
		{strings.Repeat("x", 95), fallbackSuffix, strings.Repeat("x", maxCharacterName-len(fallbackSuffix)) + fallbackSuffix},
	} {
		got := homeName(tt.base, tt.suffix)
		if got != tt.want || len([]rune(got)) > maxCharacterName {
			t.Errorf("homeName(%.10q…, %q) = %q", tt.base, tt.suffix, got)
		}
	}
}
