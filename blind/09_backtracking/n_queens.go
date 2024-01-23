package main

import (
	"fmt"
	"strings"
)

func in(el int, ht map[int]struct{}) bool {
	if _, ok := ht[el]; ok {
		return true
	}
	return false
}

func solveNQueen(n int) [][]string {
	var columns = map[int]struct{}{}
	var positiveDiagonals = map[int]struct{}{}
	var negativeDiagonals = map[int]struct{}{}

	var result [][]string
	var board  = make([][]string, n, n)

	// make the board
	for r := 0 ; r < n ; r++ {
		board[r] = make([]string, n, n)
		for c := 0 ; c < n ; c++ {
			board[r][c] = "."
		}
	}

	var backtrack func (row int)
	backtrack = func (row int) {
		if row == n { 
			// copy board
			var copy = []string{}
			for r := 0 ; r < len(board) ; r++ {
				rowStrings := strings.Join(board[r], "")
				copy = append(copy, rowStrings)
			}
			result = append(result, copy)
			return
		}

		for col := 0 ; col < n ; col++ {
			// cant use position

			if in(col, columns) || in(row + col, positiveDiagonals) || in(row - col, negativeDiagonals) {
				continue
			}

			columns[col] = struct{}{}
			positiveDiagonals[row + col] = struct{}{}
			negativeDiagonals[row - col] = struct{}{}
			board[row][col] = "Q"

			backtrack(row + 1)

			delete(columns, col)
			delete(positiveDiagonals, row + col)
			delete(negativeDiagonals, row - col)
			board[row][col] = "."
		}
	}
	backtrack(0)
	return result
}

func main() {
	fmt.Println(solveNQueen(4))
	fmt.Println(solveNQueen(1))
	fmt.Println(solveNQueen(3))
	fmt.Println(solveNQueen(2))
	fmt.Println(solveNQueen(5))
}