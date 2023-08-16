package main

import (
	"fmt"
)

func coin_change(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1 ; i < amount + 1; i++ {
		for _, coin := range coins {
			if i - coin >= 0 {
				fmt.Println("i", i, "coin", coin)
				dp[i] = min(dp[i], 1 + dp[i - coin])
				fmt.Println(dp)
			}
		}
	}
	fmt.Println("dp at end\t", dp)
	if dp[amount] != amount + 1 {
		return dp[amount]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// [1, 2, 5]
// dp[1] => min(1)
// dp[2] => min(1)
// dp[3] => dp[2] + dp[1] = min(2)
// dp[4] => dp[2] + dp[2] = min(2)
// dp[5] => min(1)
// dp[6] => dp[5] + dp[1] = min(2)
// dp[7] => dp[5] + dp[2] = min(2)
// dp[8] => dp[7] + dp[1] = min(3)
// dp[9] => dp[5] + dp[4] = min(3)
// dp[10] => dp[5] + dp[5] = min(2)
// dp[11] => dp[10] + dp[1] = min(3)

func main() {
	fmt.Println(coin_change([]int{1, 2, 5}, 11))
	fmt.Println(coin_change([]int{2}, 3))
}