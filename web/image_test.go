package web

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"gm-session-manager/auth"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/web/views"
)

// doFile posts a multipart form with one file field.
func (e testEnv) doFile(t *testing.T, path, cookie string, fields url.Values, fileField, fileName, content string) *httptest.ResponseRecorder {
	t.Helper()
	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	for k, vs := range fields {
		for _, v := range vs {
			_ = mw.WriteField(k, v)
		}
	}
	if fileField != "" {
		fw, _ := mw.CreateFormFile(fileField, fileName)
		_, _ = fw.Write([]byte(content))
	}
	_ = mw.Close()
	req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, path, &body)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.RemoteAddr = lanAddr
	req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: cookie})
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	return w
}

// TestImageAccess: images are the GM's; a player sees one (name and
// picture only) once a block they can read links it.
func TestImageAccess(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	player := e.token(t, auth.RolePlayer)

	w := e.doFile(t, fmt.Sprintf("/campaigns/%d/images", e.campaignID), gm, url.Values{
		"name": {"Red dragon"}, "url": {"https://example.com/dragon.png"}, "category": {"Monster"}, "description": {"Secret lair notes"},
	}, "", "", "")
	if w.Code != http.StatusSeeOther {
		t.Fatalf("create: %d %s", w.Code, w.Body)
	}
	imgs, _ := e.q().ListImages(ctx, sqlc.ListImagesParams{CampaignID: e.campaignID})
	if len(imgs) != 1 || imgs[0].Category != "Monster" {
		t.Fatalf("images %+v", imgs)
	}
	img := imgs[0]
	page, file := fmt.Sprintf("/images/%d", img.ID), fmt.Sprintf("/images/%d/file", img.ID)

	for _, path := range []string{fmt.Sprintf("/campaigns/%d/images", e.campaignID), views.DisplayURL(e.campaignID), page, file} {
		if w := e.do(t, http.MethodGet, path, player, nil, false); w.Code < 400 {
			t.Errorf("player %s: %d", path, w.Code)
		}
		if w := e.do(t, http.MethodGet, path, gm, nil, false); w.Code >= 400 {
			t.Errorf("gm %s: %d", path, w.Code)
		}
	}

	note, _ := e.q().CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: "generic", Title: "Lore"})
	hidden, _ := e.q().CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: "gm"})
	shared, _ := e.q().CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: "players"})
	snapshot := func(blockID int64) {
		t.Helper()
		snap := map[string]any{
			"version": 1, "doc": json.RawMessage(`{"type":"doc","content":[]}`), "text": "",
			"links": []map[string]any{{"type": "image", "id": img.ID}},
		}
		// The snapshot only lands at a version the block has reached.
		if err := e.q().SetBlockVersion(ctx, sqlc.SetBlockVersionParams{ID: blockID, Version: 1}); err != nil {
			t.Fatal(err)
		}
		if w := e.doJSON(t, fmt.Sprintf("/blocks/%d/snapshot", blockID), gm, snap); w.Code != http.StatusNoContent {
			t.Fatalf("snapshot with an image link: %d %s", w.Code, w.Body)
		}
	}

	snapshot(hidden.ID)
	if w := e.do(t, http.MethodGet, page, player, nil, false); w.Code != http.StatusNotFound {
		t.Errorf("linked from a gm block: %d", w.Code)
	}
	snapshot(shared.ID)
	w = e.do(t, http.MethodGet, page, player, nil, false)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Red dragon") {
		t.Fatalf("linked from a players block: %d", w.Code)
	}
	for _, secret := range []string{"Secret lair notes", "Monster", "example.com"} {
		if strings.Contains(w.Body.String(), secret) {
			t.Errorf("player page leaks %q", secret)
		}
	}
	if w := e.do(t, http.MethodGet, file, player, nil, false); w.Code != http.StatusFound || w.Header().Get("Location") != "https://example.com/dragon.png" {
		t.Errorf("player file: %d %q", w.Code, w.Header().Get("Location"))
	}
}

func TestImageUploadOnlyPictures(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, auth.RoleAdmin)
	path := fmt.Sprintf("/campaigns/%d/images", e.campaignID)
	if w := e.doFile(t, path, gm, url.Values{"name": {"Notes"}}, "file", "x.png", "<html>not a picture</html>"); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("html upload: %d", w.Code)
	}
	if w := e.doFile(t, path, gm, url.Values{"name": {"Nothing"}}, "", "", ""); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("no link or file: %d", w.Code)
	}
	png := "\x89PNG\r\n\x1a\n" + strings.Repeat("\x00", 32)
	if w := e.doFile(t, path, gm, url.Values{"name": {"Map"}}, "file", "map.png", png); w.Code != http.StatusSeeOther {
		t.Fatalf("png upload: %d %s", w.Code, w.Body)
	}
	imgs, _ := e.q().ListImages(t.Context(), sqlc.ListImagesParams{CampaignID: e.campaignID})
	if len(imgs) != 1 || !imgs[0].AttachmentID.Valid {
		t.Fatalf("images %+v", imgs)
	}
	w := e.do(t, http.MethodGet, fmt.Sprintf("/images/%d/file", imgs[0].ID), gm, nil, false)
	if w.Code != http.StatusOK || w.Header().Get("Content-Type") != "image/png" {
		t.Errorf("file: %d %s", w.Code, w.Header().Get("Content-Type"))
	}
}

// TestImageFileCache: a file URL with the image's current version is cached
// for good; without it, or with a stale one, it is not.
func TestImageFileCache(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	create := func(name, field, file, content string, form url.Values) sqlc.GetImageRow {
		t.Helper()
		form.Set("name", name)
		if w := e.doFile(t, fmt.Sprintf("/campaigns/%d/images", e.campaignID), gm, form, field, file, content); w.Code != http.StatusSeeOther {
			t.Fatalf("create %s: %d %s", name, w.Code, w.Body)
		}
		imgs, _ := e.q().ListImages(ctx, sqlc.ListImagesParams{CampaignID: e.campaignID, Q: name})
		img, err := e.q().GetImage(ctx, imgs[0].ID)
		if err != nil {
			t.Fatal(err)
		}
		return img
	}
	linked := create("Linked", "", "", "", url.Values{"url": {"https://example.com/a.png"}})
	uploaded := create("Uploaded", "file", "map.png", "\x89PNG\r\n\x1a\n"+strings.Repeat("\x00", 32), url.Values{})
	oldLink := imageVersion(linked.Url, linked.AttachmentID)
	if w := e.doFile(t, fmt.Sprintf("/images/%d", linked.ID), gm, url.Values{"name": {"Linked"}, "url": {"https://example.com/b.png"}}, "", "", ""); w.Code != http.StatusSeeOther {
		t.Fatalf("relink: %d", w.Code)
	}
	linked, _ = e.q().GetImage(ctx, linked.ID)

	list := e.do(t, http.MethodGet, fmt.Sprintf("/campaigns/%d/images", e.campaignID), gm, nil, false).Body.String()
	for _, img := range []sqlc.GetImageRow{linked, uploaded} {
		if want := imageItem(img).Src(); !strings.Contains(list, want) {
			t.Errorf("list lacks %s", want)
		}
	}

	for _, tc := range []struct {
		name string
		img  sqlc.GetImageRow
		v    string
		want string
	}{
		{"link current", linked, imageVersion(linked.Url, linked.AttachmentID), cacheForever},
		{"link stale", linked, oldLink, "no-store"},
		{"link bare", linked, "", "no-store"},
		{"upload current", uploaded, imageVersion(uploaded.Url, uploaded.AttachmentID), cacheForever},
		{"upload bare", uploaded, "", cacheAttachment},
	} {
		w := e.do(t, http.MethodGet, fmt.Sprintf("/images/%d/file?v=%s", tc.img.ID, tc.v), gm, nil, false)
		if got := w.Header().Get("Cache-Control"); w.Code >= 400 || got != tc.want {
			t.Errorf("%s: %d %q, want %q", tc.name, w.Code, got, tc.want)
		}
	}
}

// TestAssetCache: assets whose content can change at a URL (ours, htmx, the
// vendor aliases) revalidate by ETag; versioned vendor files are cached for good.
func TestAssetCache(t *testing.T) {
	e := newTestEnv(t)
	get := func(path, ifNoneMatch string) *httptest.ResponseRecorder {
		t.Helper()
		req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, path, nil)
		req.RemoteAddr = lanAddr
		if ifNoneMatch != "" {
			req.Header.Set("If-None-Match", ifNoneMatch)
		}
		w := httptest.NewRecorder()
		e.router.ServeHTTP(w, req)
		return w
	}
	for _, path := range []string{"/assets/js/app.js", "/assets/js/live.js", "/assets/css/output.css", "/assets/vendor/htmx-ws.js", "/assets/vendor/tiptap-core.js"} {
		w := get(path, "")
		etag := w.Header().Get("ETag")
		if w.Code != http.StatusOK || etag == "" || w.Header().Get("Cache-Control") != "no-cache" {
			t.Fatalf("%s: %d etag %q cache %q", path, w.Code, etag, w.Header().Get("Cache-Control"))
		}
		if w := get(path, etag); w.Code != http.StatusNotModified {
			t.Errorf("%s revalidated: %d", path, w.Code)
		}
		if w := get(path, `"stale"`); w.Code != http.StatusOK {
			t.Errorf("%s stale tag: %d", path, w.Code)
		}
	}
	if w := get("/assets/vendor/@floating-ui/core@1.8.0/index.js", ""); w.Code != http.StatusOK || !strings.Contains(w.Header().Get("Cache-Control"), "immutable") {
		t.Errorf("vendor: %d %q", w.Code, w.Header().Get("Cache-Control"))
	}
}

// TestImageImport: the GM maps columns, merges two spellings into one
// category and leaves a row out.
func TestImageImport(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	csvFile := "Title,Kind,Image\nMira,Person,https://x.test/mira.png\nOld Tom,person,https://x.test/tom.png\nThe Keep,Area,https://x.test/keep.png\nBroken,Area,not a link\n"
	w := e.doFile(t, fmt.Sprintf("/campaigns/%d/images/import", e.campaignID), gm, nil, "file", "all.csv", csvFile)
	if w.Code != http.StatusSeeOther {
		t.Fatalf("upload: %d %s", w.Code, w.Body)
	}
	review := w.Header().Get("Location")
	w = e.do(t, http.MethodGet, review, gm, nil, false)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "not an http(s) link") {
		t.Fatalf("review: %d", w.Code)
	}
	// Importing can take a while: the form shows a busy indicator.
	if !strings.Contains(w.Body.String(), "data-busy") || !strings.Contains(w.Body.String(), `data-slot="spinner"`) {
		t.Errorf("review has no loading indicator")
	}
	// Links can be checked in the browser: one status cell per importable row.
	if !strings.Contains(w.Body.String(), "data-check-links") || strings.Count(w.Body.String(), "data-link-status") != 3 {
		t.Errorf("review has no link check")
	}
	// Rows can be filtered by status, and the link check is the primary action.
	if !strings.Contains(w.Body.String(), `name="filter_status"`) || !strings.Contains(w.Body.String(), "Untick them") {
		t.Errorf("rows toolbar misses the status filter or the results line")
	}
	form := url.Values{
		views.ImportColName: {"0"}, views.ImportColLink: {"2"}, views.ImportColCategory: {"1"}, views.ImportColDescription: {"-1"},
		// "person" is split from "Person"; "Area" is merged into "Person"'s target.
		views.ImportCatSource: {"Area", "Person", "person"},
		views.ImportCatTarget: {"Folk", "Folk", "NPC"},
		// The broken row is refused even if ticked.
		views.ImportRowsField: {"0-3"},
	}
	// The rows refresh shows each row under its new category.
	w = e.do(t, http.MethodGet, review+"?"+form.Encode(), gm, nil, true)
	body := w.Body.String()
	if !strings.Contains(body, `data-category="NPC"`) || strings.Contains(body, `data-category="Area"`) {
		t.Errorf("rows do not follow the category targets")
	}
	// Merged values are boxed together, and categories sort by name.
	if strings.Count(body, "data-merged") != 1 || !strings.Contains(body, "Merged into Folk") {
		t.Errorf("merged categories not marked")
	}
	if strings.Index(body, `value="Folk"`) > strings.Index(body, `value="NPC"`) {
		t.Errorf("categories not sorted by name")
	}
	if w := e.do(t, http.MethodPost, review, gm, form, false); w.Code != http.StatusSeeOther {
		t.Fatalf("commit: %d %s", w.Code, w.Body)
	}
	imgs, _ := e.q().ListImages(ctx, sqlc.ListImagesParams{CampaignID: e.campaignID})
	got := map[string]string{}
	for _, i := range imgs {
		got[i.Name] = i.Category
	}
	want := map[string]string{"Mira": "Folk", "Old Tom": "NPC", "The Keep": "Folk"}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("imported %v, want %v", got, want)
	}
	w = e.do(t, http.MethodGet, review, gm, nil, false)
	if w.Code != http.StatusNotFound {
		t.Errorf("import kept after commit: %d", w.Code)
	}
	if back := fmt.Sprintf(`href="/campaigns/%d/images/import"`, e.campaignID); !strings.Contains(w.Body.String(), back) {
		t.Errorf("expired import does not link back to the upload page")
	}
}

// TestImageImportByLink: the CSV can come from a link instead of a file.
func TestImageImportByLink(t *testing.T) {
	csvServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/images.csv" {
			http.NotFound(w, r)
			return
		}
		_, _ = w.Write([]byte("name,link\nMira,https://x.test/mira.png\n"))
	}))
	defer csvServer.Close()
	e := newTestEnv(t)
	gm := e.token(t, auth.RoleAdmin)
	path := fmt.Sprintf("/campaigns/%d/images/import", e.campaignID)
	w := e.doFile(t, path, gm, url.Values{"url": {csvServer.URL + "/images.csv"}}, "", "", "")
	if w.Code != http.StatusSeeOther {
		t.Fatalf("link import: %d %s", w.Code, w.Body)
	}
	if w := e.do(t, http.MethodGet, w.Header().Get("Location"), gm, nil, false); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Mira") {
		t.Errorf("review: %d", w.Code)
	}
	for _, link := range []string{csvServer.URL + "/missing.csv", "ftp://x.test/a.csv", ""} {
		if w := e.doFile(t, path, gm, url.Values{"url": {link}}, "", "", ""); w.Code != http.StatusUnprocessableEntity {
			t.Errorf("link %q: %d", link, w.Code)
		}
	}
}

func TestGuessMapping(t *testing.T) {
	const none = noImportColumn
	tests := []struct {
		name string
		csv  string
		want views.ImportMapping
	}{
		{"header names", "\xef\xbb\xbfname,link,category\na,https://x/a,Person\n", views.ImportMapping{Name: 0, Link: 1, Category: 2, Description: none}},
		{
			"links found by their values, columns never reused", "A,B,C,D,E\nBarrel,Person,https://x/b.jpg,,Brilinta\nKeep,Area,https://x/k.jpg,,Brilinta\n",
			views.ImportMapping{Name: 0, Link: 2, Category: none, Description: none},
		},
		{
			"partial header matches", "Picture Link,Title,Kind,Short description\nhttps://x/a,Mira,Person,Innkeeper\n",
			views.ImportMapping{Name: 1, Link: 0, Category: 2, Description: 3},
		},
		{"no links still picks a link column", "name,other\nMira,x\n", views.ImportMapping{Name: 0, Link: 1, Category: none, Description: none}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header, rows, err := parseCSV(strings.NewReader(tt.csv))
			if err != nil {
				t.Fatal(err)
			}
			if got := guessMapping(header, rows); got != tt.want {
				t.Errorf("mapping %+v, want %+v", got, tt.want)
			}
		})
	}
}

// TestImageImportKeepOnly: a master file is narrowed to one campaign
// spelled two ways, and ticks keep pointing at the right CSV rows.
func TestImageImportKeepOnly(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	csvFile := "Name,Link,Campaign\n" +
		"Mira,https://x.test/0.png,Brilinta\n" +
		"Tom,https://x.test/1.png,Other\n" +
		"Keep,https://x.test/2.png,brilinta\n" +
		"Inn,https://x.test/3.png,Brilinta\n"
	w := e.doFile(t, fmt.Sprintf("/campaigns/%d/images/import", e.campaignID), gm, nil, "file", "all.csv", csvFile)
	review := w.Header().Get("Location")
	if w := e.do(t, http.MethodGet, review, gm, nil, false); !strings.Contains(w.Body.String(), "Campaign") || !strings.Contains(w.Body.String(), "Brilinta") {
		t.Fatalf("overview misses the Campaign column")
	}
	form := url.Values{
		views.ImportKeepCol: {"2"},
		views.ImportKeepVal: {"=Brilinta", "=brilinta"},
		// Inn (row 3) is unticked: a filtered-out row must not shift it.
		views.ImportRowsField: {"!3"},
	}
	body := e.do(t, http.MethodGet, review+"?"+form.Encode(), gm, nil, true).Body.String()
	if strings.Contains(body, `data-name="Tom"`) || !strings.Contains(body, "Keeping 3 of 4 rows") {
		t.Errorf("filter not applied")
	}
	if w := e.do(t, http.MethodPost, review, gm, form, false); w.Code != http.StatusSeeOther {
		t.Fatalf("commit: %d %s", w.Code, w.Body)
	}
	imgs, _ := e.q().ListImages(ctx, sqlc.ListImagesParams{CampaignID: e.campaignID})
	var names []string
	for _, i := range imgs {
		names = append(names, i.Name)
	}
	if fmt.Sprint(names) != "[Keep Mira]" {
		t.Errorf("imported %v, want [Keep Mira]", names)
	}
}

func TestKeepValues(t *testing.T) {
	rows := [][]string{{"Brilinta"}, {"other"}, {"brilinta"}, {""}, {"Brilinta"}}
	vals, tooMany := keepValues(rows, 0, map[string]bool{"brilinta": true})
	got := fmt.Sprint(vals)
	if want := "[{ 1 false} {Brilinta 2 false} {brilinta 1 true} {other 1 false}]"; tooMany || got != want {
		t.Errorf("values %s, want %s", got, want)
	}
	var many [][]string
	for i := range maxKeepValues + 1 {
		many = append(many, []string{fmt.Sprint(i)})
	}
	if vals, tooMany := keepValues(many, 0, nil); !tooMany || vals != nil {
		t.Errorf("no cap above %d values", maxKeepValues)
	}
}

func TestParseRowSet(t *testing.T) {
	tests := []struct {
		in   string
		want string // ticked rows of 6, as x and .
	}{
		{"", "......"},
		{"0-2,4", "xxx.x."},
		{"1,3,5", ".x.x.x"},
		{"!", "xxxxxx"},
		{"!1-2", "x..xxx"},
		{"4-99", "....xx"},
		{"3-1,x,-2,2-y", "......"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			var got strings.Builder
			for _, on := range parseRowSet(tt.in, 6) {
				got.WriteString(map[bool]string{true: "x", false: "."}[on])
			}
			if got.String() != tt.want {
				t.Errorf("got %s, want %s", got.String(), tt.want)
			}
		})
	}
}

func TestImportCategories(t *testing.T) {
	rows := []views.ImportRow{{Category: "Person"}, {Category: "person"}, {Category: "PERSON"}, {Category: "PERSON"}, {Category: "Area"}, {}}
	tests := []struct {
		name     string
		existing []sqlc.ImageCategory
		want     string // default target of the person spellings
	}{
		{"case variants start merged, normalized to lowercase", nil, "person"},
		{"an existing category wins", []sqlc.ImageCategory{{Name: "Person"}}, "Person"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cats := importCategories(rows, tt.existing)
			// Every value stays its own entry, so any of them can be split off.
			if len(cats) != 4 {
				t.Fatalf("categories %+v", cats)
			}
			for _, c := range cats {
				want := tt.want
				if c.Source == "Area" {
					want = "area"
				}
				if c.Target != want {
					t.Errorf("%q → %q, want %q", c.Source, c.Target, want)
				}
			}
		})
	}
}

// TestImageNPCLooks makes an NPC from an image, adds and switches looks, and
// checks players reach only revealed looks of revealed NPCs.
func TestImageNPCLooks(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	player := e.token(t, auth.RolePlayer)
	newImage := func(campaignID int64, name, desc string) sqlc.Image {
		t.Helper()
		img, err := e.q().CreateImage(ctx, sqlc.CreateImageParams{
			CampaignID: campaignID, Name: name, Url: sql.NullString{String: "https://example.com/" + name + ".png", Valid: true},
			Description: sql.NullString{String: desc, Valid: desc != ""},
		})
		if err != nil {
			t.Fatal(err)
		}
		return img
	}
	travel, gown := newImage(e.campaignID, "travel", "Wears a grey cloak"), newImage(e.campaignID, "gown", "")

	w := e.do(t, http.MethodPost, fmt.Sprintf("/images/%d/npc", travel.ID), gm, url.Values{}, true)
	if w.Code != http.StatusNoContent {
		t.Fatalf("create NPC: %d %s", w.Code, w.Body)
	}
	npcs, _ := e.q().ListEntitiesByCampaign(ctx, e.campaignID)
	if len(npcs) != 1 || npcs[0].Name != "travel" || w.Header().Get("HX-Redirect") != fmt.Sprintf("/entities/%d", npcs[0].ID) {
		t.Fatalf("NPCs %+v, redirect %q", npcs, w.Header().Get("HX-Redirect"))
	}
	npc := npcs[0]
	if !strings.Contains(npc.FieldsJson, `"description":{"value":"Wears a grey cloak","is_revealed":false}`) {
		t.Errorf("fields %s", npc.FieldsJson)
	}
	looks := func() string {
		t.Helper()
		rows, err := e.q().ListEntityImages(ctx, npc.ID)
		if err != nil {
			t.Fatal(err)
		}
		var out []string
		for _, r := range rows {
			out = append(out, fmt.Sprintf("%s:%v:%v", r.Name, r.IsCurrent, r.IsRevealed))
		}
		return strings.Join(out, " ")
	}
	if got := looks(); got != "travel:true:false" {
		t.Fatalf("looks after create: %s", got)
	}
	base := fmt.Sprintf("/entities/%d/images", npc.ID)
	if w := e.do(t, http.MethodPost, base, gm, url.Values{"image": {fmt.Sprint(gown.ID)}}, true); w.Code != http.StatusOK {
		t.Fatalf("add look: %d %s", w.Code, w.Body)
	}
	if w := e.do(t, http.MethodPost, fmt.Sprintf("%s/%d/current", base, gown.ID), gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("make current: %d %s", w.Code, w.Body)
	}
	if got := looks(); got != "gown:true:false travel:false:false" {
		t.Fatalf("looks after switch: %s", got)
	}
	other, _ := e.q().CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Elsewhere", System: "dnd5e"})
	if w := e.do(t, http.MethodPost, base, gm, url.Values{"image": {fmt.Sprint(newImage(other.ID, "stray", "").ID)}}, true); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("another campaign's image: %d", w.Code)
	}
	if w := e.do(t, http.MethodGet, fmt.Sprintf("/images/%d", gown.ID), gm, nil, false); !strings.Contains(w.Body.String(), fmt.Sprintf("/entities/%d", npc.ID)) {
		t.Error("image page does not mention the NPC")
	}

	if err := e.q().UpdateEntity(ctx, sqlc.UpdateEntityParams{ID: npc.ID, Name: npc.Name, IsRevealed: true, FieldsJson: npc.FieldsJson}); err != nil {
		t.Fatal(err)
	}
	page, file := fmt.Sprintf("/entities/%d", npc.ID), fmt.Sprintf("/images/%d/file", gown.ID)
	src := fmt.Sprintf("/images/%d/file?v=", gown.ID)
	if w := e.do(t, http.MethodGet, file, player, nil, false); w.Code != http.StatusNotFound {
		t.Errorf("hidden look file: %d", w.Code)
	}
	if body := e.do(t, http.MethodGet, page, player, nil, false).Body.String(); strings.Contains(body, src) || strings.Contains(body, "grey cloak") {
		t.Error("player NPC page leaks a hidden look or field")
	}
	if w := e.do(t, http.MethodPost, fmt.Sprintf("%s/%d/reveal", base, gown.ID), gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("reveal: %d %s", w.Code, w.Body)
	}
	if w := e.do(t, http.MethodGet, file, player, nil, false); w.Code != http.StatusFound {
		t.Errorf("revealed look file: %d", w.Code)
	}
	if body := e.do(t, http.MethodGet, page, player, nil, false).Body.String(); !strings.Contains(body, src) {
		t.Error("player NPC page misses the revealed look")
	}
	if w := e.do(t, http.MethodPost, fmt.Sprintf("%s/%d/unlink", base, gown.ID), gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("unlink: %d %s", w.Code, w.Body)
	}
	if w := e.do(t, http.MethodGet, file, player, nil, false); w.Code != http.StatusNotFound {
		t.Errorf("unlinked look file: %d", w.Code)
	}
}

// TestCharacterLooks: a player adds looks to their own character; others
// see only revealed ones, and an image is an NPC's look or a character's.
func TestCharacterLooks(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	alice := e.token(t, auth.RolePlayer)
	other, err := e.q().CreateCharacter(ctx, sqlc.CreateCharacterParams{CampaignID: e.campaignID, Name: "Bob", System: "dnd5e"})
	if err != nil {
		t.Fatal(err)
	}
	bob, _, err := e.mgr.IssuePlayer(ctx, other.ID)
	if err != nil {
		t.Fatal(err)
	}
	add := fmt.Sprintf("/characters/%d/images", e.characterID)
	addLook := func(cookie, name string) *httptest.ResponseRecorder {
		t.Helper()
		return e.doFile(t, add, cookie, url.Values{"name": {name}, "url": {"https://example.com/" + name + ".png"}}, "", "", "")
	}
	for _, name := range []string{"travel", "disguise"} {
		if w := addLook(alice, name); w.Code >= 400 {
			t.Fatalf("add %s: %d %s", name, w.Code, w.Body)
		}
	}
	rows, _ := e.q().ListCharacterImages(ctx, e.characterID)
	var looks []string
	ids := map[string]int64{}
	for _, r := range rows {
		looks = append(looks, fmt.Sprintf("%s:%v:%v", r.Name, r.IsCurrent, r.IsRevealed))
		ids[r.Name] = r.ID
	}
	if got := strings.Join(looks, " "); got != "travel:true:true disguise:false:false" {
		t.Fatalf("looks %s", got)
	}
	for _, cookie := range []string{bob, gm} {
		if w := addLook(cookie, "intruder"); w.Code != http.StatusForbidden {
			t.Errorf("someone else adds a look: %d", w.Code)
		}
		if w := e.do(t, http.MethodPost, fmt.Sprintf("%s/%d/reveal", add, ids["disguise"]), cookie, nil, false); w.Code != http.StatusForbidden {
			t.Errorf("someone else reveals a look: %d", w.Code)
		}
	}
	file := func(name string) string { return fmt.Sprintf("/images/%d/file", ids[name]) }
	for _, tc := range []struct {
		cookie, name string
		want         int
	}{{bob, "travel", http.StatusFound}, {bob, "disguise", http.StatusNotFound}, {alice, "disguise", http.StatusFound}} {
		if w := e.do(t, http.MethodGet, file(tc.name), tc.cookie, nil, false); w.Code != tc.want {
			t.Errorf("%s file: %d, want %d", tc.name, w.Code, tc.want)
		}
	}

	// Bob sees Alice's current look in their fight, while it is revealed.
	enc, _ := e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"})
	for _, ch := range []int64{e.characterID, other.ID} {
		if _, err := e.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: enc.ID, CharacterID: nullID(ch), Name: "pc"}); err != nil {
			t.Fatal(err)
		}
	}
	if err := e.q().UpdateEncounterState(ctx, sqlc.UpdateEncounterStateParams{ID: enc.ID, Status: "active", Round: 1}); err != nil {
		t.Fatal(err)
	}
	combat := fmt.Sprintf("/campaigns/%d/combat", e.campaignID)
	if body := e.do(t, http.MethodGet, combat, bob, nil, false).Body.String(); !strings.Contains(body, file("travel")+"?v=") {
		t.Error("combat misses the revealed portrait")
	}
	if w := e.do(t, http.MethodPost, fmt.Sprintf("%s/%d/current", add, ids["disguise"]), alice, nil, true); w.Code != http.StatusOK {
		t.Fatalf("make current: %d %s", w.Code, w.Body)
	}
	if body := e.do(t, http.MethodGet, combat, bob, nil, false).Body.String(); strings.Contains(body, file("disguise")) || strings.Contains(body, file("travel")) {
		t.Error("combat shows a hidden or former portrait")
	}

	// Exclusive: a character's image cannot become an NPC's look, nor the reverse.
	bind := func(imageID int64, form url.Values) int {
		return e.do(t, http.MethodPost, fmt.Sprintf("/images/%d/npc", imageID), gm, form, true).Code
	}
	if code := bind(ids["travel"], url.Values{}); code != http.StatusUnprocessableEntity {
		t.Errorf("NPC from a character's image: %d", code)
	}
	if npcs, _ := e.q().ListEntitiesByCampaign(ctx, e.campaignID); len(npcs) != 0 {
		t.Errorf("refused NPC was created: %+v", npcs)
	}
	mask, _ := e.q().CreateImage(ctx, sqlc.CreateImageParams{CampaignID: e.campaignID, Name: "mask", Url: sql.NullString{String: "https://example.com/mask.png", Valid: true}})
	if code := bind(mask.ID, url.Values{}); code != http.StatusNoContent {
		t.Fatalf("NPC from mask: %d", code)
	}
	if code := bind(mask.ID, url.Values{"character": {fmt.Sprint(other.ID)}}); code != http.StatusUnprocessableEntity {
		t.Errorf("character from an NPC's look: %d", code)
	}
	free, _ := e.q().CreateImage(ctx, sqlc.CreateImageParams{CampaignID: e.campaignID, Name: "free", Url: sql.NullString{String: "https://example.com/free.png", Valid: true}})
	if code := bind(free.ID, url.Values{"character": {fmt.Sprint(other.ID)}}); code != http.StatusNoContent {
		t.Fatalf("give to Bob: %d", code)
	}
	if rows, _ := e.q().ListCharacterImages(ctx, other.ID); len(rows) != 1 || !rows[0].IsCurrent || !rows[0].IsRevealed {
		t.Errorf("Bob's looks %+v", rows)
	}
}

// TestEntityFormFields: the NPC editor lists only stored fields; unused
// suggestions wait in the add-field picker, and added rows save.
func TestEntityFormFields(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	npc, err := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: e.campaignID, Kind: "npc", Name: "Mira", FieldsJson: `{"role":{"value":"Innkeeper"}}`})
	if err != nil {
		t.Fatal(err)
	}
	page := fmt.Sprintf("/entities/%d", npc.ID)
	body := e.do(t, http.MethodGet, page, gm, nil, false).Body.String()
	for _, want := range []string{`name="key_0" value="role"`, `name="value_0" value="Innkeeper" data-field-value`, `name="count" value="1"`, `data-templ-value="motivation"`, `data-field-key="" data-empty`} {
		if !strings.Contains(body, want) {
			t.Errorf("editor misses %s", want)
		}
	}
	if strings.Contains(body, `name="key_1"`) || strings.Contains(body, `data-templ-value="role"`) {
		t.Error("editor shows unstored fields or offers a stored key")
	}
	form := url.Values{"name": {"Mira"}, "count": {"2"}, "key_0": {"role"}, "value_0": {"Innkeeper"}, "key_1": {"Secret Name"}, "value_1": {"Lyra"}, "revealed_1": {"true"}}
	if w := e.do(t, http.MethodPost, page, gm, form, true); w.Code != http.StatusOK {
		t.Fatalf("save: %d %s", w.Code, w.Body)
	}
	if got, _ := e.q().GetEntity(ctx, npc.ID); !strings.Contains(got.FieldsJson, `"secret name":{"value":"Lyra","is_revealed":true}`) {
		t.Errorf("fields %s", got.FieldsJson)
	}
}

// TestEntityFieldLinks: [[links]] in NPC fields render as in notes, let
// players open images a revealed field links, and show in linked mentions.
func TestEntityFieldLinks(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	player := e.token(t, auth.RolePlayer)
	tavern, _ := e.q().CreateImage(ctx, sqlc.CreateImageParams{CampaignID: e.campaignID, Name: "Tavern", Url: sql.NullString{String: "https://example.com/tavern.png", Valid: true}})
	plans, _ := e.q().CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: "generic", Title: "Secret plans"})
	if _, err := e.q().CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: plans.ID, Visibility: "gm"}); err != nil {
		t.Fatal(err)
	}
	npc, _ := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: e.campaignID, Kind: "npc", Name: "Mira", FieldsJson: "{}"})
	page, file := fmt.Sprintf("/entities/%d", npc.ID), fmt.Sprintf("/images/%d/file", tavern.ID)
	save := func(revealLocation bool) {
		t.Helper()
		form := url.Values{"name": {"Mira"}, "revealed": {"true"}, "count": {"2"},
			"key_0": {"location"}, "value_0": {fmt.Sprintf("Upstairs at [[image:%d|Tavern]]", tavern.ID)},
			"key_1": {"plots"}, "value_1": {fmt.Sprintf("See [[note:%d|Secret plans]]", plans.ID)}, "revealed_1": {"true"}}
		if revealLocation {
			form.Set("revealed_0", "true")
		}
		if w := e.do(t, http.MethodPost, page, gm, form, true); w.Code != http.StatusOK {
			t.Fatalf("save: %d %s", w.Code, w.Body)
		}
	}

	save(true)
	if w := e.do(t, http.MethodGet, file, player, nil, false); w.Code != http.StatusFound {
		t.Errorf("image from a revealed field: %d", w.Code)
	}
	body := e.do(t, http.MethodGet, page, player, nil, false).Body.String()
	if !strings.Contains(body, fmt.Sprintf(`href="/images/%d"`, tavern.ID)) || !strings.Contains(body, fmt.Sprintf(`src="/images/%d/file"`, tavern.ID)) {
		t.Error("player page misses the image link")
	}
	if strings.Contains(body, fmt.Sprintf(`href="/notes/%d"`, plans.ID)) || !strings.Contains(body, "Secret plans") {
		t.Error("a note the player cannot read should show as its label only")
	}
	for _, path := range []string{fmt.Sprintf("/images/%d", tavern.ID), fmt.Sprintf("/notes/%d", plans.ID)} {
		if body := e.do(t, http.MethodGet, path, gm, nil, false).Body.String(); !strings.Contains(body, page) {
			t.Errorf("%s does not mention the NPC", path)
		}
	}

	save(false)
	if w := e.do(t, http.MethodGet, file, player, nil, false); w.Code != http.StatusNotFound {
		t.Errorf("image from a hidden field: %d", w.Code)
	}
}

// TestCombatNPCPortrait: the tracker shows an NPC's current look to the GM
// always, and to players only while both the NPC and the look are revealed.
func TestCombatNPCPortrait(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	player := e.token(t, auth.RolePlayer)
	npc, _ := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: e.campaignID, Kind: "npc", Name: "Mira", FieldsJson: "{}"})
	mask, _ := e.q().CreateImage(ctx, sqlc.CreateImageParams{CampaignID: e.campaignID, Name: "mask", Url: sql.NullString{String: "https://example.com/mask.png", Valid: true}})
	if err := e.q().LinkEntityImage(ctx, sqlc.LinkEntityImageParams{EntityID: npc.ID, ImageID: mask.ID}); err != nil {
		t.Fatal(err)
	}
	enc, _ := e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"})
	for _, cb := range []sqlc.CreateCombatantParams{{EncounterID: enc.ID, EntityID: nullID(npc.ID), Name: "Mira"}, {EncounterID: enc.ID, CharacterID: nullID(e.characterID), Name: "pc"}} {
		if _, err := e.q().CreateCombatant(ctx, cb); err != nil {
			t.Fatal(err)
		}
	}
	if err := e.q().UpdateEncounterState(ctx, sqlc.UpdateEncounterStateParams{ID: enc.ID, Status: "active", Round: 1}); err != nil {
		t.Fatal(err)
	}
	combat, src := fmt.Sprintf("/campaigns/%d/combat?encounter=%d", e.campaignID, enc.ID), fmt.Sprintf("/images/%d/file", mask.ID)
	for _, tt := range []struct {
		name                      string
		npcRevealed, lookRevealed bool
		cookie                    string
		want                      bool
	}{
		{"GM sees a hidden look", false, false, gm, true},
		{"player misses a hidden look", true, false, player, false},
		{"player misses the look of a hidden NPC", false, true, player, false},
		{"player sees a revealed look", true, true, player, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if err := e.q().UpdateEntity(ctx, sqlc.UpdateEntityParams{ID: npc.ID, Name: npc.Name, IsRevealed: tt.npcRevealed, FieldsJson: "{}"}); err != nil {
				t.Fatal(err)
			}
			if looks, _ := e.q().ListEntityImages(ctx, npc.ID); looks[0].IsRevealed != tt.lookRevealed {
				if err := e.q().ToggleEntityImageRevealed(ctx, sqlc.ToggleEntityImageRevealedParams{EntityID: npc.ID, ImageID: mask.ID}); err != nil {
					t.Fatal(err)
				}
			}
			if got := strings.Contains(e.do(t, http.MethodGet, combat, tt.cookie, nil, false).Body.String(), src); got != tt.want {
				t.Errorf("portrait shown = %v, want %v", got, tt.want)
			}
		})
	}
}
