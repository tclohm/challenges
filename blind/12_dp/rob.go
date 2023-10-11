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

func main() {
	fmt.Println(rob([]int{1,2,3,1}))
	fmt.Println(rob([]int{2,7,9,3,1}))
}