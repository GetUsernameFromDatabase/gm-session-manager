// Package dnd5e is the system module for 2024 D&D 5e: the character sheet
// stored in characters.sheet_json, the commands that change it, and the
// level-up helper. The generic core (name, hp, max_hp, temp_hp) lives in
// columns; everything system-specific lives in Sheet.
package dnd5e

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/go-playground/validator/v10"
)

// System is the characters.system value for this module.
const System = "dnd5e"

// Edition is the rules edition this module implements.
const Edition = "2024"

// Abilities lists the ability score keys in sheet order.
var Abilities = []string{"str", "dex", "con", "int", "wis", "cha"}

// Sizes are the creature sizes, smallest first.
var Sizes = []string{"Tiny", "Small", "Medium", "Large", "Huge", "Gargantuan"}

// SizeNamed is the size named like name (any case), or "".
func SizeNamed(name string) string {
	i := slices.IndexFunc(Sizes, func(s string) bool { return strings.EqualFold(s, strings.TrimSpace(name)) })
	if i < 0 {
		return ""
	}
	return Sizes[i]
}

// AbilityNames maps ability keys to display names.
var AbilityNames = map[string]string{
	"str": "Strength", "dex": "Dexterity", "con": "Constitution",
	"int": "Intelligence", "wis": "Wisdom", "cha": "Charisma",
}

// Skill is a 2024 skill and its ability.
type Skill struct {
	Key, Name, Ability string
}

// Skills lists every 2024 skill.
var Skills = []Skill{
	{"acrobatics", "Acrobatics", "dex"}, {"animal-handling", "Animal Handling", "wis"},
	{"arcana", "Arcana", "int"}, {"athletics", "Athletics", "str"},
	{"deception", "Deception", "cha"}, {"history", "History", "int"},
	{"insight", "Insight", "wis"}, {"intimidation", "Intimidation", "cha"},
	{"investigation", "Investigation", "int"}, {"medicine", "Medicine", "wis"},
	{"nature", "Nature", "int"}, {"perception", "Perception", "wis"},
	{"performance", "Performance", "cha"}, {"persuasion", "Persuasion", "cha"},
	{"religion", "Religion", "int"}, {"sleight-of-hand", "Sleight of Hand", "dex"},
	{"stealth", "Stealth", "dex"}, {"survival", "Survival", "wis"},
}

// Conditions lists the 2024 conditions, plus the trackers the combat tracker
// and sheet treat like conditions.
var Conditions = []string{
	"blinded", "charmed", "deafened", "exhaustion", "frightened", "grappled",
	"incapacitated", "invisible", "paralyzed", "petrified", "poisoned", "prone",
	"restrained", "stunned", "unconscious",
}

// Trackers are condition-like entries: concentration (note names the
// spell) and custom effects (note describes them).
const (
	TrackConcentration = "concentration"
	TrackCustom        = "custom"
)

// ConditionSummaries are one-line reminders of each condition's effect,
// shown on hover. Paraphrased from the SRD 5.2 (CC-BY 4.0, Wizards of the
// Coast).
var ConditionSummaries = map[string]string{
	"blinded":          "Can't see; fails checks that need sight. Attacks against it have advantage; its attacks have disadvantage.",
	"charmed":          "Can't attack or harm the charmer. The charmer has advantage on social checks against it.",
	"deafened":         "Can't hear; fails checks that need hearing.",
	"exhaustion":       "Each level: −2 to d20 tests and −5 ft speed. Level 6 is death. A long rest removes one level.",
	"frightened":       "Disadvantage on checks and attacks while the source is in sight; can't willingly move closer to it.",
	"grappled":         "Speed 0. Disadvantage on attacks against anyone but the grappler, who can drag it along.",
	"incapacitated":    "No actions, bonus actions or reactions; can't speak; concentration ends. Disadvantage on initiative.",
	"invisible":        "Can't be seen. Advantage on initiative and its attacks; attacks against it have disadvantage.",
	"paralyzed":        "Incapacitated, speed 0, fails Str and Dex saves. Attacks against it have advantage; hits within 5 ft are crits.",
	"petrified":        "Turned to stone: incapacitated, speed 0, fails Str and Dex saves, resists all damage, immune to poison.",
	"poisoned":         "Disadvantage on attack rolls and ability checks.",
	"prone":            "Can only crawl or stand up (half speed). Disadvantage on attacks; attacks within 5 ft have advantage, farther ones disadvantage.",
	"restrained":       "Speed 0. Disadvantage on attacks and Dex saves; attacks against it have advantage.",
	"stunned":          "Incapacitated, fails Str and Dex saves. Attacks against it have advantage.",
	"unconscious":      "Incapacitated and prone, drops what it holds, fails Str and Dex saves. Attacks have advantage; hits within 5 ft are crits.",
	TrackConcentration: "Ends on incapacitation, casting another concentration spell, or a failed Con save after damage (DC 10 or half the damage, max 30).",
}

// IsConditionSlug reports whether slug may be stored in conditions.slug.
func IsConditionSlug(slug string) bool {
	return slices.Contains(Conditions, slug) || slug == TrackConcentration || slug == TrackCustom
}

// ClassLevel is one class a character has levels in.
type ClassLevel struct {
	Slug     string `json:"slug"               validate:"required"`
	Name     string `json:"name"               validate:"required"`
	Subclass string `json:"subclass,omitempty"`
	Level    int    `json:"level"              validate:"min=1,max=20"`
	HitDie   int    `json:"hit_die"            validate:"oneof=6 8 10 12"`
	// Caster is the class's caster type: full, half, pact or none.
	Caster string `json:"caster" validate:"oneof=full half pact none"`
}

// Slot is a spell slot pool.
type Slot struct {
	Max  int `json:"max"  validate:"min=0,max=9"`
	Used int `json:"used" validate:"min=0,ltefield=Max"`
}

// Spell is a spell the character knows or has prepared.
type Spell struct {
	Slug     string `json:"slug"     validate:"required"`
	Name     string `json:"name"     validate:"required"`
	Level    int    `json:"level"    validate:"min=0,max=9"`
	Prepared bool   `json:"prepared"`
}

// Item is an inventory line. Armor, Shield and Weapon carry the stats of
// gear the character can wear or ready; whether it is worn or ready is
// Sheet.Armor, Sheet.Shield and Sheet.Weapons (see Worn and Ready).
type Item struct {
	Name   string     `json:"name"             validate:"required,max=200"`
	Qty    int        `json:"qty"              validate:"min=0"`
	Notes  string     `json:"notes,omitempty"  validate:"max=2000"`
	Armor  *WornArmor `json:"armor,omitempty"`
	Shield bool       `json:"shield,omitempty"`
	Weapon *Weapon    `json:"weapon,omitempty"`
}

// Worn reports whether an inventory item is the armor or shield worn.
func (s Sheet) Worn(it Item) bool {
	return it.Armor != nil && s.Armor != nil && strings.EqualFold(s.Armor.Name, it.Name) || it.Shield && s.Shield
}

// Ready reports whether an inventory item is a weapon on the attack list.
func (s Sheet) Ready(it Item) bool {
	return it.Weapon != nil && s.weaponIndex(it.Name) >= 0
}

func (s Sheet) weaponIndex(name string) int {
	return slices.IndexFunc(s.Weapons, func(w Weapon) bool { return strings.EqualFold(w.Name, name) })
}

// ArmorItem is the index of the inventory armor or shield named name, or -1.
func (s Sheet) ArmorItem(name string) int {
	return slices.IndexFunc(s.Inventory, func(it Item) bool {
		return (it.Armor != nil || it.Shield) && strings.EqualFold(it.Name, name)
	})
}

// ensureItem gives the inventory item named like it its gear stats, adding
// it when missing, so gear is listed once however it was added.
func ensureItem(inv []Item, it Item) []Item {
	i := findItem(inv, it.Name)
	if i < 0 {
		it.Qty = max(it.Qty, 1)
		return append(inv, it)
	}
	if it.Armor != nil {
		inv[i].Armor = it.Armor
	}
	if it.Weapon != nil {
		inv[i].Weapon = it.Weapon
	}
	inv[i].Shield = inv[i].Shield || it.Shield
	return inv
}

// shieldItem is the inventory line a shield gets when none is listed.
var shieldItem = Item{Name: "Shield", Shield: true}

// Feature is a class feature, species trait, feat or anything else the
// player wants on the sheet. The description is shown; what the sheet can
// automate is set apart (see ParseEffects): the action it takes, its uses,
// and unarmored AC. Its stat changes are modifiers whose Source is its name.
type Feature struct {
	Name        string `json:"name"                  validate:"required,max=200"`
	Source      string `json:"source,omitempty"      validate:"max=200"`
	Level       int    `json:"level,omitempty"       validate:"min=0,max=20"`
	Description string `json:"description,omitempty" validate:"max=20000"`
	// Action is what using it takes: action, bonus, reaction or magic.
	Action string `json:"action,omitempty" validate:"omitempty,oneof=action bonus reaction magic"`
	Uses   Uses   `json:"uses,omitzero"`
	// Unarmored is the ability added to 10 + Dexterity for AC while no
	// armor is worn (Unarmored Defense); NoShield also forbids a shield.
	Unarmored string `json:"unarmored,omitempty" validate:"omitempty,oneof=str dex con int wis cha"`
	NoShield  bool   `json:"no_shield,omitempty"`
}

// Uses is how often a feature can be used between rests: Fixed times, the
// Ability modifier (minimum once), the proficiency bonus, or a class table
// column (ByLevel, by level in Class). The zero value means unlimited.
type Uses struct {
	Fixed   int         `json:"fixed,omitempty"    validate:"min=0,max=99"`
	Ability string      `json:"ability,omitempty"  validate:"omitempty,oneof=str dex con int wis cha"`
	Prof    bool        `json:"prof,omitempty"`
	ByLevel map[int]int `json:"by_level,omitempty"`
	Class   string      `json:"class,omitempty"    validate:"max=200"`
	Used    int         `json:"used,omitempty"     validate:"min=0"`
	// Reset is when uses come back: short (short or long rest), long, or
	// short-one (one on a short rest, all on a long rest).
	Reset string `json:"reset,omitempty" validate:"omitempty,oneof=short long short-one"`
}

// Limited reports whether the uses are counted at all.
func (u Uses) Limited() bool { return u.Fixed > 0 || u.Ability != "" || u.Prof || len(u.ByLevel) > 0 }

// MaxUses is how many uses a feature has now.
func (s Sheet) MaxUses(u Uses) int {
	switch {
	case u.Fixed > 0:
		return u.Fixed
	case u.Ability != "":
		return max(s.Mod(u.Ability), 1)
	case u.Prof:
		return s.ProficiencyBonus()
	case len(u.ByLevel) > 0:
		level, best := s.ClassLevel(u.Class), 0
		for l, n := range u.ByLevel {
			if l <= level && n > best {
				best = n
			}
		}
		return best
	}
	return 0
}

// ClassLevel is the character's level in the named class.
func (s Sheet) ClassLevel(name string) int {
	for _, c := range s.Classes {
		if strings.EqualFold(c.Name, name) {
			return c.Level
		}
	}
	return 0
}

// WornArmor is the armor a character wears: Base plus the Dexterity
// modifier when AddDex, at most DexCap (0: no cap).
type WornArmor struct {
	Slug    string `json:"slug,omitempty"`
	Name    string `json:"name"              validate:"required,max=200"`
	Type    string `json:"type"              validate:"oneof=light medium heavy"`
	Base    int    `json:"base"              validate:"min=0,max=30"`
	AddDex  bool   `json:"add_dex,omitempty"`
	DexCap  int    `json:"dex_cap,omitempty" validate:"min=0,max=10"`
	StrReq  int    `json:"str_req,omitempty" validate:"min=0,max=30"`
	Stealth bool   `json:"stealth,omitempty"`
}

// DeathSaves tracks death saving throws.
type DeathSaves struct {
	Successes int `json:"successes" validate:"min=0,max=3"`
	Failures  int `json:"failures"  validate:"min=0,max=3"`
}

// Coins is the character's purse.
type Coins struct {
	CP int `json:"cp" validate:"min=0"`
	SP int `json:"sp" validate:"min=0"`
	EP int `json:"ep" validate:"min=0"`
	GP int `json:"gp" validate:"min=0"`
	PP int `json:"pp" validate:"min=0"`
}

// Weapon is an attack on the sheet. Ability is the attack ability, or
// "finesse" for the better of Strength and Dexterity.
type Weapon struct {
	Name       string `json:"name"                  validate:"required,max=200"`
	Ability    string `json:"ability"               validate:"oneof=str dex con int wis cha finesse"`
	Dice       string `json:"dice"                  validate:"max=50"`
	DamageType string `json:"damage_type,omitempty" validate:"max=50"`
	Proficient bool   `json:"proficient"`
	// Range is normal/long range in feet for ranged and thrown weapons.
	Range      string `json:"range,omitempty"       validate:"max=20"`
	Properties string `json:"properties,omitempty"  validate:"max=300"`
}

// CustomMod is a custom bonus for what the app does not compute itself
// (homebrew, magic items, house rules) or a feature's stat change. Target
// and Note identify it. Source names the feature it belongs to, and goes
// with it; When limits it to a state of armor (see Applies).
type CustomMod struct {
	Target string `json:"target"           validate:"required,max=250"`
	Value  int    `json:"value"            validate:"min=-50,max=50"`
	Note   string `json:"note"             validate:"required,max=200"`
	Source string `json:"source,omitempty" validate:"max=200"`
	When   string `json:"when,omitempty"   validate:"omitempty,oneof=unarmored unarmored-no-shield no-heavy-armor"`
}

// WhenLabels say when a conditional modifier applies.
var WhenLabels = map[string]string{
	"unarmored":           "while wearing no armor",
	"unarmored-no-shield": "while wearing no armor and no shield",
	"no-heavy-armor":      "while not in heavy armor",
}

// Applies reports whether a modifier's condition holds now.
func (s Sheet) Applies(m CustomMod) bool {
	switch m.When {
	case "unarmored":
		return s.Armor == nil
	case "unarmored-no-shield":
		return s.Armor == nil && !s.Shield
	case "no-heavy-armor":
		return s.Armor == nil || s.Armor.Type != "heavy"
	}
	return true
}

// Modifier targets besides abilities (the ability key), saves
// ("save.str"), skills ("skill.stealth") and single weapons
// ("attack.<name>", "damage.<name>").
const (
	ModAC      = "ac"
	ModSpeed   = "speed"
	ModInit    = "init"
	ModPassive = "passive"
	ModAttack  = "attack"
	ModDamage  = "damage"
)

// IsModTarget reports whether target names a value modifiers can change.
func IsModTarget(target string) bool {
	if slices.Contains([]string{ModAC, ModSpeed, ModInit, ModPassive, ModAttack, ModDamage}, target) || slices.Contains(Abilities, target) {
		return true
	}
	kind, rest, ok := strings.Cut(target, ".")
	switch kind {
	case "save":
		return slices.Contains(Abilities, rest)
	case "skill":
		return slices.ContainsFunc(Skills, func(sk Skill) bool { return sk.Key == rest })
	case ModAttack, ModDamage:
		return ok && rest != ""
	}
	return false
}

// Sheet is the 2024 5e part of a character, stored as sheet_json.
type Sheet struct {
	Species    string         `json:"species,omitempty"    validate:"max=200"`
	Background string         `json:"background,omitempty" validate:"max=200"`
	Backstory  string         `json:"backstory,omitempty"  validate:"max=20000"`
	Classes    []ClassLevel   `json:"classes"              validate:"dive"`
	Scores     map[string]int `json:"abilities"            validate:"dive,keys,oneof=str dex con int wis cha,endkeys,min=1,max=30"`
	// ACBonus is added to AC whatever is worn; AC is the base AC older
	// sheets set by hand, turned into ACBonus by ParseSheet.
	ACBonus int    `json:"ac_bonus,omitempty"   validate:"min=-20,max=30"`
	AC      int    `json:"ac,omitempty"         validate:"min=0,max=50"`
	Speed   int    `json:"speed"                validate:"min=0,max=500"`
	Size    string `json:"size,omitempty"       validate:"omitempty,oneof=Tiny Small Medium Large Huge Gargantuan"`
	// Vision is what the character sees in the dark ("Darkvision 60 ft");
	// empty is normal sight.
	Vision     string   `json:"vision,omitempty"     validate:"max=200"`
	SaveProfs  []string `json:"save_profs"           validate:"dive,oneof=str dex con int wis cha"`
	SkillProfs []string `json:"skill_profs"          validate:"dive,required"`
	Expertise  []string `json:"expertise"            validate:"dive,required"`
	OtherProfs []string `json:"other_profs"          validate:"dive,required,max=200"`
	Languages  []string `json:"languages"            validate:"dive,required,max=100"`
	// HitDiceUsed counts spent hit dice across all classes.
	HitDiceUsed int         `json:"hit_dice_used" validate:"min=0"`
	Slots       [9]Slot     `json:"spell_slots"   validate:"dive"`
	Pact        Slot        `json:"pact_slots"`
	PactLevel   int         `json:"pact_level"    validate:"min=0,max=5"`
	Spells      []Spell     `json:"spells"        validate:"dive"`
	Inventory   []Item      `json:"inventory"     validate:"dive"`
	Weapons     []Weapon    `json:"weapons"       validate:"dive"`
	Modifiers   []CustomMod `json:"modifiers"     validate:"dive"`
	Features    []Feature   `json:"features"      validate:"dive"`
	Feats       []string    `json:"feats"         validate:"dive,required"`
	DeathSaves  DeathSaves  `json:"death_saves"`
	Inspiration bool        `json:"inspiration"`
	Coins       Coins       `json:"coins"`
	// LastClass is the slug of the class last levelled, the default next time.
	LastClass string `json:"last_class,omitempty"`
	// Armor is the armor worn (nil: none); Shield adds 2 to AC.
	Armor  *WornArmor `json:"armor,omitempty"`
	Shield bool       `json:"shield,omitempty"`
	// SetupDone lists the creation steps already applied, which add to the
	// sheet and so run once: class, species, background, equipment, and
	// review when the player finished.
	SetupDone []string `json:"setup_done,omitempty" validate:"dive,oneof=class species background equipment review"`
	// EffectsVersion is the EffectsVersion the features last caught up with.
	EffectsVersion int `json:"effects_version,omitempty"`
}

// SetUp reports whether a creation step was applied.
func (s Sheet) SetUp(step string) bool { return slices.Contains(s.SetupDone, step) }

// NewSheet returns the sheet of a freshly created character.
func NewSheet() Sheet {
	// Every character speaks Common; sheets saved before languages get it too.
	s := Sheet{Scores: map[string]int{}, Speed: 30, Languages: []string{"Common"}}
	for _, a := range Abilities {
		s.Scores[a] = 10
	}
	return s
}

// Character is a whole character as commands see it: the generic columns
// plus the system sheet.
type Character struct {
	Name   string `validate:"required,max=100"`
	HP     int    // below 0 records overkill
	MaxHP  int    `validate:"min=0"`
	TempHP int    `validate:"min=0"`
	Sheet  Sheet
}

var validate = validator.New(validator.WithRequiredStructEnabled())

// Validate checks every field against its tags and the cross-field rules.
func (c *Character) Validate() error {
	if err := validate.Struct(c); err != nil {
		return fmt.Errorf("invalid character: %w", err)
	}
	if c.HP > c.MaxHP && c.MaxHP > 0 {
		return fmt.Errorf("invalid character: hp %d above max %d", c.HP, c.MaxHP)
	}
	return nil
}

// ParseSheet decodes sheet_json, filling defaults for fields older rows lack.
func ParseSheet(raw string) (Sheet, error) {
	s := NewSheet()
	if raw != "" && raw != "{}" {
		if err := json.Unmarshal([]byte(raw), &s); err != nil {
			return Sheet{}, fmt.Errorf("sheet: %w", err)
		}
	}
	if s.Scores == nil {
		s.Scores = map[string]int{}
	}
	for _, a := range Abilities {
		if s.Scores[a] == 0 {
			s.Scores[a] = 10
		}
	}
	s.upgrade()
	return s, nil
}

// upgrade brings sheets saved before gear lived in the inventory and AC
// was a bonus up to date; it changes nothing on a current sheet.
func (s *Sheet) upgrade() {
	if s.AC != 0 {
		// The old base AC counted only without armor or Unarmored Defense;
		// 10 was the default, which now becomes 10 + Dexterity.
		if s.AC != 10 && s.Armor == nil && s.unarmoredDefense() == nil {
			s.ACBonus = s.AC - 10 - s.Mod("dex")
		}
		s.AC = 0
	}
	if s.Armor != nil {
		a := *s.Armor
		s.Inventory = ensureItem(s.Inventory, Item{Name: a.Name, Armor: &a})
	}
	if s.Shield && !slices.ContainsFunc(s.Inventory, func(it Item) bool { return it.Shield }) {
		s.Inventory = ensureItem(s.Inventory, shieldItem)
	}
	for _, w := range s.Weapons {
		if i := findItem(s.Inventory, w.Name); i < 0 || s.Inventory[i].Weapon == nil {
			s.Inventory = ensureItem(s.Inventory, Item{Name: w.Name, Weapon: &w})
		}
	}
}

// JSON encodes the sheet for sheet_json.
func (s Sheet) JSON() (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Level is the character's total level.
func (s Sheet) Level() int {
	n := 0
	for _, c := range s.Classes {
		n += c.Level
	}
	return n
}

// ProficiencyBonus follows total level: +2 at 1-4 up to +6 at 17-20.
func (s Sheet) ProficiencyBonus() int {
	return ProficiencyAt(s.Level())
}

// ProficiencyAt is the proficiency bonus at a total character level.
func ProficiencyAt(level int) int {
	return 2 + (max(level, 1)-1)/4
}

// Mods lists the custom modifiers on target that apply now.
func (s Sheet) Mods(target string) []CustomMod {
	var out []CustomMod
	for _, m := range s.Modifiers {
		if m.Target == target && s.Applies(m) {
			out = append(out, m)
		}
	}
	return out
}

// Bonus sums the custom modifiers on target.
func (s Sheet) Bonus(target string) int {
	n := 0
	for _, m := range s.Mods(target) {
		n += m.Value
	}
	return n
}

// Score is an ability score with its custom modifiers.
func (s Sheet) Score(ability string) int {
	return s.Scores[ability] + s.Bonus(ability)
}

// Mod is the ability modifier for an ability key.
func (s Sheet) Mod(ability string) int {
	return Modifier(s.Score(ability))
}

// ArmorClass is the AC: armor, unarmored defense or 10 + Dexterity, plus
// shield, the bonus set by hand and modifiers (see Explain).
func (s Sheet) ArmorClass() int { return Sum(s.Explain(ModAC)) }

// WalkSpeed is the speed with armor penalties and modifiers.
func (s Sheet) WalkSpeed() int { return Sum(s.Explain(ModSpeed)) }

// Initiative is the Dexterity modifier with its custom modifiers.
func (s Sheet) Initiative() int { return Sum(s.Explain(ModInit)) }

// AttackMod is the ability modifier a weapon uses.
func (s Sheet) AttackMod(w Weapon) int {
	if w.Ability == "finesse" {
		return max(s.Mod("str"), s.Mod("dex"))
	}
	return s.Mod(w.Ability)
}

// AttackBonus is a weapon's bonus to hit.
func (s Sheet) AttackBonus(w Weapon) int { return Sum(s.attackParts(w)) }

// DamageLabel is a weapon's damage, such as "1d8+3 slashing".
func (s Sheet) DamageLabel(w Weapon) string { return damageLabel(w, s.damageParts(w)) }

// OffHandDamage is a light weapon's damage as a two-weapon fighting bonus
// attack (see Explain "offhand.<name>").
func (s Sheet) OffHandDamage(w Weapon) string { return damageLabel(w, s.offHandParts(w)) }

func damageLabel(w Weapon, parts []Part) string {
	out := w.Dice
	if b := Sum(parts); out == "" {
		out = fmt.Sprint(b)
	} else if b != 0 {
		out += Signed(b)
	}
	return strings.TrimSpace(out + " " + w.DamageType)
}

// Modifier converts an ability score to its modifier.
func Modifier(score int) int {
	// Floor division so 9 gives -1, not 0.
	d := score - 10
	if d < 0 {
		return (d - 1) / 2
	}
	return d / 2
}

// SaveBonus is the saving throw bonus for an ability.
func (s Sheet) SaveBonus(ability string) int { return Sum(s.Explain("save." + ability)) }

// SkillBonus is the bonus for a skill, with expertise doubling proficiency.
func (s Sheet) SkillBonus(sk Skill) int { return Sum(s.skillParts(sk)) }

// PassivePerception is 10 plus the Perception bonus.
func (s Sheet) PassivePerception() int { return Sum(s.Explain(ModPassive)) }

// HitDice summarizes total hit dice per die size.
func (s Sheet) HitDice() map[int]int {
	out := map[int]int{}
	for _, c := range s.Classes {
		out[c.HitDie] += c.Level
	}
	return out
}

// ClassSummary is "Wizard 3 / Fighter 1".
func (s Sheet) ClassSummary() string {
	parts := make([]string, 0, len(s.Classes))
	for _, c := range s.Classes {
		p := fmt.Sprintf("%s %d", c.Name, c.Level)
		if c.Subclass != "" {
			p += " (" + c.Subclass + ")"
		}
		parts = append(parts, p)
	}
	return strings.Join(parts, " / ")
}

// Signed formats a bonus as +2 or -1.
func Signed(n int) string {
	if n >= 0 {
		return fmt.Sprintf("+%d", n)
	}
	return fmt.Sprint(n)
}

// spellAbilities are the spellcasting abilities of the casting classes.
var spellAbilities = map[string]string{
	"artificer": "int", "bard": "cha", "cleric": "wis", "druid": "wis", "paladin": "cha",
	"ranger": "wis", "sorcerer": "cha", "warlock": "cha", "wizard": "int",
}

// SpellAbility is the spellcasting ability of the character's first
// casting class, or empty when it has none we know.
func (s Sheet) SpellAbility() string {
	for _, c := range s.Classes {
		if a, ok := spellAbilities[strings.ToLower(c.Name)]; ok {
			return a
		}
	}
	return ""
}

// SpellAttack is the spell attack bonus with ability.
func (s Sheet) SpellAttack(ability string) int { return s.ProficiencyBonus() + s.Mod(ability) }

// SpellDC is the spell save DC with ability.
func (s Sheet) SpellDC(ability string) int { return 8 + s.SpellAttack(ability) }
