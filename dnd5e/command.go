package dnd5e

import (
	"cmp"
	"errors"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"gm-session-manager/perm"
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
	apply     func(c *Character, env Env, fx *Effects) error
}

// Env carries what a command needs from outside the character.
type Env struct {
	// Spell is the resolved compendium spell for Command.SpellSlug.
	Spell Spell
	// Armor is the resolved compendium armor for Command.ArmorSlug;
	// IsShield when that is a shield instead.
	Armor    WornArmor
	IsShield bool
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
	s.Languages = slices.Clone(s.Languages)
	s.Spells = slices.Clone(s.Spells)
	s.Inventory = slices.Clone(s.Inventory)
	s.Features = slices.Clone(s.Features)
	s.Feats = slices.Clone(s.Feats)
	s.Weapons = slices.Clone(s.Weapons)
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
	{"hp -n | +n | =n", "hp -5", "Take damage, heal, or set hit points"},
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
	{"inspiration", "inspiration", "Toggle Heroic Inspiration"},
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
	{`mod add <target> ±n "<source>"`, `mod add ac +1 "Ring of Protection"`, "Add a custom bonus or penalty"},
	{`mod remove <target> "<source>"`, `mod remove ac "Ring of Protection"`, "Remove a custom modifier"},
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
	case `"<name>"`, "<name>", `<slug|"item">`:
		switch {
		case word(0) == "feature":
			for _, f := range s.Features {
				out = append(out, f.Name)
			}
		case word(0) == "weapon" && word(1) != "ready":
			for _, w := range s.Weapons {
				out = append(out, w.Name)
			}
		case word(0) == "prof" && word(2) == "skill":
			return Completions(s, nil, "<skill>")
		case word(0) == "prof" && word(2) == "save":
			return Abilities
		case word(0) == "prof":
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

var parsers map[string]parser

func init() {
	parsers = map[string]parser{
		"hp":          parseHP,
		"temp":        parseTemp,
		"maxhp":       parseMaxHP,
		"toggle":      parseToggle,
		"condition":   parseCondition,
		"concentrate": parseConcentrate,
		"exhaustion":  parseExhaustion,
		"deathsave":   parseDeathSave,
		"slot":        parseSlot,
		"pact":        parsePact,
		"rest":        parseRest,
		"hitdie":      parseHitDie,
		"item":        parseItem,
		"coin":        parseCoin,
		"inspiration": parseInspiration,
		"set":         parseSet,
		"prof":        parseProf,
		"spell":       parseSpell,
		"feature":     parseFeature,
		"weapon":      parseWeapon,
		"mod":         parseMod,
		"use":         func(args []string) (Command, error) { return parseUse(args, false) },
		"restore":     func(args []string) (Command, error) { return parseUse(args, true) },
		"armor":       parseArmor,
		"shield":      parseShield,
	}
}

// Amount is an HP-style change: a delta, or a set when set is true. Parse
// builds them from "+n", "-n" and "=n".
type Amount struct {
	n   int
	set bool
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

func parseHP(args []string) (Command, error) {
	if err := want(args, 1, "hp -5 | hp +3 | hp =20"); err != nil {
		return Command{}, err
	}
	if !strings.ContainsAny(args[0][:1], "+-=") {
		return Command{}, errors.New("use -n for damage, +n for healing or =n to set")
	}
	a, err := parseAmount(args[0], false)
	if err != nil {
		return Command{}, err
	}
	return scoped(perm.ScopeHP, func(c *Character, _ Env, _ *Effects) error {
		c.ApplyHP(a)
		return nil
	}), nil
}

// ApplyHP applies damage (temporary HP first), healing (capped at max) or a
// set. HP may go below 0 so overkill stays visible; healing back above 0
// resets death saves.
func (c *Character) ApplyHP(a Amount) {
	wasDown := c.HP <= 0
	switch {
	case a.set:
		c.HP = a.n
	case a.n < 0:
		dmg := -a.n
		absorbed := min(dmg, c.TempHP)
		c.TempHP -= absorbed
		c.HP -= dmg - absorbed
	default:
		c.HP += a.n
	}
	if c.MaxHP > 0 {
		c.HP = min(c.HP, c.MaxHP)
	}
	if wasDown && c.HP > 0 {
		c.Sheet.DeathSaves = DeathSaves{}
	}
}

// Damage and Heal are the HP deltas the combat tracker uses.
func Damage(n int) Amount { return Amount{n: -n} }

// Heal returns a healing Amount.
func Heal(n int) Amount { return Amount{n: n} }

// Set returns an Amount that sets the value to n.
func Set(n int) Amount { return Amount{n: n, set: true} }

func parseTemp(args []string) (Command, error) {
	if err := want(args, 1, "temp 5 | temp +2 | temp 0"); err != nil {
		return Command{}, err
	}
	a, err := parseAmount(args[0], true)
	if err != nil {
		return Command{}, err
	}
	return scoped(perm.ScopeHP, func(c *Character, _ Env, _ *Effects) error {
		c.TempHP = max(a.applyTo(c.TempHP), 0)
		return nil
	}), nil
}

func parseMaxHP(args []string) (Command, error) {
	if err := want(args, 1, "maxhp +5 | maxhp =30"); err != nil {
		return Command{}, err
	}
	a, err := parseAmount(args[0], false)
	if err != nil {
		return Command{}, err
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		c.MaxHP = max(a.applyTo(c.MaxHP), 0)
		c.HP = min(c.HP, c.MaxHP)
		return nil
	}), nil
}

func conditionSlug(s string) (string, error) {
	slug := strings.ToLower(s)
	if !IsConditionSlug(slug) || slug == "exhaustion" {
		return "", fmt.Errorf("unknown condition %q", s)
	}
	return slug, nil
}

func condCommand(op CondOp) Command {
	return scoped(perm.ScopeConditions, func(_ *Character, _ Env, fx *Effects) error {
		fx.Conditions = append(fx.Conditions, op)
		return nil
	})
}

func parseToggle(args []string) (Command, error) {
	if len(args) == 2 && strings.EqualFold(args[0], "condition") {
		args = args[1:]
	}
	if err := want(args, 1, "toggle prone"); err != nil {
		return Command{}, err
	}
	slug, err := conditionSlug(args[0])
	if err != nil {
		return Command{}, err
	}
	return condCommand(CondOp{Op: CondToggle, Slug: slug}), nil
}

func parseCondition(args []string) (Command, error) {
	const usage = "condition add poisoned [note] | condition remove poisoned"
	if len(args) < 2 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	slug, err := conditionSlug(args[1])
	if err != nil {
		return Command{}, err
	}
	switch strings.ToLower(args[0]) {
	case "add":
		return condCommand(CondOp{Op: CondAdd, Slug: slug, Note: strings.Join(args[2:], " ")}), nil
	case "remove":
		return condCommand(CondOp{Op: CondRemove, Slug: slug}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

func parseConcentrate(args []string) (Command, error) {
	if len(args) == 0 {
		return Command{}, errors.New("usage: concentrate <spell> | concentrate end")
	}
	if len(args) == 1 && strings.EqualFold(args[0], "end") {
		return condCommand(CondOp{Op: CondRemove, Slug: TrackConcentration}), nil
	}
	return condCommand(CondOp{Op: CondAdd, Slug: TrackConcentration, Note: strings.Join(args, " ")}), nil
}

func parseExhaustion(args []string) (Command, error) {
	if err := want(args, 1, "exhaustion +1 | exhaustion -1"); err != nil {
		return Command{}, err
	}
	a, err := parseAmount(args[0], false)
	if err != nil || a.set {
		return Command{}, errors.New("use +n or -n")
	}
	return condCommand(CondOp{Op: CondExhaustion, Slug: "exhaustion", Delta: a.n}), nil
}

func parseDeathSave(args []string) (Command, error) {
	if err := want(args, 1, "deathsave success | fail | reset"); err != nil {
		return Command{}, err
	}
	kind := strings.ToLower(args[0])
	if kind != "success" && kind != "fail" && kind != "reset" {
		return Command{}, errors.New("use success, fail or reset")
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		ds := &c.Sheet.DeathSaves
		switch kind {
		case "success":
			ds.Successes = min(ds.Successes+1, 3)
		case "fail":
			ds.Failures = min(ds.Failures+1, 3)
		default:
			*ds = DeathSaves{}
		}
		return nil
	}), nil
}

func slotLevel(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil || n < 1 || n > 9 {
		return 0, fmt.Errorf("slot level must be 1-9, got %q", s)
	}
	return n, nil
}

func parseSlot(args []string) (Command, error) {
	const usage = "slot use 2 | slot restore 2 | slot max 2 3"
	if len(args) < 2 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	lvl, err := slotLevel(args[1])
	if err != nil {
		return Command{}, err
	}
	op := strings.ToLower(args[0])
	switch {
	case op == "use" && len(args) == 2, op == "restore" && len(args) == 2:
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			sl := &c.Sheet.Slots[lvl-1]
			if op == "use" {
				if sl.Used >= sl.Max {
					return fmt.Errorf("no level %d slots left", lvl)
				}
				sl.Used++
			} else {
				sl.Used = max(sl.Used-1, 0)
			}
			return nil
		}), nil
	case op == "max" && len(args) == 3:
		n, err := strconv.Atoi(args[2])
		if err != nil || n < 0 || n > 9 {
			return Command{}, errors.New("slot count must be 0-9")
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			sl := &c.Sheet.Slots[lvl-1]
			sl.Max, sl.Used = n, min(sl.Used, n)
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

func parsePact(args []string) (Command, error) {
	if err := want(args, 1, "pact use | pact restore"); err != nil {
		return Command{}, err
	}
	op := strings.ToLower(args[0])
	if op != "use" && op != "restore" {
		return Command{}, errors.New("use or restore")
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		p := &c.Sheet.Pact
		if op == "use" {
			if p.Used >= p.Max {
				return errors.New("no pact slots left")
			}
			p.Used++
		} else {
			p.Used = max(p.Used-1, 0)
		}
		return nil
	}), nil
}

func parseRest(args []string) (Command, error) {
	if err := want(args, 1, "rest short | rest long"); err != nil {
		return Command{}, err
	}
	switch strings.ToLower(args[0]) {
	case "short":
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			c.Sheet.Pact.Used = 0
			for i := range c.Sheet.Features {
				switch u := &c.Sheet.Features[i].Uses; u.Reset {
				case "short":
					u.Used = 0
				case "short-one":
					u.Used = max(u.Used-1, 0)
				}
			}
			return nil
		}), nil
	case "long":
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, fx *Effects) error {
			c.LongRest()
			fx.Conditions = append(fx.Conditions, CondOp{Op: CondExhaustion, Slug: "exhaustion", Delta: -1})
			return nil
		}), nil
	}
	return Command{}, errors.New("short or long")
}

// LongRest restores HP and slots, clears temporary HP and death saves and
// regains up to half the character's hit dice (minimum one).
func (c *Character) LongRest() {
	c.HP = c.MaxHP
	c.TempHP = 0
	for i := range c.Sheet.Slots {
		c.Sheet.Slots[i].Used = 0
	}
	c.Sheet.Pact.Used = 0
	for i := range c.Sheet.Features {
		c.Sheet.Features[i].Uses.Used = 0
	}
	c.Sheet.DeathSaves = DeathSaves{}
	regain := max(c.Sheet.Level()/2, 1)
	c.Sheet.HitDiceUsed = max(c.Sheet.HitDiceUsed-regain, 0)
}

func parseHitDie(args []string) (Command, error) {
	if err := want(args, 1, "hitdie use | hitdie restore"); err != nil {
		return Command{}, err
	}
	op := strings.ToLower(args[0])
	if op != "use" && op != "restore" {
		return Command{}, errors.New("use or restore")
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		if op == "use" {
			if c.Sheet.HitDiceUsed >= c.Sheet.Level() {
				return errors.New("no hit dice left")
			}
			c.Sheet.HitDiceUsed++
		} else {
			c.Sheet.HitDiceUsed = max(c.Sheet.HitDiceUsed-1, 0)
		}
		return nil
	}), nil
}

func parseItem(args []string) (Command, error) {
	const usage = `item add "Rope" [qty] | item remove Rope [qty] | item set Rope <qty> | item note Rope <text>`
	if len(args) < 2 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	op, name := strings.ToLower(args[0]), args[1]
	if op == "note" {
		note := strings.Join(args[2:], " ")
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := findItem(c.Sheet.Inventory, name)
			if i < 0 {
				return fmt.Errorf("no item %q", name)
			}
			c.Sheet.Inventory[i].Notes = note
			return nil
		}), nil
	}
	qty := 1
	if len(args) == 3 {
		n, err := strconv.Atoi(args[2])
		// Setting to 0 removes the item; adding or removing 0 means nothing.
		if err != nil || n < 0 || n == 0 && op != "set" {
			return Command{}, errors.New("quantity must be a positive number")
		}
		qty = n
	} else if len(args) > 3 || op == "set" {
		return Command{}, fmt.Errorf("usage: %s (quote names with spaces)", usage)
	}
	switch op {
	case "set":
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := findItem(c.Sheet.Inventory, name)
			switch {
			case i < 0:
				return fmt.Errorf("no item %q", name)
			case qty == 0:
				c.Sheet.unequip(c.Sheet.Inventory[i])
				c.Sheet.Inventory = slices.Delete(c.Sheet.Inventory, i, i+1)
			default:
				c.Sheet.Inventory[i].Qty = qty
			}
			return nil
		}), nil
	case "add":
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			if i := findItem(c.Sheet.Inventory, name); i >= 0 {
				c.Sheet.Inventory[i].Qty += qty
			} else {
				c.Sheet.Inventory = append(c.Sheet.Inventory, Item{Name: name, Qty: qty})
			}
			return nil
		}), nil
	case "remove":
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := findItem(c.Sheet.Inventory, name)
			if i < 0 {
				return fmt.Errorf("no item %q", name)
			}
			c.Sheet.Inventory[i].Qty -= qty
			if c.Sheet.Inventory[i].Qty <= 0 {
				c.Sheet.unequip(c.Sheet.Inventory[i])
				c.Sheet.Inventory = slices.Delete(c.Sheet.Inventory, i, i+1)
			}
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

// unequip takes off or stows an item leaving the inventory, so the sheet
// keeps no armor or attack the character no longer has.
func (s *Sheet) unequip(it Item) {
	if s.Worn(it) {
		if it.Shield {
			s.Shield = false
		} else {
			s.Armor = nil
		}
	}
	if s.Ready(it) {
		s.Weapons = slices.Delete(s.Weapons, s.weaponIndex(it.Name), s.weaponIndex(it.Name)+1)
	}
}

func findItem(items []Item, name string) int {
	return slices.IndexFunc(items, func(it Item) bool { return strings.EqualFold(it.Name, name) })
}

// parseCoin reads one or more coin changes ("gp +10 sp =3"); all apply or
// none do.
func parseCoin(args []string) (Command, error) {
	if len(args) == 0 || len(args)%2 != 0 {
		return Command{}, errors.New("usage: coin gp +10 [sp -3 …]")
	}
	type change struct {
		kind string
		pick func(*Coins) *int
		a    Amount
	}
	var changes []change
	for i := 0; i < len(args); i += 2 {
		a, err := parseAmount(args[i+1], false)
		if err != nil {
			return Command{}, err
		}
		kind := strings.ToLower(args[i])
		pick := coinPicks[kind]
		if pick == nil {
			return Command{}, errors.New("coin must be cp, sp, ep, gp or pp")
		}
		changes = append(changes, change{kind, pick, a})
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		for _, ch := range changes {
			p := ch.pick(&c.Sheet.Coins)
			next := ch.a.applyTo(*p)
			if next < 0 {
				return fmt.Errorf("only %d %s", *p, ch.kind)
			}
			*p = next
		}
		return nil
	}), nil
}

var coinPicks = map[string]func(*Coins) *int{
	"cp": func(c *Coins) *int { return &c.CP }, "sp": func(c *Coins) *int { return &c.SP },
	"ep": func(c *Coins) *int { return &c.EP }, "gp": func(c *Coins) *int { return &c.GP },
	"pp": func(c *Coins) *int { return &c.PP },
}

func parseInspiration(args []string) (Command, error) {
	if err := want(args, 0, "inspiration"); err != nil {
		return Command{}, err
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		c.Sheet.Inspiration = !c.Sheet.Inspiration
		return nil
	}), nil
}

func parseSet(args []string) (Command, error) {
	const usage = "set name|species|background|backstory|vision <text> | set ac|speed <n> | set size <size> | set ability <str..cha> <n>"
	if len(args) < 2 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	field, value := strings.ToLower(args[0]), strings.Join(args[1:], " ")
	switch field {
	case "size":
		size := SizeNamed(value)
		if size == "" {
			return Command{}, fmt.Errorf("size must be one of %s", strings.Join(Sizes, ", "))
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			c.Sheet.Size = size
			return nil
		}), nil
	case "vision":
		if strings.EqualFold(value, "none") {
			value = ""
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			c.Sheet.Vision = value
			return nil
		}), nil
	case "name", "species", "background", "backstory":
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			switch field {
			case "name":
				c.Name = value
			case "species":
				c.Sheet.Species = value
			case "backstory":
				c.Sheet.Backstory = value
			default:
				c.Sheet.Background = value
			}
			return nil
		}), nil
	case "ac", "speed":
		n, err := strconv.Atoi(value)
		if err != nil {
			return Command{}, fmt.Errorf("%s must be a number", field)
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			if field == "ac" {
				c.Sheet.ACBonus = n
			} else {
				c.Sheet.Speed = n
			}
			return nil
		}), nil
	case "ability":
		if len(args) != 3 {
			return Command{}, fmt.Errorf("usage: %s", usage)
		}
		ab := strings.ToLower(args[1])
		n, err := strconv.Atoi(args[2])
		if err != nil || !slices.Contains(Abilities, ab) {
			return Command{}, errors.New("usage: set ability str 16")
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			c.Sheet.Scores[ab] = n
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

// Proficiency levels for "prof set".
const (
	ProfNone       = "none"
	ProfProficient = "proficient"
	ProfExpertise  = "expertise"
)

// ProfLevel is a skill's proficiency level.
func (s Sheet) ProfLevel(skill string) string {
	switch {
	case slices.Contains(s.Expertise, skill):
		return ProfExpertise
	case slices.Contains(s.SkillProfs, skill):
		return ProfProficient
	}
	return ProfNone
}

// setIn adds value to or removes it from list.
func setIn(list *[]string, value string, on bool) {
	i := slices.IndexFunc(*list, func(x string) bool { return strings.EqualFold(x, value) })
	switch {
	case on && i < 0:
		*list = append(*list, value)
	case !on && i >= 0:
		*list = slices.Delete(*list, i, i+1)
	}
}

// parseProfSet sets a skill or save to a level in one step, so going from
// none to expertise is one command: prof set skill|save <key> <level>.
func parseProfSet(args []string) (Command, error) {
	const usage = "prof set skill <skill> none|proficient|expertise | prof set save <ability> none|proficient"
	if len(args) != 4 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	kind, key, level := strings.ToLower(args[1]), strings.ToLower(args[2]), strings.ToLower(args[3])
	switch {
	case kind == "skill" && slices.ContainsFunc(Skills, func(sk Skill) bool { return sk.Key == key }) &&
		slices.Contains([]string{ProfNone, ProfProficient, ProfExpertise}, level):
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			setIn(&c.Sheet.SkillProfs, key, level != ProfNone)
			setIn(&c.Sheet.Expertise, key, level == ProfExpertise)
			return nil
		}), nil
	case kind == "save" && slices.Contains(Abilities, key) && (level == ProfNone || level == ProfProficient):
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			setIn(&c.Sheet.SaveProfs, key, level == ProfProficient)
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

func parseProf(args []string) (Command, error) {
	const usage = "prof add|remove skill|save|expertise|other|language <value> | prof set skill|save <key> <level>"
	if len(args) < 3 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	op, kind, value := strings.ToLower(args[0]), strings.ToLower(args[1]), strings.Join(args[2:], " ")
	if op == "set" {
		return parseProfSet(args)
	}
	if op != "add" && op != "remove" {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	var pick func(s *Sheet) *[]string
	switch kind {
	case "skill", "expertise":
		value = strings.ToLower(value)
		if !slices.ContainsFunc(Skills, func(sk Skill) bool { return sk.Key == value }) {
			return Command{}, fmt.Errorf("unknown skill %q", value)
		}
		pick = func(s *Sheet) *[]string { return &s.SkillProfs }
		if kind == "expertise" {
			pick = func(s *Sheet) *[]string { return &s.Expertise }
		}
	case "save":
		value = strings.ToLower(value)
		if !slices.Contains(Abilities, value) {
			return Command{}, fmt.Errorf("unknown ability %q", value)
		}
		pick = func(s *Sheet) *[]string { return &s.SaveProfs }
	case "other":
		pick = func(s *Sheet) *[]string { return &s.OtherProfs }
	case "language":
		pick = func(s *Sheet) *[]string { return &s.Languages }
	default:
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		setIn(pick(&c.Sheet), value, op == "add")
		return nil
	}), nil
}

func parseSpell(args []string) (Command, error) {
	const usage = "spell add|remove|prepare|unprepare <slug>"
	if err := want(args, 2, usage); err != nil {
		return Command{}, err
	}
	op, slug := strings.ToLower(args[0]), strings.ToLower(args[1])
	find := func(c *Character) int {
		return slices.IndexFunc(c.Sheet.Spells, func(s Spell) bool { return s.Slug == slug })
	}
	switch op {
	case "add":
		cmd := scoped(perm.ScopeSheet, func(c *Character, env Env, _ *Effects) error {
			if env.Spell.Slug != slug {
				return fmt.Errorf("unknown spell %q", slug)
			}
			if find(c) < 0 {
				c.Sheet.Spells = append(c.Sheet.Spells, env.Spell)
			}
			return nil
		})
		cmd.SpellSlug = slug
		return cmd, nil
	case "remove":
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			if i := find(c); i >= 0 {
				c.Sheet.Spells = slices.Delete(c.Sheet.Spells, i, i+1)
			}
			return nil
		}), nil
	case "prepare", "unprepare":
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := find(c)
			if i < 0 {
				return fmt.Errorf("spell %q is not on the sheet", slug)
			}
			c.Sheet.Spells[i].Prepared = op == "prepare"
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

func parseFeature(args []string) (Command, error) {
	const usage = `feature add "Name" [description] | feature set "Name" key=value… | feature remove "Name"`
	if len(args) < 2 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	op, name := strings.ToLower(args[0]), args[1]
	find := func(c *Character) int {
		return slices.IndexFunc(c.Sheet.Features, func(f Feature) bool { return strings.EqualFold(f.Name, name) })
	}
	switch op {
	case "add":
		desc := strings.Join(args[2:], " ")
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			if i := find(c); i >= 0 {
				c.Sheet.Features[i].Description = desc
			} else {
				c.Sheet.Features = append(c.Sheet.Features, Feature{Name: name, Description: desc})
			}
			return nil
		}), nil
	case "remove":
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			if i := find(c); i >= 0 {
				name := c.Sheet.Features[i].Name
				c.Sheet.Features = slices.Delete(c.Sheet.Features, i, i+1)
				c.Sheet.Modifiers = slices.DeleteFunc(c.Sheet.Modifiers, func(m CustomMod) bool { return m.Source == name })
			}
			return nil
		}), nil
	case "set":
		set, err := parseFeatureSet(args[2:])
		if err != nil {
			return Command{}, err
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := find(c)
			if i < 0 {
				return fmt.Errorf("no feature %q", name)
			}
			set(&c.Sheet.Features[i])
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

// parseFeatureSet reads key=value settings of a feature: action
// (none|action|bonus|reaction|magic), uses (none|n|ability|prof) and reset.
func parseFeatureSet(args []string) (func(f *Feature), error) {
	if len(args) == 0 {
		return nil, errors.New("usage: feature set \"Name\" action=bonus uses=2 reset=long")
	}
	var steps []func(f *Feature)
	for _, a := range args {
		k, v, ok := strings.Cut(strings.ToLower(a), "=")
		switch {
		case !ok:
			return nil, fmt.Errorf("%q is not key=value", a)
		case k == "action" && (v == "none" || ActionLabels[v] != ""):
			steps = append(steps, func(f *Feature) { f.Action = strings.TrimPrefix(v, "none") })
		case k == "uses":
			u, err := parseUses(v)
			if err != nil {
				return nil, err
			}
			steps = append(steps, func(f *Feature) {
				u.Reset = cmp.Or(f.Uses.Reset, "long")
				f.Uses = u
			})
		case k == "reset" && slices.Contains([]string{"short", "long", "short-one"}, v):
			steps = append(steps, func(f *Feature) { f.Uses.Reset = v })
		default:
			return nil, fmt.Errorf("unknown setting %q", a)
		}
	}
	return func(f *Feature) {
		for _, s := range steps {
			s(f)
		}
	}, nil
}

// parseUses reads a uses setting: none, a count, an ability key or prof.
func parseUses(v string) (Uses, error) {
	switch {
	case v == "none":
		return Uses{}, nil
	case v == "prof":
		return Uses{Prof: true}, nil
	case slices.Contains(Abilities, v):
		return Uses{Ability: v}, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil || n < 1 || n > 99 {
		return Uses{}, errors.New("uses must be none, a number, an ability (cha) or prof")
	}
	return Uses{Fixed: n}, nil
}

// parseUse spends ("use") or regains ("restore") one use of a feature.
func parseUse(args []string, restore bool) (Command, error) {
	if err := want(args, 1, `use "Feature" | restore "Feature"`); err != nil {
		return Command{}, err
	}
	name := args[0]
	return Command{Scope: perm.ScopeSheet, apply: func(c *Character, _ Env, _ *Effects) error {
		i := slices.IndexFunc(c.Sheet.Features, func(f Feature) bool { return strings.EqualFold(f.Name, name) })
		if i < 0 {
			return fmt.Errorf("no feature %q", name)
		}
		u := &c.Sheet.Features[i].Uses
		if !u.Limited() {
			return fmt.Errorf("%s has no limited uses", c.Sheet.Features[i].Name)
		}
		if restore {
			u.Used = max(u.Used-1, 0)
			return nil
		}
		if u.Used >= c.Sheet.MaxUses(*u) {
			return fmt.Errorf("no uses of %s left", c.Sheet.Features[i].Name)
		}
		u.Used++
		return nil
	}}, nil
}

// parseArmor puts on armor (a shield goes in the off hand) from the
// inventory, or from the compendium adding it to the inventory, or takes
// armor off; it stays in the inventory.
func parseArmor(args []string) (Command, error) {
	switch {
	case len(args) == 1 && strings.EqualFold(args[0], "off"):
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			c.Sheet.Armor = nil
			return nil
		}), nil
	case len(args) == 2 && strings.EqualFold(args[0], "wear"):
		cmd := scoped(perm.ScopeSheet, func(c *Character, env Env, _ *Effects) error {
			sh := &c.Sheet
			if i := sh.ArmorItem(args[1]); i >= 0 {
				if it := sh.Inventory[i]; it.Armor != nil {
					a := *it.Armor
					sh.Armor = &a
				} else {
					sh.Shield = true
				}
				return nil
			}
			if env.IsShield {
				sh.takeShield()
				return nil
			}
			a, stats := env.Armor, env.Armor
			sh.Armor = &a
			sh.Inventory = ensureItem(sh.Inventory, Item{Name: a.Name, Armor: &stats})
			return nil
		})
		cmd.ArmorSlug = strings.ToLower(args[1])
		return cmd, nil
	}
	return Command{}, errors.New("usage: armor wear <slug> | armor off")
}

func parseShield(args []string) (Command, error) {
	if err := want(args, 1, "shield on | shield off"); err != nil {
		return Command{}, err
	}
	on := strings.EqualFold(args[0], "on")
	if !on && !strings.EqualFold(args[0], "off") {
		return Command{}, errors.New("on or off")
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		if on {
			c.Sheet.takeShield()
		} else {
			c.Sheet.Shield = false
		}
		return nil
	}), nil
}

// takeShield takes up a shield, listing one in the inventory if none is.
func (s *Sheet) takeShield() {
	s.Shield = true
	if !slices.ContainsFunc(s.Inventory, func(it Item) bool { return it.Shield }) {
		s.Inventory = ensureItem(s.Inventory, shieldItem)
	}
}

func parseWeapon(args []string) (Command, error) {
	const usage = "weapon add <name> <str..cha|finesse> <dice> [damage type] [prof] | weapon ready|stow|remove <name>"
	if len(args) < 2 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	op, name := strings.ToLower(args[0]), args[1]
	find := func(c *Character) int {
		return slices.IndexFunc(c.Sheet.Weapons, func(w Weapon) bool { return strings.EqualFold(w.Name, name) })
	}
	switch {
	case op == "add" && len(args) >= 4:
		w := Weapon{Name: name, Ability: strings.ToLower(args[2]), Dice: strings.ToLower(args[3])}
		rest := args[4:]
		if n := len(rest); n > 0 && strings.EqualFold(rest[n-1], "prof") {
			w.Proficient, rest = true, rest[:n-1]
		}
		w.DamageType = strings.Join(rest, " ")
		return SaveWeapon(w), nil
	case op == "ready" && len(args) == 2:
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := findItem(c.Sheet.Inventory, name)
			if i < 0 || c.Sheet.Inventory[i].Weapon == nil {
				return fmt.Errorf("no weapon %q in the inventory", name)
			}
			if find(c) < 0 {
				c.Sheet.Weapons = append(c.Sheet.Weapons, *c.Sheet.Inventory[i].Weapon)
			}
			return nil
		}), nil
	case op == "stow" && len(args) == 2:
		// Its modifiers stay, for when it is readied again.
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := find(c)
			if i < 0 {
				return fmt.Errorf("no weapon %q", name)
			}
			c.Sheet.Weapons = slices.Delete(c.Sheet.Weapons, i, i+1)
			return nil
		}), nil
	case op == "remove" && len(args) == 2:
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := find(c)
			if i < 0 {
				return fmt.Errorf("no weapon %q", name)
			}
			w := c.Sheet.Weapons[i]
			c.Sheet.Weapons = slices.Delete(c.Sheet.Weapons, i, i+1)
			c.Sheet.Modifiers = slices.DeleteFunc(c.Sheet.Modifiers, func(m CustomMod) bool {
				return m.Target == ModAttack+"."+w.Name || m.Target == ModDamage+"."+w.Name
			})
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

// SaveWeapon readies a weapon, or replaces the one of the same name, and
// keeps its stats on its inventory item.
func SaveWeapon(w Weapon) Command {
	return NewCommand("weapon add "+strconv.Quote(w.Name), perm.ScopeSheet, func(c *Character) error {
		if i := c.Sheet.weaponIndex(w.Name); i >= 0 {
			c.Sheet.Weapons[i] = w
		} else {
			c.Sheet.Weapons = append(c.Sheet.Weapons, w)
		}
		stats := w
		c.Sheet.Inventory = ensureItem(c.Sheet.Inventory, Item{Name: w.Name, Weapon: &stats})
		return nil
	})
}

func parseMod(args []string) (Command, error) {
	const usage = "mod add <target> <+n|-n> <note> | mod remove <target> <note> | mod edit <target> <note> <new target> <+n|-n> <new note>"
	if len(args) < 3 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	op, target := strings.ToLower(args[0]), args[1]
	if !IsModTarget(target) {
		return Command{}, fmt.Errorf("unknown modifier target %q", target)
	}
	switch op {
	case "add":
		if len(args) < 4 {
			return Command{}, fmt.Errorf("usage: %s (the note says where it comes from)", usage)
		}
		m, err := newMod(target, args[2], args[3:])
		if err != nil {
			return Command{}, err
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			if i := findMod(c, m.Target, m.Note); i >= 0 {
				c.Sheet.Modifiers[i] = m
			} else {
				c.Sheet.Modifiers = append(c.Sheet.Modifiers, m)
			}
			return nil
		}), nil
	case "edit":
		if len(args) < 6 {
			return Command{}, fmt.Errorf("usage: %s", usage)
		}
		note := args[2]
		m, err := newMod(args[3], args[4], args[5:])
		if err != nil {
			return Command{}, err
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := findMod(c, target, note)
			if i < 0 {
				return fmt.Errorf("no modifier %q on %s", note, target)
			}
			if j := findMod(c, m.Target, m.Note); j >= 0 && j != i {
				return fmt.Errorf("there is already a modifier %q on %s", m.Note, m.Target)
			}
			c.Sheet.Modifiers[i] = m
			return nil
		}), nil
	case "remove":
		note := strings.Join(args[2:], " ")
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := findMod(c, target, note)
			if i < 0 {
				return fmt.Errorf("no modifier %q on %s", note, target)
			}
			c.Sheet.Modifiers = slices.Delete(c.Sheet.Modifiers, i, i+1)
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

func newMod(target, value string, note []string) (CustomMod, error) {
	if !IsModTarget(target) {
		return CustomMod{}, fmt.Errorf("unknown modifier target %q", target)
	}
	n, err := strconv.Atoi(value)
	if err != nil || n == 0 {
		return CustomMod{}, errors.New("modifier must be a non-zero number such as +1 or -2")
	}
	return CustomMod{Target: target, Value: n, Note: strings.Join(note, " ")}, nil
}

func findMod(c *Character, target, note string) int {
	return slices.IndexFunc(c.Sheet.Modifiers, func(m CustomMod) bool {
		return m.Target == target && strings.EqualFold(m.Note, note)
	})
}

// tokenize splits on whitespace, keeping "double quoted" runs together.
func tokenize(s string) ([]string, error) {
	var out []string
	var cur strings.Builder
	inQuote, have := false, false
	for _, r := range s {
		switch {
		case r == '"':
			inQuote, have = !inQuote, true
		case unicode.IsSpace(r) && !inQuote:
			if have {
				out = append(out, cur.String())
				cur.Reset()
				have = false
			}
		default:
			cur.WriteRune(r)
			have = true
		}
	}
	if inQuote {
		return nil, errors.New("unclosed quote")
	}
	if have {
		out = append(out, cur.String())
	}
	return out, nil
}
