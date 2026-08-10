//! The Changeset channel.
//!
//! Models a proposed set of file edits plus the operations computing them.
//! Files and operations are both id-keyed, order-preserving collections routed
//! through the same proven keyed algebra the Annotations channel uses.
//!
//! Note that `operations` is `Option<Vec<..>>`: the core distinguishes "no
//! operations" from "operations not reported", and the reducer's transitions
//! respect that distinction rather than collapsing it to an empty list.
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

use std::rc::Rc;

use crate::convert::{
    from_core_opt_bool, from_dseq, from_dseq_str, from_dstr, opt_ref, to_core_opt, to_dint,
    to_dseq, to_dseq_str, to_dstr,
};
use crate::generated::Changeset as core;
use crate::{Json, Outcome};

/// One file in a changeset.
#[derive(Debug, Clone, PartialEq, Eq)]
pub struct File {
    /// File identifier, unique within the changeset.
    pub id: String,
    /// Review state: `None` means not yet reviewed either way.
    pub reviewed: Option<bool>,
    /// Opaque edit payload.
    pub edit: Json,
}

impl File {
    /// A file with an unset review state.
    pub fn new(id: impl Into<String>, edit: Json) -> File {
        File {
            id: id.into(),
            reviewed: None,
            edit,
        }
    }

    fn to_core(&self) -> Rc<core::ChangesetFile> {
        Rc::new(core::ChangesetFile::ChangesetFile {
            id: to_dstr(&self.id),
            reviewed: to_core_opt(self.reviewed),
            edit: self.edit.to_core(),
        })
    }

    fn from_core(f: &core::ChangesetFile) -> File {
        File {
            id: from_dstr(f.id()),
            reviewed: from_core_opt_bool(f.reviewed()),
            edit: Json::from_core(f.edit()),
        }
    }
}

/// One operation contributing to a changeset.
#[derive(Debug, Clone, PartialEq, Eq)]
pub struct Operation {
    /// Operation identifier, unique within the changeset.
    pub id: String,
    /// Human-readable label.
    pub label: String,
    /// Scopes the operation applies to.
    pub scopes: Vec<String>,
    /// Operation status.
    pub status: String,
    /// Opaque error payload, if the operation failed.
    pub error: Option<Json>,
}

impl Operation {
    /// An operation with no scopes and no error.
    pub fn new(
        id: impl Into<String>,
        label: impl Into<String>,
        status: impl Into<String>,
    ) -> Operation {
        Operation {
            id: id.into(),
            label: label.into(),
            scopes: Vec::new(),
            status: status.into(),
            error: None,
        }
    }

    fn to_core(&self) -> Rc<core::ChangesetOperation> {
        Rc::new(core::ChangesetOperation::ChangesetOperation {
            id: to_dstr(&self.id),
            label_: to_dstr(&self.label),
            scopes: to_dseq_str(&self.scopes),
            status: to_dstr(&self.status),
            error: to_core_opt(self.error.as_ref().map(Json::to_core)),
        })
    }

    fn from_core(o: &core::ChangesetOperation) -> Operation {
        Operation {
            id: from_dstr(o.id()),
            label: from_dstr(o.label_()),
            scopes: from_dseq_str(o.scopes()),
            status: from_dstr(o.status()),
            error: opt_ref(o.error()).map(|j| Json::from_core(j)),
        }
    }
}

type CoreOpSeq = crate::dafny_runtime::Sequence<Rc<core::ChangesetOperation>>;
type CoreFileSeq = crate::dafny_runtime::Sequence<Rc<core::ChangesetFile>>;
type CoreOpt<T> = Rc<crate::generated::AhpSkeleton::Option<T>>;

fn to_core_ops(ops: &Option<Vec<Operation>>) -> CoreOpt<CoreOpSeq> {
    to_core_opt(
        ops.as_ref()
            .map(|v| to_dseq(v.iter().map(Operation::to_core).collect())),
    )
}

fn to_core_files(files: &Option<Vec<File>>) -> CoreOpt<CoreFileSeq> {
    to_core_opt(
        files
            .as_ref()
            .map(|v| to_dseq(v.iter().map(File::to_core).collect())),
    )
}

/// An action the Changeset reducer understands.
#[derive(Debug, Clone, PartialEq, Eq)]
pub enum ChangesetAction {
    /// The changeset's overall status changed.
    StatusChanged {
        /// New status.
        status: String,
        /// Opaque error payload, if the new status is a failure.
        error: Option<Json>,
    },
    /// Insert or update a file, keyed by id, preserving order.
    FileSet(File),
    /// Remove a file by id.
    FileRemoved(String),
    /// Replace the operations list wholesale.
    OperationsChanged(Option<Vec<Operation>>),
    /// Clear the changeset.
    Cleared,
    /// Update one operation's status.
    OperationStatusChanged {
        /// Target operation.
        operation_id: String,
        /// New status.
        status: String,
        /// Opaque error payload, if the new status is a failure.
        error: Option<Json>,
    },
    /// Replace files and/or operations, and set an error payload.
    ContentChanged {
        /// New file list, or `None` to leave unchanged.
        files: Option<Vec<File>>,
        /// New operation list, or `None` to leave unchanged.
        operations: Option<Vec<Operation>>,
        /// New error payload.
        error: Option<Json>,
    },
    /// Set the review flag on a batch of files at once.
    FilesReviewedChanged {
        /// Files to mark.
        file_ids: Vec<String>,
        /// The review state to set.
        reviewed: bool,
    },
    /// An action this channel does not recognize. Defined as a no-op.
    Unknown(Json),
}

impl ChangesetAction {
    fn to_core(&self) -> Rc<core::ChangesetAction> {
        Rc::new(match self {
            ChangesetAction::StatusChanged { status, error } => {
                core::ChangesetAction::StatusChanged {
                    status: to_dstr(status),
                    error: to_core_opt(error.as_ref().map(Json::to_core)),
                }
            }
            ChangesetAction::FileSet(f) => core::ChangesetAction::FileSet { file: f.to_core() },
            ChangesetAction::FileRemoved(id) => core::ChangesetAction::FileRemoved {
                fileId: to_dstr(id),
            },
            ChangesetAction::OperationsChanged(ops) => core::ChangesetAction::OperationsChanged {
                operations: to_core_ops(ops),
            },
            ChangesetAction::Cleared => core::ChangesetAction::Cleared {},
            ChangesetAction::OperationStatusChanged {
                operation_id,
                status,
                error,
            } => core::ChangesetAction::OperationStatusChanged {
                operationId: to_dstr(operation_id),
                status: to_dstr(status),
                error: to_core_opt(error.as_ref().map(Json::to_core)),
            },
            ChangesetAction::ContentChanged {
                files,
                operations,
                error,
            } => core::ChangesetAction::ContentChanged {
                files: to_core_files(files),
                operations: to_core_ops(operations),
                error: to_core_opt(error.as_ref().map(Json::to_core)),
            },
            ChangesetAction::FilesReviewedChanged { file_ids, reviewed } => {
                core::ChangesetAction::FilesReviewedChanged {
                    fileIds: to_dseq_str(file_ids),
                    reviewed: *reviewed,
                }
            }
            ChangesetAction::Unknown(raw) => core::ChangesetAction::CsUnknown {
                raw: raw.to_core(),
            },
        })
    }
}

/// A Changeset channel state.
#[derive(Clone)]
pub struct Changeset {
    inner: Rc<core::ChangesetState>,
}

impl Changeset {
    /// A changeset in status `"idle"` with no files and no reported operations.
    ///
    /// This mirrors the well-formed witness the core hands its aggregate
    /// (`Changeset.ChangesetWfWitness`).
    pub fn new() -> Changeset {
        Changeset::with_status("idle")
    }

    /// A changeset in the given status, with no files and no reported operations.
    pub fn with_status(status: impl AsRef<str>) -> Changeset {
        Changeset {
            inner: Rc::new(core::ChangesetState::ChangesetState {
                status: to_dstr(status.as_ref()),
                files: to_dseq(Vec::new()),
                operations: to_core_opt(None),
                error: to_core_opt(None),
            }),
        }
    }

    /// Apply one action through the verified reducer.
    pub fn apply(&self, action: &ChangesetAction) -> Changeset {
        Changeset {
            inner: core::_default::apply1(&self.inner, &action.to_core()),
        }
    }

    /// Apply one action, also returning the reducer's [`Outcome`].
    pub fn reduce(&self, action: &ChangesetAction) -> (Changeset, Outcome) {
        let (state, outcome) =
            core::_default::ApplyToChangeset(&self.inner, &action.to_core(), &to_dint(9999));
        (Changeset { inner: state }, Outcome::from_core(&outcome))
    }

    /// Apply a batch of actions through the core's proven kernel fold.
    pub fn apply_all(&self, actions: &[ChangesetAction]) -> Changeset {
        let seq = to_dseq(actions.iter().map(ChangesetAction::to_core).collect());
        Changeset {
            inner: core::_default::fold(&self.inner, &seq),
        }
    }

    // ---- readouts ---------------------------------------------------------

    /// The changeset's overall status.
    pub fn status(&self) -> String {
        from_dstr(self.inner.status())
    }

    /// The files in the changeset, in order.
    pub fn files(&self) -> Vec<File> {
        from_dseq(self.inner.files(), |f| File::from_core(f))
    }

    /// The operations, if the host has reported any list at all.
    ///
    /// `None` means "not reported"; `Some(vec![])` means "reported, and empty".
    pub fn operations(&self) -> Option<Vec<Operation>> {
        opt_ref(self.inner.operations())
            .map(|seq| from_dseq(seq, |o| Operation::from_core(o)))
    }

    /// The opaque error payload, if the changeset is in a failure state.
    pub fn error(&self) -> Option<Json> {
        opt_ref(self.inner.error()).map(|j| Json::from_core(j))
    }

    /// The file with the given id, if present.
    pub fn file(&self, id: &str) -> Option<File> {
        self.files().into_iter().find(|f| f.id == id)
    }
}

impl Default for Changeset {
    fn default() -> Changeset {
        Changeset::new()
    }
}

impl PartialEq for Changeset {
    fn eq(&self, other: &Changeset) -> bool {
        self.inner == other.inner
    }
}

impl Eq for Changeset {}

impl std::fmt::Debug for Changeset {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("Changeset")
            .field("status", &self.status())
            .field("files", &self.files())
            .field("operations", &self.operations())
            .field("error", &self.error())
            .finish()
    }
}
