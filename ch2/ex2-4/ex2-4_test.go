package popcount

import (
	"testing"

	"github.com/Derek701/gopl-solutions/ch2/ex2-3/popcount"
)

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount2(0x1234567890ABCDEF)
	}
}

func BenchmarkByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByShifting(0x1234567890ABCDEF)
	}
}

// Intel Core i5-4200U @ 1.60GHz
// $ go test -bench=.
// BenchmarkPopCount-4                  200000000         6.97 ns/op
// BenchmarkPopCount2-4                 100000000        18.2 ns/op
// BenchmarkByShifting-4                20000000        118 ns/op
