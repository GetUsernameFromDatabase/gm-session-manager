package web

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

// DownloadsDir is the folder under DATA_DIR whose files the GM's server hands out (the app for
// taking a character home).
const DownloadsDir = "downloads"

// UpdateStatusFile is the one line deploy/update.sh writes in DATA_DIR about the last update.
const UpdateStatusFile = "update-status"

// updateStatus reads the server's update line on every render (it changes while the app runs);
// empty off the server.
func (a *App) updateStatus() string {
	b, err := os.ReadFile(filepath.Join(a.DataDir, UpdateStatusFile))
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(b))
}

// downloadNames lists the regular files in dir; a missing dir is empty.
func downloadNames(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var names []string
	for _, e := range entries {
		if e.Type().IsRegular() {
			names = append(names, e.Name())
		}
	}
	return names, nil
}

// appDownload serves one file of the downloads folder to the GM or a player.
// Only names from the listing are served, so no path can leave the folder.
func (a *App) appDownload(c *gin.Context) error {
	w, err := a.whoami(c)
	if err != nil {
		return err
	}
	if !w.isGM() && w.character == nil {
		c.Status(http.StatusNotFound)
		return nil
	}
	dir := filepath.Join(a.DataDir, DownloadsDir)
	names, err := downloadNames(dir)
	if err != nil {
		return err
	}
	name := c.Param("file")
	if !slices.Contains(names, name) {
		c.Status(http.StatusNotFound)
		return nil
	}
	c.FileAttachment(filepath.Join(dir, name), name)
	return nil
}
