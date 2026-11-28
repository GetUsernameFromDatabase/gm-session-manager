package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
	"strings"
	"testing"
	"time"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/route"
)

// homeEnv is the player app with the sample take-home imported (the table's
// source "src1" at 150); it returns the player's cookie and the home campaign.
func homeEnv(t *testing.T) (e testEnv, cookie string, campaign int64) {
	t.Helper()
	e = newPlayerEnv(t)
	w := e.importRequest(t, sampleTakeHome(t, nil), "", nil)
	if w.Code != http.StatusNoContent {
		t.Fatalf("import: status %d: %s", w.Code, w.Body)
	}
	cookie = homeCookie(w).Value
	sess, err := e.mgr.Resolve(t.Context(), cookie)
	if err != nil {
		t.Fatal(err)
	}
	ch, err := e.q().GetCharacter(t.Context(), sess.CharacterID.Int64)
	if err != nil {
		t.Fatal(err)
	}
	return e, cookie, ch.CampaignID
}

// homeBook adds a source the way a download at home does (not the take-home's path).
func (e testEnv) homeBook(t *testing.T, key, edition string) {
	t.Helper()
	if err := e.q().UpsertCompendiumSource(t.Context(), sqlc.UpsertCompendiumSourceParams{SourceKey: key, Path: "downloads/" + key, Hash: "h-" + key, Edition: edition, Priority: compendium.DefaultRank(key, edition)}); err != nil {
		t.Fatal(err)
	}
}

func (e testEnv) priorities(t *testing.T) map[string]int64 {
	t.Helper()
	rows, err := e.q().ListCompendiumSources(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	out := map[string]int64{}
	for _, r := range rows {
		out[r.SourceKey] = r.Priority
	}
	return out
}

// homeDo sends a request to the player app as the browser at home (cookie "": a visitor).
func (e testEnv) homeDo(t *testing.T, method, path, cookie string, form url.Values) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequestWithContext(t.Context(), method, path, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if method != http.MethodGet { // an htmx GET of the 5etools picker would fetch the book list
		req.Header.Set("HX-Request", "true")
	}
	req.RemoteAddr = loopbackAddr
	if cookie != "" {
		req.AddCookie(&http.Cookie{Name: auth.PlayerCookieName, Value: cookie})
	}
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	return w
}

// sourceRoutes are the routes that add or manage sources, as method and part
// of the campaign's compendium (the Sources route is under the campaign).
var sourceRoutes = []struct{ method, path string }{
	{http.MethodPost, route.CompendiumImport},
	{http.MethodPost, route.CompendiumFolder},
	{http.MethodGet, route.CompendiumFiveTools},
	{http.MethodPost, route.CompendiumFiveTools},
	{http.MethodPost, route.CompendiumFiveToolsBookDelete},
	{http.MethodPost, route.CompendiumFiveToolsAuto},
	{http.MethodPost, route.CompendiumImportDismiss},
	{http.MethodPost, route.CompendiumOpen5e},
	{http.MethodPost, route.CompendiumSourceDelete},
	{http.MethodPost, route.CompendiumSourceOrder},
	{http.MethodPost, route.CompendiumSourceSort},
	{http.MethodPost, route.CompendiumPagePreview},
	{http.MethodPost, route.CompendiumPage},
}

func sourceRoutePath(campaign int64, part string) string {
	if part == route.CompendiumFiveToolsBookDelete {
		return route.Path(route.CampaignCompendium+part, campaign, "x")
	}
	return route.Path(route.CampaignCompendium+part, campaign)
}

// At the table only the GM adds sources: a table player is refused everywhere.
func TestSourceRoutesGMOnlyAtTable(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, enum.RolePlayer)
	for _, r := range sourceRoutes {
		w := e.do(t, r.method, sourceRoutePath(e.campaignID, r.path), player, url.Values{}, true)
		if w.Code != http.StatusForbidden {
			t.Errorf("%s %s: status %d, want 403", r.method, r.path, w.Code)
		}
	}
	w := e.do(t, http.MethodPost, route.Path(route.CampaignSources, e.campaignID), player, url.Values{"key": {"x"}, "on": {"true"}}, true)
	if w.Code != http.StatusForbidden {
		t.Errorf("campaign sources: status %d, want 403", w.Code)
	}
}

// In the player app the home player gets through; a visitor does not.
func TestSourceRoutesHomePlayer(t *testing.T) {
	e, cookie, camp := homeEnv(t)
	e.homeBook(t, "homebook", "2024")

	// A plain GET of the picker redirects to the page: past the guard, and no download.
	picker := sourceRoutePath(camp, route.CompendiumFiveTools)
	if w := e.homeDo(t, http.MethodGet, picker, cookie, nil); w.Code == http.StatusForbidden || w.Code == http.StatusNotFound {
		t.Errorf("home player GET 5etools: status %d", w.Code)
	}
	for _, r := range sourceRoutes {
		// The campaign's own check or the sources guard: a refusal either way.
		if w := e.homeDo(t, r.method, sourceRoutePath(camp, r.path), "", nil); w.Code < http.StatusBadRequest {
			t.Errorf("visitor %s %s: status %d, want a refusal", r.method, r.path, w.Code)
		}
	}

	w := e.homeDo(t, http.MethodPost, sourceRoutePath(camp, route.CompendiumSourceOrder), cookie, url.Values{"key": e.sourceKeys(t)})
	if w.Code != http.StatusOK {
		t.Fatalf("source order: status %d: %s", w.Code, w.Body)
	}
	w = e.homeDo(t, http.MethodPost, route.Path(route.CampaignSources, camp), cookie, url.Values{"key": {"homebook"}, "on": {"true"}})
	if w.Code != http.StatusOK {
		t.Fatalf("campaign sources: status %d: %s", w.Code, w.Body)
	}
	on, err := e.q().ListCampaignSources(t.Context(), camp)
	if err != nil || !slices.Contains(on, "homebook") {
		t.Errorf("campaign sources %v, %v: homebook not ticked", on, err)
	}
}

// sourceKeys is every source key in priority order: a drag that changes nothing.
func (e testEnv) sourceKeys(t *testing.T) []string {
	t.Helper()
	rows, err := e.q().ListCompendiumSources(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	var keys []string
	for _, r := range rows {
		keys = append(keys, r.SourceKey)
	}
	return keys
}

func TestSinkHomeSources(t *testing.T) {
	e, _, _ := homeEnv(t)
	ctx := t.Context()
	e.homeBook(t, "home-a", "2024") // 100: above the table's 150
	e.homeBook(t, "home-b", "")     // 300
	before := e.priorities(t)
	if err := e.q().SinkHomeSources(ctx); err != nil {
		t.Fatal(err)
	}
	after := e.priorities(t)
	table := after["src1"]
	if table != 150 || after["homebrew"] != before["homebrew"] {
		t.Errorf("the table's and homebrew's priorities changed: %v -> %v", before, after)
	}
	if after["home-a"] <= table || after["home-b"] <= after["home-a"] {
		t.Errorf("home sources %d, %d must rank below the table's %d, keeping their order", after["home-a"], after["home-b"], table)
	}
	if after["home-b"]-after["home-a"] != before["home-b"]-before["home-a"] {
		t.Errorf("the gap between home sources changed: %v -> %v", before, after)
	}
	if err := e.q().SinkHomeSources(ctx); err != nil {
		t.Fatal(err)
	}
	if again := e.priorities(t); !maps.Equal(after, again) {
		t.Errorf("a second run moved sources: %v -> %v", after, again)
	}
}

// A source downloaded at home that the take-home also lists stays the home's.
func TestTakeHomeKeepsHomeSource(t *testing.T) {
	e := newPlayerEnv(t)
	ctx := t.Context()
	e.homeBook(t, "src1", "2024")
	if err := e.upsertEntry(ctx, sqlc.UpsertCompendiumEntryParams{Kind: enum.KindSpell, Slug: "bolt", Name: "Home bolt", SourceKey: "src1", DataJson: "{}", ContentHash: "home"}); err != nil {
		t.Fatal(err)
	}
	w := e.importRequest(t, sampleTakeHome(t, nil), "", nil)
	if w.Code != http.StatusNoContent {
		t.Fatalf("import: status %d: %s", w.Code, w.Body)
	}
	src, err := e.q().GetCompendiumSource(ctx, "src1")
	if err != nil {
		t.Fatal(err)
	}
	if src.Path != "downloads/src1" || src.Hash != "h-src1" || src.Priority != 100 {
		t.Errorf("the home source was rewritten: %+v", src)
	}
	entry, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: enum.KindSpell, Slug: "bolt"})
	if err != nil || entry.Name != "Home bolt" || entry.ContentHash != "home" {
		t.Errorf("the home copy of the entry was rewritten: %+v, %v", entry, err)
	}
	sess, err := e.mgr.Resolve(ctx, homeCookie(w).Value)
	if err != nil {
		t.Fatal(err)
	}
	ch, err := e.q().GetCharacter(ctx, sess.CharacterID.Int64)
	if err != nil {
		t.Fatal(err)
	}
	if on, err := e.q().ListCampaignSources(ctx, ch.CampaignID); err != nil || !slices.Contains(on, "src1") {
		t.Errorf("the new campaign allows %v (%v), want src1", on, err)
	}
}

// The bring-back carries the home-added sources the campaign allows, with
// their character-kind entries, and nothing of the table's or homebrew.
func TestBringBackCarriesHomeSources(t *testing.T) {
	e, _, camp := homeEnv(t)
	ctx := t.Context()
	e.homeBook(t, "homebook", "2024")
	e.homeBook(t, "unticked", "2024")
	for _, p := range []sqlc.UpsertCompendiumEntryParams{
		{Kind: enum.KindSpell, Slug: "hspell", Name: "Home spell", SourceKey: "homebook", DataJson: "{}", ContentHash: "a"},
		{Kind: enum.KindMonster, Slug: "hmon", Name: "Home monster", SourceKey: "homebook", DataJson: "{}", ContentHash: "b"},
		{Kind: enum.KindSpell, Slug: "uspell", Name: "Unticked spell", SourceKey: "unticked", DataJson: "{}", ContentHash: "c"},
	} {
		if err := e.upsertEntry(ctx, p); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := e.q().AllowCampaignSource(ctx, sqlc.AllowCampaignSourceParams{CampaignID: camp, SourceKey: "homebook"}); err != nil {
		t.Fatal(err)
	}
	copies, err := e.q().ListHomeCopies(ctx)
	if err != nil || len(copies) != 1 {
		t.Fatalf("home copies %d, %v", len(copies), err)
	}
	th, _, err := e.app.bringBack(ctx, copies[0].CharacterID)
	if err != nil {
		t.Fatal(err)
	}
	var sources, entries []string
	for _, s := range th.Sources {
		sources = append(sources, s.SourceKey)
	}
	for _, en := range th.Entries {
		entries = append(entries, en.SourceKey+"/"+en.Slug)
	}
	if !slices.Equal(sources, []string{"homebook"}) || !slices.Equal(entries, []string{"homebook/hspell"}) {
		t.Errorf("sources %v, entries %v; want only homebook and its spell", sources, entries)
	}

	// The file reads back as a bring-back, and the table-side check passes.
	var buf bytes.Buffer
	if err := export.WriteTakeHome(&buf, th, t.TempDir()); err != nil {
		t.Fatal(err)
	}
	f, err := export.ReadTakeHome(bytes.NewReader(buf.Bytes()), int64(buf.Len()), export.KindBringBack)
	if err != nil {
		t.Fatal(err)
	}
	if err := checkFileSources(f); err != nil {
		t.Errorf("checkFileSources: %v", err)
	}
}

func TestCheckFileSources(t *testing.T) {
	entry := func(source string) export.TakeHomeEntry {
		return export.TakeHomeEntry{Kind: enum.KindSpell, Slug: "s", Name: "S", SourceKey: source, Data: json.RawMessage(`{}`)}
	}
	file := func(kind export.Kind, sources []string, entries ...export.TakeHomeEntry) *export.TakeHomeFile {
		f := &export.TakeHomeFile{}
		f.Manifest.Kind = kind
		for _, k := range sources {
			f.Sources = append(f.Sources, export.TakeHomeSource{SourceKey: k})
		}
		f.Entries = entries
		return f
	}
	tests := []struct {
		name    string
		file    *export.TakeHomeFile
		wantErr bool
	}{
		{"listed source", file(export.KindBringBack, []string{"a"}, entry("a")), false},
		{"unlisted source", file(export.KindBringBack, []string{"a"}, entry("b")), true},
		{"empty source key", file(export.KindTakeHome, []string{""}), true},
		{"homebrew as a source", file(export.KindTakeHome, []string{"homebrew"}), true},
		{"homebrew entry in a take-home", file(export.KindTakeHome, nil, entry("homebrew")), false},
		{"homebrew entry in a bring-back", file(export.KindBringBack, nil, entry("homebrew")), true},
	}
	for _, tc := range tests {
		err := checkFileSources(tc.file)
		var ue userError
		if (err != nil) != tc.wantErr || (err != nil && !errors.As(err, &ue)) {
			t.Errorf("%s: err = %v, want error %v (a userError)", tc.name, err, tc.wantErr)
		}
	}
}

// The table's sources stay at home; a source added at home can be deleted.
func TestHomeSourceDelete(t *testing.T) {
	e, cookie, camp := homeEnv(t)
	e.homeBook(t, "homebook", "2024")
	del := route.Path(route.CampaignCompendium+route.CompendiumSourceDelete, camp)

	w := e.homeDo(t, http.MethodPost, del, cookie, url.Values{"key": {"src1"}})
	if w.Code != http.StatusConflict && w.Code != http.StatusUnprocessableEntity {
		t.Errorf("deleting the table's source: status %d, want a refusal", w.Code)
	}
	if _, err := e.q().GetCompendiumSource(t.Context(), "src1"); err != nil {
		t.Errorf("the table's source is gone: %v", err)
	}
	w = e.homeDo(t, http.MethodPost, del, cookie, url.Values{"key": {"homebook"}})
	if w.Code != http.StatusOK {
		t.Fatalf("deleting a home source: status %d: %s", w.Code, w.Body)
	}
	if _, err := e.q().GetCompendiumSource(t.Context(), "homebook"); err == nil {
		t.Error("the home source is still there")
	}
}

func TestNetErrMsg(t *testing.T) {
	dns := &url.Error{Op: "Get", URL: "https://example.test/x", Err: &net.OpError{Op: "dial", Net: "tcp", Err: &net.DNSError{Err: "no such host", Name: "example.test"}}}
	if got := netErrMsg("Download failed", dns); !strings.Contains(got, "No internet connection") || !strings.Contains(got, "no such host") {
		t.Errorf("network error: %q", got)
	}
	if got := netErrMsg("x", fmt.Errorf("fetch %s: %w", "x", dns)); !strings.Contains(got, "No internet connection") {
		t.Errorf("wrapped network error: %q", got)
	}
	// A real refused connection, as offline at home: the server is gone.
	srv := httptest.NewServer(http.NotFoundHandler())
	srv.Close()
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, srv.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err == nil {
		_ = res.Body.Close()
		t.Fatal("the closed server answered")
	}
	if got := netErrMsg("Download failed", fmt.Errorf("fetch: %w", err)); !strings.Contains(got, "No internet connection") {
		t.Errorf("refused connection: %q", got)
	}
	// Not the network: a refused redirect is a *url.Error too.
	redirect := &url.Error{Op: "Get", URL: "https://example.test/x", Err: errors.New("redirect off example.test")}
	for _, err := range []error{errors.New("boom"), redirect} {
		if got := netErrMsg("Download failed", err); got != "Download failed: "+err.Error() {
			t.Errorf("%v: %q", err, got)
		}
	}
}

// The replace question names the sources added here that the new copy leaves
// off, but not one the new file lists itself.
func TestHomeReplaceNamesLostSources(t *testing.T) {
	e, cookie, campaign := homeEnv(t)
	ctx := t.Context()
	for _, k := range []string{"homebook", "listedbook"} {
		e.homeBook(t, k, "2014")
		if _, err := e.q().AllowCampaignSource(ctx, sqlc.AllowCampaignSourceParams{CampaignID: campaign, SourceKey: k}); err != nil {
			t.Fatal(err)
		}
	}
	newer := sampleTakeHome(t, func(th *export.TakeHome) {
		th.Manifest.CreatedAt = th.Manifest.CreatedAt.Add(48 * time.Hour)
		th.Sources = append(th.Sources, export.TakeHomeSource{SourceKey: "listedbook"})
	})
	body := e.importRequest(t, newer, cookie, nil).Body.String()
	if !strings.Contains(body, "leaves off the sources added here (homebook)") {
		t.Errorf("lost source not named: %s", body)
	}
}
