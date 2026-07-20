#!/usr/bin/env python3
"""
Make the conflux runtime translatable by Dafny's Rust backend.

The Rust backend mandates --enforce-determinism, which forbids the
assign-such-that STATEMENT (`var k :| P(k);`) unconditionally -- even when the
witness is provably unique. It does NOT forbid the let-such-that EXPRESSION
inside a compiled function.

ConfluxJsonText.SortedKeySetRuntime contains the only compiled assign-such-that
in the whole runtime (the other 25 occurrences are in ghost lemmas/functions,
which are erased before compilation and therefore unaffected).

This rewrite lifts that one statement into a compiled function `LeastKeyOf`
carrying the IDENTICAL predicate. The method's postcondition
(`ensures keys == SortedKeySet(s)`) and every downstream proof step are
unchanged, and Dafny re-verifies the result.

Copyright (c) Microsoft Corporation
Copyright (c) 2026 Josh Mouch
SPDX-License-Identifier: MIT
"""
import sys

path = sys.argv[1]
src = open(path).read()

OLD = """  method SortedKeySetRuntime(s: set<string>) returns (keys: seq<string>)
    ensures keys == SortedKeySet(s)
    decreases |s|
  {
    if s == {} {
      keys := [];
    } else {
      LexMinIsLeast(s);
      var k :| k in s && forall y :: y in s ==> !StrLt(y, k);
      LeastKeyUnique(s, k, LexMin(s));
      var tail := SortedKeySetRuntime(s - {k});
      keys := [k] + tail;
    }
  }"""

NEW = """  // Compiled least-key selection. Carries exactly the predicate that the
  // assign-such-that statement below used to assert. Expressed as a function so
  // that --enforce-determinism (mandatory for the Rust backend) accepts it: the
  // flag bans the assign-such-that STATEMENT form, not the let-such-that
  // EXPRESSION form.
  //
  // Dafny additionally requires a COMPILED let-such-that to be uniquely
  // determined. It is: LeastKeyUnique already proves that any two least keys of
  // s are equal (via StrLtTrichotomy). The assert-by below discharges that
  // obligation, so the selected string is forced to be LexMin(s) -- the same
  // value the statement form yielded.
  function LeastKeyOf(s: set<string>): string
    requires s != {}
    requires exists k :: k in s && (forall y :: y in s ==> !StrLt(y, k))
    ensures LeastKeyOf(s) in s
    ensures forall y :: y in s ==> !StrLt(y, LeastKeyOf(s))
  {
    assert forall a, b ::
        (a in s && (forall y :: y in s ==> !StrLt(y, a))) &&
        (b in s && (forall y :: y in s ==> !StrLt(y, b))) ==> a == b
    by {
      forall a, b |
          (a in s && (forall y :: y in s ==> !StrLt(y, a))) &&
          (b in s && (forall y :: y in s ==> !StrLt(y, b)))
        ensures a == b
      {
        LeastKeyUnique(s, a, b);
      }
    }
    var k :| k in s && (forall y :: y in s ==> !StrLt(y, k));
    k
  }

  method SortedKeySetRuntime(s: set<string>) returns (keys: seq<string>)
    ensures keys == SortedKeySet(s)
    decreases |s|
  {
    if s == {} {
      keys := [];
    } else {
      LexMinIsLeast(s);
      assert LexMin(s) in s && (forall y :: y in s ==> !StrLt(y, LexMin(s)));
      var k := LeastKeyOf(s);
      LeastKeyUnique(s, k, LexMin(s));
      var tail := SortedKeySetRuntime(s - {k});
      keys := [k] + tail;
    }
  }"""

if OLD not in src:
    sys.exit("FAIL: SortedKeySetRuntime pattern not found -- runtime changed; re-check this patch.")

open(path, "w").write(src.replace(OLD, NEW, 1))
print("patched: ConfluxJsonText.SortedKeySetRuntime -> LeastKeyOf (1 of 1 compiled assign-such-that)")
