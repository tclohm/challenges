package main

import "fmt"

func permute(nums []int) [][]int {


	var backtrack func(result *[][]int, nums []int)
	backtrack = func(result [][]int, nums []int) {

	}
	if len(nums) == 1 {
		return [][]int{nums}
	}

	result := [][]int{}

	for i := 0 ; i < len(nums) ; i++ {
		n := nums[0]
		nums = nums[1:]
		perms := permute(nums)	
		for _, perm := range perms  {
			perm = append(perm, n)
		}
		result = append(result, perms...)
		nums = append(nums, n)
	}

	return result

}

func main() {
	fmt.Println(permute([]int{1,2,3}))
	fmt.Println(permute([]int{8}))
}
