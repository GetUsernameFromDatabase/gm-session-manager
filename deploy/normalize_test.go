// Package deploy holds the tests of the server scripts: sql-normalize.pl, which update.sh's
// migration gate relies on, and common.sh's names shared with the Go app. Every Docker build runs
// them, so a failure blocks the deploy.
package deploy

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func normalize(t *testing.T, file string) string {
	t.Helper()
	in, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.CommandContext(t.Context(), "perl", "sql-normalize.pl")
	cmd.Stdin = bytes.NewReader(in)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("%s: %v", file, err)
	}
	return string(out)
}

// Each case is a pair testdata/<dir>/<name>.old.sql and .new.sql: add a case by adding a pair.
func TestNormalize(t *testing.T) {
	if _, err := exec.LookPath("perl"); err != nil {
		t.Skip("perl not installed")
	}
	for dir, wantEqual := range map[string]bool{"same": true, "differ": false} {
		olds, err := filepath.Glob(filepath.Join("testdata", dir, "*.old.sql"))
		if err != nil || len(olds) == 0 {
			t.Fatalf("no fixtures in testdata/%s: %v", dir, err)
		}
		for _, old := range olds {
			name := dir + "/" + strings.TrimSuffix(filepath.Base(old), ".old.sql")
			t.Run(name, func(t *testing.T) {
				a := normalize(t, old)
				b := normalize(t, strings.TrimSuffix(old, ".old.sql")+".new.sql")
				if (a == b) != wantEqual {
					t.Errorf("equal=%v, want %v\nold: %q\nnew: %q", a == b, wantEqual, a, b)
				}
			})
		}
	}
}
