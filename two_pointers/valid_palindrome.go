package main

import (
	"fmt"
	"unicode"
	"strings"
)

func isPalindrome(s string) bool {
	var ns string = ""
	for _, r := range s {
		if unicode.IsLetter(r) {
			ns += string(r)
		}
	}

	fmt.Println(strings.ToLower(ns))
	return true
}

func main() {
	s1 := "A man, a plan, a canal: Panama"
	s2 := "race a car"

	fmt.Println(isPalindrome(s1))
	fmt.Println(isPalindrome(s2))
}