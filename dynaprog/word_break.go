package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s) + 1)
	dp[len(s)] = true

	for i := len(s) - 1 ; i >= 0 ; i-- {
		for _, w := range wordDict {
			if (i + len(w)) <= len(s) && s[i: i + len(w)] == w {
				dp[i] = dp[i + len(w)]
			}
			if dp[i] {
				break
			}
		}
	}
	fmt.Println(dp)
	return dp[0]
}


func main() {
	fmt.Println(wordBreak("leetcode",[]string{"leet","code"}))
	fmt.Println(wordBreak("applepenapple",[]string{"apple","pen"}))
	fmt.Println(wordBreak("catsandog",[]string{"cats","dog","sand","and","cat"}))
}