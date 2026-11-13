package dnd5e

import (
	"reflect"
	"testing"
)

func TestText(t *testing.T) {
	if got := SplitList("Insight and Religion"); !reflect.DeepEqual(got, []string{"Insight", "Religion"}) {
		t.Errorf("SplitList: %q", got)
	}
	if Squash("In sight") != "insight" {
		t.Error("squash")
	}
	if CountWord("Two") != 2 || CountWord("3") != 3 {
		t.Error("CountWord")
	}
}
