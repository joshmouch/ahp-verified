#!/usr/bin/env bash
# Phase 4 trust-audit gate (ahp-dafny).
#
# `dafny audit` reports every trusted assumption (extern / assume / axiom /
# unproven ensures) reachable from a program. The hand-authored reducer client
# (spec/client_main.dfy — all 7 channels + OA extension) MUST audit to ZERO
# findings: every reducer, decoder, bv32 bitset, and lemma is proven, with no
# trusted surface. This gate fails if that ever regresses (a new extern/assume/
# axiom sneaks into the reducer client), enforcing the trust-taxonomy contract
# (docs/trust-taxonomy.md): the ONLY trusted seams live in the codec/replay lane
# (Std.JSON + Std.FileIO byte I/O), never in the reducers.
set -euo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
profile="$here/.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime.doo"
[ -f "$profile" ] || { echo "missing Conflux verification package: $profile" >&2; exit 1; }

set +e
out="$(dafny audit --library "$profile" "$here/spec/client_main.dfy" 2>&1)"
audit_code=$?
set -e
printf '%s\n' "$out" | grep -E 'Dafny auditor completed' || true

# `dafny audit --library` reports assumptions in the complete imported .doo,
# even when the audited reducer does not call them. Keep the AHP-owned ratchet
# at zero while classifying only the exact canonical Conflux process-probe,
# atomic-file/path, UTF-8, and command-identity declarations by
# resolved library origin and declaration identity. Any additional package
# finding, path drift, or AHP-owned finding remains a hard failure.
finding_lines="$(printf '%s\n' "$out" | grep -E 'Warning: .*Possible mitigation' || true)"
probe_re='(^|/)\.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime[.]doo\([0-9]+,[0-9]+\): Warning: (RunProcessBounded|CurrentProcessRssMb|MonotonicTimeMs): Declaration has explicit `\{:axiom\}` attribute[.] Possible mitigation: Provide a proof or test[.]$'
identity_re='(^|/)\.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime[.]doo\([0-9]+,[0-9]+\): Warning: (CanonicalJsonBytes|DecodeCanonicalJson|Sha256Digest): Function with `\{:extern\}` attribute[.] Possible mitigation: Turn into a `method` with `modifies \{\}` and test thoroughly for lack of side effects[.]$'
# conflux-extern command-identity seam FACTS (command-identity.dfy: ConfluxCommandIdentity's two
# `lemma {:axiom}` declarations). The three {:extern} identity functions above are bodyless AND
# ensures-less, so no consumer could reason about them at all; these two lemmas state the seam's
# guarantees explicitly, which is what makes the trust surface auditable instead of hidden.
#   CanonicalJsonRoundTrips — DecodeCanonicalJson(CanonicalJsonBytes(v)) == Some(v). The defining
#     property of a canonical codec, already VERIFIED for the pure layer by
#     ConfluxJsonText.ParseStringifyRoundTrip; the adapters call ConfluxJsonText.Stringify /
#     ParseJsonChecked, so the only untrusted-by-Dafny step is the strict-UTF-8 byte boundary.
#   Sha256DigestIsValid — ValidSha256Digest(Sha256Digest(bytes)). A pure FORMAT claim ("sha256:" +
#     64 lower-hex); a SHA-256 content address satisfies it by construction. No collision-resistance
#     property is asserted.
# Both are exercised at runtime on C# and JS by conflux-extern/tests/command_identity_adapter_test.dfy
# (conflux-extern/gen/test-command-identity-adapters.sh).
identity_facts_re='(^|/)\.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime[.]doo\([0-9]+,[0-9]+\): Warning: (CanonicalJsonRoundTrips|Sha256DigestIsValid): Declaration has explicit `\{:axiom\}` attribute[.] Possible mitigation: Provide a proof or test[.]$'
storage_re='(^|/)\.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime[.]doo\([0-9]+,[0-9]+\): Warning: (AtomicWriteFile|AtomicWriteBytes|InspectPath): Declaration has explicit `\{:axiom\}` attribute[.] Possible mitigation: Provide a proof or test[.]$'
utf8_re='(^|/)\.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime[.]doo\([0-9]+,[0-9]+\): Warning: (EncodeUtf8|DecodeUtf8): Function with `\{:extern\}` attribute[.] Possible mitigation: Turn into a `method` with `modifies \{\}` and test thoroughly for lack of side effects[.]$'
# conflux-extern I/O + process-lifecycle seam (io.dfy: `method {:extern} {:axiom}`) — read-stdin,
# child exit status, terminate process tree. OS-level operations Dafny cannot implement; trusted by origin.
ioproc_re='(^|/)\.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime[.]doo\([0-9]+,[0-9]+\): Warning: (ReadStdinLineWithin|ProcessExitStatus|TerminateProcessTreeWithin): Declaration has explicit `\{:axiom\}` attribute[.] Possible mitigation: Provide a proof or test[.]$'
allowed_re="(${probe_re})|(${identity_re})|(${identity_facts_re})|(${storage_re})|(${utf8_re})|(${ioproc_re})"
unexpected="$(printf '%s\n' "$finding_lines" | grep -Ev "$allowed_re" || true)"
allowed_count="$(printf '%s\n' "$finding_lines" | grep -Ec "$allowed_re" || true)"
reported_count="$(printf '%s\n' "$out" | sed -nE 's/.*Dafny auditor completed with ([0-9]+) findings.*/\1/p' | tail -1)"

identities_ok=1
for symbol in RunProcessBounded CurrentProcessRssMb MonotonicTimeMs \
              CanonicalJsonBytes DecodeCanonicalJson Sha256Digest \
              CanonicalJsonRoundTrips Sha256DigestIsValid \
              AtomicWriteFile AtomicWriteBytes InspectPath EncodeUtf8 DecodeUtf8 \
              ReadStdinLineWithin ProcessExitStatus TerminateProcessTreeWithin; do
  count="$(printf '%s\n' "$finding_lines" | grep -Ec "Warning: ${symbol}:" || true)"
  [ "$count" -eq 1 ] || identities_ok=0
done

if [ "$audit_code" -eq 0 ] && [ "$reported_count" = 16 ] &&
   [ "$allowed_count" -eq 16 ] && [ "$identities_ok" -eq 1 ] &&
   [ -z "$unexpected" ]; then
  echo "TRUST-AUDIT OK — AHP reducer has zero owned findings; exactly 16 inherited Conflux process/identity/identity-fact/storage/UTF-8/io declarations classified by resolved origin"
  exit 0
fi

echo "TRUST-AUDIT FAIL — AHP-owned or unclassified trusted assumptions were reported:"
printf '%s\n' "$out" | grep -iE 'finding|extern|assume|axiom' | head -30
[ -z "$unexpected" ] || { echo "unclassified findings:"; printf '%s\n' "$unexpected"; }
exit 1
