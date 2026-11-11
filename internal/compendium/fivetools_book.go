package compendium

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
)

// This file reads one 5etools book folder (5etools-<id>): flat files in
// 5etools' own shapes, already filtered to the book's source, normalized
// into the shapes in entry.go.

// ftRec is one 5etools record, kept raw so _copy can merge it.
type ftRec = map[string]json.RawMessage

type ftBookInfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Group     string `json:"group"`
	Published string `json:"published"`
	Publisher string `json:"publisher"`
}

// parseFiveToolsBook normalizes a 5etools book folder.
func parseFiveToolsBook(key string, files map[string][]byte) (Source, []Entry, error) {
	src, entries, _, err := parseFiveToolsBookReport(key, files)
	return src, entries, err
}

// parseFiveToolsBookReport is parseFiveToolsBook plus the number of records
// skipped: copies of entries outside the file, unknown _mod modes, records
// that do not decode.
func parseFiveToolsBookReport(key string, files map[string][]byte) (Source, []Entry, int, error) {
	body, ok := files["book.json"]
	if !ok {
		return Source{}, nil, 0, errors.New("5etools: book.json is missing")
	}
	var book ftBookInfo
	if err := json.Unmarshal(body, &book); err != nil {
		return Source{}, nil, 0, fmt.Errorf("5etools book.json: %w", err)
	}
	src := Source{
		Key: key, Edition: "2014", License: "Not SRD — private use",
		Attribution: book.Name + ", © " + cmp.Or(book.Publisher, "Wizards of the Coast"),
		Published:   book.Published,
	}
	if book.Published >= "2024-09-01" {
		src.Edition = "2024"
	}
	p := &ftBook{id: book.ID, files: files, seen: map[string]bool{}}
	for _, step := range []func() error{p.monsters, p.spells, p.items, p.species, p.backgrounds, p.feats, p.conditions, p.rules, p.classes} {
		if err := step(); err != nil {
			return Source{}, nil, 0, err
		}
	}
	return src, p.out, p.skipped, nil
}

// FiveToolsImages maps the slugs of a book folder's monsters to their
// picture links, by the rule the single monster import uses.
func FiveToolsImages(files map[string][]byte) map[string]string {
	out := map[string]string{}
	recs, _, err := ftRecords(files, "bestiary.json", keyMonster)
	if err != nil {
		return out
	}
	for _, r := range recs {
		if m, err := ftDecode[ftMonster](r); err == nil && m.image() != "" {
			out[Slugify(m.Name)] = m.image()
		}
	}
	return out
}

type ftBook struct {
	id      string // the book's source code, "XPHB"
	files   map[string][]byte
	out     []Entry
	seen    map[string]bool // kind/slug, first wins
	skipped int
}

// records reads a file's list with copies resolved, counting the skips.
func (p *ftBook) records(file, key string) ([]ftRec, error) {
	recs, skipped, err := ftRecords(p.files, file, key)
	p.skipped += skipped
	return recs, err
}

// fluff is a fluff file's text by name|source.
func (p *ftBook) fluff(file, key string) (map[string]string, error) {
	recs, _, err := ftRecords(p.files, file, key)
	out := map[string]string{}
	for _, r := range recs {
		f, err := ftDecode[ftNamed](r)
		if err == nil {
			out[f.key()] = ftEntries(f.Entries)
		}
	}
	return out, err
}

// add stores an entry unless its kind and slug came first; text is the
// search body before tokens become labels. It returns the stored entry, or
// nil when skipped.
func (p *ftBook) add(kind enum.CompendiumKind, name string, data any, text string) *Entry {
	slug := Slugify(name)
	if slug == "" || p.seen[string(kind)+"/"+slug] {
		return nil
	}
	p.seen[string(kind)+"/"+slug] = true
	p.out = append(p.out, Entry{Kind: kind, Slug: slug, Name: name, Data: data, Text: plainText(text)})
	return &p.out[len(p.out)-1]
}

// each decodes every record of a list into T and calls fn.
func each[T any](p *ftBook, file, key string, fn func(T)) error {
	recs, err := p.records(file, key)
	if err != nil {
		return err
	}
	for _, r := range recs {
		v, err := ftDecode[T](r)
		if err != nil {
			p.skipped++
			continue
		}
		fn(v)
	}
	return nil
}

// ftNamed is the part every record has.
type ftNamed struct {
	Name    string            `json:"name"`
	Source  string            `json:"source"`
	Entries []json.RawMessage `json:"entries"`
}

func (n ftNamed) key() string { return strings.ToLower(n.Name + "|" + n.Source) }

func (p *ftBook) monsters() error {
	fluff, err := p.fluff("fluff-bestiary.json", keyMonsterFluff)
	if err != nil {
		return err
	}
	return each(p, "bestiary.json", keyMonster, func(m ftMonster) {
		d := m.data()
		d.Desc = fluff[strings.ToLower(m.Name+"|"+m.Source)]
		if e := p.add(enum.KindMonster, m.Name, d, joinText(d.searchText(), d.Desc)); e != nil {
			e.Habitats = habitatsFrom(m.Environment)
		}
	})
}

type ftSpell struct {
	ftNamed
	Level  int    `json:"level"`
	School string `json:"school"`
	Time   []struct {
		Number    int    `json:"number"`
		Unit      string `json:"unit"`
		Condition string `json:"condition"`
	} `json:"time"`
	Range struct {
		Type     string     `json:"type"`
		Distance ftDistance `json:"distance"`
	} `json:"range"`
	Components struct {
		V bool            `json:"v"`
		S bool            `json:"s"`
		M json.RawMessage `json:"m"`
	} `json:"components"`
	Duration []struct {
		Type          string     `json:"type"`
		Duration      ftDistance `json:"duration"`
		Concentration bool       `json:"concentration"`
		Ends          []string   `json:"ends"`
	} `json:"duration"`
	Meta struct {
		Ritual bool `json:"ritual"`
	} `json:"meta"`
	HigherLevel []ftEntry `json:"entriesHigherLevel"`
}

// ftDistance is a range's distance or a duration's length.
type ftDistance struct {
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}

var ftSchools = map[string]string{
	"A": "abjuration", "C": "conjuration", "D": "divination", "E": "enchantment",
	"V": "evocation", "I": "illusion", "N": "necromancy", "T": "transmutation", "P": "psionic",
}

func (p *ftBook) spells() error {
	fluff, err := p.fluff("fluff-spells.json", keySpellFluff)
	if err != nil {
		return err
	}
	// spell-classes.json is {"<book>": {"<spell>": {"class": [{name, source}]}}}.
	// Classes with definedInSource are other books' (Evoker, Mage of WoW).
	var lists map[string]map[string]struct {
		Class []struct {
			Name    string `json:"name"`
			Defined string `json:"definedInSource"`
		} `json:"class"`
	}
	if body, ok := p.files["spell-classes.json"]; ok {
		if err := json.Unmarshal(body, &lists); err != nil {
			return fmt.Errorf("5etools spell-classes.json: %w", err)
		}
	}
	return each(p, "spells.json", keySpell, func(s ftSpell) {
		d := SpellData{
			Level: s.Level, School: ftSchools[s.School], Ritual: s.Meta.Ritual,
			Desc: joinNonEmpty("\n\n", ftEntries(s.Entries), fluff[s.key()]),
		}
		for _, c := range lists[s.Source][s.Name].Class {
			if slug := Slugify(c.Name); c.Defined == "" && !slices.Contains(d.Classes, slug) {
				d.Classes = append(d.Classes, slug)
			}
		}
		var hl []string
		for _, e := range s.HigherLevel {
			hl = append(hl, ftEntries(e.Entries))
		}
		d.HigherLevel = strings.Join(hl, "\n\n")
		if len(s.Time) > 0 {
			d.CastingTime = ftCastingTime(s.Time[0].Number, s.Time[0].Unit, plainTags(s.Time[0].Condition))
		}
		d.Range = ftRange(s.Range.Type, s.Range.Distance)
		var comps []string
		for c, on := range map[string]bool{"V": s.Components.V, "S": s.Components.S, "M": s.Components.M != nil} {
			if on {
				comps = append(comps, c)
			}
		}
		slices.SortFunc(comps, func(a, b string) int { return strings.Index("VSM", a) - strings.Index("VSM", b) })
		d.Components = strings.Join(comps, ", ")
		var m struct {
			Text string `json:"text"`
		}
		// The material is text, {"text": …, "cost": …} or just true.
		if json.Unmarshal(s.Components.M, &m.Text) != nil {
			_ = json.Unmarshal(s.Components.M, &m)
		}
		if m.Text != "" {
			d.Components += " (" + plainTags(m.Text) + ")"
		}
		var durs []string
		for _, du := range s.Duration {
			d.Concentration = d.Concentration || du.Concentration
			durs = append(durs, ftDuration(du.Type, du.Duration, du.Concentration, du.Ends))
		}
		d.Duration = strings.Join(durs, " or ")
		p.add(enum.KindSpell, s.Name, d, joinText(d.School, d.Desc, d.HigherLevel))
	})
}

// 5etools casting time units that name an action.
const (
	ftUnitAction   = "action"
	ftUnitBonus    = "bonus"
	ftUnitReaction = "reaction"
)

// ftCastingTime is "Action", "Bonus Action", "Reaction, which you take…"
// or "10 minutes", as the sheet reads casting times.
func ftCastingTime(n int, unit, condition string) string {
	var s string
	switch unit {
	case ftUnitAction:
		s = "Action"
	case ftUnitBonus:
		s = "Bonus Action"
	case ftUnitReaction:
		s = "Reaction"
	default:
		s = ftAmount(n, unit)
	}
	if n > 1 && (unit == ftUnitAction || unit == ftUnitBonus || unit == ftUnitReaction) {
		s = strconv.Itoa(n) + " " + s + "s"
	}
	return joinNonEmpty(", ", s, condition)
}

// ftAmount is "1 minute", "10 minutes".
func ftAmount(n int, unit string) string {
	if n != 1 && !strings.HasSuffix(unit, "s") {
		unit += "s"
	}
	return strconv.Itoa(n) + " " + unit
}

func ftRange(typ string, d ftDistance) string {
	switch typ {
	case rangeSpecial:
		return "Special"
	case rangePoint:
		switch d.Type {
		case distFeet:
			return fmt.Sprintf("%d feet", d.Amount)
		case distMiles:
			return ftAmount(d.Amount, "mile")
		}
		return label(d.Type)
	}
	// An area around the caster: "Self (15-foot cone)".
	unit := map[string]string{distFeet: "foot", distMiles: "mile"}[d.Type]
	if d.Amount > 0 && unit != "" {
		return fmt.Sprintf("Self (%d-%s %s)", d.Amount, unit, typ)
	}
	return label(typ)
}

func ftDuration(typ string, d ftDistance, conc bool, ends []string) string {
	switch typ {
	case durationInstant:
		return "Instantaneous"
	case durationTimed:
		s := ftAmount(d.Amount, d.Type)
		if conc { // the page adds "Concentration, " from the flag
			return "up to " + s
		}
		return s
	case durationPermanent:
		var until []string
		for _, e := range ends {
			until = append(until, map[string]string{endsDispel: "dispelled", endsTrigger: "triggered", endsDischarge: "discharged"}[e])
		}
		if len(until) > 0 {
			return "Until " + strings.Join(slices.DeleteFunc(until, func(s string) bool { return s == "" }), " or ")
		}
		return "Permanent"
	}
	return label(typ)
}

type ftItem struct {
	ftNamed
	Type           string            `json:"type"`
	Rarity         string            `json:"rarity"`
	ReqAttune      json.RawMessage   `json:"reqAttune"`
	Value          float64           `json:"value"`
	Weight         float64           `json:"weight"`
	Dmg1           string            `json:"dmg1"`
	Dmg2           string            `json:"dmg2"`
	DmgType        string            `json:"dmgType"`
	Property       []json.RawMessage `json:"property"`
	Range          string            `json:"range"`
	WeaponCategory string            `json:"weaponCategory"`
	AC             int               `json:"ac"`
	Strength       json.RawMessage   `json:"strength"`
	Stealth        bool              `json:"stealth"`
	Wondrous       bool              `json:"wondrous"`
	Staff          bool              `json:"staff"`
	BaseItem       string            `json:"baseItem"`
	Mastery        []json.RawMessage `json:"mastery"` // "Vex|XPHB" or {"uid": "Vex|XPHB", "note": …}
}

// ftArmorWeights names the armor type codes a variant may require.
var ftArmorWeights = map[string]string{"LA": "light armor", "MA": "medium armor", "HA": "heavy armor"}

// ftRequires is what a magic variant applies to, in words: each requirement
// ({"type":"HA|XPHB"}, {"name":"Longsword","source":"XPHB"}, {"weapon":true},
// {"scfType":"arcane"}, {"weaponCategory":"simple"}) once, in order.
func ftRequires(reqs []map[string]json.RawMessage) string {
	var out []string
	for _, req := range reqs {
		str := func(k string) string {
			var s string
			_ = json.Unmarshal(req[k], &s)
			return s
		}
		code, _, _ := strings.Cut(str(keyType), "|")
		var w string
		switch {
		case str(keyName) != "":
			w = str(keyName)
		case str(keyScfType) != "":
			w = str(keyScfType) + " focus"
		case str(keyWeaponCategory) != "":
			w = str(keyWeaponCategory) + " weapon"
		case code != "":
			w = cmp.Or(ftArmorWeights[code], strings.ReplaceAll(ftItemTypes[code], "-", " "))
		default:
			for _, k := range slices.Sorted(maps.Keys(req)) {
				if string(req[k]) == "true" {
					w = k // weapon, sword, armor
				}
			}
		}
		if w != "" && !slices.Contains(out, w) {
			out = append(out, w)
		}
	}
	return strings.Join(out, ", ")
}

// Item categories (ItemData.Category), as ftItemTypes maps 5etools codes.
const (
	itemWeapon      = "weapon"
	itemArmor       = "armor"
	itemTools       = "tools"
	itemAdventuring = "adventuring-gear"
	itemVehicle     = "vehicle"
	itemTreasure    = "treasure"
	itemShield      = "shield"
	itemAmmunition  = "ammunition"
	itemSpellFocus  = "spellcasting-focus"
	itemPotion      = "potion"
	itemRing        = "ring"
	itemRod         = "rod"
	itemScroll      = "scroll"
	itemWand        = "wand"
	itemExplosive   = "explosive"
	itemMount       = "mount"
	itemTradeGood   = "trade-good"
	itemStaff       = "staff"
	itemWondrous    = "wondrous-item"
	itemOther       = "other"
)

// ftItemTypes maps 5etools item type codes to our categories (Open5e's).
var ftItemTypes = map[string]string{
	"M": itemWeapon, "R": itemWeapon, "LA": itemArmor, "MA": itemArmor, "HA": itemArmor, "S": itemShield,
	"A": itemAmmunition, "AF": itemAmmunition, "AT": itemTools, "T": itemTools, "INS": itemTools, "GS": itemTools,
	"G": itemAdventuring, "TAH": itemAdventuring, "FD": itemAdventuring, "SCF": itemSpellFocus,
	"P": itemPotion, "RG": itemRing, "RD": itemRod, "SC": itemScroll, "WD": itemWand, "EXP": itemExplosive,
	"MNT": itemMount, "VEH": itemVehicle, "SHP": itemVehicle, "AIR": itemVehicle, "SPC": itemVehicle,
	"TG": itemTradeGood, "$": itemTreasure, "$A": itemTreasure, "$C": itemTreasure, "$G": itemTreasure,
}

var ftDamageTypes = map[string]string{
	"A": "acid", "B": "bludgeoning", "C": "cold", "F": "fire", "O": "force", "L": "lightning",
	"N": "necrotic", "P": "piercing", "I": "poison", "Y": "psychic", "R": "radiant", "S": "slashing", "T": "thunder",
}

var ftProperties = map[string]string{
	"A": "Ammunition", "AF": "Ammunition", "BF": "Burst Fire", "F": "Finesse", "H": "Heavy", "L": "Light",
	"LD": "Loading", "R": "Reach", "RLD": "Reload", "S": "Special", "T": "Thrown", "2H": "Two-Handed", "V": "Versatile",
}

// ftBaseKeys are the fields a magic item takes from its base item.
var ftBaseKeys = []string{
	keyType, keyWeaponCategory, keyDmg1, keyDmg2, keyDmgType, keyProperty, keyRange,
	keyAC, keyStrength, keyStealth, keyWeight, keyWeapon, keyArmor, keyMastery,
}

func (p *ftBook) items() error {
	fluff, err := p.fluff("fluff-items.json", keyItemFluff)
	if err != nil {
		return err
	}
	// items-base.json is every book's base items: only the book's own count.
	all, _, err := ftRecords(p.files, "items-base.json", keyBaseItem)
	if err != nil {
		return err
	}
	bases := map[string]ftRec{}
	for _, r := range all {
		if n, err := ftDecode[ftNamed](r); err == nil {
			bases[n.key()] = r
		}
	}
	norm := func(it ftItem) {
		d := ftItemData(it)
		d.Desc = joinNonEmpty("\n\n", d.Desc, fluff[it.key()])
		p.add(enum.KindItem, it.Name, d, joinText(d.Category, d.Rarity, d.Desc))
	}
	// The book's own base items (mundane gear) first, then its items.
	for _, r := range all {
		if it, err := ftDecode[ftItem](r); err == nil && strings.EqualFold(it.Source, p.id) {
			norm(it)
		}
	}
	for _, key := range []string{keyItem, keyItemGroup} {
		recs, err := p.records("items.json", key)
		if err != nil {
			return err
		}
		for _, r := range recs {
			if n, err := ftDecode[ftItem](r); err == nil && n.BaseItem != "" {
				name, src, _ := strings.Cut(n.BaseItem, "|")
				if base, ok := bases[strings.ToLower(name+"|"+cmp.Or(src, "phb"))]; ok {
					r = maps.Clone(r)
					for _, k := range ftBaseKeys {
						if _, has := r[k]; !has && base[k] != nil {
							r[k] = base[k]
						}
					}
				}
			}
			it, err := ftDecode[ftItem](r)
			if err != nil {
				p.skipped++
				continue
			}
			norm(it)
		}
	}
	return p.magicVariants()
}

// ftPlaceholder is a {=field} reference to another field of the same record.
var ftPlaceholder = regexp.MustCompile(`\{=(\w+)\}`)

// magicVariants adds one item per generic variant ("+1 Weapon"), after the
// book's own items so a real item of the same name wins.
// shortcut: no per-base expansion ("+1 Longsword"), upgrade when players need the variant as a concrete base item.
func (p *ftBook) magicVariants() error {
	return each(p, "magicvariants.json", keyMagicVariant, func(v struct {
		ftNamed
		Requires []map[string]json.RawMessage `json:"requires"`
		Inherits ftRec                        `json:"inherits"`
	}) {
		rec := ftRec{}
		maps.Copy(rec, v.Inherits)
		// Fill from inherits' own values, as JSON text so quotes stay escaped.
		raw, _ := json.Marshal(rec[keyEntries])
		filled := ftPlaceholder.ReplaceAllStringFunc(string(raw), func(m string) string {
			var s string
			if json.Unmarshal(v.Inherits[m[2:len(m)-1]], &s) != nil {
				return m
			}
			q, _ := json.Marshal(s)
			return string(q[1 : len(q)-1])
		})
		rec[keyEntries] = json.RawMessage(filled)
		it, err := ftDecode[ftItem](rec)
		if err != nil {
			p.skipped++
			return
		}
		d := ftItemData(it)
		if applies := ftRequires(v.Requires); applies != "" {
			d.Desc = joinNonEmpty("\n\n", "Applies to: "+applies, d.Desc)
		}
		p.add(enum.KindItem, v.Name, d, joinText(d.Category, d.Rarity, d.Desc))
	})
}

// ftCode is a code without its "|source" suffix ("M|XPHB" → "M").
func ftCode(s string) string {
	code, _, _ := strings.Cut(s, "|")
	return code
}

func ftItemData(it ftItem) ItemData {
	typ := ftCode(it.Type)
	d := ItemData{
		Category: ftItemTypes[typ], Rarity: label(it.Rarity), Range: it.Range,
		Martial: it.WeaponCategory == keyMartial, Desc: ftEntries(it.Entries),
		Attunement: it.ReqAttune != nil && string(it.ReqAttune) != "false" && string(it.ReqAttune) != "null",
	}
	switch {
	case it.Staff:
		d.Category = itemStaff
	case d.Category == "" && it.Wondrous:
		d.Category = itemWondrous
	case d.Category == "":
		d.Category = itemOther
	}
	if it.Rarity == keyNone {
		d.Rarity = ""
	}
	if it.Value > 0 {
		d.Cost = strconv.FormatFloat(it.Value/100, 'f', -1, 64) // gp, as the page shows it
	}
	if it.Weight > 0 {
		d.Weight = strconv.FormatFloat(it.Weight, 'f', -1, 64)
	}
	if it.Dmg1 != "" {
		d.Damage = strings.TrimSpace(it.Dmg1 + " " + ftDamageTypes[it.DmgType])
	}
	var props []string
	for _, raw := range it.Property {
		var code string
		var o struct {
			UID string `json:"uid"`
		}
		if json.Unmarshal(raw, &code) != nil && json.Unmarshal(raw, &o) == nil {
			code = o.UID
		}
		name := ftProperties[ftCode(code)]
		if name == "Versatile" && it.Dmg2 != "" {
			name += " (" + it.Dmg2 + ")"
		}
		if name != "" {
			props = append(props, name)
		}
	}
	d.Properties = strings.Join(props, ", ")
	var mastery []string
	for _, raw := range it.Mastery {
		var m struct {
			UID string `json:"uid"`
		}
		if json.Unmarshal(raw, &m.UID) == nil || json.Unmarshal(raw, &m) == nil && m.UID != "" {
			mastery = append(mastery, ftCode(m.UID))
		}
	}
	d.Mastery = strings.Join(mastery, ", ")
	if a := ftArmor(typ, it); a != nil {
		d.Armor, d.ArmorClass = a, armorText(*a)
	}
	return d
}

func ftArmor(typ string, it ftItem) *ArmorData {
	a := &ArmorData{Base: it.AC, Stealth: it.Stealth}
	switch typ {
	case "LA":
		a.Type, a.AddDex = ArmorLight, true
	case "MA":
		a.Type, a.AddDex, a.DexCap = ArmorMedium, true, 2
	case "HA":
		a.Type = ArmorHeavy
	case "S":
		a.Type = ArmorShield
	default:
		return nil
	}
	var s string
	if json.Unmarshal(it.Strength, &s) != nil {
		_ = json.Unmarshal(it.Strength, &a.StrReq)
	}
	if s != "" {
		a.StrReq, _ = strconv.Atoi(s)
	}
	return a
}

type ftRace struct {
	ftNamed
	Size        []string                     `json:"size"`
	Speed       json.RawMessage              `json:"speed"`
	RaceName    string                       `json:"raceName"`
	RaceSource  string                       `json:"raceSource"`
	Ability     []map[string]json.RawMessage `json:"ability"`
	ReprintedAs json.RawMessage              `json:"reprintedAs"`
	Overwrite   struct {
		Ability bool `json:"ability"`
	} `json:"overwrite"`
}

func (p *ftBook) species() error {
	fluff, err := p.fluff("fluff-races.json", keyRaceFluff)
	if err != nil {
		return err
	}
	races := map[string]ftRace{}
	// A subrace without fluff of its own ("Elf (Drow)") shows its race's.
	norm := func(r ftRace, name string, parts []NameDesc, desc, fluffKey string, link TextData) {
		const asi = "Ability Score Increase"
		if inc := ftRaceAbility(r.Ability); inc != "" && !strings.Contains(joinText(desc, partsText(parts)), asi) {
			parts = append([]NameDesc{{Name: asi, Desc: inc}}, parts...)
		}
		var sizes []string
		for _, s := range r.Size {
			sizes = append(sizes, ftSizes[s])
		}
		var typed []NameDesc
		if len(sizes) > 0 {
			typed = append(typed, NameDesc{Name: LabelSize, Desc: strings.Join(sizes, " or "), Type: PartSize})
		}
		var walk int
		var speeds struct {
			Walk int `json:"walk"`
		}
		if json.Unmarshal(r.Speed, &walk) == nil || json.Unmarshal(r.Speed, &speeds) == nil {
			if walk = max(walk, speeds.Walk); walk > 0 {
				typed = append(typed, NameDesc{Name: LabelSpeed, Desc: fmt.Sprintf("%d feet", walk), Type: PartSpeed})
			}
		}
		// 2014 races also tell their size in a trait; the typed fact says it.
		if len(sizes) > 0 {
			parts = slices.DeleteFunc(parts, func(p NameDesc) bool { return p.Name == LabelSize && p.Type == "" })
		}
		d := link
		d.Desc, d.Parts = joinNonEmpty("\n\n", desc, cmp.Or(fluff[strings.ToLower(name+"|"+r.Source)], fluff[fluffKey])), append(typed, parts...)
		p.add(enum.KindSpecies, name, d, joinText(d.Desc, partsText(d.Parts)))
	}
	recs, err := p.records("races.json", keyRace)
	if err != nil {
		return err
	}
	subs, err := p.records("races.json", keySubrace)
	if err != nil {
		return err
	}
	// The races of other books our subraces build on come bundled; a race
	// reprinted in 2024 is named "Elf (2014)" beside the reprint.
	used := map[string]bool{}
	for _, rec := range subs {
		if s, err := ftDecode[ftRace](rec); err == nil && s.Name != "" {
			used[strings.ToLower(s.RaceName+"|"+s.RaceSource)] = true
		}
	}
	names := map[string]string{} // race key → species name
	bundled, err := ftCopyTargets(p.files, "races.json", keyRace)
	if err != nil {
		return err
	}
	for _, rec := range bundled {
		r, err := ftDecode[ftRace](rec)
		if err != nil || !used[r.key()] {
			continue
		}
		races[r.key()] = r
		names[r.key()] = r.Name
		if ftReprintedSameName(r.Name, r.ReprintedAs) {
			names[r.key()] += " (2014)"
		}
		desc, parts := ftParts(r.Entries)
		norm(r, names[r.key()], parts, desc, r.key(), TextData{})
	}
	// A race's versions (2024 lineages) follow it and show its fluff.
	for _, rec := range recs {
		base, err := ftDecode[ftRace](rec)
		if err != nil {
			p.skipped++
			continue
		}
		races[base.key()] = base
		versions, skipped := ftVersions(rec)
		p.skipped += skipped
		for i, v := range append([]ftRec{rec}, versions...) {
			r, err := ftDecode[ftRace](v)
			if err != nil {
				p.skipped++
				continue
			}
			desc, parts := ftParts(r.Entries)
			var link TextData
			if i > 0 {
				link.VersionOf = base.Name
			}
			norm(r, r.Name, parts, desc, base.key(), link)
		}
	}
	// 5etools shows a subrace as "Race (Subrace)", complete: the race's
	// traits, those the subrace overwrites replaced in place, then its own.
	// A subrace without a name is the race itself.
	for _, rec := range subs {
		s, err := ftDecode[ftRace](rec)
		if err != nil {
			p.skipped++
			continue
		}
		if s.Name == "" {
			continue
		}
		parent := races[strings.ToLower(s.RaceName+"|"+s.RaceSource)]
		if s.Size == nil {
			s.Size = parent.Size
		}
		if s.Speed == nil {
			s.Speed = parent.Speed
		}
		if !s.Overwrite.Ability {
			s.Ability = ftMergeAbility(parent.Ability, s.Ability)
		}
		desc, parts := ftParts(ftOverwrite(parent.Entries, s.Entries))
		// "Elf (High)", not "Elf (2014) (High)"; the link names the bundled base.
		base := cmp.Or(names[parent.key()], s.RaceName)
		norm(s, s.RaceName+" ("+strings.TrimPrefix(s.Name, "Variant; ")+")", parts, desc, parent.key(), TextData{SubraceOf: Slugify(base)})
	}
	return nil
}

// ftReprintedSameName reports whether a record was reprinted under its own
// name (Elf), which would share its slug; Half-Elf became Khoravar.
func ftReprintedSameName(name string, raw json.RawMessage) bool {
	var uids []json.RawMessage
	_ = json.Unmarshal(raw, &uids)
	for _, u := range uids {
		var s string
		var o struct {
			UID string `json:"uid"`
		}
		if json.Unmarshal(u, &s) != nil && json.Unmarshal(u, &o) == nil {
			s = o.UID
		}
		if before, _, _ := strings.Cut(s, "|"); strings.EqualFold(before, name) {
			return true
		}
	}
	return false
}

// ftOverwrite is a race's entries with a subrace's: an entry whose
// data.overwrite names one of the race's takes its place, the rest follow.
func ftOverwrite(race, sub []json.RawMessage) []json.RawMessage {
	out := slices.Clone(race)
	for _, e := range sub {
		var o struct {
			Data struct {
				Overwrite string `json:"overwrite"`
			} `json:"data"`
		}
		_ = json.Unmarshal(e, &o)
		i := -1
		if o.Data.Overwrite != "" {
			i = slices.IndexFunc(out, func(r json.RawMessage) bool {
				var n ftNamed
				return json.Unmarshal(r, &n) == nil && n.Name == o.Data.Overwrite
			})
		}
		if i >= 0 {
			out[i] = e
		} else {
			out = append(out, e)
		}
	}
	return out
}

// ftMergeAbility adds a subrace's fixed increases to each of its race's
// options (Elf's Dexterity +2 and High Elf's Intelligence +1).
func ftMergeAbility(race, sub []map[string]json.RawMessage) []map[string]json.RawMessage {
	if len(race) == 0 || len(sub) == 0 {
		return append(slices.Clone(race), sub...)
	}
	out := make([]map[string]json.RawMessage, len(race))
	for i, r := range race {
		out[i] = maps.Clone(r)
		maps.Copy(out[i], sub[0])
	}
	return out
}

// ftRaceAbility is a 2014 race's ability field as display text: fixed
// increases and choices, the options joined by "or"; "" if it has none.
func ftRaceAbility(opts []map[string]json.RawMessage) string {
	names := func(from []string) string {
		if len(from) == 0 || len(from) == len(ftAbilityNames) {
			return "any ability"
		}
		out := make([]string, len(from))
		for i, k := range from {
			out[i] = cmp.Or(ftAbilityNames[k], k)
		}
		return strings.Join(out, ", ")
	}
	var alts []string
	for _, o := range opts {
		var parts []string
		for _, k := range dnd5e.Abilities {
			var by int
			if json.Unmarshal(o[k], &by) == nil && by != 0 {
				parts = append(parts, fmt.Sprintf("%s %+d", ftAbilityNames[k], by))
			}
		}
		var c struct {
			From     []string `json:"from"`
			Count    int      `json:"count"`
			Amount   int      `json:"amount"`
			Weighted *struct {
				From    []string `json:"from"`
				Weights []int    `json:"weights"`
			} `json:"weighted"`
		}
		if json.Unmarshal(o[keyChoose], &c) == nil {
			switch {
			case c.Weighted != nil && len(c.Weighted.Weights) > 0:
				var ws []string
				for _, w := range c.Weighted.Weights {
					ws = append(ws, fmt.Sprintf("%+d", w))
				}
				parts = append(parts, fmt.Sprintf("%s to different scores among %s", strings.Join(ws, " and "), names(c.Weighted.From)))
			case len(c.From) > 0:
				n := cmp.Or(c.Count, 1)
				parts = append(parts, fmt.Sprintf("%+d to %d of %s", cmp.Or(c.Amount, 1), n, names(c.From)))
			}
		}
		if len(parts) > 0 {
			alts = append(alts, strings.Join(parts, ", "))
		}
	}
	return strings.Join(alts, "; or ")
}

func (p *ftBook) backgrounds() error {
	fluff, err := p.fluff("fluff-backgrounds.json", keyBackgroundFluff)
	if err != nil {
		return err
	}
	return each(p, "backgrounds.json", keyBackground, func(b ftNamed) {
		desc, parts := ftParts(b.Entries)
		d := TextData{Desc: joinNonEmpty("\n\n", desc, fluff[b.key()]), Parts: parts}
		p.add(enum.KindBackground, b.Name, d, joinText(d.Desc, partsText(d.Parts)))
	})
}

type ftFeat struct {
	ftNamed
	Category     string                       `json:"category"`
	Prerequisite []map[string]json.RawMessage `json:"prerequisite"`
	Ability      []map[string]json.RawMessage `json:"ability"`
}

var ftFeatCategories = map[string]string{
	"G": "General", "O": "Origin", "FS": "Fighting Style", "FS:P": "Fighting Style (Paladin)",
	"FS:R": "Fighting Style (Ranger)", "EB": "Epic Boon",
}

func (p *ftBook) feats() error {
	return each(p, "feats.json", keyFeat, func(f ftFeat) {
		desc, parts := ftParts(f.Entries)
		const asi = "Ability Score Increase"
		if inc := ftFeatIncrease(f.Ability, f.Category); inc != "" && !strings.Contains(joinText(desc, partsText(parts)), asi) {
			parts = append(parts, NameDesc{Name: asi, Desc: inc})
		}
		d := TextData{Desc: desc, Parts: parts, Prerequisite: ftPrerequisite(f.Prerequisite), Category: cmp.Or(ftFeatCategories[f.Category], f.Category)}
		p.add(enum.KindFeat, f.Name, d, joinText(d.Desc, partsText(d.Parts)))
	})
}

// ftFeatIncrease is a feat's ability score increase as the sentence the
// sheet wizard reads, or "" if it has none.
func ftFeatIncrease(abilities []map[string]json.RawMessage, category string) string {
	for _, a := range abilities {
		var hidden bool
		_ = json.Unmarshal(a[keyHidden], &hidden)
		// Hidden choices (Ability Score Improvement) are already in the feat's text.
		if hidden {
			continue
		}
		limit := 20
		if category == "EB" {
			limit = 30
		}
		_ = json.Unmarshal(a[keyMax], &limit)
		var choose struct {
			From   []string        `json:"from"`
			Amount int             `json:"amount"`
			Entry  json.RawMessage `json:"entry"`
		}
		if json.Unmarshal(a[keyChoose], &choose) == nil && len(choose.From) > 0 {
			who := "one ability score of your choice"
			if n := len(choose.From); n < len(ftAbilityNames) {
				names := make([]string, n)
				for i, k := range choose.From {
					names[i] = ftAbilityNames[k]
				}
				who = "your " + names[n-1] + " score"
				if n > 1 {
					who = "your " + strings.Join(names[:n-1], ", ") + " or " + names[n-1] + " score"
				}
			}
			s := fmt.Sprintf("Increase %s by %d, to a maximum of %d.", who, cmp.Or(choose.Amount, 1), limit)
			// The book's own limit (Resilient: an ability without save proficiency) follows the sentence the wizard reads.
			if choose.Entry != nil {
				if before, _, _ := strings.Cut(ftEntry1(choose.Entry), "Increase"); strings.TrimSpace(before) != "" {
					s += " " + strings.TrimSpace(before)
				}
			}
			return s
		}
		for _, k := range slices.Sorted(maps.Keys(a)) {
			var by int
			if name, ok := ftAbilityNames[k]; ok && json.Unmarshal(a[k], &by) == nil {
				return fmt.Sprintf("Increase your %s score by %d, to a maximum of %d.", name, by, limit)
			}
		}
	}
	return ""
}

// ftPrerequisite is a feat's prerequisites as text: the parts of one
// option joined by commas, the options by "or".
func ftPrerequisite(opts []map[string]json.RawMessage) string {
	var out []string
	for _, o := range opts {
		var parts []string
		var level int
		var lv struct {
			Level int     `json:"level"`
			Class ftNamed `json:"class"`
		}
		if json.Unmarshal(o[keyLevel], &level) == nil && level > 0 {
			parts = append(parts, fmt.Sprintf("Level %d+", level))
		} else if json.Unmarshal(o[keyLevel], &lv) == nil && lv.Level > 0 {
			parts = append(parts, strings.TrimSpace(fmt.Sprintf("Level %d+ %s", lv.Level, lv.Class.Name)))
		}
		var abilities []map[string]int
		_ = json.Unmarshal(o[keyAbility], &abilities)
		var ab []string
		for _, a := range abilities {
			for _, k := range slices.Sorted(maps.Keys(a)) {
				ab = append(ab, fmt.Sprintf("%s %d+", ftAbilityNames[k], a[k]))
			}
		}
		parts = append(parts, strings.Join(ab, " or "))
		var named []ftNamed
		for _, k := range []string{keyRace, keyBackground} {
			named = named[:0]
			_ = json.Unmarshal(o[k], &named)
			var names []string
			for _, n := range named {
				names = append(names, label(n.Name))
			}
			parts = append(parts, strings.Join(names, " or "))
		}
		var feats []string
		_ = json.Unmarshal(o[keyFeat], &feats)
		for i, f := range feats {
			feats[i] = label(ftCode(f))
		}
		parts = append(parts, strings.Join(feats, ", "))
		for _, k := range []string{keySpellcasting, keySpellcasting2020, keySpellcastingFeat, keySpellcastingPrep} {
			if o[k] != nil {
				parts = append(parts, "Spellcasting or Pact Magic Feature")
				break
			}
		}
		var profs []map[string]string
		_ = json.Unmarshal(o[keyProficiency], &profs)
		for _, pr := range profs {
			if a := pr[keyArmor]; a != "" {
				parts = append(parts, label(a)+" Armor Training")
			}
			if w := pr[keyWeapon]; w != "" {
				parts = append(parts, label(w)+" Weapon Proficiency")
			}
		}
		var other string
		_ = json.Unmarshal(o[keyOther], &other)
		parts = append(parts, plainTags(other))
		if s := joinNonEmpty(", ", parts...); s != "" {
			out = append(out, s)
		}
	}
	return strings.Join(slices.Compact(out), " or ")
}

// conditions: conditions and statuses (Concentration, Surprised) are
// conditions; diseases are rules in the "Disease" category.
func (p *ftBook) conditions() error {
	for _, c := range []struct {
		key, category string
		kind          enum.CompendiumKind
	}{{keyCondition, "", enum.KindCondition}, {keyStatus, "Status", enum.KindCondition}, {keyDisease, "Disease", enum.KindRule}} {
		if err := each(p, "conditionsdiseases.json", c.key, func(n ftNamed) {
			d := TextData{Desc: ftEntries(n.Entries), Category: c.category}
			p.add(c.kind, n.Name, d, d.Desc)
		}); err != nil {
			return err
		}
	}
	return nil
}

var ftRuleTypes = map[string]string{"C": "Core", "O": "Optional", "V": "Variant", "VO": "Variant optional", "VV": "Variant variant"}

func (p *ftBook) rules() error {
	if err := each(p, "variantrules.json", keyVariantRule, func(r struct {
		ftNamed
		RuleType string `json:"ruleType"`
	}) {
		d := TextData{Desc: ftEntries(r.Entries), Category: ftRuleTypes[r.RuleType]}
		p.add(enum.KindRule, r.Name, d, d.Desc)
	}); err != nil {
		return err
	}
	// items-base.json holds every book's mastery properties: only the book's own count.
	return each(p, "items-base.json", keyItemMastery, func(m ftNamed) {
		if strings.EqualFold(m.Source, p.id) {
			d := TextData{Desc: ftEntries(m.Entries), Category: masteryCategory}
			p.add(enum.KindRule, m.Name, d, d.Desc)
		}
	})
}

// ftPartTypes are list item names the character wizard reads as typed parts.
var ftPartTypes = map[string]string{
	LabelAbilityScores: PartAbilities, LabelFeat: PartFeat, LabelSkillProficiency: PartSkills,
	LabelToolProficiency: PartTools, LabelToolProficiencies: PartTools, LabelLanguages: PartLanguages, LabelEquipment: PartEquipment,
}

// ftParts splits entries into a description and named parts: named
// entries become parts, list items named like a background benefit typed
// parts in plain text (the wizard parses them), the rest stays in desc.
func ftParts(raws []json.RawMessage) (string, []NameDesc) {
	var desc []string
	var parts []NameDesc
	for _, raw := range raws {
		var o struct {
			Type    string            `json:"type"`
			Name    string            `json:"name"`
			Entries []json.RawMessage `json:"entries"`
			Items   []json.RawMessage `json:"items"`
		}
		_ = json.Unmarshal(raw, &o)
		switch {
		case o.Type == keyList:
			var rest []string
			for _, it := range o.Items {
				var item struct {
					Name  string            `json:"name"`
					Entry json.RawMessage   `json:"entry"`
					Items []json.RawMessage `json:"entries"`
				}
				_ = json.Unmarshal(it, &item)
				name := strings.TrimSuffix(strings.TrimSpace(plainTags(item.Name)), ":")
				if typ, ok := ftPartTypes[name]; ok {
					body := ftEntries(item.Items)
					if item.Entry != nil {
						body = ftEntry1(item.Entry)
					}
					parts = append(parts, NameDesc{Name: name, Desc: plainText(body), Type: typ})
				} else {
					rest = append(rest, "- "+ftEntry1(it))
				}
			}
			desc = append(desc, strings.Join(rest, "\n"))
		case o.Name != "" && o.Entries != nil:
			parts = append(parts, NameDesc{Name: plainTags(o.Name), Desc: ftEntries(o.Entries)})
		default:
			desc = append(desc, ftEntry1(raw))
		}
	}
	return joinNonEmpty("\n\n", desc...), parts
}

// ftTable is a 5etools table as Markdown.
func ftTable(caption string, labels []string, rows [][]json.RawMessage) string {
	var b strings.Builder
	if caption != "" {
		b.WriteString("Table: " + plainTags(caption) + "\n\n")
	}
	cols := len(labels)
	for _, r := range rows {
		cols = max(cols, len(r))
	}
	head := make([]string, cols)
	for i := range head {
		if i < len(labels) {
			head[i] = plainTags(labels[i])
		}
	}
	b.WriteString("| " + strings.Join(head, " | ") + " |\n|" + strings.Repeat("---|", cols) + "\n")
	for _, r := range rows {
		cells := make([]string, cols)
		for i, c := range r {
			cells[i] = ftCell(c)
		}
		b.WriteString("| " + strings.Join(cells, " | ") + " |\n")
	}
	return strings.TrimSpace(b.String())
}

// ftCell is a table cell: text, a roll ({"roll":{"min":1,"max":5}}) or an entry.
func ftCell(raw json.RawMessage) string {
	var c struct {
		Roll *struct {
			Exact *int `json:"exact"`
			Min   int  `json:"min"`
			Max   int  `json:"max"`
		} `json:"roll"`
	}
	if json.Unmarshal(raw, &c) == nil && c.Roll != nil {
		if c.Roll.Exact != nil {
			return strconv.Itoa(*c.Roll.Exact)
		}
		return fmt.Sprintf("%d–%d", c.Roll.Min, c.Roll.Max)
	}
	// A "|" would end the cell; our tokens keep theirs.
	s := reToken.ReplaceAllStringFunc(strings.ReplaceAll(ftEntry1(raw), "\n", " "), func(t string) string {
		return strings.Replace(t, "|", tokenSep, 1)
	})
	return strings.ReplaceAll(strings.ReplaceAll(s, "|", "/"), tokenSep, "|")
}

// ftRecords reads one list of a file with copies resolved (against the list
// and the file's copy targets), returning how many could not be. A missing
// file or list is empty.
func ftRecords(files map[string][]byte, file, key string) ([]ftRec, int, error) {
	body, ok := files[file]
	if !ok {
		return nil, 0, nil
	}
	var top map[string]json.RawMessage
	if err := json.Unmarshal(body, &top); err != nil {
		return nil, 0, fmt.Errorf("5etools %s: %w", file, err)
	}
	var recs []ftRec
	if raw, ok := top[key]; ok {
		if err := json.Unmarshal(raw, &recs); err != nil {
			return nil, 0, fmt.Errorf("5etools %s %s: %w", file, key, err)
		}
	}
	// Other books' copy targets resolve copies but are not the book's own.
	var targets map[string][]ftRec
	if raw, ok := top[copyTargetsKey]; ok {
		if err := json.Unmarshal(raw, &targets); err != nil {
			return nil, 0, fmt.Errorf("5etools %s %s: %w", file, copyTargetsKey, err)
		}
	}
	index := map[string]ftRec{}
	for _, r := range append(targets[key], recs...) {
		if n, err := ftDecode[ftCopyRef](r); err == nil {
			index[n.ftNamed.key()] = r
			index[n.key()] = r
		}
	}
	out := make([]ftRec, 0, len(recs))
	skipped := 0
	for _, r := range recs {
		if res, err := ftResolve(index, r, 0); err == nil {
			out = append(out, res)
		} else {
			skipped++
		}
	}
	return out, skipped, nil
}

// ftCopyTargets are the records of another book a file carries for one list.
func ftCopyTargets(files map[string][]byte, file, key string) ([]ftRec, error) {
	body, ok := files[file]
	if !ok {
		return nil, nil
	}
	var top struct {
		Targets map[string][]ftRec `json:"_copyTargets"`
	}
	if err := json.Unmarshal(body, &top); err != nil {
		return nil, fmt.Errorf("5etools %s %s: %w", file, copyTargetsKey, err)
	}
	return top.Targets[key], nil
}

func ftDecode[T any](r ftRec) (T, error) {
	var v T
	b, err := json.Marshal(r)
	if err == nil {
		err = json.Unmarshal(b, &v)
	}
	return v, err
}

// ftNotCopied are fields a copy does not inherit (its own pictures and
// page), unless _preserve names them.
var ftNotCopied = []string{keyPage, keyHasFluff, keyHasFluffImages, keyHasToken, keySrd, keySrd52, keyBasicRules, keyBasicRules2024, keyReprintedAs, keyOtherSources, keyVersions}

var errCopy = errors.New("unresolved _copy")

// ftCopyRef is what a _copy names. Class fields count too: a book keeps its
// subclasses once per class edition under the same name and source.
type ftCopyRef struct {
	ftNamed
	ClassName         string `json:"className"`
	ClassSource       string `json:"classSource"`
	ShortName         string `json:"shortName"`
	SubclassShortName string `json:"subclassShortName"`
	SubclassSource    string `json:"subclassSource"`
	Level             any    `json:"level"`
}

func (c ftCopyRef) key() string {
	return strings.ToLower(fmt.Sprint(c.Name, "|", c.Source, "|", c.ClassName, "|", c.ClassSource, "|", c.ShortName, "|", c.SubclassShortName, "|", c.SubclassSource, "|", c.Level))
}

// ftResolve applies a record's _copy: the target's fields, then its own,
// then _mod.
func ftResolve(index map[string]ftRec, r ftRec, depth int) (ftRec, error) {
	raw, ok := r[keyCopy]
	if !ok {
		return r, nil
	}
	var c struct {
		ftCopyRef
		Mod      map[string]json.RawMessage `json:"_mod"`
		Preserve map[string]bool            `json:"_preserve"`
	}
	if depth > 4 || json.Unmarshal(raw, &c) != nil {
		return nil, errCopy
	}
	// A _copy that names fewer fields than its target falls back to name|source.
	target, ok := index[c.key()]
	if !ok {
		target, ok = index[c.ftNamed.key()]
	}
	if !ok {
		return nil, errCopy
	}
	base, err := ftResolve(index, target, depth+1)
	if err != nil {
		return nil, err
	}
	out := maps.Clone(base)
	for _, k := range ftNotCopied {
		if !c.Preserve[k] {
			delete(out, k)
		}
	}
	for k, v := range r {
		if k != keyCopy {
			out[k] = v
		}
	}
	for _, k := range slices.Sorted(maps.Keys(c.Mod)) {
		if err := ftMod(out, k, c.Mod[k]); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ftVersions are a record's _versions as records of their own: each copies r
// with its own fields and _mod. An _abstract version is a template, filled
// once per implementation with its _variables ("{{color}}") and fields.
// "Elf; Drow Lineage" is named "Elf (Drow Lineage)", as subraces are.
func ftVersions(r ftRec) ([]ftRec, int) {
	raw, ok := r[keyVersions]
	if !ok {
		return nil, 0
	}
	parent, err := ftDecode[ftNamed](r)
	var vs []ftRec
	if err != nil || json.Unmarshal(raw, &vs) != nil {
		return nil, 1
	}
	var all []ftRec
	skipped := 0
	for _, v := range vs {
		abstract, ok := v[keyAbstract]
		if !ok {
			all = append(all, v)
			continue
		}
		var impls []ftRec
		_ = json.Unmarshal(v[keyImplementations], &impls)
		for _, impl := range impls {
			var vars map[string]string
			_ = json.Unmarshal(impl[keyVariables], &vars)
			body := string(abstract)
			for k, val := range vars {
				quoted, _ := json.Marshal(val) // the value lands inside JSON strings
				body = strings.ReplaceAll(body, "{{"+k+"}}", string(quoted[1:len(quoted)-1]))
			}
			var ver ftRec
			if json.Unmarshal([]byte(body), &ver) != nil {
				skipped++
				continue
			}
			for k, x := range impl {
				if k != keyVariables {
					ver[k] = x
				}
			}
			all = append(all, ver)
		}
	}
	index := map[string]ftRec{parent.key(): r}
	var out []ftRec
	for _, v := range all {
		ref, _ := json.Marshal(map[string]json.RawMessage{keyMod: v[keyMod], keyName: r[keyName], keySource: r[keySource]})
		ver := maps.Clone(v)
		delete(ver, keyMod)
		ver[keyCopy] = ref
		res, err := ftResolve(index, ver, 0)
		if err != nil {
			skipped++
			continue
		}
		var name string
		if json.Unmarshal(res[keyName], &name) == nil {
			if base, version, ok := strings.Cut(name, "; "); ok {
				res[keyName], _ = json.Marshal(base + " (" + version + ")")
			}
		}
		out = append(out, res)
	}
	return out, skipped
}

type ftModOp struct {
	Mode    string          `json:"mode"`
	Replace json.RawMessage `json:"replace"`
	With    string          `json:"with"`
	Flags   string          `json:"flags"`
	Items   json.RawMessage `json:"items"`
	Index   int             `json:"index"`
	Names   json.RawMessage `json:"names"`
}

// reJSGroup is a JavaScript replacement group ($1), ${1} in Go.
var reJSGroup = regexp.MustCompile(`\$(\d)`)

// ftMod applies one key's _mod operations; "*" is every field. Unknown
// modes are an error, so the record is skipped rather than wrong.
func ftMod(rec ftRec, key string, raw json.RawMessage) error {
	var s string
	if json.Unmarshal(raw, &s) == nil {
		if s != keyRemove {
			return errCopy
		}
		delete(rec, key)
		return nil
	}
	var ops []ftModOp
	if json.Unmarshal(raw, &ops) != nil {
		var op ftModOp
		if err := json.Unmarshal(raw, &op); err != nil {
			return err
		}
		ops = []ftModOp{op}
	}
	for _, op := range ops {
		if op.Mode == modReplaceTxt {
			var pat string
			if json.Unmarshal(op.Replace, &pat) != nil || strings.Trim(op.Flags, "gi") != "" {
				return errCopy
			}
			if strings.Contains(op.Flags, "i") {
				pat = "(?i)" + pat
			}
			re, err := regexp.Compile(pat)
			if err != nil {
				return errCopy
			}
			with := reJSGroup.ReplaceAllString(op.With, "$${$1}")
			for k, v := range rec {
				if (key == keyModAll && k != keyName && k != keySource && !strings.HasPrefix(k, "_")) || k == key {
					var val any
					_ = json.Unmarshal(v, &val)
					rec[k], _ = json.Marshal(ftReplaceText(val, re, with))
				}
			}
			continue
		}
		var arr, items []any
		_ = json.Unmarshal(rec[key], &arr)
		if json.Unmarshal(op.Items, &items) != nil && op.Items != nil {
			var one any
			_ = json.Unmarshal(op.Items, &one)
			items = []any{one}
		}
		switch op.Mode {
		case modAppendArr:
			arr = append(arr, items...)
		case modPrependArr:
			arr = append(items, arr...)
		case modInsertArr:
			arr = slices.Insert(arr, min(max(op.Index, 0), len(arr)), items...)
		case modRemoveArr:
			var names []string
			if json.Unmarshal(op.Names, &names) != nil {
				var one string
				_ = json.Unmarshal(op.Names, &one)
				names = []string{one}
			}
			for _, it := range items {
				if s, ok := it.(string); ok {
					names = append(names, s)
				}
			}
			arr = slices.DeleteFunc(arr, func(v any) bool { return slices.Contains(names, ftName(v)) })
		case modReplaceArr:
			var name string
			_ = json.Unmarshal(op.Replace, &name)
			i := slices.IndexFunc(arr, func(v any) bool { return ftName(v) == name })
			if i < 0 {
				return errCopy
			}
			arr = slices.Replace(arr, i, i+1, items...)
		default:
			return errCopy
		}
		rec[key], _ = json.Marshal(arr)
	}
	return nil
}

// ftName is a list element's name: the string itself or its "name".
func ftName(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	if m, ok := v.(map[string]any); ok {
		s, _ := m[keyName].(string)
		return s
	}
	return ""
}

// ftReplaceText runs re over every string in v.
func ftReplaceText(v any, re *regexp.Regexp, with string) any {
	switch x := v.(type) {
	case string:
		return re.ReplaceAllString(x, with)
	case []any:
		for i := range x {
			x[i] = ftReplaceText(x[i], re, with)
		}
	case map[string]any:
		for k := range x {
			if k != keyType {
				x[k] = ftReplaceText(x[k], re, with)
			}
		}
	}
	return v
}
