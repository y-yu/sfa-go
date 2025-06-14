package sfa

import (
	"github.com/y-yu/sfa-go/common"
	"github.com/y-yu/sfa-go/dfa"
)

type SFA struct {
	I     common.State
	F     []common.State
	Rules dfa.RuleMap
}

func ToSFA(dfa dfa.DFA) SFA {
	//stateMap := StateMap{}
	panic("not implemented")
}
