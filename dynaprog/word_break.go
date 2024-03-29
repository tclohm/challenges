package main

import "fmt"

// Two approaches : DP and DFS

// DP, build a boolean array, check if the string can be segmented into dictionary words
// leverages the subproblem overlap and build a bottom up solution

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s) + 1)
	dp[len(s)] = true

	for i := len(s) - 1 ; i >= 0 ; i-- {
		for _, word := range wordDict {
			if (i + len(word)) <= len(s) && s[i: i + len(word)] == word {
				dp[i] = dp[i + len(word)]
			}
			if dp[i] {
				break
			}
		}
	}
	return dp[0]
}

func main() {
	fmt.Println(wordBreak("leetcode",[]string{"leet","code"}))
	fmt.Println(wordBreak("applepenapple",[]string{"apple","pen"}))
	fmt.Println(wordBreak("catsandog",[]string{"cats","dog","sand","and","cat"}))
}