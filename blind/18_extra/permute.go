package main

import (
	"fmt"
)

func firstPermute(nums []int) [][]int {
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

func in(n int, array []int) bool {
	for _, num := range array {
		if n == num { return true }
	}
	return false
}

func permute(nums []int) [][]int {
	
	var result = [][]int{}
	
	var backtrack func(currentPermutation []int)
	backtrack = func(currentPermutation []int) {
		if len(currentPermutation) == len(nums) {
			result = append(result, append([]int(nil), currentPermutation...))
			return
		}

		for _, num := range nums {
			if !in(num, currentPermutation) {
				currentPermutation = append(currentPermutation, num)
				backtrack(currentPermutation)
				currentPermutation = currentPermutation[:len(currentPermutation) - 1]
			}
		}
	}

	backtrack([]int{})
	return result
}

func main() {
	fmt.Println(permute([]int{1,2,3}))
	fmt.Println(permute([]int{8}))
	fmt.Println(permute([]int{0,1}))
}
