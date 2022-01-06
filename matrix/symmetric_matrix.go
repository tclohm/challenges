package main

import "fmt"

var matrix1 = [2][2] int{
	{4,5},
	{1,2},
}

func add(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var sum [2][2]int

	for i := 0 ; i < 2 ; i++ {
		for j := 0 ; j < 2 ; j++ {
			sum[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}

	return sum
}

func subtract(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var difference [2][2]int

	for i := 0 ; i < 2 ; i++ {
		for j := 0 ; j < 2 ; j++ {
			difference[i][j] = matrix1[i][j] - matrix2[i][j]
		}
	}

	return difference
}

func multiple(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var product [2][2]int

	for i := 0 ; i < 2 ; i++ {

		for j := 0 ; j < 2 ; j++ {

			var productSum int = 0

			for k := 0 ; k < 2 ; k++ {

				productSum += matrix1[i][k] * matrix2[k][j]

			}

			product[i][j] = productSum
		}

	}

	return product
}

func transpose(matrix [2][2]int) [2][2]int {
	var tmatrix [2][2]int
	for i := 0 ; i < 2 ; i++ {
		for j := 0 ; j < 2 ; j++ {
			tmatrix[i][j] = matrix[j][i]
		}
	}
	return tmatrix
}

// func determinant(matrix [2][2]int) float32 {
// 	var det float32
// 	det = det + ( (matrix[0][0] * matrix[1][1]) - (matrix[0][1] * matrix[1][0]) )
// 	return det
// }

// func inverse(matrix [2][2]int) [2][2]float64 {
// 	var det float64
// 	det = determinant(matrix)
// 	var invmatrix [][]float64

// 	invmatrix[0][0] = matrix[1][1]/det
// 	invmatrix[0][1] = -1*matrix[0][1]/det
// 	invmatrix[1][0] = -1*matrix[1][0]/det
// 	invmatrix[1][1] = matrix[0][0]/det

// 	return invmatrix
// }


func main() {

	var matrix1 = [2][2] int{
		{4, 5},
		{1, 2},
	}

	var matrix2 = [2][2] int{
		{6, 7},
		{3, 4},
	}

	var sum [2][2]int
	sum = add(matrix1, matrix2)

	fmt.Println("sum", sum)

	var difference [2][2]int
	difference = subtract(matrix1, matrix2)

	fmt.Println("subtraction", difference)

	var product [2][2]int
	product = multiple(matrix1, matrix2)

	fmt.Println("multiple", product)

	//var inv [2][2]float64
	//inv = inverse(matrix1)

	//fmt.Println("inverse", inv)
}