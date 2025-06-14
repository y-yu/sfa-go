package common

// RuleArg is a key for the map as transition function of d.
type RuleArg struct {
	From State // from state
	C    rune  // input symbol
}

// NewRuleArgs returns a new RuleArg.
func NewRuleArgs(from State, in rune) RuleArg {
	return RuleArg{
		From: from,
		C:    in,
	}
}
