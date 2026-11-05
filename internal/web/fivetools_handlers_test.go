package web

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/db/sqlc"
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
	text := "See [[compendium:feat/alert|Alert]] and [[compendium:feat/ghost|Ghost]]."
	newer := e.feat(t, "book-new", "alert", "newer")
	older := e.feat(t, "book-old", "alert", "older")
	pointer := e.feat(t, "book-new", "pointer", text)

	tests := []struct {
		name, cookie, path string
		want, notWant      []string
	}{
		{"older names the newer", gm, fmt.Sprintf("/compendium/entries/%d", older), []string{"Also in", fmt.Sprintf(`href="/compendium/entries/%d"`, newer), "book-new, 2024"}, nil},
		{"newer names the older", player, fmt.Sprintf("/compendium/entries/%d", newer), []string{"Also in", fmt.Sprintf(`href="/compendium/entries/%d"`, older)}, nil},
		{"token links the winner, missing is plain", player, fmt.Sprintf("/compendium/entries/%d", pointer),
			[]string{fmt.Sprintf(`<a href="/compendium/entries/%d">Alert</a>`, newer), "Ghost"}, []string{"Also in", "[[compendium", "/compendium/entries/" + fmt.Sprint(older)}},
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
