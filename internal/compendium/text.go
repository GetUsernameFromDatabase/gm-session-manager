package compendium

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// The homebrew editor takes traits, features and class table columns as
// plain text a GM can paste from a book; these read and write that text.

// reTraitName is a paragraph's leading "Name." (bold or italic marks allowed).
var reTraitName = regexp.MustCompile(`^[*_]*([^.*_\n]{1,60}?)\.[*_]*\s+`)

// ParseTraits reads paragraphs "Name. Description" as named parts; a
// paragraph without a name joins the one before it.
func ParseTraits(text string) []NameDesc {
	var out []NameDesc
	for _, para := range paragraphs(text) {
		para = strings.TrimSpace(strings.TrimLeft(para, "-*• \t"))
		if m := reTraitName.FindStringSubmatch(para); m != nil {
			out = append(out, NameDesc{Name: strings.TrimSpace(m[1]), Desc: strings.TrimSpace(para[len(m[0]):])})
		} else if n := len(out); n > 0 {
			out[n-1].Desc += "\n\n" + para
		} else {
			out = append(out, NameDesc{Name: "Trait", Desc: para})
		}
	}
	return out
}

// TraitsText writes untyped parts back as ParseTraits reads them.
func TraitsText(parts []NameDesc) string {
	var b []string
	for _, p := range parts {
		if p.Type == "" {
			b = append(b, p.Name+". "+p.Desc)
		}
	}
	return strings.Join(b, "\n\n")
}

var reBlankLine = regexp.MustCompile(`\n\s*\n`)

func paragraphs(text string) []string {
	var out []string
	for _, p := range reBlankLine.Split(strings.ReplaceAll(text, "\r\n", "\n"), -1) {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

// reFeatureHead is a feature heading: "Level 3: Name", "Level 2, 5, 9: Name"
// or "Level 3: Name (subclass)" for the level the subclass is chosen at.
var reFeatureHead = regexp.MustCompile(`(?i)^level\s+(\d+(?:\s*,\s*\d+)*)\s*:\s*(.+?)(\s*\(subclass\))?\s*$`)

// ParseFeatures reads class features: each starts with a heading line
// (reFeatureHead) and its description runs to the next heading.
func ParseFeatures(text string) []ClassFeature {
	var out []ClassFeature
	for _, line := range strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n") {
		if m := reFeatureHead.FindStringSubmatch(strings.TrimSpace(line)); m != nil {
			f := ClassFeature{Name: m[2], Subclass: m[3] != ""}
			for _, l := range strings.Split(m[1], ",") {
				n, _ := strconv.Atoi(strings.TrimSpace(l))
				f.Levels = append(f.Levels, n)
			}
			out = append(out, f)
		} else if n := len(out); n > 0 {
			out[n-1].Desc += line + "\n"
		}
	}
	for i := range out {
		out[i].Desc = strings.TrimSpace(out[i].Desc)
	}
	return out
}

// FeaturesText writes features back as ParseFeatures reads them.
func FeaturesText(fs []ClassFeature) string {
	var b strings.Builder
	for _, f := range fs {
		levels := make([]string, len(f.Levels))
		for i, l := range f.Levels {
			levels[i] = strconv.Itoa(l)
		}
		fmt.Fprintf(&b, "Level %s: %s", strings.Join(levels, ", "), f.Name)
		if f.Subclass {
			b.WriteString(" (subclass)")
		}
		b.WriteString("\n" + f.Desc + "\n\n")
	}
	return strings.TrimSpace(b.String())
}

// ParseTable reads class table columns, one per line: "Focus Points: -, 2,
// 3, …" gives the value at levels 1, 2, 3…; "-" or nothing skips a level.
func ParseTable(text string) map[string]map[int]string {
	out := map[string]map[int]string{}
	for _, line := range strings.Split(text, "\n") {
		name, values, ok := strings.Cut(line, ":")
		if name = strings.TrimSpace(name); !ok || name == "" {
			continue
		}
		col := map[int]string{}
		for i, v := range strings.Split(values, ",") {
			if v = strings.TrimSpace(v); v != "" && v != "-" && i < 20 {
				col[i+1] = v
			}
		}
		out[name] = col
	}
	return out
}

// TableText writes columns back as ParseTable reads them.
func TableText(table map[string]map[int]string) string {
	names := make([]string, 0, len(table))
	for n := range table {
		names = append(names, n)
	}
	slices.Sort(names)
	lines := make([]string, len(names))
	for i, n := range names {
		values := make([]string, 20)
		for l := range values {
			values[l] = table[n][l+1]
			if values[l] == "" {
				values[l] = "-"
			}
		}
		lines[i] = n + ": " + strings.Join(values, ", ")
	}
	return strings.Join(lines, "\n")
}

// classTraits are the core traits rows the character wizard reads (see
// CoreTraits), in table order.
var classTraits = []string{"Skill Proficiencies", "Armor Training", "Weapon Proficiencies", "Tool Proficiencies"}

// withCoreTraits puts a core traits table with the given rows in front of desc.
func withCoreTraits(desc string, rows map[string]string) string {
	var b strings.Builder
	for _, k := range classTraits {
		if v := strings.TrimSpace(rows[k]); v != "" {
			fmt.Fprintf(&b, "|%s|%s|\n", k, v)
		}
	}
	if b.Len() == 0 {
		return desc
	}
	return strings.TrimSpace("|Core Traits||\n|---|---|\n" + b.String() + "\n" + desc)
}

// withoutTable drops table rows from desc, leaving the prose.
func withoutTable(desc string) string {
	var keep []string
	for _, line := range strings.Split(desc, "\n") {
		if !strings.HasPrefix(strings.TrimSpace(line), "|") {
			keep = append(keep, line)
		}
	}
	return strings.TrimSpace(strings.Join(keep, "\n"))
}
