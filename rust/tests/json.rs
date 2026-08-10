//! `Json` round-trips through the core's value type.
//!
//! The marshalling layer is the crate's one piece of consequential unverified
//! code. These tests pin it: a value pushed into the core and read back must be
//! the value that went in, including the cases most likely to be lossy —
//! exact decimals, nesting, and empty containers.
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

use ahp_verified::channels::terminal::{Terminal, TerminalAction as A};
use ahp_verified::Json;

/// Push a value through the verified core and read it back.
///
/// Terminal's `claimed` transition stores its payload opaquely, so it is a
/// faithful round-trip channel for any `Json`.
fn round_trip(v: &Json) -> Json {
    Terminal::new()
        .apply(&A::Claimed(v.clone()))
        .claim()
        .expect("claimed sets the claim")
}

#[test]
fn scalars_round_trip() {
    for v in [
        Json::Null,
        Json::Bool(true),
        Json::Bool(false),
        Json::Number(0),
        Json::Number(-1),
        Json::Number(i64::MAX),
        Json::Number(i64::MIN),
        Json::string(""),
        Json::string("hello"),
    ] {
        assert_eq!(round_trip(&v), v, "round trip changed {v:?}");
    }
}

#[test]
fn non_ascii_strings_survive() {
    // Dafny strings are sequences of code points, not bytes; the conversion
    // goes through `char`, so anything outside the BMP is the risky case.
    for s in ["héllo", "日本語", "emoji: \u{1F600}", "combining: e\u{0301}"] {
        let v = Json::string(s);
        assert_eq!(round_trip(&v), v, "round trip changed {s:?}");
    }
}

#[test]
fn exact_decimals_are_not_rounded_through_a_float() {
    // The core carries decimals as mantissa/exp precisely so canonical encoding
    // stays exact. A view that converted via f64 would lose this.
    let v = Json::Decimal {
        mantissa: 123_456_789_012_345_678,
        exp: -17,
    };
    assert_eq!(round_trip(&v), v);
}

#[test]
fn empty_containers_are_distinct_from_null() {
    assert_eq!(round_trip(&Json::Array(vec![])), Json::Array(vec![]));
    assert_eq!(round_trip(&Json::object::<String>([])), Json::object::<String>([]));
    assert_ne!(round_trip(&Json::Array(vec![])), Json::Null);
    assert_ne!(round_trip(&Json::Array(vec![])), Json::object::<String>([]));
}

#[test]
fn nested_structures_round_trip() {
    let v = Json::object([
        ("kind", Json::string("session")),
        ("n", Json::Number(42)),
        ("ok", Json::Bool(true)),
        ("nothing", Json::Null),
        (
            "items",
            Json::array([
                Json::Number(1),
                Json::string("two"),
                Json::object([("deep", Json::array([Json::Bool(false), Json::Null]))]),
            ]),
        ),
    ]);
    assert_eq!(round_trip(&v), v);
}

#[test]
fn object_keys_with_awkward_characters_survive() {
    let v = Json::object([
        ("", Json::Number(1)),
        ("with space", Json::Number(2)),
        ("with\"quote", Json::Number(3)),
        ("unicode-\u{1F600}", Json::Number(4)),
    ]);
    assert_eq!(round_trip(&v), v);
}

#[test]
fn accessors_do_what_they_say() {
    let v = Json::object([
        ("s", Json::string("x")),
        ("n", Json::Number(7)),
        ("b", Json::Bool(true)),
    ]);

    assert_eq!(v.get("s").and_then(Json::as_str), Some("x"));
    assert_eq!(v.get("n").and_then(Json::as_i64), Some(7));
    assert_eq!(v.get("b").and_then(Json::as_bool), Some(true));
    assert_eq!(v.get("absent"), None);

    // Accessors are type-checked, not coercing.
    assert_eq!(v.get("s").and_then(Json::as_i64), None);
    assert_eq!(v.get("n").and_then(Json::as_str), None);
    assert_eq!(Json::Null.get("anything"), None);
}

// ------------------------------------------------------- integer boundaries

#[test]
fn integers_at_the_i64_boundary_survive_the_round_trip() {
    // Regression guard. The runtime supplies `impl From<DafnyInt> for i64` as
    // `to_i64().unwrap()`, which makes `i64::try_from(..)` infallible-with-a-
    // panic rather than checked. Reading these values back must not panic and
    // must not silently wrap.
    for n in [0_i64, 1, -1, i64::MAX, i64::MIN, i64::MAX - 1, i64::MIN + 1] {
        assert_eq!(round_trip(&Json::Number(n)), Json::Number(n), "boundary {n}");
    }
}

#[test]
fn deeply_nested_values_round_trip_without_stack_trouble() {
    let mut v = Json::Null;
    for i in 0..64 {
        v = Json::object([("depth", Json::Number(i)), ("inner", v)]);
    }
    assert_eq!(round_trip(&v), v);
}
