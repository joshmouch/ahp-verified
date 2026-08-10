//! Drives a terminal session through the verified reducers and prints the
//! resulting state, then runs the core's own eight-channel conformance corpus.
//!
//! Run with:
//!
//! ```text
//! cargo run --example terminal_session
//! ```
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

use ahp_verified::channels::terminal::{Part, Terminal, TerminalAction as A};
use ahp_verified::{Json, Outcome};

fn main() {
    // ---- a realistic session, applied one action at a time ----------------
    let session = Terminal::new().apply_all(&[
        A::TitleChanged("zsh — ahp-verified".into()),
        A::CwdChanged("/Users/dev/project".into()),
        A::Resized { cols: 120, rows: 40 },
        A::CommandDetectionAvailable,
        A::CommandExecuted {
            command_id: "cmd-1".into(),
            command_line: "cargo test".into(),
            timestamp: 1_700_000_000_000,
        },
        A::Data("   Compiling ahp-verified v0.1.0\n".into()),
        A::Data("    Finished test profile\n".into()),
        A::CommandFinished {
            command_id: "cmd-1".into(),
            code: 0,
            duration_ms: 1234,
        },
        A::Data("$ ".into()),
        A::Claimed(Json::object([
            ("kind", Json::string("session")),
            ("session", Json::string("s1")),
        ])),
    ]);

    println!("== terminal state ==");
    println!("  title      : {}", session.title());
    println!("  cwd        : {:?}", session.cwd());
    println!("  size       : {:?}", session.size());
    println!("  detection  : {:?}", session.supports_command_detection());
    println!("  exit code  : {:?}", session.exit_code());

    println!("\n== classified content ({} parts) ==", session.content_len());
    for (i, part) in session.content().iter().enumerate() {
        match part {
            Part::Unclassified { value } => {
                println!("  [{i}] unclassified: {value:?}");
            }
            Part::Command {
                command_id,
                command_line,
                output,
                is_complete,
                exit_code,
                duration_ms,
                ..
            } => {
                println!("  [{i}] command {command_id}: {command_line:?}");
                println!("        output   : {output:?}");
                println!("        complete : {is_complete} (exit {exit_code:?}, {duration_ms:?}ms)");
            }
        }
    }

    // ---- the reducer also classifies the action --------------------------
    println!("\n== action outcomes ==");
    for action in [
        A::TitleChanged("later".into()),
        A::Input,
        A::Unknown(Json::object([("type", Json::string("terminal/nope"))])),
    ] {
        let (next, outcome) = session.reduce(&action);
        let changed = if next == session { "state unchanged" } else { "state changed" };
        let label = match outcome {
            Outcome::Applied => "Applied",
            Outcome::NoOp => "NoOp",
            Outcome::OutOfScope => "OutOfScope",
        };
        println!("  {label:<10} {changed:<16} <- {action:?}");
    }

    // ---- reducers are pure ------------------------------------------------
    let exited = session.apply(&A::Exited(0));
    println!("\n== purity ==");
    println!("  original exit code : {:?}", session.exit_code());
    println!("  derived  exit code : {:?}", exited.exit_code());

    // ---- and the core's own corpus, across all eight channels -------------
    println!("\n== core conformance corpus ==");
    ahp_verified::corpus::run();
}
