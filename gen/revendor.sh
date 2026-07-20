#!/usr/bin/env bash
# Install the already-published Conflux verification and target-language package
# used by AHP. Package publication is a Conflux-owner operation: a consumer
# revendor must fail on producer or matrix drift instead of silently rebuilding
# and repointing the constellation while its own verification is in progress.
# Canonical Conflux .dfy source never enters this repository: verification uses
# conflux-runtime.doo, C# links the DLL through the props file, and JS links the
# CJS bundle using the target-specific translation record.
set -euo pipefail

here="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CONFLUX_WORKSPACE="${CONFLUX_WORKSPACE:-/Users/josh/Code/joshmouch/conflux/conflux-workspace}"
PACKAGE="$CONFLUX_WORKSPACE/packaging/runtime-package.mjs"
MATRIX="$CONFLUX_WORKSPACE/scripts/update-runtime-consumer-matrix.mjs"
INSTALL="$here/../scripts/vendor-conflux-package.sh"

[ -f "$PACKAGE" ] || { echo "Conflux package validator not found: $PACKAGE" >&2; exit 1; }
[ -f "$MATRIX" ] || { echo "Conflux consumer-matrix validator not found: $MATRIX" >&2; exit 1; }
[ -x "$INSTALL" ] || { echo "AHP package installer not found: $INSTALL" >&2; exit 1; }

node "$PACKAGE" check
node "$MATRIX" check
bash "$INSTALL" "$here"
node "$PACKAGE" check-runtime-installed "$here/.conflux/runtime/dependencies/conflux-runtime"

echo "installed immutable Conflux runtime package: verification .doo + C#/JS target artifacts; zero Conflux .dfy source"
