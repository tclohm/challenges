package main

import "fmt"
// either the smallest will be on the left side or right side
func findMin(nums []int) int {
	left_bound := 0
	right_bound := len(nums) - 1
	// left + ((right - left) / 2)
	smallest := nums[0]

	for left_bound <= right_bound {
		middle := left_bound + ((right_bound - left_bound) / 2)
		// if left is less than middle than move the left to the middle
		if nums[0] <= nums[middle] {
			left_bound = middle + 1
		} else {
			// if middle is smaller than smallest, set it as smallest and move right to middle
			if nums[middle] < smallest {
				smallest = nums[middle]
			}
			right_bound = middle - 1
		}

	}
	return smallest
}

func main() {
	fmt.Println(findMin([]int{3,4,5,1,2}))
	fmt.Println(findMin([]int{4,5,6,7,0,1,2}))
	fmt.Println(findMin([]int{11,13,15,17}))
}