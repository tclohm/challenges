package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func charReplace(str string, k int) int {
	// store characters
	charMap := make(map[rune]int)
	start, maxCounter, result := 0, 0, 0

	// loop through the length of the string with the main shifter and start as pivot
	for endIndex, endValue := range str {
		// check if end value in dictionary
		if _, ok := charMap[endValue]; ok {
			charMap[endValue] += 1
		} else {
			charMap[endValue] = 1
		}

		// max counter finds the longest streak
		maxCounter = max(maxCounter, charMap[endValue])

		// check if the replacement exceeds k
		if endIndex - start + 1 - maxCounter > k {
			// if it exceeds, you dont move end, and you move start instead
			// if replacement is less than k, you continue moving end
			startChar := rune(str[start])
			charMap[startChar] -= 1
			start += 1
		}
		result = max(endIndex - start + 1, result)
	}
	return result
}

func main() {
	s1 := "ABAB"
	s2 := "AABABBA"

	fmt.Println(charReplace(s1, 2)) // output: 4
	fmt.Println(charReplace(s2, 1)) // output: 4
}