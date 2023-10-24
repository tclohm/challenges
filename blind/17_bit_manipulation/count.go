package main

import "fmt"

func count(n int) []int {
	dp := make([]int, n + 1)
	offset := 1

	for i := 1 ; i <= n ; i++ {
		if offset * 2 == i {
			offset = i
		}
		dp[i] = 1 + dp[i - offset]
	}
	return dp
}

func main() {
	fmt.Println(count(2))
	fmt.Println(count(5))
}