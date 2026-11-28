package dnd5e

import (
	"reflect"
	"testing"
)

func TestText(t *testing.T) {
	got := SplitList("Insight and Religion")
	if !reflect.DeepEqual(got, []string{"Insight", "Religion"}) {
		t.Errorf("SplitList: %q", got)
	}
	if Squash("In sight") != "insight" {
		t.Error("squash")
	}
	if CountWord("Two") != 2 || CountWord("3") != 3 {
		t.Error("CountWord")
	}
}
