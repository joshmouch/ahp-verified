// ahp_corpus.cpp -- run the verified core's OWN conformance corpus from C++.
//
// This is the broadest available evidence that the Rust extraction of the
// verified core is faithful: it replays every channel's fixture set through the
// reducers and prints a per-channel board. A failed expectation aborts the
// process, so a clean exit with a full green board is the pass signal.
//
// Copyright (c) Microsoft Corporation
// Copyright (c) 2026 Josh Mouch
// SPDX-License-Identifier: MIT

#include "ahp_core.h"

#include <cstdio>
#include <iostream>

int main() {
  std::cout << "AHP verified core -- conformance corpus, driven from C++\n"
            << "=======================================================\n\n";
  // The core prints through Rust's stdout, which buffers independently of
  // std::cout; flush so the board does not appear before this banner.
  std::cout.flush();
  std::fflush(stdout);

  const int32_t rc = ahp_run_corpus();

  std::fflush(stdout);
  std::cout << "\nahp_run_corpus() returned " << rc << "\n";
  return rc;
}
