package web

import (
	"context"
	"encoding/json"
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	"gm-session-manager/internal/compendium"
)

// settingsFile holds the app-wide settings in DATA_DIR, not the database:
// a few switches need no migration.
const settingsFile = "settings.json"

// settings are the app-wide switches the GM sets.
type settings struct {
	// FiveToolsAutoRedownload downloads stale 5etools books again at startup.
	FiveToolsAutoRedownload bool `json:"fivetools_auto_redownload"`
}

// appState is what the app keeps in memory between requests: the settings
// and the 5etools books downloaded before compendium.DownloadVersion.
type appState struct {
	mu       sync.Mutex
	settings settings
	stale    []compendium.Book
}

func (s *appState) get() (settings, []compendium.Book) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.settings, s.stale
}

// loadSettings reads settings.json; a missing file is the defaults.
func loadSettings(dataDir string) (settings, error) {
	var st settings
	body, err := os.ReadFile(filepath.Join(dataDir, settingsFile)) //nolint:gosec // the operator's DATA_DIR
	if errors.Is(err, fs.ErrNotExist) {
		return st, nil
	}
	if err != nil {
		return st, err
	}
	return st, json.Unmarshal(body, &st)
}

// saveSettings writes settings.json through a temp file, so a crash never
// leaves half a file.
func saveSettings(dataDir string, st settings) error {
	body, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(dataDir, ".settings-*.json")
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(tmp.Name()) }() // no-op after the rename
	if _, err := tmp.Write(body); err != nil {
		_ = tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), filepath.Join(dataDir, settingsFile)) //nolint:gosec // the operator's DATA_DIR
}

// setAutoRedownload stores the GM's switch.
func (a *App) setAutoRedownload(on bool) error {
	a.state.mu.Lock()
	defer a.state.mu.Unlock()
	st := a.state.settings
	st.FiveToolsAutoRedownload = on
	if err := saveSettings(a.DataDir, st); err != nil {
		return err
	}
	a.state.settings = st
	return nil
}

// refreshStale lists the stale books again, after a download or a removal.
func (a *App) refreshStale() error {
	stale, err := compendium.StaleFiveTools(a.compendiumDir())
	a.state.mu.Lock()
	defer a.state.mu.Unlock()
	a.state.stale = stale
	return err
}

// StartUp loads the settings and the stale books, and downloads those again
// when the GM asked for it. Failures are left for the banner to show.
func (a *App) StartUp(ctx context.Context) error {
	// A broken file must not hide the stale books: warn and use the defaults.
	st, err := loadSettings(a.DataDir)
	if err != nil {
		slog.Warn("reading "+settingsFile+", using the defaults", "err", err)
		st = settings{}
	}
	a.state.mu.Lock()
	a.state.settings = st
	a.state.mu.Unlock()
	if err := a.refreshStale(); err != nil {
		return err
	}
	_, stale := a.state.get()
	if st.FiveToolsAutoRedownload && len(stale) > 0 && a.fiveTools.start(len(stale)) {
		go a.runFiveTools(context.WithoutCancel(ctx), a.compendiumDir(), stale)
	}
	return nil
}
