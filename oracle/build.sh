#!/usr/bin/env bash
# build.sh -- build ahp-oracle and self-test it against the real corpus.
#
# The oracle is a thin host over the extracted, machine-checked Dafny core
# (../../cs/src/Ahp.Core.Verified). This script:
#   1. builds it,
#   2. publishes a self-contained single-file binary (no .NET runtime needed),
#   3. runs the whole corpus through it (must AGREE 232/232), and
#   4. runs the falsifiability sweep (must catch every state-changing mutation).
#
# Steps 3 and 4 are the gate: a build that produces a binary which disagrees
# with the corpus, or which fails to catch a real mutation, exits non-zero.
#
# Usage:  ./build.sh [runtime-id]
#   runtime-id defaults to the host RID. Cross-publish with e.g.
#   linux-x64, win-x64, osx-x64, linux-arm64 (only the host RID is self-tested).
set -euo pipefail

HERE="$(cd "$(dirname "$0")" && pwd)"
PROJ="$HERE/src/AhpOracle/AhpOracle.csproj"
CORPUS="$HERE/corpus/reducers"

RID="${1:-$(dotnet --info 2>/dev/null | awk -F': *' '/^ *RID:/{print $2; exit}')}"
OUT="$HERE/dist/$RID"

echo "==> 1/4 build (Release)"
dotnet build "$PROJ" -c Release -clp:ErrorsOnly

echo "==> 2/4 publish self-contained single-file binary ($RID) -> $OUT"
dotnet publish "$PROJ" -c Release -r "$RID" \
  --self-contained true -p:PublishSingleFile=true -p:EnableCompressionInSingleFile=true \
  -o "$OUT"

BIN="$OUT/ahp-oracle"
[ -x "$BIN" ] || { echo "FAIL: $BIN not produced"; exit 1; }
echo "    binary: $(du -h "$BIN" | cut -f1)  $(file -b "$BIN")"

HOST_RID="$(dotnet --info 2>/dev/null | awk -F': *' '/^ *RID:/{print $2; exit}')"
if [ "$RID" != "$HOST_RID" ]; then
  echo "==> 3-4/4 SKIPPED self-test: cross-published $RID cannot run on host $HOST_RID"
  echo "    (built and published; run the self-test on a $RID machine)"
  exit 0
fi

echo "==> 3/4 corpus self-test (must AGREE 232/232)"
"$BIN" corpus "$CORPUS"

echo "==> 4/4 falsifiability sweep (must catch every state-changing mutation)"
python3 "$HERE/demo/falsify.py" "$BIN" "$CORPUS"

echo "==> OK: $BIN"
