package main

import (
	"fmt"
	"unicode"
	"strings"
)

func cleanup(s string) string {
	var ns string = ""
	for _, r := range s {
		if unicode.IsLetter(r) {
			ns += string(r)
		}
	}

	lowered := strings.ToLower(ns)
	return lowered
}

func isPalindrome(s string) bool {
	// create a new string with only the characters
	str := cleanup(s)
	// pointers for start and end
	start := 0
	end := len(str) - 1
	// while start is less than end, check the characters in positions to be the same
	// if not the same return false
	// increment and decrement pointers
	// return true at the end 
	for start < end  {
		first := str[start]
		last := str[end]

		if first != last {
			return false
		}

		start++
		end--
	}

	return true
}

func main() {
	s1 := "A man, a plan, a canal: Panama"
	s2 := "race a car"

	fmt.Println(isPalindrome(s1))
	fmt.Println(isPalindrome(s2))
}