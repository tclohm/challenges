package main

import "fmt"

func permutations(nums []int) [][]int {
	var list [][]int
	var used = make([]bool, len(nums), len(nums))

	var backtrack func(tmp, nums []int, used []bool)
	backtrack = func(tmp, nums []int, used []bool) {
		if len(tmp) == len(nums) {
			list = append(list, append([]int{}, tmp...))
		} else {
			for i := 0 ; i < len(nums) ; i++ {
				if used[i] || i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {continue}

				used[i] = true
				tmp = append(tmp, nums[i])
				backtrack(tmp, nums, used)
				used[i] = false
				tmp = tmp[:len(tmp) - 1]
			}
		}
	}
	backtrack([]int{}, nums, used)
	return list
}


func main() {
	fmt.Println(permutations([]int{1,1,2}))
	fmt.Println(permutations([]int{1,2,3}))
}