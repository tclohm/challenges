package main

import "fmt"

// sliding window --- T: O(n) S: O(1)
func numSubarraysWithSum(nums []int, goal int) int {
	var sumAtMost func([]int, int) int
	sumAtMost = func(nums []int, goal int) int {
		if goal < 0 { return 0 }
		total, left, count := 0, 0, 0
		for right, val := range nums {
			// add it up
			count += val
			// if we are over the goal, decrement value and move left pointer
			for count > goal {
				count -= nums[left]
				left++
			}
			// grab the total amount between the pointers
			total += right - left + 1
		}
		return total
	}
	// the difference of goal and goal - 1
    return sumAtMost(nums, goal) - sumAtMost(nums, goal - 1)
}

func main() {
	fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1}, 2))
	fmt.Println(numSubarraysWithSum([]int{0,0,0,0,0}, 0))
}