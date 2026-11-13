package dnd5e

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	bookStart    = regexp.MustCompile(`(?i)starts with (\w+) level 1 \w+ spells`)
	bookPerLevel = regexp.MustCompile(`(?i)level after 1, add (\w+) \w+ spells`)
)

// Spellbook is how many level 1+ spells the class's spellbook holds at a
// class level ("starts with six level 1 Wizard spells … add two"), or 0 if
// its features describe none.
func Spellbook(info ClassInfo, level int) int {
	count := func(re *regexp.Regexp) int {
		for _, f := range info.Features {
			if m := re.FindStringSubmatch(f.Description); m != nil {
				return CountWord(m[1])
			}
		}
		return 0
	}
	start := count(bookStart)
	if start == 0 {
		return 0
	}
	return start + count(bookPerLevel)*(level-1)
}

// SpellColumn tells whether a class table column is its cantrip total
// ("Cantrips Known") or its level 1+ spell total ("Prepared Spells").
func SpellColumn(name string) (cantrips, spells bool) {
	lk := strings.ToLower(name)
	cantrips = strings.Contains(lk, "cantrip")
	return cantrips, !cantrips && strings.Contains(lk, "spells")
}

// SpellTotals are the class table's cantrip and level 1+ spell totals at a
// class level (-1: no such column); prepared says the spell total is a
// prepared-spells one.
func SpellTotals(info ClassInfo, level int) (cantrips, spells int, prepared bool) {
	cantrips, spells = -1, -1
	for col, v := range info.ColumnsAt(level) {
		n, err := strconv.Atoi(v)
		isCantrips, isSpells := SpellColumn(col)
		switch {
		case err != nil:
		case isCantrips:
			cantrips = n
		case isSpells:
			spells, prepared = n, strings.Contains(strings.ToLower(col), "prepared")
		}
	}
	return cantrips, spells, prepared
}
