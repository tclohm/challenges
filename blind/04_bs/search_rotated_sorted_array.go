package main

import "fmt"

func search(nums []int, target int) int {
	right, left := len(nums) - 1, 0
	// is the target larger than the right, search left side
	// is the target larger than the left, search right side
	for left <= right {
		middle := left + ((right - left) / 2)
		if nums[middle] == target { return middle }
		// left sorted
		if nums[left] <= nums[middle] {
			if target > nums[middle] || target < nums[left] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			if target < nums[middle] || target > nums[right] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 3))
	fmt.Println(search([]int{1}, 0))
}