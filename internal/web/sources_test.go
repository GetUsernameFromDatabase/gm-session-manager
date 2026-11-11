package web

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"testing"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/perm"
)

// Test sources: "old-src" (2014, priority 300) is outranked by "new-src"
// (2024, priority 100).
const (
	oldSrc = "old-src"
	newSrc = "new-src"
)

func (e testEnv) setSheet(t *testing.T, sh dnd5e.Sheet) {
	t.Helper()
	raw, err := sh.JSON()
	if err != nil {
		t.Fatal(err)
	}
	if err := e.q().UpdateCharacterSheet(t.Context(), sqlc.UpdateCharacterSheetParams{ID: e.characterID, SheetJson: raw}); err != nil {
		t.Fatal(err)
	}
}

func (e testEnv) sheet(t *testing.T) dnd5e.Sheet {
	t.Helper()
	ch, err := e.q().GetCharacter(t.Context(), e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	sh, err := dnd5e.ParseSheet(ch.SheetJson)
	if err != nil {
		t.Fatal(err)
	}
	return sh
}

func TestEntryFor(t *testing.T) {
	e := newTestEnv(t)
	e.source(t, oldSrc, "2014")
	e.source(t, newSrc, "2024")
	e.feat(t, oldSrc, "tough", "old")
	e.feat(t, newSrc, "tough", "new")
	e.feat(t, oldSrc, "alert", "only old")
	tests := []struct {
		name, slug, pin, want string
		missing               bool
	}{
		{"pinned and present", "tough", oldSrc, oldSrc, false},
		{"pinned and missing falls back to priority", "alert", "gone-src", oldSrc, true},
		{"unpinned follows priority", "tough", "", newSrc, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row, missing, err := entryFor(t.Context(), e.q(), enum.KindFeat, tt.slug, tt.pin)
			if err != nil || row.SourceKey != tt.want || missing != tt.missing {
				t.Errorf("got %q missing=%v err=%v, want %q missing=%v", row.SourceKey, missing, err, tt.want, tt.missing)
			}
		})
	}
	if _, _, err := entryFor(t.Context(), e.q(), enum.KindFeat, "nope", oldSrc); !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("neither copy: %v", err)
	}
}

func TestSourceIssues(t *testing.T) {
	e := newTestEnv(t)
	e.source(t, oldSrc, "2014")
	e.source(t, newSrc, "2024")
	e.feat(t, oldSrc, "tough", "old")
	e.feat(t, newSrc, "tough", "new")
	tests := []struct {
		name, pin     string
		issue, missed bool
	}{
		{"pinned to the best copy", newSrc, false, false},
		{"pinned and outranked", oldSrc, true, false},
		{"pinned and missing", "gone-src", true, true},
		{"unpinned", "", false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := dnd5e.Sheet{Feats: []string{"Tough"}, FeatSources: map[string]string{"tough": tt.pin}}
			got, err := e.app.sourceIssues(t.Context(), sh, e.campaignID)
			if err != nil {
				t.Fatal(err)
			}
			if (len(got) == 1) != tt.issue || len(got) > 1 {
				t.Fatalf("issues %+v", got)
			}
			if tt.issue && (got[0].Missing != tt.missed || got[0].Best != newSrc || got[0].Kind != dnd5e.RefFeat || len(got[0].Copies) != 2 || got[0].Copies[1].Label != "old-src, 2014") {
				t.Errorf("issue %+v", got[0])
			}
		})
	}
}

// TestSourceWarningAfterImport: a sheet pinned to the only copy has no
// issue until a higher-priority copy arrives, for the owner and the GM only.
func TestSourceWarningAfterImport(t *testing.T) {
	e := newTestEnv(t)
	e.source(t, oldSrc, "2014")
	e.feat(t, oldSrc, "tough", "old")
	e.setSheet(t, dnd5e.Sheet{Feats: []string{"Tough"}, FeatSources: map[string]string{"tough": oldSrc}})
	issues := func(v perm.Viewer) int {
		ch, err := e.q().GetCharacter(t.Context(), e.characterID)
		if err != nil {
			t.Fatal(err)
		}
		s, err := e.app.sheetView(t.Context(), v, ch)
		if err != nil {
			t.Fatal(err)
		}
		if s.CanSources != (v.IsGM || v.CharacterID == e.characterID) {
			t.Errorf("CanSources %v for %+v", s.CanSources, v)
		}
		return len(s.SourceIssues)
	}
	owner := perm.Player(e.characterID)
	if n := issues(owner); n != 0 {
		t.Fatalf("before the import: %d issues", n)
	}
	e.source(t, newSrc, "2024")
	e.feat(t, newSrc, "tough", "new")
	for _, v := range []perm.Viewer{owner, perm.GM} {
		if n := issues(v); n != 1 {
			t.Errorf("after the import, %+v: %d issues", v, n)
		}
	}
	if n := issues(perm.Player(e.characterID + 1)); n != 0 {
		t.Errorf("another player: %d issues", n)
	}
}

func TestSourceCommandPermissions(t *testing.T) {
	e := newTestEnv(t)
	e.source(t, oldSrc, "2014")
	e.source(t, newSrc, "2024")
	e.feat(t, oldSrc, "tough", "old")
	e.feat(t, newSrc, "tough", "new")
	e.setSheet(t, dnd5e.Sheet{Feats: []string{"Tough"}, FeatSources: map[string]string{"tough": oldSrc}})
	other, err := e.q().CreateCharacter(t.Context(), sqlc.CreateCharacterParams{CampaignID: e.campaignID, Name: "Bob", System: "dnd5e"})
	if err != nil {
		t.Fatal(err)
	}
	bob, _, err := e.mgr.IssuePlayer(t.Context(), other.ID)
	if err != nil {
		t.Fatal(err)
	}
	e.allowEntrySources(t)
	cmdPath := fmt.Sprintf("/characters/%d/cmd", e.characterID)
	tests := []struct {
		name, cookie, to string
		ok               bool
	}{
		{"other player refused", bob, newSrc, false},
		{"gm", e.token(t, enum.RoleAdmin), newSrc, true},
		{"owner", e.token(t, enum.RolePlayer), oldSrc, true},
		{"source without the feat", e.token(t, enum.RoleAdmin), "gone-src", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := e.sheet(t).FeatSources["tough"]
			w := e.do(t, http.MethodPost, cmdPath, tt.cookie, url.Values{"cmd": {"source feat tough " + tt.to}}, true)
			after := e.sheet(t).FeatSources["tough"]
			if (w.Code == http.StatusOK) != tt.ok || tt.ok && after != tt.to || !tt.ok && after != before {
				t.Errorf("status %d, pin %q → %q: %s", w.Code, before, after, w.Body)
			}
		})
	}
}

// classCopy is a test class "Testclass" from source with features.
func classCopy(source, features string) sqlc.UpsertCompendiumEntryParams {
	return sqlc.UpsertCompendiumEntryParams{Kind: "class", Slug: "testclass", Name: "Testclass", SourceKey: source, ContentHash: source,
		DataJson: `{"hit_die":10,"features":[` + features + `]}`}
}

// TestRunCommandUsesCurrentSheet: a command's compendium lookups read the
// sheet inside its transaction, not the caller's earlier copy.
func TestRunCommandUsesCurrentSheet(t *testing.T) {
	e := newTestEnv(t)
	e.source(t, oldSrc, "2014")
	e.source(t, newSrc, "2024")
	for _, c := range []sqlc.UpsertCompendiumEntryParams{
		classCopy(oldSrc, `{"name":"Strike","desc":"old strike","levels":[1]}`),
		classCopy(newSrc, `{"name":"Strike","desc":"new strike","levels":[1]}`),
	} {
		if err := e.q().UpsertCompendiumEntry(t.Context(), c); err != nil {
			t.Fatal(err)
		}
	}
	stale, err := e.q().GetCharacter(t.Context(), e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	stale.SheetJson = "{}" // read before another device added the class
	e.allowEntrySources(t)
	e.setSheet(t, dnd5e.Sheet{
		Classes:  []dnd5e.ClassLevel{{Slug: "testclass", Name: "Testclass", Level: 1, HitDie: 8, Source: oldSrc}},
		Features: []dnd5e.Feature{{Name: "Strike", Source: "Testclass", Level: 1, Description: "old strike"}},
	})
	cmd, err := dnd5e.Parse("source class testclass " + newSrc)
	if err != nil {
		t.Fatal(err)
	}
	if err := e.app.runCommand(t.Context(), stale, cmd); err != nil {
		t.Fatal(err)
	}
	if sh := e.sheet(t); len(sh.Features) != 1 || sh.Features[0].Description != "new strike" || sh.Classes[0].Source != newSrc {
		t.Errorf("sheet %+v", sh)
	}
}

// TestSourceSwitchKeepsEdits: switching a class replaces untouched features,
// keeps an edited one and adds the new copy's features for levels taken,
// through the preview and the command.
func TestSourceSwitchKeepsEdits(t *testing.T) {
	e := newTestEnv(t)
	e.source(t, oldSrc, "2014")
	e.source(t, newSrc, "2024")
	for _, c := range []sqlc.UpsertCompendiumEntryParams{
		classCopy(oldSrc, `{"name":"Strike","desc":"old strike","levels":[1]},{"name":"Guard","desc":"old guard","levels":[1]}`),
		classCopy(newSrc, `{"name":"Strike","desc":"new strike","levels":[1]},{"name":"Guard","desc":"new guard","levels":[1]},{"name":"Rally","desc":"new rally","levels":[2]}`),
	} {
		if err := e.q().UpsertCompendiumEntry(t.Context(), c); err != nil {
			t.Fatal(err)
		}
	}
	e.setSheet(t, dnd5e.Sheet{
		Classes: []dnd5e.ClassLevel{{Slug: "testclass", Name: "Testclass", Level: 2, HitDie: 8, Source: oldSrc}},
		Features: []dnd5e.Feature{
			{Name: "Strike", Source: "Testclass", Level: 1, Description: "old strike"},
			{Name: "Guard", Source: "Testclass", Level: 1, Description: "my guard"},
		},
	})
	e.allowEntrySources(t)
	player := e.token(t, enum.RolePlayer)
	base := fmt.Sprintf("/characters/%d/sources/preview?", e.characterID)
	preview := url.Values{"kind": {"class"}, "slug": {"testclass"}, "source": {newSrc}}
	if w := e.do(t, http.MethodGet, base+preview.Encode(), player, nil, false); w.Code != http.StatusOK {
		t.Errorf("preview: %d %s", w.Code, w.Body)
	}
	for _, q := range []url.Values{{"kind": {"race"}, "slug": {"testclass"}, "source": {newSrc}}, {"kind": {"class"}, "slug": {"testclass"}, "source": {"gone-src"}}} {
		if w := e.do(t, http.MethodGet, base+q.Encode(), player, nil, false); w.Code != http.StatusUnprocessableEntity {
			t.Errorf("preview %v: %d", q, w.Code)
		}
	}
	env, name, err := e.app.sourceEnv(t.Context(), e.app.Store.Queries, e.sheet(t), dnd5e.SourceRef{Kind: dnd5e.RefClass, Slug: "testclass", Source: newSrc}, e.campaignID)
	if err != nil {
		t.Fatal(err)
	}
	if name == "" {
		t.Error("sourceEnv returned no name for the target copy")
	}
	plan := dnd5e.PlanSourceSwitch(e.sheet(t), dnd5e.RefClass, "testclass", env.SourceOld, env.SourceNew)
	if len(plan.Replace) != 1 || plan.Replace[0].Name != "Strike" || len(plan.Keep) != 1 || plan.Keep[0].Name != "Guard" || len(plan.Add) != 1 || plan.Add[0].Name != "Rally" {
		t.Errorf("plan %+v", plan)
	}

	w := e.do(t, http.MethodPost, fmt.Sprintf("/characters/%d/cmd", e.characterID), player, url.Values{"cmd": {"source class testclass " + newSrc}}, true)
	if w.Code != http.StatusOK {
		t.Fatalf("command: %d %s", w.Code, w.Body)
	}
	sh := e.sheet(t)
	desc := func(name string) string {
		i := slices.IndexFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == name })
		if i < 0 {
			return ""
		}
		return sh.Features[i].Description
	}
	if desc("Strike") != "new strike" || desc("Guard") != "my guard" || desc("Rally") != "new rally" {
		t.Errorf("features %+v", sh.Features)
	}
	if cl := sh.Classes[0]; cl.Source != newSrc || cl.HitDie != 10 {
		t.Errorf("class %+v", cl)
	}
}
