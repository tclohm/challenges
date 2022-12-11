package main

import (
	"fmt"
	"unicode"
)

func strip(s string) string {
	stripped := ""
	for _, r := range s {
		if unicode.IsLetter(r) {
			stripped += string(unicode.ToLower(r))
		} 
	}
	return stripped
}

func isPalindrome(s string, left, right *int) bool {
	for *left < *right {
		if s[*left] != s[*right] {
			return false
		}
		*left++
		*right--
	}
	return true
}

func isAlmostPalindrome(s string) bool {
	scrubbed := strip(s)
	var left, right = 0, len(scrubbed) - 1
	if (isPalindrome(scrubbed, &left, &right)) {
		return true
	}

	for (left < right) {
		if scrubbed[left] != scrubbed[right] {
			leftPlus := left+1
			rightMinus := right-1
			return isPalindrome(scrubbed, &leftPlus, &right) || isPalindrome(scrubbed, &left, &rightMinus)
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isAlmostPalindrome(s))

	s1 := "race a car"
	fmt.Println(isAlmostPalindrome(s1))
}