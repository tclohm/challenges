package main

import "fmt"
// sum of all numbers to the left of the index is equal to the sum of the right
func findPivot(nums []int) int {
	total := 0
	left := 0

	for i := range nums {
		total += nums[i]
	}

	for i := range nums {
		right := total - left - nums[i]
		if right == left { return i }
		left += nums[i]
	}

	return -1
}

func main() {
	fmt.Println(findPivot([]int{1,7,3,6,5,6}))
	fmt.Println(findPivot([]int{1,2,3}))
	fmt.Println(findPivot([]int{2,1,-1}))
}
