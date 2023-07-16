package main

import "fmt"

func partition(s string) [][]string {
	result := make([][]string, 0)
	current := make([]string, 0)

	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(s) {
			result = append(result, append([]string{}, current...))
		}

		for i := index ; i < len(s) ; i++ {
			if isPalindrome(s[index : i + 1]) {
				current = append(current, s[index:i + 1])
				backtrack(i + 1)
				current = current[:len(current) - 1]
			}
		}
	}
	backtrack(0)
	return result
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
	fmt.Println(partition("racecar"))
}