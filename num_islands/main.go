package main

import "fmt"

type Coord struct {
	x, y int
}

type Set struct {
	container map[Coord]bool
}

func NewSet() *Set {
	s := &Set{
		container: make(map[Coord]bool),
	}
	return s
}

func (s *Set) Exists(key Coord) bool {
	_, exists := s.container[key]
	return exists
}

func (s *Set) Add(key Coord) {
	s.container[key] = true
}

func checkBounds(r, c int, dir Coord, rows, cols int, grid [][]int) bool {
	dr, dc := dir.y, dir.x

	if 0 < r + dr && r + dr < rows &&
	0 < c + dc && c + dc < cols &&
	grid[r + dr][r + dc] == 1 {
		return true
	} else {
		return false
	}
}

func howManyIslands(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visit := NewSet()
	islands := 0

	var bfs func(int, int)
	bfs = func (row, col int) {
		var queue = []Coord{}
		visit.Add(Coord{row, col})
		queue = append(queue, Coord{row, col})

		for len(queue) != 0 {
			coord := queue[0]
			queue = queue[1:len(queue)]
			r, c := coord.y, coord.x

			var directions = []Coord{Coord{1, 0}, Coord{-1, 0}, Coord{0, 1}, Coord{0, -1}}

			for dir := 0 ; dir < len(directions) ; dir++ {
				dr, dc := directions[dir].y, directions[dir].x
				p_coord := Coord{r + dr, c + dc}
				if checkBounds(r, c, directions[dir], rows, cols, grid) &&
				!visit.Exists(p_coord) {
					queue = append(queue, p_coord)
					visit.Add(p_coord)
				}
				
			}
		}
	}

	for r := 0 ; r < rows ; r++ {
		for c := 0 ; c < cols ; c++ {
			coord := Coord{c, r}
			if grid[r][c] == 1 && !visit.Exists(coord) {
				bfs(r, c)
				islands += 1
			}
		}
	}
	return islands
}

func main() {
	var worldMap = [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0,},
		{1, 1, 0, 0, 0,},
		{0, 0, 0, 0, 0,},
	}
	fmt.Println(howManyIslands(worldMap))
}