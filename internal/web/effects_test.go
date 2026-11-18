package web

import (
	"archive/zip"
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

const shieldOfFaith = "Shield of Faith"

// effectFix is a campaign with the GM, two players (Alice, Bram), a caster
// character (Cass) and a monster caster (Mage) in an encounter.
type effectFix struct {
	testEnv
	gm, alice, bramTok string
	bram, cass         sqlc.Character
	enc                sqlc.Encounter
	mage               sqlc.Combatant
}

func newEffectFix(t *testing.T) effectFix {
	t.Helper()
	e := newTestEnv(t)
	ctx := t.Context()
	f := effectFix{testEnv: e, gm: e.token(t, enum.RoleAdmin), alice: e.token(t, enum.RolePlayer)}
	var err error
	if f.bram, err = e.q().CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{CampaignID: e.campaignID, Name: "Bram", System: "dnd5e", SheetJson: "{}"}); err != nil {
		t.Fatal(err)
	}
	if f.cass, err = e.q().CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{CampaignID: e.campaignID, Name: "Cass", System: "dnd5e", SheetJson: "{}"}); err != nil {
		t.Fatal(err)
	}
	if f.bramTok, _, err = e.mgr.IssuePlayer(ctx, f.bram.ID); err != nil {
		t.Fatal(err)
	}
	if f.enc, err = e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"}); err != nil {
		t.Fatal(err)
	}
	if f.mage, err = e.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: f.enc.ID, Name: "Mage", Hp: 10, MaxHp: 10, Ac: 12, Initiative: sql.NullInt64{Int64: 15, Valid: true}}); err != nil {
		t.Fatal(err)
	}
	return f
}

// row inserts a condition row directly.
func (f effectFix) row(t *testing.T, p sqlc.CreateConditionParams) sqlc.Condition {
	t.Helper()
	if p.ModsJson == "" {
		p.ModsJson = noMods
	}
	c, err := f.q().CreateCondition(t.Context(), p)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

// concentrating puts a concentration row on a character or, with a
// combatant id, on a combatant.
func (f effectFix) concentrating(t *testing.T, characterID, combatantID int64) sqlc.Condition {
	t.Helper()
	return f.row(t, sqlc.CreateConditionParams{CharacterID: nullID(characterID), CombatantID: nullID(combatantID), Slug: slugConcentration, Note: "Haste"})
}

// acBonusForm is the effect form for "+2 AC" held by holder (0 for none).
func acBonusForm(name string, holder int64) url.Values {
	form := url.Values{"slug": {"custom"}, "name": {name}, "mod_target": {"ac"}, "mod_value": {"2"}}
	if holder != 0 {
		form.Set("held_by", fmt.Sprint(holder))
	}
	return form
}

func (f effectFix) effect(t *testing.T, cookie string, characterID int64, form url.Values) *httptest.ResponseRecorder {
	t.Helper()
	return f.do(t, http.MethodPost, route.Path(route.CharacterEffect, characterID), cookie, form, true)
}

func (f effectFix) cmd(t *testing.T, cookie string, characterID int64, line string) {
	t.Helper()
	if w := f.do(t, http.MethodPost, route.Path(route.CharacterCmd, characterID), cookie, url.Values{"cmd": {line}}, true); w.Code != http.StatusOK {
		t.Fatalf("%s: %d %s", line, w.Code, w.Body)
	}
}

// shown is the character's sheet as the sheet page shows it, effects included.
func (f effectFix) shown(t *testing.T, id int64) dnd5e.Sheet {
	t.Helper()
	row, err := f.q().GetCharacter(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	s, err := f.app.shownSheetView(t.Context(), perm.GM, row)
	if err != nil {
		t.Fatal(err)
	}
	return s.C.Sheet
}

func (f effectFix) conds(t *testing.T, id int64) []sqlc.Condition {
	t.Helper()
	rows, err := f.q().ListConditionsByCharacter(t.Context(), nullID(id))
	if err != nil {
		t.Fatal(err)
	}
	return rows
}

func notes(conds []sqlc.Condition) []string {
	var out []string
	for _, c := range conds {
		out = append(out, c.Note)
	}
	return out
}

// TestEffectPermissions checks who may put an effect on whom, and which
// concentration row it may hang off.
func TestEffectPermissions(t *testing.T) {
	f := newEffectFix(t)
	ctx := t.Context()
	aliceConc := f.concentrating(t, f.characterID, 0)
	cassConc := f.concentrating(t, f.cass.ID, 0)
	mageConc := f.concentrating(t, 0, f.mage.ID)
	bramProne := f.row(t, sqlc.CreateConditionParams{CharacterID: nullID(f.bram.ID), Slug: "prone"})
	far, err := f.q().CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Far", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	farChar, err := f.q().CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{CampaignID: far.ID, Name: "Far", System: "dnd5e", SheetJson: "{}"})
	if err != nil {
		t.Fatal(err)
	}
	farConc := f.concentrating(t, farChar.ID, 0)

	aliceFx, bramFx, mageFx := func() string { return route.Path(route.CharacterEffect, f.characterID) },
		func() string { return route.Path(route.CharacterEffect, f.bram.ID) },
		func() string { return route.Path(route.CombatantAction, f.mage.ID, "condition") }
	tests := []struct {
		name   string
		cookie string
		path   string
		form   url.Values
		want   int
	}{
		{"GM, a character", f.gm, aliceFx(), acBonusForm("a", 0), http.StatusOK},
		{"GM, a combatant", f.gm, mageFx(), acBonusForm("a", 0), http.StatusOK},
		{"player, own character", f.alice, aliceFx(), acBonusForm("a", 0), http.StatusOK},
		// An htmx refusal shows as a 422 toast; the message tells it from a bad form.
		{"player, another character", f.alice, bramFx(), acBonusForm("a", 0), http.StatusUnprocessableEntity},
		{"GM, held by a character's concentration", f.gm, bramFx(), acBonusForm("a", cassConc.ID), http.StatusOK},
		{"GM, held by a combatant's concentration", f.gm, aliceFx(), acBonusForm("a", mageConc.ID), http.StatusOK},
		{"GM, combatant held by a character's", f.gm, mageFx(), acBonusForm("a", cassConc.ID), http.StatusOK},
		{"player, held by own concentration", f.alice, aliceFx(), acBonusForm("a", aliceConc.ID), http.StatusOK},
		{"player, held by another character's", f.alice, aliceFx(), acBonusForm("a", cassConc.ID), http.StatusUnprocessableEntity},
		{"player, held by a combatant's", f.alice, aliceFx(), acBonusForm("a", mageConc.ID), http.StatusUnprocessableEntity},
		{"holder from another campaign", f.gm, aliceFx(), acBonusForm("a", farConc.ID), http.StatusUnprocessableEntity},
		{"holder that is not concentration", f.gm, aliceFx(), acBonusForm("a", bramProne.ID), http.StatusUnprocessableEntity},
		{"holder that does not exist", f.gm, aliceFx(), acBonusForm("a", 999999), http.StatusUnprocessableEntity},
		{"concentration held by another", f.gm, aliceFx(), func() url.Values {
			form := acBonusForm("a", cassConc.ID)
			form.Set("slug", slugConcentration)
			return form
		}(), http.StatusUnprocessableEntity},
		{"unknown end", f.gm, aliceFx(), func() url.Values {
			form := acBonusForm("a", 0)
			form.Set("ends_on", "nap")
			return form
		}(), http.StatusUnprocessableEntity},
		{"unknown target", f.gm, aliceFx(), func() url.Values {
			form := acBonusForm("a", 0)
			form.Set("mod_target", "flying")
			return form
		}(), http.StatusUnprocessableEntity},
		{"custom without a name", f.gm, aliceFx(), url.Values{"slug": {"custom"}}, http.StatusUnprocessableEntity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := f.do(t, http.MethodPost, tt.path, tt.cookie, tt.form, true)
			if w.Code != tt.want {
				t.Errorf("status %d, want %d: %s", w.Code, tt.want, w.Body)
			}
			if forbidden := strings.Contains(w.Body.String(), "not your character"); forbidden {
				t.Errorf("unexpected refusal: %s", w.Body)
			}
		})
	}

	// What an accepted form stores.
	form := acBonusForm("Bless", cassConc.ID)
	form.Set("ends_on", "short_rest")
	form.Set("note", "d4 to attacks")
	if w := f.effect(t, f.gm, f.bram.ID, form); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	var got sqlc.Condition
	for _, c := range f.conds(t, f.bram.ID) {
		if strings.HasPrefix(c.Note, "Bless") {
			got = c
		}
	}
	if got.Note != "Bless — d4 to attacks" || got.EndsOn != enum.EffectEndShortRest || got.HeldBy.Int64 != cassConc.ID || !strings.Contains(got.ModsJson, `"target":"ac"`) {
		t.Errorf("stored %+v", got)
	}
}

// TestAllyEffects checks a player may put an effect on an ally in their
// campaign only held by their own concentration, and who may remove it.
func TestAllyEffects(t *testing.T) {
	f := newEffectFix(t)
	ctx := t.Context()
	aliceConc := f.concentrating(t, f.characterID, 0)
	bramConc := f.concentrating(t, f.bram.ID, 0)
	far, err := f.q().CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Far", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	farChar, err := f.q().CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{CampaignID: far.ID, Name: "Far", System: "dnd5e", SheetJson: "{}"})
	if err != nil {
		t.Fatal(err)
	}
	cassTok, _, err := f.mgr.IssuePlayer(ctx, f.cass.ID)
	if err != nil {
		t.Fatal(err)
	}

	add := []struct {
		name string
		path string
		form url.Values
		want int
	}{
		{"ally, held by own concentration", route.Path(route.CharacterEffect, f.bram.ID), acBonusForm("a", aliceConc.ID), http.StatusOK},
		{"ally, no holder", route.Path(route.CharacterEffect, f.bram.ID), acBonusForm("a", 0), http.StatusUnprocessableEntity},
		{"ally, another's concentration", route.Path(route.CharacterEffect, f.bram.ID), acBonusForm("a", bramConc.ID), http.StatusUnprocessableEntity},
		{"another campaign's character", route.Path(route.CharacterEffect, farChar.ID), acBonusForm("a", aliceConc.ID), http.StatusUnprocessableEntity}, // an htmx refusal is a 422 toast
		{"non-PC combatant", route.Path(route.CombatantAction, f.mage.ID, "condition"), acBonusForm("a", aliceConc.ID), http.StatusForbidden},
	}
	for _, tt := range add {
		t.Run(tt.name, func(t *testing.T) {
			w := f.do(t, http.MethodPost, tt.path, f.alice, tt.form, true)
			if w.Code != tt.want {
				t.Errorf("status %d, want %d: %s", w.Code, tt.want, w.Body)
			}
		})
	}
	var applied []sqlc.Condition
	for _, c := range f.conds(t, f.bram.ID) {
		if c.Slug == "custom" {
			applied = append(applied, c)
		}
	}
	if len(applied) != 1 || applied[0].AppliedBy.Int64 != f.characterID || !applied[0].AppliedBy.Valid {
		t.Fatalf("stored %+v, want one applied by Alice", applied)
	}
	if len(f.conds(t, farChar.ID)) != 0 {
		t.Error("an effect landed on another campaign's character")
	}

	remove := []struct {
		name   string
		cookie string
		want   int
	}{
		{"third player", cassTok, http.StatusUnprocessableEntity}, // an htmx refusal is a 422 toast
		{"the applier", f.alice, http.StatusOK},
	}
	for _, tt := range remove {
		t.Run("remove, "+tt.name, func(t *testing.T) {
			w := f.do(t, http.MethodPost, route.Path(route.ConditionDelete, applied[0].ID), tt.cookie, nil, true)
			if w.Code != tt.want {
				t.Errorf("status %d, want %d: %s", w.Code, tt.want, w.Body)
			}
		})
	}
	// The target's player removes another one.
	if w := f.effect(t, f.alice, f.bram.ID, acBonusForm("b", aliceConc.ID)); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	var second sqlc.Condition
	for _, c := range f.conds(t, f.bram.ID) {
		if c.Note == "b" {
			second = c
		}
	}
	if w := f.do(t, http.MethodPost, route.Path(route.ConditionDelete, second.ID), f.bramTok, nil, true); w.Code != http.StatusOK {
		t.Errorf("target's player: %d %s", w.Code, w.Body)
	}

	// Dropping the applier's concentration ends the ally's effect.
	if w := f.effect(t, f.alice, f.bram.ID, acBonusForm("c", aliceConc.ID)); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	f.cmd(t, f.alice, f.characterID, "concentrate end")
	if got := notes(f.conds(t, f.bram.ID)); slices.Contains(got, "c") {
		t.Errorf("effect survived the applier's concentration: %v", got)
	}
}

// TestRestEndsEffects checks each rest removes the effects that end on it.
func TestRestEndsEffects(t *testing.T) {
	tests := []struct {
		rest string
		want []string
	}{
		{"rest short", []string{"keep", "long"}},
		{"rest long", []string{"keep"}},
	}
	for _, tt := range tests {
		t.Run(tt.rest, func(t *testing.T) {
			f := newEffectFix(t)
			for _, r := range []struct{ name, ends string }{{"keep", ""}, {"short", "short_rest"}, {"long", "long_rest"}} {
				form := url.Values{"slug": {"custom"}, "name": {r.name}, "ends_on": {r.ends}}
				if w := f.effect(t, f.alice, f.characterID, form); w.Code != http.StatusOK {
					t.Fatalf("%s: %d %s", r.name, w.Code, w.Body)
				}
			}
			f.cmd(t, f.alice, f.characterID, tt.rest)
			got := notes(f.conds(t, f.characterID))
			if strings.Join(got, ",") != strings.Join(tt.want, ",") {
				t.Errorf("after %s: %v, want %v", tt.rest, got, tt.want)
			}
		})
	}
}

// watchSheets opens a GM websocket on the sheets of characters and returns a
// function that waits until each of them was sent something.
func watchSheets(t *testing.T, f effectFix, characterIDs ...int64) func() {
	t.Helper()
	srv := httptest.NewServer(f.router)
	t.Cleanup(srv.Close)
	var keys []string
	for _, id := range characterIDs {
		keys = append(keys, views.LiveKey(views.LiveCharacter, id))
	}
	conn := dial(t, f.testEnv, srv, fmt.Sprintf("?campaign=%d", f.campaignID), "", keys...)
	readUntil(t, conn, "the hello", func(m wsMessage) bool { return m.Type == "patch" && len(m.Views) == 0 })
	return func() {
		t.Helper()
		seen := map[string]bool{}
		readUntil(t, conn, "every held character to be sent", func(m wsMessage) bool {
			for _, k := range keys {
				if _, ok := patchOf(m, k); ok {
					seen[k] = true
				}
			}
			return len(seen) == len(keys)
		})
	}
}

// TestConcentrationCascade checks that ending a caster's concentration, in
// any way, removes the effects it holds and refreshes their owners.
func TestConcentrationCascade(t *testing.T) {
	tests := []struct {
		name string
		// onCombatant puts the caster's concentration on the monster.
		onCombatant bool
		fire        func(t *testing.T, f effectFix)
	}{
		{"concentrate end", false, func(t *testing.T, f effectFix) { f.cmd(t, f.gm, f.cass.ID, "concentrate end") }},
		{"a new concentration replaces it", false, func(t *testing.T, f effectFix) { f.cmd(t, f.gm, f.cass.ID, "concentrate Bless") }},
		{"the caster's character is deleted", false, func(t *testing.T, f effectFix) {
			f.post200(t, route.Path(route.CharacterDelete, f.cass.ID))
		}},
		{"the caster's character is removed", false, func(t *testing.T, f effectFix) {
			f.post200(t, route.Path(route.CharacterRemove, f.cass.ID))
		}},
		{"the caster's combatant is deleted", true, func(t *testing.T, f effectFix) {
			f.post200(t, route.Path(route.CombatantAction, f.mage.ID, "delete"))
		}},
		{"the encounter is deleted", true, func(t *testing.T, f effectFix) {
			f.post200(t, route.Path(route.EncounterAction, f.enc.ID, "delete"))
		}},
		{"the caster drops to 0 HP", false, func(t *testing.T, f effectFix) {
			if err := f.app.updateCharacter(t.Context(), f.cass.ID, func(_ *sqlc.Queries, dc *dnd5e.Character) error {
				dc.HP, dc.MaxHP = 10, 10
				return nil
			}); err != nil {
				t.Fatal(err)
			}
			f.cmd(t, f.gm, f.cass.ID, "hp -99")
		}},
		{"the caster's concentration runs out", true, func(t *testing.T, f effectFix) {
			ctx := t.Context()
			if _, err := f.conn.ExecContext(ctx, "UPDATE conditions SET expires_round = 1 WHERE combatant_id = ? AND slug = ?", f.mage.ID, slugConcentration); err != nil {
				t.Fatal(err)
			}
			if err := f.q().UpdateEncounterState(ctx, sqlc.UpdateEncounterStateParams{ID: f.enc.ID, Status: enum.EncounterActive, Round: 1, ActiveCombatantID: nullID(f.mage.ID)}); err != nil {
				t.Fatal(err)
			}
			f.post200(t, route.Path(route.EncounterAction, f.enc.ID, "next")) // the mage's next turn, round 2
		}},
		{"the caster monster drops to 0 HP", true, func(t *testing.T, f effectFix) {
			if w := f.do(t, http.MethodPost, route.Path(route.CombatantAction, f.mage.ID, "hp"), f.gm, url.Values{"amount": {"99"}}, true); w.Code >= http.StatusBadRequest {
				t.Fatalf("%d %s", w.Code, w.Body)
			}
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newEffectFix(t)
			var holder sqlc.Condition
			if tt.onCombatant {
				holder = f.concentrating(t, 0, f.mage.ID)
			} else {
				holder = f.concentrating(t, f.cass.ID, 0)
			}
			targets := []int64{f.characterID, f.bram.ID}
			base := map[int64]int{}
			for _, id := range targets {
				base[id] = f.shown(t, id).ArmorClass()
				if w := f.effect(t, f.gm, id, acBonusForm(shieldOfFaith, holder.ID)); w.Code != http.StatusOK {
					t.Fatalf("effect: %d %s", w.Code, w.Body)
				}
				if got := f.shown(t, id).ArmorClass(); got != base[id]+2 {
					t.Fatalf("AC with the effect %d, want %d", got, base[id]+2)
				}
			}
			wait := watchSheets(t, f, targets...)
			tt.fire(t, f)
			for _, id := range targets {
				if rows := f.conds(t, id); len(rows) != 0 {
					t.Errorf("character %d keeps %+v", id, rows)
				}
				if got := f.shown(t, id).ArmorClass(); got != base[id] {
					t.Errorf("character %d AC %d after, want %d", id, got, base[id])
				}
			}
			wait()
		})
	}
}

// post200 posts as the GM and fails the test on anything but 200.
func (f effectFix) post200(t *testing.T, path string) {
	t.Helper()
	if w := f.do(t, http.MethodPost, path, f.gm, nil, true); w.Code >= http.StatusBadRequest {
		t.Fatalf("%s: %d %s", path, w.Code, w.Body)
	}
}

// TestIncapacitatingEndsConcentration checks the five conditions end the
// holder's concentration, for characters and combatants, and others do not.
func TestIncapacitatingEndsConcentration(t *testing.T) {
	slugs := []struct {
		slug string
		ends bool
	}{
		{"incapacitated", true}, {"paralyzed", true}, {"petrified", true}, {"stunned", true}, {"unconscious", true},
		{"prone", false}, {"poisoned", false},
	}
	for _, s := range slugs {
		for _, how := range []string{"character form", "character toggle", "combatant form"} {
			t.Run(s.slug+" "+how, func(t *testing.T) {
				f := newEffectFix(t)
				switch how {
				case "combatant form":
					f.concentrating(t, 0, f.mage.ID)
					if w := f.do(t, http.MethodPost, route.Path(route.CombatantAction, f.mage.ID, "condition"), f.gm, url.Values{"slug": {s.slug}}, true); w.Code != http.StatusOK {
						t.Fatalf("%d %s", w.Code, w.Body)
					}
				case "character form":
					f.concentrating(t, f.cass.ID, 0)
					if w := f.effect(t, f.gm, f.cass.ID, url.Values{"slug": {s.slug}}); w.Code != http.StatusOK {
						t.Fatalf("%d %s", w.Code, w.Body)
					}
				default:
					f.concentrating(t, f.cass.ID, 0)
					f.cmd(t, f.gm, f.cass.ID, "toggle "+s.slug)
				}
				// By slug: SQLite reuses the deleted row's id for the new condition.
				rows := f.conds(t, f.cass.ID)
				if how == "combatant form" {
					rows, _ = f.q().ListCombatantConditions(t.Context(), nullID(f.mage.ID))
				}
				kept := slices.ContainsFunc(rows, func(c sqlc.Condition) bool { return c.Slug == slugConcentration })
				if gone := !kept; gone != s.ends {
					t.Errorf("concentration gone = %v, want %v (rows %+v)", gone, s.ends, rows)
				}
				if !slices.ContainsFunc(rows, func(c sqlc.Condition) bool { return c.Slug == s.slug }) {
					t.Errorf("%s was not added: %+v", s.slug, rows)
				}
			})
		}
	}
}

// TestEffectNeverStored checks a command on a character with an active
// effect leaves the stored sheet without its modifiers.
func TestEffectNeverStored(t *testing.T) {
	f := newEffectFix(t)
	base := f.shown(t, f.characterID).ArmorClass()
	if w := f.effect(t, f.gm, f.characterID, acBonusForm(shieldOfFaith, 0)); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	for _, line := range []string{"hp -1", "temp 3", "rest long", "item add Rope"} {
		f.cmd(t, f.alice, f.characterID, line)
	}
	ch, err := f.q().GetCharacter(t.Context(), f.characterID)
	if err != nil {
		t.Fatal(err)
	}
	stored, err := dnd5e.ParseSheet(ch.SheetJson)
	if err != nil {
		t.Fatal(err)
	}
	if len(stored.Modifiers) != 0 || strings.Contains(ch.SheetJson, shieldOfFaith) || stored.ArmorClass() != base {
		t.Errorf("stored sheet gained the effect: %v %s", stored.Modifiers, ch.SheetJson)
	}
	if got := f.shown(t, f.characterID).ArmorClass(); got != base+2 {
		t.Errorf("shown AC %d, want %d", got, base+2)
	}
	// The builder, wizard and level-up read the stored build.
	row, _ := f.q().GetCharacter(t.Context(), f.characterID)
	sv, err := f.app.sheetView(t.Context(), perm.GM, row)
	if err != nil {
		t.Fatal(err)
	}
	if len(sv.C.Sheet.Modifiers) != 0 {
		t.Errorf("sheetView (stored) has modifiers %v", sv.C.Sheet.Modifiers)
	}
}

// TestEffectShownOnSheet checks effect modifiers reach AC and a save on the
// sheet view, name themselves in the breakdown, and go with the effect.
func TestEffectShownOnSheet(t *testing.T) {
	f := newEffectFix(t)
	sheet := f.shown(t, f.characterID)
	ac, save := sheet.ArmorClass(), sheet.SaveBonus(dnd5e.Wis)
	form := acBonusForm(shieldOfFaith, 0)
	form.Add("mod_target", "save.wis")
	form.Add("mod_value", "1")
	if w := f.effect(t, f.alice, f.characterID, form); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	sheet = f.shown(t, f.characterID)
	if sheet.ArmorClass() != ac+2 || sheet.SaveBonus(dnd5e.Wis) != save+1 {
		t.Errorf("AC %d (want %d), Wis save %d (want %d)", sheet.ArmorClass(), ac+2, sheet.SaveBonus(dnd5e.Wis), save+1)
	}
	named := false
	for _, p := range sheet.Explain(dnd5e.ModAC) {
		named = named || strings.Contains(p.Label, shieldOfFaith)
	}
	if !named {
		t.Errorf("the AC breakdown does not name the effect: %+v", sheet.Explain(dnd5e.ModAC))
	}
	if w := f.do(t, http.MethodGet, route.Path(route.Character, f.characterID), f.alice, nil, false); w.Code != http.StatusOK {
		t.Errorf("sheet page: %d", w.Code)
	}
	cond := f.conds(t, f.characterID)[0]
	if w := f.do(t, http.MethodPost, route.Path(route.ConditionDelete, cond.ID), f.alice, nil, true); w.Code != http.StatusOK {
		t.Fatalf("delete: %d %s", w.Code, w.Body)
	}
	sheet = f.shown(t, f.characterID)
	if sheet.ArmorClass() != ac || sheet.SaveBonus(dnd5e.Wis) != save {
		t.Errorf("after removal AC %d (want %d), Wis save %d (want %d)", sheet.ArmorClass(), ac, sheet.SaveBonus(dnd5e.Wis), save)
	}
}

// TestMalformedEffectSkipped checks a row with broken modifiers cannot
// take the sheet down.
func TestMalformedEffectSkipped(t *testing.T) {
	sheet := effectSheet(dnd5e.Sheet{}, []sqlc.Condition{{ID: 1, Slug: "custom", Note: "Bad", ModsJson: `{"not":"a list"}`}})
	if len(sheet.Modifiers) != 0 {
		t.Errorf("modifiers %v", sheet.Modifiers)
	}
}

// TestTrackerShowsEffects checks the tracker's PC numbers include effects,
// and what a player may learn about other combatants' effects.
func TestTrackerShowsEffects(t *testing.T) {
	f := newEffectFix(t)
	ctx := t.Context()
	var pcs [2]sqlc.Combatant
	for i, id := range []int64{f.characterID, f.bram.ID} {
		var err error
		if pcs[i], err = f.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: f.enc.ID, CharacterID: nullID(id), Name: "pc"}); err != nil {
			t.Fatal(err)
		}
	}
	find := func(ev views.EncounterView, id int64) views.CombatantView {
		for _, c := range ev.Combatants {
			if c.ID == id {
				return c
			}
		}
		t.Fatalf("combatant %d not in the tracker", id)
		return views.CombatantView{}
	}
	before, _ := f.app.encounterView(ctx, perm.GM, f.enc)
	ac := find(before, pcs[0].ID).AC
	conc := f.concentrating(t, f.cass.ID, 0)
	for _, id := range []int64{f.characterID, f.bram.ID} {
		if w := f.effect(t, f.gm, id, acBonusForm(shieldOfFaith, conc.ID)); w.Code != http.StatusOK {
			t.Fatalf("%d %s", w.Code, w.Body)
		}
	}
	if w := f.do(t, http.MethodPost, route.Path(route.CombatantAction, f.mage.ID, "condition"), f.gm, acBonusForm("Mage Armor", 0), true); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}

	gm, err := f.app.encounterView(ctx, perm.GM, f.enc)
	if err != nil {
		t.Fatal(err)
	}
	if got := find(gm, pcs[0].ID).AC; got != ac+2 {
		t.Errorf("tracker AC %d, want %d", got, ac+2)
	}
	player, err := f.app.encounterView(ctx, perm.Player(f.characterID), f.enc)
	if err != nil {
		t.Fatal(err)
	}
	// Whose effect data each viewer receives.
	tests := []struct {
		name     string
		view     views.EncounterView
		combat   int64
		wantMods bool
	}{
		{"GM, a monster", gm, f.mage.ID, true},
		{"GM, a character", gm, pcs[1].ID, true},
		{"player, a monster", player, f.mage.ID, false},
		{"player, another character", player, pcs[1].ID, false},
		{"player, own character", player, pcs[0].ID, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conds := find(tt.view, tt.combat).Conditions
			if len(conds) == 0 {
				t.Fatal("no conditions")
			}
			for _, c := range conds {
				hasMods := c.ModsJson != noMods
				hasHolder := c.HeldBy.Valid
				if hasMods != tt.wantMods {
					t.Errorf("mods_json %q, want data %v", c.ModsJson, tt.wantMods)
				}
				if hasHolder && !tt.wantMods {
					t.Errorf("held_by reached the player: %+v", c)
				}
			}
		})
	}
}

// TestEncounterEndClearsTimedEffects checks ending a fight removes round-timed
// effects, those held by a timed concentration included, and keeps the rest.
func TestEncounterEndClearsTimedEffects(t *testing.T) {
	f := newEffectFix(t)
	ctx := t.Context()
	if _, err := f.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: f.enc.ID, CharacterID: nullID(f.characterID), Name: "pc", Initiative: sql.NullInt64{Int64: 5, Valid: true}}); err != nil {
		t.Fatal(err)
	}
	// Rounds count only in a running fight.
	if w := f.effect(t, f.alice, f.characterID, url.Values{"slug": {"custom"}, "name": {"Early"}, "rounds": {"3"}}); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("rounds outside a fight: %d, want 422", w.Code)
	}
	if w := f.do(t, http.MethodPost, route.Path(route.EncounterAction, f.enc.ID, "start"), f.gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("start: %d %s", w.Code, w.Body)
	}
	if w := f.effect(t, f.alice, f.characterID, url.Values{"slug": {"custom"}, "name": {"Timed"}, "rounds": {"3"}}); w.Code != http.StatusOK {
		t.Fatalf("timed: %d %s", w.Code, w.Body)
	}
	if w := f.effect(t, f.alice, f.characterID, url.Values{"slug": {"custom"}, "name": {"Untimed"}}); w.Code != http.StatusOK {
		t.Fatalf("untimed: %d %s", w.Code, w.Body)
	}
	mageTimed := f.row(t, sqlc.CreateConditionParams{CombatantID: nullID(f.mage.ID), Slug: slugConcentration, Note: "Haste", ExpiresRound: sql.NullInt64{Int64: 4, Valid: true}})
	if w := f.effect(t, f.gm, f.bram.ID, acBonusForm(shieldOfFaith, mageTimed.ID)); w.Code != http.StatusOK {
		t.Fatalf("held: %d %s", w.Code, w.Body)
	}
	wait := watchSheets(t, f, f.bram.ID)
	if w := f.do(t, http.MethodPost, route.Path(route.EncounterAction, f.enc.ID, "end"), f.gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("end: %d %s", w.Code, w.Body)
	}
	if got := notes(f.conds(t, f.characterID)); strings.Join(got, ",") != "Untimed" {
		t.Errorf("Alice keeps %v, want only Untimed", got)
	}
	if rows := f.conds(t, f.bram.ID); len(rows) != 0 {
		t.Errorf("Bram keeps %+v", rows)
	}
	if rows, _ := f.q().ListCombatantConditions(ctx, nullID(f.mage.ID)); len(rows) != 0 {
		t.Errorf("the monster keeps %+v", rows)
	}
	wait()
}

func TestDurationRounds(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"1 minute", 10},
		{"Concentration, up to 1 minute", 10},
		{"Concentration, up to 10 minutes", 100},
		{"10 minutes", 100},
		{"1 round", 1},
		{"3 rounds", 3},
		{"Concentration, up to 2 rounds", 2},
		{"up to 1 minute", 10}, // 5etools keeps concentration as a flag
		{"Up to 10 minutes", 100},
		{"1 hour", 0},
		{"Concentration, up to 1 hour", 0},
		{"Instantaneous", 0},
		{"Until dispelled", 0},
		{"1 minute or until dispelled", 0},
		{"0 rounds", 0},
		{"", 0},
	}
	for _, tt := range tests {
		if got := durationRounds(tt.in); got != tt.want {
			t.Errorf("durationRounds(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

// TestSpellPrefill checks what a compendium spell fills into the form.
func TestSpellPrefill(t *testing.T) {
	e := newTestEnv(t)
	e.source(t, "test", "2024")
	spells := []struct {
		slug, name, data string
		wantNote         string
		wantRounds       int
		wantConc         bool
	}{
		{"shield-of-faith", "Shield of Faith", `{"level":1,"duration":"Concentration, up to 10 minutes","concentration":true,"desc":"x"}`, "Concentration, up to 10 minutes", 100, true},
		{"fire-bolt", "Fire Bolt", `{"level":0,"duration":"Instantaneous","desc":"x"}`, "Instantaneous", 0, false},
		{"mage-armor", "Mage Armor", `{"level":1,"duration":"8 hours","desc":"x"}`, "8 hours", 0, false},
	}
	for _, s := range spells {
		if err := e.q().UpsertCompendiumEntry(t.Context(), sqlc.UpsertCompendiumEntryParams{Kind: enum.KindSpell, Slug: s.slug, Name: s.name, SourceKey: "test", DataJson: s.data, ContentHash: s.slug}); err != nil {
			t.Fatal(err)
		}
	}
	e.allowEntrySources(t)
	for _, s := range spells {
		t.Run(s.slug, func(t *testing.T) {
			name, note, rounds, d, err := e.app.spellPrefill(t.Context(), s.slug, e.campaignID)
			if err != nil {
				t.Fatal(err)
			}
			if name != s.name || note != s.wantNote || rounds != s.wantRounds || d.Concentration != s.wantConc {
				t.Errorf("got %q %q %d %v", name, note, rounds, d.Concentration)
			}
		})
	}
	if _, _, _, _, err := e.app.spellPrefill(t.Context(), "nope", e.campaignID); err == nil {
		t.Error("an unknown spell must be refused")
	}
	// The picked spell fills a blank name.
	f := effectFix{testEnv: e}
	f.alice = e.token(t, enum.RolePlayer)
	if w := f.effect(t, f.alice, f.characterID, url.Values{"slug": {"concentration"}, "spell": {"shield-of-faith"}}); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	if got := notes(f.conds(t, f.characterID)); len(got) != 1 || got[0] != "Shield of Faith" {
		t.Errorf("notes %v", got)
	}
}

// TestEffectNoteMaxMatchesTag: the dice are added to the note only within
// effectNoteMax, so it must stay the Note field's validator cap.
func TestEffectNoteMaxMatchesTag(t *testing.T) {
	tagHas(t, effectForm{}, "Note", fmt.Sprintf("max=%d", effectNoteMax))
}

// TestEffectFields checks the form fields the spell picker loads: filled in
// from the spell, with rounds only in a running fight.
func TestEffectFields(t *testing.T) {
	f := newEffectFix(t)
	f.source(t, "test", "2024")
	if err := f.q().UpsertCompendiumEntry(t.Context(), sqlc.UpsertCompendiumEntryParams{Kind: enum.KindSpell, Slug: "shield-of-faith", Name: shieldOfFaith, SourceKey: "test", DataJson: `{"level":1,"duration":"Concentration, up to 1 minute","concentration":true,"desc":"x"}`, ContentHash: "sof"}); err != nil {
		t.Fatal(err)
	}
	if err := f.q().UpsertCompendiumEntry(t.Context(), sqlc.UpsertCompendiumEntryParams{Kind: enum.KindSpell, Slug: "magic-weapon", Name: "Magic Weapon", SourceKey: "test", DataJson: `{"level":2,"duration":"1 hour","desc":"a +1 bonus to attack rolls and damage rolls.","higher_level":"The bonus increases to +2 with a level 3–5 spell slot. The bonus increases to +3 with a level 6+ spell slot."}`, ContentHash: "mw"}); err != nil {
		t.Fatal(err)
	}
	f.allowEntrySources(t)
	get := func(cookie, query string) *httptest.ResponseRecorder {
		return f.do(t, http.MethodGet, route.Path(route.CampaignEffectSpell, f.campaignID)+"?"+query, cookie, nil, false)
	}
	for _, tt := range []struct {
		name, cookie, query string
		code                int
		has, lacks          []string
	}{
		{"character, no fight", f.alice, fmt.Sprintf("character=%d&spell=Shield+of+Faith", f.characterID), http.StatusOK,
			[]string{`value="Shield of Faith"`, "Concentration — pick the caster below", `value="Concentration, up to 1 minute"`}, []string{`name="rounds"`}},
		{"slot 4 raises the parsed bonus", f.alice, fmt.Sprintf("character=%d&spell=magic-weapon&slot=4", f.characterID), http.StatusOK,
			[]string{"From the spell text — check before adding", `name="slot"`, `value="attack"`, `value="damage"`, `name="mod_value" value="2"`}, []string{`name="mod_value" value="1"`}},
		{"nothing parsed, no slot", f.alice, fmt.Sprintf("character=%d&spell=Shield+of+Faith", f.characterID), http.StatusOK, nil, []string{"From the spell text", `name="slot"`}},
		{"blank", f.alice, fmt.Sprintf("character=%d", f.characterID), http.StatusOK, []string{`name="note"`}, []string{`name="rounds"`, "pick the caster"}},
		{"combatant in a running fight", f.gm, fmt.Sprintf("combatant=%d&spell=shield-of-faith", f.mage.ID), http.StatusOK,
			[]string{`value="Shield of Faith"`, `name="rounds"`, `value="10"`}, nil},
		{"player on a combatant", f.alice, fmt.Sprintf("combatant=%d", f.mage.ID), http.StatusForbidden, nil, nil},
		{"player on an ally", f.alice, fmt.Sprintf("character=%d", f.cass.ID), http.StatusOK, []string{"Concentrate on a spell first"}, []string{`name="target"`, `value="nobody"`}},
		{"player on own, with targets", f.alice, fmt.Sprintf("character=%d", f.characterID), http.StatusOK, []string{`name="target"`}, nil},
		{"no target", f.gm, "", http.StatusForbidden, nil, nil},
	} {
		if tt.name == "combatant in a running fight" {
			if err := f.q().UpdateEncounterState(t.Context(), sqlc.UpdateEncounterStateParams{ID: f.enc.ID, Status: enum.EncounterActive, Round: 1}); err != nil {
				t.Fatal(err)
			}
		}
		t.Run(tt.name, func(t *testing.T) {
			w := get(tt.cookie, tt.query)
			if w.Code != tt.code {
				t.Fatalf("status %d: %s", w.Code, w.Body)
			}
			for _, s := range tt.has {
				if !strings.Contains(w.Body.String(), s) {
					t.Errorf("fields lack %q: %s", s, w.Body)
				}
			}
			for _, s := range tt.lacks {
				if strings.Contains(w.Body.String(), s) {
					t.Errorf("fields have %q", s)
				}
			}
		})
	}
}

// TestEffectModRows checks the form starts with one stat change and adds
// one row per press, up to the server's 10, with no armor-state select.
func TestEffectModRows(t *testing.T) {
	f := newEffectFix(t)
	get := func(query string) string {
		t.Helper()
		w := f.do(t, http.MethodGet, fmt.Sprintf("%s?character=%d%s", route.Path(route.CampaignEffectSpell, f.campaignID), f.characterID, query), f.alice, nil, false)
		if w.Code != http.StatusOK {
			t.Fatalf("%s: status %d: %s", query, w.Code, w.Body)
		}
		return w.Body.String()
	}
	for _, tt := range []struct {
		name, query string
		addsNext    bool
	}{
		{"fields", "", true},
		{"second row", "&row=1", true},
		{"last row", "&row=9", false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			body := get(tt.query)
			if n := strings.Count(body, `name="mod_target"`); n != 1 {
				t.Errorf("%d target selects, want 1", n)
			}
			if strings.Contains(body, `name="mod_when"`) {
				t.Error("effects offer no armor-state select")
			}
			if got := strings.Contains(body, "Stat change</button>") || strings.Contains(body, "+ Stat change"); got != tt.addsNext {
				t.Errorf("add button shown %v, want %v", got, tt.addsNext)
			}
		})
	}
	if w := f.do(t, http.MethodGet, fmt.Sprintf("%s?character=%d&row=10", route.Path(route.CampaignEffectSpell, f.campaignID), f.characterID), f.alice, nil, false); w.Code == http.StatusOK {
		t.Error("an 11th row must be refused")
	}

	// A posted armor state is ignored: an effect's change always applies.
	form := acBonusForm("Haste", 0)
	form.Set("mod_when", dnd5e.WhenUnarmored)
	if w := f.effect(t, f.alice, f.characterID, form); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	if got := f.conds(t, f.characterID); len(got) != 1 || strings.Contains(got[0].ModsJson, "when") {
		t.Errorf("stored %+v", got)
	}
}

// TestEffectModsReadOnly checks the sheet offers to edit its own modifiers
// but not an effect's, which is shown with its source.
func TestEffectModsReadOnly(t *testing.T) {
	f := newEffectFix(t)
	if err := f.app.updateCharacter(t.Context(), f.characterID, func(_ *sqlc.Queries, dc *dnd5e.Character) error {
		dc.Sheet.Modifiers = append(dc.Sheet.Modifiers, dnd5e.CustomMod{Target: dnd5e.ModSpeed, Value: 5, Note: "Boots"})
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	if w := f.effect(t, f.gm, f.characterID, acBonusForm("Haste", 0)); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	w := f.do(t, http.MethodGet, route.Path(route.Character, f.characterID), f.alice, nil, false) // only the player edits the sheet
	body := w.Body.String()
	if !strings.Contains(body, `aria-label="Edit modifier Boots"`) {
		t.Error("the sheet's own modifier cannot be edited")
	}
	if !strings.Contains(body, "Haste") || strings.Contains(body, `aria-label="Edit modifier Haste"`) {
		t.Error("an effect's modifier must show, and have no edit control")
	}
}

// TestHeldByOptions checks which concentration rows a viewer may pick.
func TestHeldByOptions(t *testing.T) {
	f := newEffectFix(t)
	ctx := t.Context()
	aliceConc := f.concentrating(t, f.characterID, 0)
	cassConc := f.concentrating(t, f.cass.ID, 0)
	mageConc := f.concentrating(t, 0, f.mage.ID)
	if _, err := f.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: f.enc.ID, CharacterID: nullID(f.bram.ID), Name: "Bram"}); err != nil {
		t.Fatal(err)
	}
	ids := func(v perm.Viewer, character, combatant int64) []int64 {
		opts, err := f.app.heldByOptions(ctx, v, f.campaignID, character, combatant)
		if err != nil {
			t.Fatal(err)
		}
		var out []int64
		for _, o := range opts {
			out = append(out, o.ID)
		}
		return out
	}
	tests := []struct {
		name string
		got  []int64
		want []int64
	}{
		{"GM, a character in the fight", ids(perm.GM, f.bram.ID, 0), []int64{aliceConc.ID, cassConc.ID, mageConc.ID}},
		{"GM, a character in no fight", ids(perm.GM, f.cass.ID, 0), []int64{aliceConc.ID, cassConc.ID}},
		{"GM, a combatant", ids(perm.GM, 0, f.mage.ID), []int64{aliceConc.ID, cassConc.ID, mageConc.ID}},
		{"player", ids(perm.Player(f.characterID), f.characterID, 0), []int64{aliceConc.ID}},
		{"player without concentration", ids(perm.Player(f.bram.ID), f.bram.ID, 0), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if fmt.Sprint(tt.got) != fmt.Sprint(tt.want) {
				t.Errorf("options %v, want %v", tt.got, tt.want)
			}
		})
	}
	opts, _ := f.app.heldByOptions(ctx, perm.GM, f.campaignID, f.cass.ID, 0)
	if len(opts) == 0 || opts[1].Name != "Cass: Haste" {
		t.Errorf("label %+v", opts)
	}
}

// TestExportCarriesNoEffects checks neither a character's download nor the
// campaign export holds an effect or its modifiers.
func TestExportCarriesNoEffects(t *testing.T) {
	f := newEffectFix(t)
	conc := f.concentrating(t, f.cass.ID, 0)
	if w := f.effect(t, f.gm, f.characterID, acBonusForm(shieldOfFaith, conc.ID)); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	w := f.do(t, http.MethodGet, route.Path(route.CharacterDownload, f.characterID), f.gm, nil, false)
	if w.Code != http.StatusOK || strings.Contains(w.Body.String(), shieldOfFaith) {
		t.Errorf("download: %d, effect in file: %v", w.Code, strings.Contains(w.Body.String(), shieldOfFaith))
	}
	w = f.do(t, http.MethodGet, route.Path(route.CampaignExport, f.campaignID), f.gm, nil, false)
	if w.Code != http.StatusOK {
		t.Fatalf("export: %d", w.Code)
	}
	zr, err := zip.NewReader(bytes.NewReader(w.Body.Bytes()), int64(w.Body.Len()))
	if err != nil {
		t.Fatal(err)
	}
	for _, zf := range zr.File {
		rc, err := zf.Open()
		if err != nil {
			t.Fatal(err)
		}
		body, _ := io.ReadAll(rc)
		_ = rc.Close()
		if strings.Contains(string(body), shieldOfFaith) {
			t.Errorf("%s holds the effect", zf.Name)
		}
	}
}

// TestEffectPresets checks saving an effect as a preset with Add, listing,
// picking and deleting: private to the GM or to the player's character, and
// to its campaign.
func TestEffectPresets(t *testing.T) {
	f := newEffectFix(t)
	ctx := t.Context()
	far, err := f.q().CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Far", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	save := func(cookie string, characterID int64, name string) {
		t.Helper()
		form := acBonusForm(name, 0)
		form.Set("save_preset", "on")
		form.Set("note", "from a test")
		if w := f.effect(t, cookie, characterID, form); w.Code != http.StatusOK {
			t.Fatalf("save %q: %d %s", name, w.Code, w.Body)
		}
	}
	list := func(campaign, owner int64) []sqlc.EffectPreset {
		t.Helper()
		rows, err := f.q().ListEffectPresets(ctx, sqlc.ListEffectPresetsParams{CampaignID: campaign, Owner: nullID(owner)})
		if err != nil {
			t.Fatal(err)
		}
		return rows
	}
	query := func(cookie, q string) *httptest.ResponseRecorder {
		return f.do(t, http.MethodGet, fmt.Sprintf("%s?character=%d&%s", route.Path(route.CampaignEffectSpell, f.campaignID), f.characterID, q), cookie, nil, false)
	}

	save(f.gm, f.characterID, "Gm Aura")
	save(f.gm, f.characterID, "Gm Aura") // same name: overwritten, not duplicated
	save(f.alice, f.characterID, "Alice Ward")
	save(f.bramTok, f.bram.ID, "Bram Ward")
	if got := list(f.campaignID, 0); len(got) != 1 || got[0].Name != "Gm Aura" || got[0].ModsJson == noMods {
		t.Fatalf("GM presets: %+v", got)
	}
	if got := list(f.campaignID, f.characterID); len(got) != 1 || got[0].Name != "Alice Ward" {
		t.Fatalf("Alice's presets: %+v", got)
	}
	gmPreset, alicePreset, bramPreset := list(f.campaignID, 0)[0], list(f.campaignID, f.characterID)[0], list(f.campaignID, f.bram.ID)[0]
	farPreset, err := f.q().UpsertEffectPreset(ctx, sqlc.UpsertEffectPresetParams{CampaignID: far.ID, Name: "Far Ward", Slug: "custom", ModsJson: noMods})
	if err != nil {
		t.Fatal(err)
	}
	if w := f.effect(t, f.gm, f.characterID, url.Values{"slug": {"custom"}, "save_preset": {"on"}}); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("a preset without a name: %d", w.Code)
	}

	for _, tt := range []struct {
		name, cookie, q string
		code            int
		has, lacks      []string
	}{
		{"GM lists only the GM's", f.gm, "presets=1", http.StatusOK, []string{"Gm Aura"}, []string{"Alice Ward", "Bram Ward"}},
		{"player lists only their own", f.alice, "presets=1", http.StatusOK, []string{"Alice Ward"}, []string{"Gm Aura", "Bram Ward"}},
		{"owner picks", f.alice, fmt.Sprintf("preset=%d", alicePreset.ID), http.StatusOK, []string{"Alice Ward", "from a test"}, nil},
		{"GM picks the GM's", f.gm, fmt.Sprintf("preset=%d", gmPreset.ID), http.StatusOK, []string{"Gm Aura"}, nil},
		{"other player picks", f.alice, fmt.Sprintf("preset=%d", bramPreset.ID), http.StatusForbidden, nil, []string{"Bram Ward"}},
		{"GM picks a player's", f.gm, fmt.Sprintf("preset=%d", alicePreset.ID), http.StatusForbidden, nil, []string{"Alice Ward"}},
		{"player picks the GM's", f.alice, fmt.Sprintf("preset=%d", gmPreset.ID), http.StatusForbidden, nil, []string{"Gm Aura"}},
		{"other campaign's", f.gm, fmt.Sprintf("preset=%d", farPreset.ID), http.StatusForbidden, nil, []string{"Far Ward"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			w := query(tt.cookie, tt.q)
			if w.Code != tt.code {
				t.Fatalf("status %d, want %d: %s", w.Code, tt.code, w.Body)
			}
			for _, s := range tt.has {
				if !strings.Contains(w.Body.String(), s) {
					t.Errorf("lacks %q", s)
				}
			}
			for _, s := range tt.lacks {
				if strings.Contains(w.Body.String(), s) {
					t.Errorf("shows %q", s)
				}
			}
		})
	}

	for _, tt := range []struct {
		name     string
		cookie   string
		campaign int64
		id       int64
		code     int
	}{
		{"other player deletes", f.alice, f.campaignID, bramPreset.ID, http.StatusForbidden},
		{"GM deletes a player's", f.gm, f.campaignID, alicePreset.ID, http.StatusForbidden},
		{"another campaign's, through this one", f.gm, f.campaignID, farPreset.ID, http.StatusNotFound},
		{"GM deletes the GM's in another campaign", f.gm, far.ID, farPreset.ID, http.StatusOK},
		{"owner deletes", f.alice, f.campaignID, alicePreset.ID, http.StatusOK},
		{"missing", f.gm, f.campaignID, 999999, http.StatusNotFound},
	} {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{"preset": {fmt.Sprint(tt.id)}}
			if got := f.do(t, http.MethodPost, route.Path(route.CampaignEffectPresetsDelete, tt.campaign), tt.cookie, form, false).Code; got != tt.code { // plain, as htmx answers every refusal 422
				t.Errorf("status %d, want %d", got, tt.code)
			}
		})
	}
	if got := list(f.campaignID, f.characterID); len(got) != 0 {
		t.Errorf("Alice's preset is still there: %+v", got)
	}
	if got := list(f.campaignID, f.bram.ID); len(got) != 1 {
		t.Errorf("Bram's preset was removed: %+v", got)
	}
}
