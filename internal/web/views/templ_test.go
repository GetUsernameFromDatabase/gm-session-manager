package views

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// templ copies top-level Go into _templ.go, so it lives in <name>_view.go instead.
var goDecl = regexp.MustCompile(`^(type|const|var|func)\b`)

func TestTemplHasNoGoDecls(t *testing.T) {
	files, err := filepath.Glob("*.templ")
	if err != nil || len(files) == 0 {
		t.Fatalf("no .templ files: %v", err)
	}
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		for i, line := range strings.Split(string(b), "\n") {
			if goDecl.MatchString(line) {
				t.Errorf("%s:%d: Go declaration belongs in %s_view.go: %s", f, i+1, strings.TrimSuffix(f, ".templ"), line)
			}
		}
	}
}
