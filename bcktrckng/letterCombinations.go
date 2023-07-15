package main

import "fmt"

func letterCombination(digits string) []string {
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}

	phoneMap := make(map[byte]string)
	phoneMap['2'] = "abc"
	phoneMap['3'] = "def"
	phoneMap['4'] = "ghi"
	phoneMap['5'] = "jkl"
	phoneMap['6'] = "mno"
	phoneMap['7'] = "pqrs"
	phoneMap['8'] = "tuv"
	phoneMap['9'] = "wzyz"
	current := ""

	var backtrack func(index int)
	backtrack = func(index int) {
		if len(current) == len(digits) {
			result = append(result, current)
			return
		}
		digit := digits[index]
		str := phoneMap[digit]

		for i := 0 ; i < len(str) ; i++ {
			current += string(str[i])
			backtrack(index + 1)
			current = current[:len(current) - 1]
		}
	}
	backtrack(0)
	return result
}

func main() {
	fmt.Println(letterCombination("23"))
	fmt.Println(letterCombination(""))
	fmt.Println(letterCombination("2"))
}