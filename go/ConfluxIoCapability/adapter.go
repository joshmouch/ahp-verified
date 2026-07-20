// Go realization of the ConfluxIoCapability extern functions.
//
// This is the host-effect boundary: process execution, filesystem reads, clocks
// and filesystem watching. It mirrors the semantics of the C# and JavaScript
// adapters shipped with the same verified core.
//
// NOTE: this file is hand-written host glue at the Dafny {:extern} trust boundary.
// It is NOT part of the formally verified core. See EXTERN.md.
//
// Copyright (c) Microsoft Corporation
// Copyright (c) 2026 Josh Mouch
// SPDX-License-Identifier: MIT

package ConfluxIoCapability

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	_dafny "github.com/joshmouch/ahp-verified/go/dafny"
)

// CompanionStruct_Default___ supplies the extern module's static members.
type CompanionStruct_Default___ struct{}

// Companion_Default___ is the singleton companion the translated core calls into.
var Companion_Default___ = CompanionStruct_Default___{}

var monotonicStart = time.Now()

// ---------- conversions ----------

func text(seq _dafny.Sequence) string { return seq.VerbatimString(false) }

func fromText(s string) _dafny.Sequence { return _dafny.UnicodeSeqOfUtf8Bytes(s) }

func emptyBytes() _dafny.Sequence { return _dafny.SeqOfBytes(nil) }

func texts(seq _dafny.Sequence) []string {
	n := seq.Cardinality()
	out := make([]string, 0, n)
	for i := uint32(0); i < n; i++ {
		if s, ok := seq.Select(i).(_dafny.Sequence); ok {
			out = append(out, text(s))
		}
	}
	return out
}

func megabytes(b int64) int {
	if b <= 0 {
		return 0
	}
	const mb = 1024 * 1024
	return int((b + mb - 1) / mb)
}

// ---------- process registry ----------

type managedProc struct {
	cmd    *exec.Cmd
	stdout *os.File
	buf    bytes.Buffer
	mu     sync.Mutex
	done   chan struct{}

	exited   bool
	exitCode int
	consumed int64
}

var (
	procMu   sync.Mutex
	procs          = map[int64]*managedProc{}
	nextProc int64 = 1
)

func lookup(h _dafny.Int) *managedProc {
	procMu.Lock()
	defer procMu.Unlock()
	return procs[h.Int64()]
}

// SpawnProcess starts a child process and returns a handle, or -1 on failure.
func (_static *CompanionStruct_Default___) SpawnProcess(cwd, cmdName, args _dafny.Sequence) _dafny.Int {
	c := exec.Command(text(cmdName), texts(args)...)
	if dir := text(cwd); dir != "" {
		c.Dir = dir
	}
	c.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	pr, pw, err := os.Pipe()
	if err != nil {
		return _dafny.IntOfInt64(-1)
	}
	c.Stdout, c.Stderr = pw, pw
	if err := c.Start(); err != nil {
		pr.Close()
		pw.Close()
		return _dafny.IntOfInt64(-1)
	}
	pw.Close()

	p := &managedProc{cmd: c, stdout: pr, done: make(chan struct{})}
	go func() {
		chunk := make([]byte, 32*1024)
		for {
			n, err := pr.Read(chunk)
			if n > 0 {
				p.mu.Lock()
				p.buf.Write(chunk[:n])
				p.mu.Unlock()
			}
			if err != nil {
				break
			}
		}
		code := 0
		if err := c.Wait(); err != nil {
			if ee, ok := err.(*exec.ExitError); ok {
				code = ee.ExitCode()
			} else {
				code = 127
			}
		}
		p.mu.Lock()
		p.exited, p.exitCode = true, code
		p.mu.Unlock()
		close(p.done)
	}()

	procMu.Lock()
	h := nextProc
	nextProc++
	procs[h] = p
	procMu.Unlock()
	return _dafny.IntOfInt64(h)
}

// CloseProcess releases a process handle, terminating it if still running.
func (_static *CompanionStruct_Default___) CloseProcess(proc _dafny.Int) {
	p := lookup(proc)
	if p == nil {
		return
	}
	killTree(p)
	procMu.Lock()
	delete(procs, proc.Int64())
	procMu.Unlock()
}

func killTree(p *managedProc) {
	if p.cmd.Process == nil {
		return
	}
	if pgid, err := syscall.Getpgid(p.cmd.Process.Pid); err == nil {
		_ = syscall.Kill(-pgid, syscall.SIGKILL)
		return
	}
	_ = p.cmd.Process.Kill()
}

// rssMbOf reports resident set size in MB for a pid, best effort.
func rssMbOf(pid int) (bool, int) {
	out, err := exec.Command("ps", "-o", "rss=", "-p", strconv.Itoa(pid)).Output()
	if err != nil {
		return false, 0
	}
	kb, err := strconv.ParseInt(strings.TrimSpace(string(out)), 10, 64)
	if err != nil {
		return false, 0
	}
	return true, megabytes(kb * 1024)
}

// ---------- clock / memory ----------

// MonotonicTimeMs returns milliseconds elapsed since process start.
func (_static *CompanionStruct_Default___) MonotonicTimeMs() _dafny.Int {
	return _dafny.IntOfInt64(int64(time.Since(monotonicStart) / time.Millisecond))
}

// CurrentProcessRssMb reports this process's resident set size in MB.
func (_static *CompanionStruct_Default___) CurrentProcessRssMb() (bool, _dafny.Int) {
	if ok, mb := rssMbOf(os.Getpid()); ok {
		return true, _dafny.IntOf(mb)
	}
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	return true, _dafny.IntOf(megabytes(int64(ms.Sys)))
}

// ProcessRssMb reports a managed child's resident set size in MB.
func (_static *CompanionStruct_Default___) ProcessRssMb(proc _dafny.Int) (bool, _dafny.Int) {
	p := lookup(proc)
	if p == nil || p.cmd.Process == nil {
		return false, _dafny.Zero
	}
	ok, mb := rssMbOf(p.cmd.Process.Pid)
	if !ok {
		return false, _dafny.Zero
	}
	return true, _dafny.IntOf(mb)
}

// ProcessExitStatus reports whether the handle is known, whether it has exited,
// and its exit code.
func (_static *CompanionStruct_Default___) ProcessExitStatus(proc _dafny.Int) (bool, bool, _dafny.Int) {
	p := lookup(proc)
	if p == nil {
		return false, false, _dafny.Zero
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	return true, p.exited, _dafny.IntOf(p.exitCode)
}

// ---------- filesystem ----------

// ReadFile reads a UTF-8 text file.
func (_static *CompanionStruct_Default___) ReadFile(path _dafny.Sequence) (bool, _dafny.Sequence) {
	b, err := os.ReadFile(text(path))
	if err != nil {
		return false, fromText("")
	}
	return true, fromText(string(b))
}

// ReadFileBytes reads a file as raw bytes.
func (_static *CompanionStruct_Default___) ReadFileBytes(path _dafny.Sequence) (bool, _dafny.Sequence) {
	b, err := os.ReadFile(text(path))
	if err != nil {
		return false, emptyBytes()
	}
	return true, _dafny.SeqOfBytes(b)
}

// ---------- process output ----------

// ReadProcessOutputWithin drains buffered child output, waiting up to waitMs.
// It returns the chunk, whether data was available, whether the process is still
// open, and the number of new bytes observed.
func (_static *CompanionStruct_Default___) ReadProcessOutputWithin(proc _dafny.Int, waitMs _dafny.Int) (_dafny.Sequence, bool, bool, _dafny.Int) {
	p := lookup(proc)
	if p == nil {
		return emptyBytes(), false, false, _dafny.Zero
	}
	deadline := time.Now().Add(time.Duration(max64(0, waitMs.Int64())) * time.Millisecond)
	for {
		p.mu.Lock()
		n := p.buf.Len()
		exited := p.exited
		if n > 0 {
			chunk := make([]byte, n)
			copy(chunk, p.buf.Next(n))
			p.consumed += int64(n)
			p.mu.Unlock()
			return _dafny.SeqOfBytes(chunk), true, !exited, _dafny.IntOfInt64(int64(n))
		}
		p.mu.Unlock()
		if exited || !time.Now().Before(deadline) {
			return emptyBytes(), false, !exited, _dafny.Zero
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// TerminateProcessTreeWithin kills a child's process group and waits for the
// exit to be observed within deadlineMs.
func (_static *CompanionStruct_Default___) TerminateProcessTreeWithin(proc _dafny.Int, deadlineMs _dafny.Int) (
	bool, bool, bool, bool, _dafny.Int, _dafny.Sequence, _dafny.Sequence, bool, _dafny.Int, bool, _dafny.Int) {

	p := lookup(proc)
	if p == nil {
		return false, false, false, false, _dafny.Zero,
			fromText("handle not found"), emptyBytes(), false, _dafny.Zero, false, _dafny.Zero
	}
	start := time.Now()
	killTree(p)
	limit := time.Duration(max64(1, deadlineMs.Int64())) * time.Millisecond

	cleanupComplete := false
	select {
	case <-p.done:
		cleanupComplete = true
	case <-time.After(limit):
	}
	elapsed := int64(time.Since(start) / time.Millisecond)

	p.mu.Lock()
	n := p.buf.Len()
	chunk := make([]byte, n)
	copy(chunk, p.buf.Next(n))
	p.consumed += int64(n)
	exited, code := p.exited, p.exitCode
	p.mu.Unlock()

	return true, true, true, cleanupComplete, _dafny.IntOfInt64(elapsed),
		fromText(""), _dafny.SeqOfBytes(chunk), n > 0, _dafny.IntOfInt64(int64(n)),
		exited, _dafny.IntOf(code)
}

// ---------- synchronous process execution ----------

// RunProcess runs a command to completion and returns exit code, stdout, stderr.
func (_static *CompanionStruct_Default___) RunProcess(cwd, cmdName, args _dafny.Sequence) (_dafny.Int, _dafny.Sequence, _dafny.Sequence) {
	c := exec.Command(text(cmdName), texts(args)...)
	if dir := text(cwd); dir != "" {
		c.Dir = dir
	}
	var stdout, stderr bytes.Buffer
	c.Stdout, c.Stderr = &stdout, &stderr

	code := 0
	if err := c.Run(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			code = ee.ExitCode()
		} else {
			return _dafny.IntOf(127), fromText(""), fromText(err.Error())
		}
	}
	return _dafny.IntOf(code), fromText(stdout.String()), fromText(stderr.String())
}

// RunProcessBounded runs a command under wall-clock, memory and output ceilings,
// reporting which ceiling (if any) tripped first.
func (_static *CompanionStruct_Default___) RunProcessBounded(
	cwd, cmdName, args _dafny.Sequence, timeLimitMs, memoryLimitMb, outputLimitMb _dafny.Int) (
	_dafny.Int, _dafny.Sequence, _dafny.Sequence, bool, bool, bool, _dafny.Int, bool, _dafny.Int,
	_dafny.Int, bool, bool, bool, bool, _dafny.Int, _dafny.Int, _dafny.Sequence, bool, _dafny.Int,
	_dafny.Int, bool, _dafny.Int, bool, bool, _dafny.Int) {

	timeLimit := max64(0, timeLimitMs.Int64())
	memLimit := max64(0, memoryLimitMb.Int64())
	outLimit := max64(0, outputLimitMb.Int64()) * 1024 * 1024

	start := time.Now()
	c := exec.Command(text(cmdName), texts(args)...)
	if dir := text(cwd); dir != "" {
		c.Dir = dir
	}
	c.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	pr, pw, err := os.Pipe()
	if err != nil {
		return failedBounded(err.Error())
	}
	c.Stdout, c.Stderr = pw, pw
	if err := c.Start(); err != nil {
		pr.Close()
		pw.Close()
		return failedBounded(err.Error())
	}
	pw.Close()

	var mu sync.Mutex
	var out bytes.Buffer
	var totalBytes int64
	outputExceeded := false

	readDone := make(chan struct{})
	go func() {
		defer close(readDone)
		chunk := make([]byte, 32*1024)
		for {
			n, err := pr.Read(chunk)
			if n > 0 {
				mu.Lock()
				totalBytes += int64(n)
				if outLimit > 0 && totalBytes > outLimit {
					outputExceeded = true
				} else {
					out.Write(chunk[:n])
				}
				mu.Unlock()
			}
			if err != nil {
				return
			}
		}
	}()

	waitDone := make(chan int, 1)
	go func() {
		code := 0
		if err := c.Wait(); err != nil {
			if ee, ok := err.(*exec.ExitError); ok {
				code = ee.ExitCode()
			} else {
				code = 127
			}
		}
		waitDone <- code
	}()

	peakRss, memPresent, memExceeded, timedOut := 0, false, false, false
	samples := int64(0)
	firstTripDim, firstTripObserved := int64(0), int64(0)
	ticker := time.NewTicker(25 * time.Millisecond)
	defer ticker.Stop()

	exitCode := 0
loop:
	for {
		select {
		case exitCode = <-waitDone:
			break loop
		case <-ticker.C:
			samples++
			if ok, mb := rssMbOf(c.Process.Pid); ok {
				memPresent = true
				if mb > peakRss {
					peakRss = mb
				}
				if memLimit > 0 && int64(mb) > memLimit && !memExceeded {
					memExceeded = true
					firstTripDim, firstTripObserved = 2, int64(mb)
					killProcessGroup(c)
				}
			}
			mu.Lock()
			tripped := outputExceeded
			observed := totalBytes
			mu.Unlock()
			if tripped && firstTripDim == 0 {
				firstTripDim, firstTripObserved = 3, observed
				killProcessGroup(c)
			}
			if timeLimit > 0 && time.Since(start) > time.Duration(timeLimit)*time.Millisecond && !timedOut {
				timedOut = true
				if firstTripDim == 0 {
					firstTripDim, firstTripObserved = 1, int64(time.Since(start)/time.Millisecond)
				}
				killProcessGroup(c)
			}
		}
	}
	<-readDone
	wallMs := int64(time.Since(start) / time.Millisecond)

	mu.Lock()
	body := out.String()
	combined := totalBytes
	tripped := outputExceeded
	mu.Unlock()

	// Once the combined-output ceiling fires, the byte observation and kill
	// verdict are authoritative; do not materialize the over-limit capture.
	if tripped {
		body = ""
	}

	return _dafny.IntOf(exitCode), fromText(body), fromText(""),
		timedOut, memExceeded, tripped,
		_dafny.IntOfInt64(wallMs), memPresent, _dafny.IntOf(peakRss),
		_dafny.IntOfInt64(combined), !memPresent, false,
		timedOut || memExceeded || tripped, true,
		_dafny.Zero, _dafny.IntOfInt64(2000), fromText(""),
		firstTripDim != 0, _dafny.Zero, _dafny.IntOfInt64(firstTripDim),
		firstTripDim != 0, _dafny.IntOfInt64(firstTripObserved),
		false, false, _dafny.IntOfInt64(samples)
}

func killProcessGroup(c *exec.Cmd) {
	if c.Process == nil {
		return
	}
	if pgid, err := syscall.Getpgid(c.Process.Pid); err == nil {
		_ = syscall.Kill(-pgid, syscall.SIGKILL)
		return
	}
	_ = c.Process.Kill()
}

func failedBounded(msg string) (
	_dafny.Int, _dafny.Sequence, _dafny.Sequence, bool, bool, bool, _dafny.Int, bool, _dafny.Int,
	_dafny.Int, bool, bool, bool, bool, _dafny.Int, _dafny.Int, _dafny.Sequence, bool, _dafny.Int,
	_dafny.Int, bool, _dafny.Int, bool, bool, _dafny.Int) {

	return _dafny.IntOf(127), fromText(""), fromText(msg),
		false, false, false, _dafny.Zero, false, _dafny.Zero,
		_dafny.Zero, true, false, false, true,
		_dafny.Zero, _dafny.IntOfInt64(2000), fromText(""),
		true, _dafny.Zero, _dafny.IntOfInt64(1),
		false, _dafny.Zero, true, false, _dafny.IntOfInt64(1)
}

// ---------- filesystem watching ----------

type watcher struct {
	roots []string
	seen  map[string]time.Time
	mu    sync.Mutex
}

var (
	watchMu   sync.Mutex
	watchers        = map[int64]*watcher{}
	nextWatch int64 = 1
)

// StartWatch begins watching the given roots for changes.
func (_static *CompanionStruct_Default___) StartWatch(roots _dafny.Sequence) _dafny.Int {
	w := &watcher{roots: texts(roots), seen: map[string]time.Time{}}
	w.scan(false)
	watchMu.Lock()
	h := nextWatch
	nextWatch++
	watchers[h] = w
	watchMu.Unlock()
	return _dafny.IntOfInt64(h)
}

// CloseWatch releases a watch handle.
func (_static *CompanionStruct_Default___) CloseWatch(w _dafny.Int) {
	watchMu.Lock()
	delete(watchers, w.Int64())
	watchMu.Unlock()
}

// scan walks the roots, returning the first path whose mtime changed.
func (w *watcher) scan(report bool) string {
	w.mu.Lock()
	defer w.mu.Unlock()
	changed := ""
	for _, root := range w.roots {
		_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil || info == nil || info.IsDir() {
				return nil
			}
			mt := info.ModTime()
			if prev, ok := w.seen[path]; !ok || !prev.Equal(mt) {
				w.seen[path] = mt
				if report && changed == "" {
					changed = path
				}
			}
			return nil
		})
	}
	return changed
}

// PollWatch waits up to waitMs for a change, returning availability and the path.
func (_static *CompanionStruct_Default___) PollWatch(w _dafny.Int, waitMs _dafny.Int) (bool, _dafny.Sequence) {
	watchMu.Lock()
	watch := watchers[w.Int64()]
	watchMu.Unlock()
	if watch == nil {
		return false, fromText("")
	}
	deadline := time.Now().Add(time.Duration(max64(0, waitMs.Int64())) * time.Millisecond)
	for {
		if path := watch.scan(true); path != "" {
			return true, fromText(path)
		}
		if !time.Now().Before(deadline) {
			return false, fromText("")
		}
		time.Sleep(20 * time.Millisecond)
	}
}
