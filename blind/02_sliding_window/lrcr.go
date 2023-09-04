package main

import "fmt"

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	left_pointer := 0
	window := 0
	mostFrequent := 0
	longest := 0

	for right_pointer, _ := range s {
		b := s[right_pointer]

		count[b] += 1
		mostFrequent = max(mostFrequent, count[b])

		window = right_pointer - left_pointer + 1
		// check if window
		if window - mostFrequent > k {
			count[b]--
			left_pointer++
		}

		window = right_pointer - left_pointer + 1

		longest = max(longest, window)

	}

	return longest
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
}