#!/usr/bin/env bash
# Everything a third party can check, with no private dependency.
#
# gen/check-all.sh re-runs the PROOFS, and that needs the Conflux runtime package,
# which is not published (see REPRODUCIBILITY.md). This script deliberately needs
# none of it. It checks the artifacts that ARE in this repository:
#
#   1. the pinned fixture corpus is intact and has the composition claimed for it
#   2. the extracted code shipped in each language actually runs that corpus green
#
# WHAT THIS PROVES: the extracted reducers in this repository agree with the fixture
# corpus in this repository.
#
# WHAT THIS DOES NOT PROVE: that the extracted code was produced from proofs that
# verify. That is the claim requiring the private dependency, and no amount of
# running this script substitutes for it. Do not cite this gate as proof of
# verification — cite it as proof of extraction-vs-corpus agreement.
#
# Each language leg is skipped (loudly, and counted) if its toolchain is absent.
# A skipped leg is never reported as a pass.
set -uo pipefail

here="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
fail=0
skipped=0
step() { printf "\n══ %s ══\n" "$1"; }
ok()   { echo "  ✔ $1"; }
bad()  { echo "  [FAIL]: $1"; fail=1; }
skip() { echo "  [SKIP]: $1"; skipped=$((skipped + 1)); }
have() { command -v "$1" >/dev/null 2>&1; }

# Every language ships the same embedded per-action corpus with these counts.
EXPECTED_TOTAL=148

step "1/6 pinned fixture corpus"
corpus_dir="$here/corpus/reducers"
if [ ! -d "$corpus_dir" ]; then
  bad "corpus directory missing: $corpus_dir"
else
  n="$(find "$corpus_dir" -name '*.json' | wc -l | tr -d ' ')"
  if [ "$n" = 232 ]; then ok "232 fixture files present"; else bad "expected 232 fixture files, found $n"; fi

  if have python3; then
    # Composition is part of the claim: the per-channel split is what the replay
    # denominators in gen/run-reducer-replay.sh are pinned to.
    python3 - "$corpus_dir" <<'PY'
import json, glob, os, sys, collections
d = sys.argv[1]
expected = {"chat":115,"session":63,"terminal":19,"changeset":16,
            "annotations":10,"root":7,"resourceWatch":2}
got = collections.Counter()
bad = []
for f in sorted(glob.glob(os.path.join(d, "*.json"))):
    try:
        got[json.load(open(f)).get("reducer")] += 1
    except Exception as e:
        bad.append(f"{os.path.basename(f)}: {e}")
if bad:
    print("  [FAIL]: unparseable fixture(s):"); [print("    ", b) for b in bad]; sys.exit(1)
if dict(got) != expected:
    print(f"  [FAIL]: composition {dict(sorted(got.items()))} != expected {expected}")
    sys.exit(1)
print(f"  ✔ composition matches: {dict(sorted(got.items()))} = {sum(got.values())}")
PY
    [ $? -eq 0 ] || bad "corpus composition"
  else
    skip "corpus composition — python3 not on PATH"
  fi
fi

step "2/6 JavaScript — extracted reducers vs corpus"
if ! have node; then
  skip "javascript — node not on PATH"
elif node "$here/js/test/smoke.cjs" > /tmp/ahp-js-corpus.txt 2>&1; then
  grep -q "corpus OK — ${EXPECTED_TOTAL}/${EXPECTED_TOTAL}" /tmp/ahp-js-corpus.txt \
    && ok "js corpus ${EXPECTED_TOTAL}/${EXPECTED_TOTAL}" \
    || { cat /tmp/ahp-js-corpus.txt; bad "js corpus total"; }
else
  cat /tmp/ahp-js-corpus.txt; bad "js smoke test"
fi

step "3/6 Python — extracted reducers vs corpus"
if ! have python3; then
  skip "python — python3 not on PATH"
else
  # check_corpus() pins every per-channel denominator and raises on any shortfall.
  if PYTHONPATH="$here/py/src" python3 -c "
import agent_host_protocol as a
r = a.check_corpus()
print('corpus OK —', sum(t for _, t in r.values()), '/', a.CORPUS_TOTAL)
" > /tmp/ahp-py-corpus.txt 2>&1; then
    grep -q "corpus OK — ${EXPECTED_TOTAL} / ${EXPECTED_TOTAL}" /tmp/ahp-py-corpus.txt \
      && ok "python corpus ${EXPECTED_TOTAL}/${EXPECTED_TOTAL}" \
      || { cat /tmp/ahp-py-corpus.txt; bad "python corpus total"; }
  else
    cat /tmp/ahp-py-corpus.txt; bad "python corpus"
  fi
fi

step "4/6 Go — extracted reducers vs corpus"
if ! have go; then
  skip "go — go not on PATH"
elif (cd "$here/go" && go run -gcflags=all=-l ./cmd/ahp) > /tmp/ahp-go-corpus.txt 2>&1; then
  grep -q "TOTAL: ${EXPECTED_TOTAL}/${EXPECTED_TOTAL} corpus fixtures green" /tmp/ahp-go-corpus.txt \
    && ok "go corpus ${EXPECTED_TOTAL}/${EXPECTED_TOTAL}" \
    || { cat /tmp/ahp-go-corpus.txt; bad "go corpus total"; }
else
  cat /tmp/ahp-go-corpus.txt; bad "go corpus"
fi

step "5/6 C# — extracted reducers vs corpus"
if ! have dotnet; then
  skip "csharp — dotnet not on PATH"
elif (cd "$here" && dotnet run --project cs/test/CorpusRunner -c Release) > /tmp/ahp-cs-corpus.txt 2>&1; then
  grep -q "TOTAL: ${EXPECTED_TOTAL}/${EXPECTED_TOTAL} corpus fixtures green" /tmp/ahp-cs-corpus.txt \
    && ok "csharp corpus ${EXPECTED_TOTAL}/${EXPECTED_TOTAL}" \
    || { cat /tmp/ahp-cs-corpus.txt; bad "csharp corpus total"; }
else
  cat /tmp/ahp-cs-corpus.txt; bad "csharp corpus"
fi

step "6/6 cross-language agreement"
# The point of extracting one proven core to N languages is that the N results are
# not independent re-implementations. If they disagree on the same corpus, that
# property has failed regardless of what the proofs say.
if [ "$fail" = 0 ] && [ "$skipped" = 0 ]; then
  ok "all four extractions agree: ${EXPECTED_TOTAL}/${EXPECTED_TOTAL} each"
else
  skip "cross-language agreement — needs all four legs to have run"
fi

echo
if [ "$fail" -ne 0 ]; then
  echo "SOME CHECKS FAILED (see [FAIL] above)."
elif [ "$skipped" -ne 0 ]; then
  echo "NO FAILURES, but $skipped check(s) were SKIPPED — this is NOT a full run."
else
  echo "ALL GREEN — the extracted code in this repository agrees with the corpus in"
  echo "this repository. This is NOT a re-run of the proofs; see REPRODUCIBILITY.md."
fi
exit $fail
