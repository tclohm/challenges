package main

import (
	"fmt"
	"sort"
)
// n log n
func heightcheck(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	count := 0

	for i, h := range heights {
		if h != expected[i] {
			count++
		}
	}
	return count
}

func linearTimeAndSpace(heights []int) int {
	count := make([]int, 101, 101)
	for h := range heights {
		count[h] += 1
	}

	expected := []int{}	
	for i := 1 ; i < 101 ; i++ {
		c := count[i]
		x := 0
		for x < c {
			expected = append(expected, i)
			x++
		}
	}
	
	total := 0
	for i := 0 ; i < len(heights) ; i++ {
		if heights[i] != expected[i] {
			total++	
		}
	}
	return total
}

func main() {
	fmt.Println(heightcheck([]int{1,1,4,2,1,3}))
	fmt.Println(heightcheck([]int{5,1,2,3,4}))	
	fmt.Println(heightcheck([]int{1,2,3,4,5}))	
}
