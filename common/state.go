// Package utils contains utility types and functions for dfa-regex-engine.
package common

import "fmt"

// State represents a state including in DFA or NFA.
// It has its number. The number can NOT be duplicate in same DFA or NFA.
// Basically, the number is set incrementally.
type State int

// NewState returns a new state with its number set.
func NewState(n int) State {
	return State(n)
}

func (s State) String() string {
	return fmt.Sprintf("q%d", s)
}
