package main

import (
	"fmt"
	"sort"
)


func minDiff(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}

	sort.Ints(nums)
	result := 10000

	for left := 0 ; left < 4 ; left++ {
		right := len(nums) - 4 + 1
		result = min(result, nums[right] - nums[left])
	}

	return result
}

func main() {
	fmt.Println(minDiff([]int{5,3,2,4}))
	fmt.Println(minDiff([]int{1,5,0,10,14}))
	fmt.Println(minDiff([]int{3,100,20}))
}
