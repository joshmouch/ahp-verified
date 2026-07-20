// Package ClientMain
// Dafny module ClientMain compiled into Go

package ClientMain

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-verified/go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-verified/go/Changeset"
	m_Chat "github.com/joshmouch/ahp-verified/go/Chat"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-verified/go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-verified/go/ConfluxOrderedLog"
	m_ConfluxPrune "github.com/joshmouch/ahp-verified/go/ConfluxPrune"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-verified/go/ConfluxSeqRoute"
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
	return "ClientMain.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _0_rootPass _dafny.Int
	_ = _0_rootPass
	var _1_rootTotal _dafny.Int
	_ = _1_rootTotal
	var _out0 _dafny.Int
	_ = _out0
	var _out1 _dafny.Int
	_ = _out1
	_out0, _out1 = m_AhpSkeleton.Companion_Default___.RunCorpus()
	_0_rootPass = _out0
	_1_rootTotal = _out1
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("ROOT CORPUS:          ").VerbatimString(false))
	_dafny.Print(_0_rootPass)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("/").VerbatimString(false))
	_dafny.Print(_1_rootTotal)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(" green against extracted code\n").VerbatimString(false))
	var _2_rwPass _dafny.Int
	_ = _2_rwPass
	var _3_rwTotal _dafny.Int
	_ = _3_rwTotal
	var _out2 _dafny.Int
	_ = _out2
	var _out3 _dafny.Int
	_ = _out3
	_out2, _out3 = m_ResourceWatch.Companion_Default___.RunCorpus()
	_2_rwPass = _out2
	_3_rwTotal = _out3
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("RESOURCEWATCH CORPUS: ").VerbatimString(false))
	_dafny.Print(_2_rwPass)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("/").VerbatimString(false))
	_dafny.Print(_3_rwTotal)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(" green against extracted code\n").VerbatimString(false))
	var _4_cvPass _dafny.Int
	_ = _4_cvPass
	var _5_cvTotal _dafny.Int
	_ = _5_cvTotal
	var _out4 _dafny.Int
	_ = _out4
	var _out5 _dafny.Int
	_ = _out5
	_out4, _out5 = m_Canvas.Companion_Default___.RunCorpus()
	_4_cvPass = _out4
	_5_cvTotal = _out5
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("CANVAS CORPUS:        ").VerbatimString(false))
	_dafny.Print(_4_cvPass)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("/").VerbatimString(false))
	_dafny.Print(_5_cvTotal)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(" green against extracted code\n").VerbatimString(false))
	var _6_csPass _dafny.Int
	_ = _6_csPass
	var _7_csTotal _dafny.Int
	_ = _7_csTotal
	var _out6 _dafny.Int
	_ = _out6
	var _out7 _dafny.Int
	_ = _out7
	_out6, _out7 = m_Changeset.Companion_Default___.RunCorpus()
	_6_csPass = _out6
	_7_csTotal = _out7
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("CHANGESET CORPUS:     ").VerbatimString(false))
	_dafny.Print(_6_csPass)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("/").VerbatimString(false))
	_dafny.Print(_7_csTotal)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(" green against extracted code\n").VerbatimString(false))
	var _8_anPass _dafny.Int
	_ = _8_anPass
	var _9_anTotal _dafny.Int
	_ = _9_anTotal
	var _out8 _dafny.Int
	_ = _out8
	var _out9 _dafny.Int
	_ = _out9
	_out8, _out9 = m_Annotations.Companion_Default___.RunCorpus()
	_8_anPass = _out8
	_9_anTotal = _out9
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("ANNOTATIONS CORPUS:   ").VerbatimString(false))
	_dafny.Print(_8_anPass)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("/").VerbatimString(false))
	_dafny.Print(_9_anTotal)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(" green against extracted code\n").VerbatimString(false))
	var _10_tPass _dafny.Int
	_ = _10_tPass
	var _11_tTotal _dafny.Int
	_ = _11_tTotal
	var _out10 _dafny.Int
	_ = _out10
	var _out11 _dafny.Int
	_ = _out11
	_out10, _out11 = m_Terminal.Companion_Default___.RunCorpus()
	_10_tPass = _out10
	_11_tTotal = _out11
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("TERMINAL CORPUS:      ").VerbatimString(false))
	_dafny.Print(_10_tPass)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("/").VerbatimString(false))
	_dafny.Print(_11_tTotal)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(" green against extracted code\n").VerbatimString(false))
	var _12_sPass _dafny.Int
	_ = _12_sPass
	var _13_sModeled _dafny.Int
	_ = _13_sModeled
	var _out12 _dafny.Int
	_ = _out12
	var _out13 _dafny.Int
	_ = _out13
	_out12, _out13 = m_Session.Companion_Default___.RunCorpus()
	_12_sPass = _out12
	_13_sModeled = _out13
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("SESSION CORPUS:       ").VerbatimString(false))
	_dafny.Print(_12_sPass)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("/").VerbatimString(false))
	_dafny.Print(_13_sModeled)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(" modeled green (of 61 total; all ~25 action TYPES now modeled)\n").VerbatimString(false))
	var _14_chPass _dafny.Int
	_ = _14_chPass
	var _15_chModeled _dafny.Int
	_ = _15_chModeled
	var _out14 _dafny.Int
	_ = _out14
	var _out15 _dafny.Int
	_ = _out15
	_out14, _out15 = m_Chat.Companion_Default___.RunCorpus()
	_14_chPass = _out14
	_15_chModeled = _out15
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("CHAT CORPUS:          ").VerbatimString(false))
	_dafny.Print(_14_chPass)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("/").VerbatimString(false))
	_dafny.Print(_15_chModeled)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(" modeled green (of 97 total; full tool-call state machine + turn lifecycle modeled)\n").VerbatimString(false))
	var _16_pass _dafny.Int
	_ = _16_pass
	_16_pass = (((((((_0_rootPass).Plus(_2_rwPass)).Plus(_4_cvPass)).Plus(_6_csPass)).Plus(_8_anPass)).Plus(_10_tPass)).Plus(_12_sPass)).Plus(_14_chPass)
	var _17_total _dafny.Int
	_ = _17_total
	_17_total = (((((((_1_rootTotal).Plus(_3_rwTotal)).Plus(_5_cvTotal)).Plus(_7_csTotal)).Plus(_9_anTotal)).Plus(_11_tTotal)).Plus(_13_sModeled)).Plus(_15_chModeled)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("TOTAL: ").VerbatimString(false))
	_dafny.Print(_16_pass)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("/").VerbatimString(false))
	_dafny.Print(_17_total)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(" corpus fixtures green (5 full AHP channels + session/chat partial)\n").VerbatimString(false))
	if !((_16_pass).Cmp(_17_total) == 0) {
		panic("spec/client_main.dfy(41,4): " + (_dafny.UnicodeSeqOfUtf8Bytes("expectation violation")).VerbatimString(false))
	}
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("MULTI-CHANNEL CLIENT OK — extract(cs,js) + corpus all green\n").VerbatimString(false))
}

// End of class Default__
