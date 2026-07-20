#
# Build manifest for the C++ distribution of the formally-verified AHP core.
#
# The native library is produced by regenerate.sh from the Dafny-extracted Go
# core; CMake does not reimplement that step, it invokes it. Configure with
# -DAHP_REGENERATE=ON to force a rebuild of the library, or leave it off to
# consume the artifacts already committed under lib/ and include/.
#
# Copyright (c) Microsoft Corporation
# Copyright (c) 2026 Josh Mouch
# SPDX-License-Identifier: MIT

cmake_minimum_required(VERSION 3.20)

project(ahp_verified
        VERSION 1.0.0
        DESCRIPTION "C ABI for the formally-verified Agent Host Protocol core"
        LANGUAGES CXX)

include(GNUInstallDirs)
include(CMakePackageConfigHelpers)

if(APPLE)
  set(AHP_SHARED_SUFFIX "dylib")
else()
  set(AHP_SHARED_SUFFIX "so")
endif()

set(AHP_STATIC_LIB "${CMAKE_CURRENT_SOURCE_DIR}/lib/libahpverified.a")
set(AHP_GENERATED_HEADER "${CMAKE_CURRENT_SOURCE_DIR}/include/libahpverified.h")

# ── Producing the verified library ───────────────────────────────────────────
# The library is a build product of the Dafny-extracted Go core. Regenerate it
# when asked, or whenever it is simply not there yet.

option(AHP_REGENERATE
       "Rebuild the native library from the Dafny-extracted Go core" OFF)

if(AHP_REGENERATE OR NOT EXISTS "${AHP_STATIC_LIB}")
  find_program(AHP_GO_EXECUTABLE go REQUIRED)
  message(STATUS "Regenerating the verified library from the extracted Go core")
  execute_process(
    COMMAND "${CMAKE_CURRENT_SOURCE_DIR}/regenerate.sh"
    WORKING_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}"
    RESULT_VARIABLE AHP_REGENERATE_RESULT
    COMMAND_ECHO STDOUT)
  if(NOT AHP_REGENERATE_RESULT EQUAL 0)
    message(FATAL_ERROR "regenerate.sh failed (exit ${AHP_REGENERATE_RESULT})")
  endif()
endif()

if(NOT EXISTS "${AHP_STATIC_LIB}")
  message(FATAL_ERROR
          "${AHP_STATIC_LIB} is missing. Run ./regenerate.sh, or configure "
          "with -DAHP_REGENERATE=ON.")
endif()

# ── The imported target ──────────────────────────────────────────────────────

add_library(ahp_verified STATIC IMPORTED GLOBAL)
set_target_properties(ahp_verified PROPERTIES
  IMPORTED_LOCATION "${AHP_STATIC_LIB}"
  INTERFACE_INCLUDE_DIRECTORIES "${CMAKE_CURRENT_SOURCE_DIR}/include")

# The archive embeds the Go runtime, which needs a threading library
# everywhere and dynamic loading plus libresolv on ELF platforms.
find_package(Threads REQUIRED)
target_link_libraries(ahp_verified INTERFACE Threads::Threads)
if(NOT APPLE AND UNIX)
  target_link_libraries(ahp_verified INTERFACE ${CMAKE_DL_LIBS} resolv)
endif()

add_library(ahp::verified ALIAS ahp_verified)

# ── The example, which doubles as the smoke test ─────────────────────────────

option(AHP_BUILD_EXAMPLES "Build the C++ example program" ON)

if(AHP_BUILD_EXAMPLES)
  add_executable(ahp_demo examples/ahp_demo.cpp)
  target_compile_features(ahp_demo PRIVATE cxx_std_17)
  target_compile_options(ahp_demo PRIVATE -Wall -Wextra)
  target_link_libraries(ahp_demo PRIVATE ahp::verified)

  enable_testing()
  # The example returns a non-zero exit status if any check fails, so running
  # it is a real assertion that the packaged library works.
  add_test(NAME ahp_demo COMMAND ahp_demo)
endif()

# ── Installation ─────────────────────────────────────────────────────────────

install(FILES
          "${CMAKE_CURRENT_SOURCE_DIR}/include/ahp_verified.h"
          "${AHP_GENERATED_HEADER}"
        DESTINATION "${CMAKE_INSTALL_INCLUDEDIR}")

install(FILES "${AHP_STATIC_LIB}"
        DESTINATION "${CMAKE_INSTALL_LIBDIR}")

install(FILES "${CMAKE_CURRENT_SOURCE_DIR}/lib/libahpverified.${AHP_SHARED_SUFFIX}"
        DESTINATION "${CMAKE_INSTALL_LIBDIR}"
        OPTIONAL)

# A package config so downstream projects can find_package(ahp_verified) and
# link ahp::verified.
configure_package_config_file(
  "${CMAKE_CURRENT_SOURCE_DIR}/cmake/ahp_verifiedConfig.cmake.in"
  "${CMAKE_CURRENT_BINARY_DIR}/ahp_verifiedConfig.cmake"
  INSTALL_DESTINATION "${CMAKE_INSTALL_LIBDIR}/cmake/ahp_verified"
  # Without PATH_VARS the PACKAGE_* placeholders in the template are never
  # substituted, and the installed config resolves its paths to nothing.
  PATH_VARS CMAKE_INSTALL_INCLUDEDIR CMAKE_INSTALL_LIBDIR)

write_basic_package_version_file(
  "${CMAKE_CURRENT_BINARY_DIR}/ahp_verifiedConfigVersion.cmake"
  COMPATIBILITY SameMajorVersion)

install(FILES
          "${CMAKE_CURRENT_BINARY_DIR}/ahp_verifiedConfig.cmake"
          "${CMAKE_CURRENT_BINARY_DIR}/ahp_verifiedConfigVersion.cmake"
        DESTINATION "${CMAKE_INSTALL_LIBDIR}/cmake/ahp_verified")
