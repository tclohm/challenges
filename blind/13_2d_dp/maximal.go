package main

import "fmt"

// T : O(m * n) M : O(m * n)
func maximalSquare(matrix [][]string) int {
	ROWS, COLS := len(matrix), len(matrix[0])
	cache := map[string]int{}
 
	var fillCache func (row, col int) int
	fillCache = func (row, col int) int {
		if row >= ROWS || col >= COLS { return 0 }

		s := fmt.Sprintf("%d%d", row, col)
		if _, ok := cache[s]; !ok {
			down := fillCache(row + 1, col)
			right := fillCache(row, col + 1)
			diagonal := fillCache(row + 1, col + 1)

			cache[s] = 0
			if matrix[row][col] == "1" {
				cache[s] = 1 + min(down, right, diagonal)
			}
		}
		return cache[s]
	}

	fillCache(0, 0)

	max := 0

	for _, v := range cache {
		if v > max {
			max = v
		}
	}

    return max * max
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < a && b < c {
		return b
	}
	return c
}

func main() {
	fmt.Println(maximalSquare([][]string{{"1","0","1","0","0"},{"1","0","1","1","1"},{"1","1","1","1","1"},{"1","0","0","1","0"}}))
	fmt.Println(maximalSquare([][]string{{"0","1"},{"1","0"}}))
	fmt.Println(maximalSquare([][]string{{"0"}}))
}