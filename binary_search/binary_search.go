package main

import "fmt"

func binary_search(nums []int, target int) int {
	right := len(nums) - 1
	left := 0

	for left < right {
		// python doesn't have this problem
		middle := left + (right - left) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

func main() {
	n1 := []int{-1,0,3,5,9,12}

	fmt.Println(binary_search(n1, 9))
	fmt.Println(binary_search(n1, 2))
}