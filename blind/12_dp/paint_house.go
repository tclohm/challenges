package main

import "fmt"

// paint all the houses, but no two adjacent houses can be the same color
func paint(costs [][]int) int {

	var result = 1000

	var painting func(index, paint int) int
	painting = func(index, paint int) int {
		if index == len(costs) { return 0 }
		var smaller = 1000
		for i, p := range costs[index] {
			if i != paint {
				smaller = min(p + painting(index + 1, i), smaller)
			}
		}
		return smaller
	}

	result = min(painting(0, 0), result)
	result = min(painting(0, 1), result)
	result = min(painting(0, 2), result)

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// red, blue, green
	fmt.Println(paint([][]int{{17,2,17},{16,16,5},{14,3,19}}))
}