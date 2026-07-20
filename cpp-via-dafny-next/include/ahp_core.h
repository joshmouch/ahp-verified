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
 *     non-UTF-8 string).
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

void ahp_terminal_free(AhpTerminalState *state);
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
