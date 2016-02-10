// Package anagram provides a function to check for anagrams.
package anagram

import (
	"reflect"
	"strings"
	"unicode"
)

// IsAnagram reports whether two strings are anagrams of each other.
func IsAnagram(s1, s2 string) bool {
	s1, s2 = strings.TrimSpace(s1), strings.TrimSpace(s2)
	if s1 == s2 {
		return false // A string is not an anagram of itself.
	}
	counts1 := make(map[rune]int)
	for _, r := range s1 {
		if unicode.IsLetter(r) {
			counts1[unicode.ToLower(r)]++
		}
	}
	counts2 := make(map[rune]int)
	for _, r := range s2 {
		if unicode.IsLetter(r) {
			counts2[unicode.ToLower(r)]++
		}
	}
	return reflect.DeepEqual(counts1, counts2)
}
