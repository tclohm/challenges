package main

import "fmt"

func partition(s string) [][]string {
	var list [][]string
	var backtrack func(tmp []string, s string, start int)
	
	backtrack = func(tmp []string, s string, start int) {
		if start == len(s) {
			list = append(list, append([]string{}, tmp...))
		} else {
			for i := start ; i < len(s) ; i++ {
				if isPalindrome(s, start, i) {
					tmp = append(tmp, s[start:i + 1])
					backtrack(tmp, s, i + 1)
					tmp = tmp[:len(tmp) - 1]
				}
			}
		}
	}

	backtrack([]string{}, s, 0)
	return list
}

func isPalindrome(s string, low, high int) bool {
	for low < high {
		if s[low] != s[high] { return false }
		low++
		high--
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))

}