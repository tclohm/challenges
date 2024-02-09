package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	index := len(nums) - 1
	current_sum := 0
	memo := map[string]int{}

	var dp func (nums []int, target int, index int, current_sum int) int
	dp = func (nums []int, target int, index int, current_sum int) int {
		if index < 0 && current_sum == target {
			return 1
		}
		if index < 0 {
			return 0
		}

		if _, ok := memo[fmt.Sprintf("%d, %d", index, current_sum)]; ok {
			return memo[fmt.Sprintf("%d, %d", index, current_sum)]
		}

		positive := dp(nums, target, index-1, current_sum + nums[index])
		negative := dp(nums, target, index-1, current_sum + -nums[index])

		memo[fmt.Sprintf("%d, %d", index, current_sum)] = positive + negative

		return memo[fmt.Sprintf("%d, %d", index, current_sum)]
	}

	return dp(nums, target, index, current_sum)
}

func main() {
	fmt.Println(findTargetSumWays([]int{1,1,1,1,1}, 3))
	fmt.Println(findTargetSumWays([]int{1}, 1))
}