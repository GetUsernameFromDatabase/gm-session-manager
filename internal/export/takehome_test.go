package export

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"gm-session-manager/internal/enum"
)

func sampleTakeHome(file string) TakeHome {
	return TakeHome{
		Manifest:  TakeHomeManifest{Format: TakeHomeFormat, AppVersion: "v1", Token: "tok", ServerID: "srv", CharacterID: 7, CreatedAt: time.Unix(1700000000, 0).UTC()},
		Character: CharacterFile{ID: 7, Name: "Aria", System: "dnd5e", Sheet: json.RawMessage(`{}`)},
		Looks:     []TakeHomeLook{{ImageID: 1, Name: "Hood", File: file, IsCurrent: true}},
		Notes: []TakeHomeNote{{ID: 3, Kind: enum.NoteGeneric, Title: "Diary", Blocks: []TakeHomeBlock{{
			ID: 4, Visibility: enum.VisOwner, Doc: json.RawMessage(`{"type":"doc"}`),
			Files: []TakeHomeBlockFile{{AttachmentID: 5, Name: "map.png", File: file}},
		}}}},
	}
}

func writeZip(t *testing.T, th TakeHome, dir string) []byte {
	t.Helper()
	var buf bytes.Buffer
	if err := WriteTakeHome(&buf, th, dir); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestTakeHomeRoundTrip(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "abc"), []byte("pic"), 0o600); err != nil {
		t.Fatal(err)
	}
	body := writeZip(t, sampleTakeHome(ZipFile("abc")), dir)
	f, err := ReadTakeHome(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		t.Fatal(err)
	}
	if f.Manifest.Token != "tok" || f.Character.Name != "Aria" || len(f.Looks) != 1 || len(f.Notes[0].Blocks[0].Files) != 1 {
		t.Errorf("round trip lost data: %+v", f.TakeHome)
	}
	got, err := f.OpenFile(f.Looks[0].File)
	if err != nil || string(got) != "pic" {
		t.Errorf("OpenFile = %q, %v", got, err)
	}
}

// A file missing on disk is dropped and listed, not fatal; a look with no
// link goes with it.
func TestTakeHomeMissingFileOnDisk(t *testing.T) {
	body := writeZip(t, sampleTakeHome(ZipFile("gone")), t.TempDir())
	f, err := ReadTakeHome(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		t.Fatal(err)
	}
	if len(f.Looks) != 0 || len(f.Notes[0].Blocks[0].Files) != 0 {
		t.Errorf("kept lost files: %+v", f.TakeHome)
	}
}

func rewriteZip(t *testing.T, body []byte, edit func(name string, data []byte) []byte) []byte {
	t.Helper()
	zr, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	zw := zip.NewWriter(&out)
	for _, zf := range zr.File {
		rc, _ := zf.Open()
		var b bytes.Buffer
		_, _ = b.ReadFrom(rc)
		_ = rc.Close()
		if data := edit(zf.Name, b.Bytes()); data != nil {
			if err := writeFile(zw, zf.Name, data); err != nil {
				t.Fatal(err)
			}
		}
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	return out.Bytes()
}

func TestReadTakeHomeRefusals(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "abc"), []byte("pic"), 0o600); err != nil {
		t.Fatal(err)
	}
	good := writeZip(t, sampleTakeHome(ZipFile("abc")), dir)
	tests := []struct {
		name string
		edit func(name string, data []byte) []byte
		want string
	}{
		{"newer format", func(n string, d []byte) []byte {
			if n == manifestFile {
				return []byte(`{"format": 2, "new_field": 1}`)
			}
			return d
		}, "format 2"},
		{"unknown key", func(n string, d []byte) []byte {
			if n == levelUpsFile {
				return []byte(`[{"id":1,"surprise":true}]`)
			}
			return d
		}, "unknown field"},
		{"missing part", func(n string, d []byte) []byte {
			if n == notesFile {
				return nil
			}
			return d
		}, "lacks notes.json"},
		{"missing referenced file", func(n string, d []byte) []byte {
			if n == ZipFile("abc") {
				return nil
			}
			return d
		}, "lacks"},
		{"oversize part", func(n string, d []byte) []byte {
			if n == looksFile {
				return append([]byte("["), bytes.Repeat([]byte(" "), maxJSONBytes+1)...)
			}
			return d
		}, "too large"},
		{"bad block visibility", func(n string, d []byte) []byte {
			if n == notesFile {
				return []byte(strings.Replace(string(d), `"owner"`, `"everyone"`, 1))
			}
			return d
		}, "not valid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := rewriteZip(t, good, tt.edit)
			_, err := ReadTakeHome(bytes.NewReader(body), int64(len(body)))
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Errorf("err = %v, want it to contain %q", err, tt.want)
			}
		})
	}
	if _, err := ReadTakeHome(strings.NewReader("not a zip"), 9); err == nil {
		t.Error("a non-zip was accepted")
	}
}
