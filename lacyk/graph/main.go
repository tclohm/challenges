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

		var curr = needle
		var out = []int{}

		for prev[curr] != -1 {
			out = append(out, curr)
			curr = prev[curr]
		}

		out = append(out, source)

		if len(out) {
			return out.reverse()
		}
	}

	return []int{}

}

func main() {
	g := graph{{0,0},{1,1}, {2,2}}
	g.bfs(0, 0)
}