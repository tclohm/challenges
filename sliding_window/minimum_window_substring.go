package main

import "fmt"

func minWindow(s, t string) string {
	if len(s) < len(t) || len(s) == 0 {
		return ""
	}

	// map to track
	count := make(map[byte]int, len(t))
	for _, char := range []byte(t) {
		count[char]++
	}

	var (
		minimum string
		minLen = 0
		left, right = 0, 0
	)

	addExtra(s[0], count)

	// sliding window
	for left <= right {

		if allMatching(count) {

			if minLen == 0 || right - left < minLen {

				minLen = right - left + 1
				minimum = s[left:right+1]

			}

			removeExt(s[left], count)
			left++

		} else {

			right++
			if right == len(s) { break }

			addExtra(s[right], count)

		}
	}

	return minimum
}

// Add a byte. The byte might not be one of s
func addExtra(s byte, dict map[byte]int) {
	if val, ok := dict[s]; ok {
		dict[s] = val - 1
	}
}

// Remove a byte.
func removeExt(s byte, dict map[byte]int) {
	if val, ok := dict[s]; ok {
		dict[s] = val + 1
	}
}

func allMatching(dict map[byte]int) bool {
	for _, val := range dict {
		if val > 0 { return false }
	}
	return true
}

func main() {

}