package _7_benchmark

import "testing"

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	for i := 0; i < b.N; i++ {
		ConcatStringByAdd(elems)
	}
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	for i := 0; i < b.N; i++ {
		ConcatStringByBytesBuffer(elems)
	}
}


/**
	```bash
	$ go test -bench=.
	goos: linux
	goarch: amd64
	pkg: golang-dev/example.com/37-benchmark
	cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
	BenchmarkConcatStringByAdd-8            12891658                93.23 ns/op
	BenchmarkConcatStringByBytesBuffer-8    21851823                57.60 ns/op
	PASS
	ok      golang-dev/example.com/37-benchmark     2.615s


	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: golang-dev/example.com/37-benchmark
	cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
	BenchmarkConcatStringByAdd-8            11808808                90.52 ns/op           16 B/op          4 allocs/op
	BenchmarkConcatStringByBytesBuffer-8    17503350                59.62 ns/op           69 B/op          2 allocs/op
	PASS
	ok      golang-dev/example.com/37-benchmark     2.288s
	```
 */