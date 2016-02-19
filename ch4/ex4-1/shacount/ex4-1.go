package shacount

import "github.com/Derek701/gopl-solutions/ch2/ex2-3/popcount"

// DiffBits returns the number of bits that are different in two SHA256 hashes.
func DiffBits(c1, c2 [32]byte) int {
	var count int
	for i := range c1 {
		count += popcount.PopCount(uint64(c1[i] ^ c2[i]))
	}
	return count
}
