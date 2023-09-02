package main

import "fmt"

func trap(heights []int) int {
	if heights == nil {
		return 0
	}

	left, right := 0, len(heights) - 1
	leftMax, rightMax := heights[left], heights[right]
	water := 0

	for left < right {
		if leftMax < rightMax {
			left++
			leftMax = max(leftMax, heights[left])
			water += leftMax - heights[left]
		} else {
			right--
			rightMax = max(rightMax, heights[right])
			water += rightMax - heights[right]
		}
	}
	return water
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trap([]int{4,2,0,3,2,5}))
}