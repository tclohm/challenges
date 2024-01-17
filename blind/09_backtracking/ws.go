package main

import "fmt"


func search(board [][]string, word string) bool {
	ROWS := len(board)
	COLS := len(board[0])

	index := 0

	var dfs func(row, col, index int) bool
	dfs = func(row, col, index int) bool {
		if row < 0 || col < 0 || row >= ROWS || col >= COLS || index == len(word) { return false }

		if board[row][col] != string(word[index]) || board[row][col] == "." { return false }

		if index == len(word) - 1 { return true }

		letter := board[row][col]

		exist := dfs(row + 1, col, index + 1) || dfs(row - 1, col, index + 1) || dfs(row, col + 1, index + 1) || dfs(row, col - 1, index + 1)

		board[row][col] = letter

		return exist

	}

	for row := 0 ; row < ROWS ; row++ {
		for col := 0 ; col < COLS ; col++ {
			if board[row][col] == string(word[0]) {
				if dfs(row, col, index) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	fmt.Println(search([][]string{{"A","B","C","E"}, {"S","F","C","S"}, {"A","D","E","E"}}, "ABCCED"))
	fmt.Println(search([][]string{{"A","B","C","E"}, {"S","F","C","S"}, {"A","D","E","E"}}, "SEE"))
	fmt.Println(search([][]string{{"A","B","C","E"}, {"S","F","C","S"}, {"A","D","E","E"}}, "ABCB"))
}