package web

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"gm-session-manager/auth"
	"gm-session-manager/db/sqlc"
)

// TestDeleteRestoreImport: a player deletes their character, which hides it
// and logs them out; the GM restores it, downloads it and imports the file
// as a new character.
func TestDeleteRestoreImport(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	player := e.token(t, auth.RolePlayer)
	other, err := e.q().CreateCharacter(ctx, sqlc.CreateCharacterParams{CampaignID: e.campaignID, Name: "Bram", System: "dnd5e"})
	if err != nil {
		t.Fatal(err)
	}
	sheet := fmt.Sprintf("/characters/%d", e.characterID)
	overview := fmt.Sprintf("/campaigns/%d", e.campaignID)

	// Another player's character cannot be deleted.
	if w := e.do(t, http.MethodPost, fmt.Sprintf("/characters/%d/remove", other.ID), player, nil, true); w.Code == http.StatusOK && w.Header().Get("HX-Redirect") != "" {
		t.Fatal("deleted another player's character")
	}
	if _, err := e.q().GetCharacter(ctx, other.ID); err != nil {
		t.Fatalf("other character gone: %v", err)
	}

	w := e.do(t, http.MethodPost, sheet+"/remove", player, nil, true)
	if w.Header().Get("HX-Redirect") != PathHome {
		t.Fatalf("remove: %d %v %s", w.Code, w.Header(), w.Body)
	}
	if _, err := e.q().GetCharacter(ctx, e.characterID); err == nil {
		t.Error("deleted character still visible")
	}
	if w := e.do(t, http.MethodGet, sheet, player, nil, false); w.Code == http.StatusOK && strings.Contains(w.Body.String(), "Alice") {
		t.Error("the deleted character's device still sees it")
	}
	body := e.do(t, http.MethodGet, overview, gm, nil, false).Body.String()
	if !strings.Contains(body, "Deleted characters") || !strings.Contains(body, fmt.Sprintf("/characters/%d/restore", e.characterID)) {
		t.Errorf("overview lacks the deleted character: %s", body)
	}

	if w := e.do(t, http.MethodPost, sheet+"/restore", player, nil, true); w.Code == http.StatusOK {
		t.Error("a visitor restored a character")
	}
	if w := e.do(t, http.MethodPost, sheet+"/restore", gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("restore: %d %s", w.Code, w.Body)
	}
	if _, err := e.q().GetCharacter(ctx, e.characterID); err != nil {
		t.Errorf("restored character hidden: %v", err)
	}

	file := e.do(t, http.MethodGet, sheet+"/download", gm, nil, false)
	if file.Code != http.StatusOK || !strings.Contains(file.Body.String(), `"name": "Alice"`) {
		t.Fatalf("download: %d %s", file.Code, file.Body)
	}
	importPath := fmt.Sprintf("/campaigns/%d/characters/import", e.campaignID)
	if w := e.doFile(t, importPath, gm, nil, "file", "alice.json", file.Body.String()); w.Code != http.StatusSeeOther {
		t.Fatalf("import: %d %s", w.Code, w.Body)
	}
	if w := e.doFile(t, importPath, gm, nil, "file", "bad.json", `{"name":"X","system":"dnd5e","sheet":{"abilities":{"str":99}}}`); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("invalid sheet imported: %d", w.Code)
	}
	chars, _ := e.q().ListCharactersByCampaign(ctx, e.campaignID)
	if len(chars) != 3 {
		t.Errorf("after import: %d characters", len(chars))
	}
}

// TestRenameDeleteCampaign: only the GM renames or deletes a campaign, and
// deleting it removes its characters and so their devices.
func TestRenameDeleteCampaign(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	player := e.token(t, auth.RolePlayer)
	overview := fmt.Sprintf("/campaigns/%d", e.campaignID)
	name := url.Values{"name": {"  Curse of Strahd  "}}

	if w := e.do(t, http.MethodPost, overview+"/rename", player, name, true); w.Code < http.StatusBadRequest {
		t.Errorf("player rename: %d", w.Code)
	}
	if w := e.do(t, http.MethodPost, overview+"/delete", player, nil, true); w.Code < http.StatusBadRequest {
		t.Errorf("player delete: %d", w.Code)
	}

	if w := e.do(t, http.MethodPost, overview+"/rename", gm, name, true); w.Header().Get("HX-Redirect") != overview {
		t.Fatalf("rename: %d %v %s", w.Code, w.Header(), w.Body)
	}
	if camp, err := e.q().GetCampaign(ctx, e.campaignID); err != nil || camp.Name != "Curse of Strahd" {
		t.Errorf("renamed campaign: %q %v", camp.Name, err)
	}

	if w := e.do(t, http.MethodPost, overview+"/delete", gm, nil, true); w.Header().Get("HX-Redirect") != PathHome {
		t.Fatalf("delete: %d %v %s", w.Code, w.Header(), w.Body)
	}
	if w := e.do(t, http.MethodGet, overview, gm, nil, false); w.Code != http.StatusNotFound {
		t.Errorf("deleted campaign: %d", w.Code)
	}
	if _, err := e.q().GetCharacterAny(ctx, e.characterID); err == nil {
		t.Error("character outlived its campaign")
	}
}
