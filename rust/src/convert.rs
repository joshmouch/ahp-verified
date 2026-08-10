//! Marshalling between Rust types and the extracted core's Dafny types.
//!
//! This module is the crate's only unverified logic of consequence. It is
//! deliberately tiny and total: every function here is a shape change, never a
//! state transition. Nothing in this file decides anything about protocol
//! behaviour — that all lives in the extracted reducers.
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

use std::rc::Rc;

use crate::dafny_runtime::{DafnyChar, DafnyInt, Sequence};
use crate::generated::AhpSkeleton::Option as CoreOption;

/// A Dafny `string` is a sequence of code points.
pub(crate) type DStr = Sequence<DafnyChar>;

/// Rust `&str` to Dafny `string`.
pub(crate) fn to_dstr(s: &str) -> DStr {
    crate::dafny_runtime::string_of(s)
}

/// Dafny `string` to Rust `String`.
pub(crate) fn from_dstr(s: &DStr) -> String {
    s.to_array().iter().map(|c| c.0).collect()
}

/// Rust `i64` to Dafny `int`.
pub(crate) fn to_dint(n: i64) -> DafnyInt {
    DafnyInt::from(n)
}

/// Dafny `int` to Rust `i64`.
///
/// Dafny integers are unbounded. Every value the AHP reducers put in an `int`
/// field originates as a wire-level integer (exit codes, terminal dimensions,
/// millisecond durations), so this conversion is total in practice; a value
/// that genuinely exceeded `i64` would be a wire-level protocol violation.
/// Rather than panic or silently wrap, out-of-range saturates, and
/// [`from_dint_checked`] is available when the distinction matters.
pub(crate) fn from_dint(n: &DafnyInt) -> i64 {
    from_dint_checked(n).unwrap_or_else(|| {
        if num::Signed::is_negative(n) {
            i64::MIN
        } else {
            i64::MAX
        }
    })
}

/// Dafny `int` to Rust `i64`, or `None` when the value does not fit.
///
/// Deliberately goes through [`num::ToPrimitive`] rather than `i64::try_from`.
/// The runtime supplies `impl From<DafnyInt> for i64` as `to_i64().unwrap()` —
/// an *infallible* conversion that panics out of range. That makes
/// `i64::try_from(..)` resolve to the blanket impl whose `Error` is
/// `Infallible`, so `.ok()` can never be `None` and the panic fires inside the
/// call instead. `to_i64` is the only genuinely checked path.
pub(crate) fn from_dint_checked(n: &DafnyInt) -> Option<i64> {
    num::ToPrimitive::to_i64(n)
}

/// Dafny `Option<T>` to Rust `Option<&T>`.
///
/// Takes `&Rc<..>` because that is exactly how the generated field accessors
/// hand these back; requiring the caller to write `&**` at every call site
/// would be noise.
pub(crate) fn opt_ref<T: crate::dafny_runtime::DafnyType>(o: &Rc<CoreOption<T>>) -> Option<&T> {
    match &**o {
        CoreOption::Some { value } => Some(value),
        CoreOption::None {} => None,
    }
}

/// Rust `Option<T>` to Dafny `Option<T>`.
pub(crate) fn to_core_opt<T: crate::dafny_runtime::DafnyType>(o: Option<T>) -> Rc<CoreOption<T>> {
    Rc::new(match o {
        Some(value) => CoreOption::Some { value },
        None => CoreOption::None {},
    })
}

/// Rust `Option<&str>` to Dafny `Option<string>`.
pub(crate) fn to_core_opt_str(o: Option<&str>) -> Rc<CoreOption<DStr>> {
    to_core_opt(o.map(to_dstr))
}

/// Dafny `Option<string>` to Rust `Option<String>`.
pub(crate) fn from_core_opt_str(o: &Rc<CoreOption<DStr>>) -> Option<String> {
    opt_ref(o).map(from_dstr)
}

/// Dafny `Option<int>` to Rust `Option<i64>`.
pub(crate) fn from_core_opt_int(o: &Rc<CoreOption<DafnyInt>>) -> Option<i64> {
    opt_ref(o).map(from_dint)
}

/// Dafny `Option<bool>` to Rust `Option<bool>`.
pub(crate) fn from_core_opt_bool(o: &Rc<CoreOption<bool>>) -> Option<bool> {
    opt_ref(o).copied()
}

/// Build a Dafny sequence from an iterator of already-converted elements.
pub(crate) fn to_dseq<T: crate::dafny_runtime::DafnyType>(items: Vec<T>) -> Sequence<T> {
    Sequence::from_array_owned(items)
}

/// Read a Dafny sequence into a Rust `Vec` via a per-element conversion.
pub(crate) fn from_dseq<T, U, F>(seq: &Sequence<T>, f: F) -> Vec<U>
where
    T: crate::dafny_runtime::DafnyType,
    F: Fn(&T) -> U,
{
    seq.to_array().iter().map(f).collect()
}

/// Dafny sequence of strings to `Vec<String>`.
pub(crate) fn from_dseq_str(seq: &Sequence<DStr>) -> Vec<String> {
    from_dseq(seq, from_dstr)
}

/// `Vec<String>`-ish to Dafny sequence of strings.
pub(crate) fn to_dseq_str<S: AsRef<str>>(items: &[S]) -> Sequence<DStr> {
    to_dseq(items.iter().map(|s| to_dstr(s.as_ref())).collect())
}

#[cfg(test)]
mod tests {
    use super::*;

    /// Dafny integers are unbounded, so a value beyond `i64` is representable.
    fn huge() -> DafnyInt {
        let mut n = DafnyInt::from(i64::MAX);
        for _ in 0..4 {
            n = n.clone() * DafnyInt::from(1_000_000_i64);
        }
        n
    }

    #[test]
    fn from_dint_checked_reports_overflow_instead_of_panicking() {
        // The naive spelling here is `i64::try_from(n.clone()).ok()`. That
        // resolves to the blanket impl over the runtime's infallible
        // `From<DafnyInt> for i64`, whose body is `to_i64().unwrap()` -- so it
        // would panic on this input rather than return None.
        assert_eq!(from_dint_checked(&huge()), None);
        assert_eq!(from_dint_checked(&DafnyInt::from(7_i64)), Some(7));
        assert_eq!(from_dint_checked(&DafnyInt::from(i64::MAX)), Some(i64::MAX));
        assert_eq!(from_dint_checked(&DafnyInt::from(i64::MIN)), Some(i64::MIN));
    }

    #[test]
    fn from_dint_saturates_rather_than_wrapping() {
        assert_eq!(from_dint(&huge()), i64::MAX);
        assert_eq!(from_dint(&(DafnyInt::from(0_i64) - huge())), i64::MIN);
        assert_eq!(from_dint(&DafnyInt::from(-5_i64)), -5);
    }

    #[test]
    fn strings_round_trip_including_astral_code_points() {
        for s in ["", "ascii", "héllo", "\u{1F600}\u{1F680}"] {
            assert_eq!(from_dstr(&to_dstr(s)), s);
        }
    }
}
