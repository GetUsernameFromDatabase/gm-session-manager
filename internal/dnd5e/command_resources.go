package dnd5e

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/internal/perm"
)

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
	case op == verbUse && len(args) == 2, op == verbRestore && len(args) == 2:
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			sl := &c.Sheet.Slots[lvl-1]
			if op == verbUse {
				if sl.Used >= sl.Max {
					return fmt.Errorf("no level %d slots left", lvl)
				}
				sl.Used++
			} else {
				sl.Used = max(sl.Used-1, 0)
			}
			return nil
		}), nil
	case op == verbMax && len(args) == 3:
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
	if op != verbUse && op != verbRestore {
		return Command{}, errors.New("use or restore")
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		p := &c.Sheet.Pact
		if op == verbUse {
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
	case ResetShort:
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, fx *Effects) error {
			fx.Conditions = append(fx.Conditions, CondOp{Op: CondEndRest, Rest: ResetShort})
			c.Sheet.Pact.Used = 0
			for i := range c.Sheet.Features {
				switch u := &c.Sheet.Features[i].Uses; u.Reset {
				case ResetShort:
					u.Used = 0
				case ResetShortOne:
					u.Used = max(u.Used-1, 0)
				}
			}
			return nil
		}), nil
	case ResetLong:
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, fx *Effects) error {
			c.LongRest()
			fx.Conditions = append(fx.Conditions, CondOp{Op: CondExhaustion, Slug: Exhaustion, Delta: -1}, CondOp{Op: CondEndRest, Rest: ResetLong})
			return nil
		}), nil
	}
	return Command{}, errors.New("short or long")
}

// LongRest restores HP and slots, clears temporary HP and death saves and
// regains all hit dice.
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
	c.Sheet.HitDiceUsed = 0
}

func parseHitDie(args []string) (Command, error) {
	if err := want(args, 1, "hitdie use | hitdie restore"); err != nil {
		return Command{}, err
	}
	op := strings.ToLower(args[0])
	if op != verbUse && op != verbRestore {
		return Command{}, errors.New("use or restore")
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		if op == verbUse {
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
