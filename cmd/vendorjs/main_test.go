package main

import (
	"slices"
	"testing"
)

func TestChangedPaths(t *testing.T) {
	const a, b = "aaaa", "bbbb"
	tests := []struct {
		name    string
		got     map[string]string
		sumText string
		want    []string
	}{
		{
			name:    "all equal",
			got:     map[string]string{"x.js": a, "y.js": b},
			sumText: a + "  x.js\n" + b + "  y.js\n",
		},
		{
			name:    "sum changed",
			got:     map[string]string{"x.js": a, "y.js": a},
			sumText: a + "  x.js\n" + b + "  y.js\n",
			want:    []string{"y.js"},
		},
		{
			name:    "new in download",
			got:     map[string]string{"x.js": a, "y.js": b},
			sumText: a + "  x.js\n",
			want:    []string{"y.js"},
		},
		{
			name:    "gone from download",
			got:     map[string]string{"x.js": a},
			sumText: a + "  x.js\n" + b + "  y.js\n",
			want:    []string{"y.js"},
		},
		{
			// Five paths, so map order without the sort would almost never pass.
			name:    "several changed, sorted",
			got:     map[string]string{"e.js": a, "a.js": a, "d.js": a, "b.js": a, "c.js": a},
			sumText: "",
			want:    []string{"a.js", "b.js", "c.js", "d.js", "e.js"},
		},
		{
			name:    "CRLF line endings",
			got:     map[string]string{"x.js": a, "y.js": b},
			sumText: a + "  x.js\r\n" + b + "  y.js\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changedPaths(tt.got, tt.sumText); !slices.Equal(got, tt.want) {
				t.Errorf("changedPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
