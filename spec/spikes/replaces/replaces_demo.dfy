// Proves the `replaceable`/`replaces` module mechanism the plan's mirror `.doo`
// seam relies on: a downstream module overrides a replaceable seam, and the
// override is observed at run time through extracted code.
module MirrorLib {
  replaceable module VersionSeam {
    function ProtocolVersion(): string
    // A proof obligation the replacement MUST satisfy — this is what makes an
    // EMPTY replaces NOT silently pass: the replacement has to discharge it.
    lemma NonEmpty() ensures |ProtocolVersion()| > 0
  }
  function MirrorDefault(): string { "ahp-1.0" }
}

// The OpenAgency extension replaces the seam with its own protocol version.
module OaVersion replaces MirrorLib.VersionSeam {
  function ProtocolVersion(): string { "ahp-oa-1.0" }
  lemma NonEmpty() ensures |ProtocolVersion()| > 0 {}
}

module Driver {
  import MirrorLib
  import MirrorLib.VersionSeam
  method Main() {
    // The override is observed: the seam now returns the OA version, not a default.
    print "seam ProtocolVersion() = ", VersionSeam.ProtocolVersion(), "\n";
    print "mirror default          = ", MirrorLib.MirrorDefault(), "\n";
    expect VersionSeam.ProtocolVersion() == "ahp-oa-1.0";
    print "REPLACES OK — extension override observed across the seam\n";
  }
}
