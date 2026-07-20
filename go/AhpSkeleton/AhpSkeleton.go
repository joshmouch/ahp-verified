// Package AhpSkeleton
// Dafny module AhpSkeleton compiled into Go

package AhpSkeleton

import (
	os "os"

	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m__System "github.com/joshmouch/ahp-verified/go/System_"
	_dafny "github.com/joshmouch/ahp-verified/go/dafny"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__
var _ m_ConfluxCodec.Dummy__
var _ m_ConfluxContract.Dummy__

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
	return "AhpSkeleton.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Pow10(k _dafny.Int) _dafny.Int {
	var _0___accumulator _dafny.Int = _dafny.One
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (k).Sign() == 0 {
		return (_dafny.One).Times(_0___accumulator)
	} else {
		_0___accumulator = (_0___accumulator).Times(_dafny.IntOfInt64(10))
		var _in0 _dafny.Int = (k).Minus(_dafny.One)
		_ = _in0
		k = _in0
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) DecInt(m _dafny.Int, e _dafny.Int) Option {
	if (e).Sign() != -1 {
		return Companion_Option_.Create_Some_((m).Times(Companion_Default___.Pow10(e)))
	} else if ((m).Modulo(Companion_Default___.Pow10((_dafny.Zero).Minus(e)))).Sign() == 0 {
		return Companion_Option_.Create_Some_((m).DivBy(Companion_Default___.Pow10((_dafny.Zero).Minus(e))))
	} else {
		return Companion_Option_.Create_None_()
	}
}
func (_static *CompanionStruct_Default___) JsonInt(j m_ConfluxCodec.Json) Option {
	if (j).Is_JNum() {
		return Companion_Option_.Create_Some_((j).Dtor_n())
	} else if (j).Is_JDec() {
		return Companion_Default___.DecInt((j).Dtor_mantissa(), (j).Dtor_exp())
	} else {
		return Companion_Option_.Create_None_()
	}
}
func (_static *CompanionStruct_Default___) AsInt(o Option) _dafny.Int {
	if (o).Is_None() {
		return _dafny.Zero
	} else {
		var _0_r Option = Companion_Default___.JsonInt((o).Dtor_value().(m_ConfluxCodec.Json))
		_ = _0_r
		if (_0_r).Is_Some() {
			return (_0_r).Dtor_value().(_dafny.Int)
		} else {
			return _dafny.Zero
		}
	}
}
func (_static *CompanionStruct_Default___) OptInt(o Option) Option {
	if (o).Is_Some() {
		return Companion_Default___.JsonInt((o).Dtor_value().(m_ConfluxCodec.Json))
	} else {
		return Companion_Option_.Create_None_()
	}
}
func (_static *CompanionStruct_Default___) EncodeStatus(s uint32) _dafny.Int {
	return _dafny.IntOfUint32(s)
}
func (_static *CompanionStruct_Default___) ApplyToRoot(s RootState, a RootAction, now _dafny.Int) _dafny.Tuple {
	var _source0 RootAction = a
	_ = _source0
	{
		if _source0.Is_RootAgentsChanged() {
			var _0_ag _dafny.Sequence = _source0.Get_().(RootAction_RootAgentsChanged).Agents
			_ = _0_ag
			return _dafny.TupleOf(func(_pat_let0_0 RootState) RootState {
				return func(_1_dt__update__tmp_h0 RootState) RootState {
					return func(_pat_let1_0 _dafny.Sequence) RootState {
						return func(_2_dt__update_hagents_h0 _dafny.Sequence) RootState {
							return Companion_RootState_.Create_RootState_(_2_dt__update_hagents_h0, (_1_dt__update__tmp_h0).Dtor_activeSessions(), (_1_dt__update__tmp_h0).Dtor_terminals(), (_1_dt__update__tmp_h0).Dtor_config())
						}(_pat_let1_0)
					}(_0_ag)
				}(_pat_let0_0)
			}(s), Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_RootActiveSessionsChanged() {
			var _3_n _dafny.Int = _source0.Get_().(RootAction_RootActiveSessionsChanged).ActiveSessions
			_ = _3_n
			return _dafny.TupleOf(func(_pat_let2_0 RootState) RootState {
				return func(_4_dt__update__tmp_h1 RootState) RootState {
					return func(_pat_let3_0 Option) RootState {
						return func(_5_dt__update_hactiveSessions_h0 Option) RootState {
							return Companion_RootState_.Create_RootState_((_4_dt__update__tmp_h1).Dtor_agents(), _5_dt__update_hactiveSessions_h0, (_4_dt__update__tmp_h1).Dtor_terminals(), (_4_dt__update__tmp_h1).Dtor_config())
						}(_pat_let3_0)
					}(Companion_Option_.Create_Some_(_3_n))
				}(_pat_let2_0)
			}(s), Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_RootTerminalsChanged() {
			var _6_t _dafny.Sequence = _source0.Get_().(RootAction_RootTerminalsChanged).Terminals
			_ = _6_t
			return _dafny.TupleOf(func(_pat_let4_0 RootState) RootState {
				return func(_7_dt__update__tmp_h2 RootState) RootState {
					return func(_pat_let5_0 Option) RootState {
						return func(_8_dt__update_hterminals_h0 Option) RootState {
							return Companion_RootState_.Create_RootState_((_7_dt__update__tmp_h2).Dtor_agents(), (_7_dt__update__tmp_h2).Dtor_activeSessions(), _8_dt__update_hterminals_h0, (_7_dt__update__tmp_h2).Dtor_config())
						}(_pat_let5_0)
					}(Companion_Option_.Create_Some_(_6_t))
				}(_pat_let4_0)
			}(s), Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_RootConfigChanged() {
			var _9_cfg _dafny.Map = _source0.Get_().(RootAction_RootConfigChanged).Config
			_ = _9_cfg
			var _10_repl bool = _source0.Get_().(RootAction_RootConfigChanged).Replace
			_ = _10_repl
			var _source1 Option = (s).Dtor_config()
			_ = _source1
			{
				if _source1.Is_None() {
					return _dafny.TupleOf(s, Companion_ReduceOutcome_.Create_NoOp_())
				}
			}
			{
				var _11_c RootConfig = _source1.Get_().(Option_Some).Value.(RootConfig)
				_ = _11_c
				var _12_newValues _dafny.Map = (func() _dafny.Map {
					if _10_repl {
						return _9_cfg
					}
					return ((_11_c).Dtor_values()).Merge(_9_cfg)
				})()
				_ = _12_newValues
				return _dafny.TupleOf(func(_pat_let6_0 RootState) RootState {
					return func(_13_dt__update__tmp_h3 RootState) RootState {
						return func(_pat_let7_0 Option) RootState {
							return func(_14_dt__update_hconfig_h0 Option) RootState {
								return Companion_RootState_.Create_RootState_((_13_dt__update__tmp_h3).Dtor_agents(), (_13_dt__update__tmp_h3).Dtor_activeSessions(), (_13_dt__update__tmp_h3).Dtor_terminals(), _14_dt__update_hconfig_h0)
							}(_pat_let7_0)
						}(Companion_Option_.Create_Some_(Companion_RootConfig_.Create_RootConfig_((_11_c).Dtor_schema(), _12_newValues)))
					}(_pat_let6_0)
				}(s), Companion_ReduceOutcome_.Create_Applied_())
			}
		}
	}
	{
		return _dafny.TupleOf(s, Companion_ReduceOutcome_.Create_NoOp_())
	}
}
func (_static *CompanionStruct_Default___) IsUnknownRoot(a RootAction) bool {
	return (a).Is_RootUnknown()
}
func (_static *CompanionStruct_Default___) ConfigPreconditionFails(s RootState, a RootAction) bool {
	return ((a).Is_RootConfigChanged()) && (((s).Dtor_config()).Is_None())
}
func (_static *CompanionStruct_Default___) FoldRoot(s RootState, actions _dafny.Sequence, now _dafny.Int) RootState {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer2 func(RootState, RootAction) RootState) func(interface{}, interface{}) interface{} {
		return func(arg2 interface{}, arg3 interface{}) interface{} {
			return coer2(arg2.(RootState), arg3.(RootAction))
		}
	}((func(_0_now _dafny.Int) func(RootState, RootAction) RootState {
		return func(_1_s_k RootState, _2_a RootAction) RootState {
			return (*(Companion_Default___.ApplyToRoot(_1_s_k, _2_a, _0_now)).IndexInt(0)).(RootState)
		}
	})(now)), s, actions).(RootState)
}
func (_static *CompanionStruct_Default___) Apply1(s RootState, a RootAction) RootState {
	return (*(Companion_Default___.ApplyToRoot(s, a, _dafny.IntOfInt64(9999))).IndexInt(0)).(RootState)
}
func (_static *CompanionStruct_Default___) RunCorpus() (_dafny.Int, _dafny.Int) {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	var total _dafny.Int = _dafny.Zero
	_ = total
	total = _dafny.IntOfInt64(7)
	pass = _dafny.Zero
	var _0_empty RootState
	_ = _0_empty
	_0_empty = Companion_RootState_.Create_RootState_(_dafny.SeqOf(), Companion_Option_.Create_None_(), Companion_Option_.Create_None_(), Companion_Option_.Create_None_())
	var _1_ag _dafny.Sequence
	_ = _1_ag
	_1_ag = _dafny.SeqOf(Companion_AgentInfo_.Create_AgentInfo_(_dafny.UnicodeSeqOfUtf8Bytes("copilot"), _dafny.UnicodeSeqOfUtf8Bytes("Copilot"), _dafny.UnicodeSeqOfUtf8Bytes("AI"), _dafny.SeqOf()))
	if (Companion_Default___.Apply1(_0_empty, Companion_RootAction_.Create_RootAgentsChanged_(_1_ag))).Equals(Companion_RootState_.Create_RootState_(_1_ag, Companion_Option_.Create_None_(), Companion_Option_.Create_None_(), Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_0_empty, Companion_RootAction_.Create_RootActiveSessionsChanged_(_dafny.IntOfInt64(5)))).Equals(Companion_RootState_.Create_RootState_(_dafny.SeqOf(), Companion_Option_.Create_Some_(_dafny.IntOfInt64(5)), Companion_Option_.Create_None_(), Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	var _2_terms _dafny.Sequence
	_ = _2_terms
	_2_terms = _dafny.SeqOf(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("resource"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("agenthost:/terminal/1")))), m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("resource"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("agenthost:/terminal/2")))))
	if (Companion_Default___.Apply1(_0_empty, Companion_RootAction_.Create_RootTerminalsChanged_(_2_terms))).Equals(Companion_RootState_.Create_RootState_(_dafny.SeqOf(), Companion_Option_.Create_None_(), Companion_Option_.Create_Some_(_2_terms), Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_0_empty, Companion_RootAction_.Create_RootUnknown_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("type"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("root/nonExistentAction"))))))).Equals(_0_empty) {
		pass = (pass).Plus(_dafny.One)
	}
	var _3_sch m_ConfluxCodec.Json
	_ = _3_sch
	_3_sch = m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("type"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("object"))))
	var _4_cfg127 RootState
	_ = _4_cfg127
	_4_cfg127 = Companion_RootState_.Create_RootState_(_dafny.SeqOf(), Companion_Option_.Create_None_(), Companion_Option_.Create_None_(), Companion_Option_.Create_Some_(Companion_RootConfig_.Create_RootConfig_(_3_sch, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("theme"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("light"))))))
	var _5_exp127 RootState
	_ = _5_exp127
	_5_exp127 = Companion_RootState_.Create_RootState_(_dafny.SeqOf(), Companion_Option_.Create_None_(), Companion_Option_.Create_None_(), Companion_Option_.Create_Some_(Companion_RootConfig_.Create_RootConfig_(_3_sch, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("theme"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("dark"))))))
	if (Companion_Default___.Apply1(_4_cfg127, Companion_RootAction_.Create_RootConfigChanged_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("theme"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("dark"))), false))).Equals(_5_exp127) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_0_empty, Companion_RootAction_.Create_RootConfigChanged_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("theme"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("dark"))), false))).Equals(_0_empty) {
		pass = (pass).Plus(_dafny.One)
	}
	var _6_cfg130 RootState
	_ = _6_cfg130
	_6_cfg130 = Companion_RootState_.Create_RootState_(_dafny.SeqOf(), Companion_Option_.Create_None_(), Companion_Option_.Create_None_(), Companion_Option_.Create_Some_(Companion_RootConfig_.Create_RootConfig_(_3_sch, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("theme"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("light"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("locale"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("en"))))))
	var _7_exp130 RootState
	_ = _7_exp130
	_7_exp130 = Companion_RootState_.Create_RootState_(_dafny.SeqOf(), Companion_Option_.Create_None_(), Companion_Option_.Create_None_(), Companion_Option_.Create_Some_(Companion_RootConfig_.Create_RootConfig_(_3_sch, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("theme"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("dark"))))))
	if (Companion_Default___.Apply1(_6_cfg130, Companion_RootAction_.Create_RootConfigChanged_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("theme"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("dark"))), true))).Equals(_7_exp130) {
		pass = (pass).Plus(_dafny.One)
	}
	return pass, total
}

// End of class Default__

// Definition of datatype Option
type Option struct {
	Data_Option_
}

func (_this Option) Get_() Data_Option_ {
	return _this.Data_Option_
}

type Data_Option_ interface {
	isOption()
}

type CompanionStruct_Option_ struct {
}

var Companion_Option_ = CompanionStruct_Option_{}

type Option_None struct {
}

func (Option_None) isOption() {}

func (CompanionStruct_Option_) Create_None_() Option {
	return Option{Option_None{}}
}

func (_this Option) Is_None() bool {
	_, ok := _this.Get_().(Option_None)
	return ok
}

type Option_Some struct {
	Value interface{}
}

func (Option_Some) isOption() {}

func (CompanionStruct_Option_) Create_Some_(Value interface{}) Option {
	return Option{Option_Some{Value}}
}

func (_this Option) Is_Some() bool {
	_, ok := _this.Get_().(Option_Some)
	return ok
}

func (CompanionStruct_Option_) Default() Option {
	return Companion_Option_.Create_None_()
}

func (_this Option) Dtor_value() interface{} {
	return _this.Get_().(Option_Some).Value
}

func (_this Option) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Option_None:
		{
			return "AhpSkeleton.Option.None"
		}
	case Option_Some:
		{
			return "AhpSkeleton.Option.Some" + "(" + _dafny.String(data.Value) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Option) Equals(other Option) bool {
	switch data1 := _this.Get_().(type) {
	case Option_None:
		{
			_, ok := other.Get_().(Option_None)
			return ok
		}
	case Option_Some:
		{
			data2, ok := other.Get_().(Option_Some)
			return ok && _dafny.AreEqual(data1.Value, data2.Value)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Option) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Option)
	return ok && _this.Equals(typed)
}

func Type_Option_() _dafny.TypeDescriptor {
	return type_Option_{}
}

type type_Option_ struct {
}

func (_this type_Option_) Default() interface{} {
	return Companion_Option_.Default()
}

func (_this type_Option_) String() string {
	return "AhpSkeleton.Option"
}
func (_this Option) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Option{}

// End of datatype Option

// Definition of datatype ReduceOutcome
type ReduceOutcome struct {
	Data_ReduceOutcome_
}

func (_this ReduceOutcome) Get_() Data_ReduceOutcome_ {
	return _this.Data_ReduceOutcome_
}

type Data_ReduceOutcome_ interface {
	isReduceOutcome()
}

type CompanionStruct_ReduceOutcome_ struct {
}

var Companion_ReduceOutcome_ = CompanionStruct_ReduceOutcome_{}

type ReduceOutcome_Applied struct {
}

func (ReduceOutcome_Applied) isReduceOutcome() {}

func (CompanionStruct_ReduceOutcome_) Create_Applied_() ReduceOutcome {
	return ReduceOutcome{ReduceOutcome_Applied{}}
}

func (_this ReduceOutcome) Is_Applied() bool {
	_, ok := _this.Get_().(ReduceOutcome_Applied)
	return ok
}

type ReduceOutcome_NoOp struct {
}

func (ReduceOutcome_NoOp) isReduceOutcome() {}

func (CompanionStruct_ReduceOutcome_) Create_NoOp_() ReduceOutcome {
	return ReduceOutcome{ReduceOutcome_NoOp{}}
}

func (_this ReduceOutcome) Is_NoOp() bool {
	_, ok := _this.Get_().(ReduceOutcome_NoOp)
	return ok
}

type ReduceOutcome_OutOfScope struct {
}

func (ReduceOutcome_OutOfScope) isReduceOutcome() {}

func (CompanionStruct_ReduceOutcome_) Create_OutOfScope_() ReduceOutcome {
	return ReduceOutcome{ReduceOutcome_OutOfScope{}}
}

func (_this ReduceOutcome) Is_OutOfScope() bool {
	_, ok := _this.Get_().(ReduceOutcome_OutOfScope)
	return ok
}

func (CompanionStruct_ReduceOutcome_) Default() ReduceOutcome {
	return Companion_ReduceOutcome_.Create_Applied_()
}

func (_ CompanionStruct_ReduceOutcome_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_ReduceOutcome_.Create_Applied_(), true
		case 1:
			return Companion_ReduceOutcome_.Create_NoOp_(), true
		case 2:
			return Companion_ReduceOutcome_.Create_OutOfScope_(), true
		default:
			return ReduceOutcome{}, false
		}
	}
}

func (_this ReduceOutcome) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case ReduceOutcome_Applied:
		{
			return "AhpSkeleton.ReduceOutcome.Applied"
		}
	case ReduceOutcome_NoOp:
		{
			return "AhpSkeleton.ReduceOutcome.NoOp"
		}
	case ReduceOutcome_OutOfScope:
		{
			return "AhpSkeleton.ReduceOutcome.OutOfScope"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ReduceOutcome) Equals(other ReduceOutcome) bool {
	switch _this.Get_().(type) {
	case ReduceOutcome_Applied:
		{
			_, ok := other.Get_().(ReduceOutcome_Applied)
			return ok
		}
	case ReduceOutcome_NoOp:
		{
			_, ok := other.Get_().(ReduceOutcome_NoOp)
			return ok
		}
	case ReduceOutcome_OutOfScope:
		{
			_, ok := other.Get_().(ReduceOutcome_OutOfScope)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ReduceOutcome) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ReduceOutcome)
	return ok && _this.Equals(typed)
}

func Type_ReduceOutcome_() _dafny.TypeDescriptor {
	return type_ReduceOutcome_{}
}

type type_ReduceOutcome_ struct {
}

func (_this type_ReduceOutcome_) Default() interface{} {
	return Companion_ReduceOutcome_.Default()
}

func (_this type_ReduceOutcome_) String() string {
	return "AhpSkeleton.ReduceOutcome"
}
func (_this ReduceOutcome) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ReduceOutcome{}

// End of datatype ReduceOutcome

// Definition of datatype ActionOrigin
type ActionOrigin struct {
	Data_ActionOrigin_
}

func (_this ActionOrigin) Get_() Data_ActionOrigin_ {
	return _this.Data_ActionOrigin_
}

type Data_ActionOrigin_ interface {
	isActionOrigin()
}

type CompanionStruct_ActionOrigin_ struct {
}

var Companion_ActionOrigin_ = CompanionStruct_ActionOrigin_{}

type ActionOrigin_ActionOrigin struct {
	ClientId  _dafny.Sequence
	ClientSeq _dafny.Int
}

func (ActionOrigin_ActionOrigin) isActionOrigin() {}

func (CompanionStruct_ActionOrigin_) Create_ActionOrigin_(ClientId _dafny.Sequence, ClientSeq _dafny.Int) ActionOrigin {
	return ActionOrigin{ActionOrigin_ActionOrigin{ClientId, ClientSeq}}
}

func (_this ActionOrigin) Is_ActionOrigin() bool {
	_, ok := _this.Get_().(ActionOrigin_ActionOrigin)
	return ok
}

func (CompanionStruct_ActionOrigin_) Default() ActionOrigin {
	return Companion_ActionOrigin_.Create_ActionOrigin_(_dafny.EmptySeq, _dafny.Zero)
}

func (_this ActionOrigin) Dtor_clientId() _dafny.Sequence {
	return _this.Get_().(ActionOrigin_ActionOrigin).ClientId
}

func (_this ActionOrigin) Dtor_clientSeq() _dafny.Int {
	return _this.Get_().(ActionOrigin_ActionOrigin).ClientSeq
}

func (_this ActionOrigin) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ActionOrigin_ActionOrigin:
		{
			return "AhpSkeleton.ActionOrigin.ActionOrigin" + "(" + data.ClientId.VerbatimString(true) + ", " + _dafny.String(data.ClientSeq) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ActionOrigin) Equals(other ActionOrigin) bool {
	switch data1 := _this.Get_().(type) {
	case ActionOrigin_ActionOrigin:
		{
			data2, ok := other.Get_().(ActionOrigin_ActionOrigin)
			return ok && data1.ClientId.Equals(data2.ClientId) && data1.ClientSeq.Cmp(data2.ClientSeq) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ActionOrigin) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ActionOrigin)
	return ok && _this.Equals(typed)
}

func Type_ActionOrigin_() _dafny.TypeDescriptor {
	return type_ActionOrigin_{}
}

type type_ActionOrigin_ struct {
}

func (_this type_ActionOrigin_) Default() interface{} {
	return Companion_ActionOrigin_.Default()
}

func (_this type_ActionOrigin_) String() string {
	return "AhpSkeleton.ActionOrigin"
}
func (_this ActionOrigin) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ActionOrigin{}

// End of datatype ActionOrigin

// Definition of datatype Envelope
type Envelope struct {
	Data_Envelope_
}

func (_this Envelope) Get_() Data_Envelope_ {
	return _this.Data_Envelope_
}

type Data_Envelope_ interface {
	isEnvelope()
}

type CompanionStruct_Envelope_ struct {
}

var Companion_Envelope_ = CompanionStruct_Envelope_{}

type Envelope_Envelope struct {
	Channel         _dafny.Sequence
	Action          interface{}
	ServerSeq       _dafny.Int
	Origin          Option
	RejectionReason Option
}

func (Envelope_Envelope) isEnvelope() {}

func (CompanionStruct_Envelope_) Create_Envelope_(Channel _dafny.Sequence, Action interface{}, ServerSeq _dafny.Int, Origin Option, RejectionReason Option) Envelope {
	return Envelope{Envelope_Envelope{Channel, Action, ServerSeq, Origin, RejectionReason}}
}

func (_this Envelope) Is_Envelope() bool {
	_, ok := _this.Get_().(Envelope_Envelope)
	return ok
}

func (CompanionStruct_Envelope_) Default(_default_A interface{}) Envelope {
	return Companion_Envelope_.Create_Envelope_(_dafny.EmptySeq, _default_A, _dafny.Zero, Companion_Option_.Default(), Companion_Option_.Default())
}

func (_this Envelope) Dtor_channel() _dafny.Sequence {
	return _this.Get_().(Envelope_Envelope).Channel
}

func (_this Envelope) Dtor_action() interface{} {
	return _this.Get_().(Envelope_Envelope).Action
}

func (_this Envelope) Dtor_serverSeq() _dafny.Int {
	return _this.Get_().(Envelope_Envelope).ServerSeq
}

func (_this Envelope) Dtor_origin() Option {
	return _this.Get_().(Envelope_Envelope).Origin
}

func (_this Envelope) Dtor_rejectionReason() Option {
	return _this.Get_().(Envelope_Envelope).RejectionReason
}

func (_this Envelope) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Envelope_Envelope:
		{
			return "AhpSkeleton.Envelope.Envelope" + "(" + data.Channel.VerbatimString(true) + ", " + _dafny.String(data.Action) + ", " + _dafny.String(data.ServerSeq) + ", " + _dafny.String(data.Origin) + ", " + _dafny.String(data.RejectionReason) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Envelope) Equals(other Envelope) bool {
	switch data1 := _this.Get_().(type) {
	case Envelope_Envelope:
		{
			data2, ok := other.Get_().(Envelope_Envelope)
			return ok && data1.Channel.Equals(data2.Channel) && _dafny.AreEqual(data1.Action, data2.Action) && data1.ServerSeq.Cmp(data2.ServerSeq) == 0 && data1.Origin.Equals(data2.Origin) && data1.RejectionReason.Equals(data2.RejectionReason)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Envelope) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Envelope)
	return ok && _this.Equals(typed)
}

func Type_Envelope_(Type_A_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_Envelope_{Type_A_}
}

type type_Envelope_ struct {
	Type_A_ _dafny.TypeDescriptor
}

func (_this type_Envelope_) Default() interface{} {
	Type_A_ := _this.Type_A_
	_ = Type_A_
	return Companion_Envelope_.Default(Type_A_.Default())
}

func (_this type_Envelope_) String() string {
	return "AhpSkeleton.Envelope"
}
func (_this Envelope) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Envelope{}

// End of datatype Envelope

// Definition of datatype AgentInfo
type AgentInfo struct {
	Data_AgentInfo_
}

func (_this AgentInfo) Get_() Data_AgentInfo_ {
	return _this.Data_AgentInfo_
}

type Data_AgentInfo_ interface {
	isAgentInfo()
}

type CompanionStruct_AgentInfo_ struct {
}

var Companion_AgentInfo_ = CompanionStruct_AgentInfo_{}

type AgentInfo_AgentInfo struct {
	Provider    _dafny.Sequence
	DisplayName _dafny.Sequence
	Description _dafny.Sequence
	Models      _dafny.Sequence
}

func (AgentInfo_AgentInfo) isAgentInfo() {}

func (CompanionStruct_AgentInfo_) Create_AgentInfo_(Provider _dafny.Sequence, DisplayName _dafny.Sequence, Description _dafny.Sequence, Models _dafny.Sequence) AgentInfo {
	return AgentInfo{AgentInfo_AgentInfo{Provider, DisplayName, Description, Models}}
}

func (_this AgentInfo) Is_AgentInfo() bool {
	_, ok := _this.Get_().(AgentInfo_AgentInfo)
	return ok
}

func (CompanionStruct_AgentInfo_) Default() AgentInfo {
	return Companion_AgentInfo_.Create_AgentInfo_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this AgentInfo) Dtor_provider() _dafny.Sequence {
	return _this.Get_().(AgentInfo_AgentInfo).Provider
}

func (_this AgentInfo) Dtor_displayName() _dafny.Sequence {
	return _this.Get_().(AgentInfo_AgentInfo).DisplayName
}

func (_this AgentInfo) Dtor_description() _dafny.Sequence {
	return _this.Get_().(AgentInfo_AgentInfo).Description
}

func (_this AgentInfo) Dtor_models() _dafny.Sequence {
	return _this.Get_().(AgentInfo_AgentInfo).Models
}

func (_this AgentInfo) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case AgentInfo_AgentInfo:
		{
			return "AhpSkeleton.AgentInfo.AgentInfo" + "(" + data.Provider.VerbatimString(true) + ", " + data.DisplayName.VerbatimString(true) + ", " + data.Description.VerbatimString(true) + ", " + _dafny.String(data.Models) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this AgentInfo) Equals(other AgentInfo) bool {
	switch data1 := _this.Get_().(type) {
	case AgentInfo_AgentInfo:
		{
			data2, ok := other.Get_().(AgentInfo_AgentInfo)
			return ok && data1.Provider.Equals(data2.Provider) && data1.DisplayName.Equals(data2.DisplayName) && data1.Description.Equals(data2.Description) && data1.Models.Equals(data2.Models)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this AgentInfo) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(AgentInfo)
	return ok && _this.Equals(typed)
}

func Type_AgentInfo_() _dafny.TypeDescriptor {
	return type_AgentInfo_{}
}

type type_AgentInfo_ struct {
}

func (_this type_AgentInfo_) Default() interface{} {
	return Companion_AgentInfo_.Default()
}

func (_this type_AgentInfo_) String() string {
	return "AhpSkeleton.AgentInfo"
}
func (_this AgentInfo) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = AgentInfo{}

// End of datatype AgentInfo

// Definition of datatype RootConfig
type RootConfig struct {
	Data_RootConfig_
}

func (_this RootConfig) Get_() Data_RootConfig_ {
	return _this.Data_RootConfig_
}

type Data_RootConfig_ interface {
	isRootConfig()
}

type CompanionStruct_RootConfig_ struct {
}

var Companion_RootConfig_ = CompanionStruct_RootConfig_{}

type RootConfig_RootConfig struct {
	Schema m_ConfluxCodec.Json
	Values _dafny.Map
}

func (RootConfig_RootConfig) isRootConfig() {}

func (CompanionStruct_RootConfig_) Create_RootConfig_(Schema m_ConfluxCodec.Json, Values _dafny.Map) RootConfig {
	return RootConfig{RootConfig_RootConfig{Schema, Values}}
}

func (_this RootConfig) Is_RootConfig() bool {
	_, ok := _this.Get_().(RootConfig_RootConfig)
	return ok
}

func (CompanionStruct_RootConfig_) Default() RootConfig {
	return Companion_RootConfig_.Create_RootConfig_(m_ConfluxCodec.Companion_Json_.Default(), _dafny.EmptyMap)
}

func (_this RootConfig) Dtor_schema() m_ConfluxCodec.Json {
	return _this.Get_().(RootConfig_RootConfig).Schema
}

func (_this RootConfig) Dtor_values() _dafny.Map {
	return _this.Get_().(RootConfig_RootConfig).Values
}

func (_this RootConfig) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RootConfig_RootConfig:
		{
			return "AhpSkeleton.RootConfig.RootConfig" + "(" + _dafny.String(data.Schema) + ", " + _dafny.String(data.Values) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RootConfig) Equals(other RootConfig) bool {
	switch data1 := _this.Get_().(type) {
	case RootConfig_RootConfig:
		{
			data2, ok := other.Get_().(RootConfig_RootConfig)
			return ok && data1.Schema.Equals(data2.Schema) && data1.Values.Equals(data2.Values)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RootConfig) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RootConfig)
	return ok && _this.Equals(typed)
}

func Type_RootConfig_() _dafny.TypeDescriptor {
	return type_RootConfig_{}
}

type type_RootConfig_ struct {
}

func (_this type_RootConfig_) Default() interface{} {
	return Companion_RootConfig_.Default()
}

func (_this type_RootConfig_) String() string {
	return "AhpSkeleton.RootConfig"
}
func (_this RootConfig) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RootConfig{}

// End of datatype RootConfig

// Definition of datatype RootState
type RootState struct {
	Data_RootState_
}

func (_this RootState) Get_() Data_RootState_ {
	return _this.Data_RootState_
}

type Data_RootState_ interface {
	isRootState()
}

type CompanionStruct_RootState_ struct {
}

var Companion_RootState_ = CompanionStruct_RootState_{}

type RootState_RootState struct {
	Agents         _dafny.Sequence
	ActiveSessions Option
	Terminals      Option
	Config         Option
}

func (RootState_RootState) isRootState() {}

func (CompanionStruct_RootState_) Create_RootState_(Agents _dafny.Sequence, ActiveSessions Option, Terminals Option, Config Option) RootState {
	return RootState{RootState_RootState{Agents, ActiveSessions, Terminals, Config}}
}

func (_this RootState) Is_RootState() bool {
	_, ok := _this.Get_().(RootState_RootState)
	return ok
}

func (CompanionStruct_RootState_) Default() RootState {
	return Companion_RootState_.Create_RootState_(_dafny.EmptySeq, Companion_Option_.Default(), Companion_Option_.Default(), Companion_Option_.Default())
}

func (_this RootState) Dtor_agents() _dafny.Sequence {
	return _this.Get_().(RootState_RootState).Agents
}

func (_this RootState) Dtor_activeSessions() Option {
	return _this.Get_().(RootState_RootState).ActiveSessions
}

func (_this RootState) Dtor_terminals() Option {
	return _this.Get_().(RootState_RootState).Terminals
}

func (_this RootState) Dtor_config() Option {
	return _this.Get_().(RootState_RootState).Config
}

func (_this RootState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RootState_RootState:
		{
			return "AhpSkeleton.RootState.RootState" + "(" + _dafny.String(data.Agents) + ", " + _dafny.String(data.ActiveSessions) + ", " + _dafny.String(data.Terminals) + ", " + _dafny.String(data.Config) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RootState) Equals(other RootState) bool {
	switch data1 := _this.Get_().(type) {
	case RootState_RootState:
		{
			data2, ok := other.Get_().(RootState_RootState)
			return ok && data1.Agents.Equals(data2.Agents) && data1.ActiveSessions.Equals(data2.ActiveSessions) && data1.Terminals.Equals(data2.Terminals) && data1.Config.Equals(data2.Config)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RootState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RootState)
	return ok && _this.Equals(typed)
}

func Type_RootState_() _dafny.TypeDescriptor {
	return type_RootState_{}
}

type type_RootState_ struct {
}

func (_this type_RootState_) Default() interface{} {
	return Companion_RootState_.Default()
}

func (_this type_RootState_) String() string {
	return "AhpSkeleton.RootState"
}
func (_this RootState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RootState{}

// End of datatype RootState

// Definition of datatype RootAction
type RootAction struct {
	Data_RootAction_
}

func (_this RootAction) Get_() Data_RootAction_ {
	return _this.Data_RootAction_
}

type Data_RootAction_ interface {
	isRootAction()
}

type CompanionStruct_RootAction_ struct {
}

var Companion_RootAction_ = CompanionStruct_RootAction_{}

type RootAction_RootAgentsChanged struct {
	Agents _dafny.Sequence
}

func (RootAction_RootAgentsChanged) isRootAction() {}

func (CompanionStruct_RootAction_) Create_RootAgentsChanged_(Agents _dafny.Sequence) RootAction {
	return RootAction{RootAction_RootAgentsChanged{Agents}}
}

func (_this RootAction) Is_RootAgentsChanged() bool {
	_, ok := _this.Get_().(RootAction_RootAgentsChanged)
	return ok
}

type RootAction_RootActiveSessionsChanged struct {
	ActiveSessions _dafny.Int
}

func (RootAction_RootActiveSessionsChanged) isRootAction() {}

func (CompanionStruct_RootAction_) Create_RootActiveSessionsChanged_(ActiveSessions _dafny.Int) RootAction {
	return RootAction{RootAction_RootActiveSessionsChanged{ActiveSessions}}
}

func (_this RootAction) Is_RootActiveSessionsChanged() bool {
	_, ok := _this.Get_().(RootAction_RootActiveSessionsChanged)
	return ok
}

type RootAction_RootTerminalsChanged struct {
	Terminals _dafny.Sequence
}

func (RootAction_RootTerminalsChanged) isRootAction() {}

func (CompanionStruct_RootAction_) Create_RootTerminalsChanged_(Terminals _dafny.Sequence) RootAction {
	return RootAction{RootAction_RootTerminalsChanged{Terminals}}
}

func (_this RootAction) Is_RootTerminalsChanged() bool {
	_, ok := _this.Get_().(RootAction_RootTerminalsChanged)
	return ok
}

type RootAction_RootConfigChanged struct {
	Config  _dafny.Map
	Replace bool
}

func (RootAction_RootConfigChanged) isRootAction() {}

func (CompanionStruct_RootAction_) Create_RootConfigChanged_(Config _dafny.Map, Replace bool) RootAction {
	return RootAction{RootAction_RootConfigChanged{Config, Replace}}
}

func (_this RootAction) Is_RootConfigChanged() bool {
	_, ok := _this.Get_().(RootAction_RootConfigChanged)
	return ok
}

type RootAction_RootUnknown struct {
	Raw m_ConfluxCodec.Json
}

func (RootAction_RootUnknown) isRootAction() {}

func (CompanionStruct_RootAction_) Create_RootUnknown_(Raw m_ConfluxCodec.Json) RootAction {
	return RootAction{RootAction_RootUnknown{Raw}}
}

func (_this RootAction) Is_RootUnknown() bool {
	_, ok := _this.Get_().(RootAction_RootUnknown)
	return ok
}

func (CompanionStruct_RootAction_) Default() RootAction {
	return Companion_RootAction_.Create_RootAgentsChanged_(_dafny.EmptySeq)
}

func (_this RootAction) Dtor_agents() _dafny.Sequence {
	return _this.Get_().(RootAction_RootAgentsChanged).Agents
}

func (_this RootAction) Dtor_activeSessions() _dafny.Int {
	return _this.Get_().(RootAction_RootActiveSessionsChanged).ActiveSessions
}

func (_this RootAction) Dtor_terminals() _dafny.Sequence {
	return _this.Get_().(RootAction_RootTerminalsChanged).Terminals
}

func (_this RootAction) Dtor_config() _dafny.Map {
	return _this.Get_().(RootAction_RootConfigChanged).Config
}

func (_this RootAction) Dtor_replace() bool {
	return _this.Get_().(RootAction_RootConfigChanged).Replace
}

func (_this RootAction) Dtor_raw() m_ConfluxCodec.Json {
	return _this.Get_().(RootAction_RootUnknown).Raw
}

func (_this RootAction) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RootAction_RootAgentsChanged:
		{
			return "AhpSkeleton.RootAction.RootAgentsChanged" + "(" + _dafny.String(data.Agents) + ")"
		}
	case RootAction_RootActiveSessionsChanged:
		{
			return "AhpSkeleton.RootAction.RootActiveSessionsChanged" + "(" + _dafny.String(data.ActiveSessions) + ")"
		}
	case RootAction_RootTerminalsChanged:
		{
			return "AhpSkeleton.RootAction.RootTerminalsChanged" + "(" + _dafny.String(data.Terminals) + ")"
		}
	case RootAction_RootConfigChanged:
		{
			return "AhpSkeleton.RootAction.RootConfigChanged" + "(" + _dafny.String(data.Config) + ", " + _dafny.String(data.Replace) + ")"
		}
	case RootAction_RootUnknown:
		{
			return "AhpSkeleton.RootAction.RootUnknown" + "(" + _dafny.String(data.Raw) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RootAction) Equals(other RootAction) bool {
	switch data1 := _this.Get_().(type) {
	case RootAction_RootAgentsChanged:
		{
			data2, ok := other.Get_().(RootAction_RootAgentsChanged)
			return ok && data1.Agents.Equals(data2.Agents)
		}
	case RootAction_RootActiveSessionsChanged:
		{
			data2, ok := other.Get_().(RootAction_RootActiveSessionsChanged)
			return ok && data1.ActiveSessions.Cmp(data2.ActiveSessions) == 0
		}
	case RootAction_RootTerminalsChanged:
		{
			data2, ok := other.Get_().(RootAction_RootTerminalsChanged)
			return ok && data1.Terminals.Equals(data2.Terminals)
		}
	case RootAction_RootConfigChanged:
		{
			data2, ok := other.Get_().(RootAction_RootConfigChanged)
			return ok && data1.Config.Equals(data2.Config) && data1.Replace == data2.Replace
		}
	case RootAction_RootUnknown:
		{
			data2, ok := other.Get_().(RootAction_RootUnknown)
			return ok && data1.Raw.Equals(data2.Raw)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RootAction) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RootAction)
	return ok && _this.Equals(typed)
}

func Type_RootAction_() _dafny.TypeDescriptor {
	return type_RootAction_{}
}

type type_RootAction_ struct {
}

func (_this type_RootAction_) Default() interface{} {
	return Companion_RootAction_.Default()
}

func (_this type_RootAction_) String() string {
	return "AhpSkeleton.RootAction"
}
func (_this RootAction) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RootAction{}

// End of datatype RootAction
