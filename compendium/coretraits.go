package compendium

import (
	"regexp"
	"strconv"
	"strings"
)

// CoreTraits reads a class's core traits table, whose rows are
// "|Primary Ability|Strength|", keyed by the first cell.
func CoreTraits(desc string) map[string]string {
	out := map[string]string{}
	for _, line := range strings.Split(desc, "\n") {
		cells := strings.Split(strings.Trim(strings.TrimSpace(line), "|"), "|")
		if len(cells) == 2 && !strings.HasPrefix(cells[0], "-") {
			out[strings.TrimSpace(cells[0])] = strings.TrimSpace(cells[1])
		}
	}
	return out
}

var chooseN = regexp.MustCompile(`(?i)^choose (?:any )?(\d+)`)

// SkillChoice reads "Choose 2: Athletics, Insight, or Stealth" as how many
// to pick and from which; "Choose any 3 skills" gives no list (any skill).
func SkillChoice(s string) (int, []string) {
	m := chooseN.FindStringSubmatch(strings.TrimSpace(s))
	if m == nil {
		return 0, nil
	}
	n, _ := strconv.Atoi(m[1])
	_, list, ok := strings.Cut(s, ":")
	if !ok {
		return n, nil
	}
	return n, SplitList(list)
}

// SplitList splits "A, B, and C" (or "A and B", "A, or B") into its items.
func SplitList(s string) []string {
	var out []string
	for _, part := range strings.Split(s, ",") {
		for _, item := range strings.Split(part, " and ") {
			item = strings.TrimSpace(item)
			item = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(item, "or "), "and "))
			if item != "" {
				out = append(out, item)
			}
		}
	}
	return out
}

// EquipmentItem is one line of a starting equipment option.
type EquipmentItem struct {
	Name string
	Qty  int
}

// EquipmentOption is one lettered choice of starting equipment.
type EquipmentOption struct {
	Letter string
	Items  []EquipmentItem
	GP     int
}

var (
	optionStart = regexp.MustCompile(`\(([A-Z])\)\s*`)
	gpAmount    = regexp.MustCompile(`^(\d+)\s*GP$`)
	leadingQty  = regexp.MustCompile(`^(\d+)\s+(.+)$`)
)

// EquipmentOptions reads "Choose A or B: (A) Greataxe, 4 Handaxes, and
// 15 GP; or (B) 75 GP" as its options.
func EquipmentOptions(s string) []EquipmentOption {
	s = strings.ReplaceAll(s, "*", "")
	idx := optionStart.FindAllStringSubmatchIndex(s, -1)
	var out []EquipmentOption
	for i, m := range idx {
		end := len(s)
		if i+1 < len(idx) {
			end = idx[i+1][0]
		}
		// Each option but the last ends "; or" or ";".
		body := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(s[m[1]:end]), "or"))
		body = strings.TrimSuffix(body, ";")
		o := EquipmentOption{Letter: s[m[2]:m[3]]}
		for _, item := range SplitList(body) {
			if g := gpAmount.FindStringSubmatch(item); g != nil {
				o.GP, _ = strconv.Atoi(g[1])
				continue
			}
			it := EquipmentItem{Name: item, Qty: 1}
			if q := leadingQty.FindStringSubmatch(item); q != nil {
				it.Qty, _ = strconv.Atoi(q[1])
				it.Name = q[2]
			}
			o.Items = append(o.Items, it)
		}
		out = append(out, o)
	}
	return out
}

// Squash lowercases s and drops spaces, so names split by the source's
// PDF extraction ("In sight", "Ar mor") still match.
func Squash(s string) string {
	return strings.ToLower(strings.Join(strings.Fields(s), ""))
}
