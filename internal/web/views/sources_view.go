package views

import (
	"fmt"

	"gm-session-manager/internal/dnd5e"
)

// SourceIssue is a sheet ref whose pinned source is gone (Missing) or
// outranked by a higher-priority copy (Best), for "Review sources".
type SourceIssue struct {
	// Kind is a dnd5e ref kind (dnd5e.RefClass …), as the source command takes it.
	Kind, Slug, Name string
	// Pinned is the sheet's source; Best is the copy priority picks, "" for none.
	Pinned, Best string
	Missing      bool
	// Copies are the sources holding the ref, best first.
	Copies []SourceCopy
}

// SourceCopy is one source holding a ref.
type SourceCopy struct {
	Key, Label string
}

// SourceLabel names a source's copy, as the "Also in" line does.
func SourceLabel(key, edition string) string { return key + commaPrefixed(edition) }

// SourcePreviewView is what switching a ref to Source does, for the dialog
// to list before Confirm posts Command to the character's command endpoint.
type SourcePreviewView struct {
	CharacterID          int64
	Kind, Slug, Source   string
	Name                 string
	SourceLabel, Command string
	Plan                 dnd5e.SwitchPlan
}

func (i SourceIssue) previewID(s SheetView) string {
	return fmt.Sprintf("source-%d-%s-%s", s.ID, i.Kind, i.Slug)
}
