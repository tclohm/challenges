package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums), len(nums))
	result[0] = nums[0]
	for i := 1 ; i < len(nums) ; i++ {
		result[i] = 1
	}
	
	for i := 0 ; i < len(result) ; i++ {
		for j := 0 ; j < len(result) ; j++ {
			if j == i { continue }
			result[i] *= nums[j]
		}
	}


	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1,2,4,6}))
	fmt.Println(productExceptSelf([]int{-1,0,1,2,3}))
}
