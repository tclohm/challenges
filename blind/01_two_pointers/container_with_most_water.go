package main

import "fmt"


func maxArea(heights []int) int {
	max_area := -1
	left, right := 0, len(heights) - 1
	for left < right {
		length := right - left
		// area = height * width
		area := min(heights[left], heights[right]) * length

		max_area = max(area, max_area)

		if heights[left] <= heights[right] {
			left++
		} else {
			right--
		}
	}

	return max_area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
	fmt.Println(maxArea([]int{1,1}))
}