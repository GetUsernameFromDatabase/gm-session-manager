package views

import (
	"cmp"
	"fmt"
	"strings"
)

// CampaignSourcesID is the Sources card on the campaign page, which every
// toggle redraws.
const CampaignSourcesID = "campaign-sources"

// CampaignSource is one imported source on the campaign's Sources card.
type CampaignSource struct {
	Key, Edition, Attribution string
	On                        bool
	// Users are the characters and NPCs in the campaign that use it; a used
	// source cannot be unticked.
	Users []string
}

// CampaignSourcesView is the campaign's allow list, in priority order.
// Homebrew is always allowed and not listed.
type CampaignSourcesView struct {
	CampaignID int64
	Sources    []CampaignSource
}

// sourceGroup is the sources of one edition.
type sourceGroup struct {
	Title   string
	Sources []CampaignSource
}

// groups splits the sources by edition, in order of first appearance.
func (v CampaignSourcesView) groups() []sourceGroup {
	var out []sourceGroup
	at := map[string]int{}
	for _, s := range v.Sources {
		i, ok := at[s.Edition]
		if !ok {
			i, at[s.Edition] = len(out), len(out)
			out = append(out, sourceGroup{Title: cmp.Or(s.Edition, labelOther)})
		}
		out[i].Sources = append(out[i].Sources, s)
	}
	return out
}

func (v CampaignSourcesView) onCount() int {
	n := 0
	for _, s := range v.Sources {
		if s.On {
			n++
		}
	}
	return n
}

// groupKeys are the group's keys the all (on) or none (off) button changes:
// those not already so, and for none only unused ones.
func (g sourceGroup) groupKeys(on bool) []string {
	var keys []string
	for _, s := range g.Sources {
		if s.On != on && (on || len(s.Users) == 0) {
			keys = append(keys, s.Key)
		}
	}
	return keys
}

func (v CampaignSourcesView) postURL() string {
	return fmt.Sprintf("/campaigns/%d/sources", v.CampaignID)
}

func sourceRowTitle(s CampaignSource, locked bool) string {
	if locked {
		return s.Key + " is used by " + strings.Join(s.Users, ", ") + "; switch their sources first."
	}
	return cmp.Or(s.Attribution, s.Key)
}
