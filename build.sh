#!/usr/bin/env bash
# The one build process, for developers and CI.
#   ./build.sh           host OS/arch -> $out/gm-session-manager[.exe]
#   ./build.sh release   every release target -> $out/*.zip|*.tar.gz + SHA256SUMS
set -Eeuo pipefail
cd "$(dirname "$0")"

app=gm-session-manager
pkg=./cmd/$app
targets="windows/amd64 linux/amd64" # other systems build from source
extras="README.md .env.example" # shipped next to the binary in each archive
out=dist                        # binaries and release archives (release.yml, .gitignore)
stage=build                     # scratch folders packed into the archives

VERSION=${VERSION:-$(git describe --tags --always --dirty 2>/dev/null || echo dev)}
export CGO_ENABLED=0

current=setup
step() { current=$*; printf '\n==> %s %s\n' "$(date +%T)" "$*"; }
trap 'echo "build.sh: failed during \"$current\" at $(date +%T)" >&2' ERR
echo "building $app $VERSION ($(go env GOVERSION), mode ${1:-host})"

# Both outputs are gitignored, so a fresh checkout (CI included) lacks them.
step "vendoring browser JS"
go run ./cmd/vendorjs
step "bundling shadcn-templ scripts"
go tool shadcn-templ bundle --silent
bundle=web$(sed -n 's/.*bundleSrc = "\(.*\)"/\1/p' web/components/scripts_bundle.go)
[ -f "$bundle" ] || { echo "missing $bundle (named in web/components/scripts_bundle.go)" >&2; exit 1; }
echo "$bundle"

step "testing"
go test ./... | { grep -v 'no test files' || true; } # pipefail keeps go test's exit status

build() { # os arch out
	step "building $1/$2"
	GOOS=$1 GOARCH=$2 go build -trimpath -ldflags "-s -w -X main.version=$VERSION" -o "$3" "$pkg"
}

rm -rf "$out" && mkdir "$out"
if [ "${1:-}" != release ]; then
	exe=$app; [ "$(go env GOOS)" = windows ] && exe=$exe.exe
	build "$(go env GOOS)" "$(go env GOARCH)" "$out/$exe"
	step "done"
	ls -lh "$out"
	exit
fi

rm -rf "$stage"
for target in $targets; do
	os=${target%/*} arch=${target#*/}
	name=$app-$VERSION-$os-$arch
	exe=$app; [ "$os" = windows ] && exe=$exe.exe
	build "$os" "$arch" "$stage/$name/$exe"
	# shellcheck disable=SC2086 # $extras is a list
	cp $extras "$stage/$name/"
	if [ "$os" = windows ]; then
		(cd "$stage" && zip -qr "$OLDPWD/$out/$name.zip" "$name")
	else
		tar -C "$stage" -czf "$out/$name.tar.gz" "$name"
	fi
done
step "writing checksums"
(cd "$out" && sha256sum -- * > SHA256SUMS)
step "done"
ls -lh "$out"
