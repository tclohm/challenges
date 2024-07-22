package main

import "fmt"

func restoreMatrix(rowSum, colSum []int) [][]int {
	var matrix = [][]int{}

	for col := 0 ; col < len(colSum) ; col++ {
		matrix = append(matrix, make([]int, len(colSum), len(colSum)))
	}

	row, col := 0, 0
	for row < len(rowSum) && col < len(colSum) {
		matrix[row][col] = min(rowSum[row], colSum[col])

		rowSum[row] -= matrix[row][col]
		colSum[col] -= matrix[row][col]

		if rowSum[row] == 0 {
			row++
		} else {
			col++
		}
	}
	
	return matrix
}

func main() {
	fmt.Println(restoreMatrix([]int{3,8}, []int{4,7}))
	fmt.Println(restoreMatrix([]int{5,7,10}, []int{8,6,8}))
}
