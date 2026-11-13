package web

// The one door from system-neutral web code (combat, sessions, campaigns,
// character files, the compendium) into D&D 5e rules: plain-typed functions,
// so neutral files never import dnd5e (depguard in .golangci.yml checks it).
// Everything else 5e lives in sheet5e_*.go.

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/export"
)

// The game system and edition new campaigns and characters are created in.
const (
	gameSystem  = dnd5e.System
	gameEdition = dnd5e.Edition
)

// newSheetJSON is a new character's empty sheet.
func newSheetJSON() (string, error) {
	s := dnd5e.NewSheet()
	s.EffectsVersion = dnd5e.EffectsVersion // nothing to catch up; NewSheet stays 0 for old rows
	return s.JSON()
}

// sheetArmorClass is the armor class a stored sheet gives.
func sheetArmorClass(sheetJSON string) (int, error) {
	s, err := dnd5e.ParseSheet(sheetJSON)
	return s.ArmorClass(), err
}

// normalizeCharacterFile checks a character file's sheet and the whole
// character as a command would leave it, and stores the sheet as parsed, so
// the row holds exactly what the app understands.
func normalizeCharacterFile(cf *export.CharacterFile) error {
	sheet, err := dnd5e.ParseSheet(string(cf.Sheet))
	if err != nil {
		return err
	}
	ch := dnd5e.Character{Name: cf.Name, HP: int(cf.HP), MaxHP: int(cf.MaxHP), TempHP: int(cf.TempHP), Sheet: sheet}
	if err := ch.Validate(); err != nil {
		return err
	}
	cf.Sheet, err = json.Marshal(sheet)
	return err
}

// runLine runs one command line on ch, as if typed on the sheet; a line
// that does not parse is the user's error.
func (a *App) runLine(ctx context.Context, ch sqlc.Character, line string, raw ...string) error {
	cmd, err := dnd5e.Parse(line, raw...)
	if err != nil {
		return userError{err}
	}
	return a.runCommand(ctx, ch, cmd)
}

// pcSheet is a PC row's live sheet; ok is false for other combatants.
func pcSheet(r sqlc.ListCombatantsRow) (dnd5e.Sheet, bool) {
	if !r.SheetJson.Valid {
		return dnd5e.Sheet{}, false
	}
	s, err := dnd5e.ParseSheet(r.SheetJson.String)
	return s, err == nil
}

// initiativeMod is the bonus a combatant adds to its initiative roll: a PC's
// from the sheet, a monster's (or a statted NPC's) from its statblock.
func initiativeMod(r sqlc.ListCombatantsRow) int {
	if s, ok := pcSheet(r); ok {
		return s.Initiative()
	}
	if r.MonsterData.Valid {
		if d, err := compendium.Decode[compendium.MonsterData](enum.KindMonster, r.MonsterData.String); err == nil {
			return d.Initiative
		}
	}
	return 0
}

// combatantDefenses are how a combatant takes damage: a PC's sheet or a
// statblock's free text, all of it resisted while Petrified.
func combatantDefenses(r sqlc.ListCombatantsRow, conds []sqlc.Condition) dnd5e.Defenses {
	var d dnd5e.Defenses
	if s, ok := pcSheet(r); ok {
		d = s.Defenses()
	} else if r.MonsterData.Valid {
		if md, err := compendium.Decode[compendium.MonsterData](enum.KindMonster, r.MonsterData.String); err == nil {
			d.List = slices.Concat(dnd5e.ParseDefenses(dnd5e.DefResist, md.Resistance), dnd5e.ParseDefenses(dnd5e.DefImmune, md.Immunity), dnd5e.ParseDefenses(dnd5e.DefVulnerable, md.Vulnerability))
		}
	}
	d.ResistAll = petrified(conds)
	return d
}

func petrified(conds []sqlc.Condition) bool {
	return slices.ContainsFunc(conds, func(c sqlc.Condition) bool { return c.Slug == dnd5e.Petrified })
}

// outOfFight reports a combatant whose turn is skipped: a monster or NPC at
// 0 HP, or a PC who failed three death saves (a PC at 0 HP still rolls them).
func outOfFight(r sqlc.ListCombatantsRow) bool {
	if s, ok := pcSheet(r); ok {
		return s.DeathSaves.Failures >= dnd5e.MaxDeathSaves
	}
	return !r.CharacterID.Valid && r.Hp <= 0 && r.MaxHp > 0
}

// damageChoice reads the HP dialog's damage type (views.damageTypeItems):
// "fire", or "bludgeoning@3" when conditional defense 3 holds; empty is
// untyped.
func damageChoice(v string, d dnd5e.Defenses) (kind string, when int, err error) {
	kind, at, conditional := strings.Cut(v, "@")
	if kind != "" && !slices.Contains(dnd5e.DamageTypes, kind) {
		return "", 0, userError{fmt.Errorf("unknown damage type %q", kind)}
	}
	if !conditional {
		return kind, -1, nil
	}
	when, err = strconv.Atoi(at)
	if err != nil || when < 0 || when >= len(d.List) || d.List[when].When == "" || !slices.Contains(d.List[when].Types, kind) {
		return "", 0, userError{errors.New("that defense changed; pick the damage type again")}
	}
	return kind, when, nil
}

// hpChange changes a creature's hit points by the rules (temporary HP
// first, never below 0, and so on).
type hpChange func(*dnd5e.Character)

func healHP(n int) hpChange { return func(m *dnd5e.Character) { m.ApplyHP(dnd5e.Heal(n)) } }

// damageHP is n damage of kind after d (when as damageChoice reads it).
func damageHP(n int, d dnd5e.Defenses, kind string, when int) hpChange {
	return func(m *dnd5e.Character) { m.ApplyHP(dnd5e.Damage(d.Taken(n, kind, when))) }
}

// setHP sets the given values and leaves nil ones alone.
func setHP(hp, maxHP, temp *int) hpChange {
	return func(dc *dnd5e.Character) {
		if maxHP != nil {
			if dc.MaxHP == 0 && hp == nil { // no stats yet: start at full
				dc.HP = *maxHP
			}
			dc.MaxHP = *maxHP
			dc.HP = min(dc.HP, dc.MaxHP)
		}
		if hp != nil {
			dc.ApplyHP(dnd5e.Set(*hp))
		}
		if temp != nil {
			dc.TempHP = *temp
		}
	}
}

// changeCharacterHP applies fn to a character in one transaction.
func (a *App) changeCharacterHP(ctx context.Context, id int64, fn hpChange) error {
	return a.updateCharacter(ctx, id, func(_ *sqlc.Queries, dc *dnd5e.Character) error {
		fn(dc)
		return nil
	})
}

// updateMonster reads, changes and writes a monster's HP in one
// transaction, so changes sent from two devices at once both land.
func (a *App) updateMonster(ctx context.Context, id int64, fn hpChange) error {
	return a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		cb, err := q.GetCombatant(ctx, id)
		if err != nil {
			return err
		}
		m := dnd5e.Character{HP: int(cb.Hp), MaxHP: int(cb.MaxHp), TempHP: int(cb.TempHp)}
		fn(&m)
		return q.SetCombatantHP(ctx, sqlc.SetCombatantHPParams{ID: id, Hp: int64(m.HP), MaxHp: int64(m.MaxHP), TempHp: int64(m.TempHP)})
	})
}

// addCondition puts the condition slug on o: exhaustion gains a level, and
// a custom effect needs its note.
func (a *App) addCondition(ctx context.Context, o condOwner, slug, note string, expires sql.NullInt64) error {
	if !dnd5e.IsConditionSlug(slug) {
		return userError{fmt.Errorf("unknown condition %q", slug)}
	}
	if slug == dnd5e.TrackCustom && note == "" {
		return userError{errors.New("describe the custom effect")}
	}
	op := dnd5e.CondOp{Op: dnd5e.CondAdd, Slug: slug, Note: note}
	if slug == dnd5e.Exhaustion {
		op = dnd5e.CondOp{Op: dnd5e.CondExhaustion, Slug: slug, Delta: 1}
	}
	return a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		return applyCondOps(ctx, q, o, []dnd5e.CondOp{op}, expires)
	})
}
