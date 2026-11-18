package web

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
)

// deathFix is the effects fixture with Cass in the encounter at 0 HP and
// the given death saves.
func deathFix(t *testing.T, ds dnd5e.DeathSaves) (effectFix, sqlc.Combatant) {
	t.Helper()
	f := newEffectFix(t)
	cb, err := f.q().CreateCombatant(t.Context(), sqlc.CreateCombatantParams{EncounterID: f.enc.ID, CharacterID: nullID(f.cass.ID), Name: f.cass.Name})
	if err != nil {
		t.Fatal(err)
	}
	if err := f.app.updateCharacter(t.Context(), f.cass.ID, func(_ *sqlc.Queries, dc *dnd5e.Character) error {
		dc.HP, dc.MaxHP, dc.Sheet.DeathSaves = 0, 20, ds
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	return f, cb
}

func (f effectFix) saves(t *testing.T) dnd5e.DeathSaves {
	t.Helper()
	row, err := f.q().GetCharacter(t.Context(), f.cass.ID)
	if err != nil {
		t.Fatal(err)
	}
	s, err := dnd5e.ParseSheet(row.SheetJson)
	if err != nil {
		t.Fatal(err)
	}
	return s.DeathSaves
}

func TestTrackerRefusesWhileDead(t *testing.T) {
	f, cb := deathFix(t, dnd5e.DeathSaves{Failures: dnd5e.MaxDeathSaves})
	hp := fmt.Sprintf("/combatants/%d/hp", cb.ID)
	for name, form := range map[string]url.Values{
		"heal":     {"op": {"heal"}, "amount": {"5"}},
		"set hp":   {"op": {"set"}, "hp": {"5"}},
		"set temp": {"op": {"set"}, "temp_hp": {"5"}},
	} {
		t.Run(name, func(t *testing.T) {
			w := f.do(t, http.MethodPost, hp, f.gm, form, true)
			if w.Code != http.StatusUnprocessableEntity || !strings.Contains(w.Body.String(), "Dead. Use revive.") {
				t.Errorf("%d %s", w.Code, w.Body)
			}
		})
	}
	// Max HP still changes, and damage still lowers HP.
	if w := f.do(t, http.MethodPost, hp, f.gm, url.Values{"op": {"damage"}, "amount": {"3"}}, true); w.Code != http.StatusOK {
		t.Errorf("damage: %d %s", w.Code, w.Body)
	}
}

// The sheet's and own row's quick form posts hp_crit, not the dialog's crit.
func TestQuickFormCritAtZeroHP(t *testing.T) {
	f, _ := deathFix(t, dnd5e.DeathSaves{})
	form := url.Values{"hp_op": {"damage"}, "hp_amount": {"2"}, "hp_crit": {"true"}}
	if w := f.do(t, http.MethodPost, fmt.Sprintf("/characters/%d/cmd", f.cass.ID), f.gm, form, true); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	if got := f.saves(t); got.Failures != 2 {
		t.Errorf("saves after a crit: %+v", got)
	}
}

func TestTrackerCritAtZeroHP(t *testing.T) {
	f, cb := deathFix(t, dnd5e.DeathSaves{})
	hp := fmt.Sprintf("/combatants/%d/hp", cb.ID)
	if w := f.do(t, http.MethodPost, hp, f.gm, url.Values{"op": {"damage"}, "amount": {"3"}, "crit": {"true"}}, true); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	if got := f.saves(t); got.Failures != 2 {
		t.Errorf("saves after a crit: %+v", got)
	}
	if w := f.do(t, http.MethodPost, hp, f.gm, url.Values{"op": {"damage"}, "amount": {"3"}}, true); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	if got := f.saves(t); got.Failures != 3 {
		t.Errorf("saves after a plain hit: %+v", got)
	}
}

// A monster or NPC at 0 HP takes damage and heals as before, crit or not.
func TestMonsterAtZeroHPUnaffected(t *testing.T) {
	f := newEffectFix(t)
	hp := fmt.Sprintf("/combatants/%d/hp", f.mage.ID)
	for _, form := range []url.Values{
		{"op": {"damage"}, "amount": {"10"}, "crit": {"true"}},
		{"op": {"damage"}, "amount": {"4"}},
		{"op": {"heal"}, "amount": {"6"}},
	} {
		if w := f.do(t, http.MethodPost, hp, f.gm, form, true); w.Code != http.StatusOK {
			t.Fatalf("%v: %d %s", form, w.Code, w.Body)
		}
	}
	if cb, _ := f.q().GetCombatant(t.Context(), f.mage.ID); cb.Hp != 6 {
		t.Errorf("HP %d, want 6", cb.Hp)
	}
}

// Dying with three failures drops the character's conditions and the
// effects its concentration holds, however the failures came.
func TestDeathClearsConditions(t *testing.T) {
	tests := []struct {
		name string
		fire func(t *testing.T, f effectFix, cb sqlc.Combatant)
	}{
		{"tracker crit damage", func(t *testing.T, f effectFix, cb sqlc.Combatant) {
			f.postForm(t, fmt.Sprintf("/combatants/%d/hp", cb.ID), url.Values{"op": {"damage"}, "amount": {"2"}, "crit": {"true"}})
		}},
		{"sheet deathsave fail", func(t *testing.T, f effectFix, _ sqlc.Combatant) {
			tok, _, err := f.mgr.IssuePlayer(t.Context(), f.cass.ID) // deathsave is the owner's
			if err != nil {
				t.Fatal(err)
			}
			f.cmd(t, tok, f.cass.ID, "deathsave fail")
		}},
		{"sheet crit command", func(t *testing.T, f effectFix, _ sqlc.Combatant) { f.cmd(t, f.gm, f.cass.ID, "hp -2 crit") }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, cb := deathFix(t, dnd5e.DeathSaves{Failures: 2})
			holder := f.concentrating(t, f.cass.ID, 0)
			f.row(t, sqlc.CreateConditionParams{CharacterID: nullID(f.cass.ID), Slug: dnd5e.Prone})
			if w := f.effect(t, f.gm, f.bram.ID, acBonusForm(shieldOfFaith, holder.ID)); w.Code != http.StatusOK {
				t.Fatalf("effect: %d %s", w.Code, w.Body)
			}
			wait := watchSheets(t, f, f.cass.ID, f.bram.ID)
			tt.fire(t, f, cb)
			if got := f.saves(t); got.Failures != dnd5e.MaxDeathSaves {
				t.Fatalf("saves %+v", got)
			}
			for _, id := range []int64{f.cass.ID, f.bram.ID} {
				if rows := f.conds(t, id); len(rows) != 0 {
					t.Errorf("character %d keeps %+v", id, rows)
				}
			}
			wait()
		})
	}
}

func (f effectFix) postForm(t *testing.T, path string, form url.Values) {
	t.Helper()
	if w := f.do(t, http.MethodPost, path, f.gm, form, true); w.Code != http.StatusOK {
		t.Fatalf("%s: %d %s", path, w.Code, w.Body)
	}
}
