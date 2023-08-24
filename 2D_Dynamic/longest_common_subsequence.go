package main

import "fmt"

// two strings return the longest common subsequence, 
// if no subsequence, return 0
// subsequence of a string is a new string generated from 
// the original with some chars deleted or not, without changing
// the relative order of the remaining characters


// abcde, ace
func longestCommonSubsequence(text1 string, text2 string) int {
	// build dp
	dp := make([][]int, len(text1) + 1)
	for i := 0 ; i < len(dp) ; i++ {
		dp[i] = make([]int, len(text2) + 1)
	}

	for index := len(text1) - 1; index >= 0 ; index-- {
		for jndex := len(text2) - 1 ; jndex >= 0 ; jndex-- {

			if text1[index] == text2[jndex] {
				dp[index][jndex] = 1 + dp[index + 1][jndex + 1]
			} else {
				dp[index][jndex] = max(dp[index][jndex + 1], dp[index + 1][jndex])
			}

		}
	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace")) // 3
	fmt.Println(longestCommonSubsequence("abc", "abc")) // 3
	fmt.Println(longestCommonSubsequence("abc", "def")) // 0
}