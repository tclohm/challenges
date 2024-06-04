package main

import "fmt"

func longestPalindrome(s string) int {
	ht := map[rune]bool{}
	res := 0
	for _, r := range s {
		if ht[r] {
			delete(ht, r)
			res += 2
		} else {
			ht[r] = true
		}
	}

	if len(ht) > 0 {
		return res + 1
	} else {
		return res
	}
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
