// Command gm-session-manager runs the LAN host for a tabletop RPG table.
package main

import (
	"context"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/config"
	"gm-session-manager/internal/db"
	"gm-session-manager/internal/join"
	"gm-session-manager/internal/web"
)

const (
	envFile         = ".env"
	databaseFile    = "gm-session-manager.db"
	tlsDir          = "tls"
	compendiumDir   = "compendium"
	shutdownTimeout = 5 * time.Second
	// readHeaderTimeout stops a client that never finishes its headers from holding a connection.
	readHeaderTimeout = 10 * time.Second
)

// version is set at build time by scripts/build.sh (-X main.version=…).
var version = "dev"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "gm-session-manager:", err)
		waitIfOwnConsole()
		os.Exit(1)
	}
}

// loopbackHostOnly refuses requests addressed to any other host name. Plain
// HTTP has no certificate to stop DNS rebinding: a web page could point its
// own domain at 127.0.0.1 and reach the player app with loopback's rights.
func loopbackHostOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host, _, err := net.SplitHostPort(r.Host)
		if err != nil {
			host = r.Host // no port: port 80
		}
		if host != "127.0.0.1" && host != "localhost" {
			http.Error(w, "open this app at 127.0.0.1", http.StatusMisdirectedRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// parseFlags reads the command line; --port, --data and --no-browser only make sense for the player app.
func parseFlags(args []string) (player, noBrowser bool, port int, data string, err error) {
	fs := flag.NewFlagSet("gm-session-manager", flag.ContinueOnError)
	fs.BoolVar(&player, "player", false, "run the player app on this computer: 127.0.0.1 only, plain HTTP, no admin login, configured by flags only")
	fs.BoolVar(&noBrowser, "no-browser", false, "do not open the player app in the browser (with --player)")
	fs.IntVar(&port, "port", config.DefaultPlayerPort, "player app port (with --player)")
	fs.StringVar(&data, "data", config.DefaultPlayerDataDir, "player app data folder (with --player)")
	if err = fs.Parse(args); err != nil {
		return false, false, 0, "", err
	}
	given := false
	fs.Visit(func(f *flag.Flag) { given = given || f.Name != "player" })
	if given && !player {
		return false, false, 0, "", errors.New("--port, --data and --no-browser apply only with --player")
	}
	return player, noBrowser, port, data, nil
}

func run() error {
	player, noBrowser, port, data, err := parseFlags(os.Args[1:])
	if errors.Is(err, flag.ErrHelp) {
		return nil // the usage is already printed
	}
	if err != nil {
		return err
	}
	var cfg config.Config
	if player {
		cfg, err = config.Player(port, data)
	} else {
		cfg, err = config.Load(envFile)
	}
	if err != nil {
		return err
	}
	if os.Getenv(gin.EnvGinMode) == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	// SIGTERM is how Linux service managers and kill stop a process.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := os.MkdirAll(cfg.DataDir, 0o750); err != nil {
		return fmt.Errorf("create data dir: %w", err)
	}
	conn, err := db.Open(ctx, filepath.Join(cfg.DataDir, databaseFile))
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	// The player app is plain HTTP on loopback: no LAN addresses, no certificate.
	var (
		lanIPs      []net.IP
		fingerprint string
		tlsConfig   *tls.Config
	)
	if cfg.Player {
		slog.Info("player app: this computer only, no login needed")
	} else {
		if lanIPs, err = join.LANAddresses(); err != nil {
			return err
		}
		cert, created, err := join.LoadOrCreateCert(filepath.Join(cfg.DataDir, tlsDir), lanIPs, time.Now())
		if err != nil {
			return fmt.Errorf("tls certificate: %w", err)
		}
		fingerprint = join.Fingerprint(cert)
		tlsConfig = &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12}
		slog.Info("tls certificate", "new", created, "sha256", fingerprint, "lan", lanIPs)
		// The take-home dialog lists these files for players to download.
		if err := os.MkdirAll(filepath.Join(cfg.DataDir, web.DownloadsDir), 0o750); err != nil {
			return fmt.Errorf("create downloads dir: %w", err)
		}
		// Bring-back files waiting for the GM live only until a restart (web.homeFiles).
		homeFiles := filepath.Join(cfg.DataDir, web.HomeFilesDir)
		if err := os.RemoveAll(homeFiles); err != nil {
			return fmt.Errorf("clear home files dir: %w", err)
		}
		if err := os.MkdirAll(homeFiles, 0o750); err != nil {
			return fmt.Errorf("create home files dir: %w", err)
		}
	}

	store := db.NewStore(conn)
	if err := compendium.CleanDownloads(filepath.Join(cfg.DataDir, compendiumDir)); err != nil {
		slog.Warn("removing unfinished 5etools downloads", "err", err)
	}
	// Each source is logged as it finishes, so a long import does not look stalled.
	reports, err := compendium.ImportProgress(ctx, store, filepath.Join(cfg.DataDir, compendiumDir), func(done, total int, _ string) {
		if done == 0 {
			slog.Info("compendium importing", "sources", total)
		}
	}, func(r compendium.Report) {
		if r.Err != nil {
			slog.Warn("compendium import failed", "source", r.Source, "err", r.Err)
		} else {
			slog.Info("compendium", "source", r.Source, "unchanged", r.Unchanged, "entries", r.Entries, "removed", r.Removed)
		}
	})
	if err != nil {
		return err
	}
	if len(reports) > 0 {
		var changed, unchanged, failed, entries int
		for _, r := range reports {
			switch {
			case r.Err != nil:
				failed++
			case r.Unchanged:
				unchanged++
			default:
				changed++
				entries += r.Entries
			}
		}
		slog.Info("compendium done", "changed", changed, "unchanged", unchanged, "failed", failed, "entries", entries)
	}
	if n, err := web.UpgradeSheets(ctx, store); err != nil {
		slog.Warn("sheet upgrade failed", "err", err)
	} else if n > 0 {
		slog.Info("sheets upgraded", "changed", n)
	}

	manager := auth.NewManager(store, cfg.AdminUsername, cfg.AdminPassword)
	if cfg.Player {
		manager = auth.NewPlayerManager(store)
	}
	app := web.New(web.Deps{
		Auth:        manager,
		Player:      cfg.Player,
		Version:     version,
		Store:       store,
		DataDir:     cfg.DataDir,
		Port:        cfg.Port,
		LANIPs:      lanIPs,
		Fingerprint: fingerprint,
	})
	defer app.Close()
	if err := app.CleanAttachments(ctx); err != nil {
		slog.Warn("cleaning attachments", "err", err)
	}
	if err := app.WriteFeedbackIndex(); err != nil {
		slog.Warn("listing unread feedback", "err", err)
	}
	if err := app.StartUp(ctx); err != nil {
		slog.Warn("settings or stale 5etools books", "err", err)
	}

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           app.Router(),
		ReadHeaderTimeout: readHeaderTimeout,
		TLSConfig:         tlsConfig,
	}
	if cfg.Player {
		// IPv4 only: "localhost" may resolve to ::1.
		srv.Addr = fmt.Sprintf("127.0.0.1:%d", cfg.Port)
		srv.Handler = loopbackHostOnly(srv.Handler)
	}
	// Websockets would otherwise keep Shutdown waiting until its timeout.
	srv.RegisterOnShutdown(app.Close)

	base, err := (&net.ListenConfig{}).Listen(ctx, "tcp", srv.Addr)
	if err != nil {
		return err
	}
	defer func() { _ = base.Close() }()

	var redirect *http.Server
	errc := make(chan error, 1)
	if cfg.Player {
		url := fmt.Sprintf("http://127.0.0.1:%d/", cfg.Port)
		go func() {
			slog.Info("listening", "url", url, "data_dir", cfg.DataDir, "version", version)
			errc <- srv.Serve(base)
		}()
		if !noBrowser {
			openBrowser(url)
		}
	} else {
		// Browsers try http:// for an address typed without a scheme, so the same
		// port redirects plain HTTP to HTTPS.
		tlsL, plainL := join.SplitTLS(base)
		redirect = &http.Server{Handler: join.RedirectHTTPS, ReadHeaderTimeout: readHeaderTimeout}
		go func() { _ = redirect.Serve(plainL) }()

		go keepAwake(ctx)
		go func() {
			slog.Info("listening", "url", fmt.Sprintf("https://localhost:%d", cfg.Port), "data_dir", cfg.DataDir, "version", version)
			errc <- srv.ServeTLS(tlsL, "", "")
		}()
	}

	select {
	case err := <-errc:
		return err
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if redirect != nil {
		_ = redirect.Shutdown(shutdownCtx)
	}
	if err := srv.Shutdown(shutdownCtx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
