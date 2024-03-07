package main

import "fmt"

func minimumLength(s string) int {
	left, right := 0, len(s) - 1
	for left < right && s[left] == s[right] {
		tmp := s[left]
		for left <= right && s[left] == tmp {
			left += 1
		}
		for left <= right && s[right] == tmp {
			right -= 1
		}
	}
	return right - left + 1
}

func main() {
	fmt.Println(minimumLength("ca"))
	fmt.Println(minimumLength("cabaabac"))
	fmt.Println(minimumLength("aabccabba"))
}