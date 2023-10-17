package main

import "fmt"

func wordbreak(s string, dict []string) bool {
	dp := make([]bool, len(s) + 1)
	dp[len(s)] = true

	for i := len(s) - 1 ; i >= 0 ; i-- {
		for _, word := range dict {
			if (i + len(word)) <= len(s) && s[i : i + len(word)] == word {
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
	fmt.Println(wordbreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordbreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordbreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}