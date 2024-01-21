package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var result []string
	var phonePad = map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}


	var backtrack func(combination, digits string)
	backtrack = func(combination string, digits string) {
		
		if digits == "" {
			result = append(result, combination) 
			return
		}

		var button = digits[0]
		var pressed = phonePad[button]

		for _, letter := range pressed {
			backtrack(combination + string(letter), digits[1:])
		}

	}

	backtrack("", digits)
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("99"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("5"))
	fmt.Println(letterCombinations("2568"))
}