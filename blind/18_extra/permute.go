package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var perms [][]int

	var backtrack func([]int, []int)
	backtrack = func(nums []int, path []int) {
		if len(nums) == 0 {
			perms = append(perms, append([]int(nil), path...))
			return
		}

		for i := 0 ; i < len(nums) ; i++ {
			newPerms := append([]int(nil), nums[:i]...)
			newPerms = append(newPerms, nums[i + 1:]...)
			newPath := append([]int(nil), path...)
			newPath = append(newPath, nums[i])
			backtrack(newPerms, newPath)
		}
	}

	backtrack(nums, []int{})
	return perms
}
func main() {
	fmt.Println(permute([]int{1,2,3}))
	fmt.Println(permute([]int{8}))
	fmt.Println(permute([]int{0,1}))
}
