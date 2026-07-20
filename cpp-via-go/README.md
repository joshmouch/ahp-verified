# AHP verified core — C++ distribution via a Go c-archive

**Status: WORKING.** A C++ program links a native library and calls the
Dafny-extracted AHP reducers. The reducers that execute are the proven ones.

[`../cpp/`](../cpp/) records why `dafny translate cpp` cannot produce this: the
Dafny 4.11.0 C++ backend supports neither unbounded `int`/`nat` (the core uses
60 `: int` and 47 `nat` annotations) nor arrow types (26 signatures), and
bounding the integers would break the fold proofs. That conclusion stands and
is not revisited here.

This package takes a different route to the same destination. Go is one of the
five targets the core already extracts to cleanly, and Go can emit a native
static archive with a C ABI. So the verified reducers are compiled to machine
code by the Go toolchain, wrapped in a small C entry-point layer, and linked
into a C++ program like any other native library.

```
spec/*.dfy  ──dafny translate go──▶  ../go/**  ──go build -buildmode=c-archive──▶  libahpverified.a
  (proven)         (mechanical)      (extracted)          (mechanical)              + libahpverified.h
                                                                                          │
                                                             bridge/ahp_bridge.go ─────────┤
                                                             (hand-written, ~200 lines)    │
                                                                                           ▼
                                                                            your C++ ──▶ ahp_verified.h
```

---

## What the trust chain actually is

This is the section to read before believing anything else in this document.

| Layer | Machine-derived from the proofs? |
|---|---|
| The reducers, codecs and identity functions that compute every result | **Yes.** `../go/**` is `dafny translate go` output, unedited. |
| Their compilation to arm64/amd64 machine code | **Yes**, by the Go compiler — the same trust you extend to any compiler. |
| `bridge/ahp_bridge.go` — the C entry points | **No. Hand-written.** ~200 lines of marshalling. |
| `include/ahp_verified.h` — the C declarations | **No. Hand-written**, checked against the generated header on every build. |
| `examples/*.cpp` | **No. Hand-written**, and not part of the library. |

**The honest claim is: a proven core reached through an unproven binding.** No
theorem covers the bridge. If it converts a string wrongly, or leaks a handle,
or maps `AHP_AVAILABILITY_STALE` to the wrong Dafny constructor, no proof
catches it — the proofs are about what happens *after* the call lands.

Three things bound that risk, none of which eliminate it:

1. **The bridge decides nothing.** Every C function is convert-in, call-verified-
   function, convert-out. There is no protocol logic in it: no field
   precedence, no merge, no validation, no defaulting. When the specification
   says a sparse `Updated` leaves `title` alone, that behaviour comes from
   `Canvas.ApplyToCanvas`, and the bridge is not consulted.
2. **The surface is deliberately small** — 16 exported functions. Opaque handles
   carry Dafny values across the boundary untouched, so no datatype is
   marshalled field-by-field and there is nowhere for a re-implementation to
   hide. The cost is that the C API is narrower than the Go one; that trade is
   intentional.
3. **The example asserts, and fails loudly.** `examples/ahp_demo.cpp` returns a
   non-zero exit status on any failed check, and CMake registers it as a test.

What is *not* claimed: that the C ABI is verified, that memory-safety of the
binding is proven, or that the C++ side is anything but ordinary code.

---

## Quick start

Requires a Go toolchain (1.25+), clang++ or g++ with C++17, and CMake 3.20+.

```bash
./regenerate.sh                        # build the library, then compile and run the example
```

or, through CMake:

```bash
cmake -S . -B build/cmake
cmake --build build/cmake
ctest --test-dir build/cmake --output-on-failure
```

`CMakeLists.txt` invokes `regenerate.sh` automatically when `lib/` is empty, so
a fresh checkout needs no separate step.

### Using it from your own project

```bash
cmake -S . -B build/cmake -DCMAKE_INSTALL_PREFIX=/usr/local
cmake --build build/cmake && cmake --install build/cmake
```

```cmake
find_package(ahp_verified REQUIRED)
target_link_libraries(my_app PRIVATE ahp::verified)
```

```cpp
#include "ahp_verified.h"

uintptr_t canvas = ahp_canvas_new("canvas-1", "acme.editors");

int outcome = -1;
uintptr_t next = ahp_canvas_apply_updated(canvas, "Design Review", "editing",
                                          nullptr, AHP_AVAILABILITY_UNCHANGED,
                                          /*now=*/1, &outcome);

char* title = ahp_canvas_title(next);   // "Design Review"
ahp_string_free(title);
ahp_canvas_release(next);
ahp_canvas_release(canvas);             // the reducer is pure; this is still valid
```

`examples/downstream/` is exactly this, as a standalone project that resolves
the package with `find_package` — it exists to prove the installed artifact is
consumable from outside this tree, not just from inside it.

---

## The C surface

Sixteen functions. Full documentation is in
[`include/ahp_verified.h`](include/ahp_verified.h).

| Group | Functions | Verified code reached |
|---|---|---|
| Canvas reducer | `ahp_canvas_new`, `ahp_canvas_apply_updated`, `ahp_canvas_apply_close_requested`, `ahp_canvas_{id,provider_id,title,activity,content_uri,availability}`, `ahp_canvas_equals`, `ahp_canvas_release` | `Canvas.ApplyToCanvas` |
| JSON codec | `ahp_json_canonicalize` | `ConfluxJsonText.ParseJsonChecked`, `.Stringify` |
| Command identity | `ahp_sha256_digest` | `ConfluxCommandIdentityCapability.Sha256Digest` |
| Conformance corpus | `ahp_run_corpus` | `Session.RunCorpus` |
| Housekeeping | `ahp_abi_version`, `ahp_string_free` | — |

**Ownership.** Every returned `char*` is heap-allocated and released with
`ahp_string_free`; `NULL` means the specification-level `Option` was `None`, or
the input was refused. Every `uintptr_t` handle is released exactly once with
`ahp_canvas_release`. Because the reducers are pure, applying an action neither
invalidates nor modifies the input handle — it returns a new one, and both must
be released.

`examples/ahp_demo.cpp` wraps both in small RAII types; that is the
recommended pattern, and it is about twenty lines.

---

## Evidence

Real command output for every claim above, under [`evidence/`](evidence/):

| | What it shows |
|---|---|
| [E1](evidence/E1-regenerate-from-scratch.txt) | `./regenerate.sh` from an empty `lib/`: both build modes, the ABI drift check, then the example compiled and run |
| [E2](evidence/E2-cpp-demo-run.txt) | The compiled C++ program's output, exit status 0 |
| [E3](evidence/E3-cmake-build-test-install.txt) | `cmake` configure, build, `ctest` and `cmake --install` |
| [E4](evidence/E4-downstream-find-package.txt) | A separate project resolving the installed package with `find_package` and calling the reducer |
| [E5](evidence/E5-symbol-provenance.txt) | `nm` over the archive: 1050 `ahp-verified/go` symbols including `Canvas.ApplyToCanvas`, against 16 hand-written `ahp_*` exports |

The C++ program's output, in full:

```
[1] Canvas reducer (Canvas.ApplyToCanvas)
    initial : canvasId=canvas-1 title=None activity=None contentUri=None availability=Ready
    Updated : canvasId=canvas-1 title=Design Review activity=editing contentUri=ahp-session:/2f9c/content/canvas-1 availability=Ready
    outcome: Applied
    degraded: canvasId=canvas-1 title=Design Review activity=editing contentUri=ahp-session:/2f9c/content/canvas-1 availability=Stale
    CloseRequested outcome: NoOp

[2] JSON codec (ConfluxJsonText.ParseJsonChecked / Stringify)
    input:     {"seq": 42, "channel":"canvas"}
    canonical: {"channel":"canvas","seq":42}

[3] Command identity (ConfluxCommandIdentity.Sha256Digest)
    digest("abc") = sha256:ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad

[4] Embedded corpus (Session.RunCorpus)
    corpus: 36/36 modeled cases passed

ALL CHECKS PASSED (0 failures)
```

Each of those is a decision made inside the verified core: the sparse update
preserving `title`, `CloseRequested` reporting `NoOp`, the canonical key
ordering, the digest matching the published SHA-256 vector.

---

## Honest limits

Shipped deliberately, because you would find them anyway.

- **The binding is not verified.** Stated above; restated here because it is the
  limit that matters most.
- **Your process embeds the Go runtime.** The archive is ~5.9 MB and brings a
  garbage collector, a scheduler and its own signal handlers into your address
  space. This is the price of the route and it is not small. It rules the
  package out for freestanding, hard-real-time and `fork`-heavy contexts, and
  means `libahpverified.a` does not compose with a second Go c-archive in the
  same binary.
- **One channel of eight is exposed.** Only Canvas has a C surface. The core
  proves host-authority laws for all 8 channels and the extracted Go exposes
  all of them; the C ABI simply has not been widened yet. That is a scope
  decision, not a limitation of the route.
- **`ahp_run_corpus` reports 36 cases, not 237.** `Session.RunCorpus` is the
  curated modeled subset embedded in the extracted core, and 36 of session's 61
  cases is what it covers. The 237-fixture replay corpus runs against C# and JS
  and is not embedded in the Go extraction, so it is not reachable from here.
  See the root [README](../README.md) "Honest limits".
- **`-gcflags=all=-l` is mandatory.** The extracted `Chat` package defeats the
  Go inliner, which is OOM-killed while compiling it. Note `all=`: applying
  `-gcflags=-l` without it covers only the package named on the command line
  and the dependency build still dies. `regenerate.sh` carries the right flags.
- **Built and tested on arm64 macOS.** The route is not platform-specific —
  `-buildmode=c-archive` is supported on Linux and Windows too, and
  `CMakeLists.txt` carries the ELF link dependencies — but only the Darwin path
  has actually been run here.
- **`include/ahp_verified.h` is hand-written and could drift** from the
  cgo-generated ABI. `regenerate.sh` compares the two symbol sets on every
  build and fails on a mismatch, which is a check rather than a proof.

---

## Layout

```
bridge/ahp_bridge.go     the C entry points          hand-written
include/ahp_verified.h   the public C header         hand-written
examples/ahp_demo.cpp    the example and smoke test  hand-written
examples/downstream/     a standalone consumer       hand-written
CMakeLists.txt           build manifest              hand-written
regenerate.sh            reproduces everything below hand-written

include/libahpverified.h  generated by cgo
lib/libahpverified.a      generated — static archive, link this
lib/libahpverified.dylib  generated — shared library, if you prefer it
```

`lib/` is git-ignored: any checkout with a Go toolchain reproduces it, and
CMake does so automatically. The generated header is committed as a record of
the ABI.

---

## License

MIT. Copyright (c) Microsoft Corporation for the protocol and the original
client; Copyright (c) 2026 Josh Mouch for the verification, the extraction and
this binding. See [LICENSE](LICENSE).

This is a third-party community project and is not an official Microsoft
product.
