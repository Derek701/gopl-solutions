package echo

import "testing"

// -- Benchmarks --

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3()
	}
}

// Intel Core i5-4200U @ 1.60GHz
// $ go test -bench=.
// BenchmarkEcho1-4         5000000               283 ns/op
// BenchmarkEcho2-4         5000000               287 ns/op
// BenchmarkEcho3-4         5000000               248 ns/op
