package views

import (
	"strings"
	"testing"
)

func TestRenderMarkdown(t *testing.T) {
	tests := []struct {
		name, in string
		want     []string
		not      []string
	}{
		{"table", "|Score|Modifier|\n|---|---|\n|1|−5|", []string{"<table>", "<th>Score</th>", "<td>−5</td>"}, nil},
		{"bold callout", "> **Round Down**\n> Whenever", []string{"<blockquote>", "<strong>Round Down</strong><br", "Whenever"}, nil},
		{"raw html dropped", "<script>alert(1)</script>", nil, []string{"<script"}},
		{"js link dropped", "[a](javascript:alert(1))", nil, []string{"javascript:"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := renderMarkdown(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			for _, w := range tt.want {
				if !strings.Contains(got, w) {
					t.Errorf("missing %q in %s", w, got)
				}
			}
			for _, n := range tt.not {
				if strings.Contains(got, n) {
					t.Errorf("unexpected %q in %s", n, got)
				}
			}
		})
	}
}
