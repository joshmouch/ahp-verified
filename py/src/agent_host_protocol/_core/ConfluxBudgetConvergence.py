import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_ as module_
import _dafny as _dafny
import System_ as System_
import ConfluxCodec as ConfluxCodec
import ConfluxContract as ConfluxContract
import AhpSkeleton as AhpSkeleton
import ResourceWatch as ResourceWatch
import Canvas as Canvas
import ConfluxOperators as ConfluxOperators
import ConfluxPrune as ConfluxPrune
import ConfluxSeqRoute as ConfluxSeqRoute
import Changeset as Changeset
import Annotations as Annotations
import Terminal as Terminal
import Session as Session
import ConfluxOrderedLog as ConfluxOrderedLog
import Chat as Chat
import ClientMain as ClientMain
import ConfluxSemanticGraphIdentity as ConfluxSemanticGraphIdentity
import ConfluxContentAddress as ConfluxContentAddress
import ConfluxSemanticGraphContract as ConfluxSemanticGraphContract
import ConfluxResourceCeilings as ConfluxResourceCeilings
import ConfluxSupervisedOperationResult as ConfluxSupervisedOperationResult
import ConfluxTree as ConfluxTree
import ConfluxMerge as ConfluxMerge
import ConfluxRoute as ConfluxRoute
import ConfluxMapValues as ConfluxMapValues
import ConfluxDelimited as ConfluxDelimited
import ConfluxFixpoint as ConfluxFixpoint
import ConfluxGenericCodec as ConfluxGenericCodec
import ConfluxResolve as ConfluxResolve
import ConfluxWatermark as ConfluxWatermark
import ConfluxStateMachine as ConfluxStateMachine
import ConfluxProduct as ConfluxProduct
import ConfluxJoin as ConfluxJoin
import ConfluxDedupe as ConfluxDedupe
import ConfluxBatchRoute as ConfluxBatchRoute
import ConfluxSeqRouteMerge as ConfluxSeqRouteMerge
import ConfluxKeyedOps as ConfluxKeyedOps
import ConfluxFilterKeys as ConfluxFilterKeys
import ConfluxChannel as ConfluxChannel
import ConfluxDurableHistory as ConfluxDurableHistory
import ConfluxCommand as ConfluxCommand
import ConfluxChannelManifest as ConfluxChannelManifest
import ConfluxConvergence as ConfluxConvergence

# Module: ConfluxBudgetConvergence

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def ApplyObservation(budget, phase, state, observation):
        source0_ = state
        if True:
            if source0_.is_SupervisionTripped:
                d_0_record_ = source0_.record
                return state
        if True:
            d_1_samplesSeen_ = source0_.samplesSeen
            d_2_verdict_ = ConfluxResourceCeilings.default__.DecideInFlight(budget, phase, observation)
            if ((d_2_verdict_).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementKill()):
                return ConfluxSupervisedOperationResult.SupervisionState_SupervisionTripped(ConfluxSupervisedOperationResult.FirstBreachV1_FirstBreachV1(d_1_samplesSeen_, d_2_verdict_, ((d_2_verdict_).observed).is_None))
            elif True:
                return ConfluxSupervisedOperationResult.SupervisionState_SupervisionRunning((d_1_samplesSeen_) + (1))

    @staticmethod
    def Drive(budget, phase, observations):
        return ConfluxContract.default__.Fold(default__.SupervisionReducer(budget, phase), ConfluxSupervisedOperationResult.SupervisionState_SupervisionRunning(0), observations)

    @staticmethod
    def SupervisionReducer(budget, phase):
        def lambda0_(d_0_budget_, d_1_phase_):
            def lambda1_(d_2_state_, d_3_observation_):
                return default__.ApplyObservation(d_0_budget_, d_1_phase_, d_2_state_, d_3_observation_)

            return lambda1_

        return lambda0_(budget, phase)

