# Build manifest for the C++ distribution of the formally-verified AHP core.
#
# The native library is produced by cargo from Rust that `dafny translate rs`
# generated out of the verified Dafny core. CMake drives cargo, then exposes the
# result as an ordinary imported target that C++ consumers link against.
#
#   cmake -S . -B build-cmake -DCMAKE_BUILD_TYPE=Release
#   cmake --build build-cmake
#   ctest --test-dir build-cmake --output-on-failure
#
# Copyright (c) Microsoft Corporation
# Copyright (c) 2026 Josh Mouch
# SPDX-License-Identifier: MIT

cmake_minimum_required(VERSION 3.20)
project(ahp_core_cpp LANGUAGES CXX VERSION 0.1.0)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_program(CARGO_EXECUTABLE cargo REQUIRED)

set(AHP_FFI_DIR "${CMAKE_CURRENT_SOURCE_DIR}/ahp-core-ffi")

# Debug CMake builds still use a release-profile Rust library: the generated core
# is large and a debug build of it is dramatically slower without being any more
# debuggable from C++.
set(AHP_CARGO_PROFILE "release")
set(AHP_STATIC_LIB
    "${AHP_FFI_DIR}/target/${AHP_CARGO_PROFILE}/${CMAKE_STATIC_LIBRARY_PREFIX}ahp_core${CMAKE_STATIC_LIBRARY_SUFFIX}")

file(GLOB_RECURSE AHP_RUST_SOURCES CONFIGURE_DEPENDS
     "${AHP_FFI_DIR}/src/*.rs"
     "${AHP_FFI_DIR}/Cargo.toml")

add_custom_command(
    OUTPUT "${AHP_STATIC_LIB}"
    COMMAND "${CARGO_EXECUTABLE}" build --release --manifest-path "${AHP_FFI_DIR}/Cargo.toml"
    DEPENDS ${AHP_RUST_SOURCES}
    COMMENT "Building verified AHP core (cargo, ${AHP_CARGO_PROFILE})"
    VERBATIM)

add_custom_target(ahp_core_build DEPENDS "${AHP_STATIC_LIB}")

add_library(ahp_core STATIC IMPORTED GLOBAL)
set_target_properties(ahp_core PROPERTIES
    IMPORTED_LOCATION "${AHP_STATIC_LIB}"
    INTERFACE_INCLUDE_DIRECTORIES "${CMAKE_CURRENT_SOURCE_DIR}/include")
add_dependencies(ahp_core ahp_core_build)

# The Rust staticlib needs the platform runtime libraries it was built against.
if(APPLE)
    target_link_libraries(ahp_core INTERFACE
        "-framework CoreFoundation" "-framework Security" resolv)
elseif(UNIX)
    target_link_libraries(ahp_core INTERFACE pthread dl m)
endif()

# ---- demo / test ------------------------------------------------------------

add_executable(ahp_terminal_demo examples/ahp_terminal_demo.cpp)
target_link_libraries(ahp_terminal_demo PRIVATE ahp_core)
target_compile_options(ahp_terminal_demo PRIVATE -Wall -Wextra)

add_executable(ahp_corpus examples/ahp_corpus.cpp)
target_link_libraries(ahp_corpus PRIVATE ahp_core)
target_compile_options(ahp_corpus PRIVATE -Wall -Wextra)

enable_testing()
add_test(NAME terminal_reducer COMMAND ahp_terminal_demo)
add_test(NAME conformance_corpus COMMAND ahp_corpus)

# ---- install ----------------------------------------------------------------

include(GNUInstallDirs)
install(FILES include/ahp_core.h DESTINATION "${CMAKE_INSTALL_INCLUDEDIR}")
install(FILES "${AHP_STATIC_LIB}" DESTINATION "${CMAKE_INSTALL_LIBDIR}")
