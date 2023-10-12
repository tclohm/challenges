package main

import "fmt"

func longestPalindrome(s string) string {
	longest := string(s[0])
	for i := 0 ; i <= len(s) ; i++ {
		for k := i ; k <= len(s) ; k++ {
			if isPalindrome(string(s[i:k])) {
				longest = max(longest, string(s[i:k]))
			}
			
		}
	}
	
	return longest
}

func isPalindrome(s string) bool {
	start, end := 0, len(s) - 1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func max(a, b string) string {
	if len(a) >= len(b) { return a }
	return b
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}