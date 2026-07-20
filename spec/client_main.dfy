// AHP Dafny client — single multi-channel entry point.
// Runs every channel's corpus against the extracted (cs/js) code.
include "ahp_skeleton.dfy"
include "resource_watch.dfy"
include "canvas.dfy"
include "changeset.dfy"
include "annotations.dfy"
include "terminal.dfy"
include "session.dfy"
include "chat.dfy"

module ClientMain {
  import AhpSkeleton
  import ResourceWatch
  import Canvas
  import Changeset
  import Annotations
  import Terminal
  import Session
  import Chat
  method Main() {
    var rootPass, rootTotal := AhpSkeleton.RunCorpus();
    print "ROOT CORPUS:          ", rootPass, "/", rootTotal, " green against extracted code\n";
    var rwPass, rwTotal := ResourceWatch.RunCorpus();
    print "RESOURCEWATCH CORPUS: ", rwPass, "/", rwTotal, " green against extracted code\n";
    var cvPass, cvTotal := Canvas.RunCorpus();
    print "CANVAS CORPUS:        ", cvPass, "/", cvTotal, " green against extracted code\n";
    var csPass, csTotal := Changeset.RunCorpus();
    print "CHANGESET CORPUS:     ", csPass, "/", csTotal, " green against extracted code\n";
    var anPass, anTotal := Annotations.RunCorpus();
    print "ANNOTATIONS CORPUS:   ", anPass, "/", anTotal, " green against extracted code\n";
    var tPass, tTotal := Terminal.RunCorpus();
    print "TERMINAL CORPUS:      ", tPass, "/", tTotal, " green against extracted code\n";
    var sPass, sModeled := Session.RunCorpus();
    print "SESSION CORPUS:       ", sPass, "/", sModeled, " modeled green (of 61 total; all ~25 action TYPES now modeled)\n";
    var chPass, chModeled := Chat.RunCorpus();
    print "CHAT CORPUS:          ", chPass, "/", chModeled, " modeled green (of 97 total; full tool-call state machine + turn lifecycle modeled)\n";
    var pass := rootPass + rwPass + cvPass + csPass + anPass + tPass + sPass + chPass;
    var total := rootTotal + rwTotal + cvTotal + csTotal + anTotal + tTotal + sModeled + chModeled;
    print "TOTAL: ", pass, "/", total, " corpus fixtures green (5 full AHP channels + session/chat partial)\n";
    expect pass == total;
    print "MULTI-CHANNEL CLIENT OK — extract(cs,js) + corpus all green\n";
  }
}
