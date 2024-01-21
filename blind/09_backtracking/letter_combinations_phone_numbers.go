package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" { return []string{} }
	var list []string
	var phone = map[byte]string{'2':"abc", '3':"def", '4':"ghi", '5':"jkl", '6':"mno", '7':"pqrs", '8':"tuv", '9':"wxyz"}

	var backtrack func(combination, digit string)
	backtrack = func(combination, digit string) {
		if digit == "" {
			list = append(list, combination)
		} else {
			button := digit[0]
			letters := phone[button]
			for _, letter := range letters {
				// next digit for combo
				backtrack(combination + string(letter), digit[1:])
			}
		}
	}
	backtrack("", digits)
	return list	
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("99"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("5"))
}