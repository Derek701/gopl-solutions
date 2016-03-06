package popcount

// ByClearing returns the population count (number of set bits) of x.
func ByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}
