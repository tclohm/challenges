package main

import "fmt"

type matrix [][]int
type Seen [][]bool

func (s *Seen) fill(columnLength, rowLength int) {
	for i := 0 ; i < columnLength ; i++ {
		var inner = make([]bool, rowLength, rowLength)
		*s = append(*s, inner)
	}
}

func dfs(m matrix, row, col int, s Seen, v *[]int) {
	// up, down, left, right
	directions := [][]int{{-1,0}, {1,0}, {0,-1}, {0,1},}
	if row < 0 || col < 0 || row >= len(m) ||
	col >= len(m[0]) || s[row][col] {
		return
	}
	*v = append(*v, m[row][col])
	s[row][col] = true
	fmt.Println(row, col)
	for i := 0 ; i < len(directions) ; i++ {
		var currentDir = directions[i]
		dfs(m, row + currentDir[0], col + currentDir[1], s, v)
	}
}

func traverseDFS(matrix matrix) []int {
	var seen = Seen{}
	rl, cl := len(matrix[0]), len(matrix)
	seen.fill(cl, rl)
	var values = []int{}
	dfs(matrix, 0, 0, seen, &values)
	return values
}

func main() {

	ourMap := [][]int{
		{1, 2, 3, 4, 5}, 
		{6, 7, 8, 9, 10}, 
		{11, 12, 13, 14, 15}, 
		{16, 17, 18, 19, 20},
	}

	fmt.Println(traverseDFS(ourMap))
}