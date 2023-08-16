package main

import (
	"fmt"
)

func coin_change(coins []int, amount int) int {
	// track how much the amount will be at every coin
	dp := make([]int, amount+1)
	// make all the amounts one higher than the amount
	for i := range dp {
		dp[i] = amount + 1
	}
	// zero is zero
	dp[0] = 0
	// for each amount
	for i := 1 ; i < amount + 1; i++ {
		// for each coin
		for _, coin := range coins {
			// if the amount - coin is greater than 0, add the 
			if i - coin >= 0 {
				dp[i] = min(dp[i], 1 + dp[i - coin])
			}
		}
	}
	
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