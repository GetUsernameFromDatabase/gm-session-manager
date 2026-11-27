package web

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"html"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

func (e testEnv) source(t testing.TB, key, edition string) {
	t.Helper()
	if err := e.q().UpsertCompendiumSource(t.Context(), sqlc.UpsertCompendiumSourceParams{SourceKey: key, Edition: edition, License: "test", Attribution: key + " attribution"}); err != nil {
		t.Fatal(err)
	}
}

// upsertEntry stores p, first adding a bare row for its source when the test
// has not made one: entries reference their source.
func (e testEnv) upsertEntry(ctx context.Context, p sqlc.UpsertCompendiumEntryParams) error {
	if _, err := e.q().GetCompendiumSource(ctx, p.SourceKey); errors.Is(err, sql.ErrNoRows) {
		if err := e.q().UpsertCompendiumSource(ctx, sqlc.UpsertCompendiumSourceParams{SourceKey: p.SourceKey}); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return e.q().UpsertCompendiumEntry(ctx, p)
}

func (e testEnv) feat(t *testing.T, source, slug, desc string) int64 {
	t.Helper()
	data := fmt.Sprintf(`{"desc":%q}`, desc)
	if err := e.upsertEntry(t.Context(), sqlc.UpsertCompendiumEntryParams{Kind: "feat", Slug: slug, Name: slug, SourceKey: source, DataJson: data, BodyText: desc, ContentHash: source + slug}); err != nil {
		t.Fatal(err)
	}
	row, err := e.q().GetCompendiumEntryBySlug(t.Context(), sqlc.GetCompendiumEntryBySlugParams{Kind: "feat", Slug: slug})
	if err != nil {
		t.Fatal(err)
	}
	if row.SourceKey != source { // the lower-priority copy
		rows, _ := e.q().ListEntryCopies(t.Context(), sqlc.ListEntryCopiesParams{Kind: "feat", Slug: slug})
		for _, r := range rows {
			if r.SourceKey == source {
				return r.ID
			}
		}
	}
	return row.ID
}

func TestFiveToolsGMOnly(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, enum.RolePlayer)
	gm := e.token(t, enum.RoleAdmin)
	e.source(t, "5etools-phb", "2014")
	tests := []struct{ method, path string }{
		{http.MethodGet, views.FiveToolsPath},
		{http.MethodPost, views.FiveToolsPath},
		{http.MethodPost, route.Compendium + route.Path(route.CompendiumFiveToolsBookDelete, "phb")},
		{http.MethodPost, route.Compendium + route.CompendiumFiveToolsAuto},
		{http.MethodPost, route.Compendium + route.CompendiumImportDismiss},
		{http.MethodPost, route.Compendium + route.CompendiumOpen5e},
		{http.MethodPost, route.Compendium + route.CompendiumSourceDelete},
		{http.MethodPost, views.SourceOrderPath},
	}
	for _, tt := range tests {
		// The same under the whole compendium and the player's campaign's.
		for _, base := range []string{"", route.Path(route.Campaign, e.campaignID)} {
			t.Run(tt.method+" "+base+tt.path, func(t *testing.T) {
				if w := e.do(t, tt.method, base+tt.path, player, url.Values{"dir": {"up"}}, true); w.Code != http.StatusForbidden {
					t.Errorf("player: %d", w.Code)
				}
			})
		}
	}
	// Reloading the picker's URL (where an import posts) opens the whole page.
	if w := e.do(t, http.MethodGet, views.FiveToolsPath, gm, nil, false); w.Code != http.StatusSeeOther || w.Header().Get("Location") != route.Compendium+"?import=fivetools" {
		t.Errorf("plain visit: %d %q", w.Code, w.Header().Get("Location"))
	}
	// An empty pick is refused before anything is downloaded.
	if w := e.do(t, http.MethodPost, views.FiveToolsPath, gm, url.Values{}, false); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Pick at least one book") {
		t.Errorf("empty pick: %d", w.Code)
	}
}

// TestSourceOrder: a drag posts every key in the new order; a list that
// misses a source (a stale page) changes nothing.
func TestSourceOrder(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, enum.RoleAdmin)
	e.source(t, "book-a", "2024")
	e.source(t, "book-b", "2014")
	order := func() string {
		rows, err := e.q().ListCompendiumSources(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		var keys []string
		for _, r := range rows {
			keys = append(keys, r.SourceKey)
		}
		return strings.Join(keys, ",")
	}
	for _, tt := range []struct {
		name string
		keys []string
		code int
		want string
	}{
		{name: "dragged", keys: []string{"book-b", "homebrew", "book-a"}, code: http.StatusOK, want: "book-b,homebrew,book-a"},
		// Refusals answer htmx with a flash message (a.fail).
		{name: "stale page", keys: []string{"book-a", "book-b"}, code: http.StatusUnprocessableEntity, want: "book-b,homebrew,book-a"},
		{name: "nothing sent", code: http.StatusUnprocessableEntity, want: "book-b,homebrew,book-a"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			w := e.do(t, http.MethodPost, views.SourceOrderPath, gm, url.Values{views.SourceKeyField: tt.keys}, true)
			if w.Code != tt.code {
				t.Fatalf("code %d, want %d: %s", w.Code, tt.code, w.Body)
			}
			if got := order(); got != tt.want {
				t.Errorf("order %s, want %s", got, tt.want)
			}
		})
	}
}

// TestFiveToolsImportRunning: the import tab shows a running import's
// progress, and a second import or a removal waits for it to finish.
func TestFiveToolsImportRunning(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, enum.RoleAdmin)
	e.source(t, "5etools-phb", "2014")
	e.app.fiveTools.start(4)
	e.app.fiveTools.update(func(j *views.FiveToolsJob) { j.Done, j.Current = 1, "Test Book" })
	body := e.do(t, http.MethodGet, route.Compendium+"?import="+views.ImportTabFiveTools, gm, nil, false).Body.String()
	if !strings.Contains(body, "Downloading 2 of 4: Test Book") || !regexp.MustCompile(`style="width:\s?25%;?"`).MatchString(body) {
		t.Error("import tab misses the progress label or bar")
	}
	if w := e.do(t, http.MethodPost, views.FiveToolsPath, gm, url.Values{views.FiveToolsField: {"phb"}}, false); !strings.Contains(w.Body.String(), errImportRunning) {
		t.Errorf("second import: %d, not refused", w.Code)
	}
	if w := e.do(t, http.MethodPost, views.FiveToolsPath+"/phb/delete", gm, nil, true); !strings.Contains(w.Body.String(), errImportRunning) {
		t.Errorf("removal during an import: %d, not refused", w.Code)
	}
	dismiss := route.Compendium + route.CompendiumImportDismiss
	if w := e.do(t, http.MethodPost, dismiss, gm, nil, true); !strings.Contains(w.Body.String(), errImportRunning) || !e.app.fiveTools.get().Running {
		t.Errorf("dismiss during an import: %d, not refused or the job dropped", w.Code)
	}
	// Done, the tab shows the summary instead.
	e.app.fiveTools.update(func(j *views.FiveToolsJob) {
		*j = views.FiveToolsJob{Msg: "5etools-phb: 3 entries imported, 0 removed"}
	})
	if body := e.do(t, http.MethodGet, route.Compendium+"?import="+views.ImportTabFiveTools, gm, nil, false).Body.String(); !strings.Contains(body, "3 entries imported") || strings.Contains(body, "Downloading") || !strings.Contains(body, `hx-post="`+dismiss+`"`) {
		t.Error("finished import shows no summary with its dismiss button, or still its progress")
	}
	// A player may not dismiss it; the GM's X clears it.
	if w := e.do(t, http.MethodPost, dismiss, e.token(t, enum.RolePlayer), nil, true); w.Code != http.StatusForbidden || e.app.fiveTools.get().Msg == "" {
		t.Errorf("player dismiss: %d", w.Code)
	}
	// A stale tab's X names an older summary and clears nothing.
	if w := e.do(t, http.MethodPost, dismiss, gm, url.Values{views.DismissWasField: {"older summary"}}, true); !strings.Contains(w.Body.String(), "changed meanwhile") || e.app.fiveTools.get().Msg == "" {
		t.Errorf("stale dismiss: %d, job %+v", w.Code, e.app.fiveTools.get())
	}
	if w := e.do(t, http.MethodPost, dismiss, gm, url.Values{views.DismissWasField: {"5etools-phb: 3 entries imported, 0 removed"}}, true); w.Code != http.StatusNoContent || e.app.fiveTools.get().Msg != "" {
		t.Errorf("GM dismiss: %d, job %+v", w.Code, e.app.fiveTools.get())
	}
	if body := e.do(t, http.MethodGet, route.Compendium+"?import="+views.ImportTabFiveTools, gm, nil, false).Body.String(); strings.Contains(body, "3 entries imported") {
		t.Error("dismissed summary still shows")
	}
}

// TestStarterSets: the GM's empty compendium offers the starting sets and
// shows the job in their place; a player's does not. Nothing here reaches
// the network: every post is refused before a download.
func TestStarterSets(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)
	page := func(cookie string) string {
		return e.do(t, http.MethodGet, route.Compendium, cookie, nil, false).Body.String()
	}
	live := `data-live="` + views.LiveKey(views.LiveImport, views.StarterLiveCompendium) + `"`
	if body := page(gm); !strings.Contains(body, "Import SRD") || !strings.Contains(body, "Import books") || !strings.Contains(body, live) {
		t.Error("GM's empty compendium has no starting sets")
	}
	if body := page(player); strings.Contains(body, "Import SRD") || strings.Contains(body, live) {
		t.Error("player's empty compendium shows the starting sets")
	}
	// A bad set is refused before the book list is fetched (a.fail answers htmx with a 422).
	if w := e.do(t, http.MethodGet, views.FiveToolsPath+"?"+views.FiveToolsSetField+"=2014", gm, nil, true); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("picker with a bad set: %d", w.Code)
	}
	if w := e.do(t, http.MethodPost, views.FiveToolsPath, gm, url.Values{views.FiveToolsSetField: {"2014"}}, false); !strings.Contains(w.Body.String(), "Pick at least one book") {
		t.Errorf("bad set posted: %d, not refused", w.Code)
	}
	// Running: the progress stands in for the rows, and both sets wait.
	e.app.fiveTools.start(0)
	if body := page(gm); !strings.Contains(body, "Starting…") || strings.Contains(body, "Import SRD") {
		t.Error("running job not shown in the empty compendium")
	}
	if w := e.do(t, http.MethodPost, route.Compendium+route.CompendiumOpen5e, gm, url.Values{views.Open5eDocField: {compendium.SRD2024Key}}, false); !strings.Contains(w.Body.String(), errImportRunning) {
		t.Errorf("SRD during an import: %d, not refused", w.Code)
	}
	if w := e.do(t, http.MethodPost, views.FiveToolsPath, gm, url.Values{views.FiveToolsSetField: {views.FiveToolsSet2024}}, false); !strings.Contains(w.Body.String(), errImportRunning) {
		t.Errorf("2024 set during an import: %d, not refused", w.Code)
	}
	if w := e.do(t, http.MethodPost, route.Compendium+route.CompendiumOpen5e, player, nil, false); w.Code != http.StatusForbidden {
		t.Errorf("player SRD import: %d", w.Code)
	}
	// Done: the summary and a way to the entries above the rows; failed: the error.
	e.app.fiveTools.update(func(j *views.FiveToolsJob) { *j = views.FiveToolsJob{Msg: "srd-2024: 3 entries imported"} })
	if body := page(gm); !strings.Contains(body, "Show entries") || !strings.Contains(body, "Import SRD") {
		t.Error("finished job shows no way to the entries, or hides the sets")
	}
	e.app.fiveTools.update(func(j *views.FiveToolsJob) { *j = views.FiveToolsJob{Msg: "Download failed: boom", Failed: true} })
	if body := page(gm); !strings.Contains(body, "Download failed: boom") || !strings.Contains(body, "Import SRD") {
		t.Error("failed job hides its error or the rows")
	}
}

// TestTickSet: ?set=2024 ticks the set's books, and only those, beside the
// downloaded ones.
func TestTickSet(t *testing.T) {
	wotc := "Wizards of the Coast"
	books := map[string]compendium.Book{
		"XPHB": {ID: "XPHB", Name: "Player's Handbook", Group: compendium.GroupCore, Publisher: wotc, Published: "2024-09-17"},
		"PHB":  {ID: "PHB", Name: "Old Handbook", Group: compendium.GroupCore, Publisher: wotc, Published: "2014-08-19"},
		"LMOP": {ID: "LMOP", Name: "Lost Mine", Group: compendium.GroupAdventure, Publisher: wotc, Published: "2025-01-01"},
	}
	want := map[string]bool{"XPHB": true}
	var list []compendium.Book
	for _, b := range books {
		list = append(list, b)
	}
	v := views.NewFiveToolsView(list, nil)
	tickSet(&v)
	for _, g := range v.Groups {
		for _, b := range g.Books {
			if b.Ticked != want[b.ID] {
				t.Errorf("%s ticked %v", b.ID, b.Ticked)
			}
		}
	}
}

// TestRunFiveToolsFinishes: a run always ends the job with its summary, so
// the next import is not refused.
func TestRunFiveToolsFinishes(t *testing.T) {
	e := newTestEnv(t)
	if !e.app.fiveTools.start(0) {
		t.Fatal("job busy before the run")
	}
	e.app.runFiveTools(t.Context(), t.TempDir(), nil, 0) // no books: nothing to download
	if j := e.app.fiveTools.get(); j.Running || !j.Failed || !strings.Contains(j.Msg, "No source folders") {
		t.Errorf("job after the run: %+v", j)
	}
	if !e.app.fiveTools.start(1) {
		t.Error("the next import is refused")
	}
}

// TestImportSummary: several sources get a total line first, which leaves
// out failed and unchanged ones; a single source does not.
func TestImportSummary(t *testing.T) {
	for _, tt := range []struct {
		name    string
		reports []compendium.Report
		want    string
		failed  bool
	}{
		{name: "one", reports: []compendium.Report{{Source: "a", Entries: 3, Removed: 1}}, want: "a: 3 entries imported, 1 removed"},
		{
			name:    "several",
			reports: []compendium.Report{{Source: "a", Entries: 3, Removed: 1}, {Source: "b", Entries: 5}, {Source: "c", Unchanged: true}, {Source: "d", Err: errors.New("bad")}},
			want:    "4 sources: 8 entries imported, 1 removed in all\na: 3 entries imported, 1 removed\nb: 5 entries imported, 0 removed\nc: unchanged\nd: failed: bad",
			failed:  true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got, failed := ImportSummary(tt.reports); got != tt.want || failed != tt.failed {
				t.Errorf("got %q %v, want %q %v", got, failed, tt.want, tt.failed)
			}
		})
	}
}

func TestFiveToolsRemove(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, enum.RoleAdmin)
	e.source(t, "5etools-phb", "2014")
	e.source(t, "5etools-xge", "2014")
	e.feat(t, "5etools-phb", "alert", "x")
	e.feat(t, "5etools-xge", "tough", "y")
	folder := filepath.Join(e.app.compendiumDir(), "5etools-phb")
	if err := os.MkdirAll(folder, 0o750); err != nil {
		t.Fatal(err)
	}
	w := e.do(t, http.MethodPost, route.Compendium+route.Path(route.CompendiumFiveToolsBookDelete, "phb"), gm, nil, true)
	if w.Code != http.StatusOK {
		t.Fatalf("remove: %d %s", w.Code, w.Body)
	}
	// A toast: the picker's notice would vanish with the reload the publish causes.
	if body := w.Body.String(); !strings.Contains(body, "Removed phb.") || !strings.Contains(body, `id="`+views.FiveToolsID+`"`) {
		t.Errorf("remove answer lacks the toast or the picker: %s", body)
	}
	if _, err := os.Stat(folder); !os.IsNotExist(err) {
		t.Errorf("folder kept: %v", err)
	}
	if _, err := e.q().GetCompendiumSource(t.Context(), "5etools-phb"); err == nil {
		t.Error("source kept")
	}
	if _, err := e.q().GetCompendiumEntryBySlug(t.Context(), sqlc.GetCompendiumEntryBySlugParams{Kind: "feat", Slug: "alert"}); err == nil {
		t.Error("entry kept")
	}
	if _, err := e.q().GetCompendiumEntryBySlug(t.Context(), sqlc.GetCompendiumEntryBySlugParams{Kind: "feat", Slug: "tough"}); err != nil {
		t.Errorf("another book's entry went too: %v", err)
	}
	// A path-like id is refused, never joined onto the folder.
	if w := e.do(t, http.MethodPost, route.Compendium+route.CompendiumFiveTools+"/..%2Fx/delete", gm, nil, true); strings.Contains(w.Body.String(), "Removed") {
		t.Errorf("bad id accepted: %d", w.Code)
	}
}

// TestEntryAlsoInAndCrossRefs: two sources hold one feat; each entry page
// names the other, and a [[compendium:...]] token links to the winning copy.
func TestEntryAlsoInAndCrossRefs(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)
	e.source(t, "book-new", "2024") // 100 wins
	e.source(t, "book-old", "2014") // 300
	text := "See [[compendium:feat/alert|Alert]], [[compendium:feat/ghost|Ghost]], [[compendium:feat/secret|Secret]] and [[compendium:feat/shared|Shared]]."
	newer := e.feat(t, "book-new", "alert", "newer")
	older := e.feat(t, "book-old", "alert", "older")
	pointer := e.feat(t, "book-new", "pointer", text)
	// Player homebrew: secret is linked to no campaign, shared to the player's.
	scoped := map[string]int64{}
	for _, slug := range []string{"secret", "shared"} {
		row, err := e.q().CreateHomebrew(t.Context(), sqlc.CreateHomebrewParams{Kind: "feat", Slug: slug, Name: slug, DataJson: `{"desc":""}`, ContentHash: slug, OwnerCharacterID: nullID(e.characterID), Scoped: true})
		if err != nil {
			t.Fatal(err)
		}
		scoped[slug] = row.ID
	}
	if err := e.q().LinkEntryCampaign(t.Context(), sqlc.LinkEntryCampaignParams{EntryID: scoped["shared"], CampaignID: e.campaignID}); err != nil {
		t.Fatal(err)
	}
	// The GM reads the whole compendium, the player their campaign's.
	mine := route.Path(route.CampaignCompendium, e.campaignID)
	link := func(base, slug, label string) string {
		return fmt.Sprintf(`<a href="%s/entries/%d">%s</a>`, base, scoped[slug], label)
	}

	tests := []struct {
		name, cookie, path string
		want, notWant      []string
	}{
		{"older names the newer", gm, route.Compendium + route.Path(route.CompendiumEntry, older), []string{"Also in", fmt.Sprintf(`href="%s%s"`, route.Compendium, route.Path(route.CompendiumEntry, newer)), "book-new, 2024"}, nil},
		{"newer names the older", player, fmt.Sprintf("%s/entries/%d", mine, newer), []string{"Also in", fmt.Sprintf(`href="%s/entries/%d"`, mine, older)}, nil},
		{"token links the winner, missing is plain", player, fmt.Sprintf("%s/entries/%d", mine, pointer),
			[]string{fmt.Sprintf(`<a href="%s/entries/%d">Alert</a>`, mine, newer), "Ghost", "Secret", link(mine, "shared", "Shared")}, []string{"Also in", "[[compendium", route.Path(route.CompendiumEntry, older), link(mine, "secret", "Secret")}},
		{"the GM sees every homebrew link", gm, route.Compendium + route.Path(route.CompendiumEntry, pointer), []string{link(route.Compendium, "secret", "Secret"), link(route.Compendium, "shared", "Shared")}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := e.do(t, http.MethodGet, tt.path, tt.cookie, nil, false).Body.String()
			for _, s := range tt.want {
				if !strings.Contains(body, s) {
					t.Errorf("missing %q", s)
				}
			}
			for _, s := range tt.notWant {
				if strings.Contains(body, s) {
					t.Errorf("unexpected %q", s)
				}
			}
			if strings.Contains(body, `href="`+route.Compendium+route.Path(route.CompendiumEntry, 0)+`"`) || strings.Contains(body, ">Ghost</a>") {
				t.Error("link to a missing entry")
			}
		})
	}
}

// TestSourceDeleteUsed: a source a sheet pins is refused and kept; once the
// pin is gone it is deleted with its folder.
func TestSourceDeleteUsed(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, enum.RoleAdmin)
	e.source(t, "my-folder", "2024")
	featID := e.feat(t, "my-folder", "alert", "x")
	folder := filepath.Join(e.app.compendiumDir(), "my-folder")
	if err := os.MkdirAll(folder, 0o750); err != nil {
		t.Fatal(err)
	}
	pin := func(sheet string) {
		if err := e.q().UpdateCharacterSheet(t.Context(), sqlc.UpdateCharacterSheetParams{SheetJson: sheet, ID: e.characterID}); err != nil {
			t.Fatal(err)
		}
	}
	pin(`{"species_source":"my-folder"}`)
	del := url.Values{"key": {"my-folder"}}
	w := e.do(t, http.MethodPost, route.Compendium+route.CompendiumSourceDelete, gm, del, true)
	if !strings.Contains(w.Body.String(), "my-folder is used by") || !strings.Contains(w.Body.String(), "Switch their sources first.") {
		t.Errorf("no refusal naming users: %d %s", w.Code, w.Body)
	}
	if got := w.Header().Get("HX-Reswap"); got != "none" {
		t.Errorf("HX-Reswap = %q, want none (the card must not be swapped for the flash)", got)
	}
	if _, err := e.q().GetCompendiumSource(t.Context(), "my-folder"); err != nil {
		t.Errorf("used source deleted: %v", err)
	}
	pin(`{}`)
	// An entity linked to one of its entries keeps it too.
	ent, err := e.q().CreateEntity(t.Context(), sqlc.CreateEntityParams{CampaignID: e.campaignID, Kind: "npc", Name: "Goblin boss", FieldsJson: "{}", CompendiumRef: sql.NullInt64{Int64: featID, Valid: true}})
	if err != nil {
		t.Fatal(err)
	}
	if w := e.do(t, http.MethodPost, route.Compendium+route.CompendiumSourceDelete, gm, del, true); !strings.Contains(w.Body.String(), "my-folder is used by Goblin boss.") {
		t.Errorf("entity link not counted: %d %s", w.Code, w.Body)
	}
	if err := e.q().DeleteEntity(t.Context(), ent.ID); err != nil {
		t.Fatal(err)
	}
	if w := e.do(t, http.MethodPost, route.Compendium+route.CompendiumSourceDelete, gm, del, true); w.Code != http.StatusOK {
		t.Fatalf("delete: %d %s", w.Code, w.Body)
	}
	if _, err := e.q().GetCompendiumSource(t.Context(), "my-folder"); err == nil {
		t.Error("source kept")
	}
	if _, err := os.Stat(folder); !os.IsNotExist(err) {
		t.Errorf("folder kept: %v", err)
	}
}

// TestSourceDeleteMany: several keys go together or not at all.
func TestSourceDeleteMany(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, enum.RoleAdmin)
	for _, k := range []string{"src-a", "src-b", "src-used"} {
		e.source(t, k, "2024")
	}
	has := func(k string) bool {
		_, err := e.q().GetCompendiumSource(t.Context(), k)
		return err == nil
	}
	post := func(keys ...string) *httptest.ResponseRecorder {
		return e.do(t, http.MethodPost, route.Compendium+route.CompendiumSourceDelete, gm, url.Values{"key": keys}, true)
	}
	if err := e.q().UpdateCharacterSheet(t.Context(), sqlc.UpdateCharacterSheetParams{SheetJson: `{"species_source":"src-used"}`, ID: e.characterID}); err != nil {
		t.Fatal(err)
	}
	if w := post("src-a", "src-used", "src-b"); w.Code == http.StatusOK || !strings.Contains(w.Body.String(), "src-used is used by") {
		t.Errorf("used key: %d %s", w.Code, w.Body)
	}
	if !has("src-a") || !has("src-b") || !has("src-used") {
		t.Error("a refused request deleted something")
	}
	if w := post("src-a", compendium.HomebrewSource); w.Code == http.StatusOK || !strings.Contains(w.Body.String(), "Homebrew cannot be deleted") {
		t.Errorf("homebrew: %d %s", w.Code, w.Body)
	}
	if !has("src-a") {
		t.Error("homebrew request deleted another key")
	}
	if w := post("src-a", "src-b", "src-a"); w.Code != http.StatusOK {
		t.Fatalf("delete: %d %s", w.Code, w.Body)
	}
	if has("src-a") || has("src-b") || !has("src-used") {
		t.Error("wrong sources left")
	}
}

// TestSourceKeyWithSlash: page imports key sources like 5etools:mtf/howler
// or aidedd:dnd/sorts.php?vf=a+b, whose characters a URL path would mangle;
// move and delete take the key from the body, which the import tab's
// buttons must carry. The longest key a page link allows must pass too.
func TestSourceKeyWithSlash(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, enum.RoleAdmin)
	e.source(t, "book-a", "2024")
	longLink := "https://www.aidedd.org/dnd/x.php?q="
	longLink += strings.Repeat("a", 500-len(longLink)) // pageForm.Page's max
	site, page, err := compendium.ParsePageURL(longLink)
	if err != nil {
		t.Fatal(err)
	}
	longKey := compendium.PageSourceKey(site.Key, page)
	if len(longKey) <= 200 {
		t.Fatalf("long key is only %d characters", len(longKey))
	}
	order := func() string {
		rows, err := e.q().ListCompendiumSources(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		var keys []string
		for _, r := range rows {
			keys = append(keys, r.SourceKey)
		}
		return strings.Join(keys, ",")
	}
	for _, key := range []string{"5etools:mtf/howler", "aidedd:dnd/sorts.php?vf=a+b", longKey} {
		t.Run(key[:min(len(key), 40)], func(t *testing.T) {
			e.source(t, key, "2024")
			body := e.do(t, http.MethodGet, route.Compendium+"?tab=import", gm, nil, false).Body.String()
			// toJSON's sorted keys; none of these keys holds & < > (json escapes them)
			if want := html.EscapeString(`{"key":"` + key + `"}`); !strings.Contains(body, want) {
				t.Errorf("import tab lacks %s", want)
			}
			if want := `data-source-key="` + html.EscapeString(key) + `"`; !strings.Contains(body, want) {
				t.Errorf("import tab lacks %s", want)
			}
			// Typed or dragged, a move posts the whole order with the key first.
			keys := slices.DeleteFunc(strings.Split(order(), ","), func(k string) bool { return k == key })
			keys = append([]string{key}, keys...)
			if w := e.do(t, http.MethodPost, views.SourceOrderPath, gm, url.Values{views.SourceKeyField: keys}, true); w.Code != http.StatusOK {
				t.Fatalf("order: %d %s", w.Code, w.Body)
			}
			if got := order(); got != strings.Join(keys, ",") {
				t.Errorf("order %s, want %s", got, strings.Join(keys, ","))
			}
			if w := e.do(t, http.MethodPost, route.Compendium+route.CompendiumSourceDelete, gm, url.Values{"key": {key}}, true); w.Code != http.StatusOK {
				t.Fatalf("delete: %d %s", w.Code, w.Body)
			}
			if _, err := e.q().GetCompendiumSource(t.Context(), key); err == nil {
				t.Error("source kept")
			}
		})
	}
}

func TestFiveToolsBack(t *testing.T) {
	global := route.Compendium
	camp3 := route.Path(route.CampaignCompendium, 3)
	tests := []struct{ name, referer, want string }{
		{"global", "https://host" + global + "?q=x", global},
		{"campaign", "https://host" + camp3, camp3},
		{"unrelated page", "https://host" + route.Path(route.Campaign, 3) + route.Images, global},
		{"campaign page", "https://host" + route.Path(route.Campaign, 3), route.Path(route.Campaign, 3)},
		{"other host", "https://elsewhere.example" + camp3, camp3},
		{"entry page", "https://host" + camp3 + route.Path(route.CompendiumEntry, 7), global},
		{"no referer", "", global},
		{"malformed", "http://[::1", global},
		{"bad id", "https://host" + route.Path(route.CampaignCompendium, "x"), global},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request = httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", nil)
			c.Request.Header.Set("Referer", tt.referer)
			if got := fiveToolsBack(c); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
