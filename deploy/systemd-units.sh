#!/usr/bin/env bash
# Installs or removes the nightly update's systemd units (README.md); run as root.
#   systemd-units.sh install   writes systemd/* into /etc/systemd/system with this checkout's path and starts the
#                              timer; an installed unit is never replaced: it shows how it differs from the repo's
#   systemd-units.sh remove    stops the timer and deletes the units; a running update finishes, and the drop-ins of
#                              systemctl edit stay, so remove then install takes the repo's units and keeps them
set -euo pipefail

here=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
repo=$(dirname "$here")
dest=/etc/systemd/system
units=(gsm-update.service gsm-update.timer)
timer=gsm-update.timer

die() { echo "systemd-units: $*" >&2; exit 1; }

# A unit as install writes it.
render() {
	local s
	s=$(<"$here/systemd/$1")
	# systemd reads % as a specifier (%h is root's home), so the path's own % is doubled.
	local r=${repo//\%/%%}
	# The quotes keep a & in the path literal (bash 5.2 replaces it with the match).
	printf '%s\n' "${s//@REPO@/"$r"}"
}

[ "$(id -u)" = 0 ] || die "run as root"
case ${1:-} in
install)
	installed=0 missing=()
	for u in "${units[@]}"; do
		if [ ! -e "$dest/$u" ]; then
			missing+=("$dest/$u")
			continue
		fi
		installed=1
		if cmp -s "$dest/$u" <(render "$u"); then
			echo "$dest/$u: installed, the same as the repo's"
		else
			echo "$dest/$u: installed, differs from the repo's:"
			diff -u "$dest/$u" <(render "$u") || true
		fi
	done
	if [ "$installed" = 1 ]; then
		# A half install: the missing one is why the nightly update does not run.
		for m in "${missing[@]}"; do
			echo "$m: not installed"
		done
		die "nothing changed. Edit the files above (then systemctl daemon-reload), or take the repo's: bash \"$0\" remove, then install"
	fi
	for u in "${units[@]}"; do
		render "$u" >"$dest/$u"
	done
	systemctl daemon-reload
	systemctl enable --now "$timer"
	printf 'installed %s\n' "${units[@]/#/$dest/}"
	echo "To change them: edit there and systemctl daemon-reload, or keep your changes apart with systemctl edit <unit>."
	;;
remove)
	if [ -e "$dest/$timer" ]; then
		systemctl disable --now "$timer"
	fi
	removed=()
	for u in "${units[@]}"; do
		if [ -e "$dest/$u" ]; then
			rm "$dest/$u"
			removed+=("$u")
		fi
	done
	systemctl daemon-reload
	echo "removed: ${removed[*]:-nothing, none was installed}"
	for u in "${units[@]}"; do
		if [ -d "$dest/$u.d" ]; then
			echo "kept your drop-ins: $dest/$u.d (rm -r it to drop them)"
		fi
	done
	;;
*)
	echo "usage: systemd-units.sh install|remove" >&2
	exit 2
	;;
esac
