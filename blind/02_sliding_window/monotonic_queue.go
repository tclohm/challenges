package main

import "fmt"

func max(nums []int) int {
	big := -10000000
	for _, i := range nums {
		if big < i {
			big = i
		}
	}
	return big
}

func slidingWindowMax(nums []int, k int) []int {
	result := []int{}
	end_buffer := len(nums) - k + 1
	for left := 0 ; left < end_buffer ; left++ {
		right := left + k
		max := max(nums[left:right])
		result = append(result, max)
	}
	return result
}

func main() {
	fmt.Println(slidingWindowMax([]int{1,3,-1,-3,5,3,6,7}, 3))
	fmt.Println(slidingWindowMax([]int{1}, 1))
}