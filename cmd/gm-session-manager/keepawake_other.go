//go:build !windows

package main

import "context"

// keepAwake is Windows only; elsewhere, run the app under caffeinate (macOS)
// or systemd-inhibit (Linux) to keep the host awake.
func keepAwake(context.Context) {}
