package main

import ("fmt")

// lucky number is an element of the matrix that is the 
// min element in row and max in col

// is the number the highest number in it's row
// if it is, check if it is the smallest number in it's col
// if it both the min and max, add to our result array
func lucky(matrix [][]int) []int {

	maxs := []int{}
	mins := []int{}

	for row := 0 ; row < len(matrix) - 1 ; row++ {
		for col := 0 ; col < len(matrix[row]) - 1 ; col ++ {
		}
 	}
	if max(row) == min(col) {
		result = append(result, number)
	}

	return result
}


func main() {
	fmt.Println(lucky([][]int{{3,7,8},{9,11,13},{15,16,17}}))
	fmt.Println(lucky([][]int{{1,10,4,2},{9,3,8,7},{15,16,17,12}))
}
