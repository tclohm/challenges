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

func part(nums []int, k int) bool {
	if sum(nums) % k == 1 {
		return false
	}
	sort.Ints(nums)
	var target = sum(nums) / k
	var used = make([]bool, len(nums), len(nums))

	var backtrack func(index int, k int, subsetSum int) bool
	backtrack = func(index int, k int, subsetSum int) bool {
		
		if k == 0 {
			return true
		}

		if subsetSum == target {
			return backtrack(0, k - 1, 0)
		}

		for j := index ; j < len(nums) ; j++ {
			if used[j] || subsetSum + nums[j] > target {
				continue
			} else {
				used[j] = true
				if backtrack(j + 1, k, subsetSum + nums[j]) {
					return true
				}
				used[j] = false
			}
		}
		return false
	}

	return backtrack(0, k, 0)
}

func main() {
	fmt.Println(partition([]int{4,3,2,3,5,2,1}, 4))
	fmt.Println(partition([]int{1,2,3,4}, 3))

	fmt.Println(part([]int{4,3,2,3,5,2,1}, 4))
	fmt.Println(part([]int{1,2,3,4}, 3))
}