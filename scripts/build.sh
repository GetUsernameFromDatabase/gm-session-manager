#!/usr/bin/env bash
# The build half of task build and task build:release (the tasks vendor, bundle, lint and test first):
#   (no argument)   this machine -> dist/$APP[.exe]
#   release         every release target -> dist/*.zip|*.tar.gz + SHA256SUMS. Linux only (CI and the
#                   update job): Task's Windows tar lacks -C and -z.
# APP and PKG come from the Taskfile. The version is stamped into the binary: $VERSION (CI sets the tag) or
# git describe.
set -euo pipefail
cd "$(git rev-parse --show-toplevel)"

VERSION=${VERSION:-$(git describe --tags --always --dirty 2>/dev/null || echo dev)}
targets="windows/amd64 linux/amd64" # other systems build from source
out=dist                            # binaries and release archives (release.yml, .gitignore)
stage=build/stage                   # scratch folders packed into the archives

build() { # os arch out
	GOOS=$1 GOARCH=$2 go build -trimpath -ldflags "-s -w -X main.version=$VERSION" -o "$3" "$PKG"
}

rm -rf "$out"
mkdir "$out"
if [ "${1:-}" != release ]; then
	build "$(go env GOOS)" "$(go env GOARCH)" "$out/$APP$(go env GOEXE)"
	ls -lh "$out"
	exit 0
fi

rm -rf "$stage"
for target in $targets; do
	os=${target%/*} arch=${target#*/}
	name=$APP-$VERSION-$os-$arch
	exe=$APP
	[ "$os" = windows ] && exe=$exe.exe
	build "$os" "$arch" "$stage/$name/$exe"
	cp README.md .env.example "$stage/$name/" # shipped next to the binary, with a player-app launcher
	if [ "$os" = windows ]; then
		printf '@echo off\r\ncd /d "%%~dp0"\r\n%s --player\r\nif errorlevel 1 pause\r\n' "$exe" >"$stage/$name/Player app.bat"
		(cd "$stage" && zip -qr "$OLDPWD/$out/$name.zip" "$name")
	else
		printf '#!/bin/sh\ncd "$(dirname "$0")" && exec ./%s --player\n' "$exe" >"$stage/$name/player-app.sh"
		chmod +x "$stage/$name/player-app.sh"
		tar -C "$stage" -czf "$out/$name.tar.gz" "$name"
	fi
done
# Not `*`: Task's shell creates SHA256SUMS before expanding the glob, so it would hash itself.
(cd "$out" && sha256sum -- *.zip *.tar.gz >SHA256SUMS)
ls -lh "$out"
