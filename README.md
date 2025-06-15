Go SFA[^sfa] based Regular Expression Matching
=================================================

## Benchmark

```console
$ go run main.go
regex: (ab|abab)*X
target string length: 1274880 byte

Go regex                 time:   51082,  result:         false
DFA                      time:   20503,  result:         false
SFA(parallel: 1)         time:   20550,  result:         false
SFA(parallel: 20)        time:   4762,   result:         false
```

[^sfa]: https://arxiv.org/abs/1405.0562