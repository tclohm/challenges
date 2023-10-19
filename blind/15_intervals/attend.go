package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start, End int
}

func canAttend(intervals []*Interval) bool {
	sort.Slice(intervals, func(a, b int) bool { return intervals[a].Start < intervals[b].Start })
	headEnd  := intervals[0].End
	for i := 1 ; i < len(intervals) ; i++  {
		if intervals[i].End < headEnd { return false }
		headEnd = intervals[i].End
	}

	return true
}

func main() {
	fmt.Println(canAttend([]*Interval{&Interval{0,30},&Interval{5,10},&Interval{15,20}}))
	fmt.Println(canAttend([]*Interval{&Interval{5,8},&Interval{9,15}}))
}