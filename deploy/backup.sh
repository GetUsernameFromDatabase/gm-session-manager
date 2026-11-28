#!/usr/bin/env bash
# Backs up DATA_DIR into DATA_DIR/backups/: a tar.gz of everything except backups/, downloads/ (and the
# downloads.new/ an update stages) and the live database files, plus a consistent copy of the database made by
# sqlite3 (safe while the app runs).
#   backup.sh nightly       nightly-YYYY-MM-DD.tar.gz (the timer's, through update.sh --nightly), keeps the newest 5;
#                           skipped when nothing in DATA_DIR changed since the newest one
#   backup.sh manual        manual-YYYY-MM-DD-HHMMSS.tar.gz (by hand, or update.sh run by hand), keeps the newest 3
#   backup.sh pre-deploy    pre-deploy-<commit>-<time>.tar.gz, keeps the newest 3
# <commit> is the deployed commit, the state the data belongs to. The .env is the checkout's,
# or ENV_FILE.
# File names come from common.sh.
set -euo pipefail

here=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
kind=${1:-}
case $kind in nightly) keep=5 ;; manual | pre-deploy) keep=3 ;; *) echo "usage: backup.sh nightly|manual|pre-deploy" >&2; exit 2 ;; esac

die() { echo "backup: $*" >&2; exit 1; }
command -v sqlite3 >/dev/null || die "sqlite3 is missing: apt install sqlite3"
# shellcheck source=deploy/common.sh
. "$here/common.sh"
db=$data/$DB_FILE
backups=$data/$BACKUPS_DIR
[ -f "$db" ] || die "$db not found"

# Run by hand, wait for a running update; run by update.sh, it holds the lock already.
[ -n "${GSM_LOCKED:-}" ] || lock

# Most nights nobody plays: identical nightlies would push the older states out of the rotation.
if [ "$kind" = nightly ]; then
	# shellcheck disable=SC2012
	last=$(ls -1t "$backups/nightly-"*.tar.gz 2>/dev/null | head -n 1 || true)
	# Not data: the updater's own files, and the root folder, which changes whenever they are rewritten.
	if [ -n "$last" ] && [ -z "$(find "$data" -mindepth 1 \( -path "$backups" -o -path "$data/$DOWNLOADS_DIR" -o -path "$data/$DOWNLOADS_DIR.new" \) -prune -o \
		-newer "$last" ! -name "$STATUS_FILE" ! -name "$DEPLOYED_FILE" ! -name "$FAILED_FILE" ! -name "$DB_FILE-shm" -print -quit)" ]; then
		echo "backup: nothing changed since $last"
		exit 0
	fi
fi

mkdir -p "$backups"
tmp=$(mktemp -d "$backups/.tmp-XXXXXX")
trap 'rm -rf "$tmp"' EXIT

# The backup's time is its start: a change made while it runs is newer than it, so the next night saves it.
touch "$tmp/start"
sqlite3 "$db" ".backup '$tmp/$DB_FILE'"
# This script runs as root: a restored root-owned copy would be read-only to the app.
chown --reference="$db" "$tmp/$DB_FILE"
case $kind in
nightly) name=nightly-$(date +%F).tar.gz ;;
manual) name=manual-$(date +%F-%H%M%S).tar.gz ;;
*)
	commit=$(head -c 7 "$data/$DEPLOYED_FILE" 2>/dev/null || true)
	name=pre-deploy-${commit:-none}-$(date +%Y%m%d-%H%M%S).tar.gz
	;;
esac
tar -czf "$tmp/$name" --no-wildcards-match-slash \
	--exclude="./$BACKUPS_DIR" --exclude="./$DOWNLOADS_DIR" --exclude="./$DOWNLOADS_DIR.new" \
	--exclude="./$DB_FILE" --exclude="./$DB_FILE-wal" --exclude="./$DB_FILE-shm" \
	-C "$data" . -C "$tmp" "$DB_FILE"
touch -r "$tmp/start" "$tmp/$name"
mv -f "$tmp/$name" "$backups/$name"

# Newest first; names have no spaces.
# shellcheck disable=SC2012
ls -1t "$backups/$kind-"*.tar.gz | tail -n +$((keep + 1)) | xargs -r rm --
echo "backup: $backups/$name"
