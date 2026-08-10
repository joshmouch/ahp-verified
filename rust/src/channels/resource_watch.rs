//! The ResourceWatch channel.
//!
//! A **passthrough** channel: it carries a watch configuration (`root`,
//! `recursive`) and forwards change notifications without mutating that
//! configuration. The core proves this — every action leaves the state equal to
//! its input — so the interesting signal here is the [`Outcome`], not the
//! state. Use [`ResourceWatch::reduce`] to see it.
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

use std::rc::Rc;

use crate::convert::{from_dstr, to_dint, to_dseq, to_dstr};
use crate::generated::ResourceWatch as core;
use crate::{Json, Outcome};

/// An action the ResourceWatch reducer understands.
#[derive(Debug, Clone, PartialEq, Eq)]
pub enum ResourceWatchAction {
    /// A change notification. Recognized; leaves the watch config unchanged.
    Changed(Json),
    /// An action this channel does not recognize. Defined as a no-op.
    Unknown(Json),
}

impl ResourceWatchAction {
    fn to_core(&self) -> Rc<core::ResourceWatchAction> {
        Rc::new(match self {
            ResourceWatchAction::Changed(changes) => core::ResourceWatchAction::RWChanged {
                changes: changes.to_core(),
            },
            ResourceWatchAction::Unknown(raw) => core::ResourceWatchAction::RWUnknown {
                raw: raw.to_core(),
            },
        })
    }
}

/// A ResourceWatch channel state: the watch configuration.
#[derive(Clone)]
pub struct ResourceWatch {
    inner: Rc<core::ResourceWatchState>,
}

impl ResourceWatch {
    /// A watch rooted at `root`, recursive or not.
    ///
    /// This channel has no `T0()` in the core — a watch is only meaningful
    /// relative to a root, so the root is required here.
    pub fn new(root: impl AsRef<str>, recursive: bool) -> ResourceWatch {
        ResourceWatch {
            inner: Rc::new(core::ResourceWatchState::ResourceWatchState {
                root: to_dstr(root.as_ref()),
                recursive,
            }),
        }
    }

    /// Apply one action through the verified reducer.
    pub fn apply(&self, action: &ResourceWatchAction) -> ResourceWatch {
        ResourceWatch {
            inner: core::_default::apply1(&self.inner, &action.to_core()),
        }
    }

    /// Apply one action, also returning the reducer's [`Outcome`].
    ///
    /// For this channel the outcome is the whole story: `Changed` yields
    /// [`Outcome::Applied`] and `Unknown` yields [`Outcome::NoOp`], while the
    /// state is provably identical either way.
    pub fn reduce(&self, action: &ResourceWatchAction) -> (ResourceWatch, Outcome) {
        let (state, outcome) =
            core::_default::ApplyToResourceWatch(&self.inner, &action.to_core(), &to_dint(9999));
        (ResourceWatch { inner: state }, Outcome::from_core(&outcome))
    }

    /// Apply a batch of actions through the core's proven kernel fold.
    pub fn apply_all(&self, actions: &[ResourceWatchAction]) -> ResourceWatch {
        let seq = to_dseq(actions.iter().map(ResourceWatchAction::to_core).collect());
        ResourceWatch {
            inner: core::_default::fold(&self.inner, &seq),
        }
    }

    // ---- readouts ---------------------------------------------------------

    /// The watch root URI.
    pub fn root(&self) -> String {
        from_dstr(self.inner.root())
    }

    /// Whether the watch descends into subdirectories.
    pub fn recursive(&self) -> bool {
        *self.inner.recursive()
    }
}

impl PartialEq for ResourceWatch {
    fn eq(&self, other: &ResourceWatch) -> bool {
        self.inner == other.inner
    }
}

impl Eq for ResourceWatch {}

impl std::fmt::Debug for ResourceWatch {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("ResourceWatch")
            .field("root", &self.root())
            .field("recursive", &self.recursive())
            .finish()
    }
}
