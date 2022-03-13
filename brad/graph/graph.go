package main

import (
	"fmt"
	"sync"
)

type Vertex struct {
	value interface{}
}

func (v *Vertex) String() string {
	return fmt.Sprintf("%v", v.value)
}

type Graph struct {
	vertices []*Vertex
	edges map[Vertex][]*Vertex
	lock sync.RWMutex
}

func (g *Graph) AddVertex(v *Vertex) {
	g.lock.Lock()
	g.vertices = append(g.vertices, v)
	g.lock.Unlock()
}

func (g *Graph) AddEdge(Vertex1, Vertex2 *Vertex) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Vertex][]*Vertex)
	}
	g.edges[*Vertex1] = append(g.edges[*Vertex1], Vertex2)
	g.edges[*Vertex2] = append(g.edges[*Vertex2], Vertex1)
	g.lock.Unlock()
}

func (g *Graph) String() {
	g.lock.RLock()
	s := ""
	for i := 0 ; i < len(g.vertices); i++ {
		s += g.vertices[i].String() + " -> "
		neighbors := g.edges[*g.vertices[i]]
		for j := 0 ; j < len(neighbors) ; j++ {
			s += neighbors[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}

func main() {
	g := Graph{}
	nA := Vertex{"A"}
    nB := Vertex{"B"}
    nC := Vertex{"C"}
    nD := Vertex{"D"}
    nE := Vertex{"E"}
    nF := Vertex{"F"}


    g.AddVertex(&nA)
    g.AddVertex(&nB)
    g.AddVertex(&nC)
    g.AddVertex(&nD)
    g.AddVertex(&nE)
    g.AddVertex(&nF)

    g.AddEdge(&nA, &nB)
    g.AddEdge(&nA, &nC)
    g.AddEdge(&nB, &nE)
    g.AddEdge(&nC, &nE)
    g.AddEdge(&nE, &nF)
    g.AddEdge(&nD, &nA)

    g.String()
}