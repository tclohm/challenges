package main

import (
	"fmt"
	// "math"
)

// find a subarray that has the largest product
// declare two variables for tracking the min and max product value
// track min for negative num because if we get two negs, they turn positive
func max_product(nums []int) int {
	min_value := nums[0]
	max_value := nums[0]
	max_product := nums[0]

	for i := 1 ; i < len(nums) ; i++ {
		if nums[i] < 0 {
			min_value, max_value = max_value, min_value 
		}

		min_value = min(nums[i], min_value * nums[i])
		max_value = max(nums[i], max_value * nums[i])

		max_product = max(max_value, max_product)
	}

	return max_product
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func main() {
	// subarray -> [2][2,3][2,3,-2][3,-2][3,-2,4][-2,4] 
	fmt.Println(max_product([]int{2, 3, -2, 4})) // output 6 because [2,3]
}