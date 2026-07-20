#!/usr/bin/env bash
#
# Regenerate the translated Go sources from the verified Dafny core.
#
# This reproduces every step used to build this module, including the two
# post-translation fixes that Dafny 4.11.0 requires for the Go backend. It
# rewrites only generated code; the hand-written adapters, tests and docs are
# left alone.
#
# Copyright (c) Microsoft Corporation
# Copyright (c) 2026 Josh Mouch
# SPDX-License-Identifier: MIT

set -euo pipefail

MODULE="github.com/joshmouch/ahp-go"
HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

CORE="${AHP_CORE:-/Users/josh/Code/joshmouch/conflux/agent-host-protocol-workspace/agent-host-protocol-core-dafny}"
LIB="${AHP_LIB:-$CORE/.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime.doo}"
STAGE="$(mktemp -d)"
trap 'rm -rf "$STAGE"' EXIT

command -v dafny    >/dev/null || { echo "dafny not on PATH"    >&2; exit 1; }
command -v go       >/dev/null || { echo "go not on PATH"       >&2; exit 1; }
# Dafny's Go backend post-processes its output with goimports.
command -v goimports >/dev/null || {
  echo "goimports not on PATH; install with:" >&2
  echo "  go install golang.org/x/tools/cmd/goimports@latest" >&2
  exit 1
}

echo "==> translating $CORE/spec/client_main.dfy"
# The conflux runtime is passed as an INPUT rather than via --library: --library
# suppresses emission, and unlike C#/JS there is no prebuilt Go runtime binary to
# link against, so its ~50 packages must be translated too.
dafny translate go "$CORE/spec/client_main.dfy" "$LIB" \
  --no-verify \
  --include-runtime \
  --go-module-name "$MODULE" \
  -o "$STAGE/out"

SRC="$STAGE/out-go"

echo "==> fix 1: module-qualify the bundled runtime's import of 'dafny'"
# --go-module-name rewrites imports in translated code but not in the runtime
# files copied in by --include-runtime, which keep GOPATH-style bare imports.
sed -i '' "s|_dafny \"dafny\"|_dafny \"$MODULE/dafny\"|" "$SRC/System_/System_.go"

echo "==> fix 2: un-shadow the 'Cap' parameter in ConfluxWatermark"
# Dafny emits a Create_Cap_ whose parameter name collides with its return type.
python3 - "$SRC/ConfluxWatermark/ConfluxWatermark.go" <<'PY'
import sys
path = sys.argv[1]
src = open(path).read()
old = '''func (CompanionStruct_Cap_) Create_Cap_(Pos _dafny.Int, Cap _dafny.Int) Cap {
	return Cap{Cap_Cap{Pos, Cap}}
}'''
new = '''func (CompanionStruct_Cap_) Create_Cap_(Pos _dafny.Int, Cap_ _dafny.Int) Cap {
	return Cap{Cap_Cap{Pos, Cap_}}
}'''
if old not in src:
    sys.exit("ConfluxWatermark Cap-shadowing pattern not found; "
             "the Dafny backend may have fixed it — re-check this step.")
open(path, 'w').write(src.replace(old, new, 1))
PY

echo "==> refreshing generated packages in $HERE"
# Replace generated trees only; adapters live inside three of these directories,
# so those are preserved explicitly.
for dir in "$SRC"/*/; do
  name="$(basename "$dir")"
  rm -f "$HERE/$name"/*.go 2>/dev/null || true
  mkdir -p "$HERE/$name"
  cp "$dir"/*.go "$HERE/$name/"
done
git -C "$HERE" checkout -- \
  ConfluxCommandIdentityCapability/adapter.go \
  ConfluxIoCapability/adapter.go \
  ConfluxHttpCapability/adapter.go 2>/dev/null || true

mkdir -p "$HERE/cmd/ahp"
cp "$SRC"/*.go "$HERE/cmd/ahp/main.go"

echo "==> building"
# -gcflags=-l disables inlining. Without it the Go compiler's inliner balloons to
# tens of gigabytes on the generated Chat package and is OOM-killed; optimization
# is otherwise left on.
cd "$HERE"
go build -gcflags 'all=-l' ./...
go test -gcflags 'all=-l' ./...

echo "==> done"
