package dnd5e

// Caster types, as stored on ClassLevel.Caster.
const (
	CasterFull = "full"
	CasterHalf = "half"
	CasterPact = "pact"
	CasterNone = "none"
)

// spellcasterSlots is the SRD 5.2 Multiclass Spellcaster table: spell slots
// per slot level (1st to 9th) by combined caster level. A single-class full
// caster uses the same table with caster level = class level.
var spellcasterSlots = [21][9]int{
	{},
	{2},
	{3},
	{4, 2},
	{4, 3},
	{4, 3, 2},
	{4, 3, 3},
	{4, 3, 3, 1},
	{4, 3, 3, 2},
	{4, 3, 3, 3, 1},
	{4, 3, 3, 3, 2},
	{4, 3, 3, 3, 2, 1},
	{4, 3, 3, 3, 2, 1},
	{4, 3, 3, 3, 2, 1, 1},
	{4, 3, 3, 3, 2, 1, 1},
	{4, 3, 3, 3, 2, 1, 1, 1},
	{4, 3, 3, 3, 2, 1, 1, 1},
	{4, 3, 3, 3, 2, 1, 1, 1, 1},
	{4, 3, 3, 3, 3, 1, 1, 1, 1},
	{4, 3, 3, 3, 3, 2, 1, 1, 1},
	{4, 3, 3, 3, 3, 2, 2, 1, 1},
}

// pactSlots is the Warlock Pact Magic table: slot count and slot level by
// warlock level.
var pactSlots = [21][2]int{
	{}, {1, 1}, {2, 1}, {2, 2}, {2, 2}, {2, 3}, {2, 3}, {2, 4}, {2, 4}, {2, 5}, {2, 5},
	{3, 5}, {3, 5}, {3, 5}, {3, 5}, {3, 5}, {3, 5}, {4, 5}, {4, 5}, {4, 5}, {4, 5},
}

// CasterLevel combines class levels for the spellcaster table: full casters
// count fully and half casters (Paladin, Ranger) count half, rounded up.
func CasterLevel(classes []ClassLevel) int {
	full, half := 0, 0
	for _, c := range classes {
		switch c.Caster {
		case CasterFull:
			full += c.Level
		case CasterHalf:
			half += c.Level
		}
	}
	return min(full+(half+1)/2, 20)
}

// SlotMaxima returns the spell slot maxima and Pact Magic slots for the
// classes. Used slots are clamped by the caller.
func SlotMaxima(classes []ClassLevel) (slots [9]int, pact, pactLevel int) {
	slots = spellcasterSlots[CasterLevel(classes)]
	for _, c := range classes {
		if c.Caster == CasterPact {
			p := pactSlots[min(c.Level, 20)]
			pact, pactLevel = p[0], p[1]
		}
	}
	return slots, pact, pactLevel
}

// RecomputeSlots sets slot maxima from the classes, keeping used counts
// within the new maxima.
func (s *Sheet) RecomputeSlots() {
	maxima, pact, pactLevel := SlotMaxima(s.Classes)
	for i := range s.Slots {
		s.Slots[i].Max = maxima[i]
		s.Slots[i].Used = min(s.Slots[i].Used, maxima[i])
	}
	s.Pact.Max = pact
	s.Pact.Used = min(s.Pact.Used, pact)
	s.PactLevel = pactLevel
}

// MaxSpellLevel is the highest slot level the character can cast.
func (s Sheet) MaxSpellLevel() int {
	best := s.PactLevel
	for i := len(s.Slots) - 1; i >= 0; i-- {
		if s.Slots[i].Max > 0 {
			return max(best, i+1)
		}
	}
	return best
}
