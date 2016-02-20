package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var ds = flag.Int("ds", 256, "digest size")

func main() {
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	// Read a single line of standard input.
	input, isPrefix, err := reader.ReadLine()
	if isPrefix || err != nil {
		fmt.Println("Error: problem reading standard input")
		os.Exit(1)
	}
	switch *ds {
	case 256:
		c1 := sha256.Sum256([]byte(input))
		fmt.Printf("%x\n", c1)
	case 384:
		c1 := sha512.Sum384([]byte(input))
		fmt.Printf("%x\n", c1)
	case 512:
		c1 := sha512.Sum512([]byte(input))
		fmt.Printf("%x\n", c1)
	default:
		fmt.Println("Error: invalid or unsupported digest size")
		os.Exit(1)
	}
}
