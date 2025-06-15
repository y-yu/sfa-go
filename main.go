package main

import (
	"fmt"
	"github.com/y-yu/sfa-go/dfa"
	"github.com/y-yu/sfa-go/dfaregex"
	"github.com/y-yu/sfa-go/sfa"
	"os"
	"regexp"
	"time"
)

func main() {
	regex := "(ab|abab)*X"
	fmt.Printf("regex: %s\n", regex)
	re := dfaregex.Compile(regex)

	d := re.GetDFA()
	s := sfa.NewSFA(*d)

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
	bytes, err := os.ReadFile("./testdata/abab.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("target string length: %d byte\n\n", len(bytes))

	goregex, _ := regexp.Compile(regex)

	measure("Go regex\t", func() bool {
		return goregex.MatchString(string(bytes))
	})

	measure("DFA\t\t", func() bool {
		return d.Match(string(bytes))
	})

	p := 1
	measure(fmt.Sprintf("SFA(parallel: %d)", p), func() bool {
		return s.Match(string(bytes), p)
	})

	p = 20
	measure(fmt.Sprintf("SFA(parallel: %d)", p), func() bool {
		return s.Match(string(bytes), p)
	})
}

func measure(name string, f func() bool) {
	startAt := time.Now()
	result := f()
	fmt.Printf("%s \t time:\t %d,\t result: \t %t\n", name, time.Now().UnixMicro()-startAt.UnixMicro(), result)
}
