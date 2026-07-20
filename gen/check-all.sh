#!/usr/bin/env bash
# Master verification gate (ahp-dafny) — proves the ENTIRE client in one command.
#
# Composes every gate built for this client (per gate-as-feature): proofs,
# both corpora (hand-transcribed + real-fixture reducer-replay), the Std.JSON
# codec, the trust audit, proof robustness, the generator drift gate, the
# extraction facade+mapper, and cross-lineage parity with the TS client.
# CI-ready: exits non-zero on the first failure. Run: bash gen/check-all.sh
set -uo pipefail

# ── net10 SDK + apphost-flake hardening (shared-tree, multi-agent robustness) ──
# In this shared worktree `dafny run --target cs` builds into obj/ trees that
# sibling agent-users also touch. Two hazards, both fixed here:
#   1. A stale homebrew dotnet@8 on PATH targets net8.0 and its CreateAppHost
#      dies chmod-ing the apphost. Force the clean net10 SDK ahead of it so every
#      `dafny run --target cs` rolls forward to net10 consistently.
#   2. CreateAppHost writes the launcher exe mode 755 (no group-write) and can
#      only chmod a file it OWNS; when a sibling agent-user owns a stale apphost
#      in the shared obj/, the build dies "Could not set file permission 755 for
#      .../apphost". `dafny run` never needs the native apphost, so disable it.
# umask 002 keeps any NEW build artifacts group-writable so concurrent agents can
# overwrite/remove them next run. Together these make the cs runs pass repeatably
# regardless of which agent-user last populated obj/.
[ -d /usr/local/share/dotnet ] && PATH="/usr/local/share/dotnet:$PATH"
export PATH
export UseAppHost=false DOTNET_NOLOGO=1 DOTNET_CLI_TELEMETRY_OPTOUT=1
umask 002

here="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# conflux enforcement floor — source the shared shim so `dafny verify` loads the --plugin (detectors a/b/c/d).
# The plugin path, manifest, on-demand build and injection all live in floor/floor-verify.sh (one shared copy).
CONFLUX_SEAM="AhpSocket"   # module(s) allowed to declare {:extern}; drives floor detector (b)
__cf_floor="$here"; while [ "$__cf_floor" != / ] && [ ! -f "$__cf_floor/conflux-workspace/floor/floor-verify.sh" ]; do __cf_floor="$(dirname "$__cf_floor")"; done
if [ -f "$__cf_floor/conflux-workspace/floor/floor-verify.sh" ]; then . "$__cf_floor/conflux-workspace/floor/floor-verify.sh"; fi
# AHP_REPO is OPTIONAL and only two of the eight steps need it: the generator drift
# gate (7/8) and cross-lineage TypeScript parity (8/8). Both compare against LIVE
# upstream, so neither can be pinned in-tree. Every other step runs against the corpus
# vendored at corpus/reducers/ (see corpus/PROVENANCE.md).
#
# When AHP_REPO is unset those two steps report SKIPPED — loudly, and they are counted
# in the final summary so a skipped run can never be mistaken for a fully green one.
FORK="${AHP_REPO:-}"
if [ -n "$FORK" ] && [ -d "$FORK/types" ] && [ -d "$FORK/types/test-cases/reducers" ]; then
  echo "[upstream] live checkout: $FORK — drift gate and TS parity WILL run"
else
  if [ -n "$FORK" ]; then
    echo "[upstream] AHP_REPO=$FORK is not an agent-host-protocol checkout (needs types/ and types/test-cases/reducers/)." >&2
  fi
  FORK=''
  echo "[upstream] no checkout supplied — steps 7/8 (generator drift) and the TS-parity"
  echo "[upstream] half of step 8/8 will be SKIPPED, not passed. To run them:"
  echo "[upstream]   export AHP_REPO=/path/to/agent-host-protocol"
fi
fail=0
skipped=0
step() { printf "\n══ %s ══\n" "$1"; }
ok() { echo "  ✔ $1"; }
bad() { echo "  [FAIL]: $1"; fail=1; }
# A skipped check is NEVER a pass. It is counted and reported in the final summary so
# that "no failures" and "everything ran" stay distinguishable.
skip() { echo "  [SKIP]: $1"; skipped=$((skipped + 1)); }
# ── the private dependency, stated plainly ──────────────────────────────────────
# Re-running the PROOFS requires the Conflux runtime package (a preverified Dafny
# `.doo` plus its translation records). Conflux is the author's own runtime library.
# It is NOT currently published, so a third party CANNOT run this gate today. That
# limitation is real, it is not worked around here, and it is documented in full at
# REPRODUCIBILITY.md ("What you cannot reproduce, and why").
#
# What a reader CAN check without it is listed in REPRODUCIBILITY.md §"What you can
# reproduce today" — the extracted code, the corpus replay through it, the fixture
# corpus itself, and every arithmetic claim about this repository's contents.
CONFLUX_PACKAGE_LINK="$here/.conflux/runtime/dependencies/conflux-runtime/current"
conflux_missing() {
  echo "" >&2
  echo "CANNOT RUN: the Conflux runtime package is not present." >&2
  echo "  looked for: $1" >&2
  echo "" >&2
  echo "  Conflux is the author's own Dafny runtime library and is not published yet." >&2
  echo "  Re-running the proofs therefore requires access that a third party does not" >&2
  echo "  have today. This is a known, disclosed limitation — see REPRODUCIBILITY.md." >&2
  echo "" >&2
  echo "  Checks that DO run without it: bash gen/check-extracted-corpus.sh" >&2
  echo "" >&2
  exit 1
}
if ! CONFLUX_PACKAGE_TARGET="$(readlink "$CONFLUX_PACKAGE_LINK")"; then
  conflux_missing "$CONFLUX_PACKAGE_LINK"
fi
case "$CONFLUX_PACKAGE_TARGET" in
  /*) CONFLUX_PACKAGE_DIR="$(cd "$CONFLUX_PACKAGE_TARGET" && pwd -P)" ;;
  *) CONFLUX_PACKAGE_DIR="$(cd "$(dirname "$CONFLUX_PACKAGE_LINK")/$CONFLUX_PACKAGE_TARGET" && pwd -P)" ;;
esac || { echo "unresolvable runtime package target $CONFLUX_PACKAGE_TARGET" >&2; exit 1; }
CONFLUX_PACKAGE_MANIFEST="$CONFLUX_PACKAGE_DIR/manifest.json"
CONFLUX_DOO="$CONFLUX_PACKAGE_DIR/conflux-runtime.doo"
CONFLUX_CS_DTR="$CONFLUX_PACKAGE_DIR/conflux-runtime-cs.dtr"
TARGET_RUNNER="$here/../scripts/run-conflux-target.sh"
verify_conflux() { dafny verify --library "$CONFLUX_DOO" "$@"; }
[ -f "$CONFLUX_DOO" ] || conflux_missing "$CONFLUX_DOO"

sha256_file() { shasum -a 256 "$1" | cut -d ' ' -f 1; }
owner_verification_identity() {
  local source_path="$1"
  shift
  local current_target source_sha manifest_sha doo_sha floor_dll_sha floor_manifest_sha option options
  current_target="$(readlink "$CONFLUX_PACKAGE_LINK")" || return 1
  [ "$current_target" = "$CONFLUX_PACKAGE_TARGET" ] || return 1
  [ "${CONFLUX_FLOOR_DISABLE:-0}" != 1 ] || return 1
  [ -f "$source_path" ] && [ -f "$CONFLUX_PACKAGE_MANIFEST" ] && [ -f "$CONFLUX_DOO" ] \
    && [ -n "${CONFLUX_FLOOR_PLUGIN:-}" ] && [ -f "$CONFLUX_FLOOR_PLUGIN" ] \
    && [ -n "${CONFLUX_FLOOR_MANIFEST:-}" ] && [ -f "$CONFLUX_FLOOR_MANIFEST" ] || return 1
  source_sha="$(sha256_file "$source_path")" || return 1
  manifest_sha="$(sha256_file "$CONFLUX_PACKAGE_MANIFEST")" || return 1
  doo_sha="$(sha256_file "$CONFLUX_DOO")" || return 1
  floor_dll_sha="$(sha256_file "$CONFLUX_FLOOR_PLUGIN")" || return 1
  floor_manifest_sha="$(sha256_file "$CONFLUX_FLOOR_MANIFEST")" || return 1
  options='<none>'
  for option in "$@"; do options="$options|$option"; done
  printf '%s\n' \
    "package_target=$current_target" \
    "package_manifest_sha256=$manifest_sha" \
    "runtime_doo_sha256=$doo_sha" \
    "floor_dll_sha256=$floor_dll_sha" \
    "floor_manifest_sha256=$floor_manifest_sha" \
    "floor_seam=$CONFLUX_SEAM" \
    "verify_time_limit=${CONFLUX_VERIFY_TIME_LIMIT:-<default>}" \
    "source_path=$source_path" \
    "source_sha256=$source_sha" \
    "verify_options=$options"
}

step "1/8 proofs — dafny verify"
if verify_conflux "$here/spec/client_main.dfy"; then ok "client_main + all channels verify (0 errors)"; else bad "dafny verify"; fi
if verify_conflux "$here/spec/version.dfy"; then ok "version registry strict-total-order proof verifies (0 errors)"; else bad "dafny verify version.dfy"; fi
if verify_conflux "$here/spec/fidelity.dfy"; then ok "Json⇄RootAction codec + round-trip proof verifies (0 errors)"; else bad "dafny verify fidelity.dfy"; fi
for ch in session chat terminal changeset annotations resource_watch canvas; do
  if [ "$ch" = chat ] || [ "$ch" = session ]; then
    verify_conflux --verification-time-limit 120 "$here/spec/codec/${ch}_codec.dfy"
  else
    verify_conflux "$here/spec/codec/${ch}_codec.dfy"
  fi && ok "${ch} Json⇄Action codec + round-trip proof verifies (0 errors)" \
     || bad "dafny verify ${ch}_codec.dfy"
done
if verify_conflux "$here/spec/ahp.dfy"; then ok "unified multi-channel Ahp (applyAhp dispatch + routing codec + action round-trip) verifies (0 errors)"; else bad "dafny verify ahp.dfy"; fi
if verify_conflux "$here/spec/command.dfy"; then ok "AHP command specialization + wire round trips verify (0 errors)"; else bad "dafny verify command.dfy"; fi
if dafny verify "$here/spec/turn_window_proposal.dfy"; then ok "typed turn-window/cursor proposal witness verifies (0 errors)"; else bad "dafny verify turn_window_proposal.dfy"; fi

step "1b/8 channel owners delegate folding to the conflux reducer kernel"
# The former *-apply-is-reducer correspondence files were deletion scaffolding. Their owning consumers now
# call ConfluxContract.Fold directly, so verify those owners rather than retaining or hard-coding bridges.
# Each owner resolves Conflux only through the preverified runtime profile .doo.
verified_owner_sources=()
verified_owner_identities=()
is_reusable_owner_source() {
  case "$1" in session|changeset|annotations) return 0 ;; *) return 1 ;; esac
}
owner_was_verified() {
  local source="$1"
  local current_identity i
  current_identity="$(owner_verification_identity "$here/spec/${source}.dfy")" || return 2
  i=0
  while [ "$i" -lt "${#verified_owner_sources[@]}" ]; do
    if [ "${verified_owner_sources[$i]}" = "$source" ]; then
      [ "${verified_owner_identities[$i]}" = "$current_identity" ] && return 0
      return 2
    fi
    i=$((i + 1))
  done
  return 1
}
for owner in root:ahp_skeleton session:session chat:chat terminal:terminal changeset:changeset annotations:annotations resourcewatch:resource_watch canvas:canvas ahp:ahp; do
  name="${owner%%:*}"
  source="${owner##*:}"
  owner_identity_before=''
  if is_reusable_owner_source "$source"; then
    if ! owner_identity_before="$(owner_verification_identity "$here/spec/${source}.dfy")"; then
      bad "verification identity unavailable before step 1b for ${source}.dfy"
      continue
    fi
  fi
  if verify_conflux "$here/spec/${source}.dfy"; then
    if is_reusable_owner_source "$source"; then
      if ! owner_identity_after="$(owner_verification_identity "$here/spec/${source}.dfy")"; then
        bad "verification identity changed during step 1b for ${source}.dfy"
        continue
      fi
      if [ "$owner_identity_before" != "$owner_identity_after" ]; then
        bad "verification identity changed during step 1b for ${source}.dfy"
        continue
      fi
      verified_owner_sources[${#verified_owner_sources[@]}]="$source"
      verified_owner_identities[${#verified_owner_identities[@]}]="$owner_identity_after"
    fi
    ok "${name} owner delegates to conflux reducer kernel (0 errors)"
  else
    bad "dafny verify ${source}.dfy"
  fi
done

step "1c/8 channel seq helpers delegate to the ONE ConfluxSeqRoute keyed-seq primitive"
# The former *-seqroute-bridge / *-sequpdate-bridge / *-sequpdatewhere-bridge correspondence files were
# deletion scaffolding: they proved each channel's hand-rolled upsert/remove/update-by-id helper was
# byte-for-byte the canonical primitive. Their owning consumers now CALL the packaged ConfluxSeqRoute primitive
# directly through the Conflux runtime `.doo` — upsertQ/upsertFile/upsertChat/upsertAnn/… delegate to SeqUpsertById, remove*
# to SeqRemoveById, changeAnswer/updateOp/updateAnn/upsertCust/updChat to SeqUpdateById, setMcp to
# SeqUpdateByIdWhere, and chat's deltaMarkdown/deltaReasoning/mapTC to SeqUpdateAllWhere. The delegation lives in
# each channel's helper def, so the generic ToMap∘helper == PrunableRoute-Install/Remove/guarded-update tie
# (proven once in the separately verified ConfluxSeqRoute package module) is
# inherited for free — no per-helper bridge retained. Re-verify the owning channels rather than hard-coding
# bridges (chat is also covered in step 1b; the light channels exercise SeqUpsertById/SeqRemoveById/
# SeqUpdateById/SeqUpdateByIdWhere delegation directly).
for src in session changeset annotations; do
  if owner_was_verified "$src"; then
    ok "${src} seq helpers delegate to ConfluxSeqRoute (0 errors)"
  else
    owner_reuse_status=$?
    if [ "$owner_reuse_status" -eq 2 ]; then
      bad "${src} seq-helper verification identity changed since step 1b"
    else
      bad "${src} seq-helper owner verification failed in step 1b"
    fi
  fi
done

step "1d/8 channel host-authority law — each channel's host canonical IS ConfluxContract.Fold (Campaign D)"
# spec/channel_laws.dfy INSTANTIATES ConfluxChannel.HostCanonicalIsFold + HostSeqAdvances at each channel's own
# apply1, proving host.canonical == CC.Fold(apply1, ·, ·) and nextSeq == |actions| — the host-authority half of
# "everything is a Channel". VERIFY-ONLY: it consumes the core surface + the streaming envelope through the SAME
# --library seam convergence.dfy uses (core .doo + conflux-runtime.doo). The core .doo is rebuilt first so the law
# verifies against the CURRENT channel reducers. ConfluxChannel is resolved only through the packaged `.doo`,
# and --library emits no target code, so channel_laws is never
# translated — it is pure proof. PeersConverge (host==client) stays in convergence-dafny (core owns no client seq).
if bash "$here/gen/build-doo.sh" >/dev/null 2>&1 \
   && dafny verify --library "$here/agent-host-protocol-core.doo" --library "$CONFLUX_DOO" "$here/spec/channel_laws.dfy"; then
  ok "all 8 channels instantiate HostCanonicalIsFold + HostSeqAdvances (0 errors)"
else bad "dafny verify channel_laws.dfy (channel host-authority law)"; fi

step "1e/8 firewall BITES — every tightened channel/codec module's internals unreachable through core.doo"
# Each ghost const names a REAL internal HIDDEN by its module's `export` clause (setBit, decodeCust, …) —
# a baseline-diff bite: it resolved under the old export-everything default and MUST NOT now. Verified
# against the shipped agent-host-protocol-core.doo; every reference must be an unresolved-member error.
biteprobe="$here/.firewall-bite-probe.dfy"
cat > "$biteprobe" <<'PROBE'
module FirewallBiteProbe {
  import Session import Chat import Terminal import Changeset import Annotations
  import RW = ResourceWatch import F = Fidelity
  import Sc = SessionCodec import Cc = ChatCodec import Tc = TerminalCodec
  import Csc = ChangesetCodec import Anc = AnnotationsCodec import RWc = ResourceWatchCodec
  ghost const b1 := Session.setBit          ghost const b2 := Chat.setBit
  ghost const b3 := Terminal.appendData     ghost const b4 := Changeset.hasOp
  ghost const b5 := Annotations.hasAnn      ghost const b6 := RW.IsUnknownRW
  ghost const b7 := F.encodeConfig          ghost const b8 := Sc.decodeCust
  ghost const b9 := Cc.EncodeToolCall       ghost const b10 := Tc.encodePart
  ghost const b11 := Csc.decodeOp           ghost const b12 := Anc.optStr
  ghost const b13 := RWc.ResourceWatchCodecC
}
PROBE
bo="$(dafny verify --library "$here/agent-host-protocol-core.doo" --library "$CONFLUX_DOO" "$biteprobe" 2>&1 || true)"
bmiss=0
for sym in setBit appendData hasOp hasAnn IsUnknownRW encodeConfig decodeCust EncodeToolCall encodePart decodeOp optStr ResourceWatchCodecC; do
  echo "$bo" | grep -q "$sym" || { echo "  [FAIL] firewall LEAK: $sym reachable through core.doo"; bmiss=1; fail=1; }
done
# Version is a standalone proof (not bundled in core.doo) — bite via source include.
vprobe="$here/.version-bite-probe.dfy"
printf 'include "%s/spec/version.dfy"\nmodule VBite { import Ver = Version\n  lemma b(xs: seq<Ver.SemVer>) { Ver.StrictlyDescendingPairwise(xs, 0, 1); } }\n' "$here" > "$vprobe"
vo="$(dafny verify "$vprobe" 2>&1 || true)"
echo "$vo" | grep -q "StrictlyDescendingPairwise" || { echo "  [FAIL] firewall LEAK: Version.StrictlyDescendingPairwise reachable"; bmiss=1; fail=1; }
rm -f "$biteprobe" "$vprobe"
[ "$bmiss" = 0 ] && ok "all 14 tightened modules' internals are walled off past their export firewall (baseline-diff bite)"

step "2/8 hand-transcribed corpus (cs+js)"
for t in cs js; do
  corpus_output="$("$TARGET_RUNNER" "$t" "$here" "$here/spec/client_main.dfy" "$here/build/client_main_$t" 2>&1)"
  if echo "$corpus_output" | grep -q 'MULTI-CHANNEL CLIENT OK'; then
    ok "corpus green on $t"
  else
    printf '%s\n' "$corpus_output" >&2
    bad "corpus $t"
  fi
done

step "2b/8 unified multi-channel demo — applyAhp fold + codec round-trip (cs+js)"
for t in cs js; do
  if "$TARGET_RUNNER" "$t" "$here" "$here/run/ahp_demo.dfy" "$here/build/ahp_demo_$t" 2>&1 | grep -q 'AHP MULTI-CHANNEL DEMO OK'; then ok "demo green on $t"; else bad "demo $t"; fi
done

step "3/8 reducer-replay — full pinned 237-fixture corpus (232 Microsoft + 5 Canvas)"
rr="$(bash "$here/gen/run-reducer-replay.sh" 2>&1)"
rr_failed=0
for pair in "ROOT:7/7" "RESOURCEWATCH:2/2" "CANVAS:5/5" "CHANGESET:16/16" "ANNOTATIONS:10/10" "TERMINAL:19/19" "SESSION:63/63" "CHAT:115/115"; do
  ch="${pair%%:*}"; n="${pair##*:}"
  if echo "$rr" | grep -q "$ch REDUCER-REPLAY: $n"; then ok "$ch $n"; else bad "$ch replay ($n)"; rr_failed=1; fi
done
[ "$rr_failed" = 0 ] || printf '%s\n' "$rr" >&2

step "4/8 Std.JSON codec round-trip (cs+js)"
if [ "$(bash "$here/gen/run-codec-roundtrip.sh" 2>&1 | grep -c 'CODEC OK')" -ge 2 ]; then ok "232/232 Microsoft corpus round-trip on cs+js"; else bad "codec"; fi

step "5/8 trust audit — zero trusted assumptions"
if bash "$here/gen/check-trust-audit.sh"; then
  ok "reducer client has zero AHP-owned trust findings; exact inherited Conflux probes classified"
else bad "trust audit"; fi

step "6/8 anti-vacuity + proof robustness"
if bash "$here/gen/check-proof-robustness.sh" 3 2>&1 | grep -q 'ANTI-VACUITY OK'; then ok "witnesses + mutation seeds"; else bad "proof robustness"; fi

step "7/8 generator drift gate"
if [ -z "$FORK" ]; then
  skip "generator drift — needs a live upstream checkout (export AHP_REPO=...)"
else
  drift="$(cd "$FORK" && NODE_PATH="$FORK/node_modules" node_modules/.bin/tsx "$here/gen/generate-dafny.ts" --out "$here/gen/mirror.generated.dfy" --check 2>&1)"
  if echo "$drift" | grep -q 'fresh'; then ok "mirror matches types/ (no drift)"; else bad "generator drift"; fi
fi

step "8/8 extraction mapper (C# primary, JS secondary) + cross-lineage parity"
# C# is the PRIMARY run/proof lineage (dafny translate cs output is directly
# consumable — real public classes/namespaces, no facade patch needed). See
# gen/CS-HARNESS-PATTERN.md for the recipe this composes.
mapper_work="$(mktemp -d "${TMPDIR:-/tmp}/ahp-core-mapper.XXXXXX")"
mapper_source="$mapper_work/ahp04_client_cs"
mapper_translate="$(dafny translate cs "$here/spec/client_main.dfy" --no-verify \
  --library "$CONFLUX_DOO" --translation-record "$CONFLUX_CS_DTR" \
  --output "$mapper_source" 2>&1)"
mapper_status=$?
mapper_output=''
if [ "$mapper_status" = 0 ]; then
  mapper_output="$(cd "$here" && dotnet run --project mapper/Mapper.csproj \
    -p:ConsumerSource="$mapper_source.cs" \
    -p:BaseIntermediateOutputPath="$mapper_work/obj/" \
    -p:MSBuildProjectExtensionsPath="$mapper_work/obj/" \
    -p:OutputPath="$mapper_work/bin/" 2>&1)"
  mapper_status=$?
fi
if [ "$mapper_status" = 0 ] && echo "$mapper_output" | grep -q 'MAPPER OK'; then
  ok "C# extraction + separate-process mapper (primary)"
else
  printf '%s\n%s\n' "$mapper_translate" "$mapper_output" >&2
  bad "C# extraction/mapper"
fi
printf '[INFO] retained mapper temp without recursive cleanup: %s\n' "$mapper_work"
# JS translation is still generated + checked (second translation for cross-lineage parity).
if node "$here/gen/build-js-facade.mjs" >/dev/null 2>&1 && node "$here/mapper/ahp-host-mapper.mjs" 2>&1 | grep -q 'MAPPER OK'; then ok "JS facade + separate-process mapper (secondary translation)"; else bad "facade/mapper"; fi
if [ -z "$FORK" ]; then
  skip "cross-lineage TS parity — needs a live upstream checkout (export AHP_REPO=...)"
else
  # 238 = 232 fixture cases + 6 suite-level unit tests (IS_CLIENT_DISPATCHABLE ×2,
  # isClientDispatchable ×2, reducer-immutability ×2). The fixture half is the 232 that
  # the Dafny replay also covers; the extra 6 are TypeScript-only and have no Dafny peer.
  tsparity="$(cd "$FORK" && npx tsx --test types/reducers.test.ts 2>&1)"
  if echo "$tsparity" | grep -q 'fail 0' && echo "$tsparity" | grep -q 'pass 238'; then ok "TS lineage 238/238 = 232 fixtures + 6 unit tests (parity)"; else bad "TS parity"; fi
fi

echo
if [ "$fail" -ne 0 ]; then
  echo "SOME CHECKS FAILED (see [FAIL] above)."
elif [ "$skipped" -ne 0 ]; then
  echo "NO FAILURES, but $skipped check(s) were SKIPPED (see [SKIP] above) — this is NOT a full green run."
  echo "The skipped checks compare against live upstream; supply AHP_REPO to run them."
else
  echo "ALL GREEN — the Dafny AHP core is fully verified end-to-end."
fi
exit $fail
