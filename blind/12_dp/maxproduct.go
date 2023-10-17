package main

import "fmt"

func maxproduct(nums []int) int {
	result, currentMin, currentMax := nums[0], 1, 1

	for i := 0 ; i < len(nums) ; i++ {
		tmp := currentMax * nums[i]
		currentMax = max(max(nums[i] * currentMax, nums[i] * currentMin), nums[i])
		currentMin = min(min(tmp, nums[i] * currentMin), nums[i])
		result = max(result, currentMax)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxproduct([]int{2,3,-2,4}))
	fmt.Println(maxproduct([]int{-2,0,-1}))
}