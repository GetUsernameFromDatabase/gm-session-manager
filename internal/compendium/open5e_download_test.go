package compendium

import (
	"context"
	"encoding/json"
	"errors"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// serveOpen5e serves doc's listing of names and its raw files; a file
// missing from files answers 404.
func serveOpen5e(t *testing.T, doc string, names []string, files map[string]string) {
	t.Helper()
	type item struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	list := []item{{Name: "README.md", Type: "file"}, {Name: "sub", Type: "dir"}}
	for _, n := range names {
		list = append(list, item{Name: n, Type: "file"})
	}
	listing, err := json.Marshal(list)
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/listing/"+doc, func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write(listing) })
	mux.HandleFunc("/srd/"+doc+"/", func(w http.ResponseWriter, r *http.Request) {
		body, ok := files[r.URL.Path[len("/srd/"+doc+"/"):]]
		if !ok {
			http.NotFound(w, r)
			return
		}
		_, _ = w.Write([]byte(body))
	})
	srv := httptest.NewServer(mux)
	restore := UseTestServer(srv.URL)
	t.Cleanup(func() { restore(); srv.Close() })
}

// writeSRD puts an old doc folder with one file in dir.
func writeSRD(t *testing.T, dir, doc string) string {
	t.Helper()
	old := filepath.Join(dir, doc, "Old.json")
	if err := os.MkdirAll(filepath.Dir(old), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(old, []byte("{}"), 0o600); err != nil {
		t.Fatal(err)
	}
	return old
}

func TestDownloadOpen5e(t *testing.T) {
	for _, doc := range []string{SRD2024Key, SRD2014Key} {
		t.Run(doc, func(t *testing.T) {
			files := map[string]string{"Document.json": `[1]`, "Spell.json": `[2]`}
			serveOpen5e(t, doc, []string{"Document.json", "Spell.json"}, files)
			dir := t.TempDir()
			old := writeSRD(t, dir, doc)
			var heard []string
			err := DownloadOpen5e(context.Background(), dir, doc, func(done, total int, name string) {
				if total != 2 || done != len(heard) {
					t.Errorf("progress %d/%d %s", done, total, name)
				}
				heard = append(heard, name)
			})
			if err != nil {
				t.Fatal(err)
			}
			if len(heard) != 2 {
				t.Errorf("heard %v", heard)
			}
			for name, body := range files {
				got, err := os.ReadFile(filepath.Join(dir, doc, name)) //nolint:gosec // test temp dir
				if err != nil || string(got) != body {
					t.Errorf("%s: %q %v", name, got, err)
				}
			}
			if _, err := os.Stat(old); !errors.Is(err, fs.ErrNotExist) {
				t.Errorf("old file kept: %v", err)
			}
			if entries, _ := os.ReadDir(dir); len(entries) != 1 {
				t.Errorf("left behind: %v", entries)
			}
		})
	}
}

func TestDownloadOpen5eUnknownDoc(t *testing.T) {
	dir := t.TempDir()
	if err := DownloadOpen5e(context.Background(), dir, "../x", nil); err == nil {
		t.Fatal("want error")
	}
	if entries, _ := os.ReadDir(dir); len(entries) != 0 {
		t.Errorf("left behind: %v", entries)
	}
}

func TestDownloadOpen5eBadName(t *testing.T) {
	serveOpen5e(t, SRD2024Key, []string{"Spell.json", "../evil.json"}, map[string]string{"Spell.json": `[]`})
	dir := t.TempDir()
	if err := DownloadOpen5e(context.Background(), dir, SRD2024Key, nil); err == nil {
		t.Fatal("want error")
	}
	if entries, _ := os.ReadDir(dir); len(entries) != 0 {
		t.Errorf("left behind: %v", entries)
	}
}

func TestDownloadOpen5eFailureKeepsOld(t *testing.T) {
	serveOpen5e(t, SRD2024Key, []string{"Document.json", "Spell.json"}, map[string]string{"Document.json": `[]`})
	dir := t.TempDir()
	old := writeSRD(t, dir, SRD2024Key)
	if err := DownloadOpen5e(context.Background(), dir, SRD2024Key, nil); err == nil {
		t.Fatal("want error")
	}
	if _, err := os.Stat(old); err != nil {
		t.Errorf("old folder lost: %v", err)
	}
	if entries, _ := os.ReadDir(dir); len(entries) != 1 {
		t.Errorf("left behind: %v", entries)
	}
}
