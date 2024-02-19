package main

import "fmt"

func lengthOfLIS(nums []int) int {
	var cache = make([]int, len(nums), len(nums))
	for i := 0 ; i < len(cache) ; i++ {
		cache[i] = 1
	}

	for i := 0 ; i < len(nums) ; i++ {
		for j := 0 ; j < i ; j++ {
			if nums[i] > nums[j] {
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
		if index >= len(sequence) { return 0 } // cant pick any more
		take := 0 
		dont := solve(index + 1, prev) // try skipping current element
		if sequence[index] > prev {
			take = 1 + solve(index + 1, sequence[index]) // or pick it if it is greater than previous picked element
		}
		return bigger(take, dont)
	}
	return solve(0, -10)
}

func nlogn(nums []int) int {
	tail := []int{}
	index := 0

	for _, n := range nums {
		index = binary(tail, n)
		if index == len(tail) {
			tail = append(tail, n)
		} else {
			tail[index] = n
		}
	}

	return len(tail)
} 

func binary(tail []int, target int) int {
	n := len(tail)
	left, right := 0, n
	for left < right {
		mid := left + (right - left) / 2
		if tail[mid] < target {
			left = mid + 1
		} else if tail[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return left
}
// are we going to include it or not include it : 2 choices
func main() {
	fmt.Println(nlogn([]int{10,9,2,5,3,7,101,18}))
	fmt.Println(nlogn([]int{0,1,0,3,2,3}))
	fmt.Println(nlogn([]int{7,7,7,7,7,7,7}))
}