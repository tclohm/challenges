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

func longer(s string) int {
	if len(s) <= 1 { return len(s) }
	var seen = map[string]int{}
	var left, longest = 0, 0

	for right := 0 ; right < len(s) ; right++ {
		var current = string(s[right])
		var prevSeen = seen[current]
		// have seen this character previous and is it greater or equal to left
		// if so increment left
		// and add right pointer value in to the map
		if prevSeen >= left {
			left = prevSeen + 1
		}
		seen[current] = right
		// check if longest is longer than right pointer - left pointer + 1
		longest = max(longest, right - left + 1)
	}

	return longest
}

func main() {
	str := "abccabb"

	fmt.Println(longest(str))
	fmt.Println(longer(str))
}