package web

import (
	"testing"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/perm"
)

// TestEntityMentionsStayInCampaign: an NPC field of another campaign that
// links a target here is not listed as a mention of it.
func TestEntityMentionsStayInCampaign(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	other, err := e.q().
		CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Other", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	target, err := e.q().
		CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: e.campaignID, Kind: "npc", Name: "Target", FieldsJson: "{}"})
	if err != nil {
		t.Fatal(err)
	}
	foreign, err := e.q().
		CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: other.ID, Kind: "npc", Name: "Foreign", FieldsJson: "{}"})
	if err != nil {
		t.Fatal(err)
	}
	link := sqlc.AddEntityFieldLinkParams{
		EntityID:       foreign.ID,
		FieldKey:       "notes",
		TargetEntityID: nullID(target.ID),
		IsRevealed:     true,
	}
	if err := e.q().AddEntityFieldLink(ctx, link); err != nil {
		t.Fatal(err)
	}
	gm := who{viewer: perm.Viewer{IsGM: true}}
	for _, tt := range []struct {
		name     string
		campaign int64
		want     int
	}{{"same campaign sees none", e.campaignID, 0}, {"its own campaign sees it", other.ID, 1}} {
		got, err := e.app.entityMentions(ctx, gm, tt.campaign, enum.LinkEntity, target.ID)
		if err != nil {
			t.Fatal(err)
		}
		if len(got) != tt.want {
			t.Errorf("%s: got %d mentions, want %d", tt.name, len(got), tt.want)
		}
	}
}
