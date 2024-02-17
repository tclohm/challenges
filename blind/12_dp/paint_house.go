package main

import "fmt"

// paint all the houses, but no two adjacent houses can be the same color
func paint(houses [][]int) int {

	var result = 1000

	var painting func(house_index, prev_cost_index int) int
	painting = func(house_index, prev_cost_index int) int {
		if house_index == len(houses) { return 0 }
		var smaller = 1000
		for i, p := range houses[house_index] {
			if i != prev_cost_index {
				smaller = min(p + painting(house_index + 1, i), smaller)
			}
		}
		return smaller
	}

	result = min(painting(0, 0), result)
	result = min(painting(0, 1), result)
	result = min(painting(0, 2), result)

	return result
}

func dyan(houses [][]int) int {
	dp := make([]int, len(houses[0]), len(houses[0]))

	for i := range houses {
		dp0 := houses[i][0] + min(dp[1], dp[2])
		dp1 := houses[i][1] + min(dp[0], dp[2])
		dp2 := houses[i][2] + min(dp[0], dp[1])
		dp = []int{dp0, dp1, dp2}
	}


	return minArray(dp)
}

func minArray(nums []int) int {
	var smaller = 1000
	for _, n := range nums {
		if n < smaller {
			smaller = n
		}
	}

	return smaller
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// red, blue, green
	fmt.Println(dyan([][]int{{17,2,17},{16,16,5},{14,3,19}}))
	fmt.Println(dyan([][]int{{1,2,3},{1,4,6}}))
}