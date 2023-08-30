package main

import (
	"unicode"
	"fmt"
)

func isPalindrome(str string) bool {
	new_string := ""

	for _, r := range str {
		if unicode.IsSpace(r) || unicode.IsPunct(r) { continue }
		new_string += string(unicode.ToLower(r))
	}

	start := 0
	end := len(new_string) - 1

	for start < end {
		if new_string[start] != new_string[end] { return false }
		start++
		end--
	}
	
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))

}