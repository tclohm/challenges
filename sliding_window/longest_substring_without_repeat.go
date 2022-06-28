package main

import "fmt"

func length_of_longest_substring(s string) string {
	var left, right int = 0, 1
	var longest string = ""

	for right < len(s) {
		if s[left] == s[right] {
			left = right
		} else {

		}
		right += 1
	}
}

func main() {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"
}