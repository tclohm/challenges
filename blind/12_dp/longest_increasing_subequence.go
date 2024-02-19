package main

import "fmt"

func lengthOfLIS(nums []int) int {
	var cache = make([]int, len(nums), len(nums))
	for i := 0 ; i < len(cache) ; i++ {
		cache[i] = 1
	}

	for i := len(nums) - 1 ; i >= 0 ; i-- {
		for j := i + 1 ; j < len(nums) ; j++ {
			if nums[i] < nums[j] {
				cache[i] = bigger(cache[i], 1 + cache[j])
			}
		}
	}

	return max(cache)
}

func bigger(a, b int) int {
	if a > b { return a }
	return b
}

func max(items []int) int {
	var largest = -100
	for _, n := range items {
		if n > largest {
			largest = n
		}
	}
	return largest
}

func main() {
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))
	fmt.Println(lengthOfLIS([]int{0,1,0,3,2,3}))
	fmt.Println(lengthOfLIS([]int{7,7,7,7,7,7,7}))
}