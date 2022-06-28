package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func trap(n []int) int {
	if len(n) < 2 {
		return 0
	}

	var left, right = 0, len(n) - 1
	var maxLeft, maxRight = n[left], n[right]

	var result int = 0

	for left < right {
		if maxLeft < maxRight {
			left += 1
			maxLeft = max(maxLeft, n[left])
			result += maxLeft - n[left]
		} else {
			right -= 1
			maxRight = max(maxRight, n[right])
			result += maxRight - n[right]
		}
	}
	return result
}

func main() {
	h := []int {0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(h))
}