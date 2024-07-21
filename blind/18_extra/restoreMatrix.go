package main

import "fmt"

func restoreMatrix(rowSum, colSum []int) [][]int {
	var matrix = make([][]int, len(rowSum))

	for i := 0 ; i < len(rowSum) ; i++ {
		arr := make([]int, len(colSum))

		for j := 0 ; j < len(colSum) ; j++ {
			arr[j] = colSum[j]
			if colSum[j] > rowSum[i] {
				arr[j] = rowSum[i]
			}
			rowSum[i], colSum[j] = rowSum[i] - arr[i], colSum[j] - arr[j]
		}

		matrix[i] = arr
	}
	return matrix
	
	return matrix
}

func main() {
	fmt.Println(restoreMatrix([]int{3,8}, []int{4,7}))
}
