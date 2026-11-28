package deploy

import (
	"os"
	"regexp"
	"testing"

	"gm-session-manager/internal/config"
	"gm-session-manager/internal/web"
)

// The scripts name in common.sh the files the Go app also reads or writes in DATA_DIR; a rename on
// one side only
// would back up, or show in the footer, a file that no longer exists.
func TestNamesMatchGo(t *testing.T) {
	src, err := os.ReadFile("common.sh")
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{}
	for _, m := range regexp.MustCompile(`(?m)^([A-Z_]+)=(\S+)$`).FindAllStringSubmatch(string(src), -1) {
		vars[m[1]] = m[2]
	}
	for name, goValue := range map[string]string{
		"DB_FILE":       config.DatabaseFile,
		"STATUS_FILE":   web.UpdateStatusFile,
		"DOWNLOADS_DIR": web.DownloadsDir,
	} {
		if vars[name] != goValue {
			t.Errorf("common.sh %s=%q, Go has %q", name, vars[name], goValue)
		}
	}
}
