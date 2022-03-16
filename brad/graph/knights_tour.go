package main

import "fmt"

type Set struct {
	row int
	col int
}

type Graph map[Set][]Set


func add_edge(g Graph, node1, node2 Set) {
	if _, ok := g[node1]; !ok {
		g[node1] = []Set{}
	}

	if _, ok := g[node2]; !ok {
		g[node2] = []Set{}
	}
	g[node1] = append(g[node1], node2)
	g[node2] = append(g[node2], node1)
}

func main() {
	g := Graph{}
	fmt.Println(g)
	n1 := Set{1, 2}
	n2 := Set{3, 4}
	add_edge(g, n1, n2)
	fmt.Println(g)
}