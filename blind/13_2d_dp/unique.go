package main

import "fmt"

func BF(m, n int) int {
    var unique func(i, j int) int
    unique = func(i, j int) int {
        if i >= m || j >= n { return 0 } // out of bounds
        if i == m - 1 && j == n - 1 { return 1 } // valid
        return unique(i + 1, j) + unique(i, j + 1) // try both down and right
    }
    return unique(0,0)
}

func BFM(m, n int) int {
    var dp = make([][]int, 101, 101)
    for i := 0 ; i < len(dp) ; i++ {
        dp[i] = make([]int, 101, 101)
    }
    
    var unique func (i, j int) int
    unique = func (i, j int) int {

        if i >= m || j >= n { return 0 }
        if i == m - 1 && j == n - 1 { return 1 }
        if dp[i][j] != 0 { return dp[i][j] }

        dp[i][j] = unique(i + 1, j) + unique(i, j + 1)
        
        return dp[i][j]
    }

    return unique(0,0)
}

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
	fmt.Println(BFM(3,7))
	fmt.Println(BFM(3,2))
}