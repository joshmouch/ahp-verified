// Bridge.cs — THE ENTIRE TRUSTED BOUNDARY OF THIS TOOL.
//
// Everything the oracle claims about AHP semantics comes from the extracted
// Dafny in Ahp.Core.Verified. This file does not decide anything about the
// protocol; it only moves bytes across the host/Dafny line:
//
//     stdin (string)
//       -> Dafny.Sequence<Rune>                     [this file]
//       -> ConfluxJsonText.ParseJson                [extracted Dafny]
//       -> Ahp.decodeAhpState / decodeAhpAction     [extracted Dafny, round-trip proven]
//       -> Ahp.foldAhp                              [extracted Dafny, host-authority law proven]
//       -> Ahp.encodeAhpState                       [extracted Dafny]
//       -> ConfluxJsonText.Stringify                [extracted Dafny, key-sorted]
//       -> string                                   [this file]
//       -> stdout
//
// Read that pipeline as the audit surface. Four of the six steps are machine-
// checked; the two on the ends are the ~40 lines below. Deliberately kept this
// small so a reviewer can read all of the untrusted-to-trusted glue in one sitting.
//
// Copyright (c) Microsoft Corporation.
// Copyright (c) 2026 Josh Mouch.
// Licensed under the MIT License.

using System;
using System.Numerics;
using System.Collections.Generic;

namespace AhpOracle
{
    /// <summary>
    /// Marshalling between host strings and the extracted Dafny JSON AST.
    /// No protocol semantics live here.
    /// </summary>
    internal static class Bridge
    {
        /// <summary>Host string -> Dafny <c>seq&lt;char&gt;</c>.</summary>
        public static Dafny.ISequence<Dafny.Rune> S(string s) =>
            Dafny.Sequence<Dafny.Rune>.UnicodeFromString(s ?? string.Empty);

        /// <summary>Dafny <c>seq&lt;char&gt;</c> -> host string.</summary>
        public static string U(Dafny.ISequence<Dafny.Rune> s) =>
            s == null ? string.Empty : s.ToVerbatimString(false);

        /// <summary>
        /// Parse JSON text using the core's own parser — not System.Text.Json.
        /// Using the extracted parser matters: it is the same one the round-trip
        /// gate exercises, so the oracle reads a fixture exactly as the proofs
        /// model it, rather than through a second, differently-behaved reader.
        /// </summary>
        /// <returns>false if the text is not well-formed JSON.</returns>
        public static bool TryParse(string text, out ConfluxCodec._IJson json, out string error)
        {
            json = ConfluxCodec.Json.create_JNull();
            error = null;
            try
            {
                var res = ConfluxJsonText.__default.ParseJsonChecked(S(text));
                if (res.is_Invalid)
                {
                    error = "not well-formed JSON (rejected by the core's ConfluxJsonText parser)";
                    return false;
                }
                json = res.dtor_value;
                return true;
            }
            catch (Exception e)
            {
                error = "JSON parse failed: " + e.Message;
                return false;
            }
        }

        /// <summary>
        /// Serialize via the core's own stringifier. Objects are emitted with
        /// <c>SortedKeys</c>, so oracle output is canonical: two runs that agree
        /// semantically agree byte-for-byte.
        /// </summary>
        public static string Stringify(ConfluxCodec._IJson j) =>
            U(ConfluxJsonText.__default.Stringify(j));

        // ---- small read-only accessors over the Dafny AST -------------------
        // Used by fixture handling and the differ. These read the AST; they never
        // construct protocol values.

        public static bool TryField(ConfluxCodec._IJson j, string key, out ConfluxCodec._IJson value)
        {
            value = ConfluxCodec.Json.create_JNull();
            if (j == null || !j.is_JObj) return false;
            var k = S(key);
            if (!j.dtor_fields.Contains(k)) return false;
            value = Dafny.Map<Dafny.ISequence<Dafny.Rune>, ConfluxCodec._IJson>.Select(j.dtor_fields, k);
            return true;
        }

        public static bool HasField(ConfluxCodec._IJson j, string key) =>
            TryField(j, key, out _);

        public static string AsString(ConfluxCodec._IJson j) =>
            (j != null && j.is_JStr) ? U(j.dtor_s) : null;

        public static IReadOnlyList<ConfluxCodec._IJson> AsArray(ConfluxCodec._IJson j)
        {
            var outp = new List<ConfluxCodec._IJson>();
            if (j == null || !j.is_JArr) return outp;
            var n = (int)j.dtor_elems.LongCount;
            for (int i = 0; i < n; i++) outp.Add(j.dtor_elems.Select(new BigInteger(i)));
            return outp;
        }

        public static IReadOnlyList<string> ObjectKeys(ConfluxCodec._IJson j)
        {
            var outp = new List<string>();
            if (j == null || !j.is_JObj) return outp;
            foreach (var k in j.dtor_fields.Keys.Elements) outp.Add(U(k));
            outp.Sort(StringComparer.Ordinal);
            return outp;
        }

        /// <summary>Build a one-key object. Used only to place a channel state
        /// under its channel key before handing it to the proven decoder.</summary>
        public static ConfluxCodec._IJson Obj1(string key, ConfluxCodec._IJson value)
        {
            var m = Dafny.Map<Dafny.ISequence<Dafny.Rune>, ConfluxCodec._IJson>.FromElements(
                new Dafny.Pair<Dafny.ISequence<Dafny.Rune>, ConfluxCodec._IJson>(S(key), value));
            return ConfluxCodec.Json.create_JObj(m);
        }

        public static ConfluxCodec._IJson Null() => ConfluxCodec.Json.create_JNull();
    }
}
