#!/usr/bin/env bash
# single-corruption.sh -- the smallest possible falsifiability demo.
#
# Take one real corpus fixture, corrupt exactly ONE modeled value in the state a
# client would claim, and show the oracle catch it with a diff that points at the
# exact field. This is the "corrupt one expected value" demo in miniature; the
# systematic version is falsify.py.
#
# Usage:  ./single-corruption.sh [path-to-ahp-oracle]
set -euo pipefail

HERE="$(cd "$(dirname "$0")" && pwd)"
ORACLE="${1:-$HERE/../dist/osx-arm64/ahp-oracle}"
CORPUS="$HERE/../corpus/reducers"
FIXTURE="$CORPUS/147-session-ready-preserves-inprogress-status.json"

WORK="$HERE/single"
mkdir -p "$WORK"

# The state an honest client would claim: the fixture's own expected state,
# verbatim. The oracle must AGREE with this.
python3 -c "import json,sys; json.dump(json.load(open('$FIXTURE'))['expected'], open('$WORK/claimed.good.json','w'), indent=2)"

# The state a buggy client claims: identical, but with one modeled field wrong.
# status 8 (InProgress in the verified enum) mis-set to 1 (a different lifecycle
# code). Exactly one leaf changed; everything else is byte-identical.
python3 -c "
import json
s = json.load(open('$WORK/claimed.good.json'))
s['status'] = 1
json.dump(s, open('$WORK/claimed.buggy.json','w'), indent=2)
"

echo '########################################################################'
echo '# 1. An honest client claims the fixture'"'"'s own expected state.'
echo '#    The oracle folds the actions through the proven reducers and AGREES.'
echo '########################################################################'
set +e
"$ORACLE" check --file "$FIXTURE" --expected "$WORK/claimed.good.json"
echo "exit=$?"
echo
echo '########################################################################'
echo '# 2. A buggy client claims status=1 where the proven fold says 8.'
echo '#    One leaf wrong -> the oracle DIVERGES and names the field.'
echo '########################################################################'
"$ORACLE" check --file "$FIXTURE" --expected "$WORK/claimed.buggy.json"
echo "exit=$?"
set -e
