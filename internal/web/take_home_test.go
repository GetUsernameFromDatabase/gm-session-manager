package web

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/db"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/route"
)

// takeHomeSeed is what seedTakeHome put on the character at the table.
type takeHomeSeed struct {
	ownerBlock, playersBlock int64
}

// seedTakeHome gives the env's character a level-up, a look with a file, an
// own homebrew entry and a note with a private and a public block, each with
// a file.
func (e testEnv) seedTakeHome(t *testing.T) takeHomeSeed {
	t.Helper()
	ctx := t.Context()
	q := e.q()
	dir := e.app.attachmentsDir()
	if err := os.MkdirAll(dir, 0o750); err != nil {
		t.Fatal(err)
	}
	file := func(name string) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("data-"+name), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	char := nullID(e.characterID)

	if err := q.CreateLevelUp(ctx, sqlc.CreateLevelUpParams{CharacterID: e.characterID, MaxHp: 5, SheetJson: "{}"}); err != nil {
		t.Fatal(err)
	}
	file("lookfile")
	att, err := q.CreateAttachment(ctx, sqlc.CreateAttachmentParams{CampaignID: e.campaignID, Path: "lookfile", Name: "look.png", Mime: "image/png", Size: 13})
	if err != nil {
		t.Fatal(err)
	}
	img, err := q.CreateImage(ctx, sqlc.CreateImageParams{CampaignID: e.campaignID, Name: "Hood", AttachmentID: nullID(att.ID)})
	if err != nil {
		t.Fatal(err)
	}
	if err := q.LinkCharacterImage(ctx, sqlc.LinkCharacterImageParams{CharacterID: e.characterID, ImageID: img.ID}); err != nil {
		t.Fatal(err)
	}
	if _, err := q.CreateHomebrew(ctx, sqlc.CreateHomebrewParams{Kind: enum.KindSpell, Slug: "zap", Name: "Zap", DataJson: "{}", ContentHash: "h1", OwnerCharacterID: char, Scoped: true}); err != nil {
		t.Fatal(err)
	}

	note, err := q.CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: enum.NoteGeneric, Title: "Diary", OwnerCharacterID: char})
	if err != nil {
		t.Fatal(err)
	}
	var s takeHomeSeed
	for _, b := range []struct {
		vis  enum.Visibility
		path string
		out  *int64
	}{{enum.VisOwner, "ownfile", &s.ownerBlock}, {enum.VisPlayers, "pubfile", &s.playersBlock}} {
		blk, err := q.CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: b.vis, AuthorCharacterID: char})
		if err != nil {
			t.Fatal(err)
		}
		*b.out = blk.ID
		file(b.path)
		if _, err := q.CreateAttachment(ctx, sqlc.CreateAttachmentParams{CampaignID: e.campaignID, BlockID: nullID(blk.ID), Path: b.path, Name: b.path + ".png", Mime: "image/png", Size: 10}); err != nil {
			t.Fatal(err)
		}
	}
	// Steps past the last snapshot: the file holds the snapshot, not them.
	if _, err := e.conn.ExecContext(ctx, `UPDATE note_blocks SET version = 3 WHERE id = ?`, s.playersBlock); err != nil {
		t.Fatal(err)
	}
	return s
}

// otherPlayer is a second character with a session.
func (e testEnv) otherPlayer(t *testing.T) string {
	t.Helper()
	ch, err := e.q().CreateCharacterWithSheet(t.Context(), sqlc.CreateCharacterWithSheetParams{CampaignID: e.campaignID, Name: "Bob", System: "dnd5e", SheetJson: "{}"})
	if err != nil {
		t.Fatal(err)
	}
	tok, _, err := e.mgr.IssuePlayer(t.Context(), ch.ID)
	if err != nil {
		t.Fatal(err)
	}
	return tok
}

func (e testEnv) takeHomeRequest(t *testing.T, method, remoteAddr, cookie string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequestWithContext(t.Context(), method, route.Path(route.CharacterTakeHome, e.characterID), nil)
	req.RemoteAddr = remoteAddr
	if cookie != "" {
		req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: cookie})
	}
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	return w
}

func readZip(t *testing.T, body []byte) *export.TakeHomeFile {
	t.Helper()
	f, err := export.ReadTakeHome(bytes.NewReader(body), int64(len(body)), export.KindTakeHome)
	if err != nil {
		t.Fatal(err)
	}
	return f
}

func hasBlock(f *export.TakeHomeFile, id int64) bool {
	for _, n := range f.Notes {
		for _, b := range n.Blocks {
			if b.ID == id {
				return true
			}
		}
	}
	return false
}

// TestTakeHomeDownload: who may take a character home, and what the file
// holds. A GM's file holds the player's private blocks and their files too.
func TestTakeHomeDownload(t *testing.T) {
	tests := []struct {
		name       string
		remoteAddr string
		who        string // "owner", "gm", "other" or "visitor"
		wantStatus int
		wantOwner  bool
	}{
		{"owner", lanAddr, "owner", http.StatusOK, true},
		{"gm", loopbackAddr, "gm", http.StatusOK, true},
		{"other player", lanAddr, "other", http.StatusForbidden, false},
		{"visitor", lanAddr, "visitor", http.StatusForbidden, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := newTestEnv(t)
			seed := e.seedTakeHome(t)
			cookie := map[string]string{"owner": e.token(t, enum.RolePlayer), "other": e.otherPlayer(t)}[tt.who]

			w := e.takeHomeRequest(t, http.MethodPost, tt.remoteAddr, cookie)
			if w.Code != tt.wantStatus {
				t.Fatalf("status %d, want %d: %s", w.Code, tt.wantStatus, w.Body)
			}
			rows, err := e.q().ListOpenTakeHomesByCampaign(t.Context(), e.campaignID)
			if err != nil {
				t.Fatal(err)
			}
			if tt.wantStatus != http.StatusOK {
				if len(rows) != 0 {
					t.Error("a refused request recorded a take-out")
				}
				return
			}
			if ct := w.Header().Get("Content-Type"); ct != "application/zip" || !strings.Contains(w.Header().Get("Content-Disposition"), "alice-take-home-") {
				t.Errorf("headers: %v", w.Header())
			}
			f := readZip(t, w.Body.Bytes())
			if hasBlock(f, seed.ownerBlock) != tt.wantOwner || !hasBlock(f, seed.playersBlock) {
				t.Errorf("blocks owner=%v players=%v, want owner=%v players=true", hasBlock(f, seed.ownerBlock), hasBlock(f, seed.playersBlock), tt.wantOwner)
			}
			sent := map[string]bool{}
			for _, n := range f.Notes {
				for _, b := range n.Blocks {
					for _, bf := range b.Files {
						sent[bf.File] = true
					}
					if b.ID == seed.playersBlock && b.Version != 0 {
						t.Errorf("file block version %d, want the snapshot's 0", b.Version)
					}
				}
			}
			if sent[export.ZipFile("ownfile")] != tt.wantOwner || !sent[export.ZipFile("pubfile")] {
				t.Errorf("files sent: %v", sent)
			}
			if len(f.Looks) != 1 || len(f.LevelUps) != 1 || len(f.Homebrew) != 1 {
				t.Errorf("looks %d, level-ups %d, homebrew %d", len(f.Looks), len(f.LevelUps), len(f.Homebrew))
			}

			if len(rows) != 1 {
				t.Fatalf("%d open take-homes, want 1", len(rows))
			}
			rec, err := e.q().LatestOpenTakeHome(t.Context(), e.characterID)
			if err != nil {
				t.Fatal(err)
			}
			var base export.TakeHomeBase
			if err := json.Unmarshal([]byte(rec.BaseJson), &base); err != nil {
				t.Fatal(err)
			}
			if _, ok := base.Blocks[seed.ownerBlock]; ok != tt.wantOwner {
				t.Errorf("base lists the private block = %v, want %v", ok, tt.wantOwner)
			}
			if v, ok := base.Blocks[seed.playersBlock]; !ok || v != 0 || rec.Token != f.Manifest.Token || f.Manifest.ServerID == "" {
				t.Errorf("base %v (want the snapshot's version 0), token %q vs %q, server %q", base.Blocks, rec.Token, f.Manifest.Token, f.Manifest.ServerID)
			}
			if serverID, err := e.q().GetServerID(t.Context()); err != nil || f.Manifest.ServerID != serverID {
				t.Errorf("manifest server %q, database %q (%v)", f.Manifest.ServerID, serverID, err)
			}
		})
	}
}

func TestTakeHomeDialog(t *testing.T) {
	tests := []struct {
		name       string
		remoteAddr string
		who        string
		wantStatus int
		wantText   string
	}{
		{"owner sees the app files", lanAddr, "owner", http.StatusOK, route.Path(route.AppDownload, "player-app.zip")},
		{"gm sees the app files and the private notes line", loopbackAddr, "gm", http.StatusOK, "private notes"},
		{"other player", lanAddr, "other", http.StatusForbidden, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := newTestEnv(t)
			dir := filepath.Join(e.app.DataDir, DownloadsDir)
			if err := os.MkdirAll(dir, 0o750); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(dir, "player-app.zip"), []byte("x"), 0o600); err != nil {
				t.Fatal(err)
			}
			cookie := map[string]string{"owner": e.token(t, enum.RolePlayer), "other": e.otherPlayer(t)}[tt.who]
			w := e.takeHomeRequest(t, http.MethodGet, tt.remoteAddr, cookie)
			if w.Code != tt.wantStatus || !strings.Contains(w.Body.String(), tt.wantText) {
				t.Errorf("status %d, body %q; want %d containing %q", w.Code, w.Body, tt.wantStatus, tt.wantText)
			}
		})
	}
}

// TestTakenHomeBadge: the owner and the GM see when the file went out, other
// players do not, and the roster shows it.
func TestTakenHomeBadge(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	row, err := e.q().GetCharacter(ctx, e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name   string
		viewer perm.Viewer
		want   bool
	}{
		{"owner", perm.Player(e.characterID), true},
		{"gm", perm.GM, true},
		{"other player", perm.Player(e.characterID + 100), false},
	}
	before, err := e.app.sheetView(ctx, perm.GM, row)
	if err != nil || !before.TakenHome.IsZero() {
		t.Fatalf("badge before a take-out: %v %v", before.TakenHome, err)
	}
	if w := e.takeHomeRequest(t, http.MethodPost, loopbackAddr, ""); w.Code != http.StatusOK {
		t.Fatalf("status %d", w.Code)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := e.app.sheetView(ctx, tt.viewer, row)
			if err != nil {
				t.Fatal(err)
			}
			if got := !s.TakenHome.IsZero(); got != tt.want {
				t.Errorf("badge = %v, want %v", got, tt.want)
			}
		})
	}
	r, err := e.app.roster(ctx, e.campaignID)
	if err != nil || len(r.Characters) != 1 || r.Characters[0].TakenHome.IsZero() {
		t.Errorf("roster badge: %+v, %v", r.Characters, err)
	}
}

// newPlayerEnv is the player app: no admin, no campaign yet.
func newPlayerEnv(t testing.TB) testEnv {
	t.Helper()
	conn, err := db.Open(t.Context(), filepath.Join(t.TempDir(), "home.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	q := db.NewStore(conn)
	mgr := auth.NewPlayerManager(q)
	app := New(Deps{Auth: mgr, Store: q, DataDir: t.TempDir(), Port: 8042, Player: true, Version: "v-home"})
	t.Cleanup(app.Close)
	return testEnv{router: app.Router(), mgr: mgr, app: app, conn: conn}
}

func TestTakeHomeRefusedInPlayerApp(t *testing.T) {
	e := newPlayerEnv(t)
	for _, method := range []string{http.MethodGet, http.MethodPost} {
		req := httptest.NewRequestWithContext(t.Context(), method, route.Path(route.CharacterTakeHome, 1), nil)
		req.RemoteAddr = loopbackAddr
		w := httptest.NewRecorder()
		e.router.ServeHTTP(w, req)
		if w.Code != http.StatusNotFound && w.Code != http.StatusForbidden {
			t.Errorf("%s: status %d, want a refusal", method, w.Code)
		}
	}
	// And the GM app answers 404 to the home pages.
	g := newTestEnv(t)
	req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, route.AtHome, nil)
	req.RemoteAddr = loopbackAddr
	w := httptest.NewRecorder()
	g.router.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("GM app home page: status %d, want 404", w.Code)
	}
}

// sampleTakeHome is a file as the table would write it.
func sampleTakeHome(t *testing.T, mutate func(*export.TakeHome)) []byte {
	t.Helper()
	dir := t.TempDir()
	for _, n := range []string{"tablelook", "tablenote"} {
		if err := os.WriteFile(filepath.Join(dir, n), []byte("data-"+n), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	sheet, err := newSheetJSON()
	if err != nil {
		t.Fatal(err)
	}
	const tableAttachment = 50
	doc := `{"type":"doc","content":[{"type":"image","attrs":{"src":"` + route.Path(route.Attachment, tableAttachment) + `"}},{"type":"paragraph","content":[{"type":"text","text":"` + route.Path(route.Attachment, tableAttachment) + `"}]}]}`
	made := time.Date(2026, 11, 24, 18, 0, 0, 0, time.UTC)
	th := export.TakeHome{
		Manifest: export.TakeHomeManifest{
			Format: export.TakeHomeFormat, Kind: export.KindTakeHome, AppVersion: "v-table", Token: "tok", ServerID: "srv", CharacterID: 7, CreatedAt: made,
			Campaign: export.TakeHomeCampaign{Name: "Table", System: "dnd5e", Edition: "2024"},
		},
		Character: export.CharacterFile{ID: 7, Name: "Aria", System: "dnd5e", HP: 0, MaxHP: 0, Sheet: json.RawMessage(sheet)},
		LevelUps:  []export.TakeHomeLevelUp{{ID: 3, MaxHP: 4, Sheet: json.RawMessage(sheet), CreatedAt: made}},
		Looks:     []export.TakeHomeLook{{ImageID: 9, AttachmentID: 8, Name: "Hood", File: export.ZipFile("tablelook"), IsRevealed: true, IsCurrent: true}},
		Notes: []export.TakeHomeNote{{ID: 4, Kind: enum.NoteGeneric, Title: "Diary", CreatedAt: made, UpdatedAt: made, Blocks: []export.TakeHomeBlock{{
			ID: 5, Position: 1, Visibility: enum.VisOwner, Doc: json.RawMessage(doc), Version: 12, ByCharacter: true,
			Files: []export.TakeHomeBlockFile{{AttachmentID: tableAttachment, Name: "map.png", File: export.ZipFile("tablenote")}},
		}}}},
		Homebrew: []export.TakeHomeEntry{{ID: 20, Kind: enum.KindSpell, Slug: "zap", Name: "Zap", SourceKey: "homebrew", Data: json.RawMessage(`{}`), ContentHash: "h1", Scoped: true}},
		Sources:  []export.TakeHomeSource{{SourceKey: "src1", Edition: "2024", Priority: 150}},
		Entries: []export.TakeHomeEntry{
			{Kind: enum.KindSpell, Slug: "bolt", Name: "Bolt", SourceKey: "src1", Data: json.RawMessage(`{}`), ContentHash: "h2"},
			{Kind: enum.KindFeat, Slug: "gm-feat", Name: "Table feat", SourceKey: "homebrew", Data: json.RawMessage(`{}`), ContentHash: "h3"},
		},
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

// importRequest posts a take-home file to the player app, as htmx does.
func (e testEnv) importRequest(t *testing.T, zipBody []byte, cookie string, fields map[string]string) *httptest.ResponseRecorder {
	t.Helper()
	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	part, err := mw.CreateFormFile("file", "take-home.zip")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = part.Write(zipBody)
	for k, v := range fields {
		_ = mw.WriteField(k, v)
	}
	_ = mw.Close()
	req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, route.AtHome, &body)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("HX-Request", "true")
	req.RemoteAddr = loopbackAddr
	if cookie != "" {
		req.AddCookie(&http.Cookie{Name: auth.PlayerCookieName, Value: cookie})
	}
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	return w
}

func homeCookie(w *httptest.ResponseRecorder) *http.Cookie {
	for _, c := range w.Result().Cookies() {
		if c.Name == auth.PlayerCookieName && c.MaxAge > 0 {
			return c
		}
	}
	return nil
}

func (e testEnv) get(t *testing.T, path, cookie string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, path, nil)
	req.RemoteAddr = loopbackAddr
	if cookie != "" {
		req.AddCookie(&http.Cookie{Name: auth.PlayerCookieName, Value: cookie})
	}
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	return w
}

// TestHomeImportOneTable: a player app keeps one table's characters, so a file
// from another server is refused and nothing of it is stored.
func TestHomeImportOneTable(t *testing.T) {
	e := newPlayerEnv(t)
	if w := e.importRequest(t, sampleTakeHome(t, nil), "", nil); w.Code != http.StatusNoContent {
		t.Fatalf("first import: status %d: %s", w.Code, w.Body)
	}
	other := sampleTakeHome(t, func(th *export.TakeHome) { th.Manifest.ServerID = "other-table" })
	w := e.importRequest(t, other, "", nil)
	if w.Code != http.StatusUnprocessableEntity || !strings.Contains(w.Body.String(), "second folder") {
		t.Errorf("other table: status %d: %s", w.Code, w.Body)
	}
	if copies, err := e.q().ListHomeCopies(t.Context()); err != nil || len(copies) != 1 {
		t.Errorf("home copies %d, want 1 (%v)", len(copies), err)
	}
}

// TestHomeImportSharedPicture: a picture copied into another note points at
// the first block's file; at home it must point at that file's new id too.
func TestHomeImportSharedPicture(t *testing.T) {
	e := newPlayerEnv(t)
	copied := `{"type":"doc","content":[{"type":"image","attrs":{"src":"` + route.Path(route.Attachment, 50) + `"}}]}`
	zipBody := sampleTakeHome(t, func(th *export.TakeHome) {
		th.Notes = append(th.Notes, export.TakeHomeNote{ID: 6, Kind: enum.NoteGeneric, Title: "Copy", CreatedAt: th.Manifest.CreatedAt, UpdatedAt: th.Manifest.CreatedAt,
			Blocks: []export.TakeHomeBlock{{ID: 7, Visibility: enum.VisOwner, Doc: json.RawMessage(copied), ByCharacter: true}}})
	})
	w := e.importRequest(t, zipBody, "", nil)
	if w.Code != http.StatusNoContent {
		t.Fatalf("status %d: %s", w.Code, w.Body)
	}
	sess, err := e.mgr.Resolve(t.Context(), homeCookie(w).Value)
	if err != nil {
		t.Fatal(err)
	}
	files, err := e.q().ListOwnNoteAttachments(t.Context(), sess.CharacterID)
	if err != nil || len(files) != 1 {
		t.Fatalf("files %+v, %v", files, err)
	}
	blocks, err := e.q().ListOwnNoteBlocks(t.Context(), sess.CharacterID)
	if err != nil || len(blocks) != 2 {
		t.Fatalf("blocks %+v, %v", blocks, err)
	}
	want := `"src":"` + route.Path(route.Attachment, files[0].ID) + `"`
	for _, b := range blocks {
		if !strings.Contains(b.DocJson, want) {
			t.Errorf("block %d doc %s, want %s", b.ID, b.DocJson, want)
		}
	}
}

func TestHomeImport(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	w := e.importRequest(t, sampleTakeHome(t, nil), "", nil)
	if w.Code != http.StatusNoContent {
		t.Fatalf("status %d: %s", w.Code, w.Body)
	}
	c := homeCookie(w)
	if c == nil || c.Secure || !c.HttpOnly {
		t.Fatalf("cookie %+v", c)
	}
	loc := w.Header().Get("HX-Redirect")

	sess, err := e.mgr.Resolve(ctx, c.Value)
	if err != nil {
		t.Fatal(err)
	}
	charID := sess.CharacterID.Int64
	if loc != route.Path(route.Character, charID) {
		t.Errorf("redirect %q", loc)
	}
	ch, err := e.q().GetCharacter(ctx, charID)
	if err != nil || ch.Name != "Aria" {
		t.Fatalf("character %+v, %v", ch, err)
	}
	camp, err := e.q().GetCampaign(ctx, ch.CampaignID)
	if err != nil || camp.Name != "Table" || camp.OpenJoin {
		t.Errorf("campaign %+v, %v", camp, err)
	}
	if n, _ := e.q().CountLevelUps(ctx, charID); n != 1 {
		t.Errorf("%d level-ups", n)
	}
	looks, err := e.q().ListCharacterImages(ctx, charID)
	if err != nil || len(looks) != 1 || !looks[0].AttachmentID.Valid || !looks[0].IsCurrent || !looks[0].IsRevealed {
		t.Fatalf("looks %+v, %v", looks, err)
	}
	att, err := e.q().GetAttachment(ctx, looks[0].AttachmentID.Int64)
	if err != nil {
		t.Fatal(err)
	}
	if got, err := os.ReadFile(filepath.Join(e.app.attachmentsDir(), att.Path)); err != nil || string(got) != "data-tablelook" {
		t.Errorf("look file %q, %v", got, err)
	}

	notes, err := e.q().ListOwnNotes(ctx, nullID(charID))
	if err != nil || len(notes) != 1 {
		t.Fatalf("notes %+v, %v", notes, err)
	}
	blocks, err := e.q().ListOwnNoteBlocks(ctx, nullID(charID))
	if err != nil || len(blocks) != 1 || blocks[0].Visibility != enum.VisOwner || blocks[0].Version != 0 || !blocks[0].AuthorCharacterID.Valid {
		t.Fatalf("blocks %+v, %v", blocks, err)
	}
	files, err := e.q().ListOwnNoteAttachments(ctx, nullID(charID))
	if err != nil || len(files) != 1 {
		t.Fatalf("block files %+v, %v", files, err)
	}
	wantSrc := route.Path(route.Attachment, files[0].ID)
	oldSrc := route.Path(route.Attachment, 50)
	// The image node points at the new file; text that only looks like the old
	// address stays as written.
	if doc := blocks[0].DocJson; !strings.Contains(doc, `"src":"`+wantSrc+`"`) || strings.Contains(doc, `"src":"`+oldSrc) || !strings.Contains(doc, `"text":"`+oldSrc+`"`) {
		t.Errorf("doc %s, want src %s", doc, wantSrc)
	}

	own, err := e.q().ListOwnHomebrew(ctx, nullID(charID))
	if err != nil || len(own) != 1 || !own[0].Scoped {
		t.Fatalf("own homebrew %+v, %v", own, err)
	}
	if linked, _ := e.q().ListEntryCampaigns(ctx, own[0].ID); len(linked) != 1 || linked[0] != camp.ID {
		t.Errorf("own homebrew linked to %v", linked)
	}
	if _, err := e.q().GetCompendiumEntryBySlugSource(ctx, sqlc.GetCompendiumEntryBySlugSourceParams{Kind: enum.KindSpell, Slug: "bolt", SourceKey: "src1"}); err != nil {
		t.Errorf("source entry: %v", err)
	}
	if src, err := e.q().GetCompendiumSource(ctx, "src1"); err != nil || src.Priority != 150 {
		t.Errorf("source %+v, %v", src, err)
	}
	if ok, _ := e.q().ListCampaignSources(ctx, camp.ID); len(ok) != 1 {
		t.Errorf("campaign sources %v", ok)
	}
	if hc, err := e.q().GetHomeCopyByTable(ctx, sqlc.GetHomeCopyByTableParams{ServerID: "srv", TableCharacterID: 7}); err != nil || hc.CharacterID != charID {
		t.Errorf("home copy %+v, %v", hc, err)
	}

	// The browser is the character, not the admin, though it is on loopback.
	if w := e.get(t, route.Campaigns+"/"+"1", ""); w.Code == http.StatusOK {
		t.Error("loopback is admin in the player app")
	}
	if w := e.get(t, route.Home, c.Value); w.Header().Get("Location") != route.Path(route.Character, charID) {
		t.Errorf("home goes to %q", w.Header().Get("Location"))
	}
	if w := e.get(t, route.Home, ""); w.Header().Get("Location") != route.AtHome {
		t.Errorf("visitor home goes to %q", w.Header().Get("Location"))
	}
	if w := e.get(t, route.AtHome, c.Value); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Playing") {
		t.Errorf("at-home page %d", w.Code)
	}
}

func TestHomeImportReplace(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	first := e.importRequest(t, sampleTakeHome(t, nil), "", nil)
	c := homeCookie(first)
	oldSess, _ := e.mgr.Resolve(ctx, c.Value)
	oldID := oldSess.CharacterID.Int64

	newer := sampleTakeHome(t, func(th *export.TakeHome) {
		th.Manifest.CreatedAt = th.Manifest.CreatedAt.Add(48 * time.Hour)
		th.Character.Name = "Aria Two"
	})
	w := e.importRequest(t, newer, c.Value, nil)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Aria Two is already on this computer") || !strings.Contains(w.Body.String(), `name="replace"`) {
		t.Fatalf("no confirmation: %d %s", w.Code, w.Body)
	}
	if _, err := e.q().GetCharacter(ctx, oldID); err != nil {
		t.Fatal("the copy went before the player confirmed")
	}

	older := sampleTakeHome(t, func(th *export.TakeHome) { th.Manifest.CreatedAt = th.Manifest.CreatedAt.Add(-48 * time.Hour) })
	if w := e.importRequest(t, older, c.Value, nil); !strings.Contains(w.Body.String(), "is older than the copy here") {
		t.Errorf("older file not named: %s", w.Body)
	}

	w = e.importRequest(t, newer, c.Value, map[string]string{"replace": "1"})
	if w.Code != http.StatusNoContent {
		t.Fatalf("replace: %d %s", w.Code, w.Body)
	}
	// Ids are reused, so count: the old campaign is gone, not kept beside the new.
	if camps, err := e.q().ListCampaigns(ctx); err != nil || len(camps) != 1 {
		t.Errorf("%d campaigns after a replace, want 1 (%v)", len(camps), err)
	}
	rows, err := e.q().ListHomeCopies(ctx)
	if err != nil || len(rows) != 1 || rows[0].CharacterName != "Aria Two" {
		t.Errorf("home copies %+v, %v", rows, err)
	}
	if hb, _ := e.q().ListHomebrew(ctx); len(hb) != 2 { // own spell and the table's feat, each once
		t.Errorf("%d homebrew entries after a replace, want 2", len(hb))
	}
}

// TestHomeReplaceKeepsSharedHomebrew: two characters of one table at home share
// a homebrew entry the first owns. Replacing the first must leave the second's
// campaign seeing it, and the new copy owns it again.
func TestHomeReplaceKeepsSharedHomebrew(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	first := e.importRequest(t, sampleTakeHome(t, nil), "", nil)
	c1 := homeCookie(first)
	if c1 == nil {
		t.Fatalf("first import: %d %s", first.Code, first.Body)
	}
	s1, _ := e.mgr.Resolve(ctx, c1.Value) // the second import ends this session
	zap := export.TakeHomeEntry{Kind: enum.KindSpell, Slug: "zap", Name: "Zap", SourceKey: "homebrew", Data: json.RawMessage(`{}`), ContentHash: "h1", Scoped: true}
	second := e.importRequest(t, sampleTakeHome(t, func(th *export.TakeHome) {
		th.Manifest.CharacterID, th.Character.Name = 8, "Brin"
		th.Homebrew = nil
		th.Entries = append(th.Entries, zap) // the table shows Aria's entry to Brin too
	}), c1.Value, nil)
	c2 := homeCookie(second)
	if c2 == nil {
		t.Fatalf("second import: %d %s", second.Code, second.Body)
	}
	s2, _ := e.mgr.Resolve(ctx, c2.Value)
	brin, err := e.q().GetCharacter(ctx, s2.CharacterID.Int64)
	if err != nil {
		t.Fatal(err)
	}
	entry, err := e.q().GetCompendiumEntryBySlugSource(ctx, sqlc.GetCompendiumEntryBySlugSourceParams{Kind: enum.KindSpell, Slug: "zap", SourceKey: "homebrew"})
	if err != nil {
		t.Fatal(err)
	}
	if entry.OwnerCharacterID.Int64 != s1.CharacterID.Int64 || !entry.Scoped {
		t.Fatalf("entry %+v: a link-only import changed its owner (want %d) or scope", entry, s1.CharacterID.Int64)
	}

	replaced := e.importRequest(t, sampleTakeHome(t, nil), "", map[string]string{"replace": "1"})
	if replaced.Code != http.StatusNoContent {
		t.Fatalf("replace: %d %s", replaced.Code, replaced.Body)
	}
	entry, err = e.q().GetCompendiumEntryBySlugSource(ctx, sqlc.GetCompendiumEntryBySlugSourceParams{Kind: enum.KindSpell, Slug: "zap", SourceKey: "homebrew"})
	if err != nil {
		t.Fatal("the shared entry was deleted with the replaced copy")
	}
	linked, err := e.q().ListEntryCampaigns(ctx, entry.ID)
	if err != nil || !slices.Contains(linked, brin.CampaignID) {
		t.Errorf("entry linked to %v, want it to include Brin's campaign %d (%v)", linked, brin.CampaignID, err)
	}
	newSess, err := e.mgr.Resolve(ctx, homeCookie(replaced).Value)
	if err != nil || entry.OwnerCharacterID.Int64 != newSess.CharacterID.Int64 {
		t.Errorf("owner %v, want the new copy %d (%v)", entry.OwnerCharacterID, newSess.CharacterID.Int64, err)
	}
}

func TestHomeImportRefusals(t *testing.T) {
	sheetWith := func(extra string) json.RawMessage {
		s, _ := newSheetJSON()
		return json.RawMessage(strings.TrimSuffix(s, "}") + extra + "}")
	}
	tests := []struct {
		name   string
		mutate func(*export.TakeHome)
		want   []string
	}{
		{"sheet with an unknown field", func(th *export.TakeHome) { th.Character.Sheet = sheetWith(`,"from_the_future":1`) }, []string{"v-table", "v-home"}},
		{"level-up from a newer app", func(th *export.TakeHome) { th.LevelUps[0].Sheet = sheetWith(`,"effects_version":99`) }, []string{"v-table", "v-home"}},
		{"unsafe link in a note", func(th *export.TakeHome) {
			th.Notes[0].Blocks[0].Doc = json.RawMessage(`{"type":"doc","content":[{"type":"text","text":"x","marks":[{"type":"link","attrs":{"href":"javascript:alert(1)"}}]}]}`)
		}, []string{"link"}},
		{"look link that is not a web address", func(th *export.TakeHome) { th.Looks[0].URL = "javascript:alert(1)" }, []string{"web address"}},
		{"entry from an unlisted source", func(th *export.TakeHome) { th.Entries[0].SourceKey = "other" }, []string{"source"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := newPlayerEnv(t)
			w := e.importRequest(t, sampleTakeHome(t, tt.mutate), "", nil)
			if w.Code != http.StatusUnprocessableEntity {
				t.Fatalf("status %d: %s", w.Code, w.Body)
			}
			for _, s := range tt.want {
				if !strings.Contains(w.Body.String(), s) {
					t.Errorf("response lacks %q: %s", s, w.Body)
				}
			}
			if homeCookie(w) != nil {
				t.Error("a refused file logged the browser in")
			}
			if rows, _ := e.q().ListHomeCopies(t.Context()); len(rows) != 0 {
				t.Error("a refused file left a home copy")
			}
		})
	}
}

func TestHomePlay(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	first := e.importRequest(t, sampleTakeHome(t, nil), "", nil)
	c1 := homeCookie(first)
	s1, _ := e.mgr.Resolve(ctx, c1.Value)
	second := e.importRequest(t, sampleTakeHome(t, func(th *export.TakeHome) { th.Manifest.CharacterID = 8; th.Character.Name = "Brin" }), c1.Value, nil)
	c2 := homeCookie(second)
	if c2 == nil {
		t.Fatalf("second import: %d %s", second.Code, second.Body)
	}
	if _, err := e.mgr.Resolve(ctx, c1.Value); err == nil {
		t.Error("importing a second character kept the first session")
	}

	play := func(id int64, cookie string) *httptest.ResponseRecorder {
		req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, route.Path(route.AtHomePlay, id), nil)
		req.RemoteAddr = loopbackAddr
		req.AddCookie(&http.Cookie{Name: auth.PlayerCookieName, Value: cookie})
		w := httptest.NewRecorder()
		e.router.ServeHTTP(w, req)
		return w
	}
	w := play(s1.CharacterID.Int64, c2.Value)
	if w.Code != http.StatusSeeOther || w.Header().Get("Location") != route.Path(route.Character, s1.CharacterID.Int64) {
		t.Fatalf("play: %d %q", w.Code, w.Header().Get("Location"))
	}
	c3 := homeCookie(w)
	if c3 == nil {
		t.Fatal("no new cookie")
	}
	if _, err := e.mgr.Resolve(ctx, c2.Value); err == nil {
		t.Error("the previous session survived the switch")
	}
	if s, err := e.mgr.Resolve(ctx, c3.Value); err != nil || s.CharacterID.Int64 != s1.CharacterID.Int64 {
		t.Errorf("new session %+v, %v", s, err)
	}

	// A character that was not taken home cannot be played from here.
	other, err := e.q().CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Other", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	stray, err := e.q().CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{CampaignID: other.ID, Name: "Stray", System: "dnd5e", SheetJson: "{}"})
	if err != nil {
		t.Fatal(err)
	}
	if w := play(stray.ID, c3.Value); w.Code != http.StatusNotFound {
		t.Errorf("playing a character with no home copy: status %d", w.Code)
	}
}

func TestCheckSheetFile(t *testing.T) {
	current, _ := newSheetJSON()
	tests := []struct {
		name string
		raw  string
		want bool // refused as newer
	}{
		{"current sheet", current, false},
		{"empty sheet", "{}", false},
		{"unknown field", `{"from_the_future":1}`, true},
		{"higher effects version", `{"effects_version":999}`, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkSheetFile(json.RawMessage(tt.raw))
			if (err != nil) != tt.want {
				t.Errorf("err = %v, want refusal %v", err, tt.want)
			}
		})
	}
}

func TestRewriteDocAttachments(t *testing.T) {
	doc := json.RawMessage(`{"type":"doc","content":[{"type":"paragraph","content":[{"type":"mention","attrs":{"id":"attachment:5","label":"a"}},{"type":"mention","attrs":{"id":"attachment:6","label":"b"}}]},{"type":"image","attrs":{"src":"` + route.Path(route.Attachment, 5) + `"}}]}`)
	got, err := rewriteDocAttachments(doc, map[int64]int64{5: 77})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"attachment:77"`, `"attachment:6"`, `"` + route.Path(route.Attachment, 77) + `"`} {
		if !bytes.Contains(got, []byte(want)) {
			t.Errorf("%s lacks %s", got, want)
		}
	}
}
