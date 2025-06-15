package sfa

import (
	"fmt"
	"github.com/y-yu/sfa-go/common"
	"github.com/y-yu/sfa-go/dfa"
	"github.com/y-yu/sfa-go/utils"
	"sync"
)

type SFA struct {
	I      common.State // initial state
	F      utils.Set
	Rules  dfa.RuleMap
	States StateStateMap
}

type stateStateMapPair struct {
	state    common.State
	stateMap StateMap
}

func NewSFA(d dfa.DFA) SFA {
	allState := d.AllStates()
	fmt.Printf("allState: %v\n", allState)

	fi := make(StateMap)
	for _, state := range allState {
		fi[state] = state
	}
	queue := []stateStateMapPair{
		{d.I, fi},
	}
	sigma := d.AllSymbol()

	sfaStates := make(StateStateMap)

	rules := make(dfa.RuleMap)
	f := utils.NewSet()
	if d.F.Contains(d.I) {
		f.Add(d.I)
	}

	for len(queue) != 0 {
		fromState := queue[0].state
		fromStateMap := queue[0].stateMap
		queue = queue[1:]

		sfaStates[fromState] = fromStateMap

		fmt.Printf("sfaStates: %v\n", sfaStates)

		for _, c := range sigma {
			fmt.Printf("c: %s\n", string(c))
			fnext := make(StateMap)

			for _, state := range allState {
				previous := fromStateMap[state]
				arg := common.NewRuleArgs(previous, c)
				fnext[state] = d.Rules[arg]
			}

			foundState, found := sfaStates.FindState(fnext)
			fmt.Printf("from %v, fnext %v, found: %t, found state %v\n", fromStateMap, fnext, found, foundState)

			var state common.State
			if !found {
				state = common.NewState(len(sfaStates))
				if d.F.Contains(fnext[d.I]) {
					f.Add(state)
				}
				queue = append(queue, stateStateMapPair{state, fnext})
				sfaStates[state] = fnext
			} else {
				state = foundState
			}

			rules[common.NewRuleArgs(fromState, c)] = state
		}
	}

	return SFA{
		I:      d.I,
		F:      f,
		Rules:  rules,
		States: sfaStates,
	}
}

func (s *SFA) ToDFA() dfa.DFA {
	return dfa.DFA{
		I:     s.I,
		F:     s.F,
		Rules: s.Rules,
	}
}

func (s *SFA) Match(str string, p int) bool {
	subStrings := make([]string, p)
	index := 0
	subLength := len(str) / p
	for i := range p - 1 {
		subStrings[i] = str[index : index+subLength]
		index += subLength
	}
	subStrings[p-1] = str[index:]

	wg := sync.WaitGroup{}
	curs := make([]common.State, p)
	//mu := sync.Mutex{}
	for i := range p {
		wg.Add(1)

		func(i int) {
			defer wg.Done()

			cur := s.I
			for _, c := range []rune(subStrings[i]) {
				key := common.NewRuleArgs(cur, c)
				next, _ := s.Rules[key]
				cur = next
			}
			//mu.Lock()
			curs[i] = cur
			//mu.Unlock()

			fmt.Printf("i %d\n", i)
		}(i)
	}
	wg.Wait()

	f := s.I
	for _, cur := range curs {
		f = s.States[cur][f]
	}

	return s.F.Contains(f)
}
