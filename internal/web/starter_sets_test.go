package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

const srdMini = "../compendium/testdata/srd-mini"

// serveStarterSets serves a tiny 5etools mirror (one 2024 book with one
// spell, one 2014 book) and the srd-mini folder as both Open5e SRDs. Its
// cleanup waits for e's import to end before restoring the real addresses,
// so a failing test never reaches the network.
func serveStarterSets(t *testing.T, e testEnv) {
	t.Helper()
	files := map[string]string{
		"/mirror/books.json": `{"book":[
			{"name":"Player's Handbook (2024)","id":"XPHB","group":"core","published":"2024-09-17","author":"Wizards RPG Team"},
			{"name":"Player's Handbook","id":"PHB","group":"core","published":"2014-08-19","author":"Wizards RPG Team"},
			{"name":"Broken Book","id":"BRK","group":"core","published":"2014-01-01","author":"Wizards RPG Team"}]}`,
		"/mirror/adventures.json":         `{"adventure":[]}`,
		"/mirror/spells/index.json":       `{"XPHB":"spells-xphb.json","BRK":"spells-brk.json"}`,
		"/mirror/spells/spells-xphb.json": `{"spell":[{"name":"Fireball","source":"XPHB","level":3,"school":"V","entries":["Boom."]}]}`,
		// BRK's download fails: its spell file is not JSON.
		"/mirror/spells/spells-brk.json": `not json`,
	}
	names, err := os.ReadDir(srdMini)
	if err != nil {
		t.Fatal(err)
	}
	var list []string
	for _, n := range names {
		list = append(list, `{"name":"`+n.Name()+`","type":"file"}`)
	}
	for _, doc := range []string{compendium.SRD2024Key, compendium.SRD2014Key} {
		files["/listing/"+doc] = "[" + strings.Join(list, ",") + "]"
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if rest, ok := strings.CutPrefix(r.URL.Path, "/srd/"); ok {
			http.ServeFile(w, r, filepath.Join(srdMini, filepath.Base(rest)))
			return
		}
		body, ok := files[r.URL.Path]
		if !ok {
			http.NotFound(w, r)
			return
		}
		_, _ = w.Write([]byte(body))
	}))
	restore := compendium.UseTestServer(srv.URL)
	t.Cleanup(func() {
		for end := time.Now().Add(20 * time.Second); e.app.fiveTools.get().Running && time.Now().Before(end); {
			time.Sleep(10 * time.Millisecond)
		}
		restore()
		srv.Close()
	})
}

// waitJob waits for the background import to end and returns its summary.
func waitJob(t *testing.T, e testEnv) string {
	t.Helper()
	for end := time.Now().Add(20 * time.Second); time.Now().Before(end); time.Sleep(10 * time.Millisecond) {
		if j := e.app.fiveTools.get(); !j.Running {
			if j.Failed {
				t.Fatalf("import failed: %s", j.Msg)
			}
			return j.Msg
		}
	}
	t.Fatal("import still running")
	return ""
}

// starterEnv is a test env with the fake servers, a GM, and a second
// campaign that no import may touch.
func starterEnv(t *testing.T) (e testEnv, gm string, other int64) {
	t.Helper()
	e = newTestEnv(t)
	serveStarterSets(t, e)
	camp, err := e.q().
		CreateCampaign(t.Context(), sqlc.CreateCampaignParams{Name: "Other", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	return e, e.token(t, enum.RoleAdmin), camp.ID
}

func (e testEnv) ticked(t *testing.T, campaign int64) []string {
	t.Helper()
	keys, err := e.q().ListCampaignSources(t.Context(), campaign)
	if err != nil {
		t.Fatal(err)
	}
	return keys
}

// pickerTicks reports, for each book id, whether the picker's box is ticked.
func pickerTicks(t *testing.T, picker string, ids ...string) map[string]bool {
	t.Helper()
	out := map[string]bool{}
	for _, id := range ids {
		tag := regexp.MustCompile(`<input[^>]*value="` + id + `"[^>]*>`).FindString(picker)
		if tag == "" {
			t.Fatalf("picker has no box for %s", id)
		}
		out[id] = strings.Contains(tag, "checked")
	}
	return out
}

// TestStarterSetsImport: each starting set posted from a campaign's
// compendium imports and ticks its sources for that campaign only; the
// ?set= picker ticks the set's books.
func TestStarterSetsImport(t *testing.T) {
	e, gm, other := starterEnv(t)
	picker := e.do(
		t,
		http.MethodGet,
		views.FiveToolsPath+"?"+views.FiveToolsSetField+"="+views.FiveToolsSet2024,
		gm,
		nil,
		true,
	).Body.String()
	if got := pickerTicks(t, picker, "XPHB", "PHB"); !got["XPHB"] || got["PHB"] {
		t.Errorf("set picker ticks %v", got)
	}
	// "Tick 2024 set" sends the boxes ticked in the browser, which stay ticked.
	q := url.Values{
		views.FiveToolsSetField: {views.FiveToolsSet2024},
		views.FiveToolsField:    {"PHB"},
	}
	picker = e.do(
		t,
		http.MethodGet,
		views.FiveToolsPath+"?"+q.Encode(),
		gm,
		nil,
		true,
	).Body.String()
	if got := pickerTicks(
		t,
		picker,
		"XPHB",
		"PHB",
		"BRK",
	); !got["XPHB"] || !got["PHB"] ||
		got["BRK"] {
		t.Errorf("set picker with a ticked book ticks %v", got)
	}

	base := route.Path(route.CampaignCompendium, e.campaignID)
	for _, tt := range []struct {
		name, path, key string
		form            url.Values
	}{
		{name: "5etools 2024 set", path: base + route.CompendiumFiveTools, key: "5etools-xphb", form: url.Values{views.FiveToolsSetField: {views.FiveToolsSet2024}}},
		{name: "SRD 5.2", path: base + route.CompendiumOpen5e, key: compendium.SRD2024Key, form: url.Values{views.Open5eDocField: {compendium.SRD2024Key}}},
		{name: "SRD 5.1", path: base + route.CompendiumOpen5e, key: compendium.SRD2014Key, form: url.Values{views.Open5eDocField: {compendium.SRD2014Key}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			w := e.do(t, http.MethodPost, tt.path, gm, tt.form, false)
			if w.Code != http.StatusSeeOther || w.Header().Get("Location") != base {
				t.Fatalf("post: %d %q, want Browse %q", w.Code, w.Header().Get("Location"), base)
			}
			waitJob(t, e)
			if keys := e.ticked(t, e.campaignID); !slices.Contains(keys, tt.key) {
				t.Errorf("campaign sources %v lack %s", keys, tt.key)
			}
			if keys := e.ticked(t, other); len(keys) != 0 {
				t.Errorf("other campaign ticked %v", keys)
			}
		})
	}
	if keys := e.ticked(t, e.campaignID); slices.Contains(keys, "5etools-phb") {
		t.Errorf("2014 book ticked: %v", keys)
	}
	if w := e.do(
		t,
		http.MethodPost,
		base+route.CompendiumOpen5e,
		gm,
		url.Values{views.Open5eDocField: {"srd-3000"}},
		false,
	); !strings.Contains(
		w.Body.String(),
		"Pick an SRD",
	) {
		t.Errorf("unknown SRD: %d, not refused", w.Code)
	}
	// From the Data folder tab, the GM comes back to that tab.
	w := e.do(
		t,
		http.MethodPost,
		base+route.CompendiumOpen5e,
		gm,
		url.Values{
			views.Open5eDocField: {compendium.SRD2014Key},
			views.ImportKey:      {views.ImportTabFolder},
		},
		false,
	)
	if want := base + "?" + views.ImportKey + "=" + views.ImportTabFolder; w.Code != http.StatusSeeOther ||
		w.Header().Get("Location") != want {
		t.Errorf("folder tab post: %d %q, want %q", w.Code, w.Header().Get("Location"), want)
	}
	waitJob(t, e)
}

// TestPartialDownloadTicks: when one picked book fails to download, the
// books already in are still ticked for the campaign; the job fails and no
// other campaign is touched.
func TestPartialDownloadTicks(t *testing.T) {
	e, gm, other := starterEnv(t)
	camp := route.Path(route.CampaignCompendium, e.campaignID)
	if w := e.do(
		t,
		http.MethodPost,
		camp+route.CompendiumFiveTools,
		gm,
		url.Values{views.FiveToolsField: {"XPHB"}},
		false,
	); w.Code != http.StatusSeeOther {
		t.Fatalf("first post: %d", w.Code)
	}
	waitJob(t, e)
	if err := e.q().
		DisallowCampaignSource(t.Context(), sqlc.DisallowCampaignSourceParams{CampaignID: e.campaignID, SourceKey: "5etools-xphb"}); err != nil {
		t.Fatal(err)
	}
	if w := e.do(
		t,
		http.MethodPost,
		camp+route.CompendiumFiveTools,
		gm,
		url.Values{views.FiveToolsField: {"XPHB", "BRK"}},
		false,
	); w.Code != http.StatusSeeOther {
		t.Fatalf("second post: %d", w.Code)
	}
	for end := time.Now().Add(20 * time.Second); e.app.fiveTools.get().Running && time.Now().Before(end); {
		time.Sleep(10 * time.Millisecond)
	}
	if j := e.app.fiveTools.get(); j.Running || !j.Failed || !strings.Contains(j.Msg, "BRK") {
		t.Errorf("job: %+v, want a failed BRK download", j)
	}
	if keys := e.ticked(t, e.campaignID); !slices.Contains(keys, "5etools-xphb") {
		t.Errorf("campaign sources %v lack 5etools-xphb", keys)
	}
	if keys := e.ticked(t, other); len(keys) != 0 {
		t.Errorf("other campaign ticked %v", keys)
	}
}

// TestImportsTickFromCampaignOnly: picked books and saved web pages tick
// their sources in the campaign whose compendium they come from, never from
// the whole compendium; a folder import never ticks, and in a campaign it
// offers the sources the campaign leaves off ("Use selected").
func TestImportsTickFromCampaignOnly(t *testing.T) {
	e, gm, other := starterEnv(t)
	camp := route.Path(route.CampaignCompendium, e.campaignID)
	book := url.Values{views.FiveToolsField: {"XPHB"}}
	page := url.Values{
		"site":  {"wikidot"},
		"page":  {"spell:fireball"},
		"kind":  {"spell"},
		"name":  {"Fireball"},
		"level": {"3"},
		"desc":  {"A bright streak."},
	}

	// From the whole compendium: imported, ticked nowhere.
	if w := e.do(
		t,
		http.MethodPost,
		views.FiveToolsPath,
		gm,
		book,
		false,
	); w.Code != http.StatusSeeOther {
		t.Fatalf("global book post: %d", w.Code)
	}
	waitJob(t, e)
	if w := e.do(
		t,
		http.MethodPost,
		route.Compendium+route.CompendiumPage,
		gm,
		page,
		false,
	); w.Code != http.StatusSeeOther {
		t.Fatalf("global page save: %d %s", w.Code, w.Body)
	}
	for _, id := range []int64{e.campaignID, other} {
		if keys := e.ticked(t, id); len(keys) != 0 {
			t.Fatalf("global imports ticked %v in campaign %d", keys, id)
		}
	}

	// From the campaign's compendium: ticked there only.
	if w := e.do(
		t,
		http.MethodPost,
		camp+route.CompendiumFiveTools,
		gm,
		book,
		false,
	); w.Code != http.StatusSeeOther ||
		w.Header().Get("Location") != camp+"?"+views.ImportKey+"="+views.ImportTabFiveTools {
		t.Fatalf("campaign book post: %d %q", w.Code, w.Header().Get("Location"))
	}
	waitJob(t, e)
	if w := e.do(
		t,
		http.MethodPost,
		camp+route.CompendiumPage,
		gm,
		page,
		false,
	); w.Code != http.StatusSeeOther {
		t.Fatalf("campaign page save: %d %s", w.Code, w.Body)
	}
	keys := e.ticked(t, e.campaignID)
	slices.Sort(keys)
	if want := []string{"5etools-xphb", "wikidot:spell:fireball"}; !slices.Equal(keys, want) {
		t.Errorf("campaign ticked %v, want %v", keys, want)
	}
	if keys := e.ticked(t, other); len(keys) != 0 {
		t.Errorf("other campaign ticked %v", keys)
	}

	// A folder import ticks nothing, and lists what the campaign leaves off.
	if err := os.CopyFS(
		filepath.Join(e.app.compendiumDir(), "srd-copy"),
		os.DirFS(srdMini),
	); err != nil {
		t.Fatal(err)
	}
	body := e.do(t, http.MethodPost, camp+route.CompendiumImport, gm, nil, false).Body.String()
	_, notice, _ := strings.Cut(body, "Not on in this campaign")
	notice, _, _ = strings.Cut(notice, "Use selected")
	offered := regexp.MustCompile(`name="key" value="([^"]+)"`).FindAllStringSubmatch(notice, -1)
	if len(offered) != 1 || offered[0][1] != "srd-copy" {
		t.Errorf("folder import notice offers %v, want just srd-copy", offered)
	}
	if keys := e.ticked(t, e.campaignID); slices.Contains(keys, "srd-copy") {
		t.Errorf("folder import ticked: %v", keys)
	}
	if body := e.do(
		t,
		http.MethodPost,
		route.Compendium+route.CompendiumImport,
		gm,
		nil,
		false,
	).Body.String(); strings.Contains(
		body,
		"Use selected",
	) {
		t.Error("global folder import offers campaign sources")
	}
	// "Use selected" ticks just the posted keys and goes back to the compendium.
	w := e.do(
		t,
		http.MethodPost,
		route.Path(route.CampaignSources, e.campaignID),
		gm,
		url.Values{
			views.SourceKeyField: {"srd-copy"},
			"on":                 {"true"},
			"back":               {views.SourcesBackCompendium},
		},
		true,
	)
	if w.Header().Get("HX-Redirect") != camp {
		t.Errorf("use selected: %d %q", w.Code, w.Header().Get("HX-Redirect"))
	}
	if keys := e.ticked(t, e.campaignID); len(keys) != 3 || !slices.Contains(keys, "srd-copy") {
		t.Errorf("after use selected: %v", keys)
	}
	if keys := e.ticked(t, other); len(keys) != 0 {
		t.Errorf("other campaign ticked %v", keys)
	}
	// The priority card's counts: the campaign shows the source's entries,
	// the other campaign, which leaves it off, none.
	for id, shows := range map[int64]bool{e.campaignID: true, other: false} {
		counts, err := e.app.entryCounts(t.Context(), id)
		if c := counts["srd-copy"]; err != nil || c.Total == 0 || (c.Visible > 0) != shows {
			t.Errorf("campaign %d counts %+v %v", id, c, err)
		}
	}
}

// TestPickerKeepsJobTicks: the books a running or failed import picked stay
// ticked when the picker reloads; a successful end forgets them.
func TestPickerKeepsJobTicks(t *testing.T) {
	e, gm, _ := starterEnv(t)
	picker := func() map[string]bool {
		return pickerTicks(
			t,
			e.do(t, http.MethodGet, views.FiveToolsPath, gm, nil, true).Body.String(),
			"XPHB",
			"PHB",
		)
	}
	if got := picker(); got["XPHB"] || got["PHB"] {
		t.Fatalf("fresh picker ticks %v", got)
	}
	e.app.fiveTools.start(1)
	e.app.fiveTools.update(func(j *views.FiveToolsJob) { j.Books = []string{"PHB"} })
	if got := picker(); !got["PHB"] || got["XPHB"] {
		t.Errorf("running: %v", got)
	}
	e.app.fiveTools.update(func(j *views.FiveToolsJob) {
		*j = views.FiveToolsJob{Msg: "Download failed: boom", Failed: true, Books: j.Books}
	})
	if got := picker(); !got["PHB"] {
		t.Errorf("failed: %v", got)
	}
	// A real run ending well clears them.
	if w := e.do(
		t,
		http.MethodPost,
		views.FiveToolsPath,
		gm,
		url.Values{views.FiveToolsField: {"XPHB"}},
		false,
	); w.Code != http.StatusSeeOther {
		t.Fatalf("post: %d", w.Code)
	}
	waitJob(t, e)
	if j := e.app.fiveTools.get(); len(j.Books) != 0 {
		t.Errorf("books kept after success: %v", j.Books)
	}
}

// TestSourcesCardStarterSets: with nothing imported, the GM's campaign page
// offers the sets in its Sources card, posting to that campaign's
// compendium and coming back to the page; the card is a live view for the
// GM; once imported, it lists the ticked source instead, live too.
func TestSourcesCardStarterSets(t *testing.T) {
	e, gm, _ := starterEnv(t)
	player := e.token(t, enum.RolePlayer)
	page := route.Path(route.Campaign, e.campaignID)
	srdPost := route.Path(route.CampaignCompendium, e.campaignID) + route.CompendiumOpen5e
	key := views.LiveKey(views.LiveCampaignSources, e.campaignID)
	live := `data-live="` + key + `"`

	body := e.do(t, http.MethodGet, page, gm, nil, false).Body.String()
	if !strings.Contains(body, live) || !strings.Contains(body, `action="`+srdPost+`"`) {
		t.Fatal("Sources card offers no starting sets posting to the campaign's compendium")
	}
	if w := e.do(
		t,
		http.MethodGet,
		page,
		player,
		nil,
		false,
	); w.Code != http.StatusForbidden ||
		strings.Contains(w.Body.String(), live) {
		t.Errorf("player campaign page: %d", w.Code)
	}
	// The page's refetch names its campaign, as its socket does.
	refetch := route.Live + "?" + url.Values{
		"v":        {key},
		"campaign": {strconv.FormatInt(e.campaignID, 10)},
	}.Encode()
	if w := e.do(
		t,
		http.MethodGet,
		refetch,
		gm,
		nil,
		true,
	); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), `id="`+views.StarterSetsID+`"`) ||
		!strings.Contains(w.Body.String(), `action="`+srdPost+`"`) {
		t.Errorf("GM refetch: %d %s", w.Code, w.Body)
	}

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		srdPost,
		strings.NewReader(url.Values{views.Open5eDocField: {compendium.SRD2024Key}}.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://host"+page)
	req.RemoteAddr = lanAddr
	req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: gm})
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	if w.Code != http.StatusSeeOther || w.Header().Get("Location") != page {
		t.Fatalf("post: %d %q, want %q", w.Code, w.Header().Get("Location"), page)
	}
	msg := waitJob(t, e)

	if w := e.do(
		t,
		http.MethodGet,
		refetch,
		gm,
		nil,
		true,
	); w.Code != http.StatusOK || strings.Contains(w.Body.String(), "Import SRD") ||
		!strings.Contains(w.Body.String(), compendium.SRD2024Key) {
		t.Errorf("refetch after the import still offers the sets, or lacks the SRD: %d", w.Code)
	}
	// With sources in, the card still shows a running import's bar and a failed one's reason.
	for _, job := range []views.FiveToolsJob{{Running: true, Total: 2, Current: "Book Two"}, {Msg: "Could not add them to the campaign", Failed: true}} {
		e.app.fiveTools.update(func(j *views.FiveToolsJob) { *j = job })
		if body := e.do(
			t,
			http.MethodGet,
			refetch,
			gm,
			nil,
			true,
		).Body.String(); !strings.Contains(body, job.Current+job.Msg) ||
			!strings.Contains(body, compendium.SRD2024Key) {
			t.Errorf("card with sources hides the job %+v", job)
		}
	}
	e.app.fiveTools.update(func(j *views.FiveToolsJob) { *j = views.FiveToolsJob{} })
	// The live card turns into the source list by itself, so the finished
	// sets in it offer no link, only the summary.
	var b strings.Builder
	if err := views.StarterSets(views.FiveToolsJob{Msg: msg}, e.campaignID).
		Render(t.Context(), &b); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(b.String(), msg) || strings.Contains(b.String(), "Show entries") ||
		strings.Contains(b.String(), "data-live") {
		t.Errorf("finished card sets: %s", b.String())
	}
}
