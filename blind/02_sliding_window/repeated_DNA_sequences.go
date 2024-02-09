package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
    // DNA Sequences == 10 letter long ('A', 'C', 'G', 'T')
    // return repeated sequences
    var duplicates, seen, result = make(map[string]bool), make(map[string]bool), []string{}

    for left := range s {
    	if left + 10 < len(s) {
			current := s[left : left + 10]
			if _, ok := seen[current]; ok {
				duplicates[current] = true
			}
			seen[current] = true
    	}
    }

    for k, _ := range duplicates {
    	result = append(result, k)
    }

    return result
}

// AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT
// AAAAACCCCC
// CCCCCAAAAA
// AAAAACCCCC
// CCCCCAAAAA

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}