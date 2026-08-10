//! The Annotations channel.
//!
//! Models a keyed collection of annotations, each holding a keyed collection of
//! entries. Both levels are **order-preserving upserts by id**: setting an
//! existing id updates in place rather than moving it to the end, and the core
//! proves key uniqueness is preserved across every transition.
//!
//! Operations naming an unknown id are proven no-ops, not errors — see
//! [`AnnotationsAction::EntrySet`].
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

use std::rc::Rc;

use crate::convert::{
    from_dseq, from_dstr, opt_ref, to_core_opt, to_core_opt_str, to_dint, to_dseq, to_dstr,
};
use crate::generated::Annotations as core;
use crate::{Json, Outcome};

/// One entry attached to an annotation.
#[derive(Debug, Clone, PartialEq, Eq)]
pub struct Entry {
    /// Entry identifier, unique within its annotation.
    pub id: String,
    /// Entry body text.
    pub text: String,
    /// Opaque host metadata.
    pub meta: Option<Json>,
}

impl Entry {
    /// An entry with no metadata.
    pub fn new(id: impl Into<String>, text: impl Into<String>) -> Entry {
        Entry {
            id: id.into(),
            text: text.into(),
            meta: None,
        }
    }

    fn to_core(&self) -> Rc<core::Entry> {
        Rc::new(core::Entry::Entry {
            id: to_dstr(&self.id),
            text: to_dstr(&self.text),
            meta: to_core_opt(self.meta.as_ref().map(Json::to_core)),
        })
    }

    fn from_core(e: &core::Entry) -> Entry {
        Entry {
            id: from_dstr(e.id()),
            text: from_dstr(e.text()),
            meta: opt_ref(e.meta()).map(|j| Json::from_core(j)),
        }
    }
}

/// One annotation.
#[derive(Debug, Clone, PartialEq, Eq)]
pub struct Annotation {
    /// Annotation identifier, unique within the channel.
    pub id: String,
    /// The turn this annotation belongs to.
    pub turn_id: String,
    /// The resource the annotation targets.
    pub resource: String,
    /// Opaque range payload within the resource.
    pub range: Option<Json>,
    /// Whether the annotation has been resolved.
    pub resolved: bool,
    /// Entries attached to this annotation, in order.
    pub entries: Vec<Entry>,
    /// Opaque host metadata.
    pub meta: Option<Json>,
}

impl Annotation {
    /// An unresolved annotation with no entries, range, or metadata.
    pub fn new(
        id: impl Into<String>,
        turn_id: impl Into<String>,
        resource: impl Into<String>,
    ) -> Annotation {
        Annotation {
            id: id.into(),
            turn_id: turn_id.into(),
            resource: resource.into(),
            range: None,
            resolved: false,
            entries: Vec::new(),
            meta: None,
        }
    }

    fn to_core(&self) -> Rc<core::Annotation> {
        Rc::new(core::Annotation::Annotation {
            id: to_dstr(&self.id),
            turnId: to_dstr(&self.turn_id),
            resource: to_dstr(&self.resource),
            range: to_core_opt(self.range.as_ref().map(Json::to_core)),
            resolved: self.resolved,
            entries: to_dseq(self.entries.iter().map(Entry::to_core).collect()),
            meta: to_core_opt(self.meta.as_ref().map(Json::to_core)),
        })
    }

    fn from_core(a: &core::Annotation) -> Annotation {
        Annotation {
            id: from_dstr(a.id()),
            turn_id: from_dstr(a.turnId()),
            resource: from_dstr(a.resource()),
            range: opt_ref(a.range()).map(|j| Json::from_core(j)),
            resolved: *a.resolved(),
            entries: from_dseq(a.entries(), |e| Entry::from_core(e)),
            meta: opt_ref(a.meta()).map(|j| Json::from_core(j)),
        }
    }
}

/// A partial update to an existing annotation. `None` fields are preserved.
#[derive(Debug, Clone, PartialEq, Eq, Default)]
pub struct AnnotationUpdate {
    /// New turn id.
    pub turn_id: Option<String>,
    /// New target resource.
    pub resource: Option<String>,
    /// New range payload.
    pub range: Option<Json>,
    /// New resolved flag.
    pub resolved: Option<bool>,
}

/// An action the Annotations reducer understands.
#[derive(Debug, Clone, PartialEq, Eq)]
pub enum AnnotationsAction {
    /// Insert or update an annotation, keyed by its id, preserving order.
    Set(Annotation),
    /// Remove an annotation by id.
    Removed(String),
    /// Insert or update an entry within an annotation.
    ///
    /// A no-op — proven, not merely conventional — when `annotation_id` names
    /// no existing annotation.
    EntrySet {
        /// Target annotation.
        annotation_id: String,
        /// The entry to insert or update.
        entry: Entry,
    },
    /// Remove an entry from an annotation.
    EntryRemoved {
        /// Target annotation.
        annotation_id: String,
        /// Entry to remove.
        entry_id: String,
    },
    /// Partially update an existing annotation's fields.
    Updated {
        /// Target annotation.
        annotation_id: String,
        /// The fields to change.
        update: AnnotationUpdate,
    },
    /// An action this channel does not recognize. Defined as a no-op.
    Unknown(Json),
}

impl AnnotationsAction {
    fn to_core(&self) -> Rc<core::AnnotationsAction> {
        Rc::new(match self {
            AnnotationsAction::Set(a) => core::AnnotationsAction::Set {
                annotation: a.to_core(),
            },
            AnnotationsAction::Removed(id) => core::AnnotationsAction::Removed {
                annotationId: to_dstr(id),
            },
            AnnotationsAction::EntrySet {
                annotation_id,
                entry,
            } => core::AnnotationsAction::EntrySet {
                annotationId: to_dstr(annotation_id),
                entry: entry.to_core(),
            },
            AnnotationsAction::EntryRemoved {
                annotation_id,
                entry_id,
            } => core::AnnotationsAction::EntryRemoved {
                annotationId: to_dstr(annotation_id),
                entryId: to_dstr(entry_id),
            },
            AnnotationsAction::Updated {
                annotation_id,
                update,
            } => core::AnnotationsAction::Updated {
                annotationId: to_dstr(annotation_id),
                turnId: to_core_opt_str(update.turn_id.as_deref()),
                resource: to_core_opt_str(update.resource.as_deref()),
                range: to_core_opt(update.range.as_ref().map(Json::to_core)),
                resolved: to_core_opt(update.resolved),
            },
            AnnotationsAction::Unknown(raw) => core::AnnotationsAction::AnUnknown {
                raw: raw.to_core(),
            },
        })
    }
}

/// An Annotations channel state.
#[derive(Clone)]
pub struct Annotations {
    inner: Rc<core::AnnotationsState>,
}

impl Annotations {
    /// The empty annotation set.
    pub fn new() -> Annotations {
        Annotations {
            inner: Rc::new(core::AnnotationsState::AnnotationsState {
                annotations: to_dseq(Vec::new()),
            }),
        }
    }

    /// Apply one action through the verified reducer.
    pub fn apply(&self, action: &AnnotationsAction) -> Annotations {
        Annotations {
            inner: core::_default::apply1(&self.inner, &action.to_core()),
        }
    }

    /// Apply one action, also returning the reducer's [`Outcome`].
    pub fn reduce(&self, action: &AnnotationsAction) -> (Annotations, Outcome) {
        let (state, outcome) =
            core::_default::ApplyToAnnotations(&self.inner, &action.to_core(), &to_dint(9999));
        (Annotations { inner: state }, Outcome::from_core(&outcome))
    }

    /// Apply a batch of actions through the core's proven kernel fold.
    pub fn apply_all(&self, actions: &[AnnotationsAction]) -> Annotations {
        let seq = to_dseq(actions.iter().map(AnnotationsAction::to_core).collect());
        Annotations {
            inner: core::_default::fold(&self.inner, &seq),
        }
    }

    // ---- readouts ---------------------------------------------------------

    /// All annotations, in insertion order.
    pub fn annotations(&self) -> Vec<Annotation> {
        from_dseq(self.inner.annotations(), |a| Annotation::from_core(a))
    }

    /// The annotation with the given id, if present.
    pub fn get(&self, id: &str) -> Option<Annotation> {
        self.annotations().into_iter().find(|a| a.id == id)
    }

    /// How many annotations the state holds.
    pub fn len(&self) -> usize {
        self.inner.annotations().cardinality_usize()
    }

    /// Whether the state holds no annotations.
    pub fn is_empty(&self) -> bool {
        self.len() == 0
    }
}

impl Default for Annotations {
    fn default() -> Annotations {
        Annotations::new()
    }
}

impl PartialEq for Annotations {
    fn eq(&self, other: &Annotations) -> bool {
        self.inner == other.inner
    }
}

impl Eq for Annotations {}

impl std::fmt::Debug for Annotations {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("Annotations")
            .field("annotations", &self.annotations())
            .finish()
    }
}
