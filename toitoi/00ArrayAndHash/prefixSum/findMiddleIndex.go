package main

import "fmt"

func sum(nums []int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}
	return total
}

func findMiddleIndex(nums []int) int {
	total := sum(nums)
	left := 0

	for i := range nums {
		right := total - left - nums[i]
		if right == left { return i }
		left += nums[i]
	}

	return -1 
}

func main() {
	fmt.Println(findMiddleIndex([]int{2,3,-1,8,4}))
	fmt.Println(findMiddleIndex([]int{1,-1,4}))
	fmt.Println(findMiddleIndex([]int{2,5}))
}
