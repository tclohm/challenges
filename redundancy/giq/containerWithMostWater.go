package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mostWater(heights []int) int {
	var most = -1
	var left, right = 0, len(heights) - 1

	for left < right {
		// width * height = area
		// height has to be the min, no overflows
		width := right - left
		height := min(heights[left], heights[right])
		// is new area bigger?
		if (height * width) > most {
			// new most
			most = height * width
		}
		// is left less than right, increment left
		// else increment right
		if heights[left] < heights[right] {
			left += 1
		} else {
			right -= 1
		}
		
	}
	return most
}

func main() {
	levels := []int{7, 1, 2, 3, 9}

	fmt.Println(mostWater(levels))
}