// Program.cs — ahp-oracle command line.
//
// Copyright (c) Microsoft Corporation.
// Copyright (c) 2026 Josh Mouch.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.IO;
using System.Text;

namespace AhpOracle
{
    internal static class Program
    {
        private const string Version = "0.1.0";

        // Exit codes are part of the contract: a CI job wants to tell "your client
        // disagrees with the proven semantics" apart from "the tool broke".
        private const int ExitAgree = 0;
        private const int ExitDiverge = 1;
        private const int ExitUsage = 2;
        private const int ExitBadInput = 3;

        private static int Main(string[] rawArgs)
        {
            Console.OutputEncoding = new UTF8Encoding(false);

            var args = new List<string>(rawArgs);
            if (args.Count == 0 || args[0] == "-h" || args[0] == "--help" || args[0] == "help")
            {
                Usage(Console.Out);
                return args.Count == 0 ? ExitUsage : ExitAgree;
            }
            if (args[0] == "--version" || args[0] == "version")
            {
                Console.WriteLine("ahp-oracle " + Version);
                return ExitAgree;
            }

            string cmd = args[0];
            args.RemoveAt(0);

            try
            {
                switch (cmd)
                {
                    case "fold": return CmdFold(args);
                    case "check": return CmdCheck(args);
                    case "corpus": return CmdCorpus(args);
                    case "route": return CmdRoute(args);
                    case "channels": return CmdChannels();
                    default:
                        Console.Error.WriteLine("ahp-oracle: unknown command \"" + cmd + "\"");
                        Usage(Console.Error);
                        return ExitUsage;
                }
            }
            catch (UsageException ue)
            {
                Console.Error.WriteLine("ahp-oracle: " + ue.Message);
                return ExitUsage;
            }
        }

        private sealed class UsageException : Exception
        {
            public UsageException(string m) : base(m) { }
        }

        // ---- fold ----------------------------------------------------------

        private static int CmdFold(List<string> args)
        {
            string file = TakeOption(args, "--file");
            bool pretty = TakeFlag(args, "--pretty");
            RejectLeftovers(args);

            string text = ReadInput(file);
            if (!Bridge.TryParse(text, out var doc, out var perr))
            {
                Console.Error.WriteLine("ahp-oracle: " + perr);
                return ExitBadInput;
            }
            if (!Oracle.TryReadRequest(doc, out var req, out var rerr))
            {
                Console.Error.WriteLine("ahp-oracle: " + rerr);
                return ExitBadInput;
            }

            var result = Oracle.Render(req, Oracle.Fold(req));
            string json = Bridge.Stringify(result);
            Console.WriteLine(pretty ? Pretty.Format(json) : json);
            return ExitAgree;
        }

        // ---- check ---------------------------------------------------------

        private static int CmdCheck(List<string> args)
        {
            string file = TakeOption(args, "--file");
            string expectedFile = TakeOption(args, "--expected");
            bool asJson = TakeFlag(args, "--json");
            bool quiet = TakeFlag(args, "--quiet");
            bool strictClock = TakeFlag(args, "--strict-clock");
            RejectLeftovers(args);

            string text = ReadInput(file);
            if (!Bridge.TryParse(text, out var doc, out var perr))
            {
                Console.Error.WriteLine("ahp-oracle: " + perr);
                return ExitBadInput;
            }
            if (!Oracle.TryReadRequest(doc, out var req, out var rerr))
            {
                Console.Error.WriteLine("ahp-oracle: " + rerr);
                return ExitBadInput;
            }

            ConfluxCodec._IJson claimed;
            string claimedSource;
            if (expectedFile != null)
            {
                string etext = File.ReadAllText(expectedFile);
                if (!Bridge.TryParse(etext, out claimed, out var eerr))
                {
                    Console.Error.WriteLine("ahp-oracle: --expected " + expectedFile + ": " + eerr);
                    return ExitBadInput;
                }
                claimedSource = expectedFile;
            }
            else if (req.HasExpected)
            {
                claimed = req.Expected;
                claimedSource = "the \"expected\" field of the input";
            }
            else
            {
                Console.Error.WriteLine(
                    "ahp-oracle: nothing to check against — supply --expected <file>, " +
                    "or include an \"expected\" field in the input document.");
                return ExitUsage;
            }

            // Both sides are decoded to domain values and compared structurally;
            // the diff is then computed on both-sides-encoded canonical JSON, so
            // every line it prints is a real semantic difference.
            var oracleState = Oracle.Fold(req);
            var claimedState = Oracle.DecodeClaimed(req, claimed);

            var oracleJson = Oracle.Render(req, oracleState);
            var claimedJson = Oracle.Render(req, claimedState);

            var raw = Oracle.SameState(oracleState, claimedState)
                ? new List<Difference>()
                : Diff.Compare(oracleJson, claimedJson);

            Oracle.PartitionClock(req.Channel, raw, out var diffs, out var clockDiffs);
            if (strictClock) { diffs = raw; clockDiffs = new List<Difference>(); }

            var unmodeled = Oracle.UnmodeledFields(req, claimed);

            if (asJson) return ReportJson(req, oracleJson, diffs, clockDiffs, unmodeled);
            return ReportText(req, oracleJson, diffs, clockDiffs, unmodeled, quiet, claimedSource);
        }

        private static int ReportText(Request req, ConfluxCodec._IJson oracle,
                                      List<Difference> diffs, List<Difference> clockDiffs,
                                      List<Difference> unmodeled, bool quiet, string claimedSource)
        {
            string scope = req.Channel ?? "all channels";

            if (diffs.Count == 0)
            {
                if (!quiet)
                {
                    Console.WriteLine("AGREES — your state matches the proven reducers.");
                    Console.WriteLine("  channel : " + scope);
                    Console.WriteLine("  actions : " + req.Actions.Count + " folded");
                    Console.WriteLine("  compared: " + claimedSource);
                    EmitClockNote(clockDiffs);
                    EmitUnmodeledNote(unmodeled);
                }
                return ExitAgree;
            }

            Console.WriteLine("DIVERGES — " + diffs.Count + " difference(s) from the proven reducers.");
            Console.WriteLine("  channel : " + scope);
            Console.WriteLine("  actions : " + req.Actions.Count + " folded");
            Console.WriteLine("  compared: " + claimedSource);
            EmitClockNote(clockDiffs);
            EmitUnmodeledNote(unmodeled);
            Console.WriteLine();

            foreach (var d in diffs)
            {
                Console.WriteLine("  " + d.Path);
                Console.WriteLine("      " + Diff.KindLabel(d.Kind));
                Console.WriteLine("      oracle: " + d.Oracle);
                Console.WriteLine("      yours:  " + d.Yours);
                Console.WriteLine();
            }

            // A whole-channel type mismatch usually means the shapes were never
            // comparable — say so rather than let the reader hunt.
            if (diffs.Count == 1 && diffs[0].Path == "(root)" && diffs[0].Kind == DiffKind.TypeMismatch)
            {
                Console.WriteLine("  Note: the two documents differ at the top level. Check that your");
                Console.WriteLine("  client is emitting the state for channel \"" + scope + "\" and not a wrapper around it.");
                Console.WriteLine();
            }

            EmitUnroutedHint(req);

            Console.WriteLine("  The oracle's state is authoritative here: it is the fold of your actions");
            Console.WriteLine("  through the machine-checked reducers. Run `ahp-oracle fold` on the same");
            Console.WriteLine("  input to see it in full.");
            return ExitDiverge;
        }

        /// <summary>
        /// Clock-derived differences are reported but do not decide the verdict.
        /// Stating this every time is deliberate: a silent carve-out in a
        /// conformance tool is indistinguishable from a bug it is hiding.
        /// </summary>
        private static void EmitClockNote(List<Difference> clockDiffs)
        {
            if (clockDiffs == null || clockDiffs.Count == 0) return;
            Console.WriteLine("  clock   : " + clockDiffs.Count + " field(s) differ but are clock-derived, so not counted:");
            foreach (var d in clockDiffs)
                Console.WriteLine("            " + d.Path + "  oracle: " + d.Oracle + "  yours: " + d.Yours);
            Console.WriteLine("            The verified reducer treats this as opaque and does not stamp it");
            Console.WriteLine("            from a clock. Re-run with --strict-clock to count it as a difference.");
        }

        /// <summary>
        /// Report the part of the input the oracle has no opinion about. Printing
        /// this on AGREES as well as on DIVERGES is the point: an unqualified
        /// "agrees" would overstate what was actually checked.
        /// </summary>
        private static void EmitUnmodeledNote(List<Difference> unmodeled)
        {
            if (unmodeled == null || unmodeled.Count == 0) return;
            Console.WriteLine("  unchecked: " + unmodeled.Count +
                              " field(s) in your state are not modeled by the verified core,");
            Console.WriteLine("             so the oracle has no opinion on them either way:");
            int shown = 0;
            foreach (var d in unmodeled)
            {
                if (shown++ == 12) { Console.WriteLine("             … and " + (unmodeled.Count - 12) + " more"); break; }
                Console.WriteLine("             " + d.Path + " = " + d.Yours);
            }
        }

        /// <summary>
        /// If any action failed to route to a channel, that is nearly always the
        /// real cause of a divergence, so it is worth surfacing above the diff noise.
        /// </summary>
        private static void EmitUnroutedHint(Request req)
        {
            var unrouted = new List<string>();
            foreach (var a in req.Actions)
                if (Oracle.IsUnroutedFallthrough(a))
                    unrouted.Add(Oracle.TypeOf(a));

            if (unrouted.Count == 0) return;

            Console.WriteLine("  " + unrouted.Count + " action(s) matched no channel prefix and were folded as");
            Console.WriteLine("  root no-ops. If you expected them to change state, the \"type\" string is");
            Console.WriteLine("  the thing to look at:");
            foreach (var t in unrouted) Console.WriteLine("      " + t);
            Console.WriteLine();
        }

        private static int ReportJson(Request req, ConfluxCodec._IJson oracle,
                                      List<Difference> diffs, List<Difference> clockDiffs,
                                      List<Difference> unmodeled)
        {
            var sb = new StringBuilder();
            sb.Append("{\"agrees\":").Append(diffs.Count == 0 ? "true" : "false");
            sb.Append(",\"channel\":").Append(Json.Str(req.Channel ?? "*"));
            sb.Append(",\"actions\":").Append(req.Actions.Count);
            sb.Append(",\"differenceCount\":").Append(diffs.Count);
            sb.Append(",\"differences\":").Append(DiffArray(diffs));
            sb.Append(",\"clockDerivedDifferences\":").Append(DiffArray(clockDiffs));
            sb.Append(",\"unmodeledFields\":").Append(DiffArray(unmodeled));
            sb.Append(",\"oracleState\":").Append(Bridge.Stringify(oracle));
            sb.Append('}');
            Console.WriteLine(sb.ToString());
            return diffs.Count == 0 ? ExitAgree : ExitDiverge;
        }

        private static string DiffArray(List<Difference> diffs)
        {
            var sb = new StringBuilder("[");
            for (int i = 0; i < (diffs?.Count ?? 0); i++)
            {
                if (i > 0) sb.Append(',');
                var d = diffs[i];
                sb.Append("{\"path\":").Append(Json.Str(d.Path));
                sb.Append(",\"kind\":").Append(Json.Str(d.Kind.ToString()));
                sb.Append(",\"oracle\":").Append(Json.Str(d.Oracle));
                sb.Append(",\"yours\":").Append(Json.Str(d.Yours));
                sb.Append('}');
            }
            return sb.Append(']').ToString();
        }

        // ---- corpus --------------------------------------------------------

        private static int CmdCorpus(List<string> args)
        {
            bool asJson = TakeFlag(args, "--json");
            bool verbose = TakeFlag(args, "--verbose");
            bool strictClock = TakeFlag(args, "--strict-clock");
            string only = TakeOption(args, "--channel");
            if (args.Count < 1) throw new UsageException("corpus needs a directory: ahp-oracle corpus <dir>");
            string dir = args[0];
            args.RemoveAt(0);
            RejectLeftovers(args);

            if (!Directory.Exists(dir))
            {
                Console.Error.WriteLine("ahp-oracle: no such directory: " + dir);
                return ExitBadInput;
            }

            var files = new List<string>(Directory.GetFiles(dir, "*.json"));
            files.Sort(StringComparer.Ordinal);

            int pass = 0, fail = 0, skipped = 0, clockNormalized = 0;
            var perChannel = new SortedDictionary<string, int[]>(StringComparer.Ordinal); // [pass, fail]
            var failures = new List<string>();

            foreach (var f in files)
            {
                string text = File.ReadAllText(f);
                if (!Bridge.TryParse(text, out var doc, out var perr))
                {
                    fail++;
                    failures.Add(Path.GetFileName(f) + ": " + perr);
                    continue;
                }
                if (!Oracle.TryReadRequest(doc, out var req, out var rerr))
                {
                    fail++;
                    failures.Add(Path.GetFileName(f) + ": " + rerr);
                    continue;
                }
                if (only != null && req.Channel != only) { skipped++; continue; }
                if (!req.HasExpected)
                {
                    fail++;
                    failures.Add(Path.GetFileName(f) + ": fixture has no \"expected\" state to check against");
                    continue;
                }

                string ch = req.Channel ?? "*";
                if (!perChannel.ContainsKey(ch)) perChannel[ch] = new int[2];

                var oracleState = Oracle.Fold(req);
                var claimedState = Oracle.DecodeClaimed(req, req.Expected);
                var raw = Oracle.SameState(oracleState, claimedState)
                    ? new List<Difference>()
                    : Diff.Compare(Oracle.Render(req, oracleState), Oracle.Render(req, claimedState));

                Oracle.PartitionClock(req.Channel, raw, out var diffs, out var clockOnly);
                if (strictClock) diffs = raw;
                if (clockOnly.Count > 0 && !strictClock) clockNormalized++;

                if (diffs.Count == 0)
                {
                    pass++; perChannel[ch][0]++;
                    if (verbose) Console.WriteLine("  ok   " + Path.GetFileName(f)
                        + (clockOnly.Count > 0 ? "   (clock-derived modifiedAt not counted)" : ""));
                }
                else
                {
                    fail++; perChannel[ch][1]++;
                    failures.Add(Path.GetFileName(f) + ": " + diffs.Count + " difference(s), first at " +
                                 diffs[0].Path + " (oracle: " + diffs[0].Oracle + " / fixture: " + diffs[0].Yours + ")");
                    if (verbose) Console.WriteLine("  FAIL " + Path.GetFileName(f));
                }
            }

            if (asJson)
            {
                var sb = new StringBuilder();
                sb.Append("{\"pass\":").Append(pass).Append(",\"fail\":").Append(fail);
                sb.Append(",\"skipped\":").Append(skipped);
                sb.Append(",\"clockNormalized\":").Append(clockNormalized);
                sb.Append(",\"total\":").Append(pass + fail);
                sb.Append(",\"byChannel\":{");
                bool first = true;
                foreach (var kv in perChannel)
                {
                    if (!first) sb.Append(',');
                    first = false;
                    sb.Append(Json.Str(kv.Key)).Append(":{\"pass\":").Append(kv.Value[0])
                      .Append(",\"fail\":").Append(kv.Value[1]).Append('}');
                }
                sb.Append("},\"failures\":[");
                for (int i = 0; i < failures.Count; i++)
                {
                    if (i > 0) sb.Append(',');
                    sb.Append(Json.Str(failures[i]));
                }
                sb.Append("]}");
                Console.WriteLine(sb.ToString());
                return fail == 0 ? ExitAgree : ExitDiverge;
            }

            Console.WriteLine();
            foreach (var kv in perChannel)
            {
                int p = kv.Value[0], fl = kv.Value[1];
                Console.WriteLine(string.Format("  {0,-14} {1,4}/{2,-4} {3}",
                    kv.Key, p, p + fl, fl == 0 ? "" : "  <-- " + fl + " DIVERGENT"));
            }
            Console.WriteLine();
            if (skipped > 0) Console.WriteLine("  (" + skipped + " fixture(s) skipped by --channel filter)");
            if (clockNormalized > 0)
                Console.WriteLine("  (" + clockNormalized + " fixture(s) differ only in the clock-derived chat modifiedAt, " +
                                  "which the pure reducer does not stamp; --strict-clock counts it)");

            if (fail == 0)
            {
                Console.WriteLine("  AGREES — " + pass + "/" + pass + " fixtures match the proven reducers.");
                return ExitAgree;
            }

            Console.WriteLine("  DIVERGES — " + fail + " of " + (pass + fail) + " fixtures disagree with the proven reducers:");
            Console.WriteLine();
            foreach (var m in failures) Console.WriteLine("    " + m);
            Console.WriteLine();
            return ExitDiverge;
        }

        // ---- route ---------------------------------------------------------

        private static int CmdRoute(List<string> args)
        {
            string file = TakeOption(args, "--file");
            RejectLeftovers(args);

            string text = ReadInput(file);
            if (!Bridge.TryParse(text, out var doc, out var perr))
            {
                Console.Error.WriteLine("ahp-oracle: " + perr);
                return ExitBadInput;
            }

            // Accept a bare action, an array of actions, or a full request document.
            var actions = new List<ConfluxCodec._IJson>();
            if (doc.is_JArr) actions.AddRange(Bridge.AsArray(doc));
            else if (Bridge.TryField(doc, "actions", out var acts) && acts.is_JArr) actions.AddRange(Bridge.AsArray(acts));
            else actions.Add(doc);

            int unrouted = 0;
            foreach (var a in actions)
            {
                string type = Oracle.TypeOf(a);
                string ch = Oracle.RouteOf(a);
                bool fell = Oracle.IsUnroutedFallthrough(a);
                if (fell) unrouted++;
                Console.WriteLine(string.Format("  {0,-42} -> {1}{2}",
                    type, ch, fell ? "  (no channel prefix matched; folded as a root no-op)" : ""));
            }
            return unrouted == 0 ? ExitAgree : ExitDiverge;
        }

        private static int CmdChannels()
        {
            foreach (var c in Oracle.Channels) Console.WriteLine(c);
            return ExitAgree;
        }

        // ---- helpers -------------------------------------------------------

        private static string ReadInput(string file)
        {
            if (file != null && file != "-") return File.ReadAllText(file);
            return Console.In.ReadToEnd();
        }

        private static string TakeOption(List<string> args, string name)
        {
            int i = args.IndexOf(name);
            if (i < 0) return null;
            if (i + 1 >= args.Count) throw new UsageException(name + " needs a value");
            string v = args[i + 1];
            args.RemoveAt(i + 1);
            args.RemoveAt(i);
            return v;
        }

        private static bool TakeFlag(List<string> args, string name)
        {
            int i = args.IndexOf(name);
            if (i < 0) return false;
            args.RemoveAt(i);
            return true;
        }

        private static void RejectLeftovers(List<string> args)
        {
            foreach (var a in args)
                if (a.StartsWith("-", StringComparison.Ordinal))
                    throw new UsageException("unrecognized option \"" + a + "\"");
        }

        private static void Usage(TextWriter w)
        {
            w.WriteLine(@"ahp-oracle " + Version + @" — conformance oracle for the Agent Host Protocol

Folds actions through the formally verified AHP reducers and tells you where
another implementation disagrees. The reducers, codecs and the dispatch layer
are extracted from a machine-checked Dafny specification; this tool only moves
JSON in and out of them.

USAGE
  ahp-oracle fold    [--file F] [--pretty]
  ahp-oracle check   [--file F] [--expected F] [--json] [--quiet]
  ahp-oracle corpus  <dir> [--channel C] [--json] [--verbose]
  ahp-oracle route   [--file F]
  ahp-oracle channels

INPUT (stdin, or --file)
  One channel — this is the AHP fixture shape, unchanged:
    { ""reducer"": ""chat"", ""initial"": {...}, ""actions"": [...], ""expected"": {...} }

  All eight channels at once:
    { ""state"": { ""root"": {...}, ""chat"": {...}, ... }, ""actions"": [...] }

  ""expected"" is optional for fold, and is what check compares against unless
  you pass --expected.

COMMANDS
  fold      Print the canonical state produced by folding the actions.
  check     Compare a claimed state against the oracle's and report every
            difference by path. This is the one to wire into your test suite.
  corpus    Run every *.json fixture in a directory and summarize by channel.
  route     Show which channel each action routes to, and flag any that match
            no channel prefix (these fold as root no-ops).
  channels  List the eight channel names.

EXIT CODES
  0  agrees
  1  diverges
  2  usage error
  3  malformed input

EXAMPLES
  # What does the proven core say this becomes?
  ahp-oracle fold --file fixture.json --pretty

  # Does my client agree? (my-client emits its final state as JSON)
  my-client replay fixture.json > mine.json
  ahp-oracle check --file fixture.json --expected mine.json

  # Whole corpus, machine readable
  ahp-oracle corpus ./corpus/reducers --json | jq .");
        }
    }

    internal static class Json
    {
        public static string Str(string s)
        {
            if (s == null) return "null";
            var sb = new StringBuilder("\"");
            foreach (char c in s)
            {
                switch (c)
                {
                    case '"': sb.Append("\\\""); break;
                    case '\\': sb.Append("\\\\"); break;
                    case '\n': sb.Append("\\n"); break;
                    case '\r': sb.Append("\\r"); break;
                    case '\t': sb.Append("\\t"); break;
                    default:
                        if (c < 0x20) sb.Append("\\u").Append(((int)c).ToString("x4"));
                        else sb.Append(c);
                        break;
                }
            }
            return sb.Append('"').ToString();
        }
    }

    /// <summary>Indentation only. Never reorders or rewrites values.</summary>
    internal static class Pretty
    {
        public static string Format(string json)
        {
            var sb = new StringBuilder();
            int depth = 0;
            bool inStr = false, esc = false;
            foreach (char c in json)
            {
                if (inStr)
                {
                    sb.Append(c);
                    if (esc) esc = false;
                    else if (c == '\\') esc = true;
                    else if (c == '"') inStr = false;
                    continue;
                }
                switch (c)
                {
                    case '"': inStr = true; sb.Append(c); break;
                    case '{':
                    case '[':
                        sb.Append(c).Append('\n').Append(new string(' ', ++depth * 2));
                        break;
                    case '}':
                    case ']':
                        sb.Append('\n').Append(new string(' ', --depth * 2)).Append(c);
                        break;
                    case ',':
                        sb.Append(c).Append('\n').Append(new string(' ', depth * 2));
                        break;
                    case ':':
                        sb.Append(": ");
                        break;
                    default:
                        sb.Append(c);
                        break;
                }
            }
            return sb.ToString();
        }
    }
}
