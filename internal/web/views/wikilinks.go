package views

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/internal/enum"
)

// wikiLinkPattern is a [[kind:id|label]] link in plain text (NPC fields),
// with an optional |left, |right, |inline or |break layout for images (default left); notes keep theirs as
// editor nodes. fields.js writes them.
// The compendium kind (cross-references in imported 5etools text) names its
// target by "<entry kind>/<slug>" instead of an id; ParseWiki finds it, but
// WikiLinks does not list it: it has no row id (the server resolves it, see
// web.resolveCompendiumLinks).
var wikiLinkPattern = regexp.MustCompile(`\[\[(?:(note|entity|image):(\d+)|(compendium):([a-z]+)/([^|\]/]+))\|([^\]|]*)(?:\|(left|right|inline|break))?\]\]`)

// LinkCompendium is the WikiLink.Kind of a compendium cross-reference.
const LinkCompendium = "compendium"

// WikiLink is one [[link]] in a text.
type WikiLink struct {
	Kind  string // note, entity or image (note_links.target_type), or compendium
	ID    int64
	Label string
	Float string // left, right, inline, break or "" (left)
	// EntryKind and Slug name a compendium link's target.
	EntryKind, Slug string
}

// Key is the link's target as the [[ picker names it, e.g. "image:5".
func (l WikiLink) Key() string {
	if l.Kind == LinkCompendium {
		return fmt.Sprintf("%s:%s/%s", l.Kind, l.EntryKind, l.Slug)
	}
	return fmt.Sprintf("%s:%d", l.Kind, l.ID)
}

// URL is the target's page ("" for a compendium link: its entry id is
// looked up at render time).
func (l WikiLink) URL() string {
	if l.Kind == LinkCompendium {
		return ""
	}
	return map[string]string{string(enum.LinkNote): "/notes/", string(enum.LinkEntity): "/entities/", string(enum.LinkImage): "/images/"}[l.Kind] + strconv.FormatInt(l.ID, 10)
}

// WikiPart is a run of plain text, or one link.
type WikiPart struct {
	Text string
	Link *WikiLink
}

// ParseWiki splits s into text and links.
func ParseWiki(s string) []WikiPart {
	var parts []WikiPart
	last := 0
	for _, m := range wikiLinkPattern.FindAllStringSubmatchIndex(s, -1) {
		if m[0] > last {
			parts = append(parts, WikiPart{Text: s[last:m[0]]})
		}
		// group g spans s[m[2g]:m[2g+1]], or is -1 when it did not take part
		group := func(g int) string {
			if m[2*g] < 0 {
				return ""
			}
			return s[m[2*g]:m[2*g+1]]
		}
		link := &WikiLink{Label: strings.TrimSpace(group(6)), Float: group(7)}
		if m[6] >= 0 {
			link.Kind, link.EntryKind, link.Slug = LinkCompendium, group(4), group(5)
		} else {
			var err error
			link.Kind = group(1)
			if link.ID, err = strconv.ParseInt(group(2), 10, 64); err != nil { // too many digits: keep it as text
				link = nil
			}
		}
		if link == nil {
			parts = append(parts, WikiPart{Text: s[m[0]:m[1]]})
		} else {
			parts = append(parts, WikiPart{Link: link})
		}
		last = m[1]
	}
	if last < len(s) {
		parts = append(parts, WikiPart{Text: s[last:]})
	}
	return parts
}

// CompendiumRefs lists the slugs s links to, by entry kind.
func CompendiumRefs(s string) map[string][]string {
	out := map[string][]string{}
	for _, p := range ParseWiki(s) {
		if l := p.Link; l != nil && l.Kind == LinkCompendium && !slices.Contains(out[l.EntryKind], l.Slug) {
			out[l.EntryKind] = append(out[l.EntryKind], l.Slug)
		}
	}
	return out
}

// ReplaceCompendiumLinks turns each compendium link in s into a Markdown
// link to its entry under base (ids is keyed like WikiLink.Key), or into its
// bare label when the entry is missing or hidden from the viewer.
func ReplaceCompendiumLinks(s string, ids map[string]int64, base string) string {
	var b strings.Builder
	for _, p := range ParseWiki(s) {
		switch l := p.Link; {
		case l == nil:
			b.WriteString(p.Text)
		case l.Kind != LinkCompendium:
			b.WriteString(l.raw())
		case ids[l.Key()] != 0:
			fmt.Fprintf(&b, "[%s](%s/entries/%d)", l.Label, base, ids[l.Key()])
		default:
			b.WriteString(l.Label)
		}
	}
	return b.String()
}

// raw is the link as written.
func (l WikiLink) raw() string {
	s := fmt.Sprintf("[[%s|%s", l.Key(), l.Label)
	if l.Float != "" {
		s += "|" + l.Float
	}
	return s + "]]"
}

// WikiLinks lists the links in s, each target once; compendium links have no
// id and are left out.
func WikiLinks(s string) []WikiLink {
	var out []WikiLink
	seen := map[string]bool{}
	for _, p := range ParseWiki(s) {
		if p.Link != nil && p.Link.Kind != LinkCompendium && !seen[p.Link.Key()] {
			seen[p.Link.Key()] = true
			out = append(out, *p.Link)
		}
	}
	return out
}
