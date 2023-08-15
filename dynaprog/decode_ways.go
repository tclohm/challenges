package main

import "fmt"

// two ways to decode
// take the first num
// take the first two nums

// f("1234") = f("234") + f("34")
// string is "", return 1
// string is "0" or start with "0", we return 0. "0" is not a valid num
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	size := len(s)
	dp := make([]int, size+1)
	dp[0], dp[1] = 1, 1

	for i := 2 ; i <= size ; i++ {
		s2 := s[i - 2:i]
		if "10" <= s2 && s2 <= "26" {
			dp[i] = dp[i - 2]
		}

		if s[i - 1] != '0' {
			dp[i] += dp[i - 1]
		}
	}

	return dp[size]
}

func main() {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("06"))
}