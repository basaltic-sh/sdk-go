#!/usr/bin/env bash
#
# Reports how this commit's exported API compares with the last release, and
# what version bump that implies.
#
# The bump is derived from the API itself rather than from commit messages.
# For an SDK generated from specifications, the specification decides whether a
# change is breaking — a removed field, a renamed operation — and whether
# someone wrote "feat:" or "fix:" is downstream of that, and often wrong.
#
#   scripts/apidiff.sh              report and exit 0
#   scripts/apidiff.sh --strict     exit 1 when the surface breaks
#
set -euo pipefail

MODULE="github.com/basaltic-sh/sdk-go"
STRICT=0
[[ "${1:-}" == "--strict" ]] && STRICT=1

command -v apidiff >/dev/null 2>&1 || {
    echo "apidiff not found. Install it with:"
    echo "  go install golang.org/x/exp/cmd/apidiff@latest"
    exit 1
}

PREV="$(git tag --list 'v*' --sort=-version:refname | head -n1)"
if [[ -z "$PREV" ]]; then
    echo "No release tag yet, so there is nothing to compare against."
    echo "The first release is v0.1.0."
    exit 0
fi

WORK="$(mktemp -d)"
BASE="$WORK/base.api"
OLD="$WORK/old"
trap 'git worktree remove --force "$OLD" >/dev/null 2>&1 || true; rm -rf "$WORK"' EXIT

git worktree add --detach --quiet "$OLD" "$PREV"
( cd "$OLD" && apidiff -m -w "$BASE" "$MODULE" )

REPORT="$(apidiff -m "$BASE" "$MODULE" || true)"

echo "Comparing the exported API against $PREV."
echo
if [[ -z "$REPORT" ]]; then
    echo "  No change to the exported API."
    echo "  Implied bump: PATCH."
    exit 0
fi
echo "$REPORT"
echo

if grep -q '^Incompatible changes:' <<<"$REPORT"; then
    CURRENT_MAJOR="$(sed 's/^v\([0-9]*\).*/\1/' <<<"$PREV")"
    echo "  The exported API BREAKS."
    if [[ "$CURRENT_MAJOR" == "0" ]]; then
        echo "  Implied bump: MINOR — v0 carries no compatibility promise, so a"
        echo "  break is allowed here. Say so in the release notes."
    else
        echo "  Implied bump: MAJOR."
        echo
        echo "  A major version in Go changes the MODULE PATH. v2 means the"
        echo "  go.mod line and every import in this repo become"
        echo "  $MODULE/v2. A bare v2.0.0 tag without that is a release"
        echo "  nothing can import. See RELEASING.md before tagging."
    fi
    [[ "$STRICT" == "1" ]] && exit 1
    exit 0
fi

echo "  The exported API only grew."
echo "  Implied bump: MINOR."
