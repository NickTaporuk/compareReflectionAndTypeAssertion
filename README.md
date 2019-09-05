### RUN
    go test -bench=.
### Go Version
    go version go1.12.5 darwin/amd64
### Comparison
    goos: darwin
    goarch: amd64
    pkg: github.com/NickTaporuk/interfaceToMap
    with loop for
        BenchmarkConvertInterfaceToMapStrings-12        20000000                63.7 ns/op
        BenchmarkInterfaceToMap-12                      20000000                80.5 ns/op
        
    only convert to structure map[string]interface
        BenchmarkConvertInterfaceToMapStrings2-12       2000000000               0.31 ns/op
        BenchmarkInterfaceToMap2-12                     100000000               17.5 ns/op

    PASS
    ok      github.com/NickTaporuk/interfaceToMap   5.399s


### Conclusion

    method BenchmarkConvertInterfaceToMapStrings work faster then reflection

