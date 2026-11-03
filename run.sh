#!/usr/bin/env bash
# Usage: ./run.sh          the real instance (.env, data/)
#        ./run.sh browser  a test instance for browser checks (data-test/, port 8443);
#                          touch data-test/restart to rebuild and restart it, data-test/stop to quit
cd "$(dirname "$0")"
if [ "$1" = browser ]; then
	export DATA_DIR=data-test PORT=8443
	# Its own binary name, so a stray one can be stopped by name without touching the real instance.
	bin=build/gm-session-manager-test$(go env GOEXE)
	mkdir -p "$DATA_DIR"
	rm -rf "$DATA_DIR/playwright" "$DATA_DIR/restart" "$DATA_DIR/stop"
	pid=
	stop_server() {
		[ -n "$pid" ] && kill "$pid" 2>/dev/null && wait "$pid" 2>/dev/null
		pid=
	}
	trap stop_server EXIT
	start=1
	while :; do
		if [ -e "$DATA_DIR/stop" ]; then
			rm -f "$DATA_DIR/stop"
			exit 0
		fi
		if [ -e "$DATA_DIR/restart" ]; then
			rm -f "$DATA_DIR/restart"
			start=1
		fi
		if [ -n "$start" ]; then
			start=
			stop_server
			if go build -o "$bin" ./cmd/gm-session-manager; then
				"$bin" &
				pid=$!
			else
				echo "run.sh: build failed; fix it and touch $DATA_DIR/restart"
			fi
		fi
		sleep 1
	done
fi
exec go run ./cmd/gm-session-manager
