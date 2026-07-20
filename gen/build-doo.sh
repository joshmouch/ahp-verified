#!/usr/bin/env bash
# Builds the provider's exported library — agent-host-protocol-core.doo — carrying
# ALL SEVEN channel reducer modules (root + session + chat + terminal + changeset +
# annotations + resourceWatch) via spec/core_lib.dfy. This .doo is the ONLY interface
# any downstream repo (client / host / convergence / openagency-protocol-extensions)
# may depend on, consumed via `dafny … --library agent-host-protocol-core.doo` (a
# verification firewall the compiler enforces). Rebuild whenever any spec/<channel>.dfy
# (or the core_lib.dfy include list) changes.
set -euo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CONFLUX_DOO="$here/.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime.doo"
[ -f "$CONFLUX_DOO" ] || { echo "missing $CONFLUX_DOO — run conflux-workspace/scripts/setup-vendoring.sh ahp" >&2; exit 1; }
dafny build --target:lib --library "$CONFLUX_DOO" "$here/spec/core_lib.dfy" -o "$here/agent-host-protocol-core.doo"
echo "built $here/agent-host-protocol-core.doo ($(wc -c < "$here/agent-host-protocol-core.doo") bytes)"
