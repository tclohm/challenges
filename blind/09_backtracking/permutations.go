package main

import "fmt"

func permutations(nums []int) [][]int {
	var result [][]int

	var backtrack func ([]int, []int, int)
	backtrack = func(nums []int, path []int, stacklevel int) {
		if len(nums) == 0 {
			result = append(result, append([]int{}, path...))
			return
		}

		for i := 0 ; i < len(nums) ; i++ {
			newNums := append([]int{}, nums[:i]...)
			fmt.Println("stack level:", stacklevel, "new nums init:", newNums)
			newNums = append(newNums, nums[i+1:]...)
			fmt.Println("stack level:", stacklevel, "new nums append, i + 1:", newNums)
			newPath := append([]int{}, path...)
			fmt.Println("stack level:", stacklevel, "new path init:", newPath)
			newPath = append(newPath, nums[i])
			fmt.Println("stack level:", stacklevel, "new nums append i:", newPath)
			backtrack(newNums, newPath, stacklevel+1)
		}
	}

	backtrack(nums, []int{}, 0)
	return result
}

func main() {
	fmt.Println(permutations([]int{0,1}))
	fmt.Println(permutations([]int{1,2,3}))
	fmt.Println(permutations([]int{1}))
}