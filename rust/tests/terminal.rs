//! Terminal channel behaviour, through the safe API.
//!
//! Every expectation here is transcribed from the core's own Dafny corpus
//! (`spec/terminal.dfy`, `Terminal.RunCorpus`) or from a law the core proves,
//! rather than from whatever this extraction happens to produce. Where the
//! Dafny fixture is cited by number, that number is the corpus fixture id.
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

use ahp_verified::channels::terminal::{Part, Terminal, TerminalAction as A};
use ahp_verified::{Json, Outcome};

fn unclassified(parts: &[Part]) -> Vec<&str> {
    parts
        .iter()
        .map(|p| match p {
            Part::Unclassified { value } => value.as_str(),
            Part::Command { output, .. } => output.as_str(),
        })
        .collect()
}

// ---------------------------------------------------------------- scalar fields

#[test]
fn scalar_transitions_land_where_the_corpus_says() {
    let t = Terminal::new();
    assert_eq!(t.apply(&A::CwdChanged("/tmp".into())).cwd().as_deref(), Some("/tmp"));
    assert_eq!(t.apply(&A::TitleChanged("zsh".into())).title(), "zsh");
    assert_eq!(t.apply(&A::Resized { cols: 80, rows: 24 }).size(), Some((80, 24)));
    assert_eq!(t.apply(&A::Exited(0)).exit_code(), Some(0));
}

#[test]
fn initial_state_has_no_optional_fields_set() {
    let t = Terminal::new();
    assert_eq!(t.title(), "bash", "T0 seeds the title, per spec/terminal.dfy T0()");
    assert_eq!(t.cwd(), None);
    assert_eq!(t.size(), None);
    assert_eq!(t.exit_code(), None);
    assert_eq!(t.supports_command_detection(), None);
    assert_eq!(t.content_len(), 0);
}

// ------------------------------------------------------- the content classifier

#[test]
fn data_extends_a_trailing_unclassified_run_rather_than_appending() {
    // Dafny fixture: apply1(content := [Unclassified("a")], TData("b")).content
    //                == [Unclassified("ab")]
    let t = Terminal::new()
        .apply(&A::Data("a".into()))
        .apply(&A::Data("b".into()));

    assert_eq!(t.content_len(), 1, "two data actions must not make two parts");
    assert_eq!(unclassified(&t.content()), vec!["ab"]);
}

#[test]
fn data_flows_into_an_incomplete_commands_output() {
    // Dafny fixture: data appends to incomplete command output.
    let t = Terminal::new()
        .apply(&A::CommandExecuted {
            command_id: "cmd-1".into(),
            command_line: "npm test".into(),
            timestamp: 1_700_000_000_000,
        })
        .apply(&A::Data("All tests passed\r\n".into()))
        .apply(&A::Data("!".into()));

    assert_eq!(t.content_len(), 1, "data must join the open command, not start a new part");
    match &t.content()[0] {
        Part::Command { command_id, output, is_complete, .. } => {
            assert_eq!(command_id, "cmd-1");
            assert_eq!(output, "All tests passed\r\n!");
            assert!(!is_complete);
        }
        other => panic!("expected a Command part, got {other:?}"),
    }
}

#[test]
fn data_after_a_completed_command_starts_a_new_unclassified_part() {
    // Dafny fixture: data after completed command -> new unclassified.
    let t = Terminal::new()
        .apply(&A::CommandExecuted {
            command_id: "cmd-1".into(),
            command_line: "echo hi".into(),
            timestamp: 1_700_000_000_000,
        })
        .apply(&A::Data("hi\r\n".into()))
        .apply(&A::CommandFinished {
            command_id: "cmd-1".into(),
            code: 0,
            duration_ms: 50,
        })
        .apply(&A::Data("$ ".into()));

    assert_eq!(t.content_len(), 2, "a finished command must not absorb later data");
    assert!(matches!(&t.content()[1], Part::Unclassified { value } if value == "$ "));
}

#[test]
fn command_executed_opens_a_command_and_announces_detection() {
    // Dafny fixture: commandExecuted appends command part + sets detection.
    let t = Terminal::new().apply(&A::CommandExecuted {
        command_id: "cmd-1".into(),
        command_line: "npm test".into(),
        timestamp: 1_700_000_000_000,
    });

    assert_eq!(t.supports_command_detection(), Some(true));
    match &t.content()[0] {
        Part::Command { command_id, command_line, output, timestamp, is_complete, exit_code, duration_ms } => {
            assert_eq!(command_id, "cmd-1");
            assert_eq!(command_line, "npm test");
            assert_eq!(output, "");
            assert_eq!(*timestamp, 1_700_000_000_000);
            assert!(!is_complete);
            assert_eq!(*exit_code, None);
            assert_eq!(*duration_ms, None);
        }
        other => panic!("expected a Command part, got {other:?}"),
    }
}

#[test]
fn command_finished_completes_the_matching_id_only() {
    let t = Terminal::new()
        .apply(&A::CommandExecuted {
            command_id: "cmd-1".into(),
            command_line: "a".into(),
            timestamp: 1,
        })
        .apply(&A::CommandFinished {
            command_id: "cmd-1".into(),
            code: 0,
            duration_ms: 1234,
        })
        .apply(&A::CommandExecuted {
            command_id: "cmd-2".into(),
            command_line: "b".into(),
            timestamp: 2,
        })
        // finishing an id that is not open must not touch the open one
        .apply(&A::CommandFinished {
            command_id: "cmd-absent".into(),
            code: 9,
            duration_ms: 9,
        });

    let content = t.content();
    match (&content[0], &content[1]) {
        (
            Part::Command { is_complete: done, exit_code, duration_ms, .. },
            Part::Command { is_complete: open, .. },
        ) => {
            assert!(done, "cmd-1 was finished");
            assert_eq!(*exit_code, Some(0));
            assert_eq!(*duration_ms, Some(1234));
            assert!(!open, "cmd-2 must remain open after an unrelated finish");
        }
        other => panic!("expected two Command parts, got {other:?}"),
    }
}

#[test]
fn cleared_empties_content_but_keeps_detection_support() {
    // Dafny fixture: cleared empties (2 fixtures: plain + with command detection).
    let t = Terminal::new()
        .apply(&A::CommandDetectionAvailable)
        .apply(&A::Data("x".into()))
        .apply(&A::Cleared);

    assert_eq!(t.content_len(), 0);
    assert_eq!(
        t.supports_command_detection(),
        Some(true),
        "cleared must not retract the host's detection capability"
    );
}

// -------------------------------------------------------------------- no-ops

#[test]
fn input_and_unknown_leave_the_state_identical() {
    // Dafny fixtures: `apply1(base, TInput) == base` and the TUnknown analogue.
    let base = Terminal::new()
        .apply(&A::TitleChanged("zsh".into()))
        .apply(&A::Data("x".into()));

    assert_eq!(base.apply(&A::Input), base);
    assert_eq!(
        base.apply(&A::Unknown(Json::object([("type", Json::string("terminal/nope"))]))),
        base
    );
}

#[test]
fn no_op_actions_are_reported_as_no_ops_not_as_applied() {
    // The state being unchanged is not the whole contract: the reducer also
    // classifies the action. A wrapper that dropped the outcome would pass the
    // test above while losing this distinction.
    let base = Terminal::new();

    let (_, applied) = base.reduce(&A::TitleChanged("zsh".into()));
    assert_eq!(applied, Outcome::Applied);

    let (_, noop) = base.reduce(&A::Input);
    assert_eq!(noop, Outcome::NoOp);

    let (_, unknown) = base.reduce(&A::Unknown(Json::Null));
    assert_eq!(unknown, Outcome::NoOp);
}

// -------------------------------------------------------------------- purity

#[test]
fn transitions_do_not_mutate_the_receiver() {
    let before = Terminal::new().apply(&A::TitleChanged("one".into()));
    let snapshot = before.clone();

    let _ = before.apply(&A::TitleChanged("two".into()));
    let _ = before.apply(&A::Exited(3));
    let _ = before.apply(&A::Data("noise".into()));

    assert_eq!(before, snapshot, "the reducer is pure; the input must survive intact");
    assert_eq!(before.title(), "one");
    assert_eq!(before.exit_code(), None);
}

#[test]
fn clone_shares_state_and_stays_equal() {
    let a = Terminal::new().apply(&A::Data("x".into()));
    let b = a.clone();
    assert_eq!(a, b);
    let _ = b.apply(&A::Data("y".into()));
    assert_eq!(a, b, "applying to a clone must not disturb either value");
}

// ------------------------------------------------ the proven kernel fold

#[test]
fn apply_all_agrees_with_folding_apply() {
    // `apply_all` routes through the core's proven `Terminal.fold`
    // (ConfluxContract.Fold). This pins the two paths together: a wrapper that
    // quietly reimplemented the fold in Rust would have to keep agreeing.
    let actions = vec![
        A::CwdChanged("/w".into()),
        A::Data("x".into()),
        A::Resized { cols: 100, rows: 40 },
        A::CommandExecuted {
            command_id: "c".into(),
            command_line: "ls".into(),
            timestamp: 7,
        },
        A::Data("out".into()),
        A::CommandFinished {
            command_id: "c".into(),
            code: 0,
            duration_ms: 5,
        },
        A::Exited(1),
    ];

    let folded = Terminal::new().apply_all(&actions);
    let looped = actions.iter().fold(Terminal::new(), |s, a| s.apply(a));

    assert_eq!(folded, looped);
}

#[test]
fn empty_batch_is_the_identity() {
    let t = Terminal::new().apply(&A::Data("x".into()));
    assert_eq!(t.apply_all(&[]), t);
}

#[test]
fn full_lifecycle_matches_the_corpus_fixture() {
    // Dafny fixture: full TERMINAL lifecycle (real fixture): cwd -> data ->
    // resize -> claim -> exit.
    let claim = Json::object([
        ("kind", Json::string("session")),
        ("session", Json::string("s1")),
    ]);

    let life = Terminal::new().apply_all(&[
        A::CwdChanged("/w".into()),
        A::Data("x".into()),
        A::Resized { cols: 100, rows: 40 },
        A::Claimed(claim.clone()),
        A::Exited(1),
    ]);

    assert_eq!(life.cwd().as_deref(), Some("/w"));
    assert_eq!(life.size(), Some((100, 40)));
    assert_eq!(life.exit_code(), Some(1));
    assert_eq!(life.claim(), Some(claim));
}

#[test]
fn claimed_replaces_rather_than_merges() {
    // Dafny fixture: `apply1(base.(claim := Some(sc)), TClaimed(CL())).claim == Some(CL())`
    let first = Json::object([("kind", Json::string("session"))]);
    let second = Json::object([("kind", Json::string("client"))]);

    let t = Terminal::new()
        .apply(&A::Claimed(first))
        .apply(&A::Claimed(second.clone()));

    assert_eq!(t.claim(), Some(second));
}
