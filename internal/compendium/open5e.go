package compendium

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// This file reads the Open5e data repository's v2 layout
// (data/v2/<publisher>/<document>/*.json, e.g. srd-2024). Each file is a
// Django fixture: a JSON array of {"model": "api_v2.x", "pk": ..., "fields":
// {...}}. Related rows (creature actions, class feature levels, benefits)
// sit in their own files and point at their parent by pk.

type fixture struct {
	Model  string          `json:"model"`
	PK     json.RawMessage `json:"pk"`
	Fields json.RawMessage `json:"fields"`
}

// fixtures groups the records of one document by model name.
type fixtures map[string][]fixture

// errNotOpen5e means the files are not Open5e fixtures.
var errNotOpen5e = errors.New("not an Open5e fixture set")

func parseFixtures(files map[string][]byte) (fixtures, error) {
	out := fixtures{}
	for name, body := range files {
		var recs []fixture
		if err := json.Unmarshal(body, &recs); err != nil {
			return nil, fmt.Errorf("%s: %w", name, errNotOpen5e)
		}
		for _, r := range recs {
			if !strings.HasPrefix(r.Model, "api_v2.") {
				return nil, fmt.Errorf("%s: model %q: %w", name, r.Model, errNotOpen5e)
			}
			out[strings.TrimPrefix(r.Model, "api_v2.")] = append(out[strings.TrimPrefix(r.Model, "api_v2.")], r)
		}
	}
	return out, nil
}

func (r fixture) pk() string {
	var s string
	if json.Unmarshal(r.PK, &s) == nil {
		return s
	}
	return string(r.PK)
}

func decode[T any](r fixture) (T, error) {
	var v T
	err := json.Unmarshal(r.Fields, &v)
	if err != nil {
		err = fmt.Errorf("%s %s: %w", r.Model, r.pk(), err)
	}
	return v, err
}

// children indexes a child model's records by their parent pk.
func children[T any](fx fixtures, model string, parentOf func(T) string) (map[string][]T, error) {
	out := map[string][]T{}
	for _, r := range fx[model] {
		v, err := decode[T](r)
		if err != nil {
			return nil, err
		}
		p := parentOf(v)
		out[p] = append(out[p], v)
	}
	return out, nil
}

type open5eDoc struct {
	Name        string   `json:"name"`
	DisplayName string   `json:"display_name"`
	Author      string   `json:"author"`
	GameSystem  string   `json:"gamesystem"`
	Licenses    []string `json:"licenses"`
	Permalink   string   `json:"permalink"`
}

var licenseNames = map[string]string{
	"cc-by-40": "CC-BY-4.0",
	"ogl10a":   "OGL-1.0a",
}

// parseOpen5e normalizes one Open5e document directory.
func parseOpen5e(sourceKey string, files map[string][]byte) (Source, []Entry, error) {
	fx, err := parseFixtures(files)
	if err != nil {
		return Source{}, nil, err
	}
	docs := fx["document"]
	if len(docs) != 1 {
		return Source{}, nil, fmt.Errorf("expected one Document record, found %d: %w", len(docs), errNotOpen5e)
	}
	doc, err := decode[open5eDoc](docs[0])
	if err != nil {
		return Source{}, nil, err
	}
	src := Source{Key: sourceKey, Edition: strings.TrimPrefix(doc.GameSystem, "5e-")}
	var licenses []string
	for _, l := range doc.Licenses {
		licenses = append(licenses, cmp.Or(licenseNames[l], l))
	}
	src.License = strings.Join(licenses, ", ")
	src.Attribution = fmt.Sprintf("This work includes material from %s by %s, available at %s, licensed under %s.",
		doc.Name, doc.Author, doc.Permalink, src.License)

	n := normalizer{fx: fx, prefix: docs[0].pk() + "_"}
	var entries []Entry
	for _, step := range []func() ([]Entry, error){
		n.monsters, n.spells, n.items, n.classes, n.species, n.backgrounds, n.feats, n.conditions, n.rules,
	} {
		es, err := step()
		if err != nil {
			return Source{}, nil, err
		}
		entries = append(entries, es...)
	}
	return src, entries, nil
}

type normalizer struct {
	fx     fixtures
	prefix string // the document pk plus "_", stripped from slugs
}

func (n normalizer) slug(pk string) string { return strings.TrimPrefix(pk, n.prefix) }

func joinText(parts ...string) string {
	var b strings.Builder
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			if b.Len() > 0 {
				b.WriteString("\n")
			}
			b.WriteString(p)
		}
	}
	return b.String()
}

func partsText(parts []NameDesc) string {
	var b strings.Builder
	for _, p := range parts {
		b.WriteString(p.Name + ". " + p.Desc + "\n")
	}
	return b.String()
}

type o5eCreature struct {
	Name           string      `json:"name"`
	Size           string      `json:"size"`
	Type           string      `json:"type"`
	Alignment      string      `json:"alignment"`
	AC             int         `json:"armor_class"`
	ACDetail       string      `json:"armor_detail"`
	HP             int         `json:"hit_points"`
	HitDice        string      `json:"hit_dice"`
	Initiative     *int        `json:"initiative_bonus"`
	CR             looseString `json:"challenge_rating"`
	Walk           *int        `json:"walk"`
	Fly            *int        `json:"fly"`
	Swim           *int        `json:"swim"`
	Climb          *int        `json:"climb"`
	Burrow         *int        `json:"burrow"`
	Darkvision     *int        `json:"darkvision_range"`
	Blindsight     *int        `json:"blindsight_range"`
	Tremorsense    *int        `json:"tremorsense_range"`
	Truesight      *int        `json:"truesight_range"`
	Passive        *int        `json:"passive_perception"`
	Languages      string      `json:"languages_desc"`
	Resistances    string      `json:"damage_resistances_display"`
	Immunities     string      `json:"damage_immunities_display"`
	CondImmunities string      `json:"condition_immunities_display"`
	Environments   []string    `json:"environments"`
}

type o5eChild struct {
	Parent string `json:"parent"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	// Type is a trait or action type, depending on the model.
	Type       *string `json:"type"`
	ActionType string  `json:"action_type"`
	Order      int     `json:"order_in_statblock"`
	SortOrder  int     `json:"order"`
}

func childParent(c o5eChild) string { return c.Parent }

var abilityKeys = map[string]string{
	"strength": "str", "dexterity": "dex", "constitution": "con",
	"intelligence": "int", "wisdom": "wis", "charisma": "cha",
}

func (n normalizer) monsters() ([]Entry, error) {
	actions, err := children(n.fx, "creatureaction", childParent)
	if err != nil {
		return nil, err
	}
	traits, err := children(n.fx, "creaturetrait", childParent)
	if err != nil {
		return nil, err
	}
	var out []Entry
	for _, r := range n.fx["creature"] {
		c, err := decode[o5eCreature](r)
		if err != nil {
			return nil, err
		}
		var raw map[string]json.RawMessage
		_ = json.Unmarshal(r.Fields, &raw)

		d := MonsterData{
			Size: c.Size, Type: c.Type, Alignment: c.Alignment, AC: c.AC, ACDetail: c.ACDetail,
			HP: c.HP, HitDice: c.HitDice, CR: formatCR(string(c.CR)), Languages: c.Languages,
			Resistance: c.Resistances, Immunity: joinNonEmpty("; ", c.Immunities, c.CondImmunities),
			Speed: map[string]int{}, Abilities: map[string]int{}, Saves: map[string]int{}, Skills: map[string]int{},
		}
		if c.Initiative != nil {
			d.Initiative = *c.Initiative
		}
		for k, v := range map[string]*int{"walk": c.Walk, "fly": c.Fly, "swim": c.Swim, "climb": c.Climb, "burrow": c.Burrow} {
			if v != nil && *v > 0 {
				d.Speed[k] = *v
			}
		}
		for long, short := range abilityKeys {
			if v := intField(raw, "ability_score_"+long); v != nil {
				d.Abilities[short] = *v
			}
			if v := intField(raw, "saving_throw_"+long); v != nil {
				d.Saves[short] = *v
			}
		}
		for k := range raw {
			if skill, ok := strings.CutPrefix(k, "skill_bonus_"); ok {
				if v := intField(raw, k); v != nil {
					d.Skills[strings.ReplaceAll(skill, "_", "-")] = *v
				}
			}
		}
		var senses []string
		for name, v := range map[string]*int{"Blindsight": c.Blindsight, "Darkvision": c.Darkvision, "Tremorsense": c.Tremorsense, "Truesight": c.Truesight} {
			if v != nil && *v > 0 {
				senses = append(senses, fmt.Sprintf("%s %d ft.", name, *v))
			}
		}
		slices.Sort(senses)
		if c.Passive != nil {
			senses = append(senses, fmt.Sprintf("Passive Perception %d", *c.Passive))
		}
		d.Senses = strings.Join(senses, ", ")

		for _, t := range traits[r.pk()] {
			d.Traits = append(d.Traits, NameDesc{Name: t.Name, Desc: t.Desc})
		}
		acts := actions[r.pk()]
		slices.SortStableFunc(acts, func(a, b o5eChild) int {
			return cmp.Or(cmp.Compare(actionRank(a.ActionType), actionRank(b.ActionType)), cmp.Compare(a.Order, b.Order))
		})
		for _, a := range acts {
			d.Actions = append(d.Actions, MonsterAction{Name: a.Name, Desc: a.Desc, Type: strings.ToLower(a.ActionType)})
		}
		out = append(out, Entry{Kind: KindMonster, Slug: n.slug(r.pk()), Name: c.Name, Data: d, Text: d.searchText()})
	}
	return out, nil
}

func actionRank(t string) int {
	return slices.Index([]string{"ACTION", "BONUS_ACTION", "REACTION", "LEGENDARY_ACTION"}, t)
}

func intField(raw map[string]json.RawMessage, key string) *int {
	var v *int
	if b, ok := raw[key]; ok {
		_ = json.Unmarshal(b, &v)
	}
	return v
}

func joinNonEmpty(sep string, parts ...string) string {
	return strings.Join(slices.DeleteFunc(parts, func(s string) bool { return strings.TrimSpace(s) == "" }), sep)
}

// formatCR turns "0.125" into "1/8" and "10.000" into "10".
func formatCR(cr string) string {
	f, err := strconv.ParseFloat(cr, 64)
	if err != nil {
		return cr
	}
	switch f {
	case 0.125:
		return "1/8"
	case 0.25:
		return "1/4"
	case 0.5:
		return "1/2"
	}
	return strconv.FormatFloat(f, 'f', -1, 64)
}

type o5eSpell struct {
	Name          string   `json:"name"`
	Desc          string   `json:"desc"`
	Level         int      `json:"level"`
	School        string   `json:"school"`
	HigherLevel   string   `json:"higher_level"`
	RangeText     string   `json:"range_text"`
	Ritual        bool     `json:"ritual"`
	CastingTime   string   `json:"casting_time"`
	Verbal        bool     `json:"verbal"`
	Somatic       bool     `json:"somatic"`
	Material      bool     `json:"material"`
	MaterialSpec  string   `json:"material_specified"`
	Duration      string   `json:"duration"`
	Concentration bool     `json:"concentration"`
	Classes       []string `json:"classes"`
}

// reFlatTable is a Markdown table some SRD spells carry on one line:
// "Table: Omens | Omen | Result | |---|---| | Weal | Good | then text".
var reFlatTable = regexp.MustCompile(`Table: ([^|]*?) (\|.*\|)`)

// unflattenTable puts a flattened table's rows back on their own lines.
// shortcut: one table per text, and an empty cell ("| |") would split its row; fine for the SRD's four.
func unflattenTable(s string) string {
	m := reFlatTable.FindStringSubmatchIndex(s)
	if m == nil {
		return s
	}
	table := strings.ReplaceAll(s[m[4]:m[5]], "| |", "|\n|")
	out := strings.TrimSpace(s[:m[0]]) + "\n\nTable: " + s[m[2]:m[3]] + "\n\n" + table
	if rest := strings.TrimSpace(s[m[1]:]); rest != "" {
		out += "\n\n" + rest
	}
	return strings.TrimSpace(out)
}

func (n normalizer) spells() ([]Entry, error) {
	var out []Entry
	for _, r := range n.fx["spell"] {
		s, err := decode[o5eSpell](r)
		if err != nil {
			return nil, err
		}
		s.Desc = unflattenTable(s.Desc)
		var comps []string
		for c, on := range map[string]bool{"V": s.Verbal, "S": s.Somatic, "M": s.Material} {
			if on {
				comps = append(comps, c)
			}
		}
		slices.SortFunc(comps, func(a, b string) int { return strings.Index("VSM", a) - strings.Index("VSM", b) })
		components := strings.Join(comps, ", ")
		if s.Material && s.MaterialSpec != "" {
			components += " (" + s.MaterialSpec + ")"
		}
		classes := make([]string, 0, len(s.Classes))
		for _, c := range s.Classes {
			classes = append(classes, n.slug(c))
		}
		out = append(out, Entry{
			Kind: KindSpell, Slug: n.slug(r.pk()), Name: s.Name,
			Data: SpellData{
				Level: s.Level, School: s.School, CastingTime: s.CastingTime, Range: s.RangeText,
				Components: components, Duration: s.Duration, Concentration: s.Concentration,
				Ritual: s.Ritual, Classes: classes, Desc: s.Desc, HigherLevel: s.HigherLevel,
			},
			Text: joinText(s.School, s.Desc, s.HigherLevel),
		})
	}
	return out, nil
}

type o5eItem struct {
	Name       string      `json:"name"`
	Desc       string      `json:"desc"`
	Category   string      `json:"category"`
	Rarity     string      `json:"rarity"`
	Attunement bool        `json:"requires_attunement"`
	Cost       looseString `json:"cost"`
	Weight     looseString `json:"weight"`
	Weapon     *string     `json:"weapon"`
	Armor      *string     `json:"armor"`
}

type o5eWeapon struct {
	Damage     string  `json:"damage_dice"`
	DamageType string  `json:"damage_type"`
	Range      float64 `json:"range"`
	LongRange  float64 `json:"long_range"`
	Simple     bool    `json:"is_simple"`
}

type o5eArmor struct {
	Base    int  `json:"ac_base"`
	AddDex  bool `json:"ac_add_dexmod"`
	DexCap  *int `json:"ac_cap_dexmod"`
	Stealth bool `json:"grants_stealth_disadvantage"`
	StrReq  *int `json:"strength_score_required"`
}

type o5eNamed struct {
	Name string `json:"name"`
}

type o5eAssignment struct {
	Weapon   string  `json:"weapon"`
	Property string  `json:"property"`
	Detail   *string `json:"detail"`
}

func byPK[T any](fx fixtures, model string) (map[string]T, error) {
	out := map[string]T{}
	for _, r := range fx[model] {
		v, err := decode[T](r)
		if err != nil {
			return nil, err
		}
		out[r.pk()] = v
	}
	return out, nil
}

func (n normalizer) items() ([]Entry, error) {
	weapons, err := byPK[o5eWeapon](n.fx, "weapon")
	if err != nil {
		return nil, err
	}
	armor, err := byPK[o5eArmor](n.fx, "armor")
	if err != nil {
		return nil, err
	}
	props, err := byPK[o5eNamed](n.fx, "weaponproperty")
	if err != nil {
		return nil, err
	}
	assigned, err := children(n.fx, "weaponpropertyassignment", func(a o5eAssignment) string { return a.Weapon })
	if err != nil {
		return nil, err
	}

	var out []Entry
	for _, model := range []string{"item", "magicitem"} {
		for _, r := range n.fx[model] {
			it, err := decode[o5eItem](r)
			if err != nil {
				return nil, err
			}
			d := ItemData{
				Category: it.Category, Rarity: it.Rarity, Attunement: it.Attunement,
				Cost: trimZeros(string(it.Cost)), Weight: trimZeros(string(it.Weight)), Desc: it.Desc,
			}
			if it.Weapon != nil {
				if w, ok := weapons[*it.Weapon]; ok {
					d.Damage = strings.TrimSpace(w.Damage + " " + w.DamageType)
					d.Martial = !w.Simple
					if w.Range > 0 {
						d.Range = fmt.Sprintf("%g/%g", w.Range, w.LongRange)
					}
				}
				var names []string
				for _, a := range assigned[*it.Weapon] {
					name := props[a.Property].Name
					if a.Detail != nil && *a.Detail != "" {
						name += " (" + *a.Detail + ")"
					}
					names = append(names, name)
				}
				slices.Sort(names)
				d.Properties = strings.Join(names, ", ")
			}
			if it.Armor != nil {
				if a, ok := armor[*it.Armor]; ok {
					d.Armor = armorData(*it.Armor, a)
					d.ArmorClass = armorText(*d.Armor)
				}
			}
			out = append(out, Entry{
				Kind: KindItem, Slug: n.slug(r.pk()), Name: it.Name, Data: d,
				Text: joinText(it.Category, it.Rarity, it.Desc),
			})
		}
	}
	return out, nil
}

// armorData is the armor's structured AC. The data has no armor type, so it
// follows from the Dexterity rule: none is heavy, a cap is medium, no cap
// light; the shield is the one piece that adds to AC.
func armorData(pk string, a o5eArmor) *ArmorData {
	d := &ArmorData{Base: a.Base, AddDex: a.AddDex, Stealth: a.Stealth}
	if a.DexCap != nil {
		d.DexCap = *a.DexCap
	}
	if a.StrReq != nil {
		d.StrReq = *a.StrReq
	}
	switch {
	case strings.HasSuffix(pk, "_shield"):
		d.Type = ArmorShield
	case !a.AddDex:
		d.Type = ArmorHeavy
	case d.DexCap > 0:
		d.Type = ArmorMedium
	default:
		d.Type = ArmorLight
	}
	return d
}

// armorText is an armor's AC as people read it.
func armorText(a ArmorData) string {
	s := strconv.Itoa(a.Base)
	switch {
	case a.Type == ArmorShield:
		s = "+" + s
	case a.AddDex && a.DexCap > 0:
		s += fmt.Sprintf(" + Dex modifier (max %d)", a.DexCap)
	case a.AddDex:
		s += " + Dex modifier"
	}
	if a.Stealth {
		s += ", Stealth disadvantage"
	}
	return s
}

func trimZeros(s string) string {
	if !strings.Contains(s, ".") {
		return s
	}
	s = strings.TrimRight(strings.TrimRight(s, "0"), ".")
	if s == "0" {
		return ""
	}
	return s
}

type o5eClass struct {
	Name       string   `json:"name"`
	HitDice    string   `json:"hit_dice"`
	CasterType string   `json:"caster_type"`
	Saves      []string `json:"saving_throws"`
	SubclassOf *string  `json:"subclass_of"`
}

type o5eFeature struct {
	Name        string  `json:"name"`
	Desc        string  `json:"desc"`
	Parent      string  `json:"parent"`
	FeatureType *string `json:"feature_type"`
	pk          string
}

type o5eFeatureItem struct {
	Level       int     `json:"level"`
	ColumnValue *string `json:"column_value"`
	Parent      string  `json:"parent"`
}

func (n normalizer) classes() ([]Entry, error) {
	// Features are grouped by class; each keeps its pk to find its level rows.
	features := map[string][]o5eFeature{}
	for _, r := range n.fx["classfeature"] {
		f, err := decode[o5eFeature](r)
		if err != nil {
			return nil, err
		}
		f.pk = r.pk()
		features[f.Parent] = append(features[f.Parent], f)
	}
	levels, err := children(n.fx, "classfeatureitem", func(i o5eFeatureItem) string { return i.Parent })
	if err != nil {
		return nil, err
	}

	var out []Entry
	for _, r := range n.fx["characterclass"] {
		c, err := decode[o5eClass](r)
		if err != nil {
			return nil, err
		}
		d := ClassData{
			HitDie: hitDie(c.HitDice), Caster: strings.ToLower(cmp.Or(c.CasterType, "none")),
			Saves: c.Saves, Table: map[string]map[int]string{},
		}
		kind := KindClass
		if c.SubclassOf != nil {
			kind, d.SubclassOf = KindSubclass, n.slug(*c.SubclassOf)
		}
		var text []string
		for _, f := range features[r.pk()] {
			ftype := ""
			if f.FeatureType != nil {
				ftype = *f.FeatureType
			}
			rows := levels[f.pk]
			switch ftype {
			case "PROFICIENCY_BONUS":
				continue // derived from level
			case "CORE_TRAITS_TABLE":
				d.Desc = f.Desc
			case "CLASS_TABLE_DATA", "SPELL_SLOTS":
				col := f.Name
				if ftype == "SPELL_SLOTS" {
					col = "Slots " + f.Name
				}
				d.Table[col] = map[int]string{}
				for _, row := range rows {
					if row.ColumnValue != nil {
						d.Table[col][row.Level] = *row.ColumnValue
					}
				}
			default:
				cf := ClassFeature{Name: f.Name, Desc: f.Desc}
				for _, row := range rows {
					cf.Levels = append(cf.Levels, row.Level)
				}
				slices.Sort(cf.Levels)
				d.Features = append(d.Features, cf)
				text = append(text, f.Name+". "+f.Desc)
			}
		}
		slices.SortStableFunc(d.Features, func(a, b ClassFeature) int {
			return cmp.Compare(firstOr(a.Levels, 99), firstOr(b.Levels, 99))
		})
		out = append(out, Entry{Kind: kind, Slug: n.slug(r.pk()), Name: c.Name, Data: d, Text: joinText(d.Desc, strings.Join(text, "\n"))})
	}
	return out, nil
}

func firstOr(s []int, def int) int {
	if len(s) == 0 {
		return def
	}
	return s[0]
}

func hitDie(s string) int {
	n, _ := strconv.Atoi(strings.TrimPrefix(strings.ToUpper(s), "D"))
	return n
}

type o5eText struct {
	Name         string  `json:"name"`
	Desc         string  `json:"desc"`
	Prerequisite string  `json:"prerequisite"`
	Type         string  `json:"type"`
	SubspeciesOf *string `json:"subspecies_of"`
}

// textEntries builds a TextData entry per record with child parts.
func (n normalizer) textEntries(kind Kind, model, childModel string) ([]Entry, error) {
	parts, err := children(n.fx, childModel, childParent)
	if err != nil {
		return nil, err
	}
	var out []Entry
	for _, r := range n.fx[model] {
		t, err := decode[o5eText](r)
		if err != nil {
			return nil, err
		}
		kids := parts[r.pk()]
		slices.SortStableFunc(kids, func(a, b o5eChild) int { return cmp.Compare(a.SortOrder, b.SortOrder) })
		d := TextData{Desc: t.Desc, Prerequisite: t.Prerequisite, Category: t.Type}
		for _, p := range kids {
			nd := NameDesc{Name: p.Name, Desc: p.Desc}
			if p.Type != nil {
				nd.Type = strings.ToLower(*p.Type)
			}
			d.Parts = append(d.Parts, nd)
		}
		out = append(out, Entry{Kind: kind, Slug: n.slug(r.pk()), Name: t.Name, Data: d, Text: joinText(t.Desc, partsText(d.Parts))})
	}
	return out, nil
}

func (n normalizer) species() ([]Entry, error) {
	return n.textEntries(KindSpecies, "species", "speciestrait")
}

func (n normalizer) backgrounds() ([]Entry, error) {
	return n.textEntries(KindBackground, "background", "backgroundbenefit")
}

func (n normalizer) feats() ([]Entry, error) {
	return n.textEntries(KindFeat, "feat", "featbenefit")
}

type o5eDescribes struct {
	Desc      string `json:"desc"`
	Describes string `json:"describes"`
}

func (n normalizer) conditions() ([]Entry, error) {
	var out []Entry
	for _, r := range n.fx["conditiondescription"] {
		c, err := decode[o5eDescribes](r)
		if err != nil {
			return nil, err
		}
		name := strings.ToUpper(c.Describes[:1]) + c.Describes[1:]
		out = append(out, Entry{Kind: KindCondition, Slug: n.slug(r.pk()), Name: name, Data: TextData{Desc: c.Desc}, Text: c.Desc})
	}
	return out, nil
}

type o5eRule struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	RuleSet string `json:"ruleset"`
	Index   int    `json:"index"`
}

func (n normalizer) rules() ([]Entry, error) {
	sets, err := byPK[o5eNamed](n.fx, "ruleset")
	if err != nil {
		return nil, err
	}
	var out []Entry
	for _, r := range n.fx["ruleset"] {
		s, err := decode[o5eRule](r)
		if err != nil {
			return nil, err
		}
		out = append(out, Entry{Kind: KindRule, Slug: n.slug(r.pk()), Name: s.Name, Data: TextData{Desc: s.Desc}, Text: s.Desc})
	}
	for _, r := range n.fx["rule"] {
		s, err := decode[o5eRule](r)
		if err != nil {
			return nil, err
		}
		cat := sets[s.RuleSet].Name
		out = append(out, Entry{Kind: KindRule, Slug: n.slug(r.pk()), Name: s.Name, Data: TextData{Desc: s.Desc, Category: cat}, Text: s.Desc})
	}
	return out, nil
}

// looseString accepts a JSON string, number or null; the Open5e data is not
// consistent about which it uses for costs, weights and ratings.
type looseString string

func (s *looseString) UnmarshalJSON(b []byte) error {
	switch {
	case string(b) == "null":
		*s = ""
	case len(b) > 0 && b[0] == '"':
		var v string
		if err := json.Unmarshal(b, &v); err != nil {
			return err
		}
		*s = looseString(v)
	default:
		*s = looseString(b)
	}
	return nil
}
