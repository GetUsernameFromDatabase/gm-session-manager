package web

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
)

// Allow-list test sources: "on-book" is ticked in the test campaign;
// "off-book" (2024, so it outranks on-book) was imported after and is off.
const (
	onBook  = "on-book"
	offBook = "off-book"
)

// sourcesEnv is a campaign with the srd-mini data and on-book ticked, and
// wizard subclasses only off-book has (Off School), both books have (Shared
// School) and homebrew has (Brew School).
func sourcesEnv(t *testing.T) testEnv {
	t.Helper()
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	e.source(t, onBook, "2014")
	if err := e.q().AllowAllSources(ctx, e.campaignID); err != nil {
		t.Fatal(err)
	}
	e.source(t, offBook, "2024")
	sub := func(source, slug, name string) {
		data := fmt.Sprintf(`{"subclass_of":"wizard","features":[{"name":"%s Trick","desc":"from %s","levels":[3]}]}`, name, source)
		if err := e.q().UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{Kind: "subclass", Slug: slug, Name: name, SourceKey: source, DataJson: data, ContentHash: source + slug, IsHomebrew: source == compendium.HomebrewSource}); err != nil {
			t.Fatal(err)
		}
	}
	sub(offBook, "off-school", "Off School")
	sub(offBook, "shared-school", "Shared School")
	sub(onBook, "shared-school", "Shared School")
	sub(compendium.HomebrewSource, "brew-school", "Brew School")
	return e
}

// wizard3 gives the test character Wizard 3 with the subclass of slug at source.
func (e testEnv) wizard3(t *testing.T, name, slug, source string) {
	t.Helper()
	e.setSheet(t, dnd5e.Sheet{Classes: []dnd5e.ClassLevel{{
		Slug: "wizard", Name: "Wizard", Level: 3, HitDie: 6, Caster: dnd5e.CasterFull,
		Subclass: name, SubclassSlug: slug, SubclassSource: source,
	}}})
}

// TestCampaignSourceVisibility: pickers and lists in a campaign offer only
// ticked sources and homebrew, falling back to a ticked copy of an entry an
// unticked source outranks; the GM without a campaign sees everything.
func TestCampaignSourceVisibility(t *testing.T) {
	e := sourcesEnv(t)
	tests := []struct {
		name     string
		campaign int64
		sub      string
		want     string // the source the picker offers; "" when hidden
	}{
		{"campaign hides an unticked source", e.campaignID, "Off School", ""},
		{"campaign falls back to the ticked copy", e.campaignID, "Shared School", onBook},
		{"campaign keeps homebrew", e.campaignID, "Brew School", compendium.HomebrewSource},
		{"campaign keeps ticked srd data", e.campaignID, "Evoker", "srd-mini"},
		{"all campaigns see the unticked source", 0, "Off School", offBook},
		{"all campaigns see the priority copy", 0, "Shared School", offBook},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subs, err := e.q().ListSubclassesOf(t.Context(), sqlc.ListSubclassesOfParams{ClassSlug: "wizard", CampaignID: tt.campaign})
			if err != nil {
				t.Fatal(err)
			}
			got := ""
			if i := slices.IndexFunc(subs, func(s sqlc.ListSubclassesOfRow) bool { return s.Name == tt.sub }); i >= 0 {
				got = subs[i].SourceKey
			}
			if got != tt.want {
				t.Errorf("%s from %q, want %q", tt.sub, got, tt.want)
			}
		})
	}

	// Choosing by name: an unticked-only subclass is refused, not dropped.
	refs := []struct {
		name     string
		campaign int64
		sub      string
		wantSlug string
		wantUser bool
	}{
		{"unticked in the campaign", e.campaignID, "Off School", "", true},
		{"ticked in the campaign", e.campaignID, "Shared School", "shared-school", false},
		{"unticked, no campaign", 0, "Off School", "off-school", false},
		{"unknown name", e.campaignID, "Nowhere School", "", false},
	}
	for _, tt := range refs {
		t.Run(tt.name, func(t *testing.T) {
			slug, _, err := subclassRef(t.Context(), e.q(), "wizard", dnd5e.ClassLevel{}, tt.sub, tt.campaign)
			var ue userError
			if slug != tt.wantSlug || errors.As(err, &ue) != tt.wantUser || err != nil && !tt.wantUser {
				t.Errorf("slug %q err %v, want %q (user error %v)", slug, err, tt.wantSlug, tt.wantUser)
			}
		})
	}

	// The compendium page: the campaign in the path, or every source.
	gm, player := e.token(t, auth.RoleAdmin), e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/campaigns/%d/compendium", e.campaignID)
	pages := []struct {
		name, cookie, path string
		want, hidden       []string
	}{
		{"player compendium", player, base + "?kind=subclass", []string{"Brew School", "Shared School"}, []string{"Off School"}},
		{"gm compendium in the campaign", gm, base + "?kind=subclass", []string{"Brew School", "Shared School", base + "/entries/"}, []string{"Off School"}},
		{"gm compendium, every campaign", gm, "/compendium?kind=subclass", []string{"Off School", "Brew School"}, nil},
		{"player search", player, base + "?q=school", []string{"Shared School"}, []string{"Off School"}},
		{"gm search in the campaign", gm, base + "?q=school", []string{"Shared School"}, []string{"Off School"}},
	}
	for _, tt := range pages {
		t.Run(tt.name, func(t *testing.T) {
			w := e.do(t, http.MethodGet, tt.path, tt.cookie, nil, false)
			if w.Code != http.StatusOK {
				t.Fatalf("status %d", w.Code)
			}
			for _, s := range tt.want {
				if !strings.Contains(w.Body.String(), s) {
					t.Errorf("page lacks %q", s)
				}
			}
			for _, s := range tt.hidden {
				if strings.Contains(w.Body.String(), s) {
					t.Errorf("page shows %q", s)
				}
			}
		})
	}
}

// TestCompendiumAddresses: the campaign is in the compendium's path, so an
// address shows the same view to whoever opens it.
func TestCompendiumAddresses(t *testing.T) {
	e := sourcesEnv(t)
	other, err := e.q().CreateCampaign(t.Context(), sqlc.CreateCampaignParams{Name: "Other", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	gm, player := e.token(t, auth.RoleAdmin), e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/campaigns/%d/compendium", e.campaignID)
	tests := []struct {
		name, cookie, path string
		htmx               bool
		code               int
		location           string
	}{
		{"player sent to their campaign", player, "/compendium?kind=subclass", false, http.StatusSeeOther, base + "?kind=subclass"},
		{"player entry sent to their campaign", player, "/compendium/entries/1", false, http.StatusSeeOther, base + "/entries/1"},
		{"player background request stays", player, "/compendium/lookup?kind=spell&slug=nothing", true, http.StatusOK, ""},
		{"gm old campaign key moves to the path", gm, "/compendium?q=school&campaign=" + fmt.Sprint(e.campaignID), false, http.StatusSeeOther, base + "?q=school"},
		{"gm whole compendium stays", gm, "/compendium?kind=subclass", false, http.StatusOK, ""},
		{"player in another campaign", player, fmt.Sprintf("/campaigns/%d/compendium", other.ID), false, http.StatusForbidden, ""},
		{"gm in a missing campaign", gm, "/campaigns/999/compendium", false, http.StatusNotFound, ""},
		{"gm in another campaign", gm, fmt.Sprintf("/campaigns/%d/compendium", other.ID), false, http.StatusOK, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := e.do(t, http.MethodGet, tt.path, tt.cookie, nil, tt.htmx)
			if w.Code != tt.code || w.Header().Get("Location") != tt.location {
				t.Errorf("got %d %q, want %d %q", w.Code, w.Header().Get("Location"), tt.code, tt.location)
			}
		})
	}

	// Links on a campaign's compendium stay in it; the GM's switcher opens
	// the same page elsewhere.
	links := []struct {
		name, cookie, path string
		want, hidden       []string
	}{
		{"gm list", gm, base + "?tab=homebrew", []string{
			`href="` + base + `/entries/`, `href="` + base + `/habitats"`, `href="` + base + `/homebrew/new"`,
			fmt.Sprintf(`href="/campaigns/%d/compendium?tab=homebrew"`, other.ID), `href="/compendium?tab=homebrew"`,
		}, []string{`href="/compendium/entries/`}},
		{"player list has no switcher", player, base, []string{`href="` + base + `/entries/`}, []string{`href="/compendium?`, "compendium-switch"}},
		{"gm whole compendium", gm, "/compendium", []string{`href="/compendium/entries/`, `href="` + base + `"`}, nil},
	}
	for _, tt := range links {
		t.Run(tt.name, func(t *testing.T) {
			body := e.do(t, http.MethodGet, tt.path, tt.cookie, nil, false).Body.String()
			for _, s := range tt.want {
				if !strings.Contains(body, s) {
					t.Errorf("page lacks %q", s)
				}
			}
			for _, s := range tt.hidden {
				if strings.Contains(body, s) {
					t.Errorf("page has %q", s)
				}
			}
		})
	}
}

// TestCampaignSourcePicks: the player cannot pick a subclass of an unticked
// source, while a sheet already pinned to one still renders it.
func TestCampaignSourcePicks(t *testing.T) {
	e := sourcesEnv(t)
	player := e.token(t, auth.RolePlayer)
	path := fmt.Sprintf("/characters/%d/subclass", e.characterID)
	tests := []struct {
		name, sub, wantSource string
		want                  int
	}{
		{"unticked source refused", "Off School", "", http.StatusUnprocessableEntity},
		{"ticked copy used", "Shared School", onBook, http.StatusOK},
		{"homebrew allowed", "Brew School", compendium.HomebrewSource, http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e.wizard3(t, "Evoker", "", "")
			form := e.do(t, http.MethodGet, path+"?class=wizard", player, nil, true)
			if form.Code != http.StatusOK || strings.Contains(form.Body.String(), "Off School") {
				t.Fatalf("change form %d offers Off School: %v", form.Code, strings.Contains(form.Body.String(), "Off School"))
			}
			w := e.do(t, http.MethodPost, path, player, url.Values{"class": {"wizard"}, "subclass": {tt.sub}}, true)
			if w.Code != tt.want {
				t.Fatalf("status %d, want %d: %s", w.Code, tt.want, w.Body)
			}
			if cl := e.sheet(t).Classes[0]; tt.want == http.StatusOK && (cl.Subclass != tt.sub || cl.SubclassSource != tt.wantSource) {
				t.Errorf("class after change = %+v", cl)
			}
		})
	}

	t.Run("pinned to an unticked source still renders", func(t *testing.T) {
		e.wizard3(t, "Off School", "off-school", offBook)
		w := e.do(t, http.MethodGet, fmt.Sprintf("/characters/%d", e.characterID), player, nil, false)
		if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Off School") {
			t.Fatalf("sheet %d lacks Off School", w.Code)
		}
	})
}

// TestSubclassPickBySlug: two sources' subclasses of one name are separate
// choices, labelled with their source, and the picked slug decides the copy.
func TestSubclassPickBySlug(t *testing.T) {
	e := sourcesEnv(t)
	if err := e.q().UpsertCompendiumEntry(t.Context(), sqlc.UpsertCompendiumEntryParams{
		Kind: "subclass", Slug: "shared-school-hb", Name: "Shared School", SourceKey: compendium.HomebrewSource,
		DataJson: `{"subclass_of":"wizard"}`, ContentHash: "hb", IsHomebrew: true,
	}); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, auth.RolePlayer)
	path := fmt.Sprintf("/characters/%d/subclass", e.characterID)
	e.wizard3(t, "Evoker", "", "")
	form := e.do(t, http.MethodGet, path+"?class=wizard", player, nil, true).Body.String()
	for _, label := range []string{"Shared School · " + onBook, "Shared School · " + compendium.HomebrewSource} {
		if !strings.Contains(form, label) {
			t.Errorf("change form lacks %q", label)
		}
	}
	for slug, source := range map[string]string{"shared-school-hb": compendium.HomebrewSource, "shared-school": onBook} {
		e.wizard3(t, "Evoker", "", "")
		if w := e.do(t, http.MethodPost, path, player, url.Values{"class": {"wizard"}, "subclass": {slug}}, true); w.Code != http.StatusOK {
			t.Fatalf("%s: status %d: %s", slug, w.Code, w.Body)
		}
		if cl := e.sheet(t).Classes[0]; cl.Subclass != "Shared School" || cl.SubclassSlug != slug || cl.SubclassSource != source {
			t.Errorf("%s: class after change = %+v", slug, cl)
		}
	}
}

// TestCampaignSourcesSet: only the GM changes the list, and never unticks a
// source someone in the campaign uses; a new campaign starts with every
// source.
func TestCampaignSourcesSet(t *testing.T) {
	e := sourcesEnv(t)
	e.wizard3(t, "Shared School", "shared-school", onBook)
	gm, player := e.token(t, auth.RoleAdmin), e.token(t, auth.RolePlayer)
	path := fmt.Sprintf("/campaigns/%d/sources", e.campaignID)
	tests := []struct {
		name   string
		cookie string
		form   url.Values
		want   int
		on     []string // keys that must be ticked afterwards
		off    []string
	}{
		{"player refused", player, url.Values{"key": {offBook}, "on": {"true"}}, http.StatusForbidden, nil, []string{offBook}},
		{"no key", gm, url.Values{"on": {"true"}}, http.StatusUnprocessableEntity, nil, nil},
		{"tick", gm, url.Values{"key": {offBook}, "on": {"true"}}, http.StatusOK, []string{offBook}, nil},
		{"used source stays", gm, url.Values{"key": {offBook, onBook}}, http.StatusConflict, []string{offBook, onBook}, nil},
		{"untick unused", gm, url.Values{"key": {offBook}}, http.StatusOK, []string{onBook}, []string{offBook}},
		{"homebrew is never stored", gm, url.Values{"key": {compendium.HomebrewSource}, "on": {"true"}}, http.StatusOK, nil, []string{compendium.HomebrewSource}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := e.do(t, http.MethodPost, path, tt.cookie, tt.form, false)
			if w.Code != tt.want {
				t.Fatalf("status %d, want %d: %s", w.Code, tt.want, w.Body)
			}
			if tt.want == http.StatusConflict && !strings.Contains(w.Body.String(), "Alice") {
				t.Errorf("refusal does not name the user: %s", w.Body)
			}
			keys, err := e.q().ListCampaignSources(t.Context(), e.campaignID)
			if err != nil {
				t.Fatal(err)
			}
			for _, k := range tt.on {
				if !slices.Contains(keys, k) {
					t.Errorf("%s off", k)
				}
			}
			for _, k := range tt.off {
				if slices.Contains(keys, k) {
					t.Errorf("%s on", k)
				}
			}
		})
	}

	t.Run("card shows the user and locks the box", func(t *testing.T) {
		w := e.do(t, http.MethodGet, fmt.Sprintf("/campaigns/%d", e.campaignID), gm, nil, false)
		body := w.Body.String()
		if w.Code != http.StatusOK || !strings.Contains(body, "Used by Alice") || !strings.Contains(body, onBook+" is used by Alice") {
			t.Fatalf("campaign page %d lacks the used source", w.Code)
		}
	})

	t.Run("new campaign gets every source", func(t *testing.T) {
		w := e.post(t, "/campaigns", loopbackAddr, url.Values{"name": {"Fresh"}}, gm)
		if w.Code != http.StatusSeeOther {
			t.Fatalf("create %d", w.Code)
		}
		var id int64
		if _, err := fmt.Sscanf(w.Header().Get("Location"), "/campaigns/%d", &id); err != nil {
			t.Fatal(err)
		}
		keys, _ := e.q().ListCampaignSources(t.Context(), id)
		all, _ := e.q().ListCompendiumSources(t.Context())
		if len(keys) != len(all)-1 || slices.Contains(keys, compendium.HomebrewSource) { // homebrew implied
			t.Errorf("new campaign sources %v of %d", keys, len(all))
		}
	})
}
