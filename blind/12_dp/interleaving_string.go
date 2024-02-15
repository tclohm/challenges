package main

import (
	"fmt"
)

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) { return false }

	// cache := make([][]bool, len(s1) + 1, len(s1) + 1)
	// for i := range s2 {
	// 	cache[i] = make([]bool, len(s2) + 1, len(s2) + 1)
	// }
	// cache[len(s2)] = make([]bool, len(s2) + 1, len(s2) + 1)

	// cache[len(s1)][len(s1)] = true

	var isInterleaved func(i, j int) bool

	isInterleaved = func(i, j int) bool {
		if i == len(s1) && j == len(s2) { return true }
	
		choose_s1, choose_s2 := false, false

		if i < len(s1) && s1[i] == s3[i + j] {
			choose_s1 = isInterleaved(i + 1, j)
		}
		if j < len(s2) && s2[j] == s3[i + j] {
			choose_s2 = isInterleaved(i, j + 1)
		}

		return choose_s1 || choose_s2
	}
	return isInterleaved(0, 0)
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(isInterleave("", "", ""))
}