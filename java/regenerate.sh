#!/usr/bin/env bash
# Regenerate src/main/java from the Dafny core.
#
# Copyright (c) Microsoft Corporation.
# Copyright (c) 2026 Josh Mouch.
# Licensed under the MIT License.
#
# Everything under src/main/java/agency/open/ahp/ is generated, with the single
# exception of CorpusRunner.java, which is hand-written and preserved by this script.

set -euo pipefail

CORE_DIR="${AHP_CORE_DIR:?set AHP_CORE_DIR to the agent-host-protocol-core-dafny checkout}"
LIBRARY="${AHP_LIBRARY:-$CORE_DIR/.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime.doo}"
OUTER_MODULE="agency.open.ahp"

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BUILD="$HERE/.regen"
DEST="$HERE/src/main/java"

command -v dafny >/dev/null || { echo "dafny not on PATH" >&2; exit 1; }
[ -f "$LIBRARY" ] || { echo "library not found: $LIBRARY" >&2; exit 1; }

echo "dafny:   $(dafny --version)"
echo "library: $LIBRARY"

rm -rf "$BUILD"
mkdir -p "$BUILD"

# Pass 1 — the conflux-runtime library. It must be translated separately: --library
# tells Dafny the library is already compiled, so its code is NOT emitted. Upstream
# ships prebuilt .dll/.cjs artifacts for C# and JavaScript, but none for the JVM, so
# without this pass the core references a ConfluxCodec package that does not exist.
echo "==> translating conflux-runtime library"
dafny translate java "$LIBRARY" \
  --no-verify --outer-module "$OUTER_MODULE" -o "$BUILD/lib"

# Pass 2 — the AHP core, linked against that library. Same --outer-module so the
# cross-references resolve to the same packages.
echo "==> translating AHP core"
( cd "$CORE_DIR" && dafny translate java spec/client_main.dfy \
    --no-verify --outer-module "$OUTER_MODULE" --library "$LIBRARY" -o "$BUILD/core" )

echo "==> merging into $DEST"
CORPUS_RUNNER="$(mktemp)"
cp "$DEST/agency/open/ahp/CorpusRunner.java" "$CORPUS_RUNNER"
rm -rf "$DEST/agency"
cp -R "$BUILD/lib-java/agency" "$DEST/"
cp -R "$BUILD/core-java/agency" "$DEST/"
cp "$CORPUS_RUNNER" "$DEST/agency/open/ahp/CorpusRunner.java"
rm -f "$CORPUS_RUNNER"
rm -rf "$BUILD"

# Prune to the dependency closure of the AHP reducers.
#
# The translated conflux-runtime also contains a service layer (ConfluxHttpService,
# ConfluxJsonRpc, ConfluxResourceSupervisor, ...) that bottoms out in {:extern}
# capability modules. Upstream ships extern bodies for C# and JavaScript only, so
# those modules cannot be built for the JVM. Nothing in the AHP protocol core needs
# them, so they are dropped rather than stubbed.
echo "==> pruning to AHP core closure"
KEEP="AhpSkeleton Annotations Canvas Changeset Chat ClientMain ConfluxCodec \
ConfluxContract ConfluxOperators ConfluxOrderedLog ConfluxPrune ConfluxSeqRoute \
ResourceWatch Session Terminal"
( cd "$DEST/agency/open/ahp"
  for d in */; do
    d="${d%/}"
    case " $KEEP " in *" $d "*) ;; *) rm -rf "$d" ;; esac
  done )

# Fail loudly if the closure is not self-contained.
DANGLING="$( cd "$DEST/agency/open/ahp" \
  && comm -23 \
       <(grep -rhoE 'agency\.open\.ahp\.[A-Z][A-Za-z0-9_]*' . --include='*.java' \
           | sed 's/.*ahp\.//' | sort -u) \
       <(ls -d */ | tr -d '/' | sort) )"
if [ -n "$DANGLING" ]; then
  echo "ERROR: closure references missing packages:" >&2
  echo "$DANGLING" >&2
  exit 1
fi

echo "==> done: $(find "$DEST" -name '*.java' | wc -l | tr -d ' ') java files"
echo "Build with a modern JDK (21+; 26 verified): mvn package"
