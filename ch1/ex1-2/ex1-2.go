package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i+1, arg) // Use i+1 to adjust for os.Args[1:]
	}
}
