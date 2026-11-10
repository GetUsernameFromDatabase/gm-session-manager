package views

import (
	"reflect"
	"testing"
)

func TestParseWiki(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []WikiPart
	}{
		{"plain", "Lives by the docks", []WikiPart{{Text: "Lives by the docks"}}},
		{"image link", "At [[image:12|The Tavern]].", []WikiPart{{Text: "At "}, {Link: &WikiLink{Kind: "image", ID: 12, Label: "The Tavern"}}, {Text: "."}}},
		{"two links", "[[note:1|Plans]][[entity:2|Mira]]", []WikiPart{{Link: &WikiLink{Kind: "note", ID: 1, Label: "Plans"}}, {Link: &WikiLink{Kind: "entity", ID: 2, Label: "Mira"}}}},
		{"inline image", "[[image:12|Tavern|inline]] by the docks", []WikiPart{{Link: &WikiLink{Kind: "image", ID: 12, Label: "Tavern", Float: "inline"}}, {Text: " by the docks"}}},
		{"compendium link", "see [[compendium:spell/fire-bolt|Fire Bolt]]!", []WikiPart{{Text: "see "}, {Link: &WikiLink{Kind: "compendium", EntryKind: "spell", Slug: "fire-bolt", Label: "Fire Bolt"}}, {Text: "!"}}},
		{"compendium needs a slug", "[[compendium:spell|Fire]]", []WikiPart{{Text: "[[compendium:spell|Fire]]"}}},
		{"unknown float stays text", "[[image:12|Tavern|up]]", []WikiPart{{Text: "[[image:12|Tavern|up]]"}}},
		{"unknown kind stays text", "[[spell:3|Fire]]", []WikiPart{{Text: "[[spell:3|Fire]]"}}},
		{"unclosed stays text", "[[image:3|Tav", []WikiPart{{Text: "[[image:3|Tav"}}},
		{"huge id stays text", "[[image:99999999999999999999|X]]", []WikiPart{{Text: "[[image:99999999999999999999|X]]"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseWiki(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseWiki(%q) = %+v, want %+v", tt.in, got, tt.want)
			}
		})
	}
	if got := WikiLinks("[[image:1|a]] [[image:1|b]] [[note:1|c]]"); len(got) != 2 {
		t.Errorf("WikiLinks repeats targets: %+v", got)
	}
}

func TestReplaceCompendiumLinks(t *testing.T) {
	in := "Cast [[compendium:spell/fire-bolt|Fire Bolt]] at [[compendium:monster/orc|orcs]], see [[note:3|Plans]]."
	if got := CompendiumRefs(in); !reflect.DeepEqual(got, map[string][]string{"spell": {"fire-bolt"}, "monster": {"orc"}}) {
		t.Errorf("CompendiumRefs = %v", got)
	}
	if got := WikiLinks(in); len(got) != 1 || got[0].Kind != "note" {
		t.Errorf("WikiLinks lists compendium links: %+v", got)
	}
	ids := map[string]int64{"compendium:spell/fire-bolt": 7}
	want := "Cast [Fire Bolt](/campaigns/2/compendium/entries/7) at orcs, see [[note:3|Plans]]."
	if got := ReplaceCompendiumLinks(in, ids, CompendiumBase(2)); got != want {
		t.Errorf("ReplaceCompendiumLinks = %q, want %q", got, want)
	}
}
