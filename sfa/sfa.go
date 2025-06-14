package sfa

import (
	"github.com/samber/lo"
	"github.com/y-yu/sfa-go/common"
	"github.com/y-yu/sfa-go/dfa"
	"github.com/y-yu/sfa-go/utils"
)

type SFA struct {
	I     common.State
	F     utils.Set
	Rules dfa.RuleMap
}

func ToSFA(d *dfa.DFA) *SFA {
	allState := d.AllStates()
	fi := make(StateMap)
	for _, state := range allState {
		fi[state] = state
	}
	queue := []StateMap{fi}
	sigma := d.AllSymbol()

	sfaStates := []sfaStatesMap{
		{dfaStateMap: fi, sfaState: d.I},
	}
	rules := make(dfa.RuleMap)
	f := utils.NewSet()
	if d.F.Contains(d.I) {
		f.Add(d.I)
	}

	for len(queue) != 0 {
		ftmp := queue[0]
		queue = queue[1:]

		fromSfaStateMap, _ := lo.Find(sfaStates, func(sm sfaStatesMap) bool {
			return sm.dfaStateMap.Equal(ftmp)
		})

		for c := range sigma.Iter() {
			fnext := make(StateMap)

			for _, state := range allState {
				previous := ftmp[state]
				arg := common.NewRuleArgs(previous, c)
				fnext[state] = d.Rules[arg]
			}

			sfaState, find := lo.Find(sfaStates, func(sm sfaStatesMap) bool {
				return sm.dfaStateMap.Equal(fnext)
			})

			var state common.State
			if !find {
				state = common.NewState(len(sfaStates))
				if d.F.Contains(fnext[d.I]) {
					f.Add(state)
				}
				queue = append(queue, fnext)
				sfaStates = append(sfaStates, sfaStatesMap{fnext, state})
			} else {
				state = sfaState.sfaState
			}

			rules[common.NewRuleArgs(fromSfaStateMap.sfaState, c)] = state
		}
	}

	return &SFA{
		I:     d.I,
		F:     f,
		Rules: rules,
	}
}

type sfaStatesMap struct {
	dfaStateMap StateMap
	sfaState    common.State
}
