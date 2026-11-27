package web

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"sync"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/web/views"
)

// fiveToolsJob is the GM's background 5etools import, in memory like join
// requests: one at a time, forgotten on restart (the books on disk import
// on start anyway). It also runs the Open5e SRD download (runOpen5e).
type fiveToolsJob struct {
	mu  sync.Mutex
	job views.FiveToolsJob
}

func (j *fiveToolsJob) get() views.FiveToolsJob {
	j.mu.Lock()
	defer j.mu.Unlock()
	return j.job
}

// start claims the job for total books (0 until a listing says); false while
// one runs.
func (j *fiveToolsJob) start(total int) bool {
	j.mu.Lock()
	defer j.mu.Unlock()
	if j.job.Running {
		return false
	}
	j.job = views.FiveToolsJob{Running: true, Total: total}
	return true
}

// dismiss forgets the last import's summary if it is still was; false while
// one runs or after a newer one replaced it.
func (j *fiveToolsJob) dismiss(was string) bool {
	j.mu.Lock()
	defer j.mu.Unlock()
	if j.job.Running || j.job.Msg != was {
		return false
	}
	j.job = views.FiveToolsJob{}
	return true
}

func (j *fiveToolsJob) update(f func(*views.FiveToolsJob)) {
	j.mu.Lock()
	defer j.mu.Unlock()
	f(&j.job)
}

// runFiveTools downloads the books, then imports every folder (runImport).
// A campaign (0 for none) gets the books' sources ticked.
func (a *App) runFiveTools(ctx context.Context, dir string, books []compendium.Book, campaign int64) {
	keys := make([]string, len(books))
	for i, b := range books {
		keys[i] = fiveToolsKey(b.ID)
	}
	a.runImport(ctx, dir, campaign, keys, func(progress func(done, total int, current string)) error {
		return compendium.DownloadFiveTools(ctx, dir, books, func(done int, b compendium.Book) {
			progress(done, len(books), b.Name)
		})
	})
}

// runOpen5e downloads an SRD (doc), then imports every folder (runImport).
// A campaign (0 for none) gets it ticked.
func (a *App) runOpen5e(ctx context.Context, dir, doc string, campaign int64) {
	a.runImport(ctx, dir, campaign, []string{doc}, func(progress func(done, total int, current string)) error {
		return compendium.DownloadOpen5e(ctx, dir, doc, progress)
	})
}

// runImport runs download, imports every folder and leaves the summary on
// the job, telling the GM's pages at each step. A campaign compendium shows
// only its ticked sources, so a starting set posted from one (campaign > 0)
// ticks its keys there.
func (a *App) runImport(ctx context.Context, dir string, campaign int64, keys []string, download func(progress func(done, total int, current string)) error) {
	tell := func(f func(*views.FiveToolsJob)) {
		a.fiveTools.update(f)
		a.Hub.Publish(0, hub.Topic{Kind: hub.TopicImport})
	}
	msg, failed := "", true
	defer func() {
		// Outside a request, a panic would take the server down.
		if p := recover(); p != nil {
			slog.Error("compendium import", "panic", p)
			msg, failed = fmt.Sprint("Import failed: ", p), true
		}
		tell(func(j *views.FiveToolsJob) {
			var books []string
			if failed {
				books = j.Books
			}
			*j = views.FiveToolsJob{Msg: msg, Failed: failed, Books: books}
		})
	}()
	err := download(func(done, total int, current string) {
		tell(func(j *views.FiveToolsJob) { j.Done, j.Total, j.Current = done, total, current })
	})
	if serr := a.refreshStale(); serr != nil {
		slog.Warn("listing stale 5etools books", "err", serr)
	}
	// Ticks what exists even after a partial failure (books already in).
	tick := func() {
		if err := a.allowSources(ctx, campaign, keys); err != nil {
			slog.Error("ticking imported sources", "campaign", campaign, "err", err)
			msg, failed = msg+"\nCould not add them to the campaign: "+err.Error(), true
		}
	}
	if err != nil {
		msg = netErrMsg("Download failed", err)
		tick()
		return
	}
	reports, err := compendium.ImportProgress(ctx, a.Store, dir, func(done, total int, source string) {
		tell(func(j *views.FiveToolsJob) { j.Importing, j.Done, j.Total, j.Current = true, done, total, source })
	}, nil)
	a.Hub.Publish(0, hub.Topic{Kind: hub.TopicCompendium})
	if err != nil {
		slog.Error("compendium import", "err", err)
		msg = "Import failed: " + err.Error()
		return
	}
	msg, failed = ImportSummary(reports)
	tick()
}

// netErrMsg words a failed download: what failed, and for a network error
// (no connection, DNS, a dropped transfer) that the app could not reach the
// site. Not *url.Error: a refused redirect or a bad certificate is one too.
func netErrMsg(what string, err error) string {
	if errors.As(err, new(*net.OpError)) {
		return "No internet connection. " + what + ": the app could not reach GitHub or the site. Connect and try again. (" + err.Error() + ")"
	}
	return what + ": " + err.Error()
}

// sinkHomeSources keeps the sources added at home below the table's in the
// player app (no-op at the table).
func (a *App) sinkHomeSources(ctx context.Context, q *sqlc.Queries) error {
	if !a.Player {
		return nil
	}
	return q.SinkHomeSources(ctx)
}

// allowSources ticks keys for campaign (none for 0); a key no source holds
// (a book that failed) is passed over.
func (a *App) allowSources(ctx context.Context, campaign int64, keys []string) error {
	if campaign <= 0 || len(keys) == 0 {
		return nil
	}
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		for _, k := range keys {
			if _, err := q.AllowCampaignSource(ctx, sqlc.AllowCampaignSourceParams{CampaignID: campaign, SourceKey: k}); err != nil {
				return err
			}
		}
		return a.sinkHomeSources(ctx, q)
	})
	if err == nil {
		a.Hub.Publish(campaign, hub.Topic{Kind: hub.TopicCompendium})
	}
	return err
}
