// AHP Dafny client — protocol-version registry, proven as a VERIFIED strict total order.
//
// Upstream source of truth: types/version/registry.ts in the AHP fork.
//   PROTOCOL_VERSION              = '0.5.2'
//   SUPPORTED_PROTOCOL_VERSIONS   = ['0.5.2', '0.5.1']   (most-preferred first)
//   compareProtocolVersions(a,b)  = (aMajor-bMajor) || (aMinor-bMinor) || (aPatch-bPatch)
//
// The upstream test (registry.test.ts) guards, at RUNTIME via string regex:
//   - PROTOCOL_VERSION is MAJOR.MINOR.PATCH
//   - SUPPORTED non-empty; every entry MAJOR.MINOR.PATCH
//   - SUPPORTED[0] == PROTOCOL_VERSION
//   - SUPPORTED strictly descending by SemVer
//   - SUPPORTED has no duplicates
//
// This module discharges the SAME guarantees DEDUCTIVELY. The key upgrades over
// the TS runtime checks:
//
//   * MAJOR.MINOR.PATCH well-formedness is a LANGUAGE GUARANTEE here, not a
//     regex: `SemVer` carries three `nat` components by construction, so every
//     value is well-formed. String-regex parsing therefore need not be modeled.
//   * "no duplicates" is not a separate check — it is a THEOREM derived, for ANY
//     strictly-descending sequence, from transitivity + antisymmetry of the order
//     (StrictlyDescendingImpliesDistinct). The constant list inherits it for free.
//
// Standalone spec: NOT included by client_main.dfy (so it does not enter the
// reducer trust-audit / proof-robustness lanes), gated separately in check-all.sh
// step 1 via `dafny verify spec/version.dfy`.

module Version {
  export
    provides compareSemVer, StrictlyDescending
    provides CompareReflexiveZero, CompareZeroIffEqual, CompareAntisymmetric
    provides CompareTransitive, CompareTotal
    provides StrictlyDescendingImpliesDistinct, UpstreamInvariants
    provides PROTOCOL_VERSION, SUPPORTED
    reveals SemVer


  // ─── SemVer — well-formed BY CONSTRUCTION ──────────────────────────────────
  // The MAJOR.MINOR.PATCH shape the upstream regex `^(\d+)\.(\d+)\.(\d+)$`
  // enforces at runtime is, here, guaranteed by the type: a SemVer is exactly
  // three natural numbers. There is no ill-formed inhabitant to reject.
  datatype SemVer = SemVer(major: nat, minor: nat, patch: nat)

  // ─── compareSemVer — lexicographic; sign matches upstream ───────────────────
  // Mirrors `(aMajor-bMajor) || (aMinor-bMinor) || (aPatch-bPatch)`: the first
  // nonzero component difference wins, and the returned value's SIGN is the
  // comparison result (negative ⇒ a<b, zero ⇒ a==b, positive ⇒ a>b).
  function compareSemVer(a: SemVer, b: SemVer): int
  {
    if a.major != b.major then a.major - b.major
    else if a.minor != b.minor then a.minor - b.minor
    else a.patch - b.patch
  }

  // ─── The deductive payoff: compareSemVer is a STRICT TOTAL ORDER ────────────
  // Each lemma below is universally quantified over arbitrary SemVer values
  // (Dafny's parametric-lemma idiom: proving it for arbitrary `a`/`b`/`c` IS the
  // `forall`), and non-vacuous (no `requires false`; every precondition is
  // satisfiable — e.g. compareSemVer(SemVer(1,0,0), SemVer(0,0,0)) > 0).

  // Reflexive at zero: a compared with itself is 0.  (forall a)
  lemma CompareReflexiveZero(a: SemVer)
    ensures compareSemVer(a, a) == 0
  {}

  // Zero iff equal: the == case of the order coincides with datatype equality.
  // (forall a, b)  — this is what makes "no duplicates" reducible to the order.
  lemma CompareZeroIffEqual(a: SemVer, b: SemVer)
    ensures compareSemVer(a, b) == 0 <==> a == b
  {}

  // Antisymmetry: a > b  iff  b < a.  (forall a, b)
  lemma CompareAntisymmetric(a: SemVer, b: SemVer)
    ensures compareSemVer(a, b) > 0 <==> compareSemVer(b, a) < 0
  {}

  // Transitivity of the strict order: a > b and b > c  ⇒  a > c.  (forall a, b, c)
  // Genuine lexicographic transitivity across all three components — this is the
  // load-bearing lemma the distinctness theorem consumes.
  lemma CompareTransitive(a: SemVer, b: SemVer, c: SemVer)
    requires compareSemVer(a, b) > 0
    requires compareSemVer(b, c) > 0
    ensures compareSemVer(a, c) > 0
  {}

  // Totality / trichotomy: for any a, b EXACTLY ONE of a<b, a==b, a>b holds, with
  // the == branch pinned to datatype equality and the </> branches to each other
  // by antisymmetry — i.e. a genuine total order, not merely "an int is </=/> 0".
  // (forall a, b)
  lemma CompareTotal(a: SemVer, b: SemVer)
    // at least one holds:
    ensures compareSemVer(a, b) < 0 || compareSemVer(a, b) == 0 || compareSemVer(a, b) > 0
    // at most one holds (mutual exclusion is trivial on an int, stated for completeness):
    ensures compareSemVer(a, b) < 0 ==> !(compareSemVer(a, b) == 0 || compareSemVer(a, b) > 0)
    ensures compareSemVer(a, b) > 0 ==> !(compareSemVer(a, b) == 0 || compareSemVer(a, b) < 0)
    // the three branches carry ORDER meaning, not just sign:
    ensures compareSemVer(a, b) == 0 <==> a == b
    ensures compareSemVer(a, b) > 0 <==> compareSemVer(b, a) < 0
  {
    CompareZeroIffEqual(a, b);
    CompareAntisymmetric(a, b);
  }

  // ─── StrictlyDescending + the no-duplicates THEOREM ─────────────────────────

  // A sequence is strictly descending when each element strictly exceeds its
  // successor under the SemVer order (adjacent-pairs form).
  predicate StrictlyDescending(vs: seq<SemVer>)
  {
    forall i | 0 < i < |vs| :: compareSemVer(vs[i-1], vs[i]) > 0
  }

  // Pairwise consequence: adjacent-descending lifts, by transitivity, to
  // "vs[i] strictly exceeds vs[j]" for ANY i < j.  (proven by induction on j-i)
  lemma StrictlyDescendingPairwise(vs: seq<SemVer>, i: int, j: int)
    requires StrictlyDescending(vs)
    requires 0 <= i < j < |vs|
    ensures compareSemVer(vs[i], vs[j]) > 0
    decreases j - i
  {
    if i + 1 == j {
      // adjacent: directly from StrictlyDescending (instantiated at index j).
    } else {
      StrictlyDescendingPairwise(vs, i, j - 1);          // vs[i]   > vs[j-1]
      // StrictlyDescending gives vs[j-1] > vs[j] (index j).
      CompareTransitive(vs[i], vs[j - 1], vs[j]);          // ⇒ vs[i] > vs[j]
    }
  }

  // THE THEOREM: strict descent ⇒ pairwise-distinct, for ANY sequence (not just
  // the 2-element constant list). "No duplicates" is thus a proven consequence of
  // the order's algebra (transitivity + antisymmetry), never a separate test.
  lemma StrictlyDescendingImpliesDistinct(vs: seq<SemVer>)
    requires StrictlyDescending(vs)
    ensures forall i, j | 0 <= i < j < |vs| :: vs[i] != vs[j]
  {
    forall i, j | 0 <= i < j < |vs|
      ensures vs[i] != vs[j]
    {
      StrictlyDescendingPairwise(vs, i, j);   // compareSemVer(vs[i], vs[j]) > 0
      CompareZeroIffEqual(vs[i], vs[j]);       // > 0 ⇒ != 0 ⇒ vs[i] != vs[j]
    }
  }

  // ─── The empirical / constant half — mirrors registry.ts VERBATIM ───────────
  const PROTOCOL_VERSION: SemVer := SemVer(0, 5, 2)
  const SUPPORTED: seq<SemVer> := [SemVer(0, 5, 2), SemVer(0, 5, 1)]

  // Every upstream invariant, proven on the ACTUAL constants:
  //   non-empty · SUPPORTED[0] == PROTOCOL_VERSION · strictly descending ·
  //   distinct (via the general theorem above, so duplicates are impossible).
  lemma UpstreamInvariants()
    ensures |SUPPORTED| > 0
    ensures SUPPORTED[0] == PROTOCOL_VERSION
    ensures StrictlyDescending(SUPPORTED)
    ensures forall i, j | 0 <= i < j < |SUPPORTED| :: SUPPORTED[i] != SUPPORTED[j]
  {
    // StrictlyDescending: the sole adjacent pair is (0.5.2, 0.5.1); majors and
    // minors tie, patch 2 - 1 = 1 > 0, so compareSemVer(...) > 0.
    assert compareSemVer(SUPPORTED[0], SUPPORTED[1]) == 1;
    assert StrictlyDescending(SUPPORTED);
    // Distinctness falls out of the general theorem — no per-element check.
    StrictlyDescendingImpliesDistinct(SUPPORTED);
  }
}
