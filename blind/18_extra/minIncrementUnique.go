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
			// how many times do we have to increment the value
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
// less efficient for small sample, more efficient for big sample
func minIncrement(nums []int) int {
	count := map[int]int{}
	incremented := 0
	largest := 0
	// fill that map
	for _, n := range nums {
		count[n]++
		if n > largest {
			largest = n
		}
	}

	for i := 0 ; i < len(nums) + largest ; i++ {
		if count[i] > 1 {
			extra := count[i] - 1
			count[i + 1] += extra
			incremented += extra
		}
	}

	return incremented
}
