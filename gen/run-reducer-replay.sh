#!/usr/bin/env bash
# Runs the Phase-2 REDUCER-REPLAY harnesses: real fixture JSON → decode → the
# actual reducer → assert reduced == expected, on BOTH extracted targets.
#
# Unlike the hand-transcribed corpus in each channel's RunCorpus, this drives the
# reducers from the REAL fixture files. root is the proven proof-of-pattern; the
# other channels follow the same decoder idiom (spec/replay/<channel>_replay.dfy).
set -euo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"          # ahp-dafny/
# Fixture source. Defaults to the corpus VENDORED INTO THIS REPOSITORY at corpus/reducers/
# (232 fixtures, all upstream-authored, byte-identical to their upstream blobs — see
# corpus/PROVENANCE.md). Pinning the corpus in-tree is what makes the per-channel
# denominators below checkable by a reader who has no upstream checkout.
#
# Set AHP_REPO to re-run against a LIVE upstream checkout instead; that is the drift
# check, and the denominators are expected to move when upstream adds fixtures.
if [ -n "${AHP_REPO:-}" ]; then
  fixtures_dir="$AHP_REPO/types/test-cases/reducers"
  echo "[corpus] live upstream checkout: $fixtures_dir"
else
  fixtures_dir="$here/corpus/reducers"
  echo "[corpus] pinned in-tree corpus: $fixtures_dir"
fi
[ -d "$fixtures_dir" ] || { echo "fixture directory not found: $fixtures_dir" >&2; exit 1; }
repo_node_modules="$here/node_modules"
runner="$here/../scripts/run-conflux-target.sh"
mkdir -p "$here/build"

run_channel() {  # $1 = channel name, $2 = fixture glob
  local ch="$1" glob="$2"
  local harness="$here/spec/replay/${ch}_replay.dfy"
  [ -f "$harness" ] || { echo "### ${ch}: harness not present — skipping"; return 0; }
  local fx
  if [[ "$glob" = /* ]]; then fx=($glob); else fx=($fixtures_dir/$glob); fi
  echo "### ${ch}: ${#fx[@]} fixtures"
  echo "--- C# ---"
  # Use the one package-aware target runner for C# as well as JavaScript. It
  # assigns each invocation its own generated/DTR/MSBuild paths, so concurrent
  # delegated accounts never compile or overwrite another gate's replay output.
  local target_output
  if ! target_output="$("$runner" cs "$here" "$harness" "$here/build/${ch}_replay" \
      --standard-libraries --translate-standard-library -- "${fx[@]}" 2>&1)"; then
    printf '%s\n' "$target_output" >&2
    return 1
  fi
  printf '%s\n' "$target_output" | grep -E 'REDUCER-REPLAY|REPLAY OK|FAIL'
  echo "--- JS ---"
  if ! target_output="$(NODE_PATH="$repo_node_modules" "$runner" js "$here" "$harness" "$here/build/${ch}_replay" \
      --standard-libraries --translate-standard-library -- "${fx[@]}" 2>&1)"; then
    printf '%s\n' "$target_output" >&2
    return 1
  fi
  printf '%s\n' "$target_output" | grep -E 'REDUCER-REPLAY|REPLAY OK|FAIL'
}

run_channel root '*root*.json'
# Fully-covered channels (same decoder idiom as root):
run_channel resourcewatch '*resourcewatch*.json'
run_channel canvas "$here/spec/replay/fixtures/canvas-*.json"
run_channel changeset '*changeset*.json'
run_channel annotations '*annotation*.json'
run_channel terminal '*terminal*.json'
# Richest channels (recursive decoders + modifiedAt normalization). These
# harnesses self-filter by the fixture 'reducer' field, so feed them the FULL
# corpus — most chat/session reducer fixtures are numbered/cross-named, not
# '*chat*'/'*session*' (e.g. tool-call fixtures like 019-tool-call-*.json).
run_channel session '*.json'
run_channel chat '*.json'
