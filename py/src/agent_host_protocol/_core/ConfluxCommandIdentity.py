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

# Module: ConfluxCommandIdentity

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def DecodeCanonicalCommand(encodeCommand, decodeCommand, bytes_):
        d_0_parsed_ = ConfluxCommandIdentityCapability.default__.DecodeCanonicalJson(bytes_)
        if (d_0_parsed_).is_None:
            return ConfluxContract.Option_None()
        elif True:
            d_1_decoded_ = decodeCommand((d_0_parsed_).value)
            if ((d_1_decoded_).is_None) or ((encodeCommand((d_1_decoded_).value)) != ((d_0_parsed_).value)):
                return ConfluxContract.Option_None()
            elif True:
                return ConfluxContract.Option_Some((d_1_decoded_).value)

    @staticmethod
    def CommandCodecFromJson(encodeCommand, decodeCommand):
        def lambda0_(d_0_encodeCommand_):
            def lambda1_(d_1_command_):
                return ConfluxCommandIdentityCapability.default__.CanonicalJsonBytes(d_0_encodeCommand_(d_1_command_))

            return lambda1_

        def lambda2_(d_2_encodeCommand_, d_3_decodeCommand_):
            def lambda3_(d_4_bytes_):
                return default__.DecodeCanonicalCommand(d_2_encodeCommand_, d_3_decodeCommand_, d_4_bytes_)

            return lambda3_

        def lambda4_(d_5_bytes_):
            return ConfluxCommandIdentityCapability.default__.Sha256Digest(d_5_bytes_)

        return ConfluxCommand.CommandIdentityCodec_CommandIdentityCodec(ConfluxCommand.default__.CanonicalCommandCodecVersion(), lambda0_(encodeCommand), lambda2_(encodeCommand, decodeCommand), lambda4_)

