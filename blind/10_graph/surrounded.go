package main

import "fmt"

// m x m board containing 'X' and 'O' capture all regions that are 4 dir
// surrounded by 'X'
func solve(board [][]byte) {
	ROWS := len(board)
	COLS := len(board[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(board) || c >= len(board[0]) || board[r][c] == 'X' || board[r][c] == '*' { return }
		board[r][c] = '*'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r+1, c-1)
	}
	for r := 0 ; r < ROWS ; r++ {
		dfs(0, r)
		dfs(ROWS-1, r)
	}

	for c := 0 ; c < COLS ; c++ {
		dfs(0, c)
		dfs(c, COLS-1)
	}

	for r := 0 ; r < ROWS ; r++ {
		for c := 0 ; c < COLS ; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}

			if board[r][c] == '*' {
				board[r][c] = 'O'
			}
		}
	}
}

func main() {
	b := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(b)
	fmt.Println(b)
}