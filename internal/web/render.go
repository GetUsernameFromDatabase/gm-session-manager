package web

import (
	"cmp"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

// viewerCtx is who a fragment is rendered for.
type viewerCtx struct {
	viewer     perm.Viewer
	campaignID int64
	pending    *hub.Pending
}

// CanSubscribe implements hub.Handler: a client may follow a block's
// collaboration steps when it may read the block.
func (a *App) CanSubscribe(ctx context.Context, c *hub.Client, blockID int64) bool {
	b, pb, err := a.loadBlock(ctx, blockID)
	return err == nil && b.CampaignID == c.CampaignID && perm.CanReadBlock(c.Viewer, pb)
}

func (a *App) waiting(p *hub.Pending) views.Waiting {
	state := map[hub.PendingState]string{hub.Waiting: "waiting", hub.Approved: "approved", hub.Denied: "denied"}[a.Hub.PendingState(p)]
	return views.Waiting{Secret: p.Secret, Name: p.Name, State: state}
}

// pendingViews groups every campaign's waiting requests and bring-back files,
// campaigns in the order of their oldest request.
func (a *App) pendingViews() []views.PendingGroup {
	var out []views.PendingGroup
	group := func(campaignID int64, name string) int {
		i := slices.IndexFunc(out, func(g views.PendingGroup) bool { return g.CampaignID == campaignID })
		if i < 0 {
			i = len(out)
			out = append(out, views.PendingGroup{CampaignID: campaignID, CampaignName: name})
		}
		return i
	}
	for _, p := range a.Hub.PendingList(0) {
		i := group(p.CampaignID, p.CampaignName)
		out[i].Requests = append(out[i].Requests, views.PendingView{ID: p.ID, CampaignID: p.CampaignID, Name: p.Name, Claim: p.Kind == hub.PendingClaim,
			Player: p.Player, Device: p.Device, Phone: p.Phone, Host: p.Host,
		})
	}
	// Files wait for review after the join requests; the GM's own too, so leaving its review page loses nothing.
	for _, f := range a.homeFiles.list() {
		i := group(f.campaignID, f.campaignName)
		out[i].BringBacks = append(out[i].BringBacks, views.BringBackNotice{ID: f.id, CampaignID: f.campaignID, Name: f.characterName, SavedAt: f.savedAt})
	}
	return out
}

// roster joins the characters, their sessions and who is connected now.
func (a *App) roster(ctx context.Context, campaignID int64) (views.Roster, error) {
	camp, err := a.Store.GetCampaign(ctx, campaignID)
	if err != nil {
		return views.Roster{}, err
	}
	chars, err := a.Store.ListCharactersByCampaign(ctx, campaignID)
	if err != nil {
		return views.Roster{}, err
	}
	sessions, err := a.Store.ListPlayerSessionsByCampaign(ctx, campaignID)
	if err != nil {
		return views.Roster{}, err
	}
	online := map[int64]string{}
	for _, ci := range a.Hub.Connected(campaignID) {
		if ci.SessionID != 0 {
			online[ci.SessionID] = ci.UserAgent
		}
	}
	campaigns, err := a.Store.ListCampaigns(ctx)
	if err != nil {
		return views.Roster{}, err
	}
	deleted, err := a.Store.ListDeletedCharacters(ctx, campaignID)
	if err != nil {
		return views.Roster{}, err
	}
	r := views.Roster{Campaign: camp, Deleted: deleted}
	for _, oc := range campaigns {
		if oc.ID != campaignID {
			r.OtherCampaigns = append(r.OtherCampaigns, oc)
		}
	}
	taken, err := a.Store.ListOpenTakeHomesByCampaign(ctx, campaignID)
	if err != nil {
		return views.Roster{}, err
	}
	for _, ch := range chars {
		rc := views.RosterCharacter{Character: ch}
		for _, t := range taken {
			if t.CharacterID == ch.ID {
				rc.TakenHome = t.CreatedAt
			}
		}
		for _, s := range sessions {
			if s.CharacterID.Int64 == ch.ID {
				ua, on := online[s.ID]
				d := views.Device{SessionID: s.ID, UserAgent: ua, Online: on}
				if !on {
					d.UserAgent, d.LastSeen = s.UserAgent, s.LastSeenAt
				}
				rc.Devices = append(rc.Devices, d)
			}
		}
		r.Characters = append(r.Characters, rc)
	}
	return r, nil
}

// playerCombat lists the started fights v's character is in and builds the
// one asked for, or the first (the live fight) when want is not among them.
func (a *App) playerCombat(ctx context.Context, v perm.Viewer, campaignID, want int64) ([]sqlc.Encounter, *views.EncounterView, error) {
	list, err := a.Store.ListCharacterEncounters(ctx, sqlc.ListCharacterEncountersParams{CampaignID: campaignID, CharacterID: nullID(v.CharacterID)})
	if err != nil || len(list) == 0 {
		return nil, nil, err
	}
	enc := list[0]
	if i := slices.IndexFunc(list, func(e sqlc.Encounter) bool { return e.ID == want }); i >= 0 {
		enc = list[i]
	}
	ev, err := a.encounterView(ctx, v, enc)
	return list, &ev, err
}

// hiddenNPCName is what players see for an NPC the GM has not revealed.
const hiddenNPCName = "Unknown creature"

// encounterView builds the turn order as v may see it. Hidden HP is zeroed
// here so it never reaches the template, let alone the browser.
func (a *App) encounterView(ctx context.Context, v perm.Viewer, enc sqlc.Encounter) (views.EncounterView, error) {
	rows, err := a.Store.ListCombatants(ctx, enc.ID)
	if err != nil {
		return views.EncounterView{}, err
	}
	conds, err := a.Store.ListConditionsByEncounter(ctx, enc.ID)
	if err != nil {
		return views.EncounterView{}, err
	}
	portraits, err := a.Store.CurrentCharacterLooks(ctx, enc.ID)
	if err != nil {
		return views.EncounterView{}, err
	}
	npcLooks, err := a.Store.CurrentEntityLooks(ctx, enc.ID)
	if err != nil {
		return views.EncounterView{}, err
	}
	ev := views.EncounterView{Encounter: enc, IsGM: v.IsGM}
	for _, r := range rows {
		rowConds := rowConditions(r, conds)
		sheet := pcSheetWith(r, rowConds) // parsed once for AC, initiative, defenses and out-of-fight; with its effects
		cv := views.CombatantView{
			ID: r.ID, Name: r.Name, AC: combatantArmorClass(r, sheet), HPVisible: r.HpVisible, CharacterID: r.CharacterID.Int64, EntityID: r.EntityID.Int64,
			Kind:   combatantKind(r),
			IsOwn:  v.IsPlayer() && r.CharacterID.Int64 == v.CharacterID,
			Active: enc.Status == enum.EncounterActive && enc.ActiveCombatantID.Int64 == r.ID,
			Out:    outOfFight(r, sheet),
		}
		if r.Initiative.Valid {
			cv.Initiative = fmt.Sprint(r.Initiative.Int64)
		}
		cv.ShowHP = perm.CanSeeCombatantHP(v, perm.Combatant{CharacterID: r.CharacterID.Int64, HPVisible: r.HpVisible})
		if cv.ShowHP {
			cv.HP, cv.MaxHP, cv.TempHP = r.Hp, r.MaxHp, r.TempHp
		}
		// One without max HP has no stats, so is never down.
		cv.Marker, cv.Down = r.Marker, r.Hp <= 0 && r.MaxHp > 0
		cv.Dead = sheet != nil && sheet.Dead()
		// A hidden NPC's name and linked monster are as secret as the NPC.
		hidden := !v.IsGM && r.EntityID.Valid && !r.EntityRevealed.Bool
		if v.IsGM {
			cv.Link, cv.InitMod = combatantLink(r, enc.CampaignID), initiativeMod(r, sheet)
		} else {
			cv.AC, cv.Link = 0, playerCombatantLink(v, r, enc.CampaignID, cv.ShowHP)
		}
		if hidden {
			cv.Name = hiddenNPCName
		} else {
			// The compendium is open to everyone; a visible look below overrides it.
			cv.Portrait = r.MonsterImage.String
		}
		for _, p := range portraits {
			look := toLook(p.ID, p.Name, p.Url, p.AttachmentID, p.IsRevealed, true)
			if p.CharacterID == r.CharacterID.Int64 && len(perm.VisibleLooks(v, p.CharacterID, []perm.Look{look})) > 0 {
				cv.Portrait = views.LookItem(look).Src()
			}
		}
		for _, p := range npcLooks {
			look := toLook(p.ID, p.Name, p.Url, p.AttachmentID, p.IsRevealed, true)
			// Players see it only when both the NPC and the look are revealed.
			if e, ok := perm.VisibleEntity(v, perm.Entity{IsRevealed: p.EntityRevealed, Images: []perm.Look{look}}); p.EntityID == r.EntityID.Int64 && ok && len(e.Images) > 0 {
				cv.Portrait = views.LookItem(look).Src()
			}
		}
		cv.Conditions = rowConds
		// A player learns the modifiers and holder of their own effects only,
		// and who applied one only when it was they.
		if !v.IsGM && !cv.IsOwn {
			for i := range cv.Conditions {
				cv.Conditions[i].ModsJson, cv.Conditions[i].HeldBy = noMods, sql.NullInt64{}
				if cv.Conditions[i].AppliedBy.Int64 != v.CharacterID {
					cv.Conditions[i].AppliedBy = sql.NullInt64{}
				}
			}
		}
		if cv.HeldLabels, err = a.heldLabels(ctx, v, cv.Conditions); err != nil {
			return ev, err
		}
		if v.IsGM || cv.IsOwn { // a player never learns another combatant's defenses
			cv.Defenses = combatantDefenses(r, sheet, cv.Conditions)
		}
		ev.Combatants = append(ev.Combatants, cv)
	}
	if v.IsGM {
		chars, err := a.Store.ListCharactersByCampaign(ctx, enc.CampaignID)
		if err != nil {
			return ev, err
		}
		for _, ch := range chars {
			ev.Characters = append(ev.Characters, views.Named{ID: ch.ID, Name: ch.Name})
		}
		ents, err := a.Store.ListEntitiesByCampaign(ctx, enc.CampaignID)
		if err != nil {
			return ev, err
		}
		for _, en := range ents {
			ev.Entities = append(ev.Entities, views.Named{ID: en.ID, Name: en.Name})
		}
	}
	return ev, nil
}

// combatantLink is the GM's page for a combatant's source: the character
// sheet, the NPC, or the monster in the campaign's compendium.
func combatantLink(r sqlc.ListCombatantsRow, campaignID int64) string {
	switch {
	case r.CharacterID.Valid:
		return route.Path(route.Character, r.CharacterID.Int64)
	case r.EntityID.Valid:
		return route.Path(route.Entity, r.EntityID.Int64)
	case r.MonsterRef.Valid:
		return views.CompendiumBase(campaignID) + route.Path(route.CompendiumEntry, r.MonsterRef.Int64)
	}
	return ""
}

// playerCombatantLink is the page a player may open for a combatant: a sheet
// they may view, a revealed NPC, or a monster whose HP the GM shows.
func playerCombatantLink(v perm.Viewer, r sqlc.ListCombatantsRow, campaignID int64, showHP bool) string {
	switch {
	case r.CharacterID.Valid:
		if perm.CanCharacter(v, r.CharacterID.Int64, perm.ScopeView) {
			return combatantLink(r, campaignID)
		}
	case r.EntityID.Valid:
		if r.EntityRevealed.Bool {
			return combatantLink(r, campaignID)
		}
	case showHP:
		return combatantLink(r, campaignID)
	}
	return ""
}

func combatantKind(r sqlc.ListCombatantsRow) string {
	switch {
	case r.CharacterID.Valid:
		return views.KindPC
	case r.EntityID.Valid:
		return views.KindNPC
	}
	return views.KindMonster
}

func (a *App) readableNotes(ctx context.Context, v perm.Viewer, campaignID int64) ([]sqlc.Note, error) {
	return a.Store.ListReadableNotes(ctx, sqlc.ListReadableNotesParams{
		CampaignID: campaignID, IsGm: v.IsGM, CharacterID: v.CharacterID, MaxResults: maxReadableNotes,
	})
}

// noteList is the notes index as v may see it, each note with its author:
// the owning player, or the GM.
func (a *App) noteList(ctx context.Context, v perm.Viewer, campaignID int64) ([]views.NoteItem, error) {
	notes, err := a.readableNotes(ctx, v, campaignID)
	if err != nil {
		return nil, err
	}
	names, err := a.characterNames(ctx, campaignID)
	if err != nil {
		return nil, err
	}
	var shared []int64 // only asked at home
	if v.AtHome {
		if shared, err = a.Store.ListNotesWithSharedBlocks(ctx, campaignID); err != nil {
			return nil, err
		}
	}
	items := make([]views.NoteItem, len(notes))
	for i, n := range notes {
		items[i] = views.NoteItem{
			Note: n, Author: author(v, names, n.OwnerCharacterID.Int64),
			CanManage: perm.CanManageNote(v, n.OwnerCharacterID.Int64) && !perm.LocksShared(v, slices.Contains(shared, n.ID)),
		}
	}
	return items, nil
}

func (a *App) characterNames(ctx context.Context, campaignID int64) (map[int64]string, error) {
	chars, err := a.Store.ListCharacterNames(ctx, campaignID)
	return nameMap(chars), err
}

func nameMap(chars []sqlc.ListCharacterNamesRow) map[int64]string {
	names := make(map[int64]string, len(chars))
	for _, c := range chars {
		names[c.ID] = c.Name
	}
	return names
}

// author is the character who wrote something (0 for the GM), linked to
// their sheet only when v may open it.
func author(v perm.Viewer, names map[int64]string, characterID int64) views.Author {
	if characterID == 0 {
		return views.Author{Name: "GM"}
	}
	au := views.Author{Name: names[characterID]}
	if perm.CanCharacter(v, characterID, perm.ScopeView) {
		au.Href = route.Path(route.Character, characterID)
	}
	return au
}

// noteView builds a note as v may see it; visible is false when v may see
// neither a block nor manage the note (so even its title stays hidden).
func (a *App) noteView(ctx context.Context, v perm.Viewer, note sqlc.Note) (views.NoteView, bool, error) {
	rows, err := a.Store.ListReadableBlocks(ctx, sqlc.ListReadableBlocksParams{NoteID: note.ID, IsGm: v.IsGM, CharacterID: v.CharacterID})
	if err != nil {
		return views.NoteView{}, false, err
	}
	manager := perm.CanManageNote(v, note.OwnerCharacterID.Int64)
	locked, err := a.noteLocked(ctx, v, note.ID)
	if err != nil {
		return views.NoteView{}, false, err
	}
	nv := views.NoteView{Note: note, CanManage: manager && !locked}
	if !manager && len(rows) == 0 {
		return nv, false, nil
	}
	nv.Visibilities = perm.BlockVisibilities(v, note.OwnerCharacterID.Int64, note.Kind == enum.NoteGameSession)
	chars, err := a.Store.ListCharacterNames(ctx, note.CampaignID)
	if err != nil {
		return nv, false, err
	}
	names := nameMap(chars)
	audRows, err := a.Store.ListNoteAudience(ctx, note.ID)
	if err != nil {
		return nv, false, err
	}
	audiences := make(map[int64][]int64)
	for _, ar := range audRows {
		audiences[ar.BlockID] = append(audiences[ar.BlockID], ar.CharacterID)
	}
	for _, r := range rows {
		audience := audiences[r.ID]
		pb := perm.Block{
			Visibility: r.Visibility, PlayersCanEdit: r.PlayersCanEdit, NoteOwnerID: note.OwnerCharacterID.Int64,
			Audience: audience, AuthorID: r.AuthorCharacterID.Int64,
		}
		if !perm.CanReadBlock(v, pb) {
			continue // the SQL and Go rules agree; this is belt and braces
		}
		bv := views.BlockView{
			ID: r.ID, Visibility: pb.Visibility, PlayersCanEdit: r.PlayersCanEdit,
			CanEdit: perm.CanEditBlock(v, pb), CanManage: perm.CanManageBlock(v, pb),
			// A NULL author is the GM, or the note's owner for blocks without one.
			Author: author(v, names, cmp.Or(pb.AuthorID, pb.NoteOwnerID)),
		}
		if v.IsGM {
			bv.Audience = audience // only the GM sees who else a secret is for
		} else {
			bv.ForYou = true
		}
		nv.Blocks = append(nv.Blocks, bv)
	}
	if v.IsGM && slices.Contains(nv.Visibilities, enum.VisListed) {
		nv.Characters = chars
	}
	return nv, true, nil
}

func (a *App) entityList(ctx context.Context, v perm.Viewer, campaignID int64) (views.EntityListView, error) {
	list := views.EntityListView{CampaignID: campaignID, IsGM: v.IsGM}
	var rows []sqlc.Entity
	var err error
	if v.IsGM {
		rows, err = a.Store.ListEntitiesByCampaign(ctx, campaignID)
	} else {
		rows, err = a.Store.ListRevealedEntities(ctx, campaignID)
	}
	if err != nil {
		return list, err
	}
	looks, err := a.Store.CurrentCampaignEntityLooks(ctx, campaignID)
	if err != nil {
		return list, err
	}
	for _, r := range rows {
		e, err := toEntity(r)
		if err != nil {
			return list, err
		}
		// Only the portrait: VisibleEntity drops it for players unless revealed.
		for _, l := range looks {
			if l.EntityID == e.ID {
				e.Images = []perm.Look{toLook(l.ID, l.Name, l.Url, l.AttachmentID, l.IsRevealed, true)}
			}
		}
		if visible, ok := perm.VisibleEntity(v, e); ok {
			list.Entities = append(list.Entities, visible)
		}
	}
	return list, nil
}

// readableLinks are the [[link]] targets in a player's visible fields they
// may open: images (the revealed field lets them), notes they can read and
// revealed NPCs.
func (a *App) readableLinks(ctx context.Context, v perm.Viewer, campaignID int64, fields perm.Fields) (map[string]bool, error) {
	out := map[string]bool{}
	var links []views.WikiLink
	for _, f := range fields {
		links = append(links, views.WikiLinks(f.Value)...)
	}
	if len(links) == 0 {
		return out, nil
	}
	notes, err := a.readableNotes(ctx, v, campaignID)
	if err != nil {
		return nil, err
	}
	npcs, err := a.Store.ListRevealedEntities(ctx, campaignID)
	if err != nil {
		return nil, err
	}
	for _, l := range links {
		switch l.Kind {
		case string(enum.LinkImage):
			out[l.Key()] = true
		case string(enum.LinkNote):
			out[l.Key()] = slices.ContainsFunc(notes, func(n sqlc.Note) bool { return n.ID == l.ID })
		case string(enum.LinkEntity):
			out[l.Key()] = slices.ContainsFunc(npcs, func(e sqlc.Entity) bool { return e.ID == l.ID })
		}
	}
	return out, nil
}

func toEntity(r sqlc.Entity) (perm.Entity, error) {
	fields, err := perm.ParseFields(r.FieldsJson)
	if err != nil {
		return perm.Entity{}, err
	}
	return perm.Entity{ID: r.ID, Name: r.Name, IsRevealed: r.IsRevealed, Fields: fields, CompendiumRef: r.CompendiumRef.Int64}, nil
}

func (a *App) entityView(ctx context.Context, v perm.Viewer, entityID int64) (views.EntityView, bool, error) {
	row, err := a.Store.GetEntity(ctx, entityID)
	if errors.Is(err, sql.ErrNoRows) {
		return views.EntityView{}, false, nil
	}
	if err != nil {
		return views.EntityView{}, false, err
	}
	e, err := toEntity(row)
	if err != nil {
		return views.EntityView{}, false, err
	}
	looks, err := a.Store.ListEntityImages(ctx, entityID)
	if err != nil {
		return views.EntityView{}, false, err
	}
	for _, l := range looks {
		e.Images = append(e.Images, toLook(l.ID, l.Name, l.Url, l.AttachmentID, l.IsRevealed, l.IsCurrent))
	}
	visible, ok := perm.VisibleEntity(v, e)
	if !ok {
		return views.EntityView{}, false, nil
	}
	ev := views.EntityView{CampaignID: row.CampaignID, E: visible, IsGM: v.IsGM}
	if !v.IsGM {
		if ev.Readable, err = a.readableLinks(ctx, v, row.CampaignID, visible.Fields); err != nil {
			return views.EntityView{}, false, err
		}
	}
	if v.IsGM {
		if row.CompendiumRef.Valid {
			entry, err := a.Store.GetCompendiumEntry(ctx, row.CompendiumRef.Int64)
			if err == nil {
				ev.Compendium = &entry
			}
		}
		if enc, err := a.Store.GetActiveEncounter(ctx, row.CampaignID); err == nil {
			ev.ActiveEncounterID = enc.ID
		}
		images, err := a.Store.ListImages(ctx, sqlc.ListImagesParams{CampaignID: row.CampaignID})
		if err != nil {
			return views.EntityView{}, false, err
		}
		for _, img := range images {
			if !slices.ContainsFunc(e.Images, func(l perm.Look) bool { return l.ImageID == img.ID }) {
				ev.ImageChoices = append(ev.ImageChoices, img)
			}
		}
	}
	return ev, true, nil
}
