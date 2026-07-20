/*******************************************************************************
 *  conflux-extern-impl — conflux-extern / socket: the .NET-primary shim for
 *  `module {:extern} ConfluxSocketFfi replaces ConfluxSocketCapability` (spec/socket.dfy).
 *
 *  Namespace = the BASE module `ConfluxSocketCapability` (a `replaces` module inherits its extern name
 *  from the Layer-1 base). `--target cs` co-translates this file. This is a trusted FFI body in the
 *  constellation: a REAL, newline-framed TCP transport over System.Net.Sockets (localhost / IPv4
 *  loopback), moved VERBATIM from the socket-demo's `AhpSocket.__default` — the only change is the
 *  namespace (AhpSocket -> ConfluxSocketCapability) so the AHP host/client run-channels bind this seam
 *  instead of declaring an {:extern} of their own. Every method is pure I/O + thread plumbing — no domain
 *  logic (gate-extern-io-only enforces it; the `conflux-extern-impl` sentinel above marks it in-home).
 *
 *  Dafny {:extern} methods are synchronous; C# sockets block naturally (AcceptSocket / Connect / Send /
 *  Receive), so the extern surface matches the Dafny call model with no async bridging. Handles (listener
 *  / connection) are opaque ints registered in process-local dictionaries; Dafny `int` marshals as
 *  System.Numerics.BigInteger and `seq<bv8>` as Dafny.ISequence<byte>.
 *
 *  ── CONCURRENCY MODEL (Stage 3, where it lives) ──
 *  The Stage-2 methods (Listen/Accept/Connect/SendLine/RecvLine/Close) are SYNCHRONOUS and single-socket.
 *
 *  The Stage-3 methods add GENUINE concurrency AT THE HOST, entirely inside this extern (the verified
 *  Host.accept / Client.receive fold stays pure Dafny above this boundary):
 *    - AcceptSubscribers(server, k): blocking-accepts K client connections into a subscriber list, and
 *      spawns ONE background RECV-THREAD per client.
 *    - Each recv-thread independently reads newline-framed command lines from its OWN client socket and
 *      pushes them into a shared, lock-guarded command queue (Monitor.Wait/PulseAll). Because K
 *      recv-threads run in parallel, commands from K clients ARRIVE CONCURRENTLY and get serialized by the
 *      queue into one arbitrary interleaving — the real source of concurrency.
 *    - DequeueCommand(): the host's single main thread BLOCKS pulling the next command off the queue;
 *      returns an EMPTY sentinel once the queue is drained AND every recv-thread has seen its client close.
 *    - Broadcast(line): the main thread fans one host-sequenced line out to ALL subscriber sockets under a
 *      lock. Reading (recv-threads) and writing (main thread) happen on the same sockets simultaneously —
 *      genuine full-duplex at the host.
 *
 *  The VERIFIED fold (Host.accept / Client.receive / AhpSkeleton.apply1) and the Fidelity codec stay ABOVE
 *  this boundary, in Dafny. This file only moves bytes + provides the thread-safe multiplex + fan-out
 *  plumbing.
 ******************************************************************************/
using System;
using System.Collections.Generic;
using System.Net;
using System.Net.Sockets;
using System.Numerics;
using System.Threading;
using Dafny;

namespace ConfluxSocketCapability {
  public partial class __default {
    private static readonly object _lock = new object();
    private static int _nextId = 1;
    private static readonly Dictionary<int, Socket> _socks = new Dictionary<int, Socket>();
    private static readonly Dictionary<int, TcpListener> _listeners = new Dictionary<int, TcpListener>();

    private static int Register(Socket s) {
      lock (_lock) { int id = _nextId++; _socks[id] = s; return id; }
    }

    // Bind + listen on 127.0.0.1:<port>; return an opaque listener handle.
    public static BigInteger Listen(BigInteger port) {
      var listener = new TcpListener(IPAddress.Loopback, (int)port);
      listener.Start();
      lock (_lock) { int id = _nextId++; _listeners[id] = listener; return new BigInteger(id); }
    }

    // BLOCKING accept of exactly one connection; return an opaque conn handle.
    public static BigInteger Accept(BigInteger server) {
      TcpListener listener;
      lock (_lock) { listener = _listeners[(int)server]; }
      Socket s = listener.AcceptSocket();          // blocks until a client connects
      s.NoDelay = true;
      return new BigInteger(Register(s));
    }

    // Connect to 127.0.0.1:<port>; retry briefly in case the host is still
    // coming up (the driver script also polls the port before launching us).
    public static BigInteger Connect(BigInteger port) {
      int p = (int)port;
      Exception last = null;
      for (int attempt = 0; attempt < 100; attempt++) {
        try {
          var s = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
          s.Connect(IPAddress.Loopback, p);
          s.NoDelay = true;
          return new BigInteger(Register(s));
        } catch (Exception e) {
          last = e;
          Thread.Sleep(50);
        }
      }
      throw new Exception("ConfluxSocketCapability.Connect: could not reach 127.0.0.1:" + p, last);
    }

    // Write payload bytes followed by a single '\n' frame terminator.
    public static void SendLine(BigInteger conn, ISequence<byte> payload) {
      Socket s;
      lock (_lock) { s = _socks[(int)conn]; }
      byte[] data = payload.CloneAsArray();
      byte[] framed = new byte[data.Length + 1];
      Array.Copy(data, framed, data.Length);
      framed[data.Length] = (byte)'\n';
      int sent = 0;
      while (sent < framed.Length) {
        sent += s.Send(framed, sent, framed.Length - sent, SocketFlags.None);
      }
    }

    // BLOCKING read of one '\n'-framed line; the terminator is stripped.
    public static ISequence<byte> RecvLine(BigInteger conn) {
      Socket s;
      lock (_lock) { s = _socks[(int)conn]; }
      var buf = new List<byte>();
      var one = new byte[1];
      while (true) {
        int n = s.Receive(one, 0, 1, SocketFlags.None);   // blocks for the next byte
        if (n == 0) break;                                // peer closed the connection
        if (one[0] == (byte)'\n') break;                  // end of frame
        buf.Add(one[0]);
      }
      return Helpers.SeqFromArray(buf.ToArray());
    }

    public static void Close(BigInteger conn) {
      Socket s = null;
      lock (_lock) { if (_socks.TryGetValue((int)conn, out s)) _socks.Remove((int)conn); }
      if (s != null) {
        try { s.Shutdown(SocketShutdown.Both); } catch { }
        s.Close();
      }
    }

    public static void CloseServer(BigInteger server) {
      TcpListener l = null;
      lock (_lock) { if (_listeners.TryGetValue((int)server, out l)) _listeners.Remove((int)server); }
      if (l != null) l.Stop();
    }

    // ══════════════════════════════════════════════════════════════════════════
    //  STAGE 3 — CONCURRENT MULTIPLEX + MULTI-SUBSCRIBER BROADCAST (host side).
    //
    //  All of the following state + methods exist to give the host GENUINE
    //  arrival-concurrency across K clients and a fan-out broadcast to all of
    //  them, while the verified Host.accept fold stays pure Dafny above. Each is
    //  process-local static state; the host binary is the only process that uses
    //  it (the client binary compiles the same file but never calls these).
    // ══════════════════════════════════════════════════════════════════════════

    // Subscriber sockets (the K connected clients), guarded by _subLock. The main
    // thread iterates this list to fan a broadcast out to everyone.
    private static readonly object _subLock = new object();
    private static readonly List<Socket> _subscribers = new List<Socket>();

    // The shared command multiplex queue. K recv-threads ENQUEUE concurrently;
    // the single host main thread DEQUEUEs. _cmdQueue, _openClients and the
    // Monitor wait/pulse are all guarded by _qLock.
    private static readonly object _qLock = new object();
    private static readonly Queue<byte[]> _cmdQueue = new Queue<byte[]>();
    private static int _openClients = 0;   // recv-threads still reading (not yet EOF)

    // Blocking-accept exactly K client connections, register each as a subscriber,
    // and spawn ONE background recv-thread per client. Returns after all K are
    // connected and their recv-threads are running. This is what makes arrivals
    // concurrent: from this point K threads read K sockets in parallel.
    public static void AcceptSubscribers(BigInteger server, BigInteger k) {
      TcpListener listener;
      lock (_lock) { listener = _listeners[(int)server]; }
      int kk = (int)k;
      lock (_qLock) { _openClients = kk; }
      for (int i = 0; i < kk; i++) {
        Socket s = listener.AcceptSocket();   // blocks until client i connects
        s.NoDelay = true;
        lock (_subLock) { _subscribers.Add(s); }
        var t = new Thread(() => RecvThread(s));
        t.IsBackground = true;
        t.Name = "ahp-recv-" + i;
        t.Start();
      }
    }

    // Per-client recv-thread: read newline-framed command lines from ONE client
    // socket and push each onto the shared queue (waking the main-thread waiter).
    // On EOF (the client closed after its last send) decrement the open-client
    // count and wake the waiter so DequeueCommand can re-check the termination
    // condition. Runs concurrently with the other K-1 recv-threads AND with the
    // main thread's Broadcast writes to this same socket (full-duplex).
    private static void RecvThread(Socket s) {
      try {
        while (true) {
          byte[] line = ReadFrame(s);
          if (line == null) break;   // clean EOF — client closed the connection
          lock (_qLock) { _cmdQueue.Enqueue(line); Monitor.PulseAll(_qLock); }
        }
      } catch { /* any socket error on this client is treated as its EOF */ }
      lock (_qLock) { _openClients--; Monitor.PulseAll(_qLock); }
    }

    // Read one '\n'-framed line from a socket. Returns the line WITHOUT the
    // terminator, or null on EOF (peer closed with no pending partial line).
    private static byte[] ReadFrame(Socket s) {
      var buf = new List<byte>();
      var one = new byte[1];
      while (true) {
        int n;
        try { n = s.Receive(one, 0, 1, SocketFlags.None); }
        catch { return buf.Count == 0 ? null : buf.ToArray(); }
        if (n == 0) return buf.Count == 0 ? null : buf.ToArray();   // EOF
        if (one[0] == (byte)'\n') return buf.ToArray();             // end of frame
        buf.Add(one[0]);
      }
    }

    // BLOCKING dequeue of the next multiplexed command, called by the host's main
    // thread. Blocks while the queue is empty AND at least one client is still
    // open. Returns the next command's bytes, or an EMPTY sequence (the sentinel)
    // once the queue is drained AND every client has closed — signalling the host
    // loop to terminate. The empty-sentinel matches RecvLine's EOF convention so
    // the Dafny side tests `|payload| == 0` identically.
    public static ISequence<byte> DequeueCommand() {
      lock (_qLock) {
        while (_cmdQueue.Count == 0 && _openClients > 0) {
          Monitor.Wait(_qLock);
        }
        if (_cmdQueue.Count > 0) {
          byte[] line = _cmdQueue.Dequeue();
          return Helpers.SeqFromArray(line);
        }
        return Helpers.SeqFromArray(new byte[0]);   // drained + all closed → sentinel
      }
    }

    // Fan one '\n'-framed line out to EVERY subscriber socket, under _subLock so a
    // concurrent AcceptSubscribers can't mutate the list mid-broadcast. This is
    // the broadcast leg: one host-sequenced action reaches all K mirrors. A
    // subscriber that has already closed is skipped (its Send throws → caught).
    public static void Broadcast(ISequence<byte> payload) {
      byte[] data = payload.CloneAsArray();
      byte[] framed = new byte[data.Length + 1];
      Array.Copy(data, framed, data.Length);
      framed[data.Length] = (byte)'\n';
      lock (_subLock) {
        foreach (var s in _subscribers) {
          try {
            int sent = 0;
            while (sent < framed.Length) sent += s.Send(framed, sent, framed.Length - sent, SocketFlags.None);
          } catch { /* a closed subscriber is skipped */ }
        }
      }
    }

    // Close every subscriber socket (called by the host after the loop ends).
    public static void CloseAll() {
      lock (_subLock) {
        foreach (var s in _subscribers) {
          try { s.Shutdown(SocketShutdown.Both); } catch { }
          try { s.Close(); } catch { }
        }
        _subscribers.Clear();
      }
    }

    // ══════════════════════════════════════════════════════════════════════════
    //  STAGE 3b — CONCURRENT CLIENT DUPLEX (client side).
    //
    //  The Stage-2/3 client was SEND-THEN-RECV: it flushed all its commands, THEN
    //  read the broadcast stream on one (main) thread. These methods let ONE client
    //  process send AND receive CONCURRENTLY: SpawnSender starts a REAL background
    //  C# Thread that writes the client's M command lines to the socket (with a
    //  tiny sleep between sends so the writes genuinely OVERLAP the main thread's
    //  reads), and returns immediately. The Dafny main thread then enters its
    //  RecvLine + verified Client.receive fold loop WHILE the sender thread is still
    //  writing — a genuine full-duplex client (send on the background thread, recv
    //  on the main thread, both on the same socket in opposite directions, which
    //  .NET supports for a single reader + single writer).
    //
    //  The timing accessors (NowMillis / GetFirstSendMillis / GetLastSendMillis /
    //  GetSendCount) let the Dafny side PROVE the overlap: if a receive fold
    //  completes before the sender's last write (firstRecv < lastSend), the two
    //  directions genuinely ran at the same time — that is impossible under
    //  send-then-recv, where every send precedes every recv.
    // ══════════════════════════════════════════════════════════════════════════

    // One shared monotonic clock, so timestamps taken on the sender thread (C#) and
    // on the main thread (via NowMillis, called from Dafny) share a timebase.
    private static readonly System.Diagnostics.Stopwatch _clock = System.Diagnostics.Stopwatch.StartNew();

    private static readonly object _senderLock = new object();
    private static readonly Dictionary<int, Thread> _senders = new Dictionary<int, Thread>();
    private static readonly Dictionary<int, long> _firstSendMs = new Dictionary<int, long>();
    private static readonly Dictionary<int, long> _lastSendMs = new Dictionary<int, long>();
    private static readonly Dictionary<int, int> _sendCount = new Dictionary<int, int>();

    // Delay (ms) between the client's successive command writes, so the sender
    // thread stays busy WHILE the main thread reads — making the overlap real and
    // observable rather than an instant flush. Overridable for tuning/tests.
    private static readonly int _sendGapMs = ReadIntEnv("AHP_CLIENT_SEND_GAP_MS", 40);
    private static int ReadIntEnv(string name, int dflt) {
      var v = Environment.GetEnvironmentVariable(name);
      return (v != null && int.TryParse(v, out int n) && n >= 0) ? n : dflt;
    }

    // Spawn ONE background thread that writes the given newline-framed lines to the
    // client's socket, sleeping _sendGapMs between successive writes. The Dafny
    // sequences are materialized into byte[][] HERE (on the calling thread) so the
    // background thread never touches Dafny objects concurrently. Returns as soon
    // as the thread is started — the sends happen off the main thread.
    public static void SpawnSender(BigInteger conn, ISequence<ISequence<byte>> lines) {
      int cid = (int)conn;
      Socket s;
      lock (_lock) { s = _socks[cid]; }

      ISequence<byte>[] outer = lines.CloneAsArray();
      var framed = new List<byte[]>(outer.Length);
      foreach (var line in outer) {
        byte[] data = line.CloneAsArray();
        byte[] f = new byte[data.Length + 1];
        Array.Copy(data, f, data.Length);
        f[data.Length] = (byte)'\n';
        framed.Add(f);
      }

      lock (_senderLock) { _firstSendMs[cid] = -1; _lastSendMs[cid] = -1; _sendCount[cid] = 0; }

      var thread = new Thread(() => {
        for (int i = 0; i < framed.Count; i++) {
          if (i > 0 && _sendGapMs > 0) Thread.Sleep(_sendGapMs);  // overlap the reader
          byte[] frame = framed[i];
          try {
            int sent = 0;
            while (sent < frame.Length) sent += s.Send(frame, sent, frame.Length - sent, SocketFlags.None);
          } catch { break; }   // socket error → stop sending; main thread will notice via fold
          long now = _clock.ElapsedMilliseconds;
          lock (_senderLock) {
            if (_firstSendMs[cid] < 0) _firstSendMs[cid] = now;
            _lastSendMs[cid] = now;
            _sendCount[cid] = _sendCount[cid] + 1;
          }
        }
      });
      thread.IsBackground = true;
      thread.Name = "ahp-send-" + cid;
      lock (_senderLock) { _senders[cid] = thread; }
      thread.Start();
    }

    // Block until this connection's sender thread has finished writing every line.
    // Called by the main thread AFTER it has folded all broadcasts (by which point
    // the sender is structurally done — the host cannot have broadcast this
    // client's last command without first receiving it), so this only joins an
    // already-finished thread; it never serializes the duplex.
    public static void JoinSender(BigInteger conn) {
      Thread t = null;
      lock (_senderLock) { _senders.TryGetValue((int)conn, out t); }
      if (t != null) t.Join();
    }

    // Shared-timebase "now" in ms, for the main thread to stamp its first recv fold.
    public static BigInteger NowMillis() { return new BigInteger(_clock.ElapsedMilliseconds); }

    public static BigInteger GetFirstSendMillis(BigInteger conn) {
      lock (_senderLock) { return new BigInteger(_firstSendMs.TryGetValue((int)conn, out long v) ? v : -1); }
    }
    public static BigInteger GetLastSendMillis(BigInteger conn) {
      lock (_senderLock) { return new BigInteger(_lastSendMs.TryGetValue((int)conn, out long v) ? v : -1); }
    }
    public static BigInteger GetSendCount(BigInteger conn) {
      lock (_senderLock) { return new BigInteger(_sendCount.TryGetValue((int)conn, out int v) ? v : 0); }
    }
  }
}
