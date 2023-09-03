package main

import "fmt"

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	longest := 0

	left := 0
	mostFrequent := 0

	for right, _ := range s {
		// byte
		b := s[right]
		// increase the integet by one and check for most frequent
		count[b] += 1
		mostFrequent = max(mostFrequent, count[b]) 
		// in the sliding window, check the length (right - left + 1)
		window := right - left + 1
		// is our window minus our most frequent letter bigger than k?
		if window - mostFrequent > k {
			// if yes, remove the byte count and move left
			count[b] -= 1
			left++
		}
		// then check the to see if our current longest is longer than our sliding window
		window = right - left + 1
		longest = max(longest, window)
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
}