package web

import (
	"cmp"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"

	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
	"gm-session-manager/hub"
	"gm-session-manager/perm"
	"gm-session-manager/web/views"
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

// pendingViews groups every campaign's waiting requests, campaigns in the
// order of their oldest request.
func (a *App) pendingViews() []views.PendingGroup {
	var out []views.PendingGroup
	for _, p := range a.Hub.PendingList(0) {
		i := slices.IndexFunc(out, func(g views.PendingGroup) bool { return g.CampaignID == p.CampaignID })
		if i < 0 {
			i = len(out)
			out = append(out, views.PendingGroup{CampaignID: p.CampaignID, CampaignName: p.CampaignName})
		}
		out[i].Requests = append(out[i].Requests, views.PendingView{ID: p.ID, CampaignID: p.CampaignID, Name: p.Name, Claim: p.Kind == hub.PendingClaim,
			Player: p.Player, OS: p.OS, Phone: p.Phone, Host: p.Host,
		})
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
	for _, ch := range chars {
		rc := views.RosterCharacter{Character: ch}
		for _, s := range sessions {
			if s.CharacterID.Int64 == ch.ID {
				ua, on := online[s.ID]
				rc.Devices = append(rc.Devices, views.Device{SessionID: s.ID, UserAgent: ua, Online: on})
			}
		}
		r.Characters = append(r.Characters, rc)
	}
	return r, nil
}

func (a *App) sheetView(ctx context.Context, v perm.Viewer, row sqlc.Character) (views.SheetView, error) {
	sheet, err := dnd5e.ParseSheet(row.SheetJson)
	if err != nil {
		return views.SheetView{}, err
	}
	conds, err := a.Store.ListConditionsByCharacter(ctx, nullID(row.ID))
	if err != nil {
		return views.SheetView{}, err
	}
	rows, err := a.Store.ListCharacterImages(ctx, row.ID)
	if err != nil {
		return views.SheetView{}, err
	}
	var looks []perm.Look
	for _, l := range rows {
		looks = append(looks, toLook(l.ID, l.Name, l.Url, l.AttachmentID, l.IsRevealed, l.IsCurrent))
	}
	spells, err := a.sheetSpells(ctx, sheet)
	if err != nil {
		return views.SheetView{}, err
	}
	levelUps, err := a.Store.CountLevelUps(ctx, row.ID)
	if err != nil {
		return views.SheetView{}, err
	}
	canSheet := perm.CanCharacter(v, row.ID, perm.ScopeSheet)
	return views.SheetView{
		ID: row.ID, CampaignID: row.CampaignID, Conditions: conds, Looks: perm.VisibleLooks(v, row.ID, looks), SpellData: spells,
		C:             dnd5e.Character{Name: row.Name, HP: int(row.Hp), MaxHP: int(row.MaxHp), TempHP: int(row.TempHp), Sheet: sheet},
		CanHP:         perm.CanCharacter(v, row.ID, perm.ScopeHP),
		CanConditions: perm.CanCharacter(v, row.ID, perm.ScopeConditions),
		CanSheet:      canSheet,
		CanRevert:     canSheet && levelUps > 0,
	}, nil
}

// sheetSpells are the compendium details of the sheet's spells, by slug.
func (a *App) sheetSpells(ctx context.Context, sheet dnd5e.Sheet) (map[string]*compendium.SpellData, error) {
	if len(sheet.Spells) == 0 {
		return nil, nil
	}
	slugs := make([]string, len(sheet.Spells))
	for i, sp := range sheet.Spells {
		slugs[i] = sp.Slug
	}
	rows, err := a.Store.ListCompendiumBySlugs(ctx, sqlc.ListCompendiumBySlugsParams{Kind: string(compendium.KindSpell), Slugs: slugs})
	if err != nil {
		return nil, err
	}
	out := map[string]*compendium.SpellData{}
	for _, r := range rows {
		if _, seen := out[r.Slug]; seen {
			continue // rows come first-choice first, as GetCompendiumEntryBySlug picks
		}
		d, err := compendium.DecodeData(compendium.KindSpell, r.DataJson)
		if err != nil {
			return nil, err
		}
		out[r.Slug] = d.(*compendium.SpellData)
	}
	return out, nil
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
		cv := views.CombatantView{
			ID: r.ID, Name: r.Name, AC: r.Ac, HPVisible: r.HpVisible, CharacterID: r.CharacterID.Int64, EntityID: r.EntityID.Int64,
			Kind:   combatantKind(r),
			IsOwn:  v.IsPlayer() && r.CharacterID.Int64 == v.CharacterID,
			Active: enc.Status == "active" && enc.ActiveCombatantID.Int64 == r.ID,
			Out:    outOfFight(r),
		}
		if s, ok := pcSheet(r); ok { // the stored AC is from when they joined
			cv.AC = int64(s.ArmorClass())
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
		if v.IsGM {
			cv.Link, cv.InitMod = combatantLink(r), initiativeMod(r)
		} else {
			cv.AC, cv.Link = 0, playerCombatantLink(v, r, cv.ShowHP)
		}
		// The compendium is open to everyone; a visible look below overrides it.
		cv.Portrait = r.MonsterImage.String
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
		if r.CharacterID.Valid {
			pc, err := a.Store.ListConditionsByCharacter(ctx, r.CharacterID)
			if err != nil {
				return views.EncounterView{}, err
			}
			cv.Conditions = pc
		} else {
			for _, c := range conds {
				if c.CombatantID.Int64 == r.ID {
					cv.Conditions = append(cv.Conditions, c)
				}
			}
		}
		ev.Combatants = append(ev.Combatants, cv)
	}
	if v.IsGM {
		if ev.Characters, err = a.Store.ListCharactersByCampaign(ctx, enc.CampaignID); err != nil {
			return ev, err
		}
		if ev.Entities, err = a.Store.ListEntitiesByCampaign(ctx, enc.CampaignID); err != nil {
			return ev, err
		}
	}
	return ev, nil
}

// combatantLink is the GM's page for a combatant's source: the character
// sheet, the NPC, or the compendium monster.
func combatantLink(r sqlc.ListCombatantsRow) string {
	switch {
	case r.CharacterID.Valid:
		return fmt.Sprintf("/characters/%d", r.CharacterID.Int64)
	case r.EntityID.Valid:
		return fmt.Sprintf("/entities/%d", r.EntityID.Int64)
	case r.MonsterRef.Valid:
		return fmt.Sprintf("/compendium/entries/%d", r.MonsterRef.Int64)
	}
	return ""
}

// playerCombatantLink is the page a player may open for a combatant: a sheet
// they may view, a revealed NPC, or a monster whose HP the GM shows.
func playerCombatantLink(v perm.Viewer, r sqlc.ListCombatantsRow, showHP bool) string {
	switch {
	case r.CharacterID.Valid:
		if perm.CanCharacter(v, r.CharacterID.Int64, perm.ScopeView) {
			return combatantLink(r)
		}
	case r.EntityID.Valid:
		if r.EntityRevealed.Bool {
			return combatantLink(r)
		}
	case showHP:
		return combatantLink(r)
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
		CampaignID: campaignID, IsGm: v.IsGM, CharacterID: v.CharacterID, MaxResults: 500,
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
	items := make([]views.NoteItem, len(notes))
	for i, n := range notes {
		items[i] = views.NoteItem{Note: n, Author: author(v, names, n.OwnerCharacterID.Int64), CanManage: perm.CanManageNote(v, n.OwnerCharacterID.Int64)}
	}
	return items, nil
}

func (a *App) characterNames(ctx context.Context, campaignID int64) (map[int64]string, error) {
	chars, err := a.Store.ListCharactersByCampaign(ctx, campaignID)
	names := make(map[int64]string, len(chars))
	for _, c := range chars {
		names[c.ID] = c.Name
	}
	return names, err
}

// author is the character who wrote something (0 for the GM), linked to
// their sheet only when v may open it.
func author(v perm.Viewer, names map[int64]string, characterID int64) views.Author {
	if characterID == 0 {
		return views.Author{Name: "GM"}
	}
	au := views.Author{Name: names[characterID]}
	if perm.CanCharacter(v, characterID, perm.ScopeView) {
		au.Href = fmt.Sprintf("/characters/%d", characterID)
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
	nv := views.NoteView{Note: note, CanManage: perm.CanManageNote(v, note.OwnerCharacterID.Int64)}
	if !nv.CanManage && len(rows) == 0 {
		return nv, false, nil
	}
	nv.Visibilities = perm.BlockVisibilities(v, note.OwnerCharacterID.Int64, note.Kind == views.KindGameSession)
	names, err := a.characterNames(ctx, note.CampaignID)
	if err != nil {
		return nv, false, err
	}
	for _, r := range rows {
		audience, err := a.Store.ListBlockAudience(ctx, r.ID)
		if err != nil {
			return nv, false, err
		}
		pb := perm.Block{
			Visibility: perm.Visibility(r.Visibility), PlayersCanEdit: r.PlayersCanEdit, NoteOwnerID: note.OwnerCharacterID.Int64,
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
	if v.IsGM && slices.Contains(nv.Visibilities, perm.VisListed) {
		if nv.Characters, err = a.Store.ListCharactersByCampaign(ctx, note.CampaignID); err != nil {
			return nv, false, err
		}
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
		case "image":
			out[l.Key()] = true
		case "note":
			out[l.Key()] = slices.ContainsFunc(notes, func(n sqlc.Note) bool { return n.ID == l.ID })
		case "entity":
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
