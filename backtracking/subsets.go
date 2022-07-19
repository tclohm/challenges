package main

import "fmt"

func subsets(nums []int) [][]int {
	result := [][]int{}
	dfs(&result, nil, nums, 0)
	return result
}

func dfs(result *[][]int, subset, nums []int, index int) {
	if index == len(nums) {
		*result = append(*result, append([]int{}, subset...))
		return
	}

	subset = append(subset, nums[index])
	dfs(result, subset, nums, index+1)
	subset = subset[:len(subset)-1]
	dfs(result, subset, nums, index+1)
}

func main() {
	nums := []int{1,2,3} // 2^n = num of subsets O(n * 2^n)
	n := []int{0}
	fmt.Println(subsets(nums))
	fmt.Println(subsets(n))
}