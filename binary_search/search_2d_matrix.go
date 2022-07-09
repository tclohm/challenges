package main

import "fmt"

// binary search look 
func searchMatrix(matrix [][]int, target int) bool {
	top := 0
	bottom := len(matrix) - 1
	last := len(matrix) - 1
	first := 0

	for top < bottom {
		row := top + (bottom - top) / 2
		if target > matrix[row][last] {
			top = row + 1
		} else if target < matrix[row][first] {
			bottom = row - 1
		} else {
			break
		}
	}

	if top > bottom {
		return false
	}

	row := top + (bottom - top) / 2
	left := 0
	right := len(matrix[first]) - 1

	for left <= right {
		middle := left + (right - left) / 2
		if target > matrix[row][middle] {
			left = middle + 1
		} else if target < matrix[row][middle] {
			right = middle - 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}
	fmt.Println(searchMatrix(matrix, 3))
	fmt.Println(searchMatrix(matrix, 13))
}