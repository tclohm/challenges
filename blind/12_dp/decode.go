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
// 1 - 26 == a-z
func bruteforce(s string) int {
	memo := make(map[int]int)
	memo[len(s)] = 1

	var dfs func (s string, i int) int
	dfs = func (s string, i int) int {
		if v, ok := memo[i]; ok {
			return v
		}
		
		length := len(s)
		if length == 0 { return 1 }
		if s[0] == '0' { return 0 }

		result := dfs(s[1:], i+1)

		if length >= 2 {
			if s[:2] >= "10" && s[:2] <= "26" {
				result += dfs(s[:2], i+2)
			}
		}

		memo[i] = result
		return result
	}
	return dfs(s, 0)
}

func main() {
	fmt.Println(decode("12"))
	fmt.Println(decode("226"))
	fmt.Println(decode("06"))
	fmt.Println("--------------")
	fmt.Println(bruteforce("12"))
	fmt.Println(bruteforce("226"))
	fmt.Println(bruteforce("06"))
}