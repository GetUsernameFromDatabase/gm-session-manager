package main

import (
	"os"
	"slices"
	"strings"
	"testing"
)

// TestVendorSumOneVersionPerPackage: the committed vendor.sum names each
// package at one version, so a -update never ships two copies of one
// (mixed ProseMirror copies break its instanceof checks). The sums only
// guard bytes, not which versions the crawl picked.
func TestVendorSumOneVersionPerPackage(t *testing.T) {
	raw, err := os.ReadFile("vendor.sum")
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]string{}
	for line := range strings.Lines(string(raw)) {
		f := strings.Fields(line)
		if len(f) != 2 {
			continue
		}
		name, version, _ := splitSpec("/npm/" + f[1])
		if version == "" {
			continue // classic scripts and alias modules carry no version
		}
		if prev, ok := seen[name]; ok && prev != version {
			t.Errorf("%s is vendored at %s and %s", name, prev, version)
		}
		seen[name] = version
	}
	if len(seen) == 0 {
		t.Fatal("vendor.sum names no versioned package")
	}
}

// TestVersionLess: the crawl keeps the highest version of each package, so
// a text comparison or a range sorting high would pick an older copy than
// some package needs.
func TestVersionLess(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		want bool
	}{
		{name: "numeric, not text", a: "1.9.0", b: "1.10.0", want: true},
		{name: "newer is not less", a: "1.10.0", b: "1.9.0", want: false},
		{name: "equal is not less", a: "2.0.4", b: "2.0.4", want: false},
		{name: "range sorts below a version", a: "%5E2.0.0", b: "0.0.1", want: true},
		{name: "version is not below a range", a: "0.0.1", b: "%5E2.0.0", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := versionLess(tt.a, tt.b); got != tt.want {
				t.Errorf("versionLess(%q, %q) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

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
