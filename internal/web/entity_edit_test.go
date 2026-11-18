package web

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/route"
)

// TestEntityCommands: each NPC edit changes only what it names, so a stale
// page on a second GM device cannot undo another device's change.
func TestEntityCommands(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, enum.RoleAdmin)
	npc, err := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: e.campaignID, Kind: "npc", Name: "Mira", FieldsJson: `{"role":{"value":"Innkeeper"}}`})
	if err != nil {
		t.Fatal(err)
	}
	base := route.Path(route.Entity, npc.ID)
	steps := []struct {
		name, path string
		form       url.Values
		code       int
		wantName   string
		revealed   bool
		fields     perm.Fields
	}{
		{"reveal", route.Reveal, url.Values{"revealed": {"true"}}, http.StatusOK, "Mira", true, perm.Fields{"role": {Value: "Innkeeper"}}},
		// A stale page renames without sending the reveal it never saw.
		{"rename", route.Rename, url.Values{"name": {" Mira Vale "}}, http.StatusOK, "Mira Vale", true, perm.Fields{"role": {Value: "Innkeeper"}}},
		{"add field", route.Fields, url.Values{"key": {" Secret Name "}}, http.StatusOK, "Mira Vale", true, perm.Fields{"role": {Value: "Innkeeper"}, "secret name": {}}},
		{"add existing keeps it", route.Fields, url.Values{"key": {"role"}}, http.StatusOK, "Mira Vale", true, perm.Fields{"role": {Value: "Innkeeper"}, "secret name": {}}},
		{"field text", route.FieldsText, url.Values{"key": {"secret name"}, "value": {" Lyra "}}, http.StatusOK, "Mira Vale", true, perm.Fields{"role": {Value: "Innkeeper"}, "secret name": {Value: "Lyra"}}},
		{"reveal field", route.FieldsReveal, url.Values{"key": {"role"}, "revealed": {"true"}}, http.StatusOK, "Mira Vale", true, perm.Fields{"role": {Value: "Innkeeper", IsRevealed: true}, "secret name": {Value: "Lyra"}}},
		// Text from a page that never saw the reveal keeps the field revealed.
		{"text keeps reveal", route.FieldsText, url.Values{"key": {"role"}, "value": {"Barkeep"}}, http.StatusOK, "Mira Vale", true, perm.Fields{"role": {Value: "Barkeep", IsRevealed: true}, "secret name": {Value: "Lyra"}}},
		{"remove field", route.FieldsRemove, url.Values{"key": {"secret name"}}, http.StatusOK, "Mira Vale", true, perm.Fields{"role": {Value: "Barkeep", IsRevealed: true}}},
		{"remove is idempotent", route.FieldsRemove, url.Values{"key": {"secret name"}}, http.StatusOK, "Mira Vale", true, perm.Fields{"role": {Value: "Barkeep", IsRevealed: true}}},
		// A stale page must not bring back a field another device removed.
		{"text on a removed field", route.FieldsText, url.Values{"key": {"secret name"}, "value": {"Lyra"}}, http.StatusUnprocessableEntity, "Mira Vale", true, perm.Fields{"role": {Value: "Barkeep", IsRevealed: true}}},
		{"reveal a removed field", route.FieldsReveal, url.Values{"key": {"secret name"}, "revealed": {"true"}}, http.StatusUnprocessableEntity, "Mira Vale", true, perm.Fields{"role": {Value: "Barkeep", IsRevealed: true}}},
		{"empty name refused", route.Rename, url.Values{"name": {""}}, http.StatusUnprocessableEntity, "Mira Vale", true, perm.Fields{"role": {Value: "Barkeep", IsRevealed: true}}},
		{"empty key refused", route.Fields, url.Values{"key": {" "}}, http.StatusUnprocessableEntity, "Mira Vale", true, perm.Fields{"role": {Value: "Barkeep", IsRevealed: true}}},
		{"hide", route.Reveal, url.Values{}, http.StatusOK, "Mira Vale", false, perm.Fields{"role": {Value: "Barkeep", IsRevealed: true}}},
	}
	for _, s := range steps {
		if w := e.do(t, http.MethodPost, base+s.path, gm, s.form, true); w.Code != s.code {
			t.Fatalf("%s: status %d %s", s.name, w.Code, w.Body)
		}
		got, err := e.q().GetEntity(ctx, npc.ID)
		if err != nil {
			t.Fatal(err)
		}
		fields, err := perm.ParseFields(got.FieldsJson)
		if err != nil {
			t.Fatal(err)
		}
		if got.Name != s.wantName || got.IsRevealed != s.revealed || fmt.Sprint(fields) != fmt.Sprint(s.fields) {
			t.Errorf("%s: name %q revealed %v fields %v", s.name, got.Name, got.IsRevealed, fields)
		}
	}

	// A long backstory still saves.
	long := strings.Repeat("Lore. ", 2000)
	if w := e.do(t, http.MethodPost, base+route.FieldsText, gm, url.Values{"key": {"role"}, "value": {long}}, true); w.Code != http.StatusOK {
		t.Errorf("long text: status %d", w.Code)
	}
	if w := e.do(t, http.MethodPost, route.Path(route.EntityRename, 9999), gm, url.Values{"name": {"Nobody"}}, false); w.Code != http.StatusNotFound {
		t.Errorf("missing NPC: status %d", w.Code)
	}
}

// TestEntityCommandsGMOnly: a player gets no NPC edit route at all, and
// nothing changes.
func TestEntityCommandsGMOnly(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	player := e.token(t, enum.RolePlayer)
	const fields = `{"role":{"value":"Innkeeper","is_revealed":true}}`
	npc, err := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: e.campaignID, Kind: "npc", Name: "Mira", IsRevealed: true, FieldsJson: fields})
	if err != nil {
		t.Fatal(err)
	}
	base := route.Path(route.Entity, npc.ID)
	for _, r := range []struct {
		path string
		form url.Values
	}{
		{route.Rename, url.Values{"name": {"Hacked"}}},
		{route.Reveal, url.Values{}},
		{route.Fields, url.Values{"key": {"loot"}}},
		{route.FieldsText, url.Values{"key": {"role"}, "value": {"Hacked"}}},
		{route.FieldsReveal, url.Values{"key": {"role"}}},
		{route.FieldsRemove, url.Values{"key": {"role"}}},
	} {
		for _, htmx := range []bool{true, false} {
			if w := e.do(t, http.MethodPost, base+r.path, player, r.form, htmx); w.Code != http.StatusForbidden && w.Code != http.StatusNotFound {
				t.Errorf("player %s (htmx %v): status %d", r.path, htmx, w.Code)
			}
		}
	}
	if got, _ := e.q().GetEntity(ctx, npc.ID); got.Name != "Mira" || !got.IsRevealed || got.FieldsJson != fields {
		t.Errorf("a player changed the NPC: %+v", got)
	}
}

// TestEntityEditor: the GM's editor posts each field on its own and offers
// unused suggestions; players never see an empty field.
func TestEntityEditor(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)
	npc, err := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{
		CampaignID: e.campaignID, Kind: "npc", Name: "Mira", IsRevealed: true,
		FieldsJson: `{"role":{"value":"Innkeeper"},"mood":{"value":"","is_revealed":true}}`,
	})
	if err != nil {
		t.Fatal(err)
	}
	page := route.Path(route.Entity, npc.ID)
	body := e.do(t, http.MethodGet, page, gm, nil, false).Body.String()
	for _, want := range []string{
		fmt.Sprintf(`hx-post="%s/fields/text"`, page), `name="key" value="role"`, `name="value" value="Innkeeper" data-field-value`,
		fmt.Sprintf(`hx-post="%s/rename"`, page), fmt.Sprintf(`data-add-field="%s/fields"`, page), `data-templ-value="motivation"`,
	} {
		if !strings.Contains(body, want) {
			t.Errorf("editor misses %s", want)
		}
	}
	if strings.Contains(body, `data-templ-value="role"`) || strings.Contains(body, `name="count"`) {
		t.Error("editor offers a stored key or still posts the whole form")
	}
	if body := e.do(t, http.MethodGet, page, player, nil, false).Body.String(); strings.Contains(body, ">mood<") {
		t.Error("player sees an empty field")
	}
}
