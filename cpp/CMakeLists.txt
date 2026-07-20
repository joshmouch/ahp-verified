# AHP verified core -- C++ target.
#
# STATUS: the AHP Dafny core does NOT translate to C++. See README.md for the
# full evidence. This build file therefore does NOT build the AHP core; no such
# source exists to build, because `dafny translate cpp` refuses to emit it.
#
# What this file DOES build is the toolchain conformance probe: a minimal Dafny
# program translated to C++ and compiled against the patched Dafny C++ runtime
# header. It exists so that the "C++ is blocked" claim stays falsifiable -- if a
# future Dafny release lifts the unbounded-integer and arrow-type restrictions,
# this target is the harness the real core drops into.

cmake_minimum_required(VERSION 3.16)

project(ahp_core_verified_cpp
        VERSION 0.5.2
        DESCRIPTION "Agent Host Protocol verified core -- C++ target (blocked; see README)"
        LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# The Dafny C++ runtime header, vendored and patched. Upstream Dafny 4.11.0
# ships a DafnyRuntime.h that does not compile at all; vendor/DafnyRuntime.h.patch
# documents the one-token fix carried here.
add_library(dafny_runtime_cpp INTERFACE)
target_include_directories(dafny_runtime_cpp INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/vendor)

# Toolchain conformance probe. Regenerate its source with:
#   dafny translate cpp smoke/probe_ok.dfy --no-verify --unicode-char false -o smoke/probe_ok
add_executable(ahp_cpp_toolchain_probe smoke/probe_ok.cpp)
target_include_directories(ahp_cpp_toolchain_probe PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/smoke)
target_link_libraries(ahp_cpp_toolchain_probe PRIVATE dafny_runtime_cpp)

enable_testing()
add_test(NAME toolchain_probe COMMAND ahp_cpp_toolchain_probe)
set_tests_properties(toolchain_probe PROPERTIES
                     PASS_REGULAR_EXPRESSION "hello from cpp")
