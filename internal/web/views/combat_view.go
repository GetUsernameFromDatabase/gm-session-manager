package views

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/internal/db/sqlc"
	selectcomp "gm-session-manager/internal/web/components/select"
	"gm-session-manager/internal/web/utils"

	"github.com/a-h/templ"
)

// MaxHPInput bounds the HP dialog's numbers, well above any 5e creature.
const MaxHPInput = 9999

// MaxMarker bounds a combatant's table marker, in characters.
const MaxMarker = 40

// chipConditions is how many conditions a player's chip shows before "+N".
const chipConditions = 2

// CombatantView is one row of the turn order as a viewer sees it.
type CombatantView struct {
	ID          int64
	Name        string
	Initiative  string
	AC          int64
	HP          int64
	MaxHP       int64
	TempHP      int64
	ShowHP      bool
	HPVisible   bool
	CharacterID int64
	EntityID    int64
	// Kind is the type of being: KindPC, KindNPC or KindMonster.
	Kind       string
	IsOwn      bool
	Active     bool
	Conditions []sqlc.Condition
	// Out is a combatant out of the fight, whose turn is skipped.
	Out bool
	// Down is at 0 HP, shown to everyone even when the HP is hidden.
	Down bool
	// InitMod is the initiative roll's bonus, GM only.
	InitMod int
	// Link is the combatant's source page, when the viewer may open it.
	Link string
	// Portrait is a character's or NPC's current look, when the viewer may see
	// it, else the compendium monster's picture.
	Portrait string
	// Marker names its piece on the physical table, GM only.
	Marker string
	// Defenses are how it takes damage, for the GM and the player's own row.
	Defenses Defenses
}

// EncounterView is the tracker for one encounter.
type EncounterView struct {
	Encounter  sqlc.Encounter
	Combatants []CombatantView
	IsGM       bool
	// Add-combatant choices, GM only.
	Characters []Named
	Entities   []Named
}

// Named is a row offered by name. It holds nothing else, so the tracker's
// hashed model (Shape) always covers what the add form shows.
type Named struct {
	ID   int64
	Name string
}

// Combatant kinds, the type of being shown on each row.
const (
	KindPC      = "pc"
	KindNPC     = "npc"
	KindMonster = "monster"
)

// kindLabel and kindTint set the kinds apart, as visibility does note blocks.
var (
	kindLabel  = map[string]string{KindPC: "PC", KindNPC: "NPC", KindMonster: "Monster"}
	kindTint   = map[string]string{KindPC: "border-l-sky-400/60", KindNPC: "border-l-amber-400/60", KindMonster: "border-l-rose-400/60"}
	kindGroup  = map[string]string{KindPC: "PCs", KindNPC: "NPCs", KindMonster: "Monsters"}
	kindBorder = map[string]string{KindPC: "border-sky-400/60", KindNPC: "border-amber-400/60", KindMonster: "border-rose-400/60"}
	kindDot    = map[string]string{KindPC: "bg-sky-400/60", KindNPC: "bg-amber-400/60", KindMonster: "bg-rose-400/60"}
)

// CombatView is the combat page.
type CombatView struct {
	CampaignID int64
	IsGM       bool
	Encounters []sqlc.Encounter
	Current    *EncounterView
	// Want is the encounter a player asked for; 0 follows the live fight.
	Want int64
}

func currentID(v CombatView) int64 {
	if v.Current == nil {
		return 0
	}
	return v.Current.Encounter.ID
}

func (e EncounterView) url(action string) string {
	return fmt.Sprintf("/encounters/%d/%s", e.Encounter.ID, action)
}

func (c CombatantView) hasMeter() bool { return c.ShowHP && c.MaxHP > 0 }

// combatantState marks whose turn it is and who is down; live updates flip
// these attributes.
func combatantState(c CombatantView) templ.Attributes {
	a := templ.Attributes{"data-active": fmt.Sprint(c.Active), "data-down": fmt.Sprint(c.down())}
	return with(a, bind("data-active="+combatantKey(c, "on"), "data-down="+combatantKey(c, "down")))
}

// initials is the first letter of the name's first two words ("Goblin 2" → "G2").
func initials(name string) string {
	words := strings.Fields(name)
	var b strings.Builder
	for _, w := range words[:min(2, len(words))] {
		b.WriteString(strings.ToUpper(string([]rune(w)[0])))
	}
	return b.String()
}

// hpCap is the highest HP the dialog accepts: the max HP, or any value while
// a monster has no stats yet.
func (c CombatantView) hpCap() int64 {
	if c.MaxHP > 0 {
		return c.MaxHP
	}
	return MaxHPInput
}

// initiativeSet reports whether every combatant has an initiative, which
// next turn needs to follow the order.
func (e EncounterView) initiativeSet() bool { return len(e.missingInitiative()) == 0 }

// rollableMissing reports whether a monster or NPC still has no initiative: the GM rolls for those.
func (e EncounterView) rollableMissing() bool {
	return slices.ContainsFunc(e.Combatants, func(c CombatantView) bool { return c.Kind != KindPC && c.Initiative == "" })
}

// missingInitiative names the combatants without an initiative.
func (e EncounterView) missingInitiative() []string {
	var names []string
	for _, c := range e.Combatants {
		if c.Initiative == "" {
			names = append(names, c.Name)
		}
	}
	return names
}

// turnNames are the previous, current and next combatants in the fight,
// passing over those out of it as next turn does.
func (e EncounterView) turnNames() (prev, cur, next string) {
	n := len(e.Combatants)
	i := slices.IndexFunc(e.Combatants, func(c CombatantView) bool { return c.Active })
	if i < 0 {
		return "", "", ""
	}
	near := func(dir int) string {
		for k := 1; k < n; k++ {
			if c := e.Combatants[((i+dir*k)%n+n)%n]; !c.Out {
				return c.Name
			}
		}
		return ""
	}
	prev = near(-1)
	// On the first turn of the fight nobody has acted yet.
	if e.Encounter.Round <= 1 && !slices.ContainsFunc(e.Combatants[:i], func(c CombatantView) bool { return !c.Out }) {
		prev = "—"
	}
	return prev, e.Combatants[i].Name, near(1)
}

// turnSkipped names those out of the fight between the previous turn and
// this one, whose turns next turn passed over; empty when there are none.
func (e EncounterView) turnSkipped() string {
	n := len(e.Combatants)
	i := slices.IndexFunc(e.Combatants, func(c CombatantView) bool { return c.Active })
	if i < 0 {
		return ""
	}
	var names []string
	for k := 1; k < n; k++ {
		c := e.Combatants[((i-k)%n+n)%n]
		if !c.Out {
			break
		}
		names = append(names, c.Name)
	}
	if len(names) == 0 {
		return ""
	}
	slices.Reverse(names)
	return fmt.Sprintf("Skipped %d (out of the fight): %s", len(names), strings.Join(names, ", "))
}

// combatantGroup is one kind of being in a player's grid.
type combatantGroup struct {
	Kind, Label string
	Cards       []CombatantView
}

// playerGroups are the others in the fight grouped by kind and sorted by
// name, so the grid does not give the initiative order away.
func (e EncounterView) playerGroups() []combatantGroup {
	var groups []combatantGroup
	for _, kind := range []string{KindPC, KindNPC, KindMonster} {
		g := combatantGroup{Kind: kind, Label: kindGroup[kind]}
		for _, c := range e.Combatants {
			if c.Kind == kind && !c.IsOwn {
				g.Cards = append(g.Cards, c)
			}
		}
		if len(g.Cards) == 0 {
			continue
		}
		slices.SortStableFunc(g.Cards, func(a, b CombatantView) int {
			ab, an := nameKey(a.Name)
			bb, bn := nameKey(b.Name)
			if c := strings.Compare(strings.ToLower(ab), strings.ToLower(bb)); c != 0 {
				return c
			}
			return an - bn
		})
		groups = append(groups, g)
	}
	return groups
}

// nameKey splits a trailing number off a name, so "Bandit 2" sorts before
// "Bandit 10".
func nameKey(name string) (string, int) {
	i := strings.LastIndexByte(name, ' ')
	if i < 0 {
		return name, 0
	}
	n, err := strconv.Atoi(name[i+1:])
	if err != nil {
		return name, 0
	}
	return name[:i], n
}

// quietDestructive styles the small × removes that repeat on every row or
// chip: muted until hovered, so a list is not a wall of red.
const quietDestructive = "text-muted-foreground hover:bg-destructive/10 hover:text-destructive dark:hover:bg-destructive/20"

// selectItems maps xs to select items with the given value and label.
func selectItems[T any](xs []T, value, label func(T) string) []selectcomp.ItemData {
	items := make([]selectcomp.ItemData, 0, len(xs))
	for _, x := range xs {
		items = append(items, selectcomp.ItemData{Value: value(x), Label: label(x)})
	}
	return items
}

// wideList sizes a select's list to its longest label, never narrower than its button nor wider than the screen.
const wideList = "w-auto min-w-(--anchor-width) max-w-[calc(100vw-2rem)]"

func visibleTitle(on bool) string {
	if on {
		return "Players see this HP (click to hide)"
	}
	return "HP hidden from players (click to show)"
}

func conditionLabel(c sqlc.Condition) string {
	label := c.Slug
	if c.Slug == customCondition {
		label = c.Note // a custom condition's note is its name
	} else if c.Note != "" {
		label += " " + c.Note
	}
	if c.ExpiresRound.Valid {
		label += fmt.Sprintf(" (to round %d)", c.ExpiresRound.Int64)
	}
	return label
}

// inFight reports whether a combatant matching f is already in the fight.
func (e EncounterView) inFight(f func(CombatantView) bool) bool {
	return slices.ContainsFunc(e.Combatants, f)
}

func inFightSuffix(in bool) string { return utils.IfElse(in, " (already in fight)", "") }
