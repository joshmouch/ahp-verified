// C ABI binding layer over the formally verified AHP core.
//
// Every function here is a thin, hand-written marshalling shell. All protocol
// behaviour — JSON parsing, action decoding, reduction, state encoding — is
// performed by the Dafny-extracted verified core (Ahp.Core.Verified). This file
// contains no protocol logic of its own and must never acquire any: its only
// jobs are UTF-8 marshalling, handle bookkeeping, and turning exceptions into
// C-style error codes.
//
// THREAD SAFETY
//   Every entry point is safe to call concurrently, because calls into the core
//   are SERIALIZED by CoreLock. Two separate reasons require this, and the
//   weaker of the two is the handle table:
//
//   1. ahp_state_apply is a read-modify-write. Taking a lock only to read, and
//      again only to write, lets two threads read the same state, reduce from
//      it independently, and each store its own result — the second overwriting
//      the first. Both calls report success and one action is silently gone.
//      The whole read-reduce-write must be one critical section.
//
//   2. The extracted core's own data structures are not thread-safe, so even
//      two READERS of one state race. Dafny's ConcatSequence memoizes on first
//      read: ImmutableElements assigns the non-volatile `elmts` field and then
//      nulls `left`/`right`. A second thread inside ComputeElements() can
//      observe those nulls. Derived states structurally share subterms, so this
//      is reachable from ordinary use of this ABI.
//
//   Serializing is therefore correctness, not caution. The cost is that
//   independent handles do not proceed in parallel; it is stated in the header.
//
// UTF-8 CONTRACT (shared by all three bindings in this repository)
//   Input strings MUST be well-formed UTF-8. Invalid input is REFUSED, never
//   repaired. Marshal.PtrToStringUTF8 silently substitutes U+FFFD for malformed
//   bytes, which merges DISTINCT inputs into ONE verified state — after which
//   the core's proven equality correctly reports two different identifiers as
//   equal. So this layer decodes strictly and fails the call instead.
//
// Copyright (c) Microsoft Corporation.
// Copyright (c) 2026 Josh Mouch.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading;

using Str = Dafny.ISequence<Dafny.Rune>;

internal static class AhpNative
{
    // ---------------------------------------------------------------------
    // Error reporting (thread-local, C-style)
    // ---------------------------------------------------------------------

    [ThreadStatic] private static IntPtr _lastError;

    private static void ClearError()
    {
        if (_lastError != IntPtr.Zero)
        {
            Marshal.FreeCoTaskMem(_lastError);
            _lastError = IntPtr.Zero;
        }
    }

    private static void SetError(string message)
    {
        ClearError();
        _lastError = Marshal.StringToCoTaskMemUTF8(message ?? "unknown error");
    }

    // ---------------------------------------------------------------------
    // Marshalling helpers
    // ---------------------------------------------------------------------

    private static Str S(string s) => Dafny.Sequence<Dafny.Rune>.UnicodeFromString(s);

    private static string P(Str s) => s.ToVerbatimString(false);

    private static IntPtr Out(string s) => Marshal.StringToCoTaskMemUTF8(s);

    /// <summary>
    /// Decoder that THROWS on malformed UTF-8 rather than substituting U+FFFD.
    /// See the UTF-8 CONTRACT at the top of this file for why that matters.
    /// </summary>
    private static readonly UTF8Encoding StrictUtf8 =
        new UTF8Encoding(encoderShouldEmitUTF8Identifier: false, throwOnInvalidBytes: true);

    /// <summary>
    /// Reads a required UTF-8 C string. Throws if null, or if the bytes are not
    /// well-formed UTF-8 — this binding refuses malformed text, it does not
    /// repair it.
    /// </summary>
    private static unsafe string In(IntPtr p, string name)
    {
        if (p == IntPtr.Zero) throw new ArgumentNullException(name, name + " must not be NULL");
        var bytes = MemoryMarshal.CreateReadOnlySpanFromNullTerminated((byte*)p);
        try
        {
            return StrictUtf8.GetString(bytes);
        }
        catch (DecoderFallbackException)
        {
            throw new FormatException(name + " is not well-formed UTF-8");
        }
    }

    /// <summary>
    /// Parses JSON text through the verified parser, failing loudly on invalid
    /// input. The core's ParseJson silently yields JNull for malformed text;
    /// ParseJsonChecked exposes the Invalid case, so we use that instead.
    /// </summary>
    private static ConfluxCodec._IJson ParseStrict(string text, string what)
    {
        var result = ConfluxJsonText.__default.ParseJsonChecked(S(text));
        if (!result.is_Parsed)
        {
            throw new FormatException(what + " is not well-formed JSON");
        }
        return result.dtor_value;
    }

    private static string Render(ConfluxCodec._IJson j) =>
        P(ConfluxJsonText.__default.Stringify(j));

    // ---------------------------------------------------------------------
    // Handle table for the streaming API
    //
    // CoreLock guards this table AND every call into the verified core. See
    // THREAD SAFETY at the top of this file: the core's own sequences are not
    // safe to read concurrently, so it is not enough to guard the table alone.
    // Every entry point below holds this lock for the whole of its work.
    // ---------------------------------------------------------------------

    private static readonly Dictionary<long, Ahp._IAhpState> _states = new();
    private static readonly object CoreLock = new();
    private static long _nextHandle;

    /// <summary>Registers a state and returns its handle. Caller holds CoreLock.</summary>
    private static long Store(Ahp._IAhpState state)
    {
        var h = ++_nextHandle;
        _states[h] = state;
        return h;
    }

    /// <summary>Resolves a handle. Caller holds CoreLock.</summary>
    private static Ahp._IAhpState Load(long handle)
    {
        if (!_states.TryGetValue(handle, out var s))
        {
            throw new ArgumentException("invalid state handle: " + handle);
        }
        return s;
    }

    // =====================================================================
    // Exported C ABI
    // =====================================================================

    [UnmanagedCallersOnly(EntryPoint = "ahp_version")]
    public static IntPtr Version()
    {
        // Static literal; caller must NOT free. Cached for pointer stability.
        return VersionHolder.Ptr;
    }

    private static class VersionHolder
    {
        internal static readonly IntPtr Ptr =
            Marshal.StringToCoTaskMemUTF8("ahp-core-verified 0.1.0 (nativeaot/osx-arm64)");
    }

    [UnmanagedCallersOnly(EntryPoint = "ahp_last_error")]
    public static IntPtr LastError() => _lastError;

    [UnmanagedCallersOnly(EntryPoint = "ahp_string_free")]
    public static void StringFree(IntPtr p)
    {
        if (p != IntPtr.Zero) Marshal.FreeCoTaskMem(p);
    }

    // --------------------------- pure JSON API ---------------------------

    /// <summary>Default (initial) AHP state, encoded as JSON.</summary>
    [UnmanagedCallersOnly(EntryPoint = "ahp_initial_state_json")]
    public static IntPtr InitialStateJson()
    {
        try
        {
            ClearError();
            lock (CoreLock)
            {
                return Out(Render(Ahp.__default.encodeAhpState(Ahp.AhpState.Default())));
            }
        }
        catch (Exception ex) { SetError(ex.Message); return IntPtr.Zero; }
    }

    /// <summary>applyAhp: state JSON x action JSON -> new state JSON.</summary>
    [UnmanagedCallersOnly(EntryPoint = "ahp_apply_json")]
    public static IntPtr ApplyJson(IntPtr stateJson, IntPtr actionJson)
    {
        try
        {
            ClearError();
            lock (CoreLock)
            {
                var state = Ahp.__default.decodeAhpState(ParseStrict(In(stateJson, "state_json"), "state_json"));
                var action = Ahp.__default.decodeAhpAction(ParseStrict(In(actionJson, "action_json"), "action_json"));
                return Out(Render(Ahp.__default.encodeAhpState(Ahp.__default.applyAhp(state, action))));
            }
        }
        catch (Exception ex) { SetError(ex.Message); return IntPtr.Zero; }
    }

    /// <summary>foldAhp: state JSON x JSON array of actions -> new state JSON.</summary>
    [UnmanagedCallersOnly(EntryPoint = "ahp_fold_json")]
    public static IntPtr FoldJson(IntPtr stateJson, IntPtr actionsJson)
    {
        try
        {
            ClearError();
            lock (CoreLock)
            {
                var state = Ahp.__default.decodeAhpState(ParseStrict(In(stateJson, "state_json"), "state_json"));
                var arr = ParseStrict(In(actionsJson, "actions_json"), "actions_json");
                if (!arr.is_JArr) throw new FormatException("actions_json must be a JSON array");

                var elems = arr.dtor_elems;
                var actions = new Ahp._IAhpAction[(int)elems.LongCount];
                for (var i = 0; i < actions.Length; i++)
                {
                    actions[i] = Ahp.__default.decodeAhpAction(elems.Select(i));
                }

                var folded = Ahp.__default.foldAhp(
                    state, Dafny.Sequence<Ahp._IAhpAction>.FromArray(actions));
                return Out(Render(Ahp.__default.encodeAhpState(folded)));
            }
        }
        catch (Exception ex) { SetError(ex.Message); return IntPtr.Zero; }
    }

    /// <summary>
    /// Codec round trip: encodeAhpAction(decodeAhpAction(j)). The verified core
    /// proves this is identity on well-formed actions; exposing it lets a C++
    /// caller check that its wire encoding is understood by the core.
    /// </summary>
    [UnmanagedCallersOnly(EntryPoint = "ahp_action_roundtrip_json")]
    public static IntPtr ActionRoundtripJson(IntPtr actionJson)
    {
        try
        {
            ClearError();
            lock (CoreLock)
            {
                var j = ParseStrict(In(actionJson, "action_json"), "action_json");
                return Out(Render(Ahp.__default.encodeAhpAction(Ahp.__default.decodeAhpAction(j))));
            }
        }
        catch (Exception ex) { SetError(ex.Message); return IntPtr.Zero; }
    }

    /// <summary>
    /// Channel an action routes to, per the core's dispatch. Returns
    /// "root/unknown" for action types the core does not recognise — the core
    /// decodes those to the defined RootUnknown catch-all rather than failing,
    /// so this is how a caller detects an unrecognised action.
    /// </summary>
    [UnmanagedCallersOnly(EntryPoint = "ahp_action_channel")]
    public static IntPtr ActionChannel(IntPtr actionJson)
    {
        try
        {
            ClearError();
            string channel;
            lock (CoreLock)
            {
                var a = Ahp.__default.decodeAhpAction(
                    ParseStrict(In(actionJson, "action_json"), "action_json"));

                channel =
                    a.is_ASession ? "session" :
                    a.is_AChat ? "chat" :
                    a.is_ATerminal ? "terminal" :
                    a.is_AChangeset ? "changeset" :
                    a.is_AAnnotations ? "annotations" :
                    a.is_AResourceWatch ? "resourceWatch" :
                    a.is_ACanvas ? "canvas" :
                    a.is_ARoot && a.dtor_rootAction.is_RootUnknown ? "root/unknown" :
                    "root";
            }

            return Out(channel);
        }
        catch (Exception ex) { SetError(ex.Message); return IntPtr.Zero; }
    }

    // ------------------------ handle-based API ---------------------------

    [UnmanagedCallersOnly(EntryPoint = "ahp_state_new")]
    public static long StateNew()
    {
        try
        {
            ClearError();
            lock (CoreLock) { return Store(Ahp.AhpState.Default()); }
        }
        catch (Exception ex) { SetError(ex.Message); return 0; }
    }

    [UnmanagedCallersOnly(EntryPoint = "ahp_state_from_json")]
    public static long StateFromJson(IntPtr stateJson)
    {
        try
        {
            ClearError();
            lock (CoreLock)
            {
                return Store(Ahp.__default.decodeAhpState(
                    ParseStrict(In(stateJson, "state_json"), "state_json")));
            }
        }
        catch (Exception ex) { SetError(ex.Message); return 0; }
    }

    /// <summary>
    /// Applies one action in place. Returns 0 on success, -1 on error.
    ///
    /// The load-reduce-store below is ONE critical section. Splitting it —
    /// locking to read, reducing outside the lock, locking again to write —
    /// loses updates: two threads read the same state, each reduces from it,
    /// and the second store overwrites the first while both calls report
    /// success. Do not narrow this lock.
    /// </summary>
    [UnmanagedCallersOnly(EntryPoint = "ahp_state_apply")]
    public static int StateApply(long handle, IntPtr actionJson)
    {
        try
        {
            ClearError();
            lock (CoreLock)
            {
                var state = Load(handle);
                var action = Ahp.__default.decodeAhpAction(
                    ParseStrict(In(actionJson, "action_json"), "action_json"));
                _states[handle] = Ahp.__default.applyAhp(state, action);
            }
            return 0;
        }
        catch (Exception ex) { SetError(ex.Message); return -1; }
    }

    [UnmanagedCallersOnly(EntryPoint = "ahp_state_to_json")]
    public static IntPtr StateToJson(long handle)
    {
        try
        {
            ClearError();
            lock (CoreLock)
            {
                return Out(Render(Ahp.__default.encodeAhpState(Load(handle))));
            }
        }
        catch (Exception ex) { SetError(ex.Message); return IntPtr.Zero; }
    }

    /// <summary>Structural equality of two states, per the core's own equality.</summary>
    [UnmanagedCallersOnly(EntryPoint = "ahp_state_equals")]
    public static int StateEquals(long a, long b)
    {
        try
        {
            ClearError();
            lock (CoreLock) { return Load(a).Equals(Load(b)) ? 1 : 0; }
        }
        catch (Exception ex) { SetError(ex.Message); return -1; }
    }

    /// <summary>
    /// Releases a state handle. Releasing a handle that was never issued, or
    /// that was already released, is a defined no-op: handles are never reused,
    /// so a stale one cannot name a live state.
    /// </summary>
    [UnmanagedCallersOnly(EntryPoint = "ahp_state_free")]
    public static void StateFree(long handle)
    {
        lock (CoreLock) { _states.Remove(handle); }
    }

    /// <summary>1 when the handle names a live state, 0 otherwise.</summary>
    [UnmanagedCallersOnly(EntryPoint = "ahp_state_valid")]
    public static int StateValid(long handle)
    {
        lock (CoreLock) { return _states.ContainsKey(handle) ? 1 : 0; }
    }
}
