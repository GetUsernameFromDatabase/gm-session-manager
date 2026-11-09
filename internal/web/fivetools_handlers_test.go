package web

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/web/views"
)

func (e testEnv) source(t *testing.T, key, edition string) {
	t.Helper()
	if err := e.q().UpsertCompendiumSource(t.Context(), sqlc.UpsertCompendiumSourceParams{SourceKey: key, Edition: edition, License: "test", Attribution: key + " attribution"}); err != nil {
		t.Fatal(err)
	}
}

func (e testEnv) feat(t *testing.T, source, slug, desc string) int64 {
	t.Helper()
	data := fmt.Sprintf(`{"desc":%q}`, desc)
	if err := e.q().UpsertCompendiumEntry(t.Context(), sqlc.UpsertCompendiumEntryParams{Kind: "feat", Slug: slug, Name: slug, SourceKey: source, DataJson: data, BodyText: desc, ContentHash: source + slug}); err != nil {
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
	player := e.token(t, auth.RolePlayer)
	gm := e.token(t, auth.RoleAdmin)
	e.source(t, "5etools-phb", "2014")
	tests := []struct{ method, path string }{
		{http.MethodGet, "/compendium/5etools"},
		{http.MethodPost, "/compendium/5etools"},
		{http.MethodPost, "/compendium/5etools/phb/delete"},
		{http.MethodPost, "/compendium/sources/5etools-phb/move"},
		{http.MethodPost, views.SourceOrderPath},
	}
	for _, tt := range tests {
		t.Run(tt.method+" "+tt.path, func(t *testing.T) {
			if w := e.do(t, tt.method, tt.path, player, url.Values{"dir": {"up"}}, false); w.Code != http.StatusForbidden {
				t.Errorf("player: %d", w.Code)
			}
		})
	}
	// Reloading the picker's URL (where an import posts) opens the whole page.
	if w := e.do(t, http.MethodGet, "/compendium/5etools", gm, nil, false); w.Code != http.StatusSeeOther || w.Header().Get("Location") != "/compendium?import=fivetools" {
		t.Errorf("plain visit: %d %q", w.Code, w.Header().Get("Location"))
	}
	// An empty pick is refused before anything is downloaded.
	if w := e.do(t, http.MethodPost, "/compendium/5etools", gm, url.Values{}, false); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Pick at least one book") {
		t.Errorf("empty pick: %d", w.Code)
	}
}

func TestSourceMove(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, auth.RoleAdmin)
	e.source(t, "srd-2024", "2024") // 200
	e.source(t, "book-a", "2024")   // 100
	e.source(t, "book-b", "2014")   // 300
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
	tests := []struct {
		name, key, dir, want string
		code                 int
	}{
		{"start", "", "", "homebrew,book-a,srd-2024,book-b", 0},
		{"down", "book-a", "down", "homebrew,srd-2024,book-a,book-b", http.StatusOK},
		{"up", "book-b", "up", "homebrew,srd-2024,book-b,book-a", http.StatusOK},
		{"first stays", "homebrew", "up", "homebrew,srd-2024,book-b,book-a", http.StatusOK},
		{"last stays", "book-a", "down", "homebrew,srd-2024,book-b,book-a", http.StatusOK},
		{"bad direction", "book-a", "sideways", "homebrew,srd-2024,book-b,book-a", http.StatusUnprocessableEntity},
		{"unknown source", "nope", "up", "homebrew,srd-2024,book-b,book-a", http.StatusUnprocessableEntity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.code != 0 {
				if w := e.do(t, http.MethodPost, "/compendium/sources/"+tt.key+"/move", gm, url.Values{"dir": {tt.dir}}, true); w.Code != tt.code {
					t.Fatalf("code %d, want %d: %s", w.Code, tt.code, w.Body)
				}
			}
			if got := order(); got != tt.want {
				t.Errorf("order %s, want %s", got, tt.want)
			}
		})
	}
}

// TestSourceOrder: a drag posts every key in the new order; a list that
// misses a source (a stale page) changes nothing.
func TestSourceOrder(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, auth.RoleAdmin)
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
	gm := e.token(t, auth.RoleAdmin)
	e.source(t, "5etools-phb", "2014")
	e.app.fiveTools.start(4)
	e.app.fiveTools.update(func(j *views.FiveToolsJob) { j.Done, j.Current = 1, "Test Book" })
	body := e.do(t, http.MethodGet, "/compendium?import="+views.ImportTabFiveTools, gm, nil, false).Body.String()
	if !strings.Contains(body, "Downloading 2 of 4: Test Book") || !regexp.MustCompile(`style="width:\s?25%;?"`).MatchString(body) {
		t.Error("import tab misses the progress label or bar")
	}
	if w := e.do(t, http.MethodPost, views.FiveToolsPath, gm, url.Values{views.FiveToolsField: {"phb"}}, false); !strings.Contains(w.Body.String(), errImportRunning) {
		t.Errorf("second import: %d, not refused", w.Code)
	}
	if w := e.do(t, http.MethodPost, views.FiveToolsPath+"/phb/delete", gm, nil, true); !strings.Contains(w.Body.String(), errImportRunning) {
		t.Errorf("removal during an import: %d, not refused", w.Code)
	}
	// Done, the tab shows the summary instead.
	e.app.fiveTools.update(func(j *views.FiveToolsJob) {
		*j = views.FiveToolsJob{Msg: "5etools-phb: 3 entries imported, 0 removed"}
	})
	if body := e.do(t, http.MethodGet, "/compendium?import="+views.ImportTabFiveTools, gm, nil, false).Body.String(); !strings.Contains(body, "3 entries imported") || strings.Contains(body, "Downloading") {
		t.Error("finished import shows no summary, or still its progress")
	}
}

// TestRunFiveToolsFinishes: a run always ends the job with its summary, so
// the next import is not refused.
func TestRunFiveToolsFinishes(t *testing.T) {
	e := newTestEnv(t)
	if !e.app.fiveTools.start(0) {
		t.Fatal("job busy before the run")
	}
	e.app.runFiveTools(t.Context(), t.TempDir(), nil) // no books: nothing to download
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
	gm := e.token(t, auth.RoleAdmin)
	e.source(t, "5etools-phb", "2014")
	e.source(t, "5etools-xge", "2014")
	e.feat(t, "5etools-phb", "alert", "x")
	e.feat(t, "5etools-xge", "tough", "y")
	folder := filepath.Join(e.app.compendiumDir(), "5etools-phb")
	if err := os.MkdirAll(folder, 0o750); err != nil {
		t.Fatal(err)
	}
	w := e.do(t, http.MethodPost, "/compendium/5etools/phb/delete", gm, nil, true)
	if w.Code != http.StatusOK {
		t.Fatalf("remove: %d %s", w.Code, w.Body)
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
	if w := e.do(t, http.MethodPost, "/compendium/5etools/..%2Fx/delete", gm, nil, true); strings.Contains(w.Body.String(), "Removed") {
		t.Errorf("bad id accepted: %d", w.Code)
	}
}

// TestEntryAlsoInAndCrossRefs: two sources hold one feat; each entry page
// names the other, and a [[compendium:...]] token links to the winning copy.
func TestEntryAlsoInAndCrossRefs(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, auth.RoleAdmin)
	player := e.token(t, auth.RolePlayer)
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
	link := func(slug, label string) string {
		return fmt.Sprintf(`<a href="/compendium/entries/%d">%s</a>`, scoped[slug], label)
	}

	tests := []struct {
		name, cookie, path string
		want, notWant      []string
	}{
		{"older names the newer", gm, fmt.Sprintf("/compendium/entries/%d", older), []string{"Also in", fmt.Sprintf(`href="/compendium/entries/%d"`, newer), "book-new, 2024"}, nil},
		{"newer names the older", player, fmt.Sprintf("/compendium/entries/%d", newer), []string{"Also in", fmt.Sprintf(`href="/compendium/entries/%d"`, older)}, nil},
		{"token links the winner, missing is plain", player, fmt.Sprintf("/compendium/entries/%d", pointer),
			[]string{fmt.Sprintf(`<a href="/compendium/entries/%d">Alert</a>`, newer), "Ghost", "Secret", link("shared", "Shared")}, []string{"Also in", "[[compendium", "/compendium/entries/" + fmt.Sprint(older), link("secret", "Secret")}},
		{"the GM sees every homebrew link", gm, fmt.Sprintf("/compendium/entries/%d", pointer), []string{link("secret", "Secret"), link("shared", "Shared")}, nil},
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
			if strings.Contains(body, `href="/compendium/entries/0"`) || strings.Contains(body, ">Ghost</a>") {
				t.Error("link to a missing entry")
			}
		})
	}
}

// TestSourceDeleteUsed: a source a sheet pins is refused and kept; once the
// pin is gone it is deleted with its folder.
func TestSourceDeleteUsed(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, auth.RoleAdmin)
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
	w := e.do(t, http.MethodPost, "/compendium/sources/my-folder/delete", gm, nil, true)
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
	if w := e.do(t, http.MethodPost, "/compendium/sources/my-folder/delete", gm, nil, true); !strings.Contains(w.Body.String(), "my-folder is used by Goblin boss.") {
		t.Errorf("entity link not counted: %d %s", w.Code, w.Body)
	}
	if err := e.q().DeleteEntity(t.Context(), ent.ID); err != nil {
		t.Fatal(err)
	}
	if w := e.do(t, http.MethodPost, "/compendium/sources/my-folder/delete", gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("delete: %d %s", w.Code, w.Body)
	}
	if _, err := e.q().GetCompendiumSource(t.Context(), "my-folder"); err == nil {
		t.Error("source kept")
	}
	if _, err := os.Stat(folder); !os.IsNotExist(err) {
		t.Errorf("folder kept: %v", err)
	}
}
