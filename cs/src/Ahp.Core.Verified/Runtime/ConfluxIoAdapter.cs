// conflux-extern / io — .NET-primary shim for `module {:extern} ConfluxIoFfi replaces
// ConfluxIoCapability`. Namespace = the BASE module `ConfluxIoCapability` (a `replaces` module inherits
// its extern name from the Layer-1 base). `--target cs` links this file. This is the ONE trusted FFI
// body in the constellation: the ReadFile/WriteFile/DeleteFile/ListDir/RunProcess bodies are the
// archai-proven arche-extern shims (IoCapabilityAdapter.cs + ProcessAdapter.cs) moved here verbatim —
// arche-extern is now DELETED, arche consumes this seam directly; AppendFile / FileStat / ReadByteRange /
// GetEnv extend them for the incremental-tail + poll + root-resolution providers. Every method is pure
// I/O — no domain logic (gate-extern-io-only enforces it).
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Numerics;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading;
using Rune = Dafny.Rune;

namespace ConfluxIoCapability {
  public partial class __default {
    static string S(Dafny.ISequence<Rune> x) => x.ToVerbatimString(false);
    static Dafny.ISequence<Rune> D(string x) => Dafny.Sequence<Rune>.UnicodeFromString(x ?? "");
    static void EnsureDir(string p) { var dir = Path.GetDirectoryName(Path.GetFullPath(p)); if (!string.IsNullOrEmpty(dir)) Directory.CreateDirectory(dir); }
    [DllImport("libc", SetLastError = true)]
    static extern int open(string path, int flags);
    [DllImport("libc", SetLastError = true)]
    static extern int fsync(int descriptor);
    [DllImport("libc", SetLastError = true)]
    static extern int close(int descriptor);
    [DllImport("libc", SetLastError = true)]
    static extern int kill(int pid, int signal);

    static void FlushParentDirectory(string target) {
      var directory = Path.GetDirectoryName(target);
      if (string.IsNullOrEmpty(directory)) return;
      if (OperatingSystem.IsWindows()) return;
      var descriptor = open(directory, 0);
      if (descriptor < 0) throw new IOException(
        "open parent directory failed", new System.ComponentModel.Win32Exception(Marshal.GetLastWin32Error()));
      try {
        if (fsync(descriptor) != 0) throw new IOException(
          "fsync parent directory failed", new System.ComponentModel.Win32Exception(Marshal.GetLastWin32Error()));
      } finally {
        close(descriptor);
      }
    }

    sealed class WatchState {
      public readonly object Gate = new object();
      public readonly Queue<string> Paths = new Queue<string>();
      public readonly List<FileSystemWatcher> Watchers = new List<FileSystemWatcher>();
    }
    static readonly object WatchGate = new object();
    static readonly Dictionary<int, WatchState> Watches = new Dictionary<int, WatchState>();
    static int NextWatch = 1;
    sealed class StdinState {
      public readonly object Gate = new object();
      public readonly List<byte> Bytes = new List<byte>();
      public readonly AutoResetEvent Ready = new AutoResetEvent(false);
      public bool Started;
      public bool Open = true;
      public string Failure = "";
    }
    static readonly StdinState Stdin = new StdinState();
    const int MaxStdinFrameBytes = 1024 * 1024;

    static bool StdinFrameExceedsLimit() {
      int newline = Stdin.Bytes.IndexOf((byte)'\n');
      return newline > MaxStdinFrameBytes || (newline < 0 && Stdin.Bytes.Count > MaxStdinFrameBytes);
    }

    static void FailOversizedStdinFrame() {
      Stdin.Bytes.Clear();
      Stdin.Open = false;
      Stdin.Failure = "standard input frame exceeds 1048576 bytes";
    }

    public static void ReadFile(Dafny.ISequence<Rune> path, out bool present, out Dafny.ISequence<Rune> content) {
      var p = S(path);
      try { if (File.Exists(p)) { content = D(File.ReadAllText(p)); present = true; return; } } catch { }
      present = false; content = D("");
    }

    public static void ReadFileBytes(Dafny.ISequence<Rune> path, out bool present,
        out Dafny.ISequence<byte> content) {
      try {
        var p = S(path);
        if (File.Exists(p)) {
          content = Dafny.Sequence<byte>.FromArray(File.ReadAllBytes(p));
          present = true;
          return;
        }
      } catch { }
      present = false;
      content = Dafny.Sequence<byte>.FromArray(Array.Empty<byte>());
    }

    public static void ReadLines(Dafny.ISequence<Rune> path, out bool present,
        out Dafny.ISequence<Dafny.ISequence<Rune>> lines) {
      try {
        string p = S(path);
        if (File.Exists(p)) {
          lines = Dafny.Sequence<Dafny.ISequence<Rune>>.FromArray(File.ReadAllLines(p).Select(D).ToArray());
          present = true;
          return;
        }
      } catch { }
      present = false;
      lines = Dafny.Sequence<Dafny.ISequence<Rune>>.FromArray(Array.Empty<Dafny.ISequence<Rune>>());
    }

    static void StartStdinReader() {
      lock (Stdin.Gate) {
        if (Stdin.Started) return;
        Stdin.Started = true;
      }
      var thread = new Thread(() => {
        try {
          var stream = Console.OpenStandardInput();
          var buffer = new byte[4096];
          while (true) {
            int count = stream.Read(buffer, 0, buffer.Length);
            lock (Stdin.Gate) {
              if (count <= 0) Stdin.Open = false;
              else {
                for (int i = 0; i < count; i++) Stdin.Bytes.Add(buffer[i]);
                if (StdinFrameExceedsLimit()) FailOversizedStdinFrame();
              }
            }
            Stdin.Ready.Set();
            if (count <= 0 || !Stdin.Open) return;
          }
        } catch (Exception error) {
          lock (Stdin.Gate) {
            Stdin.Open = false;
            Stdin.Failure = string.IsNullOrEmpty(error.Message) ? error.GetType().Name : error.Message;
          }
          Stdin.Ready.Set();
        }
      });
      thread.IsBackground = true;
      thread.Start();
    }

    static bool TryTakeStdinLine(out string value, out bool available, out bool open,
        out bool failed, out string diagnostic) {
      lock (Stdin.Gate) {
        value = ""; available = false; open = Stdin.Open; failed = false; diagnostic = "";
        if (Stdin.Failure.Length > 0) {
          open = false; failed = true; diagnostic = Stdin.Failure; return true;
        }
        if (StdinFrameExceedsLimit()) {
          FailOversizedStdinFrame();
          open = false; failed = true; diagnostic = Stdin.Failure; return true;
        }
        int newline = Stdin.Bytes.IndexOf((byte)'\n');
        int length = newline >= 0 ? newline : (!Stdin.Open && Stdin.Bytes.Count > 0 ? Stdin.Bytes.Count : -1);
        if (length >= 0) {
          var bytes = Stdin.Bytes.GetRange(0, length).ToArray();
          Stdin.Bytes.RemoveRange(0, newline >= 0 ? length + 1 : length);
          if (bytes.Length > 0 && bytes[bytes.Length - 1] == (byte)'\r') Array.Resize(ref bytes, bytes.Length - 1);
          try {
            value = new UTF8Encoding(false, true).GetString(bytes);
            available = true; open = Stdin.Open; return true;
          } catch (Exception error) {
            Stdin.Open = false;
            Stdin.Failure = string.IsNullOrEmpty(error.Message) ? error.GetType().Name : error.Message;
            open = false; failed = true; diagnostic = Stdin.Failure; return true;
          }
        }
        if (!Stdin.Open) return true;
        return false;
      }
    }

    public static void ReadStdinLineWithin(BigInteger waitMs, out Dafny.ISequence<Rune> line,
        out bool available, out bool open, out bool failed, out Dafny.ISequence<Rune> diagnostic) {
      StartStdinReader();
      int wait = waitMs > int.MaxValue ? int.MaxValue : Math.Max(0, (int)waitMs);
      var elapsed = Stopwatch.StartNew();
      string value; string detail;
      while (!TryTakeStdinLine(out value, out available, out open, out failed, out detail)) {
        int remaining = Math.Max(0, wait - (int)Math.Min(wait, elapsed.ElapsedMilliseconds));
        if (remaining <= 0 || !Stdin.Ready.WaitOne(remaining)) {
          value = ""; available = false; open = true; failed = false; detail = "";
          break;
        }
      }
      line = D(value);
      diagnostic = D(detail);
    }

    public static void WriteFile(Dafny.ISequence<Rune> path, Dafny.ISequence<Rune> content) {
      var p = S(path); EnsureDir(p); File.WriteAllText(p, S(content));
    }

    public static void AtomicWriteFile(Dafny.ISequence<Rune> path, Dafny.ISequence<Rune> content,
        out bool written, out Dafny.ISequence<Rune> diagnostic) {
      AtomicWriteBytes(path, Dafny.Sequence<byte>.FromArray(new UTF8Encoding(false, true).GetBytes(S(content))),
        out written, out diagnostic);
    }

    public static void AtomicWriteBytes(Dafny.ISequence<Rune> path, Dafny.ISequence<byte> content,
        out bool written, out Dafny.ISequence<Rune> diagnostic) {
      var target = Path.GetFullPath(S(path));
      var temporary = target + ".tmp-" + Guid.NewGuid().ToString("N");
      try {
        EnsureDir(target);
        var bytes = content.Elements.ToArray();
        using (var stream = new FileStream(temporary, FileMode.CreateNew, FileAccess.Write, FileShare.None,
            4096, FileOptions.WriteThrough)) {
          stream.Write(bytes, 0, bytes.Length);
          stream.Flush(true);
        }
        File.Move(temporary, target, true);
        FlushParentDirectory(target);
        written = true;
        diagnostic = D("");
      } catch (Exception error) {
        try { File.Delete(temporary); } catch { }
        written = false;
        diagnostic = D(error.Message);
      }
    }

    public static void InspectPath(Dafny.ISequence<Rune> path, out Dafny.ISequence<Rune> kind,
        out BigInteger mode, out Dafny.ISequence<Rune> symlinkText, out BigInteger size) {
      var p = S(path);
      kind = D("missing"); mode = BigInteger.Zero; symlinkText = D(""); size = BigInteger.Zero;
      try {
        if (!File.Exists(p) && !Directory.Exists(p) && !Path.Exists(p)) return;
        var attributes = File.GetAttributes(p);
        if ((attributes & FileAttributes.ReparsePoint) != 0) {
          var target = Directory.Exists(p) ? new DirectoryInfo(p).LinkTarget : new FileInfo(p).LinkTarget;
          kind = D("symlink");
          symlinkText = D(target ?? "");
        } else if ((attributes & FileAttributes.Directory) != 0) {
          kind = D("directory");
        } else {
          kind = D("file");
          size = new BigInteger(Math.Max(0, new FileInfo(p).Length));
        }
        if (!OperatingSystem.IsWindows()) {
          try { mode = new BigInteger((int)File.GetUnixFileMode(p)); } catch { mode = BigInteger.Zero; }
        }
      } catch {
        kind = D("inaccessible"); mode = BigInteger.Zero; symlinkText = D(""); size = BigInteger.Zero;
      }
    }

    public static void DeleteFile(Dafny.ISequence<Rune> path) { try { File.Delete(S(path)); } catch { } }

    public static void AppendFile(Dafny.ISequence<Rune> path, Dafny.ISequence<Rune> line) {
      var p = S(path); EnsureDir(p); File.AppendAllText(p, S(line) + "\n");
    }

    public static void ListDir(Dafny.ISequence<Rune> dir,
        out Dafny.ISequence<Dafny.ISequence<Rune>> names, out Dafny.ISequence<bool> isDir) {
      var nm = new List<Dafny.ISequence<Rune>>(); var fl = new List<bool>();
      try {
        var d = S(dir);
        var entries = Directory.GetFileSystemEntries(d)
          .Select(path => (name: Path.GetFileName(path), isDirectory: Directory.Exists(path)))
          .OrderBy(entry => entry.name, StringComparer.Ordinal);
        foreach (var entry in entries) { nm.Add(D(entry.name)); fl.Add(entry.isDirectory); }
      } catch { }
      names = Dafny.Sequence<Dafny.ISequence<Rune>>.FromArray(nm.ToArray());
      isDir = Dafny.Sequence<bool>.FromArray(fl.ToArray());
    }

    public static void FileStat(Dafny.ISequence<Rune> path,
        out bool present, out BigInteger mtimeMs, out BigInteger size) {
      var p = S(path);
      try {
        if (File.Exists(p)) {
          var fi = new FileInfo(p);
          mtimeMs = new BigInteger(new DateTimeOffset(File.GetLastWriteTimeUtc(p)).ToUnixTimeMilliseconds());
          size = new BigInteger(fi.Length); present = true; return;
        }
      } catch { }
      present = false; mtimeMs = BigInteger.Zero; size = BigInteger.Zero;
    }

    public static void ReadByteRange(Dafny.ISequence<Rune> path, BigInteger offset, BigInteger length,
        out bool present, out Dafny.ISequence<Rune> content) {
      var p = S(path);
      try {
        if (File.Exists(p)) {
          long off = (long)offset, len = (long)length;
          using var fs = new FileStream(p, FileMode.Open, FileAccess.Read, FileShare.ReadWrite);
          if (off < 0) off = 0; if (off > fs.Length) off = fs.Length;
          long avail = fs.Length - off; if (len < 0) len = 0; if (len > avail) len = avail;
          fs.Seek(off, SeekOrigin.Begin);
          var buf = new byte[len]; int read = 0;
          while (read < len) { int r = fs.Read(buf, read, (int)(len - read)); if (r <= 0) break; read += r; }
          content = D(Encoding.UTF8.GetString(buf, 0, read)); present = true; return;
        }
      } catch { }
      present = false; content = D("");
    }

    public static void RunProcess(Dafny.ISequence<Rune> cwd, Dafny.ISequence<Rune> cmd, Dafny.ISequence<Dafny.ISequence<Rune>> args,
                                  out BigInteger code, out Dafny.ISequence<Rune> outp, out Dafny.ISequence<Rune> err) {
      var psi = new ProcessStartInfo(S(cmd)) { RedirectStandardOutput = true, RedirectStandardError = true, UseShellExecute = false };
      var w = S(cwd); if (w.Length > 0) psi.WorkingDirectory = w;
      foreach (var a in args.Elements) psi.ArgumentList.Add(S(a));
      try {
        var pr = Process.Start(psi);
        string o = pr.StandardOutput.ReadToEnd(), e = pr.StandardError.ReadToEnd(); pr.WaitForExit();
        code = new BigInteger(pr.ExitCode); outp = D(o); err = D(e);
      } catch (Exception ex) { code = new BigInteger(127); outp = D(""); err = D(ex.Message); }
    }

    const int ProcessObservationDeadlineMs = 2000;

    static long DeadlineAfterMilliseconds(int milliseconds) {
      long now = Stopwatch.GetTimestamp();
      long ticks = (long)Math.Ceiling(milliseconds * (double)Stopwatch.Frequency / 1000.0);
      return ticks >= long.MaxValue - now ? long.MaxValue : now + Math.Max(1, ticks);
    }

    static int RemainingDeadlineMilliseconds(long deadlineTimestamp) {
      long remainingTicks = deadlineTimestamp - Stopwatch.GetTimestamp();
      if (remainingTicks <= 0) return 0;
      double remaining = Math.Ceiling(remainingTicks * 1000.0 / Stopwatch.Frequency);
      return remaining >= int.MaxValue ? int.MaxValue : Math.Max(1, (int)remaining);
    }

    static long SnapshotDeadline(long deadlineTimestamp) {
      int remaining = RemainingDeadlineMilliseconds(deadlineTimestamp);
      if (remaining <= 0) return deadlineTimestamp;
      return Math.Min(deadlineTimestamp, DeadlineAfterMilliseconds(Math.Min(50, remaining)));
    }

    static bool TrySampleTreeRssBytes(int rootPid, out long totalBytes, out HashSet<int> treePids,
        out bool rootAbsent, long deadlineTimestamp = 0) {
      totalBytes = 0;
      treePids = new HashSet<int>();
      rootAbsent = false;
      try {
        long deadline = deadlineTimestamp != 0
          ? deadlineTimestamp : DeadlineAfterMilliseconds(ProcessObservationDeadlineMs);
        if (RemainingDeadlineMilliseconds(deadline) <= 0) return false;
        var psi = new ProcessStartInfo("ps") {
          RedirectStandardOutput = true, RedirectStandardError = true, UseShellExecute = false,
        };
        psi.ArgumentList.Add("-axo");
        psi.ArgumentList.Add("pid=,ppid=,rss=");
        using var ps = Process.Start(psi);
        if (ps == null) return false;
        // Begin both drains before waiting so a full pipe cannot deadlock the observation subprocess. Every
        // wait consumes the same monotonic deadline; never block in ReadToEnd before applying that deadline.
        var stdoutRead = ps.StandardOutput.ReadToEndAsync();
        var stderrRead = ps.StandardError.ReadToEndAsync();
        int remaining = RemainingDeadlineMilliseconds(deadline);
        if (remaining <= 0 || !ps.WaitForExit(remaining)) {
          try { ps.Kill(true); } catch { }
          return false;
        }
        remaining = RemainingDeadlineMilliseconds(deadline);
        if ((!stdoutRead.IsCompleted && (remaining <= 0 || !stdoutRead.Wait(remaining))) ||
            (!stderrRead.IsCompleted &&
              ((remaining = RemainingDeadlineMilliseconds(deadline)) <= 0 || !stderrRead.Wait(remaining))) ||
            ps.ExitCode != 0) return false;
        string listing = stdoutRead.Result;
        var rows = new List<(int pid, int ppid, long rssKb)>();
        foreach (string line in listing.Split('\n')) {
          var fields = line.Split((char[])null, StringSplitOptions.RemoveEmptyEntries);
          if (fields.Length < 3 || !int.TryParse(fields[0], out int pid) ||
              !int.TryParse(fields[1], out int ppid) || !long.TryParse(fields[2], out long rssKb)) continue;
          rows.Add((pid, ppid, Math.Max(0, rssKb)));
        }
        if (!rows.Any(row => row.pid == rootPid)) {
          rootAbsent = true;
          return false;
        }
        treePids.Add(rootPid);
        bool added;
        do {
          added = false;
          foreach (var row in rows) {
            if (treePids.Contains(row.ppid) && treePids.Add(row.pid)) added = true;
          }
        } while (added);
        checked {
          foreach (var row in rows) if (treePids.Contains(row.pid)) totalBytes += row.rssKb * 1024L;
        }
        return true;
      } catch {
        totalBytes = 0;
        treePids.Clear();
        rootAbsent = false;
        return false;
      }
    }

    static void KillTrackedTree(Process root, IEnumerable<int> observedPids) {
      // Kill every boundary-observed descendant directly before the root can reparent it. On Unix,
      // Process.Kill(true) stops a process while discovering its children; if discovery fails, that process can
      // remain stopped. A direct SIGKILL cannot leave that intermediate state behind.
      foreach (int pid in observedPids.Distinct().Where(pid =>
          pid > 0 && pid != Environment.ProcessId && pid != root.Id)) {
        try {
          using var process = Process.GetProcessById(pid);
          if (!process.HasExited) {
            if (OperatingSystem.IsWindows()) process.Kill(true);
            else kill(pid, 9);
          }
        } catch { }
      }
      try { if (!root.HasExited) root.Kill(true); } catch { }
    }

    static bool TrackedTreeExited(Process root, IEnumerable<int> observedPids) {
      try { if (!root.HasExited) return false; } catch { }
      foreach (int pid in observedPids.Distinct().Where(pid =>
          pid > 0 && pid != Environment.ProcessId && pid != root.Id)) {
        try {
          using var process = Process.GetProcessById(pid);
          if (!process.HasExited) return false;
        } catch { }
      }
      return true;
    }

    static long SaturatingNonNegativeLong(BigInteger value) {
      if (value <= 0) return 0;
      return value > long.MaxValue ? long.MaxValue : (long)value;
    }

    static long SaturatingMegabytes(BigInteger value) {
      if (value <= 0) return 0;
      var maximum = new BigInteger(long.MaxValue / (1024L * 1024L));
      return value > maximum ? long.MaxValue : (long)value * 1024L * 1024L;
    }

    // Generic bounded captured-process primitive. Domain policy (warm/cold and typed verdict precedence)
    // stays in ConfluxResourceCeilings; this extern owns only OS observation and termination.
    public static void RunProcessBounded(Dafny.ISequence<Rune> cwd, Dafny.ISequence<Rune> cmd,
        Dafny.ISequence<Dafny.ISequence<Rune>> args, BigInteger timeLimitMs, BigInteger memoryLimitMb,
        BigInteger outputLimitMb, out BigInteger code, out Dafny.ISequence<Rune> outp,
        out Dafny.ISequence<Rune> err, out bool timedOut, out bool memoryExceeded,
        out bool outputExceeded, out BigInteger wallMs, out bool memoryPresent,
        out BigInteger peakRssMb, out BigInteger combinedOutputBytes,
        out bool memoryMeasurementFailed, out bool outputMeasurementFailed,
        out bool treeKillIssued, out bool cleanupComplete, out BigInteger cleanupElapsedMs,
        out BigInteger cleanupDeadlineMs, out Dafny.ISequence<Rune> cleanupDiagnostic,
        out bool firstTripPresent, out BigInteger firstTripIndex,
        out BigInteger firstTripDimension, out bool firstTripObservedPresent,
        out BigInteger firstTripObserved, out bool firstTripMeasurementFailure,
        out bool firstTripInFlight, out BigInteger samplesSeen) {
      var psi = new ProcessStartInfo(S(cmd)) {
        RedirectStandardOutput = true, RedirectStandardError = true, UseShellExecute = false,
      };
      var w = S(cwd); if (w.Length > 0) psi.WorkingDirectory = w;
      foreach (var a in args.Elements) psi.ArgumentList.Add(S(a));
      using var stdout = new MemoryStream(); using var stderr = new MemoryStream();
      long outputBytes = 0, peakBytes = 0;
      long timeLimit = SaturatingNonNegativeLong(timeLimitMs);
      long memoryLimitBytes = SaturatingMegabytes(memoryLimitMb);
      long outputLimitBytes = SaturatingMegabytes(outputLimitMb);
      bool observedOutputExceeded = false, observedMemory = false;
      bool observedMemoryFailure = false, observedOutputFailure = false;
      bool killWasIssued = false;
      int killRequested = 0, rootExitObserved = 0, observationCounter = 0;
      int firstDimension = -1, firstIndex = -1, firstObservedPresent = 0;
      int firstMeasurementFailure = 0, firstInFlight = 0;
      long firstObserved = 0, cleanupStartedTimestamp = 0;
      System.Threading.Tasks.Task killTask = null;
      var killTaskGate = new object();
      const long CleanupDeadline = 2000;
      string cleanupFailure = "";
      var observedPids = new HashSet<int>();
      var observedPidsGate = new object();
      var firstTripGate = new object();
      timedOut = false; memoryExceeded = false; outputExceeded = false;
      treeKillIssued = false; cleanupComplete = true;
      memoryMeasurementFailed = false; outputMeasurementFailed = false;
      cleanupElapsedMs = BigInteger.Zero; cleanupDeadlineMs = new BigInteger(CleanupDeadline);
      cleanupDiagnostic = D(""); firstTripPresent = false; firstTripIndex = BigInteger.Zero;
      firstTripDimension = BigInteger.Zero; firstTripObservedPresent = false;
      firstTripObserved = BigInteger.Zero; firstTripMeasurementFailure = false;
      firstTripInFlight = false; samplesSeen = BigInteger.Zero;
      code = new BigInteger(127); outp = D(""); err = D("");
      var sw = Stopwatch.StartNew();
      try {
        using var pr = new Process { StartInfo = psi };
        if (!pr.Start()) throw new InvalidOperationException("process did not start");
        void RequestTreeKill() {
          if (Interlocked.Exchange(ref killRequested, 1) != 0) return;
          killWasIssued = true;
          Interlocked.CompareExchange(ref cleanupStartedTimestamp, Stopwatch.GetTimestamp(), 0);
          int[] pids; lock (observedPidsGate) pids = observedPids.ToArray();
          lock (killTaskGate) {
            killTask = System.Threading.Tasks.Task.Run(() => KillTrackedTree(pr, pids));
          }
        }

        void RecordFirstTrip(int index, int dimension, long? observed, bool measurementFailure) {
          bool shouldKill = false;
          lock (firstTripGate) {
            if (firstDimension >= 0) return;
            firstIndex = index;
            firstDimension = dimension;
            firstObservedPresent = observed.HasValue ? 1 : 0;
            firstObserved = observed.GetValueOrDefault();
            firstMeasurementFailure = measurementFailure ? 1 : 0;
            firstInFlight = Volatile.Read(ref rootExitObserved) == 0 ? 1 : 0;
            shouldKill = firstInFlight != 0;
          }
          if (shouldKill) RequestTreeKill();
        }

        void Retain(MemoryStream sink, byte[] buffer, int count) {
          lock (sink) {
            long cap = outputLimitBytes > 0
              ? (outputLimitBytes > long.MaxValue - 8192L ? long.MaxValue : outputLimitBytes + 8192L)
              : long.MaxValue;
            int retain = (int)Math.Max(0, Math.Min(count, cap - sink.Length));
            if (retain > 0) sink.Write(buffer, 0, retain);
          }
        }

        async System.Threading.Tasks.Task DrainAsync(Stream source, MemoryStream sink) {
          var buffer = new byte[8192];
          try {
            while (true) {
              int read = await source.ReadAsync(buffer, 0, buffer.Length).ConfigureAwait(false);
              if (read <= 0) return;
              Retain(sink, buffer, read);
              long total = Interlocked.Add(ref outputBytes, read);
              int index = Interlocked.Increment(ref observationCounter) - 1;
              if (outputLimitBytes > 0 && total > outputLimitBytes) {
                Volatile.Write(ref observedOutputExceeded, true);
                long observedMb = Math.Max(0, (total + 1024L * 1024L - 1) / (1024L * 1024L));
                RecordFirstTrip(index, 2, observedMb, false);
              }
            }
          } catch {
            observedOutputFailure = true;
            if (outputLimitBytes > 0) {
              int index = Interlocked.Increment(ref observationCounter) - 1;
              RecordFirstTrip(index, 2, null, true);
            }
          }
        }

        var stdoutDrain = DrainAsync(pr.StandardOutput.BaseStream, stdout);
        var stderrDrain = DrainAsync(pr.StandardError.BaseStream, stderr);
        while (true) {
          if (pr.WaitForExit(10)) {
            Volatile.Write(ref rootExitObserved, 1);
            break;
          }
          int index = Interlocked.Increment(ref observationCounter) - 1;
          long elapsed = sw.ElapsedMilliseconds;
          // Inventory the tree before any dimension can request termination. Tree ownership is required for
          // cleanup even when RSS itself is not part of the selected policy.
          if (timeLimit > 0 || memoryLimitBytes > 0 || outputLimitBytes > 0) {
            if (TrySampleTreeRssBytes(pr.Id, out long treeBytes, out HashSet<int> treePids,
                out bool rootAbsent)) {
              lock (observedPidsGate) observedPids.UnionWith(treePids);
              if (memoryLimitBytes > 0) {
                observedMemory = true;
                peakBytes = Math.Max(peakBytes, treeBytes);
                if (peakBytes > memoryLimitBytes) {
                  memoryExceeded = true;
                  long observedMb = Math.Max(0, (peakBytes + 1024L * 1024L - 1) / (1024L * 1024L));
                  RecordFirstTrip(index, 1, observedMb, false);
                }
              }
            } else if (rootAbsent) {
              // The process was alive before the snapshot but completed before ps captured it. That is a
              // completion race, not an unavailable measurement. A command with no successful sample still
              // fails closed in the post-hoc check below.
              Volatile.Write(ref rootExitObserved, 1);
              pr.WaitForExit();
              break;
            } else if (memoryLimitBytes > 0) {
              observedMemoryFailure = true;
              RecordFirstTrip(index, 1, null, true);
            }
          }
          if (timeLimit > 0 && elapsed > timeLimit) {
            timedOut = true;
            RecordFirstTrip(index, 0, elapsed, false);
          }
          outputExceeded = Volatile.Read(ref observedOutputExceeded) ||
            (outputLimitBytes > 0 && Interlocked.Read(ref outputBytes) > outputLimitBytes);
          if (Volatile.Read(ref killRequested) != 0) break;
        }

        if (Volatile.Read(ref killRequested) != 0) {
          bool TreeSettled() {
            int[] pids; lock (observedPidsGate) pids = observedPids.ToArray();
            return TrackedTreeExited(pr, pids);
          }
          while (!TreeSettled()) {
            long elapsed = (long)Stopwatch.GetElapsedTime(
              Volatile.Read(ref cleanupStartedTimestamp)).TotalMilliseconds;
            long remaining = CleanupDeadline - elapsed;
            if (remaining <= 0) break;
            if (!pr.HasExited) pr.WaitForExit((int)Math.Min(50, remaining));
            int[] pids; lock (observedPidsGate) pids = observedPids.ToArray();
            lock (killTaskGate) {
              if (killTask == null || killTask.IsCompleted) {
                killTask = System.Threading.Tasks.Task.Run(() => KillTrackedTree(pr, pids));
              }
            }
          }
          cleanupComplete = TreeSettled();
          if (!cleanupComplete) {
            cleanupFailure = "process-tree cleanup did not settle before monotonic deadline";
          }
          cleanupElapsedMs = new BigInteger(Math.Max(0, Math.Min(CleanupDeadline,
            (long)Stopwatch.GetElapsedTime(Volatile.Read(ref cleanupStartedTimestamp)).TotalMilliseconds)));
        } else {
          cleanupComplete = pr.HasExited;
        }

        bool drainsComplete = false;
        try { drainsComplete = System.Threading.Tasks.Task.WaitAll(new[] { stdoutDrain, stderrDrain }, 2500); }
        catch { drainsComplete = false; }
        if (!drainsComplete) {
          observedOutputFailure = true;
          if (outputLimitBytes > 0) {
            int index = Interlocked.Increment(ref observationCounter) - 1;
            RecordFirstTrip(index, 2, null, true);
          }
        }
        // A process that completed before the first tree sample still has an unavailable required measurement.
        // Root exit is already observed, so this becomes post-hoc Fail and never requests a synthetic kill.
        if (memoryLimitBytes > 0 && !observedMemory) {
          observedMemoryFailure = true;
          int index = Interlocked.Increment(ref observationCounter) - 1;
          RecordFirstTrip(index, 1, null, true);
        }
        outputExceeded = outputExceeded || Volatile.Read(ref observedOutputExceeded);
        code = pr.HasExited ? new BigInteger(pr.ExitCode) : new BigInteger(127);
        lock (stdout) outp = D(Encoding.UTF8.GetString(stdout.ToArray()));
        lock (stderr) err = D(Encoding.UTF8.GetString(stderr.ToArray()));
      } catch (Exception ex) {
        code = new BigInteger(127); outp = D(""); err = D(ex.Message);
        Volatile.Write(ref rootExitObserved, 1);
        // A child that never spawned consumed no child-process memory. Treat that vacuous tree observation as
        // complete so the canonical spawn-failure mapping (127) is not overwritten by a fabricated memory
        // measurement failure. No budget trip occurred because there was no operation to supervise.
        if (memoryLimitBytes > 0) observedMemory = true;
      }
      sw.Stop(); wallMs = new BigInteger(Math.Max(0, sw.ElapsedMilliseconds));
      memoryPresent = observedMemory;
      peakRssMb = new BigInteger(Math.Max(0, (peakBytes + 1024L * 1024L - 1) / (1024L * 1024L)));
      combinedOutputBytes = new BigInteger(Math.Max(0, outputBytes));
      memoryMeasurementFailed = observedMemoryFailure;
      outputMeasurementFailed = observedOutputFailure;
      treeKillIssued = killWasIssued;
      cleanupDiagnostic = D(cleanupFailure);
      firstTripPresent = firstDimension >= 0;
      firstTripIndex = new BigInteger(Math.Max(0, firstIndex));
      firstTripDimension = new BigInteger(Math.Max(0, firstDimension));
      firstTripObservedPresent = firstObservedPresent != 0;
      firstTripObserved = new BigInteger(Math.Max(0, firstObserved));
      firstTripMeasurementFailure = firstMeasurementFailure != 0;
      firstTripInFlight = firstInFlight != 0;
      samplesSeen = new BigInteger(Math.Max(0, observationCounter));
    }

    // ── Warm interactive process (persistent stdio) — the long-lived sibling of RunProcess ────────────────
    // Opaque int handles into a process-local table (mirrors ConfluxSocketAdapter's conn registry). stderr is
    // drained on a background thread so a full stderr pipe never deadlocks the child; stdio is byte-accurate.
    private static readonly object _procLock = new object();
    private static int _nextProc = 1;
    private sealed class ProcessEntry {
      public readonly Process Process;
      public readonly Queue<byte[]> Stdout = new Queue<byte[]>();
      public readonly AutoResetEvent StdoutReady = new AutoResetEvent(false);
      public readonly HashSet<int> ObservedPids = new HashSet<int>();
      private readonly object OutputObservationGate = new object();
      public long CombinedOutputBytes;
      private long ReportedOutputBytes;
      public volatile bool StdoutClosed;
      public volatile bool StderrClosed;
      public ProcessEntry(Process process) { Process = process; }
      public long TakeObservedOutputBytes() {
        lock (OutputObservationGate) {
          long total = Interlocked.Read(ref CombinedOutputBytes);
          long delta = Math.Max(0, total - ReportedOutputBytes);
          ReportedOutputBytes = total;
          return delta;
        }
      }
      public byte[] DrainStdout() {
        lock (Stdout) {
          if (Stdout.Count == 0) return Array.Empty<byte>();
          var bytes = new List<byte>();
          while (Stdout.Count > 0) bytes.AddRange(Stdout.Dequeue());
          return bytes.ToArray();
        }
      }
    }
    private static readonly Dictionary<int, ProcessEntry> _procs = new Dictionary<int, ProcessEntry>();

    public static BigInteger SpawnProcess(Dafny.ISequence<Rune> cwd, Dafny.ISequence<Rune> cmd,
                                          Dafny.ISequence<Dafny.ISequence<Rune>> args) {
      var psi = new ProcessStartInfo(S(cmd)) {
        RedirectStandardInput = true, RedirectStandardOutput = true, RedirectStandardError = true, UseShellExecute = false,
      };
      var w = S(cwd); if (w.Length > 0) psi.WorkingDirectory = w;
      foreach (var a in args.Elements) psi.ArgumentList.Add(S(a));
      try {
        var pr = Process.Start(psi);
        var entry = new ProcessEntry(pr);
        var stderrThread = new Thread(() => {
          try {
            var s = pr.StandardError.BaseStream; var b = new byte[4096]; int n;
            while ((n = s.Read(b, 0, b.Length)) > 0) Interlocked.Add(ref entry.CombinedOutputBytes, n);
          } catch { }
          finally { entry.StderrClosed = true; entry.StdoutReady.Set(); }
        });
        stderrThread.IsBackground = true; stderrThread.Start();
        var stdoutThread = new Thread(() => {
          try {
            var s = pr.StandardOutput.BaseStream; var b = new byte[4096]; int n;
            while ((n = s.Read(b, 0, b.Length)) > 0) {
              Interlocked.Add(ref entry.CombinedOutputBytes, n);
              var chunk = new byte[n]; Array.Copy(b, chunk, n);
              lock (entry.Stdout) entry.Stdout.Enqueue(chunk);
              entry.StdoutReady.Set();
            }
          } catch { }
          finally { entry.StdoutClosed = true; entry.StdoutReady.Set(); }
        });
        stdoutThread.IsBackground = true; stdoutThread.Start();
        lock (_procLock) { int id = _nextProc++; _procs[id] = entry; return new BigInteger(id); }
      } catch { return new BigInteger(-1); }
    }

    public static void WriteStdin(BigInteger proc, Dafny.ISequence<byte> payload) {
      ProcessEntry entry; lock (_procLock) { if (!_procs.TryGetValue((int)proc, out entry)) return; }
      var pr = entry.Process;
      try { var s = pr.StandardInput.BaseStream; var data = payload.CloneAsArray(); s.Write(data, 0, data.Length); s.Flush(); } catch { }
    }

    public static Dafny.ISequence<byte> ReadStdout(BigInteger proc) {
      Dafny.ISequence<byte> chunk; bool available, open;
      ReadStdoutWithin(proc, new BigInteger(-1), out chunk, out available, out open);
      return chunk;
    }

    public static void ReadStdoutWithin(BigInteger proc, BigInteger waitMs, out Dafny.ISequence<byte> chunk,
                                        out bool available, out bool open) {
      ProcessEntry entry; lock (_procLock) { if (!_procs.TryGetValue((int)proc, out entry)) {
        chunk = Dafny.Helpers.SeqFromArray(new byte[0]); available = false; open = false; return;
      }}
      int wait = waitMs < 0 ? Timeout.Infinite : waitMs > int.MaxValue ? int.MaxValue : (int)waitMs;
      while (true) {
        lock (entry.Stdout) {
          if (entry.Stdout.Count > 0) {
            chunk = Dafny.Helpers.SeqFromArray(entry.Stdout.Dequeue()); available = true; open = true; return;
          }
          if (entry.StdoutClosed) {
            chunk = Dafny.Helpers.SeqFromArray(new byte[0]); available = false; open = false; return;
          }
        }
        if (!entry.StdoutReady.WaitOne(wait)) {
          chunk = Dafny.Helpers.SeqFromArray(new byte[0]); available = false; open = true; return;
        }
      }
    }

    public static void ReadProcessOutputWithin(BigInteger proc, BigInteger waitMs,
                                               out Dafny.ISequence<byte> chunk,
                                               out bool available, out bool open,
                                               out BigInteger combinedOutputBytesDelta) {
      ProcessEntry entry; lock (_procLock) { if (!_procs.TryGetValue((int)proc, out entry)) {
        chunk = Dafny.Helpers.SeqFromArray(new byte[0]); available = false; open = false;
        combinedOutputBytesDelta = BigInteger.Zero; return;
      }}
      ReadStdoutWithin(proc, waitMs, out chunk, out available, out open);
      combinedOutputBytesDelta = new BigInteger(entry.TakeObservedOutputBytes());
    }

    public static void ProcessRssMb(BigInteger proc, out bool present, out BigInteger rssMb) {
      ProcessEntry entry; lock (_procLock) { if (!_procs.TryGetValue((int)proc, out entry)) {
        present = false; rssMb = BigInteger.Zero; return;
      }}
      var pr = entry.Process;
      try {
        if (pr.HasExited) { present = false; rssMb = BigInteger.Zero; return; }
        if (!TrySampleTreeRssBytes(pr.Id, out long bytes, out HashSet<int> pids, out _)) {
          present = false; rssMb = BigInteger.Zero; return;
        }
        lock (entry.ObservedPids) entry.ObservedPids.UnionWith(pids);
        present = true; rssMb = new BigInteger((bytes + 1024L * 1024L - 1) / (1024L * 1024L));
      } catch { present = false; rssMb = BigInteger.Zero; }
    }

    public static void ProcessExitStatus(BigInteger proc, out bool handleFound, out bool observed,
        out BigInteger exitCode) {
      ProcessEntry entry; lock (_procLock) { if (!_procs.TryGetValue((int)proc, out entry)) {
        handleFound = false; observed = false; exitCode = BigInteger.Zero; return;
      }}
      handleFound = true;
      try {
        // HasExited can precede the final stdout/stderr callbacks. Do not expose a terminal state until both
        // native streams have reached EOF, otherwise a final poll can undercount trailing output.
        if (!entry.Process.HasExited || !entry.StdoutClosed || !entry.StderrClosed) {
          observed = false; exitCode = BigInteger.Zero; return;
        }
        observed = true; exitCode = new BigInteger(entry.Process.ExitCode);
      } catch {
        observed = false; exitCode = BigInteger.Zero;
      }
    }

    public static void CurrentProcessRssMb(out bool present, out BigInteger rssMb) {
      try {
        using var pr = Process.GetCurrentProcess();
        pr.Refresh(); long bytes = Math.Max(0, pr.WorkingSet64);
        present = true; rssMb = new BigInteger((bytes + 1024L * 1024L - 1) / (1024L * 1024L));
      } catch { present = false; rssMb = BigInteger.Zero; }
    }

    public static BigInteger CurrentProcessId() => new BigInteger(Environment.ProcessId);

    static bool ProcessIsGone(int pid) {
      if (pid <= 0 || pid == Environment.ProcessId) return false;
      try {
        using var process = Process.GetProcessById(pid);
        return process.HasExited;
      } catch (ArgumentException) {
        return true;
      } catch (InvalidOperationException) {
        return true;
      } catch {
        return false;
      }
    }

    static long ElapsedMilliseconds(long startedTimestamp, int deadlineMs) {
      double elapsed = Math.Ceiling(
        Math.Max(0, Stopwatch.GetTimestamp() - startedTimestamp) * 1000.0 / Stopwatch.Frequency);
      return Math.Max(0, Math.Min(deadlineMs, elapsed >= long.MaxValue ? long.MaxValue : (long)elapsed));
    }

    static void TerminateProcessTreeWithinCore(BigInteger proc, int deadlineMs, bool removeOnFailure,
        out bool operationSucceeded, out bool handleFound, out bool treeKillIssued,
        out bool cleanupComplete, out long cleanupElapsedMs, out string diagnostic,
        out byte[] terminalChunk, out bool terminalAvailable, out long combinedOutputBytesDelta,
        out bool exitObserved, out int exitCode) {
      var startedTimestamp = Stopwatch.GetTimestamp();
      ProcessEntry entry;
      int handle = proc < int.MinValue || proc > int.MaxValue ? int.MinValue : (int)proc;
      lock (_procLock) {
        if (!_procs.TryGetValue(handle, out entry)) {
          operationSucceeded = true; handleFound = false; treeKillIssued = false;
          cleanupComplete = false; cleanupElapsedMs = 0; diagnostic = "";
          terminalChunk = Array.Empty<byte>(); terminalAvailable = false;
          combinedOutputBytesDelta = 0; exitObserved = false; exitCode = 0; return;
        }
      }
      operationSucceeded = true;
      handleFound = true;
      treeKillIssued = false;
      cleanupComplete = false;
      cleanupElapsedMs = 0;
      diagnostic = "";
      terminalChunk = Array.Empty<byte>();
      terminalAvailable = false;
      combinedOutputBytesDelta = 0;
      exitObserved = false;
      exitCode = 0;
      var pr = entry.Process;
      try {
        long deadlineTimestamp = DeadlineAfterMilliseconds(deadlineMs);
        var tracked = new HashSet<int>();
        lock (entry.ObservedPids) tracked.UnionWith(entry.ObservedPids);
        tracked.Add(pr.Id);
        if (TrySampleTreeRssBytes(pr.Id, out _, out HashSet<int> boundaryPids, out _,
            SnapshotDeadline(deadlineTimestamp))) {
          tracked.UnionWith(boundaryPids);
          lock (entry.ObservedPids) entry.ObservedPids.UnionWith(boundaryPids);
        }
        System.Threading.Tasks.Task killTask = null;
        while (true) {
          treeKillIssued = true;
          if (killTask == null || killTask.IsCompleted) {
            var killPids = tracked.ToArray();
            killTask = System.Threading.Tasks.Task.Run(() => KillTrackedTree(pr, killPids));
          }
          if (tracked.All(ProcessIsGone) && entry.StdoutClosed && entry.StderrClosed) {
            cleanupComplete = true;
            break;
          }
          int remaining = RemainingDeadlineMilliseconds(deadlineTimestamp);
          if (remaining <= 0) break;
          // Capture descendants that appeared between the boundary sample and the root kill.
          if (TrySampleTreeRssBytes(pr.Id, out _, out HashSet<int> retryPids, out _,
              SnapshotDeadline(deadlineTimestamp))) {
            tracked.UnionWith(retryPids);
            lock (entry.ObservedPids) entry.ObservedPids.UnionWith(retryPids);
          }
          if (tracked.All(ProcessIsGone) && entry.StdoutClosed && entry.StderrClosed) {
            cleanupComplete = true;
            break;
          }
          remaining = RemainingDeadlineMilliseconds(deadlineTimestamp);
          if (remaining <= 0) break;
          Thread.Sleep(Math.Min(20, remaining));
        }
        // Include the exact deadline boundary: the final kill/stream callbacks can settle
        // between the loop's last check and its monotonic deadline observation.
        if (!cleanupComplete && tracked.All(ProcessIsGone) && entry.StdoutClosed && entry.StderrClosed) {
          cleanupComplete = true;
        }
        if (!cleanupComplete) {
          var remaining = tracked.Where(pid => !ProcessIsGone(pid)).OrderBy(pid => pid).ToArray();
          diagnostic = "process-tree termination was not confirmed within " +
            deadlineMs + "ms; remaining pids=" + string.Join(",", remaining) +
            "; stdoutClosed=" + entry.StdoutClosed + "; stderrClosed=" + entry.StderrClosed;
        }
      } catch (Exception error) {
        operationSucceeded = false;
        diagnostic = string.IsNullOrEmpty(error.Message) ? error.GetType().Name : error.Message;
      }
      cleanupElapsedMs = ElapsedMilliseconds(startedTimestamp, deadlineMs);
      terminalChunk = entry.DrainStdout();
      terminalAvailable = terminalChunk.Length > 0;
      combinedOutputBytesDelta = Math.Min(9007199254740991L, entry.TakeObservedOutputBytes());
      try {
        if (pr.HasExited && entry.StdoutClosed && entry.StderrClosed) {
          exitObserved = true;
          exitCode = pr.ExitCode;
        }
      } catch { }
      if (cleanupComplete || removeOnFailure) {
        lock (_procLock) {
          if (_procs.TryGetValue(handle, out var current) && ReferenceEquals(current, entry)) {
            _procs.Remove(handle);
          }
        }
      }
      try { entry.StdoutReady.Set(); } catch { }
      if (cleanupComplete || removeOnFailure) try { pr.Dispose(); } catch { }
      if (!cleanupComplete && string.IsNullOrEmpty(diagnostic)) {
        diagnostic = "process-tree termination could not be confirmed";
      }
    }

    public static void TerminateProcessTreeWithin(BigInteger proc, BigInteger deadlineMs,
        out bool operationSucceeded, out bool handleFound, out bool treeKillIssued,
        out bool cleanupComplete, out BigInteger cleanupElapsedMs,
        out Dafny.ISequence<Rune> diagnostic, out Dafny.ISequence<byte> terminalChunk,
        out bool terminalAvailable, out BigInteger combinedOutputBytesDelta,
        out bool exitObserved, out BigInteger exitCode) {
      int deadline = deadlineMs > int.MaxValue ? int.MaxValue : Math.Max(1, (int)deadlineMs);
      TerminateProcessTreeWithinCore(proc, deadline, false, out operationSucceeded, out handleFound,
        out treeKillIssued, out cleanupComplete, out long elapsed, out string message,
        out byte[] terminalBytes, out terminalAvailable, out long outputDelta,
        out exitObserved, out int observedExitCode);
      cleanupElapsedMs = new BigInteger(elapsed);
      diagnostic = D(message);
      terminalChunk = Dafny.Helpers.SeqFromArray(terminalBytes);
      combinedOutputBytesDelta = new BigInteger(outputDelta);
      exitCode = new BigInteger(observedExitCode);
    }

    public static void CloseProcess(BigInteger proc) {
      TerminateProcessTreeWithinCore(proc, ProcessObservationDeadlineMs, true,
        out _, out _, out _, out _, out _, out _, out _, out _, out _, out _, out _);
    }

    public static Dafny.ISequence<byte> EncodeUtf8(Dafny.ISequence<Rune> text) {
      return Dafny.Helpers.SeqFromArray(Encoding.UTF8.GetBytes(S(text)));
    }

    public static Dafny.ISequence<Rune> DecodeUtf8(Dafny.ISequence<byte> bytes) {
      return D(Encoding.UTF8.GetString(bytes.CloneAsArray()));
    }

    private static readonly Stopwatch _monotonic = Stopwatch.StartNew();
    public static BigInteger MonotonicTimeMs() => new BigInteger(Math.Max(0, _monotonic.ElapsedMilliseconds));
    public static void SetExitCode(BigInteger code) {
      try { Environment.ExitCode = code < int.MinValue ? int.MinValue : code > int.MaxValue ? int.MaxValue : (int)code; }
      catch { Environment.ExitCode = 127; }
    }
    public static void WriteStderr(Dafny.ISequence<Rune> content) {
      try { Console.Error.Write(S(content)); } catch { }
    }

    public static void GetEnv(Dafny.ISequence<Rune> key, out bool present, out Dafny.ISequence<Rune> value) {
      var v = Environment.GetEnvironmentVariable(S(key));
      present = v != null; value = D(v ?? "");
    }

    public static BigInteger StartWatch(Dafny.ISequence<Dafny.ISequence<Rune>> roots) {
      var state = new WatchState();
      foreach (var rootValue in roots.CloneAsArray()) {
        string root = S(rootValue);
        if (!Directory.Exists(root)) continue;
        try {
          var watcher = new FileSystemWatcher(root) {
            IncludeSubdirectories = true,
            NotifyFilter = NotifyFilters.FileName | NotifyFilters.DirectoryName | NotifyFilters.LastWrite,
            EnableRaisingEvents = false,
          };
          FileSystemEventHandler changed = (_, e) => {
            lock (state.Gate) { state.Paths.Enqueue(e.FullPath); Monitor.PulseAll(state.Gate); }
          };
          RenamedEventHandler renamed = (_, e) => {
            lock (state.Gate) { state.Paths.Enqueue(e.FullPath); Monitor.PulseAll(state.Gate); }
          };
          watcher.Changed += changed; watcher.Created += changed; watcher.Deleted += changed; watcher.Renamed += renamed;
          watcher.EnableRaisingEvents = true;
          state.Watchers.Add(watcher);
        } catch { }
      }
      lock (WatchGate) { int id = NextWatch++; Watches[id] = state; return new BigInteger(id); }
    }

    public static void PollWatch(BigInteger watcher, BigInteger waitMs, out bool available, out Dafny.ISequence<Rune> path) {
      WatchState state;
      lock (WatchGate) { Watches.TryGetValue((int)watcher, out state); }
      if (state == null) { available = false; path = D(""); return; }
      lock (state.Gate) {
        if (state.Paths.Count == 0 && waitMs > 0) Monitor.Wait(state.Gate, Math.Max(0, (int)waitMs));
        if (state.Paths.Count > 0) { available = true; path = D(state.Paths.Dequeue()); return; }
      }
      available = false; path = D("");
    }

    public static void CloseWatch(BigInteger watcher) {
      WatchState state;
      lock (WatchGate) { Watches.TryGetValue((int)watcher, out state); Watches.Remove((int)watcher); }
      if (state == null) return;
      foreach (var item in state.Watchers) try { item.Dispose(); } catch { }
      lock (state.Gate) Monitor.PulseAll(state.Gate);
    }
  }
}
