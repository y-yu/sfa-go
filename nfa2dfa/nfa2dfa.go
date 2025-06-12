package nfa2dfa

import (
	"github.com/y-yu/sfa-go/dfa"
	"github.com/y-yu/sfa-go/nfa"
)

// ToDFA converts a NFA into a DFA which recognizes the same formal language.
func ToDFA(nfa *nfa.NFA) *dfa.DFA {
	nfa.ToWithoutEpsilon()
	I, F, Delta := nfa.SubsetConstruction()
	return dfa.NewDFA(I, F, Delta)
}
