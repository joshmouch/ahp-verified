# C++ — the verified core, reachable from C++

**Status: shipping.** A C++ program can call the proven AHP reducers.

Getting here took a detour worth explaining, because the obvious route is a dead end.

## `dafny translate cpp` does not work, and will not

Dafny's C++ backend is experimental and cannot express this core. Three independent blockers,
each fatal on its own:

| Blocker | Why it is fatal here |
|---|---|
| No unbounded `int`/`nat` | The protocol datatypes carry 60 `: int` and 47 `nat` fields — timestamps, terminal dimensions, exit codes, the fold clock |
| No higher-order functions | The reducer core is built on them: 26 arrow signatures, including functions that return functions |
| No `Std.JSON` | The codec routes through it, and it is itself unbounded-int and higher-order heavy |

Demonstrated on `spec/version.dfy` — the simplest module in the core, which imports nothing, so
no library or entry-point choice can rescue it. Bounding every `int` to `int64` would compile,
but it would delete the unbounded arithmetic the convergence and fold proofs rest on. The binary
would build and would no longer be the verified core.

This was re-checked against the newest Dafny in existence (nightly `4.11.1+368adac`, 2026-07-07),
not merely assumed from the stable release — same three errors, same order. The full analysis,
with raw command output for every claim, is in [DAFNY-CPP-BACKEND.md](DAFNY-CPP-BACKEND.md).

That investigation also produced a fix worth upstreaming independently: Dafny 4.11.0's shipped
`DafnyRuntime.h` does not compile at all (`DafnySet<T>::disjoint` calls a nonexistent `find`).
A patch is in [`DafnyRuntime.h.patch`](vendor/DafnyRuntime.h.patch).

## What ships instead

The core extracts cleanly to several languages that *can* emit a native library with a C ABI.
C++ links that library. **The code deciding protocol behavior is still machine-derived from the
proofs** — no reducer was reimplemented in C++.

Three routes were built. All three work, each producing a native library, a C header, and a
compiled C++ program that calls real reducers.

### Recommended — [`../cpp-via-dafny-next/`](../cpp-via-dafny-next/), via the Rust backend

The strongest route, and a genuine surprise: **`dafny translate rs` already exists in Dafny
4.11.0**, and the Rust backend supports both features that kill the C++ one — unbounded integers
(`num::BigInt`) and function values (`Rc<dyn Fn>`).

- 43,946 lines of Rust emitted from the verified core; all 8 channel reducers present
- `libahp_core.a` (10.8 MB) / `libahp_core.dylib` (1.7 MB), 16 exported `ahp_*` C symbols
- **A C++ binary runs the core's own conformance corpus: 148/148 green, all 8 channels**

It is a direct Dafny backend, so the chain is one compiler shorter than the alternatives.

One source change was required, and it is *re-proven* rather than asserted: the Rust backend
mandates `--enforce-determinism`, which bans the assign-such-that statement. The single compiled
occurrence in the runtime — a JSON key-sort helper; no reducer was touched — was lifted to the
let-such-that expression form carrying an identical predicate, with uniqueness discharged by the
runtime's own pre-existing lemma. `regenerate.sh` re-verifies the entire patched runtime
(**1623 verified, 0 errors**) before it will translate, so the artifact cannot be rebuilt without
re-proving the rewrite.

### Alternate — [`../cpp-via-go/`](../cpp-via-go/), via Go `c-archive`

Unedited `dafny translate go` output plus a ~200-line cgo bridge, built with
`-buildmode=c-archive`. Ships a static archive and a shared library, with the hand-written
const-correct header gated against cgo ABI drift on every build.

### Alternate — [`../cpp-via-dotnet/`](../cpp-via-dotnet/), via .NET NativeAOT

The extracted C# published as a native shared library with `[UnmanagedCallersOnly]` entry points.
Two notable results: the Dafny C# backend emits **fully AOT-compatible code** (zero trim or
reflection warnings under `IlcTreatWarningsAsErrors`), and the resulting `.dylib` has **zero .NET
runtime dependencies**. It doubles as a cross-host check — CLR/JIT and NativeAOT produce
byte-identical reducer results.

## The trust chain, stated plainly

**Machine-derived from the proofs:** every reducer, the kernel fold, the codecs, and the JSON
parser/writer. Nothing was hand-ported.

**Hand-written and therefore not verified:** the FFI binding layer (marshalling, handle table,
error mapping), the C header, and the example programs. These contain no protocol logic — each
binding function converts in, calls a verified function, and converts out — but a marshalling or
lifetime bug there is excluded by no proof.

**Outside the proofs entirely,** as in every language binding: the Dafny backend itself, the host
toolchain, and the `{:extern}` host-capability boundary (filesystem, process, clock, sockets).
The reducers are pure and never call those, which is why the corpus runs green regardless.

In one sentence: *the reducers are proven; the binding around them is not.*

## Coverage and limits

The core models the root reducer and five channels completely; session (36 of 61 actions) and
chat (54 of 97) are modeled in part. Unmodeled actions are **not** rejected — they decode to a
defined `RootUnknown` catch-all, and the binding exposes this so a caller can detect them rather
than silently treating them as no-ops.

Only `osx-arm64` was built and tested. The C ABI and the C++ sources are portable and the
underlying toolchains support Linux and Windows, but those are *expected to work*, not
demonstrated.
