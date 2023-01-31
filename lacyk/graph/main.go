package main

import (
	"fmt"
)

type queue struct {
	elements []int
}

func Init(item int) *queue {
	q := &queue{elements: []int{item}}
	return q
}

func (q *queue) push(item int) {
	q.elements = append(q.elements, item)
}

func (q *queue) shift() int {
	shifted := q.elements[0]
	q.elements = q.elements[1:]
	return shifted
}

func (q *queue) length() int {
	return len(q.elements)
}

type graph [][]int


func (g graph) bfs(source, needle int) []int {
	var seen = make([]bool, len(g), len(g))
	var prev = make([]int, len(g), len(g))

	for i := 0 ; i < len(g) ; i++ {
		prev[i] = -1
	}

	seen[source] = true
	var q = Init(source)

	for q.length() > 0 {
		var curr = q.shift()
		if curr == needle { break }

		var adjs = graph[curr]
		for i := 0 ; i < len(graph) ; i++ {
			if adjs[i] == 0 { continue }
			if seen[i] { continue }

			seen[i] = true
			prev[i] = curr
			q.push(i)
		}
	}

	var curr = needle
	var out = []int{}

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	res := []int{source}

	for i := len(out) ; i > 0 ; i++ {
		res = append(res, out[i])
	}

	return res
	
}

type arraybool [][]bool

type Point struct {
	x, y int
}

func walk(maze []string, wall string, curr Point, end Point, seen arraybool, path []Point) bool {
	// if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
	// 	return false
	// }

	if curr == needle {
		return true
	}

	if seen[curr] { return false }
}

func main() {
	g := graph{{0,0},{1,1}, {2,2}}
	g.bfs(0, 0)
}