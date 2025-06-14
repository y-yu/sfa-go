package common

// RuleArgs is a key for the map as transition function of d.
type RuleArgs struct {
	From State // from state
	C    rune  // input symbol
}

// NewRuleArgs returns a new RuleArgs.
func NewRuleArgs(from State, in rune) RuleArgs {
	return RuleArgs{
		From: from,
		C:    in,
	}
}
