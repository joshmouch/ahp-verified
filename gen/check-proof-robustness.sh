#!/usr/bin/env bash
# Phase 4 anti-vacuity / proof-robustness gate (ahp-dafny).
#
# Two independent anti-vacuity signals:
#  1. Every channel's universal no-op / precondition-fail lemma carries a
#     non-vacuity WITNESS (`*_NonVacuous` proving `exists a :: Precondition(a)`),
#     so the lemma cannot be vacuously true. This gate asserts one exists per
#     channel spec.
#  2. `dafny measure-complexity --mutations N` re-verifies every proof under N
#     distinct random SMT seeds; a proof that only passes by luck of the solver's
#     search fails here. This gate runs it over the full client and fails on any
#     verification error.
set -euo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SEEDS="${1:-5}"

echo "== 1. non-vacuity witnesses per channel =="
missing=0
for f in ahp_skeleton resource_watch changeset annotations terminal session chat; do
  n=$(grep -cE 'NonVacuous' "$here/spec/$f.dfy" 2>/dev/null || echo 0)
  printf "  %-16s %s witness(es)\n" "$f" "$n"
  [ "$n" -ge 1 ] || { echo "  MISSING non-vacuity witness in $f.dfy"; missing=1; }
done
[ "$missing" -eq 0 ] || { echo "ANTI-VACUITY FAIL — a channel lacks a non-vacuity witness"; exit 1; }

echo "== 2. proof robustness: measure-complexity --mutations $SEEDS =="
out="$(dafny measure-complexity --mutations "$SEEDS" --library "$here/.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime.doo" "$here/spec/client_main.dfy" 2>&1)"
if echo "$out" | grep -qiE 'error|inconsistent|out of resource|timed out|not verified'; then
  echo "PROOF-ROBUSTNESS FAIL — a proof failed under a mutation seed:"
  echo "$out" | grep -iE 'error|inconsistent|resource|timed out|not verified' | head
  exit 1
fi
echo "  proofs verify consistently across $SEEDS seeds"
echo "ANTI-VACUITY OK — every channel has a non-vacuity witness; proofs robust across $SEEDS seeds"
