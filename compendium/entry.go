// Package compendium imports rules content into compendium_entries and
// searches it. Every source is normalized into the shapes in this file, so
// no data source's own layout reaches the rest of the app.
package compendium

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

// Kind is compendium_entries.kind.
type Kind string

// Entry kinds.
const (
	KindMonster    Kind = "monster"
	KindSpell      Kind = "spell"
	KindItem       Kind = "item"
	KindClass      Kind = "class"
	KindSubclass   Kind = "subclass"
	KindSpecies    Kind = "species"
	KindBackground Kind = "background"
	KindFeat       Kind = "feat"
	KindCondition  Kind = "condition"
	KindRule       Kind = "rule"
)

// Kinds lists every kind in display order.
var Kinds = []Kind{
	KindMonster, KindSpell, KindItem, KindClass, KindSubclass, KindSpecies,
	KindBackground, KindFeat, KindCondition, KindRule,
}

// ValidKind reports whether k is a known kind.
func ValidKind(k string) bool {
	for _, x := range Kinds {
		if string(x) == k {
			return true
		}
	}
	return false
}

// HomebrewSource is the source_key of homebrew entries.
const HomebrewSource = "homebrew"

// MonsterImagePrefix + slug + ".png" is the first guess at a monster's
// picture: a link only, never downloaded; misses are hidden by the page.
// https://github.com/5e-bits/5e-srd-api/blob/main/apps/api/src/controllers/api/imageController.ts
const MonsterImagePrefix = "https://dnd-5e-api-images.s3.us-west-1.amazonaws.com/monsters/"

// Habitats are the 2024 Monster Manual habitats, in display order; the
// monster_habitats CHECK lists the same values.
var Habitats = []string{
	"Any", "Arctic", "Coastal", "Desert", "Forest", "Grassland", "Hill",
	"Mountain", "Planar", "Swamp", "Underdark", "Underwater", "Urban",
}

// Entry is one normalized compendium entry ready to store.
type Entry struct {
	Kind Kind
	Slug string
	Name string
	// Data is one of the *Data types below, matching Kind.
	Data any
	// Text is the plain text indexed for search.
	Text string
}

// Source describes where a set of entries came from.
type Source struct {
	Key         string
	Edition     string
	License     string
	Attribution string
}

// NameDesc is a named block of rules text (traits, actions, benefits).
type NameDesc struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	// Type tags a species trait (size, speed) or background benefit
	// (ability_score, skill_proficiency, tool_proficiency, equipment, feat).
	Type string `json:"type,omitempty"`
}

// MonsterAction is one statblock action.
type MonsterAction struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	// Type is action, bonus_action, reaction or legendary_action.
	Type string `json:"type"`
}

// MonsterData is a creature statblock.
type MonsterData struct {
	Size       string          `json:"size,omitempty"`
	Type       string          `json:"type,omitempty"`
	Alignment  string          `json:"alignment,omitempty"`
	AC         int             `json:"ac"`
	ACDetail   string          `json:"ac_detail,omitempty"`
	HP         int             `json:"hp"`
	HitDice    string          `json:"hit_dice,omitempty"`
	Initiative int             `json:"initiative"`
	Speed      map[string]int  `json:"speed,omitempty"`
	Abilities  map[string]int  `json:"abilities,omitempty"`
	Saves      map[string]int  `json:"saves,omitempty"`
	Skills     map[string]int  `json:"skills,omitempty"`
	CR         string          `json:"cr,omitempty"`
	Senses     string          `json:"senses,omitempty"`
	Languages  string          `json:"languages,omitempty"`
	Resistance string          `json:"resistances,omitempty"`
	Immunity   string          `json:"immunities,omitempty"`
	Traits     []NameDesc      `json:"traits,omitempty"`
	Actions    []MonsterAction `json:"actions,omitempty"`
	Desc       string          `json:"desc,omitempty"`
}

// SpellData is a spell.
type SpellData struct {
	Level         int      `json:"level"`
	School        string   `json:"school,omitempty"`
	CastingTime   string   `json:"casting_time,omitempty"`
	Range         string   `json:"range,omitempty"`
	Components    string   `json:"components,omitempty"`
	Duration      string   `json:"duration,omitempty"`
	Concentration bool     `json:"concentration"`
	Ritual        bool     `json:"ritual"`
	Classes       []string `json:"classes,omitempty"`
	Desc          string   `json:"desc"`
	HigherLevel   string   `json:"higher_level,omitempty"`
}

// ItemData is mundane or magic equipment, weapons and armor included.
type ItemData struct {
	Category   string `json:"category,omitempty"`
	Rarity     string `json:"rarity,omitempty"`
	Attunement bool   `json:"attunement"`
	Cost       string `json:"cost,omitempty"`
	Weight     string `json:"weight,omitempty"`
	Damage     string `json:"damage,omitempty"`
	Properties string `json:"properties,omitempty"`
	// Range is a ranged or thrown weapon's normal/long range in feet.
	Range   string `json:"range,omitempty"`
	Martial bool   `json:"martial,omitempty"`
	// ArmorClass is the AC as people read it; Armor is the same as numbers.
	ArmorClass string     `json:"armor_class,omitempty"`
	Armor      *ArmorData `json:"armor,omitempty"`
	Desc       string     `json:"desc"`
}

// Armor types.
const (
	ArmorLight  = "light"
	ArmorMedium = "medium"
	ArmorHeavy  = "heavy"
	ArmorShield = "shield"
)

// ArmorData is a piece of armor's AC rule: Base plus the Dexterity
// modifier when AddDex, at most DexCap (0: no cap). A shield's Base adds.
type ArmorData struct {
	Type    string `json:"type"`
	Base    int    `json:"base"`
	AddDex  bool   `json:"add_dex,omitempty"`
	DexCap  int    `json:"dex_cap,omitempty"`
	StrReq  int    `json:"str_req,omitempty"`
	Stealth bool   `json:"stealth_disadvantage,omitempty"`
}

// ClassFeature is a class feature and the class levels it is gained at.
type ClassFeature struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Levels []int  `json:"levels"`
	// Subclass marks the feature where the subclass is chosen.
	Subclass bool `json:"subclass,omitempty"`
}

// ClassData is a class or subclass.
type ClassData struct {
	HitDie int `json:"hit_die,omitempty"`
	// Caster is full, half, pact or none.
	Caster     string         `json:"caster"`
	Saves      []string       `json:"saves,omitempty"`
	SubclassOf string         `json:"subclass_of,omitempty"`
	Features   []ClassFeature `json:"features,omitempty"`
	// Table maps class table columns to values by class level.
	Table map[string]map[int]string `json:"table,omitempty"`
	Desc  string                    `json:"desc,omitempty"`
}

// TextData covers species, backgrounds, feats, conditions and rules: a
// description and optional named parts.
type TextData struct {
	Desc         string     `json:"desc"`
	Parts        []NameDesc `json:"parts,omitempty"`
	Prerequisite string     `json:"prerequisite,omitempty"`
	Category     string     `json:"category,omitempty"`
}

// dataFor returns an empty Data value of the right type for kind.
func dataFor(kind Kind) any {
	switch kind {
	case KindMonster:
		return &MonsterData{}
	case KindSpell:
		return &SpellData{}
	case KindItem:
		return &ItemData{}
	case KindClass, KindSubclass:
		return &ClassData{}
	}
	return &TextData{}
}

// DecodeData parses data_json into the type for kind.
func DecodeData(kind Kind, raw string) (any, error) {
	d := dataFor(kind)
	if err := json.Unmarshal([]byte(raw), d); err != nil {
		return nil, fmt.Errorf("compendium %s data: %w", kind, err)
	}
	return d, nil
}

// encode returns data_json and the content hash covering everything stored.
func (e Entry) encode(src Source) (data, hash string, err error) {
	b, err := json.Marshal(e.Data)
	if err != nil {
		return "", "", fmt.Errorf("encode %s %s: %w", e.Kind, e.Slug, err)
	}
	h := sha256.New()
	for _, part := range []string{string(e.Kind), e.Slug, e.Name, string(b), e.Text, src.Edition, src.License, src.Attribution} {
		h.Write([]byte(part))
		h.Write([]byte{0})
	}
	return string(b), hex.EncodeToString(h.Sum(nil)), nil
}

// Slugify makes a URL-safe slug from a name.
func Slugify(name string) string {
	var b strings.Builder
	dash := false
	for _, r := range strings.ToLower(name) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
			dash = false
		case !dash && b.Len() > 0:
			b.WriteByte('-')
			dash = true
		}
	}
	return strings.TrimSuffix(b.String(), "-")
}
