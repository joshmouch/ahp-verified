// ahp_terminal_demo.cpp -- C++ driving the formally-verified AHP Terminal reducer.
//
// Every state transition below is a call into reducer code machine-generated
// from the verified Dafny core. This program asserts on the results, so it fails
// loudly if the reducer's observable behaviour ever changes.
//
// The expected values are the Dafny core's own corpus expectations (see
// spec/terminal.dfy `RunCorpus`), not values read back off this implementation.
//
// Copyright (c) Microsoft Corporation
// Copyright (c) 2026 Josh Mouch
// SPDX-License-Identifier: MIT

#include "ahp_core.h"

#include <cstdint>
#include <cstdlib>
#include <iostream>
#include <memory>
#include <string>
#include <vector>

namespace {

int failures = 0;

void check(bool ok, const std::string &what) {
  std::cout << (ok ? "  ok   " : "  FAIL ") << what << "\n";
  if (!ok) {
    ++failures;
  }
}

// RAII wrappers so the demo cannot leak even on an early return.
struct StateDeleter {
  void operator()(AhpTerminalState *s) const noexcept { ahp_terminal_free(s); }
};
using State = std::unique_ptr<AhpTerminalState, StateDeleter>;

struct StringDeleter {
  void operator()(char *s) const noexcept { ahp_string_free(s); }
};
using CStr = std::unique_ptr<char, StringDeleter>;

// Read an owned char* into a std::string; reports whether it was non-NULL.
bool owned_string(char *raw, std::string &out) {
  CStr guard(raw);
  if (!guard) {
    return false;
  }
  out.assign(guard.get());
  return true;
}

} // namespace

int main() {
  std::cout << "AHP verified core -- C++ calling the Terminal channel reducer\n"
            << "============================================================\n\n";

  State s0(ahp_terminal_initial());
  if (!s0) {
    std::cerr << "ahp_terminal_initial() returned NULL\n";
    return 1;
  }

  // --- initial state -------------------------------------------------------
  std::cout << "initial state (Dafny Terminal.T0()):\n";
  {
    std::string title;
    check(owned_string(ahp_terminal_title(s0.get()), title) && title == "bash",
          "title == \"bash\"");

    std::string cwd;
    check(!owned_string(ahp_terminal_cwd(s0.get()), cwd),
          "cwd is None (NULL, not empty string)");

    std::int64_t code = -1;
    check(ahp_terminal_exit_code(s0.get(), &code) == 0, "exitCode is None");
    check(ahp_terminal_content_len(s0.get()) == 0, "content is empty");
  }

  // --- single transitions through apply1 -----------------------------------
  std::cout << "\nsingle transitions (verified reducer apply1):\n";

  State s1(ahp_terminal_cwd_changed(s0.get(), "/tmp"));
  {
    std::string cwd;
    check(s1 && owned_string(ahp_terminal_cwd(s1.get()), cwd) && cwd == "/tmp",
          "cwdChanged(\"/tmp\")            -> cwd == \"/tmp\"");
  }

  State s2(ahp_terminal_title_changed(s1.get(), "zsh"));
  {
    std::string title;
    check(s2 && owned_string(ahp_terminal_title(s2.get()), title) && title == "zsh",
          "titleChanged(\"zsh\")           -> title == \"zsh\"");
  }

  State s3(ahp_terminal_resized(s2.get(), 120, 40));
  {
    std::int64_t cols = 0, rows = 0;
    check(s3 && ahp_terminal_size(s3.get(), &cols, &rows) == 1 && cols == 120 &&
              rows == 40,
          "resized(120, 40)              -> cols == 120, rows == 40");
  }

  State s4(ahp_terminal_exited(s3.get(), 130));
  {
    std::int64_t code = 0;
    check(s4 && ahp_terminal_exit_code(s4.get(), &code) == 1 && code == 130,
          "exited(130)                   -> exitCode == 130");
  }

  // --- purity: the reducer must not mutate its input -----------------------
  std::cout << "\npurity (reducer returns a new state; inputs untouched):\n";
  {
    std::string title0, title2;
    owned_string(ahp_terminal_title(s0.get()), title0);
    owned_string(ahp_terminal_title(s2.get()), title2);
    check(title0 == "bash" && title2 == "zsh",
          "s0 still \"bash\" after s2 became \"zsh\"");

    std::int64_t code = 0;
    check(ahp_terminal_exit_code(s3.get(), &code) == 0,
          "s3 exitCode still None after s4 exited(130)");
  }

  // --- higher-order path: the proven kernel fold ---------------------------
  // This is the capability the Dafny C++ backend cannot express at all
  // (`Fold` takes a function value). It works here because the core was
  // translated to Rust, where Dafny emits Rc<dyn Fn>.
  std::cout << "\nkernel fold (higher-order Terminal.fold over ConfluxContract.Fold):\n";
  {
    const std::vector<const char *> chunks = {"hello ", "verified ", "world"};
    State folded(
        ahp_terminal_fold_data(s0.get(), chunks.data(), chunks.size()));
    check(folded != nullptr, "fold over 3 data actions returned a state");

    // Fold must agree with three sequential apply1 calls -- that equivalence is
    // a theorem in the Dafny core; here we exercise it across the C ABI.
    State a(ahp_terminal_data(s0.get(), "hello "));
    State b(ahp_terminal_data(a.get(), "verified "));
    State c(ahp_terminal_data(b.get(), "world"));

    const std::size_t viaFold = ahp_terminal_content_len(folded.get());
    const std::size_t viaApply = ahp_terminal_content_len(c.get());
    check(viaFold == viaApply,
          "fold(s0, [3 actions]) content_len == apply1 x3 content_len (" +
              std::to_string(viaFold) + " == " + std::to_string(viaApply) + ")");
  }

  // --- defensive: NULL handling --------------------------------------------
  std::cout << "\nboundary behaviour:\n";
  check(ahp_terminal_cwd_changed(nullptr, "/x") == nullptr,
        "NULL state handle             -> NULL");
  check(ahp_terminal_cwd_changed(s0.get(), nullptr) == nullptr,
        "NULL string argument          -> NULL");

  std::cout << "\n============================================================\n";
  if (failures == 0) {
    std::cout << "ALL CHECKS PASSED -- C++ successfully drove the verified "
                 "AHP reducer.\n";
    return 0;
  }
  std::cout << failures << " CHECK(S) FAILED\n";
  return 1;
}
