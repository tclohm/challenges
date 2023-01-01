package main

import "fmt"

type matrix [][]int
type Seen [][]bool

type Queue struct {
	Coords [][]int
	Size int
}

func (q *Queue) Push(coord []int) {
	q.Coords = append(q.Coords, coord)
	q.Size += 1
}

func (q *Queue) Pop() (int, int) {
	var popped = q.Coords[0]
	row, col := popped[0], popped[1]
	q.Coords = q.Coords[1:]
	q.Size -= 1
	return row, col
}

func (q *Queue) Length() int {
	return q.Size
}

func (s *Seen) fill(columnLength, rowLength int) {
	for i := 0 ; i < columnLength ; i++ {
		var inner = make([]bool, rowLength, rowLength)
		*s = append(*s, inner)
	}
}

func dfs(m matrix, row, col int, s Seen, v *[]int) {
	// right, down, left, up
	directions := [][]int{{0,1}, {1,0}, {0,-1}, {-1,0},}
	if row < 0 || col < 0 || row >= len(m) ||
	col >= len(m[0]) || s[row][col] {
		return
	}
	*v = append(*v, m[row][col])
	s[row][col] = true
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

func traverseBFS(matrix matrix) []int {
	directions := [][]int{{0,1}, {1,0}, {0,-1}, {-1,0},}
	var seen = Seen{}
	rl, cl := len(matrix[0]), len(matrix)
	seen.fill(cl, rl)
	var values = []int{}

	var q = Queue{Size: 0}
	q.Push([]int{0,0})

	for q.Length() != 0 {
		row, col := q.Pop()
		if row < 0 || 
		row >= len(matrix) || 
		col < 0 || 
		col >= len(matrix[0]) ||
		seen[row][col] {
			continue
		}
		seen[row][col] = true
		values = append(values, matrix[row][col])

		for i := 0 ; i < len(directions) ; i++ {
			currentDirection := directions[i]
			q.Push([]int{row + currentDirection[0], col + currentDirection[1]})
		} 
	}
	return values
}

// 1's land
// 0's water
func countIslands(matrix matrix) int {
	var directions = [][]int{{0,1}, {1,0}, {0,-1}, {-1,0},}
	var islands = 0

	for row := 0 ; row < len(matrix) ; row++ {
		for col := 0 ; col < len(matrix[0]) ; col++ {
			if matrix[row][col] == 1 {
				islands += 1
				matrix[row][col] = 0

				var queue = Queue{Size: 0}
				queue.Push([]int{row, col})

				for queue.Length() != 0 {
					var currentRow, currentCol = queue.Pop()

					for i := 0 ; i < len(directions) ; i++ {
						var currentDir = directions[i]
						var nextRow = currentRow + currentDir[0]
						var nextCol = currentCol + currentDir[1]

						if nextRow < 0 ||
						nextRow >= len(matrix) ||
						nextCol < 0 ||
						nextCol >= len(matrix[0]) {
							continue
						}

						if matrix[nextRow][nextCol] == 1 {
							queue.Push([]int{nextRow, nextCol})
							matrix[nextRow][nextCol] = 0
						}
					}

				}
			}
		}
	}
	return islands
}

func orangesRotting(matrix matrix) int {
	if len(matrix) == 0 { return 0 }

	var directions = [][]int{{0,1}, {1,0}, {0,-1}, {-1,0},}
	var ROTTEN, FRESH = 2, 1

	var q = Queue{}
	var freshOranges = 0

	for row := 0 ; row < len(matrix) ; row++ {
		for col := 0 ; col < len(matrix[0]) ; col++ {
			if matrix[row][col] == ROTTEN {
				q.Push([]int{row, col})
			}
			if matrix[row][col] == FRESH {
				freshOranges += 1
			}
		}
	}

	var minutes = 0
	var queueSize = q.Size

	for q.Length() > 0 {
		if queueSize == 0 {
			minutes += 1
			queueSize = q.Size
		}

		var row, col = q.Pop()
		queueSize -= 1

		for i := 0 ; i < len(directions) ; i++ {
			var currentDir = directions[i]
			var nextRow = currentDir[0] + row
			var nextCol = currentDir[1] + col

			if nextRow < 0 || nextRow >= len(matrix) || nextCol < 0 || nextCol >= len(matrix[0]) {
				continue
			}

			if matrix[nextRow][nextCol] == FRESH {
				matrix[nextRow][nextCol] = 2
				freshOranges -= 1
				q.Push([]int{nextRow, nextCol})
			}
		}
	}

	if freshOranges > 0 {
		return -1
	}

	return minutes
}

func depthFirstSearch(matrix matrix, row, col, step int) {
	var direction = [][]int{{-1,0}, {0,1}, {1,0}, {0,-1}}
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || step > matrix[row][col] {
		return
	}
	matrix[row][col] = step
	for i := 0 ; i < len(directions) ; i++ {
		depthFirstSearch(matrix, row + directions[i][0], col + directions[i][1], step + 1)
	}
}

func wallsAndGates(matrix matrix) int {
	var WALL, GATE = -1, 0
	for row := 0 ; row < len(matrix) ; row++ {
		for col := 0 ; col < len(matrix[0]) ; col++ {
			if matrix[row][col] == GATE {
				depthFirstSearch(matrix, row, col, 0)
			}
		}
	}
	return matrix
}


func main() {

	ourMap := [][]int{
		{1, 2, 3, 4, 5}, 
		{6, 7, 8, 9, 10}, 
		{11, 12, 13, 14, 15}, 
		{16, 17, 18, 19, 20},
	}

	fmt.Println(traverseDFS(ourMap))
	fmt.Println(traverseBFS(ourMap))
	fmt.Println("===================================================")

	m := [][]int{
		{1, 1, 1, 1, 0}, 
		{1, 1, 0, 0, 1}, 
		{0, 0, 0, 1, 1}, 
		{0, 1, 0, 0, 1},
	}

	m2 := [][]int{
		{0, 1, 0, 1, 0}, 
		{1, 0, 1, 0, 1}, 
		{0, 1, 1, 1, 0}, 
		{1, 0, 1, 0, 1},
	}

	fmt.Println(countIslands(m))
	fmt.Println(countIslands(m2))
	fmt.Println("===================================================")

	// oranges
	m3 := [][]int{
		{2, 1, 1, 0, 0}, 
		{1, 1, 1, 0, 1}, 
		{0, 1, 1, 1, 1}, 
		{0, 1, 0, 0, 1},
	}
	fmt.Println("rotting oranges", m3)
	fmt.Println("rotten oranges conclusion:", orangesRotting(m3), "minutes")
	fmt.Println("===================================================")
	fmt.Println("rotting oranges", m2)
	fmt.Println("rotten oranges conclusion:", orangesRotting(m2), "minutes")

}