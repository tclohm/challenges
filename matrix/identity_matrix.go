package main

import ("fmt")

func Identity(order int) [][]float64 {
	var matrix [][]float64
	matrix = make([][]float64, order)

	for i := 0 ; i < order ; i++ {
		tmp := make([]float64, order)
		tmp[i] = 1
		matrix[i] = tmp
	}
	return matrix
}

func main() {
	fmt.Println(Identity(4))
}