package main

import "fmt"

func subset(nums []int) [][]int {
	var list [][]int

	var backtrack func(tmp, nums []int, start int)
	backtrack = func(tmp, nums []int, start int) {
		list = append(list, append([]int{}, tmp...))
		for i := start ; i < len(nums) ; i++ {
			tmp = append(tmp, nums[i])
			backtrack(tmp, nums, i + 1)
			tmp = tmp[:len(tmp) - 1]
		}
	}

	backtrack([]int{}, nums, 0)
	return list
}



func main() {
	fmt.Println(subset([]int{1,2,3}))
	fmt.Println(subset([]int{0}))
}