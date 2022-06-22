package main

import (
	"fmt"
	// "strings"
)

// "eat","tea","tan","ate","nat","bat"
// "eat"


func groupAnagrams(strings []string) [][]string {

	hm := make(map[[26]int][]string)
	var results [][]string

	for _, word := range strings {
		alphabet := build(word)
		if _, ok := hm[alphabet]; ok {
			hm[alphabet] = append(hm[alphabet], word)
		} else {
			var empty []string
			hm[alphabet] = append(empty, word)
		}
	}

	for _, value := range hm {
		results = append(results, value)
	}

	return results
}

// alphabet
func build(str string) [26]int {
	var s [26]int
	for i := 0 ; i < len(str) ; i++ {
		s[str[i] - 'a'] += 1
	}
	return s
}


func main() {
	s1 := []string{"eat","tea","tan","ate","nat","bat"}
	s2 := []string{""}
	fmt.Println(groupAnagrams(s1))
	fmt.Println(groupAnagrams(s2))
}