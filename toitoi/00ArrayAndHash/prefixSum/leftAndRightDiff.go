package main

import (
	"fmt"
	"math"
)

func leftRightDifference(nums []int) []int {
	prefix := make([]int, len(nums), len(nums))
	prefix[0] = nums[0]
	for i := 1 ; i < len(nums) ; i++ {
		prefix[i] = prefix[i - 1] + nums[i]
	}
	
	postfix := make([]int, len(nums), len(nums))
	end := len(nums) - 1
	postfix[end] = nums[end]

	for i := end - 1 ; i >= 0 ; i-- {
		postfix[i] = postfix[i + 1] + nums[i]
	}

	for i := 0 ; i < len(nums) ; i++ {
		diff := postfix[i] - prefix[i]
		prefix[i] = int(math.Abs(float64(diff)))
	}

	return prefix
}

func main() {
	fmt.Println(leftRightDifference([]int{10,4,8,3}))
	fmt.Println(leftRightDifference([]int{1}))
}
