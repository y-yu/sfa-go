// Package dfaregex provides a DFA regex engine(DFA engine).
package dfaregex

import (
	"github.com/y-yu/sfa-go/dfa"
	"github.com/y-yu/sfa-go/nfa2dfa"
	"github.com/y-yu/sfa-go/node"
	"github.com/y-yu/sfa-go/parser"
)

// Regexp has a DFA and regexp string.
type Regexp struct {
	regexp string
	d      *dfa.DFA
}

// NewRegexp return a new Regexp.
func NewRegexp(re string) *Regexp {
	psr := parser.NewParser(re)
	ast := psr.GetAST()
	frg := ast.Assemble(node.NewContext())
	nfa := frg.Build()
	d := nfa2dfa.ToDFA(nfa)
	//d.Minimize()

	return &Regexp{
		regexp: re,
		d:      d,
	}
}

// Compile is a wrapper function of NewRegexp().
func Compile(re string) *Regexp {
	return NewRegexp(re)
}

// Match returns whether the input string matches the regular expression.
func (re *Regexp) Match(s string) bool {
	return re.d.Match(s)
}

func (re *Regexp) GetDFA() *dfa.DFA {
	return re.d
}
