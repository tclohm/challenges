package main

import "fmt"

func max(nums []int) int {
	big := -10000000
	for _, i := range nums {
		if big < i {
			big = i
		}
	}
	return big
}

func slidingWindowMax(nums []int, k int) []int {
	result := []int{}
	end_buffer := len(nums) - k + 1
	for left := 0 ; left < end_buffer ; left++ {
		right := left + k
		max := max(nums[left:right])
		result = append(result, max)
	}
	return result
}

func maxSlidingWindow(nums []int, k int) [] int {
	deque := make([]int, 0, k)
	result := make([]int, 0, len(nums) - k + 1)

	for i, n := range nums {
		// push
		// remove all indices from deque that corresponds
		// to numbers less than number we are trying to push
		for len(deque) > 0 && nums[deque[len(deque) - 1]] < n {
			deque = deque[:len(deque) - 1]
		}
		deque = append(deque, i)

		if i < k - 1 { continue }

		result = append(result, nums[deque[0]])

		if deque[0] == i - k + 1 {
			deque = deque[1:]
		}
	}

	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
}