package main

import "fmt"

func containsDuplicate(nums []int) bool {
	hm := make(map[int]bool)
	for i := 0 ; i < len(nums) ; i++ {
		_, exists := hm[nums[i]]
		if exists {
			return true
		}
		hm[nums[i]] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	nums_two := []int{1, 2, 3, 4}
	nums_three := []int{1,1,1,3,3,4,3,2,4,2}

	fmt.Println(containsDuplicate(nums))
	fmt.Println(containsDuplicate(nums_two))
	fmt.Println(containsDuplicate(nums_three))
}