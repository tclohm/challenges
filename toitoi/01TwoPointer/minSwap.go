package main

import "fmt"

func count(nums []int, k int) int {
	var t int = 0
	for _, n := range nums {
		if n == k {
			t++
		}
	}
	return t
}

func minSwap(nums []int) int {
	var (
		N = len(nums)
		totalOnes = count(nums, 1)
		windowOnes = 0
		maxWindow = 0
		left = 0
	)
	// I did not know you could do this
	for right := range 2 * N {
		if nums[right % N] == 1 {
			windowOnes += 1
		}
		if right - left + 1 > totalOnes {
			windowOnes -= nums[left % N]
			left += 1
		}
		maxWindow = max(maxWindow, windowOnes)
	}

	return totalOnes - maxWindow
}

func main() {
	fmt.Println(minSwap([]int{0,1,0,1,1,0,0}))
	fmt.Println(minSwap([]int{0,1,1,1,0,0,1,1,0}))
	fmt.Println(minSwap([]int{1,1,0,0,1}))
}
