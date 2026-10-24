// Command gm-session-manager runs the LAN host for a tabletop RPG table.
package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"

	"gm-session-manager/auth"
	"gm-session-manager/compendium"
	"gm-session-manager/config"
	"gm-session-manager/db"
	"gm-session-manager/join"
	"gm-session-manager/web"
)

const (
	envFile         = ".env"
	databaseFile    = "gm-session-manager.db"
	tlsDir          = "tls"
	compendiumDir   = "compendium"
	shutdownTimeout = 5 * time.Second
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "gm-session-manager:", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load(envFile)
	if err != nil {
		return err
	}
	if os.Getenv(gin.EnvGinMode) == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := os.MkdirAll(cfg.DataDir, 0o750); err != nil {
		return fmt.Errorf("create data dir: %w", err)
	}
	conn, err := db.Open(ctx, filepath.Join(cfg.DataDir, databaseFile))
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	lanIPs, err := join.LANAddresses()
	if err != nil {
		return err
	}
	cert, created, err := join.LoadOrCreateCert(filepath.Join(cfg.DataDir, tlsDir), lanIPs, time.Now())
	if err != nil {
		return fmt.Errorf("tls certificate: %w", err)
	}
	slog.Info("tls certificate", "new", created, "sha256", join.Fingerprint(cert), "lan", lanIPs)

	store := db.NewStore(conn)
	reports, err := compendium.Import(ctx, store, filepath.Join(cfg.DataDir, compendiumDir))
	if err != nil {
		return err
	}
	for _, r := range reports {
		if r.Err != nil {
			slog.Warn("compendium import failed", "source", r.Source, "err", r.Err)
		} else {
			slog.Info("compendium", "source", r.Source, "unchanged", r.Unchanged, "entries", r.Entries, "removed", r.Removed)
		}
	}
	if n, err := web.UpgradeSheets(ctx, store); err != nil {
		slog.Warn("sheet upgrade failed", "err", err)
	} else if n > 0 {
		slog.Info("sheets upgraded", "changed", n)
	}

	app := web.New(web.Deps{
		Auth:        auth.NewManager(store, cfg.AdminUsername, cfg.AdminPassword),
		Store:       store,
		DataDir:     cfg.DataDir,
		Port:        cfg.Port,
		LANIPs:      lanIPs,
		Fingerprint: join.Fingerprint(cert),
	})
	defer app.Close()
	if err := app.CleanAttachments(ctx); err != nil {
		slog.Warn("cleaning attachments", "err", err)
	}

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           app.Router(),
		ReadHeaderTimeout: 10 * time.Second,
		TLSConfig:         &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12},
	}
	// Websockets would otherwise keep Shutdown waiting until its timeout.
	srv.RegisterOnShutdown(app.Close)

	errc := make(chan error, 1)
	go func() {
		slog.Info("listening", "url", fmt.Sprintf("https://localhost:%d", cfg.Port), "data_dir", cfg.DataDir)
		errc <- srv.ListenAndServeTLS("", "")
	}()

	select {
	case err := <-errc:
		return err
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
