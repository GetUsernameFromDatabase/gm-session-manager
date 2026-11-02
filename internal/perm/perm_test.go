package perm

import (
	"slices"
	"testing"
)

const (
	alice = int64(1)
	bob   = int64(2)
)

var (
	anon     = Viewer{}
	asAlice  = Player(alice)
	asBob    = Player(bob)
	viewers  = map[string]Viewer{"gm": GM, "alice": asAlice, "bob": asBob, "anon": anon}
	viewerOf = []string{"gm", "alice", "bob", "anon"}
)

// blockCase lists, per block, which viewers may read and edit it.
type blockCase struct {
	name     string
	block    Block
	readers  []string
	editors  []string
	managers []string
}

// blockCases is the note-block visibility table in CLAUDE.md (Permissions). The SQL
// agreement test runs the same cases against the readable-block queries.
var blockCases = []blockCase{
	{"gm", Block{Visibility: VisGM}, []string{"gm"}, []string{"gm"}, []string{"gm"}},
	{"players read-only", Block{Visibility: VisPlayers}, []string{"gm", "alice", "bob"}, []string{"gm"}, []string{"gm"}},
	{"players editable", Block{Visibility: VisPlayers, PlayersCanEdit: true}, []string{"gm", "alice", "bob"}, []string{"gm", "alice", "bob"}, []string{"gm"}},
	{"listed alice", Block{Visibility: VisListed, Audience: []int64{alice}}, []string{"gm", "alice"}, []string{"gm"}, []string{"gm"}},
	{"listed alice, editable flag ignored", Block{Visibility: VisListed, PlayersCanEdit: true, Audience: []int64{alice}}, []string{"gm", "alice"}, []string{"gm"}, []string{"gm"}},
	{"listed empty fails closed", Block{Visibility: VisListed}, []string{"gm"}, []string{"gm"}, []string{"gm"}},
	{"owner alice", Block{Visibility: VisOwner, NoteOwnerID: alice}, []string{"alice"}, []string{"alice"}, []string{"alice"}},
	{"players block in alice's note", Block{Visibility: VisPlayers, NoteOwnerID: alice}, []string{"gm", "alice", "bob"}, []string{"gm", "alice"}, []string{"alice"}},
	{"players block alice wrote in gm note", Block{Visibility: VisPlayers, AuthorID: alice}, []string{"gm", "alice", "bob"}, []string{"gm", "alice"}, []string{"gm"}},
	{"alice's block the gm hid", Block{Visibility: VisGM, AuthorID: alice}, []string{"gm"}, []string{"gm"}, []string{"gm"}},
	{"alice's block made secret for alice", Block{Visibility: VisListed, AuthorID: alice, Audience: []int64{alice}}, []string{"gm", "alice"}, []string{"gm", "alice"}, []string{"gm"}},
	{"alice's block made secret for bob", Block{Visibility: VisListed, AuthorID: alice, Audience: []int64{bob}}, []string{"gm", "bob"}, []string{"gm"}, []string{"gm"}},
	{"owner with no owner reads as nobody", Block{Visibility: VisOwner}, nil, nil, []string{"gm"}},
	{"unknown visibility", Block{Visibility: "public"}, nil, nil, []string{"gm"}},
}

func TestBlockRules(t *testing.T) {
	for _, tt := range blockCases {
		for _, who := range viewerOf {
			v := viewers[who]
			t.Run(tt.name+"/"+who, func(t *testing.T) {
				if got, want := CanReadBlock(v, tt.block), slices.Contains(tt.readers, who); got != want {
					t.Errorf("read = %v, want %v", got, want)
				}
				if got, want := CanEditBlock(v, tt.block), slices.Contains(tt.editors, who); got != want {
					t.Errorf("edit = %v, want %v", got, want)
				}
				if got, want := CanManageNote(v, tt.block.NoteOwnerID), slices.Contains(tt.managers, who); got != want {
					t.Errorf("manage = %v, want %v", got, want)
				}
				if CanEditBlock(v, tt.block) && !CanReadBlock(v, tt.block) {
					t.Error("can edit a block it cannot read")
				}
			})
		}
	}
}

func TestCanManageBlock(t *testing.T) {
	tests := []struct {
		name     string
		block    Block
		managers []string
	}{
		{"gm block in gm note", Block{Visibility: VisPlayers}, []string{"gm"}},
		{"alice's block in gm note", Block{Visibility: VisPlayers, AuthorID: alice}, []string{"gm", "alice"}},
		{"alice's block the gm hid", Block{Visibility: VisGM, AuthorID: alice}, []string{"gm"}},
		{"alice's block for listed bob", Block{Visibility: VisListed, AuthorID: alice, Audience: []int64{bob}}, []string{"gm"}},
		{"players block in alice's note", Block{Visibility: VisPlayers, NoteOwnerID: alice}, []string{"alice"}},
	}
	for _, tt := range tests {
		for _, who := range viewerOf {
			t.Run(tt.name+"/"+who, func(t *testing.T) {
				if got, want := CanManageBlock(viewers[who], tt.block), slices.Contains(tt.managers, who); got != want {
					t.Errorf("manage = %v, want %v", got, want)
				}
			})
		}
	}
}

func TestBlockVisibilities(t *testing.T) {
	tests := []struct {
		name        string
		v           Viewer
		owner       int64
		gameSession bool
		allowed     []Visibility
	}{
		{"gm in gm note", GM, 0, false, []Visibility{VisGM, VisPlayers, VisListed}},
		{"gm in gm game session", GM, 0, true, []Visibility{VisGM, VisPlayers, VisListed}},
		{"gm in player note", GM, alice, false, nil},
		{"owner in own note", asAlice, alice, false, []Visibility{VisOwner, VisPlayers}},
		{"other player in note", asBob, alice, false, nil},
		{"other player in player game session", asBob, alice, true, nil},
		{"player in gm note", asAlice, 0, false, nil},
		{"player in gm game session", asAlice, 0, true, []Visibility{VisPlayers}},
		{"anon", anon, 0, false, nil},
		{"anon in gm game session", anon, 0, true, nil},
	}
	all := []Visibility{VisGM, VisPlayers, VisListed, VisOwner, "public"}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlockVisibilities(tt.v, tt.owner, tt.gameSession); !slices.Equal(got, tt.allowed) {
				t.Errorf("got %v, want %v", got, tt.allowed)
			}
			for _, vis := range all {
				if got, want := ValidBlockVisibility(tt.v, tt.owner, tt.gameSession, vis), slices.Contains(tt.allowed, vis); got != want {
					t.Errorf("%s: valid = %v, want %v", vis, got, want)
				}
			}
		})
	}
}

func TestCanCharacter(t *testing.T) {
	tests := []struct {
		name  string
		v     Viewer
		char  int64
		scope CharacterScope
		want  bool
	}{
		{"own view", asAlice, alice, ScopeView, true},
		{"own hp", asAlice, alice, ScopeHP, true},
		{"own conditions", asAlice, alice, ScopeConditions, true},
		{"own sheet", asAlice, alice, ScopeSheet, true},
		{"other view", asBob, alice, ScopeView, false},
		{"other hp", asBob, alice, ScopeHP, false},
		{"other sheet", asBob, alice, ScopeSheet, false},
		{"gm view", GM, alice, ScopeView, true},
		{"gm hp", GM, alice, ScopeHP, true},
		{"gm conditions", GM, alice, ScopeConditions, true},
		{"gm sheet", GM, alice, ScopeSheet, false},
		{"anon view", anon, alice, ScopeView, false},
		{"anon with zero character", anon, 0, ScopeView, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanCharacter(tt.v, tt.char, tt.scope); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanDeleteAndHomebrew(t *testing.T) {
	tests := []struct {
		name      string
		v         Viewer
		char      int64
		del, brew bool
	}{
		{"own", asAlice, alice, true, true},
		{"other player's", asBob, alice, false, false},
		{"gm", GM, alice, true, true},
		{"gm, no owner", GM, 0, true, true},
		{"player, no owner", asAlice, 0, false, false},
		{"anon", anon, alice, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanDeleteCharacter(tt.v, tt.char); got != tt.del {
				t.Errorf("delete: got %v", got)
			}
			if got := CanEditHomebrew(tt.v, tt.char); got != tt.brew {
				t.Errorf("homebrew: got %v", got)
			}
		})
	}
}

func TestCanSeeEntry(t *testing.T) {
	tests := []struct {
		name   string
		scoped bool
		linked []int64
		want   bool
	}{
		{"imported or gm homebrew", false, nil, true},
		{"player homebrew, this campaign", true, []int64{1, 7}, true},
		{"player homebrew, other campaign", true, []int64{2}, false},
		{"player homebrew, unlinked everywhere", true, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanSeeEntry(7, tt.scoped, tt.linked); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanSeeCombatantHP(t *testing.T) {
	tests := []struct {
		name string
		v    Viewer
		c    Combatant
		want bool
	}{
		{"gm monster hidden", GM, Combatant{}, true},
		{"player monster hidden", asAlice, Combatant{}, false},
		{"player monster visible", asAlice, Combatant{HPVisible: true}, true},
		{"player own pc", asAlice, Combatant{CharacterID: alice}, true},
		{"player other pc", asAlice, Combatant{CharacterID: bob}, false},
		{"player other pc visible", asAlice, Combatant{CharacterID: bob, HPVisible: true}, true},
		{"anon visible", anon, Combatant{HPVisible: true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanSeeCombatantHP(tt.v, tt.c); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanReadAttachment(t *testing.T) {
	secret := &Block{Visibility: VisListed, Audience: []int64{alice}}
	tests := []struct {
		name  string
		v     Viewer
		block *Block
		want  bool
	}{
		{"gm no block", GM, nil, true},
		{"player no block", asAlice, nil, false},
		{"listed member", asAlice, secret, true},
		{"listed outsider", asBob, secret, false},
		{"owner block gm", GM, &Block{Visibility: VisOwner, NoteOwnerID: alice}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanReadAttachment(tt.v, tt.block); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVisibleEntity(t *testing.T) {
	npc := Entity{
		ID: 7, Name: "Mira", IsRevealed: true, CompendiumRef: 42,
		Fields: Fields{
			"role":       {Value: "Innkeeper", IsRevealed: true},
			"motivation": {Value: "Revenge", IsRevealed: false},
		},
		Images: []Look{{ImageID: 1, IsCurrent: true}, {ImageID: 2, IsRevealed: true}},
	}
	hidden := npc
	hidden.IsRevealed = false

	tests := []struct {
		name       string
		v          Viewer
		e          Entity
		wantOK     bool
		wantFields []string
		wantRef    int64
		wantImages []int64
	}{
		{"gm sees all", GM, npc, true, []string{"motivation", "role"}, 42, []int64{1, 2}},
		{"gm sees hidden entity", GM, hidden, true, []string{"motivation", "role"}, 42, []int64{1, 2}},
		{"player sees revealed fields and looks only", asAlice, npc, true, []string{"role"}, 0, []int64{2}},
		{"player cannot see hidden entity", asAlice, hidden, false, nil, 0, nil},
		{"anon sees nothing", anon, npc, false, nil, 0, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := VisibleEntity(tt.v, tt.e)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
			}
			keys := make([]string, 0, len(got.Fields))
			for k := range got.Fields {
				keys = append(keys, k)
			}
			slices.Sort(keys)
			if !slices.Equal(keys, tt.wantFields) {
				t.Errorf("fields %v, want %v", keys, tt.wantFields)
			}
			var images []int64
			for _, img := range got.Images {
				images = append(images, img.ImageID)
			}
			if !slices.Equal(images, tt.wantImages) {
				t.Errorf("images %v, want %v", images, tt.wantImages)
			}
			if got.CompendiumRef != tt.wantRef {
				t.Errorf("compendium ref %d, want %d", got.CompendiumRef, tt.wantRef)
			}
			if !ok && (got.Name != "" || got.ID != 0) {
				t.Error("hidden entity leaked data")
			}
		})
	}
	if _, ok := npc.Fields["motivation"]; !ok {
		t.Error("VisibleEntity modified its input")
	}
}

func TestVisibleLooks(t *testing.T) {
	looks := []Look{{ImageID: 1, IsRevealed: true, IsCurrent: true}, {ImageID: 2}}
	want := map[string][]int64{"gm": {1, 2}, "alice": {1, 2}, "bob": {1}, "anon": nil}
	for _, who := range viewerOf {
		t.Run(who, func(t *testing.T) {
			var got []int64
			for _, l := range VisibleLooks(viewers[who], alice, looks) {
				got = append(got, l.ImageID)
			}
			if !slices.Equal(got, want[who]) {
				t.Errorf("looks %v, want %v", got, want[who])
			}
		})
	}
}

func TestCanReadImage(t *testing.T) {
	players := Block{Visibility: VisPlayers}
	forAlice := Block{Visibility: VisListed, Audience: []int64{alice}}
	gmOnly := Block{Visibility: VisGM}
	tests := []struct {
		name    string
		blocks  []Block
		look    ImageLook
		readers []string
	}{
		{"unlinked is the gm's", nil, ImageLook{}, []string{"gm"}},
		{"linked from a gm block", []Block{gmOnly}, ImageLook{}, []string{"gm"}},
		{"linked from a players block", []Block{players}, ImageLook{}, []string{"gm", "alice", "bob"}},
		{"linked from a block for alice", []Block{forAlice}, ImageLook{}, []string{"gm", "alice"}},
		{"any readable link is enough", []Block{gmOnly, forAlice}, ImageLook{}, []string{"gm", "alice"}},
		{"revealed look", nil, ImageLook{Revealed: true}, []string{"gm", "alice", "bob"}},
		{"hidden look of alice's character", nil, ImageLook{Owner: alice}, []string{"gm", "alice"}},
		{"revealed look of alice's character", nil, ImageLook{Revealed: true, Owner: alice}, []string{"gm", "alice", "bob"}},
	}
	for _, tt := range tests {
		for _, who := range viewerOf {
			t.Run(tt.name+"/"+who, func(t *testing.T) {
				if got, want := CanReadImage(viewers[who], tt.blocks, tt.look), slices.Contains(tt.readers, who); got != want {
					t.Errorf("read = %v, want %v", got, want)
				}
			})
		}
	}
}
