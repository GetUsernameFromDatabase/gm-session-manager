package compendium

import (
	"reflect"
	"testing"
)

// TestMonsterFormRoundTrip: a statblock survives the homebrew form.
func TestMonsterFormRoundTrip(t *testing.T) {
	d := MonsterData{
		Size: "Large", Type: "fiend", Alignment: "chaotic evil", AC: 16, ACDetail: "natural armor", HP: 90, HitDice: "12d10 + 24", Initiative: 3,
		Speed:     map[string]int{"walk": 40, "fly": 60},
		Abilities: map[string]int{"str": 17, "dex": 16, "con": 15, "int": 5, "wis": 20, "cha": 6},
		Saves:     map[string]int{"dex": 6, "wis": 8},
		Skills:    map[string]int{"perception": 5, "sleight-of-hand": -1},
		CR:        "8", Senses: "Darkvision 60 ft., Passive Perception 15", Languages: "understands Abyssal", Resistance: "cold, fire", Immunity: "frightened",
		Traits: []NameDesc{{Name: "Pack Tactics", Desc: "Advantage when an ally is near."}},
		Actions: []MonsterAction{
			{Name: "Multiattack", Desc: "Two bites.", Type: "action"},
			{Name: "Rending Bite", Desc: "*Melee Attack Roll:* +6.", Type: "action"},
			{Name: "Leap", Desc: "Moves 20 feet.", Type: "bonus_action"},
			{Name: "Snap", Desc: "One bite.", Type: "legendary_action"},
		},
		Desc: "Howls.",
	}
	in := InputForEntry(Entry{Kind: KindMonster, Name: "Howler", Data: d})
	e, err := in.Entry()
	if err != nil {
		t.Fatal(err)
	}
	if got := e.Data.(MonsterData); !reflect.DeepEqual(got, d) {
		t.Errorf("round trip\n got %+v\nwant %+v", got, d)
	}
}

func TestMonsterFormParsers(t *testing.T) {
	speed := []struct {
		in   string
		want map[string]int
	}{
		{"40 ft., fly 60 ft. (hover)", map[string]int{"walk": 40, "fly": 60}},
		{"walk 30, swim 30", map[string]int{"walk": 30, "swim": 30}},
		{"", map[string]int{}},
	}
	for _, tt := range speed {
		if got := parseSpeed(tt.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("parseSpeed(%q) = %v", tt.in, got)
		}
	}
	bonus := []struct {
		in    string
		saves bool
		want  map[string]int
	}{
		{"Dexterity +5, Wis −1", true, map[string]int{"dex": 5, "wis": -1}},
		{"Perception +5, Sleight of Hand +4, junk", false, map[string]int{"perception": 5, "sleight-of-hand": 4}},
		{"Luck +2", true, map[string]int{}},
	}
	for _, tt := range bonus {
		if got := parseBonuses(tt.in, tt.saves); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("parseBonuses(%q) = %v", tt.in, got)
		}
	}
}
