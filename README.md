### RUN
    go test -bench=.
### Go Version
    go version go1.12.5 darwin/amd64
### Comparison

    goos: darwin
    goarch: amd64
    pkg: github.com/NickTaporuk/interfaceToMap
    BenchmarkConvertInterfaceToMapStrings-12        20000000                63.3 ns/op
    BenchmarkInterfaceToMap-12                      20000000                80.5 ns/op
    PASS
    ok      github.com/NickTaporuk/interfaceToMap   3.038s

### Conclusion

method BenchmarkConvertInterfaceToMapStrings work faster then reflection

