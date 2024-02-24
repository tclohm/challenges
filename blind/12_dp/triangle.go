package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	// go left or right
	var dfs func (x, y int) int
	dfs = func (x, y int) int {
		if y >= len(triangle) { return 0 }
		if x > len(triangle[y]) { return 0 }

		return min(triangle[y][x] + dfs(x, y + 1), triangle[y][x + 1] + dfs(x + 1, y + 1))

	}
	return triangle[0][0] + dfs(0,1)
}

func minimumDP(triangle [][]int) int {
	dp := make([]int, len(triangle) + 1)

	for i := len(triangle) - 1 ; i >= 0 ; i-- {
		for i, n := range triangle[i] {
			dp[i] = n + min(dp[i], dp[i+1])
		}
	}

	return dp[0]
}

func main() {
	fmt.Println(minimumDP([][]int{{2},{3,4},{6,5,7},{4,1,8,3}}))
	fmt.Println(minimumDP([][]int{{-10}}))
	fmt.Println(minimumDP([][]int{{2},{3,4},{6,5,7},{4,1,8,3},{3,1,4,2,0},{1,1,2,1,1,0}}))
}