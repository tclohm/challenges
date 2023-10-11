package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 { return nums[0] }
	return -1
}

func main() {
	fmt.Println([]int{2,3,2})
	fmt.Println([]int{1,2,3,1})
}

// choose 2 or 3
// 2 [3,2] --> cant choose because adjacent [3,2]
// 3 [2] --> cant choose because adjacent
// 