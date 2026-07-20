#!/usr/bin/env bash
# Durable regen-diff gate (rung 12) — mirror.generated.dfy vs the fork types/.
#
# Standalone, independently-runnable sibling of check-all.sh step 7/8. It
# REGENERATES the Dafny wire-type mirror from the fork's TypeScript `types/`
# into a SCRATCH file and byte-diffs it against the committed
# gen/mirror.generated.dfy, failing on any drift. Mirrors the AHP-TS
# `check-protocol-sync` convention: the committed generated artifact must always
# equal a fresh regen, so upstream type changes can never land silently.
#
# This does NOT reimplement generation — it invokes the one generator
# (gen/generate-dafny.ts), so there is a single source of drift-check truth.
#
# Usage:  AHP_REPO=/path/to/agent-host-protocol bash gen/check-mirror-sync.sh
# Exit:   0 = fresh (no drift) · 1 = DRIFT (regenerate + commit) · 2 = misconfigured
set -uo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"          # ahp-dafny/
# Which committed mirror to gate. There are TWO, one per pinned corpus, because a
# single generated file cannot be a fresh regen of two different types/ trees:
#   gen/mirror.generated.dfy         ← the pr-dotnet-client corpus (core lane;
#                                      no channels-canvas types). DEFAULT.
#   gen/mirror.canvas.generated.dfy  ← the PR #298 Canvas corpus (conformance
#                                      lane; run/canvas-conformance.sh gates it).
# Override with AHP_MIRROR, and pass the MATCHING AHP_REPO — pairing a mirror with
# the other corpus is a real drift, not a config error, and will (correctly) fail.
committed="${AHP_MIRROR:-$here/gen/mirror.generated.dfy}"
generator="$here/gen/generate-dafny.ts"

# The generator reads AHP_REPO (falling back to a legacy absolute path). Resolve
# and HARD-VALIDATE the fork here so a missing/mis-set AHP_REPO reports a config
# error (exit 2) instead of masquerading as drift: an absent types/ makes the
# generator emit an empty (0-type) mirror, which would byte-differ and look like
# a false drift. A regen-diff gate must never cry drift over a missing fork.
FORK="${AHP_REPO:-}"
if [ -z "$FORK" ]; then
  echo "[check-mirror-sync] AHP_REPO is unset — point it at an agent-host-protocol checkout." >&2
  echo "  e.g. AHP_REPO=/path/to/agent-host-protocol bash gen/check-mirror-sync.sh" >&2
  exit 2
fi
types_dir="$FORK/types"
if [ ! -d "$types_dir" ]; then
  echo "[check-mirror-sync] no types/ under AHP_REPO=$FORK — wrong fork path?" >&2
  exit 2
fi
# Require at least one channel type source so an empty/partial checkout can't pass
# or produce a false-drift 0-type mirror.
if ! ls "$types_dir"/**/*.ts >/dev/null 2>&1 && [ -z "$(find "$types_dir" -name '*.ts' -print -quit 2>/dev/null)" ]; then
  echo "[check-mirror-sync] AHP_REPO=$FORK has no *.ts under types/ — cannot regenerate." >&2
  exit 2
fi
if [ ! -f "$committed" ]; then
  echo "[check-mirror-sync] committed mirror missing: $committed" >&2
  exit 2
fi

# tsx + ts-morph resolve from the fork's node_modules (the generator's runtime).
tsx="$FORK/node_modules/.bin/tsx"
if [ ! -x "$tsx" ]; then
  echo "[check-mirror-sync] $tsx not found — run 'npm ci' in the fork so ts-morph/tsx resolve." >&2
  exit 2
fi

scratch="$(mktemp "${TMPDIR:-/tmp}/ahp-mirror-regen.XXXXXX.dfy")"
trap 'rm -f "$scratch"' EXIT

# Regenerate into the scratch path (normal write mode exercises the real emit path).
if ! (cd "$FORK" && AHP_REPO="$FORK" NODE_PATH="$FORK/node_modules" "$tsx" "$generator" --out "$scratch" >/dev/null 2>&1); then
  echo "[check-mirror-sync] generator failed against AHP_REPO=$FORK" >&2
  exit 2
fi

if cmp -s "$scratch" "$committed"; then
  echo "[check-mirror-sync] fresh: $committed matches a regen from $FORK/types/"
  exit 0
fi
echo "[check-mirror-sync] DRIFT: $committed differs from a fresh regen of $FORK/types/" >&2
echo "  regenerate + commit:  (cd \"$FORK\" && AHP_REPO=\"$FORK\" NODE_PATH=\"$FORK/node_modules\" node_modules/.bin/tsx \"$generator\" --out \"$committed\")" >&2
echo "  --- diff (committed → fresh), first 40 lines ---" >&2
diff "$committed" "$scratch" 2>/dev/null | head -40 >&2
exit 1
