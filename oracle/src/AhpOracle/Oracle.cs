// Oracle.cs — fold a request through the PROVEN reducers.
//
// The only interesting thing this file does is choose which proven entry point
// to call. It never computes a state transition itself.
//
// Copyright (c) Microsoft Corporation.
// Copyright (c) 2026 Josh Mouch.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;

namespace AhpOracle
{
    /// <summary>A parsed oracle request.</summary>
    internal sealed class Request
    {
        /// <summary>Channel name for channel-scoped mode, or null for unified mode.</summary>
        public string Channel;
        /// <summary>The starting state, already shaped as a unified AhpState object.</summary>
        public ConfluxCodec._IJson InitialUnified;
        /// <summary>The action list, in order.</summary>
        public List<ConfluxCodec._IJson> Actions = new List<ConfluxCodec._IJson>();
        /// <summary>Claimed result, if the document carried one. May be null.</summary>
        public ConfluxCodec._IJson Expected;
        public bool HasExpected;
    }

    internal static class Oracle
    {
        /// <summary>
        /// The eight channels of the unified AhpState, as encodeAhpState keys them.
        /// Mirrors spec/ahp.dfy's AhpState product; the oracle rejects anything else
        /// rather than silently folding it into the root no-op branch.
        /// </summary>
        public static readonly string[] Channels =
        {
            "root", "session", "chat", "terminal",
            "changeset", "annotations", "resourceWatch", "canvas"
        };

        /// <summary>Fixture files spell resourceWatch this way; the core tolerates both.</summary>
        private static string NormalizeChannel(string c) =>
            c == "resource-watch" ? "resourceWatch" : c;

        public static bool IsChannel(string c) =>
            Array.IndexOf(Channels, NormalizeChannel(c)) >= 0;

        /// <summary>
        /// Read a request document. Two accepted shapes:
        ///
        ///   channel-scoped (this is the corpus fixture shape, verbatim):
        ///     { "reducer": "chat", "initial": {...}, "actions": [...], "expected": {...} }
        ///
        ///   unified (all eight channels at once):
        ///     { "state": { "root":{...}, "chat":{...}, ... }, "actions": [...] }
        /// </summary>
        public static bool TryReadRequest(ConfluxCodec._IJson doc, out Request req, out string error)
        {
            req = new Request();
            error = null;

            if (doc == null || !doc.is_JObj)
            {
                error = "request must be a JSON object";
                return false;
            }

            bool channelScoped = Bridge.HasField(doc, "reducer");
            bool unified = Bridge.HasField(doc, "state");

            if (channelScoped && unified)
            {
                error = "request has both \"reducer\" and \"state\"; use one shape or the other";
                return false;
            }

            if (channelScoped)
            {
                Bridge.TryField(doc, "reducer", out var chJson);
                var ch = Bridge.AsString(chJson);
                if (ch == null)
                {
                    error = "\"reducer\" must be a string naming a channel";
                    return false;
                }
                ch = NormalizeChannel(ch);
                if (!IsChannel(ch))
                {
                    error = $"unknown channel \"{ch}\"; known channels are: {string.Join(", ", Channels)}";
                    return false;
                }
                req.Channel = ch;

                // "initial" is that one channel's state. Lift it under its channel
                // key so the PROVEN decodeAhpState handles it; the other seven
                // slots decode from absent (JNull) to their own defaults, and
                // channel isolation (proven in spec/ahp.dfy) guarantees the actions
                // cannot reach them.
                ConfluxCodec._IJson initial = Bridge.Null();
                Bridge.TryField(doc, "initial", out initial);
                req.InitialUnified = Bridge.Obj1(ch, initial);
            }
            else if (unified)
            {
                Bridge.TryField(doc, "state", out var st);
                if (!st.is_JObj)
                {
                    error = "\"state\" must be an object keyed by channel";
                    return false;
                }
                foreach (var k in Bridge.ObjectKeys(st))
                {
                    if (!IsChannel(k))
                    {
                        error = $"unknown channel key \"{k}\" in \"state\"; known channels are: {string.Join(", ", Channels)}";
                        return false;
                    }
                }
                req.InitialUnified = st;
            }
            else
            {
                error = "request must have either \"reducer\" (+ \"initial\") for one channel, or \"state\" for all channels";
                return false;
            }

            if (!Bridge.TryField(doc, "actions", out var actions) || !actions.is_JArr)
            {
                error = "request must have an \"actions\" array";
                return false;
            }
            foreach (var a in Bridge.AsArray(actions)) req.Actions.Add(a);

            if (Bridge.TryField(doc, "expected", out var exp))
            {
                req.Expected = exp;
                req.HasExpected = true;
            }

            return true;
        }

        /// <summary>
        /// THE ORACLE. Decode, fold — every step is extracted Dafny.
        /// </summary>
        /// <returns>The resulting state as a domain value.</returns>
        public static Ahp._IAhpState Fold(Request req)
        {
            // decodeAhpState : Json -> AhpState        (spec/ahp.dfy, round-trip proven)
            var state = Ahp.__default.decodeAhpState(req.InitialUnified);

            // decodeAhpAction : Json -> AhpAction      (routes on the "type" prefix)
            var decoded = new Ahp._IAhpAction[req.Actions.Count];
            for (int i = 0; i < req.Actions.Count; i++)
                decoded[i] = Ahp.__default.decodeAhpAction(req.Actions[i]);

            // foldAhp : AhpState -> seq<AhpAction> -> AhpState
            // This is ConfluxContract.Fold over applyAhp — the same fold the
            // host-authority law is stated about in spec/channel_laws.dfy.
            return Ahp.__default.foldAhp(state, Dafny.Sequence<Ahp._IAhpAction>.FromArray(decoded));
        }

        /// <summary>
        /// Decode a claimed state into the domain, using the same proven decoder
        /// the oracle used for the initial state.
        ///
        /// This is what makes the comparison meaningful rather than pedantic. The
        /// encoder writes every optional field explicitly as null; a client (and
        /// the upstream fixtures) will usually just omit them. Those are the same
        /// state. Comparing decoded domain values says so; comparing raw JSON text
        /// would report hundreds of differences that are purely spelling.
        ///
        /// This mirrors the project's own replay harnesses, which assert
        /// `reduced == decoded-expected` rather than comparing wire text.
        /// </summary>
        public static Ahp._IAhpState DecodeClaimed(Request req, ConfluxCodec._IJson claimed)
        {
            var unified = req.Channel == null ? claimed : Bridge.Obj1(req.Channel, claimed);
            return Ahp.__default.decodeAhpState(unified);
        }

        /// <summary>Dafny datatypes override Equals, so this is full structural equality.</summary>
        public static bool SameState(Ahp._IAhpState a, Ahp._IAhpState b) => object.Equals(a, b);

        /// <summary>
        /// Fields that no pure reducer can adjudicate, because their value comes
        /// from a clock rather than from (state, actions).
        ///
        /// Exactly one field qualifies: the chat channel's top-level modifiedAt.
        /// Upstream's reducer stamps it from the wall clock, so the fixtures carry
        /// a frozen fake value; the verified reducer models it as opaque and does
        /// not thread it. This is the same single carve-out the project's own chat
        /// replay harness takes (`reduced.(modifiedAt := "N") == expected.(...)`)
        /// and it is the ONLY one.
        ///
        /// Note what is deliberately NOT here: the session channel also has a
        /// modifiedAt, but there it is real threaded state (chatAdded/chatUpdated
        /// carry it), so it is compared for real. Normalizing it would hide bugs.
        /// </summary>
        public static bool IsClockDerivedPath(string channel, string path)
        {
            if (channel == "chat") return path == "modifiedAt";
            if (channel == null) return path == "chat.modifiedAt";
            return false;
        }

        /// <summary>
        /// Fields present in a claimed state that the verified core does not model.
        ///
        /// This matters more than it looks. The oracle's authority is exactly the
        /// modeled surface: a field the core does not carry is a field the core
        /// cannot adjudicate, so agreement on it is neither claimed nor checked.
        /// Reporting it is the difference between "we agree" and "we agree about
        /// the part I actually looked at".
        ///
        /// Detection is mechanical rather than a hand-maintained list: re-encode
        /// the decoded state and take the keys the round-trip dropped. The encoder
        /// emits every modeled field explicitly (nulls included), so anything left
        /// over in the input is by construction unmodeled.
        /// </summary>
        public static List<Difference> UnmodeledFields(Request req, ConfluxCodec._IJson rawClaimed)
        {
            var reencoded = Render(req, DecodeClaimed(req, rawClaimed));
            var outp = new List<Difference>();
            foreach (var d in Diff.Compare(reencoded, rawClaimed))
                if (d.Kind == DiffKind.ExtraInYours) outp.Add(d);
            return outp;
        }

        /// <summary>
        /// Split a raw diff into differences that reflect reducer behaviour and
        /// differences that only reflect a clock reading. Callers report both;
        /// only the first kind decides the verdict unless --strict-clock is set.
        /// </summary>
        public static void PartitionClock(string channel, List<Difference> all,
                                          out List<Difference> real, out List<Difference> clock)
        {
            real = new List<Difference>();
            clock = new List<Difference>();
            foreach (var d in all)
            {
                if (IsClockDerivedPath(channel, d.Path)) clock.Add(d);
                else real.Add(d);
            }
        }

        /// <summary>
        /// Render a state back to JSON. In channel-scoped mode only that channel's
        /// slot is returned, so output lines up with a fixture's "expected".
        /// </summary>
        public static ConfluxCodec._IJson Render(Request req, Ahp._IAhpState state)
        {
            // encodeAhpState : AhpState -> Json
            var encoded = Ahp.__default.encodeAhpState(state);
            if (req.Channel == null) return encoded;
            Bridge.TryField(encoded, req.Channel, out var slot);
            return slot;
        }

        /// <summary>
        /// Report which channel each action routes to, using the same proven
        /// router the fold uses. Useful when a client and the oracle disagree
        /// because an action is not being recognized at all.
        /// </summary>
        public static string RouteOf(ConfluxCodec._IJson action)
        {
            var a = Ahp.__default.decodeAhpAction(action);
            if (a.is_ARoot) return "root";
            if (a.is_ASession) return "session";
            if (a.is_AChat) return "chat";
            if (a.is_ATerminal) return "terminal";
            if (a.is_AChangeset) return "changeset";
            if (a.is_AAnnotations) return "annotations";
            if (a.is_AResourceWatch) return "resourceWatch";
            if (a.is_ACanvas) return "canvas";
            return "unknown";
        }

        /// <summary>
        /// True when the action's wire "type" matched no known channel prefix and
        /// therefore fell through to the root no-op branch. Distinguishing this
        /// from a genuine root action is what turns "my action did nothing" into a
        /// diagnosable message.
        /// </summary>
        public static bool IsUnroutedFallthrough(ConfluxCodec._IJson action)
        {
            var a = Ahp.__default.decodeAhpAction(action);
            if (!a.is_ARoot) return false;
            var ra = a.dtor_rootAction;
            return ra.is_RootUnknown;
        }

        public static string TypeOf(ConfluxCodec._IJson action)
        {
            if (Bridge.TryField(action, "type", out var t)) return Bridge.AsString(t) ?? "(non-string \"type\")";
            return "(no \"type\" field)";
        }
    }
}
