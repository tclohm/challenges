package main

import "fmt"

func decode(s string) int {
	dp := make([]int, len(s) + 1)
	dp[len(s)] = 1

	var dfs func (int) int
	dfs = func (i int) int {
		if dp[i] != 0 {
			return dp[i]
		} else if s[i] == '0' {
			return 0
		}

		res := dfs(i + 1)
		if i + 1 < len(s) &&
		(s[i] == '1' || s[i] == '2' && (s[i + 1] >= '0' && s[i + 1] <= '6')) {
			res += dfs(i + 2)
		}
		dp[i] = res
		return res
	}
	return dfs(0)
}

func main() {
	fmt.Println(decode("12"))
	fmt.Println(decode("226"))
	fmt.Println(decode("06"))
}