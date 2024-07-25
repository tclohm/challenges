package main

import "fmt"

func isAnagram(s, t string) bool {
	ht := map[rune]int{}
	for _, r := range s {
		ht[r]++
	}

	for _, r := range t {
		if ht[r] < 0 { return false }
		ht[r]--
		if ht[r] == 0 { delete(ht, r) }
	}

	if len(ht) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isAnagram("racecar", "carrace"))
	fmt.Println(isAnagram("jar", "jam"))
}
