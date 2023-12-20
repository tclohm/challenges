package main

import "fmt"

func smooth(img [][]int) [][]int {

	ROWS, COLS := len(img), len(img)
	result := make([][]int, COLS, ROWS)
	for i := 0 ; i < COLS ; i++ {
		result[i] = make([]int, COLS, ROWS)
	}

	for r := 0 ; r < ROWS; r++ {
		for c := 0 ; c < COLS ; c++ {
			total, count := 0, 0
			for i := r - 1 ; i < r + 2 ; i++ {
				for j := c - 1 ; j < c + 2 ; j++ {
					if i < 0 || i == ROWS || j < 0 || j == COLS {
						continue
					}
					total += img[i][j]
					count += 1
				}
			}
			result[r][c] = total / count
		}
	}

	return result
}

func main() {
	fmt.Println(smooth([][]int{{1,1,1},{1,0,1},{1,1,1}}))
}