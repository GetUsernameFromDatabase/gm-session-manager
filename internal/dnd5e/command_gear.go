package dnd5e

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/internal/perm"
)

func parseItem(args []string) (Command, error) {
	const usage = `item add "Rope" [qty] | item remove Rope [qty] | item set Rope <qty> | item note Rope <text>`
	if len(args) < 2 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	op, name := strings.ToLower(args[0]), args[1]
	if op == verbNote {
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
		if err != nil || n < 0 || n == 0 && op != verbSet {
			return Command{}, errors.New("quantity must be a positive number")
		}
		qty = n
	} else if len(args) > 3 || op == verbSet {
		return Command{}, fmt.Errorf("usage: %s (quote names with spaces)", usage)
	}
	switch op {
	case verbSet:
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
	case verbAdd:
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			if i := findItem(c.Sheet.Inventory, name); i >= 0 {
				c.Sheet.Inventory[i].Qty += qty
			} else {
				c.Sheet.Inventory = append(c.Sheet.Inventory, Item{Name: name, Qty: qty})
			}
			return nil
		}), nil
	case verbRemove:
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
	if err := want(args, 0, verbInspiration); err != nil {
		return Command{}, err
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		c.Sheet.Inspiration = !c.Sheet.Inspiration
		return nil
	}), nil
}

// parseArmor puts on armor (a shield goes in the off hand) from the
// inventory, or from the compendium adding it to the inventory, or takes
// armor off; it stays in the inventory.
func parseArmor(args []string) (Command, error) {
	switch {
	case len(args) == 1 && strings.EqualFold(args[0], verbOff):
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			c.Sheet.Armor = nil
			return nil
		}), nil
	case len(args) == 2 && strings.EqualFold(args[0], verbWear):
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
	on := strings.EqualFold(args[0], verbOn)
	if !on && !strings.EqualFold(args[0], verbOff) {
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
		return slices.IndexFunc(
			c.Sheet.Weapons,
			func(w Weapon) bool { return strings.EqualFold(w.Name, name) },
		)
	}
	switch {
	case op == verbAdd && len(args) >= 4:
		w := Weapon{Name: name, Ability: strings.ToLower(args[2]), Dice: strings.ToLower(args[3])}
		rest := args[4:]
		if n := len(rest); n > 0 && strings.EqualFold(rest[n-1], verbProf) {
			w.Proficient, rest = true, rest[:n-1]
		}
		w.DamageType = strings.Join(rest, " ")
		return SaveWeapon(w), nil
	case op == verbReady && len(args) == 2:
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
	case op == verbStow && len(args) == 2:
		// Its modifiers stay, for when it is readied again.
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := find(c)
			if i < 0 {
				return fmt.Errorf("no weapon %q", name)
			}
			c.Sheet.Weapons = slices.Delete(c.Sheet.Weapons, i, i+1)
			return nil
		}), nil
	case op == verbRemove && len(args) == 2:
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

// parseMastery chooses or drops a mastered weapon kind: mastery add|remove
// <weapon>. Adding one already chosen, or removing one not chosen, does
// nothing, as another device may have done it first.
func parseMastery(args []string) (Command, error) {
	const usage = "mastery add|remove <weapon>"
	if len(args) < 2 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	op, name := strings.ToLower(args[0]), WeaponKind(strings.Join(args[1:], " "))
	switch op {
	case verbAdd:
		cmd := scoped(perm.ScopeSheet, func(c *Character, env Env, _ *Effects) error {
			if env.MasteryWeapon == "" {
				return fmt.Errorf("no weapon %q with a mastery property in the compendium", name)
			}
			s := &c.Sheet
			if s.Masters(env.MasteryWeapon) {
				return nil
			}
			if len(s.Masteries) >= env.MasteryLimit {
				return fmt.Errorf(
					"%d of %d weapon masteries chosen: remove one first",
					len(s.Masteries),
					env.MasteryLimit,
				)
			}
			s.Masteries = append(s.Masteries, env.MasteryWeapon)
			return nil
		})
		cmd.MasteryWeapon = name
		return cmd, nil
	case verbRemove:
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			c.Sheet.Masteries = slices.DeleteFunc(
				c.Sheet.Masteries,
				func(m string) bool { return strings.EqualFold(m, name) },
			)
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

// SaveWeapon readies a weapon, or replaces the one of the same name, and
// keeps its stats on its inventory item.
func SaveWeapon(w Weapon) Command {
	return NewCommand(
		"weapon add "+strconv.Quote(w.Name),
		perm.ScopeSheet,
		func(c *Character) error {
			if i := c.Sheet.weaponIndex(w.Name); i >= 0 {
				c.Sheet.Weapons[i] = w
			} else {
				c.Sheet.Weapons = append(c.Sheet.Weapons, w)
			}
			stats := w
			c.Sheet.Inventory = ensureItem(c.Sheet.Inventory, Item{Name: w.Name, Weapon: &stats})
			return nil
		},
	)
}
