package main

import "fmt"

func islands(board [][]string) int {
	ROWS := len(board)
	COLS := len(board[0])
	if ROWS == 0 { return 0 }
	islands := 0

	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= ROWS || c < 0 || c >= COLS || board[r][c] == "0" { return }

		if board[r][c] == "1" {
			board[r][c] = "0"
			dfs(r + 1, c)
			dfs(r - 1, c)
			dfs(r    , c + 1)
			dfs(r    , c - 1)
		}
	}

	for r := 0 ; r < ROWS ; r++ {
		for c := 0 ; c < COLS ; c++ {
			if board[r][c] == "1" {
				islands += 1
				dfs(r,c)
			}
		}
	}

	return islands
}

func main() {
	fmt.Println(islands([][]string{{"1","1","1","1","0"}, {"1","1","0","1","0"}, {"1","1","0","0","0"}, {"0","0","0","0","0"}}))
	fmt.Println(islands([][]string{{"1","1","0","0","0"}, {"1","1","0","0","0"}, {"0","0","1","0","0"}, {"0","0","0","1","1"}}))
}