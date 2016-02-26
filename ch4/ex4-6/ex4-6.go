package main

import (
	"fmt"
	"unicode"
)

// squashSpace returns a slice with adjacent spaces squashed.
// The underlying array is modified during the call.
func squashSpace(bytes []byte) []byte {
	runes := []rune(string(bytes))
	out := runes[:0]
	isRunSquashed := false // Is the current run of spaces squashed?
	for _, r := range runes {
		if unicode.IsSpace(r) {
			if isRunSquashed {
				continue
			} else {
				out = append(out, ' ')
				isRunSquashed = true
			}
		} else {
			out = append(out, r)
			isRunSquashed = false
		}
	}
	return []byte(string(out))
}

func main() {
	data := []byte(" \n I \r 2 \t 3 \v 4  Π ")
	fmt.Printf("%q\n", squashSpace(data)) // " I 2 3 4 Π "
}
