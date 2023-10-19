package main

import "fmt"

type double [][]int

func insert(intervals double, newInterval []int) double {
	result := double{}
	for i, array := range intervals {
		currentStart, currentEnd := array[0], array[1]
		newStart, newEnd := newInterval[0], newInterval[1]

		if newEnd < currentStart {
			
			result = append(result, newInterval)
			result = append(result, intervals[i:]...)
			return result
		
		} else if newStart > currentEnd {

			result = append(result, array)

		} else {
			newInterval[0] = min(newStart, currentStart)
			newInterval[1] = max(newEnd, currentEnd)
		}

	}

	return intervals
}

func max(a, b int) int {
	if a > b { return a }
	return b 
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func main() {
	fmt.Println(insert(double{{1,3},{6,9}}, []int{2,5}))
}