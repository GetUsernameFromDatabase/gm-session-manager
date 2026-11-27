package views

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestHomeSourceMarks(t *testing.T) {
	spells := []Option{{Slug: "a", Name: "Alpha", Source: "homebook"}, {Slug: "b", Name: "Beta", Source: "tablebook"}}
	render := func(ctx context.Context) string {
		var buf bytes.Buffer
		if err := spellBoxes("spells", -1, spells, url.Values{}).Render(ctx, &buf); err != nil {
			t.Fatal(err)
		}
		return buf.String()
	}
	home := WithHomeSources(t.Context(), []string{"homebook"})
	if got := strings.Count(render(home), "Added at home"); got != 1 {
		t.Errorf("under a home set: %d marks, want 1 (only the home source's spell)", got)
	}
	if strings.Contains(render(t.Context()), "Added at home") {
		t.Error("without a home set (the table) nothing is marked")
	}

	labels := func(ctx context.Context) []string {
		var out []string
		for _, it := range optionItems(ctx, "none", spells, func(o Option) string { return o.Slug }) {
			out = append(out, fmt.Sprint(it.Label))
		}
		return out
	}
	if got := strings.Join(labels(home), "|"); got != "none|Alpha (added at home)|Beta" {
		t.Errorf("select labels %q", got)
	}
	if got := strings.Join(labels(t.Context()), "|"); got != "none|Alpha|Beta" {
		t.Errorf("select labels at the table %q", got)
	}
}

func TestSourceOrigin(t *testing.T) {
	home := WithHomeSources(t.Context(), []string{"homebook"})
	tests := []struct {
		name               string
		ctx                context.Context
		key                string
		wantHome, wantFrom bool
	}{
		{"added at home", home, "homebook", true, false},
		{"from the table", home, "tablebook", false, true},
		{"homebrew is neither", home, "homebrew", false, false},
		{"no set (the table)", t.Context(), "tablebook", false, false},
	}
	for _, tc := range tests {
		if got := homeSource(tc.ctx, tc.key); got != tc.wantHome {
			t.Errorf("%s: homeSource = %v", tc.name, got)
		}
		if got := fromTable(tc.ctx, tc.key); got != tc.wantFrom {
			t.Errorf("%s: fromTable = %v", tc.name, got)
		}
	}
}
