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
		Manifest:  TakeHomeManifest{Format: TakeHomeFormat, Kind: KindTakeHome, AppVersion: "v1", Token: "tok", ServerID: "srv", CharacterID: 7, CreatedAt: time.Unix(1700000000, 0).UTC()},
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
	f, err := ReadTakeHome(bytes.NewReader(body), int64(len(body)), KindTakeHome)
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
	f, err := ReadTakeHome(bytes.NewReader(body), int64(len(body)), KindTakeHome)
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
			_, err := ReadTakeHome(bytes.NewReader(body), int64(len(body)), KindTakeHome)
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Errorf("err = %v, want it to contain %q", err, tt.want)
			}
		})
	}
	if _, err := ReadTakeHome(strings.NewReader("not a zip"), 9, KindTakeHome); err == nil {
		t.Error("a non-zip was accepted")
	}
}

func sampleBringBack() TakeHome {
	th := sampleTakeHome("")
	th.Looks, th.Notes = nil, nil
	th.Manifest.Kind = KindBringBack
	th.Manifest.SavedAt = time.Unix(1700001000, 0).UTC()
	th.Maps = &HomeIDMaps{Notes: map[int64]int64{11: 3}, Blocks: map[int64]int64{12: 4}}
	return th
}

func TestBringBackRoundTrip(t *testing.T) {
	body := writeZip(t, sampleBringBack(), t.TempDir())
	f, err := ReadTakeHome(bytes.NewReader(body), int64(len(body)), KindBringBack)
	if err != nil {
		t.Fatal(err)
	}
	if f.Maps == nil || f.Maps.Notes[11] != 3 || f.Maps.Blocks[12] != 4 || f.Manifest.Kind != KindBringBack || f.Manifest.SavedAt.IsZero() {
		t.Errorf("round trip lost data: %+v %+v", f.Manifest, f.Maps)
	}
}

func TestReadTakeHomeKinds(t *testing.T) {
	dir := t.TempDir()
	takeHome := writeZip(t, sampleTakeHome(""), dir)
	bringBack := writeZip(t, sampleBringBack(), dir)
	noMaps := rewriteZip(t, bringBack, func(n string, d []byte) []byte {
		if n == mapsFile {
			return nil
		}
		return d
	})
	withMaps := writeZip(t, func() TakeHome { th := sampleTakeHome(""); th.Maps = &HomeIDMaps{}; return th }(), dir)
	tests := []struct {
		name string
		body []byte
		want Kind
		err  string
	}{
		{"bring-back where take-home wanted", bringBack, KindTakeHome, "bring it to the table"},
		{"take-home where bring-back wanted", takeHome, KindBringBack, "open it in the player app"},
		{"bring-back without maps", noMaps, KindBringBack, "lacks maps.json"},
		{"take-home with maps", withMaps, KindTakeHome, "must not hold maps.json"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadTakeHome(bytes.NewReader(tt.body), int64(len(tt.body)), tt.want)
			if err == nil || !strings.Contains(err.Error(), tt.err) {
				t.Errorf("err = %v, want it to contain %q", err, tt.err)
			}
		})
	}
}

func sampleNewCharacter() TakeHome {
	th := sampleBringBack()
	th.Maps = nil
	th.Manifest.Kind = KindNewCharacter
	th.Manifest.Token, th.Manifest.CharacterID = "", 0
	return th
}

func sampleFallback() *TakeHomeFallback {
	return &TakeHomeFallback{
		MaxHP: 31, Sheet: json.RawMessage(`{"level":3}`),
		LevelUps: []TakeHomeLevelUp{{ID: 9, MaxHP: 24, Sheet: json.RawMessage(`{}`), CreatedAt: time.Unix(1700000500, 0).UTC()}},
	}
}

func TestBringBackFallbackRoundTrip(t *testing.T) {
	th := sampleBringBack()
	th.Fallback = sampleFallback()
	body := writeZip(t, th, t.TempDir())
	f, err := ReadTakeHome(bytes.NewReader(body), int64(len(body)), KindBringBack)
	if err != nil {
		t.Fatal(err)
	}
	if f.Fallback == nil || f.Fallback.MaxHP != 31 || !strings.Contains(string(f.Fallback.Sheet), `"level": 3`) || len(f.Fallback.LevelUps) != 1 || f.Fallback.LevelUps[0].ID != 9 {
		t.Errorf("fallback lost: %+v", f.Fallback)
	}
	plain := writeZip(t, sampleBringBack(), t.TempDir())
	if f, err := ReadTakeHome(bytes.NewReader(plain), int64(len(plain)), KindBringBack); err != nil || f.Fallback != nil {
		t.Errorf("a bring-back without a fallback: %v, %+v", err, f)
	}
}

func TestNewCharacterRoundTrip(t *testing.T) {
	body := writeZip(t, sampleNewCharacter(), t.TempDir())
	f, err := ReadTakeHome(bytes.NewReader(body), int64(len(body)), KindNewCharacter)
	if err != nil {
		t.Fatal(err)
	}
	if f.Manifest.Kind != KindNewCharacter || f.Maps != nil || f.Fallback != nil || f.Character.Name != "Aria" {
		t.Errorf("round trip lost data: %+v", f.TakeHome)
	}
}

func TestReadTakeHomeNewCharacterKinds(t *testing.T) {
	dir := t.TempDir()
	withFallback := func(th TakeHome) TakeHome { th.Fallback = sampleFallback(); return th }
	newChar := sampleNewCharacter()
	withMaps := newChar
	withMaps.Maps = &HomeIDMaps{}
	withToken := newChar
	withToken.Manifest.Token = "tok"
	withID := newChar
	withID.Manifest.CharacterID = 7
	tests := []struct {
		name string
		th   TakeHome
		want Kind
		err  string
	}{
		{"new-character where take-home wanted", newChar, KindTakeHome, "give it to your GM at the table"},
		{"take-home where new-character wanted", sampleTakeHome(""), KindNewCharacter, "open it in the player app"},
		{"bring-back where new-character wanted", sampleBringBack(), KindNewCharacter, "bring-back file, not a new-character file"},
		{"new-character where bring-back wanted", newChar, KindBringBack, "new-character file, not a bring-back file"},
		{"fallback in take-home", withFallback(sampleTakeHome("")), KindTakeHome, "must not hold fallback.json"},
		{"fallback in new-character", withFallback(newChar), KindNewCharacter, "must not hold fallback.json"},
		{"maps in new-character", withMaps, KindNewCharacter, "must not hold maps.json"},
		{"token in new-character", withToken, KindNewCharacter, "must not name a table character"},
		{"character id in new-character", withID, KindNewCharacter, "must not name a table character"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := writeZip(t, tt.th, dir)
			_, err := ReadTakeHome(bytes.NewReader(body), int64(len(body)), tt.want)
			if err == nil || !strings.Contains(err.Error(), tt.err) {
				t.Errorf("err = %v, want it to contain %q", err, tt.err)
			}
		})
	}
}

func TestDropMissing(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "here"), []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	th := sampleTakeHome(ZipFile("gone"))
	th.Looks = append(th.Looks, TakeHomeLook{ImageID: 2, Name: "Link", URL: "https://example.com/a.png", File: ZipFile("gone2"), FileName: "a"}, TakeHomeLook{ImageID: 3, Name: "Kept", File: ZipFile("here")})
	missing := DropMissing(&th, dir)
	if len(th.Looks) != 2 || th.Looks[0].Name != "Link" || th.Looks[0].File != "" || th.Looks[1].Name != "Kept" {
		t.Errorf("looks = %+v", th.Looks)
	}
	if len(th.Notes[0].Blocks[0].Files) != 0 || len(missing) != 3 {
		t.Errorf("block files %+v, missing %v", th.Notes[0].Blocks[0].Files, missing)
	}
}
