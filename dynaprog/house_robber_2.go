package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	return max(robbing(nums, 0, n-1), robbing(nums, 1, n))
}

func robbing(nums []int, start, end int) int {
	dp := make([]int, len(nums))
	dp[start] = nums[stairs]
	dp[start+1] = max(nums[start], nums[start + 1])

	for i := start + 2 ; i < end ; i++ {
		dp[i] = max(nums[i] + dp[i - 2], db[i - 1])
	}

	return dp[end - 1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
}