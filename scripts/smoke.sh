#!/usr/bin/env bash
# Smoke test: build golings the way a user gets it (go install into a temporary
# GOBIN), then run the first commands a learner types in a fresh folder and check
# their exit codes.
#
#   scripts/smoke.sh                       the usual machine
#   EXPECT_NO_RACE=1 scripts/smoke.sh      a machine without the race detector
#                                          (CGO_ENABLED=0 on Linux): the exercises
#                                          that need it are skipped, not failed
set -euo pipefail

repo=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
export NO_COLOR=1

gobin="$tmp/bin"
if command -v cygpath >/dev/null 2>&1; then
  gobin=$(cygpath -w "$gobin") # Git Bash on Windows: go wants a Windows path
fi

echo "== go install ./golings"
(cd "$repo" && GOBIN="$gobin" go install ./golings)
export PATH="$tmp/bin:$PATH"

fail=0
# check <expected exit code, or "nonzero"> <text the output must contain, or ""> <command...>
check() {
  local want=$1 text=$2
  shift 2
  local out="$tmp/out.txt" code=0
  "$@" >"$out" 2>&1 || code=$?
  echo "== $* (exit $code)"
  if [ "$want" = nonzero ]; then
    [ "$code" -ne 0 ] || { echo "FAIL: expected a non-zero exit code"; fail=1; }
  else
    [ "$code" -eq "$want" ] || { echo "FAIL: expected exit code $want"; fail=1; }
  fi
  if [ -n "$text" ] && ! grep -qi -- "$text" "$out"; then
    echo "FAIL: the output does not say '$text':"
    sed 's/^/    /' "$out"
    fail=1
  fi
}

mkdir "$tmp/learner"
cd "$tmp/learner"

check 0 "golings version" golings --version
check 0 "" golings init
cd golings-exercises
first=$(sed -n 's/^name = "\(.*\)"/\1/p' info.toml | head -1)
[ -n "$first" ] || { echo "FAIL: no exercise in info.toml"; exit 1; }

check 0 "Progress" golings list
check 0 "Go toolchain" golings doctor
# The first exercise is broken on purpose: running it has to fail, and say so.
check nonzero "Failed to compile" golings run "$first"

if [ "${EXPECT_NO_RACE:-}" = 1 ]; then
  # Without the race detector, doctor warns and says which exercises it skips,
  # and running one of them is a skip with a message, not a build failure.
  check 0 "Race detector" golings doctor
  grep -qi "race detector.*will be skipped" "$tmp/out.txt" || { echo "FAIL: doctor does not warn about the race detector"; fail=1; }
  check nonzero "race detector" golings run concurrent2
  if grep -qi "failed to compile\|tests failed" "$tmp/out.txt"; then
    echo "FAIL: concurrent2 was built and failed, it should have been skipped"
    fail=1
  fi
fi

if [ "$fail" -ne 0 ]; then
  echo "smoke test FAILED"
  exit 1
fi
echo "smoke test passed"
