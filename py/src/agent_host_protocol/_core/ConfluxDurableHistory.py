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

# Module: ConfluxDurableHistory

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Append(history, entry):
        return ((history)) + (_dafny.SeqWithoutIsStrInference([entry]))

    @staticmethod
    def Replay(reducer, initial, history):
        return ConfluxContract.default__.Fold(reducer, initial, (history))

    @staticmethod
    def MakeCheckpoint(historyId, reducerSchemaVersion, coordinateOf, hashBytes, codec, state):
        d_0_bytes_ = (codec).encode(state)
        return ReplayCheckpoint_ReplayCheckpoint(1, historyId, coordinateOf(state), reducerSchemaVersion, hashBytes(d_0_bytes_), d_0_bytes_)

    @staticmethod
    def RestoreCheckpoint(expectedHistoryId, expectedReducerSchemaVersion, coordinateOf, hashBytes, codec, checkpoint):
        if (((((checkpoint).schemaVersion) != (1)) or (((checkpoint).historyId) != (expectedHistoryId))) or (((checkpoint).reducerSchemaVersion) != (expectedReducerSchemaVersion))) or (((checkpoint).stateHash) != (hashBytes((checkpoint).stateBytes))):
            return ConfluxContract.Option_None()
        elif True:
            source0_ = (codec).decode((checkpoint).stateBytes)
            if True:
                if source0_.is_None:
                    return ConfluxContract.Option_None()
            if True:
                d_0_state_ = source0_.value
                if (coordinateOf(d_0_state_)) == ((checkpoint).nextCoordinate):
                    return ConfluxContract.Option_Some(d_0_state_)
                elif True:
                    return ConfluxContract.Option_None()


class History:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_History(self) -> bool:
        return isinstance(self, History_History)

class History_History(History, NamedTuple('History', [('entries', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxDurableHistory.History.History({_dafny.string_of(self.entries)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, History_History) and self.entries == __o.entries
    def __hash__(self) -> int:
        return super().__hash__()


class StateCodec:
    @classmethod
    def default(cls, ):
        return lambda: StateCodec_StateCodec((lambda x10: _dafny.Seq({})), (lambda x11: ConfluxContract.Option.default()()))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_StateCodec(self) -> bool:
        return isinstance(self, StateCodec_StateCodec)

class StateCodec_StateCodec(StateCodec, NamedTuple('StateCodec', [('encode', Any), ('decode', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxDurableHistory.StateCodec.StateCodec({_dafny.string_of(self.encode)}, {_dafny.string_of(self.decode)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, StateCodec_StateCodec) and self.encode == __o.encode and self.decode == __o.decode
    def __hash__(self) -> int:
        return super().__hash__()


class ReplayCheckpoint:
    @classmethod
    def default(cls, ):
        return lambda: ReplayCheckpoint_ReplayCheckpoint(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ReplayCheckpoint(self) -> bool:
        return isinstance(self, ReplayCheckpoint_ReplayCheckpoint)

class ReplayCheckpoint_ReplayCheckpoint(ReplayCheckpoint, NamedTuple('ReplayCheckpoint', [('schemaVersion', Any), ('historyId', Any), ('nextCoordinate', Any), ('reducerSchemaVersion', Any), ('stateHash', Any), ('stateBytes', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxDurableHistory.ReplayCheckpoint.ReplayCheckpoint({_dafny.string_of(self.schemaVersion)}, {self.historyId.VerbatimString(True)}, {_dafny.string_of(self.nextCoordinate)}, {self.reducerSchemaVersion.VerbatimString(True)}, {self.stateHash.VerbatimString(True)}, {_dafny.string_of(self.stateBytes)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ReplayCheckpoint_ReplayCheckpoint) and self.schemaVersion == __o.schemaVersion and self.historyId == __o.historyId and self.nextCoordinate == __o.nextCoordinate and self.reducerSchemaVersion == __o.reducerSchemaVersion and self.stateHash == __o.stateHash and self.stateBytes == __o.stateBytes
    def __hash__(self) -> int:
        return super().__hash__()

