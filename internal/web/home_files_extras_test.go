package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

const (
	// Private text and titles the GM's pages must never show.
	secretEdit  = "HOME-EDIT-SECRET"
	secretBlock = "NEW-BLOCK-SECRET"
	secretTitle = "PRIVATE-TITLE"
)

func (e testEnv) count(t *testing.T, query string, args ...any) int {
	t.Helper()
	var n int
	if err := e.conn.QueryRowContext(t.Context(), query, args...).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func mentionDoc(text string, ids ...string) json.RawMessage {
	nodes := []string{fmt.Sprintf(`{"type":"text","text":%q}`, text)}
	for _, id := range ids {
		nodes = append(nodes, `{"type":"mention","attrs":{"id":"`+id+`"}}`)
	}
	return json.RawMessage(`{"type":"doc","content":[{"type":"paragraph","content":[` + strings.Join(nodes, ",") + `]}]}`)
}

// extrasEnv is a take-out of the seeded character and the table ids a
// bring-back's maps point at.
type extrasEnv struct {
	testEnv
	rec                      sqlc.TakeHome
	seed                     takeHomeSeed
	hood, diary, zap         int64
	ownerBlock, playersBlock int64
}

func newExtrasEnv(t *testing.T) extrasEnv {
	t.Helper()
	e := newTestEnv(t)
	x := extrasEnv{testEnv: e, seed: e.seedTakeHome(t)}
	x.rec = e.takeOut(t)
	ctx := t.Context()
	looks, err := e.q().ListCharacterImages(ctx, e.characterID)
	if err != nil || len(looks) != 1 {
		t.Fatalf("looks %v (%v)", looks, err)
	}
	notes, err := e.q().ListOwnNotes(ctx, nullID(e.characterID))
	if err != nil || len(notes) != 1 {
		t.Fatalf("notes %v (%v)", notes, err)
	}
	hb, err := e.q().ListOwnHomebrew(ctx, nullID(e.characterID))
	if err != nil || len(hb) != 1 {
		t.Fatalf("homebrew %v (%v)", hb, err)
	}
	x.hood, x.diary, x.zap = looks[0].ID, notes[0].ID, hb[0].ID
	x.ownerBlock, x.playersBlock = x.seed.ownerBlock, x.seed.playersBlock
	return x
}

// changes is the bring-back of a player who hid the hood, added a cloak and made
// it current, rewrote a private block, added a block and a note, and changed
// and added homebrew.
func (x extrasEnv) changes() func(*export.TakeHome) {
	return func(th *export.TakeHome) {
		th.Looks = []export.TakeHomeLook{
			{ImageID: 501, Name: "Hood", URL: "https://example.com/hood.png"},
			{ImageID: 502, Name: "Cloak", URL: "https://example.com/cloak.png", IsRevealed: true, IsCurrent: true},
		}
		edit := mentionDoc(secretEdit, fmt.Sprintf("note:-%d", x.diary), "note:602", "image:502", "image:777")
		th.Notes = []export.TakeHomeNote{
			{ID: 601, Kind: enum.NoteGeneric, Title: "Diary", Blocks: []export.TakeHomeBlock{
				{ID: 701, Position: 1, Visibility: enum.VisOwner, Doc: edit, Version: 2, Text: secretEdit},
				{ID: 702, Position: 2, Visibility: enum.VisPlayers, Doc: json.RawMessage(`{"type":"doc"}`), Version: 0},
				{ID: 703, Position: 3, Visibility: enum.VisPlayers, Doc: mentionDoc(secretBlock), Text: secretBlock},
			}},
			{ID: 602, Kind: enum.NoteGeneric, Title: secretTitle, Blocks: []export.TakeHomeBlock{{ID: 704, Visibility: enum.VisOwner, Doc: mentionDoc("n")}}},
		}
		th.Homebrew = []export.TakeHomeEntry{
			{ID: 801, Kind: enum.KindSpell, Slug: "zap", Name: "Zap!", SourceKey: homebrewSource, Data: json.RawMessage(`{}`), ContentHash: "h2"},
			{ID: 802, Kind: enum.KindSpell, Slug: "newt", Name: "Newt", SourceKey: homebrewSource, Data: json.RawMessage(`{}`), ContentHash: "n1"},
		}
		th.Maps = &export.HomeIDMaps{
			Images: map[int64]int64{501: x.hood}, Notes: map[int64]int64{601: x.diary},
			Blocks: map[int64]int64{701: x.ownerBlock, 702: x.playersBlock}, Homebrew: map[int64]int64{801: x.zap},
		}
	}
}

func TestBringBackAcceptExtras(t *testing.T) {
	x := newExtrasEnv(t)
	ctx := t.Context()
	q := x.q()
	gm := x.token(t, enum.RoleAdmin)
	// An editor left steps for the private block; the bring-back replaces the document.
	if _, err := x.conn.ExecContext(ctx, `INSERT INTO block_steps (block_id, version, client_id, step_json) VALUES (?, 1, 'c', '{}')`, x.ownerBlock); err != nil {
		t.Fatal(err)
	}
	if w := x.uploadBringBack(t, x.bringBackZip(t, x.rec, x.changes()), gm); w.Code != http.StatusNoContent {
		t.Fatalf("upload: %d %s", w.Code, w.Body)
	}
	f, was := x.waitingFile(t)
	review := route.Path(route.CampaignHomeFile, x.campaignID, f.id)

	// The GM sees names of looks and homebrew and counts of notes, never private text.
	page := x.do(t, http.MethodGet, review, gm, nil, false).Body.String()
	for _, want := range []string{"Hood: hidden", "New: Cloak", "Current: Cloak", "1 changed · 1 new", "Changed: Zap!", "New: Newt"} {
		if !strings.Contains(page, want) {
			t.Errorf("review lacks %q", want)
		}
	}
	for _, secret := range []string{secretEdit, secretBlock, secretTitle} {
		if strings.Contains(page, secret) {
			t.Errorf("the review shows private text %q", secret)
		}
	}

	w := x.do(t, http.MethodPost, review+route.Accept, gm, url.Values{formWas: {was}}, true)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "is back at the table") {
		t.Fatalf("accept: %d %s", w.Code, w.Body)
	}

	// Looks.
	looks, err := q.ListCharacterImages(ctx, x.characterID)
	if err != nil || len(looks) != 2 {
		t.Fatalf("looks %v (%v)", looks, err)
	}
	var cloak int64
	for _, l := range looks {
		switch l.Name {
		case "Hood":
			if l.IsRevealed || l.IsCurrent {
				t.Errorf("hood %+v: want hidden and not current", l)
			}
		case "Cloak":
			cloak = l.ID
			if !l.IsRevealed || !l.IsCurrent {
				t.Errorf("cloak %+v: want revealed and current", l)
			}
		}
	}

	// Notes: the private block is the file's, its steps gone; the shared one is untouched.
	b, err := q.GetBlock(ctx, x.ownerBlock)
	if err != nil {
		t.Fatal(err)
	}
	var newNote sqlc.Note
	notes, _ := q.ListOwnNotes(ctx, nullID(x.characterID))
	for _, n := range notes {
		if n.Title == secretTitle {
			newNote = n
		}
	}
	if newNote.ID == 0 || len(notes) != 2 {
		t.Fatalf("notes %v: the new note is missing", notes)
	}
	for _, want := range []string{secretEdit, fmt.Sprintf(`"note:%d"`, x.diary), fmt.Sprintf(`"note:%d"`, newNote.ID), fmt.Sprintf(`"image:%d"`, cloak), `"image:0"`} {
		if !strings.Contains(b.DocJson, want) {
			t.Errorf("the replaced block %s lacks %s", b.DocJson, want)
		}
	}
	if b.Version != 1 || b.DocVersion != 1 || x.count(t, `SELECT count(*) FROM block_steps WHERE block_id = ?`, x.ownerBlock) != 0 {
		t.Errorf("version %d/%d, steps %d: want a new version and no steps", b.Version, b.DocVersion, x.count(t, `SELECT count(*) FROM block_steps WHERE block_id = ?`, x.ownerBlock))
	}
	if pb, _ := q.GetBlock(ctx, x.playersBlock); strings.Contains(pb.DocJson, secretEdit) || pb.Version != 3 {
		t.Errorf("the shared block changed: %+v", pb)
	}
	if n := x.count(t, `SELECT count(*) FROM note_blocks WHERE note_id = ? AND visibility = 'owner' AND author_character_id = ? AND text = ?`, x.diary, x.characterID, secretBlock); n != 1 {
		t.Errorf("%d private blocks with the new text in the diary, want 1 (forced private)", n)
	}
	if n := x.count(t, `SELECT count(*) FROM note_blocks WHERE note_id = ? AND visibility = 'owner'`, newNote.ID); n != 1 {
		t.Errorf("the new note has %d private blocks, want 1", n)
	}
	// Links rebuilt from the rewritten document.
	for _, target := range []struct{ col, what string }{{"target_note_id", fmt.Sprint(x.diary)}, {"target_note_id", fmt.Sprint(newNote.ID)}, {"target_image_id", fmt.Sprint(cloak)}} {
		if n := x.count(t, `SELECT count(*) FROM note_links WHERE source_block_id = ? AND `+target.col+` = ?`, x.ownerBlock, target.what); n != 1 {
			t.Errorf("%d links from the block to %s %s, want 1", n, target.col, target.what)
		}
	}

	// Homebrew.
	zap, err := q.GetCompendiumEntry(ctx, x.zap)
	if err != nil || zap.ContentHash != "h2" || zap.Name != "Zap!" {
		t.Errorf("zap %+v (%v)", zap, err)
	}
	newt, err := q.GetCompendiumEntryBySlugSource(ctx, sqlc.GetCompendiumEntryBySlugSourceParams{Kind: enum.KindSpell, Slug: "newt", SourceKey: homebrewSource})
	if err != nil || newt.OwnerCharacterID.Int64 != x.characterID || !newt.Scoped {
		t.Errorf("newt %+v (%v)", newt, err)
	}
	if x.count(t, `SELECT count(*) FROM compendium_campaigns WHERE entry_id = ? AND campaign_id = ?`, newt.ID, x.campaignID) != 1 {
		t.Error("the new homebrew is not linked to the campaign")
	}
	if x.fileCount(t) != 0 {
		t.Error("the zip was kept")
	}
}

func (e testEnv) attachmentFileCount(t *testing.T) int {
	t.Helper()
	entries, err := os.ReadDir(e.app.attachmentsDir())
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		t.Fatal(err)
	}
	return len(entries)
}

// TestBringBackAcceptFiles: a new look's picture and a picture in a replaced
// private block are saved before the transaction, owned by their rows and
// pointed at; a stale Accept leaves no file behind.
func TestBringBackAcceptFiles(t *testing.T) {
	x := newExtrasEnv(t)
	ctx := t.Context()
	q := x.q()
	gm := x.token(t, enum.RoleAdmin)
	src := t.TempDir()
	for _, name := range []string{"cloak.png", "map.png"} {
		if err := os.WriteFile(filepath.Join(src, name), []byte("\x89PNG\r\n\x1a\n"+name), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	const homeAttachment = 901
	doc := json.RawMessage(fmt.Sprintf(`{"type":"doc","content":[{"type":"image","attrs":{"src":%q}}]}`, route.Path(route.Attachment, homeAttachment)))
	body := x.bringBackZipFiles(t, x.rec, func(th *export.TakeHome) {
		th.Looks = []export.TakeHomeLook{
			{ImageID: 501, Name: "Hood", URL: "https://example.com/hood.png"},
			{ImageID: 502, Name: "Cloak", File: export.ZipFile("cloak.png"), FileName: "cloak.png"},
		}
		th.Notes = []export.TakeHomeNote{{ID: 601, Kind: enum.NoteGeneric, Title: "Diary", Blocks: []export.TakeHomeBlock{
			{ID: 701, Position: 1, Visibility: enum.VisOwner, Doc: doc, Version: 2, Files: []export.TakeHomeBlockFile{{AttachmentID: homeAttachment, Name: "map.png", File: export.ZipFile("map.png")}}},
			{ID: 702, Position: 2, Visibility: enum.VisPlayers, Doc: json.RawMessage(`{"type":"doc"}`)},
		}}}
		th.Maps = &export.HomeIDMaps{Images: map[int64]int64{501: x.hood}, Notes: map[int64]int64{601: x.diary}, Blocks: map[int64]int64{701: x.ownerBlock, 702: x.playersBlock}}
	}, src)
	before := x.attachmentFileCount(t)
	if w := x.uploadBringBack(t, body, gm); w.Code != http.StatusNoContent {
		t.Fatalf("upload: %d %s", w.Code, w.Body)
	}
	f, was := x.waitingFile(t)
	accept := route.Path(route.CampaignHomeFile, x.campaignID, f.id) + route.Accept

	if w := x.do(t, http.MethodPost, accept, gm, url.Values{formWas: {"stale"}}, true); !strings.Contains(w.Body.String(), "review again") || x.attachmentFileCount(t) != before {
		t.Fatalf("stale accept: %d %s, %d files (want %d)", w.Code, w.Body, x.attachmentFileCount(t), before)
	}
	if w := x.do(t, http.MethodPost, accept, gm, url.Values{formWas: {was}}, true); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "is back at the table") {
		t.Fatalf("accept: %d %s", w.Code, w.Body)
	}
	if got := x.attachmentFileCount(t); got != before+2 {
		t.Errorf("%d files, want %d: the two pictures, nothing else", got, before+2)
	}

	looks, err := q.ListCharacterImages(ctx, x.characterID)
	if err != nil {
		t.Fatal(err)
	}
	var cloak sqlc.ListCharacterImagesRow
	for _, l := range looks {
		if l.Name == "Cloak" {
			cloak = l
		}
	}
	if !cloak.AttachmentID.Valid {
		t.Fatalf("cloak %+v has no picture", cloak)
	}
	att, err := q.GetAttachment(ctx, cloak.AttachmentID.Int64)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(x.app.attachmentsDir(), att.Path)); err != nil {
		t.Errorf("the cloak's file: %v", err)
	}

	var mapID int64
	if err := x.conn.QueryRowContext(ctx, `SELECT id FROM attachments WHERE block_id = ? AND name = 'map.png'`, x.ownerBlock).Scan(&mapID); err != nil {
		t.Fatalf("the block's picture: %v", err)
	}
	if b, err := q.GetBlock(ctx, x.ownerBlock); err != nil || !strings.Contains(b.DocJson, route.Path(route.Attachment, mapID)) {
		t.Errorf("the replaced block %s does not show attachment %d (%v)", b.DocJson, mapID, err)
	}
}

// TestBringBackStaleExtras: a change at the table to what the plan read
// (a look, a private block, homebrew) makes Accept ask for a new review.
func TestBringBackStaleExtras(t *testing.T) {
	tests := []struct {
		name   string
		change func(*testing.T, extrasEnv)
	}{
		{"look flag", func(t *testing.T, x extrasEnv) {
			if err := x.q().ToggleCharacterImageRevealed(t.Context(), sqlc.ToggleCharacterImageRevealedParams{CharacterID: x.characterID, ImageID: x.hood}); err != nil {
				t.Fatal(err)
			}
		}},
		{"private block", func(t *testing.T, x extrasEnv) {
			if _, err := x.conn.ExecContext(t.Context(), `UPDATE note_blocks SET version = version + 1 WHERE id = ?`, x.ownerBlock); err != nil {
				t.Fatal(err)
			}
		}},
		{"note title", func(t *testing.T, x extrasEnv) {
			if err := x.q().RenameNote(t.Context(), sqlc.RenameNoteParams{Title: "Renamed", Kind: enum.NoteGeneric, ID: x.diary}); err != nil {
				t.Fatal(err)
			}
		}},
		{"homebrew", func(t *testing.T, x extrasEnv) {
			if err := x.q().UpdateHomebrew(t.Context(), sqlc.UpdateHomebrewParams{Name: "Zap", DataJson: "{}", ContentHash: "h9", ID: x.zap}); err != nil {
				t.Fatal(err)
			}
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := newExtrasEnv(t)
			gm := x.token(t, enum.RoleAdmin)
			if w := x.uploadBringBack(t, x.bringBackZip(t, x.rec, x.changes()), gm); w.Code != http.StatusNoContent {
				t.Fatal(w.Body)
			}
			f, was := x.waitingFile(t)
			tt.change(t, x)
			w := x.do(t, http.MethodPost, route.Path(route.CampaignHomeFileAccept, x.campaignID, f.id), gm, url.Values{formWas: {was}}, true)
			if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "changed at the table meanwhile") {
				t.Fatalf("status %d: %s", w.Code, w.Body)
			}
			if x.count(t, `SELECT count(*) FROM images WHERE name = 'Cloak'`) != 0 {
				t.Error("a stale accept wrote a look")
			}
		})
	}
}

// newCharacterZip is a new-character file as the player app writes it.
func (e testEnv) newCharacterZip(t *testing.T, mutate func(*export.TakeHome)) []byte {
	t.Helper()
	return e.newCharacterZipFiles(t, mutate, t.TempDir())
}

// newCharacterZipFiles is newCharacterZip with the files the looks and blocks
// name read from dir.
func (e testEnv) newCharacterZipFiles(t *testing.T, mutate func(*export.TakeHome), dir string) []byte {
	t.Helper()
	serverID, err := e.q().GetServerID(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, 11, 25, 9, 0, 0, 0, time.UTC)
	th := export.TakeHome{
		Manifest: export.TakeHomeManifest{
			Format: export.TakeHomeFormat, Kind: export.KindNewCharacter, AppVersion: "v-home", ServerID: serverID,
			Campaign: export.TakeHomeCampaign{Name: "Home", System: gameSystem, Edition: gameEdition}, CreatedAt: now, SavedAt: now,
		},
		Character: export.CharacterFile{Name: "Bryn", System: gameSystem, HP: 8, MaxHP: 8, Sheet: homeSheet(t, nil)},
		Looks:     []export.TakeHomeLook{{ImageID: 1, Name: "Mask", URL: "https://example.com/mask.png", IsRevealed: true, IsCurrent: true}},
		Notes: []export.TakeHomeNote{
			{ID: 1, Kind: enum.NoteGeneric, Title: secretTitle, Blocks: []export.TakeHomeBlock{
				{ID: 1, Visibility: enum.VisOwner, Doc: mentionDoc(secretEdit, "image:1"), Text: secretEdit},
				{ID: 2, Visibility: enum.VisPlayers, Doc: mentionDoc(secretBlock), Text: secretBlock},
			}},
			{ID: 2, Kind: enum.NoteGeneric, Title: "Empty"},
		},
		Homebrew: []export.TakeHomeEntry{{ID: 1, Kind: enum.KindSpell, Slug: "glow-3", Name: "Glow", SourceKey: homebrewSource, Data: json.RawMessage(`{}`), ContentHash: "g1"}},
	}
	if mutate != nil {
		mutate(&th)
	}
	var buf bytes.Buffer
	if err := export.WriteTakeHome(&buf, th, dir); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestNewCharacterFromGM(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)
	add := route.Path(route.CampaignNewCharacter, e.campaignID)

	if w := e.do(t, http.MethodGet, add, gm, nil, true); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Add a character made at home") {
		t.Errorf("dialog: %d %s", w.Code, w.Body)
	}
	if w := e.doFile(t, add, player, nil, formFile, "x.zip", string(e.newCharacterZip(t, nil))); w.Code != http.StatusForbidden {
		t.Errorf("a player's upload: %d", w.Code)
	}
	if w := e.doFile(t, add, gm, nil, formFile, "x.zip", string(e.bringBackZip(t, e.takeOut(t), nil))); w.Code != http.StatusUnprocessableEntity || !strings.Contains(w.Body.String(), "not a new character file") {
		t.Errorf("a bring-back file: %d %s", w.Code, w.Body)
	}
	if w := e.doFile(t, add, gm, nil, formFile, "x.zip", string(e.newCharacterZip(t, func(th *export.TakeHome) { th.Manifest.ServerID = "elsewhere" }))); w.Code != http.StatusUnprocessableEntity || !strings.Contains(w.Body.String(), "another GM") {
		t.Errorf("another table: %d %s", w.Code, w.Body)
	}
	if len(e.app.homeFiles.list()) != 0 || e.fileCount(t) != 0 {
		t.Fatal("a refused file was kept")
	}

	// The look brings a picture: Add reads it from the zip.
	src := t.TempDir()
	if err := os.WriteFile(filepath.Join(src, "mask.png"), []byte("\x89PNG\r\n\x1a\nmask"), 0o600); err != nil {
		t.Fatal(err)
	}
	withPicture := e.newCharacterZipFiles(t, func(th *export.TakeHome) {
		th.Looks[0].File, th.Looks[0].FileName = export.ZipFile("mask.png"), "mask.png"
	}, src)
	attachmentsBefore := e.attachmentFileCount(t)
	w := e.doFile(t, add, gm, nil, formFile, "x.zip", string(withPicture))
	if w.Code != http.StatusSeeOther {
		t.Fatalf("upload: %d %s", w.Code, w.Body)
	}
	list := e.app.homeFiles.list()
	if len(list) != 1 || list[0].characterID != 0 || list[0].byPlayer || w.Header().Get("Location") != route.Path(route.CampaignHomeFile, e.campaignID, list[0].id) {
		t.Fatalf("files %+v, location %q", list, w.Header().Get("Location"))
	}
	f := list[0]
	// A second file for no character does not replace the first.
	if w := e.doFile(t, add, gm, nil, formFile, "x.zip", string(e.newCharacterZip(t, nil))); w.Code != http.StatusSeeOther || len(e.app.homeFiles.list()) != 2 {
		t.Errorf("second upload: %d, %d files", w.Code, len(e.app.homeFiles.list()))
	}

	// The popup card; the GM's own file is a card with Review and Reject.
	var html bytes.Buffer
	if err := views.PendingRequests(e.app.pendingViews()).Render(ctx, &html); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(html.String(), "Bryn, made at home") || strings.Contains(html.String(), secretTitle) {
		t.Errorf("popup: %s", html.String())
	}

	review := route.Path(route.CampaignHomeFile, e.campaignID, f.id)
	page := e.do(t, http.MethodGet, review, gm, nil, false).Body.String()
	for _, want := range []string{"Bryn", "Made at home · level 1 · uploaded by the GM", "Add to campaign", "New: Mask", "New: Glow", "1 note"} {
		if !strings.Contains(page, want) {
			t.Errorf("review lacks %q", want)
		}
	}
	for _, secret := range []string{secretEdit, secretBlock, secretTitle} {
		if strings.Contains(page, secret) {
			t.Errorf("the review shows private text %q", secret)
		}
	}

	rv, err := e.app.reviewHomeFile(ctx, e.q(), f, func(string) string { return "" })
	if err != nil {
		t.Fatal(err)
	}
	w = e.do(t, http.MethodPost, review+route.Accept, gm, url.Values{formWas: {rv.was}}, true)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Bryn joined the campaign.") {
		t.Fatalf("add: %d %s", w.Code, w.Body)
	}
	chars, _ := e.q().ListCharactersByCampaign(ctx, e.campaignID)
	var bryn sqlc.Character
	for _, c := range chars {
		if c.Name == "Bryn" {
			bryn = c
		}
	}
	if bryn.ID == 0 {
		t.Fatalf("characters %v", chars)
	}
	notes, _ := e.q().ListOwnNotes(ctx, nullID(bryn.ID))
	looks, _ := e.q().ListCharacterImages(ctx, bryn.ID)
	hb, _ := e.q().ListOwnHomebrew(ctx, nullID(bryn.ID))
	if len(notes) != 1 || notes[0].Title != secretTitle || len(looks) != 1 || !looks[0].IsCurrent || len(hb) != 1 || hb[0].Slug != "glow-3" {
		t.Fatalf("notes %v looks %v homebrew %v", notes, looks, hb)
	}
	if !looks[0].AttachmentID.Valid || e.attachmentFileCount(t) != attachmentsBefore+1 {
		t.Errorf("look %+v, %d files (want %d): the picture was not saved", looks[0], e.attachmentFileCount(t), attachmentsBefore+1)
	}
	blocks, _ := e.q().ListOwnNoteBlocks(ctx, nullID(bryn.ID))
	if len(blocks) != 1 || blocks[0].Visibility != enum.VisOwner || !strings.Contains(blocks[0].DocJson, fmt.Sprintf(`"image:%d"`, looks[0].ID)) {
		t.Errorf("blocks %+v: want the private one only, its look mention rewritten", blocks)
	}
	if e.count(t, `SELECT count(*) FROM note_links WHERE source_block_id = ? AND target_image_id = ?`, blocks[0].ID, looks[0].ID) != 1 {
		t.Error("the block's link to the look was not written")
	}
	if len(e.app.homeFiles.list()) != 1 {
		t.Error("the added file was kept")
	}
}

// A new character may pin books of this table that the file does not list:
// the review offers those the campaign lacks, and Allow turns them on.
func TestNewCharacterTableSources(t *testing.T) {
	tests := []struct {
		name     string
		known    bool // the table has the book
		allowed  bool // the campaign allows it
		pinned   bool // the sheet uses it
		listed   bool // the file lists it (added at home)
		wantCode int
		wantRow  bool
	}{
		{name: "used, listed, known", known: true, pinned: true, listed: true, wantCode: http.StatusSeeOther, wantRow: true},
		{name: "used, listed, unknown to the table", pinned: true, listed: true, wantCode: http.StatusSeeOther, wantRow: true},
		{name: "used, not allowed", known: true, pinned: true, wantCode: http.StatusSeeOther, wantRow: true},
		{name: "used, allowed", known: true, allowed: true, pinned: true, wantCode: http.StatusSeeOther},
		{name: "not used", known: true, wantCode: http.StatusSeeOther},
		{name: "used, unknown to the table", pinned: true, wantCode: http.StatusUnprocessableEntity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := newTestEnv(t)
			ctx := t.Context()
			q := e.q()
			gm := e.token(t, enum.RoleAdmin)
			if tt.known {
				if err := q.UpsertCompendiumSource(ctx, sqlc.UpsertCompendiumSourceParams{SourceKey: "tbook", Path: "books", Edition: "2024"}); err != nil {
					t.Fatal(err)
				}
			}
			if tt.allowed {
				if _, err := q.AllowCampaignSource(ctx, sqlc.AllowCampaignSourceParams{CampaignID: e.campaignID, SourceKey: "tbook"}); err != nil {
					t.Fatal(err)
				}
			}
			body := e.newCharacterZip(t, func(th *export.TakeHome) {
				if tt.pinned {
					th.Character.Sheet = homeSheet(t, func(s *dnd5e.Sheet) {
						s.Spells = []dnd5e.Spell{{Slug: "zap", Name: "Zap", Source: "tbook"}}
					})
				}
				if tt.listed {
					th.Sources = []export.TakeHomeSource{{SourceKey: "tbook", Edition: "2024"}}
					th.Entries = []export.TakeHomeEntry{{Kind: enum.KindSpell, Slug: "zap", Name: "Zap", SourceKey: "tbook", Data: json.RawMessage(`{}`), ContentHash: "z"}}
				}
			})
			w := e.doFile(t, route.Path(route.CampaignNewCharacter, e.campaignID), gm, nil, formFile, "x.zip", string(body))
			if w.Code != tt.wantCode {
				t.Fatalf("upload: %d, want %d: %s", w.Code, tt.wantCode, w.Body)
			}
			if w.Code != http.StatusSeeOther {
				if !strings.Contains(w.Body.String(), "this table does not have") || len(e.app.homeFiles.list()) != 0 {
					t.Errorf("refusal %q, %d files kept", w.Body, len(e.app.homeFiles.list()))
				}
				return
			}
			f := e.app.homeFiles.list()[0]
			form := url.Values{}
			rv, err := e.app.reviewHomeFile(ctx, q, f, form.Get)
			if err != nil {
				t.Fatal(err)
			}
			if got := len(rv.plan.Sources) == 1 && rv.plan.Sources[0].Key == "tbook"; got != tt.wantRow {
				t.Fatalf("sources %+v, want a tbook row: %v", rv.plan.Sources, tt.wantRow)
			}
			if !tt.wantRow {
				return
			}
			if rv.plan.Block == "" {
				t.Error("Add is not blocked before the source is chosen")
			}
			form.Set(views.SourceParam("tbook"), views.SourceAllow)
			if rv, err = e.app.reviewHomeFile(ctx, q, f, form.Get); err != nil || rv.plan.Block != "" {
				t.Fatalf("allowed: block %q, %v", rv.plan.Block, err)
			}
			form.Set(formWas, rv.was)
			if w := e.do(t, http.MethodPost, route.Path(route.CampaignHomeFile, e.campaignID, f.id)+route.Accept, gm, form, true); w.Code != http.StatusOK {
				t.Fatalf("add: %d %s", w.Code, w.Body)
			}
			if got, _ := q.ListCampaignSources(ctx, e.campaignID); !slices.Contains(got, "tbook") {
				t.Errorf("campaign sources %v lack tbook", got)
			}
		})
	}
}

func (e testEnv) visitorUpload(t *testing.T, zipBody []byte) *httptest.ResponseRecorder {
	t.Helper()
	return e.doFile(t, route.Path(route.JoinNewCharacter, e.campaignID), "", url.Values{"player": {"Pat"}}, formFile, "x.zip", string(zipBody))
}

func TestNewCharacterFromVisitor(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, enum.RoleAdmin)

	if w := e.visitorUpload(t, []byte("not a zip")); w.Code != http.StatusUnprocessableEntity || !strings.Contains(w.Body.String(), "not a new character file") {
		t.Errorf("bad file: %d %s", w.Code, w.Body)
	}
	if w := e.doFile(t, route.Path(route.JoinNewCharacter, e.campaignID), "", nil, formFile, "x.zip", string(e.newCharacterZip(t, nil))); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("no player name: %d", w.Code)
	}
	if w := e.doFile(t, route.Path(route.JoinNewCharacter, e.campaignID), e.token(t, enum.RolePlayer), url.Values{"player": {"Pat"}}, formFile, "x.zip", string(e.newCharacterZip(t, nil))); w.Code != http.StatusSeeOther {
		t.Errorf("a device with a role: %d, want a redirect home", w.Code)
	}
	if len(e.app.homeFiles.list()) != 0 || e.fileCount(t) != 0 {
		t.Fatal("a refused file was kept")
	}

	// Even with open join the file waits for the GM.
	if err := e.q().SetCampaignOpenJoin(ctx, sqlc.SetCampaignOpenJoinParams{OpenJoin: true, ID: e.campaignID}); err != nil {
		t.Fatal(err)
	}
	if w := e.visitorUpload(t, e.newCharacterZip(t, nil)); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Waiting for the GM") {
		t.Fatalf("upload: %d %s", w.Code, w.Body)
	}
	pending := e.app.Hub.PendingList(e.campaignID)
	if len(pending) != 1 || pending[0].Kind != hub.PendingNewCharacter || pending[0].Name != "Bryn" || pending[0].FileID == "" {
		t.Fatalf("pending %+v", pending)
	}
	p := pending[0]
	f, ok := e.app.homeFiles.get(p.FileID, e.campaignID)
	if !ok || !f.byPlayer {
		t.Fatalf("file %+v", f)
	}

	// The popup card is not a request: no approve-all bar counts it and nothing approves it.
	groups := e.app.pendingViews()
	if len(groups) != 1 || len(groups[0].Requests) != 0 || len(groups[0].NewCharacters) != 1 || len(groups[0].BringBacks) != 0 {
		t.Fatalf("groups %+v", groups)
	}
	var html bytes.Buffer
	if err := views.PendingRequests(groups).Render(ctx, &html); err != nil {
		t.Fatal(err)
	}
	card := html.String()
	for _, want := range []string{"Pat brings Bryn, made at home", route.Path(route.CampaignHomeFile, e.campaignID, f.id), route.Path(route.CampaignHomeFileReject, e.campaignID, f.id), "Review", "Deny"} {
		if !strings.Contains(card, want) {
			t.Errorf("popup lacks %q", want)
		}
	}
	if strings.Contains(card, "Approve") || strings.Contains(card, route.Path(route.CampaignRequestDecision, e.campaignID, p.ID, "approve")) {
		t.Errorf("the card can be approved: %s", card)
	}
	if e.app.Hub.Decide(e.campaignID, p.ID, true) {
		t.Fatal("Decide approved a new-character request")
	}
	if w := e.do(t, http.MethodPost, route.Path(route.CampaignRequestsDecision, e.campaignID, "approve"), gm, url.Values{"id": {p.ID}}, true); w.Code != http.StatusOK {
		t.Errorf("approve-all: %d", w.Code)
	}
	if got := e.app.Hub.PendingList(e.campaignID); len(got) != 1 {
		t.Fatalf("approve-all took the request: %+v", got)
	}

	// Add: the request is approved with the new character, which the visitor redeems.
	review := route.Path(route.CampaignHomeFile, e.campaignID, f.id)
	rv, err := e.app.reviewHomeFile(ctx, e.q(), f, func(string) string { return "" })
	if err != nil {
		t.Fatal(err)
	}
	page := e.do(t, http.MethodGet, review, gm, nil, false).Body.String()
	if !strings.Contains(page, "sent from the player") || strings.Contains(page, secretTitle) {
		t.Errorf("review: %s", page)
	}
	if w := e.do(t, http.MethodPost, review+route.Accept, gm, url.Values{formWas: {rv.was}}, true); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Bryn joined the campaign.") {
		t.Fatalf("add: %d %s", w.Code, w.Body)
	}
	w := e.do(t, http.MethodPost, route.RequestsComplete, "", url.Values{"secret": {p.Secret}}, true)
	c := sessionCookie(w)
	if c == nil {
		t.Fatalf("complete: %d %v", w.Code, w.Header())
	}
	sess, err := e.mgr.Resolve(ctx, c.Value)
	if err != nil {
		t.Fatal(err)
	}
	ch, err := e.q().GetCharacter(ctx, sess.CharacterID.Int64)
	if err != nil || ch.Name != "Bryn" {
		t.Fatalf("the visitor logged in as %+v (%v)", ch, err)
	}
}

// TestNewCharacterVisitorGone: Add is refused, and nothing is created, once the
// visitor's request is gone.
func TestNewCharacterVisitorGone(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, enum.RoleAdmin)
	if w := e.visitorUpload(t, e.newCharacterZip(t, nil)); w.Code != http.StatusOK {
		t.Fatal(w.Body)
	}
	p := e.app.Hub.PendingList(e.campaignID)[0]
	f, _ := e.app.homeFiles.get(p.FileID, e.campaignID)
	rv, err := e.app.reviewHomeFile(ctx, e.q(), f, func(string) string { return "" })
	if err != nil {
		t.Fatal(err)
	}
	if !e.app.Hub.DenyFile(e.campaignID, f.id) {
		t.Fatal("the request was not waiting")
	}
	w := e.do(t, http.MethodPost, route.Path(route.CampaignHomeFileAccept, e.campaignID, f.id), gm, url.Values{formWas: {rv.was}}, true)
	if w.Code != http.StatusUnprocessableEntity || !strings.Contains(w.Body.String(), "The player left") {
		t.Fatalf("add: %d %s", w.Code, w.Body)
	}
	if chars, _ := e.q().ListCharactersByCampaign(ctx, e.campaignID); len(chars) != 1 {
		t.Errorf("characters %v: the refused add left one behind", chars)
	}
	if e.count(t, `SELECT count(*) FROM images`) != 0 || e.count(t, `SELECT count(*) FROM compendium_entries WHERE slug = 'glow-3'`) != 0 {
		t.Error("the refused add left rows behind")
	}
	if e.fileCount(t) != 1 {
		t.Errorf("%d zips, want the waiting one", e.fileCount(t))
	}
	// Rejecting a visitor's file denies its request and drops the file.
	if w := e.visitorUpload(t, e.newCharacterZip(t, nil)); w.Code != http.StatusOK {
		t.Fatal(w.Body)
	}
	p2 := e.app.Hub.PendingList(e.campaignID)[0]
	if w := e.do(t, http.MethodPost, route.Path(route.CampaignHomeFileReject, e.campaignID, p2.FileID), gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("deny: %d %s", w.Code, w.Body)
	}
	if got := e.app.Hub.PendingList(e.campaignID); len(got) != 0 {
		t.Errorf("the request still waits: %+v", got)
	}
	if _, ok := e.app.homeFiles.get(p2.FileID, e.campaignID); ok {
		t.Error("the denied file was kept")
	}
}

func TestKeepsBodyNewCharacter(t *testing.T) {
	for _, pattern := range []string{route.JoinNewCharacter, route.CampaignNewCharacter} {
		if keepsBody(path.Base(pattern), pattern) {
			t.Errorf("the log keeps the body of %s", pattern)
		}
	}
}

// TestHomeImportMentions: at home a mention of a note or look that went home
// points at its home copy; any other is negated, so it links nowhere.
func TestHomeImportMentions(t *testing.T) {
	e := newPlayerEnv(t)
	first := mentionDoc("see", "note:6", "note:99", "image:9", "image:55", "entity:3", "attachment:5")
	zipBody := sampleTakeHome(t, func(th *export.TakeHome) {
		th.Notes[0].Blocks[0].Doc = first
		th.Notes[0].Blocks[0].Files = nil
		th.Notes = append(th.Notes, export.TakeHomeNote{ID: 6, Kind: enum.NoteGeneric, Title: "Copy", CreatedAt: th.Manifest.CreatedAt, UpdatedAt: th.Manifest.CreatedAt,
			Blocks: []export.TakeHomeBlock{{ID: 7, Visibility: enum.VisOwner, Doc: mentionDoc("x"), ByCharacter: true}}})
	})
	w := e.importRequest(t, zipBody, "", nil)
	if w.Code != http.StatusNoContent {
		t.Fatalf("status %d: %s", w.Code, w.Body)
	}
	sess, err := e.mgr.Resolve(t.Context(), homeCookie(w).Value)
	if err != nil {
		t.Fatal(err)
	}
	notes, _ := e.q().ListOwnNotes(t.Context(), sess.CharacterID)
	looks, _ := e.q().ListCharacterImages(t.Context(), sess.CharacterID.Int64)
	blocks, _ := e.q().ListOwnNoteBlocks(t.Context(), sess.CharacterID)
	var copyID int64
	for _, n := range notes {
		if n.Title == "Copy" {
			copyID = n.ID
		}
	}
	if copyID == 0 || len(looks) != 1 || len(blocks) == 0 {
		t.Fatalf("notes %v looks %v blocks %v", notes, looks, blocks)
	}
	doc := blocks[0].DocJson
	for _, want := range []string{fmt.Sprintf(`"note:%d"`, copyID), `"note:-99"`, fmt.Sprintf(`"image:%d"`, looks[0].ID), `"image:-55"`, `"entity:3"`} {
		if !strings.Contains(doc, want) {
			t.Errorf("doc %s lacks %s", doc, want)
		}
	}
}
