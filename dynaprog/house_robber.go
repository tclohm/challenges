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
	prev, current := 0, 0
	
	for _, num := range nums {
		tmp := max(num + prev, current)
		prev = current
		current = tmp
	}
	// current will be the max	
	return current
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}