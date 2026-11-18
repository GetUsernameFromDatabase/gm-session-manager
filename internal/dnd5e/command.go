package dnd5e

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"gm-session-manager/internal/perm"
)

// Command is one small change to a character, such as "hp -5" or
// "toggle prone". Commands are deltas wherever a delta makes sense, because
// several devices may hold the same character; "=n" forms are the explicit
// manual override.
type Command struct {
	// Text is the normalized command, for logs and the activity line.
	Text string
	// Scope says who may run it (see perm.CanCharacter).
	Scope perm.CharacterScope
	// SpellSlug is set when the command needs a compendium spell; the
	// caller resolves it into Env.Spell before Apply.
	SpellSlug string
	// ArmorSlug is set when the command needs compendium armor; the caller
	// resolves it into Env.Armor (and Env.IsShield) before Apply.
	ArmorSlug string
	// PrepareCaps is set when the command needs the classes' prepared-spell
	// totals; the caller resolves them into Env.PreparedCaps before Apply.
	PrepareCaps bool
	// SourceRef is set when the command re-points a ref to another source;
	// the caller resolves Env.SourceOld and Env.SourceNew (a class or
	// subclass) or Env.SourceSpell before Apply.
	SourceRef SourceRef
	// SubclassChange is set when the command swaps a class's subclass; the
	// caller resolves Env.SourceOld and Env.SourceNew before Apply.
	SubclassChange SubclassRef
	// MasteryWeapon is set when the command masters a weapon kind; the
	// caller resolves Env.MasteryWeapon and Env.MasteryLimit before Apply.
	MasteryWeapon string
	apply         func(c *Character, env Env, fx *Effects) error
}

// Env carries what a command needs from outside the character.
type Env struct {
	// Spell is the resolved compendium spell for Command.SpellSlug.
	Spell Spell
	// Armor is the resolved compendium armor for Command.ArmorSlug;
	// IsShield when that is a shield instead.
	Armor    WornArmor
	IsShield bool
	// Petrified is the character's condition, which resists all damage.
	Petrified bool
	// PreparedCaps are the prepared-spell totals by class slug, for
	// Command.PrepareCaps; a class without an entry has no cap.
	PreparedCaps map[string]int
	// SourceOld and SourceNew are the class (with the subclass's features)
	// in the pinned copy and in Command.SourceRef's (for change-subclass:
	// with the current subclass and with the new one); SourceOld is nil when
	// the pinned copy is gone. SourceSpell is the spell in the new copy.
	SourceOld   *ClassInfo
	SourceNew   ClassInfo
	SourceSpell Spell
	// MasteryWeapon is the compendium weapon Command.MasteryWeapon names, as
	// the compendium spells it; empty when none of that name has a mastery.
	// MasteryLimit is how many kinds the classes let the character master.
	MasteryWeapon string
	MasteryLimit  int
}

// NewCommand wraps a change built by code instead of parsed from text, for
// forms with more fields than a command line carries.
func NewCommand(text string, scope perm.CharacterScope, apply func(c *Character) error) Command {
	return Command{Text: text, Scope: scope, apply: func(c *Character, _ Env, _ *Effects) error { return apply(c) }}
}

// CondOp is a change to the conditions table, applied by the caller in the
// same transaction as the character update.
type CondOp struct {
	Op   CondOpKind
	Slug string
	Note string
	// Delta is used by CondExhaustion.
	Delta int
	// Rest is ResetShort or ResetLong for CondEndRest.
	Rest string
}

// CondOpKind is what a CondOp does.
type CondOpKind int

// Condition operations.
const (
	CondToggle CondOpKind = iota
	CondAdd
	CondRemove
	// CondExhaustion changes the exhaustion level (stored in the note) by
	// Delta, removing the condition at 0 and capping at 6.
	CondExhaustion
	// CondEndRest removes the effects that end on the rest named by Rest: a
	// long rest also ends the short rest ones.
	CondEndRest
	// CondClearAll removes every condition and effect of the owner (death).
	CondClearAll
)

// Effects lists what Apply could not do to the character struct itself.
type Effects struct {
	Conditions []CondOp
}

// ErrUnknownCommand is returned for input no command matches.
var ErrUnknownCommand = errors.New("unknown command")

// Apply runs the command against c and validates the result. c is left
// unchanged if it returns an error.
func (cmd Command) Apply(c *Character, env Env) (Effects, error) {
	next := c.clone()
	var fx Effects
	if err := cmd.apply(&next, env, &fx); err != nil {
		return Effects{}, err
	}
	if err := next.Validate(); err != nil {
		return Effects{}, err
	}
	*c = next
	return fx, nil
}

func (c Character) clone() Character {
	s := c.Sheet
	s.Scores = maps.Clone(c.Sheet.Scores)
	s.Classes = slices.Clone(s.Classes)
	s.SaveProfs = slices.Clone(s.SaveProfs)
	s.SkillProfs = slices.Clone(s.SkillProfs)
	s.Expertise = slices.Clone(s.Expertise)
	s.OtherProfs = slices.Clone(s.OtherProfs)
	s.ProfSources = maps.Clone(s.ProfSources)
	s.Languages = slices.Clone(s.Languages)
	s.Resistances = slices.Clone(s.Resistances)
	s.Immunities = slices.Clone(s.Immunities)
	s.Vulnerabilities = slices.Clone(s.Vulnerabilities)
	s.Spells = slices.Clone(s.Spells)
	s.Inventory = slices.Clone(s.Inventory)
	s.Features = slices.Clone(s.Features)
	s.Feats = slices.Clone(s.Feats)
	s.FeatSources = maps.Clone(s.FeatSources)
	s.Weapons = slices.Clone(s.Weapons)
	s.Masteries = slices.Clone(s.Masteries)
	s.Modifiers = slices.Clone(s.Modifiers)
	s.SetupDone = slices.Clone(s.SetupDone)
	if s.Armor != nil {
		a := *s.Armor
		s.Armor = &a
	}
	c.Sheet = s
	return c
}

// HelpEntry is one command form as the UI explains it.
type HelpEntry struct {
	Syntax, Example, Desc string
}

// Help lists the command forms for the UI, roughly by how often they are used.
var Help = []HelpEntry{
	{"hp -n [type] [crit] | +n | =n", "hp -7 fire crit", "Take damage (crit: a critical hit), heal, or set hit points"},
	{"temp n | +n", "temp 5", "Set or add temporary hit points"},
	{"maxhp +n | =n", "maxhp =30", "Change max hit points"},
	{"slot use|restore <level>", "slot use 2", "Spend or regain a spell slot"},
	{"slot max <level> <n>", "slot max 2 3", "Set how many slots a level has"},
	{"pact use|restore", "pact use", "Spend or regain a pact slot"},
	{"rest short|long", "rest long", "Rest: long restores HP, slots and hit dice"},
	{"hitdie use|restore", "hitdie use", "Spend or regain a hit die"},
	{"toggle <condition>", "toggle prone", "Turn a condition on or off"},
	{"condition add|remove <condition> [note]", "condition add poisoned spider bite", "Add or remove a condition"},
	{"concentrate <spell> | end", "concentrate Bless", "Start or end concentration"},
	{"exhaustion +n | -n", "exhaustion +1", "Change exhaustion level"},
	{"deathsave success|fail|reset", "deathsave fail", "Track death saves"},
	{"revive [n]", "revive 1", "Bring a dead character back with n hit points (default 1)"},
	{verbInspiration, verbInspiration, "Toggle Heroic Inspiration"},
	{`item add|remove "<name>" [qty]`, `item add "Rope (50 ft)"`, "Add or remove items (quote names with spaces)"},
	{`item set "<name>" <qty>`, "item set Torch 5", "Set an item's count; 0 removes it"},
	{`item note "<name>" <text>`, "item note Torch half burnt", "Note on an item"},
	{"coin <cp|sp|ep|gp|pp> +n|-n|=n …", "coin gp +10 sp -3", "Gain, spend or set coins, several at once"},
	{"spell add|remove|prepare|unprepare <slug>", "spell prepare shield", "Manage spells by compendium slug"},
	{"set ability <ab> <n>", "set ability str 16", "Set a base ability score"},
	{"set name|species|background <text>", "set species Elf", "Change details"},
	{"set ac|speed <n>", "set ac 1", "Set an AC bonus that always applies, or base speed"},
	{"set size <size> | set vision <text>|none", "set vision Darkvision 60 ft", "Set size (Tiny…Gargantuan) or vision"},
	{"set backstory <text>", "set backstory Raised by wolves", "Replace the backstory"},
	{"prof add|remove skill|save|expertise|other|language <name>", "prof add skill athletics", "Add or remove a proficiency"},
	{"prof set skill <skill> none|proficient|expertise", "prof set skill stealth expertise", "Set a skill's proficiency"},
	{`feature add "<name>" [description]`, `feature add "Second Wind"`, "Add or replace a feature"},
	{`feature set "<name>" action=bonus uses=2|cha|prof reset=short|long|short-one`, `feature set "Second Wind" action=bonus uses=2 reset=short-one`, "Set how a feature is used and how often"},
	{`feature remove "<name>"`, `feature remove "Second Wind"`, "Remove a feature and its stat changes"},
	{`use "<feature>" | restore "<feature>"`, `use "Bardic Inspiration"`, "Spend or regain one use of a feature"},
	{`armor wear <slug|"item"> | armor off`, "armor wear chain-mail", "Put on armor (or a shield) from the inventory or compendium, or take armor off"},
	{"shield on|off", "shield off", "Take up or put down a shield (+2 AC)"},
	{"weapon add <name> <ability> <dice> [type] [prof]", "weapon add Longsword str 1d8 slashing prof", "Add or replace a weapon"},
	{"weapon ready|stow <name>", "weapon stow Longsword", "Ready a weapon from the inventory as an attack, or put it away"},
	{"weapon remove <name>", "weapon remove Longsword", "Remove a weapon"},
	{"mastery add|remove <weapon>", "mastery add Longsword", "Choose or drop a weapon kind whose mastery property you use"},
	{`mod add <target> ±n "<source>"`, `mod add ac +1 "Ring of Protection"`, "Add a custom bonus or penalty"},
	{`mod remove <target> "<source>"`, `mod remove ac "Ring of Protection"`, "Remove a custom modifier"},
	{"source <kind> <slug> <source>", "source spell shield XPHB", "Point a class, subclass, spell, species, background, feat or armor at another source"},
	{"change-subclass <class> <slug> <source>", "change-subclass fighter champion XPHB", "Swap a class's subclass (owner only); its features replace the old ones"},
}

// Verbs are the command words Parse knows, for the command bar's checks.
func Verbs() []string { return slices.Sorted(maps.Keys(parsers)) }

// Completions are the values the command bar offers for slot, a Help
// placeholder such as "<condition>", after the words typed before it.
// Numbers, free text and compendium look-ups have none here.
func Completions(s Sheet, words []string, slot string) []string {
	word := func(i int) string {
		if i < len(words) {
			return strings.ToLower(words[i])
		}
		return ""
	}
	var out []string
	switch slot {
	case "<condition>":
		return Conditions
	case "<ab>":
		return Abilities
	case "<skill>":
		for _, sk := range Skills {
			out = append(out, sk.Key)
		}
	case "<spell>", "<slug>":
		// concentrate takes the spell's name, spell commands its slug.
		for _, sp := range s.Spells {
			if slot == "<spell>" {
				out = append(out, sp.Name)
			} else {
				out = append(out, sp.Slug)
			}
		}
	case `"<feature>"`:
		for _, f := range s.Features {
			out = append(out, f.Name)
		}
	case "<weapon>":
		if word(1) == verbRemove {
			return s.Masteries
		}
		for _, w := range s.Weapons {
			out = append(out, WeaponKind(w.Name))
		}
	case `"<name>"`, "<name>", `<slug|"item">`:
		switch {
		case word(0) == verbFeature:
			for _, f := range s.Features {
				out = append(out, f.Name)
			}
		case word(0) == verbWeapon && word(1) != verbReady:
			for _, w := range s.Weapons {
				out = append(out, w.Name)
			}
		case word(0) == verbProf && word(2) == ModSkill:
			return Completions(s, nil, "<skill>")
		case word(0) == verbProf && word(2) == ModSave:
			return Abilities
		case word(0) == verbProf:
		default:
			for _, it := range s.Inventory {
				out = append(out, it.Name)
			}
		}
	}
	return out
}

// Parse turns user input into a Command. raw arguments follow the input's
// as they are, so forms can pass names and free text without quoting.
func Parse(input string, raw ...string) (Command, error) {
	args, err := tokenize(input)
	if err != nil {
		return Command{}, err
	}
	args = append(args, raw...)
	if len(args) == 0 {
		return Command{}, ErrUnknownCommand
	}
	verb, rest := strings.ToLower(args[0]), args[1:]
	p, ok := parsers[verb]
	if !ok {
		return Command{}, fmt.Errorf("%w %q", ErrUnknownCommand, verb)
	}
	cmd, err := p(rest)
	if err != nil {
		return Command{}, fmt.Errorf("%s: %w", verb, err)
	}
	cmd.Text = strings.Join(args, " ")
	return cmd, nil
}

type parser func(args []string) (Command, error)

// Command names (the parsers keys).
const (
	verbHP             = "hp"
	verbTemp           = "temp"
	verbMaxHP          = "maxhp"
	verbToggle         = "toggle"
	verbCondition      = "condition"
	verbConcentrate    = "concentrate"
	verbDeathSave      = "deathsave"
	verbRevive         = "revive"
	verbSlot           = "slot"
	verbPact           = "pact"
	verbRest           = "rest"
	verbHitDie         = "hitdie"
	verbItem           = "item"
	verbCoin           = "coin"
	verbInspiration    = "inspiration"
	verbSet            = "set"
	verbProf           = "prof"
	verbSpell          = "spell"
	verbFeature        = "feature"
	verbWeapon         = "weapon"
	verbMastery        = "mastery"
	verbMod            = "mod"
	verbUse            = "use"
	verbRestore        = "restore"
	verbArmor          = "armor"
	verbShield         = "shield"
	verbSource         = "source"
	verbChangeSubclass = "change-subclass"
)

// Subcommand words the parsers switch on.
const (
	verbAdd       = "add"
	verbRemove    = "remove"
	verbPrepare   = "prepare"
	verbUnprepare = "unprepare"
	verbEnd       = "end"
	verbSuccess   = "success"
	verbFail      = "fail"
	verbCrit      = "crit"
	verbReset     = "reset"
	verbMax       = "max"
	verbNote      = "note"
	verbWear      = "wear"
	verbOn        = "on"
	verbOff       = "off"
	verbReady     = "ready"
	verbStow      = "stow"
	verbEdit      = "edit"
)

var parsers map[string]parser

func init() {
	parsers = map[string]parser{
		verbHP:          parseHP,
		verbTemp:        parseTemp,
		verbMaxHP:       parseMaxHP,
		verbToggle:      parseToggle,
		verbCondition:   parseCondition,
		verbConcentrate: parseConcentrate,
		Exhaustion:      parseExhaustion,
		verbDeathSave:   parseDeathSave,
		verbRevive:      parseRevive,
		verbSlot:        parseSlot,
		verbPact:        parsePact,
		verbRest:        parseRest,
		verbHitDie:      parseHitDie,
		verbItem:        parseItem,
		verbCoin:        parseCoin,
		verbInspiration: parseInspiration,
		verbSet:         parseSet,
		verbProf:        parseProf,
		verbSpell:       parseSpell,
		verbFeature:     parseFeature,
		verbWeapon:      parseWeapon,
		verbMastery:     parseMastery,
		verbMod:         parseMod,
		verbUse:         func(args []string) (Command, error) { return parseUse(args, false) },
		verbRestore:     func(args []string) (Command, error) { return parseUse(args, true) },
		verbArmor:       parseArmor,
		verbShield:      parseShield,
		verbSource:      parseSource,
		// Owner only (ScopeSheet), unlike source: it rebuilds the character.
		verbChangeSubclass: func(args []string) (Command, error) {
			if err := want(args, 3, "change-subclass <class> <slug> <source>"); err != nil {
				return Command{}, err
			}
			ref := SubclassRef{Class: strings.ToLower(args[0]), Slug: strings.ToLower(args[1]), Source: args[2]}
			cmd := scoped(perm.ScopeSheet, func(c *Character, env Env, _ *Effects) error { return c.ChangeSubclass(ref, env) })
			cmd.SubclassChange = ref
			return cmd, nil
		},
	}
}

// parseSource re-points a compendium ref to another source:
// source <kind> <slug> <source>.
func parseSource(args []string) (Command, error) {
	usage := "source " + strings.Join(RefKinds, "|") + " <slug> <source>"
	if err := want(args, 3, usage); err != nil {
		return Command{}, err
	}
	ref := SourceRef{Kind: strings.ToLower(args[0]), Slug: strings.ToLower(args[1]), Source: args[2]}
	if !slices.Contains(RefKinds, ref.Kind) {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	cmd := scoped(perm.ScopeSources, func(c *Character, env Env, _ *Effects) error { return c.SwitchSource(ref, env) })
	cmd.SourceRef = ref
	return cmd, nil
}

// MaxAmount bounds an HP-style number; views.MaxHPInput mirrors it for the
// combat dialog (views must not import dnd5e).
const MaxAmount = 9999

// Amount is an HP-style change: a delta, or a set when set is true. Parse
// builds them from "+n", "-n" and "=n".
type Amount struct {
	n    int
	set  bool
	crit bool // a critical hit: two death save failures at 0 HP
}

func (a Amount) applyTo(v int) int {
	if a.set {
		return a.n
	}
	return v + a.n
}

// parseAmount reads "+n", "-n", "=n" or a bare n. bareSets says whether a
// bare number sets the value (temp hp) or adds to it.
func parseAmount(s string, bareSets bool) (Amount, error) {
	set := bareSets
	switch {
	case strings.HasPrefix(s, "="):
		set, s = true, s[1:]
	case strings.HasPrefix(s, "+"), strings.HasPrefix(s, "-"):
		set = false
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return Amount{}, fmt.Errorf("%q is not a number", s)
	}
	if set && n < 0 {
		return Amount{}, errors.New("value cannot be negative")
	}
	if n > MaxAmount || n < -MaxAmount {
		return Amount{}, fmt.Errorf("amount cannot exceed %d", MaxAmount)
	}
	return Amount{n: n, set: set}, nil
}

func want(args []string, n int, usage string) error {
	if len(args) != n {
		return fmt.Errorf("usage: %s", usage)
	}
	return nil
}

func scoped(scope perm.CharacterScope, fn func(c *Character, env Env, fx *Effects) error) Command {
	return Command{Scope: scope, apply: fn}
}

// tokenize splits on whitespace, keeping "double quoted" runs together.
// Quoted runs read Go escapes, as the server quotes names with strconv.Quote.
func tokenize(s string) ([]string, error) {
	var out []string
	var cur strings.Builder
	have := false
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		switch {
		case r == '"':
			q, err := strconv.QuotedPrefix(s[i:])
			if err != nil {
				return nil, errors.New("unclosed quote")
			}
			u, _ := strconv.Unquote(q)
			cur.WriteString(u)
			have, size = true, len(q)
		case unicode.IsSpace(r):
			if have {
				out = append(out, cur.String())
				cur.Reset()
				have = false
			}
		default:
			cur.WriteRune(r)
			have = true
		}
		i += size
	}
	if have {
		out = append(out, cur.String())
	}
	return out, nil
}
