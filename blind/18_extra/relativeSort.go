package main

import (
	"fmt"
	"sort"
)

func relSort(a, b []int) []int {
	count := make(map[int]int)
	set := make(map[int]bool)
	end := []int{}

	for _, n := range b {
		set[n] = true
	}

	for _, n := range a {
		if _, ok := set[n]; !ok {
			end = append(end, n)
		} else {
			count[n]++
		}
	}
	sort.Ints(end)

	result := []int{}

	for _, n := range b {
		for count[n] > 0 {
			result = append(result, n)
			count[n]--
		}
	}

	result = append(result, end...)
	return result
}

func main() {
	fmt.Println(relSort([]int{2,3,1,3,2,4,6,7,9,2,19}, []int{2,1,4,3,9,6}))
	fmt.Println(relSort([]int{28,6,22,8,44,17}, []int{22,28,8,6}))
}
