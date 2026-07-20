# ahp-go

Go binding for the Agent Host Protocol (AHP) core, mechanically extracted from a
formally verified Dafny specification.

```
go get github.com/joshmouch/ahp-go
```

The seven AHP channel reducers, the convergence and fold proofs, and the JSON
codec are not reimplemented here — they are translated from the same proven
source that produces the C# and JavaScript bindings, so all three agree by
construction rather than by test.

## Status

`go build ./...` succeeds across all 61 packages, and the extracted client runs
the reducer conformance corpus green:

```
$ go run ./cmd/ahp
ROOT CORPUS:          7/7 green against extracted code
RESOURCEWATCH CORPUS: 2/2 green against extracted code
CANVAS CORPUS:        5/5 green against extracted code
CHANGESET CORPUS:     15/15 green against extracted code
ANNOTATIONS CORPUS:   10/10 green against extracted code
TERMINAL CORPUS:      19/19 green against extracted code
SESSION CORPUS:       36/36 modeled green (of 61 total; all ~25 action TYPES now modeled)
CHAT CORPUS:          54/54 modeled green (of 97 total; full tool-call state machine + turn lifecycle modeled)
TOTAL: 148/148 corpus fixtures green (5 full AHP channels + session/chat partial)
```

The Session and Chat lines report against the modeled subset of those two
corpora, matching the coverage the specification currently claims. The other five
channels are complete.

## Building

Build with inlining disabled:

```
go build -gcflags 'all=-l' ./...
```

This flag is required, not cosmetic. The generated `Chat` package defeats the Go
compiler's inliner, which balloons past 60 GB of resident memory and is
OOM-killed. Disabling inlining compiles the same package in about three seconds
and leaves the rest of the optimizer on. Consumers importing this module as a
library need the flag only if they hit the same failure; it is scoped per build,
not baked into the module.

## Usage

```go
import (
    m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
    m_Canvas "github.com/joshmouch/ahp-go/Canvas"
    _dafny "github.com/joshmouch/ahp-go/dafny"
)

state := m_Canvas.Companion_CanvasState_.Create_CanvasState_(
    _dafny.UnicodeSeqOfUtf8Bytes("canvas-1"),
    _dafny.UnicodeSeqOfUtf8Bytes("provider-1"),
    m_AhpSkeleton.Companion_Option_.Create_None_(),
    m_AhpSkeleton.Companion_Option_.Create_None_(),
    m_AhpSkeleton.Companion_Option_.Create_None_(),
    m_AhpSkeleton.Companion_Option_.Create_None_(),
    m_Canvas.Companion_CanvasAvailability_.Create_Ready_(),
)

action := m_Canvas.Companion_CanvasAction_.Create_Updated_(
    m_AhpSkeleton.Companion_Option_.Create_Some_(
        _dafny.UnicodeSeqOfUtf8Bytes("Design Review")),
    m_AhpSkeleton.Companion_Option_.Create_None_(),
    m_AhpSkeleton.Companion_Option_.Create_None_(),
    m_AhpSkeleton.Companion_Option_.Create_None_(),
)

next := (*m_Canvas.Companion_Default___.
    ApplyToCanvas(state, action, _dafny.IntOf(1)).IndexInt(0)).(m_Canvas.CanvasState)
```

The API is generated, so it reads like Dafny rather than like idiomatic Go:
values are immutable, strings are `_dafny.Sequence` of code points, and reducers
return tuples. An idiomatic wrapper is a separate concern from this package,
which deliberately publishes the extracted core unaltered.

See `ahp_smoke_test.go` for worked examples of the reducers, the JSON codec and
command identity.

## What carries a proof, and what does not

Clocks, process execution, the filesystem and sockets sit behind Dafny
`{:extern}` declarations: the specification stops there and the host takes over.
Three packages in this module are therefore hand-written Go rather than
translated output — `ConfluxCommandIdentityCapability`, `ConfluxIoCapability` and
`ConfluxHttpCapability`. They carry no proof and should be reviewed as ordinary
code. `EXTERN.md` names them precisely, along with the single one-line fix
applied to generated output to work around a Dafny Go-backend bug.

## Regenerating

```
./regenerate.sh
```

Re-translates from the Dafny core, reapplies both required post-translation
fixes, rebuilds and runs the tests. Point it at a different core with `AHP_CORE`.

## License

MIT. Copyright (c) Microsoft Corporation; copyright (c) 2026 Josh Mouch. The
protocol and original client are upstream MIT; the Dafny verification and this
port are the second copyright.
