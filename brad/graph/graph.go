package main

import "fmt"

type Vertex struct {
	key int
	neighbors map[*Vertex]int
}

func (v *Vertex) AddNeighbor(neighbor Vertex, weight int) {
	if v.neighbors == nil {
		v.neighbors = map[*Vertex]int{}
	}
	v.neighbors[&neighbor] = weight
}

func (v *Vertex) GetConnections() []int {
	keys := []int{}
	for vertex, _ := range v.neighbors {
		keys = append(keys, vertex.key)
	}
	return keys
}

func (v *Vertex) GetWeight(neighbor Vertex) int {
	return v.neighbors[&neighbor]
}

type Graph struct {
	vertices map[int]Vertex
}

func (g *Graph) AddVertex(vertex Vertex) {
	if g.vertices == nil {
		g.vertices = map[int]Vertex{}
	}
	g.vertices[vertex.key] = vertex
}

func (g *Graph) GetVertex(key int) Vertex {
	return g.vertices[key]
}

func (g *Graph) Contains(key int) bool {
	if _, ok := g.vertices[key]; ok {
		return true
	}
	return false
}

func (g *Graph) AddEdge(from, to, weight int) {
	if _, ok := g.vertices[from]; !ok {
		g.AddVertex(Vertex{key: from})
	}
	if _, ok := g.vertices[to]; !ok {
		g.AddVertex(Vertex{key: to})
	}
	fromVertex := g.vertices[from]
	toVertex := g.vertices[to]
	fromVertex.AddNeighbor(toVertex, weight)
}

func main() {
	g := Graph{}
	for i := 0 ; i < 6 ; i++ {
		v := Vertex{key: i}
		g.AddVertex(v)
	}

	g.AddEdge(0, 1, 5)
	g.AddEdge(0, 5, 2)
	g.AddEdge(1, 2, 4)
	g.AddEdge(2, 3, 9)
	g.AddEdge(3, 4, 7)
	g.AddEdge(3, 5, 3)
	g.AddEdge(4, 0, 1)
	g.AddEdge(5, 4, 8)
	g.AddEdge(5, 2, 1)

	fmt.Println(&g)
}