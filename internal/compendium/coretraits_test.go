package compendium

import (
	"reflect"
	"testing"
)

func TestUnflattenTable(t *testing.T) {
	in := "Shown in the Materials table. Table: Materials | Material | Duration | |----------|-------| | Gems | 10 minutes | | Adamantine or mithral | 1 minute | Then more text."
	want := "Shown in the Materials table.\n\nTable: Materials\n\n| Material | Duration |\n|----------|-------|\n| Gems | 10 minutes |\n| Adamantine or mithral | 1 minute |\n\nThen more text."
	if got := unflattenTable(in); got != want {
		t.Errorf("got %q", got)
	}
	if got := unflattenTable("No table here."); got != "No table here." {
		t.Errorf("got %q", got)
	}
}

func TestCoreTraits(t *testing.T) {
	desc := "|Core Barbarian Traits||\n|---|---|\n|Primary Ability|Strength|\n|Hit Point Die|D12 per Barbarian level|\n"
	got := CoreTraits(desc)
	if got["Primary Ability"] != "Strength" || got["Hit Point Die"] != "D12 per Barbarian level" {
		t.Errorf("%v", got)
	}
}

func TestSkillChoice(t *testing.T) {
	for _, tt := range []struct {
		in   string
		n    int
		list []string
	}{
		{"Choose 2: Animal Handling, Athletics, or Survival", 2, []string{"Animal Handling", "Athletics", "Survival"}},
		{"Choose any 3 skills", 3, nil},
		{"Insight and Religion", 0, nil},
	} {
		n, list := SkillChoice(tt.in)
		if n != tt.n || !reflect.DeepEqual(list, tt.list) {
			t.Errorf("%q: %d %q", tt.in, n, list)
		}
	}
	if got := SplitList("Insight and Religion"); !reflect.DeepEqual(got, []string{"Insight", "Religion"}) {
		t.Errorf("SplitList: %q", got)
	}
}

func TestEquipmentOptions(t *testing.T) {
	got := EquipmentOptions("Choose A, B, or C: (A) Chain Mail, 8 Javelins, Dungeoneer's Pack, and 4 GP; (B) Druidic Focus (Quarterstaff), and 11 GP; or (C) 155 GP")
	want := []EquipmentOption{
		{"A", []EquipmentItem{{"Chain Mail", 1}, {"Javelins", 8}, {"Dungeoneer's Pack", 1}}, 4},
		{"B", []EquipmentItem{{"Druidic Focus (Quarterstaff)", 1}}, 11},
		{"C", nil, 155},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v", got)
	}
	bg := EquipmentOptions("*Choose A or B:* (A) Calligrapher's Supplies, Book (prayers), 8 GP; or (B) 50 GP")
	if len(bg) != 2 || bg[0].GP != 8 || len(bg[0].Items) != 2 || bg[1].GP != 50 {
		t.Errorf("background: %+v", bg)
	}
	if Squash("In sight") != "insight" {
		t.Error("squash")
	}
}
