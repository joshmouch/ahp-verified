/*******************************************************************************
 * conflux-extern-impl — .NET binding for ConfluxWebSocketCapability.
 *
 * This file is the sole System.Net.WebSockets owner used by Dafny consumers.  Handles are opaque to
 * Dafny.  ConnectWebSocket creates client sockets; Adopt admits a server socket already accepted by an
 * HttpListener.  Both then share the exact same receive/send/close implementation.
 ******************************************************************************/
using System;
using System.Collections.Generic;
using System.IO;
using System.Net.WebSockets;
using System.Numerics;
using System.Text;
using System.Threading;
using Dafny;
using Rune = Dafny.Rune;

namespace ConfluxWebSocketCapability {
  public partial class __default {
    private sealed class Entry {
      internal readonly WebSocket Socket;
      internal readonly object ReceiveLock = new object();
      internal readonly object SendLock = new object();
      internal Entry(WebSocket socket) { Socket = socket; }
    }

    private static readonly object _tableLock = new object();
    private static readonly Dictionary<int, Entry> _entries = new Dictionary<int, Entry>();
    private static int _nextHandle = 1;

    private static string Read(ISequence<Rune> value) => value.ToVerbatimString(false);
    private static ISequence<Rune> Write(string value) => Sequence<Rune>.UnicodeFromString(value ?? "");

    private static BigInteger Register(WebSocket socket) {
      if (socket == null) throw new ArgumentNullException(nameof(socket));
      lock (_tableLock) {
        int handle = _nextHandle++;
        _entries[handle] = new Entry(socket);
        return new BigInteger(handle);
      }
    }

    // Native acceptors use this one admission point, then relinquish direct socket I/O to the capability.
    public static BigInteger Adopt(WebSocket socket) => Register(socket);

    private static bool TryEntry(BigInteger conn, out Entry entry) {
      lock (_tableLock) return _entries.TryGetValue((int)conn, out entry);
    }

    public static void ConnectWebSocket(ISequence<Rune> url, out bool connected, out BigInteger conn) {
      var socket = new ClientWebSocket();
      try {
        socket.ConnectAsync(new Uri(Read(url)), CancellationToken.None).GetAwaiter().GetResult();
        conn = Register(socket);
        connected = true;
      } catch {
        try { socket.Abort(); } catch { }
        try { socket.Dispose(); } catch { }
        conn = new BigInteger(-1);
        connected = false;
      }
    }

    public static void ReceiveWebSocketText(BigInteger conn, out bool open, out bool isText,
        out ISequence<Rune> text) {
      open = false;
      isText = false;
      text = Write("");
      if (!TryEntry(conn, out var entry)) return;

      lock (entry.ReceiveLock) {
        try {
          var bytes = new MemoryStream();
          var buffer = new byte[64 * 1024];
          WebSocketMessageType? messageType = null;
          while (true) {
            var result = entry.Socket.ReceiveAsync(
              new ArraySegment<byte>(buffer), CancellationToken.None).GetAwaiter().GetResult();
            if (result.MessageType == WebSocketMessageType.Close) return;
            messageType ??= result.MessageType;
            if (result.MessageType != messageType) return;
            bytes.Write(buffer, 0, result.Count);
            if (result.EndOfMessage) break;
          }
          open = true;
          if (messageType != WebSocketMessageType.Text) return;
          text = Write(Encoding.UTF8.GetString(bytes.ToArray()));
          isText = true;
        } catch { }
      }
    }

    public static bool SendWebSocketText(BigInteger conn, ISequence<Rune> text) {
      if (!TryEntry(conn, out var entry)) return false;
      var payload = Encoding.UTF8.GetBytes(Read(text));
      lock (entry.SendLock) {
        try {
          entry.Socket.SendAsync(new ArraySegment<byte>(payload), WebSocketMessageType.Text, true,
            CancellationToken.None).GetAwaiter().GetResult();
          return true;
        } catch { return false; }
      }
    }

    public static void CloseWebSocket(BigInteger conn) {
      Entry entry = null;
      lock (_tableLock) {
        if (_entries.TryGetValue((int)conn, out entry)) _entries.Remove((int)conn);
      }
      if (entry == null) return;
      try {
        if (entry.Socket.State == WebSocketState.Open || entry.Socket.State == WebSocketState.CloseReceived)
          entry.Socket.CloseAsync(WebSocketCloseStatus.NormalClosure, "", CancellationToken.None)
            .GetAwaiter().GetResult();
      } catch { try { entry.Socket.Abort(); } catch { } }
      try { entry.Socket.Dispose(); } catch { }
    }
  }
}
