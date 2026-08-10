//! A safe Rust view of the core's JSON value type.
//!
//! Several AHP channels carry opaque JSON payloads (claims, ranges, provider
//! error bodies, unknown-action passthrough). The verified core models these
//! with its own `ConfluxCodec.Json` datatype rather than a host JSON library,
//! so that the proofs quantify over the same values the reducers see.
//!
//! [`Json`] mirrors that datatype one-for-one. Note `Decimal`, which the core
//! carries as an exact `mantissa`/`exp` pair rather than an `f64` — the core's
//! canonical-encoding proofs depend on that exactness, so this crate preserves
//! it rather than lossily converting.
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

use std::collections::BTreeMap;
use std::rc::Rc;

use crate::convert::{from_dint, from_dstr, to_dint, to_dstr};
use crate::dafny_runtime::Map;
use crate::generated::ConfluxCodec::Json as CoreJson;

/// A JSON value as modelled by the verified core.
#[derive(Debug, Clone, PartialEq, Eq, PartialOrd, Ord)]
pub enum Json {
    /// `null`
    Null,
    /// `true` / `false`
    Bool(bool),
    /// An integer. The core models integers as unbounded; this view uses `i64`.
    Number(i64),
    /// An exact decimal, as `mantissa * 10^exp`.
    ///
    /// Kept exact rather than converted to a float: the core's canonical
    /// encoding proofs are stated over this representation.
    Decimal {
        /// The significand.
        mantissa: i64,
        /// The base-10 exponent.
        exp: i64,
    },
    /// A string.
    String(String),
    /// An array.
    Array(Vec<Json>),
    /// An object.
    ///
    /// A `BTreeMap` because the core's object model is an unordered map whose
    /// canonical encoding sorts keys; a sorted map makes the Rust view's
    /// iteration order match the canonical order rather than a hash order.
    Object(BTreeMap<String, Json>),
}

impl Json {
    /// Convenience constructor for a string value.
    pub fn string(s: impl Into<String>) -> Json {
        Json::String(s.into())
    }

    /// Convenience constructor for an object from key/value pairs.
    pub fn object<K: Into<String>>(pairs: impl IntoIterator<Item = (K, Json)>) -> Json {
        Json::Object(pairs.into_iter().map(|(k, v)| (k.into(), v)).collect())
    }

    /// Convenience constructor for an array.
    pub fn array(items: impl IntoIterator<Item = Json>) -> Json {
        Json::Array(items.into_iter().collect())
    }

    /// Borrow the value at an object key, if this is an object with that key.
    pub fn get(&self, key: &str) -> Option<&Json> {
        match self {
            Json::Object(m) => m.get(key),
            _ => None,
        }
    }

    /// Borrow the string contents, if this is a string.
    pub fn as_str(&self) -> Option<&str> {
        match self {
            Json::String(s) => Some(s),
            _ => None,
        }
    }

    /// The integer value, if this is a number.
    pub fn as_i64(&self) -> Option<i64> {
        match self {
            Json::Number(n) => Some(*n),
            _ => None,
        }
    }

    /// The boolean value, if this is a bool.
    pub fn as_bool(&self) -> Option<bool> {
        match self {
            Json::Bool(b) => Some(*b),
            _ => None,
        }
    }

    // ---- core interop -----------------------------------------------------

    pub(crate) fn to_core(&self) -> Rc<CoreJson> {
        Rc::new(match self {
            Json::Null => CoreJson::JNull {},
            Json::Bool(b) => CoreJson::JBool { b: *b },
            Json::Number(n) => CoreJson::JNum { n: to_dint(*n) },
            Json::Decimal { mantissa, exp } => CoreJson::JDec {
                mantissa: to_dint(*mantissa),
                exp: to_dint(*exp),
            },
            Json::String(s) => CoreJson::JStr { s: to_dstr(s) },
            Json::Array(items) => CoreJson::JArr {
                elems: crate::convert::to_dseq(items.iter().map(Json::to_core).collect()),
            },
            Json::Object(fields) => {
                let pairs: Vec<(crate::convert::DStr, Rc<CoreJson>)> = fields
                    .iter()
                    .map(|(k, v)| (to_dstr(k), v.to_core()))
                    .collect();
                CoreJson::JObj {
                    fields: Map::from_iterator(pairs.into_iter()),
                }
            }
        })
    }

    pub(crate) fn from_core(j: &CoreJson) -> Json {
        match j {
            CoreJson::JNull {} => Json::Null,
            CoreJson::JBool { b } => Json::Bool(*b),
            CoreJson::JNum { n } => Json::Number(from_dint(n)),
            CoreJson::JDec { mantissa, exp } => Json::Decimal {
                mantissa: from_dint(mantissa),
                exp: from_dint(exp),
            },
            CoreJson::JStr { s } => Json::String(from_dstr(s)),
            CoreJson::JArr { elems } => {
                Json::Array(crate::convert::from_dseq(elems, |e| Json::from_core(e)))
            }
            CoreJson::JObj { fields } => {
                let mut out = BTreeMap::new();
                for k in fields.keys().iter() {
                    let v = fields.get(k);
                    out.insert(from_dstr(k), Json::from_core(&v));
                }
                Json::Object(out)
            }
        }
    }
}

impl From<bool> for Json {
    fn from(b: bool) -> Json {
        Json::Bool(b)
    }
}

impl From<i64> for Json {
    fn from(n: i64) -> Json {
        Json::Number(n)
    }
}

impl From<&str> for Json {
    fn from(s: &str) -> Json {
        Json::String(s.to_owned())
    }
}

impl From<String> for Json {
    fn from(s: String) -> Json {
        Json::String(s)
    }
}
