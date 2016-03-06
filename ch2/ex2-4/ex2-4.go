package popcount

// ByShifting returns the population count (number of set bits) of x.
func ByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}
