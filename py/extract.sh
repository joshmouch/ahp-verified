#!/usr/bin/env bash
# Regenerate the extracted Python core from the Dafny specification.
#
# The Dafny sources are the source of truth; everything under
# src/agent_host_protocol/_core/ is generated output and should never be
# edited by hand.
#
# Usage:  ./extract.sh [path-to-dafny-core]

set -euo pipefail

CORE="${1:-/Users/josh/Code/joshmouch/conflux/agent-host-protocol-workspace/agent-host-protocol-core-dafny}"
HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
OUT="$HERE/build/ahp_core"
DEST="$HERE/src/agent_host_protocol/_core"

RUNTIME="$CORE/.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime.doo"

# The conflux runtime is passed as a SOURCE INPUT, not via --library.
# --library assumes a prebuilt native runtime for the target; the vendored
# package ships those for cs (.dll) and js (.cjs) only, so for Python the
# runtime's Dafny modules must be translated alongside the spec.
#
# --include-runtime bundles the Dafny Python support modules (_dafny,
# System_) so the result has no external dependencies.

rm -rf "$OUT-py"
dafny translate py \
    "$CORE/spec/client_main.dfy" \
    "$RUNTIME" \
    --no-verify \
    --include-runtime \
    -o "$OUT"

rm -rf "$DEST"
mkdir -p "$DEST"
cp "$OUT-py"/*.py "$DEST/"
cp -R "$OUT-py/_dafny" "$DEST/"
cp -R "$OUT-py/System_" "$DEST/"
rm -f "$DEST/__main__.py"
find "$DEST" -name __pycache__ -type d -exec rm -rf {} + 2>/dev/null || true

echo "Extracted $(ls "$DEST"/*.py | wc -l | tr -d ' ') modules to $DEST"
python3 "$HERE/smoke_test.py"
