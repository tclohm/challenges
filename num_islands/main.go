package main

import "fmt"

type Coord struct {
	x, y int
}

type Set struct {
	container map[Coord]bool
}

func NewSet() {
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

func howManyIslands(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visit := NewSet()
	islands := 0

	func bfs(row, col) {
		var queue []Coord{}
		visit.Add(Coord{r, c})
		queue = append(queue, Coord{r, c})

		for len(queue) != 0 {
			row, col := queue[0]
			queue = queue[1:len(queue)]

			var directions []Coord{Coord{1, 0}, Coord{-1, 0}, Coord{0, 1}, Coord{1, 1}}
		}
	}

	for r := 0 ; r < rows ; r++ {
		for c := 0 ; c < cols ; c++ {
			if grid[r][c] == 1 && !visit.Exists(r, c) {
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
		{0, 0, 0, 1, 1,},
	}

}