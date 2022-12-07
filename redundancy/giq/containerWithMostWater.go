package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func mostWater(container []int) int {
	var most = -1
	var left, right = 0, len(container) - 1
	var bookmark = []int{0,0}

	for left < right {
		width := right - left
		height := min(container[left], container[right])

		if (height * width) > most {
			most = height * width
			bookmark[0] = container[left]
			bookmark[1] = container[right]
		}

		if container[left] < container[right] {
			left += 1
		} else {
			right -= 1
		}
		
	}
	fmt.Println(bookmark)
	return most
}

func main() {
	levels := []int{7, 1, 2, 3, 9}

	fmt.Println(mostWater(levels))
}