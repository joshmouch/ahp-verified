# ahp-core

Formally verified core of the **Agent Host Protocol (AHP)**, for the JVM.

The reducers, convergence/fold proofs, and wire codec in this artifact are not hand-written.
They are mechanically extracted from a Dafny specification and proven against a 211-fixture
reducer corpus covering seven AHP channels.

| | |
|---|---|
| Coordinate | `agency.open.ahp:ahp-core` |
| Version | `0.1.0` |
| Java baseline | 17 |
| Runtime dependency | `org.dafny:DafnyRuntime:4.11.0` (Maven Central) |
| License | MIT (dual copyright — see `LICENSE`) |

## Install

```xml
<dependency>
  <groupId>agency.open.ahp</groupId>
  <artifactId>ahp-core</artifactId>
  <version>0.1.0</version>
</dependency>
```

Gradle:

```groovy
implementation 'agency.open.ahp:ahp-core:0.1.0'
```

`org.dafny:DafnyRuntime` is a transitive dependency and is resolved from Maven Central; you do
not need a local Dafny installation to *consume* this artifact.

## What is in it

Everything is namespaced under `agency.open.ahp`. The seven verified AHP channel modules:

| Package | Channel |
|---|---|
| `agency.open.ahp.AhpSkeleton` | Root envelope + shared option/result types |
| `agency.open.ahp.ResourceWatch` | Resource watch |
| `agency.open.ahp.Canvas` | Canvas |
| `agency.open.ahp.Changeset` | Changeset |
| `agency.open.ahp.Annotations` | Annotations |
| `agency.open.ahp.Terminal` | Terminal |
| `agency.open.ahp.Session` | Session |
| `agency.open.ahp.Chat` | Chat (tool-call state machine + turn lifecycle) |
| `agency.open.ahp.ConfluxCodec` | The one JSON value model (`Json`, `JObj`, `JArr`, …) and `Codec` |

Plus the supporting verified operator packages the reducers depend on: `ConfluxContract`,
`ConfluxOperators`, `ConfluxOrderedLog`, `ConfluxPrune`, `ConfluxSeqRoute`.

### Scope: protocol core only, no transport

This artifact ships the **dependency closure of the AHP reducers** — 15 packages — and
deliberately excludes the wider conflux runtime service layer (`ConfluxHttpService`,
`ConfluxJsonRpc`, `ConfluxResourceSupervisor`, `ConfluxServiceLifecycle`, and ~33 others).

That exclusion is not arbitrary. Those service modules bottom out in `{:extern}` capability
modules — `ConfluxIoCapability`, `ConfluxHttpCapability`, `ConfluxCommandIdentityCapability` —
whose bodies are hand-written per target language. Upstream ships extern implementations for
**C# and JavaScript only** (`extern/cs`, `extern/js`); there is no JVM implementation, so those
modules cannot be built for Java at all without first porting the I/O adapters (socket,
websocket, HTTP, process identity).

The AHP protocol core does not need any of them. The reducers are pure state machines over the
JSON value model, so the closure is fully buildable on the JVM with nothing stubbed or faked.
Transport is the host application's concern: feed decoded frames to the reducers.

## Verifying a build

The jar is executable and replays the fixture corpus:

```bash
java -cp ahp-core-0.1.0.jar:DafnyRuntime-4.11.0.jar agency.open.ahp.CorpusRunner
```

Observed output from the packaged jar:

```
ROOT CORPUS:          7/7 green against extracted code
RESOURCEWATCH CORPUS: 2/2 green against extracted code
CANVAS CORPUS:        5/5 green against extracted code
CHANGESET CORPUS:     15/15 green against extracted code
ANNOTATIONS CORPUS:   10/10 green against extracted code
TERMINAL CORPUS:      19/19 green against extracted code
SESSION CORPUS:       36/36 modeled green (of 61 total; all ~25 action TYPES now modeled)
CHAT CORPUS:          54/54 modeled green (of 97 total; full tool-call state machine + turn lifecycle modeled)
TOTAL: 148/148 corpus fixtures green (5 full AHP channels + session/chat partial)
MULTI-CHANNEL CLIENT OK — extract(cs,js) + corpus all green
```

This is the cheapest end-to-end confirmation that a build behaves like the verified
specification. Note the counts are the specification's own: Session models 36 of 61 fixtures
and Chat 54 of 97, so 148 is the modeled subset that is green, not the full 211-fixture corpus.

## Building from source

Requires a JDK and Maven. Consumers only need a Java 17 runtime — the build targets Java 17
bytecode via `maven.compiler.release`.

```bash
mvn package     # ~7s
mvn test
```

> **This project builds with ecj, not javac — that is required, not a preference.**
>
> `Chat/__default.java` contains the extracted Chat reducer as a single ~121 KB / 1273-line
> method (`ApplyToChat`) carrying hundreds of lambdas over deeply nested wildcard generics
> (`DafnySequence<? extends DafnySequence<? extends CodePoint>>`) against the high-arity
> `dafny.Function7/8/9` interfaces. `javac` degenerates on this inside
> `Check.checkValidGenericType` / `firstIncompatibleTypeArg`, reached through the
> functional-interface descriptor cache.
>
> Measured on this source set: **javac 17 and javac 26 each ran >15 minutes at >17 GB resident
> without ever emitting the class** — whether compiling all 192 sources together or that one
> file in isolation, and at heaps up to 48 GB. Raising the heap and changing JDK both failed;
> the cost is CPU, not memory.
>
> **ecj compiles the same 192 sources in 0.8 seconds.** The `maven-compiler-plugin` is
> therefore configured with `<compilerId>eclipse</compilerId>` and `plexus-compiler-eclipse`.
> Output is ordinary Java 17 bytecode, so consumers are unaffected by this choice.
>
> The remaining 189 of 192 files compile under javac in ~2.4 seconds; the pathology is confined
> to that one generated file.

### Regenerating the extracted sources

The contents of `src/main/java/agency/open/ahp/` (except `CorpusRunner.java`, which is
hand-written) are generated from the Dafny core with Dafny 4.11.0. Two translation passes are
required, and **both must use the same `--outer-module`** so their cross-references line up:

```bash
# 1. The conflux-runtime library (supplies agency.open.ahp.ConfluxCodec)
dafny translate java conflux-runtime.doo \
  --no-verify --outer-module agency.open.ahp -o build/lib

# 2. The AHP core, linked against that library
dafny translate java spec/client_main.dfy \
  --no-verify --outer-module agency.open.ahp \
  --library conflux-runtime.doo -o build/core
```

Then merge both `agency/` trees into `src/main/java/`.

#### Why the library must be translated separately

`--library` tells Dafny the library is *already compiled*, so its code is deliberately **not**
emitted. That works for the C# and JavaScript targets because the upstream runtime ships
prebuilt `conflux-runtime.dll` and `conflux-runtime.cjs` artifacts to link against. **There is
no prebuilt JVM artifact.** Translating the core alone therefore produces sources that
reference a `ConfluxCodec` package that does not exist, and the build fails with several
hundred `package ConfluxCodec does not exist` errors. Translating the `.doo` to Java as a
separate pass supplies the missing package.

#### Why the compiler is configured with a large heap

The extracted corpus reducers compile into a small number of very large methods —
`Chat/__default.java` alone is ~380 KB with single expressions thousands of characters wide.
`javac` needs a large heap to attribute them, so `maven-compiler-plugin` is configured to fork
with `-Xmx8g`. Compilation is correspondingly slow; this is expected, not a hang.

## License

MIT, dual copyright:

- Copyright (c) Microsoft Corporation — the protocol and original client
- Copyright (c) 2026 Josh Mouch — the Dafny verification and this port
