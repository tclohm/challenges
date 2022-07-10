package main

import "fmt"

func bruteForceSearch(nums []int, target int) int {
	for index, number := range nums {
		if number == target {
			return index
		}
	}
	return -1
}

func binarySearch(nums []int, target int) int {
	// left sorted portion
	// right sorted portion
	var left = 0
	var right = len(nums) - 1

	for left <= right {
		// find middle
		middle := left + (right - left) / 2
		// if the middle the target return the target
		if nums[middle] == target {
			return middle
		}
		// if middle is less than left
		if nums[left] <= nums[middle] {
			if target > nums[middle] || target < nums[left] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
		// if middle is less than right
		else {
			if target <= nums[middle] || target > nums[right] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return -1
}


func main() {
	n1 := []int{4,5,6,7,0,1,2}

	fmt.Println(binarySearch(n1, 0))

	fmt.Println(binarySearch(n1, 3))

	n2 := []int{1}

	fmt.Println(binarySearch(n2, 0))
}