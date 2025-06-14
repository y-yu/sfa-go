package main

import (
	"fmt"
	"github.com/y-yu/sfa-go/dfaregex"
	"github.com/y-yu/sfa-go/sfa"
)

func main() {
	regex := "(ab)*"
	fmt.Printf("regex: %s\n", regex)
	re := dfaregex.Compile(regex)

	d := re.GetDFA()
	fmt.Printf("dfa states: %v\n", d)

	s := sfa.ToSFA(d)
	fmt.Printf("dfa states: %v\n", s)

	sfa.SFA2dot(s, "abast")

	for _, s := range []string{"aaa", "ababab", "babababb"} {
		if re.Match(s) {
			fmt.Printf("%s\t=> matched.\n", s)
		} else {
			fmt.Printf("%s\t=> NOT matched.\n", s)
		}
	}
}
