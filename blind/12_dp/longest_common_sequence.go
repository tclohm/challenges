package main

import "fmt"

// smaller, longer
func getLengths(a, b string) (string, string) {
	if len(a) <= len(b) {
		return a, b
	} 
	return b, a
}

func makeDP(longer string) [][]int {
	var dp = make([][]int, len(longer), len(longer))
	for i, _ := range longer {
		dp[i] = make([]int, len(longer), len(longer))
	}
	return dp
}

func lcsBF(s1, s2 string) int {
	var longestCommonSub func (i, j int) int
	longestCommonSub = func (i, j int) int {
		if i == len(s1) || j == len(s2) { return 0 }
		if s1[i] == s2[j] {
			return 1 + longestCommonSub(i+1, j+1)
		} else {
			return max(longestCommonSub(i+1, j), longestCommonSub(i, j+1))
		}
	}
	return longestCommonSub(0, 0)
}

func lcsTopDown(s1, s2 string) int {

	var _, longer = getLengths(s1, s2)
	
	var dp = makeDP(longer)

	var longestCommonSub func (i, j int) int
	longestCommonSub = func (i, j int) int {
		if i == len(s1) || j == len(s2) { return 0 }

		if dp[i][j] != 0 {
			return dp[i][j]
		}

		if s1[i] == s2[j] {
			dp[i][j] = 1 + longestCommonSub(i + 1, j + 1)
			return dp[i][j]
		} else {
			return max(longestCommonSub(i + 1, j), longestCommonSub(i, j + 1))
		}
	}

	return longestCommonSub(0, 0)
}


func lcsBottomUpDp(s1, s2 string) int {

	var dp [][]int = make([][]int, len(s1)+1)
	for i := 0 ; i < len(s1) + 1 ; i++ {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 1 ; i < len(s1) + 1 ; i++ {
		for j := 1 ; j < len(s2) + 1 ; j++ {
			if s1[i - 1] == s2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				if dp[i][j - 1] > dp[i - 1][j] {
					dp[i][j] = dp[i][j - 1]
				} else {
					dp[i][j] = dp[i - 1][j]
				}
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func main() {
	fmt.Println(lcsBottomUpDp("abcde", "ace"))
	fmt.Println(lcsBottomUpDp("abc", "abc"))
	fmt.Println(lcsBottomUpDp("abc", "def"))
	fmt.Println(lcsBottomUpDp("eb", "def"))
	fmt.Println(lcsBottomUpDp("qwerty", "dqewferddty"))
}