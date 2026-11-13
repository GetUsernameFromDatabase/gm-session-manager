package views

import (
	"fmt"
	"slices"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/hub"
)

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
		if c, on := s.condition(slug); slug != dnd5e.Exhaustion { // its level is its own chip
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

// Shape is the sheet with its hot values fixed, hashed as JSON: the sheet's
// markup must read nothing outside SheetView.
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

func dying(s SheetView) bool { return s.C.HP <= 0 && s.C.MaxHP > 0 }

func isToggle(slug string) bool { return slices.Contains(dnd5e.Conditions, slug) }

func tempLabel(n int) string { return fmt.Sprintf("+%d temp", n) }

func slotLabel(sl dnd5e.Slot) string { return fmt.Sprintf("%d/%d", sl.Max-sl.Used, sl.Max) }

func inspirationLabel(on bool) string {
	if on {
		return "Inspiration ★"
	}
	return "Inspiration ☆"
}
