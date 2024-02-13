package main

import "fmt"

func longestSubstring(s string) string {
	var result string
	var length int

	for i, _ := range s {
		// odd
		var left, right = i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right - left + 1 > length {
				result = s[left:right+1]
				length = right - left + 1
			}
			left -= 1
			right += 1
		}
		// even
		left, right = i, i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right - left + 1 > length {
				result = s[left:right+1]
				length = right - left + 1
			}
			left -= 1
			right += 1
		}
	}
	return result
}

func main() {
	fmt.Println(longestSubstring("babad"))
	fmt.Println(longestSubstring("cbbd"))
}