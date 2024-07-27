package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums), len(nums))
	total := 1 
	for i := range nums {
		result[i] = nums[i]
		if nums[i] != 0 {
			total *= nums[i]
		}
	}
		
	for i := range nums {
		if nums[i] == 0 { continue }
		result[i] = total / nums[i]
	}

	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1,2,4,6}))
	fmt.Println(productExceptSelf([]int{-1,0,1,2,3}))
}
