# ahp-oracle — a conformance oracle for the Agent Host Protocol

**Point this at your AHP client and it tells you where you diverge from the
proven semantics.**

`ahp-oracle` folds a sequence of AHP actions through the *formally verified*
reducers — the same machine-checked Dafny that the AHP core is extracted from —
and reports the canonical resulting state. Its `check` mode compares your
client's claimed state against that canonical state and prints a precise,
per-field diff of every real disagreement.

There is no second implementation of the reducers here. The tool does exactly
two things the core does not: it moves JSON in and out of the extracted core,
and it diffs two states. Every state transition it reports comes from
`Ahp.foldAhp` in [`spec/ahp.dfy`](../../spec/ahp.dfy), which is verified (0
errors) and whose action round-trip and channel-isolation laws are proven.

If your client disagrees with this tool, one of you is wrong about the protocol,
and it is not the machine-checked side.

---

## Why you'd want this

Microsoft ships five hand-written AHP clients and there is no cross-client
oracle. Each client re-derives the reducer semantics — turn lifecycle, tool-call
state machine, changeset status transitions, MCP server upsert/keying, annotation
re-anchoring — from prose and examples. Two clients that both "pass their own
tests" can still disagree on what a stream of actions means.

This tool is the missing referee. It is not another client's opinion; it is the
fold through the proven reducers. Wire it into your client's test suite and it
will fail the moment your state stops matching the verified semantics — pointing
at the exact field.

---

## Get it

The oracle is a thin host over the extracted core in
[`../../cs/src/Ahp.Core.Verified`](../../cs/src/Ahp.Core.Verified). Build a
standalone, single-file binary (no .NET runtime required to *run* it):

```bash
./build.sh                 # host RID; also self-tests against the corpus
# or a specific target:
./build.sh linux-x64       # cross-publish (linux-arm64, win-x64, osx-x64, osx-arm64, …)
```

`build.sh` produces `dist/<rid>/ahp-oracle` and, when building for the host,
runs the whole 232-fixture corpus and the falsifiability sweep as a gate.
Observed on this machine:

```
    binary:  33M  Mach-O 64-bit executable arm64
==> 3/4 corpus self-test (must AGREE 232/232)
  AGREES — 232/232 fixtures match the proven reducers.
==> 4/4 falsifiability sweep (must catch every state-changing mutation)
  PASS: all 200 state-changing mutations were caught; …0 state-changing mutations escaped.
```

You can also run it framework-dependent without publishing:

```bash
dotnet run --project src/AhpOracle -- corpus ../../corpus/reducers
```

---

## Point it at YOUR client

The whole workflow is: replay a fixture through your client, dump the final
state as JSON, and hand both to the oracle.

```bash
# 1. Your client replays the actions in a fixture and prints its final
#    state for that channel as JSON.
my-ahp-client replay fixture.json > mine.json

# 2. The oracle folds the same fixture through the proven reducers and diffs.
ahp-oracle check --file fixture.json --expected mine.json
```

If you agree:

```
AGREES — your state matches the proven reducers.
  channel : session
  actions : 1 folded
  compared: mine.json
```

If you don't, every difference is named by path:

```
DIVERGES — 1 difference(s) from the proven reducers.
  channel : session
  actions : 1 folded

  status
      different value
      oracle: 8
      yours:  1

  The oracle's state is authoritative here: it is the fold of your actions
  through the machine-checked reducers.
```

The exit code is the contract, so you can gate CI on it:

```bash
if ! ahp-oracle check --file fixture.json --expected mine.json --quiet; then
  echo "client diverges from the verified AHP semantics"; exit 1
fi
```

For machine consumption, `--json` emits the verdict, every difference, and the
authoritative `oracleState`:

```json
{
  "agrees": false,
  "channel": "session",
  "actions": 1,
  "differenceCount": 1,
  "differences": [
    { "path": "status", "kind": "ValueMismatch", "oracle": "8", "yours": "1" }
  ],
  "clockDerivedDifferences": [],
  "unmodeledFields": [],
  "oracleState": { "status": 8, "title": "Test Session", "lifecycle": "ready", … }
}
```

Run your entire fixture directory in one shot:

```bash
ahp-oracle corpus ./my-fixtures --json | jq '.byChannel'
```

---

## Input formats

The oracle accepts the AHP fixture shape verbatim, so you can point it straight
at the shared corpus in [`../../corpus/reducers`](../../corpus/reducers).

**One channel** (the corpus fixture shape, unchanged):

```json
{ "reducer": "chat", "initial": {…}, "actions": [ … ], "expected": {…} }
```

`"expected"` is optional for `fold`, and is what `check` compares against unless
you pass `--expected`.

**All eight channels at once:**

```json
{ "state": { "root": {…}, "chat": {…}, "session": {…}, … }, "actions": [ … ] }
```

The eight channel names are `root`, `session`, `chat`, `terminal`, `changeset`,
`annotations`, `resourceWatch`, `canvas` (also spelled `resource-watch`).
`ahp-oracle channels` prints them.

---

## Commands

| command | what it does |
|---|---|
| `fold` | Print the canonical state produced by folding the actions through the proven reducers. |
| `check` | Compare a claimed state against the oracle's and report every difference by path. **This is the one to wire into your test suite.** |
| `corpus <dir>` | Run every `*.json` fixture in a directory and summarize by channel. |
| `route` | Show which channel each action routes to, and flag any that match no channel prefix (these fold as root no-ops). |
| `channels` | List the eight channel names. |

Flags: `--file F` / stdin, `--expected F`, `--json`, `--quiet`, `--pretty`,
`--strict-clock`, `--channel C` (corpus filter). `ahp-oracle --help` has the full
reference.

**Exit codes** (stable, part of the contract):

| code | meaning |
|---|---|
| `0` | agrees |
| `1` | diverges |
| `2` | usage error |
| `3` | malformed input |

`route` returns `1` if any action matched no channel prefix — the usual reason a
client and the oracle disagree is that an action isn't being recognized at all.

---

## What it adjudicates — and what it deliberately does not

A conformance oracle is only trustworthy if it is honest about the boundary of
its authority. This one adjudicates **decoded verified-core state**, not wire
spelling, and it says so at every turn.

**Unmodeled fields.** The verified core models a specific surface. A field your
client carries that the core does not model (say, a per-turn `duration`, or
customization `children`) is one the oracle *cannot* adjudicate — so it neither
claims nor checks agreement on it. It reports these under `unmodeledFields`
rather than silently ignoring them, on `AGREES` as well as `DIVERGES`. An
unqualified "agrees" would overstate what was checked.

**The one clock carve-out.** The chat channel's top-level `modifiedAt` is stamped
from a wall clock by the upstream reducer; the verified reducer models it as
opaque and does not thread it. The oracle reports a `modifiedAt` difference but
does not count it toward the verdict — and tells you it did so, with the reason,
every time:

```
  clock   : 1 field(s) differ but are clock-derived, so not counted:
            modifiedAt  oracle: "1970-01-01T00:00:01.000Z"  yours: "1970-01-01T00:00:09.999Z"
            The verified reducer treats this as opaque and does not stamp it
            from a clock. Re-run with --strict-clock to count it as a difference.
```

Pass `--strict-clock` to count it. This is the *only* such carve-out. Note what
is deliberately **not** in it: the session channel also has a `modifiedAt`, but
there it is real threaded state, so it is compared for real. Normalizing it would
hide bugs.

---

## Falsifiability — proof it would say otherwise

A tool that reports `232/232` green is worthless unless you can show it would
have said no. Two demos do that.

**One corrupted value.** [`demo/single-corruption.sh`](demo/single-corruption.sh)
takes a real fixture, claims its own verified state (AGREES), then flips exactly
one modeled field and shows the oracle catch it with a diff pointing at that
field:

```
$ ./demo/single-corruption.sh
# honest client  -> AGREES  (exit 0)
# status 8 -> 1   -> DIVERGES, "status  oracle: 8  yours: 1"  (exit 1)
```

**The systematic sweep.** [`demo/falsify.py`](demo/falsify.py) perturbs a single
leaf of *every* fixture's expected state and checks the oracle's reaction. The
claim it proves is two-directional and honest — not "every mutation is caught"
(that would be a lie, and the oracle adjudicates *state*, not bytes), but:

- every mutation that **changes the decoded verified state** is caught, with a
  diff that names the mutated leaf;
- every mutation that is **wire-only** (an unmodeled field, or a value the
  tolerant decoder normalizes to the same state) is correctly passed — and is
  *proven* wire-only, not a missed bug, because the untouched fixture and the
  mutant both fold to the same verified state;
- **zero** state-changing mutations escape.

```
CAUGHT  (mutation moved the verified state, oracle diverged) : 200
    of which the diff named the mutated leaf                 : 200
WIRE-ONLY (decodes to the SAME verified state, oracle agreed): 30
    because the mutated field is not modeled by the core     : 29
    because the decoder normalizes the value to the same state: 1
GENUINE MISSES (state changed but oracle failed to catch it) : 0
```

The single decode-normalized case is instructive: fixture 049 sets
`queuedMessages` to `null`; the mutation changes it to `"MUTATED"`; the verified
chat decoder maps *both* to the empty list. Same state, so the oracle correctly
agrees. That is the point — it checks meaning, not spelling.

The mutation is seeded by filename, so every number above is reproducible on any
machine: `python3 demo/falsify.py <binary> ../../corpus/reducers`.

---

## How it's built, and why C#

The oracle must fold through the **unified, proven reducer** `Ahp.foldAhp`
(`spec/ahp.dfy`) — the entry point whose correctness laws are stated and
verified. That module ships in exactly one of the released extractions:

- The **C#** extraction is translated from `spec/core_lib.dfy`, whose closure
  includes the `Ahp` module. So `decodeAhpState`, `decodeAhpAction`, `foldAhp`,
  and `encodeAhpState` are all present. This host is ~1,300 lines over a
  ~56,000-line extracted core; of those, the actual trusted semantic boundary is
  the ~40 lines of `Bridge.cs` (string ⇄ the core's `Json` datatype). The rest is
  argument parsing, the structural diff, and reporting — nothing that decides
  protocol semantics.
- The **Go** extraction — despite being the more convenient single-static-binary
  distribution target — is translated from `spec/client_main.dfy`, whose import
  closure is the eight *per-channel* modules and `ClientMain`, and **does not
  include `spec/ahp.dfy`**. The Go product therefore exposes the channel reducers
  but no unified `foldAhp`. Hosting the oracle on Go would require either
  regenerating the Go extraction from `core_lib.dfy` (needs the read-only core,
  the private runtime dependency, and a Dafny→Go build) or hand-writing the
  routing and state assembly in Go — which would be *unproven* glue, defeating
  the whole "thin host over the proven core" property.

So C# is the correct host today: it is the only shipped extraction that exposes
the proven unified fold. Distribution is solved without Go by publishing a
**self-contained single-file** binary (`build.sh` above) — a 33 MB standalone
executable that needs no installed runtime. It is larger than a Go binary would
be; that is the honest price of hosting on the extraction that actually carries
the verified reducer. (To make Go viable, add a `dafny translate go
spec/core_lib.dfy` target to the Go extraction's `regenerate.sh`; that is the one
change that would flip this decision.)

---

## Layout

```
build.sh                    reproducible build + corpus + falsifiability gate
src/AhpOracle/
  Program.cs                CLI: fold / check / corpus / route / channels
  Oracle.cs                 the fold — every transition is extracted Dafny
  Bridge.cs                 JSON <-> the core's Json datatype
  Diff.cs                   structural, path-named diff of two canonical states
demo/
  single-corruption.sh      one corrupted value -> one named diff
  falsify.py                per-leaf mutation sweep with honest accounting
  mutants/                  generated mutants (regenerated by falsify.py)
dist/<rid>/ahp-oracle       self-contained single-file binary
```

---

Copyright (c) Microsoft Corporation. Copyright (c) 2026 Josh Mouch.
Licensed under the MIT License.
