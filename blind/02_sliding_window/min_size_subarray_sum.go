package main

import (
	"fmt"
)

func minSubArray(nums []int, target int) int {
	l := 0
	total := 0
	count := 1000

	for r, _ := range nums {
		total += nums[r]
		for total >= target {
			count = min(r - l + 1, count)
			total -= nums[l]
			l += 1
		}
	}
	if count == 1000 {
		return 0
	}
	return count
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func main() {
	fmt.Println(minSubArray([]int{2,3,1,2,4,3}, 7))
	fmt.Println(minSubArray([]int{1,4,4}, 4))
}