package main

import "fmt"

const aLen = 6 // Define the array size for function reverse.

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
}

// reverse reverses an array of ints in place.
func reverse(aPtr *[aLen]int) {
	for i, j := 0, len(aPtr)-1; i < j; i, j = i+1, j-1 {
		aPtr[i], aPtr[j] = aPtr[j], aPtr[i]
	}
}
