package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums), len(nums))
	total := 1
	result[0] = nums[0]	
	for i := 1 ; i < len(nums) ; i++ {
		result[i] = nums[i]
		total = nums[i] * total
	}

	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1,2,4,6}))
	fmt.Println(productExceptSelf([]int{-1,0,1,2,3}))
}
