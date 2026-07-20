// Go realization of the ConfluxHttpCapability extern functions.
//
// The verified core addresses HTTP requests and servers by opaque integer handle.
// This adapter keeps a registry mapping those handles onto Go http.ResponseWriter
// values, so a Go host can serve the verified request/response logic directly.
//
// A host wires itself up by calling Register inside its own http.Handler; the
// extern calls then resolve to the right writer. Calls against
// an unregistered handle are dropped rather than panicking, which matches the
// C# and JavaScript adapters' behavior for stale handles.
//
// NOTE: this file is hand-written host glue at the Dafny {:extern} trust boundary.
// It is NOT part of the formally verified core. See EXTERN.md.
//
// Copyright (c) Microsoft Corporation
// Copyright (c) 2026 Josh Mouch
// SPDX-License-Identifier: MIT

package ConfluxHttpCapability

import (
	"net/http"
	"strings"
	"sync"

	_dafny "github.com/joshmouch/ahp-go/dafny"
)

// CompanionStruct_Default___ supplies the extern module's static members.
type CompanionStruct_Default___ struct{}

// Companion_Default___ is the singleton companion the translated core calls into.
var Companion_Default___ = CompanionStruct_Default___{}

type pending struct {
	w         http.ResponseWriter
	done      chan struct{}
	once      sync.Once
	streaming bool
}

var (
	mu       sync.Mutex
	requests       = map[int64]*pending{}
	servers        = map[int64][]*pending{}
	nextReq  int64 = 1
)

// Register associates an in-flight response writer with a fresh request handle.
// The returned wait function blocks until the verified core has responded, and
// the release function drops the registration.
func Register(serverHandle int64, w http.ResponseWriter) (handle int64, wait func(), release func()) {
	p := &pending{w: w, done: make(chan struct{})}
	mu.Lock()
	handle = nextReq
	nextReq++
	requests[handle] = p
	servers[serverHandle] = append(servers[serverHandle], p)
	mu.Unlock()

	return handle, func() { <-p.done }, func() {
		mu.Lock()
		delete(requests, handle)
		list := servers[serverHandle][:0]
		for _, q := range servers[serverHandle] {
			if q != p {
				list = append(list, q)
			}
		}
		servers[serverHandle] = list
		mu.Unlock()
		p.finish()
	}
}

func (p *pending) finish() { p.once.Do(func() { close(p.done) }) }

func lookup(h _dafny.Int) *pending {
	mu.Lock()
	defer mu.Unlock()
	return requests[h.Int64()]
}

func text(seq _dafny.Sequence) string { return seq.VerbatimString(false) }

// writeHeaders applies "Name: value" header lines from a seq<string>.
func writeHeaders(w http.ResponseWriter, headers _dafny.Sequence) {
	n := headers.Cardinality()
	for i := uint32(0); i < n; i++ {
		s, ok := headers.Select(i).(_dafny.Sequence)
		if !ok {
			continue
		}
		name, value, found := strings.Cut(text(s), ":")
		if !found {
			continue
		}
		w.Header().Add(strings.TrimSpace(name), strings.TrimSpace(value))
	}
}

// Respond writes a complete response and completes the request.
func (_static *CompanionStruct_Default___) Respond(request _dafny.Int, status _dafny.Int,
	contentType _dafny.Sequence, headers _dafny.Sequence, body _dafny.Sequence) {

	p := lookup(request)
	if p == nil {
		return
	}
	p.w.Header().Set("Content-Type", text(contentType))
	writeHeaders(p.w, headers)
	p.w.WriteHeader(int(status.Int64()))
	_, _ = p.w.Write([]byte(text(body)))
	p.finish()
}

// BeginStream writes response headers plus a first chunk and leaves the response
// open for subsequent Publish calls.
func (_static *CompanionStruct_Default___) BeginStream(request _dafny.Int,
	contentType _dafny.Sequence, headers _dafny.Sequence, body _dafny.Sequence) {

	p := lookup(request)
	if p == nil {
		return
	}
	p.streaming = true
	p.w.Header().Set("Content-Type", text(contentType))
	writeHeaders(p.w, headers)
	p.w.WriteHeader(http.StatusOK)
	_, _ = p.w.Write([]byte(text(body)))
	if f, ok := p.w.(http.Flusher); ok {
		f.Flush()
	}
}

// Publish broadcasts a chunk to every open stream on a server handle.
func (_static *CompanionStruct_Default___) Publish(server _dafny.Int, body _dafny.Sequence) {
	mu.Lock()
	targets := append([]*pending(nil), servers[server.Int64()]...)
	mu.Unlock()

	payload := []byte(text(body))
	for _, p := range targets {
		if !p.streaming {
			continue
		}
		if _, err := p.w.Write(payload); err != nil {
			p.finish()
			continue
		}
		if f, ok := p.w.(http.Flusher); ok {
			f.Flush()
		}
	}
}
