package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0, 1<<n)
	current := make([]int, 0, n)
	sort.Ints(nums)

	var backtrack func(index int)
	backtrack = func(index int) {
		result = append(result, append([]int{}, current...))
		for i := index ; i < n ; i++ {
			if i > index && nums[i] == nums[i - 1] {
				continue
			}
			current = append(current, nums[i])
			backtrack(i + 1)
			current = current[:len(current) - 1]
		}
	}

	backtrack(0)

	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 2}))
	fmt.Println(subsets([]int{0}))
}