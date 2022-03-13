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
	vA := Vertex{"A"}
    vB := Vertex{"B"}
    vC := Vertex{"C"}
    vD := Vertex{"D"}
    vE := Vertex{"E"}
    vF := Vertex{"F"}


    g.AddVertex(&vA)
    g.AddVertex(&vB)
    g.AddVertex(&vC)
    g.AddVertex(&vD)
    g.AddVertex(&vE)
    g.AddVertex(&vF)

    g.AddEdge(&vA, &vB)
    g.AddEdge(&vA, &vC)
    g.AddEdge(&vB, &vE)
    g.AddEdge(&vC, &vE)
    g.AddEdge(&vE, &vF)
    g.AddEdge(&vD, &vA)

    g.String()
}