# AHP Verified

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Dafny 4.11.0](https://img.shields.io/badge/Dafny-4.11.0-6f42c1.svg)](https://dafny.org)
[![Targets: C%23 · JS · Python · Go · JVM](https://img.shields.io/badge/targets-C%23%20%C2%B7%20JS%20%C2%B7%20Python%20%C2%B7%20Go%20%C2%B7%20JVM-informational.svg)](#packages)

**A formally verified implementation of the [Agent Host Protocol](https://github.com/microsoft/agent-host-protocol) — the reducers, codecs and protocol laws are stated as machine-checked theorems in Dafny, then mechanically extracted to C#, JavaScript, Python, Go and the JVM.**

The shipped code in every language is the *same proven artifact*, not N independent hand-written re-interpretations. This is a third-party/community project by Joshua Mouch. It is **not** an official Microsoft product.

---

## Why this exists

Every other AHP client — Microsoft's own Go, Kotlin, Rust, Swift and TypeScript clients, and any hand-written .NET one — is written by hand and checked by tests. Tests *sample* the input space; proofs *cover* it. Where the TypeScript lineage asserts that a status bitfield round-trips by exercising a handful of fixtures, this core proves it for all 2³² values, and proves that an unknown high bit survives a round-trip — a property a closed TypeScript const-enum cannot even represent.

The second consequence is drift. N hand-written clients are N chances to interpret the same wire format differently. Here the reducers are written and proved once, and each language binding is a translation of that one artifact.

The 237-fixture corpus replay is **not** the safety net — the proofs are. The corpus is the independent cross-check on top of them, and it is a far sharper oracle than per-assertion tests: it decodes each fixture into domain values, folds them through the actual proven reducer, and asserts *full structural state equality*. Building this client, that strictness drove out a batch of modeling defects that a curated hand-written corpus had passed clean.

---

## What is proven

Every module verifies with **0 errors and 0 timeouts** under Dafny 4.11.0 — **626 verification units** in total.

| Area | What is proven | Where |
|---|---|---|
| **Host-authority law** | All **8** channels instantiate the kernel lemmas at their own `apply1`: `AcceptFold(r,h,as).canonical == Fold(r,h.canonical,as)` and `nextSeq == h.nextSeq + \|as\|`. The "everything is a channel" law, generic over the reducer. | `spec/channel_laws.dfy` |
| **Codec round-trips** | `decode ∘ encode == id` per channel, unified across the dispatch layer (`AhpActionRoundTrip`, `AhpStateRoundTrip`, `AhpWireCanonicalRoundTrip`). | `spec/codec/*_codec.dfy`, `spec/ahp.dfy` |
| **Status bitfield fidelity** | `Status_RoundTrip` over **all** `bv32` values; `Status_UnknownBitPreserved` proves an unknown high bit (`0x80000000`) survives. | `spec/fidelity.dfy` |
| **Open/closed dichotomy** | `Open_PreservesUnknownKeys` and its dual `Closed_DropsUnknownKeys`. | `spec/fidelity.dfy` |
| **Unknown-variant passthrough** | Unknown variants re-encode verbatim (`RootUnknown_EncodeVerbatim`, `UnknownVariant_Verbatim`). | `spec/fidelity.dfy` |
| **Channel isolation** | A session action provably touches only session state; a root action only root state. | `spec/ahp.dfy` |
| **Version registry** | `compareSemVer` proven a strict total order, then non-emptiness, strict descent and pairwise distinctness of the shipped `SUPPORTED` list discharge as theorems — not runtime tests. | `spec/version.dfy` |

### Zero trusted assumptions in the reducer core

This is the claim to check first, because it is the one a formal-methods skeptic attacks first — *what did you assume away?*

```
$ bash gen/check-trust-audit.sh
Dafny auditor completed with 16 findings
TRUST-AUDIT OK — AHP reducer has zero owned findings; exactly 16 inherited
```

The AHP-owned trusted surface is **zero**. All 16 findings are inherited from the imported runtime library, and each is pinned by exact symbol name *and* resolved origin path — any other symbol, any other origin, or a count ≠ 16 fails the gate. It is an enumerable list, not a blanket waiver. Adding a channel adds no trusted surface.

### Corpus results

| Check | Result | Scope |
|---|---|---|
| Reducer replay | **237/237** | 232 upstream fixtures + 5 canvas, full structural state equality, on C# and JS |
| `Std.JSON` codec round-trip | **232/232** | parse → serialize → re-parse, ASTs identical (unknown-key + key-order preservation) |
| Cross-lineage parity | **238/238** | Microsoft's own TypeScript reducer suite, unmodified, on the same corpus |
| Extracted-code corpus | **148/148** | curated per-action corpus, run against the *extracted* code in each binding |

Replay denominators are hard-coded per channel in the gate, so a fixture that failed to parse would drop the count and fail the build — nothing can be silently skipped.

### Honest limits

Shipped deliberately, because you would find them anyway:

- The `RootActionRoundTrip` / `AhpActionRoundTrip` lemmas carry preconditions excluding unknown variants. That case is **not unproven** — it has its own dual law (unknown leaves re-encode verbatim), and each precondition is guarded by an inhabitation witness so neither lemma is vacuously true.
- **Anti-vacuity covers 7 of 8 channels.** Canvas is not in the gate's loop and has no non-vacuity witness yet.
- **Mutation-seed robustness covers one verification task**, not all 626. Dafny's `--verify-included-files` defaults to false, so the seed-mutation arm resolves to `Main` alone.
- The **148-fixture curated corpus is a modeled subset** for two channels (session 36 of 61, chat 54 of 97). The *replay* corpus covers both in full (session 63/63, chat 115/115).
- Proof binds Dafny-code to Dafny-model; replay binds Dafny-model to corpus. So a future TypeScript-vs-Dafny divergence localizes to the TypeScript side **only within the corpus-confirmed subset**; outside it, divergence is two-sided.
- `dafny verify spec/client_main.dfy` alone reports `1 verified` — only `Main`. The gate verifies each file individually; that is what produces 626. Cite the gate, not the one-liner.

Reproduce everything: `bash gen/check-all.sh`.

---

## Packages

**Nothing is published to a public registry yet.** Every artifact below is built and reproducible from source in this tree; the install commands are what they *will* be. Verified against the registries on 2026-07-20: NuGet, npm and PyPI all return 404 for these names.

| Target | Package | Status |
|---|---|---|
| **.NET** | `Ahp.Core.Verified` 0.1.0 | Artifact built — self-contained `net8.0` assembly, no external dependencies. Released on GitHub; registry publish pending credentials. |
| **JavaScript** | `@open-agency/ahp` 0.1.0 | Artifact built, smoke test **148/148** green. Released on GitHub; registry publish pending credentials. |
| **Python** | `agent-host-protocol` 0.1.0 | Wheel + sdist built, smoke test **148/148** green. Released on GitHub; registry publish pending credentials. |
| **Go** | `github.com/joshmouch/ahp-verified/go` | Builds and runs the corpus **148/148** green. Tagged `go/v0.1.0`. Requires a build flag — see below. |
| **JVM** | `agency.open.ahp:ahp-core` 0.1.0 | Extracted sources + `pom.xml` present; jar not built in this tree. Compilation needs `-Xmx8g` and is slow by design. |
| **C++** | — | **Blocked, irreducibly.** Dafny 4.11.0 cannot translate this core: unbounded `int`/`nat` and higher-order types are unsupported by the C++ backend. Bounding integers to `int64` would break the fold proofs. Evidence in `cpp/evidence/`. |

```bash
dotnet add package Ahp.Core.Verified          # .NET
npm install @open-agency/ahp                  # JavaScript
pip install agent-host-protocol               # Python
go get github.com/joshmouch/ahp-verified/go            # Go
```

```xml
<dependency>
  <groupId>agency.open.ahp</groupId>
  <artifactId>ahp-core</artifactId>
  <version>0.1.0</version>
</dependency>
```

**Go build flag.** The generated `Chat` package defeats the Go compiler's inliner, which balloons past 60 GB resident and is OOM-killed. Build with inlining disabled — it compiles in about three seconds and leaves the rest of the optimizer on:

```bash
go build -gcflags 'all=-l' ./...
```

---

## Example

```js
const ahp = require('@open-agency/ahp');

// The 8 channel reducers plus the shared skeleton:
//   ahp.Chat, ahp.Session, ahp.Terminal, ahp.Canvas,
//   ahp.Changeset, ahp.Annotations, ahp.ResourceWatch, ahp.AhpSkeleton

// Each channel exposes its reducer as ApplyTo<Channel>(state, action, now):
const next = ahp.Chat.__default.ApplyToChat(state, action, now);

const emptyTurns = ahp._dafny.Seq.of();
ahp.Chat.__default.hasTurn(emptyTurns, ahp._dafny.Seq.UnicodeFromString('id')); // false
```

Values crossing the boundary are Dafny representations (`_dafny.Seq` for sequences and strings, `BigNumber` for integers), not plain JS values. That is deliberate: shipping the extracted core verbatim means no hand-written marshalling layer sits between you and the proven code, so there is nothing unverified to disagree with the proofs. An idiomatic wrapper can be layered on top.

---

## How it is built

```
Microsoft AHP types/  ──generator──>  mirror.generated.dfy  (drift-gated)
                                              │
              hand-authored + proven reducers, codecs, laws
                                              │
                                      dafny verify  (626 units, 0 errors)
                                              │
                          ┌───────────────────┼───────────────────┐
                       translate            translate          translate
                          C#                  JS              Python / Go / JVM
                                              │
                    corpus replay + codec round-trip + TS parity
```

1. **Generate the wire surface.** `gen/generate-dafny.ts` — a ts-morph codegen target, sibling of upstream's own `generate-<lang>.ts` convention — reads the real AHP `types/` and emits **295 datatypes + 39 enums + 38 alias-unions + 23 alias-stubs + 2 never-types** (397 types, 726 lines). A `--check` drift gate fails on any divergence from upstream.

   Three details worth naming: alias unions whose members are all declared types become *real* discriminated unions, not opaque blobs. The 23 remaining opaque `X(raw: Json)` stubs are unions mixing primitives, string literals or inline shapes — disclosed, not hidden. And TypeScript `never` maps to an *uninhabited subset type* rather than a stub, because mapping it to a stub silently invented wire values that could be mistaken for real server-only variants.

2. **Prove.** Reducers, codecs and laws are hand-authored and machine-checked. The library exports a `.doo` with an explicit `provides`/`reveals` set; an export-firewall gate asserts that 13 module internals stay unreachable through it.

3. **Extract.** C# is the primary lineage — `dafny translate cs` output is directly consumable. Each target is then driven by a **separate OS process** through the host boundary, so a divergence between lineages surfaces immediately.

4. **Cross-check.** Corpus replay, codec round-trip, and Microsoft's own TypeScript suite, all on the same pinned fixtures.

---

## Relationship to upstream

This project tracks [microsoft/agent-host-protocol](https://github.com/microsoft/agent-host-protocol) and consumes its `types/` and fixture corpus directly — the generator reads upstream types, and the drift gate fails when they change. Upstream currently ships five first-party clients: Go, Kotlin, Rust, Swift and TypeScript.

This is an independent third-party implementation. It is not affiliated with or endorsed by Microsoft.

**Standing offer to the AHP maintainers:** the verified core is available to the project. If there is interest in a proven reducer core, a conformance oracle for the existing clients, or the fixture-replay harness as a cross-client gate, it can be contributed upstream on whatever terms the maintainers prefer. The corpus replay already runs Microsoft's own TypeScript suite unmodified, so it can serve as a differential oracle today.

### License

MIT, dual copyright:

- Copyright © Microsoft Corporation — the protocol and the original client
- Copyright © 2026 Josh Mouch — the Dafny verification and the extraction

---

## Further reading

- [The verification explainer](site/index.html) — how the proofs are structured and what each law buys you, for readers who do not write Dafny.
- Per-language notes: [.NET](cs/README.md) · [JavaScript](js/README.md) · [Python](py/README.md) · [Go](go/README.md) · [JVM](java/README.md) · [C++ blocker analysis](cpp/README.md)

Authored by Joshua Mouch.
