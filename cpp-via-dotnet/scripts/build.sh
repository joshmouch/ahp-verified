#!/usr/bin/env bash
#
# Regenerates the C++ distribution of the verified AHP core from scratch:
#   1. NativeAOT-compiles the Dafny-extracted verified core into a native
#      shared library exporting a C ABI.
#   2. Stages the library and header into the package layout.
#   3. Builds and runs the C++ demo against it.
#
# Usage:  scripts/build.sh [--cmake]
#           (default)  build the demo directly with clang++
#           --cmake    build the demo via CMake instead
#
# Copyright (c) Microsoft Corporation.
# Copyright (c) 2026 Josh Mouch.
# Licensed under the MIT License.

set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

RID="osx-arm64"
TFM="net10.0"
LIB_NAME="libahp_core.dylib"
PUBLISH_DIR="src/bin/Release/${TFM}/${RID}/publish"

USE_CMAKE=0
[[ "${1:-}" == "--cmake" ]] && USE_CMAKE=1

echo "==> 1/4  Checking toolchain"
command -v dotnet  >/dev/null || { echo "error: dotnet not found"  >&2; exit 1; }
command -v clang++ >/dev/null || { echo "error: clang++ not found" >&2; exit 1; }
echo "    dotnet  $(dotnet --version)"
echo "    clang++ $(clang++ --version | head -1)"

echo "==> 2/4  NativeAOT-compiling the verified core"
rm -rf src/bin src/obj build lib
dotnet publish -c Release src/Ahp.Core.Native.csproj

if [[ ! -f "${PUBLISH_DIR}/${LIB_NAME}" ]]; then
  echo "error: expected ${PUBLISH_DIR}/${LIB_NAME} to exist" >&2
  exit 1
fi

echo "==> 3/4  Staging package layout"
mkdir -p lib build
cp "${PUBLISH_DIR}/${LIB_NAME}" "lib/${LIB_NAME}"
echo "    lib/${LIB_NAME}  ($(wc -c < "lib/${LIB_NAME}" | tr -d ' ') bytes)"
echo "    exported symbols:"
nm -gU "lib/${LIB_NAME}" | grep ' _ahp_' | awk '{print "      " $3}' | sort

echo "==> 4/4  Building and running the C++ demo"
if [[ $USE_CMAKE -eq 1 ]]; then
  command -v cmake >/dev/null || { echo "error: cmake not found" >&2; exit 1; }
  cmake -S . -B build -DCMAKE_BUILD_TYPE=Release
  cmake --build build
  ./build/ahp_demo
else
  clang++ -std=c++17 -Wall -Wextra -O2 \
    -Iinclude examples/ahp_demo.cpp "lib/${LIB_NAME}" \
    -Wl,-rpath,@executable_path/../lib \
    -o build/ahp_demo
  ./build/ahp_demo
fi

echo
echo "Done. Artifacts:"
echo "  lib/${LIB_NAME}"
echo "  include/ahp_core.h"
echo "  build/ahp_demo"
