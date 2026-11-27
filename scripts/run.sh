#!/usr/bin/env bash
# The dev server loop behind task run, run:player, run:browser and run:before (the modes are their env: in
# Taskfile.yml; task --list describes them). Task runs it in its own shell (source), so it also works from
# PowerShell: no find, sort or bash-only set flags.
# Watching modes rebuild and restart when sources change, keeping the old server if the build fails; a changed
# migration pauses reloading. touch $MARKERS/restart to force a restart (and resume), $MARKERS/stop to quit.
# A second start of a mode takes over from the first; each quits with its terminal.
# Input env: BIN (the server, unique per mode so a stray one is stopped by name) and MARKERS (restart, stop
# and run.pid live there) are required; MODE (names the change-detection state); WATCH (rebuild when sources
# change); NOREBUILD (BIN is prebuilt); SERVER_ARGS; RESTART_ARGS (added to SERVER_ARGS after the first
# start); PREPARE (a task to run once the old server is stopped).
set +e # Task's shell starts with errexit; the loop checks its own statuses
[ -n "${BIN:-}" ] && [ -n "${MARKERS:-}" ] || { echo "run: BIN and MARKERS are required; use task run, run:player, run:browser or run:before"; exit 1; }
[ -z "${NOREBUILD:-}" ] || [ -f "$BIN" ] || { echo "run: no $BIN; run task before first"; exit 1; }

mkdir -p build "$MARKERS"
ext=
case $BIN in *.exe) ext=.exe ;; esac
args=${SERVER_ARGS:-} restart_args=${RESTART_ARGS:-}
started=

# On Windows a pid is probed with tasklist and the server stopped by name (kill does not always end a native
# process, and a stray from a killed loop goes too). Linux: the minimum.
if [ -n "${WINDIR:-}" ]; then
	alive() { tasklist /FI "PID eq $1" /NH | grep -qw "$1"; }
	stop_server() {
		taskkill /F /IM "${BIN##*/}" >/dev/null 2>&1
		# taskkill returns before the process is gone, and a running exe cannot be replaced: wait (up to 5 s).
		# CSV, because the table format cuts image names at 25 characters.
		for _ in 1 2 3 4 5 6 7 8 9 10; do
			tasklist /FI "IMAGENAME eq ${BIN##*/}" /NH /FO CSV | grep -qi "\"${BIN##*/}\"" || break
			sleep 0.5
		done
		started=
	}
else
	alive() { kill -0 "$1" 2>/dev/null; }
	stop_server() { pkill -f -- "$BIN" 2>/dev/null; started=; }
fi

# One generate-and-build at a time across loops (task run and task run:player in two terminals share the
# generated files): mkdir either takes the lock or fails. A lock whose holder died is taken over.
# shortcut: a holder killed between mkdir and writing its pid leaves a lock nobody removes (rm -rf build/.gen.lock),
# and two waiters seeing one dead holder can, rarely, both get in. Fine for two or three local loops.
lockdir=build/.gen.lock locked=
pidfile=$MARKERS/run.pid
# This loop was taken over by a newer run of its mode, or its terminal closed (Windows sends no signal).
superseded() {
	! alive "$PPID" || [ "$(cat "$pidfile" 2>/dev/null)" != "$$" ]
}
lock() {
	# Task's Windows mkdir reports an existing dir on stderr but exits 0, so success is "no output".
	until [ -z "$(mkdir "$lockdir" 2>&1)" ]; do
		# Waiting can take a whole build of the other loop: notice a takeover meanwhile.
		superseded && exit 0
		holder=$(cat "$lockdir/pid" 2>/dev/null)
		[ -n "$holder" ] && ! alive "$holder" && rm -rf "$lockdir"
		sleep 0.5
	done
	echo $$ >"$lockdir/pid"
	locked=1
}
unlock() {
	[ -n "$locked" ] && rm -rf "$lockdir"
	locked=
}

# Task keeps when each change task last ran under TASK_TEMP_DIR: one per mode, so modes do not share state.
unchanged() { TASK_TEMP_DIR=build/.task/$MODE task -s --status "$1" >/dev/null 2>&1; }

# Build first and swap only on success, so a broken build leaves the old server running.
start_server() {
	if [ -n "${WATCH:-}" ]; then
		# Before generating and building, so an edit saved meanwhile counts as the next change.
		TASK_TEMP_DIR=build/.task/$MODE task -s dev:changed dev:migrations
	fi
	if [ -z "${NOREBUILD:-}" ]; then
		new=${BIN%"$ext"}-new$ext
		lock
		task dev:build OUT="$new"
		built=$?
		unlock
		superseded && exit 0 # taken over during the build: the newer loop starts the server
		if [ "$built" != 0 ]; then
			echo "run: build failed${started:+; the old server keeps running}. Fix it (or touch $MARKERS/restart)"
			return 0
		fi
	fi
	stop_server
	if [ -z "${NOREBUILD:-}" ]; then
		mv "$new" "$BIN" || echo "run: could not replace $BIN; the old build runs. touch $MARKERS/restart to retry"
	fi
	"$BIN" $args &
	started=1
	args="$args $restart_args" restart_args=
}

rm -f "$MARKERS/restart" "$MARKERS/stop"
# One loop at a time: a newer run takes over (an older one waits for every marker).
old=$(cat "$pidfile" 2>/dev/null)
echo $$ >"$pidfile"
for _ in 1 2 3 4 5; do
	[ -n "$old" ] && alive "$old" || break
	sleep 1
done
stop_server
# After the old server is stopped: on Windows it holds its files open.
[ -z "${PREPARE:-}" ] || task -s "$PREPARE" || exit 1
trap 'unlock; stop_server' EXIT

start=1 paused=
while :; do
	superseded && exit 0
	if [ -e "$MARKERS/stop" ]; then
		rm -f "$MARKERS/stop"
		exit 0
	fi
	if [ -e "$MARKERS/restart" ]; then
		rm -f "$MARKERS/restart"
		start=1 paused=
	fi
	if [ -n "${WATCH:-}" ] && [ -z "$paused" ] && [ -z "$start" ] && ! unchanged dev:changed; then
		# A migration from half-finished work would change the real DB: wait for the user.
		if ! unchanged dev:migrations; then
			paused=1
			echo "run: migrations changed; not restarting. touch $MARKERS/restart when ready"
		else
			start=1
		fi
	fi
	if [ -n "$start" ]; then
		start=
		start_server
	fi
	sleep 2
done
