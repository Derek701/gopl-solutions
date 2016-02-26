package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	rotate(s[:], -2)
	fmt.Println(s) // "[2 3 4 5 0 1]"
}

// rotate rotates a slice of ints in place by p elements to the right or left.
func rotate(s []int, p int) {
	if len(s) == 0 || p == 0 { // check for no rotation
		return
	}
	p %= len(s) // skip unnecessary complete rotations and keep p in-bounds
	if p < 0 {
		p += len(s) // convert left rotation to the equivalent right rotation
	}
	tmp := make([]int, len(s))
	copy(tmp, s[len(s)-p:])
	copy(tmp[p:], s[:len(s)-p])
	copy(s, tmp)
}
