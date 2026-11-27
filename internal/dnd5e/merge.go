package dnd5e

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"strings"
)

// The three-way merge of a bring-back: the character as it was taken out
// (base), at the table now and at home now. The sheet's keys and the
// character's columns fall into four groups; a group changed on one side only
// takes that side, one changed on both and differing is the GM's pick.

// Merge groups.
const (
	GroupBuild  = "build"
	GroupHealth = "health"
	GroupGear   = "gear"
	GroupStory  = "story"
)

// Sides a group's value can come from.
const (
	SideHome  = "home"
	SideTable = "table"
)

// MergeUnchanged and MergeConflict are a group's states besides SideHome and
// SideTable (changed on that side only).
const (
	MergeUnchanged = "unchanged"
	MergeConflict  = "conflict"
)

// MergeGroups are the groups in the order the review shows them.
var MergeGroups = []string{GroupBuild, GroupHealth, GroupGear, GroupStory}

var groupLabels = map[string]string{GroupBuild: "Build", GroupHealth: "Health", GroupGear: "Gear & coins", GroupStory: "Story"}

// GroupLabel is a group's name for the GM.
func GroupLabel(group string) string { return groupLabels[group] }

// groupKeys are the sheet's JSON keys of each group. Besides them a group owns
// character columns: Build MaxHP, Health HP and TempHP, Story Name.
// TestMergeGroupsCoverSheet keeps them in step with Sheet.
var groupKeys = map[string][]string{
	GroupBuild: strings.Fields(`species species_ability background species_source background_source feat_sources classes abilities
		ac_bonus ac speed size vision save_profs skill_profs expertise other_profs armor_training
		prof_sources resistances immunities vulnerabilities spell_slots pact_slots pact_level spells
		masteries modifiers features feats last_class setup_done setup_snapshots effects_version`),
	GroupHealth: strings.Fields(`hit_dice_used death_saves inspiration`),
	GroupGear:   strings.Fields(`inventory weapons armor shield coins`),
	GroupStory:  strings.Fields(`backstory languages`),
}

// GroupKeys is the sheet keys of a group, for tests.
func GroupKeys(group string) []string { return groupKeys[group] }

// GroupDiff is how one group stands in a merge.
type GroupDiff struct {
	Group string
	// State is MergeUnchanged (also when both sides changed it alike), SideHome
	// or SideTable (that side only) or MergeConflict.
	State string
	// Side is where the merged character takes the group from.
	Side string
	// Home and Table describe what each side changed since the base, empty for
	// a side that did not (or when the group is unchanged).
	Home, Table []string
}

type sheetMap map[string]json.RawMessage

func toMap(s Sheet) (sheetMap, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	var m sheetMap
	return m, json.Unmarshal(b, &m)
}

// signature is a group's value, sheet keys and columns, as comparable text.
func signature(group string, c Character, m sheetMap) string {
	var b strings.Builder
	for _, k := range groupKeys[group] {
		fmt.Fprintf(&b, "%s=%s;", k, m[k])
	}
	switch group {
	case GroupBuild:
		fmt.Fprintf(&b, "max_hp=%d", c.MaxHP)
	case GroupHealth:
		fmt.Fprintf(&b, "hp=%d;temp_hp=%d", c.HP, c.TempHP)
	case GroupStory:
		fmt.Fprintf(&b, "name=%q", c.Name)
	}
	return b.String()
}

// takeGroup copies a group's sheet keys and columns from src into dst.
func takeGroup(group string, dst *Character, dm sheetMap, src Character, sm sheetMap) {
	for _, k := range groupKeys[group] {
		if v, ok := sm[k]; ok {
			dm[k] = v
		} else {
			delete(dm, k)
		}
	}
	switch group {
	case GroupBuild:
		dst.MaxHP = src.MaxHP
	case GroupHealth:
		dst.HP, dst.TempHP = src.HP, src.TempHP
	case GroupStory:
		dst.Name = src.Name
	}
}

// finish turns the merged keys into a valid character: HP never above max.
func finish(c Character, m sheetMap) (Character, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return c, err
	}
	if c.Sheet, err = ParseSheet(string(b)); err != nil {
		return c, err
	}
	if c.MaxHP > 0 {
		c.HP = min(c.HP, c.MaxHP)
	}
	return c, c.Validate()
}

// Merge merges the home character into the table's against the base. picks
// names the side (SideHome, SideTable) of each conflicting group; home is the
// default. All three sheets must come from ParseSheet.
func Merge(base, table, home Character, picks map[string]string) (Character, []GroupDiff, error) {
	bm, err := toMap(base.Sheet)
	if err != nil {
		return table, nil, err
	}
	tm, err := toMap(table.Sheet)
	if err != nil {
		return table, nil, err
	}
	hm, err := toMap(home.Sheet)
	if err != nil {
		return table, nil, err
	}
	merged, mm := table, maps.Clone(tm)
	diffs := make([]GroupDiff, 0, len(MergeGroups))
	for _, g := range MergeGroups {
		sb, st, sh := signature(g, base, bm), signature(g, table, tm), signature(g, home, hm)
		d := GroupDiff{Group: g, Side: SideTable, State: MergeUnchanged}
		switch homeChanged, tableChanged := sb != sh, sb != st; {
		case sh == st:
		case homeChanged && !tableChanged:
			d.State, d.Side = SideHome, SideHome
		case tableChanged && !homeChanged:
			d.State = SideTable
		default:
			d.State, d.Side = MergeConflict, SideHome
			if picks[g] == SideTable {
				d.Side = SideTable
			}
		}
		if d.State != MergeUnchanged {
			if sb != sh {
				d.Home = describe(g, base, home)
			}
			if sb != st {
				d.Table = describe(g, base, table)
			}
		}
		if d.Side == SideHome {
			takeGroup(g, &merged, mm, home, hm)
		}
		diffs = append(diffs, d)
	}
	merged, err = finish(merged, mm)
	return merged, diffs, err
}

// TakeBuild gives c the Build group of from (a fallback build), the rest kept.
func TakeBuild(c, from Character) (Character, error) {
	cm, err := toMap(c.Sheet)
	if err != nil {
		return c, err
	}
	fm, err := toMap(from.Sheet)
	if err != nil {
		return c, err
	}
	takeGroup(GroupBuild, &c, cm, from, fm)
	return finish(c, cm)
}

func orNone(s string) string {
	if s == "" {
		return "(none)"
	}
	return s
}

// diffNames are the names in side but not in base, and the other way round.
func diffNames(base, side []string) (added, removed []string) {
	for _, n := range side {
		if !slices.Contains(base, n) {
			added = append(added, n)
		}
	}
	for _, n := range base {
		if !slices.Contains(side, n) {
			removed = append(removed, n)
		}
	}
	return added, removed
}

// addedRemoved is the "added" and "removed" lines of a name list.
func addedRemoved(label string, base, side []string) []string {
	var out []string
	added, removed := diffNames(base, side)
	if len(added) > 0 {
		out = append(out, label+" added: "+strings.Join(added, ", "))
	}
	if len(removed) > 0 {
		out = append(out, label+" removed: "+strings.Join(removed, ", "))
	}
	return out
}

func spellNames(s Sheet) []string {
	var out []string
	for _, sp := range s.Spells {
		out = append(out, sp.Name)
	}
	return out
}

func itemNames(s Sheet) []string {
	var out []string
	for _, it := range s.Inventory {
		out = append(out, it.Name)
	}
	return out
}

// describe says in short lines what side changed in a group since base.
func describe(group string, base, side Character) []string {
	var out []string
	b, s := base.Sheet, side.Sheet
	switch group {
	case GroupBuild:
		if bl, sl := b.Level(), s.Level(); bl != sl {
			out = append(out, fmt.Sprintf("Level %d → %d", bl, sl))
		}
		for _, c := range s.Classes {
			i := slices.IndexFunc(b.Classes, func(x ClassLevel) bool { return x.Slug == c.Slug })
			switch {
			case i < 0:
				out = append(out, "Class "+c.Name+" added")
			case b.Classes[i].Subclass != c.Subclass:
				out = append(out, fmt.Sprintf("%s subclass: %s → %s", c.Name, orNone(b.Classes[i].Subclass), orNone(c.Subclass)))
			}
		}
		for _, c := range b.Classes {
			if !slices.ContainsFunc(s.Classes, func(x ClassLevel) bool { return x.Slug == c.Slug }) {
				out = append(out, "Class "+c.Name+" removed")
			}
		}
		if base.MaxHP != side.MaxHP {
			out = append(out, fmt.Sprintf("Max HP %d → %d", base.MaxHP, side.MaxHP))
		}
		out = append(out, addedRemoved("Spells", spellNames(b), spellNames(s))...)
	case GroupHealth:
		if base.HP != side.HP {
			if side.MaxHP > 0 {
				out = append(out, fmt.Sprintf("HP %d → %d / %d", base.HP, side.HP, side.MaxHP))
			} else {
				out = append(out, fmt.Sprintf("HP %d → %d", base.HP, side.HP))
			}
		}
		if base.TempHP != side.TempHP {
			out = append(out, fmt.Sprintf("Temp HP %d → %d", base.TempHP, side.TempHP))
		}
		if b.HitDiceUsed != s.HitDiceUsed {
			out = append(out, fmt.Sprintf("Hit dice spent %d → %d", b.HitDiceUsed, s.HitDiceUsed))
		}
		if b.DeathSaves != s.DeathSaves {
			out = append(out, fmt.Sprintf("Death saves %d ok, %d failed → %d ok, %d failed", b.DeathSaves.Successes, b.DeathSaves.Failures, s.DeathSaves.Successes, s.DeathSaves.Failures))
		}
		if b.Inspiration != s.Inspiration {
			out = append(out, map[bool]string{true: "Inspiration gained", false: "Inspiration used"}[s.Inspiration])
		}
	case GroupGear:
		out = append(out, addedRemoved("Items", itemNames(b), itemNames(s))...)
		for _, it := range s.Inventory {
			if i := slices.IndexFunc(b.Inventory, func(x Item) bool { return x.Name == it.Name }); i >= 0 && b.Inventory[i].Qty != it.Qty {
				out = append(out, fmt.Sprintf("%s ×%d → ×%d", it.Name, b.Inventory[i].Qty, it.Qty))
			}
		}
		for _, c := range []struct {
			unit      string
			was, then int
		}{{"cp", b.Coins.CP, s.Coins.CP}, {"sp", b.Coins.SP, s.Coins.SP}, {"ep", b.Coins.EP, s.Coins.EP}, {"gp", b.Coins.GP, s.Coins.GP}, {"pp", b.Coins.PP, s.Coins.PP}} {
			if c.was != c.then {
				out = append(out, fmt.Sprintf("%d %s → %d %s", c.was, c.unit, c.then, c.unit))
			}
		}
		if wb, ws := wornName(b), wornName(s); wb != ws {
			out = append(out, fmt.Sprintf("Armor: %s → %s", orNone(wb), orNone(ws)))
		}
		if b.Shield != s.Shield {
			out = append(out, map[bool]string{true: "Shield equipped", false: "Shield put away"}[s.Shield])
		}
	case GroupStory:
		if base.Name != side.Name {
			out = append(out, fmt.Sprintf("Name %s → %s", base.Name, side.Name))
		}
		if b.Backstory != s.Backstory {
			out = append(out, "Backstory edited")
		}
		out = append(out, addedRemoved("Languages", b.Languages, s.Languages)...)
	}
	if len(out) == 0 {
		out = append(out, "Other details changed")
	}
	return out
}

func wornName(s Sheet) string {
	if s.Armor == nil {
		return ""
	}
	return s.Armor.Name
}

// SourceUses names what on the sheet pins the compendium source key: its
// classes, species, background, feats, spells and armor.
func (s Sheet) SourceUses(key string) []string {
	var out []string
	s.eachPin(func(k, kind, name string) {
		if u := kind + " " + name; k == key && !slices.Contains(out, u) {
			out = append(out, u)
		}
	})
	return out
}

// SourceKeys is every book the sheet pins (homebrew is not one).
func (s Sheet) SourceKeys() []string {
	var out []string
	s.eachPin(func(k, _, _ string) {
		if k != "" && k != homebrewSource && !slices.Contains(out, k) {
			out = append(out, k)
		}
	})
	return out
}

// eachPin calls fn for every source pin on the sheet (key "" when unpinned),
// with what it pins.
func (s Sheet) eachPin(fn func(key, kind, name string)) {
	for _, c := range s.Classes {
		fn(c.Source, "class", c.Name)
		fn(c.SubclassSource, "subclass", c.Subclass)
	}
	fn(s.SpeciesSource, "species", s.Species)
	fn(s.BackgroundSource, "background", s.Background)
	for _, slug := range slices.Sorted(maps.Keys(s.FeatSources)) {
		name := slug // a pin whose feat left the list keeps its slug
		if i := slices.IndexFunc(s.Feats, func(n string) bool { return sameSlug(n, slug) }); i >= 0 {
			name = s.Feats[i]
		}
		fn(s.FeatSources[slug], "feat", name)
	}
	for _, sp := range s.Spells {
		fn(sp.Source, "spell", sp.Name)
	}
	for _, it := range s.Inventory {
		if it.Armor != nil {
			fn(it.Armor.Source, "armor", it.Name)
		}
	}
	if s.Armor != nil {
		fn(s.Armor.Source, "armor", s.Armor.Name)
	}
}

// homebrewSource is the compendium source key of homebrew entries.
const homebrewSource = "homebrew"

// RenameSpells gives homebrew spells the slugs renames maps them to (old to
// new), for a bring-back whose homebrew got other slugs at the table. One pass,
// so a new slug that is another's old one is not renamed again. A spell pinned
// to another source is not homebrew and stays.
func (s *Sheet) RenameSpells(renames map[string]string) {
	for i, sp := range s.Spells {
		if renamed, ok := renames[sp.Slug]; ok && (sp.Source == "" || sp.Source == homebrewSource) {
			s.Spells[i].Slug = renamed
		}
	}
}

// DropSource removes what a declined source supplied and the sheet can lose:
// its spells and armor (carried and worn). The build cannot be dropped (see
// BuildPinsSource). It returns what went, named as SourceUses does.
func (s *Sheet) DropSource(key string) []string {
	var removed []string
	note := func(kind, name string) {
		if u := kind + " " + name; !slices.Contains(removed, u) {
			removed = append(removed, u)
		}
	}
	s.Spells = slices.DeleteFunc(s.Spells, func(sp Spell) bool {
		if sp.Source == key {
			note("spell", sp.Name)
		}
		return sp.Source == key
	})
	s.Inventory = slices.DeleteFunc(s.Inventory, func(it Item) bool {
		drop := it.Armor != nil && it.Armor.Source == key
		if drop {
			note("armor", it.Name)
		}
		return drop
	})
	if s.Armor != nil && s.Armor.Source == key {
		note("armor", s.Armor.Name)
		s.Armor = nil
	}
	return removed
}
