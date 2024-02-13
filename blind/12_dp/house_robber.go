package main

import "fmt"

func max(a, b int) int {
	if a < b { return b }
	return a
}

func rob(nums []int) int {

	var robbing func(index int) int
	robbing = func(index int) int {
		if index == len(nums) { return 0 }
		if index + 2 <= len(nums) {
			return robbing(index + 2) + nums[index]
		}
		return nums[index]
	}

	return max(robbing(0), robbing(1))
}

func main() {
	fmt.Println(rob([]int{1,2,3,1}))
	fmt.Println(rob([]int{2,7,9,3,1}))
	fmt.Println(rob([]int{1,7,2,1,10,1}))
	fmt.Println(rob([]int{8,9}))
}