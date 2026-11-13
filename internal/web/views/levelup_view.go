package views

import (
	"encoding/json"
	"fmt"
	"net/url"
	"slices"
	"strings"

	"gm-session-manager/internal/dnd5e"
	selectcomp "gm-session-manager/internal/web/components/select"
	"gm-session-manager/internal/web/utils"
)

// Option is a compendium choice in a form.
type Option struct {
	Slug  string
	Name  string
	Level int
	// TaughtBy are "<group key>=<option>" rules: a lineage picked on the same page
	// already teaches this spell, so it is not offered while one of them is picked.
	TaughtBy []string
}

// SubclassPanel is one subclass's details, shown while it is picked.
type SubclassPanel struct {
	Slug     string
	Name     string
	Source   string
	Features []SubclassFeature
}

// SubclassFeature is a subclass feature and the class levels it is gained at.
type SubclassFeature struct {
	Name   string
	Desc   string
	Levels []int
}

// FirstLevel is the lowest level the feature is gained at (0 for none).
func (f SubclassFeature) FirstLevel() int {
	if len(f.Levels) == 0 {
		return 0
	}
	return slices.Min(f.Levels)
}

// LevelUpView is the level-up helper.
type LevelUpView struct {
	CharacterID int64
	Name        string
	Level       int
	// Sheet is the character before levelling, to show what changes.
	Sheet      dnd5e.Sheet
	Classes    []Option
	Selected   string
	Plan       *dnd5e.LevelUpPlan
	Subclasses []Option
	// SubclassPanels are the subclasses' details, in the same order.
	SubclassPanels []SubclassPanel
	Feats          []Option
	// FeatChoices are each feat's own choices, by slug (Grappler's ability).
	FeatChoices map[string][]ChoiceGroup
	Spells      []Option
	SpellNeed   SpellNeed
	// Groups are the class options this level adds (invocations) and the
	// feats' Spell Change.
	Groups []ChoiceGroup
	// LineageSpells are the species spells this character level teaches.
	LineageSpells []dnd5e.Spell
	// Error is shown at the top; Errors next to their fields.
	Error  string
	Errors FieldErrors
	// Form is the refused form's values, ticked again; nil on a fresh page.
	Form url.Values
}

func (v LevelUpView) fieldErr() FieldErrors { return v.Errors }

// DrawsField reports whether the page shows an error at key (see WizardView's).
func (v LevelUpView) DrawsField(key string) bool {
	if v.Plan == nil {
		return false
	}
	switch key {
	case FieldSpell:
		return len(v.Spells) > 0
	case FieldASI:
		return v.Plan.ASI
	}
	if drawsGroup(v.Groups, key) {
		return true
	}
	if v.Plan.ASI {
		for _, f := range v.Feats {
			if drawsGroup(v.FeatChoices[f.Slug], key) {
				return true
			}
		}
	}
	return false
}

func slotSummary(slots [9]int) string {
	out := ""
	for i, n := range slots {
		if n > 0 {
			if out != "" {
				out += ", "
			}
			out += fmt.Sprintf("L%d×%d", i+1, n)
		}
	}
	if out == "" {
		return valNone
	}
	return out
}

// slotMaxima is how many slots of each level s has.
func slotMaxima(s dnd5e.Sheet) [9]int {
	var out [9]int
	for i, sl := range s.Slots {
		out[i] = sl.Max
	}
	return out
}

func sortedColumns(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

// spellCounts is the class table's cantrip and spell totals ("Cantrips 2 ·
// Prepared Spells 3"); slot columns are left out.
func spellCounts(cols map[string]string) string {
	var out []string
	for _, k := range sortedColumns(cols) {
		if cantrips, spells := dnd5e.SpellColumn(k); cantrips || spells {
			out = append(out, k+" "+cols[k])
		}
	}
	return strings.Join(out, " · ")
}

// SpellNeed is how many class spells a step asks for: exactly, or at most
// when …UpTo (prepared spells). -1 is no limit (the class table has no such
// column). Book marks the spells as a spellbook's; BookShort, that the book
// was already below its size at the previous level (BookHave of BookSize).
type SpellNeed struct {
	Cantrips, Spells         int
	CantripsUpTo, SpellsUpTo bool
	Book, BookShort          bool
	BookHave, BookSize       int
}

// String is the ask ("2 cantrips and up to 1 level 1+ spell"), or "" if none.
func (n SpellNeed) String() string {
	var out []string
	if n.Cantrips > 0 {
		out = append(out, utils.IfElse(n.CantripsUpTo, "up to ", "")+Plural(n.Cantrips, "cantrip"))
	}
	if n.Spells > 0 {
		out = append(out, utils.IfElse(n.SpellsUpTo, "up to ", "")+Plural(n.Spells, "level 1+ spell")+utils.IfElse(n.Book, " for your spellbook", ""))
		if n.BookShort {
			out[len(out)-1] += fmt.Sprintf(" (your spellbook has %d of %d)", n.BookHave, n.BookSize)
		}
	}
	return strings.Join(out, " and ")
}

// Plural is "1 spell" or "2 spells".
func Plural(n int, noun string) string {
	return fmt.Sprintf("%d %s%s", n, noun, utils.IfElse(n == 1, "", "s"))
}

// SubclassChangeView is the form that swaps a class's subclass.
type SubclassChangeView struct {
	CharacterID      int64
	Class, ClassName string
	Current          string
	CurrentSlug      string
	Level            int
	Subclasses       []Option
	Panels           []SubclassPanel
}

func joinInts(n []int) string {
	s := make([]string, len(n))
	for i, v := range n {
		s[i] = fmt.Sprint(v)
	}
	return strings.Join(s, ", ")
}

// splitCantrips splits spells into cantrips and level 1+ spells, which the
// class table caps separately, keeping their order.
func splitCantrips(spells []Option) (cantrips, leveled []Option) {
	for _, sp := range spells {
		if sp.Level == 0 {
			cantrips = append(cantrips, sp)
		} else {
			leveled = append(leveled, sp)
		}
	}
	return cantrips, leveled
}

// taughtPicked reports whether a lineage picked on the form teaches the spell
// (one of its rules is picked), so the spell is hidden and not posted.
func taughtPicked(rules []string, form url.Values) bool {
	return slices.ContainsFunc(rules, func(r string) bool {
		key, val, _ := strings.Cut(r, "=")
		return slices.Contains(form[key], val)
	})
}

// taughtRules is the rules as a JSON array for app.js (never null).
func taughtRules(rules []string) string {
	b, _ := json.Marshal(append([]string{}, rules...))
	return string(b)
}

var abilityItems = append([]selectcomp.ItemData{{Value: "", Label: "—"}},
	selectItems(dnd5e.Abilities, func(a string) string { return a }, func(a string) string { return dnd5e.AbilityNames[a] })...)

// optionItems is a select over opts led by an empty choice labelled none.
func optionItems(none string, opts []Option, value func(Option) string) []selectcomp.ItemData {
	return append([]selectcomp.ItemData{{Value: "", Label: none}},
		selectItems(opts, value, func(o Option) string { return o.Name })...)
}
