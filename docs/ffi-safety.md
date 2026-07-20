# Known defects in the C ABI bindings

**Scope of this document: disclosure.** These defects are reproduced and unfixed at
the time of writing. They are published here rather than left to be discovered
because the packages table previously described the C++ route as "Shipping", and a
reader deciding whether to link this library needs to know what is actually known
about it.

## The boundary this is about

The proofs cover the **reducers**: given a state and an action, what the next state
is. They say nothing about the FFI layer wrapped around them.

That layer — the C ABI, the handle tables, the string marshalling, the refcounting —
is **hand-written and unverified**. It is where every defect below lives. The
distinction matters: none of these are counterexamples to the proofs, and none of
them are reachable from the pure-language packages (.NET, JS, Python, Go, JVM used
directly). They are reachable from C or C++ code linking one of the `cpp-via-*`
routes.

## Status by route

| Route | Backend | Status |
|---|---|---|
| `cpp-via-go` | Go c-archive | Exercised in CI. Defects S3–S5 below. |
| `cpp-via-dotnet` | .NET NativeAOT | Not in CI. Defect S1 below. Best error model of the three. |
| `cpp-via-dafny-next` | Rust | **Not in CI. Not safe under concurrency — see S2.** |

`cpp-via-dafny-next` is the route the documentation previously recommended. On the
evidence below, it is the one to avoid.

## Reproduced defects

Each was reproduced with a standalone C driver.

### S2 — Rust: undefined behaviour when handles cross threads (most severe)

Dafny's Rust backend represents shared values with `Rc`, a **non-atomic** refcount.
The binding hands those handles across the C ABI, where nothing stops a caller
putting them on two threads. Concurrent refcount updates on `Rc` are undefined
behaviour, not merely racy.

A two-thread driver crashed **18 of 20 runs**, in three distinct ways:

```
crashed 18/20
  Trace/BPT trap: 5
  Bus error: 10
  Segmentation fault: 11
```

This is not a tuning problem and cannot be fixed by callers being careful — the
type is wrong for a boundary that permits multithreaded access. Do not use this
route from multithreaded code.

### S7 — Rust: use-after-free

Releasing a handle and then using it segfaults rather than failing cleanly
(`exit=139`). No liveness check on the handle path.

### S6 — Rust: `capacity overflow` on a fold count

A fold length argument reaches an allocation path unvalidated and panics.

### S3 — Go: two distinct canvas ids compare equal (a correctness defect)

```
canvasId A bytes: 63 ff
canvasId B bytes: 63 fe
ahp_canvas_equals(A,B) = 1     <-- distinct ids reported equal
```

This one is worth separating from the rest. It is not a memory-safety issue — it is
a **wrong answer** from an exported comparison function, and a caller has no way to
detect it.

### S4 — Go: invalid or misused handles crash inside the bridge

Rather than returning an error, misuse faults inside `cpp-via-go/bridge/ahp_bridge.go`
(`:244` in `ahp_canvas_id`, `:304` in `ahp_canvas_release`).

### S5 — Go: `gobytes: length out of range` and a nil dereference in the SHA path

```
panic: runtime error: gobytes: length out of range
panic: runtime error: invalid memory address or nil pointer dereference
```

### S1 — .NET: concurrent applies silently lose updates

Two threads, 300 applies each. Sequentially this yields 600 parts:

```
sequential 2x300 applies -> 600 parts
concurrent 2x300 applies -> 367 parts
LOST UPDATES: 233 updates silently lost
```

No crash, no error, no diagnostic — the caller is simply given a wrong answer. The
handle table is not safe for concurrent mutation.

### UTF-8 handling diverges between bindings

Given the same invalid UTF-8 input, the bindings disagree:

- **Rust** rejects it (returns NULL).
- **Go** round-trips it, silently substituting replacement characters:
  `ef bf bd ef bf bd 61` (`��a`).

Neither is documented as the contract. Two bindings of the same proven core give
different observable answers for the same input, which defeats the property the
extraction exists to provide.

## What this means for the claims elsewhere

- The C++ row in the packages table now reads **prototype, not recommended for
  production**, and names this document.
- "148/148 green" for C++ remains true and is **single-threaded**. None of the
  defects above are exercised by that run.
- These defects do not affect the pure-language packages used directly.

## If you are using a C ABI route today

- Do not share handles across threads on any route.
- Prefer `cpp-via-go` (in CI) or `cpp-via-dotnet` (best error model) over
  `cpp-via-dafny-next`.
- Do not rely on `ahp_canvas_equals` (S3).
- Validate UTF-8 before it reaches the boundary; the bindings do not agree.

Reproduction drivers and captured output are retained under `_work2/ffi-safety/`.
