To benchmark and create profiles, run:
- go test -bench . -benchmem -cpuprofile prof.cpu -memprofile prof.mem        

To pprof into the cpu profile:
-  go tool pprof profiling_strings_concatenation.test prof.cpu 

To view the flamegraph and more for the cpu profile:
- go tool pprof -http ":8000" prof.cpu

To do so for memory, just replace ".cpu" with ".mem"

Benchmark results:

go test -bench . -benchmem -cpuprofile prof.cpu -memprofile prof.mem

goos: darwin
goarch: arm64
pkg: profiling_strings_concatenation
BenchmarkConcatStringsPlusEqual-10                          2583            466030 ns/op         2684247 B/op        999 allocs/op
BenchmarkConcatStringsPlusEqualImproved-10                247645              4918 ns/op           10752 B/op          2 allocs/op
BenchmarkConcatStringsBytesBuffer-10                      144388              7968 ns/op           21696 B/op          9 allocs/op
BenchmarkConcatStringsBytesBufferImproved-10              245060              5199 ns/op           10752 B/op          2 allocs/op
PASS
ok      profiling_strings_concatenation 5.263s