# A standalone project that consumes the INSTALLED package, to prove the
# distribution is usable from outside this tree:
#
#     cmake -S . -B ../../../build/install -DCMAKE_INSTALL_PREFIX=/tmp/ahp
#     cmake --build ../../../build/install
#     cmake --install ../../../build/install
#     cmake -S examples/downstream -B /tmp/downstream -DCMAKE_PREFIX_PATH=/tmp/ahp
#     cmake --build /tmp/downstream && /tmp/downstream/consumer
#
# Copyright (c) Microsoft Corporation
# Copyright (c) 2026 Josh Mouch
# SPDX-License-Identifier: MIT

cmake_minimum_required(VERSION 3.20)
project(ahp_downstream_consumer LANGUAGES CXX)

find_package(ahp_verified REQUIRED)

add_executable(consumer main.cpp)
target_compile_features(consumer PRIVATE cxx_std_17)
target_compile_options(consumer PRIVATE -Wall -Wextra)
target_link_libraries(consumer PRIVATE ahp::verified)
