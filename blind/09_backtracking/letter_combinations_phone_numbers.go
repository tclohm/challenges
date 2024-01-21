package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" { return []string{} }
	var list []string
	var phone = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	var backtrack func(combination, nextDigit string)
	backtrack = func(combination, nextDigit string) {
		if nextDigit == "" {
			list = append(list, combination)
		} else {
			// rune
			letters := phone[nextDigit[0]-'2']
			for _, letter := range letters {
				backtrack(combination + string(letter), nextDigit[1:])
			}
		}
	}
	backtrack("", digits)
	return list	
}

func main() {
	fmt.Println(letterCombinations("23"))
}