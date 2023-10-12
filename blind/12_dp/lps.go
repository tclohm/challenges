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

func optimizedPalindrome(s string) string {
	result := ""
	length := 0

	for i, _ := range s {
		// odd
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right - left + 1 > length {
				result = s[left:right + 1]
				length = right - left + 1
			}
			left -= 1
			right += 1
		}
		// even
		left, right = i, i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right - left + 1 > length {
				result = s[left:right + 1]
				length = right - left + 1
			}
			left -= 1
			right += 1
		}
	}

	return result
}

func main() {
	fmt.Println(optimizedPalindrome("babad"))
	fmt.Println(optimizedPalindrome("cbbd"))
}