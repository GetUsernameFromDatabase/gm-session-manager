package web

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"gm-session-manager/auth"
	"gm-session-manager/compendium"
	"gm-session-manager/dnd5e"
)

// TestRevertLevelUp checks level-ups revert newest first, back to the sheet
// and max HP from before each, keeping current HP up to that max.
func TestRevertLevelUp(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/characters/%d/levelup", e.characterID)
	start, _ := e.q().GetCharacter(ctx, e.characterID)
	var after []int64
	for range 2 {
		if w := e.do(t, http.MethodPost, base, player, url.Values{"class": {"wizard"}}, false); w.Code != http.StatusSeeOther {
			t.Fatalf("level up: %d %s", w.Code, w.Body)
		}
		ch, _ := e.q().GetCharacter(ctx, e.characterID)
		after = append(after, ch.MaxHp)
	}
	if after[1] <= after[0] {
		t.Fatalf("max HP did not grow: %v", after)
	}
	for i, wantMax := range []int64{after[0], start.MaxHp} {
		if w := e.do(t, http.MethodPost, base+"/revert", player, nil, true); w.Code != http.StatusOK {
			t.Fatalf("revert %d: %d %s", i+1, w.Code, w.Body)
		}
		ch, _ := e.q().GetCharacter(ctx, e.characterID)
		if ch.MaxHp != wantMax || ch.Hp > ch.MaxHp {
			t.Errorf("revert %d: hp %d/%d, want max %d", i+1, ch.Hp, ch.MaxHp, wantMax)
		}
	}
	ch, _ := e.q().GetCharacter(ctx, e.characterID)
	startSheet, _ := dnd5e.ParseSheet(start.SheetJson)
	if want, _ := startSheet.JSON(); ch.SheetJson != want {
		t.Errorf("sheet not restored:\n got %s\nwant %s", ch.SheetJson, want)
	}
	if w := e.do(t, http.MethodPost, base+"/revert", player, nil, true); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("revert with nothing left: %d, want 422", w.Code)
	}
}

// TestRevertLevelUpInterleaved checks level-ups and reverts may alternate:
// each revert returns to the sheet from just before the newest level-up.
func TestRevertLevelUpInterleaved(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/characters/%d/levelup", e.characterID)
	start, _ := e.q().GetCharacter(ctx, e.characterID)
	startSheet, _ := dnd5e.ParseSheet(start.SheetJson)
	wantStart, _ := startSheet.JSON()
	maxHP := func() int64 {
		ch, _ := e.q().GetCharacter(ctx, e.characterID)
		return ch.MaxHp
	}
	levelUp := func() {
		t.Helper()
		if w := e.do(t, http.MethodPost, base, player, url.Values{"class": {"wizard"}}, false); w.Code != http.StatusSeeOther {
			t.Fatalf("level up: %d %s", w.Code, w.Body)
		}
	}
	revert := func() {
		t.Helper()
		if w := e.do(t, http.MethodPost, base+"/revert", player, nil, true); w.Code != http.StatusOK {
			t.Fatalf("revert: %d %s", w.Code, w.Body)
		}
	}

	levelUp()
	revert()
	if n, _ := e.q().CountLevelUps(ctx, e.characterID); n != 0 || maxHP() != start.MaxHp {
		t.Fatalf("after up, revert: %d level-ups left, max HP %d, want 0 and %d", n, maxHP(), start.MaxHp)
	}

	levelUp()
	once := maxHP()
	levelUp()
	revert()
	if got := maxHP(); got != once {
		t.Errorf("after up, up, revert: max HP %d, want %d", got, once)
	}

	levelUp()
	revert()
	revert()
	if ch, _ := e.q().GetCharacter(ctx, e.characterID); ch.SheetJson != wantStart || ch.MaxHp != start.MaxHp {
		t.Errorf("not back at the start:\n got %s (max %d)\nwant %s (max %d)", ch.SheetJson, ch.MaxHp, wantStart, start.MaxHp)
	}
	if w := e.do(t, http.MethodPost, base+"/revert", player, nil, true); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("revert with nothing left: %d, want 422", w.Code)
	}
}
