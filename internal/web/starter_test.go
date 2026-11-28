package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/route"
)

// seedStarterContent gives the test campaign one book with a spell and a
// monster, a GM homebrew feat and a player's own homebrew spell.
func (e testEnv) seedStarterContent(t *testing.T) {
	t.Helper()
	ctx := t.Context()
	q := e.q()
	e.homeBook(t, "bk", "2024")
	for _, en := range []struct {
		kind enum.CompendiumKind
		slug string
	}{{enum.KindSpell, "bolt"}, {enum.KindMonster, "orc"}} {
		if err := q.UpsertCompendiumEntry(
			ctx,
			sqlc.UpsertCompendiumEntryParams{
				Kind:        en.kind,
				Slug:        en.slug,
				Name:        en.slug,
				SourceKey:   "bk",
				DataJson:    "{}",
				ContentHash: "h-" + en.slug,
			},
		); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := q.CreateHomebrew(
		ctx,
		sqlc.CreateHomebrewParams{
			Kind:        enum.KindFeat,
			Slug:        "gm-feat",
			Name:        "Table feat",
			DataJson:    "{}",
			ContentHash: "h-gm",
		},
	); err != nil {
		t.Fatal(err)
	}
	if _, err := q.CreateHomebrew(
		ctx,
		sqlc.CreateHomebrewParams{
			Kind:             enum.KindSpell,
			Slug:             "zap",
			Name:             "Zap",
			DataJson:         "{}",
			ContentHash:      "h-zap",
			OwnerCharacterID: nullID(e.characterID),
			Scoped:           true,
		},
	); err != nil {
		t.Fatal(err)
	}
	e.allowEntrySources(t)
}

func readStarter(t *testing.T, body []byte) *export.TakeHomeFile {
	t.Helper()
	f, err := export.ReadTakeHome(bytes.NewReader(body), int64(len(body)), export.KindStarter)
	if err != nil {
		t.Fatal(err)
	}
	return f
}

func TestStarterDownloadRoutes(t *testing.T) {
	e := newTestEnv(t)
	e.seedStarterContent(t)
	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)
	gmPath := route.Path(route.CampaignStarter, e.campaignID)
	joinPath := route.Path(route.JoinStarter, e.campaignID)
	setOpenJoin := func(on bool) {
		t.Helper()
		if err := e.q().
			SetCampaignOpenJoin(t.Context(), sqlc.SetCampaignOpenJoinParams{OpenJoin: on, ID: e.campaignID}); err != nil {
			t.Fatal(err)
		}
	}

	// The join page offers it, and its route serves it, only while open join is on.
	setOpenJoin(false)
	if w := e.do(t, http.MethodGet, joinPath, "", nil, false); w.Code != http.StatusNotFound {
		t.Errorf("join route with open join off: %d, want 404", w.Code)
	}
	if w := e.do(
		t,
		http.MethodGet,
		route.Path(route.JoinCampaign, e.campaignID),
		"",
		nil,
		false,
	); strings.Contains(
		w.Body.String(),
		joinPath,
	) {
		t.Error("the join page offers the starter with open join off")
	}
	setOpenJoin(true)
	w := e.do(t, http.MethodGet, joinPath, "", nil, false)
	if w.Code != http.StatusOK ||
		!strings.Contains(w.Header().Get("Content-Disposition"), "attachment") {
		t.Fatalf(
			"join route with open join on: %d %q",
			w.Code,
			w.Header().Get("Content-Disposition"),
		)
	}
	if w := e.do(
		t,
		http.MethodGet,
		route.Path(route.JoinCampaign, e.campaignID),
		"",
		nil,
		false,
	); !strings.Contains(
		w.Body.String(),
		joinPath,
	) {
		t.Error("the join page lacks the starter link with open join on")
	}
	if w := e.do(
		t,
		http.MethodGet,
		route.Path(route.JoinStarter, 999),
		"",
		nil,
		false,
	); w.Code != http.StatusNotFound {
		t.Errorf("join route of a missing campaign: %d", w.Code)
	}

	// The GM route is GM-only, whatever open join says.
	setOpenJoin(false)
	for name, cookie := range map[string]string{"visitor": "", "player": player} {
		if w := e.do(
			t,
			http.MethodGet,
			gmPath,
			cookie,
			nil,
			false,
		); w.Code < http.StatusBadRequest {
			t.Errorf("%s got the GM starter route: %d", name, w.Code)
		}
	}
	if w := e.do(t, http.MethodGet, gmPath, gm, nil, false); w.Code != http.StatusOK {
		t.Fatalf("GM: %d %s", w.Code, w.Body)
	}
}

// The file holds what any player of the campaign sees: its books' entries and
// the GM's homebrew, never a monster or a character's own homebrew, and no character.
func TestStarterContent(t *testing.T) {
	e := newTestEnv(t)
	e.seedStarterContent(t)
	gm := e.token(t, enum.RoleAdmin)
	w := e.do(t, http.MethodGet, route.Path(route.CampaignStarter, e.campaignID), gm, nil, false)
	if w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	f := readStarter(t, w.Body.Bytes())
	if m := f.Manifest; m.Kind != export.KindStarter || m.Campaign.ID != e.campaignID ||
		m.Campaign.Name != "Test" ||
		m.ServerID == "" ||
		m.Token != "" ||
		m.CharacterID != 0 {
		t.Errorf("manifest %+v", m)
	}
	if len(f.Sources) != 1 || f.Sources[0].SourceKey != "bk" {
		t.Errorf("sources %+v", f.Sources)
	}
	for _, tt := range []struct {
		slug string
		want bool
	}{{"bolt", true}, {"gm-feat", true}, {"orc", false}, {"zap", false}} {
		if got := slices.ContainsFunc(
			f.Entries,
			func(en export.TakeHomeEntry) bool { return en.Slug == tt.slug },
		); got != tt.want {
			t.Errorf("entry %q in the file: %v, want %v", tt.slug, got, tt.want)
		}
	}
}

// addNewCharacterFile uploads a new-character file as the GM and, when it is
// accepted for review, adds it. It returns the status and body of the last answer.
func (e testEnv) addNewCharacterFile(
	t *testing.T,
	gm string,
	body []byte,
) *httptest.ResponseRecorder {
	t.Helper()
	add := route.Path(route.CampaignNewCharacter, e.campaignID)
	up := e.doFile(t, add, gm, nil, formFile, "x.zip", string(body))
	if up.Code != http.StatusSeeOther {
		return up
	}
	list := e.app.homeFiles.list()
	f := list[len(list)-1]
	rv, err := e.app.reviewHomeFile(t.Context(), e.q(), f, func(string) string { return "" })
	if err != nil {
		t.Fatal(err)
	}
	return e.do(
		t,
		http.MethodPost,
		route.Path(route.CampaignHomeFile, e.campaignID, f.id)+route.Accept,
		gm,
		url.Values{formWas: {rv.was}},
		true,
	)
}

func TestMadeAtHomeTokenAtTable(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, enum.RoleAdmin)
	withToken := func(tok string) []byte {
		return e.newCharacterZip(t, func(th *export.TakeHome) { th.Manifest.MadeAtHome = tok })
	}

	w := e.addNewCharacterFile(t, gm, withToken("tok-1"))
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Bryn joined the campaign.") {
		t.Fatalf("first add: %d %s", w.Code, w.Body)
	}
	first, err := e.q().GetNewCharacterByToken(ctx, "tok-1")
	if err != nil || first.Name != "Bryn" {
		t.Fatalf("token holder %+v, %v", first, err)
	}

	// The same file again is refused, at upload.
	w = e.addNewCharacterFile(t, gm, withToken("tok-1"))
	if w.Code != http.StatusUnprocessableEntity ||
		!strings.Contains(w.Body.String(), "Bryn was already added from this file") {
		t.Fatalf("duplicate: %d %s", w.Code, w.Body)
	}
	if n := e.count(t, `SELECT count(*) FROM characters WHERE name = 'Bryn'`); n != 1 {
		t.Errorf("%d Bryns, want 1", n)
	}

	// A file with no token is not tied to anything.
	if w = e.addNewCharacterFile(t, gm, withToken("")); w.Code != http.StatusOK {
		t.Fatalf("no token: %d %s", w.Code, w.Body)
	}

	// Once the holder is deleted, the token goes to the new character.
	if err := e.q().SoftDeleteCharacter(ctx, first.CharacterID); err != nil {
		t.Fatal(err)
	}
	if w = e.addNewCharacterFile(t, gm, withToken("tok-1")); w.Code != http.StatusOK {
		t.Fatalf("after delete: %d %s", w.Code, w.Body)
	}
	second, err := e.q().GetNewCharacterByToken(ctx, "tok-1")
	if err != nil || second.CharacterID == first.CharacterID || second.DeletedAt.Valid {
		t.Errorf("token holder after takeover %+v, %v", second, err)
	}
	if n := e.count(t, `SELECT count(*) FROM new_character_tokens WHERE token = 'tok-1'`); n != 1 {
		t.Errorf("%d rows hold the token", n)
	}
}

func TestTakeHomeCarriesMadeAtHomeToken(t *testing.T) {
	e := newTestEnv(t)
	e.seedTakeHome(t)
	read := func() *export.TakeHomeFile {
		w := e.takeHomeRequest(t, http.MethodPost, loopbackAddr, "")
		if w.Code != http.StatusOK {
			t.Fatalf("%d %s", w.Code, w.Body)
		}
		return readZip(t, w.Body.Bytes())
	}
	if got := read().Manifest.MadeAtHome; got != "" {
		t.Errorf("a table character's take-home carries %q", got)
	}
	if err := e.q().
		CreateNewCharacterToken(t.Context(), sqlc.CreateNewCharacterTokenParams{CharacterID: e.characterID, Token: "tok-9"}); err != nil {
		t.Fatal(err)
	}
	if got := read().Manifest.MadeAtHome; got != "tok-9" {
		t.Errorf("made-at-home token %q, want tok-9", got)
	}
}

// starterZip is a starter file as the table writes it.
func starterZip(t *testing.T, mutate func(*export.TakeHome)) []byte {
	t.Helper()
	th := export.TakeHome{
		Manifest: export.TakeHomeManifest{
			Format:     export.TakeHomeFormat,
			Kind:       export.KindStarter,
			AppVersion: "v-table",
			ServerID:   "srv",
			CreatedAt:  time.Date(2026, 11, 24, 18, 0, 0, 0, time.UTC),
			Campaign: export.TakeHomeCampaign{
				ID:      3,
				Name:    "Table",
				System:  "dnd5e",
				Edition: "2024",
			},
		},
		Sources: []export.TakeHomeSource{
			{SourceKey: "src1", Edition: "2024", Priority: 150},
			{SourceKey: "src0", Edition: "2024", Priority: 100},
		},
		Entries: []export.TakeHomeEntry{
			{
				Kind:        enum.KindSpell,
				Slug:        "bolt",
				Name:        "Bolt",
				SourceKey:   "src1",
				Data:        json.RawMessage(`{}`),
				ContentHash: "h2",
			},
			{
				Kind:        enum.KindFeat,
				Slug:        "gm-feat",
				Name:        "Table feat",
				SourceKey:   homebrewSource,
				Data:        json.RawMessage(`{}`),
				ContentHash: "h3",
			},
		},
	}
	if mutate != nil {
		mutate(&th)
	}
	var buf bytes.Buffer
	if err := export.WriteTakeHome(&buf, th, t.TempDir()); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func (e testEnv) importStarterFile(t *testing.T, mutate func(*export.TakeHome)) {
	t.Helper()
	if w := e.importRequest(
		t,
		starterZip(t, mutate),
		"",
		nil,
	); w.Code != http.StatusNoContent ||
		w.Header().Get("HX-Redirect") != route.AtHome {
		t.Fatalf("starter import: %d %q: %s", w.Code, w.Header().Get("HX-Redirect"), w.Body)
	}
}

func (e testEnv) starterCampaign(t *testing.T) sqlc.ListHomeStartersRow {
	t.Helper()
	rows, err := e.q().ListHomeStarters(t.Context())
	if err != nil || len(rows) != 1 {
		t.Fatalf("starters %+v, %v", rows, err)
	}
	return rows[0]
}

func TestStarterRefreshInPlace(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	q := e.q()
	e.importStarterFile(t, nil)
	first := e.starterCampaign(t)
	if got, _ := q.ListCampaignSources(
		ctx,
		first.CampaignID,
	); !slices.Equal(
		slices.Sorted(slices.Values(got)),
		[]string{"src0", "src1"},
	) {
		t.Errorf("sources %v", got)
	}
	if id, err := e.app.Store.GetHomeServerID(ctx); err != nil || id != "srv" {
		t.Errorf("server id %q, %v: a starter alone must give it", id, err)
	}

	// The page lists the starter with its books in priority order.
	page := e.get(t, route.AtHome, "").Body.String()
	for _, want := range []string{"Campaigns from your GM", "2 books", "<li>src0, 2024</li><li>src1, 2024</li>", "New character…", "Remove the Table starter?"} {
		if !strings.Contains(page, want) {
			t.Errorf("At home page lacks %q", want)
		}
	}

	// A second opening: renamed, other books, the homebrew renamed (one copy).
	e.importStarterFile(t, func(th *export.TakeHome) {
		th.Manifest.Campaign.Name = "Table 2"
		th.Sources = []export.TakeHomeSource{{SourceKey: "src2", Edition: "2024", Priority: 120}}
		th.Entries = []export.TakeHomeEntry{
			{
				Kind:        enum.KindSpell,
				Slug:        "bolt2",
				Name:        "Bolt 2",
				SourceKey:   "src2",
				Data:        json.RawMessage(`{}`),
				ContentHash: "h5",
			},
			{
				Kind:        enum.KindFeat,
				Slug:        "gm-feat",
				Name:        "Table feat v2",
				SourceKey:   homebrewSource,
				Data:        json.RawMessage(`{}`),
				ContentHash: "h4",
			},
		}
	})
	second := e.starterCampaign(t)
	if second.CampaignID != first.CampaignID || second.Name != "Table 2" {
		t.Errorf("starter %+v after the refresh, was %+v", second, first)
	}
	if n := e.count(t, `SELECT count(*) FROM campaigns`); n != 1 {
		t.Errorf("%d campaigns, want the one starter", n)
	}
	if got, _ := q.ListCampaignSources(
		ctx,
		first.CampaignID,
	); !slices.Equal(
		got,
		[]string{"src2"},
	) {
		t.Errorf("sources after the refresh %v, want exactly src2", got)
	}
	if n := e.count(t, `SELECT count(*) FROM compendium_entries WHERE slug = 'gm-feat'`); n != 1 {
		t.Errorf("%d copies of the homebrew, want 1", n)
	}
	if n := e.count(
		t,
		`SELECT count(*) FROM compendium_entries WHERE slug = 'gm-feat' AND name = 'Table feat v2'`,
	); n != 1 {
		t.Error("the homebrew was not updated in place")
	}
	if n := e.count(t, `SELECT count(*) FROM compendium_campaigns`); n != 1 {
		t.Errorf("%d homebrew links, want 1", n)
	}

	// The homebrew leaves the file: none left over.
	e.importStarterFile(t, func(th *export.TakeHome) { th.Entries = th.Entries[:1] })
	if n := e.count(
		t,
		`SELECT count(*) FROM compendium_entries WHERE source_key = 'homebrew'`,
	); n != 0 {
		t.Errorf("%d orphaned homebrew entries", n)
	}
}

func TestStarterOfAnotherTableRefused(t *testing.T) {
	e := newPlayerEnv(t)
	e.importStarterFile(t, nil)
	w := e.importRequest(
		t,
		starterZip(t, func(th *export.TakeHome) { th.Manifest.ServerID = "elsewhere" }),
		"",
		nil,
	)
	if w.Code != http.StatusUnprocessableEntity ||
		!strings.Contains(w.Body.String(), "another GM") {
		t.Errorf("another table's starter: %d %s", w.Code, w.Body)
	}
	if n := e.count(t, `SELECT count(*) FROM home_starters`); n != 1 {
		t.Errorf("%d starters, want the first only", n)
	}
	// A take-home of this table is refused next to another table's starter.
	e2 := newPlayerEnv(t)
	e2.importStarterFile(t, func(th *export.TakeHome) { th.Manifest.ServerID = "elsewhere" })
	if w := e2.importRequest(
		t,
		sampleTakeHome(t, nil),
		"",
		nil,
	); w.Code != http.StatusUnprocessableEntity ||
		!strings.Contains(w.Body.String(), "another GM") {
		t.Errorf("take-home next to another table's starter: %d %s", w.Code, w.Body)
	}
}

func TestHomeNewFromStarter(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	q := e.q()
	e.importStarterFile(t, nil)
	s := e.starterCampaign(t)
	form := func(from int64) url.Values {
		return url.Values{"name": {"Nia"}, "from": {strconv.FormatInt(from, 10)}}
	}
	w := e.postPlayer(t, route.AtHomeNew, form(s.CampaignID))
	nia := e.homeCharacterByName(t, "Nia")
	if w.Code != http.StatusNoContent ||
		w.Header().Get("HX-Redirect") != wizardPath(nia.ID, wizardSteps[0].Key) {
		t.Fatalf("new from a starter: %d %q: %s", w.Code, w.Header().Get("HX-Redirect"), w.Body)
	}
	if nia.CampaignID == s.CampaignID {
		t.Error("the character shares the starter's campaign")
	}
	if got, _ := q.ListCampaignSources(ctx, nia.CampaignID); len(got) != 2 {
		t.Errorf("sources %v", got)
	}
	if n := e.count(
		t,
		`SELECT count(*) FROM compendium_campaigns WHERE campaign_id = ?`,
		nia.CampaignID,
	); n != 1 {
		t.Errorf("%d homebrew links, want the starter's one", n)
	}
	// A character's own campaign is neither a starter nor a table copy's.
	if w := e.postPlayer(
		t,
		route.AtHomeNew,
		form(nia.CampaignID),
	); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("model that is no starter or table copy: %d", w.Code)
	}

	// The new character's file has the table's id and a token that stays the same.
	f1, _ := e.homeFileOf(t, nia.ID, export.KindNewCharacter)
	f2, _ := e.homeFileOf(t, nia.ID, export.KindNewCharacter)
	if f1 == nil || f2 == nil {
		t.Fatal("no new-character file")
	}
	if f1.Manifest.ServerID != "srv" || f1.Manifest.MadeAtHome == "" ||
		f1.Manifest.MadeAtHome != f2.Manifest.MadeAtHome {
		t.Errorf(
			"tokens %q and %q, server %q",
			f1.Manifest.MadeAtHome,
			f2.Manifest.MadeAtHome,
			f1.Manifest.ServerID,
		)
	}
}

func TestStarterRemoveKeepsCharacters(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	e.importStarterFile(t, nil)
	s := e.starterCampaign(t)
	e.postPlayer(
		t,
		route.AtHomeNew,
		url.Values{"name": {"Nia"}, "from": {strconv.FormatInt(s.CampaignID, 10)}},
	)
	nia := e.homeCharacterByName(t, "Nia")

	for _, id := range []int64{999, nia.CampaignID} {
		if w := e.postPlayer(
			t,
			route.Path(route.AtHomeStarterRemove, id),
			nil,
		); w.Code != http.StatusUnprocessableEntity {
			t.Errorf("remove %d, which is no starter: %d", id, w.Code)
		}
	}
	if w := e.postPlayer(
		t,
		route.Path(route.AtHomeStarterRemove, s.CampaignID),
		nil,
	); w.Code != http.StatusNoContent {
		t.Fatalf("remove: %d %s", w.Code, w.Body)
	}
	if rows, _ := e.q().ListHomeStarters(ctx); len(rows) != 0 {
		t.Errorf("starters %+v", rows)
	}
	if _, err := e.q().GetCampaign(ctx, s.CampaignID); err == nil {
		t.Error("the starter's campaign is still here")
	}
	again := e.homeCharacterByName(t, "Nia")
	if got, _ := e.q().ListCampaignSources(ctx, again.CampaignID); len(got) != 2 {
		t.Errorf("Nia's sources after the removal: %v", got)
	}
	if n := e.count(
		t,
		`SELECT count(*) FROM compendium_campaigns WHERE campaign_id = ?`,
		again.CampaignID,
	); n != 1 {
		t.Errorf("Nia lost the homebrew link: %d", n)
	}
	if n := e.count(
		t,
		`SELECT count(*) FROM compendium_entries WHERE source_key = 'homebrew'`,
	); n != 1 {
		t.Errorf("%d homebrew rows after the removal, want the one Nia still links", n)
	}
}

// A take-home carrying the token of a character made here asks first, then
// replaces exactly that character with its campaign.
func TestTakeHomeReplacesMadeAtHome(t *testing.T) {
	e := newPlayerEnv(t)
	e.importStarterFile(t, nil)
	s := e.starterCampaign(t)
	tokenOf := func(name string) string {
		e.postPlayer(
			t,
			route.AtHomeNew,
			url.Values{"name": {name}, "from": {strconv.FormatInt(s.CampaignID, 10)}},
		)
		f, _ := e.homeFileOf(t, e.homeCharacterByName(t, name).ID, export.KindNewCharacter)
		return f.Manifest.MadeAtHome
	}
	niaToken := tokenOf("Nia")
	tokenOf("Zed")
	nia, zed := e.homeCharacterByName(t, "Nia"), e.homeCharacterByName(t, "Zed")
	// tableID picks which table character the file is, so a copy here matches only when it is 7.
	takeHome := func(token string, tableID int64) []byte {
		return sampleTakeHome(t, func(th *export.TakeHome) {
			th.Manifest.MadeAtHome, th.Manifest.CharacterID = token, tableID
		})
	}
	names := func() []string {
		rows, err := e.q().ListHomeCharacters(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		var out []string
		for _, r := range rows {
			out = append(out, r.Name)
		}
		slices.Sort(out)
		return out
	}

	// No match: opens at once and keeps both made here.
	if w := e.importRequest(t, takeHome("nobody", 7), "", nil); w.Code != http.StatusNoContent {
		t.Fatalf("no match: %d %s", w.Code, w.Body)
	}
	if got := names(); !slices.Equal(got, []string{"Aria", "Nia", "Zed"}) {
		t.Fatalf("characters %v", got)
	}

	// A match asks first and changes nothing.
	w := e.importRequest(t, takeHome(niaToken, 8), "", nil)
	if w.Code != http.StatusOK ||
		!strings.Contains(
			w.Body.String(),
			"Replaces Nia, made at home. Changes made at home since you sent them are lost.",
		) ||
		!strings.Contains(w.Body.String(), `name="replace"`) {
		t.Fatalf("confirm: %d %s", w.Code, w.Body)
	}
	if got := names(); !slices.Equal(got, []string{"Aria", "Nia", "Zed"}) {
		t.Fatalf("the confirm changed things: %v", got)
	}
	// Confirmed: exactly Nia and her campaign go.
	if w := e.importRequest(
		t,
		takeHome(niaToken, 8),
		"",
		map[string]string{formReplace: "1"},
	); w.Code != http.StatusNoContent {
		t.Fatalf("replace: %d %s", w.Code, w.Body)
	}
	if got := names(); !slices.Equal(got, []string{"Aria", "Aria", "Zed"}) {
		t.Errorf(
			"characters after the replace: %v, want Nia gone and the two Arias and Zed kept",
			got,
		)
	}
	if _, err := e.q().GetCampaign(t.Context(), nia.CampaignID); err == nil {
		t.Error("Nia's campaign survived")
	}
	for name, id := range map[string]int64{"Zed's": zed.CampaignID, "the starter's": s.CampaignID} {
		if _, err := e.q().GetCampaign(t.Context(), id); err != nil {
			t.Errorf("%s campaign: %v", name, err)
		}
	}
	// The copy that replaced Nia goes back to the table without the token (the table refuses one).
	cp, err := e.q().
		GetHomeCopyByTable(t.Context(), sqlc.GetHomeCopyByTableParams{ServerID: "srv", TableCharacterID: 8})
	if err != nil {
		t.Fatal(err)
	}
	if f, w := e.homeFileOf(t, cp.CharacterID, export.KindBringBack); f == nil {
		t.Errorf("bring-back: %d %s", w.Code, w.Body)
	}

	// A file that matches a copy here and a made-at-home character asks once, about both.
	zedToken := func() string {
		f, _ := e.homeFileOf(t, zed.ID, export.KindNewCharacter)
		return f.Manifest.MadeAtHome
	}()
	w = e.importRequest(t, takeHome(zedToken, 7), "", nil)
	body := w.Body.String()
	if w.Code != http.StatusOK || !strings.Contains(body, "already on this computer") ||
		!strings.Contains(body, "Replaces Zed, made at home") {
		t.Errorf("combined confirm: %d %s", w.Code, body)
	}
}
