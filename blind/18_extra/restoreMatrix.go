package main

import "fmt"

func restoreMatrix(rowSum, colSum []int) [][]int {
	var matrix = [][]int{rowSum, colSum}
	return matrix
}

func main() {
	fmt.Println(restoreMatrix([]int{3,8}, []int{4,7}))
}
