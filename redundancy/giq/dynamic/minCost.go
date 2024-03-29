package main

import "fmt"

// given staircase, i-th step is assigned a non-negative cost indicated by a cost array
// pay the cost for a step, climb one or two steps
// find the "minimum cost" to reach the top of the staircase
// first step can be the first or second step

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(i int, cost []int, dp []int) int {
	if i < 0 { return 0 }
	if i == 0 || i == 1 { return cost[i] }
	if dp[i] != 0 {
		return dp[i]
	}
	dp[i] = cost[i] + min(minCost(i - 1, cost, dp), minCost(i - 2, cost, dp))
	return dp[i]
}

// top down
func climbStairs(cost []int) int {
	var n = len(cost)
	var dp = make([]int, n, n)
	return min(minCost(n - 1, cost, dp), minCost(n - 2, cost, dp))
}

// bottom up
func climbing(cost []int) int {
	var n = len(cost)
	//var dp = make([]int, n, n)

	if n == 0 { return 0 }
	if n == 1 { return cost[0] }

	var dpOne = cost[0]
	var dpTwo = cost[1]

	for i := 2 ; i < n ; i++ {
		var current = cost[i] + min(dp[i-1], dp[i-2])
		dpOne = dpTwo
		dpTwo = current
	}
	return min(dpOne, dpTwo)
}

func main() {
	cost := []int{10, 15, 6, 4}
	fmt.Println(climbStairs(cost))
	fmt.Println(climbing(cost))
}