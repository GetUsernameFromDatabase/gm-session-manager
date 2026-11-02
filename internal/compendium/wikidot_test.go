package compendium

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

// wikiPage wraps content the way dnd5e.wikidot.com lays out a page. The
// fixtures copy its markup only; the text is placeholder.
func wikiPage(title, content string) []byte {
	return []byte(`<html><body><div id="page-title">` + title + `</div><div id="page-content">` + content + `</div></body></html>`)
}

// TestParsePageURL: the link's host picks the site, and only the known
// sites' pages are accepted, since the server fetches them.
func TestParsePageURL(t *testing.T) {
	for _, tt := range []struct {
		url, site, page string
		kind            Kind
		err             error // checked when not nil
	}{
		{"https://dnd5e.wikidot.com/lineage:giff", "wikidot", "lineage:giff", KindSpecies, nil},
		{"http://dnd5e.wikidot.com/monk:mercy/", "wikidot", "monk:mercy", KindSubclass, nil},
		{"https://dnd5e.wikidot.com/artificer", "wikidot", "artificer", KindClass, nil},
		{"dnd5e.wikidot.com/feat:mobile", "wikidot", "feat:mobile", KindFeat, nil},
		{"https://dnd5e.wikidot.com/spell:fireball", "wikidot", "spell:fireball", "", nil},
		{"https://www.aidedd.org/dnd/monstres.php?vo=howler", "aidedd", "howler", KindMonster, nil},
		{"www.aidedd.org/dnd/monstres.php?vo=adult-red-dragon", "aidedd", "adult-red-dragon", KindMonster, nil},
		{"https://www.aidedd.org/dnd/sorts.php?vo=fireball", "aidedd", "dnd/sorts.php?vo=fireball", "", nil},
		{"https://5e.tools/bestiary/howler-mtf.html", "5etools", "mtf/howler", KindMonster, nil},
		{"https://5e.tools/bestiary/adult-red-dragon-mm.html", "5etools", "mm/adult-red-dragon", KindMonster, nil},
		{"https://5e.tools/bestiary.html#adult%20red%20dragon_mm", "5etools", "mm/adult-red-dragon", KindMonster, nil},
		{"asdasd", "", "", "", errNotASite},
		{"lineage:giff", "", "", "", errNotASite},
		{"@evil.example/feat:mobile", "", "", "", errNotASite},
		{"https://evil.example/lineage:giff", "", "", "", errNotASite},
		{"https://dnd5e.wikidot.com.evil.example/feat:mobile", "", "", "", errNotASite},
		{"https://dnd5e.wikidot.com:8080/feat:mobile", "", "", "", errNotASite},
		{"https://user@www.aidedd.org/dnd/monstres.php?vo=howler", "", "", "", errNotASite},
		{"file:///etc/passwd", "", "", "", errNotASite},
		{"https://dnd5e.wikidot.com/a/b", "", "", "", nil},
		{"https://www.aidedd.org/", "", "", "", nil},
		{"https://5e.tools/spells.html#fireball_phb", "", "", "", nil},
		{"https://5e.tools/bestiary/howler.html", "", "", "", nil},
	} {
		s, page, err := ParsePageURL(tt.url)
		if (err == nil) != (tt.site != "") || s.Key != tt.site || page != tt.page || tt.err != nil && !errors.Is(err, tt.err) {
			t.Errorf("%s: %q %q %v", tt.url, s.Key, page, err)
		}
		if err == nil {
			if kind := (Page{Site: s, ID: page}).Kind(); kind != tt.kind {
				t.Errorf("%s: kind %q", tt.url, kind)
			}
		}
	}
}

func TestParsePlainPage(t *testing.T) {
	e, err := ParsePlainPage([]byte(`<html><head><title>Site</title></head><body><nav>x</nav><div><h1>Fireball</h1><p>A bright streak.</p><p>Each creature<br>takes damage.</p></div></body></html>`))
	if err != nil || e.Name != "Fireball" || e.Kind != "" || e.Data.(TextData).Desc != "A bright streak.\n\nEach creature\ntakes damage." {
		t.Errorf("%+v %v", e, err)
	}
	if e, err := ParsePlainPage([]byte(`<html><head><title>Only a title</title></head><body><p>Text.</p></body></html>`)); err != nil || e.Name != "Only a title" {
		t.Errorf("%+v %v", e, err)
	}
	if _, err := ParsePlainPage([]byte(`<p>No name.</p>`)); err == nil {
		t.Error("no error without a name")
	}
}

func TestParseWikidotUnknownKind(t *testing.T) {
	_, e, err := ParseWikidot("spell:fireball", wikiPage("Fireball", `<p>Source: Some Book</p><p>A bright streak.</p>`))
	if err != nil || e.Kind != "" || e.Data.(TextData).Desc != "A bright streak." {
		t.Errorf("%+v %v", e, err)
	}
}

func TestParseWikidotSpecies(t *testing.T) {
	book, e, err := ParseWikidot("lineage:giff", wikiPage("Giff", `
<p>Source: Some Book</p>
<p><em><strong>Flavor about the folk.</strong></em></p>
<h5>Giff Traits</h5>
<ul>
<li><em><strong>Size.</strong></em> You are Medium.</li>
<li><em><strong>Speed.</strong></em> Your walking speed is 30 feet.</li>
<li><em><strong>Hippo Build.</strong></em> You are strong.</li>
<li><em><strong>Languages.</strong></em> Common and one more.</li>
</ul>`))
	if err != nil {
		t.Fatal(err)
	}
	d := e.Data.(TextData)
	want := []NameDesc{
		{Name: "Size", Desc: "You are Medium.", Type: partSize},
		{Name: "Speed", Desc: "Your walking speed is 30 feet.", Type: partSpeed},
		{Name: "Hippo Build", Desc: "You are strong."},
		{Name: "Languages", Desc: "Common and one more.", Type: partLanguages},
	}
	if e.Kind != KindSpecies || e.Slug != "giff" || e.Name != "Giff" || d.Desc != "Flavor about the folk." || !reflect.DeepEqual(d.Parts, want) {
		t.Errorf("%+v %+v", e, d)
	}
	roundTrip(t, e)
	if src := pageSource(Page{Site: wikidotSite, ID: "lineage:giff", Book: book}); src.Key != "wikidot:lineage:giff" || src.Attribution != "Some Book via https://dnd5e.wikidot.com/lineage:giff" {
		t.Errorf("%+v", src)
	}
}

// dnd5e.wikidot.com's theme has no #page-title: the title is a .page-title
// element, and the document title is the last resort.
func TestParseWikidotTitle(t *testing.T) {
	content := `<div id="page-content"><p>Source: Some Book</p><p>You are quick.</p></div>`
	for _, body := range []string{
		`<html><body><div class="page-title page-header"><span>Mobile</span></div>` + content + `</body></html>`,
		`<html><head><title>Mobile - DND 5th Edition</title></head><body>` + content + `</body></html>`,
	} {
		if _, e, err := ParseWikidot("feat:mobile", []byte(body)); err != nil || e.Name != "Mobile" {
			t.Errorf("%q %v", e.Name, err)
		}
	}
	if _, _, err := ParseWikidot("feat:mobile", []byte(`<html><title>Mobile</title></html>`)); err == nil || !strings.Contains(err.Error(), "#page-content") {
		t.Errorf("want a missing content error, got %v", err)
	}
}

func TestParseWikidotFeatAndBackground(t *testing.T) {
	_, e, err := ParseWikidot("feat:mobile", wikiPage("Mobile", `<p>Source: Some Book</p>
<p>You are quick. You gain the following benefits:</p>
<ul><li>Your speed increases by 10 feet.</li><li>Dashing ignores difficult terrain.</li></ul>`))
	if err != nil {
		t.Fatal(err)
	}
	if d := e.Data.(TextData); d.Desc != "You are quick. You gain the following benefits:\n\n- Your speed increases by 10 feet.\n\n- Dashing ignores difficult terrain." {
		t.Errorf("%q", d.Desc)
	}
	roundTrip(t, e)

	_, e, err = ParseWikidot("background:outlander", wikiPage("Outlander", `<p>Source: Some Book</p>
<p>You grew up <br>in the wilds.</p>
<p><strong>Skill Proficiencies:</strong> Athletics, Survival<br />
<strong>Tool Proficiencies:</strong> One type of musical instrument<br />
<strong>Equipment:</strong> A staff</p>
<h3>Feature: Wanderer</h3>
<p>You remember maps.</p>
<table><tr><th>d8</th><th>Trait</th></tr><tr><td>1</td><td>Roll me</td></tr></table>`))
	if err != nil {
		t.Fatal(err)
	}
	want := []NameDesc{
		{Name: "Skill Proficiencies", Desc: "Athletics, Survival", Type: partSkills},
		{Name: "Tool Proficiencies", Desc: "One type of musical instrument", Type: partTools},
		{Name: "Equipment", Desc: "A staff"},
		{Name: "Wanderer", Desc: "You remember maps."},
	}
	if d := e.Data.(TextData); d.Desc != "You grew up in the wilds." || !reflect.DeepEqual(d.Parts, want) {
		t.Errorf("%+v", d)
	}
	roundTrip(t, e)
}

func TestParseWikidotSubclass(t *testing.T) {
	_, e, err := ParseWikidot("monk:mercy", wikiPage("Monk: Way of Mercy", `
<p><em>Monks of this way heal and harm.</em></p>
<p>Source: Some Book</p>
<h3>Hands of Healing</h3>
<p>At 3rd level, your touch can mend wounds.</p>
<h3>Physician's Touch</h3>
<p>Starting at 6th level, you cure more.</p>`))
	if err != nil {
		t.Fatal(err)
	}
	d := e.Data.(ClassData)
	want := []ClassFeature{
		{Name: "Hands of Healing", Desc: "At 3rd level, your touch can mend wounds.", Levels: []int{3}},
		{Name: "Physician's Touch", Desc: "Starting at 6th level, you cure more.", Levels: []int{6}},
	}
	if e.Kind != KindSubclass || e.Name != "Way of Mercy" || e.Slug != "mercy" || d.SubclassOf != "monk" || !reflect.DeepEqual(d.Features, want) {
		t.Errorf("%+v %+v", e, d)
	}
	roundTrip(t, e)
}

func TestParseWikidotClass(t *testing.T) {
	_, e, err := ParseWikidot("artificer", wikiPage("Artificer", `
<p>Source: Some Book</p>
<h1>Class Features</h1>
<h5>Hit Points</h5>
<p><strong>Hit Dice:</strong> 1d8 per artificer level</p>
<h5>Proficiencies</h5>
<p><strong>Armor:</strong> Light armor</p>
<p><strong>Saving Throws:</strong> Constitution, Intelligence</p>
<p><strong>Skills:</strong> Choose two from Arcana, History, and Investigation</p>
<table>
<tr><th colspan="6"></th><th colspan="2">Spell Slots per Spell Level</th></tr>
<tr><th>Level</th><th>Proficiency Bonus</th><th>Features</th><th>Infusions Known</th><th>Cantrips Known</th><th>1st</th><th>5th</th></tr>
<tr><td>1st</td><td>+2</td><td>Magical Tinkering, Spellcasting</td><td>-</td><td>2</td><td>2</td><td>-</td></tr>
<tr><td>2nd</td><td>+2</td><td>Infuse Item</td><td>4</td><td>2</td><td>2</td><td>-</td></tr>
<tr><td>3rd</td><td>+2</td><td>Artificer Specialist</td><td>4</td><td>2</td><td>3</td><td>-</td></tr>
<tr><td>17th</td><td>+6</td><td>-</td><td>12</td><td>4</td><td>4</td><td>1</td></tr>
</table>
<h3>Magical Tinkering</h3>
<p>At 1st level, you tinker.</p>
<h3>Spellcasting</h3>
<p>You cast spells.</p>
<h5>Tools Required</h5>
<p>You need a tool.</p>
<h3>Infuse Item</h3>
<p>At 2nd level, you infuse.</p>
<h3>Artificer Specialist</h3>
<p>At 3rd level, choose a specialty:</p>
<table><tr><th>Specialty</th></tr><tr><td><a href="/artificer:alchemist">Alchemist</a></td></tr></table>`))
	if err != nil {
		t.Fatal(err)
	}
	d := e.Data.(ClassData)
	if d.HitDie != 8 || d.Caster != "half" || !reflect.DeepEqual(d.Saves, []string{"con", "int"}) {
		t.Errorf("%+v", d)
	}
	if got := CoreTraits(d.Desc); got["Skill Proficiencies"] != "Choose 2: Arcana, History, and Investigation" || got["Armor Training"] != "Light armor" {
		t.Errorf("%v", got)
	}
	if !reflect.DeepEqual(d.Table, map[string]map[int]string{"Infusions Known": {2: "4", 3: "4", 17: "12"}, "Cantrips Known": {1: "2", 2: "2", 3: "2", 17: "4"}}) {
		t.Errorf("%v", d.Table)
	}
	var names []string
	for _, f := range d.Features {
		names = append(names, f.Name)
	}
	if !reflect.DeepEqual(names, []string{"Magical Tinkering", "Spellcasting", "Infuse Item", "Artificer Specialist"}) {
		t.Fatalf("%q", names)
	}
	if f := d.Features[1]; !reflect.DeepEqual(f.Levels, []int{1}) || !strings.Contains(f.Desc, "**Tools Required.** You need a tool.") {
		t.Errorf("%+v", f)
	}
	if f := d.Features[3]; !f.Subclass || !reflect.DeepEqual(f.Levels, []int{3}) || d.Features[2].Subclass {
		t.Errorf("%+v", d.Features)
	}
	roundTrip(t, e)
}

func TestHomebrewTextRoundTrip(t *testing.T) {
	in := HomebrewInput{
		Kind: string(KindClass), Name: "Tinkerer", HitDie: 10, Caster: "half", Saves: []string{"con", "int"},
		Skills: "Choose 2: Arcana, History", Desc: "Makes things.",
		Features: "Level 1: Tinkering\nYou tinker.\n\nLevel 3: Specialist (subclass)\nPick one.",
		Table:    "Focus Points: -, 2, 3",
	}
	e, err := in.Entry()
	if err != nil {
		t.Fatal(err)
	}
	d := e.Data.(ClassData)
	if len(d.Features) != 2 || !d.Features[1].Subclass || !reflect.DeepEqual(d.Features[1].Levels, []int{3}) || d.Table["Focus Points"][3] != "3" {
		t.Fatalf("%+v", d)
	}
	back := HomebrewInputFor(in.Kind, in.Name, "", &d)
	if back.Features != in.Features || back.Skills != in.Skills || back.Desc != in.Desc || back.HitDie != 10 {
		t.Errorf("%+v", back)
	}

	sp := HomebrewInput{Kind: string(KindSpecies), Name: "Giff", Speed: 30, Size: "Medium", Traits: "Hippo Build. You are strong.\n\n**Astral Spark.** Extra damage.\nMore lines."}
	e, err = sp.Entry()
	if err != nil {
		t.Fatal(err)
	}
	td := e.Data.(TextData)
	want := []NameDesc{
		{Name: "Size", Desc: "Medium", Type: partSize}, {Name: "Speed", Desc: "30 feet", Type: partSpeed},
		{Name: "Hippo Build", Desc: "You are strong."}, {Name: "Astral Spark", Desc: "Extra damage.\nMore lines."},
	}
	if !reflect.DeepEqual(td.Parts, want) {
		t.Errorf("%+v", td.Parts)
	}
	if back := HomebrewInputFor(sp.Kind, sp.Name, "", &td); back.Speed != 30 || back.Size != "Medium" || back.Traits != "Hippo Build. You are strong.\n\nAstral Spark. Extra damage.\nMore lines." {
		t.Errorf("%+v", back)
	}
}

// roundTrip checks that saving the review form unchanged keeps everything a
// parsed page shows in it.
func roundTrip(t *testing.T, e Entry) {
	t.Helper()
	form := InputForEntry(e)
	back, err := form.Entry()
	if err != nil {
		t.Fatal(err)
	}
	if got := InputForEntry(back); !reflect.DeepEqual(got, form) {
		t.Errorf("%s changed in the form:\n got %+v\nwant %+v", e.Name, got, form)
	}
}
