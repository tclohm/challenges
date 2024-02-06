package main

import (
	"fmt"
	"sort"
)

// return max possible frequency of elements 
// after performing at most k operation


// only increment
// [1,2,4], 5
// []
func maxFreq(nums []int, k int) int {
	sort.Ints(nums)
	var left, right = 0, 0
	var result, total = 0, 0

	for right < len(nums) {
		total += nums[right]
		// window * value

		// intuition, how much do we need to convert
		// our right pointer will be biggest since we have sorted it
		// if our window * the value is greater than our total + k (our budget)
		// DECREASE the window and total
		for nums[right] * (right - left + 1) > total + k {
			total -= nums[left]
			left += 1
		}
		// check if our result is bigger than our window
		result = max(result, right - left + 1)
		// increase our window
		right += 1
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxFreq([]int{1,2,4}, 5))
	fmt.Println(maxFreq([]int{1,4,8,13}, 5))
	fmt.Println(maxFreq([]int{3,6,9}, 2))
}