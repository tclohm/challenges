package main

import (
	"fmt"
	"sort"
)

func sortJumbled(mapping, nums []int) []int {

	pairs := map[int][2]int{}

	for i, n := range nums {
		mappedN := 0
		base := 1
		if n == 0 {
			mappedN = mapping[n]
		}
		for n > 0 {
			digit := n % 10
			n = n / 10
			mappedN += base * mapping[digit]
			base *= 10
		}
		pairs[n] = [2]int{mappedN, i}
	}

	sort.Slice(nums, func(i, j int) bool {
		if pairs[nums[i]][0] == pairs[nums[j]][0] {
			return pairs[nums[i]][1] < pairs[nums[j]][1]
		}
		return pairs[nums[i]][0] < pairs[nums[j]][0]
	})
	return nums
}

func main() {
	fmt.Println(sortJumbled([]int{8,9,4,0,2,1,3,5,7,6}, []int{991,338,38}))
	fmt.Println(sortJumbled([]int{0,1,2,3,4,5,6,7,8,9}, []int{789,456,123}))
}
