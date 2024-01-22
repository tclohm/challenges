package main

import (
	"fmt"
	"strconv"
)

func splitString(s string) bool {

	var backtrack func(index int, prev int) bool
	backtrack = func(index int, prev int) bool {
		if index == len(s) { return true }

		for i := index ; i < len(s) ; i++ {
			// value : start at index and add i + 1
			v, _ := strconv.Atoi(string(s[index: i + 1]))
			// if our prev - 1 equals value and our backtracking is true, return true
			if prev - 1 == v && backtrack(i + 1, v) {
				return true
			}
		}
		return false
	}

	// starting the run
	for i := 0 ; i < len(s) - 1 ; i++ {
		v, _ := strconv.Atoi(string(s[:i + 1]))
		if backtrack(i + 1, v) {
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