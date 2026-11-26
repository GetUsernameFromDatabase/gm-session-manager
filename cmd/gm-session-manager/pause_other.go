//go:build !windows

package main

// waitIfOwnConsole is Windows only: elsewhere a terminal stays open on its own.
func waitIfOwnConsole() {}
