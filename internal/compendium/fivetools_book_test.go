package compendium

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/db/sqlc"
)

// A mastery is a "Name|Source" code or an object with a uid (Psychic Blade).
func TestFiveToolsItemMastery(t *testing.T) {
	for _, tt := range []struct{ json, want string }{
		{`{"name":"Longsword","mastery":["Sap|XPHB"]}`, "Sap"},
		{`{"name":"Psychic Blade","mastery":[{"uid":"Vex|XPHB","note":"free"}]}`, "Vex"},
		{`{"name":"Odd","mastery":["Nick|XPHB",{"note":"no uid"}]}`, "Nick"},
		{`{"name":"Club"}`, ""},
	} {
		var it ftItem
		if err := json.Unmarshal([]byte(tt.json), &it); err != nil {
			t.Fatalf("%s: %v", tt.json, err)
		}
		if got := ftItemData(it).Mastery; got != tt.want {
			t.Errorf("%s: mastery %q, want %q", tt.json, got, tt.want)
		}
	}
}

func TestFiveToolsMagicVariants(t *testing.T) {
	files := map[string][]byte{
		"book.json":  []byte(`{"id":"DMG","name":"DMG"}`),
		"items.json": []byte(`{"item":[{"name":"Vorpal Sword","source":"DMG","rarity":"legendary","entries":["Real."]}]}`),
		"magicvariants.json": []byte(`{"magicvariant":[
			{"name":"+1 Weapon","requires":[{"weapon":true}],"inherits":{"source":"DMG","rarity":"uncommon","bonusWeapon":"+1","entries":["You have a {=bonusWeapon} bonus to attack rolls."]}},
			{"name":"Flame Tongue","requires":[{"sword":true}],"inherits":{"source":"DMG","rarity":"rare","reqAttune":true,"entries":["Aflame."]}},
			{"name":"Vorpal Sword","requires":[{"sword":true}],"inherits":{"source":"DMG","entries":["Generic."]}}]}`),
	}
	_, entries, _, err := parseFiveToolsBookReport("k", files)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]ItemData{}
	for _, e := range entries {
		got[e.Slug] = e.Data.(ItemData)
	}
	for _, tc := range []struct {
		slug, rarity, desc string
		attune             bool
	}{
		{"1-weapon", "Uncommon", "Applies to: weapon\n\nYou have a +1 bonus to attack rolls.", false},
		{"flame-tongue", "Rare", "Applies to: sword\n\nAflame.", true},
		{"vorpal-sword", "Legendary", "Real.", false}, // the book's own item wins
	} {
		d := got[tc.slug]
		if d.Rarity != tc.rarity || d.Desc != tc.desc || d.Attunement != tc.attune || d.Category != "other" {
			t.Errorf("%s = %+v, want rarity %q, desc %q, attunement %v", tc.slug, d, tc.rarity, tc.desc, tc.attune)
		}
	}
}

// fiveToolsFixture sits below its own folder so the web tests' import of
// testdata does not pick it up.
const fiveToolsFixture = "testdata/fivetools/5etools-xphb"

func TestImportFiveToolsBook(t *testing.T) {
	store := newStore(t)
	r := importOnce(t, store, filepath.Dir(fiveToolsFixture))
	if r.Entries == 0 {
		t.Fatalf("report %+v", r)
	}
	cat, err := store.GetCompendiumEntryBySlug(t.Context(), sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "cat"})
	if err != nil {
		t.Fatal(err)
	}
	if want := FiveToolsImages(readBook(t, fiveToolsFixture))["cat"]; want == "" || cat.ImageUrl.String != want {
		t.Errorf("cat image %q, want %q", cat.ImageUrl.String, want)
	}
	src, err := store.GetCompendiumSource(t.Context(), cat.SourceKey)
	if err != nil || src.Edition != "2024" || src.Published != "2024-09-17" || !strings.Contains(src.Attribution, "Wizards of the Coast") {
		t.Errorf("source %+v %v", src, err)
	}
	// The cat's environment seeds its habitats; a GM change survives re-import.
	habitats := func() []string {
		h, _ := store.ListEntryHabitats(t.Context(), cat.ID)
		slices.Sort(h)
		return h
	}
	if got := habitats(); !reflect.DeepEqual(got, []string{"Planar", "Urban"}) {
		t.Errorf("cat habitats %v", got)
	}
	if err := store.RemoveMonsterHabitat(t.Context(), sqlc.RemoveMonsterHabitatParams{EntryID: cat.ID, Habitat: "Urban"}); err != nil {
		t.Fatal(err)
	}
	// A changed hash forces the re-import.
	if err := store.UpsertCompendiumSource(t.Context(), sqlc.UpsertCompendiumSourceParams{SourceKey: src.SourceKey, Path: src.Path, Edition: src.Edition}); err != nil {
		t.Fatal(err)
	}
	if r := importOnce(t, store, filepath.Dir(fiveToolsFixture)); r.Unchanged {
		t.Fatal("re-import skipped")
	}
	if got := habitats(); !reflect.DeepEqual(got, []string{"Planar"}) {
		t.Errorf("cat habitats after re-import %v", got)
	}
	// Pickers offer base species, never versions or subraces.
	names, err := store.ListCompendiumNames(t.Context(), sqlc.ListCompendiumNamesParams{Kind: string(KindSpecies)})
	if err != nil {
		t.Fatal(err)
	}
	var slugs []string
	for _, n := range names {
		slugs = append(slugs, n.Slug)
	}
	if want := []string{"dragonborn", "elf", "elf-2014", "half-elf"}; !reflect.DeepEqual(slugs, want) {
		t.Errorf("species picks %v, want %v", slugs, want)
	}
}

// Sheets, combat and the homebrew form read entries without the entry page's
// link resolution, so they get the labels, never the tokens.
func TestTokensReadAsLabels(t *testing.T) {
	raw := `{"desc":"The target is [[compendium:condition/prone|Prone]].","actions":[{"name":"Shove","desc":"It falls [[compendium:condition/prone|prone]].","type":"action"}]}`
	d, err := DecodeData(KindMonster, raw)
	if err != nil {
		t.Fatal(err)
	}
	m := d.(*MonsterData)
	if m.Desc != "The target is Prone." || m.Actions[0].Desc != "It falls prone." {
		t.Errorf("decoded %q / %q", m.Desc, m.Actions[0].Desc)
	}
	in := InputForEntry(Entry{Kind: KindMonster, Name: "Ogre", Data: *m})
	if strings.Contains(in.Actions, "[[") || !strings.Contains(in.Actions, "It falls prone.") {
		t.Errorf("form actions %q", in.Actions)
	}
	in = InputForEntry(Entry{Kind: KindMonster, Name: "Ogre", Data: MonsterData{Desc: "See [[compendium:monster/goblin|goblins]]."}})
	if in.Desc != "See goblins." {
		t.Errorf("form desc %q", in.Desc)
	}
}

func TestStripTagsTokens(t *testing.T) {
	for _, tt := range []struct{ in, want, plain string }{
		{"{@spell Fireball|XPHB}", "[[compendium:spell/fireball|Fireball]]", "Fireball"},
		{"{@spell fireball|xphb|a ball of fire}", "[[compendium:spell/fireball|a ball of fire]]", "a ball of fire"},
		{"{@item Bag of Holding|XPHB|bags}", "[[compendium:item/bag-of-holding|bags]]", "bags"},
		{"{@creature goblin warrior|XMM}", "[[compendium:monster/goblin-warrior|goblin warrior]]", "goblin warrior"},
		{"{@condition Prone|XPHB}", "[[compendium:condition/prone|Prone]]", "Prone"},
		{"{@feat Magic Initiate|XPHB|Magic Initiate (Cleric)}", "[[compendium:feat/magic-initiate|Magic Initiate (Cleric)]]", "Magic Initiate (Cleric)"},
		{"{@background Acolyte|XPHB}", "[[compendium:background/acolyte|Acolyte]]", "Acolyte"},
		{"{@race Elf|XPHB}", "[[compendium:species/elf|Elf]]", "Elf"},
		{"{@spell Fireball||Big}", "[[compendium:spell/fireball|Big]]", "Big"},
		{"{@itemMastery Sap|XPHB}", "[[compendium:rule/sap|Sap]]", "Sap"},
		// Other tags stay plain text.
		{"{@skill Insight|XPHB} {@damage 8d6} {@dc 15}", "Insight 8d6 DC 15", "Insight 8d6 DC 15"},
		{"{@b {@spell Shield|XPHB}}", "[[compendium:spell/shield|Shield]]", "Shield"},
	} {
		got := stripTags(tt.in)
		if got != tt.want || plainText(got) != tt.plain {
			t.Errorf("%s: %q / %q, want %q / %q", tt.in, got, plainText(got), tt.want, tt.plain)
		}
	}
}

func TestHabitatsFrom(t *testing.T) {
	for _, tt := range []struct {
		in   []string
		want []string
	}{
		{nil, nil},
		{[]string{"urban", "forest"}, []string{"Forest", "Urban"}},
		{[]string{"planar, nine hells", "planar, feywild"}, []string{"Planar"}},
		{[]string{"Underdark", "moon"}, []string{"Underdark"}},
		{[]string{"any"}, Habitats},
	} {
		if got := habitatsFrom(tt.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v: %v, want %v", tt.in, got, tt.want)
		}
	}
}

func readBook(t *testing.T, dir string) map[string][]byte {
	t.Helper()
	paths, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil || len(paths) == 0 {
		t.Fatal(dir, err)
	}
	files := map[string][]byte{}
	for _, p := range paths {
		b, err := os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		files[filepath.Base(p)] = b
	}
	return files
}

func TestParseFiveToolsBook(t *testing.T) {
	files := readBook(t, fiveToolsFixture)
	src, entries, skipped, err := parseFiveToolsBookReport("5etools-xphb", files)
	if err != nil {
		t.Fatal(err)
	}
	want := Source{Key: "5etools-xphb", Edition: "2024", License: "Not SRD — private use", Attribution: "Player's Handbook (2024), © Wizards of the Coast", Published: "2024-09-17"}
	if src != want {
		t.Errorf("source %+v", src)
	}
	// Owl Spirit copies a monster of another book, as in a folder downloaded
	// before _copyTargets (TestFiveToolsCopyTargets covers the bundled case).
	if skipped != 1 {
		t.Errorf("skipped %d, want 1", skipped)
	}
	got := map[string]Entry{}
	for _, e := range entries {
		got[string(e.Kind)+"/"+e.Slug] = e
		if strings.Contains(e.Text, "[[") {
			t.Errorf("%s/%s: token in search text %q", e.Kind, e.Slug, e.Text)
		}
	}
	keys := []string{
		"monster/cat", "monster/mountain-cat", "spell/fireball", "spell/detect-magic", "spell/shield",
		"item/longsword", "item/dagger", "item/plate-armor", "item/dwarven-plate", "item/bag-of-holding",
		"item/flame-tongue-longsword", "item/absorbing-tattoo", "species/elf", "species/elf-drow",
		"species/half-elf", "species/elf-2014", "species/half-elf-drow-descent", "species/half-elf-aquatic-elf-descent", "species/elf-high",
		"species/elf-wood-elf-lineage", "species/dragonborn", "species/dragonborn-black", "species/dragonborn-red",
		"background/acolyte", "feat/alert", "feat/grappler", "condition/prone", "condition/concentration",
		"rule/sewer-plague", "rule/advantage", "rule/sap", "class/fighter", "class/wizard", "subclass/champion", "subclass/evoker",
	}
	if len(entries) != len(keys) {
		t.Errorf("%d entries, want %d", len(entries), len(keys))
	}
	for _, k := range keys {
		if _, ok := got[k]; !ok {
			t.Errorf("no %s", k)
		}
	}

	for _, tt := range []struct {
		key  string
		data any
	}{
		{"monster/mountain-cat", MonsterData{
			Size: "Medium", Type: "beast", Alignment: "unaligned", AC: 12, HP: 11, HitDice: "2d8 + 2", Initiative: 2,
			Speed: map[string]int{"walk": 40, "climb": 40}, Saves: map[string]int{}, Skills: map[string]int{"perception": 3, "stealth": 4},
			Abilities: map[string]int{"str": 3, "dex": 14, "con": 10, "int": 3, "wis": 12, "cha": 7},
			CR:        "0", Senses: "darkvision 60 ft., passive Perception 13",
			Actions: []MonsterAction{
				{Name: "Scratch", Desc: "Melee Attack Roll: +4, reach 5 ft. Hit: 1 slashing damage. The target is [[compendium:condition/grappled|grappled]] by the mountain cat.", Type: "action"},
				{Name: "Pounce", Desc: "It casts [[compendium:spell/jump|a leap]].", Type: "action"},
			},
		}},
		{"spell/fireball", SpellData{
			Level: 3, School: "evocation", CastingTime: "Action", Range: "150 feet",
			Components: "V, S, M (a ball of bat guano and sulfur)", Duration: "Instantaneous", Classes: []string{"sorcerer", "wizard"},
			Desc:        "A bright streak explodes in a 20-foot-radius Sphere. Each creature makes a Dexterity saving throw, taking 8d6 Fire damage on a failed save.\n\n| d6 | Effect |\n|---|---|\n| 1–5 | Burns |\n| 6 | Ignites |",
			HigherLevel: "The damage increases by 1d6 for each slot level above 3.",
		}},
		{"spell/detect-magic", SpellData{
			Level: 1, School: "divination", CastingTime: "Action", Range: "Self (30-foot emanation)", Components: "V, S",
			Duration: "up to 10 minutes", Concentration: true, Ritual: true, Desc: "You sense magic within 30 feet.",
		}},
		{"spell/shield", SpellData{
			Level: 1, School: "abjuration", Range: "Self", Components: "V, S", Duration: "1 round", Classes: []string{"sorcerer", "wizard"},
			CastingTime: "Reaction, which you take when you are hit by an attack roll or targeted by the Magic Missile spell",
			Desc:        "An imperceptible barrier of magical force protects you.\n\nA favorite of cautious wizards.",
		}},
		{"item/dagger", ItemData{
			Category: "weapon", Cost: "2", Weight: "1", Damage: "1d4 piercing", Properties: "Finesse, Light, Thrown", Range: "20/60",
		}},
		{"item/flame-tongue-longsword", ItemData{
			Category: "weapon", Rarity: "Rare", Attunement: true, Weight: "3", Damage: "1d8 slashing", Properties: "Versatile (1d10)", Martial: true,
			Mastery: "Sap", Desc: "You can use a Bonus Action to set the blade aflame.",
		}},
		{"item/longsword", ItemData{
			Category: "weapon", Cost: "15", Weight: "3", Damage: "1d8 slashing", Properties: "Versatile (1d10)", Martial: true, Mastery: "Sap",
		}},
		{"rule/sap", TextData{Desc: "If you hit a creature with this weapon, that creature has Disadvantage on its next attack roll.", Category: "Weapon Mastery"}},
		// _versions: the base's traits with the version's _mod, its fluff, and
		// a template filled per implementation.
		{"species/elf-wood-elf-lineage", TextData{Desc: "Elves live long lives.", Parts: []NameDesc{
			{Name: "Size", Desc: "Medium", Type: partSize}, {Name: "Speed", Desc: "35 feet", Type: partSpeed},
			{Name: "Darkvision", Desc: "You have Darkvision with a range of 60 feet."},
			{Name: "Elven Lineage (Wood Elf)", Desc: "You know the [[compendium:spell/druidcraft|Druidcraft]] cantrip."},
		}, VersionOf: "Elf"}},
		{"species/dragonborn-red", TextData{Parts: []NameDesc{
			{Name: "Size", Desc: "Medium", Type: partSize}, {Name: "Speed", Desc: "30 feet", Type: partSpeed},
			{Name: "Damage Resistance", Desc: "You have Resistance to Fire damage."},
		}, VersionOf: "Dragonborn"}},
		{"item/dwarven-plate", ItemData{
			Category: "armor", Rarity: "Very rare", Weight: "65", ArmorClass: "18, Stealth disadvantage",
			Armor: &ArmorData{Type: ArmorHeavy, Base: 18, StrReq: 15, Stealth: true},
			Desc:  "While wearing this armor, you gain a +2 bonus to Armor Class.",
		}},
		{"item/bag-of-holding", ItemData{
			Category: "wondrous-item", Rarity: "Uncommon", Weight: "5",
			Desc: "This bag has an interior space larger than its outside, like [[compendium:spell/rope-trick|Rope Trick]].",
		}},
		{"species/elf-drow", TextData{Desc: "Elves live long lives.", Parts: []NameDesc{
			{Name: "Size", Desc: "Medium", Type: partSize}, {Name: "Speed", Desc: "30 feet", Type: partSpeed},
			{Name: "Darkvision", Desc: "You have Darkvision with a range of 60 feet."}, {Name: "Trance", Desc: "You don't need to sleep."},
			{Name: "Drow Lineage", Desc: "You know the [[compendium:spell/dancing-lights|Dancing Lights]] cantrip."},
		}, SubraceOf: "elf"}},
		// Parents from another book come bundled; a reprinted one is "(2014)".
		{"species/half-elf", TextData{Parts: []NameDesc{
			{Name: "Size", Desc: "Medium", Type: partSize}, {Name: "Speed", Desc: "30 feet", Type: partSpeed},
			{Name: "Ability Score Increase", Desc: "Charisma +2, +1 to 2 of Strength, Dexterity, Constitution, Intelligence, Wisdom"},
			{Name: "Fey Ancestry", Desc: "You resist charm."}, {Name: "Skill Versatility", Desc: "You gain two skills."},
		}}},
		{"species/elf-2014", TextData{Parts: []NameDesc{
			{Name: "Size", Desc: "Medium", Type: partSize}, {Name: "Speed", Desc: "30 feet", Type: partSpeed},
			{Name: "Ability Score Increase", Desc: "Dexterity +2"}, {Name: "Keen Senses", Desc: "You have Perception."},
		}}},
		// A subrace is complete: an overwritten trait is replaced in place.
		{"species/half-elf-drow-descent", TextData{Parts: []NameDesc{
			{Name: "Size", Desc: "Medium", Type: partSize}, {Name: "Speed", Desc: "30 feet", Type: partSpeed},
			{Name: "Ability Score Increase", Desc: "Charisma +2, +1 to 2 of Strength, Dexterity, Constitution, Intelligence, Wisdom"},
			{Name: "Fey Ancestry", Desc: "You resist charm."}, {Name: "Skill Versatility", Desc: "You know Dancing Lights."},
		}, SubraceOf: "half-elf"}},
		{"species/half-elf-aquatic-elf-descent", TextData{Parts: []NameDesc{
			{Name: "Size", Desc: "Medium", Type: partSize}, {Name: "Speed", Desc: "30 feet", Type: partSpeed},
			{Name: "Ability Score Increase", Desc: "Wisdom +1"},
			{Name: "Fey Ancestry", Desc: "You resist charm."}, {Name: "Skill Versatility", Desc: "You gain two skills."}, {Name: "Swim", Desc: "You swim."},
		}, SubraceOf: "half-elf"}},
		{"species/elf-high", TextData{Parts: []NameDesc{
			{Name: "Size", Desc: "Medium", Type: partSize}, {Name: "Speed", Desc: "30 feet", Type: partSpeed},
			{Name: "Ability Score Increase", Desc: "Dexterity +2, Intelligence +1"},
			{Name: "Keen Senses", Desc: "You have Perception."}, {Name: "Cantrip", Desc: "You know one wizard cantrip."},
		}, SubraceOf: "elf-2014"}},
		{"background/acolyte", TextData{Desc: "You devoted yourself to service in a temple.", Parts: []NameDesc{
			{Name: "Ability Scores", Desc: "Intelligence, Wisdom, Charisma", Type: partAbilities},
			{Name: "Feat", Desc: "Magic Initiate (Cleric)", Type: partFeat},
			{Name: "Skill Proficiencies", Desc: "Insight and Religion", Type: partSkills},
			{Name: "Tool Proficiency", Desc: "Calligrapher's Supplies", Type: partTools},
			{Name: "Equipment", Desc: "Choose A or B: (A) Holy Symbol, 8 GP; or (B) 50 GP", Type: "equipment"},
		}}},
		{"feat/grappler", TextData{
			Prerequisite: "Level 4+, Strength 13+ or Dexterity 13+", Category: "General",
			Parts: []NameDesc{
				{Name: "Punch and Grab", Desc: "You can impose the [[compendium:condition/grappled|Grappled]] condition."},
				{Name: "Ability Score Increase", Desc: "Increase your Strength or Dexterity score by 1, to a maximum of 20."},
			},
		}},
		{"feat/alert", TextData{
			Desc: "You gain the following benefits.", Category: "Origin",
			Parts: []NameDesc{{Name: "Initiative Proficiency", Desc: "Add your Proficiency Bonus to Initiative."}},
		}},
		{"condition/prone", TextData{Desc: "- **Restricted Movement.** You can only crawl."}},
		{"condition/concentration", TextData{Desc: "Some spells require you to maintain concentration.", Category: "Status"}},
		{"rule/sewer-plague", TextData{Desc: "A generic name for diseases incubated in sewers.", Category: "Disease"}},
		{"rule/advantage", TextData{Desc: "If you have Advantage, roll two d20s and use the higher roll. See also Disadvantage.", Category: "Core"}},
	} {
		if d := got[tt.key].Data; !reflect.DeepEqual(d, tt.data) {
			t.Errorf("%s\n got %#v\nwant %#v", tt.key, d, tt.data)
		}
	}

	// Fluff joins the description and the search text.
	cat := got["monster/cat"]
	if d := cat.Data.(MonsterData); d.Desc != "Cats prowl rooftops and [[compendium:item/bag-of-holding|bags]]." || !strings.Contains(cat.Text, "rooftops and bags.") {
		t.Errorf("cat fluff %q, text %q", d.Desc, cat.Text)
	}

	images := FiveToolsImages(files)
	if want := map[string]string{"cat": fiveToolsImages + "XPHB/Cat.webp"}; !reflect.DeepEqual(images, want) {
		t.Errorf("images %v", images)
	}

	delete(files, "book.json")
	if _, _, err := parseFiveToolsBook("5etools-xphb", files); err == nil {
		t.Error("no book.json: no error")
	}
}

func TestFtFeatIncrease(t *testing.T) {
	all := `{"choose":{"from":["str","dex","con","int","wis","cha"]}`
	for _, tt := range []struct{ ability, category, want string }{
		{`[{"choose":{"from":["str","dex"]}}]`, "G", "Increase your Strength or Dexterity score by 1, to a maximum of 20."},
		{`[{"choose":{"from":["int","wis","cha"]}}]`, "G", "Increase your Intelligence, Wisdom or Charisma score by 1, to a maximum of 20."},
		{`[` + all + `}]`, "G", "Increase one ability score of your choice by 1, to a maximum of 20."},
		{`[{"con":1}]`, "O", "Increase your Constitution score by 1, to a maximum of 20."},
		{`[` + all + `}]`, "EB", "Increase one ability score of your choice by 1, to a maximum of 30."},
		{`[` + all + `,"max":30}]`, "", "Increase one ability score of your choice by 1, to a maximum of 30."},
		{`[{"choose":{"from":["str","dex","con","int","wis","cha"],"entry":"Choose one ability in which you lack saving throw proficiency. Increase the chosen ability score by 1, to a maximum of 20."}}]`, "G",
			"Increase one ability score of your choice by 1, to a maximum of 20. Choose one ability in which you lack saving throw proficiency."},
		{`[{"choose":{"from":["str","dex","con","int","wis","cha"],"amount":2},"hidden":true}]`, "G", ""},
		{``, "G", ""},
	} {
		var ab []map[string]json.RawMessage
		_ = json.Unmarshal([]byte(tt.ability), &ab)
		if got := ftFeatIncrease(ab, tt.category); got != tt.want {
			t.Errorf("%s %s: %q, want %q", tt.ability, tt.category, got, tt.want)
		}
	}
}

func TestFiveToolsEdition(t *testing.T) {
	for _, tt := range []struct{ book, edition, attribution string }{
		{`{"id":"XPHB","name":"Player's Handbook (2024)","published":"2024-09-17","publisher":"Wizards of the Coast"}`, "2024", "Player's Handbook (2024), © Wizards of the Coast"},
		{`{"id":"PHB","name":"Player's Handbook","published":"2014-08-19"}`, "2014", "Player's Handbook, © Wizards of the Coast"},
		{`{"id":"TDCSR","name":"Tal'Dorei Campaign Setting Reborn","published":"2022-01-18","publisher":"Darrington Press"}`, "2014", "Tal'Dorei Campaign Setting Reborn, © Darrington Press"},
		{`{"id":"XMM","name":"Monster Manual (2025)","published":"2025-02-18"}`, "2024", "Monster Manual (2025), © Wizards of the Coast"},
	} {
		src, _, err := parseFiveToolsBook("k", map[string][]byte{"book.json": []byte(tt.book)})
		if err != nil || src.Edition != tt.edition || src.Attribution != tt.attribution {
			t.Errorf("%s: %+v %v", tt.book, src, err)
		}
	}
}

func TestFiveToolsCell(t *testing.T) {
	for _, tt := range []struct{ raw, want string }{
		{`{"roll":{"min":1,"max":5}}`, "1–5"},
		{`"a | b"`, "a / b"},
		{`"You know the {@spell Dancing Lights|XPHB} cantrip | or not."`, "You know the [[compendium:spell/dancing-lights|Dancing Lights]] cantrip / or not."},
	} {
		if got := ftCell(json.RawMessage(tt.raw)); got != tt.want {
			t.Errorf("%s: %q, want %q", tt.raw, got, tt.want)
		}
	}
}

func TestFiveToolsCopyTargets(t *testing.T) {
	files := readBook(t, fiveToolsFixture)
	var doc map[string]json.RawMessage
	if err := json.Unmarshal(files["bestiary.json"], &doc); err != nil {
		t.Fatal(err)
	}
	doc[copyTargetsKey] = json.RawMessage(`{"monster":[{"name":"Owl","source":"XMM","size":["T"],"type":"beast","ac":[11],"hp":{"average":1,"formula":"1d4 - 1"},"cr":"0"}]}`)
	files["bestiary.json"], _ = json.Marshal(doc)
	_, entries, skipped, err := parseFiveToolsBookReport("5etools-xphb", files)
	if err != nil {
		t.Fatal(err)
	}
	if skipped != 0 {
		t.Errorf("skipped %d, want 0", skipped)
	}
	got := map[string]Entry{}
	for _, e := range entries {
		got[string(e.Kind)+"/"+e.Slug] = e
	}
	if m, ok := got["monster/owl-spirit"].Data.(MonsterData); !ok || m.HP != 1 {
		t.Errorf("owl spirit %+v", got["monster/owl-spirit"])
	}
	if _, ok := got["monster/owl"]; ok {
		t.Error("the copy target was imported")
	}
}

func TestFiveToolsCopy(t *testing.T) {
	base := `{"name":"Base","source":"B","page":3,"hasToken":true,"trait":[{"name":"A","entries":["The base bites."]},{"name":"B","entries":["x"]}]}`
	for _, tt := range []struct {
		name, mod string
		want      string // the trait names and first trait text, or "skip"
	}{
		{"plain copy", `{}`, "A,B: The base bites."},
		{"replaceTxt", `{"*":{"mode":"replaceTxt","replace":"the (base)","with":"a big $1","flags":"i"}}`, "A,B: a big base bites."},
		{"appendArr", `{"trait":{"mode":"appendArr","items":{"name":"C","entries":[]}}}`, "A,B,C: The base bites."},
		{"insertArr", `{"trait":{"mode":"insertArr","index":0,"items":[{"name":"C","entries":["c"]}]}}`, "C,A,B: c"},
		{"removeArr", `{"trait":{"mode":"removeArr","names":["A"]}}`, "B: x"},
		{"replaceArr", `{"trait":{"mode":"replaceArr","replace":"A","items":{"name":"Z","entries":["z"]}}}`, "Z,B: z"},
		{"several ops", `{"trait":[{"mode":"removeArr","names":"B"},{"mode":"prependArr","items":{"name":"P","entries":["p"]}}]}`, "P,A: p"},
		{"remove key", `{"trait":"remove"}`, ": "},
		{"unknown mode", `{"trait":{"mode":"addSpells"}}`, "skip"},
	} {
		body := `{"monster":[` + base + `,{"name":"Copy","source":"B","_copy":{"name":"Base","source":"B","_mod":` + tt.mod + `}}]}`
		recs, skipped, err := ftRecords(map[string][]byte{"f": []byte(body)}, "f", "monster")
		if err != nil {
			t.Fatal(err)
		}
		if tt.want == "skip" {
			if skipped != 1 || len(recs) != 1 {
				t.Errorf("%s: skipped %d of %d", tt.name, skipped, len(recs)+skipped)
			}
			continue
		}
		if len(recs) != 2 {
			t.Fatalf("%s: %d records", tt.name, len(recs))
		}
		m, err := ftDecode[ftMonster](recs[1])
		if err != nil {
			t.Fatal(err)
		}
		var names []string
		first := ""
		for i, tr := range m.Trait {
			names = append(names, tr.Name)
			if i == 0 {
				first = ftEntries(tr.Entries)
			}
		}
		if got := strings.Join(names, ",") + ": " + first; got != tt.want || m.Name != "Copy" || m.Page != 0 || m.HasToken {
			t.Errorf("%s: %q (name %s page %d token %v), want %q", tt.name, got, m.Name, m.Page, m.HasToken, tt.want)
		}
		if _, ok := recs[1]["_copy"]; ok {
			b, _ := json.Marshal(recs[1])
			t.Errorf("%s: _copy kept: %s", tt.name, b)
		}
	}

	for _, tt := range []struct {
		name, versions string
		want           string // the versions' names, or "" when none
		skipped        int
	}{
		{"none", `null`, "", 0},
		{"explicit", `[{"name":"Base; Tall","source":"B","speed":40}]`, "Base (Tall)", 0},
		{"template", `[{"_abstract":{"name":"Base ({{c}})","source":"B"},"_implementations":[{"_variables":{"c":"Red \"R\""}},{"_variables":{"c":"Blue"}}]}]`, `Base (Red "R"),Base (Blue)`, 0},
		{"unknown mode", `[{"name":"Base; Odd","_mod":{"trait":{"mode":"addSpells"}}},{"name":"Base; Ok"}]`, "Base (Ok)", 1},
		{"not a list", `{}`, "", 1},
	} {
		rec := ftRec{"name": json.RawMessage(`"Base"`), "source": json.RawMessage(`"B"`), "trait": json.RawMessage(`[]`)}
		if tt.versions != "null" {
			rec["_versions"] = json.RawMessage(tt.versions)
		}
		vs, skipped := ftVersions(rec)
		var names []string
		for _, v := range vs {
			n, _ := ftDecode[ftNamed](v)
			names = append(names, n.Name)
			if _, ok := v["_versions"]; ok || n.Source != "B" {
				t.Errorf("%s: %s kept _versions or lost its source %q", tt.name, n.Name, n.Source)
			}
		}
		if got := strings.Join(names, ","); got != tt.want || skipped != tt.skipped {
			t.Errorf("%s: %q skipped %d, want %q skipped %d", tt.name, got, skipped, tt.want, tt.skipped)
		}
	}

	// A subclass kept once per class edition shares its name and source; the
	// class fields pick the target.
	body := `{"subclass":[{"name":"Moon","source":"P","className":"Druid","classSource":"P","page":1},
		{"name":"Moon","source":"P","className":"Druid","classSource":"X","_copy":{"name":"Moon","source":"P","className":"Druid","classSource":"P"}}]}`
	recs, skipped, err := ftRecords(map[string][]byte{"f": []byte(body)}, "f", "subclass")
	if err != nil || skipped != 0 || len(recs) != 2 || string(recs[1]["classSource"]) != `"X"` {
		t.Errorf("class edition copy: %d records, %d skipped, %v", len(recs), skipped, err)
	}
}
