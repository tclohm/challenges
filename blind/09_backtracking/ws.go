package main

import "fmt"


func search(board [][]string, word string) bool {
	ROWS := len(board)
	COLS := len(board[0])

	DIRS := [][]int{{0,1},{0,-1},{1,0},{-1,0}}

	var dfs func(row, col int)
	dfs = func(row, col int) {

		if board[row][col] != string(word[0]) { return }
		board[row][col] = "."
		word = word[1:]
		fmt.Println(word)
		for _, d := range DIRS {
			x, y := d[0], d[1]
			fmt.Println(x, y, row, col)
			if row + y < len(board) { dfs(row + y, col) }
			if row + y >= 0 { dfs(row + y, col) }
			if col + x < len(board[0]) { dfs(row, col + x) }
			if col + y >= 0 { dfs(row, col + y) }
		}
	}

	for row := 0 ; row < ROWS ; row++ {
		for col := 0 ; col < COLS ; col++ {
			if board[row][col] == string(word[0]) {
				dfs(row, col)
			}
		}
	}

	return len(word) == 0
}

func main() {
	fmt.Println(search([][]string{{"A","B","C","E"}, {"S","F","C","S"}, {"A","D","E","E"}}, "ABCCED"))
	fmt.Println(search([][]string{{"A","B","C","E"}, {"S","F","C","S"}, {"A","D","E","E"}}, "SEE"))
	fmt.Println(search([][]string{{"A","B","C","E"}, {"S","F","C","S"}, {"A","D","E","E"}}, "ABCB"))
}