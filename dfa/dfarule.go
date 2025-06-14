// Package dfarule implements the transition function of DFA.
package dfa

import (
	"fmt"
	"github.com/y-yu/sfa-go/common"
	"reflect"
)

// RuleMap represents a transition function of d.
// The key is a pair like "(from state, input symbol)".
// The value is a destination state when "input symbol"
// is received in "from state".
type RuleMap map[common.RuleArgs]common.State

func (r RuleMap) String() string {
	s := ""

	keys := reflect.ValueOf(r).MapKeys()
	for i, k := range keys {
		from := k.FieldByName("From").Interface().(common.State)
		c := k.FieldByName("C").Interface().(rune)
		dst := r[common.NewRuleArgs(from, c)]
		s += fmt.Sprintf("%s\t--['%c']-->\t%s", from, c, dst)
		if i+1 < len(keys) {
			s += "\n"
		}
	}
	return s
}
