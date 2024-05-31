package main

import "fmt"

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	left, right := 0, 0
	totalCost := 0
	maxLength := 0

	for right < n {
		totalCost += abs(int(s[right]) - int(t[right]))

		for totalCost > maxCost {
			totalCost -= abs(int(s[left]) - int(t[left]))
			left += 1
		}

		maxLength = max(maxLength, right - left + 1)
		right += 1
	}

	return maxLength
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
	fmt.Println(equalSubstring("abcd", "cdef", 3))
	fmt.Println(equalSubstring("abcd", "acde", 0))
}
