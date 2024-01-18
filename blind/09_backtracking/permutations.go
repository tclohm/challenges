package main

import "fmt"

func permutations(nums []int) [][]int {
	var result [][]int

	var backtrack func ([]int, []int)
	backtrack = func(nums []int, path []int) {
		if len(nums) == 0 {
			result = append(result, append([]int(nil), path...))
			return
		}

		for i := 0 ; i < len(nums) ; i++ {
			newNums := append([]int{}, nums[:i]...)
			newNums = append(newNums, nums[i+1:]...)
			newPath := append([]int{}, path...)
			newPath = append(newPath, nums[i])
			backtrack(newNums, newPath)
		}
	}

	backtrack(nums, []int{})
	return result
}

func main() {
	fmt.Println(permutations([]int{0,1}))
	fmt.Println(permutations([]int{1,2,3}))
	fmt.Println(permutations([]int{1}))
}