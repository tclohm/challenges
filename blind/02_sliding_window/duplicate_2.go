package main

import "fmt"

func duplicate(nums []int, k int) bool {
	window := make(map[int]bool)
	left := 0

	for right, _ := range nums {
		// bigger than k, remove left
		if right - left > k {
			delete(window, nums[left])
			left += 1
		}
		// right in window
		if _, ok := window[nums[right]]; ok {
			return true
		}
		// move the right into window
		window[nums[right]] = true
	}

	return false
}


func main() {
	fmt.Println(duplicate([]int{1,2,3,1}, 3))
	fmt.Println(duplicate([]int{1,2,3,1}, 1))
}