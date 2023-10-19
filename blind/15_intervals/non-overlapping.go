package main

import "fmt"

func eraseOverlap(intervals [][]int) int {
	return 1
}

func main() {
	fmt.Println(eraseOverlap([][]int{{1,2},{2,3},{3,4},{1,3}}))
	fmt.Println(eraseOverlap([][]int{{1,2},{1,2},{1,2}}))
	fmt.Println(eraseOverlap([][]int{{1,2},{2,3}}))
}