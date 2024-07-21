package main

import "fmt"

func restoreMatrix(rowSum, colSum []int) [][]int {
	var matrix = [][]int{}
	for i := 0 ; i < len(rowSum) ; i++ {
		matrix = append(matrix, make([]int, len(rowSum), len(rowSum)))
	}
	return matrix
}

func main() {
	fmt.Println(restoreMatrix([]int{3,8}, []int{4,7}))
}
