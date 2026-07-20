// Package ConfluxConvergence
// Dafny module ConfluxConvergence compiled into Go

package ConfluxConvergence

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-verified/go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-verified/go/Changeset"
	m_Chat "github.com/joshmouch/ahp-verified/go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-verified/go/ClientMain"
	m_ConfluxBatchRoute "github.com/joshmouch/ahp-verified/go/ConfluxBatchRoute"
	m_ConfluxChannel "github.com/joshmouch/ahp-verified/go/ConfluxChannel"
	m_ConfluxChannelManifest "github.com/joshmouch/ahp-verified/go/ConfluxChannelManifest"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxCommand "github.com/joshmouch/ahp-verified/go/ConfluxCommand"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-verified/go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ConfluxDedupe "github.com/joshmouch/ahp-verified/go/ConfluxDedupe"
	m_ConfluxDelimited "github.com/joshmouch/ahp-verified/go/ConfluxDelimited"
	m_ConfluxDurableHistory "github.com/joshmouch/ahp-verified/go/ConfluxDurableHistory"
	m_ConfluxFilterKeys "github.com/joshmouch/ahp-verified/go/ConfluxFilterKeys"
	m_ConfluxFixpoint "github.com/joshmouch/ahp-verified/go/ConfluxFixpoint"
	m_ConfluxGenericCodec "github.com/joshmouch/ahp-verified/go/ConfluxGenericCodec"
	m_ConfluxJoin "github.com/joshmouch/ahp-verified/go/ConfluxJoin"
	m_ConfluxKeyedOps "github.com/joshmouch/ahp-verified/go/ConfluxKeyedOps"
	m_ConfluxMapValues "github.com/joshmouch/ahp-verified/go/ConfluxMapValues"
	m_ConfluxMerge "github.com/joshmouch/ahp-verified/go/ConfluxMerge"
	m_ConfluxOperators "github.com/joshmouch/ahp-verified/go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-verified/go/ConfluxOrderedLog"
	m_ConfluxProduct "github.com/joshmouch/ahp-verified/go/ConfluxProduct"
	m_ConfluxPrune "github.com/joshmouch/ahp-verified/go/ConfluxPrune"
	m_ConfluxResolve "github.com/joshmouch/ahp-verified/go/ConfluxResolve"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-verified/go/ConfluxResourceCeilings"
	m_ConfluxRoute "github.com/joshmouch/ahp-verified/go/ConfluxRoute"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-verified/go/ConfluxSeqRoute"
	m_ConfluxSeqRouteMerge "github.com/joshmouch/ahp-verified/go/ConfluxSeqRouteMerge"
	m_ConfluxStateMachine "github.com/joshmouch/ahp-verified/go/ConfluxStateMachine"
	m_ConfluxSupervisedOperationResult "github.com/joshmouch/ahp-verified/go/ConfluxSupervisedOperationResult"
	m_ConfluxTree "github.com/joshmouch/ahp-verified/go/ConfluxTree"
	m_ConfluxWatermark "github.com/joshmouch/ahp-verified/go/ConfluxWatermark"
	m_ResourceWatch "github.com/joshmouch/ahp-verified/go/ResourceWatch"
	m_Session "github.com/joshmouch/ahp-verified/go/Session"
	m__System "github.com/joshmouch/ahp-verified/go/System_"
	m_Terminal "github.com/joshmouch/ahp-verified/go/Terminal"
	_dafny "github.com/joshmouch/ahp-verified/go/dafny"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__
var _ m_ConfluxCodec.Dummy__
var _ m_ConfluxContract.Dummy__
var _ m_AhpSkeleton.Dummy__
var _ m_ResourceWatch.Dummy__
var _ m_Canvas.Dummy__
var _ m_ConfluxOperators.Dummy__
var _ m_ConfluxPrune.Dummy__
var _ m_ConfluxSeqRoute.Dummy__
var _ m_Changeset.Dummy__
var _ m_Annotations.Dummy__
var _ m_Terminal.Dummy__
var _ m_Session.Dummy__
var _ m_ConfluxOrderedLog.Dummy__
var _ m_Chat.Dummy__
var _ m_ClientMain.Dummy__
var _ m_ConfluxSemanticGraphIdentity.Dummy__
var _ m_ConfluxContentAddress.Dummy__
var _ m_ConfluxSemanticGraphContract.Dummy__
var _ m_ConfluxResourceCeilings.Dummy__
var _ m_ConfluxSupervisedOperationResult.Dummy__
var _ m_ConfluxTree.Dummy__
var _ m_ConfluxMerge.Dummy__
var _ m_ConfluxRoute.Dummy__
var _ m_ConfluxMapValues.Dummy__
var _ m_ConfluxDelimited.Dummy__
var _ m_ConfluxFixpoint.Dummy__
var _ m_ConfluxGenericCodec.Dummy__
var _ m_ConfluxResolve.Dummy__
var _ m_ConfluxWatermark.Dummy__
var _ m_ConfluxStateMachine.Dummy__
var _ m_ConfluxProduct.Dummy__
var _ m_ConfluxJoin.Dummy__
var _ m_ConfluxDedupe.Dummy__
var _ m_ConfluxBatchRoute.Dummy__
var _ m_ConfluxSeqRouteMerge.Dummy__
var _ m_ConfluxKeyedOps.Dummy__
var _ m_ConfluxFilterKeys.Dummy__
var _ m_ConfluxChannel.Dummy__
var _ m_ConfluxDurableHistory.Dummy__
var _ m_ConfluxCommand.Dummy__
var _ m_ConfluxChannelManifest.Dummy__

type Dummy__ struct{}

// Definition of class Default__
type Default__ struct {
	dummy byte
}

func New_Default___() *Default__ {
	_this := Default__{}

	return &_this
}

type CompanionStruct_Default___ struct {
}

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "ConfluxConvergence.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) FoldHost(r func(interface{}, interface{}) interface{}, h m_ConfluxChannel.HostState, xs _dafny.Sequence) m_ConfluxChannel.HostState {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer109 func(m_ConfluxChannel.HostState, interface{}) m_ConfluxChannel.HostState) func(interface{}, interface{}) interface{} {
		return func(arg128 interface{}, arg129 interface{}) interface{} {
			return coer109(arg128.(m_ConfluxChannel.HostState), arg129)
		}
	}((func(_0_r func(interface{}, interface{}) interface{}) func(m_ConfluxChannel.HostState, interface{}) m_ConfluxChannel.HostState {
		return func(_1_hs m_ConfluxChannel.HostState, _2_a interface{}) m_ConfluxChannel.HostState {
			return m_ConfluxChannel.Companion_Default___.Accept(func(coer110 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
				return func(arg130 interface{}, arg131 interface{}) interface{} {
					return coer110(arg130, arg131)
				}
			}(_0_r), _1_hs, _2_a)
		}
	})(r)), h, xs).(m_ConfluxChannel.HostState)
}
func (_static *CompanionStruct_Default___) FoldClient(r func(interface{}, interface{}) interface{}, c m_ConfluxChannel.ClientState, xs _dafny.Sequence) m_ConfluxChannel.ClientState {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer111 func(m_ConfluxChannel.ClientState, interface{}) m_ConfluxChannel.ClientState) func(interface{}, interface{}) interface{} {
		return func(arg132 interface{}, arg133 interface{}) interface{} {
			return coer111(arg132.(m_ConfluxChannel.ClientState), arg133)
		}
	}((func(_0_r func(interface{}, interface{}) interface{}) func(m_ConfluxChannel.ClientState, interface{}) m_ConfluxChannel.ClientState {
		return func(_1_cs m_ConfluxChannel.ClientState, _2_a interface{}) m_ConfluxChannel.ClientState {
			return m_ConfluxChannel.Companion_Default___.Receive(func(coer112 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
				return func(arg134 interface{}, arg135 interface{}) interface{} {
					return coer112(arg134, arg135)
				}
			}(_0_r), _1_cs, _2_a)
		}
	})(r)), c, xs).(m_ConfluxChannel.ClientState)
}
func (_static *CompanionStruct_Default___) InitClients(s0 interface{}, n _dafny.Int) _dafny.Sequence {
	return _dafny.SeqCreate((n).Uint32(), func(coer113 func(_dafny.Int) m_ConfluxChannel.ClientState) func(_dafny.Int) interface{} {
		return func(arg136 _dafny.Int) interface{} {
			return coer113(arg136)
		}
	}((func(_0_s0 interface{}) func(_dafny.Int) m_ConfluxChannel.ClientState {
		return func(_1___v4 _dafny.Int) m_ConfluxChannel.ClientState {
			return m_ConfluxChannel.Companion_ClientState_.Create_ClientState_(_0_s0, _dafny.Zero)
		}
	})(s0)))
}
func (_static *CompanionStruct_Default___) FanoutFold(r func(interface{}, interface{}) interface{}, clients _dafny.Sequence, xs _dafny.Sequence) _dafny.Sequence {
	return _dafny.SeqCreate((_dafny.IntOfUint32((clients).Cardinality())).Uint32(), func(coer114 func(_dafny.Int) m_ConfluxChannel.ClientState) func(_dafny.Int) interface{} {
		return func(arg137 _dafny.Int) interface{} {
			return coer114(arg137)
		}
	}((func(_0_r func(interface{}, interface{}) interface{}, _1_clients _dafny.Sequence, _2_xs _dafny.Sequence) func(_dafny.Int) m_ConfluxChannel.ClientState {
		return func(_3_i _dafny.Int) m_ConfluxChannel.ClientState {
			return Companion_Default___.FoldClient(func(coer115 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
				return func(arg138 interface{}, arg139 interface{}) interface{} {
					return coer115(arg138, arg139)
				}
			}(_0_r), (_1_clients).Select((_3_i).Uint32()).(m_ConfluxChannel.ClientState), _2_xs)
		}
	})(r, clients, xs)))
}
func (_static *CompanionStruct_Default___) IntAdd() func(_dafny.Int, _dafny.Int) _dafny.Int {
	return func(_0_s _dafny.Int, _1_a _dafny.Int) _dafny.Int {
		return (_0_s).Plus(_1_a)
	}
}

// End of class Default__
