package main

import "fmt"

type tuple struct {
	x, y int
}

func isPathCrossing(path string) bool {
	var directions = map[rune][]int{
		'N': {0,1},
		'S': {0,-1},
		'E': {1,0},
		'W': {-1,0},
	}

	var visited map[tuple]bool = make(map[tuple]bool)
	var x, y = 0, 0

	for _, step := range path {
		visited[tuple{x: x, y: y}] = true
		var dx, dy = directions[step][0], directions[step][1]
		x, y = x + dx, y + dy
		var point tuple = tuple{x: x, y: y}
		if _, crossing := visited[point]; crossing {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isPathCrossing("NES"))
	fmt.Println(isPathCrossing("NESWW"))
}