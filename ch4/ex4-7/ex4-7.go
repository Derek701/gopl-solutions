package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := []byte("0 I 2 3 4 Π")
	reverse(a)
	fmt.Println(string(a)) // "Π 4 3 2 I 0"
}

// reverse reverses a slice of characters in place.
func reverse(bytes []byte) {
	var r1, r2 rune
	var l1, l2 int
	for i, j := 0, len(bytes); i < j; i, j = i+l2, j-l1 {
		r1, l1 = utf8.DecodeRune(bytes[i:])
		r2, l2 = utf8.DecodeLastRune(bytes[:j])
		if l2 > l1 {
			copy(bytes[i+l2:], bytes[i+l1:j-l2])
		}
		copy(bytes[i:], []byte(string(r2)))
		copy(bytes[j-l1:], []byte(string(r1)))
	}
}
