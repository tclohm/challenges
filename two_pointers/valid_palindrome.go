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

	str := cleanup(s)
	
	start := 0
	end := len(str) - 1

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