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
		// add right pointer chars to the count
		b := s[right]
		count[b]++
		// grab our window
		window := (right - left + 1)
		// is our window - right pointer count larger than k
		if window - count[b] > k {
			// if it is, remove letters from count using the left pointer and increment
			lb := s[left]
			count[lb]--
			left++
		}
		// update our window
		window = (right - left + 1)
		// check if the longest is less than our window
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