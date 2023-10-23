package main

import "fmt"

func rotate(matrix [][]int) {
	left, right := 0, len(matrix) - 1
	for left < right {
		for index := 0 ; index < right - left ; index++ {
			top, bottom := left, right

			topLeft := matrix[top][left + index]
			// move bottom left to top left
			matrix[top][left + index] = matrix[bottom - index][left]
			// bottom  right to bottom left
			matrix[bottom - index][left] = matrix[bottom][right - index]
			// top right to bottom right
			matrix[bottom][right - index] = matrix[top + index][right]
			// top left to top right
			matrix[top + index][right] = topLeft
		}
		right -= 1
		left += 1
	}

	fmt.Println(matrix)
}

func main() {
	rotate([][]int{{1,2,3},{4,5,6},{7,8,9}})
	//rotate([][]int{{5,1,9,11},{2,4,8,10},{3,3,6,7},{15,14,12,16}})
}