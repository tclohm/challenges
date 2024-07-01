package main

import (
	"fmt"
	"sort"
)
func maxImportance(n int, roads [][]int) int {
	var (
		edgeCount = make([]int, n, n)
		label int = 1
		res int = 0
	)

	for i, _ := range roads {
		n1 := roads[i][0]
		n2 := roads[i][1]

		edgeCount[n1] += 1
		edgeCount[n2] += 1
	}

	sort.Ints(edgeCount)

	for _, count := range edgeCount {
		res += label * count
		label += 1
	}

	return res 
}
func main() {
	fmt.Println(maxImportance(5, [][]int{{0,1},{1,2},{2,3},{0,2},{1,3},{2,4}}))
	fmt.Println(maxImportance(5, [][]int{{0,1},{2,4},{1,3}}))
}
