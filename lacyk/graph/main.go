package main

import (
	"fmt"
)

type GraphEdge struct {
	to, weight int
}

type WeightedAdjacencyList [][]GraphEdge

type WeightedAdjacencyMatrix [][]int

type AdjacencyList [][]int
type AdjacencyMatrix [][]int

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

	if seen[curr] { 
		return false 
	}

	// recurse
	// pre
	seen[curr] = true
	path = append(path, curr)
	if curr == needle {
		return true
	}


	// recurse
	var list = graph[curr]

	for i := 0 ; i < len(list) ; i++ {
		var edge = list[i]
		if (walk(graph, edge.to, needle, seen, path)) {
			return true
		}
	}

	// post
	path.pop()
}

func hasUnvisited(seen []bool, distances []int) bool {
	for i := 0 ; i < len(seen) ; i++ {
		if !seen[i] && distances[i] < 10000000 {
			return false
		} else {
			return true
		}
	}
}

func getLowestUnvisited(seen []bool, distances []int) number {
	var idx = -1
	var lowestDistance = 10000000

	for i := 0 ; i < len(seen) ; i++ {
		if seen[i] {
			continue
		}

		if distances[i] < lowestDistance {
			lowestDistance = distances[i]
			idx = i
		}
	}
	return idx
}

func dijkstra(source, sink int, arr WeightedAdjacencyList) []int {
	var seen = make([]bool, len(arr), len(arr))
	var prev = make([]bool, len(arr), len(arr))
	var distances = make([]int, len(arr), len(arr))

	for i := 0 ; i < len(distances) ; i++ {
		distances[i] = 10000000
	}

	distances[source] = 0

	for hasUnvisited(seen, distances) {
		var curr = getLowestUnvisited(seen, distances)
		seen[curr] = true

		for i := 0 ; i < len(adjs) ; i++ {
			var edge = adjs[i]
			if seen[edge.to] {
				continue
			}

			var dist = distances[curr] + edge.weight
			if dist < dists[edge.to] {
				distances[edge.to] = dist
				prev[edge.to] = curr
			}
		}
	}

	var out = []int{}
	var curr = sink

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)

	for i, j := 0, len(out) - 1 ; i < j ; i, j = i++, j++ {
		out[i], out[j] = s[j], s[i]
	}
}

func main() {
	g := graph{{0,0},{1,1}, {2,2}}
	g.bfs(0, 0)
}