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

func longest_increasing_sub(sequence []int) int {
	var solve func(index int, prev int) int
	solve = func(index int, prev int) int {
		if index >= len(sequence) { return 0 }
		take := 0
		dont := solve(index + 1, prev)
		if sequence[index] > prev {
			take = 1 + solve(index + 1, sequence[index])
		}
		return bigger(take, dont)
	}
	return solve(0, -10)
}

// are we going to include it or not include it : 2 choices
func main() {
	fmt.Println(longest_increasing_sub([]int{10,9,2,5,3,7,101,18}))
	fmt.Println(longest_increasing_sub([]int{0,1,0,3,2,3}))
	fmt.Println(longest_increasing_sub([]int{7,7,7,7,7,7,7}))
}