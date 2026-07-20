#!/usr/bin/env bash
#
# Reproduces the Ahp.Core.Verified release artifact from the Dafny core.
#
# Chain: dafny translate -> dotnet build -> smoke test -> corpus -> dotnet pack
#
# Copyright (c) Microsoft Corporation.
# Copyright (c) 2026 Josh Mouch.
# Licensed under the MIT License.

set -euo pipefail

CORE="/Users/josh/Code/joshmouch/conflux/agent-host-protocol-workspace/agent-host-protocol-core-dafny"
LIB="$CORE/.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime.doo"
HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "==> 1/5 translate core_lib.dfy -> C# (self-contained)"
# The runtime .doo is passed as a normal input rather than via --library so its
# code is emitted inline; --include-runtime folds in the Dafny C# runtime. The
# result is a single assembly with no external package dependencies.
(cd "$CORE" && dafny translate cs spec/core_lib.dfy "$LIB" \
    --no-verify --include-runtime \
    -o "$HERE/build/ahp_core_full")

cp "$HERE/build/ahp_core_full.cs" "$HERE/src/Ahp.Core.Verified/Generated/AhpCore.g.cs"

echo "==> 2/5 build the class library"
dotnet build "$HERE/src/Ahp.Core.Verified" -c Release

echo "==> 3/5 smoke test the extracted core"
dotnet run --project "$HERE/test/Ahp.Core.Verified.SmokeTest" -c Release

echo "==> 4/5 run the AHP reducer corpus against the extracted C#"
(cd "$CORE" && dafny translate cs spec/client_main.dfy "$LIB" \
    --no-verify --include-runtime \
    -o "$HERE/test/CorpusRunner/CorpusRunner")
dotnet run --project "$HERE/test/CorpusRunner" -c Release

echo "==> 5/5 pack"
dotnet pack "$HERE/src/Ahp.Core.Verified" -c Release -o "$HERE/artifacts"

echo
echo "Done. Artifacts:"
ls -la "$HERE/artifacts"
