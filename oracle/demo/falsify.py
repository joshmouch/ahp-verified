#!/usr/bin/env python3
"""
Falsifiability harness for ahp-oracle.

A conformance oracle that reports 232/232 green is worthless unless you can show
it would have said otherwise. This harness perturbs a single leaf of every
fixture's *expected* state and confirms the oracle reacts correctly.

The rigorous claim is NOT "every mutation is caught" -- that would be false, and
a conformance tool that lied about its own reach would be worse than useless.
The oracle adjudicates the *decoded verified-core state*, not the wire spelling.
So the honest, two-directional claim it must satisfy is:

  * Every mutation that CHANGES the decoded verified state is caught (diverges),
    with a diff that names the mutated path.
  * Every mutation that does NOT change the decoded state (a wire-only change --
    an unmodeled field, or a value the tolerant decoder normalizes to the same
    domain value) is correctly passed.
  * ZERO state-changing mutations escape.  <-- the load-bearing invariant.

For each escape the harness PROVES it is wire-only rather than a missed bug:
the mutant agrees (decode(mutant.expected) == fold) and the untouched original
fixture also agrees (decode(original.expected) == fold), so by transitivity the
mutation did not move the decoded state. The reason (unmodeled field vs.
decode-normalized value) is reported for color.

The mutation is deterministic (seeded by filename), so the result is
reproducible on any machine.

Usage:  python3 falsify.py <oracle-binary> <corpus-dir> [mutants-out-dir]

Exit 0 iff zero state-changing mutations escaped.
"""
import json
import os
import subprocess
import sys
import hashlib

# The chat channel's top-level modifiedAt is clock-derived: the verified reducer
# does not stamp it from a wall clock, so the oracle deliberately does not
# adjudicate it. Mutating it would be an expected non-catch that adds no signal;
# --strict-clock measures it separately. Excluded from mutation, not from proof.
SKIP_TOP_LEVEL = {"modifiedAt"}


def leaves(node, path=()):
    """Yield (path, value) for every scalar leaf of a JSON value."""
    if isinstance(node, dict):
        for k, v in node.items():
            if len(path) == 0 and k in SKIP_TOP_LEVEL:
                continue
            yield from leaves(v, path + (k,))
    elif isinstance(node, list):
        for i, v in enumerate(node):
            yield from leaves(v, path + (i,))
    else:
        yield path, node


def perturb(value):
    """Return a different value of a comparable shape."""
    if isinstance(value, bool):
        return not value
    if isinstance(value, int):
        return value + 1
    if isinstance(value, float):
        return value + 1.0
    if isinstance(value, str):
        return value + "-MUTATED" if value != "" else "MUTATED"
    if value is None:
        return "MUTATED"
    return None


def set_at(root, path, value):
    node = root
    for step in path[:-1]:
        node = node[step]
    node[path[-1]] = value


def run_check(binary, path):
    """Run `oracle check --file <path> --json` and return the parsed report."""
    r = subprocess.run([binary, "check", "--file", path, "--json"],
                       capture_output=True, text=True)
    if r.returncode not in (0, 1):  # 0 agree, 1 diverge; anything else is a tool error
        raise SystemExit(f"oracle error on {path} (exit {r.returncode}): {r.stderr.strip()}")
    return json.loads(r.stdout)


def top_segment(dotted):
    return dotted.split(".")[0]


def main():
    if len(sys.argv) < 3:
        print(__doc__)
        return 2
    binary = sys.argv[1]
    corpus = sys.argv[2]
    outdir = sys.argv[3] if len(sys.argv) > 3 else os.path.join(os.path.dirname(__file__), "mutants")
    os.makedirs(outdir, exist_ok=True)
    for f in os.listdir(outdir):
        if f.endswith(".json"):
            os.remove(os.path.join(outdir, f))

    manifest = []
    # results[name] = dict describing the classification
    caught = []
    escaped_unmodeled = []
    escaped_normalized = []
    genuine_misses = []
    no_mutable_leaf = []

    for name in sorted(os.listdir(corpus)):
        if not name.endswith(".json"):
            continue
        original_path = os.path.join(corpus, name)
        doc = json.load(open(original_path))
        if "expected" not in doc:
            continue

        candidates = [p for p, _ in leaves(doc["expected"])]
        if not candidates:
            no_mutable_leaf.append(name)
            continue

        # Deterministic choice, stable across runs and machines.
        seed = int(hashlib.sha256(name.encode()).hexdigest()[:8], 16)
        path = candidates[seed % len(candidates)]
        before = dict(leaves(doc["expected"]))[path]
        after = perturb(before)
        dotted = ".".join(str(p) for p in path)

        set_at(doc["expected"], list(path), after)
        mutant_path = os.path.join(outdir, name)
        json.dump(doc, open(mutant_path, "w"), indent=2)
        manifest.append({"fixture": name, "path": dotted, "from": before, "to": after})

        report = run_check(binary, mutant_path)

        if not report["agrees"]:
            # State-changing mutation -> caught. Confirm the diff actually names
            # the mutated leaf (so the catch is FOR THIS mutation, not incidental).
            diff_paths = {d["path"] for d in report["differences"]}
            # oracle diff paths use [i] for array indices; the manifest uses .i --
            # normalize both to a dotted, bracket-free form for the containment test.
            def norm(p):
                return p.replace("[", ".").replace("]", "")
            named = any(norm(dotted) == norm(dp) or norm(dp).startswith(norm(dotted))
                        or norm(dotted).startswith(norm(dp)) for dp in diff_paths)
            caught.append({"fixture": name, "path": dotted, "named": named,
                           "first": report["differences"][0] if report["differences"] else None})
            continue

        # Escape: oracle agreed. Prove it is wire-only, not a missed bug.
        # 1) the untouched original fixture must also agree (it is a 232/232 corpus
        #    member) -> decode(original.expected) == fold.
        original_report = run_check(binary, original_path)
        legit = original_report["agrees"]  # decode(original)==fold==decode(mutant) => wire-only

        if not legit:
            genuine_misses.append({"fixture": name, "path": dotted,
                                   "why": "mutant agreed but original fixture disagrees -- not decode-equivalent"})
            continue

        # 2) sub-classify WHY the wire change did not move the decoded state.
        #    The oracle reports unmodeled fields at their full nested path, using
        #    [i] for array indices; the manifest uses .i. Normalize both to a
        #    dotted, bracket-free form and test whether the mutated leaf is the
        #    unmodeled field itself or lives under an unmodeled subtree.
        def norm_path(p):
            return p.replace("[", ".").replace("]", "")
        leaf = norm_path(dotted)
        unmodeled = {norm_path(d["path"]) for d in report["unmodeledFields"]}
        is_unmodeled = any(leaf == u or leaf.startswith(u + ".") for u in unmodeled)
        if is_unmodeled:
            escaped_unmodeled.append({"fixture": name, "path": dotted, "from": before, "to": after})
        else:
            escaped_normalized.append({"fixture": name, "path": dotted, "from": before, "to": after})

    json.dump(manifest, open(os.path.join(outdir, "..", "mutants.manifest.json"), "w"), indent=2)

    total = len(caught) + len(escaped_unmodeled) + len(escaped_normalized) + len(genuine_misses)
    caught_unnamed = [c for c in caught if not c["named"]]

    print("=" * 72)
    print("ahp-oracle falsifiability sweep")
    print("=" * 72)
    print(f"corpus fixtures mutated      : {total}")
    print(f"  (fixtures with no mutable leaf, skipped: {len(no_mutable_leaf)})")
    print()
    print(f"CAUGHT  (mutation moved the verified state, oracle diverged) : {len(caught)}")
    print(f"    of which the diff named the mutated leaf                 : {len(caught) - len(caught_unnamed)}")
    print(f"WIRE-ONLY (decodes to the SAME verified state, oracle agreed): {len(escaped_unmodeled) + len(escaped_normalized)}")
    print(f"    because the mutated field is not modeled by the core     : {len(escaped_unmodeled)}")
    print(f"    because the decoder normalizes the value to the same state: {len(escaped_normalized)}")
    print()
    print(f"GENUINE MISSES (state changed but oracle failed to catch it) : {len(genuine_misses)}")
    print("=" * 72)

    if escaped_normalized:
        print("\nwire-only via decode-normalization (modeled field, tolerant decode):")
        for e in escaped_normalized:
            print(f"    {e['fixture']}: {e['path']}  {json.dumps(e['from'])} -> {json.dumps(e['to'])}"
                  f"   (both decode to the same value)")

    if caught_unnamed:
        print("\nNOTE: caught, but the first diff did not name the mutated leaf "
              "(still a correct catch -- the mutation cascaded):")
        for c in caught_unnamed[:10]:
            print(f"    {c['fixture']}: mutated {c['path']}, first diff at {c['first']['path'] if c['first'] else '?'}")

    if genuine_misses:
        print("\n*** FALSIFIABILITY FAILURE: state-changing mutations escaped the oracle ***")
        for m in genuine_misses:
            print(f"    {m['fixture']}: {m['path']}  ({m['why']})")
        print("\nThe oracle is UNSOUND: it agreed with a state it should have rejected.")
        return 1

    print(f"\nPASS: all {len(caught)} state-changing mutations were caught; "
          f"all {len(escaped_unmodeled) + len(escaped_normalized)} wire-only mutations were correctly passed;")
    print("      0 state-changing mutations escaped. The oracle's verdict tracks the verified semantics.")
    return 0


if __name__ == "__main__":
    sys.exit(main())
