package main

import "fmt"

func maxSubArray(nums []int) int {
	result, total := nums[0], 0

	for i := 0 ; i < len(nums) ; i++ {
		total += nums[i]
		result = max(result, total)
		if total < 0 {
			total = 0
		}
	}
	return result
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5,4,-1,7,8}))
}