package main

import (
	"fmt"
)

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) { return false }

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

// DP 1D Array
func DPisInterleave(s1, s2, s3 string) bool {
	m, n := len(s1), len(s2)
	if m + n != len(s3) { return false }

	prev := make([]bool, n + 1)

	for i := 0 ; i <= m ; i++ {
		curr := make([]bool, n+1)
		for j := 0 ; j <= n ; j++ {
			if i == 0 && j == 0 {
				curr[j] = true
			} else if i == 0 {
				curr[j] = curr[j - 1] && s2[j - 1] == s3[i + j - 1]
			} else if j == 0 {
				curr[j] = prev[j] && s1[i - 1] == s3[i + j - 1]
			} else {
				curr[j] = (curr[j - 1] && (s2[j - 1] == s3[i + j - 1]) || prev[j] && (s1[i - 1] == s3[i + j - 1]))
			}
		}
		prev = curr
	}
	return prev[n]
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(isInterleave("", "", ""))
}