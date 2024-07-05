package main

import "fmt"

func permute(nums []int) [][]int {


	var dfs func(result *[][]int, nums []int) [][]int
	dfs = func(result *[][]int, nums []int) [][]int {
		if len(nums) == 1 {
			return [][]int{nums}
		}

		for i := 0 ; i < len(nums) ; i++ {
			n := nums[0]
			nums = nums[1:]
			perms := dfs(result, nums)	
			for _, perm := range perms  {
				perm = append(perm, n)
			}
			*result = append(*result, perms...)
			nums = append(nums, n)
		}
		return *result
	}

	var result = [][]int{}
	dfs(&result, nums)
	return result
}

func main() {
	fmt.Println(permute([]int{1,2,3}))
	fmt.Println(permute([]int{8}))
}
