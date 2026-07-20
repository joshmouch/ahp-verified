// AHP Dafny client — terminal channel (5th; the content-part state machine).
// content = seq of parts (unclassified | command). `data` appends to the last
// incomplete-command output / unclassified value, else a new unclassified part.
// commandExecuted appends a command part; commandFinished completes by id.
// Full coverage of the 19 terminal fixtures. claim kept opaque (reducer replaces it).
include "ahp_skeleton.dfy"

module Terminal {
  // Firewall: consumers get the terminal channel's action/state/part types + the reducer
  // surface (ApplyToTerminal/apply1/fold/T0/RunCorpus), NOT the internal seq helpers
  // (appendData/finishCommand), the CL claim literal, or the T_UnknownIsNoOp lemmas.
  // provides Wire/CC/AhpSkeleton because the exported signatures name Wire.Json /
  // CC.Fold / Option+ReduceOutcome (transitive-export rule).
  export
    provides ApplyToTerminal, apply1, T0, RunCorpus
    reveals fold   // channel_laws.TerminalIsChannel asserts Terminal.fold == CC.Fold(apply1,·) in its body
    provides AhpSkeleton, Wire, CC
    reveals TerminalState, TerminalAction, Part

  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import CC = ConfluxContract  // the one fold — `fold` calls CC.Fold (delete-and-reuse, not a bridge)
  import Ops = ConfluxOperators

  datatype Part =
    | Unclassified(value: string)
    | Command(commandId: string, commandLine: string, output: string, timestamp: int,
              isComplete: bool, exitCode: Option<int>, durationMs: Option<int>)

  datatype TerminalState = TerminalState(
    title: string, cwd: Option<string>, cols: Option<int>, rows: Option<int>,
    content: seq<Part>, claim: Option<Wire.Json>, exitCode: Option<int>, supportsCommandDetection: Option<bool>)

  datatype TerminalAction =
    | TCwdChanged(cwd: string)                                     // cwdChanged
    | TTitleChanged(title: string)                                // titleChanged
    | TResized(cols: int, rows: int)                              // resized
    | TExited(code: int)                                          // exited
    | TData(data: string)                                         // data
    | TCleared                                                    // cleared
    | TClaimed(claim: Wire.Json)                                       // claimed
    | TCommandDetectionAvailable                                  // commandDetectionAvailable
    | TCommandExecuted(commandId: string, commandLine: string, timestamp: int)  // commandExecuted
    | TCommandFinished(commandId: string, code: int, durationMs: int)           // commandFinished
    | TInput                                                      // input — no-op
    | TUnknown(raw: Wire.Json)                                         // unknown — no-op

  // data append: last incomplete command → output+=d; last unclassified → value+=d; else new unclassified.
  function appendData(content: seq<Part>, d: string): seq<Part>
  {
    if |content| == 0 then [Unclassified(d)]
    else
      var last := content[|content|-1];
      var init := content[..|content|-1];
      match last
      case Command(cid, cl, out, ts, comp, ec, dm) =>
        if !comp then init + [Command(cid, cl, out + d, ts, comp, ec, dm)] else content + [Unclassified(d)]
      case Unclassified(v) => init + [Unclassified(v + d)]
  }

  function finishCommand(content: seq<Part>, id: string, code: int, dur: int): seq<Part>
  {
    Ops.Map((p: Part) => match p
        case Command(cid, cl, out, ts, comp, ec, dm) =>
          if cid == id then Command(cid, cl, out, ts, true, Some(code), Some(dur)) else p
        case Unclassified(_) => p, content)
  }

  function ApplyToTerminal(s: TerminalState, a: TerminalAction, now: int): (TerminalState, ReduceOutcome)
  {
    match a
    case TCwdChanged(c)      => (s.(cwd := Some(c)), Applied)
    case TTitleChanged(t)    => (s.(title := t), Applied)
    case TResized(co, ro)    => (s.(cols := Some(co), rows := Some(ro)), Applied)
    case TExited(code)       => (s.(exitCode := Some(code)), Applied)
    case TData(d)            => (s.(content := appendData(s.content, d)), Applied)
    case TCleared            => (s.(content := []), Applied)
    case TClaimed(c)         => (s.(claim := Some(c)), Applied)
    case TCommandDetectionAvailable => (s.(supportsCommandDetection := Some(true)), Applied)
    case TCommandExecuted(id, cl, ts) =>
      (s.(content := s.content + [Command(id, cl, "", ts, false, None, None)], supportsCommandDetection := Some(true)), Applied)
    case TCommandFinished(id, code, dur) => (s.(content := finishCommand(s.content, id, code, dur)), Applied)
    case TInput              => (s, NoOp)      // input is a no-op in the reducer
    case TUnknown(_)         => (s, NoOp)
  }

  lemma T_UnknownIsNoOp(s: TerminalState, a: TerminalAction, now: int)
    requires a.TUnknown? || a.TInput?
    ensures ApplyToTerminal(s, a, now).0 == s && ApplyToTerminal(s, a, now).1 == NoOp {}
  lemma T_UnknownIsNoOp_NonVacuous()
    ensures exists a: TerminalAction :: a.TUnknown? || a.TInput?
  { assert (TUnknown(Wire.JNull)).TUnknown?; }

  function apply1(s: TerminalState, a: TerminalAction): TerminalState { ApplyToTerminal(s, a, 9999).0 }
  function fold(s: TerminalState, acts: seq<TerminalAction>): TerminalState
  { CC.Fold(apply1, s, acts) }   // hand-rolled left-fold of apply1 DELETED — now the kernel fold

  function T0(): TerminalState { TerminalState("bash", None, None, None, [], None, None, None) }
  function CL(): Wire.Json { Wire.JObj(map["kind" := Wire.JStr("client"), "clientId" := Wire.JStr("c1")]) }

  method RunCorpus() returns (pass: nat, total: nat) {
    total := 19; pass := 0;
    var base := T0().(claim := Some(CL()));

    // cwdChanged / titleChanged / resized / exited
    if apply1(base, TCwdChanged("/tmp")).cwd == Some("/tmp") { pass := pass+1; }
    if apply1(base, TTitleChanged("zsh")).title == "zsh" { pass := pass+1; }
    if apply1(base, TResized(80, 24)) == base.(cols := Some(80), rows := Some(24)) { pass := pass+1; }
    if apply1(base, TExited(0)).exitCode == Some(0) { pass := pass+1; }
    // input no-op + unknown no-op
    if apply1(base, TInput) == base { pass := pass+1; }
    if apply1(base, TUnknown(Wire.JObj(map["type" := Wire.JStr("terminal/nope")]))) == base { pass := pass+1; }
    // cleared empties (2 fixtures: plain + with command detection)
    if apply1(base.(content := [Unclassified("x")]), TCleared).content == [] { pass := pass+1; }
    if apply1(base.(content := [Unclassified("x")], supportsCommandDetection := Some(true)), TCleared) == base.(content := [], supportsCommandDetection := Some(true)) { pass := pass+1; }
    // claimed (3 variants — reducer just replaces the opaque claim)
    var sc := Wire.JObj(map["kind" := Wire.JStr("session"), "session" := Wire.JStr("s1")]);
    if apply1(base, TClaimed(sc)).claim == Some(sc) { pass := pass+1; }
    var scTool := Wire.JObj(map["kind" := Wire.JStr("session"), "session" := Wire.JStr("s1"), "toolCallId" := Wire.JStr("t1")]);
    if apply1(base, TClaimed(scTool)).claim == Some(scTool) { pass := pass+1; }
    if apply1(base.(claim := Some(sc)), TClaimed(CL())).claim == Some(CL()) { pass := pass+1; }
    // commandDetectionAvailable
    if apply1(base, TCommandDetectionAvailable).supportsCommandDetection == Some(true) { pass := pass+1; }
    // commandExecuted appends command part + sets detection
    var afterExec := apply1(base, TCommandExecuted("cmd-1", "npm test", 1700000000000));
    if afterExec == base.(content := [Command("cmd-1", "npm test", "", 1700000000000, false, None, None)], supportsCommandDetection := Some(true)) { pass := pass+1; }
    // commandFinished completes by id
    var withCmd := base.(content := [Command("cmd-1", "npm test", "All tests passed\r\n", 1700000000000, false, None, None)], supportsCommandDetection := Some(true));
    var doneCmd := base.(content := [Command("cmd-1", "npm test", "All tests passed\r\n", 1700000000000, true, Some(0), Some(1234))], supportsCommandDetection := Some(true));
    if apply1(withCmd, TCommandFinished("cmd-1", 0, 1234)) == doneCmd { pass := pass+1; }
    // data appends to unclassified content
    if apply1(base.(content := [Unclassified("a")]), TData("b")).content == [Unclassified("ab")] { pass := pass+1; }
    // data after completed command → new unclassified
    var complete := base.(content := [Command("cmd-1", "echo hi", "hi\r\n", 1700000000000, true, Some(0), Some(50))], supportsCommandDetection := Some(true));
    if apply1(complete, TData("$ ")).content == complete.content + [Unclassified("$ ")] { pass := pass+1; }
    // data appends to incomplete command output
    if apply1(withCmd, TData("!")).content == [Command("cmd-1", "npm test", "All tests passed\r\n!", 1700000000000, false, None, None)] { pass := pass+1; }
    // full TERMINAL lifecycle (real fixture): cwd→data→resize→claim→exit
    var life := fold(base, [TCwdChanged("/w"), TData("x"), TResized(100, 40), TClaimed(sc), TExited(1)]);
    if life.cwd == Some("/w") && life.cols == Some(100) && life.exitCode == Some(1) && life.claim == Some(sc) { pass := pass+1; }
    // 123 full COMMAND lifecycle: data→execute→data→finish→data→execute→data→finish
    var lc := fold(base.(content := []), [
      TData("$ "), TCommandExecuted("cmd-1", "echo hello", 1700000000000), TData("hello\r\n"),
      TCommandFinished("cmd-1", 0, 10), TData("$ "), TCommandExecuted("cmd-2", "exit 1", 1700000001000),
      TData(""), TCommandFinished("cmd-2", 1, 5)]);
    var expLc := [
      Unclassified("$ "),
      Command("cmd-1", "echo hello", "hello\r\n", 1700000000000, true, Some(0), Some(10)),
      Unclassified("$ "),
      Command("cmd-2", "exit 1", "", 1700000001000, true, Some(1), Some(5))];
    if lc.content == expLc && lc.supportsCommandDetection == Some(true) { pass := pass+1; }
  }
}
