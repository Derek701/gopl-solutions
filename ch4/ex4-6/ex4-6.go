package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// squashSpace returns a slice with adjacent spaces squashed.
// The underlying array is modified during the call.
func squashSpace(bytes []byte) []byte {
	var r rune
	var l, pos int
	out := bytes[:0]       // zero-length slice of original
	isRunSquashed := false // Is the current run of spaces squashed?
	for range string(bytes) {
		r, l = utf8.DecodeRune(bytes[pos:])
		if unicode.IsSpace(r) {
			if isRunSquashed {
				pos += l
				continue
			} else {
				out = append(out, ' ')
				isRunSquashed = true
			}
		} else {
			for i := 0; i < l; i++ {
				out = append(out, bytes[pos+i])
			}
			isRunSquashed = false
		}
		pos += l
	}
	return out
}

func main() {
	data := []byte(" \n I \r 2 \t 3 \v 4  Π ")
	fmt.Printf("%q\n", squashSpace(data)) // " I 2 3 4 Π "
}
