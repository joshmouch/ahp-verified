//! The Terminal channel.
//!
//! Models a terminal surface: title, working directory, dimensions, exit code,
//! and a classified content stream. The interesting proven behaviour is the
//! content classifier — appended data joins the trailing incomplete command's
//! output if there is one, else extends the trailing unclassified run, else
//! starts a new one. See [`TerminalAction::Data`].
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

use std::rc::Rc;

use crate::convert::{
    from_core_opt_bool, from_core_opt_int, from_core_opt_str, from_dint, from_dseq, from_dstr,
    to_dint, to_dseq, to_dstr,
};
use crate::generated::Terminal as core;
use crate::{Json, Outcome};

/// One classified region of the terminal's content stream.
#[derive(Debug, Clone, PartialEq, Eq)]
pub enum Part {
    /// Output not attributed to a detected command.
    Unclassified {
        /// The raw text.
        value: String,
    },
    /// Output attributed to a detected command.
    Command {
        /// The command's identifier, as reported by the host.
        command_id: String,
        /// The command line as executed.
        command_line: String,
        /// Output accumulated for this command so far.
        output: String,
        /// Host-supplied timestamp of execution.
        timestamp: i64,
        /// Whether the host has reported the command finished.
        is_complete: bool,
        /// Exit code, once finished.
        exit_code: Option<i64>,
        /// Duration in milliseconds, once finished.
        duration_ms: Option<i64>,
    },
}

impl Part {
    fn from_core(p: &core::Part) -> Part {
        match p {
            core::Part::Unclassified { value } => Part::Unclassified {
                value: from_dstr(value),
            },
            core::Part::Command {
                commandId,
                commandLine,
                output,
                timestamp,
                isComplete,
                exitCode,
                durationMs,
            } => Part::Command {
                command_id: from_dstr(commandId),
                command_line: from_dstr(commandLine),
                output: from_dstr(output),
                timestamp: from_dint(timestamp),
                is_complete: *isComplete,
                exit_code: from_core_opt_int(exitCode),
                duration_ms: from_core_opt_int(durationMs),
            },
        }
    }
}

/// An action the Terminal reducer understands.
#[derive(Debug, Clone, PartialEq, Eq)]
pub enum TerminalAction {
    /// The working directory changed.
    CwdChanged(String),
    /// The title changed.
    TitleChanged(String),
    /// The terminal was resized.
    Resized {
        /// New column count.
        cols: i64,
        /// New row count.
        rows: i64,
    },
    /// The terminal's process exited.
    Exited(i64),
    /// Output data arrived.
    ///
    /// Appending is classifier-driven, not a plain concatenation: see the
    /// module docs and [`Terminal::content`].
    Data(String),
    /// The terminal was cleared.
    Cleared,
    /// An opaque claim payload was attached.
    Claimed(Json),
    /// The host reported that it can detect command boundaries.
    CommandDetectionAvailable,
    /// A command started.
    CommandExecuted {
        /// The command's identifier.
        command_id: String,
        /// The command line as executed.
        command_line: String,
        /// Host-supplied timestamp.
        timestamp: i64,
    },
    /// A command finished.
    CommandFinished {
        /// The command's identifier.
        command_id: String,
        /// Exit code.
        code: i64,
        /// Duration in milliseconds.
        duration_ms: i64,
    },
    /// User input. Defined as a no-op on state.
    Input,
    /// An action this channel does not recognize. Defined as a no-op.
    Unknown(Json),
}

impl TerminalAction {
    fn to_core(&self) -> Rc<core::TerminalAction> {
        Rc::new(match self {
            TerminalAction::CwdChanged(cwd) => core::TerminalAction::TCwdChanged {
                cwd: to_dstr(cwd),
            },
            TerminalAction::TitleChanged(title) => core::TerminalAction::TTitleChanged {
                title: to_dstr(title),
            },
            TerminalAction::Resized { cols, rows } => core::TerminalAction::TResized {
                cols: to_dint(*cols),
                rows: to_dint(*rows),
            },
            TerminalAction::Exited(code) => core::TerminalAction::TExited {
                code: to_dint(*code),
            },
            TerminalAction::Data(data) => core::TerminalAction::TData {
                data: to_dstr(data),
            },
            TerminalAction::Cleared => core::TerminalAction::TCleared {},
            TerminalAction::Claimed(claim) => core::TerminalAction::TClaimed {
                claim: claim.to_core(),
            },
            TerminalAction::CommandDetectionAvailable => {
                core::TerminalAction::TCommandDetectionAvailable {}
            }
            TerminalAction::CommandExecuted {
                command_id,
                command_line,
                timestamp,
            } => core::TerminalAction::TCommandExecuted {
                commandId: to_dstr(command_id),
                commandLine: to_dstr(command_line),
                timestamp: to_dint(*timestamp),
            },
            TerminalAction::CommandFinished {
                command_id,
                code,
                duration_ms,
            } => core::TerminalAction::TCommandFinished {
                commandId: to_dstr(command_id),
                code: to_dint(*code),
                durationMs: to_dint(*duration_ms),
            },
            TerminalAction::Input => core::TerminalAction::TInput {},
            TerminalAction::Unknown(raw) => core::TerminalAction::TUnknown {
                raw: raw.to_core(),
            },
        })
    }
}

/// A Terminal channel state.
///
/// Values are immutable; every transition returns a new `Terminal`. Cloning is
/// cheap (the underlying state is reference-counted).
#[derive(Clone)]
pub struct Terminal {
    inner: Rc<core::TerminalState>,
}

impl Terminal {
    /// The initial state, as defined by the core's `Terminal.T0()`.
    pub fn new() -> Terminal {
        Terminal {
            inner: core::_default::T0(),
        }
    }

    /// Apply one action through the verified reducer, returning the new state.
    ///
    /// The receiver is not modified — the reducer is a pure function.
    pub fn apply(&self, action: &TerminalAction) -> Terminal {
        Terminal {
            inner: core::_default::apply1(&self.inner, &action.to_core()),
        }
    }

    /// Apply one action, also returning the reducer's [`Outcome`] for it.
    ///
    /// Use this to distinguish "the reducer processed this" from "the reducer
    /// recognized this and defines it as a no-op" — a distinction the plain
    /// [`apply`][Terminal::apply] discards.
    pub fn reduce(&self, action: &TerminalAction) -> (Terminal, Outcome) {
        // `now` is the reducer's injected clock. The Terminal reducer's
        // transitions do not read it (the core's own `apply1` passes a fixed
        // 9999 for exactly this reason); it is threaded through for uniformity
        // with clock-sensitive channels.
        let (state, outcome) =
            core::_default::ApplyToTerminal(&self.inner, &action.to_core(), &to_dint(9999));
        (Terminal { inner: state }, Outcome::from_core(&outcome))
    }

    /// Apply a batch of actions through the core's **proven kernel fold**
    /// (`Terminal.fold`, defined over `ConfluxContract.Fold`).
    ///
    /// Equivalent to folding [`apply`][Terminal::apply] over the slice, but the
    /// sequencing happens inside the verified core rather than here.
    pub fn apply_all(&self, actions: &[TerminalAction]) -> Terminal {
        let seq = to_dseq(actions.iter().map(TerminalAction::to_core).collect());
        Terminal {
            inner: core::_default::fold(&self.inner, &seq),
        }
    }

    // ---- readouts ---------------------------------------------------------

    /// The terminal title.
    pub fn title(&self) -> String {
        from_dstr(self.inner.title())
    }

    /// The working directory, if the host has reported one.
    pub fn cwd(&self) -> Option<String> {
        from_core_opt_str(self.inner.cwd())
    }

    /// The dimensions as `(cols, rows)`, if the host has reported them.
    pub fn size(&self) -> Option<(i64, i64)> {
        match (
            from_core_opt_int(self.inner.cols()),
            from_core_opt_int(self.inner.rows()),
        ) {
            (Some(c), Some(r)) => Some((c, r)),
            _ => None,
        }
    }

    /// The process exit code, if the process has exited.
    pub fn exit_code(&self) -> Option<i64> {
        from_core_opt_int(self.inner.exitCode())
    }

    /// Whether the host has reported command-detection support.
    pub fn supports_command_detection(&self) -> Option<bool> {
        from_core_opt_bool(self.inner.supportsCommandDetection())
    }

    /// The opaque claim payload, if one was attached.
    pub fn claim(&self) -> Option<Json> {
        crate::convert::opt_ref(self.inner.claim()).map(|j| Json::from_core(j))
    }

    /// The classified content stream.
    pub fn content(&self) -> Vec<Part> {
        from_dseq(self.inner.content(), |p| Part::from_core(p))
    }

    /// The number of parts in the content stream, without materialising them.
    pub fn content_len(&self) -> usize {
        self.inner.content().cardinality_usize()
    }
}

impl Default for Terminal {
    fn default() -> Terminal {
        Terminal::new()
    }
}

impl PartialEq for Terminal {
    fn eq(&self, other: &Terminal) -> bool {
        self.inner == other.inner
    }
}

impl Eq for Terminal {}

impl std::fmt::Debug for Terminal {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("Terminal")
            .field("title", &self.title())
            .field("cwd", &self.cwd())
            .field("size", &self.size())
            .field("exit_code", &self.exit_code())
            .field("content_len", &self.content_len())
            .finish()
    }
}
