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
