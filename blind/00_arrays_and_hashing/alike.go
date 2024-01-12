package main

import (
	"fmt"
	"strings"
)

func isVowel(str string) bool {
	return str == "a" || str == "e" || str == "i"  || str == "o" || str == "u"
}

// two strings are alike if they have the same number of vowels
func halvesAreAlike(s string) bool {

	under_case := strings.ToLower(s)
	
	first, second := len(under_case) / 2, len(under_case)
	first_half, second_half := under_case[:first], under_case[first:second]

	first_vowels := 0
	second_vowels := 0

	for i := 0 ; i < len(first_half) ; i++ {
		if isVowel(string(first_half[i])){
			first_vowels++
		}
	}

	for i := 0 ; i < len(second_half) ; i++ {
		if isVowel(string(second_half[i])){
			second_vowels++
		}
	}

	return first_vowels == second_vowels
}

func alike(s string) bool {
	half := len(s) / 2
	mark := 0

	for i, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			if i < half {
				mark++
			} else {
				mark--
			}
		}
	}
	return mark == 0
}

func main() {
	fmt.Println(halvesAreAlike("book"))
	fmt.Println(halvesAreAlike("textbook"))
	fmt.Println(halvesAreAlike("textBOOK"))
	fmt.Println(halvesAreAlike("BOok"))
}