#!/usr/bin/env bash
#
# Regenerate the C++-consumable native library from the verified Dafny core.
#
# Route: Dafny -> Rust (`dafny translate rs`) -> staticlib/cdylib with a C ABI ->
# consumed by C++. The Dafny C++ backend cannot compile this core (see
# ../cpp/README.md and evidence/03-nightly-cpp-and-rs.txt); the Rust backend can,
# and Rust exposes a C ABI natively.
#
# Everything below is generated. The only hand-written sources in this package
# are:
#   ahp-core-ffi/src/ffi.rs             -- the C ABI binding layer
#   ahp-core-ffi/src/_dafny_externs.rs  -- host-capability boundary stubs
#   include/ahp_core.h                  -- the C header
#   examples/ahp_terminal_demo.cpp      -- the C++ test program
#   scripts/patch-determinism.py        -- the one required source rewrite
#
# Copyright (c) Microsoft Corporation
# Copyright (c) 2026 Josh Mouch
# SPDX-License-Identifier: MIT

set -euo pipefail

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

CORE="${AHP_CORE:-/Users/josh/Code/joshmouch/conflux/agent-host-protocol-workspace/agent-host-protocol-core-dafny}"
DOO="${AHP_LIB:-$CORE/.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime.doo}"

STAGE="$HERE/build/stage"
OUT="$HERE/build/ahpcore"

command -v dafny >/dev/null || { echo "dafny not on PATH" >&2; exit 1; }
command -v cargo >/dev/null || { echo "cargo not on PATH" >&2; exit 1; }
command -v unzip >/dev/null || { echo "unzip not on PATH" >&2; exit 1; }
[ -f "$CORE/spec/client_main.dfy" ] || { echo "core not found at $CORE" >&2; exit 1; }
[ -f "$DOO" ] || { echo "conflux runtime .doo not found at $DOO" >&2; exit 1; }

echo "==> dafny: $(dafny --version)"
mkdir -p "$STAGE"

# ---------------------------------------------------------------------------
# 1. Recover the conflux runtime's Dafny SOURCE from its .doo.
#
# A .doo is a zip whose `program` member is the runtime's Dafny source. We need
# the source (not the prebuilt .doo) because the .doo was built with
# enforce-determinism=false and the Rust backend mandates it -- passing the .doo
# directly fails with:
#   "--enforce-determinism is set locally to True, but the library was built
#    with False"
# ---------------------------------------------------------------------------
echo "==> extracting conflux runtime source from $(basename "$DOO")"
rm -rf "$STAGE/doo" && mkdir -p "$STAGE/doo"
unzip -q -o "$DOO" -d "$STAGE/doo"
cp "$STAGE/doo/program" "$STAGE/conflux-runtime.dfy"

# ---------------------------------------------------------------------------
# 2. The single source rewrite the Rust backend requires.
#
# --enforce-determinism rejects the assign-such-that STATEMENT form
# unconditionally. The runtime contains exactly one in compiled code
# (ConfluxJsonText.SortedKeySetRuntime; the other 25 occurrences are in ghost
# lemmas, which are erased). The patch lifts it into a compiled function using
# the let-such-that EXPRESSION form, which the flag permits, carrying the
# identical predicate. No AHP reducer is touched.
# ---------------------------------------------------------------------------
echo "==> applying determinism rewrite"
python3 "$HERE/scripts/patch-determinism.py" "$STAGE/conflux-runtime.dfy"

# ---------------------------------------------------------------------------
# 3. RE-VERIFY. This is the load-bearing step: it proves the rewrite preserved
#    every postcondition, so the translated reducers remain proven code.
# ---------------------------------------------------------------------------
echo "==> re-verifying patched runtime (this is the proof the rewrite is sound)"
dafny verify "$STAGE/conflux-runtime.dfy" --allow-warnings

# ---------------------------------------------------------------------------
# 4. Translate the verified core to Rust.
# ---------------------------------------------------------------------------
echo "==> translating core to Rust"
rm -rf "$OUT-rust"
dafny translate rs "$CORE/spec/client_main.dfy" "$STAGE/conflux-runtime.dfy" \
  --no-verify \
  --enforce-determinism \
  --include-runtime \
  --allow-warnings \
  -o "$OUT"

# ---------------------------------------------------------------------------
# 5. Stage the generated Rust into the FFI crate.
#
# The generated module tree is included at the crate root (the Dafny Rust
# backend emits absolute `crate::<Module>::...` paths), so its three leading
# inner attributes must move to lib.rs.
# ---------------------------------------------------------------------------
echo "==> staging generated Rust into ahp-core-ffi"
cp "$OUT-rust/src/ahpcore.rs" "$HERE/ahp-core-ffi/src/ahpcore.rs"
rm -rf "$HERE/ahp-core-ffi/dafny_runtime"
cp -R "$OUT-rust/runtime" "$HERE/ahp-core-ffi/dafny_runtime"

python3 - "$HERE/ahp-core-ffi/src/ahpcore.rs" <<'PY'
import sys
path = sys.argv[1]
lines = open(path).read().split("\n")
removed = 0
while lines and lines[0].startswith("#!["):
    lines.pop(0)
    removed += 1
open(path, "w").write("\n".join(lines))
print(f"    moved {removed} inner attribute(s) to lib.rs")
PY

# ---------------------------------------------------------------------------
# 6. Build the native library and run the C++ program against it.
# ---------------------------------------------------------------------------
echo "==> building native library"
cargo build --release --manifest-path "$HERE/ahp-core-ffi/Cargo.toml"

echo "==> building + running the C++ test program"
mkdir -p "$HERE/build"
clang++ -std=c++17 -O2 -Wall -Wextra \
  -I"$HERE/include" \
  "$HERE/examples/ahp_terminal_demo.cpp" \
  "$HERE/ahp-core-ffi/target/release/libahp_core.a" \
  -o "$HERE/build/ahp_terminal_demo"

"$HERE/build/ahp_terminal_demo"

echo
echo "==> building + running the full conformance corpus"
clang++ -std=c++17 -O2 -Wall -Wextra \
  -I"$HERE/include" \
  "$HERE/examples/ahp_corpus.cpp" \
  "$HERE/ahp-core-ffi/target/release/libahp_core.a" \
  -o "$HERE/build/ahp_corpus"

"$HERE/build/ahp_corpus"

echo
echo "==> done"
echo "    static lib : $HERE/ahp-core-ffi/target/release/libahp_core.a"
echo "    shared lib : $HERE/ahp-core-ffi/target/release/libahp_core.dylib"
echo "    header     : $HERE/include/ahp_core.h"
