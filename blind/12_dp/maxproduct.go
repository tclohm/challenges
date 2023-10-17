package main

import "fmt"

func maxproduct(nums []int) int {
	result, currentMin, currentMax := nums[0], 1, 1
	// we have to find the current max and current min because 2 negatives could create a large positive
	for i := 0 ; i < len(nums) ; i++ {
		tmp := currentMax * nums[i]
		// what's bigger? currentMax * the current index or the currentMin * the current index, or current index
		currentMax = max(max(nums[i] * currentMax, nums[i] * currentMin), nums[i])
		// what's smaller? current * the current index or current index * currentMin or current index
		currentMin = min(min(tmp, nums[i] * currentMin), nums[i])
		// result, will hold max
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