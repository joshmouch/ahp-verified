// Diff.cs — structural diff between the oracle's state and a claimed state.
//
// Both sides are compared as parsed ASTs, never as text, so key order and
// whitespace never produce a false difference.
//
// The claimed side is PARSED ONLY. It is deliberately not passed through
// decode/encode: doing so would normalize away exactly the defects this tool
// exists to find.
//
// Copyright (c) Microsoft Corporation.
// Copyright (c) 2026 Josh Mouch.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.Numerics;

namespace AhpOracle
{
    internal enum DiffKind
    {
        ValueMismatch,   // same path, same JSON type, different value
        TypeMismatch,    // same path, different JSON type
        MissingInYours,  // oracle has it, the claimed state does not
        ExtraInYours,    // the claimed state has it, the oracle does not
        LengthMismatch   // arrays of different length
    }

    internal sealed class Difference
    {
        public string Path;
        public DiffKind Kind;
        public string Oracle;
        public string Yours;

        public override string ToString() => Path + ": " + Kind;
    }

    internal static class Diff
    {
        private const int MaxDifferences = 200;

        public static List<Difference> Compare(ConfluxCodec._IJson oracle, ConfluxCodec._IJson yours)
        {
            var outp = new List<Difference>();
            Walk("", oracle, yours, outp);
            return outp;
        }

        private static void Walk(string path, ConfluxCodec._IJson o, ConfluxCodec._IJson y, List<Difference> acc)
        {
            if (acc.Count >= MaxDifferences) return;

            string ot = TypeName(o), yt = TypeName(y);
            if (ot != yt)
            {
                acc.Add(new Difference
                {
                    Path = Show(path),
                    Kind = DiffKind.TypeMismatch,
                    Oracle = Render(o),
                    Yours = Render(y)
                });
                return;
            }

            switch (ot)
            {
                case "null":
                    return;

                case "boolean":
                    if (o.dtor_b != y.dtor_b) Mismatch(path, o, y, acc);
                    return;

                case "number":
                    if (!Bridge.Stringify(o).Equals(Bridge.Stringify(y), StringComparison.Ordinal))
                        Mismatch(path, o, y, acc);
                    return;

                case "string":
                    if (!Bridge.U(o.dtor_s).Equals(Bridge.U(y.dtor_s), StringComparison.Ordinal))
                        Mismatch(path, o, y, acc);
                    return;

                case "array":
                {
                    var oa = Bridge.AsArray(o);
                    var ya = Bridge.AsArray(y);
                    if (oa.Count != ya.Count)
                    {
                        acc.Add(new Difference
                        {
                            Path = Show(path),
                            Kind = DiffKind.LengthMismatch,
                            Oracle = oa.Count + " element(s)",
                            Yours = ya.Count + " element(s)"
                        });
                    }
                    int n = Math.Min(oa.Count, ya.Count);
                    for (int i = 0; i < n; i++)
                        Walk(path + "[" + i + "]", oa[i], ya[i], acc);

                    // Show the elements that exist on only one side, so a dropped
                    // or duplicated list entry is visible and not just a count.
                    for (int i = n; i < oa.Count; i++)
                        acc.Add(new Difference
                        {
                            Path = Show(path + "[" + i + "]"),
                            Kind = DiffKind.MissingInYours,
                            Oracle = Render(oa[i]),
                            Yours = "(absent)"
                        });
                    for (int i = n; i < ya.Count; i++)
                        acc.Add(new Difference
                        {
                            Path = Show(path + "[" + i + "]"),
                            Kind = DiffKind.ExtraInYours,
                            Oracle = "(absent)",
                            Yours = Render(ya[i])
                        });
                    return;
                }

                case "object":
                {
                    var okeys = Bridge.ObjectKeys(o);
                    var ykeys = new HashSet<string>(Bridge.ObjectKeys(y), StringComparer.Ordinal);

                    foreach (var k in okeys)
                    {
                        Bridge.TryField(o, k, out var ov);
                        if (!ykeys.Contains(k))
                        {
                            acc.Add(new Difference
                            {
                                Path = Show(Join(path, k)),
                                Kind = DiffKind.MissingInYours,
                                Oracle = Render(ov),
                                Yours = "(absent)"
                            });
                            continue;
                        }
                        Bridge.TryField(y, k, out var yv);
                        Walk(Join(path, k), ov, yv, acc);
                    }

                    var oset = new HashSet<string>(okeys, StringComparer.Ordinal);
                    foreach (var k in Bridge.ObjectKeys(y))
                    {
                        if (oset.Contains(k)) continue;
                        Bridge.TryField(y, k, out var yv);
                        acc.Add(new Difference
                        {
                            Path = Show(Join(path, k)),
                            Kind = DiffKind.ExtraInYours,
                            Oracle = "(absent)",
                            Yours = Render(yv)
                        });
                    }
                    return;
                }
            }
        }

        private static void Mismatch(string path, ConfluxCodec._IJson o, ConfluxCodec._IJson y, List<Difference> acc) =>
            acc.Add(new Difference
            {
                Path = Show(path),
                Kind = DiffKind.ValueMismatch,
                Oracle = Render(o),
                Yours = Render(y)
            });

        private static string Join(string path, string key) =>
            path.Length == 0 ? key : path + "." + key;

        private static string Show(string path) =>
            path.Length == 0 ? "(root)" : path;

        private static string TypeName(ConfluxCodec._IJson j)
        {
            if (j == null || j.is_JNull) return "null";
            if (j.is_JBool) return "boolean";
            if (j.is_JNum || j.is_JDec) return "number";
            if (j.is_JStr) return "string";
            if (j.is_JArr) return "array";
            if (j.is_JObj) return "object";
            return "unknown";
        }

        /// <summary>One-line rendering, elided when large — a diff line should stay readable.</summary>
        private static string Render(ConfluxCodec._IJson j)
        {
            const int limit = 120;
            string s;
            try { s = Bridge.Stringify(j); }
            catch (Exception) { s = "(unrenderable)"; }
            if (s.Length <= limit) return s;
            return s.Substring(0, limit) + "… (" + s.Length + " chars)";
        }

        public static string KindLabel(DiffKind k) => k switch
        {
            DiffKind.ValueMismatch => "different value",
            DiffKind.TypeMismatch => "different JSON type",
            DiffKind.MissingInYours => "missing from your state",
            DiffKind.ExtraInYours => "not in the oracle's state",
            DiffKind.LengthMismatch => "different array length",
            _ => k.ToString()
        };
    }
}
