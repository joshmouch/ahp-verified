# AHP verified core — C++ distribution (NativeAOT route)

A native shared library with a C ABI, and a C++ program that links it and calls
the **formally verified** Agent Host Protocol reducers.

The reducers that execute are machine-derived from a Dafny specification. No
protocol logic in this directory was written by hand.

```
$ ./build/ahp_demo
AHP verified core — C++ consumer demo
=====================================
library: ahp-core-verified 0.1.0 (nativeaot/osx-arm64)

[1] ahp_apply_json — decode + applyAhp
  PASS  initial state encodes                          (len=603)
  PASS  apply returns a state                          (ok)
  PASS  reduced state carries cols 120                 (found "cols":120)
  PASS  reduced state carries rows 40                  (found "rows":40)
...
DEMO OK — all assertions passed.
```

Full output: [`evidence/cpp-demo-output.txt`](evidence/cpp-demo-output.txt).

---

## Why this route exists

The obvious path — `dafny translate cpp` — does not work for this core, and
that is not a matter of effort. Dafny's experimental C++ backend lacks
unbounded `int`/`nat`, most higher-order functions, and `Std.JSON`. The AHP core
uses all three heavily: 60 `int` and 47 `nat` occurrences in the protocol
datatypes, and 26 arrow-typed signatures in the reducer core. The backend fails
even on `spec/version.dfy`, the simplest dependency-free module in the tree.
Reproductions are in [`../cpp/evidence/`](../cpp/evidence/).

So this distribution takes a different route to the same destination. The core
already extracts cleanly to C#. .NET 8+ can compile C# **ahead of time** into a
native shared library that exports C-callable entry points — no CLR, no JIT, no
runtime to install. C++ links that library like any other `.dylib`.

The result is a real C++ distribution: the machine code that runs the reducers
was generated from the verified specification by a chain of compilers, not
transcribed by a person.

```
    Dafny specification  (proofs: convergence, fold consistency, codec identity)
            │
            │  dafny translate cs          ← verified core, machine-extracted
            ▼
    Ahp.Core.Verified  (C#)
            │
            │  + AhpNative.cs              ← hand-written C ABI shell (NOT verified)
            │
            │  dotnet publish -p:PublishAot=true -p:NativeLib=Shared
            ▼
    libahp_core.dylib  (native arm64, C ABI, no .NET runtime)
            │
            │  clang++ … -lahp_core
            ▼
    ahp_demo  (C++)
```

---

## Trust chain — read this before relying on it

Being precise about what is and is not proven matters more than the headline.

**Machine-derived from the proofs.** Everything that decides protocol
behaviour:

- the reducers `applyAhp`, `foldAhp`, and the seven per-channel reducers;
- the codecs `decodeAhpAction`, `encodeAhpAction`, `decodeAhpState`,
  `encodeAhpState`;
- the JSON parser and writer (`ConfluxJsonText.ParseJsonChecked`, `Stringify`),
  which are themselves specified and proven in Dafny rather than delegated to a
  third-party JSON library.

These reach the `.dylib` through two compilers — Dafny's C# backend and .NET's
ILCompiler. Neither was hand-edited.

**Hand-written, and therefore NOT verified.** One file:
[`src/AhpNative.cs`](src/AhpNative.cs). It converts UTF-8 `char*` to Dafny
strings and back, keeps a handle table, and turns exceptions into C error
codes. It contains no protocol logic and must never acquire any. A bug there
could still corrupt data on the way in or out — marshalling the wrong string,
leaking memory, mishandling a handle. The proofs say nothing about it.

**Also outside the proofs**, as with any extracted code: the Dafny C# backend,
the .NET ILCompiler, the Dafny runtime library, and the assumption that the
specification says what the protocol authors meant.

**Coverage is partial.** The verified core models the root reducer and five
channels completely; `session` (36 of 61 actions) and `chat` (54 of 97) are
modeled in part. Actions the core does not model are not rejected — they decode
to a defined `RootUnknown` catch-all. Call `ahp_action_channel()` and check for
`"root/unknown"` to detect them. See [`../VERIFICATION.md`](../VERIFICATION.md).

In one sentence: **the reducers are proven; the binding around them is not.**

---

## Building

Requires .NET SDK 10 (or 8+ with a NativeAOT workload), clang++ with C++17, and
the Xcode command line tools. CMake 3.20+ is optional.

```bash
./scripts/build.sh          # publish, stage, build the demo, run it
./scripts/build.sh --cmake  # same, but build the demo through CMake
```

The script wipes `lib/`, `build/`, and the .NET intermediates first, so it is a
true from-scratch reproduction.

Via CMake directly:

```bash
cmake -S . -B build -DCMAKE_BUILD_TYPE=Release
cmake --build build
ctest --test-dir build --output-on-failure
cmake --install build --prefix /usr/local
```

`-DAHP_BUILD_NATIVE_LIB=ON` makes CMake run `dotnet publish` itself rather than
consuming a prebuilt `lib/libahp_core.dylib`.

---

## Using it from C++

```cpp
#include "ahp_core.h"
#include <cstdio>

int main() {
    char *state = ahp_initial_state_json();

    char *next = ahp_apply_json(
        state, R"({"type":"terminal/resized","cols":120,"rows":40})");

    if (!next) {
        std::fprintf(stderr, "ahp: %s\n", ahp_last_error());
        return 1;
    }

    std::printf("%s\n", next);

    ahp_string_free(next);
    ahp_string_free(state);
}
```

```bash
clang++ -std=c++17 -Iinclude app.cpp lib/libahp_core.dylib \
        -Wl,-rpath,@executable_path/lib -o app
```

### API

| Function | Purpose |
|---|---|
| `ahp_version()` | Library version. Library-owned. |
| `ahp_last_error()` | Last error on this thread. Library-owned. |
| `ahp_string_free(s)` | Frees a returned string. |
| `ahp_initial_state_json()` | Default state as JSON. |
| `ahp_apply_json(state, action)` | `applyAhp` — one action. |
| `ahp_fold_json(state, actions)` | `foldAhp` — a JSON array of actions. |
| `ahp_action_roundtrip_json(action)` | `encode(decode(a))`; proven identity. |
| `ahp_action_channel(action)` | Routing channel, or `"root/unknown"`. |
| `ahp_state_new()` / `ahp_state_from_json()` | Allocate a state handle. |
| `ahp_state_apply(h, action)` | Apply in place; `0` ok, `-1` error. |
| `ahp_state_to_json(h)` | Encode a handle's state. |
| `ahp_state_equals(a, b)` | The core's structural equality. |
| `ahp_state_free(h)` | Release a handle. |

Two styles are offered because they suit different callers. The pure JSON
functions are total functions of their arguments and hold no state, which makes
them easy to reason about. The handle functions keep decoded state inside the
library so a stream of actions does not pay to re-parse and re-encode the whole
state on every step.

### Conventions

- Strings are NUL-terminated UTF-8.
- `char*` returns transfer ownership — free with `ahp_string_free()`.
  `const char*` returns are library-owned.
- A `NULL` return (or `-1`) means failure; `ahp_last_error()` explains it.
  Errors are per-thread and valid until the next call on that thread.
- **Emitted JSON is canonical: object keys are sorted.** Round-tripping
  preserves the JSON *value*; it is byte-stable only for input already in
  canonical form. Compare outputs against outputs, not against hand-written
  input literals.
- Malformed JSON fails loudly. The core's `ParseJson` silently yields `JNull`
  on bad input; the binding uses `ParseJsonChecked` instead so callers get an
  error rather than a silent empty reduction.

---

## What was verified while building this

Claims here are backed by command output in [`evidence/`](evidence/).

**NativeAOT compiles the extracted core with no trim or reflection warnings.**
The project sets `IlcTreatWarningsAsErrors=true`, so the build fails if the
verified core or the binding needs reflection or dynamic codegen that AOT
cannot honour. It builds clean. The Dafny C# backend emits fully
AOT-compatible code — worth recording, since that was the main risk in this
route.

**The library needs no .NET runtime.** `otool -L` shows only macOS system
libraries; no `libcoreclr`, `hostfxr`, or `hostpolicy`. It is a 2.2 MB
self-contained native object.

**Both hosts agree.** The C# smoke test runs the same core on the CLR; this
C++ demo runs it AOT-compiled through the C ABI. Both report `cols=120`,
`title="build"`, `cwd="/src"`, `foldAhp` equal to successive `applyAhp`, and
identical results across repeated folds.

**The demo asserts protocol properties, not just liveness** — codec identity,
canonicalisation, idempotence, fold/apply agreement, determinism, cross-channel
routing, and loud rejection of malformed input. All 22 assertions pass; the
demo exits non-zero if any fails, so `ctest` treats it as a real test.

---

## Layout

```
cpp-via-dotnet/
├── CMakeLists.txt              build manifest
├── README.md
├── LICENSE
├── include/ahp_core.h          the C ABI
├── src/
│   ├── Ahp.Core.Native.csproj  NativeAOT shared-library project
│   └── AhpNative.cs            hand-written binding (the unverified part)
├── examples/ahp_demo.cpp       C++ consumer
├── scripts/build.sh            from-scratch reproduction
├── lib/libahp_core.dylib       generated artifact
└── evidence/                   captured command output
```

`lib/`, `build/`, and the .NET intermediates are generated; `scripts/build.sh`
recreates them.

---

## Portability

Built and verified on macOS arm64. The route itself is not macOS-specific —
NativeAOT targets `linux-x64`, `linux-arm64`, and `win-x64` too, and the C ABI
and C++ source are portable. Change `RuntimeIdentifier` in
`src/Ahp.Core.Native.csproj` and the library name in `CMakeLists.txt`
(`.so` / `.dll`). Those targets have not been built or tested here, so they are
listed as expected to work rather than as something demonstrated.

---

## License

MIT. Copyright (c) Microsoft Corporation (protocol and original client);
Copyright (c) 2026 Josh Mouch (verification, extraction, and this
distribution). See [`LICENSE`](LICENSE).
