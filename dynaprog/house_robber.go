package main

import "fmt"

// money stashed in houses
// cant take money from adjacent houses

// input [1, 2, 3, 1]
// tracking system [1, 1, 4, 4]
// output 4

// input [2, 7, 9, 3, 1]
// tracking system [2, 11, 12]

// track both of them 
// 2,9,3 = 12
// 7,3,1 = 11

func max(a, b int) int {
	if a > b { return a }
	return b
}

func rob(nums []int) int {
	one, two := 0, 0
	for i := 0 ; i < len(nums) ; i += 2 {
		one += nums[i]
	}

	for i := 1 ; i < len(nums) ; i += 2 {
		two += nums[i]
	}

	return max(two, one)
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}