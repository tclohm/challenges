package main

import "fmt"

func restoreMatrix(rowSum, colSum []int) [][]int {
	var matrix = [][]int{}

	for col := 0 ; col < len(colSum) ; col++ {
		matrix = append(matrix, make([]int, len(colSum), len(colSum)))
	}

	for row := 0 ; row < len(rowSum) ; row++ {
		matrix[row][0] = rowSum[row]
	}

	for col := 0 ; col < len(colSum) ; col++ {
		currentColSum := 0
		for row := 0 ; row < len(rowSum) ; row++ {
			currentColSum += matrix[row][col]
		}

		row := 0
		for currentColSum > colSum[col] {
			difference := currentColSum - colSum[col]
			greedyShift := min(matrix[row][col], difference)

			matrix[row][col] -= greedyShift
			matrix[row][col + 1] += greedyShift
			currentColSum -= greedyShift
			row += 1 
		}
	}

	return matrix
}

func main() {
	fmt.Println(restoreMatrix([]int{3,8}, []int{4,7}))
	fmt.Println(restoreMatrix([]int{5,7,10}, []int{8,6,8}))
}
