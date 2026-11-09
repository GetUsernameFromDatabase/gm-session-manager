package web

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/web/views"
)

// fiveToolsJob is the GM's background 5etools import, in memory like join
// requests: one at a time, forgotten on restart (the books on disk import
// on start anyway).
type fiveToolsJob struct {
	mu  sync.Mutex
	job views.FiveToolsJob
}

func (j *fiveToolsJob) get() views.FiveToolsJob {
	j.mu.Lock()
	defer j.mu.Unlock()
	return j.job
}

// start claims the job for total books; false while one runs.
func (j *fiveToolsJob) start(total int) bool {
	j.mu.Lock()
	defer j.mu.Unlock()
	if j.job.Running {
		return false
	}
	j.job = views.FiveToolsJob{Running: true, Total: total}
	return true
}

func (j *fiveToolsJob) update(f func(*views.FiveToolsJob)) {
	j.mu.Lock()
	defer j.mu.Unlock()
	f(&j.job)
}

// runFiveTools downloads the books, imports every folder and leaves the
// summary on the job, telling the GM's pages at each step.
func (a *App) runFiveTools(ctx context.Context, dir string, books []compendium.Book) {
	tell := func(f func(*views.FiveToolsJob)) {
		a.fiveTools.update(f)
		a.Hub.Publish(0, hub.Topic{Kind: hub.TopicImport})
	}
	msg, failed := "", true
	defer func() {
		// Outside a request, a panic would take the server down.
		if p := recover(); p != nil {
			slog.Error("5etools import", "panic", p)
			msg, failed = fmt.Sprint("Import failed: ", p), true
		}
		tell(func(j *views.FiveToolsJob) { *j = views.FiveToolsJob{Msg: msg, Failed: failed} })
	}()
	err := compendium.DownloadFiveTools(ctx, dir, books, func(done int, b compendium.Book) {
		tell(func(j *views.FiveToolsJob) { j.Done, j.Current = done, b.Name })
	})
	if err != nil {
		msg = "Download failed: " + err.Error()
		return
	}
	reports, err := compendium.ImportProgress(ctx, a.Store, dir, func(done, total int, source string) {
		tell(func(j *views.FiveToolsJob) { j.Importing, j.Done, j.Total, j.Current = true, done, total, source })
	}, nil)
	a.Hub.Publish(0, hub.Topic{Kind: hub.TopicCompendium})
	if err != nil {
		slog.Error("5etools import", "err", err)
		msg = "Import failed: " + err.Error()
		return
	}
	msg, failed = ImportSummary(reports)
}
