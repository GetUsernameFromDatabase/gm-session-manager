package perm

import (
	"encoding/json"
	"fmt"
)

// Field is one entry of entities.fields_json.
type Field struct {
	Value      string `json:"value"`
	IsRevealed bool   `json:"is_revealed"`
}

// Fields is the decoded fields_json.
type Fields map[string]Field

// ParseFields decodes fields_json.
func ParseFields(raw string) (Fields, error) {
	f := Fields{}
	if raw == "" {
		return f, nil
	}
	if err := json.Unmarshal([]byte(raw), &f); err != nil {
		return nil, fmt.Errorf("entity fields: %w", err)
	}
	return f, nil
}

// Entity is an entity as a viewer may see it.
type Entity struct {
	ID            int64
	Name          string
	IsRevealed    bool
	Fields        Fields
	CompendiumRef int64 // 0 when none or hidden
	Images        []Look
}

// Look is one of an NPC's or a character's looks: a linked image, revealed
// to players on its own. The current look is the portrait.
type Look struct {
	ImageID    int64
	Name       string
	Version    string
	IsRevealed bool
	IsCurrent  bool
}

// VisibleEntity returns what the viewer may see of e, and false if the
// viewer may not see it at all. Players get revealed entities only, with
// every unrevealed field removed entirely (key and value) and no compendium
// link, so nothing hidden is ever sent to their browser.
func VisibleEntity(v Viewer, e Entity) (Entity, bool) {
	if v.IsGM {
		return e, true
	}
	if !v.IsPlayer() || !e.IsRevealed {
		return Entity{}, false
	}
	out := Entity{ID: e.ID, Name: e.Name, IsRevealed: true, Fields: Fields{}}
	for k, f := range e.Fields {
		if f.IsRevealed {
			out.Fields[k] = f
		}
	}
	out.Images = VisibleLooks(v, 0, e.Images)
	return out, true
}

// VisibleLooks returns the looks v may see: all of them for the GM and the
// owning character's player, the revealed ones for other players. Hidden
// looks are removed entirely.
func VisibleLooks(v Viewer, owner int64, looks []Look) []Look {
	if v.IsGM || v.owns(owner) {
		return looks
	}
	var out []Look
	if !v.IsPlayer() {
		return out
	}
	for _, l := range looks {
		if l.IsRevealed {
			out = append(out, l)
		}
	}
	return out
}
