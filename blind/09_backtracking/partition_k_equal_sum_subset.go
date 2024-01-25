package main

import (
	"fmt"
	"sort"
)

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func partition(nums []int, k int) bool {
	sort.Ints(nums)
	target := sum(nums) / k

	if nums[len(nums) - 1] > target { return false }

	start, end := 0, len(nums) - 1
	result := 0

	for start < end {

		if nums[end] == target {
			result += 1
			end -= 1
			continue
		}

		added := nums[start] + nums[end]

		if added == target {
			result += 1
			start += 1
			end -= 1
			continue
		}
		if added > target {
			end -= 1
			continue
		}
		if added < target {
			start += 1
			continue
		}

	}
	return result == k
}

func main() {
	fmt.Println(partition([]int{4,3,2,3,5,2,1}, 4))
	fmt.Println(partition([]int{1,2,3,4}, 3))
}