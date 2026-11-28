package web

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/route"
)

func TestAppDownload(t *testing.T) {
	e := newTestEnv(t)
	dir := filepath.Join(e.app.DataDir, DownloadsDir)
	if err := os.MkdirAll(filepath.Join(dir, "sub"), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "app.zip"), []byte("zipdata"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(
		filepath.Join(e.app.DataDir, "secret.txt"),
		[]byte("no"),
		0o600,
	); err != nil {
		t.Fatal(err)
	}

	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)
	tests := []struct {
		name, file, cookie string
		want               int
	}{
		{"gm", "app.zip", gm, http.StatusOK},
		{"player", "app.zip", player, http.StatusOK},
		{"visitor", "app.zip", "", http.StatusNotFound},
		{"unknown name", "other.zip", gm, http.StatusNotFound},
		{"traversal", "..%2Fsecret.txt", gm, http.StatusNotFound},
		{"subfolder", "sub", gm, http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := e.do(
				t,
				http.MethodGet,
				strings.Replace(route.AppDownload, ":file", tt.file, 1),
				tt.cookie,
				nil,
				false,
			)
			if w.Code != tt.want {
				t.Fatalf("status %d, want %d", w.Code, tt.want)
			}
			if tt.want != http.StatusOK {
				return
			}
			if w.Body.String() != "zipdata" ||
				!strings.HasPrefix(w.Header().Get("Content-Disposition"), "attachment") {
				t.Errorf(
					"body %q, disposition %q",
					w.Body.String(),
					w.Header().Get("Content-Disposition"),
				)
			}
		})
	}
}
