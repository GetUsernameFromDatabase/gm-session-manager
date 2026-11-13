package views

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/perm"
)

// SheetView is a character as one viewer sees it.
type SheetView struct {
	ID         int64
	CampaignID int64
	C          dnd5e.Character
	Conditions []sqlc.Condition
	// Looks are the character's images the viewer may see (perm.VisibleLooks).
	Looks []perm.Look
	// What the viewer may change (see perm.CanCharacter).
	CanHP, CanConditions, CanSheet bool
	// CanRevert is set while a level-up can be undone.
	CanRevert bool
	// SpellData are the compendium details of the sheet's spells, by slug.
	SpellData map[string]*compendium.SpellData
	// CanSources lets the viewer switch sources; SourceIssues are only
	// computed then.
	CanSources   bool
	SourceIssues []SourceIssue
	// MasteryLimit is how many weapon kinds the classes let the character
	// master (dnd5e.Sheet.MasteryLimit).
	MasteryLimit int
}

func (s SheetView) cmdURL() string { return fmt.Sprintf("/characters/%d/cmd", s.ID) }

func (s SheetView) condition(slug string) (sqlc.Condition, bool) {
	i := slices.IndexFunc(s.Conditions, func(c sqlc.Condition) bool { return c.Slug == slug })
	if i < 0 {
		return sqlc.Condition{}, false
	}
	return s.Conditions[i], true
}

// defenses are the sheet's, resisting all damage while Petrified.
func (s SheetView) defenses() dnd5e.Defenses {
	d := s.C.Sheet.Defenses()
	_, d.ResistAll = s.condition(dnd5e.Petrified)
	return d
}

// tapSize lifts small sheet controls to 32px on phones.
const tapSize = "max-sm:h-8 max-sm:px-3"

func cmdVals(cmd string) string {
	return fmt.Sprintf(`{"cmd":%q}`, cmd)
}

// subtitlePart is one piece of the line under the name.
type subtitlePart struct {
	text, what string
	to         edit
}

// subtitle is species, background and classes, leaving out empty ones;
// classes change through levelling up.
func (s SheetView) subtitle() []subtitlePart {
	parts := []subtitlePart{
		{s.C.Sheet.Species, "species", edit{verbDetails, opSetSpecies, ""}},
		{s.C.Sheet.Background, "background", edit{verbDetails, opSetBackground, ""}},
		{s.C.Sheet.ClassSummary(), "classes", edit{}},
	}
	return slices.DeleteFunc(parts, func(p subtitlePart) bool { return p.text == "" })
}

func hitDice(s dnd5e.Sheet) string {
	total := s.Level()
	if total == 0 {
		return "–"
	}
	dice := s.HitDice()
	sizes := slices.Sorted(maps.Keys(dice))
	slices.Reverse(sizes)
	parts := make([]string, len(sizes))
	for i, d := range sizes {
		if len(sizes) == 1 {
			parts[i] = fmt.Sprintf("d%d", d)
		} else {
			parts[i] = fmt.Sprintf("%dd%d", dice[d], d)
		}
	}
	return fmt.Sprintf("%d/%d %s", total-s.HitDiceUsed, total, strings.Join(parts, " + "))
}

// edit names the builder form a value opens; the zero edit opens none.
type edit struct{ verb, op, subject string }

// The stats of the sheet's stat row, by the name the row shows.
const (
	statAC              = "AC"
	statInit            = "Init"
	statSpeed           = "Speed"
	statProf            = "Prof"
	statPassive         = "Passive Perception"
	statHitDice         = "Hit dice"
	statSize            = "Size"
	statVision          = "Vision"
	statResistances     = "Resistances"
	statImmunities      = "Immunities"
	statVulnerabilities = "Vulnerabilities"
)

// statWhy are the stats with a breakdown (dnd5e.Sheet.Explain target).
var statWhy = map[string]string{statAC: dnd5e.ModAC, statSpeed: dnd5e.ModSpeed, statInit: dnd5e.ModInit, statPassive: dnd5e.ModPassive}

// plainTotals are the breakdown targets that are values, not bonuses.
var plainTotals = []string{dnd5e.ModAC, dnd5e.ModSpeed, dnd5e.ModPassive, dnd5e.ExplainSpellDC}

// spareArmor is the armor and shields in the inventory not worn.
func (s SheetView) spareArmor() []dnd5e.Item {
	return slices.DeleteFunc(slices.Clone(s.C.Sheet.Inventory), func(it dnd5e.Item) bool {
		return it.Armor == nil && !it.Shield || s.C.Sheet.Worn(it)
	})
}

// spareWeapons are the weapons in the inventory not ready to attack with.
func (s SheetView) spareWeapons() []dnd5e.Item {
	return slices.DeleteFunc(slices.Clone(s.C.Sheet.Inventory), func(it dnd5e.Item) bool {
		return it.Weapon == nil || s.C.Sheet.Ready(it)
	})
}

// statTips explain the stats whose names do not say what they are for.
var statTips = map[string]string{ //nolint:gosec // G101: tooltip text, not a credential
	statPassive:         "10 + your Perception bonus: what you notice without looking. Hidden creatures' Stealth is checked against it.",
	statInit:            "Initiative: add it to a d20 when combat starts.",
	statAC:              "Armor class: an attack hits you if its roll meets or beats this.",
	statSpeed:           "How far you can move on your turn, in feet.",
	statProf:            "Proficiency bonus: added to the attacks, saves and skills you are proficient in.",
	statHitDice:         "Spend these on a short rest to heal; a long rest restores them.",
	statVision:          "What you can see in darkness and dim light without a light source.",
	statResistances:     "Damage of these types is halved.",
	statImmunities:      "Damage of these types is ignored.",
	statVulnerabilities: "Damage of these types is doubled.",
}

// capitalized are keys such as damage types as names: "fire" is "Fire".
func capitalized(keys []string) []string {
	out := make([]string, len(keys))
	for i, k := range keys {
		out[i] = strings.ToUpper(k[:1]) + k[1:]
	}
	return out
}

// hasMods reports whether any custom modifier is on one of targets.
func (s SheetView) hasMods(targets ...string) bool {
	return slices.ContainsFunc(s.C.Sheet.Modifiers, func(m dnd5e.CustomMod) bool { return slices.Contains(targets, m.Target) })
}

// commandsID is the palette that explains the commands and fills one in.
const commandsID = "command-help"

// actionKinds are what an action takes, in turn order.
var actionKinds = []struct{ key, title, tip string }{
	{dnd5e.ActionMain, "Action", "One per turn: attack, cast a spell, Dash, Dodge, Help, Hide…"},
	{dnd5e.ActionBonus, "Bonus Action", "One per turn, only for things that say so."},
	{dnd5e.ActionReaction, "Reaction", "One per round, even on others' turns, when its trigger happens."},
	{actionOther, labelOther, "Limited features that take no action of their own."},
}

// actionSections are the Actions card's sections: what each action is.
var actionSections = []struct{ key, title string }{
	{sectionAttacks, "Attacks"}, {sectionSpells, "Spells"}, {sectionFeatures, "Features"},
}

// The Actions card's section keys.
const (
	sectionAttacks  = "attacks"
	sectionSpells   = "spells"
	sectionFeatures = "features"
)

// actionEntry is one thing the character can do, with why they can.
type actionEntry struct {
	Name, Detail, Why string
	// Kind is what it takes (an actionKinds key).
	Kind string
	// Weapon, for attacks, breaks down the bonuses to hit and damage.
	Weapon string
	// OffHand marks a weapon's two-weapon fighting bonus attack.
	OffHand bool
	// SpellSlug opens the spell's text; FeatureIdx (with Uses) shows and
	// spends its uses.
	SpellSlug  string
	FeatureIdx int
	Uses       bool
}

var (
	reSpellAttack = regexp.MustCompile(`(?i)make an? (?:ranged|melee) spell attack`)
	reSpellSave   = regexp.MustCompile(`(?i)\b(strength|dexterity|constitution|intelligence|wisdom|charisma) saving throw`)
)

// actions sorts what the character can do into sections: ready weapons,
// spells they can cast now and features, each in turn order.
func (s SheetView) actions() map[string][]actionEntry {
	sh := s.C.Sheet
	out := map[string][]actionEntry{}
	for _, w := range sh.Weapons {
		detail := fmt.Sprintf("%s to hit · %s · %s", dnd5e.Signed(sh.AttackBonus(w)), sh.DamageLabel(w), weaponReach(w))
		out[sectionAttacks] = append(out[sectionAttacks], actionEntry{Name: w.Name, Detail: detail, Why: w.Properties, Kind: dnd5e.ActionMain, Weapon: w.Name})
	}
	if sh.CanDualWield() {
		for _, w := range sh.LightWeapons() {
			detail := fmt.Sprintf("%s to hit · %s · %s", dnd5e.Signed(sh.AttackBonus(w)), sh.OffHandDamage(w), weaponReach(w))
			out[sectionAttacks] = append(out[sectionAttacks], actionEntry{Name: "Off-hand: " + w.Name, Detail: detail, Why: offHandRule, Kind: dnd5e.ActionBonus, Weapon: w.Name, OffHand: true})
		}
	}
	for _, sp := range sh.Spells {
		d := s.SpellData[sp.Slug]
		if d == nil || sp.Level > 0 && !sp.Prepared {
			continue
		}
		ab := sh.SpellAbilityOf(sp)
		kind := castingKind(d.CastingTime)
		detail := strings.TrimSpace(d.CastingTime + " · " + d.Range)
		if ab != "" {
			if reSpellAttack.MatchString(d.Desc) {
				detail += " · " + dnd5e.Signed(sh.SpellAttack(ab)) + " to hit"
			} else if m := reSpellSave.FindStringSubmatch(d.Desc); m != nil {
				detail += fmt.Sprintf(" · DC %d %s save", sh.SpellDC(ab), strings.ToUpper(m[1][:3]))
			}
		}
		if sp.Level > 0 {
			detail += fmt.Sprintf(" · level %d slot", sp.Level)
		}
		out[sectionSpells] = append(out[sectionSpells], actionEntry{Name: sp.Name, Detail: detail, SpellSlug: sp.Slug, Kind: kind})
	}
	for i, f := range sh.Features {
		if f.Action == "" {
			f.Action = s.freeCastAction(f.Name)
		}
		kind := f.Action
		if kind == dnd5e.ActionMagic {
			kind = dnd5e.ActionMain
		}
		if kind == "" {
			if !f.Uses.Limited() {
				continue
			}
			kind = actionOther
		}
		detail := strings.TrimSpace(dnd5e.ActionLabels[f.Action] + " · " + f.Source)
		out[sectionFeatures] = append(out[sectionFeatures], actionEntry{Name: f.Name, Detail: strings.Trim(detail, " ·"), Why: f.Description, Kind: kind, FeatureIdx: i, Uses: f.Uses.Limited()})
	}
	for _, entries := range out {
		slices.SortStableFunc(entries, func(a, b actionEntry) int { return kindOrder(a.Kind) - kindOrder(b.Kind) })
	}
	return out
}

// freeCastAction is what a lineage spell's free cast (see dnd5e.FreeCast)
// takes: its spell's casting time, or "" when that is longer.
func (s SheetView) freeCastAction(name string) string {
	spell, ok := strings.CutSuffix(name, " (free cast)")
	if !ok {
		return ""
	}
	i := slices.IndexFunc(s.C.Sheet.Spells, func(sp dnd5e.Spell) bool { return sp.Name == spell })
	if i < 0 || s.SpellData[s.C.Sheet.Spells[i].Slug] == nil {
		return ""
	}
	return strings.TrimPrefix(castingKind(s.SpellData[s.C.Sheet.Spells[i].Slug].CastingTime), actionOther)
}

func kindOrder(kind string) int {
	return slices.IndexFunc(actionKinds, func(k struct{ key, title, tip string }) bool { return k.key == kind })
}

// offHandRule is the Light property's extra attack, paraphrased from the
// SRD 5.2 (CC-BY 4.0, Wizards of the Coast).
const offHandRule = "Bonus Action, after you take the Attack action with a different Light weapon. No ability modifier to its damage unless it is negative."

// castingKind sorts a casting time ("action", "bonus action", "reaction,
// which you take when…", "1 minute") into an Actions card group.
func castingKind(t string) string {
	t = strings.ToLower(t)
	switch {
	case strings.HasPrefix(t, "bonus"):
		return dnd5e.ActionBonus
	case strings.HasPrefix(t, "reaction"):
		return dnd5e.ActionReaction
	case strings.HasPrefix(t, "action"), strings.HasPrefix(t, "1 action"):
		return dnd5e.ActionMain
	}
	return actionOther
}

// actionOther is the Actions card group for what takes no action.
const actionOther = "other"

// weaponReach is "melee" or "ranged 80/320".
func weaponReach(w dnd5e.Weapon) string {
	if w.Range == "" {
		return "melee"
	}
	if strings.Contains(strings.ToLower(w.Properties), "thrown") {
		return "melee or thrown " + w.Range + " ft"
	}
	return "ranged " + w.Range + " ft"
}

// usesLabel is "2/3 left" for a feature's uses.
func (s SheetView) usesLabel(f dnd5e.Feature) string {
	n := s.C.Sheet.MaxUses(f.Uses)
	return fmt.Sprintf("%d/%d left", max(n-f.Uses.Used, 0), n)
}

// weaponAbility names the ability a weapon attacks with; finesse takes
// the better of STR and DEX, as AttackBonus does.
func (s SheetView) weaponAbility(w dnd5e.Weapon) string {
	if w.Ability != dnd5e.Finesse {
		return strings.ToUpper(w.Ability)
	}
	if s.C.Sheet.Mod(dnd5e.Dex) > s.C.Sheet.Mod(dnd5e.Str) {
		return "DEX (finesse)"
	}
	return "STR (finesse)"
}

// attackTargets are the modifier targets that change a weapon's kind (to
// hit or damage), including those on its ability.
func attackTargets(w dnd5e.Weapon, kind string) []string {
	t := []string{kind, kind + "." + w.Name}
	if w.Ability == dnd5e.Finesse {
		return append(t, dnd5e.Str, dnd5e.Dex)
	}
	return append(t, w.Ability)
}
