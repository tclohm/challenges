package main

import "fmt"

func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])

	// up/down, left/right, current_square
	var dfs func(i int, j int, curr int) bool

	dfs = func(i int, j int, curr int) bool {
		// check whether i and j are out of bounds and curr is the length of word
		if i < 0 || j < 0 || i >= n || j >= m || curr == len(word) {
			return false
		}
		// if current word doesn't equal word or is *
		if board[i][j] != word[curr] || board[i][j] == '*' {
			return false
		}
		// if curr == length
		if curr == len(word)-1 {
			return true
		}
		// hold the letter
		tmp := board[i][j]
		board[i][j] = '*'

		// look down,					look up,				look left, 				look right
		res := dfs(i+1, j, curr+1) || dfs(i-1, j, curr+1) || dfs(i, j-1, curr+1) || dfs(i, j+1, curr+1)

		// place the char back after dfs
		board[i][j] = tmp

		return res
	}

	// start the traverse
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// dfs at each point
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	// return false if we can't reach true
	return false
}

