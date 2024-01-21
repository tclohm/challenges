package main

import "fmt"

func letterCombinations(digits string) []string {
	var combinations []string
	var phone = map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxuz",
	}

	var punched_numbers []string
	for _, r := range digits {
		punched_numbers = append(punched_numbers, phone[r])
	}

	var backtrack func()

	return combinations
}

func main() {
	fmt.Println(letterCombinations("23"))
}