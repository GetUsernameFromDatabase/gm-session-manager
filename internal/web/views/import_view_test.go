package views

import (
	"slices"
	"testing"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
)

// TestSourcePositions: homebrew is 0 wherever it sits; the others count on
// from 1 in list order.
func TestSourcePositions(t *testing.T) {
	var sources []sqlc.CompendiumSource
	for _, k := range []string{"book-a", compendium.HomebrewSource, "book-b", "book-c"} {
		sources = append(sources, sqlc.CompendiumSource{SourceKey: k})
	}
	if got, want := sourcePositions(sources), []int{1, 0, 2, 3}; !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// TestEntryCount: the chip reads shown/total, or the total when all show;
// red for none; its title spells it out (bands: TestEntryCountClass).
func TestEntryCount(t *testing.T) {
	counts := map[string]sqlc.CountEntriesBySourceRow{
		"all":   {Total: 3, Visible: 3},
		"some":  {Total: 215, Visible: 147},
		"none":  {Total: 3, Visible: 0},
		"empty": {},
		"one":   {Total: 1, Visible: 1},
	}
	text := map[string]string{
		"all":   "3",
		"some":  "147/215",
		"none":  "0/3",
		"empty": "0",
		"one":   "1",
	}
	title := map[string]string{
		"all":   "3 entries, all shown.",
		"some":  "215 entries, 147 shown.",
		"none":  "3 entries, 0 shown.",
		"empty": "0 entries.",
		"one":   "1 entry, all shown.",
	}
	colour := map[string]string{"none": countNoneShown}
	for name, c := range counts {
		if got := entryCountText(c); got != text[name] {
			t.Errorf("%s: text %q, want %q", name, got, text[name])
		}
	}
	for name, c := range counts {
		if got := entryCountTitle(c, sqlc.CountSharedBySourceRow{}, false); got != title[name] {
			t.Errorf("%s: title %q, want %q", name, got, title[name])
		}
	}
	shared := sqlc.CountSharedBySourceRow{Shared: 68, Shown: 0}
	if got, want := entryCountTitle(
		counts["some"],
		shared,
		true,
	), "215 entries, 147 shown. 68 shared with other sources, 0 of those win."; got != want {
		t.Errorf("shared title %q, want %q", got, want)
	}
	for name, c := range counts {
		if got, want := entryCountClass(c), colour[name]; got != want {
			t.Errorf("%s: class %q, want %q", name, got, want)
		}
	}
}

// TestEntryCountClass: muted from half shown up, then amber, orange, red,
// destructive as the shown share falls; an empty source stays muted.
func TestEntryCountClass(t *testing.T) {
	bands := map[sqlc.CountEntriesBySourceRow]string{
		{Total: 0, Visible: 0}:     "",
		{Total: 957, Visible: 956}: "",
		{Total: 100, Visible: 50}:  "",
		{Total: 100, Visible: 49}:  countAmber,
		{Total: 100, Visible: 35}:  countAmber,
		{Total: 100, Visible: 34}:  countOrange,
		{Total: 100, Visible: 20}:  countOrange,
		{Total: 100, Visible: 19}:  countRed,
		{Total: 100, Visible: 1}:   countRed,
		{Total: 100, Visible: 0}:   countNoneShown,
	}
	for c, want := range bands {
		if got := entryCountClass(c); got != want {
			t.Errorf("%d/%d: class %q, want %q", c.Visible, c.Total, got, want)
		}
	}
}

func TestLostDot(t *testing.T) {
	for _, tt := range []struct {
		name           string
		total, visible int64
		pct            int
		lost           bool
		class          string
	}{
		{"nothing lost", 100, 100, 0, false, ""},
		{"empty", 0, 0, 0, false, ""},
		{"one of 400", 400, 399, minLostPercent, true, countAmber},
		{"sixty shown", 100, 60, 40, true, countAmber},
		{"forty-five shown", 100, 45, 55, true, countAmber},
		{"thirty shown", 100, 30, 70, true, countOrange},
		{"ten shown", 100, 10, 90, true, countRed},
		{"none shown", 100, 0, 100, true, countNoneShown},
	} {
		t.Run(tt.name, func(t *testing.T) {
			c := sqlc.CountEntriesBySourceRow{Total: tt.total, Visible: tt.visible}
			pct, lost := lostPercent(c)
			if pct != tt.pct || lost != tt.lost {
				t.Errorf("lostPercent %d, %v; want %d, %v", pct, lost, tt.pct, tt.lost)
			}
			if lost && lostDotClass(c) != tt.class {
				t.Errorf("class %q, want %q", lostDotClass(c), tt.class)
			}
		})
	}
}
