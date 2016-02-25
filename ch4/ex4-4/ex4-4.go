package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	rotate(s[:], 2, false)
	fmt.Println(s) // "[2 3 4 5 0 1]"
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// rotate rotates a slice of ints in place by p elements to the right or left.
func rotate(s []int, p int, r bool) {
	if r {
		reverse(s)
	}
	reverse(s[:p])
	reverse(s[p:])
	if !r {
		reverse(s)
	}
}
