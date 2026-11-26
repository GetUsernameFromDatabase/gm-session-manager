package web

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"slices"

	"github.com/gin-gonic/gin"
)

// DownloadsDir is the folder under DATA_DIR whose files the GM's server hands out (the app for taking a character home).
const DownloadsDir = "downloads"

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
