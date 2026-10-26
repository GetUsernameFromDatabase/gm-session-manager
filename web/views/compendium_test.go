package views

import (
	"testing"

	"gm-session-manager/db/sqlc"
)

// TestRowFigure checks the right-hand text and tint of each kind's row.
func TestRowFigure(t *testing.T) {
	for _, tt := range []struct {
		name            string
		row             sqlc.SearchCompendiumRow
		text, wantColor string
	}{
		{"rule category", sqlc.SearchCompendiumRow{Kind: "rule", Category: "combat"}, "combat", ""},
		{"feat category", sqlc.SearchCompendiumRow{Kind: "feat", Category: "origin"}, "origin", ""},
		{"rare item", sqlc.SearchCompendiumRow{Kind: "item", Category: "wondrous-item", Figure: "Rare"}, "wondrous item", rarityColors["rare"]},
		{"common item", sqlc.SearchCompendiumRow{Kind: "item", Category: "weapon", Figure: "common"}, "weapon", ""},
		{"item without rarity", sqlc.SearchCompendiumRow{Kind: "item", Category: "armor"}, "armor", ""},
		{"subclass", sqlc.SearchCompendiumRow{Kind: "subclass", Classes: "wizard"}, "wizard", "text-blue-700 dark:text-blue-300"},
		{"monster CR ½", sqlc.SearchCompendiumRow{Kind: "monster", Figure: "1/2"}, "CR ½", ""},
		{"monster CR ¼", sqlc.SearchCompendiumRow{Kind: "monster", Figure: "1/4"}, "CR ¼", ""},
		{"monster CR ⅛", sqlc.SearchCompendiumRow{Kind: "monster", Figure: "1/8"}, "CR ⅛", ""},
		{"monster CR 12", sqlc.SearchCompendiumRow{Kind: "monster", Figure: "12"}, "CR 12", ""},
		{"monster without CR", sqlc.SearchCompendiumRow{Kind: "monster"}, "", ""},
		{"cantrip", sqlc.SearchCompendiumRow{Kind: "spell", Figure: "0"}, "Cantrip", ""},
		{"spell level", sqlc.SearchCompendiumRow{Kind: "spell", Figure: "3"}, "Level 3", ""},
	} {
		text, color := rowFigure(tt.row)
		if text != tt.text || color != tt.wantColor {
			t.Errorf("%s: rowFigure = %q, %q; want %q, %q", tt.name, text, color, tt.text, tt.wantColor)
		}
	}
}
