#!/usr/bin/env bash
# Tags a release and pushes the tag, which runs .github/workflows/release.yml.
#   ./release.sh 0.2.0           (or v0.2.0)
#   ./release.sh 0.1.5 --force   allow a version older than the newest tag
set -Eeuo pipefail
cd "$(dirname "$0")"

die() { echo "release.sh: $*" >&2; exit 1; }
semver='^v[0-9]+\.[0-9]+\.[0-9]+$' # no pre-release suffixes
# newer a b: is version a above version b. Not sort -V: from PowerShell, Windows' sort.exe wins.
newer() {
	local -a x y; local i
	IFS=. read -ra x <<<"${1#v}"; IFS=. read -ra y <<<"${2#v}"
	for i in 0 1 2; do
		((10#${x[i]} == 10#${y[i]})) || { ((10#${x[i]} > 10#${y[i]})); return; }
	done
	return 1
}

force= v=
for arg; do
	case $arg in
	--force) force=1 ;;
	*) [ -z "$v" ] || die "one version only"; v=v${arg#v} ;;
	esac
done
[ -n "$v" ] || die "usage: ./release.sh <version> [--force]"
[[ $v =~ $semver ]] || die "$v is not a version like v1.2.3"

git fetch --verbose --tags origin
! git rev-parse -q --verify "refs/tags/$v" >/dev/null || die "$v already exists"
latest=$(git tag -l 'v*' --sort=-v:refname | grep -E "$semver" | head -1 || true)
if [ -n "$latest" ] && ! newer "$v" "$latest" && [ -z "$force" ]; then
	die "$v is not newer than $latest (use --force)"
fi
[ -z "$(git status --porcelain --untracked-files=no)" ] || die "commit or stash your changes first"
git merge-base --is-ancestor HEAD origin/main || die "HEAD is not on origin/main: push main first"

git tag -a "$v" -m "$v"
git push origin "$v"
echo "pushed $v: watch the release with 'gh run watch'"
