package web

import (
	"slices"
	"testing"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
)

func TestBooksUsed(t *testing.T) {
	sources := []sqlc.CompendiumSource{
		{SourceKey: "homebrew"}, {SourceKey: "xphb", Edition: "2024"}, {SourceKey: "tce", Edition: "2014"}, {SourceKey: "xge", Edition: "2014"},
	}
	pins := []dnd5e.Pin{
		{Source: "xphb", Kind: "class", Name: "Fighter"},
		{Source: "tce", Kind: "spell", Name: "Shadow Blade"},
		{Source: "xphb", Kind: "spell", Name: "Shield"},
		{Source: "xphb", Kind: "spell", Name: "Shield"}, // used twice, listed once
		{Source: "xge", Kind: "subclass", Name: "Rune Knight"},
	}
	used, campaign := booksUsed(pins, sources, []string{"xphb", "xge"}, []string{"xge"}, false)
	if len(used) != 3 {
		t.Fatalf("books %+v", used)
	}
	for i, want := range []struct {
		key              string
		uses             []string
		notAllowed, home bool
	}{
		{"xphb", []string{"class Fighter", "spell Shield"}, false, false},
		{"tce", []string{"spell Shadow Blade"}, true, false},
		{"xge", []string{"subclass Rune Knight"}, false, true},
	} {
		b := used[i]
		if b.Key != want.key || !slices.Equal(b.Uses, want.uses) || b.NotAllowed != want.notAllowed || b.AtHome != want.home {
			t.Errorf("book %d: %+v, want %+v", i, b, want)
		}
	}
	if used[0].Label != "xphb, 2024" {
		t.Errorf("label %q", used[0].Label)
	}
	if !slices.Equal(campaign, []string{"xphb, 2024"}) {
		t.Errorf("campaign books %v: homebrew, denied and home books stay out, priority order", campaign)
	}
	// A home book the campaign does not allow gets only the home warning.
	if used, _ := booksUsed(pins, sources, []string{"xphb"}, []string{"xge"}, false); used[2].NotAllowed || !used[2].AtHome {
		t.Errorf("home book %+v", used[2])
	}
	// No campaign allows everything.
	if used, _ := booksUsed(pins, sources, nil, nil, true); used[1].NotAllowed {
		t.Error("no campaign: nothing is disallowed")
	}
}
