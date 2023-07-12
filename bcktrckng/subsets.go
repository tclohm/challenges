package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	return dfs(nums, 0, [][]int{{}})
}

func dfs(nums []int, i int, sets [][]int) [][]int {
	if i == len(nums) {
		return sets
	}

	for _, set := range sets {
		sets = append(sets, append([]int{nums[i]}, set...))
	}
	return dfs(nums, i+1, sets)
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}