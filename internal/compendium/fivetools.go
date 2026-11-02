package compendium

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// 5e.tools refuses server fetches and builds its pages in the browser, so
// a 5e.tools bestiary link is read from the same data on the 5etools
// GitHub mirror, one JSON file per book.

const (
	fiveToolsMirror  = "https://raw.githubusercontent.com/5etools-mirror-3/5etools-src/main/data/bestiary/bestiary-"
	fiveToolsImages  = "https://raw.githubusercontent.com/5etools-mirror-3/5etools-img/main/bestiary/"
	fiveToolsMaxBody = 16 << 20
)

var fiveToolsSite = Site{
	Key: "5etools", Name: "5e.tools bestiary", Home: "https://5e.tools/bestiary.html",
	Imports: "monsters of every book, with a picture when it has one",
	Note:    "5e.tools blocks the server, so the monster is read from the 5etools GitHub mirror instead.",
	NoteURL: "https://github.com/5etools-mirror-3/5etools-src/tree/main/data/bestiary",
	hosts:   []string{"5e.tools", "www.5e.tools"},
	page:    fiveToolsPage,
	check: func(page string) (Kind, error) {
		if !reFiveToolsPage.MatchString(page) {
			return "", errors.New("5e.tools blocks the server, so only bestiary (monster) links can be read, like https://5e.tools/bestiary/howler-mtf.html")
		}
		return KindMonster, nil
	},
	link: func(page string) string {
		src, name, _ := strings.Cut(page, "/")
		return "https://5e.tools/bestiary/" + name + "-" + src + ".html"
	},
	load: func(ctx context.Context, page string) (string, string, Entry, error) {
		src, name, _ := strings.Cut(page, "/")
		body, err := fetch(ctx, fiveToolsMirror+src+".json", fiveToolsMaxBody)
		if err != nil {
			return "", "", Entry{}, err
		}
		return ParseFiveTools(body, name)
	},
	edition: func(page string) string {
		// The 2024 books' codes start with X (XMM, XPHB).
		if strings.HasPrefix(page, "x") {
			return "2024"
		}
		return "2014"
	},
}

var reFiveToolsPage = regexp.MustCompile(`^[a-z0-9]+/[a-z0-9]+(?:-[a-z0-9]+)*$`)

// fiveToolsPage is "<book>/<name slug>" of /bestiary/<name>-<book>.html or
// /bestiary.html#<name>_<book>.
func fiveToolsPage(u *url.URL) string {
	var name, src string
	if file, ok := strings.CutPrefix(u.Path, "/bestiary/"); ok {
		file = strings.TrimSuffix(file, ".html")
		if i := strings.LastIndex(file, "-"); i > 0 {
			name, src = file[:i], file[i+1:]
		}
	} else if u.Path == "/bestiary.html" {
		frag, _, _ := strings.Cut(u.Fragment, ",")
		if i := strings.LastIndex(frag, "_"); i > 0 {
			name, src = frag[:i], frag[i+1:]
		}
	}
	return strings.ToLower(src) + "/" + Slugify(name)
}

type ftMonster struct {
	Name      string            `json:"name"`
	Source    string            `json:"source"`
	Page      int               `json:"page"`
	Copy      json.RawMessage   `json:"_copy"`
	Size      []string          `json:"size"`
	Type      json.RawMessage   `json:"type"`
	Alignment []json.RawMessage `json:"alignment"`
	AC        []json.RawMessage `json:"ac"`
	HP        struct {
		Average int    `json:"average"`
		Formula string `json:"formula"`
		Special string `json:"special"`
	} `json:"hp"`
	Speed           map[string]json.RawMessage `json:"speed"`
	Str             int                        `json:"str"`
	Dex             int                        `json:"dex"`
	Con             int                        `json:"con"`
	Int             int                        `json:"int"`
	Wis             int                        `json:"wis"`
	Cha             int                        `json:"cha"`
	Save            map[string]string          `json:"save"`
	Skill           map[string]json.RawMessage `json:"skill"`
	Senses          []string                   `json:"senses"`
	Passive         json.RawMessage            `json:"passive"`
	Languages       []string                   `json:"languages"`
	CR              json.RawMessage            `json:"cr"`
	Resist          []json.RawMessage          `json:"resist"`
	Immune          []json.RawMessage          `json:"immune"`
	ConditionImmune []json.RawMessage          `json:"conditionImmune"`
	Vulnerable      []json.RawMessage          `json:"vulnerable"`
	Trait           []ftEntry                  `json:"trait"`
	Spellcasting    []ftSpellcasting           `json:"spellcasting"`
	Action          []ftEntry                  `json:"action"`
	Bonus           []ftEntry                  `json:"bonus"`
	Reaction        []ftEntry                  `json:"reaction"`
	Legendary       []ftEntry                  `json:"legendary"`
	HasFluffImages  bool                       `json:"hasFluffImages"`
	HasToken        bool                       `json:"hasToken"`
}

type ftEntry struct {
	Name    string            `json:"name"`
	Entries []json.RawMessage `json:"entries"`
}

type ftSpellcasting struct {
	Name          string                     `json:"name"`
	HeaderEntries []json.RawMessage          `json:"headerEntries"`
	FooterEntries []json.RawMessage          `json:"footerEntries"`
	Will          []string                   `json:"will"`
	Daily         map[string][]string        `json:"daily"`
	Spells        map[string]json.RawMessage `json:"spells"`
	DisplayAs     string                     `json:"displayAs"`
}

var ftSizes = map[string]string{"T": "Tiny", "S": "Small", "M": "Medium", "L": "Large", "H": "Huge", "G": "Gargantuan"}

var ftAlignments = map[string]string{
	"L": "lawful", "N": "neutral", "C": "chaotic", "G": "good", "E": "evil",
	"U": "unaligned", "A": "any alignment", "NX": "neutral", "NY": "neutral",
}

// ParseFiveTools normalizes the monster named like slug in a bestiary
// file into an entry, returning its book and picture link.
func ParseFiveTools(body []byte, slug string) (book, image string, e Entry, err error) {
	var file struct {
		Monster []ftMonster `json:"monster"`
	}
	if err := json.Unmarshal(body, &file); err != nil {
		return "", "", Entry{}, fmt.Errorf("5etools: %w", err)
	}
	var m *ftMonster
	for i := range file.Monster {
		if Slugify(file.Monster[i].Name) == slug {
			m = &file.Monster[i]
			break
		}
	}
	switch {
	case m == nil:
		return "", "", Entry{}, fmt.Errorf("5etools: no monster %q in that book", slug)
	case m.Copy != nil:
		// shortcut: monsters that copy another one with changes are refused; resolve _copy if GMs need them.
		return "", "", Entry{}, fmt.Errorf("5etools: %s is a variant of another monster; import that one instead", m.Name)
	}

	d := MonsterData{
		Type: ftType(m.Type), HP: m.HP.Average, HitDice: m.HP.Formula,
		Speed: map[string]int{}, Saves: map[string]int{}, Skills: map[string]int{},
		Abilities: map[string]int{"str": m.Str, "dex": m.Dex, "con": m.Con, "int": m.Int, "wis": m.Wis, "cha": m.Cha},
		Senses:    stripTags(strings.Join(m.Senses, ", ")), Languages: stripTags(strings.Join(m.Languages, ", ")),
		Resistance: ftList(m.Resist, "resist"),
		Immunity:   joinNonEmpty("; ", ftList(m.Immune, "immune"), ftList(m.ConditionImmune, "conditionImmune")),
		Initiative: abilityMod(m.Dex),
	}
	var sizes []string
	for _, s := range m.Size {
		sizes = append(sizes, ftSizes[s])
	}
	d.Size = strings.Join(sizes, " or ")
	var align []string
	for _, raw := range m.Alignment {
		var s string
		if json.Unmarshal(raw, &s) == nil {
			align = append(align, ftAlignments[s])
		}
	}
	d.Alignment = strings.Join(align, " ")
	if len(m.AC) > 0 {
		var ac struct {
			AC   int      `json:"ac"`
			From []string `json:"from"`
		}
		if json.Unmarshal(m.AC[0], &d.AC) != nil && json.Unmarshal(m.AC[0], &ac) == nil {
			d.AC, d.ACDetail = ac.AC, stripTags(strings.Join(ac.From, ", "))
		}
	}
	if m.HP.Special != "" {
		d.HitDice = m.HP.Special
	}
	for mode, raw := range m.Speed {
		var n int
		var o struct {
			Number int `json:"number"`
		}
		if json.Unmarshal(raw, &n) == nil || json.Unmarshal(raw, &o) == nil && o.Number > 0 {
			d.Speed[mode] = max(n, o.Number)
		}
	}
	for k, v := range m.Save {
		d.Saves[k], _ = strconv.Atoi(strings.TrimPrefix(v, "+"))
	}
	for k, raw := range m.Skill {
		var v string
		if json.Unmarshal(raw, &v) == nil {
			d.Skills[Slugify(k)], _ = strconv.Atoi(strings.TrimPrefix(v, "+"))
		}
	}
	var passive any
	if json.Unmarshal(m.Passive, &passive) == nil && passive != nil {
		d.Senses = joinNonEmpty(", ", d.Senses, fmt.Sprintf("passive Perception %v", passive))
	}
	var cr struct {
		CR string `json:"cr"`
	}
	if json.Unmarshal(m.CR, &d.CR) != nil && json.Unmarshal(m.CR, &cr) == nil {
		d.CR = cr.CR
	}
	if v := ftList(m.Vulnerable, "vulnerable"); v != "" {
		d.Desc = "Vulnerabilities: " + v
	}

	for _, t := range m.Trait {
		d.Traits = append(d.Traits, NameDesc{Name: stripTags(t.Name), Desc: ftEntries(t.Entries)})
	}
	for _, s := range m.Spellcasting {
		desc := ftSpells(s)
		if s.DisplayAs == "action" {
			d.Actions = append(d.Actions, MonsterAction{Name: s.Name, Desc: desc, Type: "action"})
		} else {
			d.Traits = append(d.Traits, NameDesc{Name: s.Name, Desc: desc})
		}
	}
	for i, list := range [][]ftEntry{m.Action, m.Bonus, m.Reaction, m.Legendary} {
		for _, a := range list {
			d.Actions = append(d.Actions, MonsterAction{Name: stripTags(a.Name), Desc: ftEntries(a.Entries), Type: actionTypes[i]})
		}
	}

	book = m.Source
	if m.Page > 0 {
		book += fmt.Sprintf(" p. %d", m.Page)
	}
	switch {
	case m.HasFluffImages:
		image = fiveToolsImages + url.PathEscape(m.Source) + "/" + url.PathEscape(m.Name) + ".webp"
	case m.HasToken:
		image = fiveToolsImages + "tokens/" + url.PathEscape(m.Source) + "/" + url.PathEscape(m.Name) + ".webp"
	}
	return book, image, Entry{Kind: KindMonster, Slug: slug, Name: m.Name, Data: d, Text: d.searchText()}, nil
}

// ftType is "fiend (demon)" from "fiend" or {"type":"fiend","tags":["demon"]}.
func ftType(raw json.RawMessage) string {
	var s string
	if json.Unmarshal(raw, &s) == nil {
		return s
	}
	var t struct {
		Type json.RawMessage   `json:"type"`
		Tags []json.RawMessage `json:"tags"`
	}
	if json.Unmarshal(raw, &t) != nil {
		return ""
	}
	_ = json.Unmarshal(t.Type, &s)
	var tags []string
	for _, raw := range t.Tags {
		var tag string
		if json.Unmarshal(raw, &tag) == nil {
			tags = append(tags, tag)
		}
	}
	if len(tags) > 0 {
		s += " (" + strings.Join(tags, ", ") + ")"
	}
	return s
}

// ftList is a damage or condition list: names, or groups with a note
// ({"resist":[…],"note":"from nonmagical attacks"}), or special text.
func ftList(raws []json.RawMessage, key string) string {
	var simple, groups []string
	for _, raw := range raws {
		var s string
		if json.Unmarshal(raw, &s) == nil {
			simple = append(simple, s)
			continue
		}
		var g map[string]json.RawMessage
		if json.Unmarshal(raw, &g) != nil {
			continue
		}
		var special, pre, note string
		_ = json.Unmarshal(g["special"], &special)
		_ = json.Unmarshal(g["preNote"], &pre)
		_ = json.Unmarshal(g["note"], &note)
		var inner []json.RawMessage
		_ = json.Unmarshal(g[key], &inner)
		groups = append(groups, strings.Join(strings.Fields(strings.Join([]string{pre, special, ftList(inner, key), note}, " ")), " "))
	}
	return stripTags(joinNonEmpty("; ", append([]string{strings.Join(simple, ", ")}, groups...)...))
}

// ftEntries is entries as Markdown paragraphs.
func ftEntries(raws []json.RawMessage) string {
	var out []string
	for _, raw := range raws {
		if t := ftEntry1(raw); t != "" {
			out = append(out, t)
		}
	}
	return strings.Join(out, "\n\n")
}

// ftEntry1 is one entry: text, a list, or named entries.
func ftEntry1(raw json.RawMessage) string {
	var s string
	if json.Unmarshal(raw, &s) == nil {
		return stripTags(s)
	}
	var o struct {
		Type    string            `json:"type"`
		Name    string            `json:"name"`
		Entry   json.RawMessage   `json:"entry"`
		Entries []json.RawMessage `json:"entries"`
		Items   []json.RawMessage `json:"items"`
	}
	if json.Unmarshal(raw, &o) != nil {
		return ""
	}
	body := ftEntries(o.Entries)
	if o.Entry != nil {
		body = ftEntry1(o.Entry)
	}
	if o.Type == "list" {
		var items []string
		for _, it := range o.Items {
			items = append(items, "- "+ftEntry1(it))
		}
		body = strings.Join(items, "\n")
	}
	if o.Name != "" {
		return "**" + stripTags(o.Name) + ".** " + body
	}
	return body
}

// ftSpells is a spellcasting block as text.
func ftSpells(s ftSpellcasting) string {
	lines := []string{ftEntries(s.HeaderEntries)}
	if len(s.Will) > 0 {
		lines = append(lines, "At will: "+stripTags(strings.Join(s.Will, ", ")))
	}
	for _, k := range slices.Sorted(maps.Keys(s.Daily)) {
		// "1e" is once a day each.
		n, each := strings.CutSuffix(k, "e")
		label := n + "/day"
		if each {
			label += " each"
		}
		lines = append(lines, label+": "+stripTags(strings.Join(s.Daily[k], ", ")))
	}
	for _, k := range slices.Sorted(maps.Keys(s.Spells)) {
		var lvl struct {
			Slots  int      `json:"slots"`
			Spells []string `json:"spells"`
		}
		if json.Unmarshal(s.Spells[k], &lvl) != nil {
			continue
		}
		label := "Level " + k
		if k == "0" {
			label = "Cantrips"
		} else if lvl.Slots > 0 {
			label += fmt.Sprintf(" (%d slots)", lvl.Slots)
		}
		lines = append(lines, label+": "+stripTags(strings.Join(lvl.Spells, ", ")))
	}
	lines = append(lines, ftEntries(s.FooterEntries))
	return strings.TrimSpace(joinNonEmpty("\n\n", lines...))
}

var reTag = regexp.MustCompile(`\{@(\w+)(?: ([^{}]*))?\}`)

var ftAttacks = map[string]string{
	"mw": "Melee Weapon Attack:", "rw": "Ranged Weapon Attack:", "mw,rw": "Melee or Ranged Weapon Attack:",
	"ms": "Melee Spell Attack:", "rs": "Ranged Spell Attack:", "ms,rs": "Melee or Ranged Spell Attack:",
	"m": "Melee Attack Roll:", "r": "Ranged Attack Roll:", "m,r": "Melee or Ranged Attack Roll:",
}

var ftAbilityNames = map[string]string{
	"str": "Strength", "dex": "Dexterity", "con": "Constitution", "int": "Intelligence", "wis": "Wisdom", "cha": "Charisma",
}

// stripTags turns 5etools tags ({@hit 6}, {@spell fireball|phb}) into
// plain text, innermost first.
func stripTags(s string) string {
	for {
		out := reTag.ReplaceAllStringFunc(s, func(m string) string {
			sub := reTag.FindStringSubmatch(m)
			tag, arg := sub[1], sub[2]
			text, _, _ := strings.Cut(arg, "|")
			switch tag {
			case "atk", "atkr":
				return ftAttacks[arg]
			case "h":
				return "Hit: "
			case "hom":
				return "Hit or Miss: "
			case "hit", "d20":
				if strings.HasPrefix(text, "-") {
					return text
				}
				return "+" + text
			case "dc":
				return "DC " + text
			case "recharge":
				if text == "" {
					return "(Recharge 6)"
				}
				return "(Recharge " + text + "–6)"
			case "actSave":
				return ftAbilityNames[text] + " Saving Throw:"
			case "actSaveFail":
				return "Failure:"
			case "actSaveSuccess":
				return "Success:"
			case "actSaveSuccessOrFail":
				return "Failure or Success:"
			case "actTrigger":
				return "Trigger:"
			case "actResponse":
				return "Response:"
			case "chance":
				return text + " percent"
			}
			return text
		})
		if out == s {
			return out
		}
		s = out
	}
}
