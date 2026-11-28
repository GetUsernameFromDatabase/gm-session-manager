package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

// waitIfOwnConsole keeps a double-clicked exe's window open after a fatal
// error: the exe is then the only process on its console.
func waitIfOwnConsole() {
	proc := syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleProcessList")
	var pids [2]uint32
	//nolint:gosec // the syscall needs a buffer pointer
	n, _, _ := proc.Call(uintptr(unsafe.Pointer(&pids[0])), uintptr(len(pids)))
	if n != 1 {
		return
	}
	fmt.Fprintln(os.Stderr, "Press Enter to close.")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}
