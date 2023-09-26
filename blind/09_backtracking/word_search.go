package main

import "fmt"

func exist(board [][]byte, word string) bool {
	ROWS := len(board)
	COLS := len(board[0])

	var dfs func (int, int, int) bool
	dfs = func (row, col, current int) bool {
		if row < 0 || col < 0 || row >= ROWS || col >= COLS || current == len(word) {
			return false
		}
		if board[row][col] != word[current] || board[row][col] == '*' {
			return false
		}
		if current == len(word) - 1 {
			return true
		}

		char := board[row][col]
		board[row][col] = '*'

		result := dfs(row + 1, col, current + 1) || dfs(row - 1, col, current + 1) || dfs(row, col + 1, current + 1) || dfs(row, col - 1, current + 1)
		
		board[row][col] = char

		return result
	}

	for r := 0 ; r < ROWS ; r++ {
		for c := 0 ; c < COLS ; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(exist([][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "ABCCED"))
	fmt.Println(exist([][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "SEE"))
}