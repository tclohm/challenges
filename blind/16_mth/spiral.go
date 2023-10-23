package main

import "fmt"

func spiral(matrix [][]int) []int {
	result := []int{}

	left, right := 0, len(matrix[0]) - 1
	top, bottom := 0, len(matrix) - 1

	for left <= right && top <= bottom {
		for column := left ; column <= right ; column++ {
			result = append(result, matrix[top][column])
		}

		top++

		for row := top ; row <= bottom ; row++ {
			result = append(result, matrix[row][right])
		}
		
		right--

		if left > right || top > bottom { break }
		
		for column := right ; column >= left ; column-- {
			result = append(result, matrix[bottom][column])
		}
		
		bottom--

		for row := bottom ; row >= top ; row-- {
			result = append(result, matrix[row][left])
		}
		
		left++
	}

	return result
}

func main() {
	fmt.Println(spiral([][]int{{1,2,3},{4,5,6},{7,8,9}}))
}