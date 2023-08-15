package main

import "fmt"

func palindromic_substrings(s string) int {
	res := 0

	for i, _ := range s {
		l, r := i, i
		// odd
		res += helper(s, l, r)
		// even
		res += helper(s, l, r+1)
	}
	return res
}

func helper(s string, left, right int) int {
	res := 0
	l, r := left, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		res += 1
		l -= 1
		r += 1
	}
	return res
} 

func main() {
	fmt.Println(palindromic_substrings("abc"))
	fmt.Println(palindromic_substrings("aaa"))
}