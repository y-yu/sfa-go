Go SFA[^sfa] based Regular Expression Matching
=================================================

## Benchmark

```console
$ go run main.go
regex: (ab|abab)*X
target string length: 1274880 byte

Go regex                 time:   76566,  result:         false
DFA                      time:   22925,  result:         false
SFA(parallel: 1)         time:   20554,  result:         false
SFA(parallel: 20)        time:   4589,   result:         false
```

[^sfa]: https://doi.org/10.1109/ICPP.2013.3