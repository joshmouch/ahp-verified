# Verification

**AHP Verified** — a formally verified implementation of the Microsoft Agent Host Protocol, written in Dafny and machine-extracted to multiple languages.

This is a third-party community project by Joshua Mouch. It is not an official Microsoft product and carries no Microsoft endorsement. MIT licensed, dual copyright: Microsoft (the protocol and the original client) and Joshua Mouch (the Dafny verification and extraction).

This document explains what "verified" means here and what it does not mean. It is deliberately specific about the limits. If you are evaluating whether to trust this code, the limits section is the part that matters.

> **On checking these claims yourself.** Some you can, some you currently cannot, and the difference is not cosmetic. The fixture corpus and the extracted code are checkable from a clean clone — `bash gen/check-extracted-corpus.sh`. **The proofs are not**: re-running them requires the Conflux runtime library, which is the author's own work and is not published. The proof sources are in `spec/` and can be read; they cannot presently be re-verified by a third party. [REPRODUCIBILITY.md](REPRODUCIBILITY.md) sets out claim by claim which is which, and what that means for the "zero owned trusted assumptions" result below.

Toolchain: Dafny 4.11.0. Every command in this document was run against the state of the repository it ships with.

---

## 1. The thesis

Every other AHP client is hand-written and tested. Upstream ships five first-party clients — Go, Kotlin, Rust, Swift, TypeScript — each an independent re-reading of the protocol by a different author in a different language. Tests sample the input space. Proofs cover it.

Verification-first means something narrower and more mechanical than "we were careful":

1. **The protocol's laws are stated as theorems in Dafny.** Not as test assertions over chosen inputs, but as universally quantified statements over all inputs, discharged by an SMT solver. "The session-status bitset round-trips" is not 40 example values; it is a statement about all 2^32 of them.

2. **The implementation is extracted from the proven spec, not written against it.** `dafny translate cs` and `dafny translate js` emit the shipped code from the same source the proofs are about. There is no second artifact to keep in sync, no "the spec says X but the implementation does Y" drift class, because the implementation *is* a mechanical translation of the proven source.

That second point is the whole argument. A hand-written client verified against a formal spec has two artifacts and a gap between them. An extracted client has one artifact and no gap. The cost is that the extracted code is not idiomatic, which Section 4 addresses honestly.

The corpus of real fixtures is not the safety net here. It is the cross-check *on top of* the proofs — the thing that catches a proof that is correct about a model that does not match reality. Both signals are load-bearing and they fail in different directions.

---

## 2. What is proven, and what is only tested

This is the section a skeptical reader should read first, because it is where a verification claim usually turns out to be smaller than advertised.

### 2.1 Every module verifies

`dafny verify` on each module, with the Conflux runtime library on the `--library` path:

| Module | File | Units | Result |
|---|---|---:|---|
| root | `spec/ahp_skeleton.dfy` | 18 | 0 errors |
| session | `spec/session.dfy` | 19 | 0 errors |
| chat | `spec/chat.dfy` | 113 | 0 errors |
| terminal | `spec/terminal.dfy` | 4 | 0 errors |
| changeset | `spec/changeset.dfy` | 13 | 0 errors |
| annotations | `spec/annotations.dfy` | 12 | 0 errors |
| resourceWatch | `spec/resource_watch.dfy` | 4 | 0 errors |
| canvas | `spec/canvas.dfy` | 14 | 0 errors |
| unified dispatch | `spec/ahp.dfy` | 76 | 0 errors |
| commands | `spec/command.dfy` | 18 | 0 errors |
| fidelity | `spec/fidelity.dfy` | 34 | 0 errors |
| version registry | `spec/version.dfy` | 14 | 0 errors |
| codecs (7 files) | `spec/codec/*_codec.dfy` | 269 | 0 errors |
| channel laws | `spec/channel_laws.dfy` | 11 | 0 errors |
| turn-window proposal | `spec/turn_window_proposal.dfy` | 6 | 0 errors |
| entry point | `spec/client_main.dfy` | 1 | 0 errors |

**626 verification units, 0 errors, 0 timeouts.**

> **One caveat, stated plainly because it would otherwise look like an overclaim.** Dafny's `--verify-included-files` defaults to false. So `dafny verify spec/client_main.dfy` on its own reports `1 verified` — it verifies `Main` and nothing else. Older text in this repository's README implies that single command verifies everything; it does not. The 626 figure comes from verifying each file individually, which is what `gen/check-all.sh` actually does (steps 1 and 1b). Cite the gate, not the one-liner.

### 2.2 The laws, by name

**Host authority — "everything is a channel."** `spec/channel_laws.dfy` instantiates two kernel lemmas at each of the 8 channels' own per-action step `apply1`:

```
AcceptFold(r, h, actions).canonical == Fold(r, h.canonical, actions)
AcceptFold(r, h, actions).nextSeq   == h.nextSeq + |actions|
```

That is: folding an action stream through the streaming envelope reaches exactly the same state as folding it through the reducer kernel, and the assigned sequence number advances by exactly one per action. Per-channel instantiations: `SessionIsChannel`, `ChatIsChannel`, `TerminalIsChannel`, `ChangesetIsChannel`, `AnnotationsIsChannel`, `ResourceWatchIsChannel`, `CanvasIsChannel`, `RootIsChannel`. A non-vacuity witness (`ChannelWitness`) folds a concrete 2-action root stream and asserts `nextSeq == 2` — falsifiable by changing the stream length.

Note what is *not* here: peer convergence (host state equals client mirror) is not instantiated in this repository. This core owns no live client sequence; that seam lives in a separate convergence module. What is proven here is the host-authority half — the part a consumer with no live client can carry.

**Fidelity — three families,** in `spec/fidelity.dfy`:

- *bv32 status algebra.* `Status_RoundTrip` (`:38`) — the encode/decode round-trip over **all** `bv32` values. `Status_WidthUniformity` (`:39`). `Status_FlagCombo72` (`:41`) pins fixture 004's flag decode. `Status_UnknownBitPreserved` (`:45`) proves an unknown high bit (`0x80000000`) survives encoding — explicitly the property a closed TypeScript const-enum cannot represent.
- *Open/closed dichotomy.* `Open_PreservesUnknownKeys` (`:58`) proves an open struct re-encodes with unknown keys intact; its dual `Closed_DropsUnknownKeys` (`:68`) proves a closed struct drops them. Having both directions matters — one alone is satisfiable by a degenerate codec.
- *Unknown-variant passthrough.* `RootUnknown_EncodeVerbatim` (`:183`) and `UnknownVariant_Verbatim` (`:185`).

**Codec round-trips.** `decode ∘ encode == id`, per channel in `spec/codec/*_codec.dfy`, unified in `spec/ahp.dfy` as `AhpActionRoundTrip` (`:467`), with the wire-fixpoint corollary `AhpWireCanonicalRoundTrip` (`:462`).

**Channel isolation.** `ApplyAhp_SessionTouchesOnlySession` (`spec/ahp.dfy:203`) and `ApplyAhp_RootTouchesOnlyRoot` (`:213`) prove dispatch updates only the addressed component and leaves the other seven byte-identical. This rules out a stub dispatch that ignores its action or clobbers unrelated state.

**Version registry as a theorem.** `spec/version.dfy` proves `compareSemVer` is a strict total order — `CompareReflexiveZero` (`:63`), `CompareZeroIffEqual` (`:69`), `CompareAntisymmetric` (`:74`), `CompareTransitive` (`:81`), `CompareTotal` (`:91`) — then discharges `StrictlyDescendingImpliesDistinct` (`:134`) and `UpstreamInvariants` (`:153`) for the shipped `SUPPORTED := [0.5.2, 0.5.1]`: non-empty, `SUPPORTED[0] == PROTOCOL_VERSION`, strictly descending, pairwise distinct. Upstream asserts these with runtime tests; here "strictly descending implies no duplicates" is proven once and holds for any future registry that satisfies the ordering.

#### The round-trip preconditions, and why they are not a hole

The first thing a careful reader should attack is `requires`. `RootActionRoundTrip` (`spec/fidelity.dfy:161`) carries `requires !a.RootUnknown?`. `AhpActionRoundTrip` carries `requires typedAhp(a) && wfAhp(a)`. A round-trip theorem with the hard case excluded is worth very little.

The excluded case is not unproven — it has its own law. The `Unknown` leaf re-encodes verbatim (`:183`, `:185`), which is the correct behavior for a variant whose payload may itself be typed-shaped and must not be re-classified. The union of the two lemmas covers the action space.

Each precondition is also guarded by an inhabitation witness — `RootActionRoundTrip_NonVacuous` (`:174`), `AhpActionRoundTrip_NonVacuous` (`spec/ahp.dfy:487`) — proving `exists a :: Precondition(a) && ...`. Without those, a precondition that no value satisfies would make the theorem vacuously true and the proof worthless.

### 2.3 What the proofs do not cover

Proofs bind Dafny code to a Dafny model. They say nothing about the four seams below, which are trusted:

| Seam | What is assumed | What guards it |
|---|---|---|
| `Std.FileIO` raw byte read | reads a file's bytes faithfully | any corruption fails `reduced == expected` on real fixtures |
| `Std.JSON` byte↔AST extern boundary | the extern byte/AST edge is faithful (the AST transform itself is proven Dafny) | `spec/codec/json_roundtrip.dfy` on the full corpus, both targets |
| `bv8 → uint8` bridge | a raw byte and a JSON byte denote the same 0–255 value | pure total Dafny; range-checked by the verifier |
| target runtime | the C#/JS runtime executes the translated code faithfully | the corpus runs on **both** extracted targets; a divergence surfaces immediately |

Beyond those, the imported Conflux runtime library contributes inherited assumptions. `dafny audit` reports **16 findings, all inherited, zero owned by AHP**:

```
$ bash gen/check-trust-audit.sh
Dafny auditor completed with 16 findings
TRUST-AUDIT OK — AHP reducer has zero owned findings; exactly 16 inherited
Conflux process/identity/identity-fact/storage/UTF-8/io declarations classified by resolved origin
```

This is not a blanket waiver. The gate pins each of the 16 by exact symbol name *and* resolved library origin: `RunProcessBounded`, `CurrentProcessRssMb`, `MonotonicTimeMs`, `CanonicalJsonBytes`, `DecodeCanonicalJson`, `Sha256Digest`, `CanonicalJsonRoundTrips`, `Sha256DigestIsValid`, `AtomicWriteFile`, `AtomicWriteBytes`, `InspectPath`, `EncodeUtf8`, `DecodeUtf8`, `ReadStdinLineWithin`, `ProcessExitStatus`, `TerminateProcessTreeWithin`. Any other symbol, a different origin path, or a count other than 16 fails the gate.

The practical consequence: **adding a channel adds zero trusted surface.** A new reducer and decoder are proven Dafny. The trusted set is fixed by the codec and I/O choices, not by protocol coverage.

**What "zero owned" does and does not mean.** All sixteen resolve to `conflux-runtime.doo`. So the trusted surface was not eliminated — it was relocated into Conflux, which is the author's own runtime library and is **not published**. A third party can read the pinned list above (`gen/check-trust-audit.sh` is in this repository) but cannot audit the declarations themselves, nor re-run the gate that produced this output. "Zero owned, sixteen inherited" is exactly true and is narrower than it first reads. See [REPRODUCIBILITY.md](REPRODUCIBILITY.md).

**Conflux is also redistributed inside the published packages** — 49 of 61 modules in the Python wheel, and proportionate amounts elsewhere. Dafny extraction inlines the library, so there is no separate dependency to install. Disclosed per package in [NOTICE](NOTICE).

### 2.4 The gates that bound the proofs — including where they fall short

Proofs can be vacuous. Exported internals can leak. Three gates address this, and two of them have limits worth stating.

**Anti-vacuity.** `gen/check-proof-robustness.sh` asserts every channel carries a `*_NonVacuous` witness and runs mutation seeds.

```
$ bash gen/check-proof-robustness.sh 3
ANTI-VACUITY OK
```

Two honest limits:

- **Canvas is not covered.** The gate's loop checks 7 channels (`ahp_skeleton`, `resource_watch`, `changeset`, `annotations`, `terminal`, `session`, `chat`). `canvas` is absent from that list, and `spec/canvas.dfy` contains zero `NonVacuous` witnesses. The gate's success message says "every channel"; it means the 7 it checks. Canvas's 14 units verify, but no witness rules out a vacuous lemma among them.
- **The mutation arm is narrow.** It runs `dafny measure-complexity --mutations N` on `client_main.dfy`, which — per the `--verify-included-files` default above — resolves to exactly one verification task. Do not read this as "all 626 proofs are robust across random SMT seeds." One is. The other 625 are not seed-mutated.

**Export firewall.** `gen/check-all.sh` (step 1e) writes a probe module referencing 13 module internals hidden behind `export` clauses — `Session.setBit`, `Chat.setBit`, `Terminal.appendData`, `Changeset.hasOp`, `Annotations.hasAnn`, `ResourceWatch.IsUnknownRW`, `Fidelity.encodeConfig`, `SessionCodec.decodeCust`, `ChatCodec.EncodeToolCall`, `TerminalCodec.encodePart`, `ChangesetCodec.decodeOp`, `AnnotationsCodec.optStr`, `ResourceWatchCodec.ResourceWatchCodecC` — plus `Version.StrictlyDescendingPairwise`, and **fails if any of them resolves** through the shipped library artifact. Every one resolved under the old export-everything default, so this is a real baseline-diff bite, not a tautology.

---

## 3. What the proofs actually buy

Four arguments. Each is falsifiable.

### 3.1 One proven core extracted to N languages eliminates a drift class

The failure mode in a multi-client protocol is not that any one client is wrong. It is that each is independently *plausible*. Six hand-written clients are six readings of the same prose. Where the prose is ambiguous — does `configChanged` merge or replace? does a tool call in `pending-confirmation` raise the input-needed bit? — each author resolves the ambiguity locally and reasonably, and the clients silently disagree. Nothing fails until two of them meet in production.

Extraction removes the re-interpretation step. The C# and the JS are translations of the same proven source, so a semantic difference between them would have to be a bug in the Dafny backend rather than a difference of opinion about the protocol. The two lineages here are themselves a cross-check on that: both are generated, both run the full corpus, and a backend-level divergence would show up as one target passing and the other failing.

The honest scope: this eliminates drift *among clients extracted from this core*. It does nothing about drift between this core and a hand-written client, which is what Section 3.4 is for.

### 3.2 Proofs cover what tests sample

`Status_RoundTrip` holds over all 2^32 `bv32` values. A test suite covers the values someone thought of. That difference is not rhetorical: `Status_UnknownBitPreserved` proves an unknown high bit survives, which is a property TypeScript's closed const-enum *structurally cannot express* — there is no place in the type to put a bit the enum does not name. You cannot write that test in that lineage; you can only write the version that passes.

Same shape for the version registry. Upstream asserts "SUPPORTED is strictly descending" and "SUPPORTED has no duplicates" as two runtime tests over one array. Here the second is a *consequence* of the first, proven once, for any array.

**On the "35 bugs" claim, which you may have seen elsewhere and should discount.** Earlier writeups of this project say strict replay drove out roughly 35 reducer bugs the curated corpus missed. Two corrections. First, characterization: those were defects in *the Dafny modeling of the protocol*, not defects in Microsoft's TypeScript client. Every one was fixed to match the fixtures — that is, to match the TypeScript lineage's behavior. Any reading of that number as "found 35 bugs in the upstream client" is wrong. Second, provenance: the figure originates in development history that predates this repository's extraction (102 commits, 11 of them fixes), so it is not mechanically checkable from what ships here. Treat it as the author's account, not as a verified number.

What *is* checkable, and is the real point: strict full-state replay is a sharper oracle than per-assertion testing. The harnesses compare whole decoded state, not sampled fields.

### 3.3 Protocol changes re-verify mechanically

The type surface is generated from upstream's own `types/` by `gen/generate-dafny.ts`, a ts-morph sibling of upstream's existing per-language codegen convention:

```
$ tsx gen/generate-dafny.ts --out gen/mirror.generated.dfy --check
fresh: gen/mirror.generated.dfy matches types/
(295 datatypes + 39 enums + 38 alias-unions + 23 alias-stubs + 2 never-types)
```

397 emitted types, 726 lines, drift gate passing. When upstream's types change, the gate fails, the mirror regenerates, and everything downstream re-verifies — the proofs either still discharge or they point at the exact obligation the change broke. That is a materially different maintenance posture from re-reading a changelog and hand-patching six clients.

Three generator details a reviewer should check, because each is a judgment call:

- **38 alias-unions are real discriminated unions.** A TypeScript alias whose members are all declared types becomes a genuine Dafny union, not an opaque blob.
- **23 remain opaque `X(raw: Json)` stubs.** Unions mixing primitives, string literals, or inline shapes. This is disclosed, not hidden — they are the weakest part of the type surface.
- **2 `never` types map to uninhabited subset types**, not to stubs. Mapping `never` to an opaque stub silently invented wire values (notably `ServerCanvasAction`) that could be mistaken for real server-only variants. The uninhabited encoding makes the impossibility explicit.

### 3.4 It is a reducer oracle any implementation can be checked against

Because the reducers are both proven against a model and replayed against the real corpus, they can adjudicate. Point any implementation at the same fixtures; where it disagrees, one side is wrong and the disagreement is localized.

The localization has an honest bound, and stating it is more useful than hiding it: proof binds Dafny-code to Dafny-model; replay binds Dafny-model to corpus. So within the corpus-confirmed subset, where the model is validated by fixtures, a divergence localizes to the other implementation. **Outside that subset, a divergence is two-sided** — the model has no fixture backing there, so either side could be at fault. The oracle is strong on 237 fixtures and silent beyond them.

---

## 4. What this costs

Verification-first is not free, and the costs are structural rather than incidental.

**Dafny expertise is required to change the core.** Adding a channel means writing a reducer, a codec, round-trip lemmas, and a non-vacuity witness, then getting them past the solver. This is a narrower talent pool than "can write TypeScript." For a protocol client that a broad contributor base is expected to maintain, that is a real argument against this approach.

**Extracted code is not idiomatic.** The C# is 14,464 lines of machine translation: Dafny's naming, Dafny's sequence and map types, `BigInteger` where a domain type would use `int`. Nobody would enjoy reading it. This is why the shipped surface uses a facade — a hand-written idiomatic layer over the extracted core, so consumers get normal types and the proven code stays untouched underneath. The facade is itself hand-written and therefore not proven; it is a thin mapping layer, but it is trusted, and it is fair to count it as trusted surface for anyone consuming through it.

**Only C# and JS are proven-extracted.** Both run end-to-end here. Go, Java, and Python are mature non-experimental Dafny backends with no known structural blocker for this core, but the extraction has not been run or validated for any of them. Classify those as engineering effort, not proof risk — and do not classify them as done.

**C++ cannot express this core, and that is unlikely to change soon.** The experimental `cpp` backend lacks BigInteger (the protocol datatypes embed `int`/`nat` directly — 66 and 47 occurrences respectively, including terminal dimensions, timestamps, exit codes, and the fold clock), most higher-order functions (37 arrow signatures in the reducer merge/upsert core, including functions returning functions), and `Std.JSON`. Bounding every integer to `int64` to satisfy the backend would break the fold proofs that unbounded arithmetic underpins — which forfeits the entire point. This is a capability gap, not a packaging gap.

**Verification time is real.** The chat codec alone needs an extended time limit (`--verification-time-limit 120`) to discharge its 140 units. A full `gen/check-all.sh` compiles and runs 8 channels across 2 targets on top of the proof work. This is a minutes-scale gate, not a seconds-scale one, and it shapes the edit-test loop accordingly.

---

## 5. Check it yourself

Everything below is re-runnable from a clean checkout. `AHP_REPO` must point at an `agent-host-protocol` checkout containing `types/` and `types/test-cases/reducers/`.

**Everything at once:**

```bash
export AHP_REPO=/path/to/agent-host-protocol
bash gen/check-all.sh
```

Eight steps: proofs, both corpora, reducer-replay, Std.JSON codec, trust audit, anti-vacuity, generator drift, extraction mapper plus cross-lineage parity. Exits non-zero on the first failure.

**Proofs, per module** — note the per-file invocation; the single-entry-point form verifies only `Main`:

```bash
DOO=.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime.doo
dafny verify --library "$DOO" spec/chat.dfy          # 113 verified, 0 errors
dafny verify --library "$DOO" spec/ahp.dfy           #  76 verified, 0 errors
dafny verify --library "$DOO" spec/fidelity.dfy      #  34 verified, 0 errors
dafny verify --library "$DOO" spec/version.dfy       #  14 verified, 0 errors
dafny verify --library "$DOO" --verification-time-limit 120 spec/codec/chat_codec.dfy   # 140

# channel laws need the core library too:
dafny verify --library agent-host-protocol-core.doo --library "$DOO" spec/channel_laws.dfy   # 11
```

**Reducer replay — 237 fixtures, both targets:**

```bash
bash gen/run-reducer-replay.sh
```

Expected, and hard-coded as denominators in `gen/check-all.sh:252`:

```
ROOT 7/7 · RESOURCEWATCH 2/2 · CANVAS 5/5 · CHANGESET 16/16
ANNOTATIONS 10/10 · TERMINAL 19/19 · SESSION 63/63 · CHAT 115/115
```

232 Microsoft fixtures plus 5 local canvas fixtures. The Microsoft denominators reconcile exactly against the corpus counted by each fixture's own `reducer` field:

```bash
cd "$AHP_REPO/types/test-cases/reducers"
for f in *.json; do python3 -c "import json;print(json.load(open('$f'))['reducer'])"; done | sort | uniq -c
#  115 chat / 63 session / 19 terminal / 16 changeset / 10 annotations / 7 root / 2 resourceWatch  = 232
```

*The obvious attack, and the answer.* The runner globs by filename, which over-selects: `*terminal*` catches 20 files but `071-root-terminalschanged.json` is a root fixture; `*changeset*` catches 18 but `145-` and `146-session-changesetschanged-*.json` are session fixtures. The chat, session, terminal, changeset, and canvas harnesses filter by the fixture's own `reducer` field and drop non-matching files from the denominator — so in principle an unparseable fixture could vanish rather than fail. **The denominators are hard-coded in the gate.** One broken chat fixture makes the harness print 114 and `gen/check-all.sh` fails. Separately, the root, resourceWatch, and annotations harnesses do *not* filter at all — they count every argument, so a parse failure there is a hard failure rather than a skip.

Each replay reads the real fixture bytes, parses with `Std.JSON`, decodes `initial`/`actions`/`expected` into domain values, folds through the actual proven reducer, and asserts full structural state equality `reduced == expected`. Not field sampling. The only normalization anywhere is chat's `modifiedAt := "N"` applied to both sides, and it is printed in the result line.

**Codec round-trip — 232 fixtures, both targets:**

```bash
bash gen/run-codec-roundtrip.sh
```

Per fixture: parse, serialize, re-parse, assert the two ASTs are identical — which is what establishes unknown-key and key-order preservation on real data. `Main` counts every argument with no filter and `expect pass == total` aborts hard.

**Trust audit and anti-vacuity:**

```bash
bash gen/check-trust-audit.sh        # 16 inherited, 0 AHP-owned
bash gen/check-proof-robustness.sh 3 # witnesses (7 of 8 channels) + mutation seeds (1 task)
```

**Generator drift:**

```bash
cd "$AHP_REPO"
AHP_REPO=$PWD NODE_PATH=$PWD/node_modules node_modules/.bin/tsx \
  <core>/gen/generate-dafny.ts --out <core>/gen/mirror.generated.dfy --check
```

**Extraction:**

```bash
PKG=.conflux/runtime/dependencies/conflux-runtime/current
dafny translate cs spec/client_main.dfy --no-verify \
  --library "$PKG/conflux-runtime.doo" \
  --translation-record "$PKG/conflux-runtime-cs.dtr" --output /tmp/out_cs   # 14,464 lines
dafny translate js spec/client_main.dfy --no-verify \
  --library "$PKG/conflux-runtime.doo" \
  --translation-record "$PKG/conflux-runtime-js.dtr" --output /tmp/out_js   #  8,137 lines
```

Both are then driven by a separate OS process across the host boundary — `mapper/Mapper.csproj` for C#, `mapper/ahp-host-mapper.mjs` for JS — each asserting `MAPPER OK`. C# is the primary lineage: `dafny translate cs` output is directly consumable, with real public classes and namespaces and no facade patch. JS is the secondary translation, consumed through a generated CommonJS facade.

**Cross-lineage parity** — Microsoft's own TypeScript reducer suite, unmodified, on the same fixtures:

```bash
cd "$AHP_REPO" && npx tsx --test types/reducers.test.ts
# tests 238   pass 238   fail 0
```

---

## 6. Where to attack this

If you are reviewing seriously, these are the places I would push, listed because you will find them anyway and it is better that they come from me.

1. **Canvas has no non-vacuity witness** and is not in the anti-vacuity gate's loop. Its 14 units verify, but nothing rules out a vacuous lemma among them. This is a genuine gap, not a documentation artifact.
2. **Mutation testing covers one verification task,** not 626. The gate's phrasing invites over-reading.
3. **23 opaque type stubs** remain in the generated surface. Unions mixing primitives and inline shapes are carried as raw JSON, so the type-level guarantees are weaker there than the datatype count suggests.
4. **The idiomatic facade is hand-written** and therefore unproven. If you consume through it rather than through the extracted core directly, it belongs in your trusted set.
5. **The oracle is bounded by the corpus.** Outside 237 fixtures, a divergence between this core and another implementation does not localize.
6. **This repository's own README is stale** on several figures — it says 7 channels, 211 replay fixtures, 217 TypeScript tests, 285 datatypes. Current and verified: 8, 237, 238, 295. `docs/parity-report.md`, `docs/trust-taxonomy.md`, and `docs/DISTRIBUTION-MATRIX.md` carry the same stale numbers. `gen/check-all.sh` is the executable source of truth; where prose and gate disagree, trust the gate. Documents describing OpenAgency extension channels refer to a module that no longer lives in this repository.
7. **`docs/DISTRIBUTION-MATRIX.md` is a research draft.** Its registry-availability findings were current as of 2026-07-18 and have not been rechecked. Do not treat any package-name availability claim there as live.

The five claims I would stand behind without qualification: zero AHP-owned trusted assumptions with all 16 inherited ones enumerable by name; 626 verification units at 0 errors; 232 upstream fixtures replaying to full structural state equality on two independently extracted targets with gate-pinned denominators; properties proven here that the TypeScript lineage structurally cannot express; and both lineages green on the same canonical corpus.

---

*Joshua Mouch. MIT licensed. Not affiliated with or endorsed by Microsoft.*
