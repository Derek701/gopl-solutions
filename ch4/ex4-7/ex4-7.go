package main

import "fmt"

func main() {
	a := []byte("0 I 2 3 4 Π")
	reverse(a)
	fmt.Println(string(a)) // "Π 4 3 2 I 0"
}

// reverse reverses the characters of a slice.
func reverse(bytes []byte) {
	runes := []rune(string(bytes))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	copy(bytes, []byte(string(runes)))
}
