package main

import "fmt"

func comb(nums []int, target int) [][]int {
	var list [][]int

	var backtrack func (tmp []int, nums []int, remain int, start int)
	backtrack = func(tmp []int, nums []int, remain, start int) {
		if remain < 0 {
			return 
		} else if remain == 0 {
			// coping tmp			
			list = append(list, append([]int{}, tmp...))
		} else {
			for i := start ; i < len(nums) ; i++ {
				tmp = append(tmp, nums[i])
				backtrack(tmp, nums, remain - nums[i], i)
				tmp = tmp[:len(tmp) - 1]
			}
		}
	}

	backtrack([]int{}, nums, target, 0)
	return list
}

func main() {
	fmt.Println(comb([]int{2,3,6,7}, 7))
	fmt.Println(comb([]int{2,3,5}, 8))
	fmt.Println(comb([]int{2}, 1))
}