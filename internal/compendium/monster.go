package compendium

import (
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Monster action types, in statblock order.
var actionTypes = []string{"action", "bonus_action", "reaction", "legendary_action"}

// searchText is a statblock's plain text for search.
func (d MonsterData) searchText() string {
	var acts []string
	for _, a := range d.Actions {
		acts = append(acts, a.Name+". "+a.Desc)
	}
	return joinText(d.Size+" "+d.Type, partsText(d.Traits), strings.Join(acts, "\n"))
}

// actionsText writes the actions of one type as ParseTraits reads them.
func actionsText(actions []MonsterAction, typ string) string {
	var parts []NameDesc
	for _, a := range actions {
		if a.Type == typ {
			parts = append(parts, NameDesc{Name: a.Name, Desc: a.Desc})
		}
	}
	return TraitsText(parts)
}

// parseActions reads each type's text (ParseTraits) into actions.
func parseActions(texts ...string) []MonsterAction {
	var out []MonsterAction
	for i, text := range texts {
		for _, p := range ParseTraits(text) {
			out = append(out, MonsterAction{Name: p.Name, Desc: p.Desc, Type: actionTypes[i]})
		}
	}
	return out
}

// bonusText writes bonuses as "dex +5, stealth +6".
func bonusText(m map[string]int) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.SortFunc(keys, func(a, b string) int {
		// Abilities in sheet order first, then skills alphabetically.
		ia, ib := slices.Index(abilityOrder, a), slices.Index(abilityOrder, b)
		if ia >= 0 && ib >= 0 {
			return ia - ib
		}
		if ia >= 0 || ib >= 0 {
			return ib - ia
		}
		return strings.Compare(a, b)
	})
	parts := make([]string, len(keys))
	for i, k := range keys {
		parts[i] = fmt.Sprintf("%s %+d", k, m[k])
	}
	return strings.Join(parts, ", ")
}

var (
	abilityOrder = []string{"str", "dex", "con", "int", "wis", "cha"}
	reBonus      = regexp.MustCompile(`^(.+?)\s*([-+−]\s*\d+)$`)
	reSpeed      = regexp.MustCompile(`(?i)(?:([a-z]+)\s+)?(\d+)\s*(?:ft\.?|feet)?`)
)

// parseBonuses reads "Dex +5, Wis +3" (saves, keyed by ability) or
// "Perception +5, Stealth +6" (skills, keyed by slug).
func parseBonuses(s string, saves bool) map[string]int {
	out := map[string]int{}
	for _, part := range strings.Split(s, ",") {
		m := reBonus.FindStringSubmatch(strings.TrimSpace(part))
		if m == nil {
			continue
		}
		n, err := strconv.Atoi(strings.NewReplacer(" ", "", "−", "-").Replace(m[2]))
		if err != nil {
			continue
		}
		key := Slugify(m[1])
		if saves {
			if key = abilityKey(key); key == "" {
				continue
			}
		}
		out[key] = n
	}
	return out
}

// abilityKey is "dex" for "dex" or "dexterity", else "".
func abilityKey(name string) string {
	if len(name) >= 3 && slices.Contains(abilityOrder, name[:3]) {
		return name[:3]
	}
	return ""
}

// speedText writes speeds as "walk 40, fly 60".
func speedText(m map[string]int) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.SortFunc(keys, func(a, b string) int {
		// Walking speed first, as statblocks write it.
		return strings.Compare(strings.Replace(a, "walk", "", 1), strings.Replace(b, "walk", "", 1))
	})
	parts := make([]string, len(keys))
	for i, k := range keys {
		parts[i] = fmt.Sprintf("%s %d", k, m[k])
	}
	return strings.Join(parts, ", ")
}

// parseSpeed reads "40 ft., fly 60 ft. (hover)" or "walk 40, fly 60"; a
// speed with no mode is walking.
func parseSpeed(s string) map[string]int {
	out := map[string]int{}
	for _, part := range strings.Split(s, ",") {
		m := reSpeed.FindStringSubmatch(part)
		if m == nil {
			continue
		}
		n, _ := strconv.Atoi(m[2])
		out[strings.ToLower(cmp.Or(m[1], "walk"))] = n
	}
	return out
}
