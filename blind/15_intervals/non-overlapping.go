package main

import (
	"fmt"
	"sort"
)

func eraseOverlap(intervals [][]int) int {
	start := 0
	end := 1

	sort.Slice(intervals, func(a, b int) bool { return intervals[a][start] < intervals[b][start] })

	removal := 0

	currentArrayEnd := intervals[0][end]

	for i := 1 ; i < len(intervals) ; i++ {
		currentArray := intervals[i]

		if currentArray[start] < currentArrayEnd {
			removal += 1
			currentArrayEnd = min(currentArray[end], currentArrayEnd)
		} else {
			currentArrayEnd = currentArray[end]
		}
	}

	return removal
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func main() {
	fmt.Println(eraseOverlap([][]int{{1,2},{2,3},{3,4},{1,3}}))
	fmt.Println(eraseOverlap([][]int{{1,2},{1,2},{1,2}}))
	fmt.Println(eraseOverlap([][]int{{1,2},{2,3}}))
}