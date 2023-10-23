package main

import "fmt"

func setZeroes(matrix [][]int) {
	
}

func main() {
	matrix := [][]int{{1,1,1},{1,0,1},{1,1,1}}
	setZeroes(matrix)
	fmt.Println(matrix)

	matrix2 := [][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}}
	setZeroes(matrix2)
	fmt.Println(matrix2)
}