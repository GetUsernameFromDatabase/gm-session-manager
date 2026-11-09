package compendium

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveSourceDir(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "compendium")
	for _, d := range []string{"book", "keep"} {
		if err := os.MkdirAll(filepath.Join(dir, d), 0o750); err != nil {
			t.Fatal(err)
		}
	}
	for _, key := range []string{"book", "../keep", "..", "", "a/b", `a\b`, "web-page"} {
		if err := RemoveSourceDir(dir, key); err != nil {
			t.Fatalf("%q: %v", key, err)
		}
	}
	if _, err := os.Stat(filepath.Join(dir, "book")); !os.IsNotExist(err) {
		t.Error("book folder not removed")
	}
	if _, err := os.Stat(filepath.Join(dir, "keep")); err != nil {
		t.Error("sibling removed by an escaping key")
	}
}
