package web

// The one door from system-neutral web code (combat, sessions, campaigns,
// character files, the compendium) into D&D 5e rules: plain-typed functions,
// so neutral files never import dnd5e (depguard in .golangci.yml checks it).
// Everything else 5e lives in sheet5e_*.go.

import (
	"bytes"
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
	"gm-session-manager/internal/web/views"
)

// The game system and edition new campaigns and characters are created in.
const (
	gameSystem  = dnd5e.System
	gameEdition = dnd5e.Edition
)

// appendedSourcePriority is a new source's priority: below every ordered source, by group rank.
func appendedSourcePriority(ctx context.Context, q *sqlc.Queries, key, edition string) (int64, error) {
	base, err := q.MaxSourcePriority(ctx)
	return base + compendium.DefaultRank(key, edition), err
}

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

// pcSheet is a PC row's live sheet, nil for other combatants. Parse it once
// per row and pass it to the helpers below.
func pcSheet(r sqlc.ListCombatantsRow) *dnd5e.Sheet {
	if !r.SheetJson.Valid {
		return nil
	}
	s, err := dnd5e.ParseSheet(r.SheetJson.String)
	if err != nil {
		return nil
	}
	return &s
}

// slugConcentration is the condition slug effects hang off (held_by).
const slugConcentration = dnd5e.TrackConcentration

// rowConditions are the conditions of a combatant row in conds (an
// encounter's): a PC's are its character's.
func rowConditions(r sqlc.ListCombatantsRow, conds []sqlc.Condition) []sqlc.Condition {
	return slices.DeleteFunc(slices.Clone(conds), func(c sqlc.Condition) bool {
		if r.CharacterID.Valid {
			return c.CharacterID != r.CharacterID
		}
		return c.CombatantID.Int64 != r.ID
	})
}

// pcSheetWith is pcSheet with the PC's active effects added, for the
// numbers the tracker shows or rolls. conds are the row's own (rowConditions).
func pcSheetWith(r sqlc.ListCombatantsRow, conds []sqlc.Condition) *dnd5e.Sheet {
	s := pcSheet(r)
	if s == nil {
		return nil
	}
	shown := effectSheet(*s, conds)
	return &shown
}

// combatantArmorClass is a PC's AC from the sheet (the stored one is from
// when they joined), else the row's.
func combatantArmorClass(r sqlc.ListCombatantsRow, s *dnd5e.Sheet) int64 {
	if s != nil {
		return int64(s.ArmorClass())
	}
	return r.Ac
}

// initiativeMod is the bonus a combatant adds to its initiative roll: a PC's
// from the sheet, a monster's (or a statted NPC's) from its statblock.
func initiativeMod(r sqlc.ListCombatantsRow, s *dnd5e.Sheet) int {
	if s != nil {
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
func combatantDefenses(r sqlc.ListCombatantsRow, s *dnd5e.Sheet, conds []sqlc.Condition) dnd5e.Defenses {
	var d dnd5e.Defenses
	if s != nil {
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
func outOfFight(r sqlc.ListCombatantsRow, s *dnd5e.Sheet) bool {
	if s != nil {
		return s.Dead()
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

// damageArgs are the arguments after "hp -n" for a damage type (blank or
// " " is untyped) and a critical hit; none for healing and temp HP, which a
// crit does not touch.
func damageArgs(damage bool, kind string, crit bool) []string {
	var args []string
	if kind = strings.TrimSpace(kind); damage && kind != "" {
		args = append(args, kind)
	}
	if damage && crit {
		args = append(args, "crit")
	}
	return args
}

// hpChange changes a creature's hit points by the rules (temporary HP
// first, never below 0, and so on).
type hpChange func(*dnd5e.Character) error

func healHP(n int) hpChange {
	return func(m *dnd5e.Character) error { return m.ApplyHP(dnd5e.Heal(n)) }
}

// damageHP is n damage of kind after d (when as damageChoice reads it).
// Monsters have no death saves, so a crit changes nothing for them.
func damageHP(n int, d dnd5e.Defenses, kind string, when int) hpChange {
	return func(m *dnd5e.Character) error { return m.ApplyHP(dnd5e.Damage(d.Taken(n, kind, when))) }
}

// setHP sets the given values and leaves nil ones alone.
func setHP(hp, maxHP, temp *int) hpChange {
	return func(dc *dnd5e.Character) error {
		if maxHP != nil {
			if dc.MaxHP == 0 && hp == nil { // no stats yet: start at full
				dc.HP = *maxHP
			}
			dc.MaxHP = *maxHP
			dc.HP = min(dc.HP, dc.MaxHP)
		}
		if hp != nil {
			if err := dc.ApplyHP(dnd5e.Set(*hp)); err != nil {
				return err
			}
		}
		if temp != nil {
			if dc.Sheet.Dead() {
				return dnd5e.ErrDead
			}
			dc.TempHP = *temp
		}
		return nil
	}
}

// changeCharacterHP applies fn to a character in one transaction; what fn
// refuses is the user's error.
func (a *App) changeCharacterHP(ctx context.Context, id int64, fn hpChange) error {
	return a.updateCharacter(ctx, id, func(_ *sqlc.Queries, dc *dnd5e.Character) error {
		if err := fn(dc); err != nil {
			return userError{err}
		}
		return nil
	})
}

// updateMonster reads, changes and writes a monster's HP in one
// transaction, so changes sent from two devices at once both land.
func (a *App) updateMonster(ctx context.Context, id int64, fn hpChange) error {
	return a.Store.InTx(ctx, func(q *sqlc.Queries) error { return monsterHP(ctx, q, id, fn) })
}

// monsterHP is updateMonster inside the caller's transaction.
func monsterHP(ctx context.Context, q *sqlc.Queries, id int64, fn hpChange) error {
	cb, err := q.GetCombatant(ctx, id)
	if err != nil {
		return err
	}
	m := dnd5e.Character{HP: int(cb.Hp), MaxHP: int(cb.MaxHp), TempHP: int(cb.TempHp)}
	if err := fn(&m); err != nil { // m has an empty sheet, so it is never dead
		return userError{err}
	}
	if cb.Hp > 0 && m.HP <= 0 { // a downed caster stops concentrating
		if err := endConcentration(ctx, q, condOwner{combatantID: id}); err != nil {
			return err
		}
	}
	return q.SetCombatantHP(ctx, sqlc.SetCombatantHPParams{ID: id, Hp: int64(m.HP), MaxHp: int64(m.MaxHP), TempHp: int64(m.TempHP)})
}

// monsterStats fills name, AC and HP from a compendium monster.
func (a *App) monsterStats(ctx context.Context, entryID int64, p *sqlc.CreateCombatantParams) error {
	row, err := a.Store.GetCompendiumEntry(ctx, entryID)
	if err != nil {
		return err
	}
	if row.Kind != enum.KindMonster {
		return sql.ErrNoRows
	}
	m, err := compendium.Decode[compendium.MonsterData](enum.KindMonster, row.DataJson)
	if err != nil {
		return err
	}
	p.MonsterRef, p.Name, p.Ac, p.Hp, p.MaxHp = nullID(row.ID), row.Name, int64(m.AC), int64(m.HP), int64(m.HP)
	return nil
}

// addHomebrewToSheet puts a new homebrew spell or item on the character.
func (a *App) addHomebrewToSheet(ctx context.Context, ch sqlc.Character, id int64) error {
	e, err := a.Store.GetCompendiumEntry(ctx, id)
	if err != nil {
		return err
	}
	line := fmt.Sprintf("item add %s", strconv.Quote(e.Name))
	if e.Kind == enum.KindSpell {
		line = "spell add " + e.Slug
	}
	return a.runLine(ctx, ch, line)
}

// homeBuildSources are the keys the sheet's build (classes, species, background,
// feats) pins.
func homeBuildSources(sheetJSON string, keys []string) ([]string, error) {
	s, err := dnd5e.ParseSheet(sheetJSON)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, k := range keys {
		if s.BuildPinsSource(k) {
			out = append(out, k)
		}
	}
	return out, nil
}

// errNewerSheet means a take-home sheet holds something this app does not know.
var errNewerSheet = errors.New("the sheet comes from a newer app")

// checkSheetFile refuses a sheet a newer app wrote: an unknown field (decoded
// strictly) or a higher EffectsVersion. Older sheets pass and UpgradeSheets
// catches them up. shortcut: new ProseMirror node types in notes are not
// detected, when a note format changes bump the zip format instead.
func checkSheetFile(raw json.RawMessage) error {
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	var s dnd5e.Sheet
	if err := dec.Decode(&s); err != nil {
		if strings.HasPrefix(err.Error(), "json: unknown field") {
			return errNewerSheet
		}
		return fmt.Errorf("sheet: %w", err)
	}
	if s.EffectsVersion > dnd5e.EffectsVersion {
		return errNewerSheet
	}
	return nil
}

// checkFileEntries checks the compendium entries of a bring-back: homebrew only
// of the kinds players write, and every entry's data readable.
func checkFileEntries(file *export.TakeHomeFile) error {
	for _, e := range file.Homebrew {
		if !slices.Contains(compendium.PlayerKinds, e.Kind) {
			return userf("the homebrew %q is a kind players cannot write", e.Name)
		}
	}
	for _, e := range slices.Concat(file.Homebrew, file.Entries) {
		if _, err := compendium.DecodeData(e.Kind, string(e.Data)); err != nil {
			return userf("the entry %q is not valid: %w", e.Name, err)
		}
	}
	return nil
}

// charState is a character as the bring-back review merges it: the columns and
// the sheet's JSON, so neutral code holds no 5e types.
type charState struct {
	Name              string
	HP, MaxHP, TempHP int64
	Sheet             string
}

func (c charState) parse() (dnd5e.Character, error) {
	s, err := dnd5e.ParseSheet(c.Sheet)
	return dnd5e.Character{Name: c.Name, HP: int(c.HP), MaxHP: int(c.MaxHP), TempHP: int(c.TempHP), Sheet: s}, err
}

func stateOf(c dnd5e.Character) (charState, error) {
	raw, err := c.Sheet.JSON()
	return charState{Name: c.Name, HP: int64(c.HP), MaxHP: int64(c.MaxHP), TempHP: int64(c.TempHP), Sheet: raw}, err
}

// Where the merged build comes from (dnd5e's sides, or the fallback).
const (
	buildFromHome     = dnd5e.SideHome
	buildFromTable    = dnd5e.SideTable
	buildFromFallback = "fallback"
)

// reviewSource is a source added at home that the campaign does not allow yet.
type reviewSource struct{ Key, Label string }

// bringBackIn is what the review merges: the character at take-out, at the
// table now and at home now, the fallback build if the file has one, and what
// the GM picked (a side per conflicting group, allow or decline per source).
type bringBackIn struct {
	Base, Table, Home charState
	Fallback          *charState
	Sources           []reviewSource
	// Form reads the page's fields (views.PickParam, views.SourceParam).
	Form func(key string) string
}

// bringBackPlan is what Accept would do, and how the page shows it.
type bringBackPlan struct {
	Merged  charState
	Groups  []views.HomeFileGroup
	Sources []views.HomeFileSource
	Removed []string
	// BuildFrom is buildFromHome, buildFromTable or buildFromFallback: whose
	// level-ups the character keeps.
	BuildFrom          string
	LevelFrom, LevelTo int
	// Block is why Accept must wait (a source not chosen, a build that needs a
	// declined source); Error is a merge that is not a valid character.
	Block, Error string
}

// planBringBack merges the three characters and applies the source choices.
// Undecided sources count as allowed, so the page can show a state before
// every choice is made.
func planBringBack(in bringBackIn) (bringBackPlan, error) {
	var p bringBackPlan
	base, err := in.Base.parse()
	if err != nil {
		return p, err
	}
	table, err := in.Table.parse()
	if err != nil {
		return p, err
	}
	home, err := in.Home.parse()
	if err != nil {
		return p, err
	}
	var fb *dnd5e.Character
	if in.Fallback != nil {
		c, err := in.Fallback.parse()
		if err != nil {
			return p, err
		}
		fb = &c
	}
	p.LevelFrom = base.Sheet.Level()
	picks := map[string]string{}
	for _, g := range dnd5e.MergeGroups {
		picks[g] = in.Form(views.PickParam(g))
	}
	merged, diffs, err := dnd5e.Merge(base, table, home, picks)
	p.BuildFrom = buildFromTable
	for _, d := range diffs {
		p.Groups = append(p.Groups, groupView(d))
		if d.Group == dnd5e.GroupBuild {
			p.BuildFrom = d.Side
		}
	}
	if err != nil {
		p.Error = "The merged character is not valid: " + err.Error()
		return p, nil
	}

	var declined []string
	// The fallback stands in only if it uses no declined source, this one included.
	fallbackFree := func(key string) bool {
		return fb != nil && !fb.Sheet.BuildPinsSource(key) && !slices.ContainsFunc(in.Sources, func(o reviewSource) bool {
			return in.Form(views.SourceParam(o.Key)) == views.SourceDecline && fb.Sheet.BuildPinsSource(o.Key)
		})
	}
	mergedJSON, err := merged.Sheet.JSON()
	if err != nil {
		return p, err
	}
	for _, rs := range in.Sources {
		choice := in.Form(views.SourceParam(rs.Key))
		if choice != views.SourceAllow && choice != views.SourceDecline {
			choice = ""
		}
		sv := views.HomeFileSource{Key: rs.Key, Label: rs.Label, Choice: choice, Uses: "Not used"}
		if uses := home.Sheet.SourceUses(rs.Key); len(uses) > 0 {
			sv.Uses = "Used: " + strings.Join(uses, ", ")
		}
		ms, err := dnd5e.ParseSheet(mergedJSON) // a copy: DropSource edits in place
		if err != nil {
			return p, err
		}
		switch {
		case ms.BuildPinsSource(rs.Key) && fallbackFree(rs.Key):
			sv.Effect = "Declining: the Build comes from the fallback (" + buildSummary(fb.Sheet) + ")"
		case ms.BuildPinsSource(rs.Key):
			sv.Effect = "Declining blocks Accept"
		default:
			if gone := ms.DropSource(rs.Key); len(gone) > 0 {
				sv.Effect = "Declining removes: " + strings.Join(gone, ", ")
			}
		}
		switch choice {
		case views.SourceDecline:
			declined = append(declined, rs.Key)
		case "":
			p.Block = "Choose Allow or Decline for each source."
		}
		p.Sources = append(p.Sources, sv)
	}

	var pinned []string
	for _, k := range declined {
		if merged.Sheet.BuildPinsSource(k) {
			pinned = append(pinned, k)
		}
	}
	if len(pinned) > 0 {
		if fb != nil && !slices.ContainsFunc(declined, fb.Sheet.BuildPinsSource) {
			if merged, err = dnd5e.TakeBuild(merged, *fb); err != nil {
				p.Error = "The fallback build is not valid: " + err.Error()
				return p, nil
			}
			p.BuildFrom = buildFromFallback
			for i, g := range p.Groups {
				if g.Key == dnd5e.GroupBuild {
					p.Groups[i].Note = "A declined source is in the build, so it comes from the fallback."
				}
			}
		} else if p.Block == "" {
			p.Block = fmt.Sprintf("%s's build uses %s and has no fallback without it. Allow it, or reject and ask the player to revert that level-up at home.", merged.Name, strings.Join(pinned, ", "))
		}
	}
	for _, k := range declined {
		for _, gone := range merged.Sheet.DropSource(k) {
			if !slices.Contains(p.Removed, gone) {
				p.Removed = append(p.Removed, gone)
			}
		}
	}
	if err := merged.Validate(); err != nil {
		p.Error = "The merged character is not valid: " + err.Error()
		return p, nil
	}
	p.LevelTo = merged.Sheet.Level()
	p.Merged, err = stateOf(merged)
	return p, err
}

// groupView is a merge group as the review shows it.
func groupView(d dnd5e.GroupDiff) views.HomeFileGroup {
	g := views.HomeFileGroup{Key: d.Group, Label: dnd5e.GroupLabel(d.Group)}
	switch d.State {
	case dnd5e.SideHome:
		g.Badge, g.Lines = titleHome, d.Home
	case dnd5e.SideTable:
		g.Badge, g.Lines = titleTable, d.Table
	case dnd5e.MergeConflict:
		g.Badge = "Both"
		g.Columns = []views.HomeFileColumn{
			{Side: dnd5e.SideHome, Title: titleHome, Lines: d.Home, Checked: d.Side == dnd5e.SideHome},
			{Side: dnd5e.SideTable, Title: titleTable, Lines: d.Table, Checked: d.Side == dnd5e.SideTable},
		}
	}
	return g
}

// buildSummary is "level 5, Wizard 5" for a sheet.
func buildSummary(s dnd5e.Sheet) string {
	parts := []string{fmt.Sprintf("level %d", s.Level())}
	for _, c := range s.Classes {
		parts = append(parts, fmt.Sprintf("%s %d", c.Name, c.Level))
	}
	return strings.Join(parts, ", ")
}

// writeMerged stores the merged character in the caller's transaction, through
// the same path as a command, so dropping to 0 HP still ends concentration.
func writeMerged(ctx context.Context, q *sqlc.Queries, id int64, m charState) error {
	return changeCharacterTx(ctx, q, id, func(_ *sqlc.Queries, dc *dnd5e.Character) error {
		c, err := m.parse()
		*dc = c
		return err
	})
}

// renameSpells applies homebrew spell slug renames (old to new) to a stored sheet.
func renameSpells(sheet string, renames map[string]string) (string, error) {
	if len(renames) == 0 {
		return sheet, nil
	}
	s, err := dnd5e.ParseSheet(sheet)
	if err != nil {
		return "", err
	}
	s.RenameSpells(renames)
	return s.JSON()
}

// sheetLevel is the character level a stored sheet has.
func sheetLevel(sheet string) (int, error) {
	s, err := dnd5e.ParseSheet(sheet)
	return s.Level(), err
}

// sheetSourceKeys is every book a sheet pins (not homebrew).
func sheetSourceKeys(sheet string) ([]string, error) {
	s, err := dnd5e.ParseSheet(sheet)
	return s.SourceKeys(), err
}

// planNewSources shows the sources a new character's file brings that the
// campaign lacks. There is no merge to fall back on, so one Decline blocks the
// add, and so does a source not chosen yet.
func planNewSources(sheet, name string, sources []reviewSource, form func(string) string) ([]views.HomeFileSource, string, error) {
	s, err := dnd5e.ParseSheet(sheet)
	if err != nil {
		return nil, "", err
	}
	var out []views.HomeFileSource
	declined, undecided := "", false
	for _, rs := range sources {
		choice := form(views.SourceParam(rs.Key))
		if choice != views.SourceAllow && choice != views.SourceDecline {
			choice = ""
		}
		sv := views.HomeFileSource{Key: rs.Key, Label: rs.Label, Choice: choice, Uses: "Not used", Effect: "Declining rejects " + name}
		if uses := s.SourceUses(rs.Key); len(uses) > 0 {
			sv.Uses = "Used: " + strings.Join(uses, ", ")
		}
		switch {
		case choice == views.SourceDecline && declined == "":
			declined = rs.Label
		case choice == "":
			undecided = true
		}
		out = append(out, sv)
	}
	switch {
	case declined != "":
		return out, fmt.Sprintf("Declining %s rejects %s.", declined, name), nil
	case undecided:
		return out, "Choose Allow or Decline for each source.", nil
	}
	return out, "", nil
}

// sourceLabel names a source for the GM.
func sourceLabel(s export.TakeHomeSource) string { return views.SourceLabel(s.SourceKey, s.Edition) }
