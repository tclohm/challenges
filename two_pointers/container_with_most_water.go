package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	result := 0
	for left := 0 ; left < len(height) ; left++ {
		for right := left + 1 ; right < len(height) ; right++ {
			area := (right - left) * min(height[left], height[right])
			result = max(result, area)
		}
	}
	return result
}

func optMaxArea(height []int) int {
	result := 0

	left, right := 0, len(height) - 1

	for left < right {
		area := (right - left) * min(height[left], height[right])
		result = max(result, area)

		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return result
}

func main() {
	// - 8 - - - - 8 - 7
	h1 := []int{1,8,6,2,5,4,8,3,7}
	h2 := []int{1,1}

	fmt.Println(optMaxArea(h1))
	fmt.Println(optMaxArea(h2))
}