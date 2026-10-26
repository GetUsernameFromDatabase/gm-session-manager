// Package views holds the templ pages and components, and the live-view
// keys and bindings they share with the hub.
package views

import (
	"fmt"
	"slices"
	"strings"

	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
	"gm-session-manager/hub"

	"github.com/a-h/templ"
)

// Live view kinds. A page marks each live fragment root with
// data-live="kind:id[:extra]"; the hub sends that view's changed data and
// live.js binds it, or refetches the view's HTML from /live.
const (
	LiveRoster       = "roster"       // GM: characters and devices, id = campaign
	LiveRequests     = "requests"     // GM: every campaign's join requests, id = 0
	LiveWaiting      = "waiting"      // visitor: own request, id = request secret
	LiveCharacter    = "character"    // id = character
	LiveEncounters   = "encounters"   // GM: encounter list, id = campaign, extra = shown encounter
	LiveEncounter    = "encounter"    // GM: one tracker, id = encounter
	LivePlayerCombat = "playercombat" // player: own fights, id = campaign, extra = asked-for encounter
	LiveNotes        = "notes"        // note list, id = campaign
	LiveNote         = "note"         // note title and block layout, id = note
	LiveEntities     = "entities"     // NPC list, id = campaign
	LiveEntity       = "entity"       // id = entity
)

// LiveKey builds a view key.
func LiveKey(kind string, id any, extra ...any) string {
	parts := []string{kind, fmt.Sprint(id)}
	for _, e := range extra {
		parts = append(parts, fmt.Sprint(e))
	}
	return strings.Join(parts, ":")
}

func live(kind string, id any, extra ...any) templ.Attributes {
	return templ.Attributes{"data-live": LiveKey(kind, id, extra...)}
}

// bind marks an element for live.js: each spec is target=key, where target
// is text, value (inputs, skipped while focused), show (hidden unless true),
// reveal (data-closed unless true, for CSS to transition) or an attribute name.
func bind(spec ...string) templ.Attributes {
	return templ.Attributes{"data-bind": strings.Join(spec, " ")}
}

// with merges attribute sets, later ones winning.
func with(sets ...templ.Attributes) templ.Attributes {
	out := templ.Attributes{}
	for _, s := range sets {
		for k, v := range s {
			out[k] = v
		}
	}
	return out
}

// Bound views give their hot values (Hot) and a copy with those values held
// fixed (Shape), whose rendering is hashed to catch every other change. Hot
// adds exactly the keys of the elements the real rendering binds, so a key
// appearing or disappearing means the markup changed shape.

// Hot adds the sheet's often-changing values.
func (s SheetView) Hot(st hub.State) {
	st["hp"] = fmt.Sprint(s.C.HP)
	if s.C.TempHP > 0 {
		st["temp"] = tempLabel(s.C.TempHP)
	}
	if dying(s) {
		st["ds.s"] = fmt.Sprintf("✓ %d/3", s.C.Sheet.DeathSaves.Successes)
		st["ds.f"] = fmt.Sprintf("✗ %d/3", s.C.Sheet.DeathSaves.Failures)
	}
	if s.CanSheet {
		st["insp"] = inspirationLabel(s.C.Sheet.Inspiration)
	}
	st["hd"] = hitDice(s.C.Sheet)
	for _, slug := range dnd5e.Conditions {
		if c, on := s.condition(slug); slug != exhaustion {
			st["cond."+slug] = on
		} else if on {
			st["exh"] = c.Note
		}
	}
	for i, sl := range s.C.Sheet.Slots {
		if sl.Max > 0 {
			st[fmt.Sprintf("slot.%d", i+1)] = slotLabel(sl)
		}
	}
	if s.C.Sheet.Pact.Max > 0 {
		st["pact"] = slotLabel(s.C.Sheet.Pact)
	}
	for i, f := range s.C.Sheet.Features {
		if f.Uses.Limited() {
			st[fmt.Sprintf("use.%d", i)] = s.usesLabel(f)
		}
	}
}

// Shape is the sheet with its hot values fixed.
func (s SheetView) Shape() SheetView {
	s.C.HP, s.C.TempHP = 0, 0
	s.C.Sheet.DeathSaves = dnd5e.DeathSaves{}
	s.C.Sheet.Inspiration = false
	s.C.Sheet.HitDiceUsed = 0
	for i := range s.C.Sheet.Slots { // an array: s holds its own copy
		s.C.Sheet.Slots[i].Used = 0
	}
	s.C.Sheet.Pact.Used = 0
	s.C.Sheet.Features = slices.Clone(s.C.Sheet.Features)
	for i := range s.C.Sheet.Features {
		s.C.Sheet.Features[i].Uses.Used = 0
	}
	s.Conditions = slices.DeleteFunc(slices.Clone(s.Conditions), func(c sqlc.Condition) bool { return isToggle(c.Slug) })
	return s
}

// Hot adds the tracker's often-changing values: the round, whose turn it
// is, and the hit points the viewer may see.
func (e EncounterView) Hot(st hub.State) {
	if !e.showsTracker() {
		return
	}
	if e.Encounter.Status == "active" {
		st["round"] = roundLabel(e.Encounter.Round)
		st["turn.prev"], st["turn.cur"], st["turn.next"] = e.turnNames()
		if e.IsGM {
			st["turn.skipped"] = e.turnSkipped()
		}
	}
	for _, c := range e.Combatants {
		st[combatantKey(c, "on")] = c.Active
		st[combatantKey(c, "down")] = c.down()
		if c.ShowHP {
			st[combatantKey(c, "hp")] = fmt.Sprintf("%d/%d", c.HP, c.MaxHP)
			if c.TempHP > 0 {
				st[combatantKey(c, "temp")] = fmt.Sprintf("+%d", c.TempHP)
			}
		}
		if e.IsGM || c.hasMeter() { // the HP dialog's inputs and the HP meter
			st[combatantKey(c, "hpval")] = c.HP
		}
		if e.IsGM {
			st[combatantKey(c, "tempval")] = c.TempHP
		}
	}
}

// Shape is the tracker with its hot values fixed.
func (e EncounterView) Shape() EncounterView {
	e.Encounter.Round = 0
	e.Combatants = slices.Clone(e.Combatants)
	for i := range e.Combatants {
		e.Combatants[i].HP, e.Combatants[i].TempHP, e.Combatants[i].Active, e.Combatants[i].Down = 0, 0, false, false
	}
	return e
}

// Hot adds the roster's hit points and which devices are online.
func (r Roster) Hot(st hub.State) {
	for _, rc := range r.Characters {
		st[fmt.Sprintf("ch.%d.hp", rc.Character.ID)] = rosterHP(rc)
		for _, d := range rc.Devices {
			st[fmt.Sprintf("dev.%d.on", d.SessionID)] = d.Online
			st[fmt.Sprintf("dev.%d.label", d.SessionID)] = deviceLabel(d)
		}
	}
}

// Shape is the roster with its hot values fixed.
func (r Roster) Shape() Roster {
	r.Characters = slices.Clone(r.Characters)
	for i, rc := range r.Characters {
		rc.Character.Hp, rc.Character.MaxHp = 0, 0
		rc.Devices = slices.Clone(rc.Devices)
		for j := range rc.Devices {
			rc.Devices[j].Online, rc.Devices[j].UserAgent = false, ""
		}
		r.Characters[i] = rc
	}
	return r
}

// exhaustion is the one condition with a level, shown as its own chip.
const exhaustion = "exhaustion"

func dying(s SheetView) bool { return s.C.HP <= 0 && s.C.MaxHP > 0 }

func (c CombatantView) down() bool { return c.Down }

func isToggle(slug string) bool { return slices.Contains(dnd5e.Conditions, slug) }

func tempLabel(n int) string { return fmt.Sprintf("+%d temp", n) }

func slotLabel(sl dnd5e.Slot) string { return fmt.Sprintf("%d/%d", sl.Max-sl.Used, sl.Max) }

func roundLabel(n int64) string { return fmt.Sprintf("Round %d", n) }

func inspirationLabel(on bool) string {
	if on {
		return "Inspiration ★"
	}
	return "Inspiration ☆"
}

func combatantKey(c CombatantView, field string) string {
	return fmt.Sprintf("c.%d.%s", c.ID, field)
}

func rosterHP(rc RosterCharacter) string {
	return fmt.Sprintf("HP %d/%d", rc.Character.Hp, rc.Character.MaxHp)
}

func deviceLabel(d Device) string {
	if d.UserAgent != "" {
		return d.UserAgent
	}
	return "offline device"
}

// showsTracker mirrors Encounter: players never see a fight still being
// prepared.
func (e EncounterView) showsTracker() bool {
	return e.IsGM || e.Encounter.Status != "preparing"
}
