/*
 * ahp_core.h -- C ABI over the formally-verified AHP core (Terminal channel).
 *
 * The state transitions behind this header are MACHINE-GENERATED from the
 * verified Dafny core (`dafny translate rs`), then compiled to a native library.
 * The binding layer that marshals across this boundary is hand-written and is
 * NOT verified. See README.md, "Trust chain".
 *
 * OWNERSHIP
 *   - Every AhpTerminalState* returned must be released with ahp_terminal_free.
 *   - Every char* returned must be released with ahp_string_free.
 *   - Transitions are PURE: they return a new state and never mutate or consume
 *     their input, which remains valid and independently owned.
 *   - Functions returning a pointer return NULL on invalid input (NULL handle,
 *     released handle, non-UTF-8 string).
 *
 * HANDLES ARE OPAQUE TOKENS, NOT POINTERS
 *   An AhpTerminalState* is a token that names an entry in a table inside the
 *   library. It is NOT a dereferenceable address: do not dereference it, do not
 *   do arithmetic on it, do not synthesize one. Tokens are never reused, so a
 *   released token can never come to name a different live state. Passing a
 *   released, never-issued, or NULL token is DEFINED behaviour, not a crash:
 *     - readouts and transitions return NULL (or 0 for the int-returning ones),
 *     - ahp_terminal_free is a no-op, so double-free is harmless.
 *   ahp_terminal_valid() reports whether a token is live.
 *
 * THREAD SAFETY
 *   Every function here is safe to call concurrently from any thread. Calls
 *   into the verified core are SERIALIZED internally by a process-wide lock:
 *   the Dafny-extracted Rust uses non-atomic reference counts and self-mutating
 *   lazy sequences, so two threads touching one state -- even two readers -- is
 *   a data race. The lock makes the C surface sound at the cost of throughput.
 *   Independent states do NOT proceed in parallel. Callers that need real
 *   parallelism should shard across processes.
 *
 * TEXT
 *   All strings are NUL-terminated UTF-8. Invalid UTF-8 is REFUSED, never
 *   repaired: such a call returns NULL. This binding will not substitute
 *   U+FFFD for malformed bytes, because that silently merges distinct inputs
 *   into one verified state.
 *
 * Copyright (c) Microsoft Corporation
 * Copyright (c) 2026 Josh Mouch
 * SPDX-License-Identifier: MIT
 */

#ifndef AHP_CORE_H
#define AHP_CORE_H

#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Opaque handle to a verified Terminal channel state. */
typedef struct AhpTerminalState AhpTerminalState;

/* ---- lifecycle ---------------------------------------------------------- */

/* Initial state (Dafny `Terminal.T0()`). */
AhpTerminalState *ahp_terminal_initial(void);

/*
 * Releases a handle. Releasing a handle that was never issued, or that was
 * already released, is a defined no-op.
 */
void ahp_terminal_free(AhpTerminalState *state);

/*
 * 1 when the handle names a live state, 0 otherwise (NULL, never issued, or
 * already released). There is no other way to test a handle from C.
 */
int32_t ahp_terminal_valid(const AhpTerminalState *state);

void ahp_string_free(char *s);

/* ---- transitions (each is a call into the verified reducer `apply1`) ----- */

AhpTerminalState *ahp_terminal_cwd_changed(const AhpTerminalState *state,
                                           const char *cwd);
AhpTerminalState *ahp_terminal_title_changed(const AhpTerminalState *state,
                                             const char *title);
AhpTerminalState *ahp_terminal_resized(const AhpTerminalState *state,
                                       int64_t cols, int64_t rows);
AhpTerminalState *ahp_terminal_exited(const AhpTerminalState *state,
                                      int64_t code);
AhpTerminalState *ahp_terminal_data(const AhpTerminalState *state,
                                    const char *data);
AhpTerminalState *ahp_terminal_cleared(const AhpTerminalState *state);

/*
 * Fold `count` data actions through the core's proven kernel fold
 * (`Terminal.fold`, itself defined over `ConfluxContract.Fold`) in a single
 * call, rather than looping apply1 on the caller's side.
 *
 * Returns NULL without allocating when `count` exceeds 1048576 (2^20), when
 * `items` is NULL and `count` is non-zero, or when any element is NULL or not
 * UTF-8. The cap exists because the true length of a C array cannot be checked
 * against the array: it refuses counts that cannot be a real array -- notably a
 * size_t underflow from `v.size() - 1` on an empty vector -- before any
 * allocation is sized from an untrusted number. A `count` that is merely wrong
 * but plausible is still the caller's responsibility.
 */
AhpTerminalState *ahp_terminal_fold_data(const AhpTerminalState *state,
                                         const char *const *items,
                                         size_t count);

/* ---- readouts ----------------------------------------------------------- */

/* Caller frees with ahp_string_free. */
char *ahp_terminal_title(const AhpTerminalState *state);

/* Caller frees with ahp_string_free. NULL when the reducer holds None. */
char *ahp_terminal_cwd(const AhpTerminalState *state);

/* Returns 1 and writes *out when present; returns 0 when the reducer holds None. */
int32_t ahp_terminal_exit_code(const AhpTerminalState *state, int64_t *out);
int32_t ahp_terminal_size(const AhpTerminalState *state, int64_t *cols,
                          int64_t *rows);

/* Number of content parts accumulated by the reducer. */
size_t ahp_terminal_content_len(const AhpTerminalState *state);

/* ---- full conformance corpus -------------------------------------------- */

/*
 * Replay the verified core's OWN conformance corpus (Dafny `ClientMain.Main`):
 * every channel's fixture set through the reducers, printing a per-channel
 * board to stdout. This exercises all eight channels, not just the Terminal
 * transitions above.
 *
 * Returns 0. A failed corpus expectation ABORTS the process rather than
 * returning non-zero -- this library is built with panic=abort because
 * unwinding across a C ABI boundary is undefined behaviour. Returning at all is
 * therefore the pass signal.
 *
 * Call fflush(stdout) first if you interleave this with C++ stream output; the
 * library writes through Rust's stdout, which buffers independently of std::cout.
 */
int32_t ahp_run_corpus(void);

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif /* AHP_CORE_H */
