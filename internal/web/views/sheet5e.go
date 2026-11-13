package views

// The one door from system-neutral views (the combat tracker, live keys)
// into D&D 5e rules, so those files never import dnd5e (depguard in
// .golangci.yml checks it).

import (
	"fmt"
	"slices"

	"gm-session-manager/internal/dnd5e"
	selectcomp "gm-session-manager/internal/web/components/select"
)

// Defenses are how a combatant takes damage.
type Defenses = dnd5e.Defenses

// customCondition is the condition slug whose note is its name.
const customCondition = dnd5e.TrackCustom

// conditionSummary is a one-line reminder of a condition's effect; custom
// effects have none.
func conditionSummary(slug string) (string, bool) {
	tip, ok := dnd5e.ConditionSummaries[slug]
	return tip, ok
}

// customToggleJS makes the note required, and says why, while "custom…" is
// the chosen condition. It runs on the form's change events.
var customToggleJS = fmt.Sprintf(`const custom = this.slug.value === %q; this.note.required = custom; this.querySelector("[data-custom-hint]").hidden = !custom`, customCondition)

func conditionItems() []selectcomp.ItemData {
	var items []selectcomp.ItemData
	for _, slug := range dnd5e.Conditions {
		items = append(items, selectcomp.ItemData{Value: slug, Label: slug})
	}
	return append(items,
		selectcomp.ItemData{Value: dnd5e.TrackConcentration, Label: "concentration"},
		selectcomp.ItemData{Value: customCondition, Label: "custom…"},
	)
}

// damageTypeItems offers untyped damage first, then each type a defense
// names, then each conditional defense's types as "<type>@<index>"
// (damageChoice reads them back).
func damageTypeItems(d Defenses) []selectcomp.ItemData {
	d.ResistAll = false // Petrified has its own note
	effect := map[int]string{0: "immune, none", 2: "resisted, half", 4: "resisted and vulnerable", 8: "vulnerable, double"}
	items := []selectcomp.ItemData{{Value: " ", Label: "Other"}} //nolint:goconst // unrelated to the other "Other" labels
	for _, t := range dnd5e.DamageTypes {
		if slices.ContainsFunc(d.List, func(def dnd5e.Defense) bool { return def.When == "" && slices.Contains(def.Types, t) }) {
			items = append(items, selectcomp.ItemData{Value: t, Label: fmt.Sprintf("%s (%s)", capitalized([]string{t})[0], effect[d.Taken(4, t, -1)])})
		}
	}
	for i, def := range d.List {
		if def.When == "" {
			continue
		}
		for _, t := range def.Types {
			items = append(items, selectcomp.ItemData{Value: fmt.Sprintf("%s@%d", t, i), Label: fmt.Sprintf("%s %s (%s)", capitalized([]string{t})[0], def.When, effect[d.Taken(4, t, i)])})
		}
	}
	return items
}
