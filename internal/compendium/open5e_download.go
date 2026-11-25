package compendium

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// The Open5e documents DownloadOpen5e fetches; each is also its folder's
// name, so its source key.
const (
	SRD2024Key = "srd-2024"
	SRD2014Key = "srd-2014"
)

// The WotC folder of the Open5e repo: its listing (GitHub contents API) and
// its raw files, each followed by a document key. Tests point both at httptest.
var (
	open5eListBase = "https://api.github.com/repos/open5e/open5e-api/contents/data/v2/wizards-of-the-coast/"
	open5eRawBase  = "https://raw.githubusercontent.com/open5e/open5e-api/staging/data/v2/wizards-of-the-coast/"
)

// UseTestServer is for tests only (it panics elsewhere): it points the
// 5etools mirror at base+"/mirror/" and the Open5e listing and files of a
// document at base+"/listing/<doc>" and base+"/srd/<doc>/"; the returned
// func restores them.
func UseTestServer(base string) (restore func()) {
	if !testing.Testing() {
		panic("compendium.UseTestServer outside a test")
	}
	old := []string{fiveToolsBase, open5eListBase, open5eRawBase}
	fiveToolsBase, open5eListBase, open5eRawBase = base+"/mirror/", base+"/listing/", base+"/srd/"
	return func() { fiveToolsBase, open5eListBase, open5eRawBase = old[0], old[1], old[2] }
}

var open5eFilePattern = regexp.MustCompile(`^[A-Za-z0-9_]+\.json$`)

// open5eFiles lists the .json files of doc's folder. A .json name outside
// open5eFilePattern refuses the run rather than reach a path.
func open5eFiles(ctx context.Context, doc string) ([]string, error) {
	body, err := fetch(ctx, open5eListBase+doc+"?ref=staging", bookMaxBody)
	if err != nil {
		return nil, err
	}
	var list []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	if err := json.Unmarshal(body, &list); err != nil {
		return nil, fmt.Errorf("open5e listing: %w", err)
	}
	var names []string
	for _, f := range list {
		if f.Type != "file" || !strings.HasSuffix(strings.ToLower(f.Name), ".json") {
			continue
		}
		if !open5eFilePattern.MatchString(f.Name) {
			return nil, fmt.Errorf("bad open5e file name %q", f.Name)
		}
		names = append(names, f.Name)
	}
	if len(names) == 0 {
		return nil, errors.New("open5e listing holds no files")
	}
	return names, nil
}

// DownloadOpen5e writes the files of doc (SRD2024Key or SRD2014Key) into
// dir/doc, replacing the old folder only once every file is in. progress, if set, hears of each
// file as it starts.
func DownloadOpen5e(ctx context.Context, dir, doc string, progress func(done, total int, name string)) error {
	if doc != SRD2024Key && doc != SRD2014Key {
		return fmt.Errorf("unknown open5e document %q", doc)
	}
	if err := os.MkdirAll(dir, 0o750); err != nil { //nolint:gosec // dir is the operator's DATA_DIR, never request input
		return err
	}
	names, err := open5eFiles(ctx, doc)
	if err != nil {
		return err
	}
	// Same prefix as a book download, so CleanDownloads clears a leftover.
	tmp, err := os.MkdirTemp(dir, tmpBookPrefix)
	if err != nil {
		return err
	}
	defer func() { _ = os.RemoveAll(tmp) }() //nolint:gosec // a temp dir under DATA_DIR; no-op after the rename
	for i, name := range names {
		if progress != nil {
			progress(i, len(names), name)
		}
		body, err := fetch(ctx, open5eRawBase+doc+"/"+name, bookMaxBody)
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(tmp, name), body, 0o600); err != nil { //nolint:gosec // name matched open5eFilePattern
			return err
		}
	}
	return replaceDir(dir, tmp, filepath.Join(dir, doc))
}
