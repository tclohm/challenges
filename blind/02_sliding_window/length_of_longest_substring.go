package main

import "fmt"

func length_of_longest_substring(s string) int {
	set := make(map[byte]bool)
	left, longest := 0, 0

	for right, _ := range s {
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