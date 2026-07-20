# Fixture corpus — provenance

`corpus/reducers/` holds 232 reducer fixtures. They are **upstream fixtures, vendored
into this repository unmodified**. Nothing here was authored for this project.

They are vendored rather than fetched because the per-channel denominators that
`gen/run-reducer-replay.sh` pins — and every "237/237" and "232/232" figure in this
repository's documentation — are only checkable if the corpus they refer to is fixed.
An unpinned corpus makes those numbers unfalsifiable: they would drift silently
whenever upstream added a fixture.

## Source

| | |
|---|---|
| Upstream project | [microsoft/agent-host-protocol](https://github.com/microsoft/agent-host-protocol) |
| Path in upstream | `types/test-cases/reducers/` |
| Snapshot taken from | branch `pr/dotnet-client` @ `0985176fddd0d845374394fab4465356d34c7faf` |
| That branch's base | upstream `main` @ `ea8c6b9445a5d4b4a084a7db99ad5e9fde7cbd50` |
| Upstream git tree hash | `08ed98689a442ca33e297b5dfb7829a31719daff` |
| Licence | MIT, Copyright (c) Microsoft Corporation — see `LICENSE` at the repository root |

`pr/dotnet-client` is a working branch in the author's **fork** of the upstream
repository. It is used here only as the snapshot point; it is strictly ahead of the
`main` it is based on and adds no fixtures of its own. Every one of the 232 files is
upstream-authored — verified by `git log` attribution across the whole set:

```
$ for f in corpus/reducers/*.json; do
    git log -1 --format='%ae' pr/dotnet-client -- "types/test-cases/reducers/$(basename $f)"
  done | sort | uniq -c | sort -rn
 113 connor@peet.io
  64 justchen@microsoft.com
  26 sasomava@microsoft.com
  10 roblourens@gmail.com
  10 besimmonds@microsoft.com
   6 3372902+lszomoru@users.noreply.github.com
   2 colbylwilliams@gmail.com
   1 don.jayamanne@outlook.com
```

Zero fixtures are authored by this project. "232 upstream fixtures" is meant
literally.

## Composition

The replay denominators are exactly this split:

| reducer | fixtures |
|---|---:|
| chat | 115 |
| session | 63 |
| terminal | 19 |
| changeset | 16 |
| annotations | 10 |
| root | 7 |
| resourceWatch | 2 |
| **total** | **232** |

The replay gate reports **237**: these 232 plus the 5 canvas fixtures at
`spec/replay/fixtures/canvas-*.json`. Canvas is not an upstream channel and its
fixtures are authored by this project — which is why they are counted separately
everywhere rather than folded into the upstream number.

## Verifying this snapshot

Integrity, with no network access and no upstream checkout:

```bash
cd corpus/reducers && shasum -a 256 -c ../SHA256SUMS
```

`corpus/SHA256SUMS` itself hashes to:

```
c3f5f6908ae78ab1b63dd80a485893815cad2c1a27569a2734745505e93d425d
```

Against a real upstream checkout, that these files are byte-identical to upstream's:

```bash
cd /path/to/agent-host-protocol
for f in /path/to/ahp-verified/corpus/reducers/*.json; do
  b=$(basename "$f")
  a=$(git show "0985176f:types/test-cases/reducers/$b" | shasum -a 256 | cut -d' ' -f1)
  c=$(shasum -a 256 "$f" | cut -d' ' -f1)
  [ "$a" = "$c" ] || echo "MISMATCH: $b"
done
```

## Drift

This snapshot is **not** automatically synchronised with upstream, and upstream moves.
At the time of vendoring, upstream `main` had 211 reducer fixtures against this
snapshot's 232 — the snapshot branch was ahead of the `main` ref available locally.
Upstream has continued to add fixtures since.

So: **these numbers describe this snapshot, not upstream's current state.** To check
against live upstream instead, point the gates at a checkout —

```bash
export AHP_REPO=/path/to/agent-host-protocol
bash gen/run-reducer-replay.sh
```

— and expect the denominators to differ. A difference is drift to be reconciled, not
a failure of the extracted code. The generator drift gate (`gen/check-all.sh` step
7/8) is the check that fails on upstream type changes; the fixture corpus has no
equivalent automatic gate, which is a real gap and is recorded as such in
`REPRODUCIBILITY.md`.
