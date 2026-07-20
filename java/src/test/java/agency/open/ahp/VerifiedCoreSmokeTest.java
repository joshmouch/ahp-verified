// Copyright (c) Microsoft Corporation.
// Copyright (c) 2026 Josh Mouch.
// Licensed under the MIT License.

package agency.open.ahp;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import agency.open.ahp.AhpSkeleton.ReduceOutcome;
import agency.open.ahp.ConfluxCodec.Json;
import agency.open.ahp.ResourceWatch.ResourceWatchAction;
import agency.open.ahp.ResourceWatch.ResourceWatchAction_RWChanged;
import agency.open.ahp.ResourceWatch.ResourceWatchAction_RWUnknown;
import agency.open.ahp.ResourceWatch.ResourceWatchState;
import dafny.DafnySequence;
import dafny.Tuple2;
import java.math.BigInteger;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

/**
 * Proves the extracted core is alive on the JVM: real datatypes are constructed and real
 * reducers are invoked, rather than merely asserting that the artifact loads.
 */
class VerifiedCoreSmokeTest {

    private static DafnySequence<? extends dafny.CodePoint> str(String s) {
        return DafnySequence.asUnicodeString(s);
    }

    @Test
    @DisplayName("JSON value model round-trips through the verified codec datatype")
    void jsonValueModelRoundTrips() {
        Json str = Json.create_JStr(str("workspace"));
        Json num = Json.create_JNum(BigInteger.valueOf(42));

        assertTrue(str.is_JStr(), "expected a JStr");
        assertTrue(num.is_JNum(), "expected a JNum");
        assertEquals(BigInteger.valueOf(42), num.dtor_n());
        assertEquals(str("workspace"), str.dtor_s());

        // Structural equality is what the proofs rely on.
        assertEquals(Json.create_JStr(str("workspace")), str);
        assertNotEquals(Json.create_JStr(str("other")), str);
    }

    @Test
    @DisplayName("ResourceWatch reducer applies a change action")
    void resourceWatchReducerApplies() {
        ResourceWatchState state = ResourceWatchState.create(str("/repo"), true);
        ResourceWatchAction action =
                new ResourceWatchAction_RWChanged(Json.create_JArr(DafnySequence.empty(Json._typeDescriptor())));

        Tuple2<ResourceWatchState, ReduceOutcome> result =
                agency.open.ahp.ResourceWatch.__default.ApplyToResourceWatch(state, action, BigInteger.ZERO);

        ReduceOutcome outcome = result.dtor__1();
        assertTrue(outcome.is_Applied() || outcome.is_NoOp(), "a known action must not be OutOfScope");

        // The watch root is invariant under a change action.
        assertEquals(str("/repo"), result.dtor__0().dtor_root());
    }

    @Test
    @DisplayName("Unknown actions are classified as unknown and leave state untouched")
    void unknownActionIsInert() {
        ResourceWatchState state = ResourceWatchState.create(str("/repo"), false);
        ResourceWatchAction unknown =
                new ResourceWatchAction_RWUnknown(Json.create_JStr(str("resourceWatch/notARealAction")));

        assertTrue(agency.open.ahp.ResourceWatch.__default.IsUnknownRW(unknown), "should be classified unknown");

        ResourceWatchState after = agency.open.ahp.ResourceWatch.__default.apply1(state, unknown);
        assertEquals(state, after, "an unknown action must be a no-op on state");
    }

    @Test
    @DisplayName("fold over an empty action sequence is the identity")
    void foldOfEmptySequenceIsIdentity() {
        ResourceWatchState state = ResourceWatchState.create(str("/repo"), true);

        ResourceWatchState folded =
                agency.open.ahp.ResourceWatch.__default.fold(
                        state, DafnySequence.empty(ResourceWatchAction._typeDescriptor()));

        assertEquals(state, folded, "folding no actions must not change state");
    }
}
