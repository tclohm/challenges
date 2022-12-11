package main

import "fmt"

func max(a, b int) int {
	if a < b { return b }
	return a
}

func longest(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	
	var longest = 0

	for left := 0 ; left < len(s) ; left++ {

		var seen, currentLength = map[string]bool{}, 0

		for right := left ; right < len(s) ; right++ {

			var currentChar = string(s[right])
		
			if !seen[currentChar] {
				currentLength++
				seen[currentChar] = true
				longest = max(longest, currentLength)
			} else {
				break
			}
		}

	}
	return longest
}

func longestPointer(s string) int {}

func main() {
	str := "abccabb"

	fmt.Println(longest(str))
}