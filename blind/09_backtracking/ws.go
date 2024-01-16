package main

import "fmt"


func search(board [][]string, word string) bool {
	ROWS := len(board)
	COLS := len(board[0])

	index := 0

	var dfs func(row, col, index int) bool
	dfs = func(row, col, index int) bool {

		if string(word[index]) != board[row][col] { return false }
		board[row][col] = "."
		index++
		if index == len(word) - 1 { return true }
		
		if row + 1 < len(board) { dfs(row + 1, col, index) }
		if row - 1 >= 0 { dfs(row - 1, col, index) }
		if col + 1 < len(board[0]) { dfs(row, col + 1, index) }
		if col - 1 >= 0 { dfs(row, col - 1, index) }
		return false
	}

	for row := 0 ; row < ROWS ; row++ {
		for col := 0 ; col < COLS ; col++ {
			if board[row][col] == string(word[0]) {
				dfs(row, col, index)
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