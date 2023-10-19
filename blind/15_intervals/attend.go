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
	headStart, headEnd  := intervals[0].Start, intervals[0].End
	for i := 1 ; i < len(intervals) ; i++  {

		if intervals[i].End < headEnd { return false }
		//0 < 5 < 30 < 40
		if headStart < intervals[i].Start && intervals[i].Start < headEnd { return false }
		
		headStart = intervals[i].Start
		headEnd = intervals[i].End
	}

	return true
}

func main() {
	fmt.Println(canAttend([]*Interval{&Interval{0,30},&Interval{5,20},&Interval{15,25}}))
	fmt.Println(canAttend([]*Interval{&Interval{5,8},&Interval{9,15}}))
}