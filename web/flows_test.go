package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"gm-session-manager/auth"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/hub"
)

// do sends a request as a device. htmx marks the request as an htmx one.
func (e testEnv) do(t *testing.T, method, path, cookie string, form url.Values, htmx bool) *httptest.ResponseRecorder {
	t.Helper()
	var body io.Reader
	if form != nil {
		body = strings.NewReader(form.Encode())
	}
	req := httptest.NewRequestWithContext(t.Context(), method, path, body)
	if form != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if htmx {
		req.Header.Set("HX-Request", "true")
	}
	req.RemoteAddr = lanAddr
	if cookie != "" {
		req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: cookie})
	}
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	return w
}

func (e testEnv) doJSON(t *testing.T, path, cookie string, body any) *httptest.ResponseRecorder {
	t.Helper()
	b, _ := json.Marshal(body)
	req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, path, strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "application/json")
	req.RemoteAddr = lanAddr
	req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: cookie})
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	return w
}

func (e testEnv) q() *sqlc.Queries { return e.app.Store.Queries }

func TestJoinOpenCampaign(t *testing.T) {
	e := newTestEnv(t)
	w := e.do(t, http.MethodPost, fmt.Sprintf("/join/%d/create", e.campaignID), "", url.Values{"name": {"Bob"}}, false)
	if w.Code != http.StatusSeeOther {
		t.Fatalf("status %d: %s", w.Code, w.Body)
	}
	c := sessionCookie(w)
	if c == nil || !e.alive(t, c.Value) {
		t.Fatal("open join did not log the device in")
	}
	// A device with a session cannot join again without logging out.
	if w := e.do(t, http.MethodPost, fmt.Sprintf("/join/%d/create", e.campaignID), c.Value, url.Values{"name": {"Eve"}}, false); sessionCookie(w) != nil {
		t.Error("second join issued another session")
	}
}

func TestClaimNeedsApproval(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, auth.RoleAdmin)

	claim := fmt.Sprintf("/join/%d/claim/%d", e.campaignID, e.characterID)
	for _, player := range []string{"", "Thirteen char"} {
		if w := e.do(t, http.MethodPost, claim, "", url.Values{"player": {player}}, false); w.Code != http.StatusUnprocessableEntity {
			t.Errorf("player %q: status %d", player, w.Code)
		}
	}
	if len(e.app.Hub.PendingList(e.campaignID)) != 0 {
		t.Fatal("a refused claim left a request")
	}
	w := e.do(t, http.MethodPost, claim, "", url.Values{"player": {"Mira"}}, false)
	if w.Code != http.StatusOK || sessionCookie(w) != nil {
		t.Fatalf("claim: status %d, cookie %v", w.Code, sessionCookie(w))
	}
	pending := e.app.Hub.PendingList(e.campaignID)
	if len(pending) != 1 || pending[0].Player != "Mira" {
		t.Fatalf("pending %+v", pending)
	}
	p, _ := e.app.Hub.PendingBySecret(pendingSecret(t, e, pending[0].ID))

	// Before approval the secret redeems nothing.
	if w := e.do(t, http.MethodPost, "/requests/complete", "", url.Values{"secret": {p.Secret}}, true); sessionCookie(w) != nil {
		t.Fatal("redeemed before approval")
	}
	// A player cannot approve; the GM can.
	player := e.token(t, auth.RolePlayer)
	approve := fmt.Sprintf("/campaigns/%d/requests/%s/approve", e.campaignID, p.ID)
	if w := e.do(t, http.MethodPost, approve, player, nil, true); w.Code != http.StatusForbidden {
		t.Errorf("player approve: %d", w.Code)
	}
	if w := e.do(t, http.MethodPost, approve, gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("approve: %d %s", w.Code, w.Body)
	}
	w = e.do(t, http.MethodPost, "/requests/complete", "", url.Values{"secret": {p.Secret}}, true)
	c := sessionCookie(w)
	if c == nil || w.Header().Get("HX-Redirect") != "/" {
		t.Fatalf("complete: %d %v", w.Code, w.Header())
	}
	sess, err := e.mgr.Resolve(t.Context(), c.Value)
	if err != nil || sess.CharacterID.Int64 != e.characterID {
		t.Fatalf("session %+v %v", sess, err)
	}

	// Kick: the session is gone and the device is anonymous again.
	if w := e.do(t, http.MethodPost, fmt.Sprintf("/sessions/%d/kick", sess.ID), gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("kick: %d %s", w.Code, w.Body)
	}
	if e.alive(t, c.Value) {
		t.Error("kicked session still works")
	}
}

// TestJoinBurst is the table scanning the code at once: approve-all takes
// only the requests the GM saw, on any GM page.
func TestJoinBurst(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, auth.RoleAdmin)
	if err := e.q().SetCampaignOpenJoin(t.Context(), sqlc.SetCampaignOpenJoinParams{ID: e.campaignID, OpenJoin: false}); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"Ann", "Ben", "Cat"} {
		e.do(t, http.MethodPost, fmt.Sprintf("/join/%d/create", e.campaignID), "", url.Values{"name": {name}, "player": {name}}, false)
	}
	p := e.app.Hub.PendingList(e.campaignID)
	if len(p) != 3 {
		t.Fatalf("pending %+v", p)
	}
	if w := e.do(t, http.MethodGet, "/compendium", gm, nil, false); !strings.Contains(w.Body.String(), "/requests/all/approve") || !strings.Contains(w.Body.String(), "Cat") {
		t.Error("compendium page does not show the requests")
	}

	// The GM saw only the first two when they clicked.
	all := fmt.Sprintf("/campaigns/%d/requests/all/", e.campaignID)
	if w := e.do(t, http.MethodPost, all+"approve", gm, url.Values{"id": {p[0].ID, p[1].ID}}, true); w.Code != http.StatusOK {
		t.Fatalf("approve all: %d %s", w.Code, w.Body)
	}
	if left := e.app.Hub.PendingList(e.campaignID); len(left) != 1 || left[0].ID != p[2].ID {
		t.Fatalf("still waiting %+v", left)
	}
	for _, r := range p[:2] {
		if w := e.do(t, http.MethodPost, "/requests/complete", "", url.Values{"secret": {r.Secret}}, true); sessionCookie(w) == nil {
			t.Errorf("%s not logged in", r.Name)
		}
		if w := e.do(t, http.MethodPost, "/requests/complete", "", url.Values{"secret": {r.Secret}}, true); sessionCookie(w) != nil {
			t.Errorf("%s redeemed twice", r.Name)
		}
	}

	if w := e.do(t, http.MethodPost, all+"deny", gm, url.Values{"id": {p[2].ID}}, true); w.Code != http.StatusOK {
		t.Fatalf("deny all: %d %s", w.Code, w.Body)
	}
	if left := e.app.Hub.PendingList(e.campaignID); len(left) != 0 {
		t.Errorf("still waiting after deny %+v", left)
	}
	if w := e.do(t, http.MethodPost, "/requests/complete", "", url.Values{"secret": {p[2].Secret}}, true); sessionCookie(w) != nil {
		t.Error("denied request logged in")
	}
}

// pendingSecret finds a request's secret the way the visitor's page does.
func pendingSecret(t *testing.T, e testEnv, id string) string {
	t.Helper()
	for _, p := range e.app.Hub.PendingList(e.campaignID) {
		if p.ID == id {
			return p.Secret
		}
	}
	t.Fatal("request not found")
	return ""
}

func TestWaitingPageCarriesSecret(t *testing.T) {
	e := newTestEnv(t)
	if err := e.q().SetCampaignOpenJoin(t.Context(), sqlc.SetCampaignOpenJoinParams{ID: e.campaignID, OpenJoin: false}); err != nil {
		t.Fatal(err)
	}
	w := e.do(t, http.MethodPost, fmt.Sprintf("/join/%d/create", e.campaignID), "", url.Values{"name": {"Zed"}, "player": {"Mira"}}, false)
	if sessionCookie(w) != nil {
		t.Fatal("closed campaign created a session without approval")
	}
	p := e.app.Hub.PendingList(e.campaignID)
	if len(p) != 1 || p[0].Kind != hub.PendingCreate {
		t.Fatalf("pending %+v", p)
	}
	if !strings.Contains(w.Body.String(), "/ws?pending="+p[0].Secret) {
		t.Error("waiting page does not connect its websocket")
	}
}

func TestCommandPermissions(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, auth.RolePlayer)
	gm := e.token(t, auth.RoleAdmin)
	other, err := e.q().CreateCharacter(t.Context(), sqlc.CreateCharacterParams{CampaignID: e.campaignID, Name: "Bob", System: "dnd5e"})
	if err != nil {
		t.Fatal(err)
	}
	bob, _, err := e.mgr.IssuePlayer(t.Context(), other.ID)
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := fmt.Sprintf("/characters/%d/cmd", e.characterID)
	tests := []struct {
		name, cookie, cmd string
		want              int
	}{
		{"owner sheet", player, "set ac 15", http.StatusOK},
		{"owner hp", player, "maxhp =20", http.StatusOK},
		{"gm hp", gm, "hp -3", http.StatusOK},
		{"gm condition", gm, "toggle prone", http.StatusOK},
		{"gm sheet refused", gm, "set ac 18", http.StatusUnprocessableEntity},
		{"other player refused", bob, "hp -1", http.StatusUnprocessableEntity},
		{"bad command", player, "dance", http.StatusUnprocessableEntity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := e.do(t, http.MethodPost, cmdPath, tt.cookie, url.Values{"cmd": {tt.cmd}}, true)
			if w.Code != tt.want {
				t.Errorf("status %d, want %d: %s", w.Code, tt.want, w.Body)
			}
		})
	}
	ch, _ := e.q().GetCharacter(t.Context(), e.characterID)
	if !strings.Contains(ch.SheetJson, `"ac_bonus":15`) || ch.MaxHp != 20 || ch.Hp != -3 {
		t.Errorf("character after commands: %+v", ch)
	}
	conds, _ := e.q().ListConditionsByCharacter(t.Context(), nullID(e.characterID))
	if len(conds) != 1 || conds[0].Slug != "prone" {
		t.Errorf("conditions %+v", conds)
	}
}

// TestHPQuickForm checks the HP card's form names a missing amount instead
// of failing as an unknown command.
func TestHPQuickForm(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, auth.RolePlayer)
	cmdPath := fmt.Sprintf("/characters/%d/cmd", e.characterID)
	for _, tt := range []struct {
		name, op, amount string
		want             int
	}{
		{"maxhp first", "", "", http.StatusOK},
		{"empty heal", "heal", "", http.StatusUnprocessableEntity},
		{"zero damage", "damage", "0", http.StatusUnprocessableEntity},
		{"bad op", "explode", "3", http.StatusUnprocessableEntity},
		{"damage", "damage", "3", http.StatusOK},
	} {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{"hp_op": {tt.op}, "hp_amount": {tt.amount}}
			if tt.op == "" {
				form = url.Values{"cmd": {"maxhp =20"}}
			}
			w := e.do(t, http.MethodPost, cmdPath, player, form, true)
			if w.Code != tt.want {
				t.Errorf("status %d, want %d: %s", w.Code, tt.want, w.Body)
			}
			if tt.want != http.StatusOK && (!strings.Contains(w.Body.String(), "hit points") || strings.Contains(w.Body.String(), "unknown command")) {
				t.Errorf("message: %s", w.Body)
			}
		})
	}
	if ch, _ := e.q().GetCharacter(t.Context(), e.characterID); ch.Hp != -3 {
		t.Errorf("hp %d, want -3", ch.Hp)
	}
}

// TestCommandBuilder renders every builder form for the owner only, and
// checks that its fields become commands with free text kept as typed.
func TestCommandBuilder(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, auth.RolePlayer)
	gm := e.token(t, auth.RoleAdmin)
	base := fmt.Sprintf("/characters/%d/builder?verb=", e.characterID)
	for _, verb := range []string{"", "hp", "conditions", "rest", "abilities", "profs", "details", "weapons", "mods", "slots", "spells", "items", "coins", "features"} {
		if w := e.do(t, http.MethodGet, base+verb, player, nil, true); w.Code != http.StatusOK {
			t.Errorf("verb %q: status %d: %s", verb, w.Code, w.Body)
		}
	}
	// The score field starts at the preselected ability's score.
	e.do(t, http.MethodPost, fmt.Sprintf("/characters/%d/cmd", e.characterID), player, url.Values{"cmd": {"set ability wis 17"}}, true)
	if w := e.do(t, http.MethodGet, base+"abilities&op=ability&subject=wis", player, nil, true); !strings.Contains(w.Body.String(), `value="17"`) {
		t.Errorf("ability not preselected: %s", w.Body)
	}
	if w := e.do(t, http.MethodGet, base+"hp", gm, nil, true); w.Code == http.StatusOK {
		t.Error("GM got the owner's builder")
	}
	cmdPath := fmt.Sprintf("/characters/%d/cmd", e.characterID)
	for _, form := range []url.Values{
		{"prefix": {"set ability"}, "subject": {"dex"}, "sign": {""}, "n": {"16"}},
		{"prefix": {"weapon add"}, "subject": {"Rapier"}, "ability": {"finesse"}, "dice": {"1d8"}, "dtype": {"piercing"}, "prof": {"prof"}},
		{"prefix": {"mod add"}, "subject": {"attack.Rapier"}, "sign": {""}, "n": {"1"}, "text": {`"Lucky" blade`}},
		{"prefix": {"set backstory"}, "text": {"Line one\nSaid \"hi\""}},
	} {
		if w := e.do(t, http.MethodPost, cmdPath, player, form, true); w.Code != http.StatusOK {
			t.Fatalf("%v: status %d: %s", form, w.Code, w.Body)
		}
	}
	// Tapping a skill or save opens just its proficiency; one tap sets it.
	for path, want := range map[string]string{
		"profs&op=skill&subject=stealth":       "Stealth",
		"profs&op=save&subject=con":            "Constitution save",
		"weapons&op=weapon-add&subject=Rapier": "weapon remove",
	} {
		if w := e.do(t, http.MethodGet, base+path, player, nil, true); !strings.Contains(w.Body.String(), want) {
			t.Errorf("%s lacks %q: %s", path, want, w.Body)
		}
	}
	for _, form := range []url.Values{
		{"prefix": {"prof set skill"}, "subject": {"stealth"}, "text": {"expertise"}},
		{"prefix": {"prof set save"}, "subject": {"con"}, "text": {"proficient"}},
	} {
		if w := e.do(t, http.MethodPost, cmdPath, player, form, true); w.Code != http.StatusOK {
			t.Fatalf("%v: status %d: %s", form, w.Code, w.Body)
		}
	}
	// Tapping the modifier opens it for editing.
	if w := e.do(t, http.MethodGet, base+"mods&op=mod-edit&subject=0", player, nil, true); !strings.Contains(w.Body.String(), `value="mod edit"`) || !strings.Contains(w.Body.String(), `Lucky`) {
		t.Errorf("edit form: %s", w.Body)
	}
	for _, form := range []url.Values{
		{"prefix": {"mod edit"}, "from": {"attack.Rapier"}, "fromnote": {`"Lucky" blade`}, "subject": {"ac"}, "sign": {""}, "n": {"2"}, "text": {"Shield spell"}},
		{"prefix": {"mod add"}, "subject": {"init"}, "sign": {""}, "n": {"5"}, "text": {"Alert"}},
		// The edit form's Remove button: its hx-vals replace the form's fields.
		{"prefix": {"mod remove"}, "subject": {"init"}, "text": {"Alert"}},
	} {
		if w := e.do(t, http.MethodPost, cmdPath, player, form, true); w.Code != http.StatusOK {
			t.Fatalf("%v: status %d: %s", form, w.Code, w.Body)
		}
	}
	ch, _ := e.q().GetCharacter(t.Context(), e.characterID)
	if !strings.Contains(ch.SheetJson, `"modifiers":[{"target":"ac","value":2,"note":"Shield spell"}]`) {
		t.Errorf("modifiers after edit and remove: %s", ch.SheetJson)
	}
	for _, want := range []string{`"dex":16`, `"save_profs":["con"]`, `"skill_profs":["stealth"],"expertise":["stealth"]`, `"name":"Rapier","ability":"finesse","dice":"1d8","damage_type":"piercing","proficient":true`, `"backstory":"Line one\nSaid \"hi\""`} {
		if !strings.Contains(ch.SheetJson, want) {
			t.Errorf("sheet lacks %s: %s", want, ch.SheetJson)
		}
	}
}

// TestServerSideStripping checks that hidden data is never in a player's
// response at all, not merely hidden on the page.
func TestServerSideStripping(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	player := e.token(t, auth.RolePlayer)

	ent, err := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{
		CampaignID: e.campaignID, Kind: "npc", Name: "Mira", IsRevealed: true,
		FieldsJson: `{"role":{"value":"Innkeeper","is_revealed":true},"secret":{"value":"Is the lich","is_revealed":false}}`,
	})
	if err != nil {
		t.Fatal(err)
	}
	body := e.do(t, http.MethodGet, fmt.Sprintf("/entities/%d", ent.ID), player, nil, false).Body.String()
	if !strings.Contains(body, "Innkeeper") || strings.Contains(body, "lich") || strings.Contains(body, ">secret<") {
		t.Errorf("entity page leaks or misses fields")
	}

	note, _ := e.q().CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: "generic", Title: "Plans"})
	secret, _ := e.q().CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: "gm"})
	if w := e.do(t, http.MethodGet, fmt.Sprintf("/blocks/%d/doc", secret.ID), player, nil, false); w.Code != http.StatusNotFound {
		t.Errorf("player read a GM block: %d", w.Code)
	}
	if w := e.do(t, http.MethodGet, fmt.Sprintf("/notes/%d", note.ID), player, nil, false); w.Code != http.StatusNotFound || strings.Contains(w.Body.String(), "Plans") {
		t.Errorf("player saw a GM-only note: %d", w.Code)
	}

	enc, _ := e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"})
	if _, err := e.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: enc.ID, Name: "Ogre", Hp: 59, MaxHp: 59, Ac: 11}); err != nil {
		t.Fatal(err)
	}
	if _, err := e.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: enc.ID, CharacterID: nullID(e.characterID), Name: "Alice"}); err != nil {
		t.Fatal(err)
	}
	cb, _ := e.q().ListCombatants(ctx, enc.ID)
	if err := e.q().UpdateEncounterState(ctx, sqlc.UpdateEncounterStateParams{ID: enc.ID, Status: "active", Round: 1, ActiveCombatantID: nullID(cb[0].ID)}); err != nil {
		t.Fatal(err)
	}
	body = e.do(t, http.MethodGet, fmt.Sprintf("/campaigns/%d/combat", e.campaignID), player, nil, false).Body.String()
	if !strings.Contains(body, "Ogre") || strings.Contains(body, "59/59") {
		t.Errorf("combat page leaks monster HP or misses the monster")
	}
}

func TestCollabConflict(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	note, _ := e.q().CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: "generic", Title: "Log"})
	b, _ := e.q().CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: "players", PlayersCanEdit: true})
	path := fmt.Sprintf("/blocks/%d/steps", b.ID)
	step := json.RawMessage(`{"stepType":"replace","from":1,"to":1,"slice":{"content":[{"type":"text","text":"a"}]}}`)

	w := e.doJSON(t, path, gm, map[string]any{"version": 0, "clientID": "c1", "steps": []json.RawMessage{step}})
	if w.Code != http.StatusOK {
		t.Fatalf("first submit: %d %s", w.Code, w.Body)
	}
	// A second client still at version 0 must rebase.
	w = e.doJSON(t, path, e.token(t, auth.RolePlayer), map[string]any{"version": 0, "clientID": "c2", "steps": []json.RawMessage{step}})
	if w.Code != http.StatusConflict {
		t.Fatalf("stale submit: %d", w.Code)
	}
	var p stepsPayload
	if err := json.Unmarshal(w.Body.Bytes(), &p); err != nil || len(p.Steps) != 1 || p.ClientIDs[0] != "c1" {
		t.Errorf("conflict payload %+v %v", p, err)
	}

	// After a snapshot at version 1 the old steps are trimmed: asking for
	// them reports Gone so the client reloads.
	snap := map[string]any{"version": 1, "doc": json.RawMessage(`{"type":"doc","content":[]}`), "text": "a", "links": []any{}}
	if w := e.doJSON(t, fmt.Sprintf("/blocks/%d/snapshot", b.ID), gm, snap); w.Code != http.StatusNoContent {
		t.Fatalf("snapshot: %d %s", w.Code, w.Body)
	}
	if w := e.do(t, http.MethodGet, path+"?since=0", gm, nil, false); w.Code != http.StatusGone {
		t.Errorf("trimmed steps: %d", w.Code)
	}
}

// TestPlayerBlocksInGameSession: players add 'players' blocks to the GM's
// game session notes and manage only the blocks they wrote.
func TestPlayerBlocksInGameSession(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	alice := e.token(t, auth.RolePlayer)
	other, err := e.q().CreateCharacter(ctx, sqlc.CreateCharacterParams{CampaignID: e.campaignID, Name: "Bob", System: "dnd5e"})
	if err != nil {
		t.Fatal(err)
	}
	bob, _, err := e.mgr.IssuePlayer(ctx, other.ID)
	if err != nil {
		t.Fatal(err)
	}
	log, _ := e.q().CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: "game_session", Title: "Session 1"})
	plain, _ := e.q().CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: "generic", Title: "Lore"})
	gmBlock, _ := e.q().CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: log.ID, Visibility: "players"})

	add := func(cookie string, noteID int64, vis string) int {
		return e.do(t, http.MethodPost, fmt.Sprintf("/notes/%d/blocks", noteID), cookie, url.Values{"visibility": {vis}}, true).Code
	}
	if got := add(alice, log.ID, "players"); got != http.StatusOK {
		t.Fatalf("player add to game session: %d", got)
	}
	for _, tt := range []struct {
		name   string
		noteID int64
		vis    string
	}{{"secret", log.ID, "listed"}, {"gm only", log.ID, "gm"}, {"generic gm note", plain.ID, "players"}} {
		if got := add(alice, tt.noteID, tt.vis); got != http.StatusUnprocessableEntity {
			t.Errorf("%s: %d, want 422", tt.name, got)
		}
	}

	blocks, _ := e.q().ListReadableBlocks(ctx, sqlc.ListReadableBlocksParams{NoteID: log.ID, IsGm: true})
	own := blocks[len(blocks)-1]
	if own.AuthorCharacterID.Int64 != e.characterID {
		t.Fatalf("author = %v", own.AuthorCharacterID)
	}
	// The author edits without players_can_edit; other players cannot.
	step := json.RawMessage(`{"stepType":"replace","from":1,"to":1,"slice":{"content":[{"type":"text","text":"a"}]}}`)
	steps := func(cookie string) int {
		return e.doJSON(t, fmt.Sprintf("/blocks/%d/steps", own.ID), cookie, map[string]any{"version": own.Version, "clientID": "c", "steps": []json.RawMessage{step}}).Code
	}
	if own.PlayersCanEdit || steps(alice) != http.StatusOK || steps(bob) != http.StatusNotFound {
		t.Error("only the author may edit a block without players_can_edit")
	}
	post := func(cookie, path string, form url.Values) int {
		return e.do(t, http.MethodPost, path, cookie, form, true).Code
	}
	access := url.Values{"visibility": {"players"}, "players_can_edit": {"true"}}
	const refused = http.StatusUnprocessableEntity // a.fail answers htmx with a flash
	for _, tt := range []struct {
		name, cookie, path string
		form               url.Values
		want               int
	}{
		{"bob re-scopes alice's block", bob, fmt.Sprintf("/blocks/%d/access", own.ID), access, refused},
		{"bob deletes alice's block", bob, fmt.Sprintf("/blocks/%d/delete", own.ID), nil, refused},
		{"alice deletes the gm's block", alice, fmt.Sprintf("/blocks/%d/delete", gmBlock.ID), nil, refused},
		{"alice moves her block", alice, fmt.Sprintf("/blocks/%d/move", own.ID), url.Values{"dir": {"up"}}, refused},
		{"alice re-scopes her block", alice, fmt.Sprintf("/blocks/%d/access", own.ID), access, http.StatusOK},
		{"alice deletes her block", alice, fmt.Sprintf("/blocks/%d/delete", own.ID), nil, http.StatusOK},
	} {
		if got := post(tt.cookie, tt.path, tt.form); got != tt.want {
			t.Errorf("%s: %d, want %d", tt.name, got, tt.want)
		}
	}
}
