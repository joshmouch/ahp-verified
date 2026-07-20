# agent-host-protocol

Formally verified reducer core for the **Agent Host Protocol (AHP)**.

The reducers in this package are not hand-written. They are extracted from a
Dafny specification in which the channel reducers, their convergence
properties and the fold laws over them are machine-checked proofs. The Python
sources here are the output of `dafny translate py` against that
specification, so the behaviour you get is the behaviour that was proven.

Pure standard library — no runtime dependencies.

## Install

```bash
pip install agent-host-protocol
```

## Channels

| Accessor | Channel | Reducer |
| --- | --- | --- |
| `root` | Agent / session / terminal registry and client config | `ApplyToRoot` |
| `session` | Session lifecycle | `ApplyToSession` |
| `chat` | Turn lifecycle and tool-call state machine | `ApplyToChat` |
| `terminal` | Terminal streams | `ApplyToTerminal` |
| `canvas` | Canvas surface | `ApplyToCanvas` |
| `changeset` | File changesets | `ApplyToChangeset` |
| `annotations` | Source annotations | `ApplyToAnnotations` |
| `resource_watch` | Resource watching | `ApplyToResourceWatch` |

Every reducer has the shape:

```
ApplyTo<Channel>(state, action, now) -> (next_state, outcome)
```

`outcome` is a `ReduceOutcome` — `Applied` when the action changed state,
`NoOp` when it was unrecognised or its precondition did not hold. Reducers
are total: an unknown action never raises, it yields `NoOp` and the original
state.

## Usage

```python
import agent_host_protocol as ahp
from agent_host_protocol import root

Seq = ahp.dafny.Seq

state = root.RootState_RootState(
    Seq([]),                 # agents
    root.Option_None(),      # activeSessions
    root.Option_None(),      # terminals
    root.Option_None(),      # config
)

action = root.RootAction_RootActiveSessionsChanged(Seq(["s-1", "s-2"]))
next_state, outcome = root.default__.ApplyToRoot(state, action, 1000)

assert outcome.is_Applied
assert next_state.activeSessions.is_Some
assert state.activeSessions.is_None   # inputs are never mutated
```

Folding a sequence of actions:

```python
final = root.default__.FoldRoot(state, Seq([action_a, action_b]), now)
```

## Self-test

The fixture corpus the Dafny build checks the extraction against ships with
the package:

```python
>>> import agent_host_protocol as ahp
>>> ahp.run_corpus()
{'root': (7, 7), 'session': (36, 36), 'chat': (54, 54), 'terminal': (19, 19),
 'canvas': (5, 5), 'changeset': (15, 15), 'annotations': (10, 10),
 'resource_watch': (2, 2)}
```

Each entry is `(passed, total)`. All eight channels report full green.

`root`, `terminal`, `canvas`, `changeset`, `annotations` and `resource_watch`
are modelled in full. `session` and `chat` are the two large channels; every
action *type* is modelled, and the fixtures covering the modelled surface all
pass. See "Coverage" below.

## Coverage

| Channel | Fixtures green | Notes |
| --- | --- | --- |
| root | 7 / 7 | complete |
| resource_watch | 2 / 2 | complete |
| canvas | 5 / 5 | complete |
| changeset | 15 / 15 | complete |
| annotations | 10 / 10 | complete |
| terminal | 19 / 19 | complete |
| session | 36 / 36 modelled | of 61 total fixtures; all ~25 action types modelled |
| chat | 54 / 54 modelled | of 97 total fixtures; full tool-call state machine and turn lifecycle modelled |

Totals: **148 / 148 modelled fixtures green.**

## Layout

The extracted modules live in the private `agent_host_protocol._core`
subpackage. They are generated with flat, absolute module names, so `_core`
is placed on `sys.path` when the package is imported. Import from
`agent_host_protocol`, not from `_core` directly — the `_core` layout tracks
the extractor and is not a stable interface.

## Related packages

The same verified core is published to other ecosystems:

| Ecosystem | Package |
| --- | --- |
| .NET | `Ahp.Core.Verified` |
| npm | `@open-agency/ahp` |
| Go | `github.com/joshmouch/ahp-go` |
| Maven | `agency.open.ahp:ahp-core` |

## License

MIT. Copyright (c) Microsoft Corporation; Copyright (c) 2026 Josh Mouch.
The protocol and the original client are MIT upstream; the Dafny
verification and this port are by Josh Mouch.
