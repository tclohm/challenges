package main

import "fmt"

func uniquePaths(m int, n int) int {
    dp := make([]int, n)
    dp[0] = 1

    for i := 0 ; i < m ; i++ {
    	for k := 1 ; k < n ; k++ {
    		dp[k] += dp[k - 1]
    	}
    }

    return dp[n - 1]
}

func main() {
	fmt.Println(uniquePaths(3,7))
	fmt.Println(uniquePaths(3,2))
}