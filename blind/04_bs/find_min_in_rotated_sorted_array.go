package main

import "fmt"

func findMin(nums []int) int {
	left_bound := 0
	right_bound := len(nums) - 1
	// left + ((right - left) / 2)
	middle := left_bound + ((right_bound - left_bound) / 2)

	for {
		if nums[left_bound] == nums[right_bound] {
			return nums[middle]
		} else if nums[left_bound] > nums[right_bound] {
			//fmt.Println("left is greater")
			left_bound = middle
			middle = left_bound + ((right_bound - left_bound) / 2)
		} else {
			//fmt.Println("right is greater")
			right_bound = middle
			middle = left_bound + ((right_bound - left_bound) / 2)
		}
	}
	return middle
}

func main() {
	fmt.Println(findMin([]int{3,4,5,1,2}))
	fmt.Println(findMin([]int{4,5,6,7,0,1,2}))
	fmt.Println(findMin([]int{11,13,15,17}))
}