package main

import "fmt"

func rob(nums []int) int {
	// 2 options: rob current house or don't rob current house
	// rob(i) = max(rob(i - 2) + currentValue, rob(i - 1))
	return robbing(nums, len(nums) - 1)
}

func robbing(nums []int , length int) int {
	if length < 0 {
		return 0
	}
	return max(robbing(nums, length - 2) + nums[length], robbing(nums, length - 1))
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func robbingMemo(nums []int) int {
	memo := make([]int, len(nums) + 1)

	for i := 0 ; i < len(memo) ; i++ {
		memo[i] = -1
	}

	return robMemo(nums, len(nums) - 1, memo)
}

func robMemo(nums []int, length int, memo []int) int {
	if length < 0 {
		return 0
	}
	if memo[length] >= 0 {
		return memo[length]
	}

	result := max(robMemo(nums, length - 2, memo) + nums[length], robMemo(nums, length - 1, memo))
	memo[length] = result
	return result
}

func main() {
	fmt.Println(robbingMemo([]int{1,2,3,1}))
	fmt.Println(robbingMemo([]int{2,7,9,3,1}))
}