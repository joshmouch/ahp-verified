/*******************************************************************************
 * conflux-extern-impl - native HTTP transport for ConfluxHttpCapability.
 * Domain routing and payload construction live in Dafny; this file only moves
 * request and response bytes and keeps event-stream connections open.
 *******************************************************************************/
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Numerics;
using System.Text;
using Dafny;
using Rune = Dafny.Rune;

namespace ConfluxHttpCapability {
  public partial class __default {
    private sealed class ServerState {
      public TcpListener Listener;
      public readonly List<NetworkStream> Streams = new List<NetworkStream>();
    }

    private sealed class RequestState {
      public int Server;
      public TcpClient Client;
      public NetworkStream Stream;
    }

    private static readonly object Gate = new object();
    private static readonly Dictionary<int, ServerState> Servers = new Dictionary<int, ServerState>();
    private static readonly Dictionary<int, RequestState> Requests = new Dictionary<int, RequestState>();
    private static int NextHandle = 1;
    private static ISequence<Rune> D(string value) => Sequence<Rune>.UnicodeFromString(value ?? "");

    private static int RegisterServer(ServerState state) {
      lock (Gate) { int id = NextHandle++; Servers[id] = state; return id; }
    }

    private static int RegisterRequest(RequestState state) {
      lock (Gate) { int id = NextHandle++; Requests[id] = state; return id; }
    }

    private static IPAddress ResolveAddress(string host) {
      if (String.IsNullOrWhiteSpace(host) || host == "0.0.0.0" || host == "*") return IPAddress.Any;
      if (host == "localhost") return IPAddress.Loopback;
      if (IPAddress.TryParse(host, out var parsed)) return parsed;
      return Dns.GetHostAddresses(host).First(address => address.AddressFamily == AddressFamily.InterNetwork);
    }

    public static void Start(ISequence<Rune> host, BigInteger port, out BigInteger server, out BigInteger boundPort) {
      var listener = new TcpListener(ResolveAddress(host.ToVerbatimString(false)), (int)port);
      listener.Start();
      int id = RegisterServer(new ServerState { Listener = listener });
      int bound = ((IPEndPoint)listener.LocalEndpoint).Port;
      server = new BigInteger(id);
      boundPort = new BigInteger(bound);
    }

    private static bool ReadRequest(NetworkStream stream, out string method, out string path, out string body) {
      method = ""; path = ""; body = "";
      stream.ReadTimeout = 5000;
      var bytes = new List<byte>();
      var one = new byte[1];
      int headerEnd = -1;
      while (bytes.Count < 2 * 1024 * 1024) {
        int read = stream.Read(one, 0, 1);
        if (read == 0) return false;
        bytes.Add(one[0]);
        int n = bytes.Count;
        if (n >= 4 && bytes[n - 4] == 13 && bytes[n - 3] == 10 && bytes[n - 2] == 13 && bytes[n - 1] == 10) {
          headerEnd = n;
          break;
        }
      }
      if (headerEnd < 0) return false;
      string header = Encoding.UTF8.GetString(bytes.ToArray(), 0, headerEnd);
      string[] lines = header.Split(new[] { "\r\n" }, StringSplitOptions.None);
      string[] requestLine = lines[0].Split(' ');
      if (requestLine.Length < 2) return false;
      method = requestLine[0];
      path = requestLine[1];
      int contentLength = 0;
      foreach (string line in lines.Skip(1)) {
        int colon = line.IndexOf(':');
        if (colon > 0 && line.Substring(0, colon).Trim().Equals("Content-Length", StringComparison.OrdinalIgnoreCase)) {
          Int32.TryParse(line.Substring(colon + 1).Trim(), out contentLength);
        }
      }
      if (contentLength < 0 || contentLength > 2 * 1024 * 1024) return false;
      var payload = new byte[contentLength];
      int offset = 0;
      while (offset < contentLength) {
        int read = stream.Read(payload, offset, contentLength - offset);
        if (read == 0) return false;
        offset += read;
      }
      body = Encoding.UTF8.GetString(payload);
      return true;
    }

    public static void AcceptWithin(BigInteger server, BigInteger waitMs, out bool available, out BigInteger request,
        out ISequence<Rune> httpMethod, out ISequence<Rune> path, out ISequence<Rune> body) {
      ServerState state;
      lock (Gate) { Servers.TryGetValue((int)server, out state); }
      if (state == null || !state.Listener.Server.Poll(Math.Max(0, (int)waitMs) * 1000, SelectMode.SelectRead)) {
        available = false; request = BigInteger.Zero; httpMethod = D(""); path = D(""); body = D(""); return;
      }
      TcpClient client = null;
      try {
        client = state.Listener.AcceptTcpClient();
        client.NoDelay = true;
        NetworkStream stream = client.GetStream();
        if (!ReadRequest(stream, out string method, out string requestPath, out string requestBody)) {
          client.Close();
          available = false; request = BigInteger.Zero; httpMethod = D(""); path = D(""); body = D(""); return;
        }
        int requestHandle = RegisterRequest(new RequestState { Server = (int)server, Client = client, Stream = stream });
        available = true; request = new BigInteger(requestHandle); httpMethod = D(method); path = D(requestPath); body = D(requestBody);
      } catch {
        try { client?.Close(); } catch { }
        available = false; request = BigInteger.Zero; httpMethod = D(""); path = D(""); body = D(""); return;
      }
    }

    private static string Reason(int status) {
      switch (status) {
        case 200: return "OK";
        case 204: return "No Content";
        case 400: return "Bad Request";
        case 404: return "Not Found";
        case 405: return "Method Not Allowed";
        default: return "Response";
      }
    }

    private static byte[] HeaderBytes(int status, string contentType, IEnumerable<string> headers, int? length, bool keepAlive) {
      var builder = new StringBuilder();
      builder.Append("HTTP/1.1 ").Append(status).Append(' ').Append(Reason(status)).Append("\r\n");
      if (!String.IsNullOrEmpty(contentType)) builder.Append("Content-Type: ").Append(contentType).Append("\r\n");
      if (length.HasValue) builder.Append("Content-Length: ").Append(length.Value).Append("\r\n");
      builder.Append("Connection: ").Append(keepAlive ? "keep-alive" : "close").Append("\r\n");
      foreach (string header in headers) builder.Append(header).Append("\r\n");
      builder.Append("\r\n");
      return Encoding.UTF8.GetBytes(builder.ToString());
    }

    private static string[] Strings(ISequence<ISequence<Rune>> values) {
      return values.CloneAsArray().Select(value => value.ToVerbatimString(false)).ToArray();
    }

    public static void Respond(BigInteger request, BigInteger status, ISequence<Rune> contentType, ISequence<ISequence<Rune>> headers, ISequence<Rune> body) {
      RequestState state;
      lock (Gate) { Requests.TryGetValue((int)request, out state); Requests.Remove((int)request); }
      if (state == null) return;
      try {
        byte[] payload = Encoding.UTF8.GetBytes(body.ToVerbatimString(false));
        byte[] head = HeaderBytes((int)status, contentType.ToVerbatimString(false), Strings(headers), payload.Length, false);
        state.Stream.Write(head, 0, head.Length);
        state.Stream.Write(payload, 0, payload.Length);
      } catch {
        // Client disconnect is an ordinary transport outcome, not a host-fatal exception.
      } finally {
        state.Client.Close();
      }
    }

    public static void BeginStream(BigInteger request, ISequence<Rune> contentType, ISequence<ISequence<Rune>> headers, ISequence<Rune> initialBody) {
      RequestState state;
      lock (Gate) { Requests.TryGetValue((int)request, out state); Requests.Remove((int)request); }
      if (state == null) return;
      try {
        byte[] head = HeaderBytes(200, contentType.ToVerbatimString(false), Strings(headers), null, true);
        byte[] payload = Encoding.UTF8.GetBytes(initialBody.ToVerbatimString(false));
        state.Stream.Write(head, 0, head.Length);
        state.Stream.Write(payload, 0, payload.Length);
        state.Stream.Flush();
        lock (Gate) { if (Servers.TryGetValue(state.Server, out var server)) server.Streams.Add(state.Stream); }
      } catch {
        state.Client.Close();
      }
    }

    public static void Publish(BigInteger server, ISequence<Rune> payload) {
      byte[] bytes = Encoding.UTF8.GetBytes(payload.ToVerbatimString(false));
      lock (Gate) {
        if (!Servers.TryGetValue((int)server, out var state)) return;
        for (int i = state.Streams.Count - 1; i >= 0; i--) {
          try { state.Streams[i].Write(bytes, 0, bytes.Length); state.Streams[i].Flush(); }
          catch { try { state.Streams[i].Close(); } catch { } state.Streams.RemoveAt(i); }
        }
      }
    }

    public static void Stop(BigInteger server) {
      ServerState state;
      lock (Gate) { Servers.TryGetValue((int)server, out state); Servers.Remove((int)server); }
      if (state == null) return;
      state.Listener.Stop();
      lock (Gate) { foreach (var stream in state.Streams) try { stream.Close(); } catch { } state.Streams.Clear(); }
    }
  }
}
