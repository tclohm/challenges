package main

import (
	"fmt"
	"sort"
)

func c2(nums []int, target int) [][]int {
	var list [][]int
	sort.Ints(nums)
	var backtrack func (tmp, nums []int, remain, start int)
	backtrack = func (tmp, nums []int, remain, start int) {
		if remain < 0 { 
			return 
		} else if remain == 0 {
			list = append(list, append([]int{}, tmp...))
		} else {
			for i := start ; i < len(nums) ; i++ {
				if i > start && nums[i] == nums[i - 1] { continue }
				tmp = append(tmp, nums[i])
				backtrack(tmp, nums, remain - nums[i], i + 1)
				tmp = tmp[:len(tmp) - 1]
			}
		}
	}

	backtrack([]int{}, nums, target, 0)
	return list
}

func main() {
	fmt.Println(c2([]int{10,1,2,7,6,1,5}, 8))
	fmt.Println(c2([]int{2,5,2,1,2}, 5))
}