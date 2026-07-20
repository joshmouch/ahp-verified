"""Smoke test for the agent-host-protocol package.

Exercises the extracted reducers directly - constructing protocol
datatypes and driving state transitions - rather than only running the
embedded fixture corpus.

Run with:  python3 smoke_test.py
"""

import sys

sys.path.insert(0, "src")

import agent_host_protocol as ahp
from agent_host_protocol import root

Root = root.RootState_RootState
NoneOpt = root.Option_None
SomeOpt = root.Option_Some
Applied = root.ReduceOutcome_Applied
NoOp = root.ReduceOutcome_NoOp
Seq = ahp.dafny.Seq

failures = []


def check(label, condition):
    status = "ok  " if condition else "FAIL"
    print(f"  [{status}] {label}")
    if not condition:
        failures.append(label)


print("1. construct an empty root state")
state = Root(Seq([]), NoneOpt(), NoneOpt(), NoneOpt())
check("state is a RootState", state.is_RootState)
check("activeSessions starts None", state.activeSessions.is_None)

print("2. apply RootActiveSessionsChanged -> Applied, field updated")
action = root.RootAction_RootActiveSessionsChanged(Seq(["s-1", "s-2"]))
next_state, outcome = root.default__.ApplyToRoot(state, action, 1000)
check("outcome is Applied", outcome.is_Applied)
check("activeSessions now Some", next_state.activeSessions.is_Some)
check("carries both sessions", len(next_state.activeSessions.value) == 2)
check("input state unchanged (immutability)", state.activeSessions.is_None)

print("3. unknown action is a no-op, state preserved")
unknown = root.RootAction_RootUnknown(Seq("some.unrecognized.method"))
same_state, outcome2 = root.default__.ApplyToRoot(next_state, unknown, 1001)
check("outcome is NoOp", outcome2.is_NoOp)
check("state identical", same_state == next_state)
check("IsUnknownRoot detects it", root.default__.IsUnknownRoot(unknown))

print("4. config precondition: change on absent config is rejected")
cfg_action = root.RootAction_RootConfigChanged(ahp.dafny.Map({}), False)
check(
    "precondition reports failure",
    root.default__.ConfigPreconditionFails(state, cfg_action),
)
_, outcome3 = root.default__.ApplyToRoot(state, cfg_action, 1002)
check("reducer returns NoOp", outcome3.is_NoOp)

print("5. FoldRoot folds a sequence of actions")
folded = root.default__.FoldRoot(
    state,
    Seq([
        root.RootAction_RootActiveSessionsChanged(Seq(["a"])),
        root.RootAction_RootActiveSessionsChanged(Seq(["a", "b", "c"])),
    ]),
    2000,
)
check("fold applied last write", len(folded.activeSessions.value) == 3)

print("6. embedded fixture corpus")
results = ahp.run_corpus()
total_passed = sum(p for p, _ in results.values())
total_all = sum(t for _, t in results.values())
for name, (passed, total) in results.items():
    check(f"{name}: {passed}/{total}", passed == total)
print(f"     corpus total: {total_passed}/{total_all}")

print()
if failures:
    print(f"SMOKE TEST FAILED - {len(failures)} check(s) failed:")
    for f in failures:
        print(f"  - {f}")
    sys.exit(1)
print("SMOKE TEST PASSED - extracted reducers are live")
