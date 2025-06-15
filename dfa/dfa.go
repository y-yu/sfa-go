// Package dfa implements Deterministic Finite Automaton(DFA).
package dfa

import (
	"github.com/samber/lo"
	"github.com/y-yu/sfa-go/common"
	"github.com/y-yu/sfa-go/utils"
	"maps"
	"slices"
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
	states := dfa.AllStates()
	n := len(states)

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
	result := lo.Uniq(allState)
	slices.Sort(result)

	return result
}

func (dfa *DFA) AllSymbol() []rune {
	result := lo.Uniq(
		lo.Map(slices.Collect(maps.Keys(dfa.Rules)), func(item common.RuleArg, _ int) rune {
			return item.C
		}),
	)
	slices.Sort(result)

	return result
}

func (d *DFA) Match(str string) bool {
	cur := d.I
	for _, c := range []rune(str) {
		key := common.NewRuleArgs(cur, c)
		cur = d.Rules[key]
	}
	return d.F.Contains(cur)
}
