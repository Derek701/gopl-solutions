// Comma prints its argument numbers with a comma at each power of 1000.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a decimal integer or floating-point string.
func comma(s string) string {
	var buf bytes.Buffer
	var dec string
	n := len(s)
	if s[0] == '-' || s[0] == '+' {
		n--
		buf.WriteByte(s[0])
		s = s[1:]
	}
	if dot := strings.LastIndexByte(s, '.'); dot >= 0 {
		n -= len(s[dot:])
		dec = s[dot:]
		s = s[:dot]
	}
	commaPos := n % 3
	if commaPos == 0 {
		commaPos = 3
	}
	for i := 0; i < n; i++ {
		if i == commaPos {
			buf.WriteByte(',')
			commaPos += 3
		}
		buf.WriteByte(s[i])
	}
	return buf.String() + dec
}
