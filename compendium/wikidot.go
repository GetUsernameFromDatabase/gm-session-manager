package compendium

import (
	"cmp"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"

	"gm-session-manager/db"
)

// The GM imports single pages of dnd5e.wikidot.com (2014 rules) by URL.
// Each page is its own source, so re-importing it updates it in place and
// it never collides with an SRD entry of the same slug. The wiki's own
// license does not cover the book text it copies, so none is claimed.

const (
	wikidotHost    = "dnd5e.wikidot.com"
	wikidotMaxBody = 2 << 20
	// WikidotSourcePrefix + page is a wikidot page's source_key.
	WikidotSourcePrefix = "wikidot:"
)

var rePage = regexp.MustCompile(`^[a-z0-9-]+(?::[a-z0-9-]+)?$`)

// wikidotClasses are the classes whose pages "<class>:<name>" are subclasses.
var wikidotClasses = []string{
	"artificer", "barbarian", "bard", "blood-hunter", "cleric", "druid", "fighter", "monk",
	"paladin", "ranger", "rogue", "sorcerer", "warlock", "wizard",
}

// WikidotBase is the address every imported page name is relative to.
const WikidotBase = "https://" + wikidotHost + "/"

// WikidotPage checks a page name ("lineage:giff") or page URL and returns
// the page name. Only pages of dnd5e.wikidot.com are accepted: the server
// fetches them.
func WikidotPage(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if !strings.Contains(raw, "://") {
		raw = WikidotBase + strings.TrimPrefix(raw, "/")
	}
	u, err := url.Parse(raw)
	if err != nil || (u.Scheme != "https" && u.Scheme != "http") || !strings.EqualFold(u.Hostname(), wikidotHost) || u.Port() != "" {
		return "", fmt.Errorf("give a page of https://%s", wikidotHost)
	}
	page := strings.ToLower(strings.Trim(u.Path, "/"))
	if !rePage.MatchString(page) {
		return "", fmt.Errorf("%q is not a wikidot page name", page)
	}
	if _, err := wikidotKind(page); err != nil {
		return "", err
	}
	return page, nil
}

func wikidotKind(page string) (Kind, error) {
	prefix, _, sub := strings.Cut(page, ":")
	switch {
	case !sub && slices.Contains(wikidotClasses, prefix):
		return KindClass, nil
	case prefix == "lineage" || prefix == "race":
		return KindSpecies, nil
	case prefix == "background":
		return KindBackground, nil
	case prefix == "feat":
		return KindFeat, nil
	case sub && slices.Contains(wikidotClasses, prefix):
		return KindSubclass, nil
	}
	return "", errors.New("only class, subclass, lineage, background and feat pages can be imported")
}

var wikidotClient = &http.Client{
	Timeout: 15 * time.Second,
	CheckRedirect: func(req *http.Request, _ []*http.Request) error {
		if !strings.EqualFold(req.URL.Hostname(), wikidotHost) {
			return errors.New("redirect off " + wikidotHost)
		}
		return nil
	},
}

func fetchWikidot(ctx context.Context, page string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://"+wikidotHost+"/"+page, nil)
	if err != nil {
		return nil, err
	}
	res, err := wikidotClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch %s: %w", page, err)
	}
	defer func() { _ = res.Body.Close() }()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch %s: %s", page, res.Status)
	}
	body, err := io.ReadAll(io.LimitReader(res.Body, wikidotMaxBody+1))
	if err != nil {
		return nil, fmt.Errorf("fetch %s: %w", page, err)
	}
	if len(body) > wikidotMaxBody {
		return nil, fmt.Errorf("fetch %s: page too large", page)
	}
	return body, nil
}

// FetchWikidot fetches and parses one page (a name or URL) without saving
// it, for the GM to review: it returns the page name, its book and entry.
func FetchWikidot(ctx context.Context, raw string) (page, book string, e Entry, err error) {
	if page, err = WikidotPage(raw); err != nil {
		return "", "", Entry{}, err
	}
	body, err := fetchWikidot(ctx, page)
	if err != nil {
		return "", "", Entry{}, err
	}
	book, e, err = ParseWikidot(page, body)
	return page, book, e, err
}

// SaveWikidot stores a reviewed page entry as the page's own source and
// returns that source's key. The slug comes from the page name, so saving
// the same page again replaces its entry.
func SaveWikidot(ctx context.Context, store *db.Store, page, book string, e Entry) (string, error) {
	page, err := WikidotPage(page)
	if err != nil {
		return "", err
	}
	if kind, _ := wikidotKind(page); e.Kind != kind {
		return "", fmt.Errorf("%s is a %s page, not a %s", page, kind, e.Kind)
	}
	_, sub, _ := strings.Cut(page, ":")
	e.Slug = Slugify(cmp.Or(sub, page))
	src := wikidotSource(page, book)
	body, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(body)
	_, err = storeSource(ctx, store, src, WikidotBase+page, hex.EncodeToString(sum[:]), []Entry{e})
	return src.Key, err
}

// WikidotSourceKey is the source key of a page.
func WikidotSourceKey(page string) string { return WikidotSourcePrefix + page }

func wikidotSource(page, book string) Source {
	return Source{
		Key: WikidotSourceKey(page), Edition: "2014", License: "Unverified (third-party wiki copy)",
		Attribution: strings.TrimSpace(strings.TrimSpace(book) + " via " + WikidotBase + page),
	}
}

// block is one piece of page content: a heading, a paragraph or list item
// (Label is its leading bold or italic "Name.", Text the rest), or a table.
type block struct {
	Heading int
	Label   string
	Text    string
	Item    bool
	Rows    [][]string
	Links   []string
}

// markdown is a paragraph block as Markdown, its label in bold.
func (b block) markdown() string {
	if b.Label != "" {
		return "**" + b.Label + ".** " + b.Text
	}
	return b.Text
}

// ParseWikidot normalizes a fetched page into an entry and returns the book
// its "Source:" line names.
func ParseWikidot(page string, body []byte) (string, Entry, error) {
	kind, err := wikidotKind(page)
	if err != nil {
		return "", Entry{}, err
	}
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return "", Entry{}, fmt.Errorf("wikidot %s: %w", page, err)
	}
	title, content := pageTitle(doc), byID(doc, "page-content")
	switch {
	case content == nil:
		return "", Entry{}, fmt.Errorf("wikidot %s: the page has no #page-content", page)
	case title == "":
		return "", Entry{}, fmt.Errorf("wikidot %s: the page has no title", page)
	}
	var blocks []block
	collect(content, &blocks)
	book, blocks := sourceLine(blocks)

	prefix, sub, _ := strings.Cut(page, ":")
	e := Entry{Kind: kind, Name: title, Slug: Slugify(cmp.Or(sub, page))}
	switch kind {
	case KindClass:
		d := wikidotClass(page, blocks)
		e.Data, e.Text = d, joinText(withoutTable(d.Desc), FeaturesText(d.Features))
	case KindSubclass:
		if _, after, ok := strings.Cut(title, ":"); ok {
			e.Name = strings.TrimSpace(after) // "Monk: Way of Mercy"
		}
		d := ClassData{Caster: "none", SubclassOf: prefix}
		d.Desc, d.Features = featuresUnder(blocks, nil, "")
		e.Data, e.Text = d, joinText(d.Desc, FeaturesText(d.Features))
	default:
		d := wikidotText(blocks)
		e.Data, e.Text = d, joinText(d.Desc, partsText(d.Parts))
	}
	return book, e, nil
}

// pageTitle is the page's title: #page-title on stock wikidot themes, the
// .page-title element on dnd5e.wikidot.com, else the document title.
func pageTitle(doc *html.Node) string {
	if t := textOf(byID(doc, "page-title")); t != "" {
		return t
	}
	if t := textOf(find(doc, func(n *html.Node) bool { return slices.Contains(strings.Fields(attr(n, "class")), "page-title") })); t != "" {
		return t
	}
	t := textOf(find(doc, func(n *html.Node) bool { return n.Data == "title" }))
	t, _, _ = strings.Cut(t, " - DND 5th Edition")
	return strings.TrimSpace(t)
}

func attr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

// find is the first element in document order that ok accepts.
func find(n *html.Node, ok func(*html.Node) bool) *html.Node {
	if n.Type == html.ElementNode && ok(n) {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if f := find(c, ok); f != nil {
			return f
		}
	}
	return nil
}

func byID(n *html.Node, id string) *html.Node {
	return find(n, func(n *html.Node) bool { return attr(n, "id") == id })
}

// textOf is the text of n with whitespace collapsed.
func textOf(n *html.Node) string {
	if n == nil {
		return ""
	}
	var b strings.Builder
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.TextNode {
			b.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return strings.Join(strings.Fields(b.String()), " ")
}

func links(n *html.Node) []string {
	var out []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				out = append(out, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		out = append(out, links(c)...)
	}
	return out
}

var emphasis = []string{"strong", "em", "b", "i"}

// leadingLabel is the bold or italic "Name." or "Name:" a paragraph opens with.
func leadingLabel(n *html.Node) string { return labelAt(n.FirstChild) }

// labelAt is the bold or italic label at c or the first non-blank node after it.
func labelAt(c *html.Node) string {
	for ; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode && strings.TrimSpace(c.Data) == "" {
			continue
		}
		if t := textOf(c); c.Type == html.ElementNode && slices.Contains(emphasis, c.Data) && (strings.HasSuffix(t, ".") || strings.HasSuffix(t, ":")) {
			return t
		}
		return ""
	}
	return ""
}

// splitAtLabels splits n at each <br> that opens a labeled line: wikidot
// puts a background's proficiencies and equipment in one paragraph.
func splitAtLabels(n *html.Node) []*html.Node {
	var kids []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		kids = append(kids, c)
	}
	out, cur := []*html.Node{n}, n
	for _, c := range kids {
		switch {
		case c.Type == html.ElementNode && c.Data == "br" && labelAt(c.NextSibling) != "":
			n.RemoveChild(c)
			cur = &html.Node{Type: html.ElementNode, Data: n.Data}
			out = append(out, cur)
		case cur != n:
			n.RemoveChild(c)
			cur.AppendChild(c)
		}
	}
	return out
}

// collect flattens content into blocks in document order.
func collect(n *html.Node, out *[]block) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}
		switch c.Data {
		case "h1", "h2", "h3", "h4", "h5", "h6":
			*out = append(*out, block{Heading: int(c.Data[1] - '0'), Text: textOf(c)})
		case "p", "li":
			if c.Data == "li" && hasChild(c, "ul", "ol") {
				collect(c, out)
				continue
			}
			for _, p := range splitAtLabels(c) {
				text, label := textOf(p), leadingLabel(p)
				if text == "" {
					continue
				}
				rest := strings.TrimSpace(strings.TrimPrefix(text, label))
				if label = strings.TrimRight(label, ".:"); rest == "" {
					label, rest = "", text
				}
				*out = append(*out, block{Label: label, Text: rest, Item: c.Data == "li", Links: links(p)})
			}
		case "table":
			b := block{Links: links(c)}
			eachElement(c, "tr", func(tr *html.Node) {
				var row []string
				for td := tr.FirstChild; td != nil; td = td.NextSibling {
					if td.Type == html.ElementNode && (td.Data == "td" || td.Data == "th") {
						row = append(row, textOf(td))
					}
				}
				b.Rows = append(b.Rows, row)
			})
			*out = append(*out, b)
		case "script", "style":
		default:
			collect(c, out)
		}
	}
}

func hasChild(n *html.Node, tags ...string) bool {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && slices.Contains(tags, c.Data) {
			return true
		}
	}
	return false
}

func eachElement(n *html.Node, tag string, f func(*html.Node)) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == tag {
			f(c)
		} else {
			eachElement(c, tag, f)
		}
	}
}

// sourceLine takes the "Source: Book" line out of blocks. A page with
// versions from several books (tabs) keeps only the first.
func sourceLine(blocks []block) (string, []block) {
	isSource := func(b block) bool { return b.Label == "Source" || strings.HasPrefix(b.Text, "Source:") }
	i := slices.IndexFunc(blocks, isSource)
	if i < 0 {
		return "", blocks
	}
	book := strings.TrimSpace(strings.TrimPrefix(blocks[i].Text, "Source:"))
	rest := slices.Delete(slices.Clone(blocks), i, i+1)
	if j := slices.IndexFunc(rest[i:], isSource); j >= 0 {
		rest = rest[:i+j]
	}
	return book, rest
}

// textTypes map 2014 trait and background labels to typed parts.
var textTypes = map[string]string{
	"size": partSize, "speed": partSpeed, "languages": partLanguages,
	"skill proficiencies": partSkills, "tool proficiencies": partTools,
}

// wikidotText reads a species, background or feat: labeled paragraphs are
// parts, prose under a heading is a part named by it ("Feature: Wanderer"),
// and the rest is the description.
func wikidotText(blocks []block) TextData {
	var d TextData
	var desc []string
	heading := ""
	for _, b := range blocks {
		switch {
		case b.Heading > 0:
			heading = strings.TrimSpace(strings.TrimPrefix(b.Text, "Feature:"))
			if strings.HasSuffix(heading, "Traits") {
				heading = ""
			}
		case b.Rows != nil:
		case b.Label == "Prerequisite":
			d.Prerequisite = b.Text
		case b.Label != "":
			name := strings.TrimSpace(strings.TrimPrefix(b.Label, "Feature:"))
			d.Parts = append(d.Parts, NameDesc{Name: name, Desc: b.Text, Type: textTypes[strings.ToLower(name)]})
		case heading != "":
			if n := len(d.Parts); n > 0 && d.Parts[n-1].Name == heading {
				d.Parts[n-1].Desc += "\n\n" + b.Text
			} else {
				d.Parts = append(d.Parts, NameDesc{Name: heading, Desc: b.Text})
			}
		case b.Item:
			desc = append(desc, "- "+b.Text)
		default:
			desc = append(desc, b.Text)
		}
	}
	d.Desc = strings.Join(desc, "\n\n")
	return d
}

var (
	reGainedAt = regexp.MustCompile(`(?i)\b(\d+)(?:st|nd|rd|th) level`)
	reHitDie   = regexp.MustCompile(`(?i)d(\d+)\b`)
	reChoose   = regexp.MustCompile(`(?i)^choose (\w+)(?: skills)? from (.+)$`)
	reOrdinal  = regexp.MustCompile(`^(\d+)(?:st|nd|rd|th)$`)
)

var countWords = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4"}

// featuresUnder reads features from headings: the prose before the first
// feature is the description, and each heading starts a feature whose level
// is levels[name] or the first "Nth level" in its text. levels, when set,
// limits features to its names (a class's Features column); deeper
// headings are parts of a feature. A feature linking to "<subclassLink>…"
// pages is where the subclass is chosen.
func featuresUnder(blocks []block, levels map[string][]int, subclassLink string) (string, []ClassFeature) {
	var desc []string
	var out []ClassFeature
	top := 0
	for _, b := range blocks {
		name := strings.ToLower(b.Text)
		if b.Heading > 0 && (levels == nil || levels[name] != nil) && (top == 0 || b.Heading <= top) {
			top = b.Heading
			out = append(out, ClassFeature{Name: b.Text, Levels: levels[name]})
			continue
		}
		if len(out) == 0 {
			if b.Heading == 0 && b.Rows == nil {
				desc = append(desc, b.markdown())
			}
			continue
		}
		f := &out[len(out)-1]
		switch {
		case b.Heading > 0:
			f.Desc += "**" + b.Text + ".** "
		case b.Rows != nil:
			f.Desc += tableText(b.Rows) + "\n\n"
		default:
			f.Desc += b.markdown() + "\n\n"
		}
		if subclassLink != "" && slices.ContainsFunc(b.Links, func(l string) bool { return strings.Contains(l, "/"+subclassLink) }) {
			f.Subclass = true
		}
	}
	for i := range out {
		out[i].Desc = strings.TrimSpace(out[i].Desc)
		if len(out[i].Levels) == 0 {
			if m := reGainedAt.FindStringSubmatch(out[i].Desc); m != nil {
				n, _ := strconv.Atoi(m[1])
				out[i].Levels = []int{n}
			}
		}
	}
	return strings.Join(desc, "\n\n"), out
}

func tableText(rows [][]string) string {
	lines := make([]string, len(rows))
	for i, r := range rows {
		lines[i] = "|" + strings.Join(r, "|") + "|"
		if i == 0 {
			lines[i] += "\n|" + strings.Repeat("---|", len(r))
		}
	}
	return strings.Join(lines, "\n")
}

// wikidotClass reads a class page: the labeled Hit Points and Proficiencies
// lines, the class table and the features.
func wikidotClass(page string, blocks []block) ClassData {
	d := ClassData{Caster: "none", Table: map[string]map[int]string{}}
	traits := map[string]string{}
	levels := map[string][]int{}
	table := false
	for _, b := range blocks {
		switch {
		case b.Label == "Hit Dice":
			if m := reHitDie.FindStringSubmatch(b.Text); m != nil {
				d.HitDie, _ = strconv.Atoi(m[1])
			}
		case b.Label == "Saving Throws":
			for _, s := range SplitList(b.Text) {
				for _, a := range []string{"str", "dex", "con", "int", "wis", "cha"} {
					if strings.HasPrefix(strings.ToLower(s), a) {
						d.Saves = append(d.Saves, a)
					}
				}
			}
		case b.Label == "Armor":
			traits[classTraits[1]] = b.Text
		case b.Label == "Weapons":
			traits[classTraits[2]] = b.Text
		case b.Label == "Tools":
			traits[classTraits[3]] = b.Text
		case b.Label == "Skills":
			traits[classTraits[0]] = b.Text
			if m := reChoose.FindStringSubmatch(b.Text); m != nil {
				traits[classTraits[0]] = "Choose " + cmp.Or(countWords[strings.ToLower(m[1])], m[1]) + ": " + m[2]
			}
		case b.Rows != nil && !table:
			table = classTable(&d, levels, b.Rows)
		}
	}
	desc, features := featuresUnder(blocks, levels, page+":")
	d.Features = features
	d.Desc = withCoreTraits(desc, traits)
	return d
}

// classTable reads the class table: the levels of each feature in the
// Features column, the other columns, and the caster type from the spell
// slot columns. It reports whether rows were the class table.
func classTable(d *ClassData, levels map[string][]int, rows [][]string) bool {
	head := slices.IndexFunc(rows, func(r []string) bool { return slices.Contains(r, "Level") && slices.Contains(r, "Features") })
	if head < 0 {
		return false
	}
	cols := rows[head]
	maxSlot := 0
	for _, r := range rows[head+1:] {
		m := reOrdinal.FindStringSubmatch(cellAt(r, 0))
		if m == nil {
			continue
		}
		level, _ := strconv.Atoi(m[1])
		for i, col := range cols {
			v := cellAt(r, i)
			empty := v == "" || v == "-" || v == "—"
			switch {
			case col == "Level" || col == "Proficiency Bonus":
			case col == "Features":
				for _, f := range strings.Split(v, ",") {
					if f = strings.ToLower(strings.TrimSpace(f)); f != "" && f != "-" {
						levels[f] = append(levels[f], level)
					}
				}
			case reOrdinal.MatchString(col):
				if !empty {
					n, _ := strconv.Atoi(reOrdinal.FindStringSubmatch(col)[1])
					maxSlot = max(maxSlot, n)
				}
			case col == "Slot Level":
				d.Caster = "pact"
			case !empty:
				if d.Table[col] == nil {
					d.Table[col] = map[int]string{}
				}
				d.Table[col][level] = v
			}
		}
	}
	switch {
	case d.Caster == "pact":
	case maxSlot == 9:
		d.Caster = "full"
	case maxSlot == 5:
		d.Caster = "half"
	}
	return true
}

func cellAt(r []string, i int) string {
	if i < len(r) {
		return strings.TrimSpace(r[i])
	}
	return ""
}
