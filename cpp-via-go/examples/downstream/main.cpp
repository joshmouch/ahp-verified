// A third-party consumer of the installed AHP package. It knows nothing about
// this repository beyond find_package(ahp_verified) and the public header, so
// it exercises the distribution the way a real downstream project would.
//
// Copyright (c) Microsoft Corporation
// Copyright (c) 2026 Josh Mouch
// SPDX-License-Identifier: MIT

#include "ahp_verified.h"

#include <cstdint>
#include <iostream>
#include <string>

int main() {
  // A canvas arrives Ready with no title.
  uintptr_t canvas = ahp_canvas_new("downstream-canvas", "third.party");

  // Fold one Updated action through the verified reducer, naming a title and
  // degrading availability while leaving activity and contentUri alone.
  int outcome = -1;
  uintptr_t next = ahp_canvas_apply_updated(canvas,
                                            /*title=*/"Shipped",
                                            /*activity=*/nullptr,
                                            /*contentURI=*/nullptr,
                                            AHP_AVAILABILITY_STALE,
                                            /*now=*/7,
                                            &outcome);

  char* title = ahp_canvas_title(next);
  char* activity = ahp_canvas_activity(next);

  std::cout << "title=" << (title ? title : "None")
            << " activity=" << (activity ? activity : "None")
            << " availability="
            << (ahp_canvas_availability(next) == AHP_AVAILABILITY_STALE ? "Stale" : "Ready")
            << " outcome="
            << (outcome == AHP_OUTCOME_APPLIED ? "Applied" : "NoOp") << "\n";

  const bool ok = title != nullptr && std::string(title) == "Shipped" &&
                  activity == nullptr &&
                  ahp_canvas_availability(next) == AHP_AVAILABILITY_STALE &&
                  outcome == AHP_OUTCOME_APPLIED;

  ahp_string_free(title);
  ahp_string_free(activity);
  ahp_canvas_release(next);
  ahp_canvas_release(canvas);

  std::cout << (ok ? "downstream consumer OK" : "downstream consumer FAILED") << "\n";
  return ok ? 0 : 1;
}
