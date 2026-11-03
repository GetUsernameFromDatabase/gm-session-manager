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

func TestOptions(t *testing.T) {
	for _, tt := range []struct {
		name, in string
		want     []NameDesc
	}{
		{"table", "Choose a lineage.\n\nTable: Elven Lineages\n\n|Lineage|Level 1|Level 3|\n|---|---|---|\n|Drow|You know the Dancing Lights cantrip.|Faerie Fire|\n|High Elf|You know the Prestidigitation cantrip.|Detect Magic|\n\n|Other|Table|\n|---|---|\n|Not|This|",
			[]NameDesc{{Name: "Drow", Desc: "You know the Dancing Lights cantrip."}, {Name: "High Elf", Desc: "You know the Prestidigitation cantrip."}}},
		{"paragraphs", "Choose one of the following options:\n\n**Forest Gnome.** You know the *Minor Illusion* cantrip.\n\n**Rock Gnome.** You know the *Mending* and *Prestidigitation* cantrips.",
			[]NameDesc{{Name: "Forest Gnome", Desc: "You know the *Minor Illusion* cantrip."}, {Name: "Rock Gnome", Desc: "You know the *Mending* and *Prestidigitation* cantrips."}}},
		{"bullets", "Choose one of the following benefits:\n\n- **Cloud's Jaunt (Cloud Giant)**. You teleport.\n- **Fire's Burn (Fire Giant)**. You burn.",
			[]NameDesc{{Name: "Cloud's Jaunt (Cloud Giant)", Desc: "You teleport."}, {Name: "Fire's Burn (Fire Giant)", Desc: "You burn."}}},
		{"sections", "In alphabetical order.\n\n### Agonizing Blast\n\n*Prerequisite: Level 2+ Warlock*\n\n**Repeatable.** Again.\n\n### **Armor of Shadows**\n\nMage Armor.",
			[]NameDesc{{Name: "Agonizing Blast", Desc: "*Prerequisite: Level 2+ Warlock*\n\n**Repeatable.** Again."}, {Name: "Armor of Shadows", Desc: "Mage Armor."}}},
		{"none", "You have proficiency in the Insight, Perception, or Survival skill.", nil},
	} {
		if got := Options(tt.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: %+v", tt.name, got)
		}
	}
	lineage := "Choose a lineage.\n\n**Table: Lineages**\n\n|Lineage|Level 1|Level 3|Level 5|\n|---|---|---|---|\n|Drow|Dancing Lights.|Faerie Fire|Darkness|\n|High Elf|Prestidigitation.|Detect Magic||\n\n|Other|Level 3|\n|---|---|\n|Not|This|"
	if got, want := LevelSpells(lineage), map[string]map[int]string{"Drow": {3: "Faerie Fire", 5: "Darkness"}, "High Elf": {3: "Detect Magic"}}; !reflect.DeepEqual(got, want) {
		t.Errorf("LevelSpells: %v", got)
	}
	if got := LevelSpells("**Forest Gnome.** You know the *Minor Illusion* cantrip."); len(got) != 0 {
		t.Errorf("LevelSpells without a table: %v", got)
	}
	if OptionLevel("*Prerequisite: Level 5+ Warlock*") != 5 || OptionLevel("Mage Armor.") != 0 {
		t.Error("OptionLevel")
	}
	for in, want := range map[string][]string{
		"You also know the Dancing Lights cantrip.":                  {"Dancing Lights"},
		"You know the *Mending* and *Prestidigitation* cantrips. X.": {"Mending", "Prestidigitation"},
		"You have Resistance to Fire damage.":                        nil,
	} {
		if got := Cantrips(in); !reflect.DeepEqual(got, want) {
			t.Errorf("Cantrips(%q) = %q", in, got)
		}
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
