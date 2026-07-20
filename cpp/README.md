# AHP verified core — C++ target

**Status: BLOCKED. The AHP Dafny core cannot be translated to C++ by Dafny 4.11.0.**

This is not a "we ran out of time" result and not a packaging gap. Two of the
blockers found here were fixable and *were* fixed; the remaining two are
properties of the Dafny C++ backend that no flag, entry point, or workaround
reaches. This document records exactly how far the chain got and where it stops,
with the real command output for each claim under [`evidence/`](evidence/).

No C++ package is published. The locked distribution matrix assigns no C++
coordinate, and inventing one for an artifact that does not exist would be worse
than shipping nothing.

## Summary

| # | Blocker | Verdict | Evidence |
|---|---|---|---|
| 1 | `--unicode-char` default rejects cpp | **FIXED** — pass `--unicode-char false` | [E1](evidence/E1-translate-default.txt) |
| 2 | `DafnyRuntime.h` does not compile (upstream bug) | **FIXED** — one-token patch, [`vendor/DafnyRuntime.h.patch`](vendor/DafnyRuntime.h.patch) | [E5](evidence/E5-runtime-header-bug.txt), [E6](evidence/E6-patched-runtime-works.txt) |
| 3 | Unbounded integers (`int` / `nat`) unsupported | **IRREDUCIBLE** | [E3](evidence/E3-version-unbounded-int.txt), [E4](evidence/E4-feature-probes.txt) |
| 4 | Arrow / higher-order types unsupported | **IRREDUCIBLE** | [E4](evidence/E4-feature-probes.txt) |
| 5 | Every shipped `.doo` library is built `--unicode-char true` | **IRREDUCIBLE** (catch-22 with #1) | [E2](evidence/E2-library-unicode-conflict.txt) |

Feature census of the core, which is what makes #3 and #4 fatal rather than
inconvenient: **60** `: int` annotations, **47** `nat`, **26** arrow (`->`)
signatures, and JSON codec (`Wire.`) use in 8 of 10 spec modules.

## Baseline: the invocation is correct

Before blaming the backend, the same command shape was proven against the C#
backend, which is one of the two proven-extracting targets:

```
$ dafny translate cs spec/client_main.dfy --no-verify \
    --library .conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime.doo \
    -o /tmp/ahp-cs-baseline

Dafny program verifier did not attempt verification
-rw-r--r--  1009723  /tmp/ahp-cs-baseline.cs
```

1,009,723 bytes of C#. So the library path, the flags, and the entry point are
all right; swapping `cs` for `cpp` is the only change in what follows.

**Library used:** `.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime.doo`
— the generic build. There is no `conflux-runtime-cpp.doo`; a repo-wide search
for `*.doo` found per-target variants only for `js`, so the generic one is the
correct and only choice for this target.

## Blocker 1 — Unicode chars (FIXED)

```
$ dafny translate cpp spec/client_main.dfy --no-verify --library <conflux-runtime.doo> -o /tmp/x
spec/client_main.dfy(3,0): Error: Feature not supported for this compilation target: Unicode chars
```

Dafny 4.x defaults `--unicode-char true`; the C++ backend requires the legacy
char model. Passing `--unicode-char false` clears this error. That fix is
carried through every command below.

## Blocker 2 — the shipped C++ runtime header does not compile (FIXED)

This one is worth stating plainly, because it means **no Dafny program at all,
of any complexity, compiles to a working C++ binary on a stock Dafny 4.11.0
install.** Even a hello-world:

```
$ g++ -std=c++17 -I $(brew --prefix dafny)/libexec/DafnyRuntimeCpp -c out_ok.cpp
DafnyRuntime.h:550:23: error: no member named 'find' in 'DafnySet<T>'
  550 |             if (other.find(elt) != other.set.end()) {
      |                 ~~~~~ ^
```

Identical on `g++ -std=c++14/17/20` and `clang++ -std=c++17`. `DafnySet<T>`
wraps a `std::set` member named `set`; every sibling method goes through it
(`contains()` calls `set.find(t)`), so `disjoint()` calling `other.find(elt)` is
a plain omission. The fix is `other.set.find(elt)` — see
[`vendor/DafnyRuntime.h.patch`](vendor/DafnyRuntime.h.patch), which is a genuine
upstream bug fix independent of AHP.

With the patched header vendored here, translation → compile → link → run works:

```
$ g++ -std=c++17 -I vendor -I smoke smoke/probe_ok.cpp -o build/ahp_cpp_toolchain_probe
exit=0
$ ./build/ahp_cpp_toolchain_probe
hello from cpp
exit=0
```

**The toolchain is therefore proven functional.** Everything that fails from
here fails on language features, not on environment.

## Blocker 3 — unbounded integers (IRREDUCIBLE)

`int` and `nat` are hard-rejected:

```
$ dafny translate cpp probe_int.dfy --no-verify --unicode-char false
probe_int.dfy(2,25): Error: Feature not supported for this compilation target: Unbounded integers
  |
2 |   datatype Thing = Thing(id: int, count: nat)
  |                          ^^
```

The sharpest demonstration is `spec/version.dfy` — the *simplest* module in the
core, a SemVer registry, and one of the few that imports nothing at all, so the
library question does not even arise:

```
$ dafny translate cpp spec/version.dfy --no-verify --unicode-char false
spec/version.dfy(49,11): Error: Feature not supported for this compilation target: Unbounded integers
   |
49 |   function compareSemVer(a: SemVer, b: SemVer): int
spec/version.dfy(51,39): Error: Feature not supported for this compilation target: Non-native numeric newtypes
   |
51 |     if a.major != b.major then a.major - b.major
```

If the easiest, dependency-free module in the core cannot cross, no amount of
entry-point selection helps.

**Why substituting bounded ints is not a workaround.** The `nat` typing is not
incidental; it carries the proof. `version.dfy`'s own header states that
`MAJOR.MINOR.PATCH` well-formedness "is a LANGUAGE GUARANTEE here, not a regex:
`SemVer` carries three `nat` components by construction". Replacing `nat` with
`i32` deletes the property the verification exists to establish. A C++ artifact
obtained that way would compile while no longer being the verified core — which
is the only thing this release is for.

## Blocker 4 — higher-order functions (IRREDUCIBLE)

Independent of the integer problem. With deliberately bounded (`i32`) types, so
that BigInteger cannot be the cause:

```
$ dafny translate cpp probe_ho32.dfy --no-verify --unicode-char false
(0,-1): Error: UserDefinedTypeName _#Func1
```

The backend cannot name the arrow type at all. The core has 26 arrow signatures;
the fold/reducer architecture is built on them.

## Blocker 5 — the library catch-22 (IRREDUCIBLE)

Blocker 1's fix collides with the dependency:

```
$ dafny translate cpp spec/client_main.dfy --no-verify --unicode-char false --library <conflux-runtime.doo>
CLI: Error: cannot load .../conflux-runtime.doo:
  --unicode-char is set locally to False, but the library was built with True
```

The core genuinely needs it — dropping `--library` yields `module ConfluxCodec
does not exist`, plus the same for `ConfluxContract`, `ConfluxOperators`,
`ConfluxOrderedLog`, and `ConfluxSeqRoute`. The identical conflict blocks
`DafnyStandardLibraries.doo`, so **Std.JSON is unreachable from C++ too**:

```
CLI: Error: cannot load /DafnyStandardLibraries.doo:
  --unicode-char is set locally to False, but the library was built with True
```

So: C++ requires `--unicode-char false`, and every shipped `.doo` on the machine
is built with `true`. Rebuilding `conflux-runtime` from source with
`--unicode-char false` would not rescue this — blockers 3 and 4 are demonstrated
on library-free code and would still stand.

## What is in this directory

```
CMakeLists.txt      Build manifest. Builds the toolchain probe, NOT the AHP core.
LICENSE             MIT, dual copyright (Microsoft Corporation; Josh Mouch).
vendor/
  DafnyRuntime.h        Patched Dafny 4.11.0 C++ runtime header.
  DafnyRuntime.h.patch  The upstream fix, as a standalone patch.
smoke/
  probe_ok.dfy      Toolchain probe source; translates + compiles + runs.
  probe_ok.cpp/.h   Its emitted C++.
  probe_int.dfy     Minimal repro of blocker 3.
  probe_ho32.dfy    Minimal repro of blocker 4.
evidence/           Raw captured output, E1-E7.
```

## Reproducing

```
cd <core>/agent-host-protocol-core-dafny
LIB=.conflux/runtime/dependencies/conflux-runtime/current/conflux-runtime.doo

dafny translate cpp spec/client_main.dfy --no-verify --library $LIB -o /tmp/x   # blocker 1
dafny translate cpp spec/client_main.dfy --no-verify --unicode-char false \
      --library $LIB -o /tmp/x                                                  # blocker 5
dafny translate cpp spec/version.dfy --no-verify --unicode-char false -o /tmp/x # blocker 3
```

## Honest limitations of this report

- **`CMakeLists.txt` is unverified.** `cmake` is not installed on this machine
  (`cmake not found`). The compile+link+run it encodes was executed directly
  with `g++` and passes ([E7](evidence/E7-package-build.txt)), but CMake itself
  has not been run.
- **The smoke test exercises the toolchain, not AHP.** It proves a Dafny-emitted
  C++ binary runs; it does not exercise a reducer, because no AHP reducer can be
  emitted. Claiming otherwise would misrepresent what was built.
- **No attempt was made to rebuild `conflux-runtime` from source** with
  `--unicode-char false`. That path was assessed and abandoned as pointless
  rather than attempted: blockers 3 and 4 are demonstrated on code that uses no
  library at all, so a rebuilt library cannot change the outcome. This is a
  reasoned skip, not a completed experiment.

## Conclusion

C++ is blocked by the Dafny 4.11.0 backend's lack of unbounded integers and
arrow types. Both are load-bearing in the AHP core, and both are load-bearing
*because of the verification* rather than by stylistic accident — `nat` is what
makes well-formedness a language guarantee, and arrow types are what the fold
architecture is made of. Any C++ artifact produced by stripping them would
compile without being the verified core.

The two fixable blockers were fixed, and the runtime-header patch is worth
upstreaming on its own merits, since it currently breaks the C++ backend for
every Dafny user. Revisit C++ if a future Dafny release adds BigInteger and
first-class function support to the backend; `CMakeLists.txt` is the harness the
real core would drop into.
