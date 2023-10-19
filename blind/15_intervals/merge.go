package main

import (
	"fmt"
	"sort"
)

type double [][]int

func merge(intervals double) [][]int {
	
	start := 0
	end := 1

	sort.SliceStable(intervals, func(i, k int) bool {
		return intervals[i][start] < intervals[k][start]
	})

	result := [][]int{intervals[0]}

	for i := 0 ; i < len(intervals) ; i++ {
		lastIndex := len(result) - 1
		currentEnd := result[lastIndex]
		current := intervals[i]

		if currentEnd[end] < current[start] {
			result = append(result, current)
		} else if currentEnd[end] < current[end] {
			result[lastIndex][end] = current[end]
		}
	}

	return result
}

func main() {
	fmt.Println(merge([][]int{{1,3},{2,6},{8,10},{15,18}}))
	fmt.Println(merge([][]int{{1,4},{4,5}}))
}