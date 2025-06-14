// Package nfabuilder implements some structures and functions to construct NFA.
package nfa

import (
	"github.com/y-yu/sfa-go/common"
	"github.com/y-yu/sfa-go/utils"
)

// Fragment represents a fragment of NFA to construct a larger NFA.
type Fragment struct {
	I     common.State // initial state
	F     utils.Set    // accept states
	Rules RuleMap
}

// NewFragment returns a new Fragment.
func NewFragment() *Fragment {
	return &Fragment{
		I:     common.NewState(0),
		F:     utils.NewSet(),
		Rules: RuleMap{},
	}
}

// AddRule add a new transition rule to the Fragment.
// Rule concept: State(from) -->[Symbol(c)]--> State(next)
func (frg *Fragment) AddRule(from common.State, c rune, next common.State) {
	r := frg.Rules
	_, ok := r[common.NewRuleArgs(from, c)]
	if ok {
		r[common.NewRuleArgs(from, c)].Add(next)
	} else {
		r[common.NewRuleArgs(from, c)] = utils.NewSet(next)
	}
}

// CreateSkeleton returns a nfa fragment which has
// same transition rule as original fragment has.
// The initial state and accept state is set to default.
func (frg *Fragment) CreateSkeleton() (Skeleton *Fragment) {
	Skeleton = NewFragment()
	Skeleton.Rules = frg.Rules
	return
}

// MergeRule returns a new NFA fragment into which the
// transition rules of original fragment and the fragment
// given in the argument are merged.
func (frg *Fragment) MergeRule(frg2 *Fragment) (synthesizedFrg *Fragment) {
	synthesizedFrg = frg.CreateSkeleton()
	for k, v := range frg2.Rules {
		_, ok := synthesizedFrg.Rules[k]
		if !ok {
			synthesizedFrg.Rules[k] = utils.NewSet()
		}
		synthesizedFrg.Rules[k] = synthesizedFrg.Rules[k].Union(v)
	}
	return
}

// Build converts NFA fragments into a NFA, and returns it.
func (frg *Fragment) Build() *NFA {
	return NewNFA(frg.I, frg.F, frg.Rules)
}
