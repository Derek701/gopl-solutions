// Wordfreq reports the frequency of each word in an input text file.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: No input text file specified as an argument.")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)   // scan for words (space-separated text)
	counts := make(map[string]int) // counts of each word in input text file
	for input.Scan() {
		word := strings.ToLower(input.Text())
		counts[word]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	file.Close()
	fmt.Print("word\t\tfreq\n----\t\t----\n")
	for s, n := range counts {
		if len(s) > 15 { // condense too-long words to fit table spacing
			s = s[0:12] + "..."
		}
		if len(s) < 8 { // pad too-short words to fit table spacing
			s += "\t"
		}
		fmt.Printf("%s\t%d\n", s, n)
	}
}
