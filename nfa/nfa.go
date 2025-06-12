// Package nfa implements Non-Deterministic Finite Automaton(NFA).
package nfa

import (
	"github.com/samber/lo"
	"github.com/y-yu/sfa-go/dfa/dfarule"
	"github.com/y-yu/sfa-go/nfa/nfarule"
	"github.com/y-yu/sfa-go/utils"
)

// NFA represents a Non-Deterministic Finite Automaton.
type NFA struct {
	I     utils.State     // initial state
	F     utils.Set       // accept states
	Rules nfarule.RuleMap // transition function
}

// NewNFA returns a new NFA.
func NewNFA(init utils.State, accepts utils.Set, rules nfarule.RuleMap) *NFA {
	return &NFA{
		I:     init,
		F:     accepts,
		Rules: rules,
	}
}

// allStates returns a set of the all "From State" in Rule.
func (nfa *NFA) allStates() utils.Set {
	states := utils.NewSet()
	for key := range nfa.Rules {
		states.Add(key.From)
	}
	return states
}

// AllSymbol returns a set of the all "Symbol" in Rule.
func (nfa *NFA) AllSymbol() utils.MapSet[rune] {
	symbols := utils.NewMapSet[rune]()
	for key := range nfa.Rules {
		symbols.Add(key.C)
	}
	return symbols
}

// CalcDst returns, according to the transition function, a set of states
// to which transition is executed when c is received in the state of argument q.
func (nfa *NFA) CalcDst(q utils.State, c rune) (utils.Set, bool) {
	s, ok := nfa.Rules[nfarule.NewRuleArgs(q, c)]
	if ok {
		return s, true
	}
	return nil, false
}

// ToWithoutEpsilon update ε-NFA to NFA whose no epsilon transitions.
func (nfa *NFA) ToWithoutEpsilon() {
	if nfa.F.IsSubset(nfa.epsilonClosure(nfa.I)) {
		nfa.F.Add(nfa.I)
	}
	nfa.Rules = nfa.removeEpsilonRule()
}

// removeEpsilonRule returns a new RuleMap removing epsilon transitions
// from original RuleMap.
func (nfa *NFA) removeEpsilonRule() (newRule nfarule.RuleMap) {
	newRule = nfarule.RuleMap{}
	states, sym := nfa.allStates(), nfa.AllSymbol()
	sym.Remove('ε')

	for q := range states.Iter() {
		for c := range sym.Iter() {
			for mid := range nfa.epsilonClosure(q).Iter() {
				dst := nfa.epsilonExpand(mid, c)
				s, ok := newRule[nfarule.NewRuleArgs(q, c)]
				if !ok {
					s = utils.NewSet()
				}
				newRule[nfarule.NewRuleArgs(q, c)] = s.Union(dst)
			}
		}
	}

	for k := range newRule {
		if newRule[k].Cardinality() == 0 {
			delete(newRule, k)
		}
	}

	return
}

// epsilonExpand returns the state set, which is a result of simulating the transitions like 'ε*->symbol->ε*'.
func (nfa *NFA) epsilonExpand(state utils.State, symbol rune) (final utils.Set) {
	first := nfa.epsilonClosure(state)

	second := utils.NewSet()
	for q := range first.Iter() {
		if dst, ok := nfa.CalcDst(q, symbol); ok {
			second = second.Union(dst)
		}
	}

	final = utils.NewSet()
	for q := range second.Iter() {
		dst := nfa.epsilonClosure(q)
		final = final.Union(dst)
	}

	return
}

// epsilonClosure returns a set of reachable states with epsilon transitions only.
func (nfa *NFA) epsilonClosure(state utils.State) (reachable utils.Set) {
	reachable = utils.NewSet(state)

	modified := true
	for modified {
		modified = false
		for q := range reachable.Iter() {
			dst, ok := nfa.CalcDst(q, 'ε')
			if !ok || reachable.IsSuperset(dst) {
				continue
			}
			reachable = reachable.Union(dst)
			modified = true
		}
	}
	return
}

// subsetConstruction implements Subset Construction.
// Returns the data for constructing the equivalent DFA from the NFA given in the argument.
// For details: https://en.wikipedia.org/wiki/Powerset_construction
func (nfa *NFA) SubsetConstruction() (dI utils.State, dF utils.Set, dRules dfarule.RuleMap) {
	I := nfa.I
	F := nfa.F
	rules := nfa.Rules

	dI = utils.NewState(0)
	dF = utils.NewSet()
	dRules = dfarule.RuleMap{}

	dStates := []dfaStatesMap{
		{utils.NewSet(I), dI},
	}

	queue := []utils.Set{utils.NewSet(I)}
	sigma := nfa.AllSymbol()
	for len(queue) != 0 {
		states := queue[0]
		queue = queue[1:] // the states set which can be reached from a NFA states.

		fromDfaStateMap, _ := lo.Find(dStates, func(ds dfaStatesMap) bool {
			return ds.nfaStateSet.Equal(states)
		})

		if F.Intersect(states).Cardinality() > 0 {
			dF.Add(fromDfaStateMap.dfaState)
		}

		for c := range sigma.Iter() {
			dnext := utils.NewSet()
			for q := range states.Iter() {
				d, ok := rules[nfarule.NewRuleArgs(q, c)]
				if ok {
					dnext = dnext.Union(d)
				}
			}

			dfaStateMap, found := lo.Find(dStates, func(ds dfaStatesMap) bool {
				return ds.nfaStateSet.Equal(dnext)
			})
			var dState utils.State
			if !found {
				queue = append(queue, dnext)
				dState = utils.NewState(len(dStates))
				dStates = append(dStates, dfaStatesMap{dnext, dState})
			} else {
				dState = dfaStateMap.dfaState
			}

			dRules[dfarule.NewRuleArgs(fromDfaStateMap.dfaState, c)] = dState
		}
	}

	return
}

// DFAStatesMap associates subsets of the NFA state set with the states of the DFA.
type dfaStatesMap struct {
	nfaStateSet utils.Set
	dfaState    utils.State
}
