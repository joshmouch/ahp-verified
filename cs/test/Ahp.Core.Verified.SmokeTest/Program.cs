// Smoke test for the Ahp.Core.Verified package.
//
// Exercises the extracted verified core end to end: JSON -> action decode ->
// reducer application -> state inspection -> action re-encode (round trip),
// plus the multi-action fold path and reducer determinism.
//
// Copyright (c) Microsoft Corporation.
// Copyright (c) 2026 Josh Mouch.
// Licensed under the MIT License.

using System;
using System.Numerics;
using Dafny;

using Str = Dafny.ISequence<Dafny.Rune>;

internal static class Program
{
    private static int _failures;

    private static Str S(string s) => Dafny.Sequence<Dafny.Rune>.UnicodeFromString(s);

    private static string P(Str s) => s.ToVerbatimString(false);

    private static void Check(string label, bool condition, string detail)
    {
        if (condition)
        {
            Console.WriteLine($"  PASS  {label}  ({detail})");
        }
        else
        {
            _failures++;
            Console.WriteLine($"  FAIL  {label}  ({detail})");
        }
    }

    private static ConfluxCodec._IJson JObj(params (string Key, ConfluxCodec._IJson Value)[] fields)
    {
        var pairs = new Dafny.Pair<Str, ConfluxCodec._IJson>[fields.Length];
        for (var i = 0; i < fields.Length; i++)
        {
            pairs[i] = new Dafny.Pair<Str, ConfluxCodec._IJson>(S(fields[i].Key), fields[i].Value);
        }

        return ConfluxCodec.Json.create_JObj(
            Dafny.Map<Str, ConfluxCodec._IJson>.FromElements(pairs));
    }

    private static ConfluxCodec._IJson JStr(string s) => ConfluxCodec.Json.create_JStr(S(s));

    private static ConfluxCodec._IJson JNum(int n) => ConfluxCodec.Json.create_JNum(new BigInteger(n));

    private static int Main()
    {
        Console.WriteLine("Ahp.Core.Verified — smoke test");
        Console.WriteLine("==============================");
        Console.WriteLine();

        // ---------------------------------------------------------------
        // 1. Decode a terminal/resized action from JSON and apply it.
        // ---------------------------------------------------------------
        Console.WriteLine("[1] decode -> applyAhp -> state");

        var resizedJson = JObj(
            ("type", JStr("terminal/resized")),
            ("cols", JNum(120)),
            ("rows", JNum(40)));

        var resizedAction = Ahp.__default.decodeAhpAction(resizedJson);
        Check("routes to the terminal channel", resizedAction.is_ATerminal,
            $"is_ATerminal={resizedAction.is_ATerminal}");

        var state0 = Ahp.AhpState.Default();
        Check("initial terminal cols is None", state0.dtor_terminal.dtor_cols.is_None,
            $"is_None={state0.dtor_terminal.dtor_cols.is_None}");

        var state1 = Ahp.__default.applyAhp(state0, resizedAction);
        var cols = state1.dtor_terminal.dtor_cols;
        var rows = state1.dtor_terminal.dtor_rows;

        Check("cols reduced to Some(120)",
            cols.is_Some && cols.dtor_value == new BigInteger(120),
            $"cols={(cols.is_Some ? cols.dtor_value.ToString() : "None")}");
        Check("rows reduced to Some(40)",
            rows.is_Some && rows.dtor_value == new BigInteger(40),
            $"rows={(rows.is_Some ? rows.dtor_value.ToString() : "None")}");
        Check("initial state is unchanged (reducer is pure)",
            state0.dtor_terminal.dtor_cols.is_None,
            $"state0 cols still None={state0.dtor_terminal.dtor_cols.is_None}");

        Console.WriteLine();

        // ---------------------------------------------------------------
        // 2. Codec round trip: encode(decode(j)) == j
        // ---------------------------------------------------------------
        Console.WriteLine("[2] codec round trip");

        var reEncoded = Ahp.__default.encodeAhpAction(resizedAction);
        Check("terminal/resized round-trips to an equal Json value",
            reEncoded.Equals(resizedJson),
            $"equal={reEncoded.Equals(resizedJson)}");

        var titleJson = JObj(
            ("type", JStr("terminal/titleChanged")),
            ("title", JStr("build")));
        var titleAction = Ahp.__default.decodeAhpAction(titleJson);
        Check("terminal/titleChanged round-trips",
            Ahp.__default.encodeAhpAction(titleAction).Equals(titleJson),
            $"equal={Ahp.__default.encodeAhpAction(titleAction).Equals(titleJson)}");

        Console.WriteLine();

        // ---------------------------------------------------------------
        // 3. foldAhp over a multi-channel action sequence.
        // ---------------------------------------------------------------
        Console.WriteLine("[3] foldAhp over a multi-channel sequence");

        var cwdAction = Ahp.__default.decodeAhpAction(JObj(
            ("type", JStr("terminal/cwdChanged")),
            ("cwd", JStr("/src"))));
        var dataAction = Ahp.__default.decodeAhpAction(JObj(
            ("type", JStr("terminal/data")),
            ("data", JStr("hello"))));

        var actions = Dafny.Sequence<Ahp._IAhpAction>.FromElements(
            resizedAction, titleAction, cwdAction, dataAction);

        var folded = Ahp.__default.foldAhp(Ahp.AhpState.Default(), actions);

        Check("folded title is \"build\"",
            P(folded.dtor_terminal.dtor_title) == "build",
            $"title=\"{P(folded.dtor_terminal.dtor_title)}\"");
        Check("folded cwd is Some(\"/src\")",
            folded.dtor_terminal.dtor_cwd.is_Some && P(folded.dtor_terminal.dtor_cwd.dtor_value) == "/src",
            $"cwd={(folded.dtor_terminal.dtor_cwd.is_Some ? "\"" + P(folded.dtor_terminal.dtor_cwd.dtor_value) + "\"" : "None")}");
        Check("folded cols survived the later actions",
            folded.dtor_terminal.dtor_cols.is_Some && folded.dtor_terminal.dtor_cols.dtor_value == new BigInteger(120),
            $"cols={(folded.dtor_terminal.dtor_cols.is_Some ? folded.dtor_terminal.dtor_cols.dtor_value.ToString() : "None")}");
        Check("folded content received the data part",
            folded.dtor_terminal.dtor_content.Count > 0,
            $"content parts={folded.dtor_terminal.dtor_content.Count}");
        Check("untouched channels stayed at their defaults",
            folded.dtor_annotations.Count == 0,
            $"annotations={folded.dtor_annotations.Count}");

        Console.WriteLine();

        // ---------------------------------------------------------------
        // 4. Determinism: folding the same sequence twice yields equal state.
        // ---------------------------------------------------------------
        Console.WriteLine("[4] reducer determinism");

        var foldedAgain = Ahp.__default.foldAhp(Ahp.AhpState.Default(), actions);
        Check("fold is deterministic on identical input",
            folded.Equals(foldedAgain),
            $"equal={folded.Equals(foldedAgain)}");

        // fold == successive apply1
        var stepwise = Ahp.AhpState.Default();
        stepwise = Ahp.__default.applyAhp(stepwise, resizedAction);
        stepwise = Ahp.__default.applyAhp(stepwise, titleAction);
        stepwise = Ahp.__default.applyAhp(stepwise, cwdAction);
        stepwise = Ahp.__default.applyAhp(stepwise, dataAction);
        Check("foldAhp agrees with successive applyAhp",
            folded.Equals(stepwise),
            $"equal={folded.Equals(stepwise)}");

        Console.WriteLine();

        // ---------------------------------------------------------------
        // 5. Per-channel reducers are reachable directly.
        // ---------------------------------------------------------------
        Console.WriteLine("[5] direct per-channel reducer access");

        var termOnly = Terminal.__default.apply1(
            folded.dtor_terminal,
            Terminal.TerminalAction.create_TExited(new BigInteger(0)));
        Check("Terminal.apply1 records the exit code",
            termOnly.dtor_exitCode.is_Some && termOnly.dtor_exitCode.dtor_value.IsZero,
            $"exitCode={(termOnly.dtor_exitCode.is_Some ? termOnly.dtor_exitCode.dtor_value.ToString() : "None")}");

        var cleared = Terminal.__default.apply1(termOnly, Terminal.TerminalAction.create_TCleared());
        Check("Terminal.apply1 clears content",
            cleared.dtor_content.Count == 0,
            $"content parts={cleared.dtor_content.Count}");

        Console.WriteLine();
        Console.WriteLine("==============================");
        if (_failures == 0)
        {
            Console.WriteLine("SMOKE TEST OK — all assertions passed.");
            return 0;
        }

        Console.WriteLine($"SMOKE TEST FAILED — {_failures} assertion(s) failed.");
        return 1;
    }
}
