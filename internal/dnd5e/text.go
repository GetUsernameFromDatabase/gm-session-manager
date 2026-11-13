package dnd5e

import (
	"cmp"
	"strconv"
	"strings"
)

// SplitList splits "A, B, and C" (or "A and B", "A, or B") into its items.
func SplitList(s string) []string {
	var out []string
	for _, part := range strings.Split(s, ",") {
		for _, item := range strings.Split(part, " and ") {
			item = strings.TrimSpace(item)
			item = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(item, "or "), "and "))
			if item != "" {
				out = append(out, item)
			}
		}
	}
	return out
}

// Squash lowercases s and drops spaces, so names split by the source's
// PDF extraction ("In sight", "Ar mor") still match.
func Squash(s string) string {
	return strings.ToLower(strings.Join(strings.Fields(s), ""))
}

// CountWords are the spelled-out counts source texts use ("two cantrips").
var CountWords = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6"}

// CountWord reads a count written as a word ("two") or digits.
func CountWord(w string) int {
	n, _ := strconv.Atoi(cmp.Or(CountWords[strings.ToLower(w)], w))
	return n
}
