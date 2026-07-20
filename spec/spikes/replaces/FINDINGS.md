> [↑ Doc index](../../../../README.md) · agent-host-protocol-workspace

# Spike: `replaces`-across-library mechanism + the empty-`replaces` footgun

Resolves the plan's flagged-UNPROVEN item (review Major 6): the mirror ships a
`replaceable` seam that the OpenAgency extension overrides, and the plan warned
that "an empty `replaces` compiles clean" (a silent footgun) — so the override
mechanism needed a first-order proof before anything relies on it.

## What's proven (`replaces_demo.dfy`, `dafny verify` → 0 errors, run → override observed)

- **The override mechanism works.** `module OaVersion replaces MirrorLib.VersionSeam`
  supplies `ProtocolVersion() == "ahp-oa-1.0"`; a `Driver` importing the seam
  observes the OVERRIDE at runtime (`ahp-oa-1.0`), not the mirror default
  (`ahp-1.0`). Confirmed through extracted C# (`dafny run --target cs`), not just
  at proof time.
- **The footgun is DEFEATED by giving the seam a proof obligation.** The
  `replaceable module VersionSeam` declares `lemma NonEmpty() ensures |ProtocolVersion()| > 0`.
  A replacement MUST discharge every declared member — so an *empty* `replaces`
  (one that omits `ProtocolVersion`/`NonEmpty`) fails resolution instead of
  silently compiling clean. **The mitigation for the silent-empty-`replaces`
  footgun is therefore: attach a non-vacuous proof obligation to every seam the
  extension is expected to fill.** The seam's contract, not a lint, forces a real
  replacement.

## Contract for the real mirror seam (Phase 6/8)

1. Every `replaceable` seam in the generated mirror carries at least one
   `ensures`-bearing lemma the replacement must discharge (so empty overrides
   can't compile).
2. Default to **add-only** overrides; gate any seam whose default is genuinely
   replaced behind an audited allowlist (per the plan's `replaces` discipline).
3. The `.doo`/`--library` packaging (build the mirror once as an immutable
   library, consume with `--library`) is the DEPLOYMENT wrapper around this
   language-level mechanism; the mechanism itself — the load-bearing risk — is
   proven here.

## Re-run

```bash
cd archai/10-arche/ahp-dafny/spec/spikes/replaces
dafny verify replaces_demo.dfy              # 0 errors — the NonEmpty obligation is discharged
dafny run replaces_demo.dfy --target cs --no-verify   # prints ahp-oa-1.0 → REPLACES OK
```
