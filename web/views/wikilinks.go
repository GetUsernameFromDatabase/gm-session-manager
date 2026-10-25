package views

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// wikiLinkPattern is a [[kind:id|label]] link in plain text (NPC fields),
// with an optional |left, |right, |inline or |break layout for images (default left); notes keep theirs as
// editor nodes. fields.js writes them.
var wikiLinkPattern = regexp.MustCompile(`\[\[(note|entity|image):(\d+)\|([^\]|]*)(?:\|(left|right|inline|break))?\]\]`)

// WikiLink is one [[link]] in a text.
type WikiLink struct {
	Kind  string // note, entity or image (note_links.target_type)
	ID    int64
	Label string
	Float string // left, right, inline, break or "" (left)
}

// Key is the link's target as the [[ picker names it, e.g. "image:5".
func (l WikiLink) Key() string { return fmt.Sprintf("%s:%d", l.Kind, l.ID) }

// URL is the target's page.
func (l WikiLink) URL() string {
	return map[string]string{"note": "/notes/", "entity": "/entities/", "image": "/images/"}[l.Kind] + strconv.FormatInt(l.ID, 10)
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
		id, err := strconv.ParseInt(s[m[4]:m[5]], 10, 64)
		if err != nil { // too many digits: keep it as text
			parts = append(parts, WikiPart{Text: s[m[0]:m[1]]})
		} else {
			link := &WikiLink{Kind: s[m[2]:m[3]], ID: id, Label: strings.TrimSpace(s[m[6]:m[7]])}
			if m[8] >= 0 {
				link.Float = s[m[8]:m[9]]
			}
			parts = append(parts, WikiPart{Link: link})
		}
		last = m[1]
	}
	if last < len(s) {
		parts = append(parts, WikiPart{Text: s[last:]})
	}
	return parts
}

// WikiLinks lists the links in s, each target once.
func WikiLinks(s string) []WikiLink {
	var out []WikiLink
	seen := map[string]bool{}
	for _, p := range ParseWiki(s) {
		if p.Link != nil && !seen[p.Link.Key()] {
			seen[p.Link.Key()] = true
			out = append(out, *p.Link)
		}
	}
	return out
}
