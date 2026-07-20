# Reproducibility

What a third party can check in this repository, what they cannot, and exactly why.

This document exists because the rest of the repository makes strong claims — 626
verification units, zero owned trusted assumptions, 237/237 corpus replay — and a
verification claim that cannot be checked is an assertion. The honest position is
that **some of these claims are checkable from a clean clone today and some are
not**, and a reader is entitled to know which is which before deciding how much
weight to give any of them.

Short version:

| Claim | Checkable from a clean clone? |
|---|---|
| The extracted code runs the fixture corpus green, in 4 languages | **Yes** |
| The fixture corpus is upstream's, unmodified | **Yes** |
| The Dafny sources say what the documentation says they say | **Yes** |
| Cross-language agreement (all extractions produce identical results) | **Yes** |
| The C++ / C ABI route builds and runs | **Yes** |
| **The proofs verify** (626 units, 0 errors) | **No — needs a dependency that is not published** |
| **Zero owned trusted assumptions** (the `dafny audit` result) | **No — same dependency** |
| Cross-lineage parity with upstream's TypeScript suite | Only with an upstream checkout |
| Generator drift against upstream `types/` | Only with an upstream checkout |

The rest of this document is the detail behind that table.

---

## What you can reproduce today

One command, no private dependency, no upstream checkout:

```bash
bash gen/check-extracted-corpus.sh
```

It checks two things and is explicit that it checks only those two:

1. **The pinned fixture corpus is intact** — 232 files, the per-channel composition
   the replay denominators are pinned to, every file parseable. Integrity is
   independently checkable with `cd corpus/reducers && shasum -a 256 -c ../SHA256SUMS`.
2. **The extracted code shipped in this repository runs that corpus green** — in
   JavaScript, Python, Go and C#, 148/148 each, with every per-channel denominator
   asserted rather than merely summed.

A leg whose toolchain is missing is reported `[SKIP]` and counted; the summary
distinguishes "no failures" from "everything ran". A skipped check is never a pass.

**What this proves:** the extracted reducers in this repository agree with the
fixture corpus in this repository, and the four language extractions agree with each
other.

**What this does not prove:** that the extracted code was produced from proofs that
verify. That is the claim below, and no amount of running this gate substitutes for
it. It is worth being precise about this, because the corpus result is the most
visible green number in the project and it is *not* the verification claim.

You can also read the proofs. `spec/` holds the Dafny sources — the reducers, the
codecs, the fidelity laws, the version registry, the channel laws, and the replay
harnesses. Reading them does not require running them, and the named theorems the
documentation cites (`Status_RoundTrip`, `Status_UnknownBitPreserved`,
`Open_PreservesUnknownKeys`, `Closed_DropsUnknownKeys`, `UpstreamInvariants`,
`AhpActionRoundTrip`, the per-channel `*IsChannel` instantiations) are all present
and greppable. `gen/` holds every gate script, so what is checked is inspectable
even where it is not runnable.

### The 16 trusted assumptions, without running anything

`gen/check-trust-audit.sh` pins each inherited finding by exact symbol name *and*
resolved origin path. Since the gate is in the repository, the list is readable
even though the gate is not runnable — so "16 inherited" is an enumerable claim, not
a number you have to take on trust:

| Seam | Symbols |
|---|---|
| process probes | `RunProcessBounded`, `CurrentProcessRssMb`, `MonotonicTimeMs` |
| command identity | `CanonicalJsonBytes`, `DecodeCanonicalJson`, `Sha256Digest` |
| identity facts | `CanonicalJsonRoundTrips`, `Sha256DigestIsValid` |
| atomic storage | `AtomicWriteFile`, `AtomicWriteBytes`, `InspectPath` |
| UTF-8 boundary | `EncodeUtf8`, `DecodeUtf8` |
| I/O + process lifecycle | `ReadStdinLineWithin`, `ProcessExitStatus`, `TerminateProcessTreeWithin` |

Sixteen symbols, and **all sixteen resolve to `conflux-runtime.doo`** — which is the
concrete form of the point made above: the trusted surface is not zero, it is
entirely inside the library you cannot currently obtain. The gate fails on any other
symbol, any other origin, or a count that is not exactly 16.

---

## What you cannot reproduce, and why

**Re-running the proofs requires the Conflux runtime package, which is not
published.** `bash gen/check-all.sh` will refuse to start without it, and says so
rather than skipping quietly.

### What Conflux is

Conflux is a Dafny runtime library — a verified reducer kernel (`Fold`), a JSON
codec, sequence-routing primitives, content addressing, and a set of capability
seams. It is **the author's own work**, not a third-party dependency. It lives in a
separate private repository and is consumed here as a preverified `.doo` package plus
translation records.

This project's reducers are written against Conflux's kernel: `ConfluxContract.Fold`
is the fold every channel delegates to, `ConfluxCodec` is the wire representation,
and `ConfluxSeqRoute` supplies the keyed-sequence primitives the channels' helpers
call. It is load-bearing, not incidental.

### Why this matters more than a normal missing dependency

The headline "zero owned trusted assumptions" is precise but easy to misread. The
trust audit reports 16 findings, all *inherited* from Conflux, and zero *owned* by
the AHP reducers. That sentence is true. What it means is that **the trusted surface
was not eliminated — it was moved into Conflux.** Since Conflux is not published, a
reader cannot audit the 16 assumptions the whole trust argument now rests on.

So the trust chain is presented as bottoming out in machine-checked Dafny. Today, for
a third party, it bottoms out in an unpublished library. That is a real limitation of
this artifact as shipped, and it is not resolved by anything in this repository.

### What is vendored into the published packages

**The published packages contain Conflux code.** This was previously undisclosed and
is the single most important thing on this page for anyone installing a package.

Dafny extraction inlines the library into the emitted source; there is no separate
Conflux package to depend on. So every artifact here ships it:

| Package | Conflux content |
|---|---|
| Python `agent-host-protocol` | 49 of 61 modules under `_core/` are `Conflux*` |
| Go `github.com/joshmouch/ahp-verified/go` | 52 of 66 tracked `.go` files |
| JavaScript `@open-agency/ahp` | 48 distinct `Conflux*` modules inside `dist/ahp.cjs` |
| .NET `Ahp.Core.Verified` | 49 of 72 namespaces in `Generated/AhpCore.g.cs`, plus 5 hand-written `Runtime/Conflux*Adapter.cs` |
| JVM `agency.open.ahp:ahp-core` | 25 files across 6 `Conflux*` packages |

The .NET package is described elsewhere as a "self-contained `net8.0` assembly, no
external dependencies". That is literally true — the `.csproj` has zero
`PackageReference` — but the reason it is true is that the runtime is *inlined*
rather than referenced. Read as "this package has few moving parts", it is
misleading. It has the same moving parts; they are compiled in.

### Licence status of the vendored code

The vendored Conflux code is the author's own, and this repository is MIT with dual
copyright (Microsoft for the protocol and original client; Josh Mouch for the
verification and extraction). The author is therefore in a position to license it —
but at the time of writing **the Conflux source repository carries no `LICENSE`
file of its own.** Anyone redistributing these packages should treat the Conflux
portion as covered by this repository's MIT grant via the copyright line for Josh
Mouch, and this is stated explicitly in `NOTICE`.

Resolving this properly — publishing Conflux, or vendoring it with its own explicit
grant, or cutting the dependency — is an open decision and is not made here.

---

## The two checks that need an upstream checkout

These compare against live upstream, so they cannot be pinned in-tree:

```bash
export AHP_REPO=/path/to/agent-host-protocol
bash gen/check-all.sh     # steps 7/8 and the parity half of 8/8 now run
```

- **Generator drift** (`gen/generate-dafny.ts --check`) — fails if upstream's
  `types/` no longer matches the generated mirror.
- **Cross-lineage parity** — runs upstream's own `types/reducers.test.ts` unmodified.
  It reports **238**, which is *not* the 232-fixture corpus: it is 232 fixture cases
  plus 6 suite-level unit tests (`IS_CLIENT_DISPATCHABLE` ×2, `isClientDispatchable`
  ×2, reducer-immutability ×2). The fixture half is the same 232 the Dafny replay
  covers; the extra 6 are TypeScript-only and have no Dafny counterpart.

Without `AHP_REPO`, both report `[SKIP]` and the run is reported as incomplete rather
than green.

---

## Known gaps in the checking apparatus itself

Stated here rather than left to be found:

- **The fixture corpus has no automatic drift gate.** The generated *types* are
  drift-gated; the *fixtures* are not. `corpus/` is a manual snapshot and upstream
  has continued to add fixtures since it was taken. See `corpus/PROVENANCE.md`.
- **CI does not run the proofs.** It cannot — same dependency. CI builds the
  committed extracted product and runs the corpus against it in five languages plus
  the C ABI route. "CI green" is therefore evidence about the *extraction*, and
  carries no information about the proofs. Do not read the badge as more than that.
- **Anti-vacuity covers 7 of 8 channels.** Canvas is verified but is not in the
  witness gate's loop.
- **Mutation-seed robustness covers one verification task, not 626.** Dafny's
  `--verify-included-files` defaults to false, so the seed-mutation arm resolves to
  `Main` alone.
- **The C ABI bindings are hand-written and unverified**, and have known
  memory-safety defects under concurrent use. The proofs cover the reducers, not the
  FFI layer wrapped around them. See `docs/ffi-safety.md`.

- **Two printed denominators in the extracted binaries are stale.**
  `spec/client_main.dfy:35,37` print "of 61 total" for session and "of 97 total" for
  chat. Those were correct against an earlier upstream snapshot; against the corpus
  pinned in `corpus/` they should read **63** and **115**. The modeled counts
  themselves (36 and 54) are correct — only the totals they are quoted against are
  out of date, so the real modeled coverage is 36-of-63 and 54-of-115, i.e. slightly
  *lower* than the printed text implies.

  This is deliberately **not** patched. `spec/` is a verbatim copy of the sources the
  shipped extraction was produced from, and that correspondence is worth more than a
  cosmetic fix — particularly since the proofs cannot currently be re-run after an
  edit (see above), so patching would mean shipping Dafny that has not been
  re-verified. It will be corrected at the next regeneration.

---

## If you are evaluating whether to trust this

The most useful thing you can do without the private dependency is read `spec/` and
run `gen/check-extracted-corpus.sh`. That tells you the model is real, the corpus is
upstream's, and the shipped code agrees with it in four languages.

It does not tell you the proofs verify. Until Conflux is resolvable, that claim rests
on the author's word plus the inspectable proof sources — which is a materially
weaker position than the rest of this repository's tone implies, and the reason this
document exists.
