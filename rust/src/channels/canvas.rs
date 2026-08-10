//! The Canvas channel.
//!
//! Models a provider-backed canvas surface: title, activity, content URI, and a
//! readiness flag. The core's canvas laws are notably strict about *identity*:
//! a canvas instance is identified by its channel URI alone, so no snapshot
//! field — not even `canvas_id` — can change which instance you are looking at.
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

use std::rc::Rc;

use crate::convert::{from_core_opt_str, from_dstr, to_core_opt, to_core_opt_str, to_dint, to_dseq};
use crate::generated::Canvas as core;
use crate::{Json, Outcome};

/// Whether the canvas snapshot is current.
#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub enum Availability {
    /// The snapshot reflects the provider's current state.
    Ready,
    /// The snapshot may be out of date.
    Stale,
}

impl Availability {
    fn to_core(self) -> Rc<core::CanvasAvailability> {
        Rc::new(match self {
            Availability::Ready => core::CanvasAvailability::Ready {},
            Availability::Stale => core::CanvasAvailability::Stale {},
        })
    }

    fn from_core(a: &core::CanvasAvailability) -> Availability {
        match a {
            core::CanvasAvailability::Ready {} => Availability::Ready,
            core::CanvasAvailability::Stale {} => Availability::Stale,
        }
    }
}

/// An action the Canvas reducer understands.
#[derive(Debug, Clone, PartialEq, Eq, Default)]
pub struct Updated {
    /// New title, or `None` to leave unchanged.
    pub title: Option<String>,
    /// New activity string, or `None` to leave unchanged.
    pub activity: Option<String>,
    /// New content URI, or `None` to leave unchanged.
    pub content_uri: Option<String>,
    /// New availability, or `None` to leave unchanged.
    pub availability: Option<Availability>,
}

/// An action the Canvas reducer understands.
#[derive(Debug, Clone, PartialEq, Eq)]
pub enum CanvasAction {
    /// A partial update. Fields left `None` are preserved, not cleared.
    Updated(Updated),
    /// The provider asked to close the canvas.
    CloseRequested,
    /// An action this channel does not recognize. Defined as a no-op.
    Unknown(Json),
}

impl CanvasAction {
    /// A partial update touching only the title.
    pub fn title(title: impl Into<String>) -> CanvasAction {
        CanvasAction::Updated(Updated {
            title: Some(title.into()),
            ..Default::default()
        })
    }

    /// A partial update touching only the availability.
    pub fn availability(availability: Availability) -> CanvasAction {
        CanvasAction::Updated(Updated {
            availability: Some(availability),
            ..Default::default()
        })
    }

    fn to_core(&self) -> Rc<core::CanvasAction> {
        Rc::new(match self {
            CanvasAction::Updated(u) => core::CanvasAction::Updated {
                title: to_core_opt_str(u.title.as_deref()),
                activity: to_core_opt_str(u.activity.as_deref()),
                contentUri: to_core_opt_str(u.content_uri.as_deref()),
                availability: to_core_opt(u.availability.map(Availability::to_core)),
            },
            CanvasAction::CloseRequested => core::CanvasAction::CloseRequested {},
            CanvasAction::Unknown(raw) => core::CanvasAction::CanvasUnknown {
                raw: raw.to_core(),
            },
        })
    }
}

/// A Canvas channel state.
#[derive(Clone)]
pub struct Canvas {
    inner: Rc<core::CanvasState>,
}

impl Canvas {
    /// The initial state, as defined by the core's `Canvas.C0()`.
    pub fn new() -> Canvas {
        Canvas {
            inner: core::_default::C0(),
        }
    }

    /// Apply one action through the verified reducer.
    pub fn apply(&self, action: &CanvasAction) -> Canvas {
        Canvas {
            inner: core::_default::apply1(&self.inner, &action.to_core()),
        }
    }

    /// Apply one action, also returning the reducer's [`Outcome`].
    pub fn reduce(&self, action: &CanvasAction) -> (Canvas, Outcome) {
        let (state, outcome) =
            core::_default::ApplyToCanvas(&self.inner, &action.to_core(), &to_dint(9999));
        (Canvas { inner: state }, Outcome::from_core(&outcome))
    }

    /// Apply a batch of actions through the core's proven kernel fold.
    pub fn apply_all(&self, actions: &[CanvasAction]) -> Canvas {
        let seq = to_dseq(actions.iter().map(CanvasAction::to_core).collect());
        Canvas {
            inner: core::_default::fold(&self.inner, &seq),
        }
    }

    // ---- readouts ---------------------------------------------------------

    /// The canvas kind identifier.
    pub fn canvas_id(&self) -> String {
        from_dstr(self.inner.canvasId())
    }

    /// The backing provider's identifier.
    pub fn provider_id(&self) -> String {
        from_dstr(self.inner.providerId())
    }

    /// The canvas title, if set.
    pub fn title(&self) -> Option<String> {
        from_core_opt_str(self.inner.title())
    }

    /// The current activity string, if set.
    pub fn activity(&self) -> Option<String> {
        from_core_opt_str(self.inner.activity())
    }

    /// The content URI, if set.
    pub fn content_uri(&self) -> Option<String> {
        from_core_opt_str(self.inner.contentUri())
    }

    /// Whether the snapshot is current.
    pub fn availability(&self) -> Availability {
        Availability::from_core(self.inner.availability())
    }
}

impl Default for Canvas {
    fn default() -> Canvas {
        Canvas::new()
    }
}

impl PartialEq for Canvas {
    fn eq(&self, other: &Canvas) -> bool {
        self.inner == other.inner
    }
}

impl Eq for Canvas {}

impl std::fmt::Debug for Canvas {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("Canvas")
            .field("canvas_id", &self.canvas_id())
            .field("provider_id", &self.provider_id())
            .field("title", &self.title())
            .field("activity", &self.activity())
            .field("content_uri", &self.content_uri())
            .field("availability", &self.availability())
            .finish()
    }
}
