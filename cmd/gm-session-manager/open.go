package main

import (
	"context"
	"log/slog"
	"os/exec"
	"runtime"
)

// openBrowser shows url in the default browser; failing only costs a click.
func openBrowser(url string) {
	// Background: the browser must outlive any request or shutdown context.
	ctx := context.Background()
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		//nolint:gosec // url is built from our own port
		cmd = exec.CommandContext(ctx, "rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.CommandContext(ctx, "open", url) //nolint:gosec // url is built from our own port
	default:
		//nolint:gosec // url is built from our own port
		cmd = exec.CommandContext(ctx, "xdg-open", url)
	}
	if err := cmd.Start(); err != nil {
		slog.Warn("cannot open the browser; open this address yourself", "url", url, "err", err)
	}
}
