package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums), len(nums))
	for i := 0 ; i < len(result) ; i++ {
		result[i] = 1
	}

	for i := 1 ; i < len(nums) ; i++  {
		result[i] = result[i - 1] * nums[i - 1]
	}

	postfix := 1

	for i := len(result) - 1 ; i >= 0 ; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1,2,4,6}))
	fmt.Println(productExceptSelf([]int{-1,0,1,2,3}))
}
