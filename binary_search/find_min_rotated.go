package main

import "fmt"

func find_min_rotated(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := left + (right - left) / 2

		if nums[left] < nums[right] {
			if nums[left] < nums[middle] {
				right = middle - 1
			} else {
				right = middle
			}
		} else {
			if nums[middle] < nums[right] {
				left = middle
			} else {
				left = middle + 1
			}
		}
	}

	return nums[right]
}

func main() {
	n1 := []int{3,4,5,1,2}
	n2 := []int{4,5,6,7,0,1,2}
	n3 := []int{11,13,15,17}

	fmt.Println(find_min_rotated(n1))
	fmt.Println(find_min_rotated(n2))
	fmt.Println(find_min_rotated(n3))
}