package sfa

import (
	"github.com/y-yu/sfa-go/dfa"
	"github.com/y-yu/sfa-go/dfa/dfarule"
	"github.com/y-yu/sfa-go/utils"
)

type SFA struct {
	I     utils.State
	F     []utils.State
	Rules dfarule.RuleMap
}

func ToSFA(dfa dfa.DFA) SFA {
	//stateMap := StateMap{}
	panic("not implemented")
}
