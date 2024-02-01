package main

import "fmt"

func length_of_longest_substring(s string) int {
	var set = map[byte]bool{}
	var longest = 0
	var left = 0

	for right, _ := range s {
		// keep deleting from the left until right is not true
		for set[s[right]] {
			delete(set, s[left])
			left++
		}
		set[s[right]] = true
		longest = max(longest, right - left + 1)
	}

	return longest
}

func max(a, b int) int {
	if a > b { return a }
	return b
}


func main() {
	fmt.Println(length_of_longest_substring("abcabcbb"))
	fmt.Println(length_of_longest_substring("bbbbb"))
	fmt.Println(length_of_longest_substring("pwwkew"))
}