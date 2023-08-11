package main

import "fmt"

// [10, 15, 20]
// 1, 2 steps

// 						0 steps
// 				10 					15
// 			15			20		20 		end
// 		20 		end end 	end end end

func min_cost(stairs []int) int {
	dp := make([]int, len(stairs))
	n := len(stairs)

	dp[0] = stairs[0]
	dp[1] = stairs[1]

	for i := 2 ; i < n ; i++ {
		dp[i] = stairs[i] + min(dp[i - 1], dp[i - 2])
	}

	return min(dp[n - 1], dp[n - 2])
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func main() {
	fmt.Println(min_cost([]int{10, 15, 20}))
	fmt.Println(min_cost([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}