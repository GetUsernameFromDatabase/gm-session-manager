package web

import (
	"net/http"
	"net/url"
	"testing"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

// TestSheetRoutesPermissions: the sheet's page routes (wizard, level-up,
// builder, completions, source preview) answer the owner and nobody else
// but the GM where perm says so. Requests are not htmx, so a refusal is the
// real 403 and not the 422 toast htmx gets.
func TestSheetRoutesPermissions(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	newPlayer := func(campaignID int64, name string) string {
		t.Helper()
		ch, err := e.q().
			CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{CampaignID: campaignID, Name: name, System: "dnd5e", SheetJson: "{}"})
		if err != nil {
			t.Fatal(err)
		}
		tok, _, err := e.mgr.IssuePlayer(ctx, ch.ID)
		if err != nil {
			t.Fatal(err)
		}
		return tok
	}
	elsewhere, err := e.q().
		CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Elsewhere", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	owner, gm := e.token(t, enum.RolePlayer), e.token(t, enum.RoleAdmin)
	strangers := []struct{ name, token string }{
		{"player in the same campaign", newPlayer(e.campaignID, "Bob")},
		{"player in another campaign", newPlayer(elsewhere.ID, "Cy")},
	}
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	e.allowEntrySources(t)
	id := e.characterID
	routes := []struct {
		name, method, path string
		gmMay              bool // perm lets the GM in: view (completions), sources (preview)
	}{
		{
			"wizard page",
			http.MethodGet,
			route.Path(route.CharacterCreateStep, id, views.StepClass),
			false,
		},
		{
			"wizard save",
			http.MethodPost,
			route.Path(route.CharacterCreateStep, id, views.StepClass),
			false,
		},
		{"level-up page", http.MethodGet, route.Path(route.CharacterLevelup, id), false},
		{
			"level-up",
			http.MethodPost,
			route.Path(route.CharacterLevelup, id),
			false,
		}, // owner: class found, no level: 422
		{"level-up revert", http.MethodPost, route.Path(route.CharacterLevelupRevert, id), false},
		{"builder", http.MethodGet, route.Path(route.CharacterBuilder, id), false},
		{"picks", http.MethodGet, route.Path(route.CharacterPicks, id), false},
		{"completions", http.MethodGet, route.Path(route.CharacterComplete, id), true},
		{"source preview", http.MethodGet, route.Path(route.CharacterSourcesPreview, id), true},
	}
	do := func(method, path, token string) int {
		var form url.Values
		if method == http.MethodPost {
			form = url.Values{"class": {"wizard"}} // the level-up needs a class to look up
		}
		return e.do(t, method, path, token, form, false).Code
	}
	for _, r := range routes {
		t.Run(r.name, func(t *testing.T) {
			// An empty form may be refused as invalid (422), but never as forbidden or missing.
			if code := do(
				r.method,
				r.path,
				owner,
			); code == http.StatusForbidden || code == http.StatusNotFound ||
				code >= http.StatusInternalServerError {
				t.Errorf("owner: %d", code)
			}
			for _, s := range strangers {
				if code := do(r.method, r.path, s.token); code != http.StatusForbidden {
					t.Errorf("%s: %d, want 403", s.name, code)
				}
			}
		})
	}
	// The GM is not a stranger everywhere: only the routes perm gives them.
	for _, r := range routes {
		t.Run("gm "+r.name, func(t *testing.T) {
			code := do(r.method, r.path, gm)
			if forbidden := code == http.StatusForbidden; forbidden == r.gmMay ||
				code >= http.StatusInternalServerError {
				t.Errorf("gm: %d, may %v", code, r.gmMay)
			}
		})
	}
}
