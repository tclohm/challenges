package main

import (
	"fmt"
	"strings"
)

func findDifferentBinaryString(nums []string) string {
	//var set = make(map[byte]bool)
	var result string
	var backtrack func (index int, current []string) string
	backtrack = func (index int, current []string) string {
		if index == len(nums) {
			result = strings.Join(current, ",")
			return result
		}

		result = backtrack(index + 1, current)
		if result != "" {
			return result
		}

		current[index] = "1"
		result = backtrack(index + 1, current)
		if result != "" {
			return result
		}
		return ""
	}
	return ""
}

func main() {
	fmt.Println(findDifferentBinaryString([]string{"01","10"}))
	fmt.Println(findDifferentBinaryString([]string{"00","01"}))
	fmt.Println(findDifferentBinaryString([]string{"111","011","001"}))
}