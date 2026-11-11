package web

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
)

// TestPlayerHomebrew: a player's homebrew spell lands on their sheet and in
// their campaign only; the GM shares it with another campaign and takes it
// back; other players may not change it.
func TestPlayerHomebrew(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)
	camp2, err := e.q().CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Other", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	ch2, err := e.q().CreateCharacter(ctx, sqlc.CreateCharacterParams{CampaignID: camp2.ID, Name: "Bram", System: "dnd5e"})
	if err != nil {
		t.Fatal(err)
	}
	other, _, err := e.mgr.IssuePlayer(ctx, ch2.ID)
	if err != nil {
		t.Fatal(err)
	}

	spell := url.Values{"kind": {"spell"}, "name": {"Glitter Bolt"}, "level": {"1"}, "casting_time": {"action"}, "desc": {"A sparkly ranged spell attack."}, "add": {"1"}}
	w := e.do(t, http.MethodPost, "/compendium/homebrew", player, spell, false)
	if w.Code != http.StatusSeeOther || w.Header().Get("Location") != fmt.Sprintf("/characters/%d", e.characterID) {
		t.Fatalf("create: %d %v %s", w.Code, w.Header(), w.Body)
	}
	slug := fmt.Sprintf("glitter-bolt-%d", e.characterID)
	ch, _ := e.q().GetCharacter(ctx, e.characterID)
	if !strings.Contains(ch.SheetJson, `"slug":"`+slug+`"`) {
		t.Errorf("spell not added to the sheet: %s", ch.SheetJson)
	}
	entry, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "spell", Slug: slug})
	if err != nil || !entry.Scoped {
		t.Fatalf("entry %+v %v", entry, err)
	}
	if w := e.do(t, http.MethodPost, "/compendium/homebrew", player, url.Values{"kind": {"monster"}, "name": {"Pet"}}, false); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("player monster: %d", w.Code)
	}

	search := "/compendium/search?kind=spell&q=glitter"
	entryPath := fmt.Sprintf("/compendium/entries/%d", entry.ID)
	sees := func(cookie string) bool {
		return strings.Contains(e.do(t, http.MethodGet, search, cookie, nil, true).Body.String(), "Glitter Bolt")
	}
	if !sees(player) || !sees(gm) || sees(other) {
		t.Errorf("before linking: author %v gm %v other %v", sees(player), sees(gm), sees(other))
	}
	if w := e.do(t, http.MethodGet, fmt.Sprintf("/campaigns/%d/compendium/entries/%d", camp2.ID, entry.ID), other, nil, false); w.Code != http.StatusNotFound {
		t.Errorf("other campaign opened the entry: %d", w.Code)
	}
	// By slug too: neither the lookup nor a command reaches it.
	lookup := func(cookie string) bool {
		return strings.Contains(e.do(t, http.MethodGet, "/compendium/lookup?kind=spell&slug="+slug, cookie, nil, true).Body.String(), "sparkly")
	}
	if !lookup(player) || lookup(other) {
		t.Errorf("lookup by slug: author %v other %v", lookup(player), lookup(other))
	}
	if w := e.do(t, http.MethodPost, fmt.Sprintf("/characters/%d/cmd", ch2.ID), other, url.Values{"cmd": {"spell add " + slug}}, true); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("other campaign added the spell by slug: %d", w.Code)
	}
	// Nor by entry number, through the command builder.
	blade := url.Values{"kind": {"item"}, "name": {"Moon Blade"}, "damage": {"1d8 radiant"}, "properties": {"Moonlit"}}
	if w := e.do(t, http.MethodPost, "/compendium/homebrew", player, blade, false); w.Code != http.StatusSeeOther {
		t.Fatalf("create item: %d %s", w.Code, w.Body)
	}
	item, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "item", Slug: fmt.Sprintf("moon-blade-%d", e.characterID)})
	if err != nil {
		t.Fatal(err)
	}
	builder := func(cookie string, characterID int64) bool {
		path := fmt.Sprintf("/characters/%d/builder?verb=weapons&op=weapon-add&entry=%d", characterID, item.ID)
		return strings.Contains(e.do(t, http.MethodGet, path, cookie, nil, true).Body.String(), "Moonlit")
	}
	if !builder(player, e.characterID) || builder(other, ch2.ID) {
		t.Errorf("builder by entry: author %v other %v", builder(player, e.characterID), builder(other, ch2.ID))
	}

	if w := e.do(t, http.MethodPost, entryPath+"/campaigns", gm, url.Values{"campaign_id": {fmt.Sprint(camp2.ID)}}, true); w.Code >= 400 {
		t.Fatalf("link: %d %s", w.Code, w.Body)
	}
	if !sees(other) {
		t.Error("linked campaign does not see the entry")
	}
	if w := e.do(t, http.MethodPost, entryPath, other, spell, false); w.Code != http.StatusForbidden {
		t.Errorf("another player edited the entry: %d", w.Code)
	}
	if w := e.do(t, http.MethodPost, fmt.Sprintf("%s/campaigns/%d/delete", entryPath, camp2.ID), gm, nil, true); w.Code >= 400 {
		t.Fatalf("unlink: %d %s", w.Code, w.Body)
	}
	if sees(other) {
		t.Error("unlinked campaign still sees the entry")
	}
	if w := e.do(t, http.MethodPost, entryPath+"/delete", player, nil, false); w.Code != http.StatusSeeOther {
		t.Errorf("author delete: %d %s", w.Code, w.Body)
	}
}

// TestScopedHomebrewPickers: homebrew scoped to campaign B is listed and
// applicable in the creation and level-up pickers for B's characters only.
func TestScopedHomebrewPickers(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	campB, err := e.q().CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "B", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	owner, err := e.q().CreateCharacter(ctx, sqlc.CreateCharacterParams{CampaignID: campB.ID, Name: "Bram", System: "dnd5e"})
	if err != nil {
		t.Fatal(err)
	}
	for _, hb := range []struct {
		kind       enum.CompendiumKind
		slug, data string
	}{
		{"class", "riftblade", `{"hit_die":8}`},
		{"feat", "rift-step", `{}`},
		{"spell", "rift-bolt", `{"level":1,"classes":["wizard"]}`},
	} {
		row, err := e.q().CreateHomebrew(ctx, sqlc.CreateHomebrewParams{Kind: hb.kind, Slug: hb.slug, Name: hb.slug, DataJson: hb.data,
			OwnerCharacterID: sql.NullInt64{Int64: owner.ID, Valid: true}, Scoped: true})
		if err != nil {
			t.Fatal(err)
		}
		if err := e.q().LinkEntryCampaign(ctx, sqlc.LinkEntryCampaignParams{EntryID: row.ID, CampaignID: campB.ID}); err != nil {
			t.Fatal(err)
		}
	}
	for _, tt := range []struct {
		name     string
		campaign int64
		want     bool
	}{
		{"other campaign", e.campaignID, false},
		{"linked campaign", campB.ID, true},
		{"everything (GM)", 0, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			listed := func(kind enum.CompendiumKind, slug string) bool {
				rows, err := e.q().ListCompendiumNames(ctx, sqlc.ListCompendiumNamesParams{Kind: kind, CampaignID: tt.campaign})
				if err != nil {
					t.Fatal(err)
				}
				return slices.ContainsFunc(rows, func(r sqlc.ListCompendiumNamesRow) bool { return r.Slug == slug })
			}
			spells, err := e.q().ListSpellsForClass(ctx, sqlc.ListSpellsForClassParams{ClassSlug: "wizard", MaxLevel: 1, CampaignID: tt.campaign})
			if err != nil {
				t.Fatal(err)
			}
			spellListed := slices.ContainsFunc(spells, func(r sqlc.ListSpellsForClassRow) bool { return r.Slug == "rift-bolt" })
			if listed("class", "riftblade") != tt.want || listed("feat", "rift-step") != tt.want || spellListed != tt.want {
				t.Errorf("listed: class %v feat %v spell %v, want %v", listed("class", "riftblade"), listed("feat", "rift-step"), spellListed, tt.want)
			}
			_, err = e.app.wizardPick(ctx, enum.KindClass, "riftblade", tt.campaign)
			var ue userError
			if (err == nil) != tt.want || (err != nil && !errors.As(err, &ue)) {
				t.Errorf("pick: %v, want applicable %v (as a user error otherwise)", err, tt.want)
			}
		})
	}
}

// TestGMHomebrewClassAndWikidot: a GM's structured homebrew class keeps its
// hit die and features; the wiki import takes wiki pages from the GM only.
func TestGMHomebrewClassAndWikidot(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)

	class := url.Values{"kind": {"class"}, "name": {"Tinkerer"}, "hit_die": {"10"}, "caster": {"half"}, "saves": {"con", "int"},
		"features": {"Level 1: Tinkering\nYou tinker."}, "table": {"Gadgets: 2, 3"}}
	if w := e.do(t, http.MethodPost, "/compendium/homebrew", gm, class, false); w.Code != http.StatusSeeOther {
		t.Fatalf("create: %d %s", w.Code, w.Body)
	}
	entry, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "class", Slug: "tinkerer"})
	if err != nil || !strings.Contains(entry.DataJson, `"hit_die":10`) || !strings.Contains(entry.DataJson, `"name":"Tinkering"`) || !strings.Contains(entry.DataJson, `"saves":["con","int"]`) {
		t.Errorf("entry %s %v", entry.DataJson, err)
	}

	// A reviewed page, as the preview form posts it back (no network).
	review := url.Values{"site": {"wikidot"}, "page": {"lineage:giff"}, "book": {"Some Book"}, "kind": {"species"}, "name": {"Giff"}, "speed": {"30"}, "traits": {"Hippo Build. You are strong."}}
	for _, path := range []string{"/compendium/page/preview", "/compendium/page", "/compendium/folder"} {
		if w := e.do(t, http.MethodPost, path, player, review, false); w.Code != http.StatusForbidden {
			t.Errorf("player %s: %d", path, w.Code)
		}
	}
	if w := e.do(t, http.MethodPost, "/compendium/folder", gm, nil, false); w.Code != http.StatusForbidden {
		t.Errorf("folder from another device: %d", w.Code)
	}
	if w := e.do(t, http.MethodPost, "/compendium/page/preview", gm, url.Values{"page": {"asdasd"}}, false); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "one of the sites listed below") || !strings.Contains(w.Body.String(), `role="alert"`) {
		t.Errorf("not a link: %d", w.Code)
	}
	review.Set("back", "q=giff&kind=species&evil=1")
	w := e.do(t, http.MethodPost, "/compendium/page", gm, review, false)
	if w.Code != http.StatusSeeOther {
		t.Fatalf("save: %d %s", w.Code, w.Body)
	}
	// The entry links back to the filtered list, wiki tab open; other keys are dropped.
	if w := e.do(t, http.MethodGet, w.Header().Get("Location"), gm, nil, false); !strings.Contains(w.Body.String(), `href="/compendium?import=page&amp;kind=species&amp;q=giff#compendium-import"`) {
		t.Errorf("back link missing: %d %s", w.Code, w.Body)
	}
	review.Del("back")
	giff, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "species", Slug: "giff"})
	if err != nil || giff.SourceKey != "wikidot:lineage:giff" || !strings.Contains(giff.DataJson, "Hippo Build") {
		t.Errorf("saved %+v %v", giff, err)
	}
	if src, err := e.q().GetCompendiumSource(ctx, "wikidot:lineage:giff"); err != nil || src.Attribution != "Some Book via https://dnd5e.wikidot.com/lineage:giff" {
		t.Errorf("source %+v %v", src, err)
	}
	review.Set("kind", "feat")
	if w := e.do(t, http.MethodPost, "/compendium/page", gm, review, false); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("kind changed: %d", w.Code)
	}

	// A page of no known kind takes the kind the GM picked.
	spell := url.Values{"site": {"wikidot"}, "page": {"spell:fireball"}, "kind": {"spell"}, "name": {"Fireball"}, "level": {"3"}, "desc": {"A bright streak."}}
	if w := e.do(t, http.MethodPost, "/compendium/page", gm, spell, false); w.Code != http.StatusSeeOther {
		t.Fatalf("save spell: %d %s", w.Code, w.Body)
	}
	if fb, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "spell", Slug: "fireball"}); err != nil || fb.SourceKey != "wikidot:spell:fireball" {
		t.Errorf("saved %+v %v", fb, err)
	}

	// A monster page keeps its statblock and picture.
	monster := url.Values{"site": {"aidedd"}, "page": {"howler"}, "kind": {"monster"}, "name": {"Howler"}, "ac": {"16"}, "hp": {"90"},
		"str": {"17"}, "monster_speed": {"walk 40"}, "actions": {"Rending Bite. Bites."}, "image_url": {"https://www.aidedd.org/dnd/images/howler.jpg"}}
	if w := e.do(t, http.MethodPost, "/compendium/page", gm, monster, false); w.Code != http.StatusSeeOther {
		t.Fatalf("save monster: %d %s", w.Code, w.Body)
	}
	howler, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "howler"})
	if err != nil || howler.SourceKey != "aidedd:howler" || howler.ImageUrl.String != monster.Get("image_url") || !strings.Contains(howler.DataJson, `"Rending Bite"`) || !strings.Contains(howler.DataJson, `"str":17`) {
		t.Errorf("saved %+v %v", howler, err)
	}
	monster.Set("site", "evil")
	if w := e.do(t, http.MethodPost, "/compendium/page", gm, monster, false); w.Code != http.StatusBadRequest {
		t.Errorf("unknown site: %d", w.Code)
	}
}
