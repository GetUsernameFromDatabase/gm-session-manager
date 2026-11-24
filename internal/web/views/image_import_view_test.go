package views

import "testing"

// TestImportFilterMarkersNeverClash: a CSV category named like a filter marker stays distinct.
func TestImportFilterMarkersNeverClash(t *testing.T) {
	v := ImportView{Cats: []ImportCategory{
		{Source: "a", Target: ImportFilterNone},
		{Source: "b", Target: ImportFilterAll},
		{Source: "c", Target: "Maps"},
	}}
	for _, target := range []string{ImportFilterNone, ImportFilterAll, "Maps"} {
		got := rowFilterCategory(ImportRow{Target: target})
		if got == ImportFilterNone || got == ImportFilterAll {
			t.Errorf("target %q filters as %q, a marker", target, got)
		}
	}
	if got := rowFilterCategory(ImportRow{}); got != ImportFilterNone {
		t.Errorf("no target = %q, want the none marker", got)
	}
	seen := map[string]bool{}
	for _, it := range v.filterItems() {
		if seen[it.Value] {
			t.Errorf("duplicate filter value %q", it.Value)
		}
		seen[it.Value] = true
	}
	if !seen[rowFilterCategory(ImportRow{Target: ImportFilterNone})] {
		t.Error("a row's filter value is not among the select's values")
	}
}
