package main

import "fmt"

func characterReplacement(s string, k int) int {
	n := len(s)
	if n == 0 || n == 1 {
		return n
	}

	count := make(map[byte]int)
	longest := 0

	left := 0

	for right, _ := range s {
		b := s[right]
		count[b]++

		window := (right - left + 1)

		if window - count[b] > k {
			// remove left pointer
			lb := s[left]
			count[lb]--
			left++
		}

		window = (right - left + 1)

		if longest < window {
			longest = window
		}
	}
	return longest
}

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
}