package main

import "fmt"

// dedupAdj returns a slice with adjacent duplicates removed.
// The underlying array is modified during the call.
func dedupAdj(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	dup := ""
	for _, s := range strings {
		if s != dup {
			out = append(out, s)
			dup = s
		}
	}
	return out
}

func main() {
	data := []string{"one", "one", "three"}
	fmt.Printf("%q\n", dedupAdj(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
}
