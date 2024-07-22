package main

import "fmt"

func restoreMatrix(rowSum, colSum []int) [][]int {
	var matrix = make([][]int, len(rowSum))

	for row := 0 ; row < len(rowSum) ; row++ {
		arr := make([]int, len(colSum))

		for col := 0 ; col < len(colSum) ; col++ {
			arr[col] = colSum[col]
			if colSum[col] > rowSum[row] {
				arr[col] = rowSum[row]
			}
			rowSum[row], colSum[col] = rowSum[row] - arr[row], colSum[col] - arr[col]
		}

		matrix[row] = arr
	}
	return matrix
}

func main() {
	fmt.Println(restoreMatrix([]int{3,8}, []int{4,7}))
}
