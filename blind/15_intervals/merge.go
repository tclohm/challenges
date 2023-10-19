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
		lastArrayIndex := len(result) - 1
		resultEndArray := result[lastArrayIndex]
		currentArray := intervals[i]
		// if the current arrays start is greater than results arrays end,
		// we append				
		if resultEndArray[end] < currentArray[start] {
			result = append(result, currentArray)
		// if the results array last number is less than current arrays end,
		// change the end
		} else if resultEndArray[end] < currentArray[end] {
			result[lastArrayIndex][end] = currentArray[end]
		}
	}

	return result
}

func main() {
	fmt.Println(merge([][]int{{1,3},{2,6},{8,10},{15,18}}))
	fmt.Println(merge([][]int{{1,4},{4,5}}))
}