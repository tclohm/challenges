package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func does(r rune, existIn []rune) bool {
	for _, char := range existIn {
		if r == char {
			return true
		}
	}
	return false
}

func length_of_longest_substring(s string) int {
	var set = []rune{}
	var left int = 0
	
	var result int = 0

	for right, char := range s {
		for does(char, set) {
				length := len(set)
				set = set[1:length]
				left += 1
		}
		result = max(result, right - left + 1)
		set = append(set, char)
	}
	return result
}

func main() {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"

	fmt.Println(length_of_longest_substring(s1))
	fmt.Println(length_of_longest_substring(s2))
	fmt.Println(length_of_longest_substring(s3))
}