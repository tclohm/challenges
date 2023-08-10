package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	return max(robbing(nums[1:]), robbing(nums[:len(nums)-1]))
}

func robbing(nums []int) int {

	prev, current := 0, 0

	for _, n := range nums {
		prev, current = current, max(prev + n, current)
	}

	return current
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
}