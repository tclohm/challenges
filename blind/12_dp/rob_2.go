package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 { return nums[0] }

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	// what's bigger, nums[0:length - 1] or nums[1:length], 
	// this just looking at rob house with more boundaries
	return max(robbing(nums, 0, len(nums)-1), robbing(nums, 1, len(nums)))
}

func robbing(nums []int, start, end int) int {
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start + 1] = max(nums[start], nums[start + 1])

	for i := start + 2 ; i < end ; i++ {
		dp[i] = max(nums[i] + dp[i - 2], dp[i - 1])
	}

	return dp[end - 1]
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func main() {
	fmt.Println(rob([]int{2,3,2}))
	fmt.Println(rob([]int{1,2,3,1}))
}

// choose 2 or 3
// 2 [3,2] --> cant choose because adjacent [3,2]
// 3 [2] --> cant choose because adjacent
// 