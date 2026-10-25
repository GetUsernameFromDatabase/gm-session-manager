package main

import (
	"context"
	"log/slog"
	"syscall"
	"time"
)

// esSystemRequired resets Windows' idle sleep timer once; without
// ES_CONTINUOUS it needs no thread of its own, so it is simply repeated.
const esSystemRequired = 0x00000001

// keepAwake stops the host from sleeping while the server runs, so players'
// devices do not lose it mid-session. The screen may still turn off.
func keepAwake(ctx context.Context) {
	proc := syscall.NewLazyDLL("kernel32.dll").NewProc("SetThreadExecutionState")
	if err := proc.Find(); err != nil {
		slog.Warn("cannot keep the computer awake", "err", err)
		return
	}
	tick := time.NewTicker(30 * time.Second)
	defer tick.Stop()
	for {
		_, _, _ = proc.Call(esSystemRequired)
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
		}
	}
}
