#!/usr/bin/env bash
# Runs the Std.JSON round-trip fidelity harness (spec/codec/json_roundtrip.dfy)
# over the full real corpus on BOTH extracted targets (cs + js).
#
# Q4 codec proof: every real AHP fixture survives parse→serialize→re-parse with
# an identical AST (unknown keys + key order preserved), proving the Std.JSON
# codec trusts only raw byte I/O and satisfies the open/closed fidelity law.
set -euo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"          # ahp-dafny/
harness="$here/spec/codec/json_roundtrip.dfy"
# Same corpus resolution as gen/run-reducer-replay.sh: the pinned in-tree corpus by
# default, a live upstream checkout when AHP_REPO is set. See corpus/PROVENANCE.md.
if [ -n "${AHP_REPO:-}" ]; then
  fixtures_dir="$AHP_REPO/types/test-cases/reducers"
else
  fixtures_dir="$here/corpus/reducers"
fi
[ -d "$fixtures_dir" ] || { echo "fixture directory not found: $fixtures_dir" >&2; exit 1; }
repo_node_modules="$here/node_modules"
runner="$here/../scripts/run-conflux-target.sh"

# Collect every corpus fixture path as args (bash 3.2-compatible glob array).
fixtures=("$fixtures_dir"/*.json)
echo "Replaying ${#fixtures[@]} corpus fixtures through Std.JSON round-trip…"

echo "=== C# ==="
"$runner" cs "$here" "$harness" "$here/build/json_roundtrip_cs" \
  --standard-libraries --translate-standard-library -- "${fixtures[@]}" \
  | grep -E 'UNKNOWN-KEY|ROUND-TRIP|CODEC|FAIL'

echo "=== JS ==="
NODE_PATH="$repo_node_modules" "$runner" js "$here" "$harness" "$here/build/json_roundtrip" \
  --standard-libraries --translate-standard-library -- "${fixtures[@]}" \
  | grep -E 'UNKNOWN-KEY|ROUND-TRIP|CODEC|FAIL'
