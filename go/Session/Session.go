// Package Session
// Dafny module Session compiled into Go

package Session

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-go/Changeset"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-go/ConfluxSeqRoute"
	m_ResourceWatch "github.com/joshmouch/ahp-go/ResourceWatch"
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
	return "Session.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) SetBit(s uint32, b uint32) uint32 {
	return (s) | (b)
}
func (_static *CompanionStruct_Default___) ClearBit(s uint32, b uint32) uint32 {
	return (s) & (^(b))
}
func (_static *CompanionStruct_Default___) ToggleEnabled(en bool) func(Cust) Cust {
	return (func(_0_en bool) func(Cust) Cust {
		return func(_1_e Cust) Cust {
			return func(_pat_let102_0 Cust) Cust {
				return func(_2_dt__update__tmp_h0 Cust) Cust {
					return func(_pat_let103_0 m_AhpSkeleton.Option) Cust {
						return func(_3_dt__update_henabled_h0 m_AhpSkeleton.Option) Cust {
							return Companion_Cust_.Create_Cust_((_2_dt__update__tmp_h0).Dtor_id(), (_2_dt__update__tmp_h0).Dtor_kind(), (_2_dt__update__tmp_h0).Dtor_uri(), (_2_dt__update__tmp_h0).Dtor_name(), _3_dt__update_henabled_h0, (_2_dt__update__tmp_h0).Dtor_state(), (_2_dt__update__tmp_h0).Dtor_channel())
						}(_pat_let103_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_0_en))
				}(_pat_let102_0)
			}(_1_e)
		}
	})(en)
}
func (_static *CompanionStruct_Default___) ChatPatch(st m_AhpSkeleton.Option, ac m_AhpSkeleton.Option, md m_AhpSkeleton.Option) func(Chat) Chat {
	return (func(_0_md m_AhpSkeleton.Option, _1_ac m_AhpSkeleton.Option, _2_st m_AhpSkeleton.Option) func(Chat) Chat {
		return func(_3_c Chat) Chat {
			return func(_pat_let104_0 Chat) Chat {
				return func(_4_dt__update__tmp_h0 Chat) Chat {
					return func(_pat_let105_0 _dafny.Sequence) Chat {
						return func(_5_dt__update_hmodifiedAt_h0 _dafny.Sequence) Chat {
							return func(_pat_let106_0 m_AhpSkeleton.Option) Chat {
								return func(_6_dt__update_hactivity_h0 m_AhpSkeleton.Option) Chat {
									return func(_pat_let107_0 _dafny.Int) Chat {
										return func(_7_dt__update_hstatus_h0 _dafny.Int) Chat {
											return Companion_Chat_.Create_Chat_((_4_dt__update__tmp_h0).Dtor_resource(), (_4_dt__update__tmp_h0).Dtor_title(), _7_dt__update_hstatus_h0, _6_dt__update_hactivity_h0, _5_dt__update_hmodifiedAt_h0, (_4_dt__update__tmp_h0).Dtor_origin())
										}(_pat_let107_0)
									}((func() _dafny.Int {
										if (_2_st).Is_Some() {
											return (_2_st).Dtor_value().(_dafny.Int)
										}
										return (_3_c).Dtor_status()
									})())
								}(_pat_let106_0)
							}((func() m_AhpSkeleton.Option {
								if (_1_ac).Is_Some() {
									return _1_ac
								}
								return (_3_c).Dtor_activity()
							})())
						}(_pat_let105_0)
					}((func() _dafny.Sequence {
						if (_0_md).Is_Some() {
							return (_0_md).Dtor_value().(_dafny.Sequence)
						}
						return (_3_c).Dtor_modifiedAt()
					})())
				}(_pat_let104_0)
			}(_3_c)
		}
	})(md, ac, st)
}
func (_static *CompanionStruct_Default___) McpPatch(st m_AhpSkeleton.Option, ch m_AhpSkeleton.Option) func(Cust) Cust {
	return (func(_0_ch m_AhpSkeleton.Option, _1_st m_AhpSkeleton.Option) func(Cust) Cust {
		return func(_2_e Cust) Cust {
			return func(_pat_let108_0 Cust) Cust {
				return func(_3_dt__update__tmp_h0 Cust) Cust {
					return func(_pat_let109_0 m_AhpSkeleton.Option) Cust {
						return func(_4_dt__update_hchannel_h0 m_AhpSkeleton.Option) Cust {
							return func(_pat_let110_0 m_AhpSkeleton.Option) Cust {
								return func(_5_dt__update_hstate_h0 m_AhpSkeleton.Option) Cust {
									return Companion_Cust_.Create_Cust_((_3_dt__update__tmp_h0).Dtor_id(), (_3_dt__update__tmp_h0).Dtor_kind(), (_3_dt__update__tmp_h0).Dtor_uri(), (_3_dt__update__tmp_h0).Dtor_name(), (_3_dt__update__tmp_h0).Dtor_enabled(), _5_dt__update_hstate_h0, _4_dt__update_hchannel_h0)
								}(_pat_let110_0)
							}(_1_st)
						}(_pat_let109_0)
					}(_0_ch)
				}(_pat_let108_0)
			}(_2_e)
		}
	})(ch, st)
}
func (_static *CompanionStruct_Default___) UpsertCust(cs _dafny.Sequence, id _dafny.Sequence, en bool) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpdateById(func(coer23 func(Cust) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg30 interface{}) interface{} {
			return coer23(arg30.(Cust))
		}
	}(Companion_Default___.CustKey()), cs, id, func(coer24 func(Cust) Cust) func(interface{}) interface{} {
		return func(arg31 interface{}) interface{} {
			return coer24(arg31.(Cust))
		}
	}(Companion_Default___.ToggleEnabled(en)))
}
func (_static *CompanionStruct_Default___) RemoveCust(cs _dafny.Sequence, id _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqRemoveById(func(coer25 func(Cust) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg32 interface{}) interface{} {
			return coer25(arg32.(Cust))
		}
	}(Companion_Default___.CustKey()), cs, id)
}
func (_static *CompanionStruct_Default___) RemoveChat(chs _dafny.Sequence, r _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqRemoveById(func(coer26 func(Chat) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg33 interface{}) interface{} {
			return coer26(arg33.(Chat))
		}
	}(Companion_Default___.ChatKey()), chs, r)
}
func (_static *CompanionStruct_Default___) UpsertChat(chs _dafny.Sequence, c Chat) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpsertById(func(coer27 func(Chat) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg34 interface{}) interface{} {
			return coer27(arg34.(Chat))
		}
	}(Companion_Default___.ChatKey()), chs, c)
}
func (_static *CompanionStruct_Default___) UpdChat(chs _dafny.Sequence, r _dafny.Sequence, st m_AhpSkeleton.Option, ac m_AhpSkeleton.Option, md m_AhpSkeleton.Option) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpdateById(func(coer28 func(Chat) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg35 interface{}) interface{} {
			return coer28(arg35.(Chat))
		}
	}(Companion_Default___.ChatKey()), chs, r, func(coer29 func(Chat) Chat) func(interface{}) interface{} {
		return func(arg36 interface{}) interface{} {
			return coer29(arg36.(Chat))
		}
	}(Companion_Default___.ChatPatch(st, ac, md)))
}
func (_static *CompanionStruct_Default___) UpsertClient(cl _dafny.Sequence, c Client) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpsertById(func(coer30 func(Client) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg37 interface{}) interface{} {
			return coer30(arg37.(Client))
		}
	}(Companion_Default___.ClientKey()), cl, c)
}
func (_static *CompanionStruct_Default___) RemoveClient(cl _dafny.Sequence, id _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqRemoveById(func(coer31 func(Client) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg38 interface{}) interface{} {
			return coer31(arg38.(Client))
		}
	}(Companion_Default___.ClientKey()), cl, id)
}
func (_static *CompanionStruct_Default___) UpsertInput(ins _dafny.Sequence, r InputReq) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpsertById(func(coer32 func(InputReq) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg39 interface{}) interface{} {
			return coer32(arg39.(InputReq))
		}
	}(Companion_Default___.InputKey()), ins, r)
}
func (_static *CompanionStruct_Default___) RemoveInput(ins _dafny.Sequence, id _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqRemoveById(func(coer33 func(InputReq) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg40 interface{}) interface{} {
			return coer33(arg40.(InputReq))
		}
	}(Companion_Default___.InputKey()), ins, id)
}
func (_static *CompanionStruct_Default___) UpsertCustFull(cs _dafny.Sequence, c Cust) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpsertById(func(coer34 func(Cust) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg41 interface{}) interface{} {
			return coer34(arg41.(Cust))
		}
	}(Companion_Default___.CustKey()), cs, c)
}
func (_static *CompanionStruct_Default___) SetMcp(cs _dafny.Sequence, id _dafny.Sequence, st m_AhpSkeleton.Option, ch m_AhpSkeleton.Option) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpdateByIdWhere(func(coer35 func(Cust) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg42 interface{}) interface{} {
			return coer35(arg42.(Cust))
		}
	}(Companion_Default___.CustKey()), func(coer36 func(Cust) bool) func(interface{}) bool {
		return func(arg43 interface{}) bool {
			return coer36(arg43.(Cust))
		}
	}(Companion_Default___.McpPred()), cs, id, func(coer37 func(Cust) Cust) func(interface{}) interface{} {
		return func(arg44 interface{}) interface{} {
			return coer37(arg44.(Cust))
		}
	}(Companion_Default___.McpPatch(st, ch)))
}
func (_static *CompanionStruct_Default___) IsMcp(cs _dafny.Sequence, id _dafny.Sequence) bool {
	return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((cs).Cardinality())), false, func(_exists_var_0 _dafny.Int) bool {
		var _0_i _dafny.Int
		_0_i = interface{}(_exists_var_0).(_dafny.Int)
		return ((((_0_i).Sign() != -1) && ((_0_i).Cmp(_dafny.IntOfUint32((cs).Cardinality())) < 0)) && (_dafny.Companion_Sequence_.Equal(((cs).Select((_0_i).Uint32()).(Cust)).Dtor_id(), id))) && (_dafny.Companion_Sequence_.Equal(((cs).Select((_0_i).Uint32()).(Cust)).Dtor_kind(), _dafny.UnicodeSeqOfUtf8Bytes("mcpServer")))
	})
}
func (_static *CompanionStruct_Default___) ApplyToSession(s SessionState, a SessionAction, now _dafny.Int) _dafny.Tuple {
	var _pat_let_tv0 = s
	_ = _pat_let_tv0
	var _pat_let_tv1 = s
	_ = _pat_let_tv1
	var _pat_let_tv2 = s
	_ = _pat_let_tv2
	var _pat_let_tv3 = s
	_ = _pat_let_tv3
	var _pat_let_tv4 = s
	_ = _pat_let_tv4
	var _pat_let_tv5 = s
	_ = _pat_let_tv5
	var _pat_let_tv6 = s
	_ = _pat_let_tv6
	var _pat_let_tv7 = s
	_ = _pat_let_tv7
	var _pat_let_tv8 = s
	_ = _pat_let_tv8
	var _pat_let_tv9 = s
	_ = _pat_let_tv9
	var _pat_let_tv10 = s
	_ = _pat_let_tv10
	var _pat_let_tv11 = s
	_ = _pat_let_tv11
	var _pat_let_tv12 = s
	_ = _pat_let_tv12
	var _pat_let_tv13 = s
	_ = _pat_let_tv13
	var _pat_let_tv14 = s
	_ = _pat_let_tv14
	var _pat_let_tv15 = s
	_ = _pat_let_tv15
	var _pat_let_tv16 = s
	_ = _pat_let_tv16
	var _pat_let_tv17 = s
	_ = _pat_let_tv17
	var _pat_let_tv18 = s
	_ = _pat_let_tv18
	var _pat_let_tv19 = s
	_ = _pat_let_tv19
	var _pat_let_tv20 = s
	_ = _pat_let_tv20
	var _source0 SessionAction = a
	_ = _source0
	{
		if _source0.Is_Ready() {
			return _dafny.TupleOf(func(_pat_let111_0 SessionState) SessionState {
				return func(_0_dt__update__tmp_h0 SessionState) SessionState {
					return func(_pat_let112_0 _dafny.Sequence) SessionState {
						return func(_1_dt__update_hlifecycle_h0 _dafny.Sequence) SessionState {
							return Companion_SessionState_.Create_SessionState_((_0_dt__update__tmp_h0).Dtor_provider(), (_0_dt__update__tmp_h0).Dtor_title(), (_0_dt__update__tmp_h0).Dtor_status(), _1_dt__update_hlifecycle_h0, (_0_dt__update__tmp_h0).Dtor_activity(), (_0_dt__update__tmp_h0).Dtor_config(), (_0_dt__update__tmp_h0).Dtor_meta(), (_0_dt__update__tmp_h0).Dtor_creationError(), (_0_dt__update__tmp_h0).Dtor_serverTools(), (_0_dt__update__tmp_h0).Dtor_changesets(), (_0_dt__update__tmp_h0).Dtor_canvases(), (_0_dt__update__tmp_h0).Dtor_openCanvases(), (_0_dt__update__tmp_h0).Dtor_defaultChat(), (_0_dt__update__tmp_h0).Dtor_activeClients(), (_0_dt__update__tmp_h0).Dtor_chats(), (_0_dt__update__tmp_h0).Dtor_customizations(), (_0_dt__update__tmp_h0).Dtor_inputNeeded())
						}(_pat_let112_0)
					}(_dafny.UnicodeSeqOfUtf8Bytes("ready"))
				}(_pat_let111_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_CreationFailed() {
			var _2_e m_ConfluxCodec.Json = _source0.Get_().(SessionAction_CreationFailed).Error
			_ = _2_e
			return _dafny.TupleOf(func(_pat_let113_0 SessionState) SessionState {
				return func(_3_dt__update__tmp_h1 SessionState) SessionState {
					return func(_pat_let114_0 m_AhpSkeleton.Option) SessionState {
						return func(_4_dt__update_hcreationError_h0 m_AhpSkeleton.Option) SessionState {
							return func(_pat_let115_0 _dafny.Sequence) SessionState {
								return func(_5_dt__update_hlifecycle_h1 _dafny.Sequence) SessionState {
									return Companion_SessionState_.Create_SessionState_((_3_dt__update__tmp_h1).Dtor_provider(), (_3_dt__update__tmp_h1).Dtor_title(), (_3_dt__update__tmp_h1).Dtor_status(), _5_dt__update_hlifecycle_h1, (_3_dt__update__tmp_h1).Dtor_activity(), (_3_dt__update__tmp_h1).Dtor_config(), (_3_dt__update__tmp_h1).Dtor_meta(), _4_dt__update_hcreationError_h0, (_3_dt__update__tmp_h1).Dtor_serverTools(), (_3_dt__update__tmp_h1).Dtor_changesets(), (_3_dt__update__tmp_h1).Dtor_canvases(), (_3_dt__update__tmp_h1).Dtor_openCanvases(), (_3_dt__update__tmp_h1).Dtor_defaultChat(), (_3_dt__update__tmp_h1).Dtor_activeClients(), (_3_dt__update__tmp_h1).Dtor_chats(), (_3_dt__update__tmp_h1).Dtor_customizations(), (_3_dt__update__tmp_h1).Dtor_inputNeeded())
								}(_pat_let115_0)
							}(_dafny.UnicodeSeqOfUtf8Bytes("creationFailed"))
						}(_pat_let114_0)
					}((func() m_AhpSkeleton.Option {
						if (_2_e).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_()) {
							return m_AhpSkeleton.Companion_Option_.Create_None_()
						}
						return m_AhpSkeleton.Companion_Option_.Create_Some_(_2_e)
					})())
				}(_pat_let113_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_TitleChanged() {
			var _6_t _dafny.Sequence = _source0.Get_().(SessionAction_TitleChanged).Title
			_ = _6_t
			return _dafny.TupleOf(func(_pat_let116_0 SessionState) SessionState {
				return func(_7_dt__update__tmp_h2 SessionState) SessionState {
					return func(_pat_let117_0 _dafny.Sequence) SessionState {
						return func(_8_dt__update_htitle_h0 _dafny.Sequence) SessionState {
							return Companion_SessionState_.Create_SessionState_((_7_dt__update__tmp_h2).Dtor_provider(), _8_dt__update_htitle_h0, (_7_dt__update__tmp_h2).Dtor_status(), (_7_dt__update__tmp_h2).Dtor_lifecycle(), (_7_dt__update__tmp_h2).Dtor_activity(), (_7_dt__update__tmp_h2).Dtor_config(), (_7_dt__update__tmp_h2).Dtor_meta(), (_7_dt__update__tmp_h2).Dtor_creationError(), (_7_dt__update__tmp_h2).Dtor_serverTools(), (_7_dt__update__tmp_h2).Dtor_changesets(), (_7_dt__update__tmp_h2).Dtor_canvases(), (_7_dt__update__tmp_h2).Dtor_openCanvases(), (_7_dt__update__tmp_h2).Dtor_defaultChat(), (_7_dt__update__tmp_h2).Dtor_activeClients(), (_7_dt__update__tmp_h2).Dtor_chats(), (_7_dt__update__tmp_h2).Dtor_customizations(), (_7_dt__update__tmp_h2).Dtor_inputNeeded())
						}(_pat_let117_0)
					}(_6_t)
				}(_pat_let116_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_ServerToolsChanged() {
			var _9_t m_ConfluxCodec.Json = _source0.Get_().(SessionAction_ServerToolsChanged).Tools
			_ = _9_t
			return _dafny.TupleOf(func(_pat_let118_0 SessionState) SessionState {
				return func(_10_dt__update__tmp_h3 SessionState) SessionState {
					return func(_pat_let119_0 m_AhpSkeleton.Option) SessionState {
						return func(_11_dt__update_hserverTools_h0 m_AhpSkeleton.Option) SessionState {
							return Companion_SessionState_.Create_SessionState_((_10_dt__update__tmp_h3).Dtor_provider(), (_10_dt__update__tmp_h3).Dtor_title(), (_10_dt__update__tmp_h3).Dtor_status(), (_10_dt__update__tmp_h3).Dtor_lifecycle(), (_10_dt__update__tmp_h3).Dtor_activity(), (_10_dt__update__tmp_h3).Dtor_config(), (_10_dt__update__tmp_h3).Dtor_meta(), (_10_dt__update__tmp_h3).Dtor_creationError(), _11_dt__update_hserverTools_h0, (_10_dt__update__tmp_h3).Dtor_changesets(), (_10_dt__update__tmp_h3).Dtor_canvases(), (_10_dt__update__tmp_h3).Dtor_openCanvases(), (_10_dt__update__tmp_h3).Dtor_defaultChat(), (_10_dt__update__tmp_h3).Dtor_activeClients(), (_10_dt__update__tmp_h3).Dtor_chats(), (_10_dt__update__tmp_h3).Dtor_customizations(), (_10_dt__update__tmp_h3).Dtor_inputNeeded())
						}(_pat_let119_0)
					}((func() m_AhpSkeleton.Option {
						if (_9_t).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_()) {
							return m_AhpSkeleton.Companion_Option_.Create_None_()
						}
						return m_AhpSkeleton.Companion_Option_.Create_Some_(_9_t)
					})())
				}(_pat_let118_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_MetaChanged() {
			var _12_m m_ConfluxCodec.Json = _source0.Get_().(SessionAction_MetaChanged).Meta
			_ = _12_m
			return _dafny.TupleOf(func(_pat_let120_0 SessionState) SessionState {
				return func(_13_dt__update__tmp_h4 SessionState) SessionState {
					return func(_pat_let121_0 m_AhpSkeleton.Option) SessionState {
						return func(_14_dt__update_hmeta_h0 m_AhpSkeleton.Option) SessionState {
							return Companion_SessionState_.Create_SessionState_((_13_dt__update__tmp_h4).Dtor_provider(), (_13_dt__update__tmp_h4).Dtor_title(), (_13_dt__update__tmp_h4).Dtor_status(), (_13_dt__update__tmp_h4).Dtor_lifecycle(), (_13_dt__update__tmp_h4).Dtor_activity(), (_13_dt__update__tmp_h4).Dtor_config(), _14_dt__update_hmeta_h0, (_13_dt__update__tmp_h4).Dtor_creationError(), (_13_dt__update__tmp_h4).Dtor_serverTools(), (_13_dt__update__tmp_h4).Dtor_changesets(), (_13_dt__update__tmp_h4).Dtor_canvases(), (_13_dt__update__tmp_h4).Dtor_openCanvases(), (_13_dt__update__tmp_h4).Dtor_defaultChat(), (_13_dt__update__tmp_h4).Dtor_activeClients(), (_13_dt__update__tmp_h4).Dtor_chats(), (_13_dt__update__tmp_h4).Dtor_customizations(), (_13_dt__update__tmp_h4).Dtor_inputNeeded())
						}(_pat_let121_0)
					}((func() m_AhpSkeleton.Option {
						if (_12_m).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_()) {
							return m_AhpSkeleton.Companion_Option_.Create_None_()
						}
						return m_AhpSkeleton.Companion_Option_.Create_Some_(_12_m)
					})())
				}(_pat_let120_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_IsReadChanged() {
			var _15_v bool = _source0.Get_().(SessionAction_IsReadChanged).IsRead
			_ = _15_v
			return _dafny.TupleOf(func(_pat_let122_0 SessionState) SessionState {
				return func(_16_dt__update__tmp_h5 SessionState) SessionState {
					return func(_pat_let123_0 uint32) SessionState {
						return func(_17_dt__update_hstatus_h0 uint32) SessionState {
							return Companion_SessionState_.Create_SessionState_((_16_dt__update__tmp_h5).Dtor_provider(), (_16_dt__update__tmp_h5).Dtor_title(), _17_dt__update_hstatus_h0, (_16_dt__update__tmp_h5).Dtor_lifecycle(), (_16_dt__update__tmp_h5).Dtor_activity(), (_16_dt__update__tmp_h5).Dtor_config(), (_16_dt__update__tmp_h5).Dtor_meta(), (_16_dt__update__tmp_h5).Dtor_creationError(), (_16_dt__update__tmp_h5).Dtor_serverTools(), (_16_dt__update__tmp_h5).Dtor_changesets(), (_16_dt__update__tmp_h5).Dtor_canvases(), (_16_dt__update__tmp_h5).Dtor_openCanvases(), (_16_dt__update__tmp_h5).Dtor_defaultChat(), (_16_dt__update__tmp_h5).Dtor_activeClients(), (_16_dt__update__tmp_h5).Dtor_chats(), (_16_dt__update__tmp_h5).Dtor_customizations(), (_16_dt__update__tmp_h5).Dtor_inputNeeded())
						}(_pat_let123_0)
					}((func() uint32 {
						if _15_v {
							return Companion_Default___.SetBit((_pat_let_tv0).Dtor_status(), Companion_Default___.B__READ())
						}
						return Companion_Default___.ClearBit((_pat_let_tv1).Dtor_status(), Companion_Default___.B__READ())
					})())
				}(_pat_let122_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_IsArchivedChanged() {
			var _18_v bool = _source0.Get_().(SessionAction_IsArchivedChanged).IsArchived
			_ = _18_v
			return _dafny.TupleOf(func(_pat_let124_0 SessionState) SessionState {
				return func(_19_dt__update__tmp_h6 SessionState) SessionState {
					return func(_pat_let125_0 uint32) SessionState {
						return func(_20_dt__update_hstatus_h1 uint32) SessionState {
							return Companion_SessionState_.Create_SessionState_((_19_dt__update__tmp_h6).Dtor_provider(), (_19_dt__update__tmp_h6).Dtor_title(), _20_dt__update_hstatus_h1, (_19_dt__update__tmp_h6).Dtor_lifecycle(), (_19_dt__update__tmp_h6).Dtor_activity(), (_19_dt__update__tmp_h6).Dtor_config(), (_19_dt__update__tmp_h6).Dtor_meta(), (_19_dt__update__tmp_h6).Dtor_creationError(), (_19_dt__update__tmp_h6).Dtor_serverTools(), (_19_dt__update__tmp_h6).Dtor_changesets(), (_19_dt__update__tmp_h6).Dtor_canvases(), (_19_dt__update__tmp_h6).Dtor_openCanvases(), (_19_dt__update__tmp_h6).Dtor_defaultChat(), (_19_dt__update__tmp_h6).Dtor_activeClients(), (_19_dt__update__tmp_h6).Dtor_chats(), (_19_dt__update__tmp_h6).Dtor_customizations(), (_19_dt__update__tmp_h6).Dtor_inputNeeded())
						}(_pat_let125_0)
					}((func() uint32 {
						if _18_v {
							return Companion_Default___.SetBit((_pat_let_tv2).Dtor_status(), Companion_Default___.B__ARCH())
						}
						return Companion_Default___.ClearBit((_pat_let_tv3).Dtor_status(), Companion_Default___.B__ARCH())
					})())
				}(_pat_let124_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_ActivityChanged() {
			var _21_ac m_AhpSkeleton.Option = _source0.Get_().(SessionAction_ActivityChanged).Activity
			_ = _21_ac
			return _dafny.TupleOf(func(_pat_let126_0 SessionState) SessionState {
				return func(_22_dt__update__tmp_h7 SessionState) SessionState {
					return func(_pat_let127_0 m_AhpSkeleton.Option) SessionState {
						return func(_23_dt__update_hactivity_h0 m_AhpSkeleton.Option) SessionState {
							return Companion_SessionState_.Create_SessionState_((_22_dt__update__tmp_h7).Dtor_provider(), (_22_dt__update__tmp_h7).Dtor_title(), (_22_dt__update__tmp_h7).Dtor_status(), (_22_dt__update__tmp_h7).Dtor_lifecycle(), _23_dt__update_hactivity_h0, (_22_dt__update__tmp_h7).Dtor_config(), (_22_dt__update__tmp_h7).Dtor_meta(), (_22_dt__update__tmp_h7).Dtor_creationError(), (_22_dt__update__tmp_h7).Dtor_serverTools(), (_22_dt__update__tmp_h7).Dtor_changesets(), (_22_dt__update__tmp_h7).Dtor_canvases(), (_22_dt__update__tmp_h7).Dtor_openCanvases(), (_22_dt__update__tmp_h7).Dtor_defaultChat(), (_22_dt__update__tmp_h7).Dtor_activeClients(), (_22_dt__update__tmp_h7).Dtor_chats(), (_22_dt__update__tmp_h7).Dtor_customizations(), (_22_dt__update__tmp_h7).Dtor_inputNeeded())
						}(_pat_let127_0)
					}(_21_ac)
				}(_pat_let126_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_ConfigChanged() {
			var _24_cfg _dafny.Map = _source0.Get_().(SessionAction_ConfigChanged).Config
			_ = _24_cfg
			var _25_repl bool = _source0.Get_().(SessionAction_ConfigChanged).Replace
			_ = _25_repl
			var _source1 m_AhpSkeleton.Option = (s).Dtor_config()
			_ = _source1
			{
				if _source1.Is_None() {
					return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
				}
			}
			{
				var _26_c SConfig = _source1.Get_().(m_AhpSkeleton.Option_Some).Value.(SConfig)
				_ = _26_c
				return _dafny.TupleOf(func(_pat_let128_0 SessionState) SessionState {
					return func(_27_dt__update__tmp_h8 SessionState) SessionState {
						return func(_pat_let129_0 m_AhpSkeleton.Option) SessionState {
							return func(_28_dt__update_hconfig_h0 m_AhpSkeleton.Option) SessionState {
								return Companion_SessionState_.Create_SessionState_((_27_dt__update__tmp_h8).Dtor_provider(), (_27_dt__update__tmp_h8).Dtor_title(), (_27_dt__update__tmp_h8).Dtor_status(), (_27_dt__update__tmp_h8).Dtor_lifecycle(), (_27_dt__update__tmp_h8).Dtor_activity(), _28_dt__update_hconfig_h0, (_27_dt__update__tmp_h8).Dtor_meta(), (_27_dt__update__tmp_h8).Dtor_creationError(), (_27_dt__update__tmp_h8).Dtor_serverTools(), (_27_dt__update__tmp_h8).Dtor_changesets(), (_27_dt__update__tmp_h8).Dtor_canvases(), (_27_dt__update__tmp_h8).Dtor_openCanvases(), (_27_dt__update__tmp_h8).Dtor_defaultChat(), (_27_dt__update__tmp_h8).Dtor_activeClients(), (_27_dt__update__tmp_h8).Dtor_chats(), (_27_dt__update__tmp_h8).Dtor_customizations(), (_27_dt__update__tmp_h8).Dtor_inputNeeded())
							}(_pat_let129_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_SConfig_.Create_SConfig_((_26_c).Dtor_schema(), (func() _dafny.Map {
							if _25_repl {
								return _24_cfg
							}
							return ((_26_c).Dtor_values()).Merge(_24_cfg)
						})())))
					}(_pat_let128_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			}
		}
	}
	{
		if _source0.Is_CustomizationsChanged() {
			var _29_cs _dafny.Sequence = _source0.Get_().(SessionAction_CustomizationsChanged).Customizations
			_ = _29_cs
			return _dafny.TupleOf(func(_pat_let130_0 SessionState) SessionState {
				return func(_30_dt__update__tmp_h9 SessionState) SessionState {
					return func(_pat_let131_0 _dafny.Sequence) SessionState {
						return func(_31_dt__update_hcustomizations_h0 _dafny.Sequence) SessionState {
							return Companion_SessionState_.Create_SessionState_((_30_dt__update__tmp_h9).Dtor_provider(), (_30_dt__update__tmp_h9).Dtor_title(), (_30_dt__update__tmp_h9).Dtor_status(), (_30_dt__update__tmp_h9).Dtor_lifecycle(), (_30_dt__update__tmp_h9).Dtor_activity(), (_30_dt__update__tmp_h9).Dtor_config(), (_30_dt__update__tmp_h9).Dtor_meta(), (_30_dt__update__tmp_h9).Dtor_creationError(), (_30_dt__update__tmp_h9).Dtor_serverTools(), (_30_dt__update__tmp_h9).Dtor_changesets(), (_30_dt__update__tmp_h9).Dtor_canvases(), (_30_dt__update__tmp_h9).Dtor_openCanvases(), (_30_dt__update__tmp_h9).Dtor_defaultChat(), (_30_dt__update__tmp_h9).Dtor_activeClients(), (_30_dt__update__tmp_h9).Dtor_chats(), _31_dt__update_hcustomizations_h0, (_30_dt__update__tmp_h9).Dtor_inputNeeded())
						}(_pat_let131_0)
					}(_29_cs)
				}(_pat_let130_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_CustomizationToggled() {
			var _32_id _dafny.Sequence = _source0.Get_().(SessionAction_CustomizationToggled).Id
			_ = _32_id
			var _33_en bool = _source0.Get_().(SessionAction_CustomizationToggled).Enabled
			_ = _33_en
			return _dafny.TupleOf(func(_pat_let132_0 SessionState) SessionState {
				return func(_34_dt__update__tmp_h10 SessionState) SessionState {
					return func(_pat_let133_0 _dafny.Sequence) SessionState {
						return func(_35_dt__update_hcustomizations_h1 _dafny.Sequence) SessionState {
							return Companion_SessionState_.Create_SessionState_((_34_dt__update__tmp_h10).Dtor_provider(), (_34_dt__update__tmp_h10).Dtor_title(), (_34_dt__update__tmp_h10).Dtor_status(), (_34_dt__update__tmp_h10).Dtor_lifecycle(), (_34_dt__update__tmp_h10).Dtor_activity(), (_34_dt__update__tmp_h10).Dtor_config(), (_34_dt__update__tmp_h10).Dtor_meta(), (_34_dt__update__tmp_h10).Dtor_creationError(), (_34_dt__update__tmp_h10).Dtor_serverTools(), (_34_dt__update__tmp_h10).Dtor_changesets(), (_34_dt__update__tmp_h10).Dtor_canvases(), (_34_dt__update__tmp_h10).Dtor_openCanvases(), (_34_dt__update__tmp_h10).Dtor_defaultChat(), (_34_dt__update__tmp_h10).Dtor_activeClients(), (_34_dt__update__tmp_h10).Dtor_chats(), _35_dt__update_hcustomizations_h1, (_34_dt__update__tmp_h10).Dtor_inputNeeded())
						}(_pat_let133_0)
					}(Companion_Default___.UpsertCust((_pat_let_tv4).Dtor_customizations(), _32_id, _33_en))
				}(_pat_let132_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_CustomizationRemoved() {
			var _36_id _dafny.Sequence = _source0.Get_().(SessionAction_CustomizationRemoved).Id
			_ = _36_id
			return _dafny.TupleOf(func(_pat_let134_0 SessionState) SessionState {
				return func(_37_dt__update__tmp_h11 SessionState) SessionState {
					return func(_pat_let135_0 _dafny.Sequence) SessionState {
						return func(_38_dt__update_hcustomizations_h2 _dafny.Sequence) SessionState {
							return Companion_SessionState_.Create_SessionState_((_37_dt__update__tmp_h11).Dtor_provider(), (_37_dt__update__tmp_h11).Dtor_title(), (_37_dt__update__tmp_h11).Dtor_status(), (_37_dt__update__tmp_h11).Dtor_lifecycle(), (_37_dt__update__tmp_h11).Dtor_activity(), (_37_dt__update__tmp_h11).Dtor_config(), (_37_dt__update__tmp_h11).Dtor_meta(), (_37_dt__update__tmp_h11).Dtor_creationError(), (_37_dt__update__tmp_h11).Dtor_serverTools(), (_37_dt__update__tmp_h11).Dtor_changesets(), (_37_dt__update__tmp_h11).Dtor_canvases(), (_37_dt__update__tmp_h11).Dtor_openCanvases(), (_37_dt__update__tmp_h11).Dtor_defaultChat(), (_37_dt__update__tmp_h11).Dtor_activeClients(), (_37_dt__update__tmp_h11).Dtor_chats(), _38_dt__update_hcustomizations_h2, (_37_dt__update__tmp_h11).Dtor_inputNeeded())
						}(_pat_let135_0)
					}(Companion_Default___.RemoveCust((_pat_let_tv5).Dtor_customizations(), _36_id))
				}(_pat_let134_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_DefaultChatChanged() {
			var _39_c m_AhpSkeleton.Option = _source0.Get_().(SessionAction_DefaultChatChanged).Chat
			_ = _39_c
			return _dafny.TupleOf(func(_pat_let136_0 SessionState) SessionState {
				return func(_40_dt__update__tmp_h12 SessionState) SessionState {
					return func(_pat_let137_0 m_AhpSkeleton.Option) SessionState {
						return func(_41_dt__update_hdefaultChat_h0 m_AhpSkeleton.Option) SessionState {
							return Companion_SessionState_.Create_SessionState_((_40_dt__update__tmp_h12).Dtor_provider(), (_40_dt__update__tmp_h12).Dtor_title(), (_40_dt__update__tmp_h12).Dtor_status(), (_40_dt__update__tmp_h12).Dtor_lifecycle(), (_40_dt__update__tmp_h12).Dtor_activity(), (_40_dt__update__tmp_h12).Dtor_config(), (_40_dt__update__tmp_h12).Dtor_meta(), (_40_dt__update__tmp_h12).Dtor_creationError(), (_40_dt__update__tmp_h12).Dtor_serverTools(), (_40_dt__update__tmp_h12).Dtor_changesets(), (_40_dt__update__tmp_h12).Dtor_canvases(), (_40_dt__update__tmp_h12).Dtor_openCanvases(), _41_dt__update_hdefaultChat_h0, (_40_dt__update__tmp_h12).Dtor_activeClients(), (_40_dt__update__tmp_h12).Dtor_chats(), (_40_dt__update__tmp_h12).Dtor_customizations(), (_40_dt__update__tmp_h12).Dtor_inputNeeded())
						}(_pat_let137_0)
					}(_39_c)
				}(_pat_let136_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_ChatAdded() {
			var _42_c Chat = _source0.Get_().(SessionAction_ChatAdded).Summary
			_ = _42_c
			return _dafny.TupleOf(func(_pat_let138_0 SessionState) SessionState {
				return func(_43_dt__update__tmp_h13 SessionState) SessionState {
					return func(_pat_let139_0 _dafny.Sequence) SessionState {
						return func(_44_dt__update_hchats_h0 _dafny.Sequence) SessionState {
							return Companion_SessionState_.Create_SessionState_((_43_dt__update__tmp_h13).Dtor_provider(), (_43_dt__update__tmp_h13).Dtor_title(), (_43_dt__update__tmp_h13).Dtor_status(), (_43_dt__update__tmp_h13).Dtor_lifecycle(), (_43_dt__update__tmp_h13).Dtor_activity(), (_43_dt__update__tmp_h13).Dtor_config(), (_43_dt__update__tmp_h13).Dtor_meta(), (_43_dt__update__tmp_h13).Dtor_creationError(), (_43_dt__update__tmp_h13).Dtor_serverTools(), (_43_dt__update__tmp_h13).Dtor_changesets(), (_43_dt__update__tmp_h13).Dtor_canvases(), (_43_dt__update__tmp_h13).Dtor_openCanvases(), (_43_dt__update__tmp_h13).Dtor_defaultChat(), (_43_dt__update__tmp_h13).Dtor_activeClients(), _44_dt__update_hchats_h0, (_43_dt__update__tmp_h13).Dtor_customizations(), (_43_dt__update__tmp_h13).Dtor_inputNeeded())
						}(_pat_let139_0)
					}(Companion_Default___.UpsertChat((_pat_let_tv6).Dtor_chats(), _42_c))
				}(_pat_let138_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_ChatRemoved() {
			var _45_r _dafny.Sequence = _source0.Get_().(SessionAction_ChatRemoved).Resource
			_ = _45_r
			return _dafny.TupleOf(func(_pat_let140_0 SessionState) SessionState {
				return func(_46_dt__update__tmp_h14 SessionState) SessionState {
					return func(_pat_let141_0 m_AhpSkeleton.Option) SessionState {
						return func(_47_dt__update_hdefaultChat_h1 m_AhpSkeleton.Option) SessionState {
							return func(_pat_let142_0 _dafny.Sequence) SessionState {
								return func(_48_dt__update_hchats_h1 _dafny.Sequence) SessionState {
									return Companion_SessionState_.Create_SessionState_((_46_dt__update__tmp_h14).Dtor_provider(), (_46_dt__update__tmp_h14).Dtor_title(), (_46_dt__update__tmp_h14).Dtor_status(), (_46_dt__update__tmp_h14).Dtor_lifecycle(), (_46_dt__update__tmp_h14).Dtor_activity(), (_46_dt__update__tmp_h14).Dtor_config(), (_46_dt__update__tmp_h14).Dtor_meta(), (_46_dt__update__tmp_h14).Dtor_creationError(), (_46_dt__update__tmp_h14).Dtor_serverTools(), (_46_dt__update__tmp_h14).Dtor_changesets(), (_46_dt__update__tmp_h14).Dtor_canvases(), (_46_dt__update__tmp_h14).Dtor_openCanvases(), _47_dt__update_hdefaultChat_h1, (_46_dt__update__tmp_h14).Dtor_activeClients(), _48_dt__update_hchats_h1, (_46_dt__update__tmp_h14).Dtor_customizations(), (_46_dt__update__tmp_h14).Dtor_inputNeeded())
								}(_pat_let142_0)
							}(Companion_Default___.RemoveChat((_pat_let_tv9).Dtor_chats(), _45_r))
						}(_pat_let141_0)
					}((func() m_AhpSkeleton.Option {
						if ((_pat_let_tv7).Dtor_defaultChat()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_45_r)) {
							return m_AhpSkeleton.Companion_Option_.Create_None_()
						}
						return (_pat_let_tv8).Dtor_defaultChat()
					})())
				}(_pat_let140_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_ChatUpdated() {
			var _49_r _dafny.Sequence = _source0.Get_().(SessionAction_ChatUpdated).Resource
			_ = _49_r
			var _50_st m_AhpSkeleton.Option = _source0.Get_().(SessionAction_ChatUpdated).Status
			_ = _50_st
			var _51_ac m_AhpSkeleton.Option = _source0.Get_().(SessionAction_ChatUpdated).Activity
			_ = _51_ac
			var _52_md m_AhpSkeleton.Option = _source0.Get_().(SessionAction_ChatUpdated).ModifiedAt
			_ = _52_md
			return _dafny.TupleOf(func(_pat_let143_0 SessionState) SessionState {
				return func(_53_dt__update__tmp_h15 SessionState) SessionState {
					return func(_pat_let144_0 _dafny.Sequence) SessionState {
						return func(_54_dt__update_hchats_h2 _dafny.Sequence) SessionState {
							return Companion_SessionState_.Create_SessionState_((_53_dt__update__tmp_h15).Dtor_provider(), (_53_dt__update__tmp_h15).Dtor_title(), (_53_dt__update__tmp_h15).Dtor_status(), (_53_dt__update__tmp_h15).Dtor_lifecycle(), (_53_dt__update__tmp_h15).Dtor_activity(), (_53_dt__update__tmp_h15).Dtor_config(), (_53_dt__update__tmp_h15).Dtor_meta(), (_53_dt__update__tmp_h15).Dtor_creationError(), (_53_dt__update__tmp_h15).Dtor_serverTools(), (_53_dt__update__tmp_h15).Dtor_changesets(), (_53_dt__update__tmp_h15).Dtor_canvases(), (_53_dt__update__tmp_h15).Dtor_openCanvases(), (_53_dt__update__tmp_h15).Dtor_defaultChat(), (_53_dt__update__tmp_h15).Dtor_activeClients(), _54_dt__update_hchats_h2, (_53_dt__update__tmp_h15).Dtor_customizations(), (_53_dt__update__tmp_h15).Dtor_inputNeeded())
						}(_pat_let144_0)
					}(Companion_Default___.UpdChat((_pat_let_tv10).Dtor_chats(), _49_r, _50_st, _51_ac, _52_md))
				}(_pat_let143_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_ChangesetsChanged() {
			var _55_cs m_AhpSkeleton.Option = _source0.Get_().(SessionAction_ChangesetsChanged).Changesets
			_ = _55_cs
			return _dafny.TupleOf(func(_pat_let145_0 SessionState) SessionState {
				return func(_56_dt__update__tmp_h16 SessionState) SessionState {
					return func(_pat_let146_0 m_AhpSkeleton.Option) SessionState {
						return func(_57_dt__update_hchangesets_h0 m_AhpSkeleton.Option) SessionState {
							return Companion_SessionState_.Create_SessionState_((_56_dt__update__tmp_h16).Dtor_provider(), (_56_dt__update__tmp_h16).Dtor_title(), (_56_dt__update__tmp_h16).Dtor_status(), (_56_dt__update__tmp_h16).Dtor_lifecycle(), (_56_dt__update__tmp_h16).Dtor_activity(), (_56_dt__update__tmp_h16).Dtor_config(), (_56_dt__update__tmp_h16).Dtor_meta(), (_56_dt__update__tmp_h16).Dtor_creationError(), (_56_dt__update__tmp_h16).Dtor_serverTools(), _57_dt__update_hchangesets_h0, (_56_dt__update__tmp_h16).Dtor_canvases(), (_56_dt__update__tmp_h16).Dtor_openCanvases(), (_56_dt__update__tmp_h16).Dtor_defaultChat(), (_56_dt__update__tmp_h16).Dtor_activeClients(), (_56_dt__update__tmp_h16).Dtor_chats(), (_56_dt__update__tmp_h16).Dtor_customizations(), (_56_dt__update__tmp_h16).Dtor_inputNeeded())
						}(_pat_let146_0)
					}(_55_cs)
				}(_pat_let145_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_CanvasesChanged() {
			var _58_cs m_ConfluxCodec.Json = _source0.Get_().(SessionAction_CanvasesChanged).Canvases
			_ = _58_cs
			return _dafny.TupleOf(func(_pat_let147_0 SessionState) SessionState {
				return func(_59_dt__update__tmp_h17 SessionState) SessionState {
					return func(_pat_let148_0 m_AhpSkeleton.Option) SessionState {
						return func(_60_dt__update_hcanvases_h0 m_AhpSkeleton.Option) SessionState {
							return Companion_SessionState_.Create_SessionState_((_59_dt__update__tmp_h17).Dtor_provider(), (_59_dt__update__tmp_h17).Dtor_title(), (_59_dt__update__tmp_h17).Dtor_status(), (_59_dt__update__tmp_h17).Dtor_lifecycle(), (_59_dt__update__tmp_h17).Dtor_activity(), (_59_dt__update__tmp_h17).Dtor_config(), (_59_dt__update__tmp_h17).Dtor_meta(), (_59_dt__update__tmp_h17).Dtor_creationError(), (_59_dt__update__tmp_h17).Dtor_serverTools(), (_59_dt__update__tmp_h17).Dtor_changesets(), _60_dt__update_hcanvases_h0, (_59_dt__update__tmp_h17).Dtor_openCanvases(), (_59_dt__update__tmp_h17).Dtor_defaultChat(), (_59_dt__update__tmp_h17).Dtor_activeClients(), (_59_dt__update__tmp_h17).Dtor_chats(), (_59_dt__update__tmp_h17).Dtor_customizations(), (_59_dt__update__tmp_h17).Dtor_inputNeeded())
						}(_pat_let148_0)
					}((func() m_AhpSkeleton.Option {
						if (_58_cs).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_()) {
							return m_AhpSkeleton.Companion_Option_.Create_None_()
						}
						return m_AhpSkeleton.Companion_Option_.Create_Some_(_58_cs)
					})())
				}(_pat_let147_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_OpenCanvasesChanged() {
			var _61_cs m_ConfluxCodec.Json = _source0.Get_().(SessionAction_OpenCanvasesChanged).OpenCanvases
			_ = _61_cs
			return _dafny.TupleOf(func(_pat_let149_0 SessionState) SessionState {
				return func(_62_dt__update__tmp_h18 SessionState) SessionState {
					return func(_pat_let150_0 m_AhpSkeleton.Option) SessionState {
						return func(_63_dt__update_hopenCanvases_h0 m_AhpSkeleton.Option) SessionState {
							return Companion_SessionState_.Create_SessionState_((_62_dt__update__tmp_h18).Dtor_provider(), (_62_dt__update__tmp_h18).Dtor_title(), (_62_dt__update__tmp_h18).Dtor_status(), (_62_dt__update__tmp_h18).Dtor_lifecycle(), (_62_dt__update__tmp_h18).Dtor_activity(), (_62_dt__update__tmp_h18).Dtor_config(), (_62_dt__update__tmp_h18).Dtor_meta(), (_62_dt__update__tmp_h18).Dtor_creationError(), (_62_dt__update__tmp_h18).Dtor_serverTools(), (_62_dt__update__tmp_h18).Dtor_changesets(), (_62_dt__update__tmp_h18).Dtor_canvases(), _63_dt__update_hopenCanvases_h0, (_62_dt__update__tmp_h18).Dtor_defaultChat(), (_62_dt__update__tmp_h18).Dtor_activeClients(), (_62_dt__update__tmp_h18).Dtor_chats(), (_62_dt__update__tmp_h18).Dtor_customizations(), (_62_dt__update__tmp_h18).Dtor_inputNeeded())
						}(_pat_let150_0)
					}((func() m_AhpSkeleton.Option {
						if (_61_cs).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_()) {
							return m_AhpSkeleton.Companion_Option_.Create_None_()
						}
						return m_AhpSkeleton.Companion_Option_.Create_Some_(_61_cs)
					})())
				}(_pat_let149_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_ActiveClientSet() {
			var _64_c Client = _source0.Get_().(SessionAction_ActiveClientSet).Client
			_ = _64_c
			return _dafny.TupleOf(func(_pat_let151_0 SessionState) SessionState {
				return func(_65_dt__update__tmp_h19 SessionState) SessionState {
					return func(_pat_let152_0 _dafny.Sequence) SessionState {
						return func(_66_dt__update_hactiveClients_h0 _dafny.Sequence) SessionState {
							return Companion_SessionState_.Create_SessionState_((_65_dt__update__tmp_h19).Dtor_provider(), (_65_dt__update__tmp_h19).Dtor_title(), (_65_dt__update__tmp_h19).Dtor_status(), (_65_dt__update__tmp_h19).Dtor_lifecycle(), (_65_dt__update__tmp_h19).Dtor_activity(), (_65_dt__update__tmp_h19).Dtor_config(), (_65_dt__update__tmp_h19).Dtor_meta(), (_65_dt__update__tmp_h19).Dtor_creationError(), (_65_dt__update__tmp_h19).Dtor_serverTools(), (_65_dt__update__tmp_h19).Dtor_changesets(), (_65_dt__update__tmp_h19).Dtor_canvases(), (_65_dt__update__tmp_h19).Dtor_openCanvases(), (_65_dt__update__tmp_h19).Dtor_defaultChat(), _66_dt__update_hactiveClients_h0, (_65_dt__update__tmp_h19).Dtor_chats(), (_65_dt__update__tmp_h19).Dtor_customizations(), (_65_dt__update__tmp_h19).Dtor_inputNeeded())
						}(_pat_let152_0)
					}(Companion_Default___.UpsertClient((_pat_let_tv11).Dtor_activeClients(), _64_c))
				}(_pat_let151_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_ActiveClientRemoved() {
			var _67_id _dafny.Sequence = _source0.Get_().(SessionAction_ActiveClientRemoved).ClientId
			_ = _67_id
			return _dafny.TupleOf(func(_pat_let153_0 SessionState) SessionState {
				return func(_68_dt__update__tmp_h20 SessionState) SessionState {
					return func(_pat_let154_0 _dafny.Sequence) SessionState {
						return func(_69_dt__update_hactiveClients_h1 _dafny.Sequence) SessionState {
							return Companion_SessionState_.Create_SessionState_((_68_dt__update__tmp_h20).Dtor_provider(), (_68_dt__update__tmp_h20).Dtor_title(), (_68_dt__update__tmp_h20).Dtor_status(), (_68_dt__update__tmp_h20).Dtor_lifecycle(), (_68_dt__update__tmp_h20).Dtor_activity(), (_68_dt__update__tmp_h20).Dtor_config(), (_68_dt__update__tmp_h20).Dtor_meta(), (_68_dt__update__tmp_h20).Dtor_creationError(), (_68_dt__update__tmp_h20).Dtor_serverTools(), (_68_dt__update__tmp_h20).Dtor_changesets(), (_68_dt__update__tmp_h20).Dtor_canvases(), (_68_dt__update__tmp_h20).Dtor_openCanvases(), (_68_dt__update__tmp_h20).Dtor_defaultChat(), _69_dt__update_hactiveClients_h1, (_68_dt__update__tmp_h20).Dtor_chats(), (_68_dt__update__tmp_h20).Dtor_customizations(), (_68_dt__update__tmp_h20).Dtor_inputNeeded())
						}(_pat_let154_0)
					}(Companion_Default___.RemoveClient((_pat_let_tv12).Dtor_activeClients(), _67_id))
				}(_pat_let153_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_InputNeededSet() {
			var _70_r InputReq = _source0.Get_().(SessionAction_InputNeededSet).Req
			_ = _70_r
			return _dafny.TupleOf(func(_pat_let155_0 SessionState) SessionState {
				return func(_71_dt__update__tmp_h21 SessionState) SessionState {
					return func(_pat_let156_0 uint32) SessionState {
						return func(_72_dt__update_hstatus_h2 uint32) SessionState {
							return func(_pat_let157_0 _dafny.Sequence) SessionState {
								return func(_73_dt__update_hinputNeeded_h0 _dafny.Sequence) SessionState {
									return Companion_SessionState_.Create_SessionState_((_71_dt__update__tmp_h21).Dtor_provider(), (_71_dt__update__tmp_h21).Dtor_title(), _72_dt__update_hstatus_h2, (_71_dt__update__tmp_h21).Dtor_lifecycle(), (_71_dt__update__tmp_h21).Dtor_activity(), (_71_dt__update__tmp_h21).Dtor_config(), (_71_dt__update__tmp_h21).Dtor_meta(), (_71_dt__update__tmp_h21).Dtor_creationError(), (_71_dt__update__tmp_h21).Dtor_serverTools(), (_71_dt__update__tmp_h21).Dtor_changesets(), (_71_dt__update__tmp_h21).Dtor_canvases(), (_71_dt__update__tmp_h21).Dtor_openCanvases(), (_71_dt__update__tmp_h21).Dtor_defaultChat(), (_71_dt__update__tmp_h21).Dtor_activeClients(), (_71_dt__update__tmp_h21).Dtor_chats(), (_71_dt__update__tmp_h21).Dtor_customizations(), _73_dt__update_hinputNeeded_h0)
								}(_pat_let157_0)
							}(Companion_Default___.UpsertInput((_pat_let_tv14).Dtor_inputNeeded(), _70_r))
						}(_pat_let156_0)
					}(Companion_Default___.SetBit((_pat_let_tv13).Dtor_status(), Companion_Default___.B__INPUT()))
				}(_pat_let155_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_InputNeededRemoved() {
			var _74_id _dafny.Sequence = _source0.Get_().(SessionAction_InputNeededRemoved).Id
			_ = _74_id
			var _75_rem _dafny.Sequence = Companion_Default___.RemoveInput((s).Dtor_inputNeeded(), _74_id)
			_ = _75_rem
			return _dafny.TupleOf(func(_pat_let158_0 SessionState) SessionState {
				return func(_76_dt__update__tmp_h22 SessionState) SessionState {
					return func(_pat_let159_0 uint32) SessionState {
						return func(_77_dt__update_hstatus_h3 uint32) SessionState {
							return func(_pat_let160_0 _dafny.Sequence) SessionState {
								return func(_78_dt__update_hinputNeeded_h1 _dafny.Sequence) SessionState {
									return Companion_SessionState_.Create_SessionState_((_76_dt__update__tmp_h22).Dtor_provider(), (_76_dt__update__tmp_h22).Dtor_title(), _77_dt__update_hstatus_h3, (_76_dt__update__tmp_h22).Dtor_lifecycle(), (_76_dt__update__tmp_h22).Dtor_activity(), (_76_dt__update__tmp_h22).Dtor_config(), (_76_dt__update__tmp_h22).Dtor_meta(), (_76_dt__update__tmp_h22).Dtor_creationError(), (_76_dt__update__tmp_h22).Dtor_serverTools(), (_76_dt__update__tmp_h22).Dtor_changesets(), (_76_dt__update__tmp_h22).Dtor_canvases(), (_76_dt__update__tmp_h22).Dtor_openCanvases(), (_76_dt__update__tmp_h22).Dtor_defaultChat(), (_76_dt__update__tmp_h22).Dtor_activeClients(), (_76_dt__update__tmp_h22).Dtor_chats(), (_76_dt__update__tmp_h22).Dtor_customizations(), _78_dt__update_hinputNeeded_h1)
								}(_pat_let160_0)
							}(_75_rem)
						}(_pat_let159_0)
					}((func() uint32 {
						if (_dafny.IntOfUint32((_75_rem).Cardinality())).Sign() == 0 {
							return Companion_Default___.ClearBit((_pat_let_tv15).Dtor_status(), Companion_Default___.B__INPUT())
						}
						return (_pat_let_tv16).Dtor_status()
					})())
				}(_pat_let158_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_CustomizationUpdated() {
			var _79_c Cust = _source0.Get_().(SessionAction_CustomizationUpdated).Customization
			_ = _79_c
			return _dafny.TupleOf(func(_pat_let161_0 SessionState) SessionState {
				return func(_80_dt__update__tmp_h23 SessionState) SessionState {
					return func(_pat_let162_0 _dafny.Sequence) SessionState {
						return func(_81_dt__update_hcustomizations_h3 _dafny.Sequence) SessionState {
							return Companion_SessionState_.Create_SessionState_((_80_dt__update__tmp_h23).Dtor_provider(), (_80_dt__update__tmp_h23).Dtor_title(), (_80_dt__update__tmp_h23).Dtor_status(), (_80_dt__update__tmp_h23).Dtor_lifecycle(), (_80_dt__update__tmp_h23).Dtor_activity(), (_80_dt__update__tmp_h23).Dtor_config(), (_80_dt__update__tmp_h23).Dtor_meta(), (_80_dt__update__tmp_h23).Dtor_creationError(), (_80_dt__update__tmp_h23).Dtor_serverTools(), (_80_dt__update__tmp_h23).Dtor_changesets(), (_80_dt__update__tmp_h23).Dtor_canvases(), (_80_dt__update__tmp_h23).Dtor_openCanvases(), (_80_dt__update__tmp_h23).Dtor_defaultChat(), (_80_dt__update__tmp_h23).Dtor_activeClients(), (_80_dt__update__tmp_h23).Dtor_chats(), _81_dt__update_hcustomizations_h3, (_80_dt__update__tmp_h23).Dtor_inputNeeded())
						}(_pat_let162_0)
					}(Companion_Default___.UpsertCustFull((_pat_let_tv17).Dtor_customizations(), _79_c))
				}(_pat_let161_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_McpServerStateChanged() {
			var _82_id _dafny.Sequence = _source0.Get_().(SessionAction_McpServerStateChanged).Id
			_ = _82_id
			var _83_st m_ConfluxCodec.Json = _source0.Get_().(SessionAction_McpServerStateChanged).State
			_ = _83_st
			var _84_ch m_AhpSkeleton.Option = _source0.Get_().(SessionAction_McpServerStateChanged).Channel
			_ = _84_ch
			if Companion_Default___.IsMcp((s).Dtor_customizations(), _82_id) {
				return _dafny.TupleOf(func(_pat_let163_0 SessionState) SessionState {
					return func(_85_dt__update__tmp_h24 SessionState) SessionState {
						return func(_pat_let164_0 _dafny.Sequence) SessionState {
							return func(_86_dt__update_hcustomizations_h4 _dafny.Sequence) SessionState {
								return Companion_SessionState_.Create_SessionState_((_85_dt__update__tmp_h24).Dtor_provider(), (_85_dt__update__tmp_h24).Dtor_title(), (_85_dt__update__tmp_h24).Dtor_status(), (_85_dt__update__tmp_h24).Dtor_lifecycle(), (_85_dt__update__tmp_h24).Dtor_activity(), (_85_dt__update__tmp_h24).Dtor_config(), (_85_dt__update__tmp_h24).Dtor_meta(), (_85_dt__update__tmp_h24).Dtor_creationError(), (_85_dt__update__tmp_h24).Dtor_serverTools(), (_85_dt__update__tmp_h24).Dtor_changesets(), (_85_dt__update__tmp_h24).Dtor_canvases(), (_85_dt__update__tmp_h24).Dtor_openCanvases(), (_85_dt__update__tmp_h24).Dtor_defaultChat(), (_85_dt__update__tmp_h24).Dtor_activeClients(), (_85_dt__update__tmp_h24).Dtor_chats(), _86_dt__update_hcustomizations_h4, (_85_dt__update__tmp_h24).Dtor_inputNeeded())
							}(_pat_let164_0)
						}(Companion_Default___.SetMcp((_pat_let_tv18).Dtor_customizations(), _82_id, (func() m_AhpSkeleton.Option {
							if (_83_st).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_()) {
								return m_AhpSkeleton.Companion_Option_.Create_None_()
							}
							return m_AhpSkeleton.Companion_Option_.Create_Some_(_83_st)
						})(), _84_ch))
					}(_pat_let163_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_McpServerStartRequested() {
			var _87_id _dafny.Sequence = _source0.Get_().(SessionAction_McpServerStartRequested).Id
			_ = _87_id
			if Companion_Default___.IsMcp((s).Dtor_customizations(), _87_id) {
				return _dafny.TupleOf(func(_pat_let165_0 SessionState) SessionState {
					return func(_88_dt__update__tmp_h25 SessionState) SessionState {
						return func(_pat_let166_0 _dafny.Sequence) SessionState {
							return func(_89_dt__update_hcustomizations_h5 _dafny.Sequence) SessionState {
								return Companion_SessionState_.Create_SessionState_((_88_dt__update__tmp_h25).Dtor_provider(), (_88_dt__update__tmp_h25).Dtor_title(), (_88_dt__update__tmp_h25).Dtor_status(), (_88_dt__update__tmp_h25).Dtor_lifecycle(), (_88_dt__update__tmp_h25).Dtor_activity(), (_88_dt__update__tmp_h25).Dtor_config(), (_88_dt__update__tmp_h25).Dtor_meta(), (_88_dt__update__tmp_h25).Dtor_creationError(), (_88_dt__update__tmp_h25).Dtor_serverTools(), (_88_dt__update__tmp_h25).Dtor_changesets(), (_88_dt__update__tmp_h25).Dtor_canvases(), (_88_dt__update__tmp_h25).Dtor_openCanvases(), (_88_dt__update__tmp_h25).Dtor_defaultChat(), (_88_dt__update__tmp_h25).Dtor_activeClients(), (_88_dt__update__tmp_h25).Dtor_chats(), _89_dt__update_hcustomizations_h5, (_88_dt__update__tmp_h25).Dtor_inputNeeded())
							}(_pat_let166_0)
						}(Companion_Default___.SetMcp((_pat_let_tv19).Dtor_customizations(), _87_id, m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("starting"))))), m_AhpSkeleton.Companion_Option_.Create_None_()))
					}(_pat_let165_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_McpServerStopRequested() {
			var _90_id _dafny.Sequence = _source0.Get_().(SessionAction_McpServerStopRequested).Id
			_ = _90_id
			if Companion_Default___.IsMcp((s).Dtor_customizations(), _90_id) {
				return _dafny.TupleOf(func(_pat_let167_0 SessionState) SessionState {
					return func(_91_dt__update__tmp_h26 SessionState) SessionState {
						return func(_pat_let168_0 _dafny.Sequence) SessionState {
							return func(_92_dt__update_hcustomizations_h6 _dafny.Sequence) SessionState {
								return Companion_SessionState_.Create_SessionState_((_91_dt__update__tmp_h26).Dtor_provider(), (_91_dt__update__tmp_h26).Dtor_title(), (_91_dt__update__tmp_h26).Dtor_status(), (_91_dt__update__tmp_h26).Dtor_lifecycle(), (_91_dt__update__tmp_h26).Dtor_activity(), (_91_dt__update__tmp_h26).Dtor_config(), (_91_dt__update__tmp_h26).Dtor_meta(), (_91_dt__update__tmp_h26).Dtor_creationError(), (_91_dt__update__tmp_h26).Dtor_serverTools(), (_91_dt__update__tmp_h26).Dtor_changesets(), (_91_dt__update__tmp_h26).Dtor_canvases(), (_91_dt__update__tmp_h26).Dtor_openCanvases(), (_91_dt__update__tmp_h26).Dtor_defaultChat(), (_91_dt__update__tmp_h26).Dtor_activeClients(), (_91_dt__update__tmp_h26).Dtor_chats(), _92_dt__update_hcustomizations_h6, (_91_dt__update__tmp_h26).Dtor_inputNeeded())
							}(_pat_let168_0)
						}(Companion_Default___.SetMcp((_pat_let_tv20).Dtor_customizations(), _90_id, m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("stopped"))))), m_AhpSkeleton.Companion_Option_.Create_None_()))
					}(_pat_let167_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
	}
}
func (_static *CompanionStruct_Default___) WfCust(c Cust) bool {
	return (!(((c).Dtor_state()).Is_Some()) || (!(((c).Dtor_state()).Dtor_value().(m_ConfluxCodec.Json)).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_()))) && (!(((c).Dtor_channel()).Is_Some()) || (!(((c).Dtor_channel()).Dtor_value().(m_ConfluxCodec.Json)).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_())))
}
func (_static *CompanionStruct_Default___) WfCusts(cs _dafny.Sequence) bool {
	return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((cs).Cardinality())), true, func(_forall_var_0 _dafny.Int) bool {
		var _0_i _dafny.Int
		_0_i = interface{}(_forall_var_0).(_dafny.Int)
		return !(((_0_i).Sign() != -1) && ((_0_i).Cmp(_dafny.IntOfUint32((cs).Cardinality())) < 0)) || (Companion_Default___.WfCust((cs).Select((_0_i).Uint32()).(Cust)))
	})
}
func (_static *CompanionStruct_Default___) SessionWf(s SessionState) bool {
	return ((((((((((!(((s).Dtor_meta()).Is_Some()) || (!(((s).Dtor_meta()).Dtor_value().(m_ConfluxCodec.Json)).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_()))) && (!(((s).Dtor_creationError()).Is_Some()) || (!(((s).Dtor_creationError()).Dtor_value().(m_ConfluxCodec.Json)).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_())))) && (!(((s).Dtor_serverTools()).Is_Some()) || (!(((s).Dtor_serverTools()).Dtor_value().(m_ConfluxCodec.Json)).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_())))) && (!(((s).Dtor_changesets()).Is_Some()) || (!(((s).Dtor_changesets()).Dtor_value().(m_ConfluxCodec.Json)).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_())))) && (!(((s).Dtor_canvases()).Is_Some()) || (!(((s).Dtor_canvases()).Dtor_value().(m_ConfluxCodec.Json)).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_())))) && (!(((s).Dtor_openCanvases()).Is_Some()) || (!(((s).Dtor_openCanvases()).Dtor_value().(m_ConfluxCodec.Json)).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_())))) && (Companion_Default___.WfCusts((s).Dtor_customizations()))) && (m_ConfluxSeqRoute.Companion_Default___.UniqueKeys(func(coer38 func(Cust) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg45 interface{}) interface{} {
			return coer38(arg45.(Cust))
		}
	}(Companion_Default___.CustKey()), (s).Dtor_customizations()))) && (m_ConfluxSeqRoute.Companion_Default___.UniqueKeys(func(coer39 func(Chat) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg46 interface{}) interface{} {
			return coer39(arg46.(Chat))
		}
	}(Companion_Default___.ChatKey()), (s).Dtor_chats()))) && (m_ConfluxSeqRoute.Companion_Default___.UniqueKeys(func(coer40 func(Client) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg47 interface{}) interface{} {
			return coer40(arg47.(Client))
		}
	}(Companion_Default___.ClientKey()), (s).Dtor_activeClients()))) && (m_ConfluxSeqRoute.Companion_Default___.UniqueKeys(func(coer41 func(InputReq) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg48 interface{}) interface{} {
			return coer41(arg48.(InputReq))
		}
	}(Companion_Default___.InputKey()), (s).Dtor_inputNeeded()))
}
func (_static *CompanionStruct_Default___) WfSessionActionInv(a SessionAction) bool {
	var _source0 SessionAction = a
	_ = _source0
	{
		if _source0.Is_CustomizationsChanged() {
			var _0_cs _dafny.Sequence = _source0.Get_().(SessionAction_CustomizationsChanged).Customizations
			_ = _0_cs
			return (Companion_Default___.WfCusts(_0_cs)) && (m_ConfluxSeqRoute.Companion_Default___.UniqueKeys(func(coer42 func(Cust) _dafny.Sequence) func(interface{}) interface{} {
				return func(arg49 interface{}) interface{} {
					return coer42(arg49.(Cust))
				}
			}(Companion_Default___.CustKey()), _0_cs))
		}
	}
	{
		if _source0.Is_CustomizationUpdated() {
			var _1_c Cust = _source0.Get_().(SessionAction_CustomizationUpdated).Customization
			_ = _1_c
			return Companion_Default___.WfCust(_1_c)
		}
	}
	{
		if _source0.Is_ChangesetsChanged() {
			var _2_cs m_AhpSkeleton.Option = _source0.Get_().(SessionAction_ChangesetsChanged).Changesets
			_ = _2_cs
			return !((_2_cs).Is_Some()) || (!((_2_cs).Dtor_value().(m_ConfluxCodec.Json)).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_()))
		}
	}
	{
		if _source0.Is_McpServerStateChanged() {
			var _3_ch m_AhpSkeleton.Option = _source0.Get_().(SessionAction_McpServerStateChanged).Channel
			_ = _3_ch
			return !((_3_ch).Is_Some()) || (!((_3_ch).Dtor_value().(m_ConfluxCodec.Json)).Equals(m_ConfluxCodec.Companion_Json_.Create_JNull_()))
		}
	}
	{
		return true
	}
}
func (_static *CompanionStruct_Default___) Apply1(s SessionState, a SessionAction) SessionState {
	return (*(Companion_Default___.ApplyToSession(s, a, _dafny.IntOfInt64(9999))).IndexInt(0)).(SessionState)
}
func (_static *CompanionStruct_Default___) Fold(s SessionState, acts _dafny.Sequence) SessionState {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer43 func(SessionState, SessionAction) SessionState) func(interface{}, interface{}) interface{} {
		return func(arg50 interface{}, arg51 interface{}) interface{} {
			return coer43(arg50.(SessionState), arg51.(SessionAction))
		}
	}(Companion_Default___.Apply1), s, acts).(SessionState)
}
func (_static *CompanionStruct_Default___) S0() SessionState {
	return Companion_SessionState_.Create_SessionState_(_dafny.UnicodeSeqOfUtf8Bytes("copilot"), _dafny.UnicodeSeqOfUtf8Bytes("Test Session"), uint32(1), _dafny.UnicodeSeqOfUtf8Bytes("creating"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), _dafny.SeqOf(), _dafny.SeqOf(), _dafny.SeqOf(), _dafny.SeqOf())
}
func (_static *CompanionStruct_Default___) RunCorpus() (_dafny.Int, _dafny.Int) {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	var modeled _dafny.Int = _dafny.Zero
	_ = modeled
	modeled = _dafny.IntOfInt64(36)
	pass = _dafny.Zero
	var _0_s SessionState
	_ = _0_s
	_0_s = Companion_Default___.S0()
	if (Companion_Default___.Apply1(func(_pat_let169_0 SessionState) SessionState {
		return func(_1_dt__update__tmp_h0 SessionState) SessionState {
			return func(_pat_let170_0 _dafny.Sequence) SessionState {
				return func(_2_dt__update_hlifecycle_h0 _dafny.Sequence) SessionState {
					return Companion_SessionState_.Create_SessionState_((_1_dt__update__tmp_h0).Dtor_provider(), (_1_dt__update__tmp_h0).Dtor_title(), (_1_dt__update__tmp_h0).Dtor_status(), _2_dt__update_hlifecycle_h0, (_1_dt__update__tmp_h0).Dtor_activity(), (_1_dt__update__tmp_h0).Dtor_config(), (_1_dt__update__tmp_h0).Dtor_meta(), (_1_dt__update__tmp_h0).Dtor_creationError(), (_1_dt__update__tmp_h0).Dtor_serverTools(), (_1_dt__update__tmp_h0).Dtor_changesets(), (_1_dt__update__tmp_h0).Dtor_canvases(), (_1_dt__update__tmp_h0).Dtor_openCanvases(), (_1_dt__update__tmp_h0).Dtor_defaultChat(), (_1_dt__update__tmp_h0).Dtor_activeClients(), (_1_dt__update__tmp_h0).Dtor_chats(), (_1_dt__update__tmp_h0).Dtor_customizations(), (_1_dt__update__tmp_h0).Dtor_inputNeeded())
				}(_pat_let170_0)
			}(_dafny.UnicodeSeqOfUtf8Bytes("creating"))
		}(_pat_let169_0)
	}(_0_s), Companion_SessionAction_.Create_Ready_())).Equals(func(_pat_let171_0 SessionState) SessionState {
		return func(_3_dt__update__tmp_h1 SessionState) SessionState {
			return func(_pat_let172_0 _dafny.Sequence) SessionState {
				return func(_4_dt__update_hlifecycle_h1 _dafny.Sequence) SessionState {
					return Companion_SessionState_.Create_SessionState_((_3_dt__update__tmp_h1).Dtor_provider(), (_3_dt__update__tmp_h1).Dtor_title(), (_3_dt__update__tmp_h1).Dtor_status(), _4_dt__update_hlifecycle_h1, (_3_dt__update__tmp_h1).Dtor_activity(), (_3_dt__update__tmp_h1).Dtor_config(), (_3_dt__update__tmp_h1).Dtor_meta(), (_3_dt__update__tmp_h1).Dtor_creationError(), (_3_dt__update__tmp_h1).Dtor_serverTools(), (_3_dt__update__tmp_h1).Dtor_changesets(), (_3_dt__update__tmp_h1).Dtor_canvases(), (_3_dt__update__tmp_h1).Dtor_openCanvases(), (_3_dt__update__tmp_h1).Dtor_defaultChat(), (_3_dt__update__tmp_h1).Dtor_activeClients(), (_3_dt__update__tmp_h1).Dtor_chats(), (_3_dt__update__tmp_h1).Dtor_customizations(), (_3_dt__update__tmp_h1).Dtor_inputNeeded())
				}(_pat_let172_0)
			}(_dafny.UnicodeSeqOfUtf8Bytes("ready"))
		}(_pat_let171_0)
	}(_0_s)) {
		pass = (pass).Plus(_dafny.One)
	}
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_CreationFailed_(m_ConfluxCodec.Companion_Json_.Create_JNull_()))).Dtor_lifecycle(), _dafny.UnicodeSeqOfUtf8Bytes("creationFailed")) {
		pass = (pass).Plus(_dafny.One)
	}
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_TitleChanged_(_dafny.UnicodeSeqOfUtf8Bytes("New")))).Dtor_title(), _dafny.UnicodeSeqOfUtf8Bytes("New")) {
		pass = (pass).Plus(_dafny.One)
	}
	var _5_tools m_ConfluxCodec.Json
	_ = _5_tools
	_5_tools = m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("name"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("bash"))))))
	if ((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_ServerToolsChanged_(_5_tools))).Dtor_serverTools()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_5_tools)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _6_m m_ConfluxCodec.Json
	_ = _6_m
	_6_m = m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("git"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("feature"))))
	if ((Companion_Default___.Apply1(func(_pat_let173_0 SessionState) SessionState {
		return func(_7_dt__update__tmp_h2 SessionState) SessionState {
			return func(_pat_let174_0 m_AhpSkeleton.Option) SessionState {
				return func(_8_dt__update_hmeta_h0 m_AhpSkeleton.Option) SessionState {
					return Companion_SessionState_.Create_SessionState_((_7_dt__update__tmp_h2).Dtor_provider(), (_7_dt__update__tmp_h2).Dtor_title(), (_7_dt__update__tmp_h2).Dtor_status(), (_7_dt__update__tmp_h2).Dtor_lifecycle(), (_7_dt__update__tmp_h2).Dtor_activity(), (_7_dt__update__tmp_h2).Dtor_config(), _8_dt__update_hmeta_h0, (_7_dt__update__tmp_h2).Dtor_creationError(), (_7_dt__update__tmp_h2).Dtor_serverTools(), (_7_dt__update__tmp_h2).Dtor_changesets(), (_7_dt__update__tmp_h2).Dtor_canvases(), (_7_dt__update__tmp_h2).Dtor_openCanvases(), (_7_dt__update__tmp_h2).Dtor_defaultChat(), (_7_dt__update__tmp_h2).Dtor_activeClients(), (_7_dt__update__tmp_h2).Dtor_chats(), (_7_dt__update__tmp_h2).Dtor_customizations(), (_7_dt__update__tmp_h2).Dtor_inputNeeded())
				}(_pat_let174_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JNull_()))
		}(_pat_let173_0)
	}(_0_s), Companion_SessionAction_.Create_MetaChanged_(_6_m))).Dtor_meta()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_6_m)) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_IsReadChanged_(true))).Dtor_status()) == (uint32(33)) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(func(_pat_let175_0 SessionState) SessionState {
		return func(_9_dt__update__tmp_h3 SessionState) SessionState {
			return func(_pat_let176_0 uint32) SessionState {
				return func(_10_dt__update_hstatus_h0 uint32) SessionState {
					return Companion_SessionState_.Create_SessionState_((_9_dt__update__tmp_h3).Dtor_provider(), (_9_dt__update__tmp_h3).Dtor_title(), _10_dt__update_hstatus_h0, (_9_dt__update__tmp_h3).Dtor_lifecycle(), (_9_dt__update__tmp_h3).Dtor_activity(), (_9_dt__update__tmp_h3).Dtor_config(), (_9_dt__update__tmp_h3).Dtor_meta(), (_9_dt__update__tmp_h3).Dtor_creationError(), (_9_dt__update__tmp_h3).Dtor_serverTools(), (_9_dt__update__tmp_h3).Dtor_changesets(), (_9_dt__update__tmp_h3).Dtor_canvases(), (_9_dt__update__tmp_h3).Dtor_openCanvases(), (_9_dt__update__tmp_h3).Dtor_defaultChat(), (_9_dt__update__tmp_h3).Dtor_activeClients(), (_9_dt__update__tmp_h3).Dtor_chats(), (_9_dt__update__tmp_h3).Dtor_customizations(), (_9_dt__update__tmp_h3).Dtor_inputNeeded())
				}(_pat_let176_0)
			}(uint32(33))
		}(_pat_let175_0)
	}(_0_s), Companion_SessionAction_.Create_IsReadChanged_(false))).Dtor_status()) == (uint32(1)) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_IsArchivedChanged_(true))).Dtor_status()) == (uint32(65)) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(func(_pat_let177_0 SessionState) SessionState {
		return func(_11_dt__update__tmp_h4 SessionState) SessionState {
			return func(_pat_let178_0 uint32) SessionState {
				return func(_12_dt__update_hstatus_h1 uint32) SessionState {
					return Companion_SessionState_.Create_SessionState_((_11_dt__update__tmp_h4).Dtor_provider(), (_11_dt__update__tmp_h4).Dtor_title(), _12_dt__update_hstatus_h1, (_11_dt__update__tmp_h4).Dtor_lifecycle(), (_11_dt__update__tmp_h4).Dtor_activity(), (_11_dt__update__tmp_h4).Dtor_config(), (_11_dt__update__tmp_h4).Dtor_meta(), (_11_dt__update__tmp_h4).Dtor_creationError(), (_11_dt__update__tmp_h4).Dtor_serverTools(), (_11_dt__update__tmp_h4).Dtor_changesets(), (_11_dt__update__tmp_h4).Dtor_canvases(), (_11_dt__update__tmp_h4).Dtor_openCanvases(), (_11_dt__update__tmp_h4).Dtor_defaultChat(), (_11_dt__update__tmp_h4).Dtor_activeClients(), (_11_dt__update__tmp_h4).Dtor_chats(), (_11_dt__update__tmp_h4).Dtor_customizations(), (_11_dt__update__tmp_h4).Dtor_inputNeeded())
				}(_pat_let178_0)
			}(uint32(65))
		}(_pat_let177_0)
	}(_0_s), Companion_SessionAction_.Create_IsArchivedChanged_(false))).Dtor_status()) == (uint32(1)) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_ActivityChanged_(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Running"))))).Equals(func(_pat_let179_0 SessionState) SessionState {
		return func(_13_dt__update__tmp_h5 SessionState) SessionState {
			return func(_pat_let180_0 m_AhpSkeleton.Option) SessionState {
				return func(_14_dt__update_hactivity_h0 m_AhpSkeleton.Option) SessionState {
					return Companion_SessionState_.Create_SessionState_((_13_dt__update__tmp_h5).Dtor_provider(), (_13_dt__update__tmp_h5).Dtor_title(), (_13_dt__update__tmp_h5).Dtor_status(), (_13_dt__update__tmp_h5).Dtor_lifecycle(), _14_dt__update_hactivity_h0, (_13_dt__update__tmp_h5).Dtor_config(), (_13_dt__update__tmp_h5).Dtor_meta(), (_13_dt__update__tmp_h5).Dtor_creationError(), (_13_dt__update__tmp_h5).Dtor_serverTools(), (_13_dt__update__tmp_h5).Dtor_changesets(), (_13_dt__update__tmp_h5).Dtor_canvases(), (_13_dt__update__tmp_h5).Dtor_openCanvases(), (_13_dt__update__tmp_h5).Dtor_defaultChat(), (_13_dt__update__tmp_h5).Dtor_activeClients(), (_13_dt__update__tmp_h5).Dtor_chats(), (_13_dt__update__tmp_h5).Dtor_customizations(), (_13_dt__update__tmp_h5).Dtor_inputNeeded())
				}(_pat_let180_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Running")))
		}(_pat_let179_0)
	}(_0_s)) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(func(_pat_let181_0 SessionState) SessionState {
		return func(_15_dt__update__tmp_h6 SessionState) SessionState {
			return func(_pat_let182_0 m_AhpSkeleton.Option) SessionState {
				return func(_16_dt__update_hactivity_h1 m_AhpSkeleton.Option) SessionState {
					return Companion_SessionState_.Create_SessionState_((_15_dt__update__tmp_h6).Dtor_provider(), (_15_dt__update__tmp_h6).Dtor_title(), (_15_dt__update__tmp_h6).Dtor_status(), (_15_dt__update__tmp_h6).Dtor_lifecycle(), _16_dt__update_hactivity_h1, (_15_dt__update__tmp_h6).Dtor_config(), (_15_dt__update__tmp_h6).Dtor_meta(), (_15_dt__update__tmp_h6).Dtor_creationError(), (_15_dt__update__tmp_h6).Dtor_serverTools(), (_15_dt__update__tmp_h6).Dtor_changesets(), (_15_dt__update__tmp_h6).Dtor_canvases(), (_15_dt__update__tmp_h6).Dtor_openCanvases(), (_15_dt__update__tmp_h6).Dtor_defaultChat(), (_15_dt__update__tmp_h6).Dtor_activeClients(), (_15_dt__update__tmp_h6).Dtor_chats(), (_15_dt__update__tmp_h6).Dtor_customizations(), (_15_dt__update__tmp_h6).Dtor_inputNeeded())
				}(_pat_let182_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("x")))
		}(_pat_let181_0)
	}(_0_s), Companion_SessionAction_.Create_ActivityChanged_(m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(_0_s) {
		pass = (pass).Plus(_dafny.One)
	}
	var _17_cfg SConfig
	_ = _17_cfg
	_17_cfg = Companion_SConfig_.Create_SConfig_(m_ConfluxCodec.Companion_Json_.Create_JNull_(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("target"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("worktree"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("baseBranch"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("main"))))
	var _18_cfgExp SConfig
	_ = _18_cfgExp
	_18_cfgExp = Companion_SConfig_.Create_SConfig_(m_ConfluxCodec.Companion_Json_.Create_JNull_(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("target"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("worktree"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("baseBranch"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("develop"))))
	var _pat_let_tv0 = _17_cfg
	_ = _pat_let_tv0
	if ((Companion_Default___.Apply1(func(_pat_let183_0 SessionState) SessionState {
		return func(_19_dt__update__tmp_h7 SessionState) SessionState {
			return func(_pat_let184_0 m_AhpSkeleton.Option) SessionState {
				return func(_20_dt__update_hconfig_h0 m_AhpSkeleton.Option) SessionState {
					return Companion_SessionState_.Create_SessionState_((_19_dt__update__tmp_h7).Dtor_provider(), (_19_dt__update__tmp_h7).Dtor_title(), (_19_dt__update__tmp_h7).Dtor_status(), (_19_dt__update__tmp_h7).Dtor_lifecycle(), (_19_dt__update__tmp_h7).Dtor_activity(), _20_dt__update_hconfig_h0, (_19_dt__update__tmp_h7).Dtor_meta(), (_19_dt__update__tmp_h7).Dtor_creationError(), (_19_dt__update__tmp_h7).Dtor_serverTools(), (_19_dt__update__tmp_h7).Dtor_changesets(), (_19_dt__update__tmp_h7).Dtor_canvases(), (_19_dt__update__tmp_h7).Dtor_openCanvases(), (_19_dt__update__tmp_h7).Dtor_defaultChat(), (_19_dt__update__tmp_h7).Dtor_activeClients(), (_19_dt__update__tmp_h7).Dtor_chats(), (_19_dt__update__tmp_h7).Dtor_customizations(), (_19_dt__update__tmp_h7).Dtor_inputNeeded())
				}(_pat_let184_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(_pat_let_tv0))
		}(_pat_let183_0)
	}(_0_s), Companion_SessionAction_.Create_ConfigChanged_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("baseBranch"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("develop"))), false))).Dtor_config()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_18_cfgExp)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _21_cfgRepl SConfig
	_ = _21_cfgRepl
	_21_cfgRepl = Companion_SConfig_.Create_SConfig_(m_ConfluxCodec.Companion_Json_.Create_JNull_(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("baseBranch"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("develop"))))
	var _pat_let_tv1 = _17_cfg
	_ = _pat_let_tv1
	if ((Companion_Default___.Apply1(func(_pat_let185_0 SessionState) SessionState {
		return func(_22_dt__update__tmp_h8 SessionState) SessionState {
			return func(_pat_let186_0 m_AhpSkeleton.Option) SessionState {
				return func(_23_dt__update_hconfig_h1 m_AhpSkeleton.Option) SessionState {
					return Companion_SessionState_.Create_SessionState_((_22_dt__update__tmp_h8).Dtor_provider(), (_22_dt__update__tmp_h8).Dtor_title(), (_22_dt__update__tmp_h8).Dtor_status(), (_22_dt__update__tmp_h8).Dtor_lifecycle(), (_22_dt__update__tmp_h8).Dtor_activity(), _23_dt__update_hconfig_h1, (_22_dt__update__tmp_h8).Dtor_meta(), (_22_dt__update__tmp_h8).Dtor_creationError(), (_22_dt__update__tmp_h8).Dtor_serverTools(), (_22_dt__update__tmp_h8).Dtor_changesets(), (_22_dt__update__tmp_h8).Dtor_canvases(), (_22_dt__update__tmp_h8).Dtor_openCanvases(), (_22_dt__update__tmp_h8).Dtor_defaultChat(), (_22_dt__update__tmp_h8).Dtor_activeClients(), (_22_dt__update__tmp_h8).Dtor_chats(), (_22_dt__update__tmp_h8).Dtor_customizations(), (_22_dt__update__tmp_h8).Dtor_inputNeeded())
				}(_pat_let186_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(_pat_let_tv1))
		}(_pat_let185_0)
	}(_0_s), Companion_SessionAction_.Create_ConfigChanged_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("baseBranch"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("develop"))), true))).Dtor_config()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_21_cfgRepl)) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_ConfigChanged_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("x"), m_ConfluxCodec.Companion_Json_.Create_JNull_()), false))).Equals(_0_s) {
		pass = (pass).Plus(_dafny.One)
	}
	var _24_pa Cust
	_ = _24_pa
	_24_pa = Companion_Cust_.Create_Cust_(_dafny.UnicodeSeqOfUtf8Bytes("plugin-a"), _dafny.UnicodeSeqOfUtf8Bytes("plugin"), _dafny.UnicodeSeqOfUtf8Bytes("https://plugins.example/a"), _dafny.UnicodeSeqOfUtf8Bytes("Plugin A"), m_AhpSkeleton.Companion_Option_.Create_Some_(true), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())
	var _25_pb Cust
	_ = _25_pb
	_25_pb = Companion_Cust_.Create_Cust_(_dafny.UnicodeSeqOfUtf8Bytes("plugin-b"), _dafny.UnicodeSeqOfUtf8Bytes("plugin"), _dafny.UnicodeSeqOfUtf8Bytes("https://plugins.example/b"), _dafny.UnicodeSeqOfUtf8Bytes("Plugin B"), m_AhpSkeleton.Companion_Option_.Create_Some_(true), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_CustomizationsChanged_(_dafny.SeqOf(_24_pa, _25_pb)))).Dtor_customizations(), _dafny.SeqOf(_24_pa, _25_pb)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _pat_let_tv2 = _24_pa
	_ = _pat_let_tv2
	var _pat_let_tv3 = _25_pb
	_ = _pat_let_tv3
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let187_0 SessionState) SessionState {
		return func(_26_dt__update__tmp_h9 SessionState) SessionState {
			return func(_pat_let188_0 _dafny.Sequence) SessionState {
				return func(_27_dt__update_hcustomizations_h0 _dafny.Sequence) SessionState {
					return Companion_SessionState_.Create_SessionState_((_26_dt__update__tmp_h9).Dtor_provider(), (_26_dt__update__tmp_h9).Dtor_title(), (_26_dt__update__tmp_h9).Dtor_status(), (_26_dt__update__tmp_h9).Dtor_lifecycle(), (_26_dt__update__tmp_h9).Dtor_activity(), (_26_dt__update__tmp_h9).Dtor_config(), (_26_dt__update__tmp_h9).Dtor_meta(), (_26_dt__update__tmp_h9).Dtor_creationError(), (_26_dt__update__tmp_h9).Dtor_serverTools(), (_26_dt__update__tmp_h9).Dtor_changesets(), (_26_dt__update__tmp_h9).Dtor_canvases(), (_26_dt__update__tmp_h9).Dtor_openCanvases(), (_26_dt__update__tmp_h9).Dtor_defaultChat(), (_26_dt__update__tmp_h9).Dtor_activeClients(), (_26_dt__update__tmp_h9).Dtor_chats(), _27_dt__update_hcustomizations_h0, (_26_dt__update__tmp_h9).Dtor_inputNeeded())
				}(_pat_let188_0)
			}(_dafny.SeqOf(_pat_let_tv2, _pat_let_tv3))
		}(_pat_let187_0)
	}(_0_s), Companion_SessionAction_.Create_CustomizationToggled_(_dafny.UnicodeSeqOfUtf8Bytes("plugin-a"), false))).Dtor_customizations(), _dafny.SeqOf(func(_pat_let189_0 Cust) Cust {
		return func(_28_dt__update__tmp_h10 Cust) Cust {
			return func(_pat_let190_0 m_AhpSkeleton.Option) Cust {
				return func(_29_dt__update_henabled_h0 m_AhpSkeleton.Option) Cust {
					return Companion_Cust_.Create_Cust_((_28_dt__update__tmp_h10).Dtor_id(), (_28_dt__update__tmp_h10).Dtor_kind(), (_28_dt__update__tmp_h10).Dtor_uri(), (_28_dt__update__tmp_h10).Dtor_name(), _29_dt__update_henabled_h0, (_28_dt__update__tmp_h10).Dtor_state(), (_28_dt__update__tmp_h10).Dtor_channel())
				}(_pat_let190_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(false))
		}(_pat_let189_0)
	}(_24_pa), _25_pb)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _pat_let_tv4 = _24_pa
	_ = _pat_let_tv4
	var _pat_let_tv5 = _25_pb
	_ = _pat_let_tv5
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let191_0 SessionState) SessionState {
		return func(_30_dt__update__tmp_h11 SessionState) SessionState {
			return func(_pat_let192_0 _dafny.Sequence) SessionState {
				return func(_31_dt__update_hcustomizations_h1 _dafny.Sequence) SessionState {
					return Companion_SessionState_.Create_SessionState_((_30_dt__update__tmp_h11).Dtor_provider(), (_30_dt__update__tmp_h11).Dtor_title(), (_30_dt__update__tmp_h11).Dtor_status(), (_30_dt__update__tmp_h11).Dtor_lifecycle(), (_30_dt__update__tmp_h11).Dtor_activity(), (_30_dt__update__tmp_h11).Dtor_config(), (_30_dt__update__tmp_h11).Dtor_meta(), (_30_dt__update__tmp_h11).Dtor_creationError(), (_30_dt__update__tmp_h11).Dtor_serverTools(), (_30_dt__update__tmp_h11).Dtor_changesets(), (_30_dt__update__tmp_h11).Dtor_canvases(), (_30_dt__update__tmp_h11).Dtor_openCanvases(), (_30_dt__update__tmp_h11).Dtor_defaultChat(), (_30_dt__update__tmp_h11).Dtor_activeClients(), (_30_dt__update__tmp_h11).Dtor_chats(), _31_dt__update_hcustomizations_h1, (_30_dt__update__tmp_h11).Dtor_inputNeeded())
				}(_pat_let192_0)
			}(_dafny.SeqOf(_pat_let_tv4, _pat_let_tv5))
		}(_pat_let191_0)
	}(_0_s), Companion_SessionAction_.Create_CustomizationRemoved_(_dafny.UnicodeSeqOfUtf8Bytes("plugin-a")))).Dtor_customizations(), _dafny.SeqOf(_25_pb)) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_DefaultChatChanged_(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("c1"))))).Dtor_defaultChat()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("c1"))) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(func(_pat_let193_0 SessionState) SessionState {
		return func(_32_dt__update__tmp_h12 SessionState) SessionState {
			return func(_pat_let194_0 m_AhpSkeleton.Option) SessionState {
				return func(_33_dt__update_hdefaultChat_h0 m_AhpSkeleton.Option) SessionState {
					return Companion_SessionState_.Create_SessionState_((_32_dt__update__tmp_h12).Dtor_provider(), (_32_dt__update__tmp_h12).Dtor_title(), (_32_dt__update__tmp_h12).Dtor_status(), (_32_dt__update__tmp_h12).Dtor_lifecycle(), (_32_dt__update__tmp_h12).Dtor_activity(), (_32_dt__update__tmp_h12).Dtor_config(), (_32_dt__update__tmp_h12).Dtor_meta(), (_32_dt__update__tmp_h12).Dtor_creationError(), (_32_dt__update__tmp_h12).Dtor_serverTools(), (_32_dt__update__tmp_h12).Dtor_changesets(), (_32_dt__update__tmp_h12).Dtor_canvases(), (_32_dt__update__tmp_h12).Dtor_openCanvases(), _33_dt__update_hdefaultChat_h0, (_32_dt__update__tmp_h12).Dtor_activeClients(), (_32_dt__update__tmp_h12).Dtor_chats(), (_32_dt__update__tmp_h12).Dtor_customizations(), (_32_dt__update__tmp_h12).Dtor_inputNeeded())
				}(_pat_let194_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("ahp-chat:/old")))
		}(_pat_let193_0)
	}(_0_s), Companion_SessionAction_.Create_DefaultChatChanged_(m_AhpSkeleton.Companion_Option_.Create_None_()))).Dtor_defaultChat()).Equals(m_AhpSkeleton.Companion_Option_.Create_None_()) {
		pass = (pass).Plus(_dafny.One)
	}
	var _34_c1 Chat
	_ = _34_c1
	_34_c1 = Companion_Chat_.Create_Chat_(_dafny.UnicodeSeqOfUtf8Bytes("ahp-chat:/c1"), _dafny.UnicodeSeqOfUtf8Bytes("Chat 1"), _dafny.One, m_AhpSkeleton.Companion_Option_.Create_None_(), _dafny.UnicodeSeqOfUtf8Bytes("t1"), m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("user")))))
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_ChatAdded_(_34_c1))).Dtor_chats(), _dafny.SeqOf(_34_c1)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _35_c1r Chat
	_ = _35_c1r
	var _36_dt__update__tmp_h13 Chat = _34_c1
	_ = _36_dt__update__tmp_h13
	var _37_dt__update_hmodifiedAt_h0 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("t2")
	_ = _37_dt__update_hmodifiedAt_h0
	var _38_dt__update_hstatus_h2 _dafny.Int = _dafny.IntOfInt64(8)
	_ = _38_dt__update_hstatus_h2
	var _39_dt__update_htitle_h0 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("Chat 1 (renamed)")
	_ = _39_dt__update_htitle_h0
	_35_c1r = Companion_Chat_.Create_Chat_((_36_dt__update__tmp_h13).Dtor_resource(), _39_dt__update_htitle_h0, _38_dt__update_hstatus_h2, (_36_dt__update__tmp_h13).Dtor_activity(), _37_dt__update_hmodifiedAt_h0, (_36_dt__update__tmp_h13).Dtor_origin())
	var _pat_let_tv6 = _34_c1
	_ = _pat_let_tv6
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let195_0 SessionState) SessionState {
		return func(_40_dt__update__tmp_h14 SessionState) SessionState {
			return func(_pat_let196_0 _dafny.Sequence) SessionState {
				return func(_41_dt__update_hchats_h0 _dafny.Sequence) SessionState {
					return Companion_SessionState_.Create_SessionState_((_40_dt__update__tmp_h14).Dtor_provider(), (_40_dt__update__tmp_h14).Dtor_title(), (_40_dt__update__tmp_h14).Dtor_status(), (_40_dt__update__tmp_h14).Dtor_lifecycle(), (_40_dt__update__tmp_h14).Dtor_activity(), (_40_dt__update__tmp_h14).Dtor_config(), (_40_dt__update__tmp_h14).Dtor_meta(), (_40_dt__update__tmp_h14).Dtor_creationError(), (_40_dt__update__tmp_h14).Dtor_serverTools(), (_40_dt__update__tmp_h14).Dtor_changesets(), (_40_dt__update__tmp_h14).Dtor_canvases(), (_40_dt__update__tmp_h14).Dtor_openCanvases(), (_40_dt__update__tmp_h14).Dtor_defaultChat(), (_40_dt__update__tmp_h14).Dtor_activeClients(), _41_dt__update_hchats_h0, (_40_dt__update__tmp_h14).Dtor_customizations(), (_40_dt__update__tmp_h14).Dtor_inputNeeded())
				}(_pat_let196_0)
			}(_dafny.SeqOf(_pat_let_tv6))
		}(_pat_let195_0)
	}(_0_s), Companion_SessionAction_.Create_ChatAdded_(_35_c1r))).Dtor_chats(), _dafny.SeqOf(_35_c1r)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _42_c2 Chat
	_ = _42_c2
	_42_c2 = Companion_Chat_.Create_Chat_(_dafny.UnicodeSeqOfUtf8Bytes("ahp-chat:/c2"), _dafny.UnicodeSeqOfUtf8Bytes("Chat 2"), _dafny.IntOfInt64(8), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Thinking")), _dafny.UnicodeSeqOfUtf8Bytes("t2"), m_ConfluxCodec.Companion_Json_.Create_JNull_())
	var _pat_let_tv7 = _34_c1
	_ = _pat_let_tv7
	var _pat_let_tv8 = _42_c2
	_ = _pat_let_tv8
	var _pat_let_tv9 = _42_c2
	_ = _pat_let_tv9
	if (Companion_Default___.Apply1(func(_pat_let197_0 SessionState) SessionState {
		return func(_43_dt__update__tmp_h15 SessionState) SessionState {
			return func(_pat_let198_0 m_AhpSkeleton.Option) SessionState {
				return func(_44_dt__update_hdefaultChat_h1 m_AhpSkeleton.Option) SessionState {
					return func(_pat_let199_0 _dafny.Sequence) SessionState {
						return func(_45_dt__update_hchats_h1 _dafny.Sequence) SessionState {
							return Companion_SessionState_.Create_SessionState_((_43_dt__update__tmp_h15).Dtor_provider(), (_43_dt__update__tmp_h15).Dtor_title(), (_43_dt__update__tmp_h15).Dtor_status(), (_43_dt__update__tmp_h15).Dtor_lifecycle(), (_43_dt__update__tmp_h15).Dtor_activity(), (_43_dt__update__tmp_h15).Dtor_config(), (_43_dt__update__tmp_h15).Dtor_meta(), (_43_dt__update__tmp_h15).Dtor_creationError(), (_43_dt__update__tmp_h15).Dtor_serverTools(), (_43_dt__update__tmp_h15).Dtor_changesets(), (_43_dt__update__tmp_h15).Dtor_canvases(), (_43_dt__update__tmp_h15).Dtor_openCanvases(), _44_dt__update_hdefaultChat_h1, (_43_dt__update__tmp_h15).Dtor_activeClients(), _45_dt__update_hchats_h1, (_43_dt__update__tmp_h15).Dtor_customizations(), (_43_dt__update__tmp_h15).Dtor_inputNeeded())
						}(_pat_let199_0)
					}(_dafny.SeqOf(_pat_let_tv7, _pat_let_tv8))
				}(_pat_let198_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("ahp-chat:/c1")))
		}(_pat_let197_0)
	}(_0_s), Companion_SessionAction_.Create_ChatRemoved_(_dafny.UnicodeSeqOfUtf8Bytes("ahp-chat:/c1")))).Equals(func(_pat_let200_0 SessionState) SessionState {
		return func(_46_dt__update__tmp_h16 SessionState) SessionState {
			return func(_pat_let201_0 m_AhpSkeleton.Option) SessionState {
				return func(_47_dt__update_hdefaultChat_h2 m_AhpSkeleton.Option) SessionState {
					return func(_pat_let202_0 _dafny.Sequence) SessionState {
						return func(_48_dt__update_hchats_h2 _dafny.Sequence) SessionState {
							return Companion_SessionState_.Create_SessionState_((_46_dt__update__tmp_h16).Dtor_provider(), (_46_dt__update__tmp_h16).Dtor_title(), (_46_dt__update__tmp_h16).Dtor_status(), (_46_dt__update__tmp_h16).Dtor_lifecycle(), (_46_dt__update__tmp_h16).Dtor_activity(), (_46_dt__update__tmp_h16).Dtor_config(), (_46_dt__update__tmp_h16).Dtor_meta(), (_46_dt__update__tmp_h16).Dtor_creationError(), (_46_dt__update__tmp_h16).Dtor_serverTools(), (_46_dt__update__tmp_h16).Dtor_changesets(), (_46_dt__update__tmp_h16).Dtor_canvases(), (_46_dt__update__tmp_h16).Dtor_openCanvases(), _47_dt__update_hdefaultChat_h2, (_46_dt__update__tmp_h16).Dtor_activeClients(), _48_dt__update_hchats_h2, (_46_dt__update__tmp_h16).Dtor_customizations(), (_46_dt__update__tmp_h16).Dtor_inputNeeded())
						}(_pat_let202_0)
					}(_dafny.SeqOf(_pat_let_tv9))
				}(_pat_let201_0)
			}(m_AhpSkeleton.Companion_Option_.Create_None_())
		}(_pat_let200_0)
	}(_0_s)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _49_c1u Chat
	_ = _49_c1u
	var _50_dt__update__tmp_h17 Chat = _34_c1
	_ = _50_dt__update__tmp_h17
	var _51_dt__update_hmodifiedAt_h1 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("t2")
	_ = _51_dt__update_hmodifiedAt_h1
	var _52_dt__update_hactivity_h2 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Waiting for approval"))
	_ = _52_dt__update_hactivity_h2
	var _53_dt__update_hstatus_h3 _dafny.Int = _dafny.IntOfInt64(24)
	_ = _53_dt__update_hstatus_h3
	_49_c1u = Companion_Chat_.Create_Chat_((_50_dt__update__tmp_h17).Dtor_resource(), (_50_dt__update__tmp_h17).Dtor_title(), _53_dt__update_hstatus_h3, _52_dt__update_hactivity_h2, _51_dt__update_hmodifiedAt_h1, (_50_dt__update__tmp_h17).Dtor_origin())
	var _pat_let_tv10 = _34_c1
	_ = _pat_let_tv10
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let203_0 SessionState) SessionState {
		return func(_54_dt__update__tmp_h18 SessionState) SessionState {
			return func(_pat_let204_0 _dafny.Sequence) SessionState {
				return func(_55_dt__update_hchats_h3 _dafny.Sequence) SessionState {
					return Companion_SessionState_.Create_SessionState_((_54_dt__update__tmp_h18).Dtor_provider(), (_54_dt__update__tmp_h18).Dtor_title(), (_54_dt__update__tmp_h18).Dtor_status(), (_54_dt__update__tmp_h18).Dtor_lifecycle(), (_54_dt__update__tmp_h18).Dtor_activity(), (_54_dt__update__tmp_h18).Dtor_config(), (_54_dt__update__tmp_h18).Dtor_meta(), (_54_dt__update__tmp_h18).Dtor_creationError(), (_54_dt__update__tmp_h18).Dtor_serverTools(), (_54_dt__update__tmp_h18).Dtor_changesets(), (_54_dt__update__tmp_h18).Dtor_canvases(), (_54_dt__update__tmp_h18).Dtor_openCanvases(), (_54_dt__update__tmp_h18).Dtor_defaultChat(), (_54_dt__update__tmp_h18).Dtor_activeClients(), _55_dt__update_hchats_h3, (_54_dt__update__tmp_h18).Dtor_customizations(), (_54_dt__update__tmp_h18).Dtor_inputNeeded())
				}(_pat_let204_0)
			}(_dafny.SeqOf(_pat_let_tv10))
		}(_pat_let203_0)
	}(_0_s), Companion_SessionAction_.Create_ChatUpdated_(_dafny.UnicodeSeqOfUtf8Bytes("ahp-chat:/c1"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.IntOfInt64(24)), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Waiting for approval")), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("t2"))))).Dtor_chats(), _dafny.SeqOf(_49_c1u)) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_ChangesetsChanged_(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf()))))).Dtor_changesets()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf()))) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(func(_pat_let205_0 SessionState) SessionState {
		return func(_56_dt__update__tmp_h19 SessionState) SessionState {
			return func(_pat_let206_0 m_AhpSkeleton.Option) SessionState {
				return func(_57_dt__update_hchangesets_h0 m_AhpSkeleton.Option) SessionState {
					return Companion_SessionState_.Create_SessionState_((_56_dt__update__tmp_h19).Dtor_provider(), (_56_dt__update__tmp_h19).Dtor_title(), (_56_dt__update__tmp_h19).Dtor_status(), (_56_dt__update__tmp_h19).Dtor_lifecycle(), (_56_dt__update__tmp_h19).Dtor_activity(), (_56_dt__update__tmp_h19).Dtor_config(), (_56_dt__update__tmp_h19).Dtor_meta(), (_56_dt__update__tmp_h19).Dtor_creationError(), (_56_dt__update__tmp_h19).Dtor_serverTools(), _57_dt__update_hchangesets_h0, (_56_dt__update__tmp_h19).Dtor_canvases(), (_56_dt__update__tmp_h19).Dtor_openCanvases(), (_56_dt__update__tmp_h19).Dtor_defaultChat(), (_56_dt__update__tmp_h19).Dtor_activeClients(), (_56_dt__update__tmp_h19).Dtor_chats(), (_56_dt__update__tmp_h19).Dtor_customizations(), (_56_dt__update__tmp_h19).Dtor_inputNeeded())
				}(_pat_let206_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf())))
		}(_pat_let205_0)
	}(_0_s), Companion_SessionAction_.Create_ChangesetsChanged_(m_AhpSkeleton.Companion_Option_.Create_None_()))).Dtor_changesets()).Equals(m_AhpSkeleton.Companion_Option_.Create_None_()) {
		pass = (pass).Plus(_dafny.One)
	}
	var _58_canvasDecls m_ConfluxCodec.Json
	_ = _58_canvasDecls
	_58_canvasDecls = m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("providerId"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("acme.editors"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("canvasId"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("markdown"))))))
	if ((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_CanvasesChanged_(_58_canvasDecls))).Dtor_canvases()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_58_canvasDecls)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _59_openRefs m_ConfluxCodec.Json
	_ = _59_openRefs
	_59_openRefs = m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("channel"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("ahp-canvas:/canvas-1"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("canvasId"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("markdown"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("availability"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("ready"))))))
	if ((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_OpenCanvasesChanged_(_59_openRefs))).Dtor_openCanvases()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_59_openRefs)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _60_cl Client
	_ = _60_cl
	_60_cl = Companion_Client_.Create_Client_(_dafny.UnicodeSeqOfUtf8Bytes("vscode-1"), m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("displayName"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("VS Code")))))
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_ActiveClientSet_(_60_cl))).Dtor_activeClients(), _dafny.SeqOf(_60_cl)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _pat_let_tv11 = _60_cl
	_ = _pat_let_tv11
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let207_0 SessionState) SessionState {
		return func(_61_dt__update__tmp_h20 SessionState) SessionState {
			return func(_pat_let208_0 _dafny.Sequence) SessionState {
				return func(_62_dt__update_hactiveClients_h0 _dafny.Sequence) SessionState {
					return Companion_SessionState_.Create_SessionState_((_61_dt__update__tmp_h20).Dtor_provider(), (_61_dt__update__tmp_h20).Dtor_title(), (_61_dt__update__tmp_h20).Dtor_status(), (_61_dt__update__tmp_h20).Dtor_lifecycle(), (_61_dt__update__tmp_h20).Dtor_activity(), (_61_dt__update__tmp_h20).Dtor_config(), (_61_dt__update__tmp_h20).Dtor_meta(), (_61_dt__update__tmp_h20).Dtor_creationError(), (_61_dt__update__tmp_h20).Dtor_serverTools(), (_61_dt__update__tmp_h20).Dtor_changesets(), (_61_dt__update__tmp_h20).Dtor_canvases(), (_61_dt__update__tmp_h20).Dtor_openCanvases(), (_61_dt__update__tmp_h20).Dtor_defaultChat(), _62_dt__update_hactiveClients_h0, (_61_dt__update__tmp_h20).Dtor_chats(), (_61_dt__update__tmp_h20).Dtor_customizations(), (_61_dt__update__tmp_h20).Dtor_inputNeeded())
				}(_pat_let208_0)
			}(_dafny.SeqOf(_pat_let_tv11, Companion_Client_.Create_Client_(_dafny.UnicodeSeqOfUtf8Bytes("cli-1"), m_ConfluxCodec.Companion_Json_.Create_JNull_())))
		}(_pat_let207_0)
	}(_0_s), Companion_SessionAction_.Create_ActiveClientRemoved_(_dafny.UnicodeSeqOfUtf8Bytes("vscode-1")))).Dtor_activeClients(), _dafny.SeqOf(Companion_Client_.Create_Client_(_dafny.UnicodeSeqOfUtf8Bytes("cli-1"), m_ConfluxCodec.Companion_Json_.Create_JNull_()))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _63_ir InputReq
	_ = _63_ir
	_63_ir = Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("req-1"), m_ConfluxCodec.Companion_Json_.Create_JNull_())
	if ((Companion_Default___.Apply1(func(_pat_let209_0 SessionState) SessionState {
		return func(_64_dt__update__tmp_h21 SessionState) SessionState {
			return func(_pat_let210_0 uint32) SessionState {
				return func(_65_dt__update_hstatus_h4 uint32) SessionState {
					return Companion_SessionState_.Create_SessionState_((_64_dt__update__tmp_h21).Dtor_provider(), (_64_dt__update__tmp_h21).Dtor_title(), _65_dt__update_hstatus_h4, (_64_dt__update__tmp_h21).Dtor_lifecycle(), (_64_dt__update__tmp_h21).Dtor_activity(), (_64_dt__update__tmp_h21).Dtor_config(), (_64_dt__update__tmp_h21).Dtor_meta(), (_64_dt__update__tmp_h21).Dtor_creationError(), (_64_dt__update__tmp_h21).Dtor_serverTools(), (_64_dt__update__tmp_h21).Dtor_changesets(), (_64_dt__update__tmp_h21).Dtor_canvases(), (_64_dt__update__tmp_h21).Dtor_openCanvases(), (_64_dt__update__tmp_h21).Dtor_defaultChat(), (_64_dt__update__tmp_h21).Dtor_activeClients(), (_64_dt__update__tmp_h21).Dtor_chats(), (_64_dt__update__tmp_h21).Dtor_customizations(), (_64_dt__update__tmp_h21).Dtor_inputNeeded())
				}(_pat_let210_0)
			}(uint32(8))
		}(_pat_let209_0)
	}(_0_s), Companion_SessionAction_.Create_InputNeededSet_(_63_ir))).Dtor_status()) == (uint32(24)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _66_i1 InputReq
	_ = _66_i1
	_66_i1 = Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("a"), m_ConfluxCodec.Companion_Json_.Create_JNull_())
	var _67_i2 InputReq
	_ = _67_i2
	_67_i2 = Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("b"), m_ConfluxCodec.Companion_Json_.Create_JNull_())
	var _pat_let_tv12 = _66_i1
	_ = _pat_let_tv12
	var _pat_let_tv13 = _67_i2
	_ = _pat_let_tv13
	if ((Companion_Default___.Apply1(func(_pat_let211_0 SessionState) SessionState {
		return func(_68_dt__update__tmp_h22 SessionState) SessionState {
			return func(_pat_let212_0 _dafny.Sequence) SessionState {
				return func(_69_dt__update_hinputNeeded_h0 _dafny.Sequence) SessionState {
					return func(_pat_let213_0 uint32) SessionState {
						return func(_70_dt__update_hstatus_h5 uint32) SessionState {
							return Companion_SessionState_.Create_SessionState_((_68_dt__update__tmp_h22).Dtor_provider(), (_68_dt__update__tmp_h22).Dtor_title(), _70_dt__update_hstatus_h5, (_68_dt__update__tmp_h22).Dtor_lifecycle(), (_68_dt__update__tmp_h22).Dtor_activity(), (_68_dt__update__tmp_h22).Dtor_config(), (_68_dt__update__tmp_h22).Dtor_meta(), (_68_dt__update__tmp_h22).Dtor_creationError(), (_68_dt__update__tmp_h22).Dtor_serverTools(), (_68_dt__update__tmp_h22).Dtor_changesets(), (_68_dt__update__tmp_h22).Dtor_canvases(), (_68_dt__update__tmp_h22).Dtor_openCanvases(), (_68_dt__update__tmp_h22).Dtor_defaultChat(), (_68_dt__update__tmp_h22).Dtor_activeClients(), (_68_dt__update__tmp_h22).Dtor_chats(), (_68_dt__update__tmp_h22).Dtor_customizations(), _69_dt__update_hinputNeeded_h0)
						}(_pat_let213_0)
					}(uint32(24))
				}(_pat_let212_0)
			}(_dafny.SeqOf(_pat_let_tv12, _pat_let_tv13))
		}(_pat_let211_0)
	}(_0_s), Companion_SessionAction_.Create_InputNeededRemoved_(_dafny.UnicodeSeqOfUtf8Bytes("a")))).Dtor_status()) == (uint32(24)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _71_mc Cust
	_ = _71_mc
	_71_mc = Companion_Cust_.Create_Cust_(_dafny.UnicodeSeqOfUtf8Bytes("mcp-1"), _dafny.UnicodeSeqOfUtf8Bytes("mcpServer"), _dafny.UnicodeSeqOfUtf8Bytes("file:///w"), _dafny.UnicodeSeqOfUtf8Bytes("MCP"), m_AhpSkeleton.Companion_Option_.Create_Some_(true), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())
	var _pat_let_tv14 = _71_mc
	_ = _pat_let_tv14
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let214_0 SessionState) SessionState {
		return func(_72_dt__update__tmp_h23 SessionState) SessionState {
			return func(_pat_let215_0 _dafny.Sequence) SessionState {
				return func(_73_dt__update_hcustomizations_h2 _dafny.Sequence) SessionState {
					return Companion_SessionState_.Create_SessionState_((_72_dt__update__tmp_h23).Dtor_provider(), (_72_dt__update__tmp_h23).Dtor_title(), (_72_dt__update__tmp_h23).Dtor_status(), (_72_dt__update__tmp_h23).Dtor_lifecycle(), (_72_dt__update__tmp_h23).Dtor_activity(), (_72_dt__update__tmp_h23).Dtor_config(), (_72_dt__update__tmp_h23).Dtor_meta(), (_72_dt__update__tmp_h23).Dtor_creationError(), (_72_dt__update__tmp_h23).Dtor_serverTools(), (_72_dt__update__tmp_h23).Dtor_changesets(), (_72_dt__update__tmp_h23).Dtor_canvases(), (_72_dt__update__tmp_h23).Dtor_openCanvases(), (_72_dt__update__tmp_h23).Dtor_defaultChat(), (_72_dt__update__tmp_h23).Dtor_activeClients(), (_72_dt__update__tmp_h23).Dtor_chats(), _73_dt__update_hcustomizations_h2, (_72_dt__update__tmp_h23).Dtor_inputNeeded())
				}(_pat_let215_0)
			}(_dafny.SeqOf(_pat_let_tv14))
		}(_pat_let214_0)
	}(_0_s), Companion_SessionAction_.Create_McpServerStartRequested_(_dafny.UnicodeSeqOfUtf8Bytes("mcp-1")))).Dtor_customizations(), _dafny.SeqOf(func(_pat_let216_0 Cust) Cust {
		return func(_74_dt__update__tmp_h24 Cust) Cust {
			return func(_pat_let217_0 m_AhpSkeleton.Option) Cust {
				return func(_75_dt__update_hchannel_h0 m_AhpSkeleton.Option) Cust {
					return func(_pat_let218_0 m_AhpSkeleton.Option) Cust {
						return func(_76_dt__update_hstate_h0 m_AhpSkeleton.Option) Cust {
							return Companion_Cust_.Create_Cust_((_74_dt__update__tmp_h24).Dtor_id(), (_74_dt__update__tmp_h24).Dtor_kind(), (_74_dt__update__tmp_h24).Dtor_uri(), (_74_dt__update__tmp_h24).Dtor_name(), (_74_dt__update__tmp_h24).Dtor_enabled(), _76_dt__update_hstate_h0, _75_dt__update_hchannel_h0)
						}(_pat_let218_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("starting"))))))
				}(_pat_let217_0)
			}(m_AhpSkeleton.Companion_Option_.Create_None_())
		}(_pat_let216_0)
	}(_71_mc))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _pat_let_tv15 = _71_mc
	_ = _pat_let_tv15
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let219_0 SessionState) SessionState {
		return func(_77_dt__update__tmp_h25 SessionState) SessionState {
			return func(_pat_let220_0 _dafny.Sequence) SessionState {
				return func(_78_dt__update_hcustomizations_h3 _dafny.Sequence) SessionState {
					return Companion_SessionState_.Create_SessionState_((_77_dt__update__tmp_h25).Dtor_provider(), (_77_dt__update__tmp_h25).Dtor_title(), (_77_dt__update__tmp_h25).Dtor_status(), (_77_dt__update__tmp_h25).Dtor_lifecycle(), (_77_dt__update__tmp_h25).Dtor_activity(), (_77_dt__update__tmp_h25).Dtor_config(), (_77_dt__update__tmp_h25).Dtor_meta(), (_77_dt__update__tmp_h25).Dtor_creationError(), (_77_dt__update__tmp_h25).Dtor_serverTools(), (_77_dt__update__tmp_h25).Dtor_changesets(), (_77_dt__update__tmp_h25).Dtor_canvases(), (_77_dt__update__tmp_h25).Dtor_openCanvases(), (_77_dt__update__tmp_h25).Dtor_defaultChat(), (_77_dt__update__tmp_h25).Dtor_activeClients(), (_77_dt__update__tmp_h25).Dtor_chats(), _78_dt__update_hcustomizations_h3, (_77_dt__update__tmp_h25).Dtor_inputNeeded())
				}(_pat_let220_0)
			}(_dafny.SeqOf(_pat_let_tv15))
		}(_pat_let219_0)
	}(_0_s), Companion_SessionAction_.Create_McpServerStateChanged_(_dafny.UnicodeSeqOfUtf8Bytes("mcp-1"), m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("ready")))), m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("ch")))))).Dtor_customizations(), _dafny.SeqOf(func(_pat_let221_0 Cust) Cust {
		return func(_79_dt__update__tmp_h26 Cust) Cust {
			return func(_pat_let222_0 m_AhpSkeleton.Option) Cust {
				return func(_80_dt__update_hchannel_h1 m_AhpSkeleton.Option) Cust {
					return func(_pat_let223_0 m_AhpSkeleton.Option) Cust {
						return func(_81_dt__update_hstate_h1 m_AhpSkeleton.Option) Cust {
							return Companion_Cust_.Create_Cust_((_79_dt__update__tmp_h26).Dtor_id(), (_79_dt__update__tmp_h26).Dtor_kind(), (_79_dt__update__tmp_h26).Dtor_uri(), (_79_dt__update__tmp_h26).Dtor_name(), (_79_dt__update__tmp_h26).Dtor_enabled(), _81_dt__update_hstate_h1, _80_dt__update_hchannel_h1)
						}(_pat_let223_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("ready"))))))
				}(_pat_let222_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("ch"))))
		}(_pat_let221_0)
	}(_71_mc))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _pat_let_tv16 = _24_pa
	_ = _pat_let_tv16
	var _pat_let_tv17 = _24_pa
	_ = _pat_let_tv17
	if (Companion_Default___.Apply1(func(_pat_let224_0 SessionState) SessionState {
		return func(_82_dt__update__tmp_h27 SessionState) SessionState {
			return func(_pat_let225_0 _dafny.Sequence) SessionState {
				return func(_83_dt__update_hcustomizations_h4 _dafny.Sequence) SessionState {
					return Companion_SessionState_.Create_SessionState_((_82_dt__update__tmp_h27).Dtor_provider(), (_82_dt__update__tmp_h27).Dtor_title(), (_82_dt__update__tmp_h27).Dtor_status(), (_82_dt__update__tmp_h27).Dtor_lifecycle(), (_82_dt__update__tmp_h27).Dtor_activity(), (_82_dt__update__tmp_h27).Dtor_config(), (_82_dt__update__tmp_h27).Dtor_meta(), (_82_dt__update__tmp_h27).Dtor_creationError(), (_82_dt__update__tmp_h27).Dtor_serverTools(), (_82_dt__update__tmp_h27).Dtor_changesets(), (_82_dt__update__tmp_h27).Dtor_canvases(), (_82_dt__update__tmp_h27).Dtor_openCanvases(), (_82_dt__update__tmp_h27).Dtor_defaultChat(), (_82_dt__update__tmp_h27).Dtor_activeClients(), (_82_dt__update__tmp_h27).Dtor_chats(), _83_dt__update_hcustomizations_h4, (_82_dt__update__tmp_h27).Dtor_inputNeeded())
				}(_pat_let225_0)
			}(_dafny.SeqOf(_pat_let_tv16))
		}(_pat_let224_0)
	}(_0_s), Companion_SessionAction_.Create_McpServerStateChanged_(_dafny.UnicodeSeqOfUtf8Bytes("plugin-a"), m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("ready")))), m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("nope")))))).Equals(func(_pat_let226_0 SessionState) SessionState {
		return func(_84_dt__update__tmp_h28 SessionState) SessionState {
			return func(_pat_let227_0 _dafny.Sequence) SessionState {
				return func(_85_dt__update_hcustomizations_h5 _dafny.Sequence) SessionState {
					return Companion_SessionState_.Create_SessionState_((_84_dt__update__tmp_h28).Dtor_provider(), (_84_dt__update__tmp_h28).Dtor_title(), (_84_dt__update__tmp_h28).Dtor_status(), (_84_dt__update__tmp_h28).Dtor_lifecycle(), (_84_dt__update__tmp_h28).Dtor_activity(), (_84_dt__update__tmp_h28).Dtor_config(), (_84_dt__update__tmp_h28).Dtor_meta(), (_84_dt__update__tmp_h28).Dtor_creationError(), (_84_dt__update__tmp_h28).Dtor_serverTools(), (_84_dt__update__tmp_h28).Dtor_changesets(), (_84_dt__update__tmp_h28).Dtor_canvases(), (_84_dt__update__tmp_h28).Dtor_openCanvases(), (_84_dt__update__tmp_h28).Dtor_defaultChat(), (_84_dt__update__tmp_h28).Dtor_activeClients(), (_84_dt__update__tmp_h28).Dtor_chats(), _85_dt__update_hcustomizations_h5, (_84_dt__update__tmp_h28).Dtor_inputNeeded())
				}(_pat_let227_0)
			}(_dafny.SeqOf(_pat_let_tv17))
		}(_pat_let226_0)
	}(_0_s)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _pat_let_tv18 = _71_mc
	_ = _pat_let_tv18
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let228_0 SessionState) SessionState {
		return func(_86_dt__update__tmp_h29 SessionState) SessionState {
			return func(_pat_let229_0 _dafny.Sequence) SessionState {
				return func(_87_dt__update_hcustomizations_h6 _dafny.Sequence) SessionState {
					return Companion_SessionState_.Create_SessionState_((_86_dt__update__tmp_h29).Dtor_provider(), (_86_dt__update__tmp_h29).Dtor_title(), (_86_dt__update__tmp_h29).Dtor_status(), (_86_dt__update__tmp_h29).Dtor_lifecycle(), (_86_dt__update__tmp_h29).Dtor_activity(), (_86_dt__update__tmp_h29).Dtor_config(), (_86_dt__update__tmp_h29).Dtor_meta(), (_86_dt__update__tmp_h29).Dtor_creationError(), (_86_dt__update__tmp_h29).Dtor_serverTools(), (_86_dt__update__tmp_h29).Dtor_changesets(), (_86_dt__update__tmp_h29).Dtor_canvases(), (_86_dt__update__tmp_h29).Dtor_openCanvases(), (_86_dt__update__tmp_h29).Dtor_defaultChat(), (_86_dt__update__tmp_h29).Dtor_activeClients(), (_86_dt__update__tmp_h29).Dtor_chats(), _87_dt__update_hcustomizations_h6, (_86_dt__update__tmp_h29).Dtor_inputNeeded())
				}(_pat_let229_0)
			}(_dafny.SeqOf(_pat_let_tv18))
		}(_pat_let228_0)
	}(_0_s), Companion_SessionAction_.Create_McpServerStopRequested_(_dafny.UnicodeSeqOfUtf8Bytes("mcp-1")))).Dtor_customizations(), _dafny.SeqOf(func(_pat_let230_0 Cust) Cust {
		return func(_88_dt__update__tmp_h30 Cust) Cust {
			return func(_pat_let231_0 m_AhpSkeleton.Option) Cust {
				return func(_89_dt__update_hchannel_h2 m_AhpSkeleton.Option) Cust {
					return func(_pat_let232_0 m_AhpSkeleton.Option) Cust {
						return func(_90_dt__update_hstate_h2 m_AhpSkeleton.Option) Cust {
							return Companion_Cust_.Create_Cust_((_88_dt__update__tmp_h30).Dtor_id(), (_88_dt__update__tmp_h30).Dtor_kind(), (_88_dt__update__tmp_h30).Dtor_uri(), (_88_dt__update__tmp_h30).Dtor_name(), (_88_dt__update__tmp_h30).Dtor_enabled(), _90_dt__update_hstate_h2, _89_dt__update_hchannel_h2)
						}(_pat_let232_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("stopped"))))))
				}(_pat_let231_0)
			}(m_AhpSkeleton.Companion_Option_.Create_None_())
		}(_pat_let230_0)
	}(_71_mc))) {
		pass = (pass).Plus(_dafny.One)
	}
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(_0_s, Companion_SessionAction_.Create_CustomizationUpdated_(_71_mc))).Dtor_customizations(), _dafny.SeqOf(_71_mc)) {
		pass = (pass).Plus(_dafny.One)
	}
	return pass, modeled
}
func (_static *CompanionStruct_Default___) CustKey() func(Cust) _dafny.Sequence {
	return func(_0_x Cust) _dafny.Sequence {
		return (_0_x).Dtor_id()
	}
}
func (_static *CompanionStruct_Default___) ChatKey() func(Chat) _dafny.Sequence {
	return func(_0_x Chat) _dafny.Sequence {
		return (_0_x).Dtor_resource()
	}
}
func (_static *CompanionStruct_Default___) ClientKey() func(Client) _dafny.Sequence {
	return func(_0_x Client) _dafny.Sequence {
		return (_0_x).Dtor_clientId()
	}
}
func (_static *CompanionStruct_Default___) InputKey() func(InputReq) _dafny.Sequence {
	return func(_0_x InputReq) _dafny.Sequence {
		return (_0_x).Dtor_id()
	}
}
func (_static *CompanionStruct_Default___) McpPred() func(Cust) bool {
	return func(_0_c Cust) bool {
		return _dafny.Companion_Sequence_.Equal((_0_c).Dtor_kind(), _dafny.UnicodeSeqOfUtf8Bytes("mcpServer"))
	}
}
func (_static *CompanionStruct_Default___) B__READ() uint32 {
	return uint32(32)
}
func (_static *CompanionStruct_Default___) B__ARCH() uint32 {
	return uint32(64)
}
func (_static *CompanionStruct_Default___) B__INPUT() uint32 {
	return uint32(16)
}

// End of class Default__

// Definition of datatype Cust
type Cust struct {
	Data_Cust_
}

func (_this Cust) Get_() Data_Cust_ {
	return _this.Data_Cust_
}

type Data_Cust_ interface {
	isCust()
}

type CompanionStruct_Cust_ struct {
}

var Companion_Cust_ = CompanionStruct_Cust_{}

type Cust_Cust struct {
	Id      _dafny.Sequence
	Kind    _dafny.Sequence
	Uri     _dafny.Sequence
	Name    _dafny.Sequence
	Enabled m_AhpSkeleton.Option
	State   m_AhpSkeleton.Option
	Channel m_AhpSkeleton.Option
}

func (Cust_Cust) isCust() {}

func (CompanionStruct_Cust_) Create_Cust_(Id _dafny.Sequence, Kind _dafny.Sequence, Uri _dafny.Sequence, Name _dafny.Sequence, Enabled m_AhpSkeleton.Option, State m_AhpSkeleton.Option, Channel m_AhpSkeleton.Option) Cust {
	return Cust{Cust_Cust{Id, Kind, Uri, Name, Enabled, State, Channel}}
}

func (_this Cust) Is_Cust() bool {
	_, ok := _this.Get_().(Cust_Cust)
	return ok
}

func (CompanionStruct_Cust_) Default() Cust {
	return Companion_Cust_.Create_Cust_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default())
}

func (_this Cust) Dtor_id() _dafny.Sequence {
	return _this.Get_().(Cust_Cust).Id
}

func (_this Cust) Dtor_kind() _dafny.Sequence {
	return _this.Get_().(Cust_Cust).Kind
}

func (_this Cust) Dtor_uri() _dafny.Sequence {
	return _this.Get_().(Cust_Cust).Uri
}

func (_this Cust) Dtor_name() _dafny.Sequence {
	return _this.Get_().(Cust_Cust).Name
}

func (_this Cust) Dtor_enabled() m_AhpSkeleton.Option {
	return _this.Get_().(Cust_Cust).Enabled
}

func (_this Cust) Dtor_state() m_AhpSkeleton.Option {
	return _this.Get_().(Cust_Cust).State
}

func (_this Cust) Dtor_channel() m_AhpSkeleton.Option {
	return _this.Get_().(Cust_Cust).Channel
}

func (_this Cust) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Cust_Cust:
		{
			return "Session.Cust.Cust" + "(" + data.Id.VerbatimString(true) + ", " + data.Kind.VerbatimString(true) + ", " + data.Uri.VerbatimString(true) + ", " + data.Name.VerbatimString(true) + ", " + _dafny.String(data.Enabled) + ", " + _dafny.String(data.State) + ", " + _dafny.String(data.Channel) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Cust) Equals(other Cust) bool {
	switch data1 := _this.Get_().(type) {
	case Cust_Cust:
		{
			data2, ok := other.Get_().(Cust_Cust)
			return ok && data1.Id.Equals(data2.Id) && data1.Kind.Equals(data2.Kind) && data1.Uri.Equals(data2.Uri) && data1.Name.Equals(data2.Name) && data1.Enabled.Equals(data2.Enabled) && data1.State.Equals(data2.State) && data1.Channel.Equals(data2.Channel)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Cust) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Cust)
	return ok && _this.Equals(typed)
}

func Type_Cust_() _dafny.TypeDescriptor {
	return type_Cust_{}
}

type type_Cust_ struct {
}

func (_this type_Cust_) Default() interface{} {
	return Companion_Cust_.Default()
}

func (_this type_Cust_) String() string {
	return "Session.Cust"
}
func (_this Cust) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Cust{}

// End of datatype Cust

// Definition of datatype Chat
type Chat struct {
	Data_Chat_
}

func (_this Chat) Get_() Data_Chat_ {
	return _this.Data_Chat_
}

type Data_Chat_ interface {
	isChat()
}

type CompanionStruct_Chat_ struct {
}

var Companion_Chat_ = CompanionStruct_Chat_{}

type Chat_Chat struct {
	Resource   _dafny.Sequence
	Title      _dafny.Sequence
	Status     _dafny.Int
	Activity   m_AhpSkeleton.Option
	ModifiedAt _dafny.Sequence
	Origin     m_ConfluxCodec.Json
}

func (Chat_Chat) isChat() {}

func (CompanionStruct_Chat_) Create_Chat_(Resource _dafny.Sequence, Title _dafny.Sequence, Status _dafny.Int, Activity m_AhpSkeleton.Option, ModifiedAt _dafny.Sequence, Origin m_ConfluxCodec.Json) Chat {
	return Chat{Chat_Chat{Resource, Title, Status, Activity, ModifiedAt, Origin}}
}

func (_this Chat) Is_Chat() bool {
	_, ok := _this.Get_().(Chat_Chat)
	return ok
}

func (CompanionStruct_Chat_) Default() Chat {
	return Companion_Chat_.Create_Chat_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.Zero, m_AhpSkeleton.Companion_Option_.Default(), _dafny.EmptySeq, m_ConfluxCodec.Companion_Json_.Default())
}

func (_this Chat) Dtor_resource() _dafny.Sequence {
	return _this.Get_().(Chat_Chat).Resource
}

func (_this Chat) Dtor_title() _dafny.Sequence {
	return _this.Get_().(Chat_Chat).Title
}

func (_this Chat) Dtor_status() _dafny.Int {
	return _this.Get_().(Chat_Chat).Status
}

func (_this Chat) Dtor_activity() m_AhpSkeleton.Option {
	return _this.Get_().(Chat_Chat).Activity
}

func (_this Chat) Dtor_modifiedAt() _dafny.Sequence {
	return _this.Get_().(Chat_Chat).ModifiedAt
}

func (_this Chat) Dtor_origin() m_ConfluxCodec.Json {
	return _this.Get_().(Chat_Chat).Origin
}

func (_this Chat) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Chat_Chat:
		{
			return "Session.Chat.Chat" + "(" + data.Resource.VerbatimString(true) + ", " + data.Title.VerbatimString(true) + ", " + _dafny.String(data.Status) + ", " + _dafny.String(data.Activity) + ", " + data.ModifiedAt.VerbatimString(true) + ", " + _dafny.String(data.Origin) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Chat) Equals(other Chat) bool {
	switch data1 := _this.Get_().(type) {
	case Chat_Chat:
		{
			data2, ok := other.Get_().(Chat_Chat)
			return ok && data1.Resource.Equals(data2.Resource) && data1.Title.Equals(data2.Title) && data1.Status.Cmp(data2.Status) == 0 && data1.Activity.Equals(data2.Activity) && data1.ModifiedAt.Equals(data2.ModifiedAt) && data1.Origin.Equals(data2.Origin)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Chat) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Chat)
	return ok && _this.Equals(typed)
}

func Type_Chat_() _dafny.TypeDescriptor {
	return type_Chat_{}
}

type type_Chat_ struct {
}

func (_this type_Chat_) Default() interface{} {
	return Companion_Chat_.Default()
}

func (_this type_Chat_) String() string {
	return "Session.Chat"
}
func (_this Chat) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Chat{}

// End of datatype Chat

// Definition of datatype Client
type Client struct {
	Data_Client_
}

func (_this Client) Get_() Data_Client_ {
	return _this.Data_Client_
}

type Data_Client_ interface {
	isClient()
}

type CompanionStruct_Client_ struct {
}

var Companion_Client_ = CompanionStruct_Client_{}

type Client_Client struct {
	ClientId _dafny.Sequence
	Raw      m_ConfluxCodec.Json
}

func (Client_Client) isClient() {}

func (CompanionStruct_Client_) Create_Client_(ClientId _dafny.Sequence, Raw m_ConfluxCodec.Json) Client {
	return Client{Client_Client{ClientId, Raw}}
}

func (_this Client) Is_Client() bool {
	_, ok := _this.Get_().(Client_Client)
	return ok
}

func (CompanionStruct_Client_) Default() Client {
	return Companion_Client_.Create_Client_(_dafny.EmptySeq, m_ConfluxCodec.Companion_Json_.Default())
}

func (_this Client) Dtor_clientId() _dafny.Sequence {
	return _this.Get_().(Client_Client).ClientId
}

func (_this Client) Dtor_raw() m_ConfluxCodec.Json {
	return _this.Get_().(Client_Client).Raw
}

func (_this Client) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Client_Client:
		{
			return "Session.Client.Client" + "(" + data.ClientId.VerbatimString(true) + ", " + _dafny.String(data.Raw) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Client) Equals(other Client) bool {
	switch data1 := _this.Get_().(type) {
	case Client_Client:
		{
			data2, ok := other.Get_().(Client_Client)
			return ok && data1.ClientId.Equals(data2.ClientId) && data1.Raw.Equals(data2.Raw)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Client) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Client)
	return ok && _this.Equals(typed)
}

func Type_Client_() _dafny.TypeDescriptor {
	return type_Client_{}
}

type type_Client_ struct {
}

func (_this type_Client_) Default() interface{} {
	return Companion_Client_.Default()
}

func (_this type_Client_) String() string {
	return "Session.Client"
}
func (_this Client) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Client{}

// End of datatype Client

// Definition of datatype InputReq
type InputReq struct {
	Data_InputReq_
}

func (_this InputReq) Get_() Data_InputReq_ {
	return _this.Data_InputReq_
}

type Data_InputReq_ interface {
	isInputReq()
}

type CompanionStruct_InputReq_ struct {
}

var Companion_InputReq_ = CompanionStruct_InputReq_{}

type InputReq_InputReq struct {
	Id  _dafny.Sequence
	Raw m_ConfluxCodec.Json
}

func (InputReq_InputReq) isInputReq() {}

func (CompanionStruct_InputReq_) Create_InputReq_(Id _dafny.Sequence, Raw m_ConfluxCodec.Json) InputReq {
	return InputReq{InputReq_InputReq{Id, Raw}}
}

func (_this InputReq) Is_InputReq() bool {
	_, ok := _this.Get_().(InputReq_InputReq)
	return ok
}

func (CompanionStruct_InputReq_) Default() InputReq {
	return Companion_InputReq_.Create_InputReq_(_dafny.EmptySeq, m_ConfluxCodec.Companion_Json_.Default())
}

func (_this InputReq) Dtor_id() _dafny.Sequence {
	return _this.Get_().(InputReq_InputReq).Id
}

func (_this InputReq) Dtor_raw() m_ConfluxCodec.Json {
	return _this.Get_().(InputReq_InputReq).Raw
}

func (_this InputReq) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case InputReq_InputReq:
		{
			return "Session.InputReq.InputReq" + "(" + data.Id.VerbatimString(true) + ", " + _dafny.String(data.Raw) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this InputReq) Equals(other InputReq) bool {
	switch data1 := _this.Get_().(type) {
	case InputReq_InputReq:
		{
			data2, ok := other.Get_().(InputReq_InputReq)
			return ok && data1.Id.Equals(data2.Id) && data1.Raw.Equals(data2.Raw)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this InputReq) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(InputReq)
	return ok && _this.Equals(typed)
}

func Type_InputReq_() _dafny.TypeDescriptor {
	return type_InputReq_{}
}

type type_InputReq_ struct {
}

func (_this type_InputReq_) Default() interface{} {
	return Companion_InputReq_.Default()
}

func (_this type_InputReq_) String() string {
	return "Session.InputReq"
}
func (_this InputReq) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = InputReq{}

// End of datatype InputReq

// Definition of datatype SConfig
type SConfig struct {
	Data_SConfig_
}

func (_this SConfig) Get_() Data_SConfig_ {
	return _this.Data_SConfig_
}

type Data_SConfig_ interface {
	isSConfig()
}

type CompanionStruct_SConfig_ struct {
}

var Companion_SConfig_ = CompanionStruct_SConfig_{}

type SConfig_SConfig struct {
	Schema m_ConfluxCodec.Json
	Values _dafny.Map
}

func (SConfig_SConfig) isSConfig() {}

func (CompanionStruct_SConfig_) Create_SConfig_(Schema m_ConfluxCodec.Json, Values _dafny.Map) SConfig {
	return SConfig{SConfig_SConfig{Schema, Values}}
}

func (_this SConfig) Is_SConfig() bool {
	_, ok := _this.Get_().(SConfig_SConfig)
	return ok
}

func (CompanionStruct_SConfig_) Default() SConfig {
	return Companion_SConfig_.Create_SConfig_(m_ConfluxCodec.Companion_Json_.Default(), _dafny.EmptyMap)
}

func (_this SConfig) Dtor_schema() m_ConfluxCodec.Json {
	return _this.Get_().(SConfig_SConfig).Schema
}

func (_this SConfig) Dtor_values() _dafny.Map {
	return _this.Get_().(SConfig_SConfig).Values
}

func (_this SConfig) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SConfig_SConfig:
		{
			return "Session.SConfig.SConfig" + "(" + _dafny.String(data.Schema) + ", " + _dafny.String(data.Values) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SConfig) Equals(other SConfig) bool {
	switch data1 := _this.Get_().(type) {
	case SConfig_SConfig:
		{
			data2, ok := other.Get_().(SConfig_SConfig)
			return ok && data1.Schema.Equals(data2.Schema) && data1.Values.Equals(data2.Values)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SConfig) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SConfig)
	return ok && _this.Equals(typed)
}

func Type_SConfig_() _dafny.TypeDescriptor {
	return type_SConfig_{}
}

type type_SConfig_ struct {
}

func (_this type_SConfig_) Default() interface{} {
	return Companion_SConfig_.Default()
}

func (_this type_SConfig_) String() string {
	return "Session.SConfig"
}
func (_this SConfig) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SConfig{}

// End of datatype SConfig

// Definition of datatype SessionState
type SessionState struct {
	Data_SessionState_
}

func (_this SessionState) Get_() Data_SessionState_ {
	return _this.Data_SessionState_
}

type Data_SessionState_ interface {
	isSessionState()
}

type CompanionStruct_SessionState_ struct {
}

var Companion_SessionState_ = CompanionStruct_SessionState_{}

type SessionState_SessionState struct {
	Provider       _dafny.Sequence
	Title          _dafny.Sequence
	Status         uint32
	Lifecycle      _dafny.Sequence
	Activity       m_AhpSkeleton.Option
	Config         m_AhpSkeleton.Option
	Meta           m_AhpSkeleton.Option
	CreationError  m_AhpSkeleton.Option
	ServerTools    m_AhpSkeleton.Option
	Changesets     m_AhpSkeleton.Option
	Canvases       m_AhpSkeleton.Option
	OpenCanvases   m_AhpSkeleton.Option
	DefaultChat    m_AhpSkeleton.Option
	ActiveClients  _dafny.Sequence
	Chats          _dafny.Sequence
	Customizations _dafny.Sequence
	InputNeeded    _dafny.Sequence
}

func (SessionState_SessionState) isSessionState() {}

func (CompanionStruct_SessionState_) Create_SessionState_(Provider _dafny.Sequence, Title _dafny.Sequence, Status uint32, Lifecycle _dafny.Sequence, Activity m_AhpSkeleton.Option, Config m_AhpSkeleton.Option, Meta m_AhpSkeleton.Option, CreationError m_AhpSkeleton.Option, ServerTools m_AhpSkeleton.Option, Changesets m_AhpSkeleton.Option, Canvases m_AhpSkeleton.Option, OpenCanvases m_AhpSkeleton.Option, DefaultChat m_AhpSkeleton.Option, ActiveClients _dafny.Sequence, Chats _dafny.Sequence, Customizations _dafny.Sequence, InputNeeded _dafny.Sequence) SessionState {
	return SessionState{SessionState_SessionState{Provider, Title, Status, Lifecycle, Activity, Config, Meta, CreationError, ServerTools, Changesets, Canvases, OpenCanvases, DefaultChat, ActiveClients, Chats, Customizations, InputNeeded}}
}

func (_this SessionState) Is_SessionState() bool {
	_, ok := _this.Get_().(SessionState_SessionState)
	return ok
}

func (CompanionStruct_SessionState_) Default() SessionState {
	return Companion_SessionState_.Create_SessionState_(_dafny.EmptySeq, _dafny.EmptySeq, 0, _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this SessionState) Dtor_provider() _dafny.Sequence {
	return _this.Get_().(SessionState_SessionState).Provider
}

func (_this SessionState) Dtor_title() _dafny.Sequence {
	return _this.Get_().(SessionState_SessionState).Title
}

func (_this SessionState) Dtor_status() uint32 {
	return _this.Get_().(SessionState_SessionState).Status
}

func (_this SessionState) Dtor_lifecycle() _dafny.Sequence {
	return _this.Get_().(SessionState_SessionState).Lifecycle
}

func (_this SessionState) Dtor_activity() m_AhpSkeleton.Option {
	return _this.Get_().(SessionState_SessionState).Activity
}

func (_this SessionState) Dtor_config() m_AhpSkeleton.Option {
	return _this.Get_().(SessionState_SessionState).Config
}

func (_this SessionState) Dtor_meta() m_AhpSkeleton.Option {
	return _this.Get_().(SessionState_SessionState).Meta
}

func (_this SessionState) Dtor_creationError() m_AhpSkeleton.Option {
	return _this.Get_().(SessionState_SessionState).CreationError
}

func (_this SessionState) Dtor_serverTools() m_AhpSkeleton.Option {
	return _this.Get_().(SessionState_SessionState).ServerTools
}

func (_this SessionState) Dtor_changesets() m_AhpSkeleton.Option {
	return _this.Get_().(SessionState_SessionState).Changesets
}

func (_this SessionState) Dtor_canvases() m_AhpSkeleton.Option {
	return _this.Get_().(SessionState_SessionState).Canvases
}

func (_this SessionState) Dtor_openCanvases() m_AhpSkeleton.Option {
	return _this.Get_().(SessionState_SessionState).OpenCanvases
}

func (_this SessionState) Dtor_defaultChat() m_AhpSkeleton.Option {
	return _this.Get_().(SessionState_SessionState).DefaultChat
}

func (_this SessionState) Dtor_activeClients() _dafny.Sequence {
	return _this.Get_().(SessionState_SessionState).ActiveClients
}

func (_this SessionState) Dtor_chats() _dafny.Sequence {
	return _this.Get_().(SessionState_SessionState).Chats
}

func (_this SessionState) Dtor_customizations() _dafny.Sequence {
	return _this.Get_().(SessionState_SessionState).Customizations
}

func (_this SessionState) Dtor_inputNeeded() _dafny.Sequence {
	return _this.Get_().(SessionState_SessionState).InputNeeded
}

func (_this SessionState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SessionState_SessionState:
		{
			return "Session.SessionState.SessionState" + "(" + data.Provider.VerbatimString(true) + ", " + data.Title.VerbatimString(true) + ", " + _dafny.String(data.Status) + ", " + data.Lifecycle.VerbatimString(true) + ", " + _dafny.String(data.Activity) + ", " + _dafny.String(data.Config) + ", " + _dafny.String(data.Meta) + ", " + _dafny.String(data.CreationError) + ", " + _dafny.String(data.ServerTools) + ", " + _dafny.String(data.Changesets) + ", " + _dafny.String(data.Canvases) + ", " + _dafny.String(data.OpenCanvases) + ", " + _dafny.String(data.DefaultChat) + ", " + _dafny.String(data.ActiveClients) + ", " + _dafny.String(data.Chats) + ", " + _dafny.String(data.Customizations) + ", " + _dafny.String(data.InputNeeded) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SessionState) Equals(other SessionState) bool {
	switch data1 := _this.Get_().(type) {
	case SessionState_SessionState:
		{
			data2, ok := other.Get_().(SessionState_SessionState)
			return ok && data1.Provider.Equals(data2.Provider) && data1.Title.Equals(data2.Title) && data1.Status == data2.Status && data1.Lifecycle.Equals(data2.Lifecycle) && data1.Activity.Equals(data2.Activity) && data1.Config.Equals(data2.Config) && data1.Meta.Equals(data2.Meta) && data1.CreationError.Equals(data2.CreationError) && data1.ServerTools.Equals(data2.ServerTools) && data1.Changesets.Equals(data2.Changesets) && data1.Canvases.Equals(data2.Canvases) && data1.OpenCanvases.Equals(data2.OpenCanvases) && data1.DefaultChat.Equals(data2.DefaultChat) && data1.ActiveClients.Equals(data2.ActiveClients) && data1.Chats.Equals(data2.Chats) && data1.Customizations.Equals(data2.Customizations) && data1.InputNeeded.Equals(data2.InputNeeded)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SessionState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SessionState)
	return ok && _this.Equals(typed)
}

func Type_SessionState_() _dafny.TypeDescriptor {
	return type_SessionState_{}
}

type type_SessionState_ struct {
}

func (_this type_SessionState_) Default() interface{} {
	return Companion_SessionState_.Default()
}

func (_this type_SessionState_) String() string {
	return "Session.SessionState"
}
func (_this SessionState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SessionState{}

// End of datatype SessionState

// Definition of datatype SessionAction
type SessionAction struct {
	Data_SessionAction_
}

func (_this SessionAction) Get_() Data_SessionAction_ {
	return _this.Data_SessionAction_
}

type Data_SessionAction_ interface {
	isSessionAction()
}

type CompanionStruct_SessionAction_ struct {
}

var Companion_SessionAction_ = CompanionStruct_SessionAction_{}

type SessionAction_Ready struct {
}

func (SessionAction_Ready) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_Ready_() SessionAction {
	return SessionAction{SessionAction_Ready{}}
}

func (_this SessionAction) Is_Ready() bool {
	_, ok := _this.Get_().(SessionAction_Ready)
	return ok
}

type SessionAction_CreationFailed struct {
	Error m_ConfluxCodec.Json
}

func (SessionAction_CreationFailed) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_CreationFailed_(Error m_ConfluxCodec.Json) SessionAction {
	return SessionAction{SessionAction_CreationFailed{Error}}
}

func (_this SessionAction) Is_CreationFailed() bool {
	_, ok := _this.Get_().(SessionAction_CreationFailed)
	return ok
}

type SessionAction_TitleChanged struct {
	Title _dafny.Sequence
}

func (SessionAction_TitleChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_TitleChanged_(Title _dafny.Sequence) SessionAction {
	return SessionAction{SessionAction_TitleChanged{Title}}
}

func (_this SessionAction) Is_TitleChanged() bool {
	_, ok := _this.Get_().(SessionAction_TitleChanged)
	return ok
}

type SessionAction_ServerToolsChanged struct {
	Tools m_ConfluxCodec.Json
}

func (SessionAction_ServerToolsChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_ServerToolsChanged_(Tools m_ConfluxCodec.Json) SessionAction {
	return SessionAction{SessionAction_ServerToolsChanged{Tools}}
}

func (_this SessionAction) Is_ServerToolsChanged() bool {
	_, ok := _this.Get_().(SessionAction_ServerToolsChanged)
	return ok
}

type SessionAction_MetaChanged struct {
	Meta m_ConfluxCodec.Json
}

func (SessionAction_MetaChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_MetaChanged_(Meta m_ConfluxCodec.Json) SessionAction {
	return SessionAction{SessionAction_MetaChanged{Meta}}
}

func (_this SessionAction) Is_MetaChanged() bool {
	_, ok := _this.Get_().(SessionAction_MetaChanged)
	return ok
}

type SessionAction_IsReadChanged struct {
	IsRead bool
}

func (SessionAction_IsReadChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_IsReadChanged_(IsRead bool) SessionAction {
	return SessionAction{SessionAction_IsReadChanged{IsRead}}
}

func (_this SessionAction) Is_IsReadChanged() bool {
	_, ok := _this.Get_().(SessionAction_IsReadChanged)
	return ok
}

type SessionAction_IsArchivedChanged struct {
	IsArchived bool
}

func (SessionAction_IsArchivedChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_IsArchivedChanged_(IsArchived bool) SessionAction {
	return SessionAction{SessionAction_IsArchivedChanged{IsArchived}}
}

func (_this SessionAction) Is_IsArchivedChanged() bool {
	_, ok := _this.Get_().(SessionAction_IsArchivedChanged)
	return ok
}

type SessionAction_ActivityChanged struct {
	Activity m_AhpSkeleton.Option
}

func (SessionAction_ActivityChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_ActivityChanged_(Activity m_AhpSkeleton.Option) SessionAction {
	return SessionAction{SessionAction_ActivityChanged{Activity}}
}

func (_this SessionAction) Is_ActivityChanged() bool {
	_, ok := _this.Get_().(SessionAction_ActivityChanged)
	return ok
}

type SessionAction_ConfigChanged struct {
	Config  _dafny.Map
	Replace bool
}

func (SessionAction_ConfigChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_ConfigChanged_(Config _dafny.Map, Replace bool) SessionAction {
	return SessionAction{SessionAction_ConfigChanged{Config, Replace}}
}

func (_this SessionAction) Is_ConfigChanged() bool {
	_, ok := _this.Get_().(SessionAction_ConfigChanged)
	return ok
}

type SessionAction_CustomizationsChanged struct {
	Customizations _dafny.Sequence
}

func (SessionAction_CustomizationsChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_CustomizationsChanged_(Customizations _dafny.Sequence) SessionAction {
	return SessionAction{SessionAction_CustomizationsChanged{Customizations}}
}

func (_this SessionAction) Is_CustomizationsChanged() bool {
	_, ok := _this.Get_().(SessionAction_CustomizationsChanged)
	return ok
}

type SessionAction_CustomizationToggled struct {
	Id      _dafny.Sequence
	Enabled bool
}

func (SessionAction_CustomizationToggled) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_CustomizationToggled_(Id _dafny.Sequence, Enabled bool) SessionAction {
	return SessionAction{SessionAction_CustomizationToggled{Id, Enabled}}
}

func (_this SessionAction) Is_CustomizationToggled() bool {
	_, ok := _this.Get_().(SessionAction_CustomizationToggled)
	return ok
}

type SessionAction_CustomizationRemoved struct {
	Id _dafny.Sequence
}

func (SessionAction_CustomizationRemoved) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_CustomizationRemoved_(Id _dafny.Sequence) SessionAction {
	return SessionAction{SessionAction_CustomizationRemoved{Id}}
}

func (_this SessionAction) Is_CustomizationRemoved() bool {
	_, ok := _this.Get_().(SessionAction_CustomizationRemoved)
	return ok
}

type SessionAction_DefaultChatChanged struct {
	Chat m_AhpSkeleton.Option
}

func (SessionAction_DefaultChatChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_DefaultChatChanged_(Chat m_AhpSkeleton.Option) SessionAction {
	return SessionAction{SessionAction_DefaultChatChanged{Chat}}
}

func (_this SessionAction) Is_DefaultChatChanged() bool {
	_, ok := _this.Get_().(SessionAction_DefaultChatChanged)
	return ok
}

type SessionAction_ChatAdded struct {
	Summary Chat
}

func (SessionAction_ChatAdded) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_ChatAdded_(Summary Chat) SessionAction {
	return SessionAction{SessionAction_ChatAdded{Summary}}
}

func (_this SessionAction) Is_ChatAdded() bool {
	_, ok := _this.Get_().(SessionAction_ChatAdded)
	return ok
}

type SessionAction_ChatRemoved struct {
	Resource _dafny.Sequence
}

func (SessionAction_ChatRemoved) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_ChatRemoved_(Resource _dafny.Sequence) SessionAction {
	return SessionAction{SessionAction_ChatRemoved{Resource}}
}

func (_this SessionAction) Is_ChatRemoved() bool {
	_, ok := _this.Get_().(SessionAction_ChatRemoved)
	return ok
}

type SessionAction_ChatUpdated struct {
	Resource   _dafny.Sequence
	Status     m_AhpSkeleton.Option
	Activity   m_AhpSkeleton.Option
	ModifiedAt m_AhpSkeleton.Option
}

func (SessionAction_ChatUpdated) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_ChatUpdated_(Resource _dafny.Sequence, Status m_AhpSkeleton.Option, Activity m_AhpSkeleton.Option, ModifiedAt m_AhpSkeleton.Option) SessionAction {
	return SessionAction{SessionAction_ChatUpdated{Resource, Status, Activity, ModifiedAt}}
}

func (_this SessionAction) Is_ChatUpdated() bool {
	_, ok := _this.Get_().(SessionAction_ChatUpdated)
	return ok
}

type SessionAction_ChangesetsChanged struct {
	Changesets m_AhpSkeleton.Option
}

func (SessionAction_ChangesetsChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_ChangesetsChanged_(Changesets m_AhpSkeleton.Option) SessionAction {
	return SessionAction{SessionAction_ChangesetsChanged{Changesets}}
}

func (_this SessionAction) Is_ChangesetsChanged() bool {
	_, ok := _this.Get_().(SessionAction_ChangesetsChanged)
	return ok
}

type SessionAction_ActiveClientSet struct {
	Client Client
}

func (SessionAction_ActiveClientSet) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_ActiveClientSet_(Client Client) SessionAction {
	return SessionAction{SessionAction_ActiveClientSet{Client}}
}

func (_this SessionAction) Is_ActiveClientSet() bool {
	_, ok := _this.Get_().(SessionAction_ActiveClientSet)
	return ok
}

type SessionAction_ActiveClientRemoved struct {
	ClientId _dafny.Sequence
}

func (SessionAction_ActiveClientRemoved) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_ActiveClientRemoved_(ClientId _dafny.Sequence) SessionAction {
	return SessionAction{SessionAction_ActiveClientRemoved{ClientId}}
}

func (_this SessionAction) Is_ActiveClientRemoved() bool {
	_, ok := _this.Get_().(SessionAction_ActiveClientRemoved)
	return ok
}

type SessionAction_CanvasesChanged struct {
	Canvases m_ConfluxCodec.Json
}

func (SessionAction_CanvasesChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_CanvasesChanged_(Canvases m_ConfluxCodec.Json) SessionAction {
	return SessionAction{SessionAction_CanvasesChanged{Canvases}}
}

func (_this SessionAction) Is_CanvasesChanged() bool {
	_, ok := _this.Get_().(SessionAction_CanvasesChanged)
	return ok
}

type SessionAction_OpenCanvasesChanged struct {
	OpenCanvases m_ConfluxCodec.Json
}

func (SessionAction_OpenCanvasesChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_OpenCanvasesChanged_(OpenCanvases m_ConfluxCodec.Json) SessionAction {
	return SessionAction{SessionAction_OpenCanvasesChanged{OpenCanvases}}
}

func (_this SessionAction) Is_OpenCanvasesChanged() bool {
	_, ok := _this.Get_().(SessionAction_OpenCanvasesChanged)
	return ok
}

type SessionAction_InputNeededSet struct {
	Req InputReq
}

func (SessionAction_InputNeededSet) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_InputNeededSet_(Req InputReq) SessionAction {
	return SessionAction{SessionAction_InputNeededSet{Req}}
}

func (_this SessionAction) Is_InputNeededSet() bool {
	_, ok := _this.Get_().(SessionAction_InputNeededSet)
	return ok
}

type SessionAction_InputNeededRemoved struct {
	Id _dafny.Sequence
}

func (SessionAction_InputNeededRemoved) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_InputNeededRemoved_(Id _dafny.Sequence) SessionAction {
	return SessionAction{SessionAction_InputNeededRemoved{Id}}
}

func (_this SessionAction) Is_InputNeededRemoved() bool {
	_, ok := _this.Get_().(SessionAction_InputNeededRemoved)
	return ok
}

type SessionAction_CustomizationUpdated struct {
	Customization Cust
}

func (SessionAction_CustomizationUpdated) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_CustomizationUpdated_(Customization Cust) SessionAction {
	return SessionAction{SessionAction_CustomizationUpdated{Customization}}
}

func (_this SessionAction) Is_CustomizationUpdated() bool {
	_, ok := _this.Get_().(SessionAction_CustomizationUpdated)
	return ok
}

type SessionAction_McpServerStateChanged struct {
	Id      _dafny.Sequence
	State   m_ConfluxCodec.Json
	Channel m_AhpSkeleton.Option
}

func (SessionAction_McpServerStateChanged) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_McpServerStateChanged_(Id _dafny.Sequence, State m_ConfluxCodec.Json, Channel m_AhpSkeleton.Option) SessionAction {
	return SessionAction{SessionAction_McpServerStateChanged{Id, State, Channel}}
}

func (_this SessionAction) Is_McpServerStateChanged() bool {
	_, ok := _this.Get_().(SessionAction_McpServerStateChanged)
	return ok
}

type SessionAction_McpServerStartRequested struct {
	Id _dafny.Sequence
}

func (SessionAction_McpServerStartRequested) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_McpServerStartRequested_(Id _dafny.Sequence) SessionAction {
	return SessionAction{SessionAction_McpServerStartRequested{Id}}
}

func (_this SessionAction) Is_McpServerStartRequested() bool {
	_, ok := _this.Get_().(SessionAction_McpServerStartRequested)
	return ok
}

type SessionAction_McpServerStopRequested struct {
	Id _dafny.Sequence
}

func (SessionAction_McpServerStopRequested) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_McpServerStopRequested_(Id _dafny.Sequence) SessionAction {
	return SessionAction{SessionAction_McpServerStopRequested{Id}}
}

func (_this SessionAction) Is_McpServerStopRequested() bool {
	_, ok := _this.Get_().(SessionAction_McpServerStopRequested)
	return ok
}

type SessionAction_SUnknown struct {
	Raw m_ConfluxCodec.Json
}

func (SessionAction_SUnknown) isSessionAction() {}

func (CompanionStruct_SessionAction_) Create_SUnknown_(Raw m_ConfluxCodec.Json) SessionAction {
	return SessionAction{SessionAction_SUnknown{Raw}}
}

func (_this SessionAction) Is_SUnknown() bool {
	_, ok := _this.Get_().(SessionAction_SUnknown)
	return ok
}

func (CompanionStruct_SessionAction_) Default() SessionAction {
	return Companion_SessionAction_.Create_Ready_()
}

func (_this SessionAction) Dtor_error() m_ConfluxCodec.Json {
	return _this.Get_().(SessionAction_CreationFailed).Error
}

func (_this SessionAction) Dtor_title() _dafny.Sequence {
	return _this.Get_().(SessionAction_TitleChanged).Title
}

func (_this SessionAction) Dtor_tools() m_ConfluxCodec.Json {
	return _this.Get_().(SessionAction_ServerToolsChanged).Tools
}

func (_this SessionAction) Dtor_meta() m_ConfluxCodec.Json {
	return _this.Get_().(SessionAction_MetaChanged).Meta
}

func (_this SessionAction) Dtor_isRead() bool {
	return _this.Get_().(SessionAction_IsReadChanged).IsRead
}

func (_this SessionAction) Dtor_isArchived() bool {
	return _this.Get_().(SessionAction_IsArchivedChanged).IsArchived
}

func (_this SessionAction) Dtor_activity() m_AhpSkeleton.Option {
	switch data := _this.Get_().(type) {
	case SessionAction_ActivityChanged:
		return data.Activity
	default:
		return data.(SessionAction_ChatUpdated).Activity
	}
}

func (_this SessionAction) Dtor_config() _dafny.Map {
	return _this.Get_().(SessionAction_ConfigChanged).Config
}

func (_this SessionAction) Dtor_replace() bool {
	return _this.Get_().(SessionAction_ConfigChanged).Replace
}

func (_this SessionAction) Dtor_customizations() _dafny.Sequence {
	return _this.Get_().(SessionAction_CustomizationsChanged).Customizations
}

func (_this SessionAction) Dtor_id() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case SessionAction_CustomizationToggled:
		return data.Id
	case SessionAction_CustomizationRemoved:
		return data.Id
	case SessionAction_InputNeededRemoved:
		return data.Id
	case SessionAction_McpServerStateChanged:
		return data.Id
	case SessionAction_McpServerStartRequested:
		return data.Id
	default:
		return data.(SessionAction_McpServerStopRequested).Id
	}
}

func (_this SessionAction) Dtor_enabled() bool {
	return _this.Get_().(SessionAction_CustomizationToggled).Enabled
}

func (_this SessionAction) Dtor_chat() m_AhpSkeleton.Option {
	return _this.Get_().(SessionAction_DefaultChatChanged).Chat
}

func (_this SessionAction) Dtor_summary() Chat {
	return _this.Get_().(SessionAction_ChatAdded).Summary
}

func (_this SessionAction) Dtor_resource() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case SessionAction_ChatRemoved:
		return data.Resource
	default:
		return data.(SessionAction_ChatUpdated).Resource
	}
}

func (_this SessionAction) Dtor_status() m_AhpSkeleton.Option {
	return _this.Get_().(SessionAction_ChatUpdated).Status
}

func (_this SessionAction) Dtor_modifiedAt() m_AhpSkeleton.Option {
	return _this.Get_().(SessionAction_ChatUpdated).ModifiedAt
}

func (_this SessionAction) Dtor_changesets() m_AhpSkeleton.Option {
	return _this.Get_().(SessionAction_ChangesetsChanged).Changesets
}

func (_this SessionAction) Dtor_client() Client {
	return _this.Get_().(SessionAction_ActiveClientSet).Client
}

func (_this SessionAction) Dtor_clientId() _dafny.Sequence {
	return _this.Get_().(SessionAction_ActiveClientRemoved).ClientId
}

func (_this SessionAction) Dtor_canvases() m_ConfluxCodec.Json {
	return _this.Get_().(SessionAction_CanvasesChanged).Canvases
}

func (_this SessionAction) Dtor_openCanvases() m_ConfluxCodec.Json {
	return _this.Get_().(SessionAction_OpenCanvasesChanged).OpenCanvases
}

func (_this SessionAction) Dtor_req() InputReq {
	return _this.Get_().(SessionAction_InputNeededSet).Req
}

func (_this SessionAction) Dtor_customization() Cust {
	return _this.Get_().(SessionAction_CustomizationUpdated).Customization
}

func (_this SessionAction) Dtor_state() m_ConfluxCodec.Json {
	return _this.Get_().(SessionAction_McpServerStateChanged).State
}

func (_this SessionAction) Dtor_channel() m_AhpSkeleton.Option {
	return _this.Get_().(SessionAction_McpServerStateChanged).Channel
}

func (_this SessionAction) Dtor_raw() m_ConfluxCodec.Json {
	return _this.Get_().(SessionAction_SUnknown).Raw
}

func (_this SessionAction) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SessionAction_Ready:
		{
			return "Session.SessionAction.Ready"
		}
	case SessionAction_CreationFailed:
		{
			return "Session.SessionAction.CreationFailed" + "(" + _dafny.String(data.Error) + ")"
		}
	case SessionAction_TitleChanged:
		{
			return "Session.SessionAction.TitleChanged" + "(" + data.Title.VerbatimString(true) + ")"
		}
	case SessionAction_ServerToolsChanged:
		{
			return "Session.SessionAction.ServerToolsChanged" + "(" + _dafny.String(data.Tools) + ")"
		}
	case SessionAction_MetaChanged:
		{
			return "Session.SessionAction.MetaChanged" + "(" + _dafny.String(data.Meta) + ")"
		}
	case SessionAction_IsReadChanged:
		{
			return "Session.SessionAction.IsReadChanged" + "(" + _dafny.String(data.IsRead) + ")"
		}
	case SessionAction_IsArchivedChanged:
		{
			return "Session.SessionAction.IsArchivedChanged" + "(" + _dafny.String(data.IsArchived) + ")"
		}
	case SessionAction_ActivityChanged:
		{
			return "Session.SessionAction.ActivityChanged" + "(" + _dafny.String(data.Activity) + ")"
		}
	case SessionAction_ConfigChanged:
		{
			return "Session.SessionAction.ConfigChanged" + "(" + _dafny.String(data.Config) + ", " + _dafny.String(data.Replace) + ")"
		}
	case SessionAction_CustomizationsChanged:
		{
			return "Session.SessionAction.CustomizationsChanged" + "(" + _dafny.String(data.Customizations) + ")"
		}
	case SessionAction_CustomizationToggled:
		{
			return "Session.SessionAction.CustomizationToggled" + "(" + data.Id.VerbatimString(true) + ", " + _dafny.String(data.Enabled) + ")"
		}
	case SessionAction_CustomizationRemoved:
		{
			return "Session.SessionAction.CustomizationRemoved" + "(" + data.Id.VerbatimString(true) + ")"
		}
	case SessionAction_DefaultChatChanged:
		{
			return "Session.SessionAction.DefaultChatChanged" + "(" + _dafny.String(data.Chat) + ")"
		}
	case SessionAction_ChatAdded:
		{
			return "Session.SessionAction.ChatAdded" + "(" + _dafny.String(data.Summary) + ")"
		}
	case SessionAction_ChatRemoved:
		{
			return "Session.SessionAction.ChatRemoved" + "(" + data.Resource.VerbatimString(true) + ")"
		}
	case SessionAction_ChatUpdated:
		{
			return "Session.SessionAction.ChatUpdated" + "(" + data.Resource.VerbatimString(true) + ", " + _dafny.String(data.Status) + ", " + _dafny.String(data.Activity) + ", " + _dafny.String(data.ModifiedAt) + ")"
		}
	case SessionAction_ChangesetsChanged:
		{
			return "Session.SessionAction.ChangesetsChanged" + "(" + _dafny.String(data.Changesets) + ")"
		}
	case SessionAction_ActiveClientSet:
		{
			return "Session.SessionAction.ActiveClientSet" + "(" + _dafny.String(data.Client) + ")"
		}
	case SessionAction_ActiveClientRemoved:
		{
			return "Session.SessionAction.ActiveClientRemoved" + "(" + data.ClientId.VerbatimString(true) + ")"
		}
	case SessionAction_CanvasesChanged:
		{
			return "Session.SessionAction.CanvasesChanged" + "(" + _dafny.String(data.Canvases) + ")"
		}
	case SessionAction_OpenCanvasesChanged:
		{
			return "Session.SessionAction.OpenCanvasesChanged" + "(" + _dafny.String(data.OpenCanvases) + ")"
		}
	case SessionAction_InputNeededSet:
		{
			return "Session.SessionAction.InputNeededSet" + "(" + _dafny.String(data.Req) + ")"
		}
	case SessionAction_InputNeededRemoved:
		{
			return "Session.SessionAction.InputNeededRemoved" + "(" + data.Id.VerbatimString(true) + ")"
		}
	case SessionAction_CustomizationUpdated:
		{
			return "Session.SessionAction.CustomizationUpdated" + "(" + _dafny.String(data.Customization) + ")"
		}
	case SessionAction_McpServerStateChanged:
		{
			return "Session.SessionAction.McpServerStateChanged" + "(" + data.Id.VerbatimString(true) + ", " + _dafny.String(data.State) + ", " + _dafny.String(data.Channel) + ")"
		}
	case SessionAction_McpServerStartRequested:
		{
			return "Session.SessionAction.McpServerStartRequested" + "(" + data.Id.VerbatimString(true) + ")"
		}
	case SessionAction_McpServerStopRequested:
		{
			return "Session.SessionAction.McpServerStopRequested" + "(" + data.Id.VerbatimString(true) + ")"
		}
	case SessionAction_SUnknown:
		{
			return "Session.SessionAction.SUnknown" + "(" + _dafny.String(data.Raw) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SessionAction) Equals(other SessionAction) bool {
	switch data1 := _this.Get_().(type) {
	case SessionAction_Ready:
		{
			_, ok := other.Get_().(SessionAction_Ready)
			return ok
		}
	case SessionAction_CreationFailed:
		{
			data2, ok := other.Get_().(SessionAction_CreationFailed)
			return ok && data1.Error.Equals(data2.Error)
		}
	case SessionAction_TitleChanged:
		{
			data2, ok := other.Get_().(SessionAction_TitleChanged)
			return ok && data1.Title.Equals(data2.Title)
		}
	case SessionAction_ServerToolsChanged:
		{
			data2, ok := other.Get_().(SessionAction_ServerToolsChanged)
			return ok && data1.Tools.Equals(data2.Tools)
		}
	case SessionAction_MetaChanged:
		{
			data2, ok := other.Get_().(SessionAction_MetaChanged)
			return ok && data1.Meta.Equals(data2.Meta)
		}
	case SessionAction_IsReadChanged:
		{
			data2, ok := other.Get_().(SessionAction_IsReadChanged)
			return ok && data1.IsRead == data2.IsRead
		}
	case SessionAction_IsArchivedChanged:
		{
			data2, ok := other.Get_().(SessionAction_IsArchivedChanged)
			return ok && data1.IsArchived == data2.IsArchived
		}
	case SessionAction_ActivityChanged:
		{
			data2, ok := other.Get_().(SessionAction_ActivityChanged)
			return ok && data1.Activity.Equals(data2.Activity)
		}
	case SessionAction_ConfigChanged:
		{
			data2, ok := other.Get_().(SessionAction_ConfigChanged)
			return ok && data1.Config.Equals(data2.Config) && data1.Replace == data2.Replace
		}
	case SessionAction_CustomizationsChanged:
		{
			data2, ok := other.Get_().(SessionAction_CustomizationsChanged)
			return ok && data1.Customizations.Equals(data2.Customizations)
		}
	case SessionAction_CustomizationToggled:
		{
			data2, ok := other.Get_().(SessionAction_CustomizationToggled)
			return ok && data1.Id.Equals(data2.Id) && data1.Enabled == data2.Enabled
		}
	case SessionAction_CustomizationRemoved:
		{
			data2, ok := other.Get_().(SessionAction_CustomizationRemoved)
			return ok && data1.Id.Equals(data2.Id)
		}
	case SessionAction_DefaultChatChanged:
		{
			data2, ok := other.Get_().(SessionAction_DefaultChatChanged)
			return ok && data1.Chat.Equals(data2.Chat)
		}
	case SessionAction_ChatAdded:
		{
			data2, ok := other.Get_().(SessionAction_ChatAdded)
			return ok && data1.Summary.Equals(data2.Summary)
		}
	case SessionAction_ChatRemoved:
		{
			data2, ok := other.Get_().(SessionAction_ChatRemoved)
			return ok && data1.Resource.Equals(data2.Resource)
		}
	case SessionAction_ChatUpdated:
		{
			data2, ok := other.Get_().(SessionAction_ChatUpdated)
			return ok && data1.Resource.Equals(data2.Resource) && data1.Status.Equals(data2.Status) && data1.Activity.Equals(data2.Activity) && data1.ModifiedAt.Equals(data2.ModifiedAt)
		}
	case SessionAction_ChangesetsChanged:
		{
			data2, ok := other.Get_().(SessionAction_ChangesetsChanged)
			return ok && data1.Changesets.Equals(data2.Changesets)
		}
	case SessionAction_ActiveClientSet:
		{
			data2, ok := other.Get_().(SessionAction_ActiveClientSet)
			return ok && data1.Client.Equals(data2.Client)
		}
	case SessionAction_ActiveClientRemoved:
		{
			data2, ok := other.Get_().(SessionAction_ActiveClientRemoved)
			return ok && data1.ClientId.Equals(data2.ClientId)
		}
	case SessionAction_CanvasesChanged:
		{
			data2, ok := other.Get_().(SessionAction_CanvasesChanged)
			return ok && data1.Canvases.Equals(data2.Canvases)
		}
	case SessionAction_OpenCanvasesChanged:
		{
			data2, ok := other.Get_().(SessionAction_OpenCanvasesChanged)
			return ok && data1.OpenCanvases.Equals(data2.OpenCanvases)
		}
	case SessionAction_InputNeededSet:
		{
			data2, ok := other.Get_().(SessionAction_InputNeededSet)
			return ok && data1.Req.Equals(data2.Req)
		}
	case SessionAction_InputNeededRemoved:
		{
			data2, ok := other.Get_().(SessionAction_InputNeededRemoved)
			return ok && data1.Id.Equals(data2.Id)
		}
	case SessionAction_CustomizationUpdated:
		{
			data2, ok := other.Get_().(SessionAction_CustomizationUpdated)
			return ok && data1.Customization.Equals(data2.Customization)
		}
	case SessionAction_McpServerStateChanged:
		{
			data2, ok := other.Get_().(SessionAction_McpServerStateChanged)
			return ok && data1.Id.Equals(data2.Id) && data1.State.Equals(data2.State) && data1.Channel.Equals(data2.Channel)
		}
	case SessionAction_McpServerStartRequested:
		{
			data2, ok := other.Get_().(SessionAction_McpServerStartRequested)
			return ok && data1.Id.Equals(data2.Id)
		}
	case SessionAction_McpServerStopRequested:
		{
			data2, ok := other.Get_().(SessionAction_McpServerStopRequested)
			return ok && data1.Id.Equals(data2.Id)
		}
	case SessionAction_SUnknown:
		{
			data2, ok := other.Get_().(SessionAction_SUnknown)
			return ok && data1.Raw.Equals(data2.Raw)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SessionAction) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SessionAction)
	return ok && _this.Equals(typed)
}

func Type_SessionAction_() _dafny.TypeDescriptor {
	return type_SessionAction_{}
}

type type_SessionAction_ struct {
}

func (_this type_SessionAction_) Default() interface{} {
	return Companion_SessionAction_.Default()
}

func (_this type_SessionAction_) String() string {
	return "Session.SessionAction"
}
func (_this SessionAction) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SessionAction{}

// End of datatype SessionAction
