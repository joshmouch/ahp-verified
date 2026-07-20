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

# Module: ConfluxDecimalText

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Digit(value):
        return _dafny.CodePoint(chr((ord(_dafny.CodePoint('0'))) + (value)))

    @staticmethod
    def NatText(value):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (value) < (10):
                    return (_dafny.SeqWithoutIsStrInference([default__.Digit(value)])) + (d_0___accumulator_)
                elif True:
                    d_0___accumulator_ = (_dafny.SeqWithoutIsStrInference([default__.Digit(_dafny.euclidian_modulus(value, 10))])) + (d_0___accumulator_)
                    in0_ = _dafny.euclidian_division(value, 10)
                    value = in0_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def ParseDigits(text, position, value):
        while True:
            with _dafny.label():
                if (((position) == (len(text))) or (((text)[position]) < (_dafny.CodePoint('0')))) or (((text)[position]) > (_dafny.CodePoint('9'))):
                    return value
                elif True:
                    in0_ = text
                    in1_ = (position) + (1)
                    in2_ = ((value) * (10)) + ((ord((text)[position])) - (ord(_dafny.CodePoint('0'))))
                    text = in0_
                    position = in1_
                    value = in2_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def ParseNat(text):
        return default__.ParseDigits(text, 0, 0)

