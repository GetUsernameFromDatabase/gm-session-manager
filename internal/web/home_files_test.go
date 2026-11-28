package web

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"mime/multipart"
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

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

// takeOut takes the env's character home as the GM and returns the take-out.
func (e testEnv) takeOut(t *testing.T) sqlc.TakeHome {
	t.Helper()
	if w := e.takeHomeRequest(t, http.MethodPost, loopbackAddr, ""); w.Code != http.StatusOK {
		t.Fatalf("take home: status %d: %s", w.Code, w.Body)
	}
	rec, err := e.q().LatestOpenTakeHome(t.Context(), e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	return rec
}

// homeSheet is a sheet at home: a level 1 wizard, changed by change.
func homeSheet(t *testing.T, change func(*dnd5e.Sheet)) json.RawMessage {
	t.Helper()
	s, err := dnd5e.ParseSheet("")
	if err != nil {
		t.Fatal(err)
	}
	s.Classes = []dnd5e.ClassLevel{
		{Slug: "wizard", Name: "Wizard", Level: 1, HitDie: 6, Caster: dnd5e.CasterFull},
	}
	if change != nil {
		change(&s)
	}
	raw, err := s.JSON()
	if err != nil {
		t.Fatal(err)
	}
	return json.RawMessage(raw)
}

// bringBackZip is the bring-back file of the env's character for take-out rec,
// as the player app writes it, changed by mutate.
func (e testEnv) bringBackZip(
	t *testing.T,
	rec sqlc.TakeHome,
	mutate func(*export.TakeHome),
) []byte {
	t.Helper()
	return e.bringBackZipFiles(t, rec, mutate, t.TempDir())
}

// bringBackZipFiles is bringBackZip with the files the looks and blocks name
// read from dir.
func (e testEnv) bringBackZipFiles(
	t *testing.T,
	rec sqlc.TakeHome,
	mutate func(*export.TakeHome),
	dir string,
) []byte {
	t.Helper()
	serverID, err := e.q().GetServerID(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	th := export.TakeHome{
		Manifest: export.TakeHomeManifest{
			Format:     export.TakeHomeFormat,
			Kind:       export.KindBringBack,
			AppVersion: "v-home",
			Token:      rec.Token,
			ServerID:   serverID,
			Campaign: export.TakeHomeCampaign{
				Name:    "Test",
				System:  gameSystem,
				Edition: gameEdition,
			},
			CharacterID: e.characterID,
			CreatedAt:   rec.CreatedAt.UTC(),
			SavedAt:     time.Date(2026, 11, 25, 9, 0, 0, 0, time.UTC),
		},
		Character: export.CharacterFile{
			ID:     99,
			Name:   "Alice",
			System: gameSystem,
			HP:     12,
			MaxHP:  12,
			Sheet:  homeSheet(t, func(s *dnd5e.Sheet) { s.Coins.GP = 5 }),
		},
		Maps: &export.HomeIDMaps{},
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

// uploadBringBack posts a file to the character's bring-back route.
func (e testEnv) uploadBringBack(
	t *testing.T,
	body []byte,
	cookie string,
) *httptest.ResponseRecorder {
	t.Helper()
	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	part, err := mw.CreateFormFile(formFile, "bring-back.zip")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = part.Write(body)
	_ = mw.Close()
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		route.Path(route.CharacterBringBack, e.characterID),
		&buf,
	)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("HX-Request", "true")
	req.RemoteAddr = lanAddr
	req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: cookie})
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	return w
}

// waitingFile is the one file waiting, with the review's stamp.
func (e testEnv) waitingFile(t *testing.T) (homeFile, string) {
	t.Helper()
	list := e.app.homeFiles.list()
	if len(list) != 1 {
		t.Fatalf("%d files waiting, want 1", len(list))
	}
	rv, err := e.app.reviewBringBack(t.Context(), e.q(), list[0], func(string) string { return "" })
	if err != nil {
		t.Fatal(err)
	}
	return list[0], rv.was
}

func (e testEnv) fileCount(t *testing.T) int {
	t.Helper()
	entries, err := os.ReadDir(filepath.Join(e.app.DataDir, HomeFilesDir))
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		t.Fatal(err)
	}
	return len(entries)
}

func TestBringBackUploadRefusals(t *testing.T) {
	tests := []struct {
		name   string
		setup  func(*testing.T, testEnv, sqlc.TakeHome)
		mutate func(*export.TakeHome)
		zip    func(*testing.T) []byte // replaces the file entirely
		want   string
	}{
		{
			name: "a take-home file",
			zip:  func(t *testing.T) []byte { return sampleTakeHome(t, nil) },
			want: "not a bring-back file",
		},
		{
			name:   "another table",
			mutate: func(th *export.TakeHome) { th.Manifest.ServerID = "elsewhere" },
			want:   "belongs to another GM",
		},
		{
			name:   "another character",
			mutate: func(th *export.TakeHome) { th.Manifest.CharacterID++ },
			want:   "another character",
		},
		{
			name:   "unknown take-out",
			mutate: func(th *export.TakeHome) { th.Manifest.Token = "nope" },
			want:   "take-out is unknown",
		},
		{
			name: "take-out of another character",
			setup: func(t *testing.T, e testEnv, _ sqlc.TakeHome) {
				other, err := e.q().
					CreateCharacterWithSheet(t.Context(), sqlc.CreateCharacterWithSheetParams{CampaignID: e.campaignID, Name: "Bob", System: gameSystem, SheetJson: "{}"})
				if err != nil {
					t.Fatal(err)
				}
				if _, err := e.q().
					CreateTakeHome(t.Context(), sqlc.CreateTakeHomeParams{CharacterID: other.ID, Token: "bobs", BaseJson: "{}"}); err != nil {
					t.Fatal(err)
				}
			},
			mutate: func(th *export.TakeHome) { th.Manifest.Token = "bobs" },
			want:   "take-out is unknown",
		},
		{
			name: "already brought back",
			setup: func(t *testing.T, e testEnv, rec sqlc.TakeHome) {
				if _, err := e.q().ReturnTakeHome(t.Context(), rec.ID); err != nil {
					t.Fatal(err)
				}
			},
			want: "already brought back",
		},
		{
			name:   "newer sheet",
			mutate: func(th *export.TakeHome) { th.Character.Sheet = json.RawMessage(`{"brand_new_field":1}`) },
			want:   "needs updating",
		},
		{name: "invalid entry", mutate: func(th *export.TakeHome) {
			th.Sources = []export.TakeHomeSource{{SourceKey: "b"}}
			th.Entries = []export.TakeHomeEntry{
				{
					Kind:      enum.KindSpell,
					Slug:      "s",
					Name:      "S",
					SourceKey: "b",
					Data:      json.RawMessage(`[1]`),
				},
			}
		}, want: "is not valid"},
		{name: "homebrew of a kind players cannot write", mutate: func(th *export.TakeHome) {
			th.Homebrew = []export.TakeHomeEntry{
				{
					Kind:      enum.KindFeat,
					Slug:      "f",
					Name:      "F",
					SourceKey: "homebrew",
					Data:      json.RawMessage(`{}`),
				},
			}
		}, want: "cannot write"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := newTestEnv(t)
			rec := e.takeOut(t)
			if tt.setup != nil {
				tt.setup(t, e, rec)
			}
			body := e.bringBackZip(t, rec, tt.mutate)
			if tt.zip != nil {
				body = tt.zip(t)
			}
			w := e.uploadBringBack(t, body, e.token(t, enum.RoleAdmin))
			if w.Code != http.StatusUnprocessableEntity ||
				!strings.Contains(w.Body.String(), tt.want) {
				t.Errorf("status %d, body %q; want 422 containing %q", w.Code, w.Body, tt.want)
			}
			if n := len(e.app.homeFiles.list()); n != 0 || e.fileCount(t) != 0 {
				t.Errorf("%d entries and %d files left behind", n, e.fileCount(t))
			}
		})
	}
}

func TestBringBackPlayerUpload(t *testing.T) {
	e := newTestEnv(t)
	rec := e.takeOut(t)
	player := e.token(t, enum.RolePlayer)
	other := e.otherPlayer(t)
	// htmx gets every refusal as 422 (a toast).
	if w := e.uploadBringBack(
		t,
		e.bringBackZip(t, rec, nil),
		other,
	); w.Code != http.StatusUnprocessableEntity ||
		len(e.app.homeFiles.list()) != 0 {
		t.Errorf(
			"another player's upload: status %d, %d files waiting",
			w.Code,
			len(e.app.homeFiles.list()),
		)
	}
	w := e.uploadBringBack(t, e.bringBackZip(t, rec, nil), player)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Sent to the GM") {
		t.Fatalf("status %d: %s", w.Code, w.Body)
	}
	f, _ := e.waitingFile(t)
	if !f.byPlayer || f.characterName != "Alice" || e.fileCount(t) != 1 {
		t.Errorf("entry %+v, %d files", f, e.fileCount(t))
	}

	// The GM's popup shows it, with a way to review and to reject.
	var html bytes.Buffer
	if err := views.PendingRequests(e.app.pendingViews()).Render(t.Context(), &html); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Alice is back from home", route.Path(route.CampaignHomeFile, e.campaignID, f.id), route.Path(route.CampaignHomeFileReject, e.campaignID, f.id)} {
		if !strings.Contains(html.String(), want) {
			t.Errorf("popup lacks %q", want)
		}
	}

	// A newer file of the same character replaces it.
	if w := e.uploadBringBack(t, e.bringBackZip(t, rec, nil), player); w.Code != http.StatusOK {
		t.Fatalf("second upload: status %d", w.Code)
	}
	if again, _ := e.waitingFile(t); again.id == f.id || e.fileCount(t) != 1 {
		t.Errorf("the old file stayed: %d files", e.fileCount(t))
	}

	// The GM's own upload is listed too, under the character's name.
	if w := e.uploadBringBack(
		t,
		e.bringBackZip(t, rec, nil),
		e.token(t, enum.RoleAdmin),
	); w.Code != http.StatusNoContent {
		t.Fatalf("GM upload: status %d", w.Code)
	}
	if g := e.app.pendingViews(); len(g) != 1 || len(g[0].BringBacks) != 1 ||
		g[0].BringBacks[0].Name != "Alice" {
		t.Errorf("popup after the GM's upload: %+v", g)
	}
}

func TestHomeFilesCap(t *testing.T) {
	var s homeFiles
	for i := range maxHomeFiles {
		if _, err := s.put(
			homeFile{kind: export.KindBringBack, characterID: int64(i + 1)},
		); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := s.put(
		homeFile{kind: export.KindBringBack, characterID: 100},
	); err == nil ||
		!strings.Contains(err.Error(), "Too many files") {
		t.Errorf("the 11th file: %v", err)
	}
	// A replacement does not count against the cap.
	if _, err := s.put(homeFile{kind: export.KindBringBack, characterID: 1}); err != nil {
		t.Errorf("replacing: %v", err)
	}
}

// TestHomeFilesClaim: one Add at a time per file; a released file can be claimed again.
func TestHomeFilesClaim(t *testing.T) {
	var s homeFiles
	id, err := s.put(homeFile{kind: export.KindNewCharacter})
	if err != nil {
		t.Fatal(err)
	}
	if !s.claim(id) || s.claim(id) {
		t.Fatal("a file must be claimed once")
	}
	s.release(id)
	if !s.claim(id) {
		t.Error("a released file cannot be claimed")
	}
	if s.claim("gone") {
		t.Error("claimed a file that is not waiting")
	}
}

func TestHomeFileAccess(t *testing.T) {
	e := newTestEnv(t)
	rec := e.takeOut(t)
	gm, player := e.token(t, enum.RoleAdmin), e.token(t, enum.RolePlayer)
	if w := e.uploadBringBack(t, e.bringBackZip(t, rec, nil), player); w.Code != http.StatusOK {
		t.Fatal(w.Body)
	}
	f, was := e.waitingFile(t)
	page := route.Path(route.CampaignHomeFile, e.campaignID, f.id)
	tests := []struct {
		name, method, path, cookie string
		form                       url.Values
		want                       int
	}{
		{"gm reads", http.MethodGet, page, gm, nil, http.StatusOK},
		{"player reads", http.MethodGet, page, player, nil, http.StatusForbidden},
		{"visitor reads", http.MethodGet, page, "", nil, http.StatusForbidden},
		{
			"gm, another campaign",
			http.MethodGet,
			route.Path(route.CampaignHomeFile, e.campaignID+1, f.id),
			gm,
			nil,
			http.StatusNotFound,
		},
		{
			"gm, unknown file",
			http.MethodGet,
			route.Path(route.CampaignHomeFile, e.campaignID, "nope"),
			gm,
			nil,
			http.StatusNotFound,
		},
		{
			"player accepts",
			http.MethodPost,
			page + route.Accept,
			player,
			url.Values{formWas: {was}},
			http.StatusForbidden,
		},
		{
			"visitor accepts",
			http.MethodPost,
			page + route.Accept,
			"",
			url.Values{formWas: {was}},
			http.StatusForbidden,
		},
		{"player rejects", http.MethodPost, page + route.Reject, player, nil, http.StatusForbidden},
		{"visitor rejects", http.MethodPost, page + route.Reject, "", nil, http.StatusForbidden},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if w := e.do(
				t,
				tt.method,
				tt.path,
				tt.cookie,
				tt.form,
				tt.method == http.MethodPost,
			); w.Code != tt.want {
				t.Errorf("status %d, want %d", w.Code, tt.want)
			}
		})
	}
	if len(e.app.homeFiles.list()) != 1 {
		t.Error("a refused request dropped the file")
	}
	// The page names the character and what changed.
	w := e.do(t, http.MethodGet, page, gm, nil, false)
	for _, want := range []string{"Alice", "Changes", "Gear &amp; coins", "Unchanged"} {
		if !strings.Contains(w.Body.String(), want) {
			t.Errorf("review lacks %q", want)
		}
	}
}

func TestBringBackAccept(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	e.takeOut(t) // an older trip: accepting the newer one closes it too
	rec := e.takeOut(t)
	gm := e.token(t, enum.RoleAdmin)
	zipBody := e.bringBackZip(t, rec, func(th *export.TakeHome) {
		th.Character.Sheet = homeSheet(t, func(s *dnd5e.Sheet) {
			s.Classes[0].Level = 2
			s.Spells = []dnd5e.Spell{{Slug: "shield", Name: "Shield", Level: 1}}
			s.Coins.GP = 5
		})
		th.Character.MaxHP = 12
		th.LevelUps = []export.TakeHomeLevelUp{
			{
				ID:        1,
				MaxHP:     8,
				Sheet:     homeSheet(t, nil),
				CreatedAt: time.Date(2026, 11, 25, 8, 0, 0, 0, time.UTC),
			},
		}
	})
	// At the table meanwhile: coins spent (the home spent differently).
	table := homeTableSheet(t, func(s *dnd5e.Sheet) { s.Coins.GP = 30 })
	if err := e.q().
		UpdateCharacter(ctx, sqlc.UpdateCharacterParams{ID: e.characterID, Name: "Alice", SheetJson: table}); err != nil {
		t.Fatal(err)
	}
	if w := e.uploadBringBack(t, zipBody, gm); w.Code != http.StatusNoContent {
		t.Fatalf("GM upload: status %d: %s", w.Code, w.Body)
	}
	f, was := e.waitingFile(t)
	review := route.Path(route.CampaignHomeFile, e.campaignID, f.id)

	// Gear is changed on both sides: the page asks which, and offers both.
	w := e.do(t, http.MethodGet, review, gm, nil, false)
	for _, want := range []string{"Both", "30 gp", "pick_gear"} {
		if !strings.Contains(w.Body.String(), want) {
			t.Errorf("review lacks %q", want)
		}
	}

	accept := url.Values{formWas: {was}, views.PickParam(dnd5e.GroupGear): {dnd5e.SideTable}}
	w = e.do(t, http.MethodPost, review+route.Accept, gm, accept, true)
	if w.Code != http.StatusOK ||
		!strings.Contains(w.Body.String(), "Alice is back at the table.") {
		t.Fatalf("accept: status %d: %s", w.Code, w.Body)
	}
	ch, err := e.q().GetCharacter(ctx, e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	s, err := dnd5e.ParseSheet(ch.SheetJson)
	if err != nil {
		t.Fatal(err)
	}
	if ch.MaxHp != 12 || ch.Hp != 12 || s.Level() != 2 || len(s.Spells) != 1 || s.Coins.GP != 30 {
		t.Errorf(
			"character: hp %d/%d, level %d, %d spells, %d gp",
			ch.Hp,
			ch.MaxHp,
			s.Level(),
			len(s.Spells),
			s.Coins.GP,
		)
	}
	levelUps, err := e.q().ListLevelUps(ctx, e.characterID)
	if err != nil || len(levelUps) != 1 || levelUps[0].MaxHp != 8 {
		t.Errorf("level-ups %v (%v), want the file's", levelUps, err)
	}
	if _, err := e.q().LatestOpenTakeHome(ctx, e.characterID); !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("the take-out is still open: %v", err)
	}
	if len(e.app.homeFiles.list()) != 0 || e.fileCount(t) != 0 {
		t.Error("the file was kept after accepting")
	}
	// The same file cannot be used again, from the page or as a new upload.
	if w := e.do(
		t,
		http.MethodPost,
		review+route.Accept,
		gm,
		accept,
		true,
	); w.Code != http.StatusUnprocessableEntity ||
		!strings.Contains(w.Body.String(), "no longer waiting") {
		t.Errorf("second accept: status %d: %s", w.Code, w.Body)
	}
	if w := e.uploadBringBack(
		t,
		zipBody,
		gm,
	); w.Code != http.StatusUnprocessableEntity ||
		!strings.Contains(w.Body.String(), "already brought back") {
		t.Errorf("second upload: status %d: %s", w.Code, w.Body)
	}
}

// homeTableSheet is the table's sheet after the take-out, changed by change.
func homeTableSheet(t *testing.T, change func(*dnd5e.Sheet)) string {
	t.Helper()
	s, err := dnd5e.ParseSheet("{}")
	if err != nil {
		t.Fatal(err)
	}
	change(&s)
	raw, err := s.JSON()
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func TestBringBackReject(t *testing.T) {
	for _, fromPage := range []bool{false, true} {
		e := newTestEnv(t)
		rec := e.takeOut(t)
		gm := e.token(t, enum.RoleAdmin)
		if w := e.uploadBringBack(
			t,
			e.bringBackZip(t, rec, nil),
			e.token(t, enum.RolePlayer),
		); w.Code != http.StatusOK {
			t.Fatal(w.Body)
		}
		f, _ := e.waitingFile(t)
		form := url.Values{}
		if fromPage {
			form.Set(formPage, "1")
		}
		w := e.do(
			t,
			http.MethodPost,
			route.Path(route.CampaignHomeFileReject, e.campaignID, f.id),
			gm,
			form,
			true,
		)
		switch {
		case fromPage && w.Header().Get("HX-Redirect") != route.Path(route.Character, e.characterID):
			t.Errorf("page reject: redirect %q", w.Header().Get("HX-Redirect"))
		case !fromPage && !strings.Contains(w.Body.String(), "Rejected. The take-out stays open."):
			t.Errorf("popup reject: status %d: %s", w.Code, w.Body)
		}
		if _, err := e.q().LatestOpenTakeHome(t.Context(), e.characterID); err != nil {
			t.Errorf("the take-out closed on reject: %v", err)
		}
		if len(e.app.homeFiles.list()) != 0 || e.fileCount(t) != 0 {
			t.Error("the file was kept after rejecting")
		}
	}
}

func TestBringBackStale(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	rec := e.takeOut(t)
	gm := e.token(t, enum.RoleAdmin)
	if w := e.uploadBringBack(t, e.bringBackZip(t, rec, nil), gm); w.Code != http.StatusNoContent {
		t.Fatal(w.Body)
	}
	f, was := e.waitingFile(t)
	if err := e.q().
		UpdateCharacter(ctx, sqlc.UpdateCharacterParams{ID: e.characterID, Name: "Alice", Hp: 3, MaxHp: 9, SheetJson: "{}"}); err != nil {
		t.Fatal(err)
	}
	w := e.do(
		t,
		http.MethodPost,
		route.Path(route.CampaignHomeFileAccept, e.campaignID, f.id),
		gm,
		url.Values{formWas: {was}},
		true,
	)
	if w.Code != http.StatusOK ||
		!strings.Contains(w.Body.String(), "changed at the table meanwhile") {
		t.Fatalf("status %d: %s", w.Code, w.Body)
	}
	ch, err := e.q().GetCharacter(ctx, e.characterID)
	if err != nil || ch.Hp != 3 || ch.MaxHp != 9 {
		t.Errorf("a stale accept wrote: %+v (%v)", ch, err)
	}
	if _, err := e.q().LatestOpenTakeHome(ctx, e.characterID); err != nil {
		t.Errorf("the take-out closed: %v", err)
	}
	if len(e.app.homeFiles.list()) != 1 {
		t.Error("the file was dropped")
	}
}

// TestBringBackSources: a source added at home must be chosen, a declined one
// leaves the build from the fallback or blocks Accept, an allowed one is
// imported and ticked.
func TestBringBackSources(t *testing.T) {
	pinned := func(t *testing.T) json.RawMessage {
		return homeSheet(t, func(s *dnd5e.Sheet) {
			s.Classes[0].Level, s.Classes[0].Subclass, s.Classes[0].SubclassSource = 2, "Bladesinging", "newbook"
			s.Spells = []dnd5e.Spell{{Slug: "boom", Name: "Booming Blade", Source: "newbook"}}
		})
	}
	withSource := func(t *testing.T, fallback bool) func(*export.TakeHome) {
		return func(th *export.TakeHome) {
			th.Character.Sheet, th.Character.MaxHP = pinned(t), 12
			th.Sources = []export.TakeHomeSource{
				{SourceKey: "newbook", Edition: "2024", License: "L", Attribution: "A"},
			}
			th.Entries = []export.TakeHomeEntry{
				{
					Kind:        enum.KindSpell,
					Slug:        "boom",
					Name:        "Booming Blade",
					SourceKey:   "newbook",
					Data:        json.RawMessage(`{}`),
					ContentHash: "h",
				},
			}
			th.LevelUps = []export.TakeHomeLevelUp{
				{ID: 1, MaxHP: 8, Sheet: homeSheet(t, nil), CreatedAt: time.Now().UTC()},
			}
			if fallback {
				th.Fallback = &export.TakeHomeFallback{
					MaxHP: 7,
					Sheet: homeSheet(t, nil),
					LevelUps: []export.TakeHomeLevelUp{
						{ID: 2, MaxHP: 3, Sheet: homeSheet(t, nil), CreatedAt: time.Now().UTC()},
					},
				}
			}
		}
	}
	tests := []struct {
		name     string
		fallback bool
		choice   string
		wantCode int
		wantText string
		check    func(t *testing.T, e testEnv)
	}{
		{
			name:     "not chosen",
			choice:   "",
			wantCode: http.StatusUnprocessableEntity,
			wantText: "Choose Allow or Decline",
		},
		{
			name:     "declined, build pins it, no fallback",
			choice:   views.SourceDecline,
			wantCode: http.StatusUnprocessableEntity,
			wantText: "has no fallback without it",
		},
		{
			name:     "declined with a fallback",
			fallback: true,
			choice:   views.SourceDecline,
			wantCode: http.StatusOK,
			wantText: "is back at the table",
			check: func(t *testing.T, e testEnv) {
				ch, err := e.q().GetCharacter(t.Context(), e.characterID)
				if err != nil {
					t.Fatal(err)
				}
				s, _ := dnd5e.ParseSheet(ch.SheetJson)
				levelUps, _ := e.q().ListLevelUps(t.Context(), e.characterID)
				if s.Level() != 1 || ch.MaxHp != 7 || s.BuildPinsSource("newbook") ||
					len(s.Spells) != 0 ||
					len(levelUps) != 1 ||
					levelUps[0].MaxHp != 3 {
					t.Errorf(
						"level %d, max HP %d, %d spells, level-ups %v: want the fallback",
						s.Level(),
						ch.MaxHp,
						len(s.Spells),
						levelUps,
					)
				}
				if _, err := e.q().
					GetCompendiumSource(t.Context(), "newbook"); !errors.Is(
					err,
					sql.ErrNoRows,
				) {
					t.Errorf("a declined source was imported: %v", err)
				}
			},
		},
		{
			name:     "allowed",
			choice:   views.SourceAllow,
			wantCode: http.StatusOK,
			wantText: "is back at the table",
			check: func(t *testing.T, e testEnv) {
				src, err := e.q().GetCompendiumSource(t.Context(), "newbook")
				if err != nil || src.Path != bringBackSourcePath {
					t.Fatalf("source %+v (%v)", src, err)
				}
				if _, err := e.q().
					GetCompendiumEntryBySlug(t.Context(), sqlc.GetCompendiumEntryBySlugParams{Kind: enum.KindSpell, Slug: "boom"}); err != nil {
					t.Errorf("the source's entry was not imported: %v", err)
				}
				keys, err := e.q().ListCampaignSources(t.Context(), e.campaignID)
				if err != nil || !slices.Contains(keys, "newbook") {
					t.Errorf("campaign sources %v (%v)", keys, err)
				}
				ch, _ := e.q().GetCharacter(t.Context(), e.characterID)
				if s, _ := dnd5e.ParseSheet(ch.SheetJson); s.Level() != 2 || len(s.Spells) != 1 {
					t.Errorf(
						"the home build was not kept: level %d, %d spells",
						s.Level(),
						len(s.Spells),
					)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := newTestEnv(t)
			rec := e.takeOut(t)
			gm := e.token(t, enum.RoleAdmin)
			if w := e.uploadBringBack(
				t,
				e.bringBackZip(t, rec, withSource(t, tt.fallback)),
				gm,
			); w.Code != http.StatusNoContent {
				t.Fatalf("upload: status %d: %s", w.Code, w.Body)
			}
			f, was := e.waitingFile(t)
			review := route.Path(route.CampaignHomeFile, e.campaignID, f.id)
			form := url.Values{formWas: {was}}
			if tt.choice != "" {
				form.Set(views.SourceParam("newbook"), tt.choice)
			}
			// The page names what the source is used for and what declining does.
			page := e.do(t, http.MethodGet, review+"?"+form.Encode(), gm, nil, false).Body.String()
			for _, want := range []string{"newbook", "Used: subclass Bladesinging, spell Booming Blade"} {
				if !strings.Contains(page, want) {
					t.Errorf("review lacks %q", want)
				}
			}
			w := e.do(t, http.MethodPost, review+route.Accept, gm, form, true)
			if w.Code != tt.wantCode || !strings.Contains(w.Body.String(), tt.wantText) {
				t.Fatalf(
					"status %d, body %q; want %d containing %q",
					w.Code,
					w.Body,
					tt.wantCode,
					tt.wantText,
				)
			}
			if tt.wantCode != http.StatusOK {
				if _, err := e.q().LatestOpenTakeHome(t.Context(), e.characterID); err != nil {
					t.Errorf("a refused accept closed the take-out: %v", err)
				}
				return
			}
			tt.check(t, e)
		})
	}
}

func TestKeepsBodyBringBack(t *testing.T) {
	for _, pattern := range []string{route.CharacterBringBack, route.CampaignHomeFileAccept, route.CampaignHomeFileReject} {
		if keepsBody(path.Base(pattern), pattern) {
			t.Errorf("the log keeps the body of %s", pattern)
		}
	}
	if !keepsBody(route.CharacterCmd, route.CharacterCmd) {
		t.Error("the log lost the commands' bodies")
	}
}
