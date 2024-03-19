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
	end := points[0][1]
	for i := 1 ; i < len(points) ; i++ {
		// overlapping?
		if points[i][0] > end {
			count++
			end = points[i][1]
		}
	}
	return count
}

func main() {
	fmt.Println(findMinArrowShots([][]int{{10,16},{2,8},{1,6},{7,12}}))
	fmt.Println(findMinArrowShots([][]int{{1,2},{3,4},{5,6},{7,8}}))
	fmt.Println(findMinArrowShots([][]int{{1,2},{2,3},{3,4},{4,5}}))
}