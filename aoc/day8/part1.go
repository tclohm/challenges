package main

import (
	"fmt"
	"bufio"
	"os"
)

// 0, short - 9, tallest

// left, up, right, down

// how many trees are visible from outside the grid

// we can see all the exterior trees

type point struct {
	x, y int
}

type queue struct {
	trees []point
}

func (q *queue) push(tree point) {
	q.trees = append(q.trees, tree)
}

func (q *queue) shift() (int, int) {
	tree := q.trees[0]
	q.trees = q.trees[1:]
	return tree.x, tree.y
}

func (q *queue) length() int {
	return len(q.trees)
}

func main() {
	// left, up, right, down
	//directions := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0},}

	file, _ := os.Open("file.txt")

	defer file.Close()

	sc := bufio.NewScanner(file)

	forest := [][]rune{}
	visited := [][]bool{}
	q := queue{}

	// build the forest and visited 2d array
	for sc.Scan() {
		row := []rune{}

		for _, tree := range sc.Text() {
			row = append(row, tree)
		}
		forest = append(forest, row)
		visited = append(visited, make([]bool, len(sc.Text()), len(sc.Text())))
	}

	// first traverse
	row := 0
	col := 0

	for {
		// right
		for col < len(forest[0]) - 1 {
			q.push(point{col, row})
			visited[row][col] = true
			col += 1
			
		}

		// down
		for col == len(forest[0]) - 1 && row <= len(forest) - 1 {
			q.push(point{col, row})
			visited[row][col] = true
			row += 1
			
		} 

		// left
		row = len(forest) - 1
		col = len(forest[0]) - 2
		
		for col >= 0 && row <= len(forest) {
			q.push(point{col, row})
			visited[row][col] = true
			col -= 1
		}

		// up
		
		col = 0
		row = len(forest) - 2
		
		for col == 0 && row <= len(forest) - 1 && row > 0 {
			q.push(point{col, row})
			visited[row][col] = true
			row -= 1
		}

		break
	}

	// for q.Length() > 0 {
	// 	col, row := q.pop()
	// }
}