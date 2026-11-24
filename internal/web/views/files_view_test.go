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
		{n: 1<<20 - 1<<9 - 1, want: "1023 KB"},
		{n: 1<<20 - 1<<9, want: "1.0 MB"},
		{n: 5 << 20, want: "5.0 MB"},
		{n: 1<<30 - 1<<20/20 - 1, want: "1023.9 MB"},
		{n: 1<<30 - 1<<20/20, want: "1.0 GB"},
		{n: 3 << 30, want: "3.0 GB"},
	} {
		got := fileSize(tc.n)
		if got != tc.want {
			t.Errorf("fileSize(%d) = %q, want %q", tc.n, got, tc.want)
		}
	}
}
