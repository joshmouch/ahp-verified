#!/usr/bin/env bash
#
# Regenerates the native AHP library and C header from the Dafny-extracted Go
# core, then compiles and runs the C++ example against them.
#
# Everything under lib/ and include/libahpverified.h is a build product of this
# script; the only hand-written sources in this package are bridge/ahp_bridge.go,
# include/ahp_verified.h and examples/ahp_demo.cpp.
#
# Copyright (c) Microsoft Corporation
# Copyright (c) 2026 Josh Mouch
# SPDX-License-Identifier: MIT

set -euo pipefail

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BRIDGE="$HERE/bridge"
LIB="$HERE/lib"
BUILD="$HERE/build"
INCLUDE="$HERE/include"

# The Dafny-extracted Chat reducer is large enough that the Go compiler's
# inliner is killed by the OS while compiling it. Inlining must be disabled for
# every package in the build (`all=`), not just the one named on the command
# line, or the dependency build still dies. Serial compilation keeps peak
# memory bounded.
GOFLAGS_VERIFIED=(-p=1 -gcflags=all=-l)

step() { printf '\n== %s\n' "$1"; }

step "Checking the toolchain"
for tool in go clang++; do
  command -v "$tool" >/dev/null || { echo "error: $tool not found on PATH" >&2; exit 1; }
done
go version
clang++ --version | head -1

step "Building the static archive (c-archive)"
mkdir -p "$LIB" "$INCLUDE" "$BUILD"
( cd "$BRIDGE" && go build "${GOFLAGS_VERIFIED[@]}" \
    -buildmode=c-archive -o "$LIB/libahpverified.a" . )

step "Building the shared library (c-shared)"
case "$(uname -s)" in
  Darwin) SHARED_EXT=dylib ;;
  *)      SHARED_EXT=so    ;;
esac
( cd "$BRIDGE" && go build "${GOFLAGS_VERIFIED[@]}" \
    -buildmode=c-shared -o "$LIB/libahpverified.$SHARED_EXT" . )

# Both build modes emit the same header; keep one copy, in include/.
mv -f "$LIB/libahpverified.h" "$INCLUDE/libahpverified.h"
rm -f "$LIB/libahpverified.h"

step "Verifying include/ahp_verified.h matches the generated ABI"
# include/ahp_verified.h is hand-written so that it can be const-correct, which
# the cgo-generated header is not. That makes drift possible, so every symbol
# it declares must be present in the generated header and vice versa.
symbols_of() { grep -oE '\bahp_[a-z0-9_]+\(' "$1" | tr -d '(' | sort -u; }
declared="$(symbols_of "$INCLUDE/ahp_verified.h")"
generated="$(symbols_of "$INCLUDE/libahpverified.h")"
if [ "$declared" != "$generated" ]; then
  echo "error: include/ahp_verified.h has drifted from the generated ABI" >&2
  diff <(echo "$declared") <(echo "$generated") \
    --label "declared in ahp_verified.h" --label "exported by the bridge" -u >&2 || true
  exit 1
fi
echo "$(echo "$declared" | wc -l | tr -d ' ') exported symbols agree"

step "Compiling the C++ example"
clang++ -std=c++17 -O2 -Wall -Wextra \
  -I"$INCLUDE" "$HERE/examples/ahp_demo.cpp" "$LIB/libahpverified.a" \
  -o "$BUILD/ahp_demo"

step "Running the C++ example"
"$BUILD/ahp_demo"

step "Artifacts"
ls -la "$LIB" "$INCLUDE" "$BUILD"
