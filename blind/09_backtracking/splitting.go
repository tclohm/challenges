package main

import (
	"fmt"
	"strconv"
)

func splitString(s string) bool {
	if s == "" { return false }

	var backtrack func (index int, previousValue int) bool
	backtrack = func (index int, previousValue int) bool {
		if index == len(s) { return true }
		for i := index ; i < len(s) ; i++ {
			currentValue, err := strconv.Atoi(string(s[index:i + 1]))
			if err != nil { return false }
			if previousValue - 1 == currentValue && backtrack(i + 1, currentValue) {
				return true
			}
		}
		return false
	}

	// comparing two values
	for i := 0 ; i < len(s) - 1 ; i++ {
		val, err := strconv.Atoi(string(s[:i+1]))
		if err != nil { return false }
		if backtrack(i + 1, val) {
			return true
		}
	}
	return false
}

func main() {
	// s splits into two or more non-empty substrings
	// numbers values are descending
	// diff numbers are by one
	fmt.Println(splitString("1234"))
	fmt.Println(splitString("050043"))
	fmt.Println(splitString("9080701"))
}