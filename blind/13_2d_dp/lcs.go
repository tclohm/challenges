package main

import "fmt"

func lcs(text1, text2 string) int {
	dp := make([][]int, len(text1) + 1)

	for i := 0 ; i < len(dp) ; i++ {
		dp[i] = make([]int, len(text2) + 1)
	}

	for i := len(text1) - 1 ; i >= 0 ; i-- {
		for k := len(text2) - 1 ; k >= 0 ; k-- {
			if text1[i] == text2[k] {
				dp[i][k] = 1 + dp[i + 1][k + 1]
			} else {
				dp[i][k] = max(dp[i][k + 1], dp[i + 1][k])
			}
		}
	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lcs("abcde", "ace"))
	fmt.Println(lcs("abc", "abc"))
	fmt.Println(lcs("abc", "def"))
}