# @open-agency/ahp

Formally-verified [Agent Host Protocol](https://github.com/microsoft/agent-host-protocol) (AHP)
channel reducers for JavaScript, extracted from a machine-checked Dafny core.

This package is **not** a hand-written reimplementation of the protocol. It is the
JavaScript extraction of a Dafny development in which the seven AHP channel
reducers, their convergence/fold properties, and a `Std.JSON` codec are proven
correct, then translated to source by the Dafny compiler. The reducer logic you
call here is the same logic the proofs are about.

## Install

```sh
npm install @open-agency/ahp
```

Requires Node.js >= 18. The only runtime dependency is `bignumber.js`, which backs
the Dafny arbitrary-precision integer representation (AHP protocol datatypes use
unbounded `int`/`nat`, so fixed-width JS numbers are not sufficient).

## Usage

The package is CommonJS.

```js
const ahp = require('@open-agency/ahp');

// The seven AHP channel reducers, plus the shared skeleton:
//   ahp.Chat, ahp.Session, ahp.Terminal, ahp.Canvas,
//   ahp.Changeset, ahp.Annotations, ahp.ResourceWatch, ahp.AhpSkeleton

// Each channel exposes its reducer as ApplyTo<Channel>(state, action, now):
const next = ahp.Chat.__default.ApplyToChat(state, action, now);

// The Dafny runtime handles are exported so you can build argument values:
//   ahp._dafny      sequences, maps, sets, integer helpers
//   ahp.BigNumber   arbitrary-precision integers
//   ahp.ConfluxCodec  the proven JSON codec surface
const emptyTurns = ahp._dafny.Seq.of();
ahp.Chat.__default.hasTurn(emptyTurns, ahp._dafny.Seq.UnicodeFromString('id')); // false
```

Values crossing the boundary are Dafny representations (`_dafny.Seq` for
sequences and strings, `BigNumber` for integers), not plain JS values. This is a
deliberate consequence of shipping the extracted core verbatim: no hand-written
marshalling layer sits between you and the proven code, so there is nothing
unverified to disagree with the proofs. A separate idiomatic wrapper can be
layered on top if you want plain-JS ergonomics.

## Verification status

Each channel module carries `RunCorpus()`, the curated in-specification fixture
corpus, which runs entirely in-process:

```
$ npm test
AhpSkeleton    7/7
ResourceWatch  2/2
Canvas         5/5
Changeset      15/15
Annotations    10/10
Terminal       19/19
Session        36/36
Chat           54/54
corpus OK — 148/148 fixtures green against extracted code
```

Beyond that in-process corpus, the core is also driven by a reducer-replay
harness that reads the real upstream fixture JSON, decodes it with the proven
`Std.JSON` codec, folds it through the actual reducer, and asserts full-state
`reduced == expected`. That harness requires a checkout of the upstream
`agent-host-protocol` fixtures and so is not part of this package's `npm test`.
See the build notes below for the measured results and a known divergence.

## Building from source

`dist/ahp.cjs` is generated, not hand-edited. To rebuild it you need Dafny 4.11.0
and a checkout of the Dafny core:

```sh
AHP_DAFNY_CORE=/path/to/agent-host-protocol-core-dafny node scripts/build.mjs
```

The build translates `spec/client_main.dfy` with `dafny translate js`, linking the
Conflux runtime as a library, then converts the translated bundle into a CommonJS
library (prepend runtime, strip the auto-run entry point, append an export
surface derived from the bundle's own module bindings). See
`scripts/build.mjs` for the details.

## License

MIT. See [LICENSE](LICENSE).

Copyright (c) Microsoft Corporation — the protocol and original client.
Copyright (c) 2026 Josh Mouch — the Dafny verification and this port.
