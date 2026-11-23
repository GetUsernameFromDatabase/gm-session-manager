package views

import "testing"

func TestFileSize(t *testing.T) {
	for _, tc := range []struct {
		n    int64
		want string
	}{
		{n: 0, want: "0 B"},
		{n: 300, want: "300 B"},
		{n: 1023, want: "1023 B"},
		{n: 1024, want: "1 KB"},
		{n: 1536, want: "2 KB"},
		{n: 5 << 20, want: "5.0 MB"},
	} {
		got := fileSize(tc.n)
		if got != tc.want {
			t.Errorf("fileSize(%d) = %q, want %q", tc.n, got, tc.want)
		}
	}
}
