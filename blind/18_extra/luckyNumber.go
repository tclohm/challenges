package main

import ("fmt")

// lucky number is an element of the matrix that is the 
// min element in row and max in col

// is the number the highest number in it's row
// if it is, check if it is the smallest number in it's col
// if it both the min and max, add to our result array
func lucky(matrix [][]int) []int {
	var result = []int{}
	for row := 0 ; row < len(matrix) ; row++ {
		col, smallest := min(matrix[row])
		biggest := max(matrix, col)
		if biggest == smallest {
			result = append(result, smallest)
		}
	}
	return result
}

func min(rows []int) (int, int) {
	smallest := 100000
	col := 0
	for i, n := range rows {
		if n < smallest {
			smallest = n
			col = i
		}
	}
	// return number
	return col, smallest
}

func max(matrix [][]int, col int) int {
	big := -10000
	for row := 0 ; row < len(matrix) ; row++ {
		if matrix[row][col] > big {
			big = matrix[row][col]
		}
	}
	return big
}

func main() {
	fmt.Println(lucky([][]int{{3,7,8},{9,11,13},{15,16,17}}))
	fmt.Println(lucky([][]int{{1,10,4,2},{9,3,8,7},{15,16,17,12}}))
	fmt.Println(lucky([][]int{{7,8},{1,2}}))
}
