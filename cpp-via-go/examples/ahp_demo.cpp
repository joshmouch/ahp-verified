// Demonstration C++ client of the formally-verified AHP core.
//
// Every value this program prints is computed by the Dafny-extracted reducers
// and codecs reached through the C ABI — none of the protocol behaviour is
// reimplemented here. The program exits non-zero if any check fails, so it
// doubles as a link-and-run smoke test for the packaged artifact.
//
// Copyright (c) Microsoft Corporation
// Copyright (c) 2026 Josh Mouch
// SPDX-License-Identifier: MIT

#include "ahp_verified.h"

#include <cstdint>
#include <iostream>
#include <memory>
#include <string>

namespace {

// Owns a char* returned by the library and releases it with ahp_string_free.
class OwnedStr {
 public:
  explicit OwnedStr(char* p) : p_(p) {}
  ~OwnedStr() { ahp_string_free(p_); }
  OwnedStr(const OwnedStr&) = delete;
  OwnedStr& operator=(const OwnedStr&) = delete;

  bool present() const { return p_ != nullptr; }
  std::string value() const { return p_ ? std::string(p_) : std::string(); }
  // Renders an absent Option the way the specification names it.
  std::string display() const { return p_ ? std::string(p_) : std::string("None"); }

 private:
  char* p_;
};

// Owns a Canvas handle and releases it with ahp_canvas_release.
class OwnedCanvas {
 public:
  explicit OwnedCanvas(uintptr_t h) : h_(h) {}
  ~OwnedCanvas() { ahp_canvas_release(h_); }
  OwnedCanvas(const OwnedCanvas&) = delete;
  OwnedCanvas& operator=(const OwnedCanvas&) = delete;

  uintptr_t get() const { return h_; }

 private:
  uintptr_t h_;
};

int failures = 0;

void check(bool ok, const std::string& what) {
  if (ok) {
    std::cout << "    ok    " << what << "\n";
  } else {
    std::cout << "    FAIL  " << what << "\n";
    ++failures;
  }
}

const char* outcomeName(int code) {
  switch (code) {
    case AHP_OUTCOME_APPLIED: return "Applied";
    case AHP_OUTCOME_NOOP: return "NoOp";
    case AHP_OUTCOME_OUT_OF_SCOPE: return "OutOfScope";
    default: return "?";
  }
}

const char* availabilityName(int code) {
  return code == AHP_AVAILABILITY_READY ? "Ready" : "Stale";
}

void describe(const char* label, uintptr_t canvas) {
  OwnedStr id(ahp_canvas_id(canvas));
  OwnedStr title(ahp_canvas_title(canvas));
  OwnedStr activity(ahp_canvas_activity(canvas));
  OwnedStr uri(ahp_canvas_content_uri(canvas));
  std::cout << "    " << label << ": canvasId=" << id.value()
            << " title=" << title.display()
            << " activity=" << activity.display()
            << " contentUri=" << uri.display()
            << " availability=" << availabilityName(ahp_canvas_availability(canvas))
            << "\n";
}

}  // namespace

int main() {
  {
    OwnedStr abi(ahp_abi_version());
    std::cout << "AHP verified core — C++ client via the Go c-archive bridge\n"
              << "binding ABI version " << abi.value() << "\n\n";
  }

  // ── 1. The Canvas channel reducer ──────────────────────────────────────
  std::cout << "[1] Canvas reducer (Canvas.ApplyToCanvas)\n";

  OwnedCanvas fresh(ahp_canvas_new("canvas-1", "acme.editors"));
  describe("initial ", fresh.get());
  {
    OwnedStr title(ahp_canvas_title(fresh.get()));
    check(!title.present(), "a fresh canvas carries no title");
  }

  int outcome = -1;
  OwnedCanvas updated(ahp_canvas_apply_updated(
      fresh.get(),
      /*title=*/"Design Review",
      /*activity=*/"editing",
      /*contentUri=*/"ahp-session:/2f9c/content/canvas-1",
      /*availability=*/AHP_AVAILABILITY_UNCHANGED,
      /*now=*/1,
      &outcome));
  describe("Updated ", updated.get());
  std::cout << "    outcome: " << outcomeName(outcome) << "\n";

  {
    OwnedStr title(ahp_canvas_title(updated.get()));
    OwnedStr activity(ahp_canvas_activity(updated.get()));
    check(title.value() == "Design Review", "Updated set the title");
    check(activity.value() == "editing", "Updated set the activity");
    check(outcome == AHP_OUTCOME_APPLIED, "Updated reports Applied");
    check(ahp_canvas_availability(updated.get()) == AHP_AVAILABILITY_READY,
          "an unchanged availability is preserved");
  }

  // Purity: the reducer must not have touched the state it was handed.
  {
    OwnedStr originalTitle(ahp_canvas_title(fresh.get()));
    check(!originalTitle.present(), "the reducer left its input state untouched");
  }

  // A sparse update touches only the fields it names.
  int sparseOutcome = -1;
  OwnedCanvas staled(ahp_canvas_apply_updated(
      updated.get(),
      /*title=*/nullptr,
      /*activity=*/nullptr,
      /*contentUri=*/nullptr,
      /*availability=*/AHP_AVAILABILITY_STALE,
      /*now=*/2,
      &sparseOutcome));
  describe("degraded", staled.get());
  {
    OwnedStr title(ahp_canvas_title(staled.get()));
    check(title.value() == "Design Review", "a sparse update preserved the title");
    check(ahp_canvas_availability(staled.get()) == AHP_AVAILABILITY_STALE,
          "a sparse update degraded availability to Stale");
  }

  // CloseRequested is state-preserving per the specification.
  int closeOutcome = -1;
  OwnedCanvas closed(
      ahp_canvas_apply_close_requested(staled.get(), /*now=*/3, &closeOutcome));
  std::cout << "    CloseRequested outcome: " << outcomeName(closeOutcome) << "\n";
  check(closeOutcome == AHP_OUTCOME_NOOP, "CloseRequested reports NoOp");
  check(ahp_canvas_equals(closed.get(), staled.get()) == 1,
        "CloseRequested preserved the state");

  // ── 2. The verified JSON codec ─────────────────────────────────────────
  std::cout << "\n[2] JSON codec (ConfluxJsonText.ParseJsonChecked / Stringify)\n";
  {
    const char* input = "{\"seq\": 42, \"channel\":\"canvas\"}";
    OwnedStr canonical(ahp_json_canonicalize(input));
    std::cout << "    input:     " << input << "\n"
              << "    canonical: " << canonical.display() << "\n";
    check(canonical.present(), "well-formed JSON was accepted");

    OwnedStr again(ahp_json_canonicalize(canonical.value().c_str()));
    check(again.present() && again.value() == canonical.value(),
          "canonicalization is idempotent");

    OwnedStr rejected(ahp_json_canonicalize("{\"unterminated\": "));
    check(!rejected.present(), "malformed JSON was refused");
  }

  // ── 3. Verified command identity ───────────────────────────────────────
  std::cout << "\n[3] Command identity (ConfluxCommandIdentity.Sha256Digest)\n";
  {
    const char* data = "abc";
    OwnedStr digest(ahp_sha256_digest(data, 3));
    std::cout << "    digest(\"abc\") = " << digest.display() << "\n";
    check(digest.value() ==
              "sha256:ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
          "digest matches the published SHA-256 test vector");
  }

  // ── 4. The embedded conformance corpus ─────────────────────────────────
  std::cout << "\n[4] Embedded corpus (Session.RunCorpus)\n";
  {
    int64_t passed = 0;
    int64_t modeled = 0;
    int allPassed = ahp_run_corpus(&passed, &modeled);
    std::cout << "    corpus: " << passed << "/" << modeled << " modeled cases passed\n";
    check(modeled > 0, "the corpus reported a non-empty modeled-case count");
    check(allPassed == 1, "every modeled corpus case passed");
  }

  std::cout << "\n" << (failures == 0 ? "ALL CHECKS PASSED" : "CHECKS FAILED")
            << " (" << failures << " failure" << (failures == 1 ? "" : "s") << ")\n";
  return failures == 0 ? 0 : 1;
}
