// Package dfa implements Deterministic Finite Automaton(DFA).
package dfa

import (
	"github.com/samber/lo"
	"github.com/y-yu/sfa-go/common"
	"github.com/y-yu/sfa-go/utils"
)

// DFA represents a Deterministic Finite Automaton.
type DFA struct {
	I     common.State // initial state
	F     utils.Set
	Rules RuleMap // transition function
}

// NewDFA returns a new dfa.
func NewDFA(init common.State, accepts utils.Set, rules RuleMap) *DFA {
	return &DFA{
		I:     init,
		F:     accepts,
		Rules: rules,
	}
}

// Minimize minimizes the DFA.
func (dfa *DFA) Minimize() {
	states := utils.NewSet(dfa.I)
	for _, v := range dfa.Rules {
		states.Add(v)
	}
	n := states.Cardinality()

	eqMap := map[common.State]common.State{}
	for i := 0; i < n; i++ {
		q1 := common.NewState(i)
		for j := i + 1; j < n; j++ {
			q2 := common.NewState(j)
			if !dfa.isEquivalent(q1, q2) {
				continue
			}
			if _, ok := eqMap[q2]; ok {
				continue
			}
			eqMap[q2] = q1
			dfa.mergeState(q1, q2)
		}
	}
}

func (dfa *DFA) replaceState(to, from common.State) {
	if dfa.I == from {
		dfa.I = to
	}

	rules := dfa.Rules
	for arg, dst := range rules {
		if dst == from {
			rules[arg] = to
		}
	}
}

func (dfa *DFA) deleteState(q common.State) {
	rules := dfa.Rules
	for arg := range rules {
		if arg.From == q {
			delete(rules, arg)
		}
	}
}

func (dfa *DFA) mergeState(to, from common.State) {
	dfa.replaceState(to, from)
	dfa.deleteState(from)
}

func (dfa *DFA) isEquivalent(q1, q2 common.State) bool {
	if !((dfa.F.Contains(q1) && dfa.F.Contains(q2)) ||
		(!dfa.F.Contains(q1) && !dfa.F.Contains(q2))) {
		return false
	}

	rules := dfa.Rules
	for k := range rules {
		if k.From != q1 {
			continue
		}
		if rules[common.NewRuleArgs(q1, k.C)] != rules[common.NewRuleArgs(q2, k.C)] {
			return false
		}
	}
	return true
}

func (dfa *DFA) AllStates() []common.State {
	allState := []common.State{dfa.I}
	for k, v := range dfa.Rules {
		allState = append(allState, v, k.From)
	}
	return lo.Uniq(allState)
}

// Runtime has a pointer to d and saves current state for
// simulating d transitions.
type Runtime struct {
	d   *DFA
	cur common.State
}

// GetRuntime returns a new Runtime for simulating d transitions.
func (dfa *DFA) GetRuntime() *Runtime {
	return NewRuntime(dfa)
}

// NewRuntime returns a new runtime for DFA.
func NewRuntime(d *DFA) (r *Runtime) {
	r = &Runtime{
		d: d,
	}
	r.cur = d.I
	return
}

// transit execute a transition with a symbol, and returns whether
// the transition is success (or not).
func (r *Runtime) transit(c rune) bool {
	key := common.NewRuleArgs(r.cur, c)
	_, ok := r.d.Rules[key]
	if ok {
		r.cur = r.d.Rules[key]
		return true
	}
	return false
}

// isAccept returns whether current status is in accept states.
func (r *Runtime) isAccept() bool {
	accepts := r.d.F
	if accepts.Contains(r.cur) {
		return true
	}
	return false
}

// Matching returns whether the string given is accepted (or not) by
// simulating the all transitions.
func (r *Runtime) Matching(str string) bool {
	r.cur = r.d.I
	for _, c := range []rune(str) {
		if !r.transit(c) {
			return false // if the transition failed, the input "str" is rejected.
		}
	}
	return r.isAccept()
}
