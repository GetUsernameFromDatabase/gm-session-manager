package perm

import (
	"slices"
	"testing"

	"gm-session-manager/internal/enum"
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
	{"gm", Block{Visibility: enum.VisGM}, []string{"gm"}, []string{"gm"}, []string{"gm"}},
	{
		"players read-only",
		Block{Visibility: enum.VisPlayers},
		[]string{"gm", "alice", "bob"},
		[]string{"gm"},
		[]string{"gm"},
	},
	{
		"players editable",
		Block{Visibility: enum.VisPlayers, PlayersCanEdit: true},
		[]string{"gm", "alice", "bob"},
		[]string{"gm", "alice", "bob"},
		[]string{"gm"},
	},
	{
		"listed alice",
		Block{Visibility: enum.VisListed, Audience: []int64{alice}},
		[]string{"gm", "alice"},
		[]string{"gm"},
		[]string{"gm"},
	},
	{
		"listed alice, editable flag ignored",
		Block{Visibility: enum.VisListed, PlayersCanEdit: true, Audience: []int64{alice}},
		[]string{"gm", "alice"},
		[]string{"gm"},
		[]string{"gm"},
	},
	{
		"listed empty fails closed",
		Block{Visibility: enum.VisListed},
		[]string{"gm"},
		[]string{"gm"},
		[]string{"gm"},
	},
	{
		"owner alice",
		Block{Visibility: enum.VisOwner, NoteOwnerID: alice},
		[]string{"alice"},
		[]string{"alice"},
		[]string{"alice"},
	},
	{
		"players block in alice's note",
		Block{Visibility: enum.VisPlayers, NoteOwnerID: alice},
		[]string{"gm", "alice", "bob"},
		[]string{"gm", "alice"},
		[]string{"alice"},
	},
	{
		"players block alice wrote in gm note",
		Block{Visibility: enum.VisPlayers, AuthorID: alice},
		[]string{"gm", "alice", "bob"},
		[]string{"gm", "alice"},
		[]string{"gm"},
	},
	{
		"alice's block the gm hid",
		Block{Visibility: enum.VisGM, AuthorID: alice},
		[]string{"gm"},
		[]string{"gm"},
		[]string{"gm"},
	},
	{
		"alice's block made secret for alice",
		Block{Visibility: enum.VisListed, AuthorID: alice, Audience: []int64{alice}},
		[]string{"gm", "alice"},
		[]string{"gm", "alice"},
		[]string{"gm"},
	},
	{
		"alice's block made secret for bob",
		Block{Visibility: enum.VisListed, AuthorID: alice, Audience: []int64{bob}},
		[]string{"gm", "bob"},
		[]string{"gm"},
		[]string{"gm"},
	},
	{
		"owner with no owner reads as nobody",
		Block{Visibility: enum.VisOwner},
		nil,
		nil,
		[]string{"gm"},
	},
	{"unknown visibility", Block{Visibility: "public"}, nil, nil, []string{"gm"}},
}

func TestBlockRules(t *testing.T) {
	for _, tt := range blockCases {
		for _, who := range viewerOf {
			v := viewers[who]
			t.Run(tt.name+"/"+who, func(t *testing.T) {
				got := CanReadBlock(v, tt.block)
				want := slices.Contains(tt.readers, who)
				if got != want {
					t.Errorf("read = %v, want %v", got, want)
				}
				got = CanEditBlock(v, tt.block)
				want = slices.Contains(tt.editors, who)
				if got != want {
					t.Errorf("edit = %v, want %v", got, want)
				}
				got = CanManageNote(v, tt.block.NoteOwnerID)
				want = slices.Contains(tt.managers, who)
				if got != want {
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
		{"gm block in gm note", Block{Visibility: enum.VisPlayers}, []string{"gm"}},
		{
			"alice's block in gm note",
			Block{Visibility: enum.VisPlayers, AuthorID: alice},
			[]string{"gm", "alice"},
		},
		{
			"alice's block the gm hid",
			Block{Visibility: enum.VisGM, AuthorID: alice},
			[]string{"gm"},
		},
		{
			"alice's block for listed bob",
			Block{Visibility: enum.VisListed, AuthorID: alice, Audience: []int64{bob}},
			[]string{"gm"},
		},
		{
			"players block in alice's note",
			Block{Visibility: enum.VisPlayers, NoteOwnerID: alice},
			[]string{"alice"},
		},
	}
	for _, tt := range tests {
		for _, who := range viewerOf {
			t.Run(tt.name+"/"+who, func(t *testing.T) {
				got := CanManageBlock(viewers[who], tt.block)
				want := slices.Contains(tt.managers, who)
				if got != want {
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
		allowed     []enum.Visibility
	}{
		{
			"gm in gm note",
			GM,
			0,
			false,
			[]enum.Visibility{enum.VisGM, enum.VisPlayers, enum.VisListed},
		},
		{
			"gm in gm game session",
			GM,
			0,
			true,
			[]enum.Visibility{enum.VisGM, enum.VisPlayers, enum.VisListed},
		},
		{"gm in player note", GM, alice, false, nil},
		{
			"owner in own note",
			asAlice,
			alice,
			false,
			[]enum.Visibility{enum.VisOwner, enum.VisPlayers},
		},
		{"other player in note", asBob, alice, false, nil},
		{"other player in player game session", asBob, alice, true, nil},
		{"player in gm note", asAlice, 0, false, nil},
		{"player in gm game session", asAlice, 0, true, []enum.Visibility{enum.VisPlayers}},
		{"anon", anon, 0, false, nil},
		{"anon in gm game session", anon, 0, true, nil},
	}
	all := []enum.Visibility{enum.VisGM, enum.VisPlayers, enum.VisListed, enum.VisOwner, "public"}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlockVisibilities(
				tt.v,
				tt.owner,
				tt.gameSession,
			); !slices.Equal(
				got,
				tt.allowed,
			) {
				t.Errorf("got %v, want %v", got, tt.allowed)
			}
			for _, vis := range all {
				got := ValidBlockVisibility(tt.v, tt.owner, tt.gameSession, vis)
				want := slices.Contains(tt.allowed, vis)
				if got != want {
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
		{"own sources", asAlice, alice, ScopeSources, true},
		{"other sources", asBob, alice, ScopeSources, false},
		{"other view", asBob, alice, ScopeView, false},
		{"other hp", asBob, alice, ScopeHP, false},
		{"other sheet", asBob, alice, ScopeSheet, false},
		{"gm view", GM, alice, ScopeView, true},
		{"gm hp", GM, alice, ScopeHP, true},
		{"gm conditions", GM, alice, ScopeConditions, true},
		{"gm sources", GM, alice, ScopeSources, true},
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

func TestCanUsePreset(t *testing.T) {
	tests := []struct {
		name  string
		v     Viewer
		owner int64
		want  bool
	}{
		{"gm, gm preset", GM, 0, true},
		{"gm, player preset", GM, alice, false},
		{"owner player", asAlice, alice, true},
		{"other player", asBob, alice, false},
		{"player, gm preset", asAlice, 0, false},
		{"visitor, gm preset", anon, 0, false},
		{"visitor, player preset", anon, alice, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanUsePreset(tt.v, tt.owner); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanAddEffect(t *testing.T) {
	tests := []struct {
		name string
		v    Viewer
		same bool
		want bool
	}{
		{"gm", GM, false, true},
		{"owner", asAlice, false, true},
		{"ally player, same campaign", asBob, true, true},
		{"player, other campaign", asBob, false, false},
		{"visitor, same campaign", anon, true, false},
		{"visitor", anon, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanAddEffect(tt.v, alice, tt.same); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanRemoveCondition(t *testing.T) {
	cass := int64(3)
	tests := []struct {
		name      string
		v         Viewer
		appliedBy int64
		want      bool
	}{
		{"gm", GM, 0, true},
		{"target's player", asAlice, 0, true},
		{"applier", asBob, bob, true},
		{"third player", Player(cass), bob, false},
		{"third player, nobody applied", Player(cass), 0, false},
		{"applier-less never grants", asBob, 0, false},
		{"visitor", anon, bob, false},
		{"visitor, nobody applied", anon, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanRemoveCondition(tt.v, alice, tt.appliedBy); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanChangeSubclass(t *testing.T) {
	tests := []struct {
		name string
		v    Viewer
		want bool
	}{
		{"owner", asAlice, true},
		{"other player", asBob, false},
		{"gm", GM, false},
		{"anon", anon, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanChangeSubclass(tt.v, alice); got != tt.want {
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
	secret := &Block{Visibility: enum.VisListed, Audience: []int64{alice}}
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
		{"owner block gm", GM, &Block{Visibility: enum.VisOwner, NoteOwnerID: alice}, false},
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
		{
			"gm sees hidden entity",
			GM,
			hidden,
			true,
			[]string{"motivation", "role"},
			42,
			[]int64{1, 2},
		},
		{
			"player sees revealed fields and looks only",
			asAlice,
			npc,
			true,
			[]string{"role"},
			0,
			[]int64{2},
		},
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
	players := Block{Visibility: enum.VisPlayers}
	forAlice := Block{Visibility: enum.VisListed, Audience: []int64{alice}}
	gmOnly := Block{Visibility: enum.VisGM}
	tests := []struct {
		name    string
		blocks  []Block
		look    ImageLook
		readers []string
	}{
		{"unlinked is the gm's", nil, ImageLook{}, []string{"gm"}},
		{"linked from a gm block", []Block{gmOnly}, ImageLook{}, []string{"gm"}},
		{
			"linked from a players block",
			[]Block{players},
			ImageLook{},
			[]string{"gm", "alice", "bob"},
		},
		{"linked from a block for alice", []Block{forAlice}, ImageLook{}, []string{"gm", "alice"}},
		{
			"any readable link is enough",
			[]Block{gmOnly, forAlice},
			ImageLook{},
			[]string{"gm", "alice"},
		},
		{"revealed look", nil, ImageLook{Revealed: true}, []string{"gm", "alice", "bob"}},
		{"hidden look of alice's character", nil, ImageLook{Owner: alice}, []string{"gm", "alice"}},
		{
			"revealed look of alice's character",
			nil,
			ImageLook{Revealed: true, Owner: alice},
			[]string{"gm", "alice", "bob"},
		},
	}
	for _, tt := range tests {
		for _, who := range viewerOf {
			t.Run(tt.name+"/"+who, func(t *testing.T) {
				got := CanReadImage(viewers[who], tt.blocks, tt.look)
				want := slices.Contains(tt.readers, who)
				if got != want {
					t.Errorf("read = %v, want %v", got, want)
				}
			})
		}
	}
}

func TestCanTakeHome(t *testing.T) {
	tests := []struct {
		name string
		v    Viewer
		want bool
	}{
		{"gm", GM, true},
		{"owner", asAlice, true},
		{"other player", asBob, false},
		{"visitor", anon, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanTakeHome(tt.v, alice); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// At home (the player app) only private blocks of the player's own notes can
// be changed; everything shared is read-only, even what the player wrote.
func TestAtHomeBlockRules(t *testing.T) {
	home := Viewer{CharacterID: alice, AtHome: true}
	for _, vis := range []enum.Visibility{enum.VisGM, enum.VisPlayers, enum.VisListed, enum.VisOwner} {
		for _, noteOwner := range []int64{alice, bob, 0} {
			for _, authored := range []bool{false, true} {
				b := Block{
					Visibility:     vis,
					NoteOwnerID:    noteOwner,
					PlayersCanEdit: true,
					Audience:       []int64{alice},
				}
				if authored {
					b.AuthorID = alice
				}
				want := vis == enum.VisOwner && noteOwner == alice
				if got := CanEditBlock(home, b); got != want {
					t.Errorf(
						"%s note %d authored %v: edit = %v, want %v",
						vis,
						noteOwner,
						authored,
						got,
						want,
					)
				}
				if got := CanManageBlock(home, b); got != want {
					t.Errorf(
						"%s note %d authored %v: manage = %v, want %v",
						vis,
						noteOwner,
						authored,
						got,
						want,
					)
				}
			}
		}
	}
}

func TestAtHomeBlockVisibilities(t *testing.T) {
	home := Viewer{CharacterID: alice, AtHome: true}
	for _, gameSession := range []bool{false, true} {
		if got := BlockVisibilities(
			home,
			alice,
			gameSession,
		); !slices.Equal(
			got,
			[]enum.Visibility{enum.VisOwner},
		) {
			t.Errorf("own note: %v", got)
		}
		if got := BlockVisibilities(home, bob, gameSession); got != nil {
			t.Errorf("other's note: %v", got)
		}
		if got := BlockVisibilities(home, 0, gameSession); got != nil {
			t.Errorf("gm note: %v", got)
		}
	}
}

func TestLocksShared(t *testing.T) {
	for _, v := range []Viewer{GM, asAlice, {CharacterID: alice, AtHome: true}} {
		for _, has := range []bool{false, true} {
			if got, want := LocksShared(v, has), v.AtHome && has; got != want {
				t.Errorf("%+v has %v: %v, want %v", v, has, got, want)
			}
		}
	}
}

func TestCanAddSources(t *testing.T) {
	home := Player(alice)
	home.AtHome = true
	tests := []struct {
		name string
		v    Viewer
		want bool
	}{
		{"gm", GM, true},
		{"table player", asAlice, false},
		{"home player", home, true},
		{"visitor", anon, false},
	}
	for _, tc := range tests {
		if got := CanAddSources(tc.v); got != tc.want {
			t.Errorf("%s: CanAddSources = %v, want %v", tc.name, got, tc.want)
		}
	}
}
