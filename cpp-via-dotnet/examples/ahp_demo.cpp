// ahp_demo.cpp — C++ consumer of the formally verified AHP core.
//
// Exercises the verified reducers through the C ABI: decode, apply, fold,
// codec round trip, channel routing, and the handle-based streaming API.
// Every protocol result printed here was computed by the Dafny-extracted
// reducers inside libahp_core.dylib.
//
// Copyright (c) Microsoft Corporation.
// Copyright (c) 2026 Josh Mouch.
// Licensed under the MIT License.

#include "ahp_core.h"

#include <cstdio>
#include <cstring>
#include <memory>
#include <string>

namespace {

int g_failures = 0;

/// RAII wrapper for strings owned by the library.
struct AhpString {
    char *p{nullptr};
    explicit AhpString(char *raw) : p(raw) {}
    ~AhpString() { ahp_string_free(p); }
    AhpString(const AhpString &) = delete;
    AhpString &operator=(const AhpString &) = delete;
    explicit operator bool() const { return p != nullptr; }
    std::string str() const { return p ? std::string(p) : std::string(); }
};

/// RAII wrapper for state handles.
struct AhpState {
    ahp_state_t h{0};
    explicit AhpState(ahp_state_t handle) : h(handle) {}
    ~AhpState() { ahp_state_free(h); }
    AhpState(const AhpState &) = delete;
    AhpState &operator=(const AhpState &) = delete;
};

void check(const char *label, bool ok, const std::string &detail) {
    if (ok) {
        std::printf("  PASS  %-46s (%s)\n", label, detail.c_str());
    } else {
        ++g_failures;
        std::printf("  FAIL  %-46s (%s)\n", label, detail.c_str());
    }
}

/// True if `haystack` contains `needle`.
bool has(const std::string &haystack, const char *needle) {
    return haystack.find(needle) != std::string::npos;
}

} // namespace

int main() {
    std::printf("AHP verified core — C++ consumer demo\n");
    std::printf("=====================================\n");
    std::printf("library: %s\n\n", ahp_version());

    // -----------------------------------------------------------------
    // 1. Apply a single terminal/resized action.
    // -----------------------------------------------------------------
    std::printf("[1] ahp_apply_json — decode + applyAhp\n");

    AhpString initial(ahp_initial_state_json());
    check("initial state encodes", static_cast<bool>(initial),
          "len=" + std::to_string(initial.str().size()));

    const char *resized =
        R"({"type":"terminal/resized","cols":120,"rows":40})";

    AhpString applied(ahp_apply_json(initial.p, resized));
    check("apply returns a state", static_cast<bool>(applied),
          applied ? "ok" : std::string(ahp_last_error() ? ahp_last_error() : "?"));

    if (applied) {
        const std::string s = applied.str();
        check("reduced state carries cols 120", has(s, "\"cols\":120"),
              has(s, "\"cols\":120") ? "found \"cols\":120" : "missing");
        check("reduced state carries rows 40", has(s, "\"rows\":40"),
              has(s, "\"rows\":40") ? "found \"rows\":40" : "missing");
    }

    // -----------------------------------------------------------------
    // 2. Malformed input must fail loudly, not silently reduce to nothing.
    // -----------------------------------------------------------------
    std::printf("\n[2] malformed JSON is rejected\n");

    AhpString bad(ahp_apply_json(initial.p, "{not json"));
    check("bad action JSON returns NULL", !bad,
          bad ? "unexpectedly succeeded"
              : std::string(ahp_last_error() ? ahp_last_error() : "(no error set)"));

    // -----------------------------------------------------------------
    // 3. Codec round trip — the core's proven identity property.
    // -----------------------------------------------------------------
    std::printf("\n[3] ahp_action_roundtrip_json — codec identity\n");

    // The core's JSON writer emits object keys in sorted order, so the round
    // trip is the identity on the JSON *value*, not on arbitrary key order.
    // Feeding an already-canonical action therefore must return it byte for
    // byte -- that is a true identity check.
    const char *canonical =
        R"({"cols":120,"rows":40,"type":"terminal/resized"})";
    AhpString rt(ahp_action_roundtrip_json(canonical));
    check("canonical action round-trips byte-identically",
          rt && rt.str() == canonical, rt ? rt.str() : "NULL");

    // And for an action written with keys in a different order, the round trip
    // must land on that same canonical form, and stay there.
    AhpString rtA(ahp_action_roundtrip_json(resized));
    AhpString rtB(rtA ? ahp_action_roundtrip_json(rtA.p) : nullptr);
    check("key order is canonicalised", rtA && rtA.str() == canonical,
          rtA ? rtA.str() : "NULL");
    check("round trip is idempotent", rtA && rtB && rtA.str() == rtB.str(),
          rtB ? rtB.str() : "NULL");

    const char *title = R"({"title":"build","type":"terminal/titleChanged"})";
    AhpString rt2(ahp_action_roundtrip_json(title));
    check("terminal/titleChanged round-trips", rt2 && rt2.str() == title,
          rt2 ? rt2.str() : "NULL");

    // -----------------------------------------------------------------
    // 4. Channel routing, including the unknown-action catch-all.
    // -----------------------------------------------------------------
    std::printf("\n[4] ahp_action_channel — dispatch routing\n");

    struct { const char *json; const char *want; } routes[] = {
        {R"({"type":"terminal/resized","cols":80,"rows":24})", "terminal"},
        {R"({"type":"chat/messageAdded","id":"m1","role":"user"})", "chat"},
        {R"({"type":"canvas/cleared"})", "canvas"},
        {R"({"type":"totally/madeUp"})", "root/unknown"},
    };

    for (const auto &r : routes) {
        AhpString ch(ahp_action_channel(r.json));
        const bool ok = ch && ch.str() == r.want;
        check(r.want, ok, ch ? ("routed to " + ch.str()) : "NULL");
    }

    // -----------------------------------------------------------------
    // 5. foldAhp over a multi-action sequence.
    // -----------------------------------------------------------------
    std::printf("\n[5] ahp_fold_json — foldAhp over a sequence\n");

    const char *actions = R"([
        {"type":"terminal/resized","cols":120,"rows":40},
        {"type":"terminal/titleChanged","title":"build"},
        {"type":"terminal/cwdChanged","cwd":"/src"},
        {"type":"terminal/data","data":"hello"}
    ])";

    AhpString folded(ahp_fold_json(initial.p, actions));
    check("fold returns a state", static_cast<bool>(folded),
          folded ? "ok" : std::string(ahp_last_error() ? ahp_last_error() : "?"));

    if (folded) {
        const std::string s = folded.str();
        check("folded title is \"build\"", has(s, "\"build\""), "title present");
        check("folded cwd is \"/src\"", has(s, "\"/src\""), "cwd present");
        check("folded cols survived later actions", has(s, "\"cols\":120"),
              "cols 120 retained");
    }

    // -----------------------------------------------------------------
    // 6. Handle API + the fold/apply agreement the core proves.
    // -----------------------------------------------------------------
    std::printf("\n[6] handle API — successive applyAhp equals foldAhp\n");

    AhpState stepwise(ahp_state_new());
    check("state handle allocated", stepwise.h != 0,
          "handle=" + std::to_string(stepwise.h));

    const char *seq[] = {
        R"({"type":"terminal/resized","cols":120,"rows":40})",
        R"({"type":"terminal/titleChanged","title":"build"})",
        R"({"type":"terminal/cwdChanged","cwd":"/src"})",
        R"({"type":"terminal/data","data":"hello"})",
    };
    bool all_applied = true;
    for (const char *a : seq) {
        if (ahp_state_apply(stepwise.h, a) != 0) all_applied = false;
    }
    check("all four actions applied", all_applied,
          all_applied ? "4/4" : std::string(ahp_last_error() ? ahp_last_error() : "?"));

    AhpString stepwiseJson(ahp_state_to_json(stepwise.h));
    check("stepwise state equals folded state",
          stepwiseJson && folded && stepwiseJson.str() == folded.str(),
          stepwiseJson && folded
              ? (stepwiseJson.str() == folded.str() ? "byte-identical" : "DIFFERENT")
              : "NULL");

    // Same comparison through the core's own structural equality.
    AhpState a(ahp_state_from_json(stepwiseJson.p));
    AhpState b(ahp_state_from_json(folded.p));
    const int eq = ahp_state_equals(a.h, b.h);
    check("ahp_state_equals agrees", eq == 1,
          "ahp_state_equals=" + std::to_string(eq));

    // -----------------------------------------------------------------
    // 7. Determinism.
    // -----------------------------------------------------------------
    std::printf("\n[7] reducer determinism\n");

    AhpString foldedAgain(ahp_fold_json(initial.p, actions));
    check("folding the same input twice is identical",
          foldedAgain && folded && foldedAgain.str() == folded.str(),
          foldedAgain && folded
              ? (foldedAgain.str() == folded.str() ? "byte-identical" : "DIFFERENT")
              : "NULL");

    // -----------------------------------------------------------------
    std::printf("\n=====================================\n");
    if (g_failures == 0) {
        std::printf("DEMO OK — all assertions passed.\n");
        return 0;
    }
    std::printf("DEMO FAILED — %d assertion(s) failed.\n", g_failures);
    return 1;
}
