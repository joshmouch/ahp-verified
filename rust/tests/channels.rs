//! Canvas, ResourceWatch, Annotations and Changeset behaviour.
//!
//! As with the terminal suite, expectations are transcribed from each channel's
//! Dafny corpus (`RunCorpus` in the corresponding `spec/*.dfy`) rather than
//! from observed output.
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

use ahp_verified::channels::annotations::{
    Annotation, Annotations, AnnotationsAction as An, AnnotationUpdate, Entry,
};
use ahp_verified::channels::canvas::{Availability, Canvas, CanvasAction as C, Updated};
use ahp_verified::channels::changeset::{
    Changeset, ChangesetAction as Cs, File, Operation,
};
use ahp_verified::channels::resource_watch::{ResourceWatch, ResourceWatchAction as Rw};
use ahp_verified::{Json, Outcome};

// ================================================================== canvas

#[test]
fn canvas_partial_update_preserves_untouched_fields() {
    // Dafny fixture: `Updated(Some("Renamed"), None, Some(uri), None)` changes
    // exactly title and contentUri, leaving activity and availability alone.
    let base = Canvas::new().apply(&C::Updated(Updated {
        title: Some("Draft".into()),
        activity: Some("idle".into()),
        content_uri: Some("ahp-session:/2f9c/content/canvas-1".into()),
        availability: Some(Availability::Ready),
    }));

    let renamed = base.apply(&C::Updated(Updated {
        title: Some("Renamed".into()),
        content_uri: Some("https://example.com/docs/renamed.html".into()),
        ..Default::default()
    }));

    assert_eq!(renamed.title().as_deref(), Some("Renamed"));
    assert_eq!(
        renamed.content_uri().as_deref(),
        Some("https://example.com/docs/renamed.html")
    );
    // The load-bearing half: None means "leave alone", not "clear".
    assert_eq!(renamed.activity().as_deref(), Some("idle"));
    assert_eq!(renamed.availability(), Availability::Ready);
}

#[test]
fn canvas_update_can_touch_only_the_right_hand_fields() {
    // Dafny fixture: `Updated(None, Some("error"), None, Some(Stale))`.
    let base = Canvas::new().apply(&C::Updated(Updated {
        title: Some("Draft".into()),
        activity: Some("idle".into()),
        ..Default::default()
    }));

    let errored = base.apply(&C::Updated(Updated {
        activity: Some("error".into()),
        availability: Some(Availability::Stale),
        ..Default::default()
    }));

    assert_eq!(errored.activity().as_deref(), Some("error"));
    assert_eq!(errored.availability(), Availability::Stale);
    assert_eq!(errored.title().as_deref(), Some("Draft"));
}

#[test]
fn canvas_close_and_unknown_are_no_ops() {
    // Dafny fixtures: CloseRequested and CanvasUnknown both return `s`.
    let base = Canvas::new().apply(&C::title("Draft"));

    assert_eq!(base.apply(&C::CloseRequested), base);
    assert_eq!(
        base.apply(&C::Unknown(Json::object([(
            "type",
            Json::string("canvas/nonExistentAction")
        )]))),
        base
    );
}

#[test]
fn canvas_identity_fields_are_not_reachable_by_any_action() {
    // The core's canvas laws state that instance identity is the channel URI
    // alone, so no snapshot transition may alter canvas_id or provider_id.
    let base = Canvas::new();
    let (id, provider) = (base.canvas_id(), base.provider_id());

    let churned = base.apply_all(&[
        C::title("a"),
        C::availability(Availability::Stale),
        C::Updated(Updated {
            content_uri: Some("x".into()),
            activity: Some("y".into()),
            ..Default::default()
        }),
        C::CloseRequested,
    ]);

    assert_eq!(churned.canvas_id(), id);
    assert_eq!(churned.provider_id(), provider);
}

// =========================================================== resource watch

#[test]
fn resource_watch_is_a_passthrough_but_still_classifies() {
    // The core proves both actions leave the state equal to their input, so the
    // outcome is the only observable difference. A wrapper that returned only
    // the state would make these two indistinguishable.
    let w = ResourceWatch::new("file:///workspace", true);

    let changes = Json::object([(
        "items",
        Json::array([Json::object([
            ("uri", Json::string("file:///workspace/a.txt")),
            ("type", Json::string("added")),
        ])]),
    )]);

    let (after_changed, changed_outcome) = w.reduce(&Rw::Changed(changes));
    let (after_unknown, unknown_outcome) = w.reduce(&Rw::Unknown(Json::Null));

    assert_eq!(after_changed, w, "changed is a passthrough");
    assert_eq!(after_unknown, w, "unknown is a passthrough");
    assert_eq!(changed_outcome, Outcome::Applied);
    assert_eq!(unknown_outcome, Outcome::NoOp);
}

#[test]
fn resource_watch_config_survives_a_batch() {
    let w = ResourceWatch::new("file:///workspace", false);
    let after = w.apply_all(&[
        Rw::Changed(Json::Null),
        Rw::Unknown(Json::Null),
        Rw::Changed(Json::Null),
    ]);

    assert_eq!(after.root(), "file:///workspace");
    assert!(!after.recursive());
    assert_eq!(after, w);
}

// ============================================================== annotations

fn a1() -> Annotation {
    Annotation {
        entries: vec![Entry::new("c-1", "original")],
        range: Some(Json::object([("start", Json::Null)])),
        ..Annotation::new("t-1", "turn-1", "file:///src/a.ts")
    }
}

#[test]
fn annotations_set_appends_a_new_id_and_replaces_an_existing_one() {
    // Dafny fixtures 210 (append) and 211 (replace).
    let st = Annotations::new().apply(&An::Set(a1()));

    let a2 = Annotation {
        entries: vec![Entry::new("c-2", "x")],
        ..Annotation::new("t-2", "turn-2", "file:///src/b.ts")
    };
    let appended = st.apply(&An::Set(a2.clone()));
    assert_eq!(appended.len(), 2);
    assert_eq!(appended.annotations()[1], a2);

    // Replacing t-1 wholesale drops its range and flips resolved.
    let a1r = Annotation {
        range: None,
        resolved: true,
        ..a1()
    };
    let replaced = st.apply(&An::Set(a1r.clone()));
    assert_eq!(replaced.len(), 1, "set on an existing id replaces, not appends");
    assert_eq!(replaced.annotations()[0], a1r);
}

#[test]
fn annotations_upsert_preserves_position() {
    // The core routes both keyed collections through an order-preserving
    // upsert. Updating the first of three must not move it to the end.
    let st = Annotations::new().apply_all(&[
        An::Set(Annotation::new("a", "t", "r")),
        An::Set(Annotation::new("b", "t", "r")),
        An::Set(Annotation::new("c", "t", "r")),
    ]);

    let updated = st.apply(&An::Set(Annotation {
        resolved: true,
        ..Annotation::new("a", "t", "r")
    }));

    let ids: Vec<String> = updated.annotations().into_iter().map(|a| a.id).collect();
    assert_eq!(ids, vec!["a", "b", "c"], "upsert must preserve order");
    assert!(updated.get("a").unwrap().resolved);
}

#[test]
fn annotations_operations_on_unknown_ids_are_proven_no_ops() {
    // Dafny fixtures 212 (removed unknown), 214 (entrySet unknown), 219
    // (updated unknown). These are no-ops by proof, not errors.
    let st = Annotations::new().apply(&An::Set(a1()));

    assert_eq!(st.apply(&An::Removed("nope".into())), st);
    assert_eq!(
        st.apply(&An::EntrySet {
            annotation_id: "nope".into(),
            entry: Entry::new("c-2", "reply"),
        }),
        st
    );
    assert_eq!(
        st.apply(&An::Updated {
            annotation_id: "nope".into(),
            update: AnnotationUpdate {
                resolved: Some(true),
                ..Default::default()
            },
        }),
        st
    );
}

#[test]
fn annotations_entry_set_and_remove_round_trip() {
    // Dafny fixtures 213 and 215.
    let st = Annotations::new().apply(&An::Set(a1()));
    let e2 = Entry::new("c-2", "reply");

    let with_two = st.apply(&An::EntrySet {
        annotation_id: "t-1".into(),
        entry: e2.clone(),
    });
    assert_eq!(with_two.get("t-1").unwrap().entries, vec![Entry::new("c-1", "original"), e2]);

    let back = with_two.apply(&An::EntryRemoved {
        annotation_id: "t-1".into(),
        entry_id: "c-2".into(),
    });
    assert_eq!(back, st, "removing the added entry returns the original state");
}

#[test]
fn annotations_update_changes_only_named_fields() {
    // Dafny fixtures 216 (resolve, preserve rest) and 218 (reanchor, preserve
    // resolved + entries).
    let st = Annotations::new().apply(&An::Set(a1()));

    let resolved = st.apply(&An::Updated {
        annotation_id: "t-1".into(),
        update: AnnotationUpdate {
            resolved: Some(true),
            ..Default::default()
        },
    });
    let got = resolved.get("t-1").unwrap();
    assert!(got.resolved);
    assert_eq!(got.entries, a1().entries, "entries survive a resolve");
    assert_eq!(got.range, a1().range, "range survives a resolve");

    let new_range = Json::object([("start", Json::Bool(true))]);
    let reanchored = st.apply(&An::Updated {
        annotation_id: "t-1".into(),
        update: AnnotationUpdate {
            turn_id: Some("turn-2".into()),
            resource: Some("file:///src/b.ts".into()),
            range: Some(new_range.clone()),
            resolved: None,
        },
    });
    let got = reanchored.get("t-1").unwrap();
    assert_eq!(got.turn_id, "turn-2");
    assert_eq!(got.resource, "file:///src/b.ts");
    assert_eq!(got.range, Some(new_range));
    assert!(!got.resolved, "an unnamed field must not change");
    assert_eq!(got.entries, a1().entries);
}

// ================================================================ changeset

#[test]
fn changeset_status_change_clears_a_stale_error() {
    // Dafny fixture: apply1(ChangesetState("computing", [], None, Some(JStr("boom"))),
    //                       StatusChanged("ready", None))
    //                == ChangesetState("ready", [], None, None)
    let errored = Changeset::with_status("computing").apply(&Cs::StatusChanged {
        status: "computing".into(),
        error: Some(Json::string("boom")),
    });
    assert_eq!(errored.error(), Some(Json::string("boom")));

    let ready = errored.apply(&Cs::StatusChanged {
        status: "ready".into(),
        error: None,
    });
    assert_eq!(ready.status(), "ready");
    assert_eq!(ready.error(), None, "a clean status must not carry the old error");
}

#[test]
fn changeset_distinguishes_unreported_operations_from_an_empty_list() {
    // The core models `operations` as Option<seq<..>>; collapsing None to an
    // empty Vec in the Rust view would erase a distinction the reducer keeps.
    let fresh = Changeset::new();
    assert_eq!(fresh.operations(), None, "a fresh changeset has not reported operations");

    let empty = fresh.apply(&Cs::OperationsChanged(Some(vec![])));
    assert_eq!(
        empty.operations(),
        Some(vec![]),
        "reporting an empty list is not the same as reporting nothing"
    );
    assert_ne!(fresh.operations(), empty.operations());

    let cleared = empty.apply(&Cs::OperationsChanged(None));
    assert_eq!(cleared.operations(), None);
}

#[test]
fn changeset_files_are_keyed_and_order_preserving() {
    let cs = Changeset::new().apply_all(&[
        Cs::FileSet(File::new("f1", Json::string("edit-1"))),
        Cs::FileSet(File::new("f2", Json::string("edit-2"))),
        Cs::FileSet(File::new("f3", Json::string("edit-3"))),
    ]);
    assert_eq!(cs.files().len(), 3);

    let updated = cs.apply(&Cs::FileSet(File::new("f1", Json::string("edit-1b"))));
    let ids: Vec<String> = updated.files().into_iter().map(|f| f.id).collect();
    assert_eq!(ids, vec!["f1", "f2", "f3"], "upsert must preserve order");
    assert_eq!(updated.file("f1").unwrap().edit, Json::string("edit-1b"));

    let removed = updated.apply(&Cs::FileRemoved("f2".into()));
    let ids: Vec<String> = removed.files().into_iter().map(|f| f.id).collect();
    assert_eq!(ids, vec!["f1", "f3"]);
}

#[test]
fn changeset_batch_review_marks_only_named_files() {
    let cs = Changeset::new().apply_all(&[
        Cs::FileSet(File::new("f1", Json::Null)),
        Cs::FileSet(File::new("f2", Json::Null)),
        Cs::FileSet(File::new("f3", Json::Null)),
        Cs::FilesReviewedChanged {
            file_ids: vec!["f1".into(), "f3".into()],
            reviewed: true,
        },
    ]);

    assert_eq!(cs.file("f1").unwrap().reviewed, Some(true));
    assert_eq!(cs.file("f3").unwrap().reviewed, Some(true));
    assert_eq!(
        cs.file("f2").unwrap().reviewed,
        None,
        "an unnamed file keeps its unset review state"
    );
}

#[test]
fn changeset_operation_status_change_targets_one_operation() {
    let cs = Changeset::new()
        .apply(&Cs::OperationsChanged(Some(vec![
            Operation::new("op-1", "scan", "running"),
            Operation::new("op-2", "apply", "running"),
        ])))
        .apply(&Cs::OperationStatusChanged {
            operation_id: "op-1".into(),
            status: "failed".into(),
            error: Some(Json::string("nope")),
        });

    let ops = cs.operations().unwrap();
    assert_eq!(ops[0].status, "failed");
    assert_eq!(ops[0].error, Some(Json::string("nope")));
    assert_eq!(ops[1].status, "running", "the sibling operation is untouched");
    assert_eq!(ops[1].error, None);
}

#[test]
fn changeset_unknown_is_a_no_op() {
    let cs = Changeset::new().apply(&Cs::FileSet(File::new("f1", Json::Null)));
    assert_eq!(cs.apply(&Cs::Unknown(Json::string("who knows"))), cs);
}
