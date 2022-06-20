package main

import "fmt"

type Coord struct {
	y, x int
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

type Queue struct {
	container []Coord
}

func NewQueue() *Queue {
	q := &Queue{
		container: []Coord{},
	}
	return q
}

func (q *Queue) Push(item Coord) {
	q.container = append(q.container, item)
}

func (q *Queue) PopLeft() Coord {
	item := q.container[0]
	q.container = q.container[1:len(q.container)]
	return item
}

func (q *Queue) Length() int {
	return len(q.container)
}

func checkBounds(r, c int, dir []int, rows, cols int, grid [][]int) bool {
	dr, dc := dir[0], dir[1]

	add_r := r + dr
	add_c := c + dc
	
	if 0 <= add_r && add_r < rows &&
	0 <= add_c && add_c < cols &&
	grid[add_r][add_c] == 1 {
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

	var bfs func(int, int, Set)
	bfs = func (row, col int, visit Set) {
		var queue = NewQueue()
		visit.Add(Coord{x: col, y: row})
		queue.Push(Coord{x: col, y:row})

		for len(queue.container) > 0 {
			coord := queue.PopLeft()
			r, c := coord.y, coord.x

			var directions = [][]int{{1, 0},{-1, 0},{0, 1},{0, -1}}

			for index := 0 ; index < len(directions) ; index++ {
				dr, dc := directions[index][0], directions[index][1]
				p_coord := Coord{x: c + dc, y: r + dr}
				if checkBounds(r, c, directions[index], rows, cols, grid) &&
				!visit.Exists(p_coord) {
					queue.Push(p_coord)
					visit.Add(p_coord)
				}
			}
		}
	}

	for r := 0 ; r < rows ; r++ {
		for c := 0 ; c < cols ; c++ {
			coord := Coord{r, c}
			if grid[r][c] == 1 && !visit.Exists(coord) {
				fmt.Println("\nbfs occurs")
				bfs(r, c, *visit)
				islands += 1
			}
		}
	}

	return islands
}

func main() {
	var worldMap = [][]int{
		{1, 1, 1, 1, 0,},
		{1, 1, 0, 1, 0,},
		{1, 1, 0, 0, 0,},
		{0, 0, 0, 0, 0,},
	}
	// var worldMap = [][]int{
	// 	{1,0,1},
	// 	{1,1,1},
	// }
	fmt.Println(howManyIslands(worldMap))
}