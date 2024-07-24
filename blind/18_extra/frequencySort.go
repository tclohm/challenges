package main

import (
	"fmt"
	"sort"
)

func frequency(nums []int) []int {
	ht := map[int]int{}

	for _, n := range nums {
		ht[n]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if ht[nums[i]] == ht[nums[j]] {
			return nums[i] > nums[j]
		}
		return ht[nums[i]] < ht[nums[j]]
	})

	return nums
}


func main() {
	fmt.Println(frequency([]int{1,1,2,2,2,3}))
	fmt.Println(frequency([]int{2,3,1,3,2}))
	fmt.Println(frequency([]int{-1,1,-6,4,5,-6,1,4,1}))

}
