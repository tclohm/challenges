package main

import "fmt"

// smaller, longer
func getLengths(a, b string) (string, string) {
	if len(a) <= len(b) {
		return a, b
	} 
	return b, a
}

func lcs(s1, s2 string) int {
	var smaller, longer = getLengths(s1, s2)
	var prev int
	var counter int = 0

	for i := 0 ; i < len(smaller) ; i++ {
		prev = i
		for j := i ; j < len(longer) ; j++ {
			if smaller[i] == longer[j] && j >= prev {
				counter++
			}
		}
	}

	return counter
}

func main() {
	fmt.Println(lcs("abcde", "ace"))
	fmt.Println(lcs("abc", "abc"))
	fmt.Println(lcs("abc", "def"))
	fmt.Println(lcs("eb", "def"))
	fmt.Println(lcs("qwerty", "dqewferddty"))
}