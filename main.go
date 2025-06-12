package main

import (
	"fmt"
	"github.com/y-yu/sfa-go/dfaregex"
)

func main() {
	regex := "(ab)*"
	fmt.Printf("regex: %s\n", regex)
	re := dfaregex.Compile(regex)

	fmt.Printf("dfa states: %v\n", re.GetDFA())

	for _, s := range []string{"aaa", "ababab", "babababb"} {
		if re.Match(s) {
			fmt.Printf("%s\t=> matched.\n", s)
		} else {
			fmt.Printf("%s\t=> NOT matched.\n", s)
		}
	}
}
