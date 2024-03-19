package main

import (
	"fmt"
	"sort"
)

// 2D int array
// point = [start_x, end_x]

// array shot at x if start_x <= x end_x
// no limit of # of arrows shot

// return the min number of arrows that must be shot to burst all balloons

// sort the slice of points by x_end in ascending order
// simulate an arrow that shoots positions with x = x_end 
// of current point[i] and greedily pass through all points
// that are also shot by that arrow

// time : O(n log n) , space : O(1)
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func (i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 1
	prev := points[0]
	for i := 1 ; i < len(points) ; i++ {
		current := points[i]
		// start overlapping with prev end?
		if current[0] > prev[1] {
			count++
			prev = current
		}
	}
	return count
}

func main() {
	fmt.Println(findMinArrowShots([][]int{{10,16},{2,8},{1,6},{7,12}}))
	fmt.Println(findMinArrowShots([][]int{{1,2},{3,4},{5,6},{7,8}}))
	fmt.Println(findMinArrowShots([][]int{{1,2},{2,3},{3,4},{4,5}}))
}