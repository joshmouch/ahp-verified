# CMake build for the C++ distribution of the formally verified AHP core.
#
# The native library is produced by .NET NativeAOT from the Dafny-extracted
# verified core. This build can either (a) invoke `dotnet publish` to generate
# it, or (b) consume a prebuilt libahp_core.dylib from lib/.
#
# Copyright (c) Microsoft Corporation.
# Copyright (c) 2026 Josh Mouch.
# Licensed under the MIT License.

cmake_minimum_required(VERSION 3.20)
project(ahp_core_cpp VERSION 0.1.0 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

option(AHP_BUILD_NATIVE_LIB
       "Run `dotnet publish` to regenerate libahp_core.dylib from the verified core"
       OFF)

if(APPLE)
  set(AHP_LIB_NAME "libahp_core.dylib")
elseif(WIN32)
  set(AHP_LIB_NAME "ahp_core.dll")
else()
  set(AHP_LIB_NAME "libahp_core.so")
endif()

set(AHP_LIB_PATH "${CMAKE_CURRENT_SOURCE_DIR}/lib/${AHP_LIB_NAME}")

# ---------------------------------------------------------------------------
# Optionally regenerate the native library from the verified core.
# ---------------------------------------------------------------------------
if(AHP_BUILD_NATIVE_LIB)
  find_program(DOTNET_EXECUTABLE dotnet REQUIRED)
  add_custom_command(
    OUTPUT "${AHP_LIB_PATH}"
    COMMAND "${CMAKE_COMMAND}" -E echo "Publishing verified core via NativeAOT..."
    COMMAND "${DOTNET_EXECUTABLE}" publish -c Release
            "${CMAKE_CURRENT_SOURCE_DIR}/src/Ahp.Core.Native.csproj"
    COMMAND "${CMAKE_COMMAND}" -E make_directory "${CMAKE_CURRENT_SOURCE_DIR}/lib"
    COMMAND "${CMAKE_COMMAND}" -E copy
            "${CMAKE_CURRENT_SOURCE_DIR}/src/bin/Release/net10.0/osx-arm64/publish/${AHP_LIB_NAME}"
            "${AHP_LIB_PATH}"
    DEPENDS "${CMAKE_CURRENT_SOURCE_DIR}/src/AhpNative.cs"
            "${CMAKE_CURRENT_SOURCE_DIR}/src/Ahp.Core.Native.csproj"
    COMMENT "Generating ${AHP_LIB_NAME} from the verified Dafny core"
    VERBATIM)
  add_custom_target(ahp_native_lib DEPENDS "${AHP_LIB_PATH}")
else()
  if(NOT EXISTS "${AHP_LIB_PATH}")
    message(FATAL_ERROR
      "${AHP_LIB_NAME} not found at ${AHP_LIB_PATH}.\n"
      "Either run scripts/build.sh, or configure with -DAHP_BUILD_NATIVE_LIB=ON.")
  endif()
  add_custom_target(ahp_native_lib)
endif()

# ---------------------------------------------------------------------------
# Imported target for the native library.
# ---------------------------------------------------------------------------
add_library(ahp_core SHARED IMPORTED GLOBAL)
set_target_properties(ahp_core PROPERTIES
  IMPORTED_LOCATION "${AHP_LIB_PATH}"
  INTERFACE_INCLUDE_DIRECTORIES "${CMAKE_CURRENT_SOURCE_DIR}/include")
add_dependencies(ahp_core ahp_native_lib)

# ---------------------------------------------------------------------------
# C++ demo — links the library and calls the verified reducers.
# ---------------------------------------------------------------------------
add_executable(ahp_demo examples/ahp_demo.cpp)
target_link_libraries(ahp_demo PRIVATE ahp_core)
target_compile_options(ahp_demo PRIVATE -Wall -Wextra)

# Locate the dylib relative to the installed binary and to the build tree.
set_target_properties(ahp_demo PROPERTIES
  BUILD_RPATH   "${CMAKE_CURRENT_SOURCE_DIR}/lib"
  INSTALL_RPATH "@executable_path/../lib")

enable_testing()
add_test(NAME ahp_demo COMMAND ahp_demo)

# ---------------------------------------------------------------------------
# Install layout: include/ahp_core.h + lib/<native library> + bin/ahp_demo
# ---------------------------------------------------------------------------
install(FILES include/ahp_core.h DESTINATION include)
install(FILES "${AHP_LIB_PATH}" DESTINATION lib)
install(TARGETS ahp_demo RUNTIME DESTINATION bin)
