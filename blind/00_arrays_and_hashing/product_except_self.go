package main

import "fmt"

// travel forward and back
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	pre := 1
	for i, _ := range nums {
		result[i] = pre
		pre = pre * nums[i]
	}

	post := 1
	for i := len(nums) - 1 ; i >= 0 ; i-- {
		result[i] = result[i] * post
		post = post * nums[i]
	}
	
	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1,2,3,4}))
	fmt.Println(productExceptSelf([]int{-1,1,0,-3,3}))
}