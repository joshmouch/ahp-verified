/*
 * ahp_core.h — C ABI for the formally verified Agent Host Protocol core.
 *
 * The functions declared here are implemented by libahp_core.dylib, a
 * .NET NativeAOT shared library. All protocol behaviour behind them is
 * performed by reducers extracted from a machine-checked Dafny specification.
 * The marshalling shell that adapts those reducers to this C ABI is
 * hand-written and is NOT covered by the proofs. See README.md, "Trust chain".
 *
 * Conventions
 *   - All strings are NUL-terminated UTF-8.
 *   - Functions returning `char*` transfer ownership to the caller, who must
 *     release the result with ahp_string_free(). They return NULL on failure.
 *   - `const char*` returns are owned by the library; do NOT free them.
 *   - On any failure, ahp_last_error() describes the most recent error on the
 *     calling thread. The pointer is valid until the next call on that thread.
 *   - Errors are reported per thread; the handle table is safe to use from
 *     multiple threads.
 *   - JSON emitted by this library is canonical: object keys are written in
 *     sorted order. Round-tripping therefore preserves the JSON *value*, and
 *     is byte-stable only for input already in canonical form. Compare
 *     outputs against outputs, not against hand-written input.
 *
 * Copyright (c) Microsoft Corporation.
 * Copyright (c) 2026 Josh Mouch.
 * Licensed under the MIT License.
 */

#ifndef AHP_CORE_H
#define AHP_CORE_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

/* ---------------------------------------------------------------------
 * Library information and memory management
 * ------------------------------------------------------------------ */

/** Version string of the library. Owned by the library; do not free. */
const char *ahp_version(void);

/**
 * Description of the most recent failure on the calling thread, or NULL if
 * no call on this thread has failed. Owned by the library; do not free.
 * Valid only until the next ahp_* call on this thread.
 */
const char *ahp_last_error(void);

/** Releases a string returned by any ahp_*_json / ahp_action_channel call. */
void ahp_string_free(char *s);

/* ---------------------------------------------------------------------
 * Pure JSON API — stateless; each call is a total function of its inputs
 * ------------------------------------------------------------------ */

/** The initial (default) AHP state, encoded as JSON. Caller frees. */
char *ahp_initial_state_json(void);

/**
 * Applies one action to one state (the core's applyAhp).
 * @param state_json   AHP state, as produced by any state-returning call.
 * @param action_json  A single AHP action object.
 * @return             The resulting state as JSON, or NULL on error.
 */
char *ahp_apply_json(const char *state_json, const char *action_json);

/**
 * Folds a sequence of actions over a state (the core's foldAhp).
 * @param actions_json  A JSON *array* of action objects.
 * @return              The resulting state as JSON, or NULL on error.
 */
char *ahp_fold_json(const char *state_json, const char *actions_json);

/**
 * Codec round trip: encodeAhpAction(decodeAhpAction(action_json)).
 * The core proves this is the identity on well-formed actions, so a caller can
 * compare the result against its input to confirm the core fully understood
 * the wire form it sent.
 * @return  The re-encoded action as JSON, or NULL on error.
 */
char *ahp_action_roundtrip_json(const char *action_json);

/**
 * The channel an action routes to: one of "root", "session", "chat",
 * "terminal", "changeset", "annotations", "resourceWatch", "canvas", or
 * "root/unknown".
 *
 * Note: the core does not reject unrecognised action types — it decodes them
 * to a defined catch-all. "root/unknown" is therefore how a caller detects an
 * action type the core does not model. Caller frees; NULL on error.
 */
char *ahp_action_channel(const char *action_json);

/* ---------------------------------------------------------------------
 * Handle API — keeps decoded state in the library across calls, so a stream
 * of actions does not pay to re-parse and re-encode the whole state each time.
 * ------------------------------------------------------------------ */

/** Opaque state handle. 0 is never a valid handle and signals failure. */
typedef int64_t ahp_state_t;

/** Creates a handle holding the initial state. Returns 0 on failure. */
ahp_state_t ahp_state_new(void);

/** Creates a handle from encoded state JSON. Returns 0 on failure. */
ahp_state_t ahp_state_from_json(const char *state_json);

/**
 * Applies one action to the state behind `handle`, in place.
 * @return 0 on success, -1 on failure (see ahp_last_error()).
 */
int ahp_state_apply(ahp_state_t handle, const char *action_json);

/** Encodes the state behind `handle` as JSON. Caller frees; NULL on error. */
char *ahp_state_to_json(ahp_state_t handle);

/**
 * Structural equality of two states, using the core's own equality.
 * @return 1 if equal, 0 if not, -1 on error.
 */
int ahp_state_equals(ahp_state_t a, ahp_state_t b);

/** Releases a state handle. Safe to call with an already-released handle. */
void ahp_state_free(ahp_state_t handle);

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif /* AHP_CORE_H */
