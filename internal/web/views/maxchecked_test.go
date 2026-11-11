package views

import (
	"strings"
	"testing"

	"github.com/a-h/templ"

	"gm-session-manager/internal/compendium"
)

func renderHTML(t *testing.T, c templ.Component) string {
	t.Helper()
	var b strings.Builder
	if err := c.Render(t.Context(), &b); err != nil {
		t.Fatal(err)
	}
	return b.String()
}

// TestChoiceOptionsMaxChecked checks only a group of several picks is capped.
func TestChoiceOptionsMaxChecked(t *testing.T) {
	opts := []compendium.NameDesc{{Name: "Alpha"}, {Name: "Beta"}, {Name: "Gamma"}}
	for _, tc := range []struct {
		count int
		want  string // "" for no attribute
	}{{2, `data-max-checked="2"`}, {1, ""}, {0, ""}} {
		html := renderHTML(t, choiceOptions(ChoiceGroup{Key: "pick", Title: "Pick", Count: tc.count, Options: opts}, nil, nil))
		if tc.want == "" && strings.Contains(html, "data-max-checked") {
			t.Errorf("count %d: unexpected cap in %s", tc.count, html)
		}
		if tc.want != "" && !strings.Contains(html, tc.want) {
			t.Errorf("count %d: no %s in %s", tc.count, tc.want, html)
		}
	}
}

// TestSkillPickMaxChecked checks the class step's skill fieldset is capped.
func TestSkillPickMaxChecked(t *testing.T) {
	v := WizardView{Step: "class", Choice: &WizardChoice{
		Name:      "Fighter",
		SkillPick: &SkillPick{Count: 2, Options: []string{"athletics", "history", "insight"}},
	}}
	if html := renderHTML(t, wizardPick(v)); !strings.Contains(html, `data-max-checked="2"`) {
		t.Errorf("no cap in %s", html)
	}
}

// TestSpellChecklistMaxChecked checks each spell list's cap and counter.
func TestSpellChecklistMaxChecked(t *testing.T) {
	cantrip := Option{Slug: "light", Name: "Light", Level: 0}
	leveled := Option{Slug: "bless", Name: "Bless", Level: 1}
	for _, tc := range []struct {
		name   string
		spells []Option
		need   SpellNeed
		has    []string
		hasNot []string
	}{
		{"both capped", []Option{cantrip, leveled}, SpellNeed{Cantrips: 2, Spells: 1},
			[]string{`id="spell-cantrips" data-max-checked="2"`, `id="spell-leveled" data-max-checked="1"`, `id="spell-cantrips-count"`, `id="spell-leveled-count"`}, nil},
		{"no limit", []Option{cantrip, leveled}, SpellNeed{Cantrips: -1, Spells: -1},
			nil, []string{"data-max-checked", "spell-cantrips-count", "spell-leveled-count"}},
		{"only cantrips offered", []Option{cantrip}, SpellNeed{Cantrips: 2, Spells: 1},
			[]string{`id="spell-cantrips-count"`}, []string{"spell-leveled-count"}},
	} {
		html := renderHTML(t, spellChecklist(tc.spells, nil, tc.need, nil, nil))
		for _, s := range tc.has {
			if !strings.Contains(html, s) {
				t.Errorf("%s: no %s in %s", tc.name, s, html)
			}
		}
		for _, s := range tc.hasNot {
			if strings.Contains(html, s) {
				t.Errorf("%s: unexpected %s in %s", tc.name, s, html)
			}
		}
	}
}
