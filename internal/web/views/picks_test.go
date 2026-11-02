package views

import (
	"strings"
	"testing"
)

// TestPickRowActs checks a pick's button posts its action with no hx-target
// for it to inherit, which would make htmx drop the request.
func TestPickRowActs(t *testing.T) {
	var b strings.Builder
	picks := []BuilderPick{{EntryID: 7, Slug: "plate", Name: "Plate"}}
	if err := BuilderPicks(1, "armor", picks).Render(t.Context(), &b); err != nil {
		t.Fatal(err)
	}
	html := b.String()
	if !strings.Contains(html, `hx-post="/characters/1/cmd"`) {
		t.Errorf("no action in %s", html)
	}
	if strings.Contains(html, "hx-target") {
		t.Errorf("a target the button would inherit in %s", html)
	}
	if !strings.Contains(html, `/compendium/entries/7?preview=1`) {
		t.Errorf("no details loader in %s", html)
	}
}
