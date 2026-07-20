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
import ConfluxBudgetConvergence as ConfluxBudgetConvergence
import ConfluxExtern as ConfluxExtern
import ConfluxCommandIdentityCapability as ConfluxCommandIdentityCapability
import ConfluxDecimalText as ConfluxDecimalText
import ConfluxJsonText as ConfluxJsonText
import ConfluxCommandIdentity as ConfluxCommandIdentity
import ConfluxIoCapability as ConfluxIoCapability
import ConfluxResourceSupervisor as ConfluxResourceSupervisor
import ConfluxSocketCapability as ConfluxSocketCapability
import ConfluxWebSocketCapability as ConfluxWebSocketCapability
import ConfluxHttpCapability as ConfluxHttpCapability
import ConfluxServiceLifecycle as ConfluxServiceLifecycle

# Module: ConfluxSequence

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def StartsAt(haystack, needle, at):
        return (((at) + (len(needle))) <= (len(haystack))) and ((_dafny.SeqWithoutIsStrInference((haystack)[at:(at) + (len(needle)):])) == (needle))

    @staticmethod
    def FindSubsequenceFrom(haystack, needle, at):
        while True:
            with _dafny.label():
                if default__.StartsAt(haystack, needle, at):
                    return at
                elif (at) == (len(haystack)):
                    return -1
                elif True:
                    in0_ = haystack
                    in1_ = needle
                    in2_ = (at) + (1)
                    haystack = in0_
                    needle = in1_
                    at = in2_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def Contains(haystack, needle):
        return (default__.FindSubsequenceFrom(haystack, needle, 0)) >= (0)

    @staticmethod
    def EndsWith(haystack, suffix):
        return ((len(suffix)) <= (len(haystack))) and ((_dafny.SeqWithoutIsStrInference((haystack)[(len(haystack)) - (len(suffix))::])) == (suffix))

    @staticmethod
    def FindIndexWhere(p, items):
        if (len(items)) == (0):
            return -1
        elif p((items)[0]):
            return 0
        elif True:
            d_0_tail_ = default__.FindIndexWhere(p, _dafny.SeqWithoutIsStrInference((items)[1::]))
            if (d_0_tail_) < (0):
                return -1
            elif True:
                return (d_0_tail_) + (1)

