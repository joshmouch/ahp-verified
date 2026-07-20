#!/usr/bin/env bash
# Assert that the retired AHP codec correspondence bridges remain absent. Live codecs import
# ConfluxCodec directly, and gen/check-all.sh verifies their current round-trip proofs.
# Run: bash gen/verify-codec-bridges.sh
set -uo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

shopt -s nullglob
bridges=("$here"/spec/*-codec-is-codec.dfy)
shopt -u nullglob
if [ "${#bridges[@]}" -eq 0 ]; then
  echo "PASS — retired codec bridges remain absent"
  exit 0
fi

echo "[FAIL] retired codec bridges reappeared:" >&2
for bridge in "${bridges[@]}"; do printf '  %s\n' "$bridge" >&2; done
exit 1
