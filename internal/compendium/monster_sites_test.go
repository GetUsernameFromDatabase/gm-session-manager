package compendium

import (
	"reflect"
	"testing"
)

// The fixtures copy each site's markup or data shape; the text is placeholder.

func TestParseAidedd(t *testing.T) {
	page := []byte(`<html><body><div class="bloc"><div class="jaune">
<h1>Howler</h1><div class="trad">[ <a href="x">Hurleur</a> ]</div>
<div class="type">Large fiend (demon), chaotic evil</div>
<div class="red"><strong>Armor Class</strong> 16 (natural armor)<br><strong>Hit Points</strong> 90 (12d10 + 24)<br><strong>Speed</strong> 40 ft., climb 20 ft.<br>
<div class="carac"><strong>STR</strong><br>17 (+3)</div><div class="carac"><strong>DEX</strong><br>16 (+3)</div><div class="carac"><strong>CON</strong><br>15 (+2)</div>
<div class="carac"><strong>INT</strong><br>5 (-3)</div><div class="carac"><strong>WIS</strong><br>20 (+5)</div><div class="carac"><strong>CHA</strong><br>6 (-2)</div>
<strong>Saving Throws</strong> Dex +6<br><strong>Skills</strong> Perception +5<br><strong>Damage Resistances</strong> cold, fire<br>
<strong>Condition Immunities</strong> frightened<br><strong>Senses</strong> darkvision 60 ft., passive Perception 15<br>
<strong>Languages</strong> understands Abyssal<br><strong>Challenge</strong> 8 (3,900 XP)<br><strong>Proficiency Bonus</strong> +3</div>
<p><strong><em>Pack Tactics</em></strong>. Trait text.</p>
<div class="rub">Actions</div>
<p><strong><em>Multiattack</em></strong>.</p>
<p><strong><em>Rending Bite</em></strong>. Bite text.<br>More bite text.</p>
<div class="rub">Reactions</div>
<p><strong><em>Snap</em></strong>. Snap text.</p>
</div></div>
<div class="picture"><img src="/dnd/images/howler.jpg" alt="Howler"></div></body></html>`)
	e, image, err := ParseAidedd(page)
	if err != nil {
		t.Fatal(err)
	}
	want := MonsterData{
		Size: "Large", Type: "fiend (demon)", Alignment: "chaotic evil", AC: 16, ACDetail: "natural armor", HP: 90, HitDice: "12d10 + 24", Initiative: 3,
		Speed:     map[string]int{"walk": 40, "climb": 20},
		Abilities: map[string]int{"str": 17, "dex": 16, "con": 15, "int": 5, "wis": 20, "cha": 6},
		Saves:     map[string]int{"dex": 6}, Skills: map[string]int{"perception": 5},
		CR: "8", Senses: "darkvision 60 ft., passive Perception 15", Languages: "understands Abyssal", Resistance: "cold, fire", Immunity: "frightened",
		Traits: []NameDesc{{Name: "Pack Tactics", Desc: "Trait text."}},
		Actions: []MonsterAction{
			{Name: "Multiattack", Type: "action"},
			{Name: "Rending Bite", Desc: "Bite text.\nMore bite text.", Type: "action"},
			{Name: "Snap", Desc: "Snap text.", Type: "reaction"},
		},
	}
	if e.Name != "Howler" || e.Kind != KindMonster || !reflect.DeepEqual(e.Data, want) {
		t.Errorf("%s\n got %+v\nwant %+v", e.Name, e.Data, want)
	}
	if image != "https://www.aidedd.org/dnd/images/howler.jpg" {
		t.Errorf("image %q", image)
	}
}

func TestParseFiveTools(t *testing.T) {
	file := []byte(`{"monster":[
{"name":"Other","source":"MTF","size":["M"],"type":"beast"},
{"name":"Howler","source":"MTF","page":210,"size":["L"],"type":{"type":"fiend","tags":["demon"]},"alignment":["C","E"],
 "ac":[{"ac":16,"from":["natural armor"]}],"hp":{"average":90,"formula":"12d10 + 24"},"speed":{"walk":40,"fly":{"number":30,"condition":"(hover)"},"canHover":true},
 "str":17,"dex":16,"con":15,"int":5,"wis":20,"cha":6,"save":{"dex":"+6"},"skill":{"perception":"+8","sleight of hand":"+2"},
 "senses":["darkvision 60 ft."],"passive":15,"languages":["understands Abyssal but can't speak"],"cr":{"cr":"8","lair":"9"},
 "resist":["cold","fire",{"resist":["bludgeoning","piercing"],"note":"from nonmagical attacks","cond":true}],"conditionImmune":["frightened"],
 "trait":[{"name":"Pack Tactics","entries":["Has advantage unless the ally is {@condition incapacitated}."]}],
 "spellcasting":[{"name":"Innate Spellcasting","headerEntries":["Casts:"],"will":["{@spell darkness}"],"daily":{"1e":["{@spell fear}","{@spell sleep|phb}"]}}],
 "action":[{"name":"Rending Bite","entries":["{@atk mw} {@hit 6} to hit, reach 5 ft. {@h}13 ({@damage 2d8 + 4}) piercing damage."]},
  {"name":"Howl {@recharge 4}","entries":["{@actSave wis} {@dc 16}.",{"type":"list","items":["One.","Two."]}]}],
 "legendary":[{"name":"Lunge","entries":["Moves."]}],
 "hasFluffImages":true},
{"name":"Copycat","source":"MTF","_copy":{"name":"Howler","source":"MTF"}}
]}`)
	book, image, e, err := ParseFiveTools(file, "howler")
	if err != nil {
		t.Fatal(err)
	}
	want := MonsterData{
		Size: "Large", Type: "fiend (demon)", Alignment: "chaotic evil", AC: 16, ACDetail: "natural armor", HP: 90, HitDice: "12d10 + 24", Initiative: 3,
		Speed:     map[string]int{"walk": 40, "fly": 30},
		Abilities: map[string]int{"str": 17, "dex": 16, "con": 15, "int": 5, "wis": 20, "cha": 6},
		Saves:     map[string]int{"dex": 6}, Skills: map[string]int{"perception": 8, "sleight-of-hand": 2},
		CR: "8", Senses: "darkvision 60 ft., passive Perception 15", Languages: "understands Abyssal but can't speak",
		Resistance: "cold, fire; bludgeoning, piercing from nonmagical attacks", Immunity: "frightened",
		Traits: []NameDesc{
			{Name: "Pack Tactics", Desc: "Has advantage unless the ally is incapacitated."},
			{Name: "Innate Spellcasting", Desc: "Casts:\n\nAt will: darkness\n\n1/day each: fear, sleep"},
		},
		Actions: []MonsterAction{
			{Name: "Rending Bite", Desc: "Melee Weapon Attack: +6 to hit, reach 5 ft. Hit: 13 (2d8 + 4) piercing damage.", Type: "action"},
			{Name: "Howl (Recharge 4–6)", Desc: "Wisdom Saving Throw: DC 16.\n\n- One.\n- Two.", Type: "action"},
			{Name: "Lunge", Desc: "Moves.", Type: "legendary_action"},
		},
	}
	if book != "MTF p. 210" || e.Name != "Howler" || !reflect.DeepEqual(e.Data, want) {
		t.Errorf("%s %s\n got %+v\nwant %+v", book, e.Name, e.Data, want)
	}
	if image != fiveToolsImages+"MTF/Howler.webp" {
		t.Errorf("image %q", image)
	}
	for _, slug := range []string{"copycat", "nobody"} {
		if _, _, _, err := ParseFiveTools(file, slug); err == nil {
			t.Errorf("%s: no error", slug)
		}
	}
}
