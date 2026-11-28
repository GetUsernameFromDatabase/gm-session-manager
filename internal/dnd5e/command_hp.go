package dnd5e

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/internal/perm"
)

func parseHP(args []string) (Command, error) {
	const usage = "hp -5 | hp -5 fire | hp -5 crit | hp +3 | hp =20"
	crit := len(args) > 1 && strings.EqualFold(args[len(args)-1], verbCrit)
	if crit {
		args = args[:len(args)-1]
	}
	if len(args) == 2 && strings.HasPrefix(args[0], "-") {
		return parseDamage(args, usage, crit)
	}
	if err := want(args, 1, usage); err != nil {
		return Command{}, err
	}
	if args[0] == "" || !strings.ContainsAny(args[0][:1], "+-=") {
		return Command{}, errors.New("use -n for damage, +n for healing or =n to set")
	}
	a, err := parseAmount(args[0], false)
	if err != nil {
		return Command{}, err
	}
	if crit && (a.set || a.n >= 0) {
		return Command{}, errors.New("crit only applies to damage")
	}
	return scoped(perm.ScopeHP, func(c *Character, env Env, _ *Effects) error {
		if !a.set && a.n < 0 { // untyped damage still meets Petrified's resistance
			a = Damage(c.defenses(env).Taken(-a.n, "", -1))
			a.crit = crit
		}
		return c.ApplyHP(a)
	}), nil
}

// parseDamage reads "hp -9 fire": damage of a type, after the sheet's
// defenses.
func parseDamage(args []string, usage string, crit bool) (Command, error) {
	a, err := parseAmount(args[0], false)
	if err != nil {
		return Command{}, err
	}
	kind := strings.ToLower(args[1])
	if !slices.Contains(DamageTypes, kind) {
		return Command{}, fmt.Errorf("unknown damage type %q (usage: %s)", args[1], usage)
	}
	return scoped(perm.ScopeHP, func(c *Character, env Env, _ *Effects) error {
		a := Damage(c.defenses(env).Taken(-a.n, kind, -1))
		a.crit = crit
		return c.ApplyHP(a)
	}), nil
}

func (c *Character) defenses(env Env) Defenses {
	d := c.Sheet.Defenses()
	d.ResistAll = env.Petrified
	return d
}

// ErrDead is what a dead character's healing, HP set, temp HP, hit dice and
// rests answer; revive is the way back.
var ErrDead error = userText("Dead. Use revive.")

// userText is an error whose text is a sentence shown to the user as is.
type userText string

func (e userText) Error() string { return string(e) }

// ApplyHP applies damage (temporary HP first), healing (capped at max) or a
// set. HP may go below 0 so overkill stays visible; healing back above 0
// resets death saves. Damage at 0 HP, or massive damage, costs a PC death
// save failures; a monster's Character has an empty sheet whose saves are
// never stored, so it is unaffected.
func (c *Character) ApplyHP(a Amount) error {
	if c.Sheet.Dead() && (a.set || a.n >= 0) {
		return ErrDead
	}
	wasDown := c.HP <= 0
	switch {
	case a.set:
		c.HP = a.n
	case a.n < 0:
		dmg := -a.n
		absorbed := min(dmg, c.TempHP)
		c.TempHP -= absorbed
		through, before := dmg-absorbed, c.HP
		c.HP -= through
		c.failDeathSaves(before, through, a.crit)
	default:
		c.HP = max(c.HP, 0) + a.n // healing starts from 0, not from the overkill
	}
	if c.MaxHP > 0 {
		c.HP = min(c.HP, c.MaxHP)
	}
	if wasDown && c.HP > 0 {
		c.Sheet.DeathSaves = DeathSaves{}
	}
	return nil
}

// failDeathSaves records what one hit of through damage (after temp HP) did
// to a character who had before HP. Only this hit counts, never the stored
// overkill.
func (c *Character) failDeathSaves(before, through int, crit bool) {
	ds := &c.Sheet.DeathSaves
	if through == 0 || c.Sheet.Dead() {
		return
	}
	overflow := through
	if before > 0 { // dropping to 0 is only dying, unless it overflows max HP
		overflow -= before
	}
	if c.MaxHP > 0 && overflow >= c.MaxHP {
		ds.Failures = MaxDeathSaves
		return
	}
	if before > 0 {
		return
	}
	if ds.Successes >= MaxDeathSaves { // stable: a new hit starts the count over
		*ds = DeathSaves{}
	}
	n := 1
	if crit {
		n = 2
	}
	ds.Failures = min(ds.Failures+n, MaxDeathSaves)
}

// Dead is a PC with three failed death saves.
func (s Sheet) Dead() bool { return s.DeathSaves.Failures >= MaxDeathSaves }

// CritDamage is n damage from a critical hit: two death save failures at 0 HP.
// Only tests call it today; it stays beside Damage, Heal and Set because
// Amount's fields are unexported, so code outside this package can build a
// crit no other way.
func CritDamage(n int) Amount { return Amount{n: -n, crit: true} }

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
		if c.Sheet.Dead() {
			return ErrDead
		}
		c.TempHP = max(a.applyTo(c.TempHP), 0)
		return nil
	}), nil
}

// parseRevive reads "revive [n]": a dead character comes back with n HP
// (default 1, up to max) and clean death saves.
func parseRevive(args []string) (Command, error) {
	n := 1
	if len(args) > 1 {
		return Command{}, errors.New("usage: revive | revive 5")
	}
	if len(args) == 1 {
		var err error
		if n, err = strconv.Atoi(args[0]); err != nil || n < 1 {
			return Command{}, errors.New("n must be a positive whole number")
		}
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		if !c.Sheet.Dead() {
			return userText("Not dead.")
		}
		c.HP = n
		if c.MaxHP > 0 {
			c.HP = min(n, c.MaxHP)
		}
		c.Sheet.DeathSaves = DeathSaves{}
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
	if !IsConditionSlug(slug) || slug == Exhaustion {
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
	if len(args) == 2 && strings.EqualFold(args[0], verbCondition) {
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
	case verbAdd:
		return condCommand(CondOp{Op: CondAdd, Slug: slug, Note: strings.Join(args[2:], " ")}), nil
	case verbRemove:
		return condCommand(CondOp{Op: CondRemove, Slug: slug}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

func parseConcentrate(args []string) (Command, error) {
	if len(args) == 0 {
		return Command{}, errors.New("usage: concentrate <spell> | concentrate end")
	}
	if len(args) == 1 && strings.EqualFold(args[0], verbEnd) {
		return condCommand(CondOp{Op: CondRemove, Slug: TrackConcentration}), nil
	}
	return condCommand(
		CondOp{Op: CondAdd, Slug: TrackConcentration, Note: strings.Join(args, " ")},
	), nil
}

func parseExhaustion(args []string) (Command, error) {
	if err := want(args, 1, "exhaustion +1 | exhaustion -1"); err != nil {
		return Command{}, err
	}
	a, err := parseAmount(args[0], false)
	if err != nil || a.set {
		return Command{}, errors.New("use +n or -n")
	}
	return condCommand(CondOp{Op: CondExhaustion, Slug: Exhaustion, Delta: a.n}), nil
}

func parseDeathSave(args []string) (Command, error) {
	if err := want(args, 1, "deathsave success | fail | reset"); err != nil {
		return Command{}, err
	}
	kind := strings.ToLower(args[0])
	if kind != verbSuccess && kind != verbFail && kind != verbReset {
		return Command{}, errors.New("use success, fail or reset")
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		ds := &c.Sheet.DeathSaves
		switch kind {
		case verbSuccess:
			ds.Successes = min(ds.Successes+1, MaxDeathSaves)
		case verbFail:
			ds.Failures = min(
				ds.Failures+1,
				MaxDeathSaves,
			) // the web layer clears conditions on reaching dead
		default:
			*ds = DeathSaves{}
		}
		return nil
	}), nil
}
