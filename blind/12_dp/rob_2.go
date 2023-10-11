package main

import "fmt"

func rob(nums []int) int {
	highest := nums[0]
	i := 0

	if i + 1 == len(nums) {
		return 0
	}

	// max(nums[i], nums[i+1])

	// nums[step]
	// nums[step + 1] 
	
}

func main() {
	fmt.Println([]int{2,3,2})
	fmt.Println([]int{1,2,3,1})
}

// choose 2 or 3
// 2 [3,2] --> cant choose because adjacent [3,2]
// 3 [2] --> cant choose because adjacent
// 