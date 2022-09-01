package main

var directions = [][]int{
	[]int{0, 1},
	[]int{0, -1},
	[]int{1, 0},
	[]int{-1, 0},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))

	for y := 0 ; y < len(board) ; y++ {
		visited[y] = make([]bool, len(board[y]))
	}

	for y := 0 ; y < len(board) ; y++ {
		for x := 0 ; x < len(board[y]) ; x++ {
			if dfs(board, visited, word, y, x) { return true }
		}
	}

	return false
}

func dfs(board [][]byte, visited [][]bool, word string, y, x int) bool {
	if len(board) == 0 { return true }

	if y < 0 || x < 0 || y >= len(board) || x >= len(board[0]) || visited[y][x] || board[y][x] != word[0] {
		return false
	}

	visited[y][x] = true
	for _, d := range directions {
		if dfs(board, visited, word[1:], y+d[0], x+d[1]) {
			return true
		}
	}
	visited[y][x] = false
	return false
}