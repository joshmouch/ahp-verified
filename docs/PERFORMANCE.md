# How slow is the verified core?

"Formally verified" invites the question "so how slow is it?" This page answers it with
numbers, honestly — including a section on what the numbers do **not** mean, which is where
most benchmark write-ups go wrong.

Two questions, two benchmarks:

* **Benchmark A — does verification make one language slower than another?**
  The same reducer workload, extracted from *one* Dafny source to C#, JavaScript, Go and
  Python, timed on each runtime. Fair by construction: no hand-written per-language adapter
  sits between the harness and the reducers, so no adapter bias can be measured instead of
  the core.
* **Benchmark B — how slow is the verified core versus the hand-written reducer it replaces?**
  The real upstream fixture corpus driven through the verified core (extracted to JavaScript)
  **and** through Microsoft's own TypeScript reducer, on the same fixtures, in the same
  process.

## Machine and method

| | |
|---|---|
| Machine | Apple M5 Max, 18 cores (6 performance + 12 efficiency), 128 GB, macOS 26.5.2 (25F84), arm64 |
| Toolchains | .NET 10.0.9 · Node v26.5.0 · Go 1.26.5 · CPython 3.12.9 |
| Method | per-leg warm-up before any timing; medians over many reps; spread reported (p25/p75, min, max) |
| **Caveat — shared machine** | The runs were taken with a background load average of **~6–8** (other work on the same box). This is disclosed, not hidden: it is why the C# and JavaScript medians wobble ±16–20% run-to-run (JIT tiering + core scheduling), while Go and Python stay within ~3%. A quiet, dedicated machine would tighten the noisy legs and would likely lower every absolute number. Read these as medians on a *busy* laptop-class machine, not as a spec sheet. |

Raw per-rep JSON, the environment snapshot for every round, and the correctness-gate output
are all under [`raw/`](raw/); the harness that produced them is under [`harness/`](harness/).
Every leg runs a **correctness gate before any timing** — a benchmark of a broken reducer
measures nothing — and prints `N/N fixtures green`.

---

## Benchmark A — cross-language reducer cost

**Workload.** One rep = `RunCorpus()` on all 8 channel modules: 148 embedded reducer checks
(each applies one or more actions to a state and asserts the result), extracted identically
to every target. All 8 channels participate — including the two richest, `session` and `chat`
(here as modeled subsets: 36 of 61 and 54 of 97 checks). This is the workload where every
channel runs the *same* computation in every language.

7 rounds × 300 reps each. The **headline is the median across the 7 round-medians** (each
round-median already absorbs within-round preemption spikes); the "round range" columns are
the min and max of those 7 round-medians, i.e. the honest run-to-run spread on this loaded
machine. Time is for the whole 148-check rep.

| Runtime | median | round-min | round-max | run-to-run spread | vs fastest | per check |
|---|---:|---:|---:|---:|---:|---:|
| **Go** 1.26.5 | **0.380 ms** | 0.370 | 0.385 | ±3.8% | **1.00×** | 2.6 µs |
| **C#** .NET 10.0.9 | **0.686 ms** | 0.660 | 0.795 | ±20.5% | **1.81×** | 4.6 µs |
| **JavaScript** Node 26 | **1.379 ms** | 1.354 | 1.576 | ±16.4% | **3.63×** | 9.3 µs |
| **Python** CPython 3.12 | **2.935 ms** | 2.911 | 3.001 | ±3.1% | **7.72×** | 19.8 µs |

The C#/JS spread is JIT tiering plus core scheduling on a loaded box, not measurement error —
the seven individual round-medians and the per-rep p25/p75/min/max for each leg are in
[`raw/SUMMARY-benchmark-A.txt`](raw/SUMMARY-benchmark-A.txt). (For reference, the fastest
single rep observed for each leg — the least-preempted, closest to intrinsic cost — was
Go 0.356 ms, C# 0.633 ms, JS 1.271 ms, Python 2.755 ms, preserving the same ordering.)

**Per-channel median (ms)** — `chat` dominates every leg, as expected (it is the largest
reducer and gets the most checks):

| lang | root | resourceWatch | canvas | changeset | annotations | terminal | session | chat |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| checks | 7 | 2 | 5 | 15 | 10 | 19 | 36 | 54 |
| cs | 0.019 | 0.005 | 0.017 | 0.047 | 0.021 | 0.051 | 0.099 | 0.449 |
| go | 0.007 | 0.002 | 0.005 | 0.026 | 0.022 | 0.035 | 0.057 | 0.228 |
| js | 0.015 | 0.004 | 0.011 | 0.048 | 0.031 | 0.062 | 0.152 | 1.047 |
| py | 0.044 | 0.012 | 0.040 | 0.143 | 0.134 | 0.207 | 0.354 | 1.967 |

The ordering is unsurprising and has **nothing to do with verification**: Go and C# compile
to native code with value-type structs; JavaScript is JITed but boxes everything; CPython
interprets bytecode with boxed integers. The verified core inherits each runtime's execution
model. See "What these numbers do not mean."

---

## Benchmark B — verified core vs Microsoft's own reducer

**Workload.** The **real** upstream corpus — 232 fixtures vendored byte-identical from
`microsoft/agent-host-protocol` (see [`corpus/PROVENANCE.md`](../../corpus/PROVENANCE.md)) —
driven through two engines in one Node process:

* **upstream reduce** — the shipped Microsoft TypeScript reducer folding actions over state.
  Its state shape *is* the fixture JSON, so it decodes nothing.
* **verified reduce** — the extracted verified `fold` over the *same* actions, operating on
  Dafny immutable datatypes (arbitrary-precision `BigNumber` integers, sequences, maps).
* **verified decode** — JSON → Dafny datatypes. A real cost the verified core imposes on an
  integrator that upstream does not pay, so it is **measured separately** rather than hidden
  inside the reduce number.

`Date.now` is pinned to `9999` for both engines (the value the fixtures were generated
under); engine log chatter is silenced inside the timed region so upstream isn't charged for
I/O the verified core doesn't do. Median over 3000 reps; the ratios reproduced within 0.6%
across repeated runs.

**Coverage — honest boundary.** The verified core's fixture decoders in this benchmark cover
**5 of the 7 channels — root, resourceWatch, terminal, changeset, annotations = 54 of the 232
fixtures.** `chat` (115) and `session` (63) have recursive decoders that are not ported into
this benchmark scaffolding, so they are **not** in this head-to-head. (Upstream can reduce all
232; the verified *reducers* for chat/session are exercised in Benchmark A, just not from the
real JSON here.) So this is a real head-to-head on 54 real fixtures, not a claim over the whole
corpus.

Whole covered set (54 fixtures), per-rep median:

| | time | vs upstream reduce |
|---|---:|---:|
| upstream reduce | **0.0041 ms** | 1× |
| verified reduce | **0.0936 ms** | **22.9×** |
| verified decode | 0.226 ms | — |
| **verified decode + reduce** | **0.320 ms** | **78×** |

* **Reduce vs reduce: the verified fold is ~23× slower** than the hand-written TypeScript
  reducer on identical fixtures. This is the pure engine cost — immutable datatype allocation
  per action and arbitrary-precision integer arithmetic instead of native mutation, exactly
  the price of the machine-checked properties.
* **Integrate vs reduce: ~78×** once the JSON→datatype decode is included, and **decode is
  ~71% of that** — the dominant cost is getting data *into* the verified representation, not
  the reduction itself.

Per-channel reduce ratio (verified/upstream): root 37×, resourceWatch 8×, terminal 20×,
changeset 17×, annotations 45×. Full breakdown in
[`raw/SUMMARY-benchmark-B.txt`](raw/SUMMARY-benchmark-B.txt).

**A correctness finding, not a speed one:** the verified core reproduces every one of the 54
fixtures' expected states. The vendored upstream TypeScript **dist** misses 2 changeset
fixtures (`filesReviewChanged`) — its shipped build predates that action. This is a version
skew between the vendored dist and the newer corpus snapshot, not a defect in either engine's
speed; it is noted because it surfaced from the same gate.

### Absolute scale — the whole real corpus

The hand-written engine folds the **entire 232-fixture real corpus in 0.064 ms** (median;
p25–p75 0.063–0.066; min 0.061, max 0.212). The verified core's covered slice, decode
included, is 0.32 ms. **Both numbers are a rounding error.** The multipliers are large; the
absolute costs are not.

---

## What these numbers do **not** mean

1. **They do not mean the verified core is "slow."** Even at 78×, reducing the covered real
   corpus takes **~0.32 ms**, and the fastest engine does the *whole* 232-fixture corpus in
   **64 µs**. A protocol reducer processes one small message at a time; the per-message cost
   is sub-microsecond to a few microseconds in every language measured. Verification buys its
   guarantees at a **constant factor on an operation that is already essentially free** — not
   at a cost any user or agent would perceive.

2. **Benchmark A is not a language shootout.** It measures *the same extracted computation* on
   four runtimes. Go winning and Python trailing reflects native-compiled-value-types vs
   interpreted-boxed-integers — a property of the runtimes, not of Dafny, verification, or the
   protocol. Do not read "the verified core is 7× slower in Python" as a verification cost; it
   is CPython's cost, and it applies to any Python code shaped like this.

3. **A and B are different workloads and are not comparable to each other.** A is 148
   hand-embedded checks (session/chat as *modeled subsets*). B is the real 232-file corpus but
   only 5 of 7 channels on the verified side. Neither is "the whole real corpus through the
   verified core in every language" — that combination was not run. Don't add, divide, or
   cross-reference the two tables' numbers.

4. **The 23× / 78× ratios are JavaScript-vs-JavaScript** (verified core extracted to JS vs
   upstream native TS, both on Node 26). They are **not** a claim about C# or Go. Benchmark A
   shows the same verified reducers run ~2× (C#) to ~3.6× (Go) faster than on JS, so a
   compiled-runtime head-to-head would very likely show a smaller multiplier — but that run
   was not done, so no number is claimed for it.

5. **The decode gap overstates upstream's real-world advantage.** Upstream "decodes nothing"
   only because its state shape *is* the wire JSON and the fixtures arrive pre-parsed. A real
   upstream integration still has to `JSON.parse` and validate incoming bytes — work not
   counted here. The cleaner engine comparison is **reduce-vs-reduce (23×)**; the 78× is the
   worst case for the verified core (it pays full decode; upstream pays zero).

6. **These are medians on a busy shared machine.** Background load was ~6–8; C#/JS medians
   move ±16–20% between rounds. Treat the leading two significant figures as real and the last
   as noise. This is disclosed so the numbers are trusted, not so they look better.

7. **This measures the reducer core only** — the part that was proven. Not JSON parsing from
   bytes, not transport, not I/O, not a full client. The verified state-transition function is
   what carries the proofs and what is timed here.

### Where the cost comes from

As predicted going in, and confirmed by the per-channel and decode/reduce split: immutable
Dafny datatype allocation on every action (no in-place mutation), arbitrary-precision
`BigNumber` integer arithmetic instead of native machine ints, and sequence/map operations
that copy rather than mutate. Those are precisely the mechanisms that make the proofs
possible — totality, structural equality, no silent overflow — so the cost and the guarantee
are the same thing seen from two sides.

## Reproducing

See [`README.md`](README.md) for exact commands. In short: `harness/run_all.sh` drives
Benchmark A (all four legs, sequential); `harness/h2h/bench_h2h.mjs` and
`harness/h2h/bench_upstream_full.mjs` drive Benchmark B; `harness/analyze.py` and
`harness/analyze_h2h.py` aggregate `raw/*.json` into the tables above.
