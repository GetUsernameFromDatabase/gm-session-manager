// Package views holds the templ pages and components, and the live-view
// keys and bindings they share with the hub.
package views

import (
	"database/sql"
	"fmt"
	"slices"
	"strings"

	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"

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
	LiveImport       = "import"       // GM: the background 5etools import, id = 0
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
// appearing or disappearing means the markup changed shape. The sheet's are
// in character_live.go.

// Hot adds the tracker's often-changing values: the round, whose turn it
// is, and the hit points the viewer may see.
func (e EncounterView) Hot(st hub.State) {
	if !e.showsTracker() {
		return
	}
	if e.Encounter.Status == enum.EncounterActive {
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

// Shape is the tracker with its hot values fixed, hashed as JSON; that
// includes the turn pointer, which the rows' Active show.
func (e EncounterView) Shape() EncounterView {
	e.Encounter.Round, e.Encounter.ActiveCombatantID = 0, sql.NullInt64{}
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

func (c CombatantView) down() bool { return c.Down }

func roundLabel(n int64) string { return fmt.Sprintf("Round %d", n) }

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
	return e.IsGM || e.Encounter.Status != enum.EncounterPreparing
}
