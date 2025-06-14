// Package nfarule implements the transition function of NFA.
package nfa

import (
	"fmt"
	"github.com/y-yu/sfa-go/common"
	"github.com/y-yu/sfa-go/utils"
	"reflect"
)

// RuleMap represents a transition function of NFA.
// The key is a pair like "(from state, input symbol)".
// The value is a set of transition destination states
// when "input symbol" is received in "from state".
type RuleMap map[common.RuleArg]utils.Set

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
