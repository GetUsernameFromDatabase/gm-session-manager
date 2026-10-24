package compendium

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// www.aidedd.org has 2014 monster statblocks with a picture. Monsters
// outside the SRD show their action names without the text. Its other
// pages are read as plain text. A monster's page id is its slug; any other
// page's is its path and query.

const (
	aideddHome = "https://www.aidedd.org/"
	aideddBase = aideddHome + "dnd/monstres.php?vo="
)

var aideddSite = Site{
	Key: "aidedd", Name: "AideDD", Home: "https://www.aidedd.org/dnd-filters/monsters.php",
	Imports: "2014 monsters, with a picture",
	hosts:   []string{"www.aidedd.org", "aidedd.org"},
	page: func(u *url.URL) string {
		if vo := strings.ToLower(u.Query().Get("vo")); u.Path == "/dnd/monstres.php" && reSlug.MatchString(vo) {
			return vo
		}
		return strings.TrimPrefix(u.Path, "/") + cutEmpty("?", u.RawQuery)
	},
	check: func(page string) (Kind, error) {
		switch {
		case reSlug.MatchString(page):
			return KindMonster, nil
		case reAideddPath.MatchString(page):
			return "", nil
		}
		return "", fmt.Errorf("%q is not a www.aidedd.org page", page)
	},
	link: aideddLink,
	load: func(ctx context.Context, page string) (string, string, Entry, error) {
		body, err := fetch(ctx, aideddLink(page), wikidotMaxBody)
		if err != nil {
			return "", "", Entry{}, err
		}
		if !reSlug.MatchString(page) {
			e, err := ParsePlainPage(body)
			return "", "", e, err
		}
		e, image, err := ParseAidedd(body)
		return "", image, e, err
	},
	edition: always("2014"),
}

var (
	reSlug       = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	reAideddPath = regexp.MustCompile(`^[A-Za-z0-9._/-]+(?:\?[A-Za-z0-9._=&%+-]*)?$`)
)

func aideddLink(page string) string {
	if reSlug.MatchString(page) {
		return aideddBase + page
	}
	return aideddHome + page
}

// cutEmpty is prefix+s, or "" when s is.
func cutEmpty(prefix, s string) string {
	if s == "" {
		return ""
	}
	return prefix + s
}

var (
	reStatNum  = regexp.MustCompile(`^(\d+)\s*(?:\((.+)\))?`)
	reEntry    = regexp.MustCompile(`(?s)^([^.\n]{1,60})\.(?:\s+(.*))?$`)
	reTypeLine = regexp.MustCompile(`^(\w+)\s+([^,]+?)(?:,\s*(.+))?$`)
)

// aideddSections name the action sections, as action types.
var aideddSections = map[string]string{
	"actions": "action", "bonus actions": "bonus_action", "reactions": "reaction", "legendary actions": "legendary_action",
}

// ParseAidedd normalizes a monster page into an entry and its picture link.
func ParseAidedd(body []byte) (Entry, string, error) {
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return Entry{}, "", fmt.Errorf("aidedd: %w", err)
	}
	h1 := find(doc, func(n *html.Node) bool { return n.Data == "h1" })
	if h1 == nil || h1.Parent == nil {
		return Entry{}, "", fmt.Errorf("aidedd: the page has no monster")
	}
	name := textOf(h1)
	lines := textLines(h1.Parent)
	lines = lines[slices.Index(lines, name)+1:]
	d := MonsterData{Abilities: map[string]int{}}
	section, typed := "", false
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		label, rest := statLabel(line)
		switch {
		case strings.HasPrefix(line, "["):
			// A link to the page in French.
		case !typed:
			typed = true
			if m := reTypeLine.FindStringSubmatch(line); m != nil {
				d.Size, d.Type, d.Alignment = m[1], m[2], m[3]
			}
		case slices.Contains(abilityOrder, strings.ToLower(line)) && i+1 < len(lines):
			d.Abilities[strings.ToLower(line)] = firstInt(lines[i+1])
			i++
		case aideddSections[strings.ToLower(line)] != "":
			section = aideddSections[strings.ToLower(line)]
		case label != "" && section == "" && len(d.Traits) == 0:
			setStat(&d, label, rest)
		default:
			addEntry(&d, section, line)
		}
	}
	d.Initiative = abilityMod(d.Abilities["dex"])
	img := find(doc, func(n *html.Node) bool { return n.Data == "img" && strings.Contains(attr(n, "src"), "/dnd/images/") })
	image := ""
	if img != nil {
		if u, err := url.Parse(aideddBase); err == nil {
			if ref, err := u.Parse(attr(img, "src")); err == nil && ref.Scheme == "https" {
				image = ref.String()
			}
		}
	}
	return Entry{Kind: KindMonster, Slug: Slugify(name), Name: name, Data: d, Text: d.searchText()}, image, nil
}

// aideddLabels are the statblock lines a monster has before its traits.
var aideddLabels = []string{
	"Armor Class", "Hit Points", "Speed", "Saving Throws", "Skills", "Damage Vulnerabilities", "Damage Resistances",
	"Damage Immunities", "Condition Immunities", "Senses", "Languages", "Challenge", "Proficiency Bonus",
}

func statLabel(line string) (string, string) {
	for _, l := range aideddLabels {
		if rest, ok := strings.CutPrefix(line, l+" "); ok {
			return l, strings.TrimSpace(rest)
		}
	}
	return "", ""
}

func setStat(d *MonsterData, label, v string) {
	switch label {
	case "Armor Class", "Hit Points":
		m := reStatNum.FindStringSubmatch(v)
		if m == nil {
			return
		}
		n, _ := strconv.Atoi(m[1])
		if label == "Armor Class" {
			d.AC, d.ACDetail = n, m[2]
		} else {
			d.HP, d.HitDice = n, m[2]
		}
	case "Speed":
		d.Speed = parseSpeed(v)
	case "Saving Throws":
		d.Saves = parseBonuses(v, true)
	case "Skills":
		d.Skills = parseBonuses(v, false)
	case "Damage Resistances":
		d.Resistance = v
	case "Damage Immunities", "Condition Immunities":
		d.Immunity = joinNonEmpty("; ", d.Immunity, v)
	case "Damage Vulnerabilities":
		d.Desc = joinText(d.Desc, "Vulnerabilities: "+v)
	case "Senses":
		d.Senses = v
	case "Languages":
		d.Languages = v
	case "Challenge":
		d.CR, _, _ = strings.Cut(v, " ")
	}
}

// addEntry adds "Name. Text" as a trait or an action of the section, or
// continues the last one.
func addEntry(d *MonsterData, section, line string) {
	m := reEntry.FindStringSubmatch(line)
	switch {
	case m != nil && section == "":
		d.Traits = append(d.Traits, NameDesc{Name: m[1], Desc: m[2]})
	case m != nil:
		d.Actions = append(d.Actions, MonsterAction{Name: m[1], Desc: m[2], Type: section})
	case section != "" && len(d.Actions) > 0 && d.Actions[len(d.Actions)-1].Type == section:
		a := &d.Actions[len(d.Actions)-1]
		a.Desc = strings.TrimSpace(a.Desc + "\n" + line)
	case section == "" && len(d.Traits) > 0:
		t := &d.Traits[len(d.Traits)-1]
		t.Desc = strings.TrimSpace(t.Desc + "\n" + line)
	default:
		d.Desc = joinText(d.Desc, line)
	}
}

// abilityMod is a score's modifier (0 for no score); score/2 floors where
// (score-10)/2 would round odd scores below 10 up.
func abilityMod(score int) int {
	if score == 0 {
		return 0
	}
	return score/2 - 5
}

// textLines is n's text, one line per <br> and block element; a paragraph
// is one line, its <br>s kept as newlines.
func textLines(n *html.Node) []string {
	const keptBreak = "\x00"
	var b strings.Builder
	var walk func(*html.Node, bool)
	walk = func(n *html.Node, inP bool) {
		switch {
		case n.Type == html.TextNode:
			b.WriteString(strings.ReplaceAll(n.Data, "\n", " "))
		case n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style"):
			return
		case n.Type == html.ElementNode && n.Data == "br" && inP:
			b.WriteString(keptBreak)
		case n.Type == html.ElementNode && n.Data == "br":
			b.WriteString("\n")
		}
		block := n.Type == html.ElementNode && slices.Contains([]string{"div", "p", "h1", "h2", "h3", "h4", "h5", "h6", "li", "tr", "ul", "table"}, n.Data)
		if block {
			b.WriteString("\n")
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c, inP || n.Data == "p")
		}
		if block {
			b.WriteString("\n")
		}
	}
	walk(n, false)
	var out []string
	for _, l := range strings.Split(b.String(), "\n") {
		var parts []string
		for _, p := range strings.Split(l, keptBreak) {
			if p = strings.Join(strings.Fields(p), " "); p != "" {
				parts = append(parts, p)
			}
		}
		if len(parts) > 0 {
			out = append(out, strings.Join(parts, "\n"))
		}
	}
	return out
}
