package main

import (
	"fmt"
	"sort"
)

// sort n log n
func makeUnique(nums []int) int {
	sort.Ints(nums)
	incremented := 0

	for right := 1 ; right < len(nums) ; right++ {
		left := right - 1
		// is left greater or equal to right
		if nums[left] >= nums[right] {
			delta := nums[left] - nums[right]
			incremented += 1 + delta			
			nums[right] = nums[left] + 1
		}
	}
	return incremented
}

func main() {
	fmt.Println(makeUnique([]int{1,2,2}))
	fmt.Println(makeUnique([]int{3,2,1,2,1,7}))

}
