package main

import (
	"fmt"
	"github.com/y-yu/sfa-go/dfa"
	"github.com/y-yu/sfa-go/dfaregex"
	"github.com/y-yu/sfa-go/sfa"
	"time"
)

func main() {
	regex := "(aa|a)*X"
	fmt.Printf("regex: %s\n", regex)
	re := dfaregex.Compile(regex)

	d := re.GetDFA()
	s := sfa.NewSFA(*d)

	fmt.Printf("sfa: %s\n", s)

	dfa.DFA2dot(*d, "dfa")
	dfa.DFA2dot(s.ToDFA(), "sfa")
	/*
		dfa.DFA2dot(s.ToDFA(), "ab_ast")

		fmt.Println("\n\n")

		for _, str := range []string{"aaa", "ababab", "babababb"} {
			if re.Match(str) {
				fmt.Printf("%s\t=> matched.\n", str)
			} else {
				fmt.Printf("%s\t=> NOT matched.\n", str)
			}

			if s.Match(str, 3) {
				fmt.Printf("SFA: %s\t=> matched.\n", str)
			} else {
				fmt.Printf("SFA: %s\t=> NOT matched.\n", str)
			}
		}
	*/

	/*
		bytes, err := os.ReadFile("str.txt")
		if err != nil {
			panic(err)
		}

		fmt.Println("\n")

		goregex, _ := regexp.Compile(regex)

		measure("Go regex", func() bool {
			return goregex.MatchString(string(bytes))
		})

		measure("DFA", func() bool {
			return re.Match(string(bytes))
		})

		measure("SFA", func() bool {
			return s.Match(string(bytes), 1)
		})
	*/
}

func measure(name string, f func() bool) {
	startAt := time.Now()
	result := f()
	fmt.Printf("%s time:\t %d, result: \t %t\n", name, time.Now().UnixMilli()-startAt.UnixMilli(), result)
}
