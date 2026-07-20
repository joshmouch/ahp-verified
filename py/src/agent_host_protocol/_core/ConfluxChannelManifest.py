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

# Module: ConfluxChannelManifest

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Canonical(m, actions):
        return ConfluxContract.default__.Fold((m).apply, (m).initial, actions)

    @staticmethod
    def View(m, actions):
        return (m).readModel(default__.Canonical(m, actions))


class ChannelManifest:
    @classmethod
    def default(cls, default_S, default_W):
        return lambda: ChannelManifest_ChannelManifest(default_S(), (lambda x19, x20: default_S()), (lambda x21: ConfluxContract.Option.default()()), (lambda x22: default_W()))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ChannelManifest(self) -> bool:
        return isinstance(self, ChannelManifest_ChannelManifest)

class ChannelManifest_ChannelManifest(ChannelManifest, NamedTuple('ChannelManifest', [('initial', Any), ('apply', Any), ('decode', Any), ('readModel', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxChannelManifest.ChannelManifest.ChannelManifest({_dafny.string_of(self.initial)}, {_dafny.string_of(self.apply)}, {_dafny.string_of(self.decode)}, {_dafny.string_of(self.readModel)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChannelManifest_ChannelManifest) and self.initial == __o.initial and self.apply == __o.apply and self.decode == __o.decode and self.readModel == __o.readModel
    def __hash__(self) -> int:
        return super().__hash__()

