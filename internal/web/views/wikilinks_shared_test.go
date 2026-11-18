package views

import (
	"regexp"
	"strings"
	"testing"
)

// oldJSToken is fields.js's TOKEN before it shared the Go pattern.
var oldJSToken = regexp.MustCompile(`\[\[(note|entity|image):(\d+)\|([^\]|]*)(?:\|(left|right|inline|break))?\]\]`)

// TestWikiLinkPatternMatchesOldToken: the shared pattern reads every token the
// old browser pattern did, the same way, and adds the compendium kind.
func TestWikiLinkPatternMatchesOldToken(t *testing.T) {
	shared := regexp.MustCompile(WikiLinkPattern)
	rows := []struct {
		name, text string
		compendium bool // only the shared pattern matches
		none       bool // neither matches
	}{
		{name: "note", text: "[[note:1|a]]"},
		{name: "entity", text: "[[entity:22|Bob]]"},
		{name: "image", text: "[[image:3|pic]]"},
		{name: "left", text: "[[image:3|pic|left]]"},
		{name: "right", text: "[[image:3|pic|right]]"},
		{name: "inline", text: "[[image:3|pic|inline]]"},
		{name: "break", text: "[[image:3|pic|break]]"},
		{name: "spaces", text: "x [[note:4|two  words here]] y"},
		{name: "empty label", text: "[[note:5|]]"},
		{name: "two", text: "[[note:1|a]] and [[entity:2|b|right]]"},
		{name: "compendium", text: "[[compendium:spell/fireball|Fireball]]", compendium: true},
		{name: "compendium float", text: "[[compendium:spell/fireball|Fireball|inline]]", compendium: true},
		{name: "bad id", text: "[[note:x|a]]", none: true},
		{name: "no label", text: "[[note:1]]", none: true},
		{name: "bad float", text: "[[note:1|a|up]]", none: true},
		{name: "bad entry kind", text: "[[compendium:Spell/x|a]]", none: true},
		// a "[" in the label is allowed; both read the same span
		{name: "nested", text: "[[note:1|[a]]]"},
		{name: "nested open", text: "[[[note:1|a]]"},
	}
	group := func(re *regexp.Regexp, m []string, name string) string { return m[re.SubexpIndex(name)] }
	for _, r := range rows {
		t.Run(r.name, func(t *testing.T) {
			old := oldJSToken.FindAllStringSubmatch(r.text, -1)
			got := shared.FindAllStringSubmatch(r.text, -1)
			if r.none && (len(old) > 0 || len(got) > 0) {
				t.Fatalf("matched malformed %q: old %v shared %v", r.text, old, got)
			}
			if r.compendium {
				if len(old) > 0 || len(got) != 1 || group(shared, got[0], "compendium") != "compendium" {
					t.Fatalf("compendium %q: old %v shared %v", r.text, old, got)
				}
				return
			}
			if len(got) != len(old) {
				t.Fatalf("%q: old %d matches, shared %d", r.text, len(old), len(got))
			}
			for i, o := range old {
				g := got[i]
				if g[0] != o[0] ||
					group(shared, g, "kind") != o[1] || group(shared, g, "id") != o[2] ||
					group(shared, g, "label") != o[3] || group(shared, g, "float") != o[4] {
					t.Errorf("%q: old %q, shared %q", r.text, o, g)
				}
			}
		})
	}
}

// TestWikiLinkPatternIsJSSafe: the pattern is also compiled by the browser, so
// it may use only syntax both engines read the same.
func TestWikiLinkPatternIsJSSafe(t *testing.T) {
	for _, bad := range []string{"(?P<", "(?i", "(?s", "(?m", "(?U", `\A`, `\z`, `\Q`, `\E`, "[[:", `\p{`, `\P{`, `\pL`} {
		if strings.Contains(WikiLinkPattern, bad) {
			t.Errorf("pattern uses %q, which JS lacks or reads differently", bad)
		}
	}
}
