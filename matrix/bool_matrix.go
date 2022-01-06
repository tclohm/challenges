package main

import "fmt"

func change(matrix [3][3]int) [3][3]int {
	var Rows [3]int
	var Columns [3]int

	var matrix_changed [3][3]int

	for i := 0 ; i < 3 ; i++ {
		for j := 0 ; j < 3 ; j++ {
			if matrix[i][j] == 1 {
				Rows[i] = 1
				Columns[j] = 1
			}
		}
	}

	for i := 0 ; i < 3 ; i++ {
		for j := 0 ; j < 3 ; j++ {
			if Rows[i] == 1 || Columns[j] == 1 {
				matrix_changed[i][j] = 1
			}
		}
	}

	return matrix_changed
}

func printMatrix(matrix [3][3]int) {
	for i := 0 ; i < 3 ; i++ {

		for j := 0 ; j < 3 ; j++ {

			fmt.Printf("%d", matrix[i][j])

		}

		fmt.Printf("\n")

	}

}

func main() {
	var matrix = [3][3]int {{1,0,0},{0,0,0},{0,0,0}}
	printMatrix(matrix)
	matrix = change(matrix)
	printMatrix(matrix)
}