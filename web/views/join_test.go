package views

import "testing"

func TestFingerprintPairs(t *testing.T) {
	for in, want := range map[string]string{
		"":       "",
		"c36ae2": "C3 6A E2",
		"abc":    "AB C",
	} {
		if got := fingerprintPairs(in); got != want {
			t.Errorf("fingerprintPairs(%q) = %q, want %q", in, got, want)
		}
	}
}
