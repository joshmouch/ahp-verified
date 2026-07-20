// Package ConfluxCommandIdentity
// Dafny module ConfluxCommandIdentity compiled into Go

package ConfluxCommandIdentity

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-go/Changeset"
	m_Chat "github.com/joshmouch/ahp-go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-go/ClientMain"
	m_ConfluxBatchRoute "github.com/joshmouch/ahp-go/ConfluxBatchRoute"
	m_ConfluxBudgetConvergence "github.com/joshmouch/ahp-go/ConfluxBudgetConvergence"
	m_ConfluxChannel "github.com/joshmouch/ahp-go/ConfluxChannel"
	m_ConfluxChannelManifest "github.com/joshmouch/ahp-go/ConfluxChannelManifest"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxCommand "github.com/joshmouch/ahp-go/ConfluxCommand"
	m_ConfluxCommandIdentityCapability "github.com/joshmouch/ahp-go/ConfluxCommandIdentityCapability"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m_ConfluxConvergence "github.com/joshmouch/ahp-go/ConfluxConvergence"
	m_ConfluxDecimalText "github.com/joshmouch/ahp-go/ConfluxDecimalText"
	m_ConfluxDedupe "github.com/joshmouch/ahp-go/ConfluxDedupe"
	m_ConfluxDelimited "github.com/joshmouch/ahp-go/ConfluxDelimited"
	m_ConfluxDurableHistory "github.com/joshmouch/ahp-go/ConfluxDurableHistory"
	m_ConfluxExtern "github.com/joshmouch/ahp-go/ConfluxExtern"
	m_ConfluxFilterKeys "github.com/joshmouch/ahp-go/ConfluxFilterKeys"
	m_ConfluxFixpoint "github.com/joshmouch/ahp-go/ConfluxFixpoint"
	m_ConfluxGenericCodec "github.com/joshmouch/ahp-go/ConfluxGenericCodec"
	m_ConfluxJoin "github.com/joshmouch/ahp-go/ConfluxJoin"
	m_ConfluxJsonText "github.com/joshmouch/ahp-go/ConfluxJsonText"
	m_ConfluxKeyedOps "github.com/joshmouch/ahp-go/ConfluxKeyedOps"
	m_ConfluxMapValues "github.com/joshmouch/ahp-go/ConfluxMapValues"
	m_ConfluxMerge "github.com/joshmouch/ahp-go/ConfluxMerge"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-go/ConfluxOrderedLog"
	m_ConfluxProduct "github.com/joshmouch/ahp-go/ConfluxProduct"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ConfluxResolve "github.com/joshmouch/ahp-go/ConfluxResolve"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-go/ConfluxResourceCeilings"
	m_ConfluxRoute "github.com/joshmouch/ahp-go/ConfluxRoute"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-go/ConfluxSeqRoute"
	m_ConfluxSeqRouteMerge "github.com/joshmouch/ahp-go/ConfluxSeqRouteMerge"
	m_ConfluxStateMachine "github.com/joshmouch/ahp-go/ConfluxStateMachine"
	m_ConfluxSupervisedOperationResult "github.com/joshmouch/ahp-go/ConfluxSupervisedOperationResult"
	m_ConfluxTree "github.com/joshmouch/ahp-go/ConfluxTree"
	m_ConfluxWatermark "github.com/joshmouch/ahp-go/ConfluxWatermark"
	m_ResourceWatch "github.com/joshmouch/ahp-go/ResourceWatch"
	m_Session "github.com/joshmouch/ahp-go/Session"
	m__System "github.com/joshmouch/ahp-go/System_"
	m_Terminal "github.com/joshmouch/ahp-go/Terminal"
	_dafny "github.com/joshmouch/ahp-go/dafny"
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
var _ m_ConfluxConvergence.Dummy__
var _ m_ConfluxBudgetConvergence.Dummy__
var _ m_ConfluxExtern.Dummy__
var _ m_ConfluxDecimalText.Dummy__
var _ m_ConfluxJsonText.Dummy__

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
	return "ConfluxCommandIdentity.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) DecodeCanonicalCommand(encodeCommand func(interface{}) m_ConfluxCodec.Json, decodeCommand func(m_ConfluxCodec.Json) m_ConfluxCodec.Option, bytes _dafny.Sequence) m_ConfluxContract.Option {
	var _0_parsed m_ConfluxCodec.Option = m_ConfluxCommandIdentityCapability.Companion_Default___.DecodeCanonicalJson(bytes)
	_ = _0_parsed
	if (_0_parsed).Is_None() {
		return m_ConfluxContract.Companion_Option_.Create_None_()
	} else {
		var _1_decoded m_ConfluxCodec.Option = (decodeCommand)((_0_parsed).Dtor_value().(m_ConfluxCodec.Json))
		_ = _1_decoded
		if ((_1_decoded).Is_None()) || (!((encodeCommand)((_1_decoded).Dtor_value())).Equals((_0_parsed).Dtor_value().(m_ConfluxCodec.Json))) {
			return m_ConfluxContract.Companion_Option_.Create_None_()
		} else {
			return m_ConfluxContract.Companion_Option_.Create_Some_((_1_decoded).Dtor_value())
		}
	}
}
func (_static *CompanionStruct_Default___) CommandCodecFromJson(encodeCommand func(interface{}) m_ConfluxCodec.Json, decodeCommand func(m_ConfluxCodec.Json) m_ConfluxCodec.Option) m_ConfluxCommand.CommandIdentityCodec {
	return m_ConfluxCommand.Companion_CommandIdentityCodec_.Create_CommandIdentityCodec_(m_ConfluxCommand.Companion_Default___.CanonicalCommandCodecVersion(), func(coer118 func(interface{}) _dafny.Sequence) func(interface{}) _dafny.Sequence {
		return func(arg143 interface{}) _dafny.Sequence {
			return coer118(arg143)
		}
	}((func(_0_encodeCommand func(interface{}) m_ConfluxCodec.Json) func(interface{}) _dafny.Sequence {
		return func(_1_command interface{}) _dafny.Sequence {
			return m_ConfluxCommandIdentityCapability.Companion_Default___.CanonicalJsonBytes((_0_encodeCommand)(_1_command))
		}
	})(encodeCommand)), func(coer119 func(_dafny.Sequence) m_ConfluxContract.Option) func(_dafny.Sequence) m_ConfluxContract.Option {
		return func(arg144 _dafny.Sequence) m_ConfluxContract.Option {
			return coer119(arg144)
		}
	}((func(_2_encodeCommand func(interface{}) m_ConfluxCodec.Json, _3_decodeCommand func(m_ConfluxCodec.Json) m_ConfluxCodec.Option) func(_dafny.Sequence) m_ConfluxContract.Option {
		return func(_4_bytes _dafny.Sequence) m_ConfluxContract.Option {
			return Companion_Default___.DecodeCanonicalCommand(func(coer120 func(interface{}) m_ConfluxCodec.Json) func(interface{}) m_ConfluxCodec.Json {
				return func(arg145 interface{}) m_ConfluxCodec.Json {
					return coer120(arg145)
				}
			}(_2_encodeCommand), func(coer121 func(m_ConfluxCodec.Json) m_ConfluxCodec.Option) func(m_ConfluxCodec.Json) m_ConfluxCodec.Option {
				return func(arg146 m_ConfluxCodec.Json) m_ConfluxCodec.Option {
					return coer121(arg146)
				}
			}(_3_decodeCommand), _4_bytes)
		}
	})(encodeCommand, decodeCommand)), func(_5_bytes _dafny.Sequence) _dafny.Sequence {
		return m_ConfluxCommandIdentityCapability.Companion_Default___.Sha256Digest(_5_bytes)
	})
}

// End of class Default__
