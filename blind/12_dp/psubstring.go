package main

import "fmt"

func countSubstrings(s string) int {
	length := len(s)
	var palindrome func (int, int) int
	palindrome = func (left, right int) int {
		count := 0
		for left >= 0 && right < length && s[left] == s[right] {
			count++
			left--
			right++
		}
		return count
	}

	count := 0
	for index, _ := range s {
		// even
		count += palindrome(index, index)
		// odd
		count += palindrome(index, index+1)
	}
	return count
}

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}